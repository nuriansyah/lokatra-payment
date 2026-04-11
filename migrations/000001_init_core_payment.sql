-- =============================================================================
-- LOKATRA-PAYMENT  |  Database Schema  v1.0.0
-- =============================================================================
-- Compliance coverage:
--   PCI-DSS v4   — no raw PAN stored; tokenisation only; audit trail on all
--                  sensitive table access; encrypted vault references
--   ISO 27001    — full immutable audit log; data classification tags;
--                  access-control hooks via meta_created_by / reviewed_by
--   SOC 2 Type II— change-data capture columns (meta_updated_at/by);
--                  availability tracked via circuit-breaker state table
--   GDPR / UU PDP— personal data minimised to what is strictly necessary;
--                  gdpr_erasure_requested_at on payment_methods + purchaser data;
--                  data-retention policy table drives scheduled purge jobs
--   RFC 7807     — error detail stored as structured JSON in payment_events
-- =============================================================================
-- Relation to sibling services:
--   lokatra-auth     → supplies user_id / merchant_id (UUID, NOT resolved here)
--   lokatra-core-biz → supplies order_id / event_id  (UUID, NOT resolved here)
--   All cross-service FKs are stored as plain UUID columns (no FK constraint)
--   to maintain service autonomy and allow independent deployment.
-- =============================================================================

-- ─────────────────────────────────────────────────────────────────────────────
-- 0.  SHARED TYPES
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TYPE payment_currency AS ENUM (
    'IDR','USD','SGD','MYR','PHP','THB','AED','EUR','GBP','JPY'
);

CREATE TYPE payment_status_enum AS ENUM (
    'INITIATED',       -- intent created
    'PENDING',         -- sent to PSP, awaiting async confirm
    'AUTHORISED',      -- authorised but not captured (card pre-auth)
    'CAPTURED',        -- funds captured
    'PARTIALLY_CAPTURED',
    'COMPLETED',       -- settled end-to-end
    'FAILED',          -- PSP hard failure
    'CANCELLED',       -- cancelled before capture
    'EXPIRED',         -- payment window lapsed
    'REFUNDING',       -- refund in-flight
    'REFUNDED',        -- full refund completed
    'PARTIALLY_REFUNDED',
    'DISPUTED',        -- chargeback / dispute opened
    'CHARGEBACK_WON',
    'CHARGEBACK_LOST'
);

CREATE TYPE payment_method_type_enum AS ENUM (
    'CARD',            -- credit / debit
    'VIRTUAL_ACCOUNT', -- bank VA (Mandiri, BCA, BNI, BRI, Permata …)
    'QRIS',            -- QR Indonesian Standard
    'EWALLET',         -- GoPay, OVO, Dana, ShopeePay, LinkAja
    'DIRECT_DEBIT',    -- bank direct debit
    'BANK_TRANSFER',   -- manual transfer
    'PAYLATER',        -- Kredivo, Akulaku, Spaylater
    'CRYPTO',          -- future-proofing
    'VOUCHER',         -- gift card / voucher code
    'POINTS',          -- loyalty points redemption
    'CASH_ON_DELIVERY'
);

CREATE TYPE psp_enum AS ENUM (
    'MIDTRANS',
    'XENDIT',
    'STRIPE',
    'DOKU',
    'DANA',
    'OVO',
    'GOPAY',
    'SHOPEE_PAY',
    'LINK_AJA',
    'FLIP',
    'INTERNAL'         -- lokatra internal balance / voucher
);

CREATE TYPE refund_status_enum AS ENUM (
    'PENDING','PROCESSING','SUCCEEDED','FAILED','CANCELLED'
);

CREATE TYPE dispute_status_enum AS ENUM (
    'OPEN','EVIDENCE_SUBMITTED','WON','LOST','CLOSED_BY_ISSUER'
);

CREATE TYPE payout_status_enum AS ENUM (
    'SCHEDULED','PROCESSING','COMPLETED','FAILED','ON_HOLD','REVERSED'
);

CREATE TYPE routing_strategy_enum AS ENUM (
    'LOWEST_COST',         -- pick PSP with cheapest MDR
    'HIGHEST_SUCCESS_RATE',
    'ROUND_ROBIN',
    'GEO_PREFERRED',       -- PSP closest to user region
    'MANUAL',              -- hard-pinned via routing rule
    'WATERFALL'            -- try in priority order until success
);

CREATE TYPE settlement_status_enum AS ENUM (
    'PENDING','PROCESSING','COMPLETED','FAILED','DISPUTED'
);

CREATE TYPE webhook_status_enum AS ENUM (
    'RECEIVED','PROCESSING','PROCESSED','FAILED','IGNORED'
);


-- ─────────────────────────────────────────────────────────────────────────────
-- 1.  IDEMPOTENCY  (RFC 7231 / Stripe convention)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS idempotency_keys (
    id                  UUID            PRIMARY KEY,
    idempotency_key     VARCHAR(128)    NOT NULL UNIQUE,
    -- which merchant / caller owns this key
    merchant_id         UUID            NOT NULL,   -- ref: lokatra-auth merchants
    request_path        VARCHAR(512)    NOT NULL,
    request_body_hash   VARCHAR(64)     NOT NULL,   -- SHA-256 of canonical body
    response_status     SMALLINT,
    response_body       JSONB,
    locked_at           TIMESTAMPTZ,               -- in-flight lock
    locked_until        TIMESTAMPTZ,
    completed_at        TIMESTAMPTZ,
    expires_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW() + INTERVAL '24 hours',
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE INDEX IF NOT EXISTS idx_idem_merchant ON idempotency_keys (merchant_id, expires_at);


-- ─────────────────────────────────────────────────────────────────────────────
-- 3.  INTELLIGENT ROUTING RULES
-- ─────────────────────────────────────────────────────────────────────────────

-- A routing profile binds to a merchant (or is global when merchant_id IS NULL)
CREATE TABLE IF NOT EXISTS routing_profiles (
    id                  UUID            PRIMARY KEY,
    merchant_id         UUID,           -- NULL = platform default
    name                VARCHAR(128)    NOT NULL,
    strategy            routing_strategy_enum NOT NULL DEFAULT 'HIGHEST_SUCCESS_RATE',
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    fallback_profile_id UUID            REFERENCES routing_profiles(id),
    notes               TEXT,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

-- Individual routing rules within a profile (evaluated in priority ASC)
CREATE TABLE IF NOT EXISTS routing_rules (
    id                  UUID            PRIMARY KEY,
    profile_id          UUID            NOT NULL REFERENCES routing_profiles(id) ON DELETE CASCADE,
    priority            SMALLINT        NOT NULL DEFAULT 10,
    name                VARCHAR(128)    NOT NULL,
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    -- Conditions (all must match; NULL = wildcard)
    match_payment_method payment_method_type_enum,
    match_currency      payment_currency,
    match_amount_min    DECIMAL(18,2),
    match_amount_max    DECIMAL(18,2),
    match_user_country  CHAR(2),        -- ISO 3166-1 alpha-2
    match_card_bin      VARCHAR(8),     -- first 6-8 digits for BIN routing
    match_product_type  VARCHAR(64),    -- e.g. 'TICKET','ACCOMMODATION','EXPERIENCE'
    -- Action
    cost_weight         DECIMAL(5,4),   -- for LOWEST_COST scoring
    notes               TEXT,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE INDEX IF NOT EXISTS idx_routing_rules_profile ON routing_rules (profile_id, priority, is_active);

-- Immutable record of every routing decision made
CREATE TABLE IF NOT EXISTS routing_decisions (
    id                  UUID            PRIMARY KEY,
    payment_intent_id   UUID            NOT NULL,  -- FK set after intent table exists
    profile_id          UUID            REFERENCES routing_profiles(id),
    rule_id             UUID            REFERENCES routing_rules(id),
    strategy_used       routing_strategy_enum,
    candidate_psps      JSONB           NOT NULL DEFAULT '[]',  -- [{psp_account_id, score, reason}]
    decision_reason     TEXT,
    decided_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE INDEX IF NOT EXISTS idx_routing_decisions_intent ON routing_decisions (payment_intent_id);


-- ─────────────────────────────────────────────────────────────────────────────
-- 4.  PAYMENT METHODS  (PCI-DSS token vault reference)
-- ─────────────────────────────────────────────────────────────────────────────

-- GDPR / PCI-DSS:
--   Raw card numbers (PAN) are NEVER stored here.
--   `token_ref` is a vault or PSP network token reference.
CREATE TABLE IF NOT EXISTS payment_methods (
    id                      UUID            PRIMARY KEY,
    -- Cross-service references (no FK — service boundary)
    user_id                 UUID,           -- ref: lokatra-auth users (nullable for guest)
    merchant_id             UUID,           -- ref: lokatra-auth merchants
    -- Classification
    method_type             payment_method_type_enum NOT NULL,
    psp                     psp_enum        NOT NULL,
    -- Tokenised instrument reference (PCI-DSS compliant)
    token_ref               VARCHAR(512),   -- PSP network token or vault path
    token_type              VARCHAR(32),    -- e.g. 'NETWORK_TOKEN','VAULT_TOKEN','EWALLET_REF'
    token_expires_at        TIMESTAMPTZ,
    -- Card display info (non-sensitive, safe to store)
    card_brand              VARCHAR(32),    -- VISA, MASTERCARD, AMEX, JCB …
    card_last_four          VARCHAR(4),
    card_exp_month          SMALLINT        CHECK (card_exp_month BETWEEN 1 AND 12),
    card_exp_year           SMALLINT,
    card_country            CHAR(2),
    card_funding_type       VARCHAR(16),    -- CREDIT, DEBIT, PREPAID
    card_bin                VARCHAR(8),     -- for routing; NOT the full PAN
    -- E-wallet / VA details
    wallet_account_ref      VARCHAR(255),   -- hashed / truncated reference
    va_bank_code            VARCHAR(32),
    -- General
    display_label           VARCHAR(128),   -- "Visa …4242"
    is_default              BOOLEAN         NOT NULL DEFAULT FALSE,
    is_active               BOOLEAN         NOT NULL DEFAULT TRUE,
    verified_at             TIMESTAMPTZ,
    fingerprint             VARCHAR(128),   -- dedup: same card across tokens
    -- GDPR: right-to-erasure support
    gdpr_erasure_requested_at TIMESTAMPTZ,
    gdpr_erased_at          TIMESTAMPTZ,
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID,
    CONSTRAINT chk_pm_erased CHECK (
        gdpr_erased_at IS NULL OR gdpr_erasure_requested_at IS NOT NULL
    )
);

CREATE INDEX IF NOT EXISTS idx_pm_user       ON payment_methods (user_id, is_active);
CREATE INDEX IF NOT EXISTS idx_pm_merchant   ON payment_methods (merchant_id);
CREATE INDEX IF NOT EXISTS idx_pm_fingerprint ON payment_methods (fingerprint) WHERE fingerprint IS NOT NULL;


-- ─────────────────────────────────────────────────────────────────────────────
-- 5.  PAYMENT INTENTS
--     Models the *desire* to pay before a charge attempt is made.
--     Inspired by Stripe PaymentIntent / Adyen PaymentSession.
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payment_intents (
    id                      UUID            PRIMARY KEY,
    intent_code             VARCHAR(32)     NOT NULL UNIQUE,   -- PAY-20240412-XXXXXX
    -- Merchant / product context
    merchant_id             UUID            NOT NULL,  -- ref: lokatra-auth
    -- Cross-service reference: the order / booking this intent is for
    order_id                UUID,           -- ref: lokatra-core-biz orders
    order_type              VARCHAR(64),    -- 'EVENT_TICKET','ACCOMMODATION','EXPERIENCE' …
    -- Amounts
    amount                  DECIMAL(18,2)   NOT NULL CHECK (amount > 0),
    currency                payment_currency NOT NULL,
    tax_amount              DECIMAL(18,2)   NOT NULL DEFAULT 0,
    discount_amount         DECIMAL(18,2)   NOT NULL DEFAULT 0,
    tip_amount              DECIMAL(18,2)   NOT NULL DEFAULT 0,
    -- Payer info (GDPR: minimal, pseudonymised where possible)
    user_id                 UUID,           -- ref: lokatra-auth (NULL = guest checkout)
    customer_name           VARCHAR(255),
    customer_email          VARCHAR(255),
    customer_phone          VARCHAR(50),
    customer_ip             INET,
    customer_country        CHAR(2),
    -- Payment method selection
    payment_method_id       UUID            REFERENCES payment_methods(id),
    payment_method_type     payment_method_type_enum,
    -- State machine
    status                  payment_status_enum NOT NULL DEFAULT 'INITIATED',
    -- Routing
    routing_profile_id      UUID            REFERENCES routing_profiles(id),
    -- Expiry
    expires_at              TIMESTAMPTZ     NOT NULL DEFAULT NOW() + INTERVAL '30 minutes',
    -- 3DS / SCA
    requires_3ds            BOOLEAN         NOT NULL DEFAULT FALSE,
    three_ds_version        VARCHAR(8),
    -- Capture mode
    capture_mode            VARCHAR(16)     NOT NULL DEFAULT 'AUTOMATIC'
                                CHECK (capture_mode IN ('AUTOMATIC','MANUAL')),
    -- Metadata / description
    description             TEXT,
    statement_descriptor    VARCHAR(22),    -- what appears on bank statement
    metadata                JSONB           NOT NULL DEFAULT '{}',
    -- Promo / voucher
    promo_code              VARCHAR(64),
    promo_discount_amount   DECIMAL(18,2)   NOT NULL DEFAULT 0,
    -- Idempotency
    idempotency_key_id      UUID            REFERENCES idempotency_keys(id),
    -- Timestamps
    confirmed_at            TIMESTAMPTZ,
    cancelled_at            TIMESTAMPTZ,
    cancellation_reason     TEXT,
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID,
    CONSTRAINT chk_intent_amount CHECK (
        amount >= (tax_amount + discount_amount)
    )
);

CREATE INDEX IF NOT EXISTS idx_intent_merchant ON payment_intents (merchant_id, status, meta_created_at DESC);
CREATE INDEX IF NOT EXISTS idx_intent_order    ON payment_intents (order_id) WHERE order_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_intent_user     ON payment_intents (user_id) WHERE user_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_intent_status   ON payment_intents (status, expires_at);

-- Line items inside the intent (supports split-product carts)
CREATE TABLE IF NOT EXISTS payment_intent_items (
    id                  UUID            PRIMARY KEY,
    intent_id           UUID            NOT NULL REFERENCES payment_intents(id) ON DELETE CASCADE,
    product_id          UUID            NOT NULL,   -- ref: core-biz product/ticket
    product_type        VARCHAR(64)     NOT NULL,
    product_name        VARCHAR(255)    NOT NULL,
    quantity            INT             NOT NULL DEFAULT 1 CHECK (quantity > 0),
    unit_price          DECIMAL(18,2)   NOT NULL,
    discount_amount     DECIMAL(18,2)   NOT NULL DEFAULT 0,
    total_price         DECIMAL(18,2)   NOT NULL,
    seller_merchant_id  UUID,           -- for marketplace split
    metadata            JSONB           NOT NULL DEFAULT '{}',
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE INDEX IF NOT EXISTS idx_pii_intent ON payment_intent_items (intent_id);


-- ─────────────────────────────────────────────────────────────────────────────
-- 6.  PAYMENTS  (actual charge attempts against a PSP)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payments (
    id                      UUID            PRIMARY KEY,
    payment_code            VARCHAR(32)     NOT NULL UNIQUE,  -- TXN-20240412-XXXXXX
    intent_id               UUID            NOT NULL REFERENCES payment_intents(id),
    -- Attempt tracking
    attempt_number          SMALLINT        NOT NULL DEFAULT 1,
    -- PSP linkage
    psp                     psp_enum        NOT NULL,
    psp_transaction_id      VARCHAR(255),   -- PSP's own ID
    psp_reference           VARCHAR(255),   -- PSP order/session reference
    psp_raw_request         JSONB,          -- redacted PSP request payload
    psp_raw_response        JSONB,          -- redacted PSP response
    -- Amounts (may differ from intent if FX applied)
    amount                  DECIMAL(18,2)   NOT NULL,
    currency                payment_currency NOT NULL,
    amount_in_settlement_currency DECIMAL(18,2),
    settlement_currency     payment_currency,
    fx_rate                 DECIMAL(18,8),
    fx_rate_snapshot_id     UUID,           -- ref: fx_rate_snapshots
    -- Method used
    payment_method_id       UUID            REFERENCES payment_methods(id),
    payment_method_type     payment_method_type_enum NOT NULL,
    -- Status
    status                  payment_status_enum NOT NULL DEFAULT 'PENDING',
    failure_code            VARCHAR(64),    -- PSP error code
    failure_message         TEXT,
    failure_category        VARCHAR(32),    -- INSUFFICIENT_FUNDS, CARD_DECLINED …
    -- Capture
    authorised_at           TIMESTAMPTZ,
    authorised_amount       DECIMAL(18,2),
    captured_at             TIMESTAMPTZ,
    captured_amount         DECIMAL(18,2),
    -- Fees
    processing_fee          DECIMAL(18,2),
    processing_fee_currency payment_currency,
    -- Fraud / risk
    risk_score_id           UUID,           -- ref: risk_assessments
    -- Metadata
    description             TEXT,
    metadata                JSONB           NOT NULL DEFAULT '{}',
    -- Timestamps
    completed_at            TIMESTAMPTZ,
    cancelled_at            TIMESTAMPTZ,
    expired_at              TIMESTAMPTZ,
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID
);

CREATE INDEX IF NOT EXISTS idx_payments_intent      ON payments (intent_id);
CREATE INDEX IF NOT EXISTS idx_payments_psp         ON payments (psp_transaction_id);
CREATE INDEX IF NOT EXISTS idx_payments_status      ON payments (status, meta_created_at DESC);
CREATE INDEX IF NOT EXISTS idx_payments_code        ON payments (payment_code);

-- Payment state transitions — append-only FSM log
CREATE TABLE IF NOT EXISTS payment_state_transitions (
    id                  UUID            PRIMARY KEY,
    payment_id          UUID            NOT NULL REFERENCES payments(id),
    from_status         payment_status_enum,
    to_status           payment_status_enum NOT NULL,
    triggered_by        VARCHAR(64)     NOT NULL,  -- 'PSP_WEBHOOK','API_CALL','SCHEDULER'
    actor_id            UUID,
    psp_event_id        VARCHAR(255),
    reason              TEXT,
    metadata            JSONB           NOT NULL DEFAULT '{}',
    transitioned_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL
    -- immutable — no updated_at
);

CREATE INDEX IF NOT EXISTS idx_pst_payment ON payment_state_transitions (payment_id, transitioned_at DESC);


-- ─────────────────────────────────────────────────────────────────────────────
-- 7.  PSP WEBHOOKS  (inbound events from PSPs)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS psp_webhooks (
    id                  UUID            PRIMARY KEY,
    psp                 psp_enum        NOT NULL,
    psp_event_id        VARCHAR(255),   -- PSP's idempotency key for event
    psp_event_type      VARCHAR(128)    NOT NULL,  -- e.g. 'payment.success'
    received_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    headers             JSONB           NOT NULL DEFAULT '{}',
    raw_payload         JSONB           NOT NULL,
    hmac_valid          BOOLEAN,        -- result of HMAC verification
    status              webhook_status_enum NOT NULL DEFAULT 'RECEIVED',
    processing_attempts INT             NOT NULL DEFAULT 0,
    last_error          TEXT,
    processed_at        TIMESTAMPTZ,
    -- resolved linkage (populated after processing)
    resolved_payment_id UUID            REFERENCES payments(id),
    resolved_intent_id  UUID            REFERENCES payment_intents(id),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID,
    UNIQUE (psp_account_id, psp_event_id)
);

CREATE INDEX IF NOT EXISTS idx_webhook_status ON psp_webhooks (status, received_at);


-- ─────────────────────────────────────────────────────────────────────────────
-- 8.  VIRTUAL ACCOUNT & QRIS  (async payment instruments)
-- ─────────────────────────────────────────────────────────────────────────────

-- VA numbers expire; each payment intent gets at most one VA assignment
CREATE TABLE IF NOT EXISTS virtual_account_assignments (
    id                  UUID            PRIMARY KEY,
    intent_id           UUID            NOT NULL REFERENCES payment_intents(id),
    bank_code           VARCHAR(32)     NOT NULL,  -- BCA, BNI, MANDIRI, BRI, PERMATA …
    va_number           VARCHAR(32)     NOT NULL,
    va_number_masked    VARCHAR(32),
    expires_at          TIMESTAMPTZ     NOT NULL,
    is_reusable         BOOLEAN         NOT NULL DEFAULT FALSE,
    paid_at             TIMESTAMPTZ,
    psp_transaction_id  VARCHAR(255),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID,
    UNIQUE (intent_id, bank_code)
);

CREATE TABLE IF NOT EXISTS qris_assignments (
    id                  UUID            PRIMARY KEY,
    intent_id           UUID            NOT NULL REFERENCES payment_intents(id),
    qr_string           TEXT            NOT NULL,  -- QRIS string (not PAN-sensitive)
    qr_url              TEXT,                      -- hosted image URL
    expires_at          TIMESTAMPTZ     NOT NULL,
    paid_at             TIMESTAMPTZ,
    psp_transaction_id  VARCHAR(255),
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID,
    UNIQUE (intent_id)
);


-- ─────────────────────────────────────────────────────────────────────────────
-- 9.  REFUNDS
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS refunds (
    id                      UUID            PRIMARY KEY,
    refund_code             VARCHAR(32)     NOT NULL UNIQUE,  -- RFD-20240412-XXXXXX
    payment_id              UUID            NOT NULL REFERENCES payments(id),
    intent_id               UUID            NOT NULL REFERENCES payment_intents(id),
    amount                  DECIMAL(18,2)   NOT NULL CHECK (amount > 0),
    currency                payment_currency NOT NULL,
    reason                  VARCHAR(64)     NOT NULL,  -- CUSTOMER_REQUEST, DUPLICATE, FRAUD …
    reason_detail           TEXT,
    status                  refund_status_enum NOT NULL DEFAULT 'PENDING',
    -- PSP
    psp_refund_id           VARCHAR(255),
    psp_raw_response        JSONB,
    -- Reviewer
    requested_by            UUID            NOT NULL,   -- user_id from lokatra-auth
    reviewed_by             UUID,
    reviewed_at             TIMESTAMPTZ,
    review_notes            TEXT,
    -- Disbursement
    refunded_at             TIMESTAMPTZ,
    estimated_arrival       TIMESTAMPTZ,
    failure_reason          TEXT,
    -- Idempotency
    idempotency_key_id      UUID            REFERENCES idempotency_keys(id),
    -- Metadata
    metadata                JSONB           NOT NULL DEFAULT '{}',
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID
);

CREATE INDEX IF NOT EXISTS idx_refunds_payment ON refunds (payment_id, status);
CREATE INDEX IF NOT EXISTS idx_refunds_intent  ON refunds (intent_id);


-- ─────────────────────────────────────────────────────────────────────────────
-- 10. DISPUTES / CHARGEBACKS
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS disputes (
    id                      UUID            PRIMARY KEY,
    dispute_code            VARCHAR(32)     NOT NULL UNIQUE,
    payment_id              UUID            NOT NULL REFERENCES payments(id),
    psp_dispute_id          VARCHAR(255)    NOT NULL,
    dispute_type            VARCHAR(32)     NOT NULL,  -- CHARGEBACK, INQUIRY, RETRIEVAL
    reason_code             VARCHAR(64),    -- Visa/MC reason codes
    reason_description      TEXT,
    amount                  DECIMAL(18,2)   NOT NULL,
    currency                payment_currency NOT NULL,
    status                  dispute_status_enum NOT NULL DEFAULT 'OPEN',
    opened_at               TIMESTAMPTZ     NOT NULL,
    respond_by              TIMESTAMPTZ,
    resolved_at             TIMESTAMPTZ,
    outcome                 VARCHAR(32),    -- WON, LOST
    evidence_due_at         TIMESTAMPTZ,
    evidence_submitted_at   TIMESTAMPTZ,
    evidence_files          JSONB           NOT NULL DEFAULT '[]',  -- [{url, type, uploaded_at}]
    notes                   TEXT,
    handled_by              UUID,
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID
);

CREATE INDEX IF NOT EXISTS idx_disputes_payment ON disputes (payment_id);
CREATE INDEX IF NOT EXISTS idx_disputes_status  ON disputes (status, respond_by);


-- ─────────────────────────────────────────────────────────────────────────────
-- 11. FX / MULTI-CURRENCY
-- ─────────────────────────────────────────────────────────────────────────────

-- Rate snapshots (point-in-time; never updated — new row per refresh)
CREATE TABLE IF NOT EXISTS fx_rate_snapshots (
    id                  UUID            PRIMARY KEY,
    base_currency       payment_currency NOT NULL,
    quote_currency      payment_currency NOT NULL,
    rate                DECIMAL(24,10)  NOT NULL,
    provider            VARCHAR(64)     NOT NULL,  -- 'BANK_INDONESIA','ECB','OPEN_EXCHANGE'
    valid_from          TIMESTAMPTZ     NOT NULL,
    valid_until         TIMESTAMPTZ,
    is_indicative       BOOLEAN         NOT NULL DEFAULT FALSE,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    UNIQUE (base_currency, quote_currency, valid_from, provider)
);

CREATE INDEX IF NOT EXISTS idx_fx_lookup ON fx_rate_snapshots (base_currency, quote_currency, valid_from DESC);


-- ─────────────────────────────────────────────────────────────────────────────
-- 12. SETTLEMENT
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS settlement_batches (
    id                  UUID            PRIMARY KEY,
    batch_code          VARCHAR(32)     NOT NULL UNIQUE,
    merchant_id         UUID            NOT NULL,
    period_from         DATE            NOT NULL,
    period_to           DATE            NOT NULL,
    gross_amount        DECIMAL(18,2)   NOT NULL DEFAULT 0,
    fee_amount          DECIMAL(18,2)   NOT NULL DEFAULT 0,
    tax_amount          DECIMAL(18,2)   NOT NULL DEFAULT 0,
    net_amount          DECIMAL(18,2)   NOT NULL DEFAULT 0,
    currency            payment_currency NOT NULL,
    status              settlement_status_enum NOT NULL DEFAULT 'PENDING',
    psp_settlement_ref  VARCHAR(255),
    settled_at          TIMESTAMPTZ,
    reconciled_at       TIMESTAMPTZ,
    reconciled_by       UUID,
    discrepancy_amount  DECIMAL(18,2)   NOT NULL DEFAULT 0,
    discrepancy_notes   TEXT,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE TABLE IF NOT EXISTS settlement_items (
    id                  UUID            PRIMARY KEY,
    batch_id            UUID            NOT NULL REFERENCES settlement_batches(id) ON DELETE CASCADE,
    payment_id          UUID            NOT NULL REFERENCES payments(id),
    gross_amount        DECIMAL(18,2)   NOT NULL,
    fee_amount          DECIMAL(18,2)   NOT NULL DEFAULT 0,
    net_amount          DECIMAL(18,2)   NOT NULL,
    currency            payment_currency NOT NULL,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    UNIQUE (batch_id, payment_id)
);

CREATE INDEX IF NOT EXISTS idx_settlement_items_batch ON settlement_items (batch_id);


-- ─────────────────────────────────────────────────────────────────────────────
-- 13. PAYOUTS (disbursements to merchants / vendors — marketplace model)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payouts (
    id                      UUID            PRIMARY KEY,
    payout_code             VARCHAR(32)     NOT NULL UNIQUE,  -- PYT-20240412-XXXXXX
    merchant_id             UUID            NOT NULL,   -- recipient
    settlement_batch_id     UUID            REFERENCES settlement_batches(id),
    amount                  DECIMAL(18,2)   NOT NULL CHECK (amount > 0),
    currency                payment_currency NOT NULL,
    fee_amount              DECIMAL(18,2)   NOT NULL DEFAULT 0,
    net_amount              DECIMAL(18,2)   NOT NULL,
    status                  payout_status_enum NOT NULL DEFAULT 'SCHEDULED',
    psp_disbursement_id     VARCHAR(255),
    psp_raw_response        JSONB,
    -- Destination bank
    bank_code               VARCHAR(32),
    bank_account_ref        VARCHAR(512),   -- vault reference, NOT raw account number
    bank_account_name       VARCHAR(255),
    -- Schedule
    scheduled_for           TIMESTAMPTZ,
    processed_at            TIMESTAMPTZ,
    completed_at            TIMESTAMPTZ,
    failed_at               TIMESTAMPTZ,
    failure_reason          TEXT,
    -- Approval (for large amounts or first-time recipients)
    requires_approval       BOOLEAN         NOT NULL DEFAULT FALSE,
    approved_by             UUID,
    approved_at             TIMESTAMPTZ,
    rejection_reason        TEXT,
    -- Idempotency
    idempotency_key_id      UUID            REFERENCES idempotency_keys(id),
    metadata                JSONB           NOT NULL DEFAULT '{}',
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    meta_deleted_at         TIMESTAMPTZ,
    meta_deleted_by         UUID
);

CREATE INDEX IF NOT EXISTS idx_payouts_merchant ON payouts (merchant_id, status, scheduled_for);

-- Line items within a payout (which payments contributed)
CREATE TABLE IF NOT EXISTS payout_items (
    id                  UUID            PRIMARY KEY,
    payout_id           UUID            NOT NULL REFERENCES payouts(id) ON DELETE CASCADE,
    payment_id          UUID            NOT NULL REFERENCES payments(id),
    commission_amount   DECIMAL(18,2)   NOT NULL DEFAULT 0,  -- platform commission deducted
    net_amount          DECIMAL(18,2)   NOT NULL,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    UNIQUE (payout_id, payment_id)
);


-- ─────────────────────────────────────────────────────────────────────────────
-- 14. RISK & FRAUD ASSESSMENT  (SOC2 / ISO 27001)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS risk_assessments (
    id                      UUID            PRIMARY KEY,
    intent_id               UUID            NOT NULL REFERENCES payment_intents(id),
    payment_id              UUID            REFERENCES payments(id),
    -- Scores
    overall_score           DECIMAL(5,4)    NOT NULL CHECK (overall_score BETWEEN 0 AND 1),
    velocity_score          DECIMAL(5,4),
    device_score            DECIMAL(5,4),
    geo_score               DECIMAL(5,4),
    behaviour_score         DECIMAL(5,4),
    -- Decision
    decision                VARCHAR(16)     NOT NULL CHECK (decision IN ('ALLOW','REVIEW','BLOCK')),
    decision_reason         TEXT,
    rule_triggers           JSONB           NOT NULL DEFAULT '[]',  -- [{rule_id, name, score}]
    -- Device / session fingerprint (no PII)
    device_fingerprint      VARCHAR(128),
    ip_address              INET,
    ip_country              CHAR(2),
    ip_is_vpn               BOOLEAN,
    ip_is_datacenter        BOOLEAN,
    user_agent_hash         VARCHAR(64),
    -- 3DS result
    three_ds_eci            VARCHAR(4),
    three_ds_authentication_value VARCHAR(255),
    -- External signals
    bin_risk_level          VARCHAR(16),    -- from BIN intelligence
    assessed_at             TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    assessed_by_engine      VARCHAR(64)     NOT NULL DEFAULT 'lokatra-risk-v1',
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    UNIQUE (intent_id)
);

CREATE INDEX IF NOT EXISTS idx_risk_intent   ON risk_assessments (intent_id);
CREATE INDEX IF NOT EXISTS idx_risk_decision ON risk_assessments (decision, assessed_at DESC);


-- ─────────────────────────────────────────────────────────────────────────────
-- 15. PROMOTIONS / VOUCHERS
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS promo_campaigns (
    id                  UUID            PRIMARY KEY,
    merchant_id         UUID            NOT NULL,
    code_prefix         VARCHAR(16),
    name                VARCHAR(128)    NOT NULL,
    discount_type       VARCHAR(16)     NOT NULL CHECK (discount_type IN ('PERCENTAGE','FIXED','FREE_SHIPPING')),
    discount_value      DECIMAL(18,4)   NOT NULL,
    max_discount_amount DECIMAL(18,2),
    min_order_amount    DECIMAL(18,2),
    currency            payment_currency,
    valid_from          TIMESTAMPTZ     NOT NULL,
    valid_until         TIMESTAMPTZ,
    max_uses_total      INT,
    max_uses_per_user   INT             NOT NULL DEFAULT 1,
    current_uses        INT             NOT NULL DEFAULT 0,
    applicable_to       VARCHAR(32)     NOT NULL DEFAULT 'ALL'
                            CHECK (applicable_to IN ('ALL','EVENT_TICKET','ACCOMMODATION','EXPERIENCE')),
    is_active           BOOLEAN         NOT NULL DEFAULT TRUE,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID,
    meta_deleted_at     TIMESTAMPTZ,
    meta_deleted_by     UUID
);

CREATE TABLE IF NOT EXISTS promo_codes (
    id                  UUID            PRIMARY KEY,
    campaign_id         UUID            NOT NULL REFERENCES promo_campaigns(id),
    code                VARCHAR(64)     NOT NULL UNIQUE,
    is_single_use       BOOLEAN         NOT NULL DEFAULT TRUE,
    assigned_user_id    UUID,           -- ref: lokatra-auth (NULL = public code)
    used_count          INT             NOT NULL DEFAULT 0,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID
);

-- Record every redemption
CREATE TABLE IF NOT EXISTS promo_redemptions (
    id                  UUID            PRIMARY KEY,
    promo_code_id       UUID            NOT NULL REFERENCES promo_codes(id),
    intent_id           UUID            NOT NULL REFERENCES payment_intents(id),
    user_id             UUID,
    discount_applied    DECIMAL(18,2)   NOT NULL,
    redeemed_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    UNIQUE (promo_code_id, intent_id)
);


-- ─────────────────────────────────────────────────────────────────────────────
-- 16. PAYMENT EVENTS  (append-only domain event log — Event Sourcing)
--     Enables full replay, SOC2 evidence, GDPR data lineage.
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payment_events (
    id                  UUID            PRIMARY KEY,
    event_type          VARCHAR(128)    NOT NULL,
    -- e.g. 'payment_intent.created', 'payment.captured', 'refund.initiated'
    aggregate_type      VARCHAR(64)     NOT NULL,  -- 'PaymentIntent','Payment','Refund'…
    aggregate_id        UUID            NOT NULL,
    sequence_number     BIGSERIAL,                  -- global ordering
    payload             JSONB           NOT NULL,   -- redacted, no raw PAN
    correlation_id      UUID,                       -- tie request chain together
    causation_id        UUID,                       -- event that caused this event
    actor_type          VARCHAR(32),    -- 'USER','SYSTEM','PSP','SCHEDULER'
    actor_id            UUID,
    ip_address          INET,
    user_agent          TEXT,
    event_at            TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    schema_version      SMALLINT        NOT NULL DEFAULT 1,
    -- PCI-DSS: data classification
    data_class          VARCHAR(16)     NOT NULL DEFAULT 'INTERNAL'
                            CHECK (data_class IN ('PUBLIC','INTERNAL','CONFIDENTIAL','RESTRICTED')),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL
    -- NO update/delete columns — this table is append-only
);

CREATE INDEX IF NOT EXISTS idx_pe_aggregate  ON payment_events (aggregate_type, aggregate_id, sequence_number);
CREATE INDEX IF NOT EXISTS idx_pe_type       ON payment_events (event_type, event_at DESC);
CREATE INDEX IF NOT EXISTS idx_pe_correlation ON payment_events (correlation_id) WHERE correlation_id IS NOT NULL;


-- ─────────────────────────────────────────────────────────────────────────────
-- 17. AUDIT LOG  (ISO 27001 A.12.4 — monitoring & logging)
--     All mutations to sensitive tables are written here by DB triggers or
--     the application layer before commit.
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payment_audit_logs (
    id                  UUID            PRIMARY KEY,
    table_name          VARCHAR(128)    NOT NULL,
    record_id           UUID            NOT NULL,
    operation           VARCHAR(8)      NOT NULL CHECK (operation IN ('INSERT','UPDATE','DELETE','SELECT')),
    changed_fields      TEXT[],
    old_values          JSONB,          -- PCI-DSS: PAN masked before write
    new_values          JSONB,
    actor_id            UUID,
    actor_type          VARCHAR(32),
    ip_address          INET,
    session_id          UUID,           -- ref: lokatra-auth user_sessions
    reason              TEXT,
    data_class          VARCHAR(16)     NOT NULL DEFAULT 'INTERNAL',
    event_at            TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL
    -- append-only: no meta_updated_* / meta_deleted_*
);

CREATE INDEX IF NOT EXISTS idx_audit_table  ON payment_audit_logs (table_name, record_id);
CREATE INDEX IF NOT EXISTS idx_audit_actor  ON payment_audit_logs (actor_id, event_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_time   ON payment_audit_logs (event_at DESC);


-- ─────────────────────────────────────────────────────────────────────────────
-- 18. GDPR / UU PDP — DATA RETENTION POLICIES
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS data_retention_policies (
    id                  UUID            PRIMARY KEY,
    table_name          VARCHAR(128)    NOT NULL UNIQUE,
    data_class          VARCHAR(16)     NOT NULL,
    retention_days      INT             NOT NULL,  -- regulatory minimum
    purge_strategy      VARCHAR(32)     NOT NULL CHECK (purge_strategy IN (
                            'DELETE',        -- hard delete
                            'ANONYMISE',     -- replace PII with tokens
                            'ARCHIVE'        -- move to cold storage
                        )),
    legal_basis         VARCHAR(64),    -- 'CONTRACT','LEGAL_OBLIGATION','LEGITIMATE_INTEREST'
    notes               TEXT,
    last_purge_run_at   TIMESTAMPTZ,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID
);

-- Track individual erasure/portability requests (GDPR Art. 17 & 20)
CREATE TABLE IF NOT EXISTS data_subject_requests (
    id                  UUID            PRIMARY KEY,
    user_id             UUID            NOT NULL,  -- ref: lokatra-auth
    merchant_id         UUID,
    request_type        VARCHAR(32)     NOT NULL CHECK (request_type IN (
                            'ERASURE',        -- right to be forgotten
                            'PORTABILITY',    -- data export
                            'RECTIFICATION',  -- correction
                            'RESTRICTION'     -- limit processing
                        )),
    status              VARCHAR(32)     NOT NULL DEFAULT 'PENDING'
                            CHECK (status IN ('PENDING','IN_PROGRESS','COMPLETED','REJECTED')),
    received_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    due_by              TIMESTAMPTZ     NOT NULL,  -- GDPR: 30 days
    completed_at        TIMESTAMPTZ,
    completed_by        UUID,
    rejection_reason    TEXT,
    affected_tables     TEXT[],
    export_url          TEXT,           -- signed URL for portability export
    notes               TEXT,
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID
);


-- ─────────────────────────────────────────────────────────────────────────────
-- 19. PSP PERFORMANCE METRICS  (rolling window — drives routing decisions)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS psp_performance_metrics (
    id                      UUID            PRIMARY KEY,
    payment_method_type     payment_method_type_enum,
    currency                payment_currency,
    window_start            TIMESTAMPTZ     NOT NULL,
    window_end              TIMESTAMPTZ     NOT NULL,
    total_attempts          INT             NOT NULL DEFAULT 0,
    success_count           INT             NOT NULL DEFAULT 0,
    failure_count           INT             NOT NULL DEFAULT 0,
    timeout_count           INT             NOT NULL DEFAULT 0,
    success_rate            DECIMAL(6,5)    GENERATED ALWAYS AS (
                                CASE WHEN total_attempts = 0 THEN 0
                                ELSE success_count::DECIMAL / total_attempts END
                            ) STORED,
    avg_latency_ms          INT,
    p99_latency_ms          INT,
    total_volume            DECIMAL(18,2)   NOT NULL DEFAULT 0,
    meta_created_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by         UUID            NOT NULL,
    meta_updated_at         TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by         UUID,
    UNIQUE (psp_account_id, payment_method_type, currency, window_start)
);

CREATE INDEX IF NOT EXISTS idx_psp_metrics_lookup
    ON psp_performance_metrics (psp_account_id, window_start DESC);


-- ─────────────────────────────────────────────────────────────────────────────
-- 20. NOTIFICATION QUEUE  (outbound payment status notifications)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS payment_notifications (
    id                  UUID            PRIMARY KEY,
    event_id            UUID            NOT NULL REFERENCES payment_events(id),
    recipient_type      VARCHAR(32)     NOT NULL CHECK (recipient_type IN (
                            'MERCHANT_WEBHOOK','USER_EMAIL','USER_SMS','USER_PUSH','INTERNAL'
                        )),
    recipient_id        UUID            NOT NULL,  -- merchant_id or user_id
    webhook_url         TEXT,
    payload             JSONB           NOT NULL DEFAULT '{}',
    status              VARCHAR(16)     NOT NULL DEFAULT 'PENDING'
                            CHECK (status IN ('PENDING','SENDING','SENT','FAILED','SKIPPED')),
    attempts            SMALLINT        NOT NULL DEFAULT 0,
    max_attempts        SMALLINT        NOT NULL DEFAULT 5,
    next_retry_at       TIMESTAMPTZ,
    last_attempt_at     TIMESTAMPTZ,
    last_error          TEXT,
    sent_at             TIMESTAMPTZ,
    response_code       SMALLINT,
    response_body       TEXT,
    signature_header    TEXT,          -- HMAC-SHA256 for merchant webhook
    meta_created_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_created_by     UUID            NOT NULL,
    meta_updated_at     TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID
);

CREATE INDEX IF NOT EXISTS idx_notif_pending ON payment_notifications (status, next_retry_at)
    WHERE status IN ('PENDING','FAILED');


-- ─────────────────────────────────────────────────────────────────────────────
-- 21. DEFERRED FOREIGN KEYS  (cross-table UUIDs resolved after table creation)
-- ─────────────────────────────────────────────────────────────────────────────

ALTER TABLE routing_decisions
    ADD CONSTRAINT fk_rd_intent
        FOREIGN KEY (payment_intent_id) REFERENCES payment_intents(id);

ALTER TABLE payments
    ADD CONSTRAINT fk_pay_risk
        FOREIGN KEY (risk_score_id) REFERENCES risk_assessments(id);

ALTER TABLE payments
    ADD CONSTRAINT fk_pay_fx
        FOREIGN KEY (fx_rate_snapshot_id) REFERENCES fx_rate_snapshots(id);


-- ─────────────────────────────────────────────────────────────────────────────
-- 22. CORE INDEXES (additional performance)
-- ─────────────────────────────────────────────────────────────────────────────

CREATE INDEX IF NOT EXISTS idx_intent_merchant_date
    ON payment_intents (merchant_id, meta_created_at DESC)
    WHERE meta_deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_payments_created
    ON payments (meta_created_at DESC)
    WHERE meta_deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_pm_user_active
    ON payment_methods (user_id, is_active, method_type)
    WHERE meta_deleted_at IS NULL AND gdpr_erased_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_promo_code_lookup
    ON promo_codes (code) WHERE meta_updated_at IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_data_subject_requests_user
    ON data_subject_requests (user_id, status, due_by);


-- ─────────────────────────────────────────────────────────────────────────────
-- 23. ROW-LEVEL SECURITY HINTS
--     Enable per-merchant data isolation (enforce at application / PostgREST layer)
-- ─────────────────────────────────────────────────────────────────────────────
-- Example (apply after setting app.current_merchant_id in session):
--
-- ALTER TABLE payment_intents ENABLE ROW LEVEL SECURITY;
-- CREATE POLICY rls_intents_merchant ON payment_intents
--     USING (merchant_id = current_setting('app.current_merchant_id')::UUID);
--
-- Repeat for: payment_methods, payments, refunds, payouts,
--             settlement_batches, promo_campaigns.
-- ─────────────────────────────────────────────────────────────────────────────


-- ─────────────────────────────────────────────────────────────────────────────
-- 24. DEFAULT RETENTION POLICIES
-- ─────────────────────────────────────────────────────────────────────────────

INSERT INTO data_retention_policies
    (id, table_name, data_class, retention_days, purge_strategy, legal_basis, notes,
     meta_created_at, meta_created_by, meta_updated_at, meta_updated_by)
VALUES
    (gen_random_uuid(), 'payment_intents',    'CONFIDENTIAL', 2555, 'ANONYMISE', 'LEGAL_OBLIGATION', 'OJK 7 year; anonymise PII after retention', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'payments',           'CONFIDENTIAL', 2555, 'ANONYMISE', 'LEGAL_OBLIGATION', 'OJK 7 year', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'payment_methods',    'RESTRICTED',   365,  'DELETE',    'CONTRACT',         'Delete on contract end or GDPR erasure', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'payment_audit_logs', 'RESTRICTED',   2555, 'ARCHIVE',   'LEGAL_OBLIGATION', 'ISO 27001 / SOC2 — move to cold storage', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'payment_events',     'CONFIDENTIAL', 2555, 'ARCHIVE',   'LEGAL_OBLIGATION', 'Event sourcing archive', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'psp_webhooks',       'INTERNAL',     90,   'DELETE',    'LEGITIMATE_INTEREST', 'Short-lived; 90d for dispute window', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'idempotency_keys',   'INTERNAL',     1,    'DELETE',    'LEGITIMATE_INTEREST', 'TTL 24h already on row; purge daily', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL),
    (gen_random_uuid(), 'risk_assessments',   'CONFIDENTIAL', 365,  'ANONYMISE', 'LEGITIMATE_INTEREST', 'Anonymise IP/device after 1 year', NOW(), '00000000-0000-0000-0000-000000000000', NOW(), NULL)
ON CONFLICT (table_name) DO NOTHING;


-- =============================================================================
-- END OF SCHEMA  |  lokatra-payment  v1.0.0
-- =============================================================================
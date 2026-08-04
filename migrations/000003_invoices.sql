-- =============================================================================
-- LOKATRA-PAYMENT  |  Invoices  v1.0.0
-- =============================================================================
-- Coverage:
--   PPh 23/26 (Income Tax) — stored as separate columns per line item and invoice
--   PPN/VAT 11%            — stored as separate columns per line item and invoice
--   All monetary values use decimal(19,4) — never float
--   Soft delete on every table
--   Immutable audit trail via invoice_status_events
-- =============================================================================

-- ── ENUMS ─────────────────────────────────────────────────────────────────────

CREATE TYPE invoice_status_enum AS ENUM (
    'draft', 'issued', 'partially_paid', 'paid', 'overdue', 'voided', 'written_off'
);

CREATE TYPE payment_link_status_enum AS ENUM ('active', 'disabled', 'expired');

CREATE TYPE invoice_payment_status_enum AS ENUM ('pending', 'processing', 'succeeded', 'failed', 'refunded');

-- ── 1. INVOICES ───────────────────────────────────────────────────────────────

CREATE TABLE invoices (
    id                      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_code            text NOT NULL UNIQUE,

    -- Parties (cross-service FKs are plain UUIDs per service autonomy policy)
    merchant_id             uuid NOT NULL,
    customer_id             uuid NULL,
    customer_name           text NULL,
    customer_email          text NULL,
    customer_phone          text NULL,

    -- Amounts — all decimal(19,4), never float
    subtotal                decimal(19,4) NOT NULL DEFAULT 0 CHECK (subtotal >= 0),
    discount_amount         decimal(19,4) NOT NULL DEFAULT 0 CHECK (discount_amount >= 0),
    pph23_amount            decimal(19,4) NOT NULL DEFAULT 0 CHECK (pph23_amount >= 0),
    pph23_rate              decimal(5,2)  NOT NULL DEFAULT 0 CHECK (pph23_rate >= 0 AND pph23_rate <= 100),
    ppn_amount              decimal(19,4) NOT NULL DEFAULT 0 CHECK (ppn_amount >= 0),
    ppn_rate                decimal(5,2)  NOT NULL DEFAULT 11 CHECK (ppn_rate >= 0 AND ppn_rate <= 100),
    total_amount            decimal(19,4) NOT NULL DEFAULT 0 CHECK (total_amount > 0),
    paid_amount             decimal(19,4) NOT NULL DEFAULT 0 CHECK (paid_amount >= 0),
    remaining_amount        decimal(19,4) NOT NULL DEFAULT 0 CHECK (remaining_amount >= 0),
    currency                char(3) NOT NULL DEFAULT 'IDR',

    -- Status
    status                  invoice_status_enum NOT NULL DEFAULT 'draft',

    -- Dates
    issued_at               timestamptz NULL,
    due_at                  timestamptz NULL,
    paid_at                 timestamptz NULL,
    voided_at               timestamptz NULL,
    void_reason             text NULL,

    -- External reference
    source_service          text NULL,
    source_type             text NULL,
    source_id               uuid NULL,

    -- Flexible data
    description             text NULL,
    notes                   text NULL,
    terms                   text NULL,

    -- Audit
    meta_created_at         timestamptz NOT NULL DEFAULT now(),
    meta_created_by         uuid NOT NULL,
    meta_updated_at         timestamptz NOT NULL DEFAULT now(),
    meta_updated_by         uuid NULL,
    meta_deleted_at         timestamptz NULL,
    meta_deleted_by         uuid NULL
);

CREATE INDEX idx_invoices_code        ON invoices(invoice_code);
CREATE INDEX idx_invoices_merchant    ON invoices(merchant_id, status);
CREATE INDEX idx_invoices_customer    ON invoices(customer_id, status);
CREATE INDEX idx_invoices_due         ON invoices(due_at) WHERE status IN ('issued', 'partially_paid');
CREATE INDEX idx_invoices_source      ON invoices(source_service, source_type, source_id);

-- ── 2. INVOICE LINE ITEMS ─────────────────────────────────────────────────────

CREATE TABLE invoice_line_items (
    id                      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id              uuid NOT NULL REFERENCES invoices(id),
    line_no                 int NOT NULL CHECK (line_no > 0),

    -- Item details
    name                    text NOT NULL,
    description             text NULL,
    quantity                decimal(10,4) NOT NULL DEFAULT 1 CHECK (quantity > 0),
    unit_price              decimal(19,4) NOT NULL DEFAULT 0 CHECK (unit_price >= 0),

    -- Discount
    discount_percent        decimal(5,2) NULL CHECK (discount_percent IS NULL OR (discount_percent >= 0 AND discount_percent <= 100)),
    discount_amount         decimal(19,4) NOT NULL DEFAULT 0 CHECK (discount_amount >= 0),

    -- PPN (VAT)
    ppn_rate                decimal(5,2) NOT NULL DEFAULT 11 CHECK (ppn_rate >= 0 AND ppn_rate <= 100),
    ppn_amount              decimal(19,4) NOT NULL DEFAULT 0 CHECK (ppn_amount >= 0),

    -- PPh 23 (Income Tax)
    pph23_rate              decimal(5,2) NOT NULL DEFAULT 0 CHECK (pph23_rate >= 0 AND pph23_rate <= 100),
    pph23_amount            decimal(19,4) NOT NULL DEFAULT 0 CHECK (pph23_amount >= 0),

    -- Totals
    subtotal                decimal(19,4) NOT NULL DEFAULT 0 CHECK (subtotal >= 0),
    total_amount            decimal(19,4) NOT NULL DEFAULT 0 CHECK (total_amount >= 0),
    currency                char(3) NOT NULL DEFAULT 'IDR',

    -- Metadata
    sku                     text NULL,
    category                text NULL,

    -- Audit
    meta_created_at         timestamptz NOT NULL DEFAULT now(),
    meta_created_by         uuid NOT NULL,
    meta_updated_at         timestamptz NOT NULL DEFAULT now(),
    meta_updated_by         uuid NULL,
    meta_deleted_at         timestamptz NULL,
    meta_deleted_by         uuid NULL,

    UNIQUE(invoice_id, line_no)
);

CREATE INDEX idx_invoice_line_items_invoice ON invoice_line_items(invoice_id);

-- ── 3. INVOICE PAYMENT LINKS ─────────────────────────────────────────────────

CREATE TABLE invoice_payment_links (
    id                      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id              uuid NOT NULL REFERENCES invoices(id),
    provider_code           citext NOT NULL,

    -- Link details
    link_url                text NOT NULL,
    link_type               text NOT NULL DEFAULT 'redirect',

    -- Status
    status                  payment_link_status_enum NOT NULL DEFAULT 'active',

    -- Tracking
    click_count             int NOT NULL DEFAULT 0,
    payment_count           int NOT NULL DEFAULT 0,
    expires_at              timestamptz NULL,

    -- Audit
    meta_created_at         timestamptz NOT NULL DEFAULT now(),
    meta_created_by         uuid NOT NULL,
    meta_updated_at         timestamptz NOT NULL DEFAULT now(),
    meta_updated_by         uuid NULL,
    meta_deleted_at         timestamptz NULL,
    meta_deleted_by         uuid NULL,

    UNIQUE(invoice_id, provider_code)
);

CREATE INDEX idx_invoice_payment_links_invoice ON invoice_payment_links(invoice_id);

-- ── 4. INVOICE PAYMENTS (bridge: invoice ↔ payment_intent) ────────────────────

CREATE TABLE invoice_payments (
    id                      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id              uuid NOT NULL REFERENCES invoices(id),
    payment_intent_id       uuid NOT NULL REFERENCES payment_intents(id),
    provider_code           citext NOT NULL,

    amount                  decimal(19,4) NOT NULL CHECK (amount > 0),
    currency                char(3) NOT NULL DEFAULT 'IDR',
    status                  invoice_payment_status_enum NOT NULL DEFAULT 'pending',

    -- Audit
    meta_created_at         timestamptz NOT NULL DEFAULT now(),
    meta_created_by         uuid NOT NULL,
    meta_updated_at         timestamptz NOT NULL DEFAULT now(),
    meta_updated_by         uuid NULL
);

CREATE INDEX idx_invoice_payments_invoice ON invoice_payments(invoice_id);
CREATE INDEX idx_invoice_payments_intent  ON invoice_payments(payment_intent_id);

-- ── 5. INVOICE STATUS EVENTS (immutable audit trail) ─────────────────────────

CREATE TABLE invoice_status_events (
    id                      uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id              uuid NOT NULL REFERENCES invoices(id),
    event_type              text NOT NULL,
    old_status              text NULL,
    new_status              text NULL,
    reason                  text NULL,
    actor_id                uuid NULL,

    -- Audit (append-only: meta_updated_at/meta_updated_by/meta_deleted_at/meta_deleted_by omitted by design)
    meta_created_at         timestamptz NOT NULL DEFAULT now(),
    meta_created_by         uuid NOT NULL
);

CREATE INDEX idx_invoice_status_events_invoice ON invoice_status_events(invoice_id, meta_created_at);

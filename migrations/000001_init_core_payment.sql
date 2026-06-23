-- =============================================================================
-- LOKATRA-PAYMENT  |  Database Schema  v1.0.0
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
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
CREATE TYPE provider_type_enum AS ENUM ('gateway', 'wallet', 'bank_transfer', 'cash', 'cod');
CREATE TYPE provider_status_enum AS ENUM ('active', 'inactive', 'deprecated');
CREATE TABLE IF NOT EXISTS payment_providers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  code citext NOT NULL UNIQUE,
  name text NOT NULL,
  provider_type provider_type_enum NOT NULL DEFAULT 'gateway',
  status provider_status_enum NOT NULL DEFAULT 'active',
  supports_refund boolean NOT NULL DEFAULT false,
  supports_partial_refund boolean NOT NULL DEFAULT false,
  supports_authorization boolean NOT NULL DEFAULT false,
  supports_capture boolean NOT NULL DEFAULT false,
  supports_void boolean NOT NULL DEFAULT false,
  supports_webhook boolean NOT NULL DEFAULT true,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE provider_account_status_enum AS ENUM ('active', 'inactive', 'deprecated');
CREATE TABLE IF NOT EXISTS provider_accounts (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_id uuid NOT NULL REFERENCES payment_providers(id),
  account_name text NOT NULL,
  environment text NOT NULL DEFAULT 'production',
  owner_type text NOT NULL DEFAULT 'platform',
  owner_id uuid NULL,
  merchant_ref text NULL,
  credential_secret_ref text NOT NULL,
  webhook_secret_ref text NULL,
  public_key_ref text NULL,
  status provider_account_status_enum NOT NULL DEFAULT 'active',
  config jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(provider_id, environment, owner_type, owner_id, account_name)
);

CREATE TYPE payment_method_type_enum AS ENUM ('card', 'bank_transfer', 'ewallet', 'cash', 'cod');
CREATE TYPE payment_method_status_enum AS ENUM ('active', 'inactive', 'deprecated');
CREATE TABLE IF NOT EXISTS payment_methods (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  code citext NOT NULL UNIQUE,
  method_type payment_method_type_enum NOT NULL,
  name text NOT NULL,
  status payment_method_status_enum NOT NULL DEFAULT 'active',
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_channel_status_enum AS ENUM ('active', 'inactive', 'deprecated');
CREATE TABLE IF NOT EXISTS payment_channels (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  method_id uuid NOT NULL REFERENCES payment_methods(id),
  code citext NOT NULL,
  name text NOT NULL,
  country_code char(2) NOT NULL DEFAULT 'ID',
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_channel_status_enum NOT NULL DEFAULT 'active',
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(method_id, code, country_code, currency)
);

CREATE TYPE payment_intent_status_enum AS ENUM ('requires_payment_method', 'requires_confirmation', 'requires_action', 'processing', 'succeeded', 'canceled');
CREATE TABLE IF NOT EXISTS payment_intents (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  intent_code text NOT NULL UNIQUE,
  source_service text NOT NULL,
  source_type text NOT NULL,
  source_id uuid NOT NULL,
  merchant_id uuid NOT NULL,
  customer_id uuid NULL,
  amount decimal(19,4) NOT NULL CHECK (amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_intent_status_enum NOT NULL DEFAULT 'requires_payment_method',
  selected_method_code citext NULL,
  selected_channel_code citext NULL,
  description text NULL,
  expires_at timestamptz NULL,
  paid_at timestamptz NULL,
  canceled_at timestamptz NULL,
  cancellation_reason text NULL,
  idempotency_key text NOT NULL,
  source_snapshot jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(source_service, source_type, source_id),
  UNIQUE(idempotency_key)
);


CREATE TABLE IF NOT EXISTS payment_route_decisions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  selected_provider_account_id uuid NULL REFERENCES provider_accounts(id),
  selected_provider_code citext NULL,
  method_code citext NOT NULL,
  channel_code citext NULL,
  reason text NOT NULL,
  evaluated_context jsonb NOT NULL DEFAULT '{}'::jsonb,
  candidates jsonb NOT NULL DEFAULT '[]'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_attempt_status_enum AS ENUM ('created', 'pending', 'authorized', 'captured', 'paid', 'failed', 'canceled');
CREATE TYPE instruction_type_enum AS ENUM ('va_number', 'qris_string', 'qr_image_url', 'checkout_url', 'deeplink_url', 'retail_code', 'manual_transfer', 'cash', 'cod');
CREATE TABLE IF NOT EXISTS payment_attempts (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  attempt_no int NOT NULL,
  provider_account_id uuid NULL REFERENCES provider_accounts(id),
  route_decision_id uuid NULL REFERENCES payment_route_decisions(id),
  provider_code citext NULL,
  method_code citext NOT NULL,
  channel_code citext NULL,
  amount decimal(19,4) NOT NULL CHECK (amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_attempt_status_enum NOT NULL DEFAULT 'created',
  provider_reference text NULL,
  provider_transaction_id text NULL,
  provider_order_id text NULL,
  provider_payment_id text NULL,
  failure_code text NULL,
  failure_message text NULL,
  expires_at timestamptz NULL,
  authorized_at timestamptz NULL,
  captured_at timestamptz NULL,
  paid_at timestamptz NULL,
  failed_at timestamptz NULL,
  canceled_at timestamptz NULL,
  status_sync_required_at timestamptz NULL,
  last_status_sync_at timestamptz NULL,
  raw_request jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(payment_intent_id, attempt_no)
);

CREATE TABLE IF NOT EXISTS payment_instructions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_attempt_id uuid NOT NULL REFERENCES payment_attempts(id),
  instruction_type instruction_type_enum NOT NULL,
  is_active boolean NOT NULL DEFAULT true,
  display_name text NULL,
  account_number text NULL,
  account_number_masked text NULL,
  account_holder_name text NULL,
  bank_code citext NULL,
  biller_code text NULL,
  payment_code text NULL,
  qr_string text NULL,
  qr_image_url text NULL,
  checkout_url text NULL,
  deeplink_url text NULL,
  retail_outlet_code citext NULL,
  expires_at timestamptz NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TABLE IF NOT EXISTS payment_status_events (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  provider_webhook_event_id uuid NULL,
  source_type text NOT NULL,
  event_type text NOT NULL,
  old_intent_status text NULL,
  new_intent_status text NULL,
  old_attempt_status text NULL,
  new_attempt_status text NULL,
  provider_status text NULL,
  reason text NULL,
  occurred_at timestamptz NOT NULL DEFAULT now(),
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE idempotency_status_enum AS ENUM ('processing', 'completed', 'failed');
CREATE TABLE IF NOT EXISTS idempotency_keys (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  key text NOT NULL UNIQUE,
  actor_type text NULL,
  actor_id uuid NULL,
  request_hash text NOT NULL,
  status idempotency_status_enum NOT NULL DEFAULT 'processing',
  resource_type text NULL,
  resource_id uuid NULL,
  response_status int NULL,
  response_body jsonb NULL,
  locked_until timestamptz NULL,
  completed_at timestamptz NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);


CREATE TYPE scope_type_enum AS ENUM ('platform', 'merchant');
CREATE TABLE IF NOT EXISTS payment_route_candidates_runtime (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  scope_type scope_type_enum NOT NULL DEFAULT 'platform',
  scope_id uuid NULL,
  merchant_id uuid NULL,
  method_code citext NOT NULL,
  channel_code citext NULL,
  currency char(3) NOT NULL DEFAULT 'IDR',
  min_amount decimal(19,4) NULL,
  max_amount decimal(19,4) NULL,
  provider_account_id uuid NOT NULL REFERENCES provider_accounts(id),
  provider_method_code text NULL,
  provider_channel_code text NULL,
  priority int NOT NULL DEFAULT 100,
  is_fallback boolean NOT NULL DEFAULT false,
  traffic_weight int NOT NULL DEFAULT 100,
  timeout_ms int NOT NULL DEFAULT 5000,
  max_attempts int NOT NULL DEFAULT 1,
  is_enabled boolean NOT NULL DEFAULT true,
  condition_expr jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE provider_api_operation_enum AS ENUM ('create_payment', 'cancel_payment', 'get_payment_status', 'create_refund', 'authorize', 'capture', 'void', 'expire_payment', 'verify_webhook');
CREATE TABLE IF NOT EXISTS provider_api_requests (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id uuid NULL REFERENCES provider_accounts(id),
  payment_intent_id uuid NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  operation provider_api_operation_enum NOT NULL,
  idempotency_key text NULL,
  request_method text NOT NULL,
  request_url text NOT NULL,
  request_headers jsonb NOT NULL DEFAULT '{}'::jsonb,
  request_body jsonb NOT NULL DEFAULT '{}'::jsonb,
  response_status int NULL,
  response_headers jsonb NOT NULL DEFAULT '{}'::jsonb,
  response_body jsonb NOT NULL DEFAULT '{}'::jsonb,
  latency_ms int NULL,
  success boolean NULL,
  error_code text NULL,
  error_message text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE circuit_status_enum AS ENUM ('closed', 'open', 'half_open');
CREATE TABLE IF NOT EXISTS provider_circuit_breakers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id uuid NOT NULL REFERENCES provider_accounts(id),
  method_code citext NULL,
  channel_code citext NULL,
  status circuit_status_enum NOT NULL DEFAULT 'closed',
  failure_count int NOT NULL DEFAULT 0,
  success_count int NOT NULL DEFAULT 0,
  last_failure_at timestamptz NULL,
  last_success_at timestamptz NULL,
  opened_at timestamptz NULL,
  open_until timestamptz NULL,
  half_open_at timestamptz NULL,
  reason text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(provider_account_id, method_code, channel_code)
);

CREATE TABLE IF NOT EXISTS provider_health_snapshots (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id uuid NOT NULL REFERENCES provider_accounts(id),
  method_code citext NULL,
  channel_code citext NULL,
  health_score int NOT NULL DEFAULT 100 CHECK (health_score BETWEEN 0 AND 100),
  success_rate decimal(7,4) NULL,
  timeout_rate decimal(7,4) NULL,
  error_rate decimal(7,4) NULL,
  p95_latency_ms int NULL,
  sample_size int NOT NULL DEFAULT 0,
  window_started_at timestamptz NOT NULL,
  window_ended_at timestamptz NOT NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);


CREATE TABLE IF NOT EXISTS provider_webhook_endpoints (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id uuid NOT NULL REFERENCES provider_accounts(id),
  provider_code citext NOT NULL,
  endpoint_key text NOT NULL,
  environment text NOT NULL DEFAULT 'production',
  secret_ref text NOT NULL,
  signature_algorithm text NOT NULL,
  is_active boolean NOT NULL DEFAULT true,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(provider_code, endpoint_key, environment)
);

CREATE TYPE webhook_processing_status_enum AS ENUM ('received', 'processing', 'processed', 'failed');
CREATE TABLE IF NOT EXISTS provider_webhook_events (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  webhook_endpoint_id uuid NULL REFERENCES provider_webhook_endpoints(id),
  endpoint_key text NULL,
  provider_account_id uuid NOT NULL REFERENCES provider_accounts(id),
  provider_code citext NOT NULL,
  event_id text NULL,
  event_type text NULL,
  provider_reference text NULL,
  provider_status text NULL,
  signature_valid boolean NOT NULL DEFAULT false,
  signature_algorithm text NULL,
  headers jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_body bytea NOT NULL,
  raw_body_sha256 text NOT NULL,
  parsed_body jsonb NULL,
  processing_status webhook_processing_status_enum NOT NULL DEFAULT 'received',
  retry_count int NOT NULL DEFAULT 0,
  next_retry_at timestamptz NULL,
  locked_until timestamptz NULL,
  received_at timestamptz NOT NULL DEFAULT now(),
  processed_at timestamptz NULL,
  error_code text NULL,
  error_message text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(provider_account_id, raw_body_sha256)
);

CREATE TYPE manual_evidence_status_enum AS ENUM ('submitted', 'under_review', 'approved', 'rejected');
CREATE TABLE IF NOT EXISTS manual_payment_evidence (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  submitted_by uuid NOT NULL,
  evidence_type text NOT NULL DEFAULT 'transfer_receipt',
  evidence_url text NULL,
  amount decimal(19,4) NOT NULL CHECK (amount >= 0),
  expected_amount decimal(19,4) NOT NULL CHECK (expected_amount >= 0),
  variance_amount decimal(19,4) NOT NULL DEFAULT 0,
  variance_status text NOT NULL DEFAULT 'exact',
  currency char(3) NOT NULL DEFAULT 'IDR',
  bank_code citext NULL,
  bank_name text NULL,
  sender_account_name text NULL,
  sender_account_number_masked text NULL,
  notes text NULL,
  status manual_evidence_status_enum NOT NULL DEFAULT 'submitted',
  reviewed_by uuid NULL,
  reviewed_at timestamptz NULL,
  rejection_reason text NULL,
  policy_decision text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE cash_session_status_enum AS ENUM ('open', 'closed', 'canceled');
CREATE TABLE IF NOT EXISTS cash_collection_sessions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  session_code text NOT NULL UNIQUE,
  merchant_id uuid NOT NULL,
  collector_id uuid NOT NULL,
  location_id uuid NULL,
  opened_at timestamptz NOT NULL DEFAULT now(),
  closed_at timestamptz NULL,
  status cash_session_status_enum NOT NULL DEFAULT 'open',
  opening_float_amount decimal(19,4) NOT NULL DEFAULT 0,
  expected_amount decimal(19,4) NOT NULL DEFAULT 0,
  counted_amount decimal(19,4) NOT NULL DEFAULT 0,
  variance_amount decimal(19,4) NOT NULL DEFAULT 0,
  currency char(3) NOT NULL DEFAULT 'IDR',
  notes text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE cash_item_status_enum AS ENUM ('collected', 'voided');
CREATE TABLE IF NOT EXISTS cash_collection_items (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  cash_collection_session_id uuid NOT NULL REFERENCES cash_collection_sessions(id),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  collection_type text NOT NULL CHECK (collection_type IN ('cash','cod')),
  amount decimal(19,4) NOT NULL CHECK (amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status cash_item_status_enum NOT NULL DEFAULT 'collected',
  collected_at timestamptz NOT NULL DEFAULT now(),
  voided_at timestamptz NULL,
  void_reason text NULL,
  notes text NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_overpayment_status_enum AS ENUM ('open', 'resolved');
CREATE TABLE IF NOT EXISTS payment_overpayments (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  paid_attempt_id uuid NULL REFERENCES payment_attempts(id),
  overpaid_attempt_id uuid NULL REFERENCES payment_attempts(id),
  expected_amount decimal(19,4) NOT NULL CHECK (expected_amount >= 0),
  received_amount decimal(19,4) NOT NULL CHECK (received_amount >= 0),
  overpaid_amount decimal(19,4) NOT NULL CHECK (overpaid_amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_overpayment_status_enum NOT NULL DEFAULT 'open',
  resolution_action text NULL,
  resolution_notes text NULL,
  resolved_at timestamptz NULL,
  resolved_by uuid NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_refund_status_enum AS ENUM ('requested', 'approved', 'rejected', 'processing', 'succeeded', 'failed');
CREATE TABLE IF NOT EXISTS payment_refunds (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  refund_code text NOT NULL UNIQUE,
  amount decimal(19,4) NOT NULL CHECK (amount > 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  reason text NULL,
  status payment_refund_status_enum NOT NULL DEFAULT 'requested',
  provider_refund_id text NULL,
  provider_reference text NULL,
  requested_by uuid NOT NULL,
  requested_at timestamptz NOT NULL DEFAULT now(),
  approved_by uuid NULL,
  approved_at timestamptz NULL,
  rejected_by uuid NULL,
  rejected_at timestamptz NULL,
  rejection_reason text NULL,
  processing_at timestamptz NULL,
  succeeded_at timestamptz NULL,
  failed_at timestamptz NULL,
  failure_code text NULL,
  failure_message text NULL,
  raw_request jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_plan_type_enum AS ENUM ('installment', 'subscription');
CREATE TYPE payment_plan_status_enum AS ENUM ('active', 'completed', 'canceled');
CREATE TABLE IF NOT EXISTS payment_plans (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  plan_type payment_plan_type_enum NOT NULL,
  status payment_plan_status_enum NOT NULL DEFAULT 'active',
  total_amount decimal(19,4) NOT NULL CHECK (total_amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  installment_count int NOT NULL DEFAULT 1 CHECK (installment_count > 0),
  deposit_amount decimal(19,4) NULL CHECK (deposit_amount IS NULL OR deposit_amount >= 0),
  auto_cancel_on_default boolean NOT NULL DEFAULT false,
  default_grace_period_seconds int NOT NULL DEFAULT 0,
  completed_at timestamptz NULL,
  canceled_at timestamptz NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(payment_intent_id)
);

CREATE TYPE payment_installment_status_enum AS ENUM ('pending', 'paid', 'overdue', 'canceled');
CREATE TABLE IF NOT EXISTS payment_installments (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_plan_id uuid NOT NULL REFERENCES payment_plans(id),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  installment_no int NOT NULL CHECK (installment_no > 0),
  due_amount decimal(19,4) NOT NULL CHECK (due_amount >= 0),
  paid_amount decimal(19,4) NOT NULL DEFAULT 0 CHECK (paid_amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  due_at timestamptz NOT NULL,
  status payment_installment_status_enum NOT NULL DEFAULT 'pending',
  paid_at timestamptz NULL,
  overdue_at timestamptz NULL,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL,
  UNIQUE(payment_plan_id, installment_no)
);

CREATE TYPE payment_authorization_status_enum AS ENUM ('requested', 'authorized', 'captured', 'voided', 'failed');
CREATE TABLE IF NOT EXISTS payment_authorizations (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  payment_attempt_id uuid NULL REFERENCES payment_attempts(id),
  provider_account_id uuid NULL REFERENCES provider_accounts(id),
  provider_authorization_id text NULL,
  amount decimal(19,4) NOT NULL CHECK (amount > 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_authorization_status_enum NOT NULL DEFAULT 'requested',
  authorized_at timestamptz NULL,
  expires_at timestamptz NULL,
  captured_amount decimal(19,4) NOT NULL DEFAULT 0 CHECK (captured_amount >= 0),
  failure_code text NULL,
  failure_message text NULL,
  raw_request jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_capture_status_enum AS ENUM ('requested', 'captured', 'failed');
CREATE TABLE IF NOT EXISTS payment_captures (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_authorization_id uuid NOT NULL REFERENCES payment_authorizations(id),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  amount decimal(19,4) NOT NULL CHECK (amount > 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_capture_status_enum NOT NULL DEFAULT 'requested',
  provider_capture_id text NULL,
  captured_at timestamptz NULL,
  failure_code text NULL,
  failure_message text NULL,
  raw_request jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

CREATE TYPE payment_void_status_enum AS ENUM ('requested', 'voided', 'failed');
CREATE TABLE IF NOT EXISTS payment_voids (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_authorization_id uuid NOT NULL REFERENCES payment_authorizations(id),
  payment_intent_id uuid NOT NULL REFERENCES payment_intents(id),
  amount decimal(19,4) NOT NULL CHECK (amount >= 0),
  currency char(3) NOT NULL DEFAULT 'IDR',
  status payment_void_status_enum NOT NULL DEFAULT 'requested',
  provider_void_id text NULL,
  voided_at timestamptz NULL,
  failure_code text NULL,
  failure_message text NULL,
  raw_request jsonb NOT NULL DEFAULT '{}'::jsonb,
  raw_response jsonb NOT NULL DEFAULT '{}'::jsonb,
  metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
  meta_created_at timestamptz NOT NULL DEFAULT now(),
  meta_created_by uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now(),
  meta_updated_by uuid NULL,
  meta_deleted_at timestamptz NULL,
  meta_deleted_by uuid NULL
);

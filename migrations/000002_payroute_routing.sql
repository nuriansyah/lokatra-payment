-- =============================================================================
-- PAYROUTE — Intelligent Routing & Orchestration Tables
-- Migration: 000002_payroute_routing.sql
-- Version: 1.0.0
-- =============================================================================
-- This migration adds DB-driven routing rules, health monitoring,
-- MDR cost data, audit logging, and kill-switch configuration.
--
-- PCI-DSS v4: No new card data storage; orchestration only.
-- SOC 2 CC6/CC7/CC8: Audit log, access control, change management.
-- RFC 7807: Structured error storage preserved from v1.
-- =============================================================================

-- ---------------------------------------------------------------------------
-- 1. ENUMS
-- ---------------------------------------------------------------------------

CREATE TYPE payroute_routing_strategy AS ENUM (
  'priority',
  'success_rate',
  'cost_aware',
  'weighted',
  'combined'
);

CREATE TYPE payroute_rule_status AS ENUM (
  'draft',
  'active',
  'rolling_out',
  'disabled',
  'archived'
);

CREATE TYPE payroute_rule_scope AS ENUM (
  'global',
  'method',
  'method_channel',
  'method_amount_band'
);

-- ---------------------------------------------------------------------------
-- 2. ROUTING RULES — Core orchestration configuration
-- ---------------------------------------------------------------------------
-- Each rule maps a scope (method/channel/amount) to an ordered list of PSPs
-- with a routing strategy. Rules are versioned for rollback support.

CREATE TABLE payroute_routing_rules (
  id                    uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name                  text NOT NULL,
  description           text NULL,

  -- Scope definition
  scope                 payroute_rule_scope NOT NULL DEFAULT 'global',
  scope_config          jsonb NOT NULL DEFAULT '{}'::jsonb,
  -- scope_config shape:
  --   { "method_code": "virtual_account", "channel_code": "bca_va",
  --     "currency": "IDR", "min_amount": 0, "max_amount": 50000000 }

  -- Strategy
  strategy              payroute_routing_strategy NOT NULL DEFAULT 'priority',
  strategy_config       jsonb NOT NULL DEFAULT '{}'::jsonb,
  -- strategy_config shape:
  --   { "window_seconds": 3600, "min_sample_size": 30,
  --     "mdr_weight": 0.3, "success_rate_weight": 0.7,
  --     "shadow_percentage": 0 }

  -- PSP priority list
  psp_list              jsonb NOT NULL DEFAULT '[]'::jsonb,
  -- psp_list shape:
  --   [{ "provider_account_id": "uuid", "priority": 1,
  --      "traffic_weight": 100, "is_fallback": false,
  --      "max_attempts": 2, "timeout_ms": 5000,
  --      "provider_method_code": "VIRTUAL_ACCOUNT",
  --      "provider_channel_code": "BCA" }]

  -- Versioning
  version               int NOT NULL DEFAULT 1,

  -- Status & rollout
  status                payroute_rule_status NOT NULL DEFAULT 'draft',
  rollout_percentage    int NOT NULL DEFAULT 100
                        CHECK (rollout_percentage BETWEEN 0 AND 100),
  rollout_started_at    timestamptz NULL,
  rollout_rollback_version int NULL,

  -- Audit
  created_by            uuid NOT NULL,
  approved_by           uuid NULL,

  -- Soft delete
  meta_created_at       timestamptz NOT NULL DEFAULT now(),
  meta_updated_at       timestamptz NOT NULL DEFAULT now(),
  meta_deleted_at       timestamptz NULL,
  meta_deleted_by       uuid NULL
);

-- Uniqueness: only one active rule per scope combination
CREATE UNIQUE INDEX idx_routing_rules_unique_active
  ON payroute_routing_rules(scope, (scope_config->>'method_code'), (scope_config->>'channel_code'))
  WHERE status IN ('active', 'rolling_out') AND meta_deleted_at IS NULL;

CREATE INDEX idx_routing_rules_status
  ON payroute_routing_rules(status) WHERE meta_deleted_at IS NULL;

CREATE INDEX idx_routing_rules_lookup
  ON payroute_routing_rules(status, (scope_config->>'method_code'))
  WHERE meta_deleted_at IS NULL;

-- ---------------------------------------------------------------------------
-- 3. ROUTING RULE VERSIONS — Immutable snapshots for rollback
-- ---------------------------------------------------------------------------

CREATE TABLE payroute_routing_rule_versions (
  id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  rule_id         uuid NOT NULL REFERENCES payroute_routing_rules(id),
  version         int NOT NULL,
  snapshot        jsonb NOT NULL,
  changed_by      uuid NOT NULL,
  change_reason   text NULL,
  meta_created_at timestamptz NOT NULL DEFAULT now(),

  UNIQUE(rule_id, version)
);

CREATE INDEX idx_rule_versions_rule_id ON payroute_routing_rule_versions(rule_id);

-- ---------------------------------------------------------------------------
-- 4. MDR RATES — Merchant Discount Rates for cost-aware routing
-- ---------------------------------------------------------------------------

CREATE TABLE payroute_mdr_rates (
  id                    uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id   uuid NOT NULL REFERENCES provider_accounts(id),
  payment_method_id     uuid NOT NULL REFERENCES payment_methods(id),
  payment_channel_id    uuid NULL REFERENCES payment_channels(id),

  percentage            decimal(7,4) NOT NULL,
  fixed_fee             decimal(19,4) NOT NULL DEFAULT 0,
  min_amount            decimal(19,4) NULL,
  max_amount            decimal(19,4) NULL,
  currency              char(3) NOT NULL DEFAULT 'IDR',

  effective_from        timestamptz NOT NULL DEFAULT now(),
  effective_to          timestamptz NULL,

  meta_created_at       timestamptz NOT NULL DEFAULT now(),
  meta_updated_at       timestamptz NOT NULL DEFAULT now(),
  meta_deleted_at       timestamptz NULL
);

CREATE INDEX idx_mdr_rates_lookup
  ON payroute_mdr_rates(provider_account_id, payment_method_id, effective_from DESC)
  WHERE meta_deleted_at IS NULL AND effective_to IS NULL;

-- ---------------------------------------------------------------------------
-- 5. PSP HEALTH — Pre-computed health aggregates per scope
-- ---------------------------------------------------------------------------

CREATE TABLE payroute_psp_health (
  id                    uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  provider_account_id   uuid NOT NULL REFERENCES provider_accounts(id),
  method_code           citext NOT NULL,
  channel_code          citext NULL,
  currency              char(3) NOT NULL DEFAULT 'IDR',

  window_start          timestamptz NOT NULL,
  window_end            timestamptz NOT NULL,

  total_attempts        int NOT NULL DEFAULT 0,
  successful            int NOT NULL DEFAULT 0,
  failed_system         int NOT NULL DEFAULT 0,
  failed_decline        int NOT NULL DEFAULT 0,

  success_rate          decimal(7,4) NOT NULL DEFAULT 0,
  avg_latency_ms        int NULL,
  p95_latency_ms        int NULL,
  sample_size           int NOT NULL DEFAULT 0,

  meta_created_at       timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_psp_health_lookup
  ON payroute_psp_health(provider_account_id, method_code, window_start DESC);

CREATE INDEX idx_psp_health_scope
  ON payroute_psp_health(method_code, channel_code, window_start DESC);

-- ---------------------------------------------------------------------------
-- 6. AUDIT LOG — Immutable, append-only audit trail
-- ---------------------------------------------------------------------------

CREATE TABLE payroute_audit_log (
  id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  actor_user_id   uuid NOT NULL,
  actor_email     text NULL,
  actor_role      text NULL,

  action          text NOT NULL,
  target_entity   text NOT NULL,
  target_id       uuid NULL,

  before_value    jsonb NULL,
  after_value     jsonb NULL,
  reason          text NULL,

  ip_address      inet NULL,
  user_agent      text NULL,

  meta_created_at timestamptz NOT NULL DEFAULT now()
);

-- Partition-friendly indexes (append-only table)
CREATE INDEX idx_audit_log_created
  ON payroute_audit_log(meta_created_at DESC);

CREATE INDEX idx_audit_log_target
  ON payroute_audit_log(target_entity, target_id);

CREATE INDEX idx_audit_log_actor
  ON payroute_audit_log(actor_user_id, meta_created_at DESC);

-- ---------------------------------------------------------------------------
-- 7. PAYROUTE CONFIG — System-wide settings and kill-switch
-- ---------------------------------------------------------------------------

CREATE TABLE payroute_config (
  key             text PRIMARY KEY,
  value           jsonb NOT NULL,
  description     text NULL,
  updated_by      uuid NOT NULL,
  meta_updated_at timestamptz NOT NULL DEFAULT now()
);

-- Seed default config
INSERT INTO payroute_config (key, value, description, updated_by) VALUES
  ('circuit_breaker.default_threshold', '3', 'Default failure count to open circuit', '00000000-0000-0000-0000-000000000000'),
  ('circuit_breaker.default_cooldown_seconds', '30', 'Default cooldown before probe', '00000000-0000-0000-0000-000000000000'),
  ('circuit_breaker.half_open_probe_percent', '5', 'Percentage of traffic for half-open probe', '00000000-0000-0000-0000-000000000000'),
  ('health.window_seconds', '3600', 'Rolling window for health computation', '00000000-0000-0000-0000-000000000000'),
  ('health.min_sample_size', '30', 'Minimum attempts before health data is trusted', '00000000-0000-0000-0000-000000000000'),
  ('health.compute_interval_seconds', '60', 'How often health aggregates are recomputed', '00000000-0000-0000-0000-000000000000'),
  ('routing.cache_ttl_seconds', '60', 'TTL for in-memory routing rule cache', '00000000-0000-0000-0000-000000000000'),
  ('routing.max_attempts', '3', 'Default max provider attempts per transaction', '00000000-0000-0000-0000-000000000000'),
  ('routing.cost_aware.mdr_staleness_days', '30', 'MDR data older than this is flagged', '00000000-0000-0000-0000-000000000000'),
  ('kill_switch.psp_disabled', '[]', 'List of disabled PSP account IDs (JSON array)', '00000000-0000-0000-0000-000000000000');

-- ---------------------------------------------------------------------------
-- 8. RLS POLICIES — Row-Level Security for multi-tenant isolation
-- ---------------------------------------------------------------------------

-- Audit log is append-only (no UPDATE/DELETE)
ALTER TABLE payroute_audit_log ENABLE ROW LEVEL SECURITY;
CREATE POLICY audit_log_insert ON payroute_audit_log
  FOR INSERT WITH CHECK (true);
CREATE POLICY audit_log_select ON payroute_audit_log
  FOR SELECT USING (true);

-- Routing rules: platform-wide (no tenant isolation needed at DB level)
ALTER TABLE payroute_routing_rules ENABLE ROW LEVEL SECURITY;
CREATE POLICY routing_rules_all ON payroute_routing_rules
  FOR ALL USING (true) WITH CHECK (true);

-- ---------------------------------------------------------------------------
-- 9. TRIGGERS — Auto-update meta_updated_at
-- ---------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION payroute_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.meta_updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_routing_rules_updated
  BEFORE UPDATE ON payroute_routing_rules
  FOR EACH ROW EXECUTE FUNCTION payroute_update_timestamp();

CREATE TRIGGER trg_mdr_rates_updated
  BEFORE UPDATE ON payroute_mdr_rates
  FOR EACH ROW EXECUTE FUNCTION payroute_update_timestamp();

CREATE TRIGGER trg_config_updated
  BEFORE UPDATE ON payroute_config
  FOR EACH ROW EXECUTE FUNCTION payroute_update_timestamp();

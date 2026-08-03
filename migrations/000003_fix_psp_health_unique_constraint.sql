-- Fix: ON CONFLICT requires a UNIQUE constraint, not just a regular index.
-- The original migration created a regular index which causes:
--   pq: there is no unique or exclusion constraint matching the ON CONFLICT specification (42P10)

DROP INDEX IF EXISTS idx_psp_health_lookup;

CREATE UNIQUE INDEX idx_psp_health_lookup
  ON payroute_psp_health(provider_account_id, method_code, window_start DESC);

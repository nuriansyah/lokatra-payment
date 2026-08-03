-- Fix: Drop FK constraint on provider_webhook_events.provider_account_id
-- and make it nullable so webhooks can be received even when the provider
-- account UUID from config doesn't exist in the provider_accounts table.

ALTER TABLE provider_webhook_events
  DROP CONSTRAINT IF EXISTS provider_webhook_events_provider_account_id_fkey,
  ALTER COLUMN provider_account_id DROP NOT NULL;

DROP INDEX IF EXISTS idx_provider_webhook_events_unique_account_body;
CREATE UNIQUE INDEX idx_provider_webhook_events_unique_account_body
  ON provider_webhook_events(provider_account_id, raw_body_sha256)
  WHERE provider_account_id IS NOT NULL;

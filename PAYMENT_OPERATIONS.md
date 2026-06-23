# Payment Operations API

All paths are relative to `/v1`. Mutation requests accept at most 1 MiB of JSON
and reject unknown fields. Create endpoints accept `Idempotency-Key`; the JSON
`idempotencyKey` field takes precedence when both are supplied.

## Public and service routes

| Method | Path | Purpose |
| --- | --- | --- |
| POST | `/payment-intents/` | Create an idempotent payment intent |
| GET | `/payment-intents/{paymentIntentID}` | Read one intent |
| POST | `/payment-intents/{paymentIntentID}/confirm` | Confirm an intent |
| POST | `/payment-intents/{paymentIntentID}/cancel` | Cancel an intent |
| POST | `/refunds` | Request an idempotent refund |
| POST | `/webhooks/{provider}` | Verified Midtrans or Xendit callback |

## Admin routes

Admin routes require `X-Internal-Admin-Token`. The service fails closed with
`503` when `INTERNAL_PAYMENT_ADMIN_TOKEN` is not configured.

| Resource | Actions |
| --- | --- |
| `/admin/refunds/{id}/{action}` | `approve`, `reject`, `process`, `succeed`, `fail` |
| `/admin/webhooks/{id}/{action}` | `retry`, `ignore` |
| `/admin/manual-payment-evidence/{id}/{action}` | `review`, `approve`, `reject` |
| `/admin/overpayments/{id}/{action}` | `refund`, `credit_balance`, `apply_next_invoice`, `write_off` |
| `/admin/cash-sessions/{id}/{action}` | `close`, `cancel` |
| `/admin/payment-installments/{id}/{action}` | `pay`, `mark-overdue`, `cancel` |
| `/admin/payment-authorizations/{id}/{action}` | `authorize`, `capture`, `void`, `fail` |

Actions use an `actorId` in the body for audit attribution. Invalid or repeated
terminal transitions return `409 Conflict`; unknown actions return `400`.

## Provider configuration

Each provider is independently enabled with:

- `EXTERNALS_PROVIDERS_<PROVIDER>_ENABLED`

Only enabled providers are registered. Midtrans webhooks use SHA-512 signature
verification. Xendit webhooks use the configured callback token and optional
HMAC secret before the event enters the durable webhook inbox.

### Intelligent routing and fallback

Routing precedence is: exact `method + channel`, method wildcard, then the
default provider chain. Example configuration:

```env
INTERNAL_PAYMENT_ROUTING_DEFAULT_PROVIDERS=xendit,durianpay,midtrans
INTERNAL_PAYMENT_ROUTING_MAX_ATTEMPTS=3
INTERNAL_PAYMENT_ROUTING_FAILURE_THRESHOLD=3
INTERNAL_PAYMENT_ROUTING_COOLDOWN_SECONDS=30
INTERNAL_PAYMENT_ROUTING_RETRY_BACKOFF_MILLIS=100
INTERNAL_PAYMENT_ROUTING_RULES_JSON=[{"method":"qris","channel":"qris","providers":["xendit","durianpay"],"maxAttempts":3},{"method":"virtual_account","channel":"*","providers":["xendit","durianpay","midtrans"],"maxAttempts":3},{"method":"virtual_account","channel":"bca_va","providers":["midtrans","xendit"],"maxAttempts":2}]

EXTERNALS_PROVIDERS_DURIANPAY_ENABLED=true
EXTERNALS_PROVIDERS_DURIANPAY_ACCOUNT_ID=<provider-account-uuid>
EXTERNALS_PROVIDERS_DURIANPAY_BASE_URL=https://api.durianpay.id
EXTERNALS_PROVIDERS_DURIANPAY_API_KEY=<secret>
EXTERNALS_PROVIDERS_DURIANPAY_WEBHOOK_SECRET=<secret>
```

Only providers whose configured capabilities match method, channel, and
currency become candidates. Retryable failures use bounded exponential
backoff. After the configured consecutive-failure threshold, the provider's
circuit opens for that method/channel and traffic moves to the next candidate.
The circuit state and payment execution lock use Redis when configured, making
the behavior consistent across replicas; local development falls back to
concurrency-safe in-memory implementations.

Every routing decision, provider attempt, and returned payment instruction is
persisted. A repeated confirmation continues attempt numbering rather than
overwriting previous history.

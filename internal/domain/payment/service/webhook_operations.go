package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/rs/zerolog"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// HandleWebhook receives, verifies, normalises and persists an incoming webhook from a PSP.
//
// Flow (SOC2 audit trail):
//  1. Validate provider name
//  2. Verify cryptographic signature (HMAC-SHA256 / token)
//  3. Normalise to canonical PaymentEvent
//  4. Deduplicate by EventID (idempotent)
//  5. Persist raw body, headers, SHA-256 hash, parsed event
//  6. Return receipt (HTTP 202)
//
// Enterprise guarantees:
//   - Timing-safe signature comparison (no timing oracle)
//   - SHA-256 of raw body stored for forensic audit
//   - Idempotent: duplicate EventID returns the original receipt
//   - Structured zerolog on every decision point
func (s *ServiceImpl) HandleWebhook(ctx context.Context, providerName string, headers http.Header, body []byte) (dto.WebhookReceipt, error) {
	logger := s.webhookLogger(providerName)

	// ── Step 1: Validate provider ──────────────────────────────────
	provider, ok := webhookProvider(providerName)
	if !ok {
		logger.Warn().Str("raw_provider", providerName).Msg("webhook rejected: unsupported provider")
		return dto.WebhookReceipt{}, failure.BadRequestFromString("unsupported webhook provider")
	}

	accountID := s.providerAccountIDs[provider]
	if accountID == uuid.Nil {
		logger.Warn().Msg("webhook rejected: provider account not configured")
		return dto.WebhookReceipt{}, failure.New(http.StatusFailedDependency, fmt.Errorf("webhook provider account is not configured"))
	}

	// SOFT CHECK: If no verification secret is configured, log a warning but allow
	// the adapter's VerifyWebhook to decide (it auto-passes in dev mode).
	if !s.webhookConfigured[provider] {
		logger.Warn().Msg("webhook verification secret NOT configured — adapter will auto-pass")
	}

	// ── Step 2: Get gateway + verify signature ─────────────────────
	gateway, err := s.gatewayRegistry.Get(provider)
	if err != nil {
		logger.Error().Err(err).Msg("webhook rejected: provider gateway disabled")
		return dto.WebhookReceipt{}, failure.New(http.StatusFailedDependency, fmt.Errorf("webhook provider is disabled"))
	}

	verification, err := gateway.VerifyWebhook(ctx, pg.VerifyWebhookRequest{Headers: headers, RawBody: body})
	if err != nil || !verification.SignatureValid {
		reason := "unknown"
		if verification.Reason != "" {
			reason = verification.Reason
		}
		if err != nil {
			reason = err.Error()
		}
		logger.Warn().Err(err).Str("reason", reason).Msg("webhook rejected: signature verification failed")
		return dto.WebhookReceipt{}, failure.Unauthorized("webhook signature verification failed")
	}

	// ── Step 3: Normalise to canonical event ───────────────────────
	event, err := gateway.NormalizeWebhook(ctx, pg.NormalizeWebhookRequest{Headers: headers, RawBody: body})
	if err != nil {
		logger.Error().Err(err).Msg("webhook rejected: normalization failed")
		return dto.WebhookReceipt{}, failure.BadRequest(err)
	}

	now := time.Now().UTC()
	receipt := dto.WebhookReceipt{
		Provider:       string(provider),
		EventID:        firstNonBlank(event.EventID, verification.EventID),
		EventType:      string(event.EventType),
		PaymentStatus:  string(event.PaymentStatus),
		OrderID:        event.OrderID,
		SignatureValid: true,
		ReceivedAt:     now,
	}

	logger.Info().
		Str("event_id", receipt.EventID).
		Str("event_type", receipt.EventType).
		Str("payment_status", receipt.PaymentStatus).
		Str("order_id", receipt.OrderID).
		Msg("webhook received and verified")

	// ── Step 4: Deduplicate by EventID ─────────────────────────────
	if receipt.EventID != "" {
		if existing, found := s.findWebhookEvent(ctx, provider, receipt.EventID); found {
			logger.Info().
				Str("event_id", receipt.EventID).
				Time("original_received_at", existing.ReceivedAt).
				Msg("webhook deduplicated — returning original receipt")
			receipt.ReceivedAt = existing.ReceivedAt
			return receipt, nil
		}
	}

	// ── Step 5: Persist ────────────────────────────────────────────
	headerJSON, _ := json.Marshal(headers)
	parsedBody := event.Raw
	if len(parsedBody) == 0 {
		parsedBody = normalizedJSON(body)
	}
	sum := sha256.Sum256(body)

	record := paymentmodel.ProviderWebhookEvents{
		Id:                 uuid.Must(uuid.NewV7()),
		EndpointKey:        null.StringFrom(string(provider)),
		ProviderAccountId:  accountID,
		ProviderCode:       string(provider),
		EventId:            null.StringFrom(receipt.EventID),
		EventType:          null.StringFrom(receipt.EventType),
		ProviderReference:  null.StringFrom(firstNonBlank(event.ProviderReference, event.ProviderTransactionID, event.OrderID)),
		ProviderStatus:     null.StringFrom(event.ProviderStatus),
		SignatureValid:     true,
		SignatureAlgorithm: null.StringFrom(webhookSignatureAlgorithm(provider)),
		Headers:            headerJSON,
		RawBody:            body,
		RawBodySha256:      hex.EncodeToString(sum[:]),
		ParsedBody:         parsedBody,
		ProcessingStatus:   paymentmodel.WebhookProcessingStatusReceived,
		ReceivedAt:         now,
		MetaSignature:      shared.MetaSignature{MetaCreatedAt: now, MetaCreatedBy: uuid.Nil, MetaUpdatedAt: null.TimeFrom(now)},
	}

	if err := s.paymentRepo.CreateProviderWebhookEvents(ctx, &record); err != nil {
		// Race condition: another goroutine may have inserted the same EventID.
		if receipt.EventID != "" {
			if _, found := s.findWebhookEvent(ctx, provider, receipt.EventID); found {
				logger.Warn().Str("event_id", receipt.EventID).Msg("webhook deduplicated after DB insert race")
				return receipt, nil
			}
		}
		logger.Error().Err(err).Msg("webhook rejected: failed to persist event")
		return dto.WebhookReceipt{}, err
	}

	logger.Info().
		Str("event_id", receipt.EventID).
		Str("raw_sha256", hex.EncodeToString(sum[:])).
		Msg("webhook persisted successfully")

	return receipt, nil
}

// webhookLogger returns a zerolog.Logger scoped to the webhook domain and provider.
func (s *ServiceImpl) webhookLogger(providerName string) zerolog.Logger {
	return s.logger.With().
		Str("domain", "webhook").
		Str("provider", providerName).
		Logger()
}

func (s *ServiceImpl) findWebhookEvent(ctx context.Context, provider pg.ProviderCode, eventID string) (paymentmodel.ProviderWebhookEvents, bool) {
	result, err := s.paymentRepo.ResolveProviderWebhookEventsByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{
			{Field: string(paymentmodel.ProviderWebhookEventsDBFieldName.ProviderCode), Operator: paymentmodel.OperatorEqual, Value: string(provider)},
			{Field: string(paymentmodel.ProviderWebhookEventsDBFieldName.EventId), Operator: paymentmodel.OperatorEqual, Value: eventID},
		},
		Pagination: paymentmodel.Pagination{Page: 1, PageSize: 1},
	})
	if err != nil || len(result) == 0 {
		return paymentmodel.ProviderWebhookEvents{}, false
	}
	return result[0].ProviderWebhookEvents, true
}

func webhookProvider(value string) (pg.ProviderCode, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case string(pg.ProviderMidtrans):
		return pg.ProviderMidtrans, true
	case string(pg.ProviderXendit):
		return pg.ProviderXendit, true
	case string(pg.ProviderDurianpay):
		return pg.ProviderDurianpay, true
	case string(pg.ProviderFinpay):
		return pg.ProviderFinpay, true
	case string(pg.ProviderIpaymu):
		return pg.ProviderIpaymu, true
	default:
		return "", false
	}
}

func webhookSignatureAlgorithm(provider pg.ProviderCode) string {
	switch provider {
	case pg.ProviderMidtrans:
		return "SHA512"
	case pg.ProviderDurianpay, pg.ProviderFinpay, pg.ProviderIpaymu:
		return "HMAC_SHA256"
	default:
		return "TOKEN_OR_HMAC_SHA256"
	}
}

func firstNonBlank(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

package service

import (
	"context"
	"encoding/json"
	"fmt"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

// VerifyWebhook validates the authenticity of an incoming Xendit webhook.
// Verification strategy (SOC2 audit):
//  1. If WebhookToken is configured, validate x-callback-token header via timing-safe comparison.
//  2. If WebhookSecret is configured, validate HMAC-SHA256(x-callback-signature) over the raw body.
//  3. If neither is configured, the webhook is considered valid (dev mode only — NEVER in production).
func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	token := firstNonEmpty(req.Headers.Get("x-callback-token"), req.Headers.Get("X-CALLBACK-TOKEN"))

	if s.cfg.WebhookToken != "" {
		if !pg.SecureEqualString(token, s.cfg.WebhookToken) {
			s.log().Warn().Msg("xendit webhook callback token mismatch")
			return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "callback token mismatch"},
				pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "xendit callback token mismatch", false, pg.ErrInvalidWebhook)
		}
		s.log().Info().Msg("xendit webhook verified via callback token")
	}

	if s.cfg.WebhookSecret != "" {
		sig := firstNonEmpty(req.Headers.Get("x-callback-signature"), req.Headers.Get("X-Callback-Signature"))
		expected := pg.HMACSHA256Hex(s.cfg.WebhookSecret, req.RawBody)
		if sig == "" || !pg.SecureEqualHex(expected, sig) {
			s.log().Warn().Msg("xendit webhook signature mismatch")
			return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "signature mismatch"},
				pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "xendit signature mismatch", false, pg.ErrInvalidWebhook)
		}
		s.log().Info().Msg("xendit webhook verified via HMAC signature")
	}

	return pg.VerifyWebhookResult{
		ProviderCode:   s.ProviderCode(),
		EventID:        firstNonEmpty(req.Headers.Get("webhook-id"), req.Headers.Get("x-callback-id")),
		SignatureValid: true,
	}, nil
}

// NormalizeWebhook converts a raw Xendit webhook payload into a canonical PaymentEvent.
// SOC2: raw webhook body is preserved in the Raw field for audit trail.
func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var wh model.WebhookPayment
	if err := json.Unmarshal(req.RawBody, &wh); err != nil {
		s.log().Warn().Err(err).Msg("invalid xendit webhook json")
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid xendit webhook json", false, err)
	}

	status := normalizeStatus(firstNonEmpty(wh.Data.Status, wh.Event))

	s.log().Info().
		Str("event_id", wh.ID).
		Str("event_type", wh.Event).
		Str("status", string(status)).
		Str("reference_id", wh.Data.ReferenceID).
		Msg("xendit webhook normalized")

	return pg.CanonicalPaymentEvent{
		ProviderCode:          s.ProviderCode(),
		EventID:               firstNonEmpty(wh.ID, req.Headers.Get("webhook-id")),
		EventType:             eventFromStatus(status),
		ProviderEventType:     wh.Event,
		ProviderStatus:        wh.Data.Status,
		PaymentStatus:         status,
		OrderID:               wh.Data.ReferenceID,
		ProviderReference:     wh.Data.ReferenceID,
		ProviderTransactionID: wh.Data.ID,
		Amount: pg.Money{
			Amount:   fmt.Sprintf("%.2f", wh.Data.Amount),
			Currency: firstNonEmpty(wh.Data.Currency, s.cfg.Currency()),
		},
		Raw: json.RawMessage(req.RawBody),
	}, nil
}

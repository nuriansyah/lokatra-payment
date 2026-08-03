package service

import (
	"context"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

// RefundPayment creates a refund against an existing Xendit Payment Request.
// Idempotency is enforced by Xendit when the Idempotency-Key header is set.
func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		s.log().Warn().Err(err).Str("amount", req.Amount.Amount).Msg("invalid refund amount")
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}

	paymentID := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference)
	s.log().Info().
		Str("order_id", req.OrderID).
		Str("refund_id", req.RefundID).
		Str("payment_request_id", paymentID).
		Str("amount", req.Amount.Amount).
		Str("reason", req.Reason).
		Str("idempotency_key", req.IdempotencyKey).
		Msg("xendit refund request")

	payload := model.RefundRequest{
		ReferenceID:      req.RefundID,
		PaymentRequestID: paymentID,
		Currency:         firstNonEmpty(req.Amount.Currency, s.cfg.Currency()),
		Amount:           amount,
		Reason:           req.Reason,
	}

	var out model.RefundResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("refunds", "/refunds"))
	if err != nil {
		s.log().Error().Err(err).Str("refund_id", req.RefundID).Msg("xendit refund request failed")
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit refund request failed", true, err)
	}
	if resp.IsError() {
		s.log().Error().Int("status", resp.StatusCode()).Str("refund_id", req.RefundID).Msg("xendit refund returned error")
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().
		Str("refund_id", req.RefundID).
		Str("xendit_refund_id", out.ID).
		Str("status", out.Status).
		Msg("xendit refund created successfully")

	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: out.ID, OrderID: req.OrderID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

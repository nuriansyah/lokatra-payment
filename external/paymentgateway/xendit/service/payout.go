package service

import (
	"context"
	"strings"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

// CreatePayout initiates a disbursement via Xendit's POST /payouts endpoint.
// Idempotency is enforced by Xendit when the Idempotency-Key header is set.
func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		s.log().Warn().Err(err).Str("amount", req.Amount.Amount).Msg("invalid payout amount")
		return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid payout amount", false, err)
	}

	s.log().Info().
		Str("external_id", req.ExternalID).
		Str("bank_code", req.BankCode).
		Str("amount", req.Amount.Amount).
		Str("currency", req.Amount.Currency).
		Str("idempotency_key", req.IdempotencyKey).
		Msg("xendit payout request")

	payload := model.PayoutRequest{
		ReferenceID:       req.ExternalID,
		ChannelCode:       strings.ToUpper(req.BankCode),
		ChannelProperties: map[string]any{"account_number": req.AccountNumber, "account_holder_name": req.AccountName},
		Amount:            amount,
		Currency:          firstNonEmpty(req.Amount.Currency, s.cfg.Currency()),
		Description:       req.Description,
		Metadata:          req.Metadata,
	}

	var out model.PayoutResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("payouts", "/payouts"))
	if err != nil {
		s.log().Error().Err(err).Str("external_id", req.ExternalID).Msg("xendit payout request failed")
		return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payout request failed", true, err)
	}
	if resp.IsError() {
		s.log().Error().Int("status", resp.StatusCode()).Str("external_id", req.ExternalID).Msg("xendit payout returned error")
		return pg.CreatePayoutResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().
		Str("external_id", req.ExternalID).
		Str("xendit_payout_id", out.ID).
		Str("status", out.Status).
		Msg("xendit payout created successfully")

	return pg.CreatePayoutResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

// GetPayoutStatus retrieves the current status of a Payout from Xendit.
// Lookup order: ProviderPayoutID > PayoutID.
func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	id := firstNonEmpty(req.ProviderPayoutID, req.PayoutID)
	if id == "" {
		return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payout id is required", false, pg.ErrInvalidRequest)
	}

	s.log().Info().Str("id", id).Msg("xendit get payout status")

	var out model.PayoutResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payout_status", "/payouts/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		s.log().Error().Err(err).Str("id", id).Msg("xendit payout status request failed")
		return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payout status request failed", true, err)
	}
	if resp.IsError() {
		s.log().Error().Int("status", resp.StatusCode()).Str("id", id).Msg("xendit payout status returned error")
		return pg.GetPayoutStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().Str("id", id).Str("status", out.Status).Msg("xendit payout status retrieved")

	return pg.GetPayoutStatusResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

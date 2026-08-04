package service

import (
	"context"
	"time"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

// CreatePaymentLink generates a payment link (invoice) via Xendit's POST /v2/invoices endpoint.
func (s *ServiceImpl) CreatePaymentLink(ctx context.Context, req pg.CreatePaymentLinkRequest) (string, *time.Time, error) {
	s.log().Info().
		Str("invoice_id", req.InvoiceID).
		Str("amount", req.Amount).
		Str("currency", req.Currency).
		Msg("xendit create payment link request")

	amount, err := pg.AmountToFloat64(req.Amount)
	if err != nil || amount <= 0 {
		return "", nil, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}

	payload := model.CreateInvoiceRequest{
		ExternalID:     req.InvoiceID,
		Amount:         amount,
		Currency:       req.Currency,
		Description:    req.Description,
		CustomerID:     req.CustomerID,
		PaymentMethods: []string{"BANK_BCA", "BANK_MANDIRI", "BANK_BNI", "BANK_BRI", "QRIS"},
	}

	var out model.InvoiceResponse
	resp, err := s.auth(
		s.client.R().
			SetContext(ctx).
			SetBody(payload).
			SetResult(&out),
	).Post(s.cfg.Endpoint("invoices", "/v2/invoices"))
	if err != nil {
		s.log().Error().Err(err).Str("invoice_id", req.InvoiceID).Msg("xendit invoice request failed")
		return "", nil, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit invoice request failed", true, err)
	}

	if resp.IsError() {
		s.log().Error().
			Int("status", resp.StatusCode()).
			Str("body", pg.BodyString(resp)).
			Str("invoice_id", req.InvoiceID).
			Msg("xendit invoice returned error")
		return "", nil, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().
		Str("invoice_id", req.InvoiceID).
		Str("xendit_invoice_id", out.ID).
		Str("invoice_url", out.InvoiceURL).
		Msg("xendit invoice created successfully")

	var expiresAt *time.Time
	if out.ExpiryDate != "" {
		if t, err := time.Parse("2006-01-02", out.ExpiryDate); err == nil {
			expiresAt = &t
		}
	}

	return out.InvoiceURL, expiresAt, nil
}

// Compile-time check that ServiceImpl implements PaymentLinkCreator.
var _ pg.PaymentLinkCreator = (*ServiceImpl)(nil)

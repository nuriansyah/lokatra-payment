package service

import (
	"context"
	"fmt"
	"strings"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

// CreatePayment creates a Payment Request via Xendit's POST /payment_requests endpoint.
// Idempotency is enforced by Xendit when the Idempotency-Key header is set.
func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	s.log().Info().
		Str("order_id", req.OrderID).
		Str("method", string(req.Method)).
		Str("channel_code", req.ChannelCode).
		Str("amount", req.Amount.Amount).
		Str("currency", req.Amount.Currency).
		Str("idempotency_key", req.IdempotencyKey).
		Msg("xendit create payment request")

	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		s.log().Warn().Err(err).Str("amount", req.Amount.Amount).Msg("invalid payment amount")
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}

	payload := model.CreatePaymentRequest{
		ReferenceID:   req.OrderID,
		Currency:      firstNonEmpty(req.Amount.Currency, s.cfg.Currency()),
		Amount:        amount,
		Country:       "ID",
		Description:   req.Description,
		Metadata:      req.Metadata,
		PaymentMethod: buildPaymentMethod(req),
	}

	if req.Customer.Email != "" || req.Customer.Phone != "" || req.Customer.Name != "" {
		custRefID := firstNonEmpty(req.Customer.ExternalID, req.Customer.ID, req.OrderID)
		payload.Customer = &model.Customer{
			ReferenceID:    custRefID,
			Type:           "INDIVIDUAL",
			Email:          req.Customer.Email,
			MobileNumber:   req.Customer.Phone,
			IndividualDetail: &model.IndividualDetail{GivenNames: req.Customer.Name},
		}
	}

	var out model.PaymentRequestResponse
	resp, err := s.auth(
		s.client.R().
			SetContext(ctx).
			SetHeader("Idempotency-Key", req.IdempotencyKey).
			SetBody(payload).
			SetResult(&out),
	).Post(s.cfg.Endpoint("payment_requests", "/payment_requests"))
	if err != nil {
		s.log().Error().Err(err).Str("order_id", req.OrderID).Msg("xendit payment request failed")
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payment request failed", true, err)
	}

	if resp.IsError() {
		s.log().Error().
			Int("status", resp.StatusCode()).
			Str("body", pg.BodyString(resp)).
			Str("order_id", req.OrderID).
			Msg("xendit payment returned error")
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().
		Str("order_id", req.OrderID).
		Str("xendit_id", out.ID).
		Str("status", out.Status).
		Msg("xendit payment created successfully")

	return pg.CreatePaymentResponse{
		ProviderCode:          s.ProviderCode(),
		ProviderReference:     out.ReferenceID,
		ProviderTransactionID: out.ID,
		ProviderPaymentID:     out.ID,
		OrderID:               out.ReferenceID,
		Status:                normalizeStatus(out.Status),
		Instructions:          buildInstructions(out),
		Raw:                   pg.RawJSON(out),
	}, nil
}

// GetPaymentStatus retrieves the current status of a Payment Request from Xendit.
// Lookup order: ProviderTransactionID > ProviderReference > OrderID.
func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment request id/reference is required", false, pg.ErrInvalidRequest)
	}

	s.log().Info().Str("id", id).Msg("xendit get payment status")

	var out model.PaymentRequestResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payment_request_status", "/payment_requests/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		s.log().Error().Err(err).Str("id", id).Msg("xendit status request failed")
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit status request failed", true, err)
	}
	if resp.IsError() {
		s.log().Error().Int("status", resp.StatusCode()).Str("id", id).Msg("xendit status returned error")
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().
		Str("id", id).
		Str("status", out.Status).
		Float64("amount", out.Amount).
		Msg("xendit payment status retrieved")

	return pg.GetPaymentStatusResponse{
		ProviderCode:          s.ProviderCode(),
		ProviderReference:     out.ReferenceID,
		ProviderTransactionID: out.ID,
		OrderID:               out.ReferenceID,
		Status:                normalizeStatus(out.Status),
		Amount: pg.Money{
			Amount:   fmt.Sprintf("%.2f", out.Amount),
			Currency: firstNonEmpty(out.Currency, s.cfg.Currency()),
		},
		Raw: pg.RawJSON(out),
	}, nil
}

// CancelPayment cancels an existing Payment Request.
// Idempotency is enforced by Xendit when the Idempotency-Key header is set.
func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment request id/reference is required", false, pg.ErrInvalidRequest)
	}

	s.log().Info().
		Str("id", id).
		Str("reason", req.Reason).
		Str("idempotency_key", req.IdempotencyKey).
		Msg("xendit cancel payment request")

	var out model.PaymentRequestResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payment_request_cancel", "/payment_requests/{id}/cancel"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetResult(&out)).Post(path)
	if err != nil {
		s.log().Error().Err(err).Str("id", id).Msg("xendit cancel request failed")
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit cancel request failed", true, err)
	}
	if resp.IsError() {
		s.log().Error().Int("status", resp.StatusCode()).Str("id", id).Msg("xendit cancel returned error")
		return pg.CancelPaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}

	s.log().Info().Str("id", id).Str("status", out.Status).Msg("xendit payment canceled")

	return pg.CancelPaymentResponse{ProviderCode: s.ProviderCode(), OrderID: out.ReferenceID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

// buildPaymentMethod maps the canonical request into Xendit's PaymentMethodRequest format.
func buildPaymentMethod(req pg.CreatePaymentRequest) model.PaymentMethodRequest {
	channel := strings.ToUpper(strings.TrimSuffix(strings.TrimSuffix(req.ChannelCode, "_va"), "_virtual_account"))
	pm := model.PaymentMethodRequest{ReferenceID: req.OrderID, Reusability: "ONE_TIME_USE"}
	switch req.Method {
	case pg.PaymentMethodVirtualAccount:
		pm.Type = "VIRTUAL_ACCOUNT"
		pm.VirtualAccount = &model.VirtualAccount{ChannelCode: channel, ChannelProperties: map[string]any{"customer_name": req.Customer.Name}}
	case pg.PaymentMethodQRIS:
		pm.Type = "QR_CODE"
		pm.QR = &model.QR{ChannelCode: firstNonEmpty(channel, "QRIS")}
	case pg.PaymentMethodEWallet:
		pm.Type = "EWALLET"
		pm.EWallet = &model.EWallet{ChannelCode: channel}
	case pg.PaymentMethodRetailOutlet:
		pm.Type = "OVER_THE_COUNTER"
		pm.RetailOutlet = &model.RetailOutlet{ChannelCode: channel}
	default:
		pm.Type = strings.ToUpper(string(req.Method))
	}
	return pm
}

// buildInstructions extracts payment instructions (VA number, QR string, checkout URL, deeplink)
// from the Xendit Payment Request response.
func buildInstructions(out model.PaymentRequestResponse) []pg.PaymentInstruction {
	var res []pg.PaymentInstruction
	pm := out.PaymentMethod
	if pm.VirtualAccount != nil {
		exp := parseTime(pm.VirtualAccount.ChannelProperties.ExpiresAt)
		res = append(res, pg.PaymentInstruction{
			Type:          "va_number",
			DisplayName:   pm.VirtualAccount.ChannelCode + " Virtual Account",
			AccountNumber: pm.VirtualAccount.ChannelProperties.VirtualAccountNumber,
			ExpiresAt:     exp,
		})
	}
	if pm.QR != nil {
		exp := parseTime(pm.QR.ChannelProperties.ExpiresAt)
		res = append(res, pg.PaymentInstruction{
			Type:        "qr_string",
			DisplayName: "QRIS",
			QRString:    pm.QR.ChannelProperties.QRString,
			ExpiresAt:   exp,
		})
	}
	for _, a := range out.Actions {
		ins := pg.PaymentInstruction{
			Type:         "checkout_url",
			DisplayName:  a.Action,
			CheckoutURL:  a.URL,
			ProviderData: map[string]any{"method": a.Method, "url_type": a.URLType},
		}
		if strings.Contains(strings.ToLower(a.Action), "deeplink") {
			ins.Type = "deeplink"
			ins.DeeplinkURL = a.URL
		}
		res = append(res, ins)
	}
	return res
}

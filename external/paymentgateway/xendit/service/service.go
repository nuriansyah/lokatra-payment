package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
)

type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
}

func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{client: pg.NewRestyClient(cfg), cfg: cfg}
}
func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderXendit }

func (s *ServiceImpl) auth(req *resty.Request) *resty.Request {
	return req.SetHeader("Authorization", pg.BasicAuthValue(s.cfg.APIKey, ""))
}

func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: s.ProviderCode(), Items: []pg.Capability{
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bca_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "mandiri_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bni_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bri_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodEWallet, ChannelCode: "ovo", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodEWallet, ChannelCode: "dana", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
	}}, nil
}

func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}
	payload := model.CreatePaymentRequest{ReferenceID: req.OrderID, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), Amount: amount, Country: "ID", Description: req.Description, Metadata: req.Metadata, PaymentMethod: paymentMethod(req)}
	if req.Customer.Email != "" || req.Customer.Phone != "" || req.Customer.Name != "" {
		payload.Customer = &model.Customer{ReferenceID: firstNonEmpty(req.Customer.ExternalID, req.Customer.ID), Type: "INDIVIDUAL", Email: req.Customer.Email, MobileNumber: req.Customer.Phone, IndividualDetail: &model.IndividualDetail{GivenNames: req.Customer.Name}}
	}
	var out model.PaymentRequestResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("payment_requests", "/payment_requests"))
	if err != nil {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payment request failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePaymentResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.ReferenceID, ProviderTransactionID: out.ID, ProviderPaymentID: out.ID, OrderID: out.ReferenceID, Status: normalizeStatus(out.Status), Instructions: instructions(out), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment request id/reference is required", false, pg.ErrInvalidRequest)
	}
	var out model.PaymentRequestResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payment_request_status", "/payment_requests/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit status request failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPaymentStatusResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.ReferenceID, ProviderTransactionID: out.ID, OrderID: out.ReferenceID, Status: normalizeStatus(out.Status), Amount: pg.Money{Amount: fmt.Sprintf("%.2f", out.Amount), Currency: firstNonEmpty(out.Currency, s.cfg.Currency())}, Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment request id/reference is required", false, pg.ErrInvalidRequest)
	}
	var out model.PaymentRequestResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payment_request_cancel", "/payment_requests/{id}/cancel"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetResult(&out)).Post(path)
	if err != nil {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit cancel request failed", true, err)
	}
	if resp.IsError() {
		return pg.CancelPaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CancelPaymentResponse{ProviderCode: s.ProviderCode(), OrderID: out.ReferenceID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}
	payload := model.RefundRequest{ReferenceID: req.RefundID, PaymentRequestID: firstNonEmpty(req.ProviderTransactionID, req.ProviderReference), Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), Amount: amount, Reason: req.Reason}
	var out model.RefundResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("refunds", "/refunds"))
	if err != nil {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit refund request failed", true, err)
	}
	if resp.IsError() {
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: out.ID, OrderID: req.OrderID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid payout amount", false, err)
	}
	payload := model.PayoutRequest{ReferenceID: req.ExternalID, ChannelCode: strings.ToUpper(req.BankCode), ChannelProperties: map[string]any{"account_number": req.AccountNumber, "account_holder_name": req.AccountName}, Amount: amount, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), Description: req.Description, Metadata: req.Metadata}
	var out model.PayoutResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("payouts", "/payouts"))
	if err != nil {
		return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payout request failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePayoutResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePayoutResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	id := firstNonEmpty(req.ProviderPayoutID, req.PayoutID)
	if id == "" {
		return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payout id is required", false, pg.ErrInvalidRequest)
	}
	var out model.PayoutResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payout_status", "/payouts/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "xendit payout status request failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPayoutStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPayoutStatusResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	token := firstNonEmpty(req.Headers.Get("x-callback-token"), req.Headers.Get("X-CALLBACK-TOKEN"))
	if s.cfg.WebhookToken != "" && !pg.SecureEqualString(token, s.cfg.WebhookToken) {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "callback token mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "xendit callback token mismatch", false, pg.ErrInvalidWebhook)
	}
	// Some newer Xendit products also send webhook-id and/or HMAC headers. If configured, verify HMAC SHA256 over raw body.
	if s.cfg.WebhookSecret != "" {
		sig := firstNonEmpty(req.Headers.Get("x-callback-signature"), req.Headers.Get("X-Callback-Signature"))
		if sig == "" || !pg.SecureEqualHex(pg.HMACSHA256Hex(s.cfg.WebhookSecret, req.RawBody), sig) {
			return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "signature mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "xendit signature mismatch", false, pg.ErrInvalidWebhook)
		}
	}
	return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: firstNonEmpty(req.Headers.Get("webhook-id"), req.Headers.Get("x-callback-id")), SignatureValid: true}, nil
}

func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var wh model.WebhookPayment
	if err := json.Unmarshal(req.RawBody, &wh); err != nil {
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid xendit webhook json", false, err)
	}
	status := normalizeStatus(firstNonEmpty(wh.Data.Status, wh.Event))
	return pg.CanonicalPaymentEvent{ProviderCode: s.ProviderCode(), EventID: firstNonEmpty(wh.ID, req.Headers.Get("webhook-id")), EventType: eventFromStatus(status), ProviderEventType: wh.Event, ProviderStatus: wh.Data.Status, PaymentStatus: status, OrderID: wh.Data.ReferenceID, ProviderReference: wh.Data.ReferenceID, ProviderTransactionID: wh.Data.ID, Amount: pg.Money{Amount: fmt.Sprintf("%.2f", wh.Data.Amount), Currency: firstNonEmpty(wh.Data.Currency, s.cfg.Currency())}, Raw: json.RawMessage(req.RawBody)}, nil
}

func paymentMethod(req pg.CreatePaymentRequest) model.PaymentMethodRequest {
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

func instructions(out model.PaymentRequestResponse) []pg.PaymentInstruction {
	var res []pg.PaymentInstruction
	pm := out.PaymentMethod
	if pm.VirtualAccount != nil {
		exp := parseTime(pm.VirtualAccount.ChannelProperties.ExpiresAt)
		res = append(res, pg.PaymentInstruction{Type: "va_number", DisplayName: pm.VirtualAccount.ChannelCode + " Virtual Account", AccountNumber: pm.VirtualAccount.ChannelProperties.VirtualAccountNumber, ExpiresAt: exp})
	}
	if pm.QR != nil {
		exp := parseTime(pm.QR.ChannelProperties.ExpiresAt)
		res = append(res, pg.PaymentInstruction{Type: "qr_string", DisplayName: "QRIS", QRString: pm.QR.ChannelProperties.QRString, ExpiresAt: exp})
	}
	for _, a := range out.Actions {
		ins := pg.PaymentInstruction{Type: "checkout_url", DisplayName: a.Action, CheckoutURL: a.URL, ProviderData: map[string]any{"method": a.Method, "url_type": a.URLType}}
		if strings.Contains(strings.ToLower(a.Action), "deeplink") {
			ins.Type = "deeplink"
			ins.DeeplinkURL = a.URL
		}
		res = append(res, ins)
	}
	return res
}

func normalizeStatus(s string) pg.PaymentStatus {
	switch strings.ToUpper(s) {
	case "SUCCEEDED", "SUCCESS", "COMPLETED", "PAID", "CAPTURED":
		return pg.PaymentStatusSucceeded
	case "PENDING", "REQUIRES_ACTION", "AWAITING_CAPTURE":
		return pg.PaymentStatusPending
	case "FAILED", "VOIDED":
		return pg.PaymentStatusFailed
	case "EXPIRED":
		return pg.PaymentStatusExpired
	case "CANCELED", "CANCELLED":
		return pg.PaymentStatusCanceled
	case "REFUNDED":
		return pg.PaymentStatusRefunded
	case "PARTIALLY_REFUNDED":
		return pg.PaymentStatusPartiallyRefunded
	}
	return pg.PaymentStatusUnknown
}
func eventFromStatus(s pg.PaymentStatus) pg.EventType {
	switch s {
	case pg.PaymentStatusSucceeded:
		return pg.EventPaymentSucceeded
	case pg.PaymentStatusPending:
		return pg.EventPaymentPending
	case pg.PaymentStatusFailed:
		return pg.EventPaymentFailed
	case pg.PaymentStatusExpired:
		return pg.EventPaymentExpired
	case pg.PaymentStatusCanceled:
		return pg.EventPaymentCanceled
	case pg.PaymentStatusRefunded:
		return pg.EventPaymentRefunded
	case pg.PaymentStatusPartiallyRefunded:
		return pg.EventPaymentPartiallyRefunded
	}
	return pg.EventUnknown
}
func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}
func parseTime(s string) *time.Time {
	if s == "" {
		return nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return &t
	}
	return nil
}

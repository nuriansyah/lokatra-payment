package service

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/durianpay/model"
)

type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
}

func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{client: pg.NewRestyClient(cfg), cfg: cfg}
}
func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderDurianpay }
func (s *ServiceImpl) auth(r *resty.Request) *resty.Request {
	return r.SetHeader("Authorization", pg.BasicAuthValue(s.cfg.APIKey, ""))
}
func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: s.ProviderCode(), Items: []pg.Capability{{Method: pg.PaymentMethodPaymentPage, ChannelCode: "checkout", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsPayout: false, SupportsExpiry: true}, {Method: pg.PaymentMethodVirtualAccount, ChannelCode: "va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}, {Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}}}, nil
}
func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	payload := model.CreatePaymentRequest{Amount: req.Amount.Amount, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), OrderRefID: req.OrderID, Customer: &model.Customer{ID: firstNonEmpty(req.Customer.ExternalID, req.Customer.ID), Name: req.Customer.Name, Email: req.Customer.Email, Phone: req.Customer.Phone}, PaymentMethod: string(req.Method), PaymentChannel: req.ChannelCode, Description: req.Description, CallbackURL: req.CallbackURL, ReturnURL: req.ReturnURL, Metadata: req.Metadata}
	if req.ExpiryAt != nil {
		payload.ExpiryAt = req.ExpiryAt.Format(time.RFC3339)
	}
	var out model.PaymentResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("create_payment", "/v1/payments"))
	if err != nil {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay create payment failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePaymentResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.OrderRefID, ProviderTransactionID: out.ID, ProviderPaymentID: out.ID, OrderID: out.OrderRefID, Status: normalizeStatus(out.Status), Instructions: instructions(out), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment id/reference required", false, pg.ErrInvalidRequest)
	}
	var out model.PaymentResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("status", "/v1/payments/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay status failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPaymentStatusResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.OrderRefID, ProviderTransactionID: out.ID, OrderID: out.OrderRefID, Status: normalizeStatus(out.Status), Amount: pg.Money{Amount: out.Amount, Currency: firstNonEmpty(out.Currency, s.cfg.Currency())}, Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	id := firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID)
	if id == "" {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "payment id/reference required", false, pg.ErrInvalidRequest)
	}
	var out model.PaymentResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("cancel", "/v1/payments/{id}/cancel"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetResult(&out)).Post(path)
	if err != nil {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay cancel failed", true, err)
	}
	if resp.IsError() {
		return pg.CancelPaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CancelPaymentResponse{ProviderCode: s.ProviderCode(), OrderID: out.OrderRefID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	payload := model.RefundRequest{PaymentID: firstNonEmpty(req.ProviderTransactionID, req.ProviderReference, req.OrderID), RefundRefID: req.RefundID, Amount: req.Amount.Amount, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), Reason: req.Reason}
	var out model.RefundResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("refund", "/v1/refunds"))
	if err != nil {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay refund failed", true, err)
	}
	if resp.IsError() {
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: out.ID, OrderID: req.OrderID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	payload := model.PayoutRequest{ReferenceID: req.ExternalID, Amount: req.Amount.Amount, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), BankCode: req.BankCode, AccountNumber: req.AccountNumber, AccountName: req.AccountName, Description: req.Description, Metadata: req.Metadata}
	var out model.PayoutResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("payout", "/v1/payouts"))
	if err != nil {
		return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay payout failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePayoutResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePayoutResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	id := firstNonEmpty(req.ProviderPayoutID, req.PayoutID)
	var out model.PayoutResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("payout_status", "/v1/payouts/{id}"), "{id}", id)
	resp, err := s.auth(s.client.R().SetContext(ctx).SetResult(&out)).Get(path)
	if err != nil {
		return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "durianpay payout status failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPayoutStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPayoutStatusResponse{ProviderCode: s.ProviderCode(), ProviderPayoutID: out.ID, Status: normalizeStatus(out.Status), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	if s.cfg.WebhookSecret == "" {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: true, Reason: "webhook secret not configured"}, nil
	}
	sig := firstNonEmpty(req.Headers.Get("X-Signature"), req.Headers.Get("x-signature"))
	if !pg.SecureEqualHex(pg.HMACSHA256Hex(s.cfg.WebhookSecret, req.RawBody), sig) {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "signature mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "durianpay signature mismatch", false, pg.ErrInvalidWebhook)
	}
	return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: req.Headers.Get("X-Event-Id"), SignatureValid: true}, nil
}
func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var wh model.Webhook
	if err := json.Unmarshal(req.RawBody, &wh); err != nil {
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid durianpay webhook", false, err)
	}
	st := normalizeStatus(wh.Data.Status)
	return pg.CanonicalPaymentEvent{ProviderCode: s.ProviderCode(), EventID: wh.ID, EventType: eventFromStatus(st), ProviderEventType: wh.Event, ProviderStatus: wh.Data.Status, PaymentStatus: st, OrderID: wh.Data.OrderRefID, ProviderReference: wh.Data.OrderRefID, ProviderTransactionID: wh.Data.ID, Amount: pg.Money{Amount: wh.Data.Amount, Currency: firstNonEmpty(wh.Data.Currency, s.cfg.Currency())}, Raw: json.RawMessage(req.RawBody)}, nil
}
func instructions(out model.PaymentResponse) []pg.PaymentInstruction {
	var res []pg.PaymentInstruction
	if out.VA != "" {
		res = append(res, pg.PaymentInstruction{Type: "va_number", DisplayName: out.PaymentChannel, AccountNumber: out.VA})
	}
	if out.PaymentCode != "" {
		res = append(res, pg.PaymentInstruction{Type: "payment_code", DisplayName: out.PaymentChannel, PaymentCode: out.PaymentCode})
	}
	if out.QRString != "" {
		res = append(res, pg.PaymentInstruction{Type: "qr_string", DisplayName: "QRIS", QRString: out.QRString})
	}
	if firstNonEmpty(out.PaymentURL, out.CheckoutURL) != "" {
		res = append(res, pg.PaymentInstruction{Type: "checkout_url", DisplayName: "Checkout", CheckoutURL: firstNonEmpty(out.PaymentURL, out.CheckoutURL)})
	}
	return res
}
func normalizeStatus(s string) pg.PaymentStatus {
	switch strings.ToLower(s) {
	case "paid", "success", "succeeded", "completed":
		return pg.PaymentStatusSucceeded
	case "pending", "processing", "created":
		return pg.PaymentStatusPending
	case "failed", "failure":
		return pg.PaymentStatusFailed
	case "expired":
		return pg.PaymentStatusExpired
	case "canceled", "cancelled":
		return pg.PaymentStatusCanceled
	case "refunded":
		return pg.PaymentStatusRefunded
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

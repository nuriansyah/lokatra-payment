package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/finpay/model"
)

type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
}

func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{client: pg.NewRestyClient(cfg), cfg: cfg}
}
func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderFinpay }
func (s *ServiceImpl) auth(r *resty.Request) *resty.Request {
	return r.SetHeader("Authorization", pg.BasicAuthValue(s.cfg.MerchantID, s.cfg.MerchantKey)).SetHeader("X-Merchant-Id", s.cfg.MerchantID)
}
func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: s.ProviderCode(), Items: []pg.Capability{{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bca_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}, {Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}, {Method: pg.PaymentMethodEWallet, ChannelCode: "ewallet", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}}}, nil
}
func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	payload := model.PaymentRequest{MerchantID: s.cfg.MerchantID, MerchantTransID: req.OrderID, Amount: req.Amount.Amount, Currency: firstNonEmpty(req.Amount.Currency, s.cfg.Currency()), PaymentChannel: req.ChannelCode, Description: req.Description, CustomerName: req.Customer.Name, CustomerEmail: req.Customer.Email, CustomerPhone: req.Customer.Phone, CallbackURL: req.CallbackURL, ReturnURL: req.ReturnURL, Metadata: req.Metadata}
	if req.ExpiryAt != nil {
		payload.ExpiryTime = req.ExpiryAt.Format(time.RFC3339)
	}
	var out model.PaymentResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("create_payment", "/payments"))
	if err != nil {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "finpay create payment failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePaymentResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.MerchantTransID, ProviderTransactionID: out.TransactionID, ProviderPaymentID: out.TransactionID, OrderID: out.MerchantTransID, Status: normalizeStatus(out.Status, out.ResponseCode), Instructions: instructions(out), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	payload := model.StatusRequest{MerchantID: s.cfg.MerchantID, MerchantTransID: firstNonEmpty(req.OrderID, req.ProviderReference), TransactionID: req.ProviderTransactionID}
	var out model.PaymentResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("status", "/payments/status"))
	if err != nil {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "finpay status failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPaymentStatusResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.MerchantTransID, ProviderTransactionID: out.TransactionID, OrderID: out.MerchantTransID, Status: normalizeStatus(out.Status, out.ResponseCode), Amount: pg.Money{Amount: out.Amount, Currency: firstNonEmpty(out.Currency, s.cfg.Currency())}, Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "finpay cancel is provider/product-specific; configure endpoint before enabling", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	payload := model.RefundRequest{MerchantID: s.cfg.MerchantID, MerchantTransID: firstNonEmpty(req.OrderID, req.ProviderReference), RefundID: req.RefundID, Amount: req.Amount.Amount, Reason: req.Reason}
	var out model.RefundResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out)).Post(s.cfg.Endpoint("refund", "/payments/refund"))
	if err != nil {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "finpay refund failed", true, err)
	}
	if resp.IsError() {
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: req.RefundID, OrderID: out.MerchantTransID, Status: normalizeStatus(out.Status, out.ResponseCode), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "finpay payout is not implemented in this adapter", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "finpay payout status is not implemented in this adapter", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	var n model.Notification
	if err := json.Unmarshal(req.RawBody, &n); err != nil {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "invalid json"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid finpay webhook", false, err)
	}
	if s.cfg.WebhookSecret != "" {
		expected := pg.HMACSHA256Hex(s.cfg.WebhookSecret, []byte(n.MerchantTransID+n.TransactionID+n.Amount))
		if n.Signature != "" && !pg.SecureEqualHex(expected, n.Signature) {
			return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, SignatureValid: false, Reason: "signature mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "finpay signature mismatch", false, pg.ErrInvalidWebhook)
		}
	}
	return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, SignatureValid: true}, nil
}
func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var n model.Notification
	if err := json.Unmarshal(req.RawBody, &n); err != nil {
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid finpay webhook", false, err)
	}
	st := normalizeStatus(n.Status, n.ResponseCode)
	return pg.CanonicalPaymentEvent{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, EventType: eventFromStatus(st), ProviderEventType: n.PaymentChannel, ProviderStatus: n.Status, PaymentStatus: st, OrderID: n.MerchantTransID, ProviderReference: n.MerchantTransID, ProviderTransactionID: n.TransactionID, Amount: pg.Money{Amount: n.Amount, Currency: firstNonEmpty(n.Currency, s.cfg.Currency())}, Raw: json.RawMessage(req.RawBody)}, nil
}
func instructions(out model.PaymentResponse) []pg.PaymentInstruction {
	var res []pg.PaymentInstruction
	if out.PaymentCode != "" {
		res = append(res, pg.PaymentInstruction{Type: "payment_code", DisplayName: out.PaymentChannel, PaymentCode: out.PaymentCode})
	}
	if out.PaymentURL != "" {
		res = append(res, pg.PaymentInstruction{Type: "checkout_url", DisplayName: out.PaymentChannel, CheckoutURL: out.PaymentURL})
	}
	if out.QRString != "" {
		res = append(res, pg.PaymentInstruction{Type: "qr_string", DisplayName: "QRIS", QRString: out.QRString})
	}
	return res
}
func normalizeStatus(status, code string) pg.PaymentStatus {
	s := strings.ToLower(status)
	c := strings.ToUpper(code)
	if c == "00" || s == "paid" || s == "success" || s == "settlement" || s == "succeeded" {
		return pg.PaymentStatusSucceeded
	}
	if s == "pending" || s == "process" || s == "processing" {
		return pg.PaymentStatusPending
	}
	if s == "expired" {
		return pg.PaymentStatusExpired
	}
	if s == "cancel" || s == "canceled" || s == "cancelled" {
		return pg.PaymentStatusCanceled
	}
	if s == "refund" || s == "refunded" {
		return pg.PaymentStatusRefunded
	}
	if c != "" && c != "00" {
		return pg.PaymentStatusFailed
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

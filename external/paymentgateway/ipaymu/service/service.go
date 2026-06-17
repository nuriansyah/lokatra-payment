package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/ipaymu/model"
)

type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
}

func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{client: pg.NewRestyClient(cfg), cfg: cfg}
}
func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderIpaymu }
func (s *ServiceImpl) auth(r *resty.Request, body []byte) *resty.Request {
	sig := pg.HMACSHA256Hex(s.cfg.APIKey, body)
	return r.SetHeader("va", s.cfg.MerchantID).SetHeader("signature", sig).SetHeader("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
}
func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: s.ProviderCode(), Items: []pg.Capability{{Method: pg.PaymentMethodPaymentPage, ChannelCode: "redirect", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}, {Method: pg.PaymentMethodVirtualAccount, ChannelCode: "va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}, {Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true}}}, nil
}
func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	products := []string{req.Description}
	qty := []int64{1}
	prices := []string{req.Amount.Amount}
	if len(req.Items) > 0 {
		products = nil
		qty = nil
		prices = nil
		for _, it := range req.Items {
			products = append(products, it.Name)
			qty = append(qty, it.Quantity)
			prices = append(prices, it.Price)
		}
	}
	payload := model.RedirectPaymentRequest{Product: products, Qty: qty, Price: prices, ReturnURL: req.ReturnURL, CancelURL: req.FailureURL, NotifyURL: req.CallbackURL, ReferenceID: req.OrderID, BuyerName: req.Customer.Name, BuyerEmail: req.Customer.Email, BuyerPhone: req.Customer.Phone, PaymentMethod: string(req.Method), PaymentChannel: req.ChannelCode}
	if req.ExpiryAt != nil {
		payload.Expired = req.ExpiryAt.Format("2006-01-02 15:04:05")
	}
	body := pg.RawJSON(payload)
	var out model.PaymentResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out), body).Post(s.cfg.Endpoint("redirect_payment", "/api/v2/payment"))
	if err != nil {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "ipaymu create payment failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePaymentResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.Data.ReferenceID, ProviderTransactionID: firstNonEmpty(out.Data.TransactionID, out.Data.SessionID), ProviderPaymentID: out.Data.SessionID, OrderID: out.Data.ReferenceID, Status: normalizeStatus(firstNonEmpty(out.Data.Status, fmt.Sprint(out.Status))), Instructions: instructions(out.Data), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	payload := model.StatusRequest{TransactionID: req.ProviderTransactionID, ReferenceID: firstNonEmpty(req.OrderID, req.ProviderReference), Account: s.cfg.MerchantID}
	body := pg.RawJSON(payload)
	var out model.PaymentResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetBody(payload).SetResult(&out), body).Post(s.cfg.Endpoint("status", "/api/v2/transaction"))
	if err != nil {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "ipaymu status failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPaymentStatusResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.Data.ReferenceID, ProviderTransactionID: firstNonEmpty(out.Data.TransactionID, out.Data.SessionID), OrderID: out.Data.ReferenceID, Status: normalizeStatus(firstNonEmpty(out.Data.Status, fmt.Sprint(out.Status))), Amount: pg.Money{Amount: out.Data.Total, Currency: s.cfg.Currency()}, Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "ipaymu cancel is not implemented in this adapter", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	payload := model.RefundRequest{TransactionID: firstNonEmpty(req.ProviderTransactionID, req.ProviderReference), ReferenceID: req.RefundID, Amount: req.Amount.Amount, Reason: req.Reason}
	body := pg.RawJSON(payload)
	var out model.RefundResponse
	resp, err := s.auth(s.client.R().SetContext(ctx).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out), body).Post(s.cfg.Endpoint("refund", "/api/v2/refund"))
	if err != nil {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "ipaymu refund failed", true, err)
	}
	if resp.IsError() {
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: firstNonEmpty(out.Data.TransactionID, out.Data.SessionID), OrderID: req.OrderID, Status: normalizeStatus(firstNonEmpty(out.Data.Status, fmt.Sprint(out.Status))), Raw: pg.RawJSON(out)}, nil
}
func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "ipaymu payout is provider-contract specific; not enabled by default", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "ipaymu payout status is not implemented", false, pg.ErrUnsupportedOperation)
}
func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	var wh model.Webhook
	if err := json.Unmarshal(req.RawBody, &wh); err != nil {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "invalid json"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid ipaymu webhook", false, err)
	}
	if s.cfg.WebhookSecret != "" {
		expected := pg.HMACSHA256Hex(s.cfg.WebhookSecret, []byte(wh.TrxID+wh.ReferenceID+wh.Amount+wh.Status))
		if wh.Signature != "" && !pg.SecureEqualHex(expected, wh.Signature) {
			return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: wh.TrxID, SignatureValid: false, Reason: "signature mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "ipaymu signature mismatch", false, pg.ErrInvalidWebhook)
		}
	}
	return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: wh.TrxID, SignatureValid: true}, nil
}
func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var wh model.Webhook
	if err := json.Unmarshal(req.RawBody, &wh); err != nil {
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid ipaymu webhook", false, err)
	}
	st := normalizeStatus(firstNonEmpty(wh.Status, wh.StatusCode))
	return pg.CanonicalPaymentEvent{ProviderCode: s.ProviderCode(), EventID: wh.TrxID, EventType: eventFromStatus(st), ProviderEventType: wh.PaymentMethod, ProviderStatus: wh.Status, PaymentStatus: st, OrderID: wh.ReferenceID, ProviderReference: wh.ReferenceID, ProviderTransactionID: wh.TrxID, Amount: pg.Money{Amount: wh.Amount, Currency: s.cfg.Currency()}, Raw: json.RawMessage(req.RawBody)}, nil
}
func instructions(d model.PaymentData) []pg.PaymentInstruction {
	var res []pg.PaymentInstruction
	if d.URL != "" {
		res = append(res, pg.PaymentInstruction{Type: "checkout_url", DisplayName: "iPaymu Checkout", CheckoutURL: d.URL})
	}
	if d.PaymentNo != "" {
		res = append(res, pg.PaymentInstruction{Type: "payment_code", DisplayName: "Payment Code", PaymentCode: d.PaymentNo})
	}
	if d.QRString != "" {
		res = append(res, pg.PaymentInstruction{Type: "qr_string", DisplayName: "QRIS", QRString: d.QRString})
	}
	return res
}
func normalizeStatus(s string) pg.PaymentStatus {
	switch strings.ToLower(s) {
	case "1", "paid", "success", "succeeded", "berhasil":
		return pg.PaymentStatusSucceeded
	case "0", "pending", "process", "processing":
		return pg.PaymentStatusPending
	case "-1", "failed", "failure", "gagal":
		return pg.PaymentStatusFailed
	case "expired":
		return pg.PaymentStatusExpired
	case "cancel", "canceled", "cancelled":
		return pg.PaymentStatusCanceled
	case "refund", "refunded":
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

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
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/midtrans/model"
)

// ServiceImpl implements Midtrans Core API via Resty.
type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
}

func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{client: pg.NewRestyClient(cfg), cfg: cfg}
}

func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderMidtrans }

func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: s.ProviderCode(), Items: []pg.Capability{
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bca_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bni_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bri_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "permata_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		{Method: pg.PaymentMethodEWallet, ChannelCode: "gopay", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
	}}, nil
}

func (s *ServiceImpl) CreatePayment(ctx context.Context, req pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if err != nil || amount <= 0 {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "invalid amount", false, err)
	}
	payload := model.ChargeRequest{
		PaymentType:       s.paymentType(req.Method, req.ChannelCode),
		TransactionDetail: model.TransactionDetail{OrderID: req.OrderID, GrossAmount: amount},
		CustomerDetails:   toCustomer(req.Customer),
		ItemDetails:       toItems(req.Items),
		CustomField1:      req.PaymentIntentID,
		CustomField2:      req.AttemptID,
	}
	if payload.PaymentType == "bank_transfer" {
		payload.BankTransfer = &model.BankTransfer{Bank: bankCode(req.ChannelCode)}
	}
	if payload.PaymentType == "gopay" {
		payload.Gopay = &model.Gopay{EnableCallback: req.ReturnURL != "", CallbackURL: req.ReturnURL}
	}
	if payload.PaymentType == "qris" {
		payload.Qris = &model.Qris{Acquirer: s.cfg.Extra["qris_acquirer"]}
	}
	if req.ExpiryAt != nil {
		dur := int64(time.Until(*req.ExpiryAt).Minutes())
		if dur < 1 {
			dur = 1
		}
		payload.Expiry = &model.Expiry{OrderTime: time.Now().Format("2006-01-02 15:04:05 -0700"), ExpiryDuration: dur, Unit: "minute"}
	}

	var out model.ChargeResponse
	resp, err := s.client.R().
		SetContext(ctx).
		SetHeader("Authorization", pg.BasicAuthValue(s.cfg.ServerKey, "")).
		SetHeader("Idempotency-Key", req.IdempotencyKey).
		SetBody(payload).
		SetResult(&out).
		Post(s.cfg.Endpoint("charge", "/v2/charge"))
	if err != nil {
		return pg.CreatePaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "midtrans charge request failed", true, err)
	}
	if resp.IsError() {
		return pg.CreatePaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CreatePaymentResponse{
		ProviderCode:          s.ProviderCode(),
		ProviderReference:     out.OrderID,
		ProviderTransactionID: out.TransactionID,
		ProviderPaymentID:     out.TransactionID,
		OrderID:               out.OrderID,
		Status:                normalizeStatus(out.TransactionStatus, out.FraudStatus),
		Instructions:          instructions(out),
		Raw:                   pg.RawJSON(out),
	}, nil
}

func (s *ServiceImpl) GetPaymentStatus(ctx context.Context, req pg.GetPaymentStatusRequest) (pg.GetPaymentStatusResponse, error) {
	orderID := firstNonEmpty(req.OrderID, req.ProviderReference)
	if orderID == "" {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "order id is required", false, pg.ErrInvalidRequest)
	}
	var out model.StatusResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("status", "/v2/{order_id}/status"), "{order_id}", orderID)
	resp, err := s.client.R().SetContext(ctx).SetHeader("Authorization", pg.BasicAuthValue(s.cfg.ServerKey, "")).SetResult(&out).Get(path)
	if err != nil {
		return pg.GetPaymentStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "midtrans status request failed", true, err)
	}
	if resp.IsError() {
		return pg.GetPaymentStatusResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.GetPaymentStatusResponse{ProviderCode: s.ProviderCode(), ProviderReference: out.OrderID, ProviderTransactionID: out.TransactionID, OrderID: out.OrderID, Status: normalizeStatus(out.TransactionStatus, out.FraudStatus), Amount: pg.Money{Amount: out.GrossAmount, Currency: firstNonEmpty(out.Currency, s.cfg.Currency())}, Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) CancelPayment(ctx context.Context, req pg.CancelPaymentRequest) (pg.CancelPaymentResponse, error) {
	orderID := firstNonEmpty(req.OrderID, req.ProviderReference)
	if orderID == "" {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "order id is required", false, pg.ErrInvalidRequest)
	}
	var out model.StatusResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("cancel", "/v2/{order_id}/cancel"), "{order_id}", orderID)
	resp, err := s.client.R().SetContext(ctx).SetHeader("Authorization", pg.BasicAuthValue(s.cfg.ServerKey, "")).SetHeader("Idempotency-Key", req.IdempotencyKey).SetResult(&out).Post(path)
	if err != nil {
		return pg.CancelPaymentResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "midtrans cancel request failed", true, err)
	}
	if resp.IsError() {
		return pg.CancelPaymentResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.CancelPaymentResponse{ProviderCode: s.ProviderCode(), OrderID: out.OrderID, Status: normalizeStatus(out.TransactionStatus, out.FraudStatus), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) RefundPayment(ctx context.Context, req pg.RefundRequest) (pg.RefundResponse, error) {
	orderID := firstNonEmpty(req.OrderID, req.ProviderReference)
	amount, err := pg.AmountToFloat64(req.Amount.Amount)
	if orderID == "" || err != nil || amount <= 0 {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeInvalidRequest, 0, "order id and valid amount are required", false, err)
	}
	payload := model.RefundRequest{RefundKey: req.RefundID, Amount: amount, Reason: req.Reason}
	var out model.RefundResponse
	path := strings.ReplaceAll(s.cfg.Endpoint("refund", "/v2/{order_id}/refund"), "{order_id}", orderID)
	resp, err := s.client.R().SetContext(ctx).SetHeader("Authorization", pg.BasicAuthValue(s.cfg.ServerKey, "")).SetHeader("Idempotency-Key", req.IdempotencyKey).SetBody(payload).SetResult(&out).Post(path)
	if err != nil {
		return pg.RefundResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeProviderTimeout, 0, "midtrans refund request failed", true, err)
	}
	if resp.IsError() {
		return pg.RefundResponse{}, pg.ErrorFromHTTP(s.ProviderCode(), resp.StatusCode(), pg.BodyString(resp))
	}
	return pg.RefundResponse{ProviderCode: s.ProviderCode(), ProviderRefundID: out.RefundKey, OrderID: out.OrderID, Status: normalizeStatus(out.TransactionStatus, ""), Raw: pg.RawJSON(out)}, nil
}

func (s *ServiceImpl) CreatePayout(ctx context.Context, req pg.CreatePayoutRequest) (pg.CreatePayoutResponse, error) {
	return pg.CreatePayoutResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "midtrans payout is not implemented in this adapter", false, pg.ErrUnsupportedOperation)
}

func (s *ServiceImpl) GetPayoutStatus(ctx context.Context, req pg.GetPayoutStatusRequest) (pg.GetPayoutStatusResponse, error) {
	return pg.GetPayoutStatusResponse{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeUnsupported, http.StatusNotImplemented, "midtrans payout status is not implemented in this adapter", false, pg.ErrUnsupportedOperation)
}

func (s *ServiceImpl) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	var n model.Notification
	if err := json.Unmarshal(req.RawBody, &n); err != nil {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), SignatureValid: false, Reason: "invalid json"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid midtrans webhook json", false, err)
	}
	expected := pg.SHA512Hex(n.OrderID + n.StatusCode + n.GrossAmount + s.cfg.ServerKey)
	ok := n.SignatureKey != "" && pg.SecureEqualHex(expected, n.SignatureKey)
	if !ok {
		return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, SignatureValid: false, Reason: "signature mismatch"}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "midtrans signature mismatch", false, pg.ErrInvalidWebhook)
	}
	return pg.VerifyWebhookResult{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, SignatureValid: true}, nil
}

func (s *ServiceImpl) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	var n model.Notification
	if err := json.Unmarshal(req.RawBody, &n); err != nil {
		return pg.CanonicalPaymentEvent{}, pg.NewGatewayError(s.ProviderCode(), pg.ErrorCodeWebhookInvalid, 0, "invalid midtrans webhook json", false, err)
	}
	status := normalizeStatus(n.TransactionStatus, n.FraudStatus)
	return pg.CanonicalPaymentEvent{ProviderCode: s.ProviderCode(), EventID: n.TransactionID, EventType: eventFromStatus(status), ProviderEventType: n.PaymentType, ProviderStatus: n.TransactionStatus, PaymentStatus: status, OrderID: n.OrderID, ProviderReference: n.OrderID, ProviderTransactionID: n.TransactionID, Amount: pg.Money{Amount: n.GrossAmount, Currency: firstNonEmpty(n.Currency, s.cfg.Currency())}, Raw: json.RawMessage(req.RawBody)}, nil
}

func toCustomer(c pg.Customer) *model.CustomerDetails {
	if c.Name == "" && c.Email == "" && c.Phone == "" {
		return nil
	}
	return &model.CustomerDetails{FirstName: c.Name, Email: c.Email, Phone: c.Phone}
}

func toItems(items []pg.Item) []model.ItemDetail {
	if len(items) == 0 {
		return nil
	}
	out := make([]model.ItemDetail, 0, len(items))
	for _, it := range items {
		price, _ := pg.AmountToFloat64(it.Price)
		out = append(out, model.ItemDetail{ID: it.ID, Price: price, Quantity: it.Quantity, Name: it.Name, Category: it.Category})
	}
	return out
}

func (s *ServiceImpl) paymentType(method pg.PaymentMethod, channel string) string {
	switch method {
	case pg.PaymentMethodVirtualAccount:
		return "bank_transfer"
	case pg.PaymentMethodQRIS:
		return "qris"
	case pg.PaymentMethodEWallet:
		if strings.Contains(channel, "gopay") || channel == "" {
			return "gopay"
		}
	}
	return string(method)
}

func bankCode(channel string) string {
	c := strings.ToLower(channel)
	for _, suffix := range []string{"_va", "_virtual_account"} {
		c = strings.TrimSuffix(c, suffix)
	}
	if c == "bca" || c == "bni" || c == "bri" || c == "permata" || c == "mandiri" || c == "cimb" {
		return c
	}
	return c
}

func normalizeStatus(status, fraud string) pg.PaymentStatus {
	s := strings.ToLower(status)
	f := strings.ToLower(fraud)
	switch s {
	case "capture":
		if f == "challenge" {
			return pg.PaymentStatusManualReview
		}
		return pg.PaymentStatusSucceeded
	case "settlement":
		return pg.PaymentStatusSucceeded
	case "pending":
		return pg.PaymentStatusPending
	case "deny", "failure":
		return pg.PaymentStatusFailed
	case "cancel":
		return pg.PaymentStatusCanceled
	case "expire":
		return pg.PaymentStatusExpired
	case "refund":
		return pg.PaymentStatusRefunded
	case "partial_refund":
		return pg.PaymentStatusPartiallyRefunded
	case "authorize":
		return pg.PaymentStatusAuthorized
	}
	return pg.PaymentStatusUnknown
}

func eventFromStatus(status pg.PaymentStatus) pg.EventType {
	switch status {
	case pg.PaymentStatusSucceeded, pg.PaymentStatusCaptured:
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
	default:
		return pg.EventUnknown
	}
}

func instructions(out model.ChargeResponse) []pg.PaymentInstruction {
	var result []pg.PaymentInstruction
	for _, va := range out.VANumbers {
		result = append(result, pg.PaymentInstruction{Type: "va_number", DisplayName: fmt.Sprintf("%s Virtual Account", strings.ToUpper(va.Bank)), AccountNumber: va.VANumber, ProviderData: map[string]any{"bank": va.Bank}})
	}
	if out.PermataVANumber != "" {
		result = append(result, pg.PaymentInstruction{Type: "va_number", DisplayName: "Permata Virtual Account", AccountNumber: out.PermataVANumber})
	}
	if out.QRString != "" {
		result = append(result, pg.PaymentInstruction{Type: "qr_string", DisplayName: "QRIS", QRString: out.QRString})
	}
	for _, a := range out.Actions {
		ins := pg.PaymentInstruction{Type: "checkout_url", DisplayName: a.Name, CheckoutURL: a.URL, ProviderData: map[string]any{"method": a.Method}}
		if strings.Contains(strings.ToLower(a.Name), "qr") {
			ins.Type = "qr_image_url"
			ins.QRImageURL = a.URL
		}
		if strings.Contains(strings.ToLower(a.Name), "deeplink") {
			ins.Type = "deeplink"
			ins.DeeplinkURL = a.URL
		}
		result = append(result, ins)
	}
	return result
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

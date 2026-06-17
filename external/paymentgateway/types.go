package paymentgateway

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// ProviderCode is a stable canonical provider identifier used by lokatra-payment.
type ProviderCode string

const (
	ProviderMidtrans  ProviderCode = "midtrans"
	ProviderXendit    ProviderCode = "xendit"
	ProviderFinpay    ProviderCode = "finpay"
	ProviderDurianpay ProviderCode = "durianpay"
	ProviderIpaymu    ProviderCode = "ipaymu"
)

// PaymentMethod is a high-level payment method family.
type PaymentMethod string

const (
	PaymentMethodVirtualAccount PaymentMethod = "virtual_account"
	PaymentMethodQRIS           PaymentMethod = "qris"
	PaymentMethodEWallet        PaymentMethod = "ewallet"
	PaymentMethodCard           PaymentMethod = "card"
	PaymentMethodRetailOutlet   PaymentMethod = "retail_outlet"
	PaymentMethodPaymentPage    PaymentMethod = "payment_page"
	PaymentMethodManualTransfer PaymentMethod = "manual_transfer"
	PaymentMethodCash           PaymentMethod = "cash"
)

// PaymentStatus is Lokatra canonical payment status. Provider-specific status must be normalized into this enum.
type PaymentStatus string

const (
	PaymentStatusUnknown            PaymentStatus = "unknown"
	PaymentStatusCreated            PaymentStatus = "created"
	PaymentStatusPending            PaymentStatus = "pending"
	PaymentStatusRequiresAction     PaymentStatus = "requires_action"
	PaymentStatusSucceeded          PaymentStatus = "succeeded"
	PaymentStatusFailed             PaymentStatus = "failed"
	PaymentStatusExpired            PaymentStatus = "expired"
	PaymentStatusCanceled           PaymentStatus = "canceled"
	PaymentStatusAuthorized         PaymentStatus = "authorized"
	PaymentStatusCaptured           PaymentStatus = "captured"
	PaymentStatusRefunded           PaymentStatus = "refunded"
	PaymentStatusPartiallyRefunded  PaymentStatus = "partially_refunded"
	PaymentStatusDisputed           PaymentStatus = "disputed"
	PaymentStatusManualReview       PaymentStatus = "manual_review"
	PaymentStatusStatusSyncRequired PaymentStatus = "status_sync_required"
)

// EventType is Lokatra canonical webhook/provider event type.
type EventType string

const (
	EventPaymentCreated           EventType = "payment.created"
	EventPaymentPending           EventType = "payment.pending"
	EventPaymentSucceeded         EventType = "payment.succeeded"
	EventPaymentFailed            EventType = "payment.failed"
	EventPaymentExpired           EventType = "payment.expired"
	EventPaymentCanceled          EventType = "payment.canceled"
	EventPaymentRefunded          EventType = "payment.refunded"
	EventPaymentPartiallyRefunded EventType = "payment.partially_refunded"
	EventPaymentDisputed          EventType = "payment.disputed"
	EventPayoutSucceeded          EventType = "payout.succeeded"
	EventPayoutFailed             EventType = "payout.failed"
	EventUnknown                  EventType = "unknown"
)

// Money must be sent as string in API JSON for precision, but gateway adapters often need numeric values.
type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Customer is a minimal gateway customer snapshot. Do not store sensitive card data here.
type Customer struct {
	ID          string `json:"id,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// Item is a minimal line item snapshot sent to providers that support itemization.
type Item struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category,omitempty"`
}

// CreatePaymentRequest is the canonical request consumed by provider adapters.
type CreatePaymentRequest struct {
	PaymentIntentID string            `json:"paymentIntentId"`
	AttemptID       string            `json:"attemptId"`
	OrderID         string            `json:"orderId"`
	Amount          Money             `json:"amount"`
	Method          PaymentMethod     `json:"method"`
	ChannelCode     string            `json:"channelCode,omitempty"` // e.g. bca_va, mandiri_va, qris, gopay
	Description     string            `json:"description,omitempty"`
	Customer        Customer          `json:"customer,omitempty"`
	Items           []Item            `json:"items,omitempty"`
	ExpiryAt        *time.Time        `json:"expiryAt,omitempty"`
	CallbackURL     string            `json:"callbackUrl,omitempty"`
	ReturnURL       string            `json:"returnUrl,omitempty"`
	FailureURL      string            `json:"failureUrl,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	IdempotencyKey  string            `json:"idempotencyKey,omitempty"`
}

type PaymentInstruction struct {
	Type          string         `json:"type"` // va_number, qr_string, checkout_url, deeplink, retail_code
	DisplayName   string         `json:"displayName,omitempty"`
	AccountNumber string         `json:"accountNumber,omitempty"`
	BillerCode    string         `json:"billerCode,omitempty"`
	PaymentCode   string         `json:"paymentCode,omitempty"`
	QRString      string         `json:"qrString,omitempty"`
	QRImageURL    string         `json:"qrImageUrl,omitempty"`
	CheckoutURL   string         `json:"checkoutUrl,omitempty"`
	DeeplinkURL   string         `json:"deeplinkUrl,omitempty"`
	ExpiresAt     *time.Time     `json:"expiresAt,omitempty"`
	ProviderData  map[string]any `json:"providerData,omitempty"`
}

type CreatePaymentResponse struct {
	ProviderCode          ProviderCode         `json:"providerCode"`
	ProviderReference     string               `json:"providerReference,omitempty"`
	ProviderTransactionID string               `json:"providerTransactionId,omitempty"`
	ProviderPaymentID     string               `json:"providerPaymentId,omitempty"`
	OrderID               string               `json:"orderId"`
	Status                PaymentStatus        `json:"status"`
	Instructions          []PaymentInstruction `json:"instructions,omitempty"`
	Raw                   json.RawMessage      `json:"raw,omitempty"`
}

type GetPaymentStatusRequest struct {
	OrderID               string `json:"orderId,omitempty"`
	ProviderReference     string `json:"providerReference,omitempty"`
	ProviderTransactionID string `json:"providerTransactionId,omitempty"`
}

type GetPaymentStatusResponse struct {
	ProviderCode          ProviderCode    `json:"providerCode"`
	ProviderReference     string          `json:"providerReference,omitempty"`
	ProviderTransactionID string          `json:"providerTransactionId,omitempty"`
	OrderID               string          `json:"orderId,omitempty"`
	Status                PaymentStatus   `json:"status"`
	Amount                Money           `json:"amount,omitempty"`
	PaidAt                *time.Time      `json:"paidAt,omitempty"`
	Raw                   json.RawMessage `json:"raw,omitempty"`
}

type CancelPaymentRequest struct {
	OrderID               string `json:"orderId,omitempty"`
	ProviderReference     string `json:"providerReference,omitempty"`
	ProviderTransactionID string `json:"providerTransactionId,omitempty"`
	Reason                string `json:"reason,omitempty"`
	IdempotencyKey        string `json:"idempotencyKey,omitempty"`
}

type CancelPaymentResponse struct {
	ProviderCode ProviderCode    `json:"providerCode"`
	OrderID      string          `json:"orderId,omitempty"`
	Status       PaymentStatus   `json:"status"`
	Raw          json.RawMessage `json:"raw,omitempty"`
}

type RefundRequest struct {
	OrderID               string `json:"orderId,omitempty"`
	ProviderReference     string `json:"providerReference,omitempty"`
	ProviderTransactionID string `json:"providerTransactionId,omitempty"`
	RefundID              string `json:"refundId"`
	Amount                Money  `json:"amount"`
	Reason                string `json:"reason,omitempty"`
	IdempotencyKey        string `json:"idempotencyKey,omitempty"`
}

type RefundResponse struct {
	ProviderCode     ProviderCode    `json:"providerCode"`
	ProviderRefundID string          `json:"providerRefundId,omitempty"`
	OrderID          string          `json:"orderId,omitempty"`
	Status           PaymentStatus   `json:"status"`
	Raw              json.RawMessage `json:"raw,omitempty"`
}

type CreatePayoutRequest struct {
	PayoutID       string            `json:"payoutId"`
	ExternalID     string            `json:"externalId"`
	Amount         Money             `json:"amount"`
	BankCode       string            `json:"bankCode,omitempty"`
	AccountNumber  string            `json:"accountNumber,omitempty"`
	AccountName    string            `json:"accountName,omitempty"`
	Description    string            `json:"description,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	IdempotencyKey string            `json:"idempotencyKey,omitempty"`
}

type CreatePayoutResponse struct {
	ProviderCode     ProviderCode    `json:"providerCode"`
	ProviderPayoutID string          `json:"providerPayoutId,omitempty"`
	Status           PaymentStatus   `json:"status"`
	Raw              json.RawMessage `json:"raw,omitempty"`
}

type GetPayoutStatusRequest struct {
	PayoutID         string `json:"payoutId,omitempty"`
	ProviderPayoutID string `json:"providerPayoutId,omitempty"`
}

type GetPayoutStatusResponse struct {
	ProviderCode     ProviderCode    `json:"providerCode"`
	ProviderPayoutID string          `json:"providerPayoutId,omitempty"`
	Status           PaymentStatus   `json:"status"`
	Raw              json.RawMessage `json:"raw,omitempty"`
}

type VerifyWebhookRequest struct {
	Headers http.Header `json:"-"`
	RawBody []byte      `json:"-"`
}

type VerifyWebhookResult struct {
	ProviderCode   ProviderCode `json:"providerCode"`
	SignatureValid bool         `json:"signatureValid"`
	EventID        string       `json:"eventId,omitempty"`
	Reason         string       `json:"reason,omitempty"`
}

type NormalizeWebhookRequest struct {
	Headers http.Header `json:"-"`
	RawBody []byte      `json:"-"`
}

type CanonicalPaymentEvent struct {
	ProviderCode          ProviderCode    `json:"providerCode"`
	EventID               string          `json:"eventId,omitempty"`
	EventType             EventType       `json:"eventType"`
	ProviderEventType     string          `json:"providerEventType,omitempty"`
	ProviderStatus        string          `json:"providerStatus,omitempty"`
	PaymentStatus         PaymentStatus   `json:"paymentStatus"`
	OrderID               string          `json:"orderId,omitempty"`
	ProviderReference     string          `json:"providerReference,omitempty"`
	ProviderTransactionID string          `json:"providerTransactionId,omitempty"`
	Amount                Money           `json:"amount,omitempty"`
	OccurredAt            *time.Time      `json:"occurredAt,omitempty"`
	Raw                   json.RawMessage `json:"raw,omitempty"`
}

type Capability struct {
	Method                PaymentMethod `json:"method"`
	ChannelCode           string        `json:"channelCode,omitempty"`
	Currency              string        `json:"currency"`
	SupportsRefund        bool          `json:"supportsRefund"`
	SupportsPartialRefund bool          `json:"supportsPartialRefund"`
	SupportsPayout        bool          `json:"supportsPayout"`
	SupportsExpiry        bool          `json:"supportsExpiry"`
}

type CapabilitiesRequest struct{}

type CapabilitiesResponse struct {
	ProviderCode ProviderCode `json:"providerCode"`
	Items        []Capability `json:"items"`
}

// PaymentGateway is the stable adapter interface consumed by payment orchestration.
type PaymentGateway interface {
	ProviderCode() ProviderCode
	Capabilities(ctx context.Context, req CapabilitiesRequest) (CapabilitiesResponse, error)
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (CreatePaymentResponse, error)
	GetPaymentStatus(ctx context.Context, req GetPaymentStatusRequest) (GetPaymentStatusResponse, error)
	CancelPayment(ctx context.Context, req CancelPaymentRequest) (CancelPaymentResponse, error)
	RefundPayment(ctx context.Context, req RefundRequest) (RefundResponse, error)
	CreatePayout(ctx context.Context, req CreatePayoutRequest) (CreatePayoutResponse, error)
	GetPayoutStatus(ctx context.Context, req GetPayoutStatusRequest) (GetPayoutStatusResponse, error)
	VerifyWebhook(ctx context.Context, req VerifyWebhookRequest) (VerifyWebhookResult, error)
	NormalizeWebhook(ctx context.Context, req NormalizeWebhookRequest) (CanonicalPaymentEvent, error)
}

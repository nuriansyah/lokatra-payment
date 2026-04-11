package dto

import (
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/shopspring/decimal"
)

// PaymentFlowRequestDTO represents the HTTP request body for payment flow operations.
// @Description Payment flow request for obtaining a routing quote or executing a payment charge.
// The contract is PSP-agnostic and is used by both /flows/quote and /flows/execute.
type PaymentFlowRequestDTO struct {
	// Unique identifier of the actor initiating the payment.
	// @Required true
	ActorID string `json:"actorId" validate:"required,uuid" example:"8e5aa8f0-8f31-4f8b-b2cc-f9276f4e35c8"`

	// Unique identifier of the merchant processing the payment.
	// @Required true
	MerchantID string `json:"merchantId" validate:"required,uuid" example:"3da7f699-6c84-4eff-b157-b467f6fa62d4"`

	// Payment use case for routing and policy decisions.
	// @Required true
	UseCase string `json:"useCase" validate:"required,oneof=EVENT_PAYMENT TOUR_TRIP GENERIC_TRANSACTION" example:"EVENT_PAYMENT"`

	// Order identifier in merchant's system.
	OrderID string `json:"orderId,omitempty" example:"b69c66a4-3d2a-4ec3-8011-185860f4e11d"`

	// Type of order (e.g., EVENT_TICKET, TOUR_BOOKING).
	OrderType string `json:"orderType,omitempty" example:"EVENT_TICKET"`

	// Product type being purchased.
	ProductType string `json:"productType,omitempty" example:"EVENT_TICKET"`

	// Payment amount in smallest currency unit (e.g., cents for USD, rupiah for IDR).
	// @Required true
	Amount decimal.Decimal `json:"amount" validate:"required" example:"350000"`

	// ISO 4217 currency code.
	// @Required true
	Currency string `json:"currency" validate:"required,len=3" example:"IDR"`

	// Tax amount applied to the transaction.
	TaxAmount decimal.Decimal `json:"taxAmount,omitempty" example:"35000"`

	// Discount amount applied to the transaction.
	DiscountAmount decimal.Decimal `json:"discountAmount,omitempty" example:"0"`

	// Tip amount provided by customer.
	TipAmount decimal.Decimal `json:"tipAmount,omitempty" example:"0"`

	// Unique identifier of the customer in merchant's system.
	UserID string `json:"userId,omitempty" example:"user-123"`

	// Full name of the customer.
	CustomerName string `json:"customerName,omitempty" validate:"max=255" example:"Integration Test User"`

	// Email address of the customer.
	CustomerEmail string `json:"customerEmail,omitempty" validate:"email" example:"integration-test@example.com"`

	// Phone number of the customer.
	CustomerPhone string `json:"customerPhone,omitempty" validate:"max=20" example:"+628111111111"`

	// ISO 3166-1 alpha-2 country code of the customer.
	CustomerCountry string `json:"customerCountry,omitempty" validate:"len=2" example:"ID"`

	// Customer's IP address for risk assessment.
	CustomerIP string `json:"customerIp,omitempty" validate:"ipv4|ipv6" example:"192.168.1.1"`

	// Payment method identifier in merchant's system.
	PaymentMethodID string `json:"paymentMethodId,omitempty" example:"pm-123"`

	// Card BIN (Bank Identification Number) if applicable.
	CardBIN string `json:"cardBin,omitempty" validate:"max=6" example:"512345"`

	// Payment method type (e.g., CARD, VIRTUAL_ACCOUNT, QRIS, EWALLET).
	// @Required true
	PaymentMethodType string `json:"paymentMethodType" validate:"required,oneof=CARD VIRTUAL_ACCOUNT QRIS EWALLET DIRECT_DEBIT BANK_TRANSFER PAYLATER VOUCHER POINTS CASH_ON_DELIVERY" example:"CARD"`

	// Indicate whether 3D Secure authentication is required.
	Requires3DS bool `json:"requires3ds,omitempty" example:"true"`

	// Statement descriptor shown on customer's bank statement.
	StatementDescriptor string `json:"statementDescriptor,omitempty" validate:"max=22" example:"LOKATRA EVENT"`

	// Description of the transaction for customer reference.
	Description string `json:"description,omitempty" validate:"max=255" example:"2x VIP event ticket"`

	// Payment capture mode (AUTOMATIC or MANUAL).
	CaptureMode string `json:"captureMode,omitempty" validate:"oneof=AUTOMATIC MANUAL" example:"AUTOMATIC"`

	// Optional routing profile ID to enforce specific provider selection.
	RoutingProfileID string `json:"routingProfileId,omitempty" validate:"uuid" example:"routing-profile-123"`

	// Idempotency key for retry safety. If provided, duplicate requests return same result.
	IdempotencyKey string `json:"idempotencyKey,omitempty" validate:"max=128" example:"ik-event-ticket-12345"`

	// Correlation ID for tracing related transactions.
	CorrelationID string `json:"correlationId,omitempty" validate:"max=128" example:"corr-123456"`

	// Additional metadata as key-value pairs for extensibility.
	Metadata map[string]interface{} `json:"metadata,omitempty" swaggertype:"object"`
}

// WebhookEventRequestDTO represents a normalized webhook event contract.
// @Description Normalized webhook payload accepted by the Midtrans and Xendit webhook endpoints.
// The request body carries provider metadata, a normalized status, and a raw payload map for audit.
type WebhookEventRequestDTO struct {
	// Event type identifier (e.g., "payment.authorized", "payment.captured", "payment.failed").
	EventType string `json:"eventType" validate:"required" example:"payment.captured"`

	// Provider identifier (e.g., "MIDTRANS", "XENDIT").
	Provider string `json:"provider" validate:"required,oneof=MIDTRANS XENDIT" example:"MIDTRANS"`

	// Transaction ID in the provider's system.
	TransactionID string `json:"transactionId" validate:"required" example:"txn-12345"`

	// External reference for idempotency (prevents duplicate processing).
	ExternalID string `json:"externalId,omitempty" validate:"max=128" example:"ext-ref-12345"`

	// Payment status from the provider (e.g., "settlement", "capture", "deny").
	Status string `json:"status" validate:"required" example:"settlement"`

	// ISO timestamp when the event occurred at the provider.
	EventTimestamp string `json:"eventTimestamp" validate:"required" example:"2026-04-11T13:21:20Z"`

	// Raw event payload from provider for audit and debugging.
	RawPayload map[string]interface{} `json:"rawPayload" swaggertype:"object"`

	// Signature for webhook authenticity verification.
	Signature string `json:"signature" validate:"required" example:"sha512=abcdef123456"`

	// Signature method used by the provider contract.
	SignatureMethod string `json:"signatureMethod,omitempty" validate:"omitempty,oneof=SHA512 SHA256 SHA1 MD5 HMACSHA256" example:"SHA512"`
}

type PaymentUseCase string

const (
	PaymentUseCaseEventPayment       PaymentUseCase = "EVENT_PAYMENT"
	PaymentUseCaseTourOrTrip         PaymentUseCase = "TOUR_TRIP"
	PaymentUseCaseGenericTransaction PaymentUseCase = "GENERIC_TRANSACTION"
)

type PaymentCaptureMode string

const (
	PaymentCaptureModeAutomatic PaymentCaptureMode = "AUTOMATIC"
	PaymentCaptureModeManual    PaymentCaptureMode = "MANUAL"
)

type PaymentFlowRequest struct {
	ActorID             uuid.UUID               `json:"actorId" validate:"required"`
	MerchantID          uuid.UUID               `json:"merchantId" validate:"required"`
	UseCase             PaymentUseCase          `json:"useCase" validate:"required"`
	OrderID             string                  `json:"orderId,omitempty"`
	OrderType           string                  `json:"orderType,omitempty"`
	ProductType         string                  `json:"productType,omitempty"`
	Amount              decimal.Decimal         `json:"amount" validate:"required"`
	Currency            model.PaymentCurrency   `json:"currency" validate:"required"`
	TaxAmount           decimal.Decimal         `json:"taxAmount,omitempty"`
	DiscountAmount      decimal.Decimal         `json:"discountAmount,omitempty"`
	TipAmount           decimal.Decimal         `json:"tipAmount,omitempty"`
	UserID              string                  `json:"userId,omitempty"`
	CustomerName        string                  `json:"customerName,omitempty"`
	CustomerEmail       string                  `json:"customerEmail,omitempty"`
	CustomerPhone       string                  `json:"customerPhone,omitempty"`
	CustomerCountry     string                  `json:"customerCountry,omitempty"`
	CustomerIP          string                  `json:"customerIp,omitempty"`
	PaymentMethodID     string                  `json:"paymentMethodId,omitempty"`
	CardBIN             string                  `json:"cardBin,omitempty"`
	PaymentMethodType   model.PaymentMethodType `json:"paymentMethodType" validate:"required"`
	Requires3DS         bool                    `json:"requires3ds"`
	StatementDescriptor string                  `json:"statementDescriptor,omitempty"`
	Description         string                  `json:"description,omitempty"`
	CaptureMode         PaymentCaptureMode      `json:"captureMode,omitempty"`
	RoutingProfileID    string                  `json:"routingProfileId,omitempty"`
	IdempotencyKey      string                  `json:"idempotencyKey,omitempty"`
	CorrelationID       string                  `json:"correlationId,omitempty"`
	Metadata            map[string]any          `json:"metadata,omitempty"`
}

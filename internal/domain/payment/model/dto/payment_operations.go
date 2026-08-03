package dto

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// CreatePaymentIntentRequest
// ---------------------------------------------------------------------------

type CreatePaymentIntentRequest struct {
	ActorID            uuid.UUID       `json:"actorId"`
	MerchantID         uuid.UUID       `json:"merchantId"`
	CustomerID         uuid.UUID       `json:"customerId,omitempty"`
	SourceService      string          `json:"sourceService"`
	SourceType         string          `json:"sourceType"`
	SourceID           uuid.UUID       `json:"sourceId"`
	Amount             decimal.Decimal `json:"amount"`
	Currency           string          `json:"currency"`
	PaymentMethodCode  string          `json:"paymentMethodCode,omitempty"`
	PaymentChannelCode string          `json:"paymentChannelCode,omitempty"`
	Description        string          `json:"description,omitempty"`
	ExpiresAt          *time.Time      `json:"expiresAt,omitempty"`
	IdempotencyKey     string          `json:"idempotencyKey,omitempty"`
	SourceSnapshot     json.RawMessage `json:"sourceSnapshot,omitempty"`
}

func (d CreatePaymentIntentRequest) Validate() error {
	if d.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidActorID, "actorId is required")
	}
	if d.MerchantID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidMerchantID, "merchantId is required")
	}
	if d.SourceID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidSourceID, "sourceId is required")
	}
	if d.SourceService == "" {
		return failure.WithCode(shared.ErrInvalidSourceService, "sourceService is required")
	}
	if d.SourceType == "" {
		return failure.WithCode(shared.ErrInvalidSourceType, "sourceType is required")
	}
	if d.Amount.LessThanOrEqual(decimal.Zero) {
		return failure.WithCode(shared.ErrInvalidAmount, "amount must be positive")
	}
	if d.Currency == "" {
		return failure.WithCode(shared.ErrInvalidCurrency, "currency is required")
	}
	if d.ExpiresAt != nil && !d.ExpiresAt.After(time.Now().UTC()) {
		return failure.WithCode(shared.ErrInvalidExpiryDate, "expiresAt must be in the future")
	}
	if len(d.SourceSnapshot) > 0 && !json.Valid(d.SourceSnapshot) {
		return failure.WithCode(shared.ErrInvalidMetadata, "sourceSnapshot must contain valid JSON")
	}
	return nil
}

// ---------------------------------------------------------------------------
// CreateRefundRequest
// ---------------------------------------------------------------------------

type CreateRefundRequest struct {
	ActorID          uuid.UUID       `json:"actorId"`
	PaymentIntentID  uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptID uuid.UUID       `json:"paymentAttemptId,omitempty"`
	Amount           decimal.Decimal `json:"amount"`
	Currency         string          `json:"currency"`
	Reason           string          `json:"reason"`
	IdempotencyKey   string          `json:"idempotencyKey,omitempty"`
}

func (d CreateRefundRequest) Validate() error {
	if d.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidActorID, "actorId is required")
	}
	if d.PaymentIntentID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "paymentIntentId is required")
	}
	if d.Amount.LessThanOrEqual(decimal.Zero) {
		return failure.WithCode(shared.ErrInvalidAmount, "refund amount must be positive")
	}
	return nil
}

// ---------------------------------------------------------------------------
// OpenCashSessionRequest
// ---------------------------------------------------------------------------

type OpenCashSessionRequest struct {
	ActorID            uuid.UUID       `json:"actorId"`
	MerchantID         uuid.UUID       `json:"merchantId"`
	CollectorID        uuid.UUID       `json:"collectorId"`
	LocationID         uuid.UUID       `json:"locationId,omitempty"`
	OpeningFloatAmount decimal.Decimal `json:"openingFloatAmount"`
	Currency           string          `json:"currency"`
	Notes              string          `json:"notes,omitempty"`
}

func (d OpenCashSessionRequest) Validate() error {
	if d.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidActorID, "actorId is required")
	}
	if d.MerchantID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidMerchantID, "merchantId is required")
	}
	if d.CollectorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidCollectorID, "collectorId is required")
	}
	if d.OpeningFloatAmount.IsNegative() {
		return failure.WithCode(shared.ErrInvalidAmount, "openingFloatAmount cannot be negative")
	}
	if d.Currency == "" {
		return failure.WithCode(shared.ErrInvalidCurrency, "currency is required")
	}
	return nil
}

// ---------------------------------------------------------------------------
// ActionCommand
// ---------------------------------------------------------------------------

type ActionCommand struct {
	ActorID        uuid.UUID       `json:"actorId"`
	Reason         string          `json:"reason,omitempty"`
	Notes          string          `json:"notes,omitempty"`
	FailureCode    string          `json:"failureCode,omitempty"`
	FailureMessage string          `json:"failureMessage,omitempty"`
	Amount         decimal.Decimal `json:"amount,omitempty"`
}

func (d ActionCommand) Validate() error {
	if d.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidActorID, "actorId is required")
	}
	return nil
}

// ---------------------------------------------------------------------------
// Webhook / Routing DTOs
// ---------------------------------------------------------------------------

type WebhookReceipt struct {
	Provider       string    `json:"provider"`
	EventID        string    `json:"eventId,omitempty"`
	EventType      string    `json:"eventType"`
	PaymentStatus  string    `json:"paymentStatus"`
	OrderID        string    `json:"orderId,omitempty"`
	SignatureValid bool      `json:"signatureValid"`
	ReceivedAt     time.Time `json:"receivedAt"`
}

type RoutingCandidateResponse struct {
	ProviderCode string    `json:"providerCode"`
	AccountID    uuid.UUID `json:"accountId"`
	Priority     int       `json:"priority"`
	MaxAttempts  int       `json:"maxAttempts"`
	Reason       string    `json:"reason"`
	Skipped      bool      `json:"skipped"`
	SkipReason   string    `json:"skipReason,omitempty"`
}

type ProviderAttemptResponse struct {
	ProviderCode string        `json:"providerCode"`
	AccountID    uuid.UUID     `json:"accountId"`
	Attempt      int           `json:"attempt"`
	StartedAt    time.Time     `json:"startedAt"`
	Duration     time.Duration `json:"duration"`
	Error        string        `json:"error,omitempty"`
}

type RoutingExecutionResponse struct {
	Selected   *RoutingCandidateResponse  `json:"selected,omitempty"`
	Candidates []RoutingCandidateResponse `json:"candidates"`
	Attempts   []ProviderAttemptResponse  `json:"attempts"`
	Payment    *pg.CreatePaymentResponse  `json:"payment,omitempty"`
}

type PaymentIntentActionResponse struct {
	Intent  PaymentIntentsResponse    `json:"intent"`
	Routing *RoutingExecutionResponse `json:"routing,omitempty"`
}

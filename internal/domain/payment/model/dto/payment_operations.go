package dto

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/shopspring/decimal"
)

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
	Metadata           json.RawMessage `json:"metadata,omitempty"`
}

type CreateRefundRequest struct {
	ActorID          uuid.UUID       `json:"actorId"`
	PaymentIntentID  uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptID uuid.UUID       `json:"paymentAttemptId,omitempty"`
	Amount           decimal.Decimal `json:"amount"`
	Currency         string          `json:"currency"`
	Reason           string          `json:"reason"`
	IdempotencyKey   string          `json:"idempotencyKey,omitempty"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
}

type OpenCashSessionRequest struct {
	ActorID            uuid.UUID       `json:"actorId"`
	MerchantID         uuid.UUID       `json:"merchantId"`
	CollectorID        uuid.UUID       `json:"collectorId"`
	LocationID         uuid.UUID       `json:"locationId,omitempty"`
	OpeningFloatAmount decimal.Decimal `json:"openingFloatAmount"`
	Currency           string          `json:"currency"`
	Notes              string          `json:"notes,omitempty"`
	Metadata           json.RawMessage `json:"metadata,omitempty"`
}

// ActionCommand contains audit data and action-specific values. Callers cannot
// submit an arbitrary status; each aggregate controls its allowed transitions.
type ActionCommand struct {
	ActorID        uuid.UUID       `json:"actorId"`
	Reason         string          `json:"reason,omitempty"`
	Notes          string          `json:"notes,omitempty"`
	FailureCode    string          `json:"failureCode,omitempty"`
	FailureMessage string          `json:"failureMessage,omitempty"`
	Amount         decimal.Decimal `json:"amount,omitempty"`
}

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

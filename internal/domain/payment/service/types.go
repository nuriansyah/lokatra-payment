package service

import (
	"time"

	"github.com/gofrs/uuid"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	routingmodel "github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	"github.com/shopspring/decimal"
)

// PaymentUseCase lets the routing engine stay configuration-first while still
// supporting different transactional scenarios.
type PaymentUseCase string

const (
	PaymentUseCaseEventPayment       PaymentUseCase = "EVENT_PAYMENT"
	PaymentUseCaseTourOrTrip         PaymentUseCase = "TOUR_TRIP"
	PaymentUseCaseGenericTransaction PaymentUseCase = "GENERIC_TRANSACTION"
)

// PaymentCaptureMode controls how the PSP flow should behave.
type PaymentCaptureMode string

const (
	PaymentCaptureModeAutomatic PaymentCaptureMode = "AUTOMATIC"
	PaymentCaptureModeManual    PaymentCaptureMode = "MANUAL"
)

// ProviderConfig describes one configurable PSP account backed by an adapter.
type ProviderConfig struct {
	PSP                 paymentmodel.Psp
	AccountID           uuid.UUID
	AccountLabel        string
	Enabled             bool
	SupportedMethods    []paymentmodel.PaymentMethodType
	SupportedCurrencies []paymentmodel.PaymentCurrency
}

func (p ProviderConfig) supports(method paymentmodel.PaymentMethodType, currency paymentmodel.PaymentCurrency) bool {
	if !p.Enabled {
		return false
	}
	if len(p.SupportedMethods) > 0 && !containsPaymentMethod(p.SupportedMethods, method) {
		return false
	}
	if len(p.SupportedCurrencies) > 0 && !containsCurrency(p.SupportedCurrencies, currency) {
		return false
	}
	return true
}

// RoutingPolicy is a declarative rule evaluated before any database fallback.
type RoutingPolicy struct {
	Name                string
	Enabled             bool
	Priority            int
	UseCases            []PaymentUseCase
	MerchantIDs         []uuid.UUID
	OrderTypes          []string
	ProductTypes        []string
	PaymentMethods      []paymentmodel.PaymentMethodType
	Currencies          []paymentmodel.PaymentCurrency
	Countries           []string
	MinAmount           *decimal.Decimal
	MaxAmount           *decimal.Decimal
	PreferredPSP        paymentmodel.Psp
	PreferredAccount    uuid.UUID
	Strategy            routingmodel.RoutingStrategy
	DatabaseFallback    bool
	CandidateAccountIDs []uuid.UUID
}

// PaymentServiceConfig is a configuration-first routing catalog.
type PaymentServiceConfig struct {
	UseDatabaseFallback bool
	DefaultUseCase      PaymentUseCase
	DefaultStrategy     routingmodel.RoutingStrategy
	Providers           []ProviderConfig
	Policies            []RoutingPolicy
}

// PaymentFlowRequest is the normalized request accepted by the facade.
type PaymentFlowRequest struct {
	ActorID             uuid.UUID                      `json:"actorId" validate:"required"`
	MerchantID          uuid.UUID                      `json:"merchantId" validate:"required"`
	UseCase             PaymentUseCase                 `json:"useCase" validate:"required"`
	OrderID             string                         `json:"orderId,omitempty"`
	OrderType           string                         `json:"orderType,omitempty"`
	ProductType         string                         `json:"productType,omitempty"`
	Amount              decimal.Decimal                `json:"amount" validate:"required"`
	Currency            paymentmodel.PaymentCurrency   `json:"currency" validate:"required"`
	TaxAmount           decimal.Decimal                `json:"taxAmount,omitempty"`
	DiscountAmount      decimal.Decimal                `json:"discountAmount,omitempty"`
	TipAmount           decimal.Decimal                `json:"tipAmount,omitempty"`
	UserID              string                         `json:"userId,omitempty"`
	CustomerName        string                         `json:"customerName,omitempty"`
	CustomerEmail       string                         `json:"customerEmail,omitempty"`
	CustomerPhone       string                         `json:"customerPhone,omitempty"`
	CustomerCountry     string                         `json:"customerCountry,omitempty"`
	CustomerIP          string                         `json:"customerIp,omitempty"`
	PaymentMethodID     string                         `json:"paymentMethodId,omitempty"`
	CardBIN             string                         `json:"cardBin,omitempty"`
	PaymentMethodType   paymentmodel.PaymentMethodType `json:"paymentMethodType" validate:"required"`
	Requires3DS         bool                           `json:"requires3ds"`
	StatementDescriptor string                         `json:"statementDescriptor,omitempty"`
	Description         string                         `json:"description,omitempty"`
	CaptureMode         PaymentCaptureMode             `json:"captureMode,omitempty"`
	RoutingProfileID    string                         `json:"routingProfileId,omitempty"`
	IdempotencyKey      string                         `json:"idempotencyKey,omitempty"`
	CorrelationID       string                         `json:"correlationId,omitempty"`
	Metadata            map[string]any                 `json:"metadata,omitempty"`
}

// RoutingCandidate is a scored PSP candidate used in the final decision.
type RoutingCandidate struct {
	PSP          paymentmodel.Psp `json:"psp"`
	AccountID    uuid.UUID        `json:"accountId"`
	AccountLabel string           `json:"accountLabel"`
	Score        int              `json:"score"`
	Reason       string           `json:"reason"`
}

// RoutingDecision is the final decision emitted by the engine.
type RoutingDecision struct {
	Strategy        routingmodel.RoutingStrategy `json:"strategy"`
	PolicyName      string                       `json:"policyName,omitempty"`
	ProfileID       uuid.UUID                    `json:"profileId,omitempty"`
	ProfileMatched  bool                         `json:"profileMatched"`
	RuleID          uuid.UUID                    `json:"ruleId,omitempty"`
	RuleMatched     bool                         `json:"ruleMatched"`
	PSP             paymentmodel.Psp             `json:"psp"`
	PSPAccountID    uuid.UUID                    `json:"pspAccountId"`
	PSPAccountLabel string                       `json:"pspAccountLabel,omitempty"`
	Reason          string                       `json:"reason,omitempty"`
	FallbackUsed    bool                         `json:"fallbackUsed"`
	Candidates      []RoutingCandidate           `json:"candidates,omitempty"`
}

// FlowStep is the explanation layer returned by the facade for observability.
type FlowStep struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PaymentFlowQuote is the public response of the routing facade.
type PaymentFlowQuote struct {
	Intent    *paymentmodel.PaymentIntents `json:"intent,omitempty"`
	Payment   *paymentmodel.Payments       `json:"payment,omitempty"`
	Decision  RoutingDecision              `json:"decision"`
	Steps     []FlowStep                   `json:"steps,omitempty"`
	CreatedAt time.Time                    `json:"createdAt"`
}

// PaymentExecutionResult is returned after a charge attempt.
type PaymentExecutionResult struct {
	Quote       PaymentFlowQuote             `json:"quote"`
	Charge      paymentmodel.ChargeResult    `json:"charge"`
	FinalStatus paymentmodel.PaymentStatus   `json:"finalStatus"`
	StatePath   []paymentmodel.PaymentStatus `json:"statePath,omitempty"`
}

// WebhookEvent is the service-level webhook contract.
type WebhookEvent struct {
	EventType       string
	Provider        string
	TransactionID   string
	ExternalID      string
	Status          string
	EventTimestamp  string
	RawPayload      map[string]interface{}
	Signature       string
	SignatureMethod string
}

type WebhookError struct {
	Code      string
	Message   string
	RequestID string
}

type WebhookResult struct {
	Provider      string
	Success       bool
	EventID       string
	TransactionID string
	PaymentID     string
	IntentID      string
	FinalStatus   string
	Message       string
	ProcessedAt   time.Time
	Error         *WebhookError
}

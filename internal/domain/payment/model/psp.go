package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

// GatewayDescriptor describes a PSP account exposed by a provider adapter.
type GatewayDescriptor struct {
	PSP                 Psp                 `json:"psp"`
	AccountID           uuid.UUID           `json:"accountId"`
	AccountLabel        string              `json:"accountLabel"`
	Enabled             bool                `json:"enabled"`
	SupportedMethods    []PaymentMethodType `json:"supportedMethods,omitempty"`
	SupportedCurrencies []PaymentCurrency   `json:"supportedCurrencies,omitempty"`
}

// ChargeRequest is the normalized request sent to a provider adapter.
type ChargeRequest struct {
	MerchantID          uuid.UUID         `json:"merchantId"`
	IntentID            uuid.UUID         `json:"intentId"`
	PaymentID           uuid.UUID         `json:"paymentId"`
	PaymentCode         string            `json:"paymentCode"`
	IntentCode          string            `json:"intentCode"`
	PSP                 Psp               `json:"psp"`
	Amount              decimal.Decimal   `json:"amount"`
	Currency            PaymentCurrency   `json:"currency"`
	PaymentMethodType   PaymentMethodType `json:"paymentMethodType"`
	CustomerName        string            `json:"customerName,omitempty"`
	CustomerEmail       string            `json:"customerEmail,omitempty"`
	CustomerPhone       string            `json:"customerPhone,omitempty"`
	CustomerCountry     string            `json:"customerCountry,omitempty"`
	Description         string            `json:"description,omitempty"`
	StatementDescriptor string            `json:"statementDescriptor,omitempty"`
	Metadata            json.RawMessage   `json:"metadata,omitempty"`
	UseCase             string            `json:"useCase,omitempty"`
	Requires3DS         bool              `json:"requires3ds"`
	CaptureMode         string            `json:"captureMode,omitempty"`
	CorrelationID       uuid.UUID         `json:"correlationId,omitempty"`
}

// ChargeResult is the normalized result returned by a provider adapter.
type ChargeResult struct {
	PSPTransactionID string          `json:"pspTransactionId,omitempty"`
	PSPReference     string          `json:"pspReference,omitempty"`
	RawRequest       json.RawMessage `json:"rawRequest,omitempty"`
	RawResponse      json.RawMessage `json:"rawResponse,omitempty"`
	NextStatus       PaymentStatus   `json:"nextStatus"`
	RequiresAction   bool            `json:"requiresAction"`
	FailureCode      string          `json:"failureCode,omitempty"`
	FailureMessage   string          `json:"failureMessage,omitempty"`
	FailureCategory  string          `json:"failureCategory,omitempty"`
	AuthorisedAmount decimal.Decimal `json:"authorisedAmount,omitempty"`
	CapturedAmount   decimal.Decimal `json:"capturedAmount,omitempty"`
	ProcessingFee    decimal.Decimal `json:"processingFee,omitempty"`
}

// PaymentGateway is the adapter contract implemented by each PSP integration.
type PaymentGateway interface {
	PSP() Psp
	Descriptor() GatewayDescriptor
	Supports(method PaymentMethodType, currency PaymentCurrency) bool
	Charge(ctx context.Context, request ChargeRequest) (ChargeResult, error)
}

// AssertGateway is a small helper used by provider constructors to fail fast
// when an adapter does not satisfy the bridge contract.
func AssertGateway(gateway PaymentGateway) PaymentGateway {
	if gateway == nil {
		panic(fmt.Errorf("payment gateway cannot be nil"))
	}
	return gateway
}

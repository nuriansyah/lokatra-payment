package service

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func uuidOrPanic() uuid.UUID {
	id, _ := uuid.NewV7()
	return id
}

func timeOrPanic(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func TestCreatePaymentIntentRequestRejectsMalformedJSON(t *testing.T) {
	req := dto.CreatePaymentIntentRequest{
		ActorID:        uuidOrPanic(),
		MerchantID:     uuidOrPanic(),
		SourceID:       uuidOrPanic(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.NewFromInt(10000),
		Currency:       "IDR",
		IdempotencyKey: "intent-key",
		SourceSnapshot: []byte(`{invalid json}`),
	}
	err := req.Validate()
	require.Error(t, err)
	require.ErrorContains(t, err, "sourceSnapshot must contain valid JSON")
}

func TestCreatePaymentIntentRequestRejectsInvalidActorID(t *testing.T) {
	req := dto.CreatePaymentIntentRequest{
		MerchantID:     uuidOrPanic(),
		SourceID:       uuidOrPanic(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.NewFromInt(10000),
		Currency:       "IDR",
		IdempotencyKey: "intent-key",
	}
	err := req.Validate()
	require.Error(t, err)
	require.Equal(t, "ERR-VAL-01", failure.GetErrorCode(err))
}

func TestCreatePaymentIntentRequestRejectsZeroAmount(t *testing.T) {
	req := dto.CreatePaymentIntentRequest{
		ActorID:        uuidOrPanic(),
		MerchantID:     uuidOrPanic(),
		SourceID:       uuidOrPanic(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.Zero,
		Currency:       "IDR",
		IdempotencyKey: "intent-key",
	}
	err := req.Validate()
	require.Error(t, err)
}

func TestCreatePaymentIntentRequestRejectsMissingCurrency(t *testing.T) {
	req := dto.CreatePaymentIntentRequest{
		ActorID:        uuidOrPanic(),
		MerchantID:     uuidOrPanic(),
		SourceID:       uuidOrPanic(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.NewFromInt(10000),
		IdempotencyKey: "intent-key",
	}
	err := req.Validate()
	require.Error(t, err)
}

func TestCreatePaymentIntentRequestRejectsPastExpiry(t *testing.T) {
	past := timeOrPanic("2020-01-01T00:00:00Z")
	req := dto.CreatePaymentIntentRequest{
		ActorID:        uuidOrPanic(),
		MerchantID:     uuidOrPanic(),
		SourceID:       uuidOrPanic(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.NewFromInt(10000),
		Currency:       "IDR",
		IdempotencyKey: "intent-key",
		ExpiresAt:      &past,
	}
	err := req.Validate()
	require.Error(t, err)
	require.ErrorContains(t, err, "expiresAt must be in the future")
}

func TestCreateRefundRequestRejectsMissingActor(t *testing.T) {
	req := dto.CreateRefundRequest{
		PaymentIntentID: uuidOrPanic(),
		Amount:          decimal.NewFromInt(1000),
		IdempotencyKey:  "refund-key",
	}
	err := req.Validate()
	require.Error(t, err)
	require.Equal(t, "ERR-VAL-01", failure.GetErrorCode(err))
}

func TestCreateRefundRequestRejectsZeroAmount(t *testing.T) {
	req := dto.CreateRefundRequest{
		ActorID:         uuidOrPanic(),
		PaymentIntentID: uuidOrPanic(),
		Amount:          decimal.Zero,
		IdempotencyKey:  "refund-key",
	}
	err := req.Validate()
	require.Error(t, err)
	require.Equal(t, "ERR-VAL-06", failure.GetErrorCode(err))
}

func TestOpenCashSessionRequestRejectsNegativeAmount(t *testing.T) {
	req := dto.OpenCashSessionRequest{
		ActorID:            uuidOrPanic(),
		MerchantID:         uuidOrPanic(),
		CollectorID:        uuidOrPanic(),
		OpeningFloatAmount: decimal.NewFromInt(-100),
		Currency:           "IDR",
	}
	err := req.Validate()
	require.Error(t, err)
}

func TestActionCommandRejectsZeroActor(t *testing.T) {
	req := dto.ActionCommand{}
	err := req.Validate()
	require.Error(t, err)
	require.Equal(t, "ERR-VAL-01", failure.GetErrorCode(err))
}

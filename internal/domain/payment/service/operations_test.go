package service

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentIntentRejectsMalformedJSONFields(t *testing.T) {
	service := &ServiceImpl{}
	_, err := service.CreatePaymentIntent(context.Background(), dto.CreatePaymentIntentRequest{
		ActorID:        mustUUID(),
		MerchantID:     mustUUID(),
		SourceID:       mustUUID(),
		SourceService:  "orders",
		SourceType:     "order",
		Amount:         decimal.NewFromInt(10000),
		Currency:       "IDR",
		IdempotencyKey: "intent-key",
		Metadata:       json.RawMessage(`{"invalid"`),
	})
	require.ErrorContains(t, err, "metadata must contain valid JSON")
}

func TestCreateRefundRejectsMalformedMetadata(t *testing.T) {
	service := &ServiceImpl{}
	_, err := service.CreateRefund(context.Background(), dto.CreateRefundRequest{
		ActorID:         mustUUID(),
		PaymentIntentID: mustUUID(),
		Amount:          decimal.NewFromInt(1000),
		IdempotencyKey:  "refund-key",
		Metadata:        json.RawMessage(`[`),
	})
	require.ErrorContains(t, err, "metadata must contain valid JSON")
}

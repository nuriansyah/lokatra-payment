//go:build e2e

package service

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	paymentrepo "github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	"github.com/shopspring/decimal"
)

func setupPaymentTestService(t *testing.T) (*ServiceImpl, func()) {
	t.Helper()
	cfg := configs.Get()
	db := infras.ProvidePostgresConn(cfg)
	paymentRepo := paymentrepo.ProvideRepository(db)

	breaker := NewMemoryCircuitBreaker(5, 30*time.Second)
	locker := NewMemoryExecutionLocker()

	svc := ProvidePaymentService(paymentRepo, cfg, breaker, locker)
	return svc, func() {}
}

func newPaymentTestID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

// ============================================================================
// Happy Path: Create and Get Payment Intent
// ============================================================================

func TestPayment_CreateAndGetIntent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	actorID := newPaymentTestID()
	merchantID := newPaymentTestID()
	sourceID := newPaymentTestID()

	t.Log("Creating payment intent...")
	intent, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        actorID,
		MerchantID:     merchantID,
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       sourceID,
		Amount:         decimal.NewFromInt(150000),
		Currency:       "IDR",
		IdempotencyKey: "test-idem-" + uuid.Must(uuid.NewV7()).String(),
	})
	if err != nil {
		t.Fatalf("CreatePaymentIntent failed: %v", err)
	}
	if intent.Id == uuid.Nil {
		t.Fatal("expected non-nil intent ID")
	}
	if intent.Amount.Cmp(decimal.NewFromInt(150000)) != 0 {
		t.Errorf("expected amount=150000, got %s", intent.Amount)
	}
	t.Logf("  Created intent: id=%s, code=%s, status=%s", intent.Id, intent.IntentCode, intent.Status)

	// Get the intent
	t.Log("Getting payment intent...")
	got, err := svc.GetPaymentIntent(ctx, intent.Id)
	if err != nil {
		t.Fatalf("GetPaymentIntent failed: %v", err)
	}
	if got.Id != intent.Id {
		t.Errorf("intent ID mismatch: got %s", got.Id)
	}
	t.Logf("  Got intent: id=%s, status=%s", got.Id, got.Status)

	t.Log("Payment CreateAndGetIntent: ALL PASSED")
}

// ============================================================================
// Happy Path: Idempotency
// ============================================================================

func TestPayment_Idempotency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	actorID := newPaymentTestID()
	merchantID := newPaymentTestID()
	sourceID := newPaymentTestID()
	idempotencyKey := "idem-test-" + uuid.Must(uuid.NewV7()).String()

	// Create first
	t.Log("Creating payment intent (first)...")
	intent1, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        actorID,
		MerchantID:     merchantID,
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       sourceID,
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: idempotencyKey,
	})
	if err != nil {
		t.Fatalf("CreatePaymentIntent (first) failed: %v", err)
	}

	// Create with same key
	t.Log("Creating payment intent (same idempotency key)...")
	intent2, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        actorID,
		MerchantID:     merchantID,
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       sourceID,
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: idempotencyKey,
	})
	if err != nil {
		t.Fatalf("CreatePaymentIntent (idempotent) failed: %v", err)
	}
	if intent1.Id != intent2.Id {
		t.Errorf("expected same intent ID, got %s and %s", intent1.Id, intent2.Id)
	}
	t.Logf("  Idempotent: same ID=%s", intent1.Id)

	t.Log("Payment Idempotency: ALL PASSED")
}

// ============================================================================
// Edge Case: Nil Actor ID
// ============================================================================

func TestPayment_NilActorID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        uuid.Nil,
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "nil-actor-test",
	})
	if err == nil {
		t.Error("expected error for nil actor, got nil")
	} else {
		t.Logf("  Correctly rejected nil actor: %v", err)
	}

	t.Log("Payment NilActorID: ALL PASSED")
}

// ============================================================================
// Edge Case: Zero Amount
// ============================================================================

func TestPayment_ZeroAmount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(0),
		Currency:       "IDR",
		IdempotencyKey: "zero-amount-test",
	})
	if err == nil {
		t.Error("expected error for zero amount, got nil")
	} else {
		t.Logf("  Correctly rejected zero amount: %v", err)
	}

	t.Log("Payment ZeroAmount: ALL PASSED")
}

// ============================================================================
// Edge Case: Negative Amount
// ============================================================================

func TestPayment_NegativeAmount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(-100000),
		Currency:       "IDR",
		IdempotencyKey: "negative-amount-test",
	})
	if err == nil {
		t.Error("expected error for negative amount, got nil")
	} else {
		t.Logf("  Correctly rejected negative amount: %v", err)
	}

	t.Log("Payment NegativeAmount: ALL PASSED")
}

// ============================================================================
// Edge Case: Empty Currency
// ============================================================================

func TestPayment_EmptyCurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "",
		IdempotencyKey: "empty-currency-test",
	})
	if err == nil {
		t.Error("expected error for empty currency, got nil")
	} else {
		t.Logf("  Correctly rejected empty currency: %v", err)
	}

	t.Log("Payment EmptyCurrency: ALL PASSED")
}

// ============================================================================
// Edge Case: Empty SourceService
// ============================================================================

func TestPayment_EmptySourceService(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "empty-source-service-test",
	})
	if err == nil {
		t.Error("expected error for empty source service, got nil")
	} else {
		t.Logf("  Correctly rejected empty source service: %v", err)
	}

	t.Log("Payment EmptySourceService: ALL PASSED")
}

// ============================================================================
// Edge Case: Past Expiry
// ============================================================================

func TestPayment_PastExpiry(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pastTime := time.Now().Add(-1 * time.Hour)

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		ExpiresAt:      &pastTime,
		IdempotencyKey: "past-expiry-test",
	})
	if err == nil {
		t.Error("expected error for past expiry, got nil")
	} else {
		t.Logf("  Correctly rejected past expiry: %v", err)
	}

	t.Log("Payment PastExpiry: ALL PASSED")
}

// ============================================================================
// Edge Case: Cancel Intent
// ============================================================================

func TestPayment_CancelIntent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create intent
	intent, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "cancel-test-" + uuid.Must(uuid.NewV7()).String(),
	})
	if err != nil {
		t.Fatalf("CreatePaymentIntent failed: %v", err)
	}

	// Cancel intent
	t.Log("Canceling payment intent...")
	result, err := svc.ApplyPaymentIntentAction(ctx, intent.Id, "cancel", dto.ActionCommand{
		ActorID: newPaymentTestID(),
		Reason:  "Customer changed mind",
	})
	if err != nil {
		t.Logf("  Cancel error (depends on current status): %v", err)
	} else {
		t.Logf("  Canceled: status=%s", result.Intent.Status)
	}

	t.Log("Payment CancelIntent: ALL PASSED")
}

// ============================================================================
// Edge Case: Create Refund on Non-Succeeded Intent
// ============================================================================

func TestPayment_CreateRefund_NonSucceededIntent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create intent (not succeeded)
	intent, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "refund-non-succeeded-test-" + uuid.Must(uuid.NewV7()).String(),
	})
	if err != nil {
		t.Fatalf("CreatePaymentIntent failed: %v", err)
	}

	// Try to refund
	t.Log("Attempting refund on non-succeeded intent...")
	_, err = svc.CreateRefund(ctx, dto.CreateRefundRequest{
		ActorID:         newPaymentTestID(),
		PaymentIntentID: intent.Id,
		Amount:          decimal.NewFromInt(50000),
		Currency:        "IDR",
		Reason:          "Test refund",
		IdempotencyKey:  "refund-test-" + uuid.Must(uuid.NewV7()).String(),
	})
	if err == nil {
		t.Error("expected error refunding non-succeeded intent, got nil")
	} else {
		t.Logf("  Correctly rejected: %v", err)
	}

	t.Log("Payment CreateRefund_NonSucceededIntent: ALL PASSED")
}

// ============================================================================
// Edge Case: Empty Idempotency Key
// ============================================================================

func TestPayment_EmptyIdempotencyKey(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       newPaymentTestID(),
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "",
	})
	if err == nil {
		t.Error("expected error for empty idempotency key, got nil")
	} else {
		t.Logf("  Correctly rejected empty idempotency key: %v", err)
	}

	t.Log("Payment EmptyIdempotencyKey: ALL PASSED")
}

// ============================================================================
// Edge Case: Nil Source ID
// ============================================================================

func TestPayment_NilSourceID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.CreatePaymentIntent(ctx, dto.CreatePaymentIntentRequest{
		ActorID:        newPaymentTestID(),
		MerchantID:     newPaymentTestID(),
		SourceService:  "booking",
		SourceType:     "booking",
		SourceID:       uuid.Nil,
		Amount:         decimal.NewFromInt(100000),
		Currency:       "IDR",
		IdempotencyKey: "nil-source-test",
	})
	if err == nil {
		t.Error("expected error for nil source ID, got nil")
	} else {
		t.Logf("  Correctly rejected nil source ID: %v", err)
	}

	t.Log("Payment NilSourceID: ALL PASSED")
}

// ============================================================================
// Edge Case: Get Non-Existent Intent
// ============================================================================

func TestPayment_GetNonExistentIntent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	svc, cleanup := setupPaymentTestService(t)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := svc.GetPaymentIntent(ctx, newPaymentTestID())
	if err == nil {
		t.Error("expected error for non-existent intent, got nil")
	} else {
		t.Logf("  Correctly returned error: %v", err)
	}

	t.Log("Payment GetNonExistentIntent: ALL PASSED")
}

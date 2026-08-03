package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/service"
	"github.com/rs/zerolog"
)

// TestLiveCreatePayment makes a REAL call to Xendit's API.
// Run with:
//
//	go test -v -run TestLiveCreatePayment -tags=live ./external/paymentgateway/xendit/service/
//
// Prerequisites:
//   - XENDIT_API_KEY env var must be set (or .env must have EXTERNALS.PROVIDERS.XENDIT.API_KEY)
//   - Xendit API key must be a VALID development key
//
// The payment will be created in PENDING state. You can then:
//   1. Check your Xendit Dashboard → Payment Requests to see it
//   2. Cancel it via the API if needed
func TestLiveCreatePayment(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live test in short mode")
	}

	apiKey := os.Getenv("XENDIT_API_KEY")
	if apiKey == "" {
		t.Skip("XENDIT_API_KEY not set, skipping live test")
	}

	cfg := pg.ProviderConfig{
		Code:            pg.ProviderXendit,
		BaseURL:         "https://api.xendit.co",
		APIKey:          apiKey,
		DefaultCurrency: "IDR",
		HTTP: pg.HTTPClientConfig{
			Timeout:    30 * time.Second,
			RetryCount: 1,
			EnableDebug: true,
		},
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	svc := service.ProvideServiceWithLogger(cfg, logger)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	orderID := fmt.Sprintf("LOKATRA-SMOKE-%d", time.Now().UnixMilli())
	idempotencyKey := fmt.Sprintf("idem-%d", time.Now().UnixMilli())

	t.Logf("Order ID:      %s", orderID)
	t.Logf("Idempotency:   %s", idempotencyKey)
	t.Logf("Base URL:      %s", cfg.BaseURL)
	t.Log("────────────────────────────────────────")

	// Debug: log the exact payload being sent
	payload := model.CreatePaymentRequest{
		ReferenceID: orderID,
		Currency:    "IDR",
		Amount:      10000,
		Country:     "ID",
		Description: "Smoke test - Lokatra Payment Orchestration",
		PaymentMethod: model.PaymentMethodRequest{
			Type:        "VIRTUAL_ACCOUNT",
			Reusability: "ONE_TIME_USE",
			ReferenceID: orderID,
			VirtualAccount: &model.VirtualAccount{
				ChannelCode:       "BCA",
				ChannelProperties: map[string]any{"customer_name": "Smoke Test User"},
			},
		},
		Customer: &model.Customer{
			ReferenceID:  orderID,
			Type:         "INDIVIDUAL",
			Email:        "smoke-test@lokatra.dev",
			IndividualDetail: &model.IndividualDetail{GivenNames: "Smoke Test"},
		},
	}
	payloadJSON, _ := json.MarshalIndent(payload, "", "  ")
	t.Logf("Request Payload:\n%s", string(payloadJSON))
	t.Log("────────────────────────────────────────")

	resp, err := svc.CreatePayment(ctx, pg.CreatePaymentRequest{
		OrderID:        orderID,
		Amount:         pg.Money{Amount: "10000", Currency: "IDR"},
		Method:         pg.PaymentMethodVirtualAccount,
		ChannelCode:    "bca_va",
		Description:    "Smoke test - Lokatra Payment Orchestration",
		IdempotencyKey: idempotencyKey,
		Customer: pg.Customer{
			Name:  "Smoke Test User",
			Email: "smoke-test@lokatra.dev",
		},
	})

	t.Log("────────────────────────────────────────")

	if err != nil {
		t.Fatalf("CreatePayment failed: %v", err)
	}

	t.Logf("Provider Code:         %s", resp.ProviderCode)
	t.Logf("Provider Payment ID:   %s", resp.ProviderPaymentID)
	t.Logf("Provider Reference:    %s", resp.ProviderReference)
	t.Logf("Order ID:              %s", resp.OrderID)
	t.Logf("Status:                %s", resp.Status)

	raw, _ := json.MarshalIndent(resp.Raw, "", "  ")
	t.Logf("Raw Response:\n%s", string(raw))

	if len(resp.Instructions) > 0 {
		t.Log("────────────────────────────────────────")
		t.Log("Payment Instructions:")
		for i, ins := range resp.Instructions {
			t.Logf("  [%d] Type: %s", i, ins.Type)
			if ins.QRString != "" {
				t.Logf("      QR String: %s", truncate(ins.QRString, 80))
			}
			if ins.AccountNumber != "" {
				t.Logf("      Account:   %s", ins.AccountNumber)
			}
			if ins.CheckoutURL != "" {
				t.Logf("      URL:       %s", ins.CheckoutURL)
			}
			if ins.DeeplinkURL != "" {
				t.Logf("      Deeplink:  %s", ins.DeeplinkURL)
			}
			if ins.ExpiresAt != nil {
				t.Logf("      Expires:   %s", ins.ExpiresAt.Format(time.RFC3339))
			}
		}
	}

	t.Log("────────────────────────────────────────")
	t.Log("✅ Payment created successfully!")
	t.Log("   → Check your Xendit Dashboard:")
	t.Log("     https://dashboard.xendit.co/payments/payment-requests")
	t.Logf("     Search for Order ID: %s", orderID)
}

// TestLiveGetPaymentStatus fetches the status of a payment request by ID.
// Run with:
//
//	go test -v -run TestLiveGetPaymentStatus -tags=live ./external/paymentgateway/xendit/service/
func TestLiveGetPaymentStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live test in short mode")
	}

	apiKey := os.Getenv("XENDIT_API_KEY")
	paymentID := os.Getenv("XENDIT_PAYMENT_ID")
	if apiKey == "" {
		t.Skip("XENDIT_API_KEY not set, skipping live test")
	}
	if paymentID == "" {
		t.Skip("XENDIT_PAYMENT_ID not set, skipping live test")
	}

	cfg := pg.ProviderConfig{
		Code:            pg.ProviderXendit,
		BaseURL:         "https://api.xendit.co",
		APIKey:          apiKey,
		DefaultCurrency: "IDR",
		HTTP: pg.HTTPClientConfig{
			Timeout: 30 * time.Second,
		},
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	svc := service.ProvideServiceWithLogger(cfg, logger)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := svc.GetPaymentStatus(ctx, pg.GetPaymentStatusRequest{
		ProviderTransactionID: paymentID,
	})
	if err != nil {
		t.Fatalf("GetPaymentStatus failed: %v", err)
	}

	t.Logf("Payment ID:    %s", resp.ProviderTransactionID)
	t.Logf("Reference:     %s", resp.ProviderReference)
	t.Logf("Order ID:      %s", resp.OrderID)
	t.Logf("Status:        %s", resp.Status)
	t.Logf("Amount:        %s %s", resp.Amount.Amount, resp.Amount.Currency)
}

// TestLiveCancelPayment cancels a payment request by ID.
// Run with:
//
//	go test -v -run TestLiveCancelPayment -tags=live ./external/paymentgateway/xendit/service/
func TestLiveCancelPayment(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live test in short mode")
	}

	apiKey := os.Getenv("XENDIT_API_KEY")
	paymentID := os.Getenv("XENDIT_PAYMENT_ID")
	if apiKey == "" || paymentID == "" {
		t.Skip("XENDIT_API_KEY or XENDIT_PAYMENT_ID not set, skipping live test")
	}

	cfg := pg.ProviderConfig{
		Code:            pg.ProviderXendit,
		BaseURL:         "https://api.xendit.co",
		APIKey:          apiKey,
		DefaultCurrency: "IDR",
		HTTP: pg.HTTPClientConfig{
			Timeout: 30 * time.Second,
		},
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	svc := service.ProvideServiceWithLogger(cfg, logger)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := svc.CancelPayment(ctx, pg.CancelPaymentRequest{
		ProviderTransactionID: paymentID,
		IdempotencyKey:        fmt.Sprintf("cancel-%d", time.Now().UnixMilli()),
	})
	if err != nil {
		t.Fatalf("CancelPayment failed: %v", err)
	}

	t.Logf("Order ID:  %s", resp.OrderID)
	t.Logf("Status:    %s", resp.Status)
	t.Log("✅ Payment canceled successfully!")
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

func init() {
	// Load .env file for live tests
	if _, err := os.Stat(".env"); err == nil {
		data, err := os.ReadFile(".env")
		if err == nil {
			for _, line := range strings.Split(string(data), "\n") {
				line = strings.TrimSpace(line)
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					val := strings.TrimSpace(parts[1])
					os.Setenv(key, val)
				}
			}
		}
	}
}

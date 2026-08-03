package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
)

func TestWebhookProviderRecognisesAllSupportedProviders(t *testing.T) {
	tests := []struct {
		input    string
		expected pg.ProviderCode
		valid    bool
	}{
		{"xendit", pg.ProviderXendit, true},
		{"XENDIT", pg.ProviderXendit, true},
		{" midtrans ", pg.ProviderMidtrans, true},
		{"durianpay", pg.ProviderDurianpay, true},
		{"finpay", pg.ProviderFinpay, true},
		{"ipaymu", pg.ProviderIpaymu, true},
		{"unknown", "", false},
		{"", "", false},
		{" stripe ", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			code, ok := webhookProvider(tt.input)
			if tt.valid {
				require.True(t, ok)
				require.Equal(t, tt.expected, code)
			} else {
				require.False(t, ok)
				require.Empty(t, code)
			}
		})
	}
}

func TestWebhookSignatureAlgorithmMapsProvidersCorrectly(t *testing.T) {
	require.Equal(t, "SHA512", webhookSignatureAlgorithm(pg.ProviderMidtrans))
	require.Equal(t, "HMAC_SHA256", webhookSignatureAlgorithm(pg.ProviderDurianpay))
	require.Equal(t, "HMAC_SHA256", webhookSignatureAlgorithm(pg.ProviderFinpay))
	require.Equal(t, "HMAC_SHA256", webhookSignatureAlgorithm(pg.ProviderIpaymu))
	require.Equal(t, "TOKEN_OR_HMAC_SHA256", webhookSignatureAlgorithm(pg.ProviderXendit))
	require.Equal(t, "TOKEN_OR_HMAC_SHA256", webhookSignatureAlgorithm("unknown"))
}

type mockWebhookGateway struct {
	pg.PaymentGateway
	code      pg.ProviderCode
	verify    func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error)
	normalize func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error)
}

func (m *mockWebhookGateway) ProviderCode() pg.ProviderCode { return m.code }

func (m *mockWebhookGateway) VerifyWebhook(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
	if m.verify != nil {
		return m.verify(ctx, req)
	}
	return pg.VerifyWebhookResult{SignatureValid: true}, nil
}

func (m *mockWebhookGateway) NormalizeWebhook(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
	if m.normalize != nil {
		return m.normalize(ctx, req)
	}
	return pg.CanonicalPaymentEvent{}, nil
}

func (m *mockWebhookGateway) Capabilities(ctx context.Context, req pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: m.code}, nil
}

type mockWebhookRepo struct {
	repository.Repository
	createErr    error
	existingByID map[string]paymentmodel.ProviderWebhookEvents
}

func newMockWebhookRepo() *mockWebhookRepo {
	return &mockWebhookRepo{existingByID: make(map[string]paymentmodel.ProviderWebhookEvents)}
}

func (m *mockWebhookRepo) CreateProviderWebhookEvents(ctx context.Context, record *paymentmodel.ProviderWebhookEvents, fields ...repository.ProviderWebhookEventsField) error {
	return m.createErr
}

func (m *mockWebhookRepo) ResolveProviderWebhookEventsByFilter(ctx context.Context, filter paymentmodel.Filter) ([]paymentmodel.ProviderWebhookEventsFilterResult, error) {
	for _, f := range filter.FilterFields {
		if string(f.Field) == string(paymentmodel.ProviderWebhookEventsDBFieldName.EventId) {
			val, _ := f.Value.(string)
			if existing, ok := m.existingByID[val]; ok {
				return []paymentmodel.ProviderWebhookEventsFilterResult{
					{ProviderWebhookEvents: existing},
				}, nil
			}
		}
	}
	return nil, nil
}

func buildWebhookTestService(gateway pg.PaymentGateway, repo repository.Repository, webhookSecret string) *ServiceImpl {
	registry := pg.NewRegistry()
	_ = registry.Register(gateway)

	accountIDs := map[pg.ProviderCode]uuid.UUID{
		pg.ProviderXendit:    uuid.Must(uuid.NewV7()),
		pg.ProviderMidtrans:  uuid.Must(uuid.NewV7()),
		pg.ProviderDurianpay: uuid.Must(uuid.NewV7()),
		pg.ProviderFinpay:    uuid.Must(uuid.NewV7()),
		pg.ProviderIpaymu:    uuid.Must(uuid.NewV7()),
	}
	webhookConfigured := map[pg.ProviderCode]bool{
		pg.ProviderXendit: webhookSecret != "",
	}

	return &ServiceImpl{
		paymentRepo:        repo,
		gatewayRegistry:    registry,
		providerAccountIDs: accountIDs,
		webhookConfigured:  webhookConfigured,
		routingEngine:      NewRoutingEngine(registry, accountIDs, NewMemoryCircuitBreaker(3, time.Minute), RoutingConfig{DefaultProviders: []pg.ProviderCode{pg.ProviderXendit}, MaxAttempts: 1}, repo, nil),
		executionLocker:    NewMemoryExecutionLocker(),
		logger:             zerolog.Nop(),
	}
}

func TestHandleWebhookRejectsUnsupportedProvider(t *testing.T) {
	svc := buildWebhookTestService(&mockWebhookGateway{code: pg.ProviderXendit}, newMockWebhookRepo(), "")

	_, err := svc.HandleWebhook(context.Background(), "stripe", http.Header{}, []byte(`{}`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported")
}

func TestHandleWebhookRejectsWhenAccountNotConfigured(t *testing.T) {
	registry := pg.NewRegistry()
	_ = registry.Register(&mockWebhookGateway{code: pg.ProviderXendit})

	svc := &ServiceImpl{
		paymentRepo:        newMockWebhookRepo(),
		gatewayRegistry:    registry,
		providerAccountIDs: map[pg.ProviderCode]uuid.UUID{pg.ProviderXendit: uuid.Nil},
		webhookConfigured:  map[pg.ProviderCode]bool{},
		executionLocker:    NewMemoryExecutionLocker(),
		logger:             zerolog.Nop(),
	}

	_, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, []byte(`{}`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "not configured")
}

func TestHandleWebhookRejectsWhenSignatureVerificationFails(t *testing.T) {
	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			return pg.VerifyWebhookResult{SignatureValid: false, Reason: "bad signature"},
				fmt.Errorf("signature mismatch")
		},
	}
	svc := buildWebhookTestService(gateway, newMockWebhookRepo(), "test-secret")

	_, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, []byte(`{"id":"evt-1"}`))
	require.Error(t, err)
	require.Contains(t, err.Error(), "signature")
}

func TestHandleWebhookAcceptsValidHMACSignature(t *testing.T) {
	secret := "my-webhook-secret"
	body := []byte(`{"id":"evt-valid-hmac","event":"payment.succeeded","data":{"id":"pr-123","reference_id":"order-001","status":"SUCCEEDED","amount":100000,"currency":"IDR"}}`)
	sig := pg.HMACSHA256Hex(secret, body)

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			expected := pg.HMACSHA256Hex(secret, req.RawBody)
			if sig == "" || expected != sig {
				return pg.VerifyWebhookResult{SignatureValid: false, Reason: "mismatch"}, nil
			}
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt-valid-hmac"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt-valid-hmac",
				EventType:     pg.EventPaymentSucceeded,
				PaymentStatus: pg.PaymentStatusSucceeded,
				OrderID:       "order-001",
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, newMockWebhookRepo(), secret)

	headers := http.Header{"X-Callback-Signature": []string{sig}}
	receipt, err := svc.HandleWebhook(context.Background(), "xendit", headers, body)
	require.NoError(t, err)
	require.True(t, receipt.SignatureValid)
	require.Equal(t, "evt-valid-hmac", receipt.EventID)
	require.Equal(t, string(pg.EventPaymentSucceeded), receipt.EventType)
	require.Equal(t, string(pg.PaymentStatusSucceeded), receipt.PaymentStatus)
	require.Equal(t, "order-001", receipt.OrderID)
}

func TestHandleWebhookDeduplicatesByEventID(t *testing.T) {
	body := []byte(`{"id":"evt-dup","event":"payment.succeeded","data":{"id":"pr-456","reference_id":"order-002","status":"SUCCEEDED","amount":50000,"currency":"IDR"}}`)
	repo := newMockWebhookRepo()

	existingTime := time.Now().UTC().Add(-time.Hour)
	repo.existingByID["evt-dup"] = paymentmodel.ProviderWebhookEvents{
		Id:         uuid.Must(uuid.NewV7()),
		EventId:    null.StringFrom("evt-dup"),
		ReceivedAt: existingTime,
	}

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt-dup"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt-dup",
				EventType:     pg.EventPaymentSucceeded,
				PaymentStatus: pg.PaymentStatusSucceeded,
				OrderID:       "order-002",
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, repo, "")

	receipt, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, body)
	require.NoError(t, err)
	require.Equal(t, "evt-dup", receipt.EventID)
	require.Equal(t, existingTime, receipt.ReceivedAt)
}

func TestHandleWebhookPersistsNewEvent(t *testing.T) {
	body := []byte(`{"id":"evt-new","event":"payment.pending","data":{"id":"pr-789","reference_id":"order-003","status":"PENDING","amount":25000,"currency":"IDR"}}`)
	repo := newMockWebhookRepo()

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt-new"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt-new",
				EventType:     pg.EventPaymentCreated,
				PaymentStatus: pg.PaymentStatusPending,
				OrderID:       "order-003",
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, repo, "")

	receipt, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, body)
	require.NoError(t, err)
	require.Equal(t, "evt-new", receipt.EventID)
	require.True(t, receipt.SignatureValid)
	require.Equal(t, "order-003", receipt.OrderID)
	require.Equal(t, string(pg.PaymentStatusPending), receipt.PaymentStatus)
}

func TestHandleWebhookReturnsErrorWhenPersistFails(t *testing.T) {
	body := []byte(`{"id":"evt-fail-persist","event":"payment.succeeded","data":{"id":"pr-111","reference_id":"order-004","status":"SUCCEEDED","amount":10000,"currency":"IDR"}}`)
	repo := newMockWebhookRepo()
	repo.createErr = fmt.Errorf("database connection refused")

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt-fail-persist"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt-fail-persist",
				EventType:     pg.EventPaymentSucceeded,
				PaymentStatus: pg.PaymentStatusSucceeded,
				OrderID:       "order-004",
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, repo, "")

	_, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, body)
	require.Error(t, err)
	require.Contains(t, err.Error(), "database connection refused")
}

func TestHandleWebhookAutoPassesWhenNoSecretConfigured(t *testing.T) {
	body := []byte(`{"id":"evt-no-secret","event":"payment.succeeded","data":{"id":"pr-222","reference_id":"order-005","status":"SUCCEEDED","amount":75000,"currency":"IDR"}}`)
	repo := newMockWebhookRepo()

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt-no-secret"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt-no-secret",
				EventType:     pg.EventPaymentSucceeded,
				PaymentStatus: pg.PaymentStatusSucceeded,
				OrderID:       "order-005",
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, repo, "")

	receipt, err := svc.HandleWebhook(context.Background(), "xendit", http.Header{}, body)
	require.NoError(t, err)
	require.True(t, receipt.SignatureValid)
	require.Equal(t, "evt-no-secret", receipt.EventID)
}

func TestHandleWebhookFinpayAndIpaymuRecognized(t *testing.T) {
	for _, providerName := range []string{"finpay", "ipaymu"} {
		t.Run(providerName, func(t *testing.T) {
			registry := pg.NewRegistry()
			gw := &mockWebhookGateway{
				code: pg.ProviderCode(providerName),
				verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
					return pg.VerifyWebhookResult{SignatureValid: true}, nil
				},
				normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
					return pg.CanonicalPaymentEvent{
						ProviderCode:  pg.ProviderCode(providerName),
						EventID:       "evt-fin-ipay",
						EventType:     pg.EventPaymentSucceeded,
						PaymentStatus: pg.PaymentStatusSucceeded,
						OrderID:       "order-fin-ipay",
						Raw:           json.RawMessage(req.RawBody),
					}, nil
				},
			}
			_ = registry.Register(gw)

			accountIDs := map[pg.ProviderCode]uuid.UUID{
				pg.ProviderCode(providerName): uuid.Must(uuid.NewV7()),
			}

			svc := &ServiceImpl{
				paymentRepo:        newMockWebhookRepo(),
				gatewayRegistry:    registry,
				providerAccountIDs: accountIDs,
				webhookConfigured:  map[pg.ProviderCode]bool{},
				routingEngine:      NewRoutingEngine(registry, accountIDs, NewMemoryCircuitBreaker(3, time.Minute), RoutingConfig{DefaultProviders: []pg.ProviderCode{pg.ProviderXendit}, MaxAttempts: 1}, nil, nil),
				executionLocker:    NewMemoryExecutionLocker(),
				logger:             zerolog.Nop(),
			}

			body := []byte(`{"id":"evt-fin-ipay","event":"payment.succeeded","data":{"id":"pr-555","reference_id":"order-fin-ipay","status":"SUCCEEDED","amount":200000,"currency":"IDR"}}`)
			receipt, err := svc.HandleWebhook(context.Background(), providerName, http.Header{}, body)
			require.NoError(t, err)
			require.True(t, receipt.SignatureValid)
			require.Equal(t, "evt-fin-ipay", receipt.EventID)
		})
	}
}

func TestHandleWebhookXenditRealisticHMACPayload(t *testing.T) {
	secret := "xendit-production-webhook-secret"
	body := []byte(`{"id":"evt_5b1f8d50e10422b420771e5f","event":"payment.succeeded","data":{"id":"pr_5b1f8d50e10422b420771e5e","reference_id":"LOKATRA-WEBHOOK-001","status":"SUCCEEDED","amount":150000,"currency":"IDR"},"created_at":"2025-07-15T08:30:00.000Z","updated_at":"2025-07-15T08:31:00.000Z"}`)
	sig := pg.HMACSHA256Hex(secret, body)

	gateway := &mockWebhookGateway{
		code: pg.ProviderXendit,
		verify: func(ctx context.Context, req pg.VerifyWebhookRequest) (pg.VerifyWebhookResult, error) {
			expected := pg.HMACSHA256Hex(secret, req.RawBody)
			if expected != sig {
				return pg.VerifyWebhookResult{SignatureValid: false, Reason: "HMAC mismatch"}, nil
			}
			return pg.VerifyWebhookResult{SignatureValid: true, EventID: "evt_5b1f8d50e10422b420771e5f"}, nil
		},
		normalize: func(ctx context.Context, req pg.NormalizeWebhookRequest) (pg.CanonicalPaymentEvent, error) {
			return pg.CanonicalPaymentEvent{
				ProviderCode:  pg.ProviderXendit,
				EventID:       "evt_5b1f8d50e10422b420771e5f",
				EventType:     pg.EventPaymentSucceeded,
				PaymentStatus: pg.PaymentStatusSucceeded,
				OrderID:       "LOKATRA-WEBHOOK-001",
				Amount:        pg.Money{Amount: "150000.00", Currency: "IDR"},
				Raw:           json.RawMessage(req.RawBody),
			}, nil
		},
	}
	svc := buildWebhookTestService(gateway, newMockWebhookRepo(), secret)

	headers := http.Header{"X-Callback-Signature": []string{sig}}
	receipt, err := svc.HandleWebhook(context.Background(), "xendit", headers, body)
	require.NoError(t, err)
	require.True(t, receipt.SignatureValid)
	require.Equal(t, "evt_5b1f8d50e10422b420771e5f", receipt.EventID)
	require.Equal(t, "LOKATRA-WEBHOOK-001", receipt.OrderID)
	require.Equal(t, string(pg.PaymentStatusSucceeded), receipt.PaymentStatus)
}

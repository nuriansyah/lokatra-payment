package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyWebhookRequiresConfiguredHMACSignature(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{WebhookSecret: "webhook-secret"})
	body := []byte(`{"id":"evt-1"}`)

	result, err := svc.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: http.Header{},
		RawBody: body,
	})
	require.Error(t, err)
	require.False(t, result.SignatureValid)

	headers := http.Header{}
	headers.Set("X-Callback-Signature", pg.HMACSHA256Hex("webhook-secret", body))
	result, err = svc.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: headers,
		RawBody: body,
	})
	require.NoError(t, err)
	require.True(t, result.SignatureValid)
}

func TestVerifyWebhookCallbackToken(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{WebhookToken: "my-token"})

	result, err := svc.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: http.Header{"X-Callback-Token": []string{"my-token"}},
		RawBody: []byte(`{"id":"evt-1"}`),
	})
	require.NoError(t, err)
	require.True(t, result.SignatureValid)

	result, err = svc.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: http.Header{"X-Callback-Token": []string{"wrong-token"}},
		RawBody: []byte(`{"id":"evt-1"}`),
	})
	require.Error(t, err)
	require.False(t, result.SignatureValid)
}

func TestVerifyWebhookNoSecretConfigured(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	result, err := svc.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: http.Header{},
		RawBody: []byte(`{"id":"evt-1"}`),
	})
	require.NoError(t, err)
	require.True(t, result.SignatureValid)
}

func TestCreatePayment_Success(t *testing.T) {
	mockResp := model.PaymentRequestResponse{
		ID:          "pr-xendit-123",
		ReferenceID: "order-001",
		Status:      "PENDING",
		Currency:    "IDR",
		Amount:      100000,
		PaymentMethod: model.PaymentMethodResponse{
			Type: "QR_CODE",
			QR: &model.QRResponse{
				ChannelCode: "QRIS",
				ChannelProperties: model.QRProperties{
					QRString:  "00020101...",
					ExpiresAt: "2026-08-04T00:00:00Z",
				},
			},
		},
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/payment_requests", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "idem-123", r.Header.Get("Idempotency-Key"))

		var body model.CreatePaymentRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "order-001", body.ReferenceID)
		assert.Equal(t, 100000.0, body.Amount)
		assert.Equal(t, "IDR", body.Currency)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "100000", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
		ChannelCode: "qris",
		Customer: pg.Customer{Name: "John"},
		IdempotencyKey: "idem-123",
	})

	require.NoError(t, err)
	assert.Equal(t, pg.ProviderXendit, resp.ProviderCode)
	assert.Equal(t, "order-001", resp.OrderID)
	assert.Equal(t, "pr-xendit-123", resp.ProviderTransactionID)
	assert.Equal(t, pg.PaymentStatusPending, resp.Status)
	require.Len(t, resp.Instructions, 1)
	assert.Equal(t, "qr_string", resp.Instructions[0].Type)
}

func TestCreatePayment_InvalidAmount(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "0", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid amount")
}

func TestCreatePayment_NegativeAmount(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "-100", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid amount")
}

func TestCreatePayment_WithCustomer(t *testing.T) {
	var receivedBody model.CreatePaymentRequest

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&receivedBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "pr-1", ReferenceID: "ord-1", Status: "PENDING",
			PaymentMethod: model.PaymentMethodResponse{Type: "VIRTUAL_ACCOUNT"},
		})
	}))
	defer server.Close()

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID:    "ord-1",
		Amount:     pg.Money{Amount: "50000", Currency: "IDR"},
		Method:     pg.PaymentMethodVirtualAccount,
		ChannelCode: "bca_va",
		Customer: pg.Customer{
			Name:  "Budi",
			Email: "budi@example.com",
			Phone: "+6281234567890",
		},
	})

	require.NoError(t, err)
	require.NotNil(t, receivedBody.Customer)
	assert.Equal(t, "INDIVIDUAL", receivedBody.Customer.Type)
	assert.Equal(t, "Budi", receivedBody.Customer.IndividualDetail.GivenNames)
	assert.Equal(t, "budi@example.com", receivedBody.Customer.Email)
}

func TestCreatePayment_HTTPError(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
	}))
	defer server.Close()

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "100000", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.Error(t, err)
	var gwErr *pg.GatewayError
	require.ErrorAs(t, err, &gwErr)
	assert.Equal(t, pg.ErrorCodeUnauthorized, gwErr.Code)
}

func TestCreatePayment_ServerError(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"internal"}`)
	}))
	defer server.Close()

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "100000", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.Error(t, err)
	var gwErr *pg.GatewayError
	require.ErrorAs(t, err, &gwErr)
	assert.True(t, gwErr.Retryable)
}

func TestGetPaymentStatus_Success(t *testing.T) {
	mockResp := model.PaymentRequestResponse{
		ID:          "pr-123",
		ReferenceID: "order-001",
		Status:      "SUCCEEDED",
		Currency:    "IDR",
		Amount:      100000,
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/payment_requests/pr-123", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.GetPaymentStatus(context.Background(), pg.GetPaymentStatusRequest{
		ProviderTransactionID: "pr-123",
	})
	require.NoError(t, err)
	assert.Equal(t, pg.PaymentStatusSucceeded, resp.Status)
	assert.Equal(t, "order-001", resp.OrderID)
	assert.Equal(t, "100000.00", resp.Amount.Amount)
}

func TestGetPaymentStatus_NoID(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.GetPaymentStatus(context.Background(), pg.GetPaymentStatusRequest{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "required")
}

func TestGetPaymentStatus_NotFound(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":"not found"}`)
	}))
	defer server.Close()

	_, err := svc.GetPaymentStatus(context.Background(), pg.GetPaymentStatusRequest{
		ProviderTransactionID: "nonexistent",
	})
	require.Error(t, err)
	var gwErr *pg.GatewayError
	require.ErrorAs(t, err, &gwErr)
	assert.Equal(t, pg.ErrorCodeNotFound, gwErr.Code)
}

func TestCancelPayment_Success(t *testing.T) {
	mockResp := model.PaymentRequestResponse{
		ID:          "pr-123",
		ReferenceID: "order-001",
		Status:      "CANCELED",
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Contains(t, r.URL.Path, "/payment_requests/pr-123/cancel")
		assert.Equal(t, "idem-cancel", r.Header.Get("Idempotency-Key"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.CancelPayment(context.Background(), pg.CancelPaymentRequest{
		ProviderTransactionID: "pr-123",
		IdempotencyKey:        "idem-cancel",
	})
	require.NoError(t, err)
	assert.Equal(t, pg.PaymentStatusCanceled, resp.Status)
	assert.Equal(t, "order-001", resp.OrderID)
}

func TestCancelPayment_NoID(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.CancelPayment(context.Background(), pg.CancelPaymentRequest{})
	require.Error(t, err)
}

func TestRefundPayment_Success(t *testing.T) {
	mockResp := model.RefundResponse{
		ID:          "rf-123",
		ReferenceID: "order-001",
		Status:      "SUCCEEDED",
		Amount:      50000,
		Currency:    "IDR",
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/refunds", r.URL.Path)

		var body model.RefundRequest
		json.NewDecoder(r.Body).Decode(&body)
		assert.Equal(t, 50000.0, body.Amount)
		assert.Equal(t, "IDR", body.Currency)
		assert.Equal(t, "rf-001", body.ReferenceID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.RefundPayment(context.Background(), pg.RefundRequest{
		OrderID:               "order-001",
		ProviderTransactionID: "pr-123",
		RefundID:              "rf-001",
		Amount:                pg.Money{Amount: "50000", Currency: "IDR"},
		Reason:                "customer request",
	})
	require.NoError(t, err)
	assert.Equal(t, pg.PaymentStatusSucceeded, resp.Status)
	assert.Equal(t, "rf-123", resp.ProviderRefundID)
}

func TestRefundPayment_InvalidAmount(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.RefundPayment(context.Background(), pg.RefundRequest{
		OrderID:  "order-001",
		RefundID: "rf-001",
		Amount:   pg.Money{Amount: "0", Currency: "IDR"},
	})
	require.Error(t, err)
}

func TestNormalizeWebhook_Pending(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	webhookBody := model.WebhookPayment{
		ID:    "evt-1",
		Event: "payment.pending",
		Data: model.PaymentRequestResponse{
			ID:          "pr-1",
			ReferenceID: "ord-1",
			Status:      "PENDING",
			Amount:      100000,
			Currency:    "IDR",
		},
	}
	body, _ := json.Marshal(webhookBody)

	evt, err := svc.NormalizeWebhook(context.Background(), pg.NormalizeWebhookRequest{
		Headers: http.Header{"webhook-id": []string{"evt-1"}},
		RawBody: body,
	})
	require.NoError(t, err)
	assert.Equal(t, pg.EventPaymentPending, evt.EventType)
	assert.Equal(t, pg.PaymentStatusPending, evt.PaymentStatus)
	assert.Equal(t, "ord-1", evt.OrderID)
	assert.Equal(t, "evt-1", evt.EventID)
}

func TestNormalizeWebhook_InvalidJSON(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.NormalizeWebhook(context.Background(), pg.NormalizeWebhookRequest{
		RawBody: []byte(`{invalid`),
	})
	require.Error(t, err)
}

func TestCapabilities(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{DefaultCurrency: "IDR"})

	resp, err := svc.Capabilities(context.Background(), pg.CapabilitiesRequest{})
	require.NoError(t, err)
	assert.Equal(t, pg.ProviderXendit, resp.ProviderCode)
	assert.NotEmpty(t, resp.Items)

	methods := make(map[pg.PaymentMethod]bool)
	for _, item := range resp.Items {
		methods[item.Method] = true
	}
	assert.True(t, methods[pg.PaymentMethodVirtualAccount])
	assert.True(t, methods[pg.PaymentMethodQRIS])
	assert.True(t, methods[pg.PaymentMethodEWallet])
}

func TestCreatePayment_VirtualAccount(t *testing.T) {
	var receivedBody model.CreatePaymentRequest

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&receivedBody)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "pr-va-1", ReferenceID: "ord-va-1", Status: "PENDING",
			PaymentMethod: model.PaymentMethodResponse{
				Type: "VIRTUAL_ACCOUNT",
				VirtualAccount: &model.VirtualAccountResponse{
					ChannelCode: "BCA",
					ChannelProperties: model.VAProperties{
						VirtualAccountNumber: "1234567890",
						ExpiresAt:            "2026-08-04T00:00:00Z",
					},
				},
			},
		})
	}))
	defer server.Close()

	resp, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID:    "ord-va-1",
		Amount:     pg.Money{Amount: "250000", Currency: "IDR"},
		Method:     pg.PaymentMethodVirtualAccount,
		ChannelCode: "bca_va",
		Customer:   pg.Customer{Name: "Siti"},
	})

	require.NoError(t, err)
	assert.Equal(t, "VIRTUAL_ACCOUNT", receivedBody.PaymentMethod.Type)
	require.Len(t, resp.Instructions, 1)
	assert.Equal(t, "va_number", resp.Instructions[0].Type)
	assert.Equal(t, "1234567890", resp.Instructions[0].AccountNumber)
}

func TestGetPaymentStatus_FallbackPath(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/payment_requests/pr-fallback", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "pr-fallback", ReferenceID: "ord-fb", Status: "PENDING",
			PaymentMethod: model.PaymentMethodResponse{Type: "QR_CODE"},
		})
	}))
	defer server.Close()

	resp, err := svc.GetPaymentStatus(context.Background(), pg.GetPaymentStatusRequest{
		ProviderTransactionID: "pr-fallback",
	})
	require.NoError(t, err)
	assert.Equal(t, "ord-fb", resp.OrderID)
}

func TestRateLimited(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, `{"error":"rate limited"}`)
	}))
	defer server.Close()

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "order-001",
		Amount:  pg.Money{Amount: "100000", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.Error(t, err)
	var gwErr *pg.GatewayError
	require.ErrorAs(t, err, &gwErr)
	assert.Equal(t, pg.ErrorCodeRateLimited, gwErr.Code)
	assert.True(t, gwErr.Retryable)
}

func TestAuthorizationHeaderFormat(t *testing.T) {
	var authHeader string

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "pr-1", ReferenceID: "o-1", Status: "PENDING",
			PaymentMethod: model.PaymentMethodResponse{Type: "QR_CODE"},
		})
	}), pg.ProviderConfig{
		APIKey:  "test-api-key-12345",
		Extra:   map[string]string{},
	})
	defer server.Close()

	_, err := svc.CreatePayment(context.Background(), pg.CreatePaymentRequest{
		OrderID: "o-1",
		Amount:  pg.Money{Amount: "100000", Currency: "IDR"},
		Method:  pg.PaymentMethodQRIS,
	})
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(authHeader, "Basic "), "should use Basic auth")
}

func TestProviderCode(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})
	assert.Equal(t, pg.ProviderXendit, svc.ProviderCode())
}

func TestDefaultCurrencyFallback(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})
	assert.Equal(t, "IDR", svc.cfg.Currency())
}

func TestDefaultEndpointFallbacks(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	assert.Equal(t, "/payment_requests", svc.cfg.Endpoint("payment_requests", "/payment_requests"))
	assert.Equal(t, "/payment_requests/{id}", svc.cfg.Endpoint("payment_request_status", "/payment_requests/{id}"))
	assert.Equal(t, "/payment_requests/{id}/cancel", svc.cfg.Endpoint("payment_request_cancel", "/payment_requests/{id}/cancel"))
	assert.Equal(t, "/refunds", svc.cfg.Endpoint("refunds", "/refunds"))
	assert.Equal(t, "/payouts", svc.cfg.Endpoint("payouts", "/payouts"))
	assert.Equal(t, "/payouts/{id}", svc.cfg.Endpoint("payout_status", "/payouts/{id}"))
}

func TestCustomEndpoints(t *testing.T) {
	cfg := pg.ProviderConfig{
		Endpoints: map[string]string{
			"payment_requests": "/v2/payment_requests",
			"refunds":          "/v2/refunds",
		},
	}

	assert.Equal(t, "/v2/payment_requests", cfg.Endpoint("payment_requests", "/payment_requests"))
	assert.Equal(t, "/v2/refunds", cfg.Endpoint("refunds", "/refunds"))
	assert.Equal(t, "/payment_requests/{id}", cfg.Endpoint("payment_request_status", "/payment_requests/{id}"))
}

func TestGetPaymentStatus_FallbackToOrderID(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/payment_requests/order-fb-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "pr-1", ReferenceID: "order-fb-1", Status: "SUCCEEDED",
			Amount: 50000, Currency: "IDR",
			PaymentMethod: model.PaymentMethodResponse{Type: "QR_CODE"},
		})
	}))
	defer server.Close()

	resp, err := svc.GetPaymentStatus(context.Background(), pg.GetPaymentStatusRequest{
		OrderID: "order-fb-1",
	})
	require.NoError(t, err)
	assert.Equal(t, "order-fb-1", resp.OrderID)
	assert.Equal(t, pg.PaymentStatusSucceeded, resp.Status)
}

func TestCancelPayment_FallbackToReference(t *testing.T) {
	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/payment_requests/ref-1/cancel", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.PaymentRequestResponse{
			ID: "ref-1", ReferenceID: "ord-2", Status: "CANCELED",
			PaymentMethod: model.PaymentMethodResponse{Type: "QR_CODE"},
		})
	}))
	defer server.Close()

	resp, err := svc.CancelPayment(context.Background(), pg.CancelPaymentRequest{
		ProviderReference: "ref-1",
	})
	require.NoError(t, err)
	assert.Equal(t, pg.PaymentStatusCanceled, resp.Status)
}

func TestCreatePayout_Success(t *testing.T) {
	mockResp := model.PayoutResponse{
		ID:          "pay-123",
		ReferenceID: "ext-001",
		Status:      "PENDING",
		Amount:      200000,
		Currency:    "IDR",
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/payouts", r.URL.Path)

		var body model.PayoutRequest
		json.NewDecoder(r.Body).Decode(&body)
		assert.Equal(t, "BCA", body.ChannelCode)
		assert.Equal(t, 200000.0, body.Amount)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.CreatePayout(context.Background(), pg.CreatePayoutRequest{
		PayoutID:       "po-001",
		ExternalID:     "ext-001",
		Amount:         pg.Money{Amount: "200000", Currency: "IDR"},
		BankCode:       "bca",
		AccountNumber:  "1234567890",
		AccountName:    "Budi",
		IdempotencyKey: "idem-po",
	})
	require.NoError(t, err)
	assert.Equal(t, "pay-123", resp.ProviderPayoutID)
	assert.Equal(t, pg.PaymentStatusPending, resp.Status)
}

func TestGetPayoutStatus_Success(t *testing.T) {
	mockResp := model.PayoutResponse{
		ID:          "pay-123",
		ReferenceID: "ext-001",
		Status:      "SUCCEEDED",
	}

	svc, server := newTestService(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/payouts/pay-123", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	resp, err := svc.GetPayoutStatus(context.Background(), pg.GetPayoutStatusRequest{
		ProviderPayoutID: "pay-123",
	})
	require.NoError(t, err)
	assert.Equal(t, pg.PaymentStatusSucceeded, resp.Status)
}

func TestGetPayoutStatus_NoID(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.GetPayoutStatus(context.Background(), pg.GetPayoutStatusRequest{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "payout id is required")
}

func TestRefundPayment_InvalidAmountString(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.RefundPayment(context.Background(), pg.RefundRequest{
		OrderID:  "order-001",
		RefundID: "rf-001",
		Amount:   pg.Money{Amount: "abc", Currency: "IDR"},
	})
	require.Error(t, err)
}

func TestCreatePayout_InvalidAmount(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	_, err := svc.CreatePayout(context.Background(), pg.CreatePayoutRequest{
		PayoutID:   "po-001",
		ExternalID: "ext-001",
		Amount:     pg.Money{Amount: "-500", Currency: "IDR"},
	})
	require.Error(t, err)
}

func TestNormalizeWebhook_FailedStatus(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	webhookBody := model.WebhookPayment{
		ID:    "evt-2",
		Event: "payment.failed",
		Data: model.PaymentRequestResponse{
			ID:          "pr-2",
			ReferenceID: "ord-2",
			Status:      "FAILED",
			Amount:      75000,
			Currency:    "IDR",
		},
	}
	body, _ := json.Marshal(webhookBody)

	evt, err := svc.NormalizeWebhook(context.Background(), pg.NormalizeWebhookRequest{
		Headers: http.Header{"webhook-id": []string{"evt-2"}},
		RawBody: body,
	})
	require.NoError(t, err)
	assert.Equal(t, pg.EventPaymentFailed, evt.EventType)
	assert.Equal(t, pg.PaymentStatusFailed, evt.PaymentStatus)
}

func TestNormalizeWebhook_SucceededStatus(t *testing.T) {
	svc := ProvideService(pg.ProviderConfig{})

	webhookBody := model.WebhookPayment{
		ID:    "evt-3",
		Event: "payment.succeeded",
		Data: model.PaymentRequestResponse{
			ID:          "pr-3",
			ReferenceID: "ord-3",
			Status:      "SUCCEEDED",
			Amount:      200000,
			Currency:    "IDR",
		},
	}
	body, _ := json.Marshal(webhookBody)

	evt, err := svc.NormalizeWebhook(context.Background(), pg.NormalizeWebhookRequest{
		Headers: http.Header{"webhook-id": []string{"evt-3"}},
		RawBody: body,
	})
	require.NoError(t, err)
	assert.Equal(t, pg.EventPaymentSucceeded, evt.EventType)
	assert.Equal(t, pg.PaymentStatusSucceeded, evt.PaymentStatus)
}

func newTestService(t *testing.T, handler http.Handler, overrides ...pg.ProviderConfig) (*ServiceImpl, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)

	cfg := pg.ProviderConfig{
		BaseURL:         server.URL,
		APIKey:          "test-key",
		DefaultCurrency: "IDR",
		HTTP: pg.HTTPClientConfig{
			RetryCount: 0,
		},
	}
	if len(overrides) > 0 {
		cfg = overrides[0]
		cfg.BaseURL = server.URL
		if cfg.APIKey == "" {
			cfg.APIKey = "test-key"
		}
		if cfg.DefaultCurrency == "" {
			cfg.DefaultCurrency = "IDR"
		}
	}

	svc := ProvideService(cfg)
	return svc, server
}

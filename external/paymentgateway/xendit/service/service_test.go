package service

import (
	"context"
	"net/http"
	"testing"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/stretchr/testify/require"
)

func TestVerifyWebhookRequiresConfiguredHMACSignature(t *testing.T) {
	service := ProvideService(pg.ProviderConfig{WebhookSecret: "webhook-secret"})
	body := []byte(`{"id":"evt-1"}`)

	result, err := service.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: http.Header{},
		RawBody: body,
	})
	require.Error(t, err)
	require.False(t, result.SignatureValid)

	headers := http.Header{}
	headers.Set("X-Callback-Signature", pg.HMACSHA256Hex("webhook-secret", body))
	result, err = service.VerifyWebhook(context.Background(), pg.VerifyWebhookRequest{
		Headers: headers,
		RawBody: body,
	})
	require.NoError(t, err)
	require.True(t, result.SignatureValid)
}

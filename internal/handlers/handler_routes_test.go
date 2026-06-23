package handlers

import (
	"net/http"
	"sort"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestHandlerRegistersPaymentOperationRoutes(t *testing.T) {
	router := chi.NewRouter()
	(&Handler{}).Router(router)

	var routes []string
	err := chi.Walk(router, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		routes = append(routes, method+" "+route)
		return nil
	})
	require.NoError(t, err)
	sort.Strings(routes)

	require.Subset(t, routes, []string{
		"GET /payment-intents/{paymentIntentID}",
		"POST /admin/cash-sessions",
		"POST /admin/cash-sessions/{cashSessionID}/{action}",
		"POST /admin/manual-payment-evidence/{evidenceID}/{action}",
		"POST /admin/overpayments/{overpaymentID}/{action}",
		"POST /admin/payment-authorizations/{authorizationID}/{action}",
		"POST /admin/payment-installments/{installmentID}/{action}",
		"POST /admin/refunds/{refundID}/{action}",
		"POST /admin/webhooks/{webhookID}/{action}",
		"POST /payment-intents/",
		"POST /payment-intents/{paymentIntentID}/{action}",
		"POST /refunds",
		"POST /webhooks/{provider}",
	})
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/stretchr/testify/require"
)

func TestRequireAdminTokenFailsClosed(t *testing.T) {
	handler := RequireAdminToken(&configs.Config{})(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("protected handler must not run")
	}))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/admin/refunds/id/approve", nil))
	require.Equal(t, http.StatusServiceUnavailable, response.Code)
}

func TestRequireAdminToken(t *testing.T) {
	config := &configs.Config{}
	config.Internal.Payment.AdminToken = "secret-token"
	handler := RequireAdminToken(config)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	unauthorized := httptest.NewRecorder()
	handler.ServeHTTP(unauthorized, httptest.NewRequest(http.MethodPost, "/admin", nil))
	require.Equal(t, http.StatusUnauthorized, unauthorized.Code)

	request := httptest.NewRequest(http.MethodPost, "/admin", nil)
	request.Header.Set(InternalAdminTokenHeader, "secret-token")
	authorized := httptest.NewRecorder()
	handler.ServeHTTP(authorized, request)
	require.Equal(t, http.StatusNoContent, authorized.Code)
}

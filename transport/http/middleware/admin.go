package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/nuriansyah/lokatra-payment/configs"
)

const InternalAdminTokenHeader = "X-Internal-Admin-Token"

// RequireAdminToken protects operational endpoints until the service is wired
// to the platform authorizer. It deliberately fails closed when unconfigured.
func RequireAdminToken(cfg *configs.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg == nil || strings.TrimSpace(cfg.Internal.Payment.AdminToken) == "" {
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}
			expected := []byte(cfg.Internal.Payment.AdminToken)
			provided := []byte(r.Header.Get(InternalAdminTokenHeader))
			if len(provided) != len(expected) || subtle.ConstantTimeCompare(provided, expected) != 1 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

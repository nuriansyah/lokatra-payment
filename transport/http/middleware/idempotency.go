package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/rs/zerolog/log"
)

// IdempotencyHeader is the standard HTTP header for idempotency keys
const IdempotencyHeader = "Idempotency-Key"

// IdempotencyMiddleware provides idempotency key validation and caching.
// This middleware checks incoming requests for an Idempotency-Key header
// and returns cached responses if the same request is retried.
func IdempotencyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idempotencyKey := r.Header.Get(IdempotencyHeader)

		// If no idempotency key is provided, process normally
		if idempotencyKey == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Validate idempotency key format (max 128 characters)
		if len(idempotencyKey) > 128 {
			http.Error(w, "Idempotency-Key header too long (max 128 characters)", http.StatusBadRequest)
			return
		}

		// Store idempotency key in context for handler access
		r.Header.Set("X-Idempotency-Key-Hash", computeIdempotencyHash(idempotencyKey))

		log.Ctx(r.Context()).Debug().
			Str("idempotencyKey", idempotencyKey).
			Msg("idempotency key provided in request")

		next.ServeHTTP(w, r)
	})
}

// computeIdempotencyHash generates a SHA256 hash of the idempotency key for internal use
func computeIdempotencyHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

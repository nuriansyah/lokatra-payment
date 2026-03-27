package middleware

import (
	"context"
	"net/http"

	uuid "github.com/gofrs/uuid"
)

// Key to use when setting the request ID.
type ctxRequestIDKey string

const RequestIDKey ctxRequestIDKey = "RequestID"

// RequestID is a middleware that injects a request ID into the context of each
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RequestIDKey, uuid.Must(uuid.NewV7()).String())
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

// GetRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

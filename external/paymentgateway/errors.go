package paymentgateway

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrProviderNotConfigured = errors.New("payment provider not configured")
	ErrUnsupportedOperation  = errors.New("payment provider operation is unsupported")
	ErrInvalidWebhook        = errors.New("invalid provider webhook")
	ErrInvalidRequest        = errors.New("invalid payment gateway request")
)

type ErrorCode string

const (
	ErrorCodeInvalidRequest      ErrorCode = "PAYMENT_GATEWAY_INVALID_REQUEST"
	ErrorCodeUnauthorized        ErrorCode = "PAYMENT_GATEWAY_UNAUTHORIZED"
	ErrorCodeForbidden           ErrorCode = "PAYMENT_GATEWAY_FORBIDDEN"
	ErrorCodeNotFound            ErrorCode = "PAYMENT_GATEWAY_NOT_FOUND"
	ErrorCodeRateLimited         ErrorCode = "PAYMENT_GATEWAY_RATE_LIMITED"
	ErrorCodeProviderTimeout     ErrorCode = "PAYMENT_GATEWAY_TIMEOUT"
	ErrorCodeProviderUnavailable ErrorCode = "PAYMENT_GATEWAY_UNAVAILABLE"
	ErrorCodeProviderError       ErrorCode = "PAYMENT_GATEWAY_PROVIDER_ERROR"
	ErrorCodeWebhookInvalid      ErrorCode = "PAYMENT_GATEWAY_WEBHOOK_INVALID"
	ErrorCodeUnsupported         ErrorCode = "PAYMENT_GATEWAY_UNSUPPORTED"
)

type GatewayError struct {
	Code           ErrorCode    `json:"code"`
	ProviderCode   ProviderCode `json:"providerCode,omitempty"`
	HTTPStatusCode int          `json:"httpStatusCode,omitempty"`
	Message        string       `json:"message"`
	Retryable      bool         `json:"retryable"`
	Cause          error        `json:"-"`
}

func (e *GatewayError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *GatewayError) Unwrap() error { return e.Cause }

func NewGatewayError(provider ProviderCode, code ErrorCode, status int, msg string, retryable bool, cause error) error {
	return &GatewayError{Code: code, ProviderCode: provider, HTTPStatusCode: status, Message: msg, Retryable: retryable, Cause: cause}
}

func ErrorFromHTTP(provider ProviderCode, status int, body string) error {
	code := ErrorCodeProviderError
	retryable := false
	switch {
	case status == http.StatusUnauthorized:
		code = ErrorCodeUnauthorized
	case status == http.StatusForbidden:
		code = ErrorCodeForbidden
	case status == http.StatusNotFound:
		code = ErrorCodeNotFound
	case status == http.StatusTooManyRequests:
		code = ErrorCodeRateLimited
		retryable = true
	case status == http.StatusRequestTimeout || status == 499:
		code = ErrorCodeProviderTimeout
		retryable = true
	case status >= 500:
		code = ErrorCodeProviderUnavailable
		retryable = true
	case status >= 400:
		code = ErrorCodeInvalidRequest
	}
	if body == "" {
		body = http.StatusText(status)
	}
	return NewGatewayError(provider, code, status, body, retryable, nil)
}

func IsRetryable(err error) bool {
	var gw *GatewayError
	if errors.As(err, &gw) {
		return gw.Retryable
	}
	return false
}

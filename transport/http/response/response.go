package response

import (
	"encoding/json"
	"net/http"

	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/logger"
)

// Base is the base object of all responses
type Base struct {
	HTTPCode   *int                             `json:"-"`
	RequestID  *string                          `json:"requestId,omitempty"`
	Data       *interface{}                     `json:"data,omitempty"`
	Metadata   *interface{}                     `json:"metadata,omitempty"`
	Error      *string                          `json:"error,omitempty"`
	ErrorCode  *string                          `json:"errorCode,omitempty"`
	Errors     failure.PartialFailureErrorField `json:"errors,omitempty"`
	Message    *string                          `json:"message,omitempty"`
	StatusCode *int                             `json:"statusCode,omitempty"`
}

// NoContent sends a response without any content
func NoContent(w http.ResponseWriter) {
	respond(w, http.StatusNoContent, nil)
}

func SetHTTPCode(httpCode int) func(*Base) {
	return func(b *Base) {
		b.HTTPCode = &httpCode
	}
}

func SetStatusCode(statusCode int) func(*Base) {
	return func(b *Base) {
		b.StatusCode = &statusCode
	}
}

func SetError(err error) func(*Base) {
	return func(b *Base) {
		errStr := err.Error()
		b.Error = &errStr
	}
}
func SetMessage(msg string) func(*Base) {
	return func(b *Base) {
		b.Message = &msg
	}
}
func SetRequestID(requestId string) func(*Base) {
	return func(b *Base) {
		b.RequestID = &requestId
	}
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string, opts ...func(*Base)) {
	base := Base{Message: &message, StatusCode: &code, HTTPCode: &code}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, *base.HTTPCode, base)
}

// WithJSON sends a response containing a JSON object
func WithJSON(w http.ResponseWriter, code int, jsonPayload interface{}, opts ...func(*Base)) {
	base := Base{Data: &jsonPayload, StatusCode: &code, HTTPCode: &code}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, *base.HTTPCode, base)
}

// WithJSONErrs sends a response containing a JSON objec and errors
func WithJSONErrs(w http.ResponseWriter, code int, jsonPayload interface{}, errs failure.PartialFailureErrorField, opts ...func(*Base)) {
	base := Base{Data: &jsonPayload, StatusCode: &code, HTTPCode: &code, Errors: errs}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, *base.HTTPCode, base)
}

// WithFiles send a response containing Files
func WithFiles(w http.ResponseWriter, attachment []byte, contentType string) {
	w.Header().Set("Content-Type", contentType)
	_, err := w.Write(attachment)
	if err != nil {
		logger.ErrorWithStack(err)
	}
}

// WithPNGImage send a response containing Images with PNG Extension
func WithPNGImage(w http.ResponseWriter, attachment []byte) {
	WithFiles(w, attachment, "image/png")
}

// WithMetadata sends a response containing a JSON object with metadata
func WithMetadata(w http.ResponseWriter, code int, jsonPayload interface{}, metadata interface{}, opts ...func(*Base)) {
	base := Base{Data: &jsonPayload, Metadata: &metadata, StatusCode: &code, HTTPCode: &code}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, *base.HTTPCode, base)
}

// WithError sends a response with an error message
func WithError(w http.ResponseWriter, err error, opts ...func(*Base)) {
	code := failure.GetCode(err)
	errCode := failure.GetErrorCode(err)
	errMsg := err.Error()
	base := Base{Error: &errMsg, StatusCode: &code, HTTPCode: &code, ErrorCode: &errCode}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, *base.HTTPCode, base)
}

// WithError sends a response with an error message
func WithErrors(w http.ResponseWriter, errs failure.PartialFailureErrorField, opts ...func(*Base)) {
	errors := make(failure.PartialFailureErrorField)
	var code int
	for index, err := range errs {
		code = failure.GetCode(err)
		errors[index] = err
	}
	base := Base{Errors: errors, StatusCode: &code, HTTPCode: &code}
	for _, opt := range opts {
		opt(&base)
	}
	respond(w, code, base)
}

// WithPreparingShutdown sends a default response for when the server is preparing to shut down
func WithPreparingShutdown(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER PREPARING TO SHUT DOWN")
}

// WithUnhealthy sends a default response for when the server is unhealthy
func WithUnhealthy(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER UNHEALTHY")
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logger.ErrorWithStack(err)
	}
}

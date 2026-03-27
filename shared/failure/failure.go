package failure

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/nuriansyah/lokatra-payment/shared"
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternalError    = errors.New("internal server error")
	ErrUnimplemented    = errors.New("unimplemented")
	ErrForbidden        = errors.New("forbidden")
	ErrNotFound         = errors.New("not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrConflict         = errors.New("conflict")
	ErrFailedDependency = errors.New("failed dependency")
)

// Failure is a wrapper for error messages and codes using standard HTTP response codes.
type Failure struct {
	Code        int    `json:"code"`
	ErrorCode   string `json:"errorCode"`
	Message     string `json:"message"`
	originalErr error
}

// Error returns the error code and message in a formatted string.
func (e *Failure) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.Code), e.Message)
}

func (e *Failure) OriginalError() error {
	return e.originalErr
}

// New returns a new Failure with code for custom error.
func Init(code int, err error) *Failure {
	if err != nil {
		return &Failure{
			Code:    code,
			Message: err.Error(),
		}
	}
	return nil
}

// New returns a new Failure with code for custom error.
func New(code int, err error) error {
	if err != nil {
		errCode, _ := shared.ParseErrorCode(err.Error())
		return &Failure{
			Code:      code,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

// BadRequest returns a new Failure with code for bad requests.
func BadRequest(err error) error {
	if err != nil {
		errCode, _ := shared.ParseErrorCode(err.Error())
		return &Failure{
			Code:      http.StatusBadRequest,
			ErrorCode: errCode,
			Message:   err.Error(),
		}
	}
	return nil
}

// BadRequestFromString returns a new Failure with code for bad requests with message set from string.
func BadRequestFromString(msg string) error {
	errCode, _ := shared.ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusBadRequest,
		Message:   msg,
		ErrorCode: errCode,
	}
}

// Unauthorized returns a new Failure with code for unauthorized requests.
func Unauthorized(msg string) error {
	errCode, _ := shared.ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusUnauthorized,
		Message:   msg,
		ErrorCode: errCode,
	}
}

// UnprocessableEntity return a new Failure with code for unprocessable entity situations.
func UnprocessableEntity(err error) error {
	if err != nil {
		errCode, _ := shared.ParseErrorCode(err.Error())
		return &Failure{
			Code:      http.StatusUnprocessableEntity,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

// RequestStillProcessing returns a new Failure with code for processing.
func RequestStillProcessing(err error) error {
	if err != nil {
		errCode, _ := shared.ParseErrorCode(err.Error())
		return &Failure{
			Code:      http.StatusProcessing,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

// InternalError returns a new Failure with code for internal error and message derived from an error interface.
func InternalError(err error) error {
	if err != nil {
		errCode, _ := shared.ParseErrorCode(err.Error())
		originalErr := err
		if failure, ok := err.(*Failure); ok {
			originalErr = failure.OriginalError()
		}
		return &Failure{
			Code:        http.StatusInternalServerError,
			Message:     err.Error(),
			ErrorCode:   errCode,
			originalErr: originalErr,
		}
	}
	return nil
}

// Unimplemented returns a new Failure with code for unimplemented method.
func Unimplemented(methodName string) error {
	return &Failure{
		Code:    http.StatusNotImplemented,
		Message: methodName,
	}
}

// NotFound returns a new Failure with code for entity not found.
func NotFound(domainName string) error {
	errCode, _ := shared.ParseErrorCode(domainName)
	return &Failure{
		Code:      http.StatusNotFound,
		Message:   domainName,
		ErrorCode: errCode,
	}
}

// Conflict returns a new Failure with code for conflict situations.
func Conflict(operationName string, domainName string, message string) error {
	return &Failure{
		Code:    http.StatusConflict,
		Message: fmt.Sprintf("%s on %s: %s", operationName, domainName, message),
	}
}

// GetCode returns the error code of an error interface.
func GetCode(err error) int {
	if f, ok := err.(*Failure); ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

// GetErrorCode returns the error code of an error interface.
func GetErrorCode(err error) string {
	if f, ok := err.(*Failure); ok {
		return f.ErrorCode
	}
	return ""
}

type PartialFailureErrorField map[int]*Failure

func InitPartialFailureErrorField() PartialFailureErrorField {
	res := map[int]*Failure{}
	return res
}
func (p *PartialFailureErrorField) Add(index int, err error) *PartialFailureErrorField {
	code := GetCode(err)
	errCode := GetErrorCode(err)
	errMsg := err.Error()
	(*p)[index] = &Failure{
		Code:      code,
		ErrorCode: errCode,
		Message:   errMsg,
	}
	return p
}
func (p *PartialFailureErrorField) AddBreakError(err error) *PartialFailureErrorField {
	code := GetCode(err)
	errCode := GetErrorCode(err)
	errMsg := err.Error()
	(*p)[-1] = &Failure{
		Code:      code,
		ErrorCode: errCode,
		Message:   errMsg,
	}
	return p
}
func (p *PartialFailureErrorField) FirstError() error {
	for _, err := range *p {
		return err
	}
	return nil
}

func (p *PartialFailureErrorField) BreakError() error {
	if (*p)[-1] != nil {
		return (*p)[-1]
	}
	return nil
}

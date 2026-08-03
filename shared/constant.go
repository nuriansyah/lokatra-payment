package shared

import (
	"strings"
)

const (
	ErrorInternalSystem               = "internal system error"
	ErrorInvalidRequest               = "invalid request"
	ErrorNotFound                     = "not found"
	ErrorPQConstrainViolated          = "database constraint violated"
	ErrorPQDuplicateConstrainViolated = "database duplicate constraint violated"
)

// Request validation error codes — used in DTO Validate() methods.
const (
	ErrInvalidActorID      = "ERR-VAL-01"
	ErrInvalidMerchantID   = "ERR-VAL-02"
	ErrInvalidSourceService = "ERR-VAL-03"
	ErrInvalidSourceType   = "ERR-VAL-04"
	ErrInvalidSourceID     = "ERR-VAL-05"
	ErrInvalidAmount       = "ERR-VAL-06"
	ErrInvalidCurrency     = "ERR-VAL-07"
	ErrInvalidID           = "ERR-VAL-08"
	ErrInvalidCollectorID  = "ERR-VAL-09"
	ErrInvalidExpiryDate   = "ERR-VAL-10"
	ErrInvalidMetadata     = "ERR-VAL-11"
)

const (
	ErrPayInvalidAmount                        = "ERR-PAY-01"
	ErrPayUnsupportedCurrency                  = "ERR-PAY-02"
	ErrPayMethodDetailsMismatch                = "ERR-PAY-04"
	ErrPayAllPSPUnavailable                    = "ERR-PAY-05"
	ErrPayPaymentDeclined                      = "ERR-PAY-06"
	ErrPayFailoverSuccess                      = "ERR-PAY-07"
	ErrPayUnknownPSPError                      = "ERR-PAY-08"
	ErrPayIdempotencyKeyReusedDifferentPayload = "ERR-PAY-09"
	ErrPayIdempotencyKeyRequired               = "ERR-PAY-10"
)

const (
	ErrRouteEmptyPSPList           = "ERR-ROUTE-01"
	ErrRouteScopeConflict          = "ERR-ROUTE-04"
	ErrRouteMDRDataStale           = "ERR-ROUTE-07"
	ErrRouteOptimisticLockConflict = "ERR-ROUTE-09"
	ErrRoutePermissionDenied       = "ERR-ROUTE-11"
)

func ProblemStatusForCode(code string) int {
	switch code {
	// Validation errors → 400
	case ErrInvalidActorID, ErrInvalidMerchantID, ErrInvalidSourceService,
		ErrInvalidSourceType, ErrInvalidSourceID, ErrInvalidAmount,
		ErrInvalidCurrency, ErrInvalidID, ErrInvalidCollectorID,
		ErrInvalidExpiryDate, ErrInvalidMetadata:
		return 400
	case ErrPayInvalidAmount, ErrPayUnsupportedCurrency, ErrPayMethodDetailsMismatch:
		return 422
	case ErrPayAllPSPUnavailable:
		return 503
	case ErrPayPaymentDeclined:
		return 402
	case ErrPayUnknownPSPError:
		return 502
	case ErrPayIdempotencyKeyReusedDifferentPayload:
		return 409
	case ErrPayIdempotencyKeyRequired:
		return 400
	case ErrRouteEmptyPSPList, ErrRouteScopeConflict:
		return 422
	case ErrRouteMDRDataStale:
		return 422
	case ErrRouteOptimisticLockConflict:
		return 409
	case ErrRoutePermissionDenied:
		return 403
	default:
		return 500
	}
}

func IsHardDecline(pspErrorCode string) bool {
	switch strings.ToUpper(pspErrorCode) {
	case "CARD_DECLINED", "INSUFFICIENT_FUNDS", "EXPIRED_CARD", "INCORRECT_CVC",
		"INVALID_CARD_NUMBER", "DO_NOT_HONOR", "RESTRICTED_CARD",
		"TRANSACTION_NOT_PERMITTED", "EXCEEDS_WITHDRAWAL_LIMIT",
		"INVALID_AMOUNT", "PAYMENT_METHOD_NOT_SUPPORTED",
		"AUTHENTICATION_REQUIRED", "DO_NOT_RETRY":
		return true
	default:
		return false
	}
}

func IsSystemFailure(pspErrorCode string) bool {
	switch strings.ToUpper(pspErrorCode) {
	case "TIMEOUT", "GATEWAY_TIMEOUT", "SERVICE_UNAVAILABLE",
		"INTERNAL_SERVER_ERROR", "BAD_GATEWAY", "NETWORK_ERROR",
		"CONNECTION_RESET", "CONNECTION_REFUSED":
		return true
	default:
		return false
	}
}

func ParseErrorCode(c string) (string, bool) {
	index := strings.Index(c, ":")
	if index == -1 {
		return "", false
	}

	return c[:index], true
}

package service

import (
	"strings"
	"time"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
)

// normalizeStatus maps Xendit's raw status strings to Lokatra's canonical PaymentStatus enum.
func normalizeStatus(s string) pg.PaymentStatus {
	switch strings.ToUpper(s) {
	case "SUCCEEDED", "SUCCESS", "COMPLETED", "PAID", "CAPTURED":
		return pg.PaymentStatusSucceeded
	case "PENDING", "REQUIRES_ACTION", "AWAITING_CAPTURE":
		return pg.PaymentStatusPending
	case "FAILED", "VOIDED":
		return pg.PaymentStatusFailed
	case "EXPIRED":
		return pg.PaymentStatusExpired
	case "CANCELED", "CANCELLED":
		return pg.PaymentStatusCanceled
	case "REFUNDED":
		return pg.PaymentStatusRefunded
	case "PARTIALLY_REFUNDED":
		return pg.PaymentStatusPartiallyRefunded
	}
	return pg.PaymentStatusUnknown
}

// eventFromStatus maps a canonical PaymentStatus to the corresponding canonical EventType.
func eventFromStatus(s pg.PaymentStatus) pg.EventType {
	switch s {
	case pg.PaymentStatusSucceeded:
		return pg.EventPaymentSucceeded
	case pg.PaymentStatusPending:
		return pg.EventPaymentPending
	case pg.PaymentStatusFailed:
		return pg.EventPaymentFailed
	case pg.PaymentStatusExpired:
		return pg.EventPaymentExpired
	case pg.PaymentStatusCanceled:
		return pg.EventPaymentCanceled
	case pg.PaymentStatusRefunded:
		return pg.EventPaymentRefunded
	case pg.PaymentStatusPartiallyRefunded:
		return pg.EventPaymentPartiallyRefunded
	}
	return pg.EventUnknown
}

// firstNonEmpty returns the first non-blank string from the variadic list.
func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

// parseTime parses an RFC3339 string into *time.Time, returning nil on failure.
func parseTime(s string) *time.Time {
	if s == "" {
		return nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return &t
	}
	return nil
}

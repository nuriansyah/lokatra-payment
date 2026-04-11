package service

import (
	"fmt"

	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
)

// PaymentStateMachine centralizes the allowed lifecycle transitions.
type PaymentStateMachine struct {
	allowed map[paymentmodel.PaymentStatus]map[paymentmodel.PaymentStatus]struct{}
}

func NewPaymentStateMachine() *PaymentStateMachine {
	return &PaymentStateMachine{
		allowed: map[paymentmodel.PaymentStatus]map[paymentmodel.PaymentStatus]struct{}{
			paymentmodel.PaymentStatusInitiated: {
				paymentmodel.PaymentStatusPending:   {},
				paymentmodel.PaymentStatusFailed:    {},
				paymentmodel.PaymentStatusCancelled: {},
				paymentmodel.PaymentStatusExpired:   {},
			},
			paymentmodel.PaymentStatusPending: {
				paymentmodel.PaymentStatusAuthorised: {},
				paymentmodel.PaymentStatusCaptured:   {},
				paymentmodel.PaymentStatusFailed:     {},
				paymentmodel.PaymentStatusCancelled:  {},
				paymentmodel.PaymentStatusExpired:    {},
			},
			paymentmodel.PaymentStatusAuthorised: {
				paymentmodel.PaymentStatusCaptured:          {},
				paymentmodel.PaymentStatusPartiallyCaptured: {},
				paymentmodel.PaymentStatusFailed:            {},
				paymentmodel.PaymentStatusCancelled:         {},
			},
			paymentmodel.PaymentStatusCaptured: {
				paymentmodel.PaymentStatusCompleted:         {},
				paymentmodel.PaymentStatusRefunding:         {},
				paymentmodel.PaymentStatusPartiallyRefunded: {},
				paymentmodel.PaymentStatusDisputed:          {},
			},
			paymentmodel.PaymentStatusPartiallyCaptured: {
				paymentmodel.PaymentStatusCompleted: {},
				paymentmodel.PaymentStatusRefunding: {},
				paymentmodel.PaymentStatusFailed:    {},
			},
			paymentmodel.PaymentStatusCompleted: {
				paymentmodel.PaymentStatusRefunding: {},
				paymentmodel.PaymentStatusDisputed:  {},
			},
			paymentmodel.PaymentStatusRefunding: {
				paymentmodel.PaymentStatusRefunded:          {},
				paymentmodel.PaymentStatusPartiallyRefunded: {},
				paymentmodel.PaymentStatusFailed:            {},
			},
			paymentmodel.PaymentStatusDisputed: {
				paymentmodel.PaymentStatusChargebackWon:  {},
				paymentmodel.PaymentStatusChargebackLost: {},
			},
		},
	}
}

func (sm *PaymentStateMachine) CanTransition(from, to paymentmodel.PaymentStatus) bool {
	if sm == nil {
		return false
	}
	if from == to {
		return true
	}
	allowedNext, ok := sm.allowed[from]
	if !ok {
		return false
	}
	_, ok = allowedNext[to]
	return ok
}

func (sm *PaymentStateMachine) Transition(from, to paymentmodel.PaymentStatus) (paymentmodel.PaymentStatus, error) {
	if sm.CanTransition(from, to) {
		return to, nil
	}
	return from, fmt.Errorf("invalid payment transition from %s to %s", from, to)
}

func (sm *PaymentStateMachine) NormalizeGatewayStatus(status paymentmodel.PaymentStatus, requiresAction bool) paymentmodel.PaymentStatus {
	switch status {
	case paymentmodel.PaymentStatusAuthorised,
		paymentmodel.PaymentStatusCaptured,
		paymentmodel.PaymentStatusFailed,
		paymentmodel.PaymentStatusCancelled,
		paymentmodel.PaymentStatusExpired,
		paymentmodel.PaymentStatusPending:
		return status
	default:
		if requiresAction {
			return paymentmodel.PaymentStatusPending
		}
		return paymentmodel.PaymentStatusFailed
	}
}

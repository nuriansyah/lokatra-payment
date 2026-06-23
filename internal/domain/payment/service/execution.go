package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PaymentIntentActionResult struct {
	Intent  paymentmodel.PaymentIntents
	Routing *RoutingResult
}

func (s *ServiceImpl) executePaymentIntent(ctx context.Context, intent paymentmodel.PaymentIntents, actorID uuid.UUID) (PaymentIntentActionResult, error) {
	method, ok := normalizePaymentMethod(intent.SelectedMethodCode.String)
	if !ok {
		return PaymentIntentActionResult{Intent: intent}, fmt.Errorf("unsupported payment method %q", intent.SelectedMethodCode.String)
	}
	intent.Status = paymentmodel.PaymentIntentStatusProcessing
	setUpdateSignature(&intent.MetaSignature, actorID, time.Now().UTC())
	if err := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); err != nil {
		return PaymentIntentActionResult{Intent: intent}, err
	}
	routingRequest := RoutingRequest{
		Method:   method,
		Channel:  intent.SelectedChannelCode.String,
		Currency: intent.Currency,
		GatewayCall: pg.CreatePaymentRequest{
			PaymentIntentID: intent.Id.String(),
			AttemptID:       mustUUID().String(),
			OrderID:         intent.IntentCode,
			Amount:          pg.Money{Amount: intent.Amount.StringFixed(2), Currency: intent.Currency},
			Method:          method,
			ChannelCode:     intent.SelectedChannelCode.String,
			Description:     intent.Description.String,
			ExpiryAt:        intent.ExpiresAt.Ptr(),
			IdempotencyKey:  intent.IdempotencyKey,
		},
	}
	routingResult, routeErr := s.routingEngine.Execute(ctx, routingRequest)
	now := time.Now().UTC()
	candidateJSON, _ := json.Marshal(routingResult.Candidates)
	evaluatedJSON, _ := json.Marshal(map[string]any{"method": method, "channel": routingRequest.Channel, "currency": routingRequest.Currency})
	decision := paymentmodel.PaymentRouteDecisions{
		Id:                        mustUUID(),
		PaymentIntentId:           intent.Id,
		SelectedProviderAccountId: optionalUUID(routingResult.Selected.AccountID),
		SelectedProviderCode:      optionalString(string(routingResult.Selected.ProviderCode)),
		MethodCode:                string(method),
		ChannelCode:               optionalString(routingRequest.Channel),
		Reason:                    routingDecisionReason(routingResult, routeErr),
		EvaluatedContext:          evaluatedJSON,
		Candidates:                candidateJSON,
		Metadata:                  json.RawMessage(`{}`),
		MetaSignature:             shared.MetaSignature{MetaCreatedAt: now, MetaCreatedBy: actorID},
	}
	if err := s.paymentRepo.CreatePaymentRouteDecisions(ctx, &decision); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, err
	}

	nextAttemptNo, err := s.nextPaymentAttemptNumber(ctx, intent.Id)
	if err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, err
	}
	var successfulAttemptID uuid.UUID
	for index, providerAttempt := range routingResult.Attempts {
		attempt := paymentmodel.PaymentAttempts{
			Id:                mustUUID(),
			PaymentIntentId:   intent.Id,
			AttemptNo:         nextAttemptNo + index,
			ProviderAccountId: optionalUUID(providerAttempt.AccountID),
			RouteDecisionId:   nuuid.From(decision.Id),
			ProviderCode:      null.StringFrom(string(providerAttempt.ProviderCode)),
			MethodCode:        string(method),
			ChannelCode:       optionalString(routingRequest.Channel),
			Amount:            intent.Amount,
			Currency:          intent.Currency,
			Status:            paymentmodel.PaymentAttemptStatusFailed,
			FailureMessage:    optionalString(providerAttempt.Error),
			RawRequest:        mustJSON(routingRequest.GatewayCall),
			Metadata:          mustJSON(map[string]any{"durationMs": providerAttempt.Duration.Milliseconds(), "providerAttempt": providerAttempt.Attempt}),
			MetaSignature:     shared.MetaSignature{MetaCreatedAt: providerAttempt.StartedAt, MetaCreatedBy: actorID},
		}
		if providerAttempt.Error == "" && providerAttempt.ProviderCode == routingResult.Selected.ProviderCode {
			attempt.Status = paymentAttemptStatus(routingResult.Payment.Status)
			attempt.ProviderReference = optionalString(routingResult.Payment.ProviderReference)
			attempt.ProviderTransactionId = optionalString(routingResult.Payment.ProviderTransactionID)
			attempt.ProviderOrderId = optionalString(routingResult.Payment.OrderID)
			attempt.ProviderPaymentId = optionalString(routingResult.Payment.ProviderPaymentID)
			attempt.RawResponse = routingResult.Payment.Raw
			successfulAttemptID = attempt.Id
		}
		if err := s.paymentRepo.CreatePaymentAttempts(ctx, &attempt); err != nil {
			return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, err
		}
	}

	if routeErr != nil {
		intent.Status = paymentmodel.PaymentIntentStatusRequiresConfirmation
		setUpdateSignature(&intent.MetaSignature, actorID, time.Now().UTC())
		if updateErr := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); updateErr != nil {
			return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, updateErr
		}
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, routeErr
	}
	if err := s.persistPaymentInstructions(ctx, successfulAttemptID, routingResult.Payment.Instructions, actorID); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, err
	}
	intent.Status = paymentIntentStatus(routingResult.Payment.Status)
	if intent.Status == paymentmodel.PaymentIntentStatusSucceeded {
		intent.PaidAt = null.TimeFrom(now)
	}
	setUpdateSignature(&intent.MetaSignature, actorID, now)
	if err := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, err
	}
	return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}, nil
}

func (s *ServiceImpl) nextPaymentAttemptNumber(ctx context.Context, intentID uuid.UUID) (int, error) {
	result, err := s.paymentRepo.ResolvePaymentAttemptsByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{{Field: string(paymentmodel.PaymentAttemptsDBFieldName.PaymentIntentId), Operator: paymentmodel.OperatorEqual, Value: intentID}},
		Sorts:        []paymentmodel.Sort{{Field: string(paymentmodel.PaymentAttemptsDBFieldName.AttemptNo), Order: paymentmodel.SortDesc}},
		Pagination:   paymentmodel.Pagination{Page: 1, PageSize: 1},
	})
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 1, nil
	}
	return result[0].PaymentAttempts.AttemptNo + 1, nil
}

func (s *ServiceImpl) persistPaymentInstructions(ctx context.Context, attemptID uuid.UUID, instructions []pg.PaymentInstruction, actorID uuid.UUID) error {
	for _, instruction := range instructions {
		record := paymentmodel.PaymentInstructions{
			Id:               mustUUID(),
			PaymentAttemptId: attemptID,
			InstructionType:  instruction.Type,
			IsActive:         true,
			DisplayName:      optionalString(instruction.DisplayName),
			AccountNumber:    optionalString(instruction.AccountNumber),
			BillerCode:       optionalString(instruction.BillerCode),
			PaymentCode:      optionalString(instruction.PaymentCode),
			QrString:         optionalString(instruction.QRString),
			QrImageUrl:       optionalString(instruction.QRImageURL),
			CheckoutUrl:      optionalString(instruction.CheckoutURL),
			DeeplinkUrl:      optionalString(instruction.DeeplinkURL),
			Metadata:         mustJSON(instruction.ProviderData),
			MetaSignature:    shared.MetaSignature{MetaCreatedAt: time.Now().UTC(), MetaCreatedBy: actorID},
		}
		if instruction.ExpiresAt != nil {
			record.ExpiresAt = null.TimeFrom(instruction.ExpiresAt.UTC())
		}
		if err := s.paymentRepo.CreatePaymentInstructions(ctx, &record); err != nil {
			return err
		}
	}
	return nil
}

func normalizePaymentMethod(value string) (pg.PaymentMethod, bool) {
	method := pg.PaymentMethod(strings.ToLower(strings.TrimSpace(value)))
	switch method {
	case pg.PaymentMethodVirtualAccount, pg.PaymentMethodQRIS, pg.PaymentMethodEWallet, pg.PaymentMethodCard, pg.PaymentMethodRetailOutlet, pg.PaymentMethodPaymentPage, pg.PaymentMethodManualTransfer, pg.PaymentMethodCash:
		return method, true
	default:
		return "", false
	}
}

func paymentAttemptStatus(status pg.PaymentStatus) paymentmodel.PaymentAttemptStatus {
	switch status {
	case pg.PaymentStatusSucceeded:
		return paymentmodel.PaymentAttemptStatusPaid
	case pg.PaymentStatusAuthorized:
		return paymentmodel.PaymentAttemptStatusAuthorized
	case pg.PaymentStatusCaptured:
		return paymentmodel.PaymentAttemptStatusCaptured
	case pg.PaymentStatusFailed, pg.PaymentStatusExpired:
		return paymentmodel.PaymentAttemptStatusFailed
	case pg.PaymentStatusCanceled:
		return paymentmodel.PaymentAttemptStatusCanceled
	default:
		return paymentmodel.PaymentAttemptStatusPending
	}
}

func paymentIntentStatus(status pg.PaymentStatus) paymentmodel.PaymentIntentStatus {
	switch status {
	case pg.PaymentStatusSucceeded, pg.PaymentStatusCaptured:
		return paymentmodel.PaymentIntentStatusSucceeded
	case pg.PaymentStatusRequiresAction:
		return paymentmodel.PaymentIntentStatusRequiresAction
	case pg.PaymentStatusFailed, pg.PaymentStatusExpired, pg.PaymentStatusCanceled:
		return paymentmodel.PaymentIntentStatusCanceled
	default:
		return paymentmodel.PaymentIntentStatusProcessing
	}
}

func routingDecisionReason(result RoutingResult, err error) string {
	if err != nil {
		return "routing_exhausted: " + err.Error()
	}
	return fmt.Sprintf("selected_%s_priority_%d", result.Selected.ProviderCode, result.Selected.Priority)
}

func mustJSON(value any) json.RawMessage {
	encoded, _ := json.Marshal(value)
	return encoded
}

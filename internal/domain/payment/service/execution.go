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
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PaymentIntentActionResult struct {
	Intent  paymentmodel.PaymentIntents
	Routing *RoutingResult
}

func (a PaymentIntentActionResult) ToResponse() dto.PaymentIntentActionResponse {
	resp := dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(a.Intent)}
	if a.Routing == nil {
		return resp
	}
	routing := &dto.RoutingExecutionResponse{
		Candidates: make([]dto.RoutingCandidateResponse, 0, len(a.Routing.Candidates)),
		Attempts:   make([]dto.ProviderAttemptResponse, 0, len(a.Routing.Attempts)),
	}
	for _, c := range a.Routing.Candidates {
		routing.Candidates = append(routing.Candidates, dto.RoutingCandidateResponse{
			ProviderCode: string(c.ProviderCode), AccountID: c.AccountID, Priority: c.Priority,
			MaxAttempts: c.MaxAttempts, Reason: c.Reason, Skipped: c.Skipped, SkipReason: c.SkipReason,
		})
	}
	for _, a := range a.Routing.Attempts {
		routing.Attempts = append(routing.Attempts, dto.ProviderAttemptResponse{
			ProviderCode: string(a.ProviderCode), AccountID: a.AccountID, Attempt: a.Attempt,
			StartedAt: a.StartedAt, Duration: a.Duration, Error: a.Error,
		})
	}
	if a.Routing.Selected.ProviderCode != "" {
		s := a.Routing.Selected
		routing.Selected = &dto.RoutingCandidateResponse{
			ProviderCode: string(s.ProviderCode), AccountID: s.AccountID, Priority: s.Priority,
			MaxAttempts: s.MaxAttempts, Reason: s.Reason, Skipped: s.Skipped, SkipReason: s.SkipReason,
		}
		p := a.Routing.Payment
		routing.Payment = &p
	}
	resp.Routing = routing
	return resp
}

func (s *ServiceImpl) executePaymentIntent(ctx context.Context, intent paymentmodel.PaymentIntents, actorID uuid.UUID) (dto.PaymentIntentActionResponse, error) {
	method, ok := normalizePaymentMethod(intent.SelectedMethodCode.String)
	if !ok {
		return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intent)}, fmt.Errorf("unsupported payment method %q", intent.SelectedMethodCode.String)
	}
	intent.Status = paymentmodel.PaymentIntentStatusProcessing
	intent.SetSignatureMetaUpdate(actorID)

	if err := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); err != nil {
		return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intent)}, err
	}
	routingRequest := RoutingRequest{
		Method:   method,
		Channel:  intent.SelectedChannelCode.String,
		Currency: intent.Currency,
		Amount:   intent.Amount.StringFixed(2),
		GatewayCall: pg.CreatePaymentRequest{
			PaymentIntentID: intent.Id.String(),
			AttemptID:       uuid.Must(uuid.NewV7()).String(),
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
	evaluatedJSON, _ := json.Marshal(map[string]any{
		"method":   method,
		"channel":  routingRequest.Channel,
		"currency": routingRequest.Currency,
		"amount":   routingRequest.Amount,
		"strategy": routingResult.Strategy,
		"ruleId":   routingResult.RuleID,
	})
	decision := paymentmodel.PaymentRouteDecisions{
		Id:                        uuid.Must(uuid.NewV7()),
		PaymentIntentId:           intent.Id,
		SelectedProviderAccountId: nuuid.From(routingResult.Selected.AccountID),
		SelectedProviderCode:      null.StringFrom(string(routingResult.Selected.ProviderCode)),
		MethodCode:                string(method),
		ChannelCode:               null.StringFrom(routingRequest.Channel),
		Reason:                    routingDecisionReason(routingResult, routeErr),
		EvaluatedContext:          evaluatedJSON,
		Candidates:                candidateJSON,
		MetaSignature:             shared.MetaSignature{MetaCreatedAt: now, MetaCreatedBy: actorID},
	}
	if err := s.paymentRepo.CreatePaymentRouteDecisions(ctx, &decision); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), err
	}

	nextAttemptNo, err := s.nextPaymentAttemptNumber(ctx, intent.Id)
	if err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), err
	}
	var successfulAttemptID uuid.UUID
	for index, providerAttempt := range routingResult.Attempts {
		attempt := paymentmodel.PaymentAttempts{
			Id:                uuid.Must(uuid.NewV7()),
			PaymentIntentId:   intent.Id,
			AttemptNo:         nextAttemptNo + index,
			ProviderAccountId: nuuid.From(providerAttempt.AccountID),
			RouteDecisionId:   nuuid.From(decision.Id),
			ProviderCode:      null.StringFrom(string(providerAttempt.ProviderCode)),
			MethodCode:        string(method),
			ChannelCode:       null.StringFrom(routingRequest.Channel),
			Amount:            intent.Amount,
			Currency:          intent.Currency,
			Status:            paymentmodel.PaymentAttemptStatusFailed,
			FailureMessage:    null.StringFrom(providerAttempt.Error),
			RawRequest:        mustJSON(routingRequest.GatewayCall),
			MetaSignature:     shared.MetaSignature{MetaCreatedAt: providerAttempt.StartedAt, MetaCreatedBy: actorID},
		}
		if providerAttempt.Error == "" && providerAttempt.ProviderCode == routingResult.Selected.ProviderCode {
			attempt.Status = paymentAttemptStatus(routingResult.Payment.Status)
			attempt.ProviderReference = null.StringFrom(routingResult.Payment.ProviderReference)
			attempt.ProviderTransactionId = null.StringFrom(routingResult.Payment.ProviderTransactionID)
			attempt.ProviderOrderId = null.StringFrom(routingResult.Payment.OrderID)
			attempt.ProviderPaymentId = null.StringFrom(routingResult.Payment.ProviderPaymentID)
			attempt.RawResponse = routingResult.Payment.Raw
			successfulAttemptID = attempt.Id
		}
		if err := s.paymentRepo.CreatePaymentAttempts(ctx, &attempt); err != nil {
			return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), err
		}
	}

	if routeErr != nil {
		intent.Status = paymentmodel.PaymentIntentStatusRequiresConfirmation
		intent.SetSignatureMetaUpdate(actorID)
		if updateErr := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); updateErr != nil {
			return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), updateErr
		}
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), routeErr
	}
	if err := s.persistPaymentInstructions(ctx, successfulAttemptID, routingResult.Payment.Instructions, actorID); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), err
	}
	intent.Status = paymentIntentStatus(routingResult.Payment.Status)
	if intent.Status == paymentmodel.PaymentIntentStatusSucceeded {
		intent.PaidAt = null.TimeFrom(now)
	}

	intent.SetSignatureMetaUpdate(actorID)

	if err := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); err != nil {
		return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), err
	}
	return PaymentIntentActionResult{Intent: intent, Routing: &routingResult}.ToResponse(), nil
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
			Id:               uuid.Must(uuid.NewV7()),
			PaymentAttemptId: attemptID,
			InstructionType:  instruction.Type,
			IsActive:         true,
			DisplayName:      null.StringFrom(instruction.DisplayName),
			AccountNumber:    null.StringFrom(instruction.AccountNumber),
			BillerCode:       null.StringFrom(instruction.BillerCode),
			PaymentCode:      null.StringFrom(instruction.PaymentCode),
			QrString:         null.StringFrom(instruction.QRString),
			QrImageUrl:       null.StringFrom(instruction.QRImageURL),
			CheckoutUrl:      null.StringFrom(instruction.CheckoutURL),
			DeeplinkUrl:      null.StringFrom(instruction.DeeplinkURL),
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
	strategy := result.Strategy
	if strategy == "" {
		strategy = "priority"
	}
	return fmt.Sprintf("selected_%s_priority_%d_strategy_%s", result.Selected.ProviderCode, result.Selected.Priority, strategy)
}

func mustJSON(value any) json.RawMessage {
	encoded, _ := json.Marshal(value)
	return encoded
}

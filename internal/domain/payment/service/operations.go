package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

func (s *ServiceImpl) CreatePaymentIntent(ctx context.Context, request dto.CreatePaymentIntentRequest) (paymentmodel.PaymentIntents, error) {
	if request.ActorID == uuid.Nil || request.MerchantID == uuid.Nil || request.SourceID == uuid.Nil {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("actorId, merchantId, and sourceId are required")
	}
	if strings.TrimSpace(request.SourceService) == "" || strings.TrimSpace(request.SourceType) == "" {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("sourceService and sourceType are required")
	}
	if request.Amount.LessThanOrEqual(decimal.Zero) || strings.TrimSpace(request.Currency) == "" || strings.TrimSpace(request.IdempotencyKey) == "" {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("positive amount, currency, and idempotencyKey are required")
	}
	if request.ExpiresAt != nil && !request.ExpiresAt.After(time.Now().UTC()) {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("expiresAt must be in the future")
	}
	if !validOptionalJSON(request.SourceSnapshot) || !validOptionalJSON(request.Metadata) {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("sourceSnapshot and metadata must contain valid JSON")
	}
	if existing, found := s.findIntentByIdempotencyKey(ctx, request.MerchantID, request.IdempotencyKey); found {
		if existing.SourceId != request.SourceID || !existing.Amount.Equal(request.Amount) || !strings.EqualFold(existing.Currency, request.Currency) {
			return paymentmodel.PaymentIntents{}, failure.Conflict("create", "payment intent", "idempotency key was used with a different request")
		}
		return existing, nil
	}

	status := paymentmodel.PaymentIntentStatusRequiresPaymentMethod
	if strings.TrimSpace(request.PaymentMethodCode) != "" {
		status = paymentmodel.PaymentIntentStatusRequiresConfirmation
	}
	now := time.Now().UTC()
	intent := paymentmodel.PaymentIntents{
		Id:                  mustUUID(),
		IntentCode:          operationCode("pi"),
		SourceService:       strings.TrimSpace(request.SourceService),
		SourceType:          strings.TrimSpace(request.SourceType),
		SourceId:            request.SourceID,
		MerchantId:          request.MerchantID,
		CustomerId:          optionalUUID(request.CustomerID),
		Amount:              request.Amount,
		Currency:            strings.ToUpper(strings.TrimSpace(request.Currency)),
		Status:              status,
		SelectedMethodCode:  optionalString(request.PaymentMethodCode),
		SelectedChannelCode: optionalString(request.PaymentChannelCode),
		Description:         optionalString(request.Description),
		IdempotencyKey:      strings.TrimSpace(request.IdempotencyKey),
		SourceSnapshot:      normalizedJSON(request.SourceSnapshot),
		Metadata:            normalizedJSON(request.Metadata),
		MetaSignature:       newCreateSignature(request.ActorID, now),
	}
	if request.ExpiresAt != nil {
		intent.ExpiresAt = null.TimeFrom(request.ExpiresAt.UTC())
	}
	if err := s.paymentRepo.CreatePaymentIntents(ctx, &intent); err != nil {
		if existing, found := s.findIntentByIdempotencyKey(ctx, request.MerchantID, request.IdempotencyKey); found {
			return existing, nil
		}
		return paymentmodel.PaymentIntents{}, err
	}
	return intent, nil
}

func (s *ServiceImpl) findIntentByIdempotencyKey(ctx context.Context, merchantID uuid.UUID, key string) (paymentmodel.PaymentIntents, bool) {
	result, err := s.paymentRepo.ResolvePaymentIntentsByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{
			{Field: string(paymentmodel.PaymentIntentsDBFieldName.MerchantId), Operator: paymentmodel.OperatorEqual, Value: merchantID},
			{Field: string(paymentmodel.PaymentIntentsDBFieldName.IdempotencyKey), Operator: paymentmodel.OperatorEqual, Value: strings.TrimSpace(key)},
		},
		Pagination: paymentmodel.Pagination{Page: 1, PageSize: 1},
	})
	if err != nil || len(result) == 0 {
		return paymentmodel.PaymentIntents{}, false
	}
	return result[0].PaymentIntents, true
}

func (s *ServiceImpl) GetPaymentIntent(ctx context.Context, id uuid.UUID) (paymentmodel.PaymentIntents, error) {
	if id == uuid.Nil {
		return paymentmodel.PaymentIntents{}, failure.BadRequestFromString("payment intent id is required")
	}
	return s.paymentRepo.ResolvePaymentIntentsByID(ctx, paymentmodel.PaymentIntentsPrimaryID{Id: id})
}

func (s *ServiceImpl) ApplyPaymentIntentAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (PaymentIntentActionResult, error) {
	if err := validateActor(command.ActorID); err != nil {
		return PaymentIntentActionResult{}, err
	}
	unlock, acquired, err := s.executionLocker.TryLock(ctx, id.String(), 2*time.Minute)
	if err != nil {
		return PaymentIntentActionResult{}, failure.InternalError(err)
	}
	if !acquired {
		return PaymentIntentActionResult{}, failure.Conflict(action, "payment intent", "another action is already in progress")
	}
	defer unlock()
	intent, err := s.GetPaymentIntent(ctx, id)
	if err != nil {
		return PaymentIntentActionResult{}, err
	}
	now := time.Now().UTC()
	switch action {
	case "confirm":
		if intent.Status != paymentmodel.PaymentIntentStatusRequiresConfirmation {
			return PaymentIntentActionResult{Intent: intent}, invalidAction(action, string(intent.Status))
		}
		return s.executePaymentIntent(ctx, intent, command.ActorID)
	case "cancel":
		if intent.Status == paymentmodel.PaymentIntentStatusSucceeded || intent.Status == paymentmodel.PaymentIntentStatusCanceled {
			return PaymentIntentActionResult{Intent: intent}, invalidAction(action, string(intent.Status))
		}
		intent.Status = paymentmodel.PaymentIntentStatusCanceled
		intent.CanceledAt = null.TimeFrom(now)
		intent.CancellationReason = optionalString(command.Reason)
	default:
		return PaymentIntentActionResult{Intent: intent}, unsupportedAction(action)
	}
	setUpdateSignature(&intent.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), &intent); err != nil {
		return PaymentIntentActionResult{Intent: intent}, err
	}
	return PaymentIntentActionResult{Intent: intent}, nil
}

func (s *ServiceImpl) CreateRefund(ctx context.Context, request dto.CreateRefundRequest) (paymentmodel.PaymentRefunds, error) {
	if request.ActorID == uuid.Nil || request.PaymentIntentID == uuid.Nil || request.Amount.LessThanOrEqual(decimal.Zero) || strings.TrimSpace(request.IdempotencyKey) == "" {
		return paymentmodel.PaymentRefunds{}, failure.BadRequestFromString("actorId, paymentIntentId, positive amount, and idempotencyKey are required")
	}
	if !validOptionalJSON(request.Metadata) {
		return paymentmodel.PaymentRefunds{}, failure.BadRequestFromString("metadata must contain valid JSON")
	}
	intent, err := s.GetPaymentIntent(ctx, request.PaymentIntentID)
	if err != nil {
		return paymentmodel.PaymentRefunds{}, err
	}
	if intent.Status != paymentmodel.PaymentIntentStatusSucceeded {
		return paymentmodel.PaymentRefunds{}, invalidAction("refund", string(intent.Status))
	}
	if request.Amount.GreaterThan(intent.Amount) {
		return paymentmodel.PaymentRefunds{}, failure.BadRequestFromString("refund amount exceeds payment intent amount")
	}
	currency := strings.ToUpper(strings.TrimSpace(request.Currency))
	if currency == "" {
		currency = intent.Currency
	}
	if !strings.EqualFold(currency, intent.Currency) {
		return paymentmodel.PaymentRefunds{}, failure.BadRequestFromString("refund currency must match payment intent currency")
	}
	now := time.Now()
	refundCode := stableOperationCode("rf", request.PaymentIntentID.String(), request.IdempotencyKey)
	if existing, found := s.findRefundByCode(ctx, refundCode); found {
		return existing, nil
	}
	refund := paymentmodel.PaymentRefunds{
		Id:               uuid.Must(uuid.NewV7()),
		PaymentIntentId:  request.PaymentIntentID,
		PaymentAttemptId: optionalUUID(request.PaymentAttemptID),
		RefundCode:       refundCode,
		Amount:           request.Amount,
		Currency:         currency,
		Reason:           optionalString(request.Reason),
		Status:           paymentmodel.PaymentRefundStatusRequested,
		RequestedBy:      request.ActorID,
		RequestedAt:      now,
		Metadata:         mergeIdempotencyMetadata(request.Metadata, request.IdempotencyKey),
		MetaSignature:    newCreateSignature(request.ActorID, now),
	}
	err = s.paymentRepo.CreatePaymentRefunds(ctx, &refund)
	if err != nil {
		// The unique refund code closes the race between concurrent retries.
		if existing, found := s.findRefundByCode(ctx, refundCode); found {
			return existing, nil
		}
		return paymentmodel.PaymentRefunds{}, err
	}
	return refund, nil
}

func (s *ServiceImpl) findRefundByCode(ctx context.Context, code string) (paymentmodel.PaymentRefunds, bool) {
	result, err := s.paymentRepo.ResolvePaymentRefundsByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{{Field: string(paymentmodel.PaymentRefundsDBFieldName.RefundCode), Operator: paymentmodel.OperatorEqual, Value: code}},
		Pagination:   paymentmodel.Pagination{Page: 1, PageSize: 1},
	})
	if err != nil || len(result) == 0 {
		return paymentmodel.PaymentRefunds{}, false
	}
	return result[0].PaymentRefunds, true
}

func (s *ServiceImpl) ApplyRefundAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.PaymentRefunds, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.PaymentRefunds{}, err
	}
	item, err := s.paymentRepo.ResolvePaymentRefundsByID(ctx, paymentmodel.PaymentRefundsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	now := time.Now().UTC()
	switch action {
	case "approve":
		if item.Status != paymentmodel.PaymentRefundStatusRequested {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.ApprovedBy, item.ApprovedAt = paymentmodel.PaymentRefundStatusApproved, nuuid.From(command.ActorID), null.TimeFrom(now)
	case "reject":
		if item.Status != paymentmodel.PaymentRefundStatusRequested {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.RejectedBy, item.RejectedAt = paymentmodel.PaymentRefundStatusRejected, nuuid.From(command.ActorID), null.TimeFrom(now)
		item.RejectionReason = optionalString(command.Reason)
	case "process":
		if item.Status != paymentmodel.PaymentRefundStatusApproved {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.ProcessingAt = paymentmodel.PaymentRefundStatusProcessing, null.TimeFrom(now)
	case "succeed":
		if item.Status != paymentmodel.PaymentRefundStatusProcessing {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.SucceededAt = paymentmodel.PaymentRefundStatusSucceeded, null.TimeFrom(now)
	case "fail":
		if item.Status != paymentmodel.PaymentRefundStatusProcessing {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.FailedAt = paymentmodel.PaymentRefundStatusFailed, null.TimeFrom(now)
		item.FailureCode, item.FailureMessage = optionalString(command.FailureCode), optionalString(command.FailureMessage)
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdatePaymentRefundsByID(ctx, item.ToPaymentRefundsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyWebhookAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.ProviderWebhookEvents, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.ProviderWebhookEvents{}, err
	}
	item, err := s.paymentRepo.ResolveProviderWebhookEventsByID(ctx, paymentmodel.ProviderWebhookEventsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	now := time.Now().UTC()
	switch action {
	case "retry":
		if item.ProcessingStatus != paymentmodel.WebhookProcessingStatusFailed {
			return item, invalidAction(action, string(item.ProcessingStatus))
		}
		item.ProcessingStatus, item.NextRetryAt, item.ErrorCode, item.ErrorMessage = paymentmodel.WebhookProcessingStatusReceived, null.Time{}, null.String{}, null.String{}
	case "ignore":
		if item.ProcessingStatus == paymentmodel.WebhookProcessingStatusProcessed {
			return item, invalidAction(action, string(item.ProcessingStatus))
		}
		item.ProcessingStatus, item.ProcessedAt = paymentmodel.WebhookProcessingStatusProcessed, null.TimeFrom(now)
		item.ErrorMessage = optionalString(command.Reason)
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdateProviderWebhookEventsByID(ctx, item.ToProviderWebhookEventsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyManualEvidenceAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.ManualPaymentEvidence, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.ManualPaymentEvidence{}, err
	}
	item, err := s.paymentRepo.ResolveManualPaymentEvidenceByID(ctx, paymentmodel.ManualPaymentEvidencePrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	now := time.Now().UTC()
	switch action {
	case "review":
		if item.Status != paymentmodel.ManualEvidenceStatusSubmitted {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status = paymentmodel.ManualEvidenceStatusUnderReview
	case "approve", "reject":
		if item.Status != paymentmodel.ManualEvidenceStatusUnderReview {
			return item, invalidAction(action, string(item.Status))
		}
		item.ReviewedBy, item.ReviewedAt = nuuid.From(command.ActorID), null.TimeFrom(now)
		if action == "approve" {
			item.Status = paymentmodel.ManualEvidenceStatusApproved
		} else {
			item.Status, item.RejectionReason = paymentmodel.ManualEvidenceStatusRejected, optionalString(command.Reason)
		}
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdateManualPaymentEvidenceByID(ctx, item.ToManualPaymentEvidencePrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyOverpaymentAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.PaymentOverpayments, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.PaymentOverpayments{}, err
	}
	item, err := s.paymentRepo.ResolvePaymentOverpaymentsByID(ctx, paymentmodel.PaymentOverpaymentsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	if item.Status == "resolved" {
		return item, invalidAction(action, item.Status)
	}
	allowed := map[string]bool{"refund": true, "credit_balance": true, "apply_next_invoice": true, "write_off": true}
	if !allowed[action] {
		return item, unsupportedAction(action)
	}
	now := time.Now().UTC()
	item.Status, item.ResolutionAction, item.ResolutionNotes = "resolved", null.StringFrom(action), optionalString(command.Notes)
	item.ResolvedAt, item.ResolvedBy = null.TimeFrom(now), nuuid.From(command.ActorID)
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdatePaymentOverpaymentsByID(ctx, item.ToPaymentOverpaymentsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) OpenCashSession(ctx context.Context, request dto.OpenCashSessionRequest) (paymentmodel.CashCollectionSessions, error) {
	if request.ActorID == uuid.Nil || request.MerchantID == uuid.Nil || request.CollectorID == uuid.Nil || request.OpeningFloatAmount.IsNegative() || strings.TrimSpace(request.Currency) == "" {
		return paymentmodel.CashCollectionSessions{}, failure.BadRequestFromString("actorId, merchantId, collectorId, non-negative openingFloatAmount, and currency are required")
	}
	if !validOptionalJSON(request.Metadata) {
		return paymentmodel.CashCollectionSessions{}, failure.BadRequestFromString("metadata must contain valid JSON")
	}
	now := time.Now().UTC()
	item := paymentmodel.CashCollectionSessions{Id: mustUUID(), SessionCode: operationCode("cash"), MerchantId: request.MerchantID, CollectorId: request.CollectorID, LocationId: optionalUUID(request.LocationID), OpenedAt: now, Status: paymentmodel.CashSessionStatusOpen, OpeningFloatAmount: request.OpeningFloatAmount, ExpectedAmount: decimal.Zero, CountedAmount: decimal.Zero, VarianceAmount: decimal.Zero, Currency: strings.ToUpper(request.Currency), Notes: optionalString(request.Notes), Metadata: normalizedJSON(request.Metadata), MetaSignature: newCreateSignature(request.ActorID, now)}
	if err := s.paymentRepo.CreateCashCollectionSessions(ctx, &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyCashSessionAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.CashCollectionSessions, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.CashCollectionSessions{}, err
	}
	item, err := s.paymentRepo.ResolveCashCollectionSessionsByID(ctx, paymentmodel.CashCollectionSessionsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	if item.Status != paymentmodel.CashSessionStatusOpen {
		return item, invalidAction(action, string(item.Status))
	}
	now := time.Now().UTC()
	switch action {
	case "close":
		if command.Amount.IsNegative() {
			return item, failure.BadRequestFromString("amount cannot be negative")
		}
		item.Status, item.ClosedAt, item.CountedAmount = paymentmodel.CashSessionStatusClosed, null.TimeFrom(now), command.Amount
		item.VarianceAmount = command.Amount.Sub(item.ExpectedAmount)
	case "cancel":
		item.Status, item.ClosedAt = paymentmodel.CashSessionStatusCanceled, null.TimeFrom(now)
		item.Notes = optionalString(command.Reason)
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdateCashCollectionSessionsByID(ctx, item.ToCashCollectionSessionsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyInstallmentAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.PaymentInstallments, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.PaymentInstallments{}, err
	}
	item, err := s.paymentRepo.ResolvePaymentInstallmentsByID(ctx, paymentmodel.PaymentInstallmentsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	if item.Status != paymentmodel.PaymentInstallmentStatusPending {
		return item, invalidAction(action, string(item.Status))
	}
	now := time.Now().UTC()
	switch action {
	case "pay":
		paid := command.Amount
		if paid.IsZero() {
			paid = item.DueAmount
		}
		if paid.LessThan(item.DueAmount) {
			return item, failure.BadRequestFromString("paid amount must cover due amount")
		}
		item.Status, item.PaidAmount, item.PaidAt = paymentmodel.PaymentInstallmentStatusPaid, paid, null.TimeFrom(now)
	case "mark-overdue":
		item.Status, item.OverdueAt = paymentmodel.PaymentInstallmentStatusOverdue, null.TimeFrom(now)
	case "cancel":
		item.Status = paymentmodel.PaymentInstallmentStatusCanceled
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdatePaymentInstallmentsByID(ctx, item.ToPaymentInstallmentsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (s *ServiceImpl) ApplyAuthorizationAction(ctx context.Context, id uuid.UUID, action string, command dto.ActionCommand) (paymentmodel.PaymentAuthorizations, error) {
	if err := validateActor(command.ActorID); err != nil {
		return paymentmodel.PaymentAuthorizations{}, err
	}
	item, err := s.paymentRepo.ResolvePaymentAuthorizationsByID(ctx, paymentmodel.PaymentAuthorizationsPrimaryID{Id: id})
	if err != nil {
		return item, err
	}
	now := time.Now().UTC()
	switch action {
	case "authorize":
		if item.Status != paymentmodel.PaymentAuthorizationStatusRequested {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.AuthorizedAt = paymentmodel.PaymentAuthorizationStatusAuthorized, null.TimeFrom(now)
	case "capture":
		if item.Status != paymentmodel.PaymentAuthorizationStatusAuthorized {
			return item, invalidAction(action, string(item.Status))
		}
		amount := command.Amount
		if amount.IsZero() {
			amount = item.Amount
		}
		if amount.IsNegative() || amount.GreaterThan(item.Amount.Sub(item.CapturedAmount)) {
			return item, failure.BadRequestFromString("capture amount exceeds remaining authorization")
		}
		item.Status, item.CapturedAmount = paymentmodel.PaymentAuthorizationStatusCaptured, item.CapturedAmount.Add(amount)
	case "void":
		if item.Status != paymentmodel.PaymentAuthorizationStatusRequested && item.Status != paymentmodel.PaymentAuthorizationStatusAuthorized {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status = paymentmodel.PaymentAuthorizationStatusVoided
	case "fail":
		if item.Status == paymentmodel.PaymentAuthorizationStatusCaptured || item.Status == paymentmodel.PaymentAuthorizationStatusVoided {
			return item, invalidAction(action, string(item.Status))
		}
		item.Status, item.FailureCode, item.FailureMessage = paymentmodel.PaymentAuthorizationStatusFailed, optionalString(command.FailureCode), optionalString(command.FailureMessage)
	default:
		return item, unsupportedAction(action)
	}
	setUpdateSignature(&item.MetaSignature, command.ActorID, now)
	if err := s.paymentRepo.UpdatePaymentAuthorizationsByID(ctx, item.ToPaymentAuthorizationsPrimaryID(), &item); err != nil {
		return item, err
	}
	return item, nil
}

func validateActor(actorID uuid.UUID) error {
	if actorID == uuid.Nil {
		return failure.BadRequestFromString("actorId is required")
	}
	return nil
}

func mustUUID() uuid.UUID { id, _ := uuid.NewV4(); return id }
func operationCode(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(mustUUID().String(), "-", "")[:20])
}
func stableOperationCode(prefix string, parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, ":")))
	return fmt.Sprintf("%s_%x", prefix, sum[:10])
}
func optionalUUID(id uuid.UUID) nuuid.NUUID {
	if id == uuid.Nil {
		return nuuid.NUUID{}
	}
	return nuuid.From(id)
}
func optionalString(value string) null.String {
	value = strings.TrimSpace(value)
	if value == "" {
		return null.String{}
	}
	return null.StringFrom(value)
}
func normalizedJSON(value json.RawMessage) json.RawMessage {
	if len(value) == 0 {
		return json.RawMessage(`{}`)
	}
	return value
}
func validOptionalJSON(value json.RawMessage) bool { return len(value) == 0 || json.Valid(value) }
func newCreateSignature(actor uuid.UUID, now time.Time) shared.MetaSignature {
	return shared.MetaSignature{
		MetaCreatedAt: now,
		MetaCreatedBy: actor,
		MetaUpdatedAt: null.TimeFrom(now),
		MetaUpdatedBy: &actor,
	}
}
func setUpdateSignature(meta *shared.MetaSignature, actor uuid.UUID, now time.Time) {
	meta.MetaUpdatedAt, meta.MetaUpdatedBy = null.TimeFrom(now), &actor
}
func unsupportedAction(action string) error {
	return failure.BadRequestFromString(fmt.Sprintf("unsupported action %q", action))
}
func invalidAction(action, status string) error {
	return failure.Conflict(action, "payment resource", fmt.Sprintf("action is not allowed from status %s", status))
}

func mergeIdempotencyMetadata(raw json.RawMessage, key string) json.RawMessage {
	metadata := map[string]any{}
	if len(raw) > 0 {
		_ = json.Unmarshal(raw, &metadata)
	}
	metadata["idempotencyKey"] = strings.TrimSpace(key)
	encoded, _ := json.Marshal(metadata)
	return encoded
}

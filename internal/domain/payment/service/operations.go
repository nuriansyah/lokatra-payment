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

func (s *ServiceImpl) CreatePaymentIntent(ctx context.Context, req dto.CreatePaymentIntentRequest) (dto.PaymentIntentsResponse, error) {
	if strings.TrimSpace(req.IdempotencyKey) == "" {
		return dto.PaymentIntentsResponse{}, failure.WithCode(shared.ErrPayIdempotencyKeyRequired, "idempotencyKey is required")
	}

	if err := validateCurrency(req.Currency); err != nil {
		return dto.PaymentIntentsResponse{}, err
	}
	if err := validateAmount(req.Amount, req.Currency); err != nil {
		return dto.PaymentIntentsResponse{}, err
	}

	if existing, found := s.findIntentByIdempotencyKey(ctx, req.MerchantID, req.IdempotencyKey); found {
		if existing.SourceId != req.SourceID || !existing.Amount.Equal(req.Amount) || !strings.EqualFold(existing.Currency, req.Currency) {
			return dto.PaymentIntentsResponse{}, failure.WithCode(shared.ErrPayIdempotencyKeyReusedDifferentPayload, "idempotency key was used with a different request")
		}
		return dto.NewPaymentIntentsResponse(existing), nil
	}

	status := paymentmodel.PaymentIntentStatusRequiresPaymentMethod
	if strings.TrimSpace(req.PaymentMethodCode) != "" {
		status = paymentmodel.PaymentIntentStatusRequiresConfirmation
	}

	now := time.Now()
	intent := paymentmodel.PaymentIntents{
		Id:                  uuid.Must(uuid.NewV7()),
		IntentCode:          operationCode("pi"),
		SourceService:       strings.TrimSpace(req.SourceService),
		SourceType:          strings.TrimSpace(req.SourceType),
		SourceId:            req.SourceID,
		MerchantId:          req.MerchantID,
		CustomerId:          nuuid.From(req.CustomerID),
		Amount:              req.Amount,
		Currency:            strings.ToUpper(strings.TrimSpace(req.Currency)),
		Status:              status,
		SelectedMethodCode:  null.StringFrom(req.PaymentMethodCode),
		SelectedChannelCode: null.StringFrom(req.PaymentChannelCode),
		Description:         null.StringFrom(req.Description),
		IdempotencyKey:      strings.TrimSpace(req.IdempotencyKey),
		SourceSnapshot:      normalizedJSON(req.SourceSnapshot),
		MetaSignature: shared.MetaSignature{
			MetaCreatedBy: req.ActorID,
			MetaCreatedAt: now,
		},
	}
	if req.ExpiresAt != nil {
		intent.ExpiresAt = null.TimeFrom(req.ExpiresAt.UTC())
	}
	if err := s.paymentRepo.CreatePaymentIntents(ctx, &intent); err != nil {
		if existing, found := s.findIntentByIdempotencyKey(ctx, req.MerchantID, req.IdempotencyKey); found {
			return dto.NewPaymentIntentsResponse(existing), nil
		}
		return dto.PaymentIntentsResponse{}, err
	}
	return dto.NewPaymentIntentsResponse(intent), nil
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

func (s *ServiceImpl) GetPaymentIntent(ctx context.Context, id uuid.UUID) (dto.PaymentIntentsResponse, error) {
	if id == uuid.Nil {
		return dto.PaymentIntentsResponse{}, failure.WithCode(shared.ErrInvalidID, "payment intent id is required")
	}
	intent, err := s.paymentRepo.ResolvePaymentIntentsByID(ctx, paymentmodel.PaymentIntentsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentIntentsResponse{}, err
	}
	return dto.NewPaymentIntentsResponse(intent), nil
}

func (s *ServiceImpl) ApplyPaymentIntentAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.PaymentIntentActionResponse, error) {
	unlock, acquired, err := s.executionLocker.TryLock(ctx, id.String(), 2*time.Minute)
	if err != nil {
		return dto.PaymentIntentActionResponse{}, failure.InternalError(err)
	}
	if !acquired {
		return dto.PaymentIntentActionResponse{}, failure.Conflict(action, "payment intent", "another action is already in progress")
	}
	defer unlock()

	intentModel, err := s.paymentRepo.ResolvePaymentIntentsByID(ctx, paymentmodel.PaymentIntentsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentIntentActionResponse{}, err
	}
	now := time.Now().UTC()

	switch action {
	case "confirm":
		if intentModel.Status != paymentmodel.PaymentIntentStatusRequiresConfirmation {
			return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intentModel)}, invalidAction(action, string(intentModel.Status))
		}
		return s.executePaymentIntent(ctx, intentModel, req.ActorID)
	case "cancel":
		if intentModel.Status == paymentmodel.PaymentIntentStatusSucceeded || intentModel.Status == paymentmodel.PaymentIntentStatusCanceled {
			return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intentModel)}, invalidAction(action, string(intentModel.Status))
		}
		intentModel.Status = paymentmodel.PaymentIntentStatusCanceled
		intentModel.CanceledAt = null.TimeFrom(now)
		intentModel.CancellationReason = null.StringFrom(req.Reason)
	default:
		return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intentModel)}, unsupportedAction(action)
	}

	intentModel.SetSignatureMetaUpdate(req.ActorID)

	err = s.paymentRepo.UpdatePaymentIntentsByID(ctx, intentModel.ToPaymentIntentsPrimaryID(), &intentModel)
	if err != nil {
		return dto.PaymentIntentActionResponse{}, err
	}
	return dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(intentModel)}, nil
}

func (s *ServiceImpl) CreateRefund(ctx context.Context, req dto.CreateRefundRequest) (dto.PaymentRefundsResponse, error) {
	if strings.TrimSpace(req.IdempotencyKey) == "" {
		return dto.PaymentRefundsResponse{}, failure.WithCode(shared.ErrPayIdempotencyKeyRequired, "idempotencyKey is required")
	}

	intentResp, err := s.GetPaymentIntent(ctx, req.PaymentIntentID)
	if err != nil {
		return dto.PaymentRefundsResponse{}, err
	}
	if intentResp.Status != paymentmodel.PaymentIntentStatusSucceeded {
		return dto.PaymentRefundsResponse{}, invalidAction("refund", string(intentResp.Status))
	}
	if req.Amount.GreaterThan(intentResp.Amount) {
		return dto.PaymentRefundsResponse{}, failure.BadRequestFromString("refund amount exceeds payment intent amount")
	}
	currency := strings.ToUpper(strings.TrimSpace(req.Currency))
	if currency == "" {
		currency = intentResp.Currency
	}
	if !strings.EqualFold(currency, intentResp.Currency) {
		return dto.PaymentRefundsResponse{}, failure.BadRequestFromString("refund currency must match payment intent currency")
	}

	now := time.Now()
	refundCode := stableOperationCode("rf", req.PaymentIntentID.String(), req.IdempotencyKey)
	if existing, found := s.findRefundByCode(ctx, refundCode); found {
		return dto.NewPaymentRefundsResponse(existing), nil
	}
	refund := paymentmodel.PaymentRefunds{
		Id:               uuid.Must(uuid.NewV7()),
		PaymentIntentId:  req.PaymentIntentID,
		PaymentAttemptId: nuuid.From(req.PaymentAttemptID),
		RefundCode:       refundCode,
		Amount:           req.Amount,
		Currency:         currency,
		Reason:           null.StringFrom(req.Reason),
		Status:           paymentmodel.PaymentRefundStatusRequested,
		RequestedBy:      req.ActorID,
		RequestedAt:      now,
		MetaSignature: shared.MetaSignature{
			MetaCreatedBy: req.ActorID,
			MetaCreatedAt: now,
		},
	}
	err = s.paymentRepo.CreatePaymentRefunds(ctx, &refund)
	if err != nil {
		if existing, found := s.findRefundByCode(ctx, refundCode); found {
			return dto.NewPaymentRefundsResponse(existing), nil
		}
		return dto.PaymentRefundsResponse{}, err
	}
	return dto.NewPaymentRefundsResponse(refund), nil
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

func (s *ServiceImpl) ApplyRefundAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.PaymentRefundsResponse, error) {
	item, err := s.paymentRepo.ResolvePaymentRefundsByID(ctx, paymentmodel.PaymentRefundsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentRefundsResponse{}, err
	}
	now := time.Now().UTC()
	switch action {
	case "approve":
		if item.Status != paymentmodel.PaymentRefundStatusRequested {
			return dto.NewPaymentRefundsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.ApprovedBy, item.ApprovedAt = paymentmodel.PaymentRefundStatusApproved, nuuid.From(req.ActorID), null.TimeFrom(now)
	case "reject":
		if item.Status != paymentmodel.PaymentRefundStatusRequested {
			return dto.NewPaymentRefundsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.RejectedBy, item.RejectedAt = paymentmodel.PaymentRefundStatusRejected, nuuid.From(req.ActorID), null.TimeFrom(now)
		item.RejectionReason = null.StringFrom(req.Reason)
	case "process":
		if item.Status != paymentmodel.PaymentRefundStatusApproved {
			return dto.NewPaymentRefundsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.ProcessingAt = paymentmodel.PaymentRefundStatusProcessing, null.TimeFrom(now)
	case "succeed":
		if item.Status != paymentmodel.PaymentRefundStatusProcessing {
			return dto.NewPaymentRefundsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.SucceededAt = paymentmodel.PaymentRefundStatusSucceeded, null.TimeFrom(now)
	case "fail":
		if item.Status != paymentmodel.PaymentRefundStatusProcessing {
			return dto.NewPaymentRefundsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.FailedAt = paymentmodel.PaymentRefundStatusFailed, null.TimeFrom(now)
		item.FailureCode, item.FailureMessage = null.StringFrom(req.FailureCode), null.StringFrom(req.FailureMessage)
	default:
		return dto.NewPaymentRefundsResponse(item), unsupportedAction(action)
	}

	item.SetSignatureMetaUpdate(req.ActorID)

	err = s.paymentRepo.UpdatePaymentRefundsByID(ctx, item.ToPaymentRefundsPrimaryID(), &item)
	if err != nil {
		return dto.PaymentRefundsResponse{}, err
	}

	return dto.NewPaymentRefundsResponse(item), nil
}

func (s *ServiceImpl) ApplyWebhookAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.ProviderWebhookEventsResponse, error) {
	item, err := s.paymentRepo.ResolveProviderWebhookEventsByID(ctx, paymentmodel.ProviderWebhookEventsPrimaryID{Id: id})
	if err != nil {
		return dto.ProviderWebhookEventsResponse{}, err
	}
	now := time.Now().UTC()
	switch action {
	case "retry":
		if item.ProcessingStatus != paymentmodel.WebhookProcessingStatusFailed {
			return dto.NewProviderWebhookEventsResponse(item), invalidAction(action, string(item.ProcessingStatus))
		}
		item.ProcessingStatus, item.NextRetryAt, item.ErrorCode, item.ErrorMessage = paymentmodel.WebhookProcessingStatusReceived, null.Time{}, null.String{}, null.String{}
	case "ignore":
		if item.ProcessingStatus == paymentmodel.WebhookProcessingStatusProcessed {
			return dto.NewProviderWebhookEventsResponse(item), invalidAction(action, string(item.ProcessingStatus))
		}
		item.ProcessingStatus, item.ProcessedAt = paymentmodel.WebhookProcessingStatusProcessed, null.TimeFrom(now)
		item.ErrorMessage = null.StringFrom(req.Reason)
	default:
		return dto.NewProviderWebhookEventsResponse(item), unsupportedAction(action)
	}

	item.SetSignatureMetaUpdate(req.ActorID)

	err = s.paymentRepo.UpdateProviderWebhookEventsByID(ctx, item.ToProviderWebhookEventsPrimaryID(), &item)
	if err != nil {
		return dto.ProviderWebhookEventsResponse{}, err
	}
	return dto.NewProviderWebhookEventsResponse(item), nil
}

func (s *ServiceImpl) ApplyManualEvidenceAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.ManualPaymentEvidenceResponse, error) {
	item, err := s.paymentRepo.ResolveManualPaymentEvidenceByID(ctx, paymentmodel.ManualPaymentEvidencePrimaryID{Id: id})
	if err != nil {
		return dto.ManualPaymentEvidenceResponse{}, err
	}
	now := time.Now().UTC()
	switch action {
	case "review":
		if item.Status != paymentmodel.ManualEvidenceStatusSubmitted {
			return dto.NewManualPaymentEvidenceResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status = paymentmodel.ManualEvidenceStatusUnderReview
	case "approve", "reject":
		if item.Status != paymentmodel.ManualEvidenceStatusUnderReview {
			return dto.NewManualPaymentEvidenceResponse(item), invalidAction(action, string(item.Status))
		}
		item.ReviewedBy, item.ReviewedAt = nuuid.From(req.ActorID), null.TimeFrom(now)
		if action == "approve" {
			item.Status = paymentmodel.ManualEvidenceStatusApproved
		} else {
			item.Status, item.RejectionReason = paymentmodel.ManualEvidenceStatusRejected, null.StringFrom(req.Reason)
		}
	default:
		return dto.NewManualPaymentEvidenceResponse(item), unsupportedAction(action)
	}

	item.SetSignatureMetaUpdate(req.ActorID)

	err = s.paymentRepo.UpdateManualPaymentEvidenceByID(ctx, item.ToManualPaymentEvidencePrimaryID(), &item)
	if err != nil {
		return dto.ManualPaymentEvidenceResponse{}, err
	}
	return dto.NewManualPaymentEvidenceResponse(item), nil
}

func (s *ServiceImpl) ApplyOverpaymentAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.PaymentOverpaymentsResponse, error) {
	item, err := s.paymentRepo.ResolvePaymentOverpaymentsByID(ctx, paymentmodel.PaymentOverpaymentsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentOverpaymentsResponse{}, err
	}
	if item.Status == "resolved" {
		return dto.NewPaymentOverpaymentsResponse(item), invalidAction(action, item.Status)
	}
	allowed := map[string]bool{"refund": true, "credit_balance": true, "apply_next_invoice": true, "write_off": true}
	if !allowed[action] {
		return dto.NewPaymentOverpaymentsResponse(item), unsupportedAction(action)
	}

	now := time.Now().UTC()
	item.Status, item.ResolutionAction, item.ResolutionNotes = "resolved", null.StringFrom(action), null.StringFrom(req.Notes)
	item.ResolvedAt, item.ResolvedBy = null.TimeFrom(now), nuuid.From(req.ActorID)
	item.SetSignatureMetaUpdate(req.ActorID)

	if err := s.paymentRepo.UpdatePaymentOverpaymentsByID(ctx, item.ToPaymentOverpaymentsPrimaryID(), &item); err != nil {
		return dto.PaymentOverpaymentsResponse{}, err
	}
	return dto.NewPaymentOverpaymentsResponse(item), nil
}

func (s *ServiceImpl) OpenCashSession(ctx context.Context, req dto.OpenCashSessionRequest) (dto.CashCollectionSessionsResponse, error) {
	now := time.Now().UTC()
	item := paymentmodel.CashCollectionSessions{
		Id:                 uuid.Must(uuid.NewV7()),
		SessionCode:        operationCode("cash"),
		MerchantId:         req.MerchantID,
		CollectorId:        req.CollectorID,
		LocationId:         nuuid.From(req.LocationID),
		OpenedAt:           now,
		Status:             paymentmodel.CashSessionStatusOpen,
		OpeningFloatAmount: req.OpeningFloatAmount,
		ExpectedAmount:     decimal.Zero,
		CountedAmount:      decimal.Zero,
		VarianceAmount:     decimal.Zero,
		Currency:           strings.ToUpper(req.Currency),
		Notes:              null.StringFrom(req.Notes),
	}

	item.SetSignatureMetaCreate(req.ActorID)

	err := s.paymentRepo.CreateCashCollectionSessions(ctx, &item)
	if err != nil {
		return dto.CashCollectionSessionsResponse{}, err
	}
	return dto.NewCashCollectionSessionsResponse(item), nil
}

func (s *ServiceImpl) ApplyCashSessionAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.CashCollectionSessionsResponse, error) {
	item, err := s.paymentRepo.ResolveCashCollectionSessionsByID(ctx, paymentmodel.CashCollectionSessionsPrimaryID{Id: id})
	if err != nil {
		return dto.CashCollectionSessionsResponse{}, err
	}
	if item.Status != paymentmodel.CashSessionStatusOpen {
		return dto.NewCashCollectionSessionsResponse(item), invalidAction(action, string(item.Status))
	}
	now := time.Now().UTC()
	switch action {
	case "close":
		if req.Amount.IsNegative() {
			return dto.CashCollectionSessionsResponse{}, failure.BadRequestFromString("amount cannot be negative")
		}
		item.Status, item.ClosedAt, item.CountedAmount = paymentmodel.CashSessionStatusClosed, null.TimeFrom(now), req.Amount
		item.VarianceAmount = req.Amount.Sub(item.ExpectedAmount)
	case "cancel":
		item.Status, item.ClosedAt = paymentmodel.CashSessionStatusCanceled, null.TimeFrom(now)
		item.Notes = null.StringFrom(req.Reason)
	default:
		return dto.NewCashCollectionSessionsResponse(item), unsupportedAction(action)
	}
	item.SetSignatureMetaUpdate(req.ActorID)
	if err := s.paymentRepo.UpdateCashCollectionSessionsByID(ctx, item.ToCashCollectionSessionsPrimaryID(), &item); err != nil {
		return dto.CashCollectionSessionsResponse{}, err
	}
	return dto.NewCashCollectionSessionsResponse(item), nil
}

func (s *ServiceImpl) ApplyInstallmentAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.PaymentInstallmentsResponse, error) {
	item, err := s.paymentRepo.ResolvePaymentInstallmentsByID(ctx, paymentmodel.PaymentInstallmentsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentInstallmentsResponse{}, err
	}
	if item.Status != paymentmodel.PaymentInstallmentStatusPending {
		return dto.NewPaymentInstallmentsResponse(item), invalidAction(action, string(item.Status))
	}
	now := time.Now().UTC()
	switch action {
	case "pay":
		paid := req.Amount
		if paid.IsZero() {
			paid = item.DueAmount
		}
		if paid.LessThan(item.DueAmount) {
			return dto.PaymentInstallmentsResponse{}, failure.BadRequestFromString("paid amount must cover due amount")
		}
		item.Status, item.PaidAmount, item.PaidAt = paymentmodel.PaymentInstallmentStatusPaid, paid, null.TimeFrom(now)
	case "mark-overdue":
		item.Status, item.OverdueAt = paymentmodel.PaymentInstallmentStatusOverdue, null.TimeFrom(now)
	case "cancel":
		item.Status = paymentmodel.PaymentInstallmentStatusCanceled
	default:
		return dto.NewPaymentInstallmentsResponse(item), unsupportedAction(action)
	}

	item.SetSignatureMetaUpdate(req.ActorID)
	if err := s.paymentRepo.UpdatePaymentInstallmentsByID(ctx, item.ToPaymentInstallmentsPrimaryID(), &item); err != nil {
		return dto.PaymentInstallmentsResponse{}, err
	}
	return dto.NewPaymentInstallmentsResponse(item), nil
}

func (s *ServiceImpl) ApplyAuthorizationAction(ctx context.Context, id uuid.UUID, action string, req dto.ActionCommand) (dto.PaymentAuthorizationsResponse, error) {
	item, err := s.paymentRepo.ResolvePaymentAuthorizationsByID(ctx, paymentmodel.PaymentAuthorizationsPrimaryID{Id: id})
	if err != nil {
		return dto.PaymentAuthorizationsResponse{}, err
	}
	now := time.Now().UTC()
	switch action {
	case "authorize":
		if item.Status != paymentmodel.PaymentAuthorizationStatusRequested {
			return dto.NewPaymentAuthorizationsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.AuthorizedAt = paymentmodel.PaymentAuthorizationStatusAuthorized, null.TimeFrom(now)
	case "capture":
		if item.Status != paymentmodel.PaymentAuthorizationStatusAuthorized {
			return dto.NewPaymentAuthorizationsResponse(item), invalidAction(action, string(item.Status))
		}
		amount := req.Amount
		if amount.IsZero() {
			amount = item.Amount
		}
		if amount.IsNegative() || amount.GreaterThan(item.Amount.Sub(item.CapturedAmount)) {
			return dto.PaymentAuthorizationsResponse{}, failure.BadRequestFromString("capture amount exceeds remaining authorization")
		}
		item.Status, item.CapturedAmount = paymentmodel.PaymentAuthorizationStatusCaptured, item.CapturedAmount.Add(amount)
	case "void":
		if item.Status != paymentmodel.PaymentAuthorizationStatusRequested && item.Status != paymentmodel.PaymentAuthorizationStatusAuthorized {
			return dto.NewPaymentAuthorizationsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status = paymentmodel.PaymentAuthorizationStatusVoided
	case "fail":
		if item.Status == paymentmodel.PaymentAuthorizationStatusCaptured || item.Status == paymentmodel.PaymentAuthorizationStatusVoided {
			return dto.NewPaymentAuthorizationsResponse(item), invalidAction(action, string(item.Status))
		}
		item.Status, item.FailureCode, item.FailureMessage = paymentmodel.PaymentAuthorizationStatusFailed, null.StringFrom(req.FailureCode), null.StringFrom(req.FailureMessage)
	default:
		return dto.NewPaymentAuthorizationsResponse(item), unsupportedAction(action)
	}

	item.SetSignatureMetaUpdate(req.ActorID)

	err = s.paymentRepo.UpdatePaymentAuthorizationsByID(ctx, item.ToPaymentAuthorizationsPrimaryID(), &item)
	if err != nil {
		return dto.PaymentAuthorizationsResponse{}, err
	}
	return dto.NewPaymentAuthorizationsResponse(item), nil
}

func operationCode(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(uuid.Must(uuid.NewV7()).String(), "-", "")[:20])
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
func mustUUID() uuid.UUID {
	id, _ := uuid.NewV7()
	return id
}

func normalizedJSON(value json.RawMessage) json.RawMessage {
	if len(value) == 0 {
		return json.RawMessage(`{}`)
	}
	return value
}
func validOptionalJSON(value json.RawMessage) bool { return len(value) == 0 || json.Valid(value) }

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

var supportedCurrencies = map[string]bool{
	"IDR": true, "USD": true, "SGD": true, "MYR": true, "THB": true,
	"PHP": true, "VND": true, "AUD": true, "EUR": true, "GBP": true,
	"JPY": true, "KRW": true, "CNY": true, "HKD": true, "TWD": true,
	"INR": true, "BND": true, "KHR": true, "LAK": true, "MMK": true,
}

var maxAmountByCurrency = map[string]decimal.Decimal{
	"IDR": decimal.NewFromInt(500_000_000),
	"USD": decimal.NewFromFloat(50_000),
	"SGD": decimal.NewFromFloat(65_000),
}

func validateCurrency(currency string) error {
	code := strings.ToUpper(strings.TrimSpace(currency))
	if code == "" {
		return failure.WithCode(shared.ErrPayUnsupportedCurrency, "currency is required")
	}
	if !supportedCurrencies[code] {
		return failure.WithCode(shared.ErrPayUnsupportedCurrency, "unsupported currency: "+code)
	}
	return nil
}

func validateAmount(amount decimal.Decimal, currency string) error {
	if amount.IsNegative() || amount.IsZero() {
		return failure.WithCode(shared.ErrPayInvalidAmount, "amount must be positive")
	}
	if max, ok := maxAmountByCurrency[strings.ToUpper(currency)]; ok {
		if amount.GreaterThan(max) {
			return failure.WithCode(shared.ErrPayInvalidAmount, "amount exceeds maximum for currency "+strings.ToUpper(currency))
		}
	}
	return nil
}

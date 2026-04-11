package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/guregu/null"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type webhookProviderStrategy interface {
	Provider() paymentmodel.Psp
	Verify(event WebhookEvent, cfg *paymentServiceSecrets) error
	NormalizeStatus(event WebhookEvent) paymentmodel.PaymentStatus
	BuildProcessedMessage(finalStatus paymentmodel.PaymentStatus) string
}

type midtransStatusReconciler interface {
	CheckTransactionStatus(ctx context.Context, reference string) (paymentmodel.ChargeResult, error)
}

type paymentServiceSecrets struct {
	midtransServerKey string
	xenditSecretKey   string
}

func (s *ServiceImpl) VerifyMidtransSignature(ctx context.Context, event WebhookEvent) error {
	strategy := midtransWebhookStrategy{}
	if err := strategy.Verify(event, s.secrets()); err != nil {
		log.Ctx(ctx).Warn().Err(err).Str("provider", event.Provider).Msg("[Webhook][Midtrans] signature verification failed")
		return err
	}
	return nil
}

func (s *ServiceImpl) VerifyXenditSignature(ctx context.Context, event WebhookEvent) error {
	strategy := xenditWebhookStrategy{}
	if err := strategy.Verify(event, s.secrets()); err != nil {
		log.Ctx(ctx).Warn().Err(err).Str("provider", event.Provider).Msg("[Webhook][Xendit] signature verification failed")
		return err
	}
	return nil
}

func (s *ServiceImpl) ProcessWebhookEvent(ctx context.Context, event WebhookEvent) (result WebhookResult, err error) {
	provider, err := parseWebhookProvider(event.Provider)
	if err != nil {
		return WebhookResult{}, err
	}

	processor, ok := s.webhookProcessor(provider)
	if !ok {
		return WebhookResult{}, failure.NotFound("webhook processor")
	}

	txRepo, err := s.paymentRepo.BeginTx(ctx)
	if err != nil {
		return WebhookResult{}, err
	}

	defer func() {
		if err != nil {
			if rollbackErr := txRepo.Rollback(ctx); rollbackErr != nil {
				log.Ctx(ctx).Error().Err(rollbackErr).Msg("[Webhook] rollback failed")
			}
			return
		}
		if commitErr := txRepo.Commit(ctx); commitErr != nil {
			err = commitErr
		}
	}()

	paymentRecord, intentRecord, err := s.processWebhookTransaction(ctx, txRepo, processor, event)
	if err != nil {
		return WebhookResult{}, err
	}

	processedAt := time.Now().UTC()
	result = WebhookResult{
		Provider:      string(provider),
		Success:       true,
		EventID:       webhookEventID(event),
		TransactionID: event.TransactionID,
		PaymentID:     paymentRecord.Id.String(),
		IntentID:      intentRecord.Id.String(),
		FinalStatus:   string(paymentRecord.Status),
		Message:       processor.BuildProcessedMessage(paymentRecord.Status),
		ProcessedAt:   processedAt,
	}

	log.Ctx(ctx).Info().
		Str("provider", string(provider)).
		Str("transactionId", event.TransactionID).
		Str("paymentId", paymentRecord.Id.String()).
		Str("status", string(paymentRecord.Status)).
		Msg("[Webhook] processed successfully")

	return result, nil
}

func (s *ServiceImpl) processWebhookTransaction(ctx context.Context, repo repository.Repository, processor webhookProviderStrategy, event WebhookEvent) (*paymentmodel.Payments, *paymentmodel.PaymentIntents, error) {
	if event.TransactionID == "" {
		return nil, nil, failure.BadRequestFromString("transactionId is required")
	}

	if err := processor.Verify(event, s.secrets()); err != nil {
		return nil, nil, err
	}

	paymentRecord, err := resolvePaymentByPSPTransactionID(ctx, repo, processor.Provider(), event.TransactionID)
	if err != nil {
		return nil, nil, err
	}

	intentRecord, err := repo.ResolvePaymentIntentsByID(ctx, paymentmodel.PaymentIntentsPrimaryID{Id: paymentRecord.IntentId})
	if err != nil {
		return nil, nil, err
	}

	normalizedStatus := processor.NormalizeStatus(event)
	normalizedStatus = s.reconcileWebhookStatus(ctx, processor.Provider(), event, normalizedStatus)
	finalPaymentStatus, err := s.transitionWebhookStatus(paymentRecord.Status, normalizedStatus)
	if err != nil {
		return nil, nil, err
	}

	processedAt := time.Now().UTC()
	rawEventPayload, _ := json.Marshal(event)
	paymentRecord.Status = finalPaymentStatus
	paymentRecord.PspTransactionId = null.StringFrom(event.TransactionID)
	paymentRecord.PspReference = null.StringFrom(firstNonEmpty(event.ExternalID, paymentRecord.PspReference.String, event.TransactionID))
	paymentRecord.PspRawResponse = rawEventPayload
	paymentRecord.MetaUpdatedAt = processedAt
	paymentRecord.MetaUpdatedBy = nuuid.From(intentRecord.MetaCreatedBy)

	applyPaymentWebhookSideEffects(paymentRecord, event, normalizedStatus, processedAt)
	applyIntentWebhookSideEffects(&intentRecord, paymentRecord, normalizedStatus, processedAt, event)

	if err := repo.UpdatePaymentsByID(ctx, paymentRecord.ToPaymentsPrimaryID(), paymentRecord); err != nil {
		return nil, nil, err
	}
	if err := repo.UpdatePaymentIntentsByID(ctx, intentRecord.ToPaymentIntentsPrimaryID(), &intentRecord); err != nil {
		return nil, nil, err
	}

	return paymentRecord, &intentRecord, nil
}

func (s *ServiceImpl) reconcileWebhookStatus(ctx context.Context, provider paymentmodel.Psp, event WebhookEvent, current paymentmodel.PaymentStatus) paymentmodel.PaymentStatus {
	if provider != paymentmodel.PspMidtrans {
		return current
	}
	if !shouldReconcileMidtransStatus(current) {
		return current
	}

	gateway, ok := s.gatewayRegistry.LookupByPSP(paymentmodel.PspMidtrans)
	if !ok || gateway == nil {
		return current
	}
	reconciler, ok := gateway.(midtransStatusReconciler)
	if !ok {
		return current
	}

	for _, reference := range []string{strings.TrimSpace(event.ExternalID), strings.TrimSpace(event.TransactionID)} {
		if reference == "" {
			continue
		}
		statusResult, err := reconciler.CheckTransactionStatus(ctx, reference)
		if err != nil {
			log.Ctx(ctx).Warn().Err(err).Str("provider", string(provider)).Str("reference", reference).Msg("[Webhook] midtrans reconciliation call failed")
			continue
		}
		if statusResult.NextStatus == "" {
			continue
		}
		if statusResult.NextStatus != current {
			log.Ctx(ctx).Info().
				Str("provider", string(provider)).
				Str("reference", reference).
				Str("webhookStatus", string(current)).
				Str("reconciledStatus", string(statusResult.NextStatus)).
				Msg("[Webhook] midtrans status reconciled from provider API")
		}
		return statusResult.NextStatus
	}

	return current
}

func shouldReconcileMidtransStatus(status paymentmodel.PaymentStatus) bool {
	switch status {
	case paymentmodel.PaymentStatusPending, paymentmodel.PaymentStatusAuthorised:
		return true
	default:
		return false
	}
}

func (s *ServiceImpl) webhookProcessor(provider paymentmodel.Psp) (webhookProviderStrategy, bool) {
	switch provider {
	case paymentmodel.PspMidtrans:
		return midtransWebhookStrategy{}, true
	case paymentmodel.PspXendit:
		return xenditWebhookStrategy{}, true
	default:
		return nil, false
	}
}

func (s *ServiceImpl) secrets() *paymentServiceSecrets {
	if s.cfg == nil {
		return &paymentServiceSecrets{}
	}
	return &paymentServiceSecrets{
		midtransServerKey: s.cfg.Externals.Providers.Midtrans.ServerKey,
		xenditSecretKey:   s.cfg.Externals.Providers.Xendit.SecretKey,
	}
}

func (s *ServiceImpl) transitionWebhookStatus(current, target paymentmodel.PaymentStatus) (paymentmodel.PaymentStatus, error) {
	if current == target {
		return target, nil
	}
	if s.stateMachine != nil && s.stateMachine.CanTransition(current, target) {
		return target, nil
	}
	if isWebhookTerminalStatus(current) {
		return current, nil
	}
	return current, fmt.Errorf("invalid payment transition from %s to %s", current, target)
}

func resolvePaymentByPSPTransactionID(ctx context.Context, repo repository.Repository, provider paymentmodel.Psp, transactionID string) (*paymentmodel.Payments, error) {
	filter := paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{
			{Field: string(paymentmodel.PaymentsDBFieldName.Psp), Operator: paymentmodel.OperatorEqual, Value: provider},
			{Field: string(paymentmodel.PaymentsDBFieldName.PspTransactionId), Operator: paymentmodel.OperatorEqual, Value: transactionID},
		},
		Sorts:      []paymentmodel.Sort{{Field: string(paymentmodel.PaymentsDBFieldName.MetaCreatedAt), Order: paymentmodel.SortDesc}},
		Pagination: paymentmodel.Pagination{Page: 1, PageSize: 1},
	}
	result, err := repo.ResolvePaymentsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, failure.NotFound("payment")
	}
	payment := result[0].Payments
	return &payment, nil
}

func applyPaymentWebhookSideEffects(payment *paymentmodel.Payments, event WebhookEvent, normalizedStatus paymentmodel.PaymentStatus, processedAt time.Time) {
	switch normalizedStatus {
	case paymentmodel.PaymentStatusCaptured:
		payment.AuthorisedAt = null.TimeFrom(processedAt)
		payment.CapturedAt = null.TimeFrom(processedAt)
		payment.AuthorisedAmount = decimal.NullDecimal{Decimal: payment.Amount, Valid: true}
		payment.CapturedAmount = decimal.NullDecimal{Decimal: payment.Amount, Valid: true}
		payment.FailureCode = null.String{}
		payment.FailureMessage = null.String{}
		payment.FailureCategory = null.String{}
	case paymentmodel.PaymentStatusAuthorised:
		payment.AuthorisedAt = null.TimeFrom(processedAt)
		payment.AuthorisedAmount = decimal.NullDecimal{Decimal: payment.Amount, Valid: true}
	case paymentmodel.PaymentStatusFailed, paymentmodel.PaymentStatusCancelled, paymentmodel.PaymentStatusExpired:
		payment.FailureCode = null.StringFrom(event.Status)
		payment.FailureMessage = null.StringFrom(event.EventType)
		payment.FailureCategory = null.StringFrom(strings.ToUpper(string(normalizedStatus)))
	case paymentmodel.PaymentStatusRefunding, paymentmodel.PaymentStatusRefunded, paymentmodel.PaymentStatusPartiallyRefunded:
		payment.FailureCode = null.String{}
		payment.FailureMessage = null.String{}
		payment.FailureCategory = null.String{}
	}
}

func applyIntentWebhookSideEffects(intent *paymentmodel.PaymentIntents, payment *paymentmodel.Payments, normalizedStatus paymentmodel.PaymentStatus, processedAt time.Time, event WebhookEvent) {
	intent.Status = normalizeIntentStatus(normalizedStatus)
	intent.MetaUpdatedAt = processedAt
	intent.MetaUpdatedBy = payment.MetaUpdatedBy

	switch normalizedStatus {
	case paymentmodel.PaymentStatusCaptured:
		if !intent.ConfirmedAt.Valid {
			intent.ConfirmedAt = null.TimeFrom(processedAt)
		}
		intent.CancellationReason = null.String{}
		intent.CancelledAt = null.Time{}
	case paymentmodel.PaymentStatusAuthorised, paymentmodel.PaymentStatusPending:
		if !intent.ConfirmedAt.Valid && event.Status != "" {
			intent.ConfirmedAt = null.TimeFrom(processedAt)
		}
	case paymentmodel.PaymentStatusCancelled, paymentmodel.PaymentStatusExpired, paymentmodel.PaymentStatusFailed:
		intent.CancelledAt = null.TimeFrom(processedAt)
		intent.CancellationReason = null.StringFrom(firstNonEmpty(event.EventType, event.Status))
	}
}

func normalizeIntentStatus(status paymentmodel.PaymentStatus) paymentmodel.PaymentStatus {
	switch status {
	case paymentmodel.PaymentStatusCaptured:
		return paymentmodel.PaymentStatusCompleted
	case paymentmodel.PaymentStatusAuthorised, paymentmodel.PaymentStatusPending:
		return paymentmodel.PaymentStatusPending
	case paymentmodel.PaymentStatusFailed:
		return paymentmodel.PaymentStatusFailed
	case paymentmodel.PaymentStatusCancelled:
		return paymentmodel.PaymentStatusCancelled
	case paymentmodel.PaymentStatusExpired:
		return paymentmodel.PaymentStatusExpired
	case paymentmodel.PaymentStatusRefunding:
		return paymentmodel.PaymentStatusRefunding
	case paymentmodel.PaymentStatusRefunded:
		return paymentmodel.PaymentStatusRefunded
	case paymentmodel.PaymentStatusPartiallyRefunded:
		return paymentmodel.PaymentStatusPartiallyRefunded
	case paymentmodel.PaymentStatusDisputed:
		return paymentmodel.PaymentStatusDisputed
	case paymentmodel.PaymentStatusChargebackWon:
		return paymentmodel.PaymentStatusChargebackWon
	case paymentmodel.PaymentStatusChargebackLost:
		return paymentmodel.PaymentStatusChargebackLost
	default:
		return paymentmodel.PaymentStatusPending
	}
}

func isWebhookTerminalStatus(status paymentmodel.PaymentStatus) bool {
	switch status {
	case paymentmodel.PaymentStatusFailed,
		paymentmodel.PaymentStatusCancelled,
		paymentmodel.PaymentStatusExpired,
		paymentmodel.PaymentStatusRefunded,
		paymentmodel.PaymentStatusPartiallyRefunded,
		paymentmodel.PaymentStatusDisputed,
		paymentmodel.PaymentStatusChargebackWon,
		paymentmodel.PaymentStatusChargebackLost:
		return true
	default:
		return false
	}
}

func webhookEventID(event WebhookEvent) string {
	return firstNonEmpty(event.ExternalID, event.TransactionID, event.EventType)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func parseWebhookProvider(value string) (paymentmodel.Psp, error) {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case string(paymentmodel.PspMidtrans):
		return paymentmodel.PspMidtrans, nil
	case string(paymentmodel.PspXendit):
		return paymentmodel.PspXendit, nil
	default:
		return "", failure.BadRequestFromString("provider must be MIDTRANS or XENDIT")
	}
}

type midtransWebhookStrategy struct{}

func (midtransWebhookStrategy) Provider() paymentmodel.Psp {
	return paymentmodel.PspMidtrans
}

func (midtransWebhookStrategy) Verify(event WebhookEvent, cfg *paymentServiceSecrets) error {
	if cfg == nil || strings.TrimSpace(cfg.midtransServerKey) == "" {
		return failure.InternalError(fmt.Errorf("midtrans server key is not configured"))
	}
	if !strings.EqualFold(strings.TrimSpace(event.SignatureMethod), "SHA512") && strings.TrimSpace(event.SignatureMethod) != "" {
		return failure.Unauthorized("midtrans webhook signature method must be SHA512")
	}

	// Midtrans signature formula: SHA512(order_id + status_code + gross_amount + ServerKey)
	orderId := event.ExternalID // order_id normalized from raw payload
	if orderId == "" {
		// Fallback: try to extract from raw payload
		orderId = rawPayloadString(event.RawPayload, "order_id")
	}
	if orderId == "" {
		return failure.BadRequestFromString("order_id is required for midtrans signature verification")
	}

	// Extract status_code from raw payload (NOT transaction_status)
	statusCode := rawPayloadString(event.RawPayload, "status_code")
	if statusCode == "" {
		return failure.BadRequestFromString("status_code is required for midtrans signature verification")
	}

	// Extract gross_amount
	grossAmount := rawPayloadString(event.RawPayload, "gross_amount")
	if grossAmount == "" {
		grossAmount = rawPayloadString(event.RawPayload, "amount")
	}
	if grossAmount == "" {
		return failure.BadRequestFromString("gross_amount is required for midtrans signature verification")
	}

	// Construct the signing string: order_id + status_code + gross_amount + ServerKey
	signingString := orderId + statusCode + grossAmount + cfg.midtransServerKey
	expected := sha512.Sum512([]byte(signingString))

	if !constantTimeEquals(normalizeSignature(event.Signature), hex.EncodeToString(expected[:])) {
		return failure.Unauthorized("midtrans webhook signature verification failed")
	}
	return nil
}

func (midtransWebhookStrategy) NormalizeStatus(event WebhookEvent) paymentmodel.PaymentStatus {
	switch strings.ToLower(strings.TrimSpace(event.Status)) {
	case "capture", "settlement", "completed", "success", "succeeded":
		return paymentmodel.PaymentStatusCaptured
	case "authorize", "authorized":
		return paymentmodel.PaymentStatusAuthorised
	case "pending":
		return paymentmodel.PaymentStatusPending
	case "cancel", "cancelled", "canceled":
		return paymentmodel.PaymentStatusCancelled
	case "expire", "expired":
		return paymentmodel.PaymentStatusExpired
	case "refund", "refunded":
		return paymentmodel.PaymentStatusRefunding
	case "partial_refund", "partially_refunded":
		return paymentmodel.PaymentStatusPartiallyRefunded
	case "dispute":
		return paymentmodel.PaymentStatusDisputed
	case "chargeback_won":
		return paymentmodel.PaymentStatusChargebackWon
	case "chargeback_lost":
		return paymentmodel.PaymentStatusChargebackLost
	default:
		return paymentmodel.PaymentStatusPending
	}
}

func (midtransWebhookStrategy) BuildProcessedMessage(finalStatus paymentmodel.PaymentStatus) string {
	return fmt.Sprintf("midtrans webhook processed with final status %s", finalStatus)
}

type xenditWebhookStrategy struct{}

func (xenditWebhookStrategy) Provider() paymentmodel.Psp {
	return paymentmodel.PspXendit
}

func (xenditWebhookStrategy) Verify(event WebhookEvent, cfg *paymentServiceSecrets) error {
	if cfg == nil || strings.TrimSpace(cfg.xenditSecretKey) == "" {
		return failure.InternalError(fmt.Errorf("xendit secret key is not configured"))
	}
	if !strings.EqualFold(strings.TrimSpace(event.SignatureMethod), "HMACSHA256") && !strings.EqualFold(strings.TrimSpace(event.SignatureMethod), "SHA256") && strings.TrimSpace(event.SignatureMethod) != "" {
		return failure.Unauthorized("xendit webhook signature method must be HMACSHA256 or SHA256")
	}
	signingString := strings.Join([]string{event.TransactionID, event.EventType, event.Status, event.ExternalID}, "|")
	mac := hmac.New(sha256.New, []byte(cfg.xenditSecretKey))
	_, _ = mac.Write([]byte(signingString))
	expected := hex.EncodeToString(mac.Sum(nil))
	if !constantTimeEquals(normalizeSignature(event.Signature), expected) {
		return failure.Unauthorized("xendit webhook signature verification failed")
	}
	return nil
}

func (xenditWebhookStrategy) NormalizeStatus(event WebhookEvent) paymentmodel.PaymentStatus {
	switch strings.ToUpper(strings.TrimSpace(event.Status)) {
	case "CAPTURE", "CAPTURED", "SETTLE", "SETTLED", "PAID", "SUCCESS", "SUCCEEDED":
		return paymentmodel.PaymentStatusCaptured
	case "AUTHORIZED", "AUTHORIZE":
		return paymentmodel.PaymentStatusAuthorised
	case "PENDING":
		return paymentmodel.PaymentStatusPending
	case "FAILED", "DECLINED", "DENY":
		return paymentmodel.PaymentStatusFailed
	case "CANCELLED", "CANCELED":
		return paymentmodel.PaymentStatusCancelled
	case "EXPIRED":
		return paymentmodel.PaymentStatusExpired
	case "REFUND", "REFUNDED":
		return paymentmodel.PaymentStatusRefunding
	case "PARTIALLY_REFUNDED", "PARTIAL_REFUND":
		return paymentmodel.PaymentStatusPartiallyRefunded
	case "DISPUTED":
		return paymentmodel.PaymentStatusDisputed
	case "CHARGEBACK_WON":
		return paymentmodel.PaymentStatusChargebackWon
	case "CHARGEBACK_LOST":
		return paymentmodel.PaymentStatusChargebackLost
	default:
		return paymentmodel.PaymentStatusPending
	}
}

func (xenditWebhookStrategy) BuildProcessedMessage(finalStatus paymentmodel.PaymentStatus) string {
	return fmt.Sprintf("xendit webhook processed with final status %s", finalStatus)
}

func rawPayloadString(raw map[string]interface{}, key string) string {
	if raw == nil {
		return ""
	}
	value, ok := raw[key]
	if !ok || value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	case fmt.Stringer:
		return typed.String()
	case float64:
		return fmt.Sprintf("%v", typed)
	default:
		encoded, err := json.Marshal(typed)
		if err != nil {
			return ""
		}
		return string(encoded)
	}
}

func normalizeSignature(signature string) string {
	trimmed := strings.TrimSpace(signature)
	trimmed = strings.TrimPrefix(trimmed, "sha512=")
	trimmed = strings.TrimPrefix(trimmed, "SHA512=")
	trimmed = strings.TrimPrefix(trimmed, "sha256=")
	trimmed = strings.TrimPrefix(trimmed, "SHA256=")
	return strings.ToLower(trimmed)
}

func constantTimeEquals(left, right string) bool {
	if len(left) != len(right) {
		return false
	}
	return hmac.Equal([]byte(left), []byte(right))
}

func normalizeWebhookResponseError(message string) *WebhookError {
	return &WebhookError{
		Code:    strings.ToUpper(strings.ReplaceAll(message, " ", "_")),
		Message: message,
	}
}

func (s *ServiceImpl) webhookErrorResponse(provider paymentmodel.Psp, event WebhookEvent, message string) WebhookResult {
	return WebhookResult{
		Provider:      string(provider),
		Success:       false,
		EventID:       webhookEventID(event),
		TransactionID: event.TransactionID,
		Message:       message,
		ProcessedAt:   time.Now().UTC(),
		Error:         normalizeWebhookResponseError(message),
	}
}

func (s *ServiceImpl) webhookSignatureErrorResponse(provider paymentmodel.Psp, event WebhookEvent, message string) WebhookResult {
	resp := s.webhookErrorResponse(provider, event, message)
	resp.Error = &WebhookError{Code: "SIGNATURE_VERIFICATION_FAILED", Message: message}
	return resp
}

func (s *ServiceImpl) processWebhookValidationError(provider paymentmodel.Psp, event WebhookEvent, err error) (WebhookResult, error) {
	if err == nil {
		return WebhookResult{}, nil
	}
	return s.webhookErrorResponse(provider, event, err.Error()), err
}

func webhookSignatureMethodOrDefault(provider paymentmodel.Psp, method string) string {
	if strings.TrimSpace(method) != "" {
		return strings.ToUpper(strings.TrimSpace(method))
	}
	switch provider {
	case paymentmodel.PspMidtrans:
		return "SHA512"
	case paymentmodel.PspXendit:
		return "HMACSHA256"
	default:
		return ""
	}
}

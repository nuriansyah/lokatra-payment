package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/netip"
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type PaymentService interface {
	ResolvePaymentFlow(ctx context.Context, request PaymentFlowRequest) (PaymentFlowQuote, error)
	ExecutePayment(ctx context.Context, request PaymentFlowRequest) (PaymentExecutionResult, error)
	VerifyMidtransSignature(ctx context.Context, event WebhookEvent) error
	VerifyXenditSignature(ctx context.Context, event WebhookEvent) error
	ProcessWebhookEvent(ctx context.Context, event WebhookEvent) (WebhookResult, error)
}

func (s *ServiceImpl) ResolvePaymentFlow(ctx context.Context, request PaymentFlowRequest) (PaymentFlowQuote, error) {
	validatedRequest, err := s.validateRequest(request)
	if err != nil {
		return PaymentFlowQuote{}, err
	}

	decision, err := s.routingEngine.Resolve(ctx, validatedRequest)
	if err != nil {
		return PaymentFlowQuote{}, err
	}

	builder := NewPaymentFlowBuilder(validatedRequest).
		WithDecision(decision).
		WithConfig(s.config)

	quote, err := builder.BuildQuote()
	if err != nil {
		return PaymentFlowQuote{}, err
	}
	return quote, nil
}

func (s *ServiceImpl) ExecutePayment(ctx context.Context, request PaymentFlowRequest) (PaymentExecutionResult, error) {
	validatedRequest, err := s.validateRequest(request)
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	idempotencyKey := strings.TrimSpace(validatedRequest.IdempotencyKey)
	if idempotencyKey == "" {
		return s.executePaymentCore(ctx, validatedRequest)
	}
	if s.idempotencyRepo == nil || s.mutex == nil {
		return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf("idempotency dependencies are not initialized"))
	}

	requestHash, err := hashExecuteRequest(validatedRequest)
	if err != nil {
		return PaymentExecutionResult{}, failure.InternalError(err)
	}

	lockName := fmt.Sprintf("payment:execute:%s:%s", validatedRequest.MerchantID.String(), idempotencyKey)
	mtx := s.mutex.NewMutex(
		lockName,
		redsync.WithExpiry(30*time.Second),
		redsync.WithTries(1),
	)
	if err := mtx.Lock(); err != nil {
		return PaymentExecutionResult{}, failure.Conflict("execute", "payment_flow", "request is already being processed")
	}
	defer func() {
		_, _ = mtx.Unlock()
	}()

	record, claimed, err := s.idempotencyRepo.ClaimExecuteIdempotency(
		ctx,
		idempotencyKey,
		validatedRequest.MerchantID,
		"/v1/payments/flows/execute",
		requestHash,
		validatedRequest.ActorID,
		time.Now().UTC().Add(2*time.Minute),
	)
	if err != nil {
		return PaymentExecutionResult{}, err
	}
	if !claimed {
		if record.CompletedAt.Valid && record.ResponseStatus.Valid && len(record.ResponseBody) > 0 {
			if int(record.ResponseStatus.Int64) == http.StatusCreated {
				var stored struct {
					Data       PaymentExecutionResult `json:"data"`
					StatusCode int                    `json:"statusCode"`
				}
				if err := json.Unmarshal(record.ResponseBody, &stored); err == nil {
					return stored.Data, nil
				}
			}

			var storedErr struct {
				Error      string `json:"error"`
				StatusCode int    `json:"statusCode"`
			}
			if err := json.Unmarshal(record.ResponseBody, &storedErr); err == nil {
				msg := strings.TrimSpace(storedErr.Error)
				if msg == "" {
					msg = "idempotent request previously failed"
				}
				if storedErr.StatusCode == 0 {
					storedErr.StatusCode = int(record.ResponseStatus.Int64)
				}
				return PaymentExecutionResult{}, failure.New(storedErr.StatusCode, errors.New(msg))
			}
		}
		return PaymentExecutionResult{}, failure.Conflict("execute", "payment_flow", "idempotency key already processed")
	}

	result, err := s.executePaymentCore(ctx, validatedRequest)
	if err != nil {
		code := failure.GetCode(err)
		errPayload, _ := json.Marshal(map[string]any{
			"error":      err.Error(),
			"errorCode":  failure.GetErrorCode(err),
			"statusCode": code,
		})
		_, _ = s.idempotencyRepo.CompleteExecuteIdempotency(ctx, record.Id, code, errPayload, validatedRequest.ActorID)
		return PaymentExecutionResult{}, err
	}

	payload, _ := json.Marshal(map[string]any{
		"data":       result,
		"statusCode": http.StatusCreated,
	})
	_, _ = s.idempotencyRepo.CompleteExecuteIdempotency(ctx, record.Id, http.StatusCreated, payload, validatedRequest.ActorID)

	return result, nil
}

func (s *ServiceImpl) executePaymentCore(ctx context.Context, request PaymentFlowRequest) (PaymentExecutionResult, error) {
	validatedRequest, err := s.validateRequest(request)
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	decision, err := s.routingEngine.Resolve(ctx, validatedRequest)
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	builder := NewPaymentFlowBuilder(validatedRequest).
		WithDecision(decision).
		WithConfig(s.config)

	quote, err := builder.BuildQuote()
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	intent, paymentAttempt, err := builder.BuildDomainRecords()
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	repo := s.paymentRepo
	if repo == nil {
		return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}

	if s.paymentRepo != nil {
		err = repo.CreatePaymentIntents(ctx, intent)
		if err != nil {
			if failure.GetCode(err) == http.StatusInternalServerError {
				log.Ctx(ctx).Error().Err(err).Msg("[ExecutePayment][CreatePaymentIntents] failed to create payment intents")
				return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
			}
			log.Ctx(ctx).Warn().Err(err).Msg("[ExecutePayment][CreatePaymentIntents] failed to create payment intents")
			return PaymentExecutionResult{}, err
		}
		err = repo.CreatePayments(ctx, paymentAttempt)
		if err != nil {
			if failure.GetCode(err) == http.StatusInternalServerError {
				log.Ctx(ctx).Error().Err(err).Msg("[ExecutePayment][CreatePayments] failed to create payments")
				return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
			}
			log.Ctx(ctx).Warn().Err(err).Msg("[ExecutePayment][CreatePayments] failed to create payments")
			return PaymentExecutionResult{}, err
		}
	}

	gateway, ok := s.gatewayRegistry.LookupByAccountID(decision.PSPAccountID)
	if !ok {
		gateway, ok = s.gatewayRegistry.LookupByPSP(decision.PSP)
	}
	if !ok {
		return PaymentExecutionResult{}, failure.NotFound("payment gateway")
	}

	chargeResult, err := gateway.Charge(ctx, paymentmodel.ChargeRequest{
		MerchantID:          intent.MerchantId,
		IntentID:            intent.Id,
		PaymentID:           paymentAttempt.Id,
		PaymentCode:         paymentAttempt.PaymentCode,
		IntentCode:          intent.IntentCode,
		PSP:                 decision.PSP,
		Amount:              paymentAttempt.Amount,
		Currency:            paymentAttempt.Currency,
		PaymentMethodType:   paymentAttempt.PaymentMethodType,
		CustomerName:        request.CustomerName,
		CustomerEmail:       request.CustomerEmail,
		CustomerPhone:       request.CustomerPhone,
		CustomerCountry:     request.CustomerCountry,
		Description:         request.Description,
		StatementDescriptor: request.StatementDescriptor,
		Metadata:            paymentAttempt.Metadata,
		UseCase:             string(request.UseCase),
		Requires3DS:         request.Requires3DS,
		CaptureMode:         string(request.CaptureMode),
	})
	if err != nil {
		return PaymentExecutionResult{}, err
	}

	finalStatus, err := s.stateMachine.Transition(paymentmodel.PaymentStatusPending, s.stateMachine.NormalizeGatewayStatus(chargeResult.NextStatus, chargeResult.RequiresAction))
	if err != nil {
		return PaymentExecutionResult{}, err
	}
	paymentAttempt.Status = finalStatus
	paymentAttempt.PspTransactionId = null.StringFrom(chargeResult.PSPTransactionID)
	paymentAttempt.PspReference = null.StringFrom(chargeResult.PSPReference)
	paymentAttempt.PspRawRequest = chargeResult.RawRequest
	paymentAttempt.PspRawResponse = chargeResult.RawResponse
	paymentAttempt.AuthorisedAmount = decimal.NullDecimal{Decimal: chargeResult.AuthorisedAmount, Valid: true}
	paymentAttempt.CapturedAmount = decimal.NullDecimal{Decimal: chargeResult.CapturedAmount, Valid: true}
	paymentAttempt.ProcessingFee = decimal.NullDecimal{Decimal: chargeResult.ProcessingFee, Valid: true}
	paymentAttempt.MetaUpdatedAt = time.Now().UTC()
	paymentAttempt.MetaUpdatedBy = nuuid.From(request.ActorID)

	if s.paymentRepo != nil {
		err = repo.UpdatePaymentsByID(ctx, paymentAttempt.ToPaymentsPrimaryID(), paymentAttempt)
		if err != nil {
			if failure.GetCode(err) == http.StatusInternalServerError {
				log.Ctx(ctx).Error().Err(err).Msg("[ExecutePayment][UpdatePaymentsByID] failed to update payments")
				return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
			}
			log.Ctx(ctx).Warn().Err(err).Msg("[ExecutePayment][UpdatePaymentsByID] failed to update payments")
			return PaymentExecutionResult{}, err
		}
		intent.Status = paymentmodel.PaymentStatusPending
		if finalStatus == paymentmodel.PaymentStatusCaptured || finalStatus == paymentmodel.PaymentStatusCompleted {
			intent.Status = paymentmodel.PaymentStatusCompleted
		}
		err = repo.UpdatePaymentIntentsByID(ctx, intent.ToPaymentIntentsPrimaryID(), intent)
		if err != nil {
			if failure.GetCode(err) == http.StatusInternalServerError {
				log.Ctx(ctx).Error().Err(err).Msg("[ExecutePayment][UpdatePaymentIntentsByID] failed to update payment intents")
				return PaymentExecutionResult{}, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
			}
			log.Ctx(ctx).Warn().Err(err).Msg("[ExecutePayment][UpdatePaymentIntentsByID] failed to update payment intents")
			return PaymentExecutionResult{}, err
		}
	}

	statePath := []paymentmodel.PaymentStatus{paymentmodel.PaymentStatusPending, finalStatus}
	return PaymentExecutionResult{
		Quote:       quote,
		Charge:      chargeResult,
		FinalStatus: finalStatus,
		StatePath:   statePath,
	}, nil
}

func hashExecuteRequest(request PaymentFlowRequest) (string, error) {
	payload := map[string]any{
		"actorId":             request.ActorID.String(),
		"merchantId":          request.MerchantID.String(),
		"useCase":             request.UseCase,
		"orderId":             request.OrderID,
		"orderType":           request.OrderType,
		"productType":         request.ProductType,
		"amount":              request.Amount.String(),
		"currency":            request.Currency,
		"taxAmount":           request.TaxAmount.String(),
		"discountAmount":      request.DiscountAmount.String(),
		"tipAmount":           request.TipAmount.String(),
		"userId":              request.UserID,
		"customerName":        request.CustomerName,
		"customerEmail":       request.CustomerEmail,
		"customerPhone":       request.CustomerPhone,
		"customerCountry":     request.CustomerCountry,
		"customerIp":          request.CustomerIP,
		"paymentMethodId":     request.PaymentMethodID,
		"paymentMethodType":   request.PaymentMethodType,
		"requires3ds":         request.Requires3DS,
		"statementDescriptor": request.StatementDescriptor,
		"description":         request.Description,
		"captureMode":         request.CaptureMode,
		"routingProfileId":    request.RoutingProfileID,
		"idempotencyKey":      request.IdempotencyKey,
		"correlationId":       request.CorrelationID,
		"metadata":            request.Metadata,
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(encoded)
	return fmt.Sprintf("%x", sum[:]), nil
}

func (s *ServiceImpl) validateRequest(request PaymentFlowRequest) (PaymentFlowRequest, error) {
	if request.ActorID == uuid.Nil {
		return PaymentFlowRequest{}, failure.BadRequestFromString("actorId is required")
	}
	if request.MerchantID == uuid.Nil {
		return PaymentFlowRequest{}, failure.BadRequestFromString("merchantId is required")
	}
	if request.Amount.LessThanOrEqual(decimal.Zero) {
		return PaymentFlowRequest{}, failure.BadRequestFromString("amount must be greater than zero")
	}
	if request.Currency == "" {
		return PaymentFlowRequest{}, failure.BadRequestFromString("currency is required")
	}
	if request.PaymentMethodType == "" {
		return PaymentFlowRequest{}, failure.BadRequestFromString("paymentMethodType is required")
	}
	if request.UseCase == "" {
		request.UseCase = s.config.DefaultUseCase
	}
	if request.CaptureMode == "" {
		request.CaptureMode = PaymentCaptureModeAutomatic
	}
	if request.Metadata == nil {
		request.Metadata = map[string]any{}
	}
	return request, nil
}

type PaymentFlowBuilder struct {
	request  PaymentFlowRequest
	config   PaymentServiceConfig
	decision RoutingDecision
	quote    *PaymentFlowQuote
	intent   *paymentmodel.PaymentIntents
	payment  *paymentmodel.Payments
}

func NewPaymentFlowBuilder(request PaymentFlowRequest) *PaymentFlowBuilder {
	return &PaymentFlowBuilder{request: request}
}

func (b *PaymentFlowBuilder) WithDecision(decision RoutingDecision) *PaymentFlowBuilder {
	b.decision = decision
	return b
}

func (b *PaymentFlowBuilder) WithConfig(config PaymentServiceConfig) *PaymentFlowBuilder {
	b.config = config
	return b
}

func (b *PaymentFlowBuilder) BuildQuote() (PaymentFlowQuote, error) {
	intent, paymentAttempt, err := b.BuildDomainRecords()
	if err != nil {
		return PaymentFlowQuote{}, err
	}
	steps := []FlowStep{
		{Name: "resolve-routing", Description: "Evaluate declarative policies and database fallback to select the PSP account."},
		{Name: "build-intent", Description: "Materialize the payment intent and provider attempt records."},
		{Name: "dispatch-charge", Description: "Send the charge request to the provider adapter through the bridge interface."},
		{Name: "apply-state", Description: "Advance the payment lifecycle using the state machine."},
	}
	return PaymentFlowQuote{Intent: intent, Payment: paymentAttempt, Decision: b.decision, Steps: steps, CreatedAt: time.Now().UTC()}, nil
}

func (b *PaymentFlowBuilder) BuildDomainRecords() (*paymentmodel.PaymentIntents, *paymentmodel.Payments, error) {
	if b.intent != nil && b.payment != nil {
		return b.intent, b.payment, nil
	}
	intent, err := b.buildIntent()
	if err != nil {
		return nil, nil, err
	}
	paymentAttempt, err := b.buildPaymentAttempt(intent)
	if err != nil {
		return nil, nil, err
	}
	b.intent = intent
	b.payment = paymentAttempt
	return intent, paymentAttempt, nil
}

func (b *PaymentFlowBuilder) buildIntent() (*paymentmodel.PaymentIntents, error) {
	intentID, err := uuid.NewV4()
	if err != nil {
		return nil, failure.InternalError(err)
	}
	intentCode := businessCode("PAY")
	metadataJSON, err := json.Marshal(b.request.Metadata)
	if err != nil {
		return nil, failure.InternalError(err)
	}

	var orderID nuuid.NUUID
	if b.request.OrderID != "" {
		orderID = nuuid.FromString(b.request.OrderID)
	}
	var userID nuuid.NUUID
	if b.request.UserID != "" {
		userID = nuuid.FromString(b.request.UserID)
	}
	var paymentMethodID nuuid.NUUID
	if b.request.PaymentMethodID != "" {
		paymentMethodID = nuuid.FromString(b.request.PaymentMethodID)
	}
	var routingProfileID nuuid.NUUID
	if b.request.RoutingProfileID != "" {
		routingProfileID = nuuid.FromString(b.request.RoutingProfileID)
	}
	var correlationID nuuid.NUUID
	if b.request.CorrelationID != "" {
		correlationID = nuuid.FromString(b.request.CorrelationID)
	}

	intent := &paymentmodel.PaymentIntents{
		Id:                  intentID,
		IntentCode:          intentCode,
		MerchantId:          b.request.MerchantID,
		OrderId:             orderID,
		OrderType:           null.StringFrom(b.request.OrderType),
		Amount:              b.request.Amount,
		Currency:            b.request.Currency,
		TaxAmount:           b.request.TaxAmount,
		DiscountAmount:      b.request.DiscountAmount,
		TipAmount:           b.request.TipAmount,
		UserId:              userID,
		CustomerName:        null.StringFrom(b.request.CustomerName),
		CustomerEmail:       null.StringFrom(b.request.CustomerEmail),
		CustomerPhone:       null.StringFrom(b.request.CustomerPhone),
		CustomerCountry:     null.StringFrom(b.request.CustomerCountry),
		PaymentMethodId:     paymentMethodID,
		PaymentMethodType:   b.request.PaymentMethodType,
		Status:              paymentmodel.PaymentStatusInitiated,
		RoutingProfileId:    routingProfileID,
		ExpiresAt:           time.Now().UTC().Add(30 * time.Minute),
		Requires3ds:         b.request.Requires3DS,
		Description:         null.StringFrom(b.request.Description),
		StatementDescriptor: null.StringFrom(b.request.StatementDescriptor),
		Metadata:            metadataJSON,
		PromoDiscountAmount: decimal.Zero,
		IdempotencyKeyId:    correlationID,
		MetaCreatedAt:       time.Now().UTC(),
		MetaCreatedBy:       b.request.ActorID,
		MetaUpdatedAt:       time.Now().UTC(),
		MetaUpdatedBy:       nuuid.From(b.request.ActorID),
	}
	if b.request.CustomerIP != "" {
		parsedIP, err := netip.ParseAddr(b.request.CustomerIP)
		if err == nil {
			intent.CustomerIp = &parsedIP
		}
	}
	return intent, nil
}

func (b *PaymentFlowBuilder) buildPaymentAttempt(intent *paymentmodel.PaymentIntents) (*paymentmodel.Payments, error) {
	paymentID, err := uuid.NewV4()
	if err != nil {
		return nil, failure.InternalError(err)
	}
	paymentCode := businessCode("TXN")
	accountID := b.decision.PSPAccountID
	if accountID == uuid.Nil {
		return nil, failure.NotFound("psp account")
	}
	payment := &paymentmodel.Payments{
		Id:                    paymentID,
		PaymentCode:           paymentCode,
		IntentId:              intent.Id,
		AttemptNumber:         1,
		Psp:                   b.decision.PSP,
		Amount:                intent.Amount,
		Currency:              intent.Currency,
		SettlementCurrency:    intent.Currency,
		PaymentMethodId:       intent.PaymentMethodId,
		PaymentMethodType:     intent.PaymentMethodType,
		Status:                paymentmodel.PaymentStatusPending,
		ProcessingFeeCurrency: intent.Currency,
		Description:           intent.Description,
		Metadata:              intent.Metadata,
		MetaCreatedAt:         time.Now().UTC(),
		MetaCreatedBy:         b.request.ActorID,
		MetaUpdatedAt:         time.Now().UTC(),
		MetaUpdatedBy:         nuuid.From(b.request.ActorID),
	}
	return payment, nil
}

func businessCode(prefix string) string {
	stamp := time.Now().UTC().Format("20060102")
	shortID, _ := uuid.NewV4()
	compact := strings.ReplaceAll(shortID.String(), "-", "")
	if len(compact) > 6 {
		compact = compact[:6]
	}
	return fmt.Sprintf("%s-%s-%s", prefix, stamp, strings.ToUpper(compact))
}

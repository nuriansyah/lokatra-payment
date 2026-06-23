package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func (s *ServiceImpl) HandleWebhook(ctx context.Context, providerName string, headers http.Header, body []byte) (dto.WebhookReceipt, error) {
	provider, ok := webhookProvider(providerName)
	if !ok {
		return dto.WebhookReceipt{}, failure.BadRequestFromString("unsupported webhook provider")
	}
	accountID := s.providerAccountIDs[provider]
	if accountID == uuid.Nil {
		return dto.WebhookReceipt{}, failure.New(http.StatusFailedDependency, fmt.Errorf("webhook provider account is not configured"))
	}
	if !s.webhookConfigured[provider] {
		return dto.WebhookReceipt{}, failure.New(http.StatusFailedDependency, fmt.Errorf("webhook verification secret is not configured"))
	}
	gateway, err := s.gatewayRegistry.Get(provider)
	if err != nil {
		return dto.WebhookReceipt{}, failure.New(http.StatusFailedDependency, fmt.Errorf("webhook provider is disabled"))
	}
	verification, err := gateway.VerifyWebhook(ctx, pg.VerifyWebhookRequest{Headers: headers, RawBody: body})
	if err != nil || !verification.SignatureValid {
		return dto.WebhookReceipt{}, failure.Unauthorized("webhook signature verification failed")
	}
	event, err := gateway.NormalizeWebhook(ctx, pg.NormalizeWebhookRequest{Headers: headers, RawBody: body})
	if err != nil {
		return dto.WebhookReceipt{}, failure.BadRequest(err)
	}
	now := time.Now().UTC()
	receipt := dto.WebhookReceipt{
		Provider:       string(provider),
		EventID:        firstNonBlank(event.EventID, verification.EventID),
		EventType:      string(event.EventType),
		PaymentStatus:  string(event.PaymentStatus),
		OrderID:        event.OrderID,
		SignatureValid: true,
		ReceivedAt:     now,
	}
	if receipt.EventID != "" {
		if existing, found := s.findWebhookEvent(ctx, provider, receipt.EventID); found {
			receipt.ReceivedAt = existing.ReceivedAt
			return receipt, nil
		}
	}
	headerJSON, _ := json.Marshal(headers)
	parsedBody := event.Raw
	if len(parsedBody) == 0 {
		parsedBody = normalizedJSON(body)
	}
	sum := sha256.Sum256(body)
	record := paymentmodel.ProviderWebhookEvents{
		Id:                 mustUUID(),
		EndpointKey:        null.StringFrom(string(provider)),
		ProviderAccountId:  accountID,
		ProviderCode:       string(provider),
		EventId:            optionalString(receipt.EventID),
		EventType:          optionalString(receipt.EventType),
		ProviderReference:  optionalString(firstNonBlank(event.ProviderReference, event.ProviderTransactionID, event.OrderID)),
		ProviderStatus:     optionalString(event.ProviderStatus),
		SignatureValid:     true,
		SignatureAlgorithm: optionalString(webhookSignatureAlgorithm(provider)),
		Headers:            headerJSON,
		RawBody:            body,
		RawBodySha256:      hex.EncodeToString(sum[:]),
		ParsedBody:         parsedBody,
		ProcessingStatus:   paymentmodel.WebhookProcessingStatusReceived,
		ReceivedAt:         now,
		Metadata:           json.RawMessage(`{}`),
		MetaSignature:      shared.MetaSignature{MetaCreatedAt: now, MetaCreatedBy: uuid.Nil},
	}
	if err := s.paymentRepo.CreateProviderWebhookEvents(ctx, &record); err != nil {
		if receipt.EventID != "" {
			if _, found := s.findWebhookEvent(ctx, provider, receipt.EventID); found {
				return receipt, nil
			}
		}
		return dto.WebhookReceipt{}, err
	}
	return receipt, nil
}

func (s *ServiceImpl) findWebhookEvent(ctx context.Context, provider pg.ProviderCode, eventID string) (paymentmodel.ProviderWebhookEvents, bool) {
	result, err := s.paymentRepo.ResolveProviderWebhookEventsByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{
			{Field: string(paymentmodel.ProviderWebhookEventsDBFieldName.ProviderCode), Operator: paymentmodel.OperatorEqual, Value: string(provider)},
			{Field: string(paymentmodel.ProviderWebhookEventsDBFieldName.EventId), Operator: paymentmodel.OperatorEqual, Value: eventID},
		},
		Pagination: paymentmodel.Pagination{Page: 1, PageSize: 1},
	})
	if err != nil || len(result) == 0 {
		return paymentmodel.ProviderWebhookEvents{}, false
	}
	return result[0].ProviderWebhookEvents, true
}

func webhookProvider(value string) (pg.ProviderCode, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case string(pg.ProviderMidtrans):
		return pg.ProviderMidtrans, true
	case string(pg.ProviderXendit):
		return pg.ProviderXendit, true
	case string(pg.ProviderDurianpay):
		return pg.ProviderDurianpay, true
	default:
		return "", false
	}
}

func webhookSignatureAlgorithm(provider pg.ProviderCode) string {
	if provider == pg.ProviderMidtrans {
		return "SHA512"
	}
	if provider == pg.ProviderDurianpay {
		return "HMAC_SHA256"
	}
	return "TOKEN_OR_HMAC_SHA256"
}

func firstNonBlank(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/transport/http/response"
	"github.com/rs/zerolog/log"
)

// HandleMidtransWebhook godoc
// @Summary Handle Midtrans webhook event
// @Description Process incoming webhook from Midtrans payment service provider. Accepts Midtrans native webhook format, normalizes it, verifies signature, and updates payment status. Idempotent by transaction ID.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param event body map[string]interface{} true "Raw Midtrans webhook payload"
// @Success 200 {object} map[string]string "Webhook processed successfully"
// @Failure 400 {object} dto.ErrorResponseDTO "Invalid webhook payload"
// @Failure 401 {object} dto.ErrorResponseDTO "Webhook signature verification failed"
// @Failure 500 {object} dto.ErrorResponseDTO "Internal server error during webhook processing"
// @Router /v1/payments/webhooks/midtrans [post]
func (h *Handler) HandleMidtransWebhook(w http.ResponseWriter, r *http.Request) {
	// Accept raw Midtrans webhook payload
	var rawPayload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&rawPayload); err != nil {
		log.Error().Err(err).Msg("[Webhook][Midtrans] failed to decode request body")
		response.WithError(w, err)
		return
	}
	defer r.Body.Close()

	// Normalize Midtrans payload to standard webhook format
	event, err := normalizeMidtransWebhook(rawPayload)
	if err != nil {
		log.Error().Err(err).Msg("[Webhook][Midtrans] failed to normalize webhook payload")
		response.WithError(w, err)
		return
	}

	if err := shared.GetValidator().Struct(event); err != nil {
		log.Error().Err(err).Msg("[Webhook][Midtrans] request validation failed")
		response.WithError(w, err)
		return
	}

	// Verify Midtrans signature
	serviceEvent := toServiceWebhookEvent(event)
	if err := h.PaymentService.VerifyMidtransSignature(r.Context(), serviceEvent); err != nil {
		log.Warn().Err(err).Str("transactionId", event.TransactionID).Msg("[Webhook][Midtrans] signature verification failed")
		response.WithJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "webhook signature verification failed",
		})
		return
	}

	// Process webhook event
	result, err := h.PaymentService.ProcessWebhookEvent(r.Context(), serviceEvent)
	if err != nil {
		log.Error().Err(err).Str("transactionId", event.TransactionID).Str("provider", event.Provider).Msg("[Webhook][Midtrans] failed to process webhook event")
		response.WithError(w, err)
		return
	}

	log.Info().Str("transactionId", event.TransactionID).Str("paymentId", result.PaymentID).Str("status", result.FinalStatus).Msg("[Webhook][Midtrans] webhook processed successfully")
	response.WithJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "webhook processed",
	})
}

// HandleXenditWebhook godoc
// @Summary Handle Xendit webhook event
// @Description Process incoming webhook from Xendit payment service provider. Accepts Xendit native webhook format, normalizes it, verifies signature, and updates payment status.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param event body map[string]interface{} true "Raw Xendit webhook payload"
// @Success 200 {object} map[string]string "Webhook processed successfully"
// @Failure 400 {object} dto.ErrorResponseDTO "Invalid webhook payload"
// @Failure 401 {object} dto.ErrorResponseDTO "Webhook signature verification failed"
// @Failure 500 {object} dto.ErrorResponseDTO "Internal server error during webhook processing"
// @Router /v1/payments/webhooks/xendit [post]
func (h *Handler) HandleXenditWebhook(w http.ResponseWriter, r *http.Request) {
	// Accept raw Xendit webhook payload
	var rawPayload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&rawPayload); err != nil {
		log.Error().Err(err).Msg("[Webhook][Xendit] failed to decode request body")
		response.WithError(w, err)
		return
	}
	defer r.Body.Close()

	// Normalize Xendit payload to standard webhook format
	event, err := normalizeXenditWebhook(rawPayload)
	if err != nil {
		log.Error().Err(err).Msg("[Webhook][Xendit] failed to normalize webhook payload")
		response.WithError(w, err)
		return
	}

	if err := shared.GetValidator().Struct(event); err != nil {
		log.Error().Err(err).Msg("[Webhook][Xendit] request validation failed")
		response.WithError(w, err)
		return
	}

	// Verify Xendit signature
	serviceEvent := toServiceWebhookEvent(event)
	if err := h.PaymentService.VerifyXenditSignature(r.Context(), serviceEvent); err != nil {
		log.Warn().Err(err).Str("transactionId", event.TransactionID).Msg("[Webhook][Xendit] signature verification failed")
		response.WithJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "webhook signature verification failed",
		})
		return
	}

	// Process webhook event
	result, err := h.PaymentService.ProcessWebhookEvent(r.Context(), serviceEvent)
	if err != nil {
		log.Error().Err(err).Str("transactionId", event.TransactionID).Str("provider", event.Provider).Msg("[Webhook][Xendit] failed to process webhook event")
		response.WithError(w, err)
		return
	}

	log.Info().Str("transactionId", event.TransactionID).Str("paymentId", result.PaymentID).Str("status", result.FinalStatus).Msg("[Webhook][Xendit] webhook processed successfully")
	response.WithJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "webhook processed",
	})
}

// HandleWebhookTest godoc
// @Summary Test webhook endpoint
// @Description Health check endpoint for webhook delivery services to verify endpoint availability. Always returns 200 for configured webhooks.
// @Tags Webhooks
// @Produce json
// @Success 200 {object} map[string]string "Webhook endpoint is operational"
// @Router /v1/payments/webhooks/health [get]
func (h *Handler) HandleWebhookTest(w http.ResponseWriter, r *http.Request) {
	response.WithJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "webhook endpoint is operational",
	})
}

func toServiceWebhookEvent(event dto.WebhookEventRequestDTO) service.WebhookEvent {
	return service.WebhookEvent{
		EventType:       event.EventType,
		Provider:        event.Provider,
		TransactionID:   event.TransactionID,
		ExternalID:      event.ExternalID,
		Status:          event.Status,
		EventTimestamp:  event.EventTimestamp,
		RawPayload:      event.RawPayload,
		Signature:       event.Signature,
		SignatureMethod: event.SignatureMethod,
	}
}

func toWebhookEventResponseDTO(result service.WebhookResult) dto.WebhookEventResponseDTO {
	resp := dto.WebhookEventResponseDTO{
		Provider:      result.Provider,
		Success:       result.Success,
		EventID:       result.EventID,
		TransactionID: result.TransactionID,
		PaymentID:     result.PaymentID,
		IntentID:      result.IntentID,
		FinalStatus:   result.FinalStatus,
		Message:       result.Message,
		ProcessedAt:   result.ProcessedAt,
	}
	if result.Error != nil {
		resp.Error = &dto.WebhookErrorDTO{
			Code:      result.Error.Code,
			Message:   result.Error.Message,
			RequestID: result.Error.RequestID,
		}
	}
	return resp
}

// normalizeMidtransWebhook converts Midtrans native webhook format to normalized WebhookEventRequestDTO
func normalizeMidtransWebhook(raw map[string]interface{}) (dto.WebhookEventRequestDTO, error) {
	// Midtrans sends fields like: transaction_id, transaction_status, signature_key, gross_amount, etc.
	// For signature verification (SHA512), use: order_id + status_code + gross_amount + ServerKey
	transactionID := getStringField(raw, "transaction_id")
	transactionStatus := getStringField(raw, "transaction_status")
	signatureKey := getStringField(raw, "signature_key")
	transactionTime := getStringField(raw, "transaction_time")

	// Build the normalized event
	event := dto.WebhookEventRequestDTO{
		EventType:       "payment." + transactionStatus, // e.g., "payment.settlement"
		Provider:        "MIDTRANS",
		TransactionID:   transactionID,
		ExternalID:      getStringField(raw, "order_id"), // order_id is used in signature
		Status:          transactionStatus,
		EventTimestamp:  transactionTime, // Use transaction_time (always present), not settlement_time (can be null)
		RawPayload:      raw,
		Signature:       signatureKey,
		SignatureMethod: "SHA512",
	}
	return event, nil
}

// normalizeXenditWebhook converts Xendit native webhook format to normalized WebhookEventRequestDTO
func normalizeXenditWebhook(raw map[string]interface{}) (dto.WebhookEventRequestDTO, error) {
	// Xendit sends different field names, adjust as needed based on Xendit webhook format
	externalID := getStringField(raw, "reference_id")
	status := getStringField(raw, "status")

	event := dto.WebhookEventRequestDTO{
		EventType:       "payment." + status,
		Provider:        "XENDIT",
		TransactionID:   getStringField(raw, "id"),
		ExternalID:      externalID,
		Status:          status,
		EventTimestamp:  getStringField(raw, "created"),
		RawPayload:      raw,
		Signature:       getStringField(raw, "verification_code"),
		SignatureMethod: "HMACSHA256",
	}
	return event, nil
}

// getStringField safely extracts a string value from a map
func getStringField(m map[string]interface{}, key string) string {
	if val, exists := m[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

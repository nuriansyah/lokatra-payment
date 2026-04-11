package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
	"github.com/nuriansyah/lokatra-payment/transport/http/response"
)

// @Summary Get payment routing quote
// @Description Obtain a payment flow quote showing which PSP will be selected for the given payment parameters. This is a read-only operation that does not create any payment records.
// @Tags Payments
// @Accept json
// @Produce json
// @Security LokatraAuth
// @Param request body dto.PaymentFlowRequestDTO true "Payment flow request"
// @Success 200 {object} dto.PaymentFlowQuoteResponseDTO "Payment quote generated successfully"
// @Failure 400 {object} dto.ErrorResponseDTO "Invalid request parameters"
// @Failure 401 {object} dto.ErrorResponseDTO "Unauthorized"
// @Failure 500 {object} dto.ErrorResponseDTO "Internal server error"
// @Router /v1/payments/flows/quote [post]
func (h *Handler) quotePaymentFlow(w http.ResponseWriter, r *http.Request) {
	request, ok := h.decodePaymentFlowRequest(w, r)
	if !ok {
		return
	}

	quote, err := h.PaymentService.ResolvePaymentFlow(r.Context(), request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, quote)
}

// executePaymentFlow godoc
// @Summary Execute a payment with PSP routing and charge
// @Description Execute a full payment flow: create payment intent, route to appropriate PSP, attempt charge, and return execution result with complete telemetry. Creates immutable payment records in database.
// @Tags Payments
// @Accept json
// @Produce json
// @Security LokatraAuth
// @Param request body dto.PaymentFlowRequestDTO true "Payment flow request"
// @Success 201 {object} dto.PaymentExecutionResponseDTO "Payment executed with result"
// @Failure 400 {object} dto.ErrorResponseDTO "Invalid request parameters"
// @Failure 401 {object} dto.ErrorResponseDTO "Unauthorized"
// @Failure 409 {object} dto.ErrorResponseDTO "Conflict (idempotency key already processed)"
// @Failure 500 {object} dto.ErrorResponseDTO "Internal server error"
// @Router /v1/payments/flows/execute [post]
func (h *Handler) executePaymentFlow(w http.ResponseWriter, r *http.Request) {
	request, ok := h.decodePaymentFlowRequest(w, r)
	if !ok {
		return
	}

	if strings.TrimSpace(request.IdempotencyKey) == "" {
		request.IdempotencyKey = strings.TrimSpace(r.Header.Get(middleware.IdempotencyHeader))
	}

	result, err := h.PaymentService.ExecutePayment(r.Context(), request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusCreated, result)
}

func (h *Handler) decodePaymentFlowRequest(w http.ResponseWriter, r *http.Request) (service.PaymentFlowRequest, bool) {
	var request service.PaymentFlowRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.WithError(w, err)
		return service.PaymentFlowRequest{}, false
	}
	if err := shared.GetValidator().Struct(request); err != nil {
		response.WithError(w, err)
		return service.PaymentFlowRequest{}, false
	}
	return request, true
}

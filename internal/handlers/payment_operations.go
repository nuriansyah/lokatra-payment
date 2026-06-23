package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
	"github.com/nuriansyah/lokatra-payment/transport/http/response"
)

const maxOperationBodyBytes = 1 << 20

func (h *Handler) CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	var request dto.CreatePaymentIntentRequest
	if !decodeOperationBody(w, r, &request) {
		return
	}
	if strings.TrimSpace(request.IdempotencyKey) == "" {
		request.IdempotencyKey = r.Header.Get(middleware.IdempotencyHeader)
	}
	result, err := h.PaymentService.CreatePaymentIntent(r.Context(), request)
	writeOperationResult(w, http.StatusCreated, dto.NewPaymentIntentsResponse(result), err)
}

func (h *Handler) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	id, ok := operationID(w, r, "paymentIntentID")
	if !ok {
		return
	}
	result, err := h.PaymentService.GetPaymentIntent(r.Context(), id)
	writeOperationResult(w, http.StatusOK, dto.NewPaymentIntentsResponse(result), err)
}

func (h *Handler) PaymentIntentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "paymentIntentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyPaymentIntentAction(r.Context(), id, action, command)
		return newPaymentIntentActionResponse(result), err
	})
}

func newPaymentIntentActionResponse(result service.PaymentIntentActionResult) dto.PaymentIntentActionResponse {
	response := dto.PaymentIntentActionResponse{Intent: dto.NewPaymentIntentsResponse(result.Intent)}
	if result.Routing == nil {
		return response
	}
	routing := &dto.RoutingExecutionResponse{
		Candidates: make([]dto.RoutingCandidateResponse, 0, len(result.Routing.Candidates)),
		Attempts:   make([]dto.ProviderAttemptResponse, 0, len(result.Routing.Attempts)),
	}
	for _, candidate := range result.Routing.Candidates {
		routing.Candidates = append(routing.Candidates, dto.RoutingCandidateResponse{
			ProviderCode: string(candidate.ProviderCode), AccountID: candidate.AccountID, Priority: candidate.Priority,
			MaxAttempts: candidate.MaxAttempts, Reason: candidate.Reason, Skipped: candidate.Skipped, SkipReason: candidate.SkipReason,
		})
	}
	for _, attempt := range result.Routing.Attempts {
		routing.Attempts = append(routing.Attempts, dto.ProviderAttemptResponse{
			ProviderCode: string(attempt.ProviderCode), AccountID: attempt.AccountID, Attempt: attempt.Attempt,
			StartedAt: attempt.StartedAt, Duration: attempt.Duration, Error: attempt.Error,
		})
	}
	if result.Routing.Selected.ProviderCode != "" {
		selected := result.Routing.Selected
		routing.Selected = &dto.RoutingCandidateResponse{
			ProviderCode: string(selected.ProviderCode), AccountID: selected.AccountID, Priority: selected.Priority,
			MaxAttempts: selected.MaxAttempts, Reason: selected.Reason, Skipped: selected.Skipped, SkipReason: selected.SkipReason,
		}
		payment := result.Routing.Payment
		routing.Payment = &payment
	}
	response.Routing = routing
	return response
}

func (h *Handler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	provider := strings.ToLower(strings.TrimSpace(chi.URLParam(r, "provider")))
	r.Body = http.MaxBytesReader(w, r.Body, maxOperationBodyBytes)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	result, err := h.PaymentService.HandleWebhook(r.Context(), provider, r.Header.Clone(), body)
	writeOperationResult(w, http.StatusAccepted, result, err)
}

func (h *Handler) CreateRefund(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateRefundRequest
	if !decodeOperationBody(w, r, &request) {
		return
	}
	if strings.TrimSpace(request.IdempotencyKey) == "" {
		request.IdempotencyKey = r.Header.Get(middleware.IdempotencyHeader)
	}
	result, err := h.PaymentService.CreateRefund(r.Context(), request)
	writeOperationResult(w, http.StatusCreated, dto.NewPaymentRefundsResponse(result), err)
}

func (h *Handler) RefundAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "refundID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyRefundAction(r.Context(), id, action, command)
		return dto.NewPaymentRefundsResponse(result), err
	})
}

func (h *Handler) AdminWebhookAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "webhookID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyWebhookAction(r.Context(), id, action, command)
		return dto.NewProviderWebhookEventsResponse(result), err
	})
}

func (h *Handler) ManualPaymentEvidenceAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "evidenceID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyManualEvidenceAction(r.Context(), id, action, command)
		return dto.NewManualPaymentEvidenceResponse(result), err
	})
}

func (h *Handler) OverpaymentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "overpaymentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyOverpaymentAction(r.Context(), id, action, command)
		return dto.NewPaymentOverpaymentsResponse(result), err
	})
}

func (h *Handler) OpenCashCollectionSession(w http.ResponseWriter, r *http.Request) {
	var request dto.OpenCashSessionRequest
	if !decodeOperationBody(w, r, &request) {
		return
	}
	result, err := h.PaymentService.OpenCashSession(r.Context(), request)
	writeOperationResult(w, http.StatusCreated, dto.NewCashCollectionSessionsResponse(result), err)
}

func (h *Handler) CashCollectionSessionAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "cashSessionID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyCashSessionAction(r.Context(), id, action, command)
		return dto.NewCashCollectionSessionsResponse(result), err
	})
}

func (h *Handler) PaymentInstallmentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "installmentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyInstallmentAction(r.Context(), id, action, command)
		return dto.NewPaymentInstallmentsResponse(result), err
	})
}

func (h *Handler) PaymentAuthorizationAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "authorizationID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		result, err := h.PaymentService.ApplyAuthorizationAction(r.Context(), id, action, command)
		return dto.NewPaymentAuthorizationsResponse(result), err
	})
}

func (h *Handler) handleAction(w http.ResponseWriter, r *http.Request, param string, apply func(uuid.UUID, string, dto.ActionCommand) (any, error)) {
	id, ok := operationID(w, r, param)
	if !ok {
		return
	}
	action := strings.ToLower(strings.TrimSpace(chi.URLParam(r, "action")))
	if action == "" {
		response.WithError(w, failure.BadRequestFromString("action is required"))
		return
	}
	var command dto.ActionCommand
	if r.ContentLength != 0 && !decodeOperationBody(w, r, &command) {
		return
	}
	result, err := apply(id, action, command)
	writeOperationResult(w, http.StatusOK, result, err)
}

func operationID(w http.ResponseWriter, r *http.Request, param string) (uuid.UUID, bool) {
	id, err := uuid.FromString(chi.URLParam(r, param))
	if err != nil || id == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("%s must be a valid UUID", param)))
		return uuid.Nil, false
	}
	return id, true
}

func decodeOperationBody(w http.ResponseWriter, r *http.Request, destination any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxOperationBodyBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return false
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		response.WithError(w, failure.BadRequestFromString("request body must contain one JSON object"))
		return false
	}
	return true
}

func writeOperationResult(w http.ResponseWriter, status int, result any, err error) {
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, status, result)
}

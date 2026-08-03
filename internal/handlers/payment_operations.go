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
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
	"github.com/nuriansyah/lokatra-payment/transport/http/response"
)

const maxOperationBodyBytes = 1 << 20

func (h *Handler) CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req dto.CreatePaymentIntentRequest
	if err := decoder.Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	if strings.TrimSpace(req.IdempotencyKey) == "" {
		req.IdempotencyKey = r.Header.Get(middleware.IdempotencyHeader)
	}

	result, err := h.PaymentService.CreatePaymentIntent(r.Context(), req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, result)
}

func (h *Handler) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	id, ok := operationID(w, r, "paymentIntentID")
	if !ok {
		return
	}

	result, err := h.PaymentService.GetPaymentIntent(r.Context(), id)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

func (h *Handler) PaymentIntentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "paymentIntentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyPaymentIntentAction(r.Context(), id, action, command)
	})
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
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusAccepted, result)
}

func (h *Handler) CreateRefund(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req dto.CreateRefundRequest
	if err := decoder.Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	if strings.TrimSpace(req.IdempotencyKey) == "" {
		req.IdempotencyKey = r.Header.Get(middleware.IdempotencyHeader)
	}

	result, err := h.PaymentService.CreateRefund(r.Context(), req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, result)
}

func (h *Handler) RefundAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "refundID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyRefundAction(r.Context(), id, action, command)
	})
}

func (h *Handler) AdminWebhookAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "webhookID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyWebhookAction(r.Context(), id, action, command)
	})
}

func (h *Handler) ManualPaymentEvidenceAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "evidenceID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyManualEvidenceAction(r.Context(), id, action, command)
	})
}

func (h *Handler) OverpaymentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "overpaymentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyOverpaymentAction(r.Context(), id, action, command)
	})
}

func (h *Handler) OpenCashCollectionSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req dto.OpenCashSessionRequest
	if err := decoder.Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.PaymentService.OpenCashSession(r.Context(), req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, result)
}

func (h *Handler) CashCollectionSessionAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "cashSessionID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyCashSessionAction(r.Context(), id, action, command)
	})
}

func (h *Handler) PaymentInstallmentAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "installmentID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyInstallmentAction(r.Context(), id, action, command)
	})
}

func (h *Handler) PaymentAuthorizationAction(w http.ResponseWriter, r *http.Request) {
	h.handleAction(w, r, "authorizationID", func(id uuid.UUID, action string, command dto.ActionCommand) (any, error) {
		return h.PaymentService.ApplyAuthorizationAction(r.Context(), id, action, command)
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
	if r.ContentLength != 0 {
		if !decodeOperationBody(w, r, &command) {
			return
		}
		if err := command.Validate(); err != nil {
			response.WithError(w, err)
			return
		}
	}
	result, err := apply(id, action, command)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, result)
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

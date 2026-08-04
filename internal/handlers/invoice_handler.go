package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/transport/http/response"
)

const maxInvoiceBodyBytes = 1 << 20

type InvoiceHandler struct {
	paymentService *service.ServiceImpl
}

func NewInvoiceHandler(paymentService *service.ServiceImpl) *InvoiceHandler {
	return &InvoiceHandler{paymentService: paymentService}
}

func (h *InvoiceHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.CreateInvoice)
	r.Get("/", h.ListInvoices)
	r.Get("/{invoiceID}", h.GetInvoice)
	r.Post("/{invoiceID}/issue", h.IssueInvoice)
	r.Post("/{invoiceID}/void", h.VoidInvoice)
	r.Post("/{invoiceID}/payments", h.RecordPayment)
	r.Post("/{invoiceID}/payment-links", h.GeneratePaymentLinks)
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
	var req dto.CreateInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	actorID, _ := uuid.FromString(r.Header.Get("X-Actor-ID"))
	req.ActorID = actorID

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.paymentService.CreateInvoice(r.Context(), req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, result)
}

func (h *InvoiceHandler) ListInvoices(w http.ResponseWriter, r *http.Request) {
	filter := model.Filter{
		Pagination: model.Pagination{Page: 1, PageSize: 20},
	}

	if r.ContentLength != 0 {
		r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
		_ = json.NewDecoder(r.Body).Decode(&filter)
	}

	results, _, err := h.paymentService.ListInvoices(r.Context(), filter)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, results)
}

func (h *InvoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID, err := uuid.FromString(chi.URLParam(r, "invoiceID"))
	if err != nil || invoiceID == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("invoiceID must be a valid UUID")))
		return
	}

	result, err := h.paymentService.GetInvoice(r.Context(), invoiceID)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

func (h *InvoiceHandler) IssueInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID, err := uuid.FromString(chi.URLParam(r, "invoiceID"))
	if err != nil || invoiceID == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("invoiceID must be a valid UUID")))
		return
	}

	actorID, _ := uuid.FromString(r.Header.Get("X-Actor-ID"))

	result, err := h.paymentService.IssueInvoice(r.Context(), invoiceID, actorID)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

func (h *InvoiceHandler) VoidInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID, err := uuid.FromString(chi.URLParam(r, "invoiceID"))
	if err != nil || invoiceID == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("invoiceID must be a valid UUID")))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
	var req dto.VoidInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	actorID, _ := uuid.FromString(r.Header.Get("X-Actor-ID"))
	req.ActorID = actorID

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.paymentService.VoidInvoice(r.Context(), invoiceID, req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

func (h *InvoiceHandler) RecordPayment(w http.ResponseWriter, r *http.Request) {
	invoiceID, err := uuid.FromString(chi.URLParam(r, "invoiceID"))
	if err != nil || invoiceID == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("invoiceID must be a valid UUID")))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
	var req dto.RecordPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	actorID, _ := uuid.FromString(r.Header.Get("X-Actor-ID"))
	req.ActorID = actorID

	if err := req.Validate(); err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.paymentService.RecordInvoicePayment(r.Context(), req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

func (h *InvoiceHandler) GeneratePaymentLinks(w http.ResponseWriter, r *http.Request) {
	invoiceID, err := uuid.FromString(chi.URLParam(r, "invoiceID"))
	if err != nil || invoiceID == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("invoiceID must be a valid UUID")))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
	var req dto.GeneratePaymentLinksRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	actorID, _ := uuid.FromString(r.Header.Get("X-Actor-ID"))
	req.ActorID = actorID

	result, err := h.paymentService.GenerateInvoicePaymentLinks(r.Context(), invoiceID, req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, result)
}

// operationID is a helper to parse UUID from URL params.
func invoiceOperationID(w http.ResponseWriter, r *http.Request, param string) (uuid.UUID, bool) {
	id, err := uuid.FromString(chi.URLParam(r, param))
	if err != nil || id == uuid.Nil {
		response.WithError(w, failure.BadRequest(fmt.Errorf("%s must be a valid UUID", param)))
		return uuid.Nil, false
	}
	return id, true
}

// decodeInvoiceBody is a helper to decode request body.
func decodeInvoiceBody(w http.ResponseWriter, r *http.Request, destination any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxInvoiceBodyBytes)
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

// Ensure shared import is used
var _ = shared.ErrInvalidID

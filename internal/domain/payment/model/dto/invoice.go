package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ──────────────────────────────────────────────────────────────────────────────
// Request DTOs
// ──────────────────────────────────────────────────────────────────────────────

type CreateInvoiceRequest struct {
	ActorID       uuid.UUID
	MerchantID    uuid.UUID
	CustomerID    uuid.UUID
	CustomerName  string
	CustomerEmail string
	CustomerPhone string
	Currency      string
	Description   string
	Notes         string
	Terms         string
	DueAt         *time.Time
	Pph23Rate     *decimal.Decimal
	PpnRate       *decimal.Decimal
	LineItems     []LineItemRequest
}

type LineItemRequest struct {
	Name            string
	Description     string
	Quantity        decimal.Decimal
	UnitPrice       decimal.Decimal
	DiscountPercent *decimal.Decimal
	Pph23Rate       *decimal.Decimal
	PpnRate         *decimal.Decimal
	SKU             string
	Category        string
}

type RecordPaymentRequest struct {
	ActorID         uuid.UUID
	PaymentIntentID uuid.UUID
	ProviderCode    string
	Amount          decimal.Decimal
	Currency        string
}

type GeneratePaymentLinksRequest struct {
	ActorID  uuid.UUID
	Provider []string
}

type VoidInvoiceRequest struct {
	ActorID uuid.UUID
	Reason  string
}

// ──────────────────────────────────────────────────────────────────────────────
// Response DTOs
// ──────────────────────────────────────────────────────────────────────────────

type InvoiceResponse struct {
	Id              uuid.UUID        `json:"id"`
	InvoiceCode     string           `json:"invoiceCode"`
	MerchantId      uuid.UUID        `json:"merchantId"`
	CustomerId      uuid.UUID        `json:"customerId"`
	CustomerName    string           `json:"customerName,omitempty"`
	CustomerEmail   string           `json:"customerEmail,omitempty"`
	CustomerPhone   string           `json:"customerPhone,omitempty"`
	Subtotal        decimal.Decimal  `json:"subtotal"`
	DiscountAmount  decimal.Decimal  `json:"discountAmount"`
	Pph23Amount     decimal.Decimal  `json:"pph23Amount"`
	Pph23Rate       decimal.Decimal  `json:"pph23Rate"`
	PpnAmount       decimal.Decimal  `json:"ppnAmount"`
	PpnRate         decimal.Decimal  `json:"ppnRate"`
	TotalAmount     decimal.Decimal  `json:"totalAmount"`
	PaidAmount      decimal.Decimal  `json:"paidAmount"`
	RemainingAmount decimal.Decimal  `json:"remainingAmount"`
	Currency        string           `json:"currency"`
	Status          string           `json:"status"`
	IssuedAt        *time.Time       `json:"issuedAt,omitempty"`
	DueAt           *time.Time       `json:"dueAt,omitempty"`
	PaidAt          *time.Time       `json:"paidAt,omitempty"`
	VoidedAt        *time.Time       `json:"voidedAt,omitempty"`
	VoidReason      string           `json:"voidReason,omitempty"`
	Description     string           `json:"description,omitempty"`
	Notes           string           `json:"notes,omitempty"`
	Terms           string           `json:"terms,omitempty"`
	MetaCreatedAt   time.Time        `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID        `json:"metaCreatedBy"`
}

type InvoiceLineItemResponse struct {
	Id              uuid.UUID       `json:"id"`
	LineNo          int             `json:"lineNo"`
	Name            string          `json:"name"`
	Description     string          `json:"description,omitempty"`
	Quantity        decimal.Decimal `json:"quantity"`
	UnitPrice       decimal.Decimal `json:"unitPrice"`
	DiscountPercent decimal.Decimal `json:"discountPercent"`
	DiscountAmount  decimal.Decimal `json:"discountAmount"`
	PpnRate         decimal.Decimal `json:"ppnRate"`
	PpnAmount       decimal.Decimal `json:"ppnAmount"`
	Pph23Rate       decimal.Decimal `json:"pph23Rate"`
	Pph23Amount     decimal.Decimal `json:"pph23Amount"`
	Subtotal        decimal.Decimal `json:"subtotal"`
	TotalAmount     decimal.Decimal `json:"totalAmount"`
	Currency        string          `json:"currency"`
	Sku             string          `json:"sku,omitempty"`
	Category        string          `json:"category,omitempty"`
}

type InvoicePaymentLinkResponse struct {
	Id          uuid.UUID  `json:"id"`
	ProviderCode string    `json:"providerCode"`
	LinkUrl     string     `json:"linkUrl"`
	LinkType    string     `json:"linkType"`
	Status      string     `json:"status"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
}

type InvoicePaymentResponse struct {
	Id              uuid.UUID       `json:"id"`
	PaymentIntentId uuid.UUID       `json:"paymentIntentId"`
	ProviderCode    string          `json:"providerCode"`
	Amount          decimal.Decimal `json:"amount"`
	Currency        string          `json:"currency"`
	Status          string          `json:"status"`
}

type InvoiceDetailResponse struct {
	InvoiceResponse
	LineItems    []InvoiceLineItemResponse    `json:"lineItems"`
	PaymentLinks []InvoicePaymentLinkResponse `json:"paymentLinks"`
	Payments     []InvoicePaymentResponse     `json:"payments"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Mappers
// ──────────────────────────────────────────────────────────────────────────────

func NewInvoiceResponse(inv paymentmodel.Invoices) InvoiceResponse {
	resp := InvoiceResponse{
		Id:              inv.Id,
		InvoiceCode:     inv.InvoiceCode,
		MerchantId:      inv.MerchantId,
		CustomerId:      inv.CustomerId.UUID,
		Subtotal:        inv.Subtotal,
		DiscountAmount:  inv.DiscountAmount,
		Pph23Amount:     inv.Pph23Amount,
		Pph23Rate:       inv.Pph23Rate,
		PpnAmount:       inv.PpnAmount,
		PpnRate:         inv.PpnRate,
		TotalAmount:     inv.TotalAmount,
		PaidAmount:      inv.PaidAmount,
		RemainingAmount: inv.RemainingAmount,
		Currency:        inv.Currency,
		Status:          string(inv.Status),
		MetaCreatedAt:   inv.MetaCreatedAt,
		MetaCreatedBy:   inv.MetaCreatedBy,
	}
	if inv.CustomerName.Valid {
		resp.CustomerName = inv.CustomerName.String
	}
	if inv.CustomerEmail.Valid {
		resp.CustomerEmail = inv.CustomerEmail.String
	}
	if inv.CustomerPhone.Valid {
		resp.CustomerPhone = inv.CustomerPhone.String
	}
	if inv.IssuedAt.Valid {
		resp.IssuedAt = &inv.IssuedAt.Time
	}
	if inv.DueAt.Valid {
		resp.DueAt = &inv.DueAt.Time
	}
	if inv.PaidAt.Valid {
		resp.PaidAt = &inv.PaidAt.Time
	}
	if inv.VoidedAt.Valid {
		resp.VoidedAt = &inv.VoidedAt.Time
	}
	if inv.VoidReason.Valid {
		resp.VoidReason = inv.VoidReason.String
	}
	if inv.Description.Valid {
		resp.Description = inv.Description.String
	}
	if inv.Notes.Valid {
		resp.Notes = inv.Notes.String
	}
	if inv.Terms.Valid {
		resp.Terms = inv.Terms.String
	}
	return resp
}

func NewInvoiceLineItemResponse(item paymentmodel.InvoiceLineItems) InvoiceLineItemResponse {
	resp := InvoiceLineItemResponse{
		Id:              item.Id,
		LineNo:          item.LineNo,
		Name:            item.Name,
		Quantity:        item.Quantity,
		UnitPrice:       item.UnitPrice,
		DiscountPercent: item.DiscountPercent,
		DiscountAmount:  item.DiscountAmount,
		PpnRate:         item.PpnRate,
		PpnAmount:       item.PpnAmount,
		Pph23Rate:       item.Pph23Rate,
		Pph23Amount:     item.Pph23Amount,
		Subtotal:        item.Subtotal,
		TotalAmount:     item.TotalAmount,
		Currency:        item.Currency,
	}
	if item.Description.Valid {
		resp.Description = item.Description.String
	}
	if item.Sku.Valid {
		resp.Sku = item.Sku.String
	}
	if item.Category.Valid {
		resp.Category = item.Category.String
	}
	return resp
}

func NewInvoicePaymentLinkResponse(link paymentmodel.InvoicePaymentLinks) InvoicePaymentLinkResponse {
	resp := InvoicePaymentLinkResponse{
		Id:           link.Id,
		ProviderCode: link.ProviderCode,
		LinkUrl:      link.LinkUrl,
		LinkType:     link.LinkType,
		Status:       string(link.Status),
	}
	if link.ExpiresAt.Valid {
		resp.ExpiresAt = &link.ExpiresAt.Time
	}
	return resp
}

func NewInvoicePaymentResponse(pay paymentmodel.InvoicePayments) InvoicePaymentResponse {
	return InvoicePaymentResponse{
		Id:              pay.Id,
		PaymentIntentId: pay.PaymentIntentId,
		ProviderCode:    pay.ProviderCode,
		Amount:          pay.Amount,
		Currency:        pay.Currency,
		Status:          string(pay.Status),
	}
}

// ──────────────────────────────────────────────────────────────────────────────
// Validation
// ──────────────────────────────────────────────────────────────────────────────

func (r *CreateInvoiceRequest) Validate() error {
	if r.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "actorId is required")
	}
	if r.MerchantID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "merchantId is required")
	}
	currency := strings.ToUpper(strings.TrimSpace(r.Currency))
	if currency == "" {
		return failure.BadRequestFromString("currency is required")
	}
	if len(r.LineItems) == 0 {
		return failure.BadRequestFromString("at least one line item is required")
	}
	for i, item := range r.LineItems {
		if strings.TrimSpace(item.Name) == "" {
			return failure.BadRequestFromString(fmt.Sprintf("line item[%d].name is required", i))
		}
		if item.Quantity.IsNegative() || item.Quantity.IsZero() {
			return failure.BadRequestFromString(fmt.Sprintf("line item[%d].quantity must be positive", i))
		}
		if item.UnitPrice.IsNegative() {
			return failure.BadRequestFromString(fmt.Sprintf("line item[%d].unitPrice must not be negative", i))
		}
	}
	return nil
}

func (r *RecordPaymentRequest) Validate() error {
	if r.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "actorId is required")
	}
	if r.PaymentIntentID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "paymentIntentId is required")
	}
	if strings.TrimSpace(r.ProviderCode) == "" {
		return failure.BadRequestFromString("providerCode is required")
	}
	if r.Amount.IsNegative() || r.Amount.IsZero() {
		return failure.BadRequestFromString("amount must be positive")
	}
	return nil
}

func (r *VoidInvoiceRequest) Validate() error {
	if r.ActorID == uuid.Nil {
		return failure.WithCode(shared.ErrInvalidID, "actorId is required")
	}
	if strings.TrimSpace(r.Reason) == "" {
		return failure.BadRequestFromString("reason is required to void an invoice")
	}
	return nil
}

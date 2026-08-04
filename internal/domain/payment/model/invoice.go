package model

import (
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type InvoicesDBFieldNameType string

type invoicesDBFieldName struct {
	Id               InvoicesDBFieldNameType
	InvoiceCode      InvoicesDBFieldNameType
	MerchantId       InvoicesDBFieldNameType
	CustomerId       InvoicesDBFieldNameType
	CustomerName     InvoicesDBFieldNameType
	CustomerEmail    InvoicesDBFieldNameType
	CustomerPhone    InvoicesDBFieldNameType
	Subtotal         InvoicesDBFieldNameType
	DiscountAmount   InvoicesDBFieldNameType
	Pph23Amount      InvoicesDBFieldNameType
	Pph23Rate        InvoicesDBFieldNameType
	PpnAmount        InvoicesDBFieldNameType
	PpnRate          InvoicesDBFieldNameType
	TotalAmount      InvoicesDBFieldNameType
	PaidAmount       InvoicesDBFieldNameType
	RemainingAmount  InvoicesDBFieldNameType
	Currency         InvoicesDBFieldNameType
	Status           InvoicesDBFieldNameType
	IssuedAt         InvoicesDBFieldNameType
	DueAt            InvoicesDBFieldNameType
	PaidAt           InvoicesDBFieldNameType
	VoidedAt         InvoicesDBFieldNameType
	VoidReason       InvoicesDBFieldNameType
	SourceService    InvoicesDBFieldNameType
	SourceType       InvoicesDBFieldNameType
	SourceId         InvoicesDBFieldNameType
	Description      InvoicesDBFieldNameType
	Notes            InvoicesDBFieldNameType
	Terms            InvoicesDBFieldNameType
	MetaCreatedAt    InvoicesDBFieldNameType
	MetaCreatedBy    InvoicesDBFieldNameType
	MetaUpdatedAt    InvoicesDBFieldNameType
	MetaUpdatedBy    InvoicesDBFieldNameType
	MetaDeletedAt    InvoicesDBFieldNameType
	MetaDeletedBy    InvoicesDBFieldNameType
}

var InvoicesDBFieldName = invoicesDBFieldName{
	Id:              "id",
	InvoiceCode:     "invoice_code",
	MerchantId:      "merchant_id",
	CustomerId:      "customer_id",
	CustomerName:    "customer_name",
	CustomerEmail:   "customer_email",
	CustomerPhone:   "customer_phone",
	Subtotal:        "subtotal",
	DiscountAmount:  "discount_amount",
	Pph23Amount:     "pph23_amount",
	Pph23Rate:       "pph23_rate",
	PpnAmount:       "ppn_amount",
	PpnRate:         "ppn_rate",
	TotalAmount:     "total_amount",
	PaidAmount:      "paid_amount",
	RemainingAmount: "remaining_amount",
	Currency:        "currency",
	Status:          "status",
	IssuedAt:        "issued_at",
	DueAt:           "due_at",
	PaidAt:          "paid_at",
	VoidedAt:        "voided_at",
	VoidReason:      "void_reason",
	SourceService:   "source_service",
	SourceType:      "source_type",
	SourceId:        "source_id",
	Description:     "description",
	Notes:           "notes",
	Terms:           "terms",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

type InvoiceStatus string

const (
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusIssued        InvoiceStatus = "issued"
	InvoiceStatusPartiallyPaid InvoiceStatus = "partially_paid"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusOverdue       InvoiceStatus = "overdue"
	InvoiceStatusVoided        InvoiceStatus = "voided"
	InvoiceStatusWrittenOff    InvoiceStatus = "written_off"
)

type Invoices struct {
	Id              uuid.UUID        `db:"id"`
	InvoiceCode     string           `db:"invoice_code"`
	MerchantId      uuid.UUID        `db:"merchant_id"`
	CustomerId      nuuid.NUUID       `db:"customer_id"`
	CustomerName    null.String      `db:"customer_name"`
	CustomerEmail   null.String      `db:"customer_email"`
	CustomerPhone   null.String      `db:"customer_phone"`
	Subtotal        decimal.Decimal  `db:"subtotal"`
	DiscountAmount  decimal.Decimal  `db:"discount_amount"`
	Pph23Amount     decimal.Decimal  `db:"pph23_amount"`
	Pph23Rate       decimal.Decimal  `db:"pph23_rate"`
	PpnAmount       decimal.Decimal  `db:"ppn_amount"`
	PpnRate         decimal.Decimal  `db:"ppn_rate"`
	TotalAmount     decimal.Decimal  `db:"total_amount"`
	PaidAmount      decimal.Decimal  `db:"paid_amount"`
	RemainingAmount decimal.Decimal  `db:"remaining_amount"`
	Currency        string           `db:"currency"`
	Status          InvoiceStatus    `db:"status"`
	IssuedAt        null.Time        `db:"issued_at"`
	DueAt           null.Time        `db:"due_at"`
	PaidAt          null.Time        `db:"paid_at"`
	VoidedAt        null.Time        `db:"voided_at"`
	VoidReason      null.String      `db:"void_reason"`
	SourceService   null.String      `db:"source_service"`
	SourceType      null.String      `db:"source_type"`
	SourceId        uuid.UUID        `db:"source_id"`
	Description     null.String      `db:"description"`
	Notes           null.String      `db:"notes"`
	Terms           null.String      `db:"terms"`

	shared.MetaSignature
}

type InvoicesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Invoices) ToInvoicesPrimaryID() InvoicesPrimaryID {
	return InvoicesPrimaryID{Id: d.Id}
}

type InvoicesList []*Invoices

type InvoicesFilterResult struct {
	Invoices
}

type InvoicesFilterResultList []InvoicesFilterResult

package model

import (
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/shopspring/decimal"
)

type InvoicePaymentsDBFieldNameType string

type invoicePaymentsDBFieldName struct {
	Id              InvoicePaymentsDBFieldNameType
	InvoiceId       InvoicePaymentsDBFieldNameType
	PaymentIntentId InvoicePaymentsDBFieldNameType
	ProviderCode    InvoicePaymentsDBFieldNameType
	Amount          InvoicePaymentsDBFieldNameType
	Currency        InvoicePaymentsDBFieldNameType
	Status          InvoicePaymentsDBFieldNameType
	MetaCreatedAt   InvoicePaymentsDBFieldNameType
	MetaCreatedBy   InvoicePaymentsDBFieldNameType
	MetaUpdatedAt   InvoicePaymentsDBFieldNameType
	MetaUpdatedBy   InvoicePaymentsDBFieldNameType
}

var InvoicePaymentsDBFieldName = invoicePaymentsDBFieldName{
	Id:              "id",
	InvoiceId:       "invoice_id",
	PaymentIntentId: "payment_intent_id",
	ProviderCode:    "provider_code",
	Amount:          "amount",
	Currency:        "currency",
	Status:          "status",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
}

type InvoicePaymentStatus string

const (
	InvoicePaymentStatusPending    InvoicePaymentStatus = "pending"
	InvoicePaymentStatusProcessing InvoicePaymentStatus = "processing"
	InvoicePaymentStatusSucceeded  InvoicePaymentStatus = "succeeded"
	InvoicePaymentStatusFailed     InvoicePaymentStatus = "failed"
	InvoicePaymentStatusRefunded   InvoicePaymentStatus = "refunded"
)

type InvoicePayments struct {
	Id              uuid.UUID           `db:"id"`
	InvoiceId       uuid.UUID           `db:"invoice_id"`
	PaymentIntentId uuid.UUID           `db:"payment_intent_id"`
	ProviderCode    string              `db:"provider_code"`
	Amount          decimal.Decimal     `db:"amount"`
	Currency        string              `db:"currency"`
	Status          InvoicePaymentStatus `db:"status"`

	shared.MetaSignature
}

type InvoicePaymentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d InvoicePayments) ToInvoicePaymentsPrimaryID() InvoicePaymentsPrimaryID {
	return InvoicePaymentsPrimaryID{Id: d.Id}
}

type InvoicePaymentsList []*InvoicePayments

type InvoicePaymentsFilterResult struct {
	InvoicePayments
}

type InvoicePaymentsFilterResultList []InvoicePaymentsFilterResult

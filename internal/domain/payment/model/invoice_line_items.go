package model

import (
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/shopspring/decimal"
)

type InvoiceLineItemsDBFieldNameType string

type invoiceLineItemsDBFieldName struct {
	Id             InvoiceLineItemsDBFieldNameType
	InvoiceId      InvoiceLineItemsDBFieldNameType
	LineNo         InvoiceLineItemsDBFieldNameType
	Name           InvoiceLineItemsDBFieldNameType
	Description    InvoiceLineItemsDBFieldNameType
	Quantity       InvoiceLineItemsDBFieldNameType
	UnitPrice      InvoiceLineItemsDBFieldNameType
	DiscountPercent InvoiceLineItemsDBFieldNameType
	DiscountAmount InvoiceLineItemsDBFieldNameType
	PpnRate        InvoiceLineItemsDBFieldNameType
	PpnAmount      InvoiceLineItemsDBFieldNameType
	Pph23Rate      InvoiceLineItemsDBFieldNameType
	Pph23Amount    InvoiceLineItemsDBFieldNameType
	Subtotal       InvoiceLineItemsDBFieldNameType
	TotalAmount    InvoiceLineItemsDBFieldNameType
	Currency       InvoiceLineItemsDBFieldNameType
	Sku            InvoiceLineItemsDBFieldNameType
	Category       InvoiceLineItemsDBFieldNameType
	MetaCreatedAt  InvoiceLineItemsDBFieldNameType
	MetaCreatedBy  InvoiceLineItemsDBFieldNameType
	MetaUpdatedAt  InvoiceLineItemsDBFieldNameType
	MetaUpdatedBy  InvoiceLineItemsDBFieldNameType
	MetaDeletedAt  InvoiceLineItemsDBFieldNameType
	MetaDeletedBy  InvoiceLineItemsDBFieldNameType
}

var InvoiceLineItemsDBFieldName = invoiceLineItemsDBFieldName{
	Id:              "id",
	InvoiceId:       "invoice_id",
	LineNo:          "line_no",
	Name:            "name",
	Description:     "description",
	Quantity:        "quantity",
	UnitPrice:       "unit_price",
	DiscountPercent: "discount_percent",
	DiscountAmount:  "discount_amount",
	PpnRate:         "ppn_rate",
	PpnAmount:       "ppn_amount",
	Pph23Rate:       "pph23_rate",
	Pph23Amount:     "pph23_amount",
	Subtotal:        "subtotal",
	TotalAmount:     "total_amount",
	Currency:        "currency",
	Sku:             "sku",
	Category:        "category",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

type InvoiceLineItems struct {
	Id              uuid.UUID        `db:"id"`
	InvoiceId       uuid.UUID        `db:"invoice_id"`
	LineNo          int              `db:"line_no"`
	Name            string           `db:"name"`
	Description     null.String      `db:"description"`
	Quantity        decimal.Decimal  `db:"quantity"`
	UnitPrice       decimal.Decimal  `db:"unit_price"`
	DiscountPercent decimal.Decimal  `db:"discount_percent"`
	DiscountAmount  decimal.Decimal  `db:"discount_amount"`
	PpnRate         decimal.Decimal  `db:"ppn_rate"`
	PpnAmount       decimal.Decimal  `db:"ppn_amount"`
	Pph23Rate       decimal.Decimal  `db:"pph23_rate"`
	Pph23Amount     decimal.Decimal  `db:"pph23_amount"`
	Subtotal        decimal.Decimal  `db:"subtotal"`
	TotalAmount     decimal.Decimal  `db:"total_amount"`
	Currency        string           `db:"currency"`
	Sku             null.String      `db:"sku"`
	Category        null.String      `db:"category"`

	shared.MetaSignature
}

type InvoiceLineItemsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d InvoiceLineItems) ToInvoiceLineItemsPrimaryID() InvoiceLineItemsPrimaryID {
	return InvoiceLineItemsPrimaryID{Id: d.Id}
}

type InvoiceLineItemsList []*InvoiceLineItems

type InvoiceLineItemsFilterResult struct {
	InvoiceLineItems
}

type InvoiceLineItemsFilterResultList []InvoiceLineItemsFilterResult

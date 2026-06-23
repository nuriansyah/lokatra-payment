package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type TaxInvoicesDBFieldNameType string

type taxInvoicesDBFieldName struct {
	Id              TaxInvoicesDBFieldNameType
	InvoiceNo       TaxInvoicesDBFieldNameType
	SourceType      TaxInvoicesDBFieldNameType
	SourceId        TaxInvoicesDBFieldNameType
	MerchantPartyId TaxInvoicesDBFieldNameType
	CustomerPartyId TaxInvoicesDBFieldNameType
	CurrencyCode    TaxInvoicesDBFieldNameType
	TaxableAmount   TaxInvoicesDBFieldNameType
	TaxAmount       TaxInvoicesDBFieldNameType
	TotalAmount     TaxInvoicesDBFieldNameType
	InvoiceStatus   TaxInvoicesDBFieldNameType
	IssuedAt        TaxInvoicesDBFieldNameType
	Metadata        TaxInvoicesDBFieldNameType
	MetaCreatedAt   TaxInvoicesDBFieldNameType
	MetaCreatedBy   TaxInvoicesDBFieldNameType
	MetaUpdatedAt   TaxInvoicesDBFieldNameType
	MetaUpdatedBy   TaxInvoicesDBFieldNameType
	MetaDeletedAt   TaxInvoicesDBFieldNameType
	MetaDeletedBy   TaxInvoicesDBFieldNameType
}

var TaxInvoicesDBFieldName = taxInvoicesDBFieldName{
	Id:              "id",
	InvoiceNo:       "invoice_no",
	SourceType:      "source_type",
	SourceId:        "source_id",
	MerchantPartyId: "merchant_party_id",
	CustomerPartyId: "customer_party_id",
	CurrencyCode:    "currency_code",
	TaxableAmount:   "taxable_amount",
	TaxAmount:       "tax_amount",
	TotalAmount:     "total_amount",
	InvoiceStatus:   "invoice_status",
	IssuedAt:        "issued_at",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewTaxInvoicesDBFieldNameFromStr(field string) (dbField TaxInvoicesDBFieldNameType, found bool) {
	switch field {

	case string(TaxInvoicesDBFieldName.Id):
		return TaxInvoicesDBFieldName.Id, true

	case string(TaxInvoicesDBFieldName.InvoiceNo):
		return TaxInvoicesDBFieldName.InvoiceNo, true

	case string(TaxInvoicesDBFieldName.SourceType):
		return TaxInvoicesDBFieldName.SourceType, true

	case string(TaxInvoicesDBFieldName.SourceId):
		return TaxInvoicesDBFieldName.SourceId, true

	case string(TaxInvoicesDBFieldName.MerchantPartyId):
		return TaxInvoicesDBFieldName.MerchantPartyId, true

	case string(TaxInvoicesDBFieldName.CustomerPartyId):
		return TaxInvoicesDBFieldName.CustomerPartyId, true

	case string(TaxInvoicesDBFieldName.CurrencyCode):
		return TaxInvoicesDBFieldName.CurrencyCode, true

	case string(TaxInvoicesDBFieldName.TaxableAmount):
		return TaxInvoicesDBFieldName.TaxableAmount, true

	case string(TaxInvoicesDBFieldName.TaxAmount):
		return TaxInvoicesDBFieldName.TaxAmount, true

	case string(TaxInvoicesDBFieldName.TotalAmount):
		return TaxInvoicesDBFieldName.TotalAmount, true

	case string(TaxInvoicesDBFieldName.InvoiceStatus):
		return TaxInvoicesDBFieldName.InvoiceStatus, true

	case string(TaxInvoicesDBFieldName.IssuedAt):
		return TaxInvoicesDBFieldName.IssuedAt, true

	case string(TaxInvoicesDBFieldName.Metadata):
		return TaxInvoicesDBFieldName.Metadata, true

	case string(TaxInvoicesDBFieldName.MetaCreatedAt):
		return TaxInvoicesDBFieldName.MetaCreatedAt, true

	case string(TaxInvoicesDBFieldName.MetaCreatedBy):
		return TaxInvoicesDBFieldName.MetaCreatedBy, true

	case string(TaxInvoicesDBFieldName.MetaUpdatedAt):
		return TaxInvoicesDBFieldName.MetaUpdatedAt, true

	case string(TaxInvoicesDBFieldName.MetaUpdatedBy):
		return TaxInvoicesDBFieldName.MetaUpdatedBy, true

	case string(TaxInvoicesDBFieldName.MetaDeletedAt):
		return TaxInvoicesDBFieldName.MetaDeletedAt, true

	case string(TaxInvoicesDBFieldName.MetaDeletedBy):
		return TaxInvoicesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxInvoicesFilterJoins = map[string]JoinSpec{}

var TaxInvoicesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"invoice_no": {
		SourcePath:        "invoice_no",
		DefaultOutputPath: "invoiceNo",
		Column:            "invoice_no",
		SQLAlias:          "invoice_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"customer_party_id": {
		SourcePath:        "customer_party_id",
		DefaultOutputPath: "customerPartyId",
		Column:            "customer_party_id",
		SQLAlias:          "customer_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"taxable_amount": {
		SourcePath:        "taxable_amount",
		DefaultOutputPath: "taxableAmount",
		Column:            "taxable_amount",
		SQLAlias:          "taxable_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_amount": {
		SourcePath:        "tax_amount",
		DefaultOutputPath: "taxAmount",
		Column:            "tax_amount",
		SQLAlias:          "tax_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"total_amount": {
		SourcePath:        "total_amount",
		DefaultOutputPath: "totalAmount",
		Column:            "total_amount",
		SQLAlias:          "total_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"invoice_status": {
		SourcePath:        "invoice_status",
		DefaultOutputPath: "invoiceStatus",
		Column:            "invoice_status",
		SQLAlias:          "invoice_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"issued_at": {
		SourcePath:        "issued_at",
		DefaultOutputPath: "issuedAt",
		Column:            "issued_at",
		SQLAlias:          "issued_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metadata": {
		SourcePath:        "metadata",
		DefaultOutputPath: "metadata",
		Column:            "metadata",
		SQLAlias:          "metadata",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_at": {
		SourcePath:        "meta_created_at",
		DefaultOutputPath: "metaCreatedAt",
		Column:            "meta_created_at",
		SQLAlias:          "meta_created_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_by": {
		SourcePath:        "meta_created_by",
		DefaultOutputPath: "metaCreatedBy",
		Column:            "meta_created_by",
		SQLAlias:          "meta_created_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_at": {
		SourcePath:        "meta_updated_at",
		DefaultOutputPath: "metaUpdatedAt",
		Column:            "meta_updated_at",
		SQLAlias:          "meta_updated_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_by": {
		SourcePath:        "meta_updated_by",
		DefaultOutputPath: "metaUpdatedBy",
		Column:            "meta_updated_by",
		SQLAlias:          "meta_updated_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_at": {
		SourcePath:        "meta_deleted_at",
		DefaultOutputPath: "metaDeletedAt",
		Column:            "meta_deleted_at",
		SQLAlias:          "meta_deleted_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_by": {
		SourcePath:        "meta_deleted_by",
		DefaultOutputPath: "metaDeletedBy",
		Column:            "meta_deleted_by",
		SQLAlias:          "meta_deleted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewTaxInvoicesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxInvoicesFilterFields[field]
	return
}

type TaxInvoicesFilterResult struct {
	TaxInvoices
	FilterCount int `db:"count"`
}

func ValidateTaxInvoicesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxInvoicesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxInvoicesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxInvoicesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxInvoicesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxInvoicesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxInvoicesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxInvoicesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type InvoiceStatus string

const (
	InvoiceStatusDraft  InvoiceStatus = "draft"
	InvoiceStatusIssued InvoiceStatus = "issued"
	InvoiceStatusVoid   InvoiceStatus = "void"
)

type TaxInvoices struct {
	Id              uuid.UUID       `db:"id"`
	InvoiceNo       string          `db:"invoice_no"`
	SourceType      string          `db:"source_type"`
	SourceId        uuid.UUID       `db:"source_id"`
	MerchantPartyId nuuid.NUUID     `db:"merchant_party_id"`
	CustomerPartyId nuuid.NUUID     `db:"customer_party_id"`
	CurrencyCode    string          `db:"currency_code"`
	TaxableAmount   decimal.Decimal `db:"taxable_amount"`
	TaxAmount       decimal.Decimal `db:"tax_amount"`
	TotalAmount     decimal.Decimal `db:"total_amount"`
	InvoiceStatus   InvoiceStatus   `db:"invoice_status"`
	IssuedAt        null.Time       `db:"issued_at"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type TaxInvoicesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d TaxInvoices) ToTaxInvoicesPrimaryID() TaxInvoicesPrimaryID {
	return TaxInvoicesPrimaryID{
		Id: d.Id,
	}
}

type TaxInvoicesList []*TaxInvoices

type TaxInvoicesFilterResultList []*TaxInvoicesFilterResult

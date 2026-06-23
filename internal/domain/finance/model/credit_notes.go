package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type CreditNotesDBFieldNameType string

type creditNotesDBFieldName struct {
	Id            CreditNotesDBFieldNameType
	CreditNoteNo  CreditNotesDBFieldNameType
	TaxInvoiceId  CreditNotesDBFieldNameType
	SourceType    CreditNotesDBFieldNameType
	SourceId      CreditNotesDBFieldNameType
	CurrencyCode  CreditNotesDBFieldNameType
	TaxableAmount CreditNotesDBFieldNameType
	TaxAmount     CreditNotesDBFieldNameType
	TotalAmount   CreditNotesDBFieldNameType
	ReasonCode    CreditNotesDBFieldNameType
	ReasonDetail  CreditNotesDBFieldNameType
	IssuedAt      CreditNotesDBFieldNameType
	Metadata      CreditNotesDBFieldNameType
	MetaCreatedAt CreditNotesDBFieldNameType
	MetaCreatedBy CreditNotesDBFieldNameType
	MetaUpdatedAt CreditNotesDBFieldNameType
	MetaUpdatedBy CreditNotesDBFieldNameType
	MetaDeletedAt CreditNotesDBFieldNameType
	MetaDeletedBy CreditNotesDBFieldNameType
}

var CreditNotesDBFieldName = creditNotesDBFieldName{
	Id:            "id",
	CreditNoteNo:  "credit_note_no",
	TaxInvoiceId:  "tax_invoice_id",
	SourceType:    "source_type",
	SourceId:      "source_id",
	CurrencyCode:  "currency_code",
	TaxableAmount: "taxable_amount",
	TaxAmount:     "tax_amount",
	TotalAmount:   "total_amount",
	ReasonCode:    "reason_code",
	ReasonDetail:  "reason_detail",
	IssuedAt:      "issued_at",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewCreditNotesDBFieldNameFromStr(field string) (dbField CreditNotesDBFieldNameType, found bool) {
	switch field {

	case string(CreditNotesDBFieldName.Id):
		return CreditNotesDBFieldName.Id, true

	case string(CreditNotesDBFieldName.CreditNoteNo):
		return CreditNotesDBFieldName.CreditNoteNo, true

	case string(CreditNotesDBFieldName.TaxInvoiceId):
		return CreditNotesDBFieldName.TaxInvoiceId, true

	case string(CreditNotesDBFieldName.SourceType):
		return CreditNotesDBFieldName.SourceType, true

	case string(CreditNotesDBFieldName.SourceId):
		return CreditNotesDBFieldName.SourceId, true

	case string(CreditNotesDBFieldName.CurrencyCode):
		return CreditNotesDBFieldName.CurrencyCode, true

	case string(CreditNotesDBFieldName.TaxableAmount):
		return CreditNotesDBFieldName.TaxableAmount, true

	case string(CreditNotesDBFieldName.TaxAmount):
		return CreditNotesDBFieldName.TaxAmount, true

	case string(CreditNotesDBFieldName.TotalAmount):
		return CreditNotesDBFieldName.TotalAmount, true

	case string(CreditNotesDBFieldName.ReasonCode):
		return CreditNotesDBFieldName.ReasonCode, true

	case string(CreditNotesDBFieldName.ReasonDetail):
		return CreditNotesDBFieldName.ReasonDetail, true

	case string(CreditNotesDBFieldName.IssuedAt):
		return CreditNotesDBFieldName.IssuedAt, true

	case string(CreditNotesDBFieldName.Metadata):
		return CreditNotesDBFieldName.Metadata, true

	case string(CreditNotesDBFieldName.MetaCreatedAt):
		return CreditNotesDBFieldName.MetaCreatedAt, true

	case string(CreditNotesDBFieldName.MetaCreatedBy):
		return CreditNotesDBFieldName.MetaCreatedBy, true

	case string(CreditNotesDBFieldName.MetaUpdatedAt):
		return CreditNotesDBFieldName.MetaUpdatedAt, true

	case string(CreditNotesDBFieldName.MetaUpdatedBy):
		return CreditNotesDBFieldName.MetaUpdatedBy, true

	case string(CreditNotesDBFieldName.MetaDeletedAt):
		return CreditNotesDBFieldName.MetaDeletedAt, true

	case string(CreditNotesDBFieldName.MetaDeletedBy):
		return CreditNotesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var CreditNotesFilterJoins = map[string]JoinSpec{}

var CreditNotesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"credit_note_no": {
		SourcePath:        "credit_note_no",
		DefaultOutputPath: "creditNoteNo",
		Column:            "credit_note_no",
		SQLAlias:          "credit_note_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_invoice_id": {
		SourcePath:        "tax_invoice_id",
		DefaultOutputPath: "taxInvoiceId",
		Column:            "tax_invoice_id",
		SQLAlias:          "tax_invoice_id",
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
	"reason_code": {
		SourcePath:        "reason_code",
		DefaultOutputPath: "reasonCode",
		Column:            "reason_code",
		SQLAlias:          "reason_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason_detail": {
		SourcePath:        "reason_detail",
		DefaultOutputPath: "reasonDetail",
		Column:            "reason_detail",
		SQLAlias:          "reason_detail",
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

func NewCreditNotesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = CreditNotesFilterFields[field]
	return
}

type CreditNotesFilterResult struct {
	CreditNotes
	FilterCount int `db:"count"`
}

func ValidateCreditNotesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewCreditNotesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewCreditNotesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewCreditNotesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateCreditNotesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateCreditNotesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewCreditNotesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateCreditNotesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type CreditNotes struct {
	Id            uuid.UUID       `db:"id"`
	CreditNoteNo  string          `db:"credit_note_no"`
	TaxInvoiceId  uuid.UUID       `db:"tax_invoice_id"`
	SourceType    string          `db:"source_type"`
	SourceId      uuid.UUID       `db:"source_id"`
	CurrencyCode  string          `db:"currency_code"`
	TaxableAmount decimal.Decimal `db:"taxable_amount"`
	TaxAmount     decimal.Decimal `db:"tax_amount"`
	TotalAmount   decimal.Decimal `db:"total_amount"`
	ReasonCode    string          `db:"reason_code"`
	ReasonDetail  null.String     `db:"reason_detail"`
	IssuedAt      time.Time       `db:"issued_at"`
	Metadata      json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type CreditNotesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d CreditNotes) ToCreditNotesPrimaryID() CreditNotesPrimaryID {
	return CreditNotesPrimaryID{
		Id: d.Id,
	}
}

type CreditNotesList []*CreditNotes

type CreditNotesFilterResultList []*CreditNotesFilterResult

package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type ProviderStatementLinesDBFieldNameType string

type providerStatementLinesDBFieldName struct {
	Id              ProviderStatementLinesDBFieldNameType
	StatementFileId ProviderStatementLinesDBFieldNameType
	LineNo          ProviderStatementLinesDBFieldNameType
	ProviderRef     ProviderStatementLinesDBFieldNameType
	TransactionType ProviderStatementLinesDBFieldNameType
	BookingDate     ProviderStatementLinesDBFieldNameType
	ValueDate       ProviderStatementLinesDBFieldNameType
	CurrencyCode    ProviderStatementLinesDBFieldNameType
	GrossAmount     ProviderStatementLinesDBFieldNameType
	FeeAmount       ProviderStatementLinesDBFieldNameType
	NetAmount       ProviderStatementLinesDBFieldNameType
	RawLine         ProviderStatementLinesDBFieldNameType
	LineHash        ProviderStatementLinesDBFieldNameType
	Metadata        ProviderStatementLinesDBFieldNameType
	MetaCreatedAt   ProviderStatementLinesDBFieldNameType
	MetaCreatedBy   ProviderStatementLinesDBFieldNameType
	MetaUpdatedAt   ProviderStatementLinesDBFieldNameType
	MetaUpdatedBy   ProviderStatementLinesDBFieldNameType
	MetaDeletedAt   ProviderStatementLinesDBFieldNameType
	MetaDeletedBy   ProviderStatementLinesDBFieldNameType
}

var ProviderStatementLinesDBFieldName = providerStatementLinesDBFieldName{
	Id:              "id",
	StatementFileId: "statement_file_id",
	LineNo:          "line_no",
	ProviderRef:     "provider_ref",
	TransactionType: "transaction_type",
	BookingDate:     "booking_date",
	ValueDate:       "value_date",
	CurrencyCode:    "currency_code",
	GrossAmount:     "gross_amount",
	FeeAmount:       "fee_amount",
	NetAmount:       "net_amount",
	RawLine:         "raw_line",
	LineHash:        "line_hash",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewProviderStatementLinesDBFieldNameFromStr(field string) (dbField ProviderStatementLinesDBFieldNameType, found bool) {
	switch field {

	case string(ProviderStatementLinesDBFieldName.Id):
		return ProviderStatementLinesDBFieldName.Id, true

	case string(ProviderStatementLinesDBFieldName.StatementFileId):
		return ProviderStatementLinesDBFieldName.StatementFileId, true

	case string(ProviderStatementLinesDBFieldName.LineNo):
		return ProviderStatementLinesDBFieldName.LineNo, true

	case string(ProviderStatementLinesDBFieldName.ProviderRef):
		return ProviderStatementLinesDBFieldName.ProviderRef, true

	case string(ProviderStatementLinesDBFieldName.TransactionType):
		return ProviderStatementLinesDBFieldName.TransactionType, true

	case string(ProviderStatementLinesDBFieldName.BookingDate):
		return ProviderStatementLinesDBFieldName.BookingDate, true

	case string(ProviderStatementLinesDBFieldName.ValueDate):
		return ProviderStatementLinesDBFieldName.ValueDate, true

	case string(ProviderStatementLinesDBFieldName.CurrencyCode):
		return ProviderStatementLinesDBFieldName.CurrencyCode, true

	case string(ProviderStatementLinesDBFieldName.GrossAmount):
		return ProviderStatementLinesDBFieldName.GrossAmount, true

	case string(ProviderStatementLinesDBFieldName.FeeAmount):
		return ProviderStatementLinesDBFieldName.FeeAmount, true

	case string(ProviderStatementLinesDBFieldName.NetAmount):
		return ProviderStatementLinesDBFieldName.NetAmount, true

	case string(ProviderStatementLinesDBFieldName.RawLine):
		return ProviderStatementLinesDBFieldName.RawLine, true

	case string(ProviderStatementLinesDBFieldName.LineHash):
		return ProviderStatementLinesDBFieldName.LineHash, true

	case string(ProviderStatementLinesDBFieldName.Metadata):
		return ProviderStatementLinesDBFieldName.Metadata, true

	case string(ProviderStatementLinesDBFieldName.MetaCreatedAt):
		return ProviderStatementLinesDBFieldName.MetaCreatedAt, true

	case string(ProviderStatementLinesDBFieldName.MetaCreatedBy):
		return ProviderStatementLinesDBFieldName.MetaCreatedBy, true

	case string(ProviderStatementLinesDBFieldName.MetaUpdatedAt):
		return ProviderStatementLinesDBFieldName.MetaUpdatedAt, true

	case string(ProviderStatementLinesDBFieldName.MetaUpdatedBy):
		return ProviderStatementLinesDBFieldName.MetaUpdatedBy, true

	case string(ProviderStatementLinesDBFieldName.MetaDeletedAt):
		return ProviderStatementLinesDBFieldName.MetaDeletedAt, true

	case string(ProviderStatementLinesDBFieldName.MetaDeletedBy):
		return ProviderStatementLinesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderStatementLinesFilterJoins = map[string]JoinSpec{}

var ProviderStatementLinesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_file_id": {
		SourcePath:        "statement_file_id",
		DefaultOutputPath: "statementFileId",
		Column:            "statement_file_id",
		SQLAlias:          "statement_file_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"line_no": {
		SourcePath:        "line_no",
		DefaultOutputPath: "lineNo",
		Column:            "line_no",
		SQLAlias:          "line_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_ref": {
		SourcePath:        "provider_ref",
		DefaultOutputPath: "providerRef",
		Column:            "provider_ref",
		SQLAlias:          "provider_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"transaction_type": {
		SourcePath:        "transaction_type",
		DefaultOutputPath: "transactionType",
		Column:            "transaction_type",
		SQLAlias:          "transaction_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"booking_date": {
		SourcePath:        "booking_date",
		DefaultOutputPath: "bookingDate",
		Column:            "booking_date",
		SQLAlias:          "booking_date",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"value_date": {
		SourcePath:        "value_date",
		DefaultOutputPath: "valueDate",
		Column:            "value_date",
		SQLAlias:          "value_date",
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
	"gross_amount": {
		SourcePath:        "gross_amount",
		DefaultOutputPath: "grossAmount",
		Column:            "gross_amount",
		SQLAlias:          "gross_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_amount": {
		SourcePath:        "fee_amount",
		DefaultOutputPath: "feeAmount",
		Column:            "fee_amount",
		SQLAlias:          "fee_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"net_amount": {
		SourcePath:        "net_amount",
		DefaultOutputPath: "netAmount",
		Column:            "net_amount",
		SQLAlias:          "net_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"raw_line": {
		SourcePath:        "raw_line",
		DefaultOutputPath: "rawLine",
		Column:            "raw_line",
		SQLAlias:          "raw_line",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"line_hash": {
		SourcePath:        "line_hash",
		DefaultOutputPath: "lineHash",
		Column:            "line_hash",
		SQLAlias:          "line_hash",
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

func NewProviderStatementLinesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderStatementLinesFilterFields[field]
	return
}

type ProviderStatementLinesFilterResult struct {
	ProviderStatementLines
	FilterCount int `db:"count"`
}

func ValidateProviderStatementLinesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderStatementLinesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderStatementLinesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderStatementLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderStatementLinesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderStatementLinesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderStatementLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderStatementLinesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderStatementLines struct {
	Id              uuid.UUID       `db:"id"`
	StatementFileId uuid.UUID       `db:"statement_file_id"`
	LineNo          int             `db:"line_no"`
	ProviderRef     null.String     `db:"provider_ref"`
	TransactionType string          `db:"transaction_type"`
	BookingDate     null.Time       `db:"booking_date"`
	ValueDate       null.Time       `db:"value_date"`
	CurrencyCode    string          `db:"currency_code"`
	GrossAmount     decimal.Decimal `db:"gross_amount"`
	FeeAmount       decimal.Decimal `db:"fee_amount"`
	NetAmount       decimal.Decimal `db:"net_amount"`
	RawLine         json.RawMessage `db:"raw_line"`
	LineHash        null.String     `db:"line_hash"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ProviderStatementLinesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderStatementLines) ToProviderStatementLinesPrimaryID() ProviderStatementLinesPrimaryID {
	return ProviderStatementLinesPrimaryID{
		Id: d.Id,
	}
}

type ProviderStatementLinesList []*ProviderStatementLines

type ProviderStatementLinesFilterResultList []*ProviderStatementLinesFilterResult

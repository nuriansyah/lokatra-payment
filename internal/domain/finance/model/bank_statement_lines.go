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

type BankStatementLinesDBFieldNameType string

type bankStatementLinesDBFieldName struct {
	Id              BankStatementLinesDBFieldNameType
	StatementFileId BankStatementLinesDBFieldNameType
	LineNo          BankStatementLinesDBFieldNameType
	BankRef         BankStatementLinesDBFieldNameType
	TransactionType BankStatementLinesDBFieldNameType
	BookingDate     BankStatementLinesDBFieldNameType
	ValueDate       BankStatementLinesDBFieldNameType
	CurrencyCode    BankStatementLinesDBFieldNameType
	DebitAmount     BankStatementLinesDBFieldNameType
	CreditAmount    BankStatementLinesDBFieldNameType
	NetAmount       BankStatementLinesDBFieldNameType
	RawLine         BankStatementLinesDBFieldNameType
	LineHash        BankStatementLinesDBFieldNameType
	Metadata        BankStatementLinesDBFieldNameType
	MetaCreatedAt   BankStatementLinesDBFieldNameType
	MetaCreatedBy   BankStatementLinesDBFieldNameType
	MetaUpdatedAt   BankStatementLinesDBFieldNameType
	MetaUpdatedBy   BankStatementLinesDBFieldNameType
	MetaDeletedAt   BankStatementLinesDBFieldNameType
	MetaDeletedBy   BankStatementLinesDBFieldNameType
}

var BankStatementLinesDBFieldName = bankStatementLinesDBFieldName{
	Id:              "id",
	StatementFileId: "statement_file_id",
	LineNo:          "line_no",
	BankRef:         "bank_ref",
	TransactionType: "transaction_type",
	BookingDate:     "booking_date",
	ValueDate:       "value_date",
	CurrencyCode:    "currency_code",
	DebitAmount:     "debit_amount",
	CreditAmount:    "credit_amount",
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

func NewBankStatementLinesDBFieldNameFromStr(field string) (dbField BankStatementLinesDBFieldNameType, found bool) {
	switch field {

	case string(BankStatementLinesDBFieldName.Id):
		return BankStatementLinesDBFieldName.Id, true

	case string(BankStatementLinesDBFieldName.StatementFileId):
		return BankStatementLinesDBFieldName.StatementFileId, true

	case string(BankStatementLinesDBFieldName.LineNo):
		return BankStatementLinesDBFieldName.LineNo, true

	case string(BankStatementLinesDBFieldName.BankRef):
		return BankStatementLinesDBFieldName.BankRef, true

	case string(BankStatementLinesDBFieldName.TransactionType):
		return BankStatementLinesDBFieldName.TransactionType, true

	case string(BankStatementLinesDBFieldName.BookingDate):
		return BankStatementLinesDBFieldName.BookingDate, true

	case string(BankStatementLinesDBFieldName.ValueDate):
		return BankStatementLinesDBFieldName.ValueDate, true

	case string(BankStatementLinesDBFieldName.CurrencyCode):
		return BankStatementLinesDBFieldName.CurrencyCode, true

	case string(BankStatementLinesDBFieldName.DebitAmount):
		return BankStatementLinesDBFieldName.DebitAmount, true

	case string(BankStatementLinesDBFieldName.CreditAmount):
		return BankStatementLinesDBFieldName.CreditAmount, true

	case string(BankStatementLinesDBFieldName.NetAmount):
		return BankStatementLinesDBFieldName.NetAmount, true

	case string(BankStatementLinesDBFieldName.RawLine):
		return BankStatementLinesDBFieldName.RawLine, true

	case string(BankStatementLinesDBFieldName.LineHash):
		return BankStatementLinesDBFieldName.LineHash, true

	case string(BankStatementLinesDBFieldName.Metadata):
		return BankStatementLinesDBFieldName.Metadata, true

	case string(BankStatementLinesDBFieldName.MetaCreatedAt):
		return BankStatementLinesDBFieldName.MetaCreatedAt, true

	case string(BankStatementLinesDBFieldName.MetaCreatedBy):
		return BankStatementLinesDBFieldName.MetaCreatedBy, true

	case string(BankStatementLinesDBFieldName.MetaUpdatedAt):
		return BankStatementLinesDBFieldName.MetaUpdatedAt, true

	case string(BankStatementLinesDBFieldName.MetaUpdatedBy):
		return BankStatementLinesDBFieldName.MetaUpdatedBy, true

	case string(BankStatementLinesDBFieldName.MetaDeletedAt):
		return BankStatementLinesDBFieldName.MetaDeletedAt, true

	case string(BankStatementLinesDBFieldName.MetaDeletedBy):
		return BankStatementLinesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var BankStatementLinesFilterJoins = map[string]JoinSpec{}

var BankStatementLinesFilterFields = map[string]FilterFieldSpec{
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
	"bank_ref": {
		SourcePath:        "bank_ref",
		DefaultOutputPath: "bankRef",
		Column:            "bank_ref",
		SQLAlias:          "bank_ref",
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
	"debit_amount": {
		SourcePath:        "debit_amount",
		DefaultOutputPath: "debitAmount",
		Column:            "debit_amount",
		SQLAlias:          "debit_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"credit_amount": {
		SourcePath:        "credit_amount",
		DefaultOutputPath: "creditAmount",
		Column:            "credit_amount",
		SQLAlias:          "credit_amount",
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

func NewBankStatementLinesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = BankStatementLinesFilterFields[field]
	return
}

type BankStatementLinesFilterResult struct {
	BankStatementLines
	FilterCount int `db:"count"`
}

func ValidateBankStatementLinesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewBankStatementLinesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewBankStatementLinesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewBankStatementLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateBankStatementLinesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateBankStatementLinesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewBankStatementLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateBankStatementLinesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type BankStatementLines struct {
	Id              uuid.UUID       `db:"id"`
	StatementFileId uuid.UUID       `db:"statement_file_id"`
	LineNo          int             `db:"line_no"`
	BankRef         null.String     `db:"bank_ref"`
	TransactionType string          `db:"transaction_type"`
	BookingDate     null.Time       `db:"booking_date"`
	ValueDate       null.Time       `db:"value_date"`
	CurrencyCode    string          `db:"currency_code"`
	DebitAmount     decimal.Decimal `db:"debit_amount"`
	CreditAmount    decimal.Decimal `db:"credit_amount"`
	NetAmount       decimal.Decimal `db:"net_amount"`
	RawLine         json.RawMessage `db:"raw_line"`
	LineHash        null.String     `db:"line_hash"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type BankStatementLinesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d BankStatementLines) ToBankStatementLinesPrimaryID() BankStatementLinesPrimaryID {
	return BankStatementLinesPrimaryID{
		Id: d.Id,
	}
}

type BankStatementLinesList []*BankStatementLines

type BankStatementLinesFilterResultList []*BankStatementLinesFilterResult

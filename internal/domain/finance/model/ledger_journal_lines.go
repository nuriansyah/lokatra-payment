package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type LedgerJournalLinesDBFieldNameType string

type ledgerJournalLinesDBFieldName struct {
	Id               LedgerJournalLinesDBFieldNameType
	JournalEntryId   LedgerJournalLinesDBFieldNameType
	LineNo           LedgerJournalLinesDBFieldNameType
	AccountId        LedgerJournalLinesDBFieldNameType
	LineSide         LedgerJournalLinesDBFieldNameType
	Amount           LedgerJournalLinesDBFieldNameType
	CurrencyCode     LedgerJournalLinesDBFieldNameType
	FxRateLockId     LedgerJournalLinesDBFieldNameType
	AmountReporting  LedgerJournalLinesDBFieldNameType
	ReferencePartyId LedgerJournalLinesDBFieldNameType
	Dimensions       LedgerJournalLinesDBFieldNameType
	Metadata         LedgerJournalLinesDBFieldNameType
	MetaCreatedAt    LedgerJournalLinesDBFieldNameType
	MetaCreatedBy    LedgerJournalLinesDBFieldNameType
	MetaUpdatedAt    LedgerJournalLinesDBFieldNameType
	MetaUpdatedBy    LedgerJournalLinesDBFieldNameType
	MetaDeletedAt    LedgerJournalLinesDBFieldNameType
	MetaDeletedBy    LedgerJournalLinesDBFieldNameType
}

var LedgerJournalLinesDBFieldName = ledgerJournalLinesDBFieldName{
	Id:               "id",
	JournalEntryId:   "journal_entry_id",
	LineNo:           "line_no",
	AccountId:        "account_id",
	LineSide:         "line_side",
	Amount:           "amount",
	CurrencyCode:     "currency_code",
	FxRateLockId:     "fx_rate_lock_id",
	AmountReporting:  "amount_reporting",
	ReferencePartyId: "reference_party_id",
	Dimensions:       "dimensions",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewLedgerJournalLinesDBFieldNameFromStr(field string) (dbField LedgerJournalLinesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerJournalLinesDBFieldName.Id):
		return LedgerJournalLinesDBFieldName.Id, true

	case string(LedgerJournalLinesDBFieldName.JournalEntryId):
		return LedgerJournalLinesDBFieldName.JournalEntryId, true

	case string(LedgerJournalLinesDBFieldName.LineNo):
		return LedgerJournalLinesDBFieldName.LineNo, true

	case string(LedgerJournalLinesDBFieldName.AccountId):
		return LedgerJournalLinesDBFieldName.AccountId, true

	case string(LedgerJournalLinesDBFieldName.LineSide):
		return LedgerJournalLinesDBFieldName.LineSide, true

	case string(LedgerJournalLinesDBFieldName.Amount):
		return LedgerJournalLinesDBFieldName.Amount, true

	case string(LedgerJournalLinesDBFieldName.CurrencyCode):
		return LedgerJournalLinesDBFieldName.CurrencyCode, true

	case string(LedgerJournalLinesDBFieldName.FxRateLockId):
		return LedgerJournalLinesDBFieldName.FxRateLockId, true

	case string(LedgerJournalLinesDBFieldName.AmountReporting):
		return LedgerJournalLinesDBFieldName.AmountReporting, true

	case string(LedgerJournalLinesDBFieldName.ReferencePartyId):
		return LedgerJournalLinesDBFieldName.ReferencePartyId, true

	case string(LedgerJournalLinesDBFieldName.Dimensions):
		return LedgerJournalLinesDBFieldName.Dimensions, true

	case string(LedgerJournalLinesDBFieldName.Metadata):
		return LedgerJournalLinesDBFieldName.Metadata, true

	case string(LedgerJournalLinesDBFieldName.MetaCreatedAt):
		return LedgerJournalLinesDBFieldName.MetaCreatedAt, true

	case string(LedgerJournalLinesDBFieldName.MetaCreatedBy):
		return LedgerJournalLinesDBFieldName.MetaCreatedBy, true

	case string(LedgerJournalLinesDBFieldName.MetaUpdatedAt):
		return LedgerJournalLinesDBFieldName.MetaUpdatedAt, true

	case string(LedgerJournalLinesDBFieldName.MetaUpdatedBy):
		return LedgerJournalLinesDBFieldName.MetaUpdatedBy, true

	case string(LedgerJournalLinesDBFieldName.MetaDeletedAt):
		return LedgerJournalLinesDBFieldName.MetaDeletedAt, true

	case string(LedgerJournalLinesDBFieldName.MetaDeletedBy):
		return LedgerJournalLinesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerJournalLinesFilterJoins = map[string]JoinSpec{}

var LedgerJournalLinesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"journal_entry_id": {
		SourcePath:        "journal_entry_id",
		DefaultOutputPath: "journalEntryId",
		Column:            "journal_entry_id",
		SQLAlias:          "journal_entry_id",
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
	"account_id": {
		SourcePath:        "account_id",
		DefaultOutputPath: "accountId",
		Column:            "account_id",
		SQLAlias:          "account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"line_side": {
		SourcePath:        "line_side",
		DefaultOutputPath: "lineSide",
		Column:            "line_side",
		SQLAlias:          "line_side",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
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
	"fx_rate_lock_id": {
		SourcePath:        "fx_rate_lock_id",
		DefaultOutputPath: "fxRateLockId",
		Column:            "fx_rate_lock_id",
		SQLAlias:          "fx_rate_lock_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount_reporting": {
		SourcePath:        "amount_reporting",
		DefaultOutputPath: "amountReporting",
		Column:            "amount_reporting",
		SQLAlias:          "amount_reporting",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reference_party_id": {
		SourcePath:        "reference_party_id",
		DefaultOutputPath: "referencePartyId",
		Column:            "reference_party_id",
		SQLAlias:          "reference_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dimensions": {
		SourcePath:        "dimensions",
		DefaultOutputPath: "dimensions",
		Column:            "dimensions",
		SQLAlias:          "dimensions",
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

func NewLedgerJournalLinesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerJournalLinesFilterFields[field]
	return
}

type LedgerJournalLinesFilterResult struct {
	LedgerJournalLines
	FilterCount int `db:"count"`
}

func ValidateLedgerJournalLinesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerJournalLinesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerJournalLinesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerJournalLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerJournalLinesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerJournalLinesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerJournalLinesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerJournalLinesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type LineSide string

const (
	LineSideDebit  LineSide = "debit"
	LineSideCredit LineSide = "credit"
)

type LedgerJournalLines struct {
	Id               uuid.UUID           `db:"id"`
	JournalEntryId   uuid.UUID           `db:"journal_entry_id"`
	LineNo           int                 `db:"line_no"`
	AccountId        uuid.UUID           `db:"account_id"`
	LineSide         LineSide            `db:"line_side"`
	Amount           decimal.Decimal     `db:"amount"`
	CurrencyCode     string              `db:"currency_code"`
	FxRateLockId     nuuid.NUUID         `db:"fx_rate_lock_id"`
	AmountReporting  decimal.NullDecimal `db:"amount_reporting"`
	ReferencePartyId nuuid.NUUID         `db:"reference_party_id"`
	Dimensions       json.RawMessage     `db:"dimensions"`
	Metadata         json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type LedgerJournalLinesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerJournalLines) ToLedgerJournalLinesPrimaryID() LedgerJournalLinesPrimaryID {
	return LedgerJournalLinesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalLinesList []*LedgerJournalLines

type LedgerJournalLinesFilterResultList []*LedgerJournalLinesFilterResult

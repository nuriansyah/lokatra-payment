package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerBooksDBFieldNameType string

type ledgerBooksDBFieldName struct {
	Id            LedgerBooksDBFieldNameType
	BookCode      LedgerBooksDBFieldNameType
	BookTypeCode  LedgerBooksDBFieldNameType
	OwnerPartyId  LedgerBooksDBFieldNameType
	CurrencyCode  LedgerBooksDBFieldNameType
	BookStatus    LedgerBooksDBFieldNameType
	CloseCutoffTz LedgerBooksDBFieldNameType
	Metadata      LedgerBooksDBFieldNameType
	MetaCreatedAt LedgerBooksDBFieldNameType
	MetaCreatedBy LedgerBooksDBFieldNameType
	MetaUpdatedAt LedgerBooksDBFieldNameType
	MetaUpdatedBy LedgerBooksDBFieldNameType
	MetaDeletedAt LedgerBooksDBFieldNameType
	MetaDeletedBy LedgerBooksDBFieldNameType
}

var LedgerBooksDBFieldName = ledgerBooksDBFieldName{
	Id:            "id",
	BookCode:      "book_code",
	BookTypeCode:  "book_type_code",
	OwnerPartyId:  "owner_party_id",
	CurrencyCode:  "currency_code",
	BookStatus:    "book_status",
	CloseCutoffTz: "close_cutoff_tz",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewLedgerBooksDBFieldNameFromStr(field string) (dbField LedgerBooksDBFieldNameType, found bool) {
	switch field {

	case string(LedgerBooksDBFieldName.Id):
		return LedgerBooksDBFieldName.Id, true

	case string(LedgerBooksDBFieldName.BookCode):
		return LedgerBooksDBFieldName.BookCode, true

	case string(LedgerBooksDBFieldName.BookTypeCode):
		return LedgerBooksDBFieldName.BookTypeCode, true

	case string(LedgerBooksDBFieldName.OwnerPartyId):
		return LedgerBooksDBFieldName.OwnerPartyId, true

	case string(LedgerBooksDBFieldName.CurrencyCode):
		return LedgerBooksDBFieldName.CurrencyCode, true

	case string(LedgerBooksDBFieldName.BookStatus):
		return LedgerBooksDBFieldName.BookStatus, true

	case string(LedgerBooksDBFieldName.CloseCutoffTz):
		return LedgerBooksDBFieldName.CloseCutoffTz, true

	case string(LedgerBooksDBFieldName.Metadata):
		return LedgerBooksDBFieldName.Metadata, true

	case string(LedgerBooksDBFieldName.MetaCreatedAt):
		return LedgerBooksDBFieldName.MetaCreatedAt, true

	case string(LedgerBooksDBFieldName.MetaCreatedBy):
		return LedgerBooksDBFieldName.MetaCreatedBy, true

	case string(LedgerBooksDBFieldName.MetaUpdatedAt):
		return LedgerBooksDBFieldName.MetaUpdatedAt, true

	case string(LedgerBooksDBFieldName.MetaUpdatedBy):
		return LedgerBooksDBFieldName.MetaUpdatedBy, true

	case string(LedgerBooksDBFieldName.MetaDeletedAt):
		return LedgerBooksDBFieldName.MetaDeletedAt, true

	case string(LedgerBooksDBFieldName.MetaDeletedBy):
		return LedgerBooksDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerBooksFilterJoins = map[string]JoinSpec{}

var LedgerBooksFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_code": {
		SourcePath:        "book_code",
		DefaultOutputPath: "bookCode",
		Column:            "book_code",
		SQLAlias:          "book_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_type_code": {
		SourcePath:        "book_type_code",
		DefaultOutputPath: "bookTypeCode",
		Column:            "book_type_code",
		SQLAlias:          "book_type_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"owner_party_id": {
		SourcePath:        "owner_party_id",
		DefaultOutputPath: "ownerPartyId",
		Column:            "owner_party_id",
		SQLAlias:          "owner_party_id",
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
	"book_status": {
		SourcePath:        "book_status",
		DefaultOutputPath: "bookStatus",
		Column:            "book_status",
		SQLAlias:          "book_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"close_cutoff_tz": {
		SourcePath:        "close_cutoff_tz",
		DefaultOutputPath: "closeCutoffTz",
		Column:            "close_cutoff_tz",
		SQLAlias:          "close_cutoff_tz",
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

func NewLedgerBooksFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerBooksFilterFields[field]
	return
}

type LedgerBooksFilterResult struct {
	LedgerBooks
	FilterCount int `db:"count"`
}

func ValidateLedgerBooksFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerBooksFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerBooksFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerBooksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerBooksFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerBooksFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerBooksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerBooksFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type BookStatus string

const (
	BookStatusActive BookStatus = "active"
	BookStatusLocked BookStatus = "locked"
	BookStatusClosed BookStatus = "closed"
)

type LedgerBooks struct {
	Id            uuid.UUID       `db:"id"`
	BookCode      string          `db:"book_code"`
	BookTypeCode  string          `db:"book_type_code"`
	OwnerPartyId  uuid.UUID       `db:"owner_party_id"`
	CurrencyCode  string          `db:"currency_code"`
	BookStatus    BookStatus      `db:"book_status"`
	CloseCutoffTz string          `db:"close_cutoff_tz"`
	Metadata      json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerBooksPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerBooks) ToLedgerBooksPrimaryID() LedgerBooksPrimaryID {
	return LedgerBooksPrimaryID{
		Id: d.Id,
	}
}

type LedgerBooksList []*LedgerBooks

type LedgerBooksFilterResultList []*LedgerBooksFilterResult

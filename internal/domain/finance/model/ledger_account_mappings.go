package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerAccountMappingsDBFieldNameType string

type ledgerAccountMappingsDBFieldName struct {
	Id            LedgerAccountMappingsDBFieldNameType
	MappingCode   LedgerAccountMappingsDBFieldNameType
	BookId        LedgerAccountMappingsDBFieldNameType
	SourceType    LedgerAccountMappingsDBFieldNameType
	SourceSubtype LedgerAccountMappingsDBFieldNameType
	AccountId     LedgerAccountMappingsDBFieldNameType
	Priority      LedgerAccountMappingsDBFieldNameType
	IsActive      LedgerAccountMappingsDBFieldNameType
	Conditions    LedgerAccountMappingsDBFieldNameType
	Metadata      LedgerAccountMappingsDBFieldNameType
	MetaCreatedAt LedgerAccountMappingsDBFieldNameType
	MetaCreatedBy LedgerAccountMappingsDBFieldNameType
	MetaUpdatedAt LedgerAccountMappingsDBFieldNameType
	MetaUpdatedBy LedgerAccountMappingsDBFieldNameType
	MetaDeletedAt LedgerAccountMappingsDBFieldNameType
	MetaDeletedBy LedgerAccountMappingsDBFieldNameType
}

var LedgerAccountMappingsDBFieldName = ledgerAccountMappingsDBFieldName{
	Id:            "id",
	MappingCode:   "mapping_code",
	BookId:        "book_id",
	SourceType:    "source_type",
	SourceSubtype: "source_subtype",
	AccountId:     "account_id",
	Priority:      "priority",
	IsActive:      "is_active",
	Conditions:    "conditions",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewLedgerAccountMappingsDBFieldNameFromStr(field string) (dbField LedgerAccountMappingsDBFieldNameType, found bool) {
	switch field {

	case string(LedgerAccountMappingsDBFieldName.Id):
		return LedgerAccountMappingsDBFieldName.Id, true

	case string(LedgerAccountMappingsDBFieldName.MappingCode):
		return LedgerAccountMappingsDBFieldName.MappingCode, true

	case string(LedgerAccountMappingsDBFieldName.BookId):
		return LedgerAccountMappingsDBFieldName.BookId, true

	case string(LedgerAccountMappingsDBFieldName.SourceType):
		return LedgerAccountMappingsDBFieldName.SourceType, true

	case string(LedgerAccountMappingsDBFieldName.SourceSubtype):
		return LedgerAccountMappingsDBFieldName.SourceSubtype, true

	case string(LedgerAccountMappingsDBFieldName.AccountId):
		return LedgerAccountMappingsDBFieldName.AccountId, true

	case string(LedgerAccountMappingsDBFieldName.Priority):
		return LedgerAccountMappingsDBFieldName.Priority, true

	case string(LedgerAccountMappingsDBFieldName.IsActive):
		return LedgerAccountMappingsDBFieldName.IsActive, true

	case string(LedgerAccountMappingsDBFieldName.Conditions):
		return LedgerAccountMappingsDBFieldName.Conditions, true

	case string(LedgerAccountMappingsDBFieldName.Metadata):
		return LedgerAccountMappingsDBFieldName.Metadata, true

	case string(LedgerAccountMappingsDBFieldName.MetaCreatedAt):
		return LedgerAccountMappingsDBFieldName.MetaCreatedAt, true

	case string(LedgerAccountMappingsDBFieldName.MetaCreatedBy):
		return LedgerAccountMappingsDBFieldName.MetaCreatedBy, true

	case string(LedgerAccountMappingsDBFieldName.MetaUpdatedAt):
		return LedgerAccountMappingsDBFieldName.MetaUpdatedAt, true

	case string(LedgerAccountMappingsDBFieldName.MetaUpdatedBy):
		return LedgerAccountMappingsDBFieldName.MetaUpdatedBy, true

	case string(LedgerAccountMappingsDBFieldName.MetaDeletedAt):
		return LedgerAccountMappingsDBFieldName.MetaDeletedAt, true

	case string(LedgerAccountMappingsDBFieldName.MetaDeletedBy):
		return LedgerAccountMappingsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerAccountMappingsFilterJoins = map[string]JoinSpec{}

var LedgerAccountMappingsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"mapping_code": {
		SourcePath:        "mapping_code",
		DefaultOutputPath: "mappingCode",
		Column:            "mapping_code",
		SQLAlias:          "mapping_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_id": {
		SourcePath:        "book_id",
		DefaultOutputPath: "bookId",
		Column:            "book_id",
		SQLAlias:          "book_id",
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
	"source_subtype": {
		SourcePath:        "source_subtype",
		DefaultOutputPath: "sourceSubtype",
		Column:            "source_subtype",
		SQLAlias:          "source_subtype",
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
	"priority": {
		SourcePath:        "priority",
		DefaultOutputPath: "priority",
		Column:            "priority",
		SQLAlias:          "priority",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_active": {
		SourcePath:        "is_active",
		DefaultOutputPath: "isActive",
		Column:            "is_active",
		SQLAlias:          "is_active",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"conditions": {
		SourcePath:        "conditions",
		DefaultOutputPath: "conditions",
		Column:            "conditions",
		SQLAlias:          "conditions",
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

func NewLedgerAccountMappingsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerAccountMappingsFilterFields[field]
	return
}

type LedgerAccountMappingsFilterResult struct {
	LedgerAccountMappings
	FilterCount int `db:"count"`
}

func ValidateLedgerAccountMappingsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerAccountMappingsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerAccountMappingsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerAccountMappingsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerAccountMappingsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerAccountMappingsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerAccountMappingsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerAccountMappingsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type LedgerAccountMappings struct {
	Id            uuid.UUID       `db:"id"`
	MappingCode   string          `db:"mapping_code"`
	BookId        uuid.UUID       `db:"book_id"`
	SourceType    string          `db:"source_type"`
	SourceSubtype null.String     `db:"source_subtype"`
	AccountId     uuid.UUID       `db:"account_id"`
	Priority      int             `db:"priority"`
	IsActive      bool            `db:"is_active"`
	Conditions    json.RawMessage `db:"conditions"`
	Metadata      json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerAccountMappingsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerAccountMappings) ToLedgerAccountMappingsPrimaryID() LedgerAccountMappingsPrimaryID {
	return LedgerAccountMappingsPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountMappingsList []*LedgerAccountMappings

type LedgerAccountMappingsFilterResultList []*LedgerAccountMappingsFilterResult

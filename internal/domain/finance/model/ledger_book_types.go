package model

import (
	"encoding/json"
	"fmt"

	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerBookTypesDBFieldNameType string

type ledgerBookTypesDBFieldName struct {
	Code          LedgerBookTypesDBFieldNameType
	Description   LedgerBookTypesDBFieldNameType
	Metadata      LedgerBookTypesDBFieldNameType
	MetaCreatedAt LedgerBookTypesDBFieldNameType
	MetaCreatedBy LedgerBookTypesDBFieldNameType
	MetaUpdatedAt LedgerBookTypesDBFieldNameType
	MetaUpdatedBy LedgerBookTypesDBFieldNameType
	MetaDeletedAt LedgerBookTypesDBFieldNameType
	MetaDeletedBy LedgerBookTypesDBFieldNameType
}

var LedgerBookTypesDBFieldName = ledgerBookTypesDBFieldName{
	Code:          "code",
	Description:   "description",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewLedgerBookTypesDBFieldNameFromStr(field string) (dbField LedgerBookTypesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerBookTypesDBFieldName.Code):
		return LedgerBookTypesDBFieldName.Code, true

	case string(LedgerBookTypesDBFieldName.Description):
		return LedgerBookTypesDBFieldName.Description, true

	case string(LedgerBookTypesDBFieldName.Metadata):
		return LedgerBookTypesDBFieldName.Metadata, true

	case string(LedgerBookTypesDBFieldName.MetaCreatedAt):
		return LedgerBookTypesDBFieldName.MetaCreatedAt, true

	case string(LedgerBookTypesDBFieldName.MetaCreatedBy):
		return LedgerBookTypesDBFieldName.MetaCreatedBy, true

	case string(LedgerBookTypesDBFieldName.MetaUpdatedAt):
		return LedgerBookTypesDBFieldName.MetaUpdatedAt, true

	case string(LedgerBookTypesDBFieldName.MetaUpdatedBy):
		return LedgerBookTypesDBFieldName.MetaUpdatedBy, true

	case string(LedgerBookTypesDBFieldName.MetaDeletedAt):
		return LedgerBookTypesDBFieldName.MetaDeletedAt, true

	case string(LedgerBookTypesDBFieldName.MetaDeletedBy):
		return LedgerBookTypesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerBookTypesFilterJoins = map[string]JoinSpec{}

var LedgerBookTypesFilterFields = map[string]FilterFieldSpec{
	"code": {
		SourcePath:        "code",
		DefaultOutputPath: "code",
		Column:            "code",
		SQLAlias:          "code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"description": {
		SourcePath:        "description",
		DefaultOutputPath: "description",
		Column:            "description",
		SQLAlias:          "description",
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

func NewLedgerBookTypesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerBookTypesFilterFields[field]
	return
}

type LedgerBookTypesFilterResult struct {
	LedgerBookTypes
	FilterCount int `db:"count"`
}

func ValidateLedgerBookTypesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerBookTypesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerBookTypesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerBookTypesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerBookTypesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerBookTypesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerBookTypesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerBookTypesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type LedgerBookTypes struct {
	Code        string          `db:"code"`
	Description string          `db:"description"`
	Metadata    json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerBookTypesPrimaryID struct {
	Code string `db:"code"`
}

func (d LedgerBookTypes) ToLedgerBookTypesPrimaryID() LedgerBookTypesPrimaryID {
	return LedgerBookTypesPrimaryID{
		Code: d.Code,
	}
}

type LedgerBookTypesList []*LedgerBookTypes

type LedgerBookTypesFilterResultList []*LedgerBookTypesFilterResult

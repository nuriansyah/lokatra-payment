package model

import (
	"encoding/json"
	"fmt"

	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceCurrenciesDBFieldNameType string

type financeCurrenciesDBFieldName struct {
	Code          FinanceCurrenciesDBFieldNameType
	DecimalCode   FinanceCurrenciesDBFieldNameType
	Exponent      FinanceCurrenciesDBFieldNameType
	IsActive      FinanceCurrenciesDBFieldNameType
	Metadata      FinanceCurrenciesDBFieldNameType
	MetaCreatedAt FinanceCurrenciesDBFieldNameType
	MetaCreatedBy FinanceCurrenciesDBFieldNameType
	MetaUpdatedAt FinanceCurrenciesDBFieldNameType
	MetaUpdatedBy FinanceCurrenciesDBFieldNameType
	MetaDeletedAt FinanceCurrenciesDBFieldNameType
	MetaDeletedBy FinanceCurrenciesDBFieldNameType
}

var FinanceCurrenciesDBFieldName = financeCurrenciesDBFieldName{
	Code:          "code",
	DecimalCode:   "decimal_code",
	Exponent:      "exponent",
	IsActive:      "is_active",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewFinanceCurrenciesDBFieldNameFromStr(field string) (dbField FinanceCurrenciesDBFieldNameType, found bool) {
	switch field {

	case string(FinanceCurrenciesDBFieldName.Code):
		return FinanceCurrenciesDBFieldName.Code, true

	case string(FinanceCurrenciesDBFieldName.DecimalCode):
		return FinanceCurrenciesDBFieldName.DecimalCode, true

	case string(FinanceCurrenciesDBFieldName.Exponent):
		return FinanceCurrenciesDBFieldName.Exponent, true

	case string(FinanceCurrenciesDBFieldName.IsActive):
		return FinanceCurrenciesDBFieldName.IsActive, true

	case string(FinanceCurrenciesDBFieldName.Metadata):
		return FinanceCurrenciesDBFieldName.Metadata, true

	case string(FinanceCurrenciesDBFieldName.MetaCreatedAt):
		return FinanceCurrenciesDBFieldName.MetaCreatedAt, true

	case string(FinanceCurrenciesDBFieldName.MetaCreatedBy):
		return FinanceCurrenciesDBFieldName.MetaCreatedBy, true

	case string(FinanceCurrenciesDBFieldName.MetaUpdatedAt):
		return FinanceCurrenciesDBFieldName.MetaUpdatedAt, true

	case string(FinanceCurrenciesDBFieldName.MetaUpdatedBy):
		return FinanceCurrenciesDBFieldName.MetaUpdatedBy, true

	case string(FinanceCurrenciesDBFieldName.MetaDeletedAt):
		return FinanceCurrenciesDBFieldName.MetaDeletedAt, true

	case string(FinanceCurrenciesDBFieldName.MetaDeletedBy):
		return FinanceCurrenciesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceCurrenciesFilterJoins = map[string]JoinSpec{}

var FinanceCurrenciesFilterFields = map[string]FilterFieldSpec{
	"code": {
		SourcePath:        "code",
		DefaultOutputPath: "code",
		Column:            "code",
		SQLAlias:          "code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"decimal_code": {
		SourcePath:        "decimal_code",
		DefaultOutputPath: "decimalCode",
		Column:            "decimal_code",
		SQLAlias:          "decimal_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"exponent": {
		SourcePath:        "exponent",
		DefaultOutputPath: "exponent",
		Column:            "exponent",
		SQLAlias:          "exponent",
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

func NewFinanceCurrenciesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceCurrenciesFilterFields[field]
	return
}

type FinanceCurrenciesFilterResult struct {
	FinanceCurrencies
	FilterCount int `db:"count"`
}

func ValidateFinanceCurrenciesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceCurrenciesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceCurrenciesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceCurrenciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceCurrenciesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceCurrenciesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceCurrenciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceCurrenciesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FinanceCurrencies struct {
	Code        string          `db:"code"`
	DecimalCode null.Int        `db:"decimal_code"`
	Exponent    int16           `db:"exponent"`
	IsActive    bool            `db:"is_active"`
	Metadata    json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceCurrenciesPrimaryID struct {
	Code string `db:"code"`
}

func (d FinanceCurrencies) ToFinanceCurrenciesPrimaryID() FinanceCurrenciesPrimaryID {
	return FinanceCurrenciesPrimaryID{
		Code: d.Code,
	}
}

type FinanceCurrenciesList []*FinanceCurrencies

type FinanceCurrenciesFilterResultList []*FinanceCurrenciesFilterResult

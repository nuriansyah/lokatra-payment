package model

import (
	"encoding/json"
	"fmt"

	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerAccountTypesDBFieldNameType string

type ledgerAccountTypesDBFieldName struct {
	Code             LedgerAccountTypesDBFieldNameType
	NormalSide       LedgerAccountTypesDBFieldNameType
	Category         LedgerAccountTypesDBFieldNameType
	Description      LedgerAccountTypesDBFieldNameType
	IsControlAccount LedgerAccountTypesDBFieldNameType
	Metadata         LedgerAccountTypesDBFieldNameType
	MetaCreatedAt    LedgerAccountTypesDBFieldNameType
	MetaCreatedBy    LedgerAccountTypesDBFieldNameType
	MetaUpdatedAt    LedgerAccountTypesDBFieldNameType
	MetaUpdatedBy    LedgerAccountTypesDBFieldNameType
	MetaDeletedAt    LedgerAccountTypesDBFieldNameType
	MetaDeletedBy    LedgerAccountTypesDBFieldNameType
}

var LedgerAccountTypesDBFieldName = ledgerAccountTypesDBFieldName{
	Code:             "code",
	NormalSide:       "normal_side",
	Category:         "category",
	Description:      "description",
	IsControlAccount: "is_control_account",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewLedgerAccountTypesDBFieldNameFromStr(field string) (dbField LedgerAccountTypesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerAccountTypesDBFieldName.Code):
		return LedgerAccountTypesDBFieldName.Code, true

	case string(LedgerAccountTypesDBFieldName.NormalSide):
		return LedgerAccountTypesDBFieldName.NormalSide, true

	case string(LedgerAccountTypesDBFieldName.Category):
		return LedgerAccountTypesDBFieldName.Category, true

	case string(LedgerAccountTypesDBFieldName.Description):
		return LedgerAccountTypesDBFieldName.Description, true

	case string(LedgerAccountTypesDBFieldName.IsControlAccount):
		return LedgerAccountTypesDBFieldName.IsControlAccount, true

	case string(LedgerAccountTypesDBFieldName.Metadata):
		return LedgerAccountTypesDBFieldName.Metadata, true

	case string(LedgerAccountTypesDBFieldName.MetaCreatedAt):
		return LedgerAccountTypesDBFieldName.MetaCreatedAt, true

	case string(LedgerAccountTypesDBFieldName.MetaCreatedBy):
		return LedgerAccountTypesDBFieldName.MetaCreatedBy, true

	case string(LedgerAccountTypesDBFieldName.MetaUpdatedAt):
		return LedgerAccountTypesDBFieldName.MetaUpdatedAt, true

	case string(LedgerAccountTypesDBFieldName.MetaUpdatedBy):
		return LedgerAccountTypesDBFieldName.MetaUpdatedBy, true

	case string(LedgerAccountTypesDBFieldName.MetaDeletedAt):
		return LedgerAccountTypesDBFieldName.MetaDeletedAt, true

	case string(LedgerAccountTypesDBFieldName.MetaDeletedBy):
		return LedgerAccountTypesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerAccountTypesFilterJoins = map[string]JoinSpec{}

var LedgerAccountTypesFilterFields = map[string]FilterFieldSpec{
	"code": {
		SourcePath:        "code",
		DefaultOutputPath: "code",
		Column:            "code",
		SQLAlias:          "code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"normal_side": {
		SourcePath:        "normal_side",
		DefaultOutputPath: "normalSide",
		Column:            "normal_side",
		SQLAlias:          "normal_side",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"category": {
		SourcePath:        "category",
		DefaultOutputPath: "category",
		Column:            "category",
		SQLAlias:          "category",
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
	"is_control_account": {
		SourcePath:        "is_control_account",
		DefaultOutputPath: "isControlAccount",
		Column:            "is_control_account",
		SQLAlias:          "is_control_account",
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

func NewLedgerAccountTypesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerAccountTypesFilterFields[field]
	return
}

type LedgerAccountTypesFilterResult struct {
	LedgerAccountTypes
	FilterCount int `db:"count"`
}

func ValidateLedgerAccountTypesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerAccountTypesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerAccountTypesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerAccountTypesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerAccountTypesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerAccountTypesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerAccountTypesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerAccountTypesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type Category string

const (
	CategoryAsset         Category = "asset"
	CategoryLiability     Category = "liability"
	CategoryEquity        Category = "equity"
	CategoryRevenue       Category = "revenue"
	CategoryExpense       Category = "expense"
	CategoryContraAsset   Category = "contra_asset"
	CategoryContraRevenue Category = "contra_revenue"
)

type NormalSide string

const (
	NormalSideDebit  NormalSide = "debit"
	NormalSideCredit NormalSide = "credit"
)

type LedgerAccountTypes struct {
	Code             string          `db:"code"`
	NormalSide       NormalSide      `db:"normal_side"`
	Category         Category        `db:"category"`
	Description      string          `db:"description"`
	IsControlAccount bool            `db:"is_control_account"`
	Metadata         json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerAccountTypesPrimaryID struct {
	Code string `db:"code"`
}

func (d LedgerAccountTypes) ToLedgerAccountTypesPrimaryID() LedgerAccountTypesPrimaryID {
	return LedgerAccountTypesPrimaryID{
		Code: d.Code,
	}
}

type LedgerAccountTypesList []*LedgerAccountTypes

type LedgerAccountTypesFilterResultList []*LedgerAccountTypesFilterResult

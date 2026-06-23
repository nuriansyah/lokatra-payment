package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type FeeRulesDBFieldNameType string

type feeRulesDBFieldName struct {
	Id               FeeRulesDBFieldNameType
	FeeRuleVersionId FeeRulesDBFieldNameType
	RuleName         FeeRulesDBFieldNameType
	MinAmount        FeeRulesDBFieldNameType
	MaxAmount        FeeRulesDBFieldNameType
	PercentageRate   FeeRulesDBFieldNameType
	FlatAmount       FeeRulesDBFieldNameType
	CapAmount        FeeRulesDBFieldNameType
	FloorAmount      FeeRulesDBFieldNameType
	TaxInclusive     FeeRulesDBFieldNameType
	SortOrder        FeeRulesDBFieldNameType
	Metadata         FeeRulesDBFieldNameType
	MetaCreatedAt    FeeRulesDBFieldNameType
	MetaCreatedBy    FeeRulesDBFieldNameType
	MetaUpdatedAt    FeeRulesDBFieldNameType
	MetaUpdatedBy    FeeRulesDBFieldNameType
	MetaDeletedAt    FeeRulesDBFieldNameType
	MetaDeletedBy    FeeRulesDBFieldNameType
}

var FeeRulesDBFieldName = feeRulesDBFieldName{
	Id:               "id",
	FeeRuleVersionId: "fee_rule_version_id",
	RuleName:         "rule_name",
	MinAmount:        "min_amount",
	MaxAmount:        "max_amount",
	PercentageRate:   "percentage_rate",
	FlatAmount:       "flat_amount",
	CapAmount:        "cap_amount",
	FloorAmount:      "floor_amount",
	TaxInclusive:     "tax_inclusive",
	SortOrder:        "sort_order",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewFeeRulesDBFieldNameFromStr(field string) (dbField FeeRulesDBFieldNameType, found bool) {
	switch field {

	case string(FeeRulesDBFieldName.Id):
		return FeeRulesDBFieldName.Id, true

	case string(FeeRulesDBFieldName.FeeRuleVersionId):
		return FeeRulesDBFieldName.FeeRuleVersionId, true

	case string(FeeRulesDBFieldName.RuleName):
		return FeeRulesDBFieldName.RuleName, true

	case string(FeeRulesDBFieldName.MinAmount):
		return FeeRulesDBFieldName.MinAmount, true

	case string(FeeRulesDBFieldName.MaxAmount):
		return FeeRulesDBFieldName.MaxAmount, true

	case string(FeeRulesDBFieldName.PercentageRate):
		return FeeRulesDBFieldName.PercentageRate, true

	case string(FeeRulesDBFieldName.FlatAmount):
		return FeeRulesDBFieldName.FlatAmount, true

	case string(FeeRulesDBFieldName.CapAmount):
		return FeeRulesDBFieldName.CapAmount, true

	case string(FeeRulesDBFieldName.FloorAmount):
		return FeeRulesDBFieldName.FloorAmount, true

	case string(FeeRulesDBFieldName.TaxInclusive):
		return FeeRulesDBFieldName.TaxInclusive, true

	case string(FeeRulesDBFieldName.SortOrder):
		return FeeRulesDBFieldName.SortOrder, true

	case string(FeeRulesDBFieldName.Metadata):
		return FeeRulesDBFieldName.Metadata, true

	case string(FeeRulesDBFieldName.MetaCreatedAt):
		return FeeRulesDBFieldName.MetaCreatedAt, true

	case string(FeeRulesDBFieldName.MetaCreatedBy):
		return FeeRulesDBFieldName.MetaCreatedBy, true

	case string(FeeRulesDBFieldName.MetaUpdatedAt):
		return FeeRulesDBFieldName.MetaUpdatedAt, true

	case string(FeeRulesDBFieldName.MetaUpdatedBy):
		return FeeRulesDBFieldName.MetaUpdatedBy, true

	case string(FeeRulesDBFieldName.MetaDeletedAt):
		return FeeRulesDBFieldName.MetaDeletedAt, true

	case string(FeeRulesDBFieldName.MetaDeletedBy):
		return FeeRulesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FeeRulesFilterJoins = map[string]JoinSpec{}

var FeeRulesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_rule_version_id": {
		SourcePath:        "fee_rule_version_id",
		DefaultOutputPath: "feeRuleVersionId",
		Column:            "fee_rule_version_id",
		SQLAlias:          "fee_rule_version_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rule_name": {
		SourcePath:        "rule_name",
		DefaultOutputPath: "ruleName",
		Column:            "rule_name",
		SQLAlias:          "rule_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"min_amount": {
		SourcePath:        "min_amount",
		DefaultOutputPath: "minAmount",
		Column:            "min_amount",
		SQLAlias:          "min_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"max_amount": {
		SourcePath:        "max_amount",
		DefaultOutputPath: "maxAmount",
		Column:            "max_amount",
		SQLAlias:          "max_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"percentage_rate": {
		SourcePath:        "percentage_rate",
		DefaultOutputPath: "percentageRate",
		Column:            "percentage_rate",
		SQLAlias:          "percentage_rate",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"flat_amount": {
		SourcePath:        "flat_amount",
		DefaultOutputPath: "flatAmount",
		Column:            "flat_amount",
		SQLAlias:          "flat_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"cap_amount": {
		SourcePath:        "cap_amount",
		DefaultOutputPath: "capAmount",
		Column:            "cap_amount",
		SQLAlias:          "cap_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"floor_amount": {
		SourcePath:        "floor_amount",
		DefaultOutputPath: "floorAmount",
		Column:            "floor_amount",
		SQLAlias:          "floor_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_inclusive": {
		SourcePath:        "tax_inclusive",
		DefaultOutputPath: "taxInclusive",
		Column:            "tax_inclusive",
		SQLAlias:          "tax_inclusive",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"sort_order": {
		SourcePath:        "sort_order",
		DefaultOutputPath: "sortOrder",
		Column:            "sort_order",
		SQLAlias:          "sort_order",
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

func NewFeeRulesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FeeRulesFilterFields[field]
	return
}

type FeeRulesFilterResult struct {
	FeeRules
	FilterCount int `db:"count"`
}

func ValidateFeeRulesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFeeRulesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFeeRulesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFeeRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFeeRulesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFeeRulesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFeeRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFeeRulesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FeeRules struct {
	Id               uuid.UUID           `db:"id"`
	FeeRuleVersionId uuid.UUID           `db:"fee_rule_version_id"`
	RuleName         string              `db:"rule_name"`
	MinAmount        decimal.NullDecimal `db:"min_amount"`
	MaxAmount        decimal.NullDecimal `db:"max_amount"`
	PercentageRate   decimal.NullDecimal `db:"percentage_rate"`
	FlatAmount       decimal.NullDecimal `db:"flat_amount"`
	CapAmount        decimal.NullDecimal `db:"cap_amount"`
	FloorAmount      decimal.NullDecimal `db:"floor_amount"`
	TaxInclusive     bool                `db:"tax_inclusive"`
	SortOrder        int                 `db:"sort_order"`
	Metadata         json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type FeeRulesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FeeRules) ToFeeRulesPrimaryID() FeeRulesPrimaryID {
	return FeeRulesPrimaryID{
		Id: d.Id,
	}
}

type FeeRulesList []*FeeRules

type FeeRulesFilterResultList []*FeeRulesFilterResult

package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FeeRuleSetsDBFieldNameType string

type feeRuleSetsDBFieldName struct {
	Id             FeeRuleSetsDBFieldNameType
	FeeProfileId   FeeRuleSetsDBFieldNameType
	RuleSetCode    FeeRuleSetsDBFieldNameType
	Precedence     FeeRuleSetsDBFieldNameType
	EffectiveFrom  FeeRuleSetsDBFieldNameType
	EffectiveUntil FeeRuleSetsDBFieldNameType
	RuleSetStatus  FeeRuleSetsDBFieldNameType
	Metadata       FeeRuleSetsDBFieldNameType
	MetaCreatedAt  FeeRuleSetsDBFieldNameType
	MetaCreatedBy  FeeRuleSetsDBFieldNameType
	MetaUpdatedAt  FeeRuleSetsDBFieldNameType
	MetaUpdatedBy  FeeRuleSetsDBFieldNameType
	MetaDeletedAt  FeeRuleSetsDBFieldNameType
	MetaDeletedBy  FeeRuleSetsDBFieldNameType
}

var FeeRuleSetsDBFieldName = feeRuleSetsDBFieldName{
	Id:             "id",
	FeeProfileId:   "fee_profile_id",
	RuleSetCode:    "rule_set_code",
	Precedence:     "precedence",
	EffectiveFrom:  "effective_from",
	EffectiveUntil: "effective_until",
	RuleSetStatus:  "rule_set_status",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewFeeRuleSetsDBFieldNameFromStr(field string) (dbField FeeRuleSetsDBFieldNameType, found bool) {
	switch field {

	case string(FeeRuleSetsDBFieldName.Id):
		return FeeRuleSetsDBFieldName.Id, true

	case string(FeeRuleSetsDBFieldName.FeeProfileId):
		return FeeRuleSetsDBFieldName.FeeProfileId, true

	case string(FeeRuleSetsDBFieldName.RuleSetCode):
		return FeeRuleSetsDBFieldName.RuleSetCode, true

	case string(FeeRuleSetsDBFieldName.Precedence):
		return FeeRuleSetsDBFieldName.Precedence, true

	case string(FeeRuleSetsDBFieldName.EffectiveFrom):
		return FeeRuleSetsDBFieldName.EffectiveFrom, true

	case string(FeeRuleSetsDBFieldName.EffectiveUntil):
		return FeeRuleSetsDBFieldName.EffectiveUntil, true

	case string(FeeRuleSetsDBFieldName.RuleSetStatus):
		return FeeRuleSetsDBFieldName.RuleSetStatus, true

	case string(FeeRuleSetsDBFieldName.Metadata):
		return FeeRuleSetsDBFieldName.Metadata, true

	case string(FeeRuleSetsDBFieldName.MetaCreatedAt):
		return FeeRuleSetsDBFieldName.MetaCreatedAt, true

	case string(FeeRuleSetsDBFieldName.MetaCreatedBy):
		return FeeRuleSetsDBFieldName.MetaCreatedBy, true

	case string(FeeRuleSetsDBFieldName.MetaUpdatedAt):
		return FeeRuleSetsDBFieldName.MetaUpdatedAt, true

	case string(FeeRuleSetsDBFieldName.MetaUpdatedBy):
		return FeeRuleSetsDBFieldName.MetaUpdatedBy, true

	case string(FeeRuleSetsDBFieldName.MetaDeletedAt):
		return FeeRuleSetsDBFieldName.MetaDeletedAt, true

	case string(FeeRuleSetsDBFieldName.MetaDeletedBy):
		return FeeRuleSetsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FeeRuleSetsFilterJoins = map[string]JoinSpec{}

var FeeRuleSetsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_profile_id": {
		SourcePath:        "fee_profile_id",
		DefaultOutputPath: "feeProfileId",
		Column:            "fee_profile_id",
		SQLAlias:          "fee_profile_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rule_set_code": {
		SourcePath:        "rule_set_code",
		DefaultOutputPath: "ruleSetCode",
		Column:            "rule_set_code",
		SQLAlias:          "rule_set_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"precedence": {
		SourcePath:        "precedence",
		DefaultOutputPath: "precedence",
		Column:            "precedence",
		SQLAlias:          "precedence",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"effective_from": {
		SourcePath:        "effective_from",
		DefaultOutputPath: "effectiveFrom",
		Column:            "effective_from",
		SQLAlias:          "effective_from",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"effective_until": {
		SourcePath:        "effective_until",
		DefaultOutputPath: "effectiveUntil",
		Column:            "effective_until",
		SQLAlias:          "effective_until",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rule_set_status": {
		SourcePath:        "rule_set_status",
		DefaultOutputPath: "ruleSetStatus",
		Column:            "rule_set_status",
		SQLAlias:          "rule_set_status",
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

func NewFeeRuleSetsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FeeRuleSetsFilterFields[field]
	return
}

type FeeRuleSetsFilterResult struct {
	FeeRuleSets
	FilterCount int `db:"count"`
}

func ValidateFeeRuleSetsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFeeRuleSetsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFeeRuleSetsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFeeRuleSetsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFeeRuleSetsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFeeRuleSetsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFeeRuleSetsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFeeRuleSetsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RuleSetStatus string

const (
	RuleSetStatusActive   RuleSetStatus = "active"
	RuleSetStatusInactive RuleSetStatus = "inactive"
)

type FeeRuleSets struct {
	Id             uuid.UUID       `db:"id"`
	FeeProfileId   uuid.UUID       `db:"fee_profile_id"`
	RuleSetCode    string          `db:"rule_set_code"`
	Precedence     int             `db:"precedence"`
	EffectiveFrom  time.Time       `db:"effective_from"`
	EffectiveUntil null.Time       `db:"effective_until"`
	RuleSetStatus  RuleSetStatus   `db:"rule_set_status"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FeeRuleSetsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FeeRuleSets) ToFeeRuleSetsPrimaryID() FeeRuleSetsPrimaryID {
	return FeeRuleSetsPrimaryID{
		Id: d.Id,
	}
}

type FeeRuleSetsList []*FeeRuleSets

type FeeRuleSetsFilterResultList []*FeeRuleSetsFilterResult

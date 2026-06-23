package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FeeRuleVersionsDBFieldNameType string

type feeRuleVersionsDBFieldName struct {
	Id            FeeRuleVersionsDBFieldNameType
	RuleSetId     FeeRuleVersionsDBFieldNameType
	VersionNo     FeeRuleVersionsDBFieldNameType
	FormulaType   FeeRuleVersionsDBFieldNameType
	AppliesTo     FeeRuleVersionsDBFieldNameType
	PayerType     FeeRuleVersionsDBFieldNameType
	RecipientType FeeRuleVersionsDBFieldNameType
	Conditions    FeeRuleVersionsDBFieldNameType
	IsCurrent     FeeRuleVersionsDBFieldNameType
	MetaCreatedAt FeeRuleVersionsDBFieldNameType
	MetaCreatedBy FeeRuleVersionsDBFieldNameType
	MetaUpdatedAt FeeRuleVersionsDBFieldNameType
	MetaUpdatedBy FeeRuleVersionsDBFieldNameType
	MetaDeletedAt FeeRuleVersionsDBFieldNameType
	MetaDeletedBy FeeRuleVersionsDBFieldNameType
}

var FeeRuleVersionsDBFieldName = feeRuleVersionsDBFieldName{
	Id:            "id",
	RuleSetId:     "rule_set_id",
	VersionNo:     "version_no",
	FormulaType:   "formula_type",
	AppliesTo:     "applies_to",
	PayerType:     "payer_type",
	RecipientType: "recipient_type",
	Conditions:    "conditions",
	IsCurrent:     "is_current",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewFeeRuleVersionsDBFieldNameFromStr(field string) (dbField FeeRuleVersionsDBFieldNameType, found bool) {
	switch field {

	case string(FeeRuleVersionsDBFieldName.Id):
		return FeeRuleVersionsDBFieldName.Id, true

	case string(FeeRuleVersionsDBFieldName.RuleSetId):
		return FeeRuleVersionsDBFieldName.RuleSetId, true

	case string(FeeRuleVersionsDBFieldName.VersionNo):
		return FeeRuleVersionsDBFieldName.VersionNo, true

	case string(FeeRuleVersionsDBFieldName.FormulaType):
		return FeeRuleVersionsDBFieldName.FormulaType, true

	case string(FeeRuleVersionsDBFieldName.AppliesTo):
		return FeeRuleVersionsDBFieldName.AppliesTo, true

	case string(FeeRuleVersionsDBFieldName.PayerType):
		return FeeRuleVersionsDBFieldName.PayerType, true

	case string(FeeRuleVersionsDBFieldName.RecipientType):
		return FeeRuleVersionsDBFieldName.RecipientType, true

	case string(FeeRuleVersionsDBFieldName.Conditions):
		return FeeRuleVersionsDBFieldName.Conditions, true

	case string(FeeRuleVersionsDBFieldName.IsCurrent):
		return FeeRuleVersionsDBFieldName.IsCurrent, true

	case string(FeeRuleVersionsDBFieldName.MetaCreatedAt):
		return FeeRuleVersionsDBFieldName.MetaCreatedAt, true

	case string(FeeRuleVersionsDBFieldName.MetaCreatedBy):
		return FeeRuleVersionsDBFieldName.MetaCreatedBy, true

	case string(FeeRuleVersionsDBFieldName.MetaUpdatedAt):
		return FeeRuleVersionsDBFieldName.MetaUpdatedAt, true

	case string(FeeRuleVersionsDBFieldName.MetaUpdatedBy):
		return FeeRuleVersionsDBFieldName.MetaUpdatedBy, true

	case string(FeeRuleVersionsDBFieldName.MetaDeletedAt):
		return FeeRuleVersionsDBFieldName.MetaDeletedAt, true

	case string(FeeRuleVersionsDBFieldName.MetaDeletedBy):
		return FeeRuleVersionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FeeRuleVersionsFilterJoins = map[string]JoinSpec{}

var FeeRuleVersionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rule_set_id": {
		SourcePath:        "rule_set_id",
		DefaultOutputPath: "ruleSetId",
		Column:            "rule_set_id",
		SQLAlias:          "rule_set_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"version_no": {
		SourcePath:        "version_no",
		DefaultOutputPath: "versionNo",
		Column:            "version_no",
		SQLAlias:          "version_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"formula_type": {
		SourcePath:        "formula_type",
		DefaultOutputPath: "formulaType",
		Column:            "formula_type",
		SQLAlias:          "formula_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"applies_to": {
		SourcePath:        "applies_to",
		DefaultOutputPath: "appliesTo",
		Column:            "applies_to",
		SQLAlias:          "applies_to",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payer_type": {
		SourcePath:        "payer_type",
		DefaultOutputPath: "payerType",
		Column:            "payer_type",
		SQLAlias:          "payer_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"recipient_type": {
		SourcePath:        "recipient_type",
		DefaultOutputPath: "recipientType",
		Column:            "recipient_type",
		SQLAlias:          "recipient_type",
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
	"is_current": {
		SourcePath:        "is_current",
		DefaultOutputPath: "isCurrent",
		Column:            "is_current",
		SQLAlias:          "is_current",
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

func NewFeeRuleVersionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FeeRuleVersionsFilterFields[field]
	return
}

type FeeRuleVersionsFilterResult struct {
	FeeRuleVersions
	FilterCount int `db:"count"`
}

func ValidateFeeRuleVersionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFeeRuleVersionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFeeRuleVersionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFeeRuleVersionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFeeRuleVersionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFeeRuleVersionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFeeRuleVersionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFeeRuleVersionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type AppliesTo string

const (
	AppliesToGmv                  AppliesTo = "gmv"
	AppliesToNet                  AppliesTo = "net"
	AppliesToCommissionableAmount AppliesTo = "commissionable_amount"
	AppliesToPayout               AppliesTo = "payout"
	AppliesToRefund               AppliesTo = "refund"
)

type FormulaType string

const (
	FormulaTypeFlat       FormulaType = "flat"
	FormulaTypePercentage FormulaType = "percentage"
	FormulaTypeTiered     FormulaType = "tiered"
	FormulaTypeHybrid     FormulaType = "hybrid"
)

type PayerType string

const (
	PayerTypeCustomer PayerType = "customer"
	PayerTypeMerchant PayerType = "merchant"
	PayerTypePlatform PayerType = "platform"
)

type RecipientType string

const (
	RecipientTypePlatform     RecipientType = "platform"
	RecipientTypeProvider     RecipientType = "provider"
	RecipientTypeTaxAuthority RecipientType = "tax_authority"
	RecipientTypeAffiliate    RecipientType = "affiliate"
)

type FeeRuleVersions struct {
	Id            uuid.UUID       `db:"id"`
	RuleSetId     uuid.UUID       `db:"rule_set_id"`
	VersionNo     int             `db:"version_no"`
	FormulaType   FormulaType     `db:"formula_type"`
	AppliesTo     AppliesTo       `db:"applies_to"`
	PayerType     PayerType       `db:"payer_type"`
	RecipientType RecipientType   `db:"recipient_type"`
	Conditions    json.RawMessage `db:"conditions"`
	IsCurrent     bool            `db:"is_current"`

	shared.MetaSignature
}
type FeeRuleVersionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FeeRuleVersions) ToFeeRuleVersionsPrimaryID() FeeRuleVersionsPrimaryID {
	return FeeRuleVersionsPrimaryID{
		Id: d.Id,
	}
}

type FeeRuleVersionsList []*FeeRuleVersions

type FeeRuleVersionsFilterResultList []*FeeRuleVersionsFilterResult

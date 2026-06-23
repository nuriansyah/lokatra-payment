package model

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
	"time"
)

type TaxRulesDBFieldNameType string

type taxRulesDBFieldName struct {
	Id                 TaxRulesDBFieldNameType
	RuleCode           TaxRulesDBFieldNameType
	CountryCode        TaxRulesDBFieldNameType
	TaxType            TaxRulesDBFieldNameType
	Rate               TaxRulesDBFieldNameType
	TaxInclusive       TaxRulesDBFieldNameType
	LiabilityPartyId   TaxRulesDBFieldNameType
	BeneficiaryPartyId TaxRulesDBFieldNameType
	Priority           TaxRulesDBFieldNameType
	IsActive           TaxRulesDBFieldNameType
	ValidFrom          TaxRulesDBFieldNameType
	ValidUntil         TaxRulesDBFieldNameType
	Metadata           TaxRulesDBFieldNameType
	MetaCreatedAt      TaxRulesDBFieldNameType
	MetaCreatedBy      TaxRulesDBFieldNameType
	MetaUpdatedAt      TaxRulesDBFieldNameType
	MetaUpdatedBy      TaxRulesDBFieldNameType
	MetaDeletedAt      TaxRulesDBFieldNameType
	MetaDeletedBy      TaxRulesDBFieldNameType
}

var TaxRulesDBFieldName = taxRulesDBFieldName{
	Id:                 "id",
	RuleCode:           "rule_code",
	CountryCode:        "country_code",
	TaxType:            "tax_type",
	Rate:               "rate",
	TaxInclusive:       "tax_inclusive",
	LiabilityPartyId:   "liability_party_id",
	BeneficiaryPartyId: "beneficiary_party_id",
	Priority:           "priority",
	IsActive:           "is_active",
	ValidFrom:          "valid_from",
	ValidUntil:         "valid_until",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewTaxRulesDBFieldNameFromStr(field string) (dbField TaxRulesDBFieldNameType, found bool) {
	switch field {

	case string(TaxRulesDBFieldName.Id):
		return TaxRulesDBFieldName.Id, true

	case string(TaxRulesDBFieldName.RuleCode):
		return TaxRulesDBFieldName.RuleCode, true

	case string(TaxRulesDBFieldName.CountryCode):
		return TaxRulesDBFieldName.CountryCode, true

	case string(TaxRulesDBFieldName.TaxType):
		return TaxRulesDBFieldName.TaxType, true

	case string(TaxRulesDBFieldName.Rate):
		return TaxRulesDBFieldName.Rate, true

	case string(TaxRulesDBFieldName.TaxInclusive):
		return TaxRulesDBFieldName.TaxInclusive, true

	case string(TaxRulesDBFieldName.LiabilityPartyId):
		return TaxRulesDBFieldName.LiabilityPartyId, true

	case string(TaxRulesDBFieldName.BeneficiaryPartyId):
		return TaxRulesDBFieldName.BeneficiaryPartyId, true

	case string(TaxRulesDBFieldName.Priority):
		return TaxRulesDBFieldName.Priority, true

	case string(TaxRulesDBFieldName.IsActive):
		return TaxRulesDBFieldName.IsActive, true

	case string(TaxRulesDBFieldName.ValidFrom):
		return TaxRulesDBFieldName.ValidFrom, true

	case string(TaxRulesDBFieldName.ValidUntil):
		return TaxRulesDBFieldName.ValidUntil, true

	case string(TaxRulesDBFieldName.Metadata):
		return TaxRulesDBFieldName.Metadata, true

	case string(TaxRulesDBFieldName.MetaCreatedAt):
		return TaxRulesDBFieldName.MetaCreatedAt, true

	case string(TaxRulesDBFieldName.MetaCreatedBy):
		return TaxRulesDBFieldName.MetaCreatedBy, true

	case string(TaxRulesDBFieldName.MetaUpdatedAt):
		return TaxRulesDBFieldName.MetaUpdatedAt, true

	case string(TaxRulesDBFieldName.MetaUpdatedBy):
		return TaxRulesDBFieldName.MetaUpdatedBy, true

	case string(TaxRulesDBFieldName.MetaDeletedAt):
		return TaxRulesDBFieldName.MetaDeletedAt, true

	case string(TaxRulesDBFieldName.MetaDeletedBy):
		return TaxRulesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxRulesFilterJoins = map[string]JoinSpec{}

var TaxRulesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rule_code": {
		SourcePath:        "rule_code",
		DefaultOutputPath: "ruleCode",
		Column:            "rule_code",
		SQLAlias:          "rule_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"country_code": {
		SourcePath:        "country_code",
		DefaultOutputPath: "countryCode",
		Column:            "country_code",
		SQLAlias:          "country_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_type": {
		SourcePath:        "tax_type",
		DefaultOutputPath: "taxType",
		Column:            "tax_type",
		SQLAlias:          "tax_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rate": {
		SourcePath:        "rate",
		DefaultOutputPath: "rate",
		Column:            "rate",
		SQLAlias:          "rate",
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
	"liability_party_id": {
		SourcePath:        "liability_party_id",
		DefaultOutputPath: "liabilityPartyId",
		Column:            "liability_party_id",
		SQLAlias:          "liability_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"beneficiary_party_id": {
		SourcePath:        "beneficiary_party_id",
		DefaultOutputPath: "beneficiaryPartyId",
		Column:            "beneficiary_party_id",
		SQLAlias:          "beneficiary_party_id",
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
	"valid_from": {
		SourcePath:        "valid_from",
		DefaultOutputPath: "validFrom",
		Column:            "valid_from",
		SQLAlias:          "valid_from",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"valid_until": {
		SourcePath:        "valid_until",
		DefaultOutputPath: "validUntil",
		Column:            "valid_until",
		SQLAlias:          "valid_until",
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

func NewTaxRulesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxRulesFilterFields[field]
	return
}

type TaxRulesFilterResult struct {
	TaxRules
	FilterCount int `db:"count"`
}

func ValidateTaxRulesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxRulesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxRulesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxRulesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxRulesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxRulesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type TaxRules struct {
	Id                 uuid.UUID       `db:"id"`
	RuleCode           string          `db:"rule_code"`
	CountryCode        string          `db:"country_code"`
	TaxType            string          `db:"tax_type"`
	Rate               decimal.Decimal `db:"rate"`
	TaxInclusive       bool            `db:"tax_inclusive"`
	LiabilityPartyId   nuuid.NUUID     `db:"liability_party_id"`
	BeneficiaryPartyId nuuid.NUUID     `db:"beneficiary_party_id"`
	Priority           int             `db:"priority"`
	IsActive           bool            `db:"is_active"`
	ValidFrom          time.Time       `db:"valid_from"`
	ValidUntil         null.Time       `db:"valid_until"`
	Metadata           json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type TaxRulesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d TaxRules) ToTaxRulesPrimaryID() TaxRulesPrimaryID {
	return TaxRulesPrimaryID{
		Id: d.Id,
	}
}

type TaxRulesList []*TaxRules

type TaxRulesFilterResultList []*TaxRulesFilterResult

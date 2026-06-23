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

type TaxBreakdownsDBFieldNameType string

type taxBreakdownsDBFieldName struct {
	Id                 TaxBreakdownsDBFieldNameType
	SourceType         TaxBreakdownsDBFieldNameType
	SourceId           TaxBreakdownsDBFieldNameType
	TaxRuleId          TaxBreakdownsDBFieldNameType
	CurrencyCode       TaxBreakdownsDBFieldNameType
	TaxableAmount      TaxBreakdownsDBFieldNameType
	TaxAmount          TaxBreakdownsDBFieldNameType
	LiabilityPartyId   TaxBreakdownsDBFieldNameType
	BeneficiaryPartyId TaxBreakdownsDBFieldNameType
	BreakdownStatus    TaxBreakdownsDBFieldNameType
	Metadata           TaxBreakdownsDBFieldNameType
	MetaCreatedAt      TaxBreakdownsDBFieldNameType
	MetaCreatedBy      TaxBreakdownsDBFieldNameType
	MetaUpdatedAt      TaxBreakdownsDBFieldNameType
	MetaUpdatedBy      TaxBreakdownsDBFieldNameType
	MetaDeletedAt      TaxBreakdownsDBFieldNameType
	MetaDeletedBy      TaxBreakdownsDBFieldNameType
}

var TaxBreakdownsDBFieldName = taxBreakdownsDBFieldName{
	Id:                 "id",
	SourceType:         "source_type",
	SourceId:           "source_id",
	TaxRuleId:          "tax_rule_id",
	CurrencyCode:       "currency_code",
	TaxableAmount:      "taxable_amount",
	TaxAmount:          "tax_amount",
	LiabilityPartyId:   "liability_party_id",
	BeneficiaryPartyId: "beneficiary_party_id",
	BreakdownStatus:    "breakdown_status",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewTaxBreakdownsDBFieldNameFromStr(field string) (dbField TaxBreakdownsDBFieldNameType, found bool) {
	switch field {

	case string(TaxBreakdownsDBFieldName.Id):
		return TaxBreakdownsDBFieldName.Id, true

	case string(TaxBreakdownsDBFieldName.SourceType):
		return TaxBreakdownsDBFieldName.SourceType, true

	case string(TaxBreakdownsDBFieldName.SourceId):
		return TaxBreakdownsDBFieldName.SourceId, true

	case string(TaxBreakdownsDBFieldName.TaxRuleId):
		return TaxBreakdownsDBFieldName.TaxRuleId, true

	case string(TaxBreakdownsDBFieldName.CurrencyCode):
		return TaxBreakdownsDBFieldName.CurrencyCode, true

	case string(TaxBreakdownsDBFieldName.TaxableAmount):
		return TaxBreakdownsDBFieldName.TaxableAmount, true

	case string(TaxBreakdownsDBFieldName.TaxAmount):
		return TaxBreakdownsDBFieldName.TaxAmount, true

	case string(TaxBreakdownsDBFieldName.LiabilityPartyId):
		return TaxBreakdownsDBFieldName.LiabilityPartyId, true

	case string(TaxBreakdownsDBFieldName.BeneficiaryPartyId):
		return TaxBreakdownsDBFieldName.BeneficiaryPartyId, true

	case string(TaxBreakdownsDBFieldName.BreakdownStatus):
		return TaxBreakdownsDBFieldName.BreakdownStatus, true

	case string(TaxBreakdownsDBFieldName.Metadata):
		return TaxBreakdownsDBFieldName.Metadata, true

	case string(TaxBreakdownsDBFieldName.MetaCreatedAt):
		return TaxBreakdownsDBFieldName.MetaCreatedAt, true

	case string(TaxBreakdownsDBFieldName.MetaCreatedBy):
		return TaxBreakdownsDBFieldName.MetaCreatedBy, true

	case string(TaxBreakdownsDBFieldName.MetaUpdatedAt):
		return TaxBreakdownsDBFieldName.MetaUpdatedAt, true

	case string(TaxBreakdownsDBFieldName.MetaUpdatedBy):
		return TaxBreakdownsDBFieldName.MetaUpdatedBy, true

	case string(TaxBreakdownsDBFieldName.MetaDeletedAt):
		return TaxBreakdownsDBFieldName.MetaDeletedAt, true

	case string(TaxBreakdownsDBFieldName.MetaDeletedBy):
		return TaxBreakdownsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxBreakdownsFilterJoins = map[string]JoinSpec{}

var TaxBreakdownsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_rule_id": {
		SourcePath:        "tax_rule_id",
		DefaultOutputPath: "taxRuleId",
		Column:            "tax_rule_id",
		SQLAlias:          "tax_rule_id",
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
	"taxable_amount": {
		SourcePath:        "taxable_amount",
		DefaultOutputPath: "taxableAmount",
		Column:            "taxable_amount",
		SQLAlias:          "taxable_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_amount": {
		SourcePath:        "tax_amount",
		DefaultOutputPath: "taxAmount",
		Column:            "tax_amount",
		SQLAlias:          "tax_amount",
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
	"breakdown_status": {
		SourcePath:        "breakdown_status",
		DefaultOutputPath: "breakdownStatus",
		Column:            "breakdown_status",
		SQLAlias:          "breakdown_status",
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

func NewTaxBreakdownsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxBreakdownsFilterFields[field]
	return
}

type TaxBreakdownsFilterResult struct {
	TaxBreakdowns
	FilterCount int `db:"count"`
}

func ValidateTaxBreakdownsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxBreakdownsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxBreakdownsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxBreakdownsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxBreakdownsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxBreakdownsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxBreakdownsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxBreakdownsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type TaxBreakdownsBreakdownStatus string

const (
	TaxBreakdownsBreakdownStatusComputed TaxBreakdownsBreakdownStatus = "computed"
	TaxBreakdownsBreakdownStatusPosted   TaxBreakdownsBreakdownStatus = "posted"
	TaxBreakdownsBreakdownStatusVoid     TaxBreakdownsBreakdownStatus = "void"
)

type TaxBreakdowns struct {
	Id                 uuid.UUID                    `db:"id"`
	SourceType         string                       `db:"source_type"`
	SourceId           uuid.UUID                    `db:"source_id"`
	TaxRuleId          nuuid.NUUID                  `db:"tax_rule_id"`
	CurrencyCode       string                       `db:"currency_code"`
	TaxableAmount      decimal.Decimal              `db:"taxable_amount"`
	TaxAmount          decimal.Decimal              `db:"tax_amount"`
	LiabilityPartyId   nuuid.NUUID                  `db:"liability_party_id"`
	BeneficiaryPartyId nuuid.NUUID                  `db:"beneficiary_party_id"`
	BreakdownStatus    TaxBreakdownsBreakdownStatus `db:"breakdown_status"`
	Metadata           json.RawMessage              `db:"metadata"`

	shared.MetaSignature
}
type TaxBreakdownsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d TaxBreakdowns) ToTaxBreakdownsPrimaryID() TaxBreakdownsPrimaryID {
	return TaxBreakdownsPrimaryID{
		Id: d.Id,
	}
}

type TaxBreakdownsList []*TaxBreakdowns

type TaxBreakdownsFilterResultList []*TaxBreakdownsFilterResult

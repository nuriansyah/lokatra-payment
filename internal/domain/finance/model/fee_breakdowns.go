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

type FeeBreakdownsDBFieldNameType string

type feeBreakdownsDBFieldName struct {
	Id               FeeBreakdownsDBFieldNameType
	SourceType       FeeBreakdownsDBFieldNameType
	SourceId         FeeBreakdownsDBFieldNameType
	FeeProfileId     FeeBreakdownsDBFieldNameType
	FeeRuleId        FeeBreakdownsDBFieldNameType
	CurrencyCode     FeeBreakdownsDBFieldNameType
	BaseAmount       FeeBreakdownsDBFieldNameType
	FeeAmount        FeeBreakdownsDBFieldNameType
	TaxAmount        FeeBreakdownsDBFieldNameType
	RecipientPartyId FeeBreakdownsDBFieldNameType
	BreakdownStatus  FeeBreakdownsDBFieldNameType
	Metadata         FeeBreakdownsDBFieldNameType
	MetaCreatedAt    FeeBreakdownsDBFieldNameType
	MetaCreatedBy    FeeBreakdownsDBFieldNameType
	MetaUpdatedAt    FeeBreakdownsDBFieldNameType
	MetaUpdatedBy    FeeBreakdownsDBFieldNameType
	MetaDeletedAt    FeeBreakdownsDBFieldNameType
	MetaDeletedBy    FeeBreakdownsDBFieldNameType
}

var FeeBreakdownsDBFieldName = feeBreakdownsDBFieldName{
	Id:               "id",
	SourceType:       "source_type",
	SourceId:         "source_id",
	FeeProfileId:     "fee_profile_id",
	FeeRuleId:        "fee_rule_id",
	CurrencyCode:     "currency_code",
	BaseAmount:       "base_amount",
	FeeAmount:        "fee_amount",
	TaxAmount:        "tax_amount",
	RecipientPartyId: "recipient_party_id",
	BreakdownStatus:  "breakdown_status",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewFeeBreakdownsDBFieldNameFromStr(field string) (dbField FeeBreakdownsDBFieldNameType, found bool) {
	switch field {

	case string(FeeBreakdownsDBFieldName.Id):
		return FeeBreakdownsDBFieldName.Id, true

	case string(FeeBreakdownsDBFieldName.SourceType):
		return FeeBreakdownsDBFieldName.SourceType, true

	case string(FeeBreakdownsDBFieldName.SourceId):
		return FeeBreakdownsDBFieldName.SourceId, true

	case string(FeeBreakdownsDBFieldName.FeeProfileId):
		return FeeBreakdownsDBFieldName.FeeProfileId, true

	case string(FeeBreakdownsDBFieldName.FeeRuleId):
		return FeeBreakdownsDBFieldName.FeeRuleId, true

	case string(FeeBreakdownsDBFieldName.CurrencyCode):
		return FeeBreakdownsDBFieldName.CurrencyCode, true

	case string(FeeBreakdownsDBFieldName.BaseAmount):
		return FeeBreakdownsDBFieldName.BaseAmount, true

	case string(FeeBreakdownsDBFieldName.FeeAmount):
		return FeeBreakdownsDBFieldName.FeeAmount, true

	case string(FeeBreakdownsDBFieldName.TaxAmount):
		return FeeBreakdownsDBFieldName.TaxAmount, true

	case string(FeeBreakdownsDBFieldName.RecipientPartyId):
		return FeeBreakdownsDBFieldName.RecipientPartyId, true

	case string(FeeBreakdownsDBFieldName.BreakdownStatus):
		return FeeBreakdownsDBFieldName.BreakdownStatus, true

	case string(FeeBreakdownsDBFieldName.Metadata):
		return FeeBreakdownsDBFieldName.Metadata, true

	case string(FeeBreakdownsDBFieldName.MetaCreatedAt):
		return FeeBreakdownsDBFieldName.MetaCreatedAt, true

	case string(FeeBreakdownsDBFieldName.MetaCreatedBy):
		return FeeBreakdownsDBFieldName.MetaCreatedBy, true

	case string(FeeBreakdownsDBFieldName.MetaUpdatedAt):
		return FeeBreakdownsDBFieldName.MetaUpdatedAt, true

	case string(FeeBreakdownsDBFieldName.MetaUpdatedBy):
		return FeeBreakdownsDBFieldName.MetaUpdatedBy, true

	case string(FeeBreakdownsDBFieldName.MetaDeletedAt):
		return FeeBreakdownsDBFieldName.MetaDeletedAt, true

	case string(FeeBreakdownsDBFieldName.MetaDeletedBy):
		return FeeBreakdownsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FeeBreakdownsFilterJoins = map[string]JoinSpec{}

var FeeBreakdownsFilterFields = map[string]FilterFieldSpec{
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
	"fee_profile_id": {
		SourcePath:        "fee_profile_id",
		DefaultOutputPath: "feeProfileId",
		Column:            "fee_profile_id",
		SQLAlias:          "fee_profile_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_rule_id": {
		SourcePath:        "fee_rule_id",
		DefaultOutputPath: "feeRuleId",
		Column:            "fee_rule_id",
		SQLAlias:          "fee_rule_id",
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
	"base_amount": {
		SourcePath:        "base_amount",
		DefaultOutputPath: "baseAmount",
		Column:            "base_amount",
		SQLAlias:          "base_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_amount": {
		SourcePath:        "fee_amount",
		DefaultOutputPath: "feeAmount",
		Column:            "fee_amount",
		SQLAlias:          "fee_amount",
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
	"recipient_party_id": {
		SourcePath:        "recipient_party_id",
		DefaultOutputPath: "recipientPartyId",
		Column:            "recipient_party_id",
		SQLAlias:          "recipient_party_id",
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

func NewFeeBreakdownsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FeeBreakdownsFilterFields[field]
	return
}

type FeeBreakdownsFilterResult struct {
	FeeBreakdowns
	FilterCount int `db:"count"`
}

func ValidateFeeBreakdownsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFeeBreakdownsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFeeBreakdownsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFeeBreakdownsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFeeBreakdownsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFeeBreakdownsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFeeBreakdownsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFeeBreakdownsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FeeBreakdownsBreakdownStatus string

const (
	FeeBreakdownsBreakdownStatusComputed FeeBreakdownsBreakdownStatus = "computed"
	FeeBreakdownsBreakdownStatusPosted   FeeBreakdownsBreakdownStatus = "posted"
	FeeBreakdownsBreakdownStatusReversed FeeBreakdownsBreakdownStatus = "reversed"
)

type FeeBreakdowns struct {
	Id               uuid.UUID                    `db:"id"`
	SourceType       string                       `db:"source_type"`
	SourceId         uuid.UUID                    `db:"source_id"`
	FeeProfileId     uuid.UUID                    `db:"fee_profile_id"`
	FeeRuleId        uuid.UUID                    `db:"fee_rule_id"`
	CurrencyCode     string                       `db:"currency_code"`
	BaseAmount       decimal.Decimal              `db:"base_amount"`
	FeeAmount        decimal.Decimal              `db:"fee_amount"`
	TaxAmount        decimal.Decimal              `db:"tax_amount"`
	RecipientPartyId nuuid.NUUID                  `db:"recipient_party_id"`
	BreakdownStatus  FeeBreakdownsBreakdownStatus `db:"breakdown_status"`
	Metadata         json.RawMessage              `db:"metadata"`

	shared.MetaSignature
}
type FeeBreakdownsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FeeBreakdowns) ToFeeBreakdownsPrimaryID() FeeBreakdownsPrimaryID {
	return FeeBreakdownsPrimaryID{
		Id: d.Id,
	}
}

type FeeBreakdownsList []*FeeBreakdowns

type FeeBreakdownsFilterResultList []*FeeBreakdownsFilterResult

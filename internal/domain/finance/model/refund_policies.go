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

type RefundPoliciesDBFieldNameType string

type refundPoliciesDBFieldName struct {
	Id                         RefundPoliciesDBFieldNameType
	PolicyCode                 RefundPoliciesDBFieldNameType
	MerchantPartyId            RefundPoliciesDBFieldNameType
	PolicyScope                RefundPoliciesDBFieldNameType
	AllowPartial               RefundPoliciesDBFieldNameType
	AllowPostPayout            RefundPoliciesDBFieldNameType
	FeeReturnMode              RefundPoliciesDBFieldNameType
	TaxReturnMode              RefundPoliciesDBFieldNameType
	RequiresApprovalOverAmount RefundPoliciesDBFieldNameType
	PolicyStatus               RefundPoliciesDBFieldNameType
	Metadata                   RefundPoliciesDBFieldNameType
	MetaCreatedAt              RefundPoliciesDBFieldNameType
	MetaCreatedBy              RefundPoliciesDBFieldNameType
	MetaUpdatedAt              RefundPoliciesDBFieldNameType
	MetaUpdatedBy              RefundPoliciesDBFieldNameType
	MetaDeletedAt              RefundPoliciesDBFieldNameType
	MetaDeletedBy              RefundPoliciesDBFieldNameType
}

var RefundPoliciesDBFieldName = refundPoliciesDBFieldName{
	Id:                         "id",
	PolicyCode:                 "policy_code",
	MerchantPartyId:            "merchant_party_id",
	PolicyScope:                "policy_scope",
	AllowPartial:               "allow_partial",
	AllowPostPayout:            "allow_post_payout",
	FeeReturnMode:              "fee_return_mode",
	TaxReturnMode:              "tax_return_mode",
	RequiresApprovalOverAmount: "requires_approval_over_amount",
	PolicyStatus:               "policy_status",
	Metadata:                   "metadata",
	MetaCreatedAt:              "meta_created_at",
	MetaCreatedBy:              "meta_created_by",
	MetaUpdatedAt:              "meta_updated_at",
	MetaUpdatedBy:              "meta_updated_by",
	MetaDeletedAt:              "meta_deleted_at",
	MetaDeletedBy:              "meta_deleted_by",
}

func NewRefundPoliciesDBFieldNameFromStr(field string) (dbField RefundPoliciesDBFieldNameType, found bool) {
	switch field {

	case string(RefundPoliciesDBFieldName.Id):
		return RefundPoliciesDBFieldName.Id, true

	case string(RefundPoliciesDBFieldName.PolicyCode):
		return RefundPoliciesDBFieldName.PolicyCode, true

	case string(RefundPoliciesDBFieldName.MerchantPartyId):
		return RefundPoliciesDBFieldName.MerchantPartyId, true

	case string(RefundPoliciesDBFieldName.PolicyScope):
		return RefundPoliciesDBFieldName.PolicyScope, true

	case string(RefundPoliciesDBFieldName.AllowPartial):
		return RefundPoliciesDBFieldName.AllowPartial, true

	case string(RefundPoliciesDBFieldName.AllowPostPayout):
		return RefundPoliciesDBFieldName.AllowPostPayout, true

	case string(RefundPoliciesDBFieldName.FeeReturnMode):
		return RefundPoliciesDBFieldName.FeeReturnMode, true

	case string(RefundPoliciesDBFieldName.TaxReturnMode):
		return RefundPoliciesDBFieldName.TaxReturnMode, true

	case string(RefundPoliciesDBFieldName.RequiresApprovalOverAmount):
		return RefundPoliciesDBFieldName.RequiresApprovalOverAmount, true

	case string(RefundPoliciesDBFieldName.PolicyStatus):
		return RefundPoliciesDBFieldName.PolicyStatus, true

	case string(RefundPoliciesDBFieldName.Metadata):
		return RefundPoliciesDBFieldName.Metadata, true

	case string(RefundPoliciesDBFieldName.MetaCreatedAt):
		return RefundPoliciesDBFieldName.MetaCreatedAt, true

	case string(RefundPoliciesDBFieldName.MetaCreatedBy):
		return RefundPoliciesDBFieldName.MetaCreatedBy, true

	case string(RefundPoliciesDBFieldName.MetaUpdatedAt):
		return RefundPoliciesDBFieldName.MetaUpdatedAt, true

	case string(RefundPoliciesDBFieldName.MetaUpdatedBy):
		return RefundPoliciesDBFieldName.MetaUpdatedBy, true

	case string(RefundPoliciesDBFieldName.MetaDeletedAt):
		return RefundPoliciesDBFieldName.MetaDeletedAt, true

	case string(RefundPoliciesDBFieldName.MetaDeletedBy):
		return RefundPoliciesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundPoliciesFilterJoins = map[string]JoinSpec{}

var RefundPoliciesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"policy_code": {
		SourcePath:        "policy_code",
		DefaultOutputPath: "policyCode",
		Column:            "policy_code",
		SQLAlias:          "policy_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"policy_scope": {
		SourcePath:        "policy_scope",
		DefaultOutputPath: "policyScope",
		Column:            "policy_scope",
		SQLAlias:          "policy_scope",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"allow_partial": {
		SourcePath:        "allow_partial",
		DefaultOutputPath: "allowPartial",
		Column:            "allow_partial",
		SQLAlias:          "allow_partial",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"allow_post_payout": {
		SourcePath:        "allow_post_payout",
		DefaultOutputPath: "allowPostPayout",
		Column:            "allow_post_payout",
		SQLAlias:          "allow_post_payout",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_return_mode": {
		SourcePath:        "fee_return_mode",
		DefaultOutputPath: "feeReturnMode",
		Column:            "fee_return_mode",
		SQLAlias:          "fee_return_mode",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_return_mode": {
		SourcePath:        "tax_return_mode",
		DefaultOutputPath: "taxReturnMode",
		Column:            "tax_return_mode",
		SQLAlias:          "tax_return_mode",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"requires_approval_over_amount": {
		SourcePath:        "requires_approval_over_amount",
		DefaultOutputPath: "requiresApprovalOverAmount",
		Column:            "requires_approval_over_amount",
		SQLAlias:          "requires_approval_over_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"policy_status": {
		SourcePath:        "policy_status",
		DefaultOutputPath: "policyStatus",
		Column:            "policy_status",
		SQLAlias:          "policy_status",
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

func NewRefundPoliciesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundPoliciesFilterFields[field]
	return
}

type RefundPoliciesFilterResult struct {
	RefundPolicies
	FilterCount int `db:"count"`
}

func ValidateRefundPoliciesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundPoliciesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundPoliciesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundPoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundPoliciesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundPoliciesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundPoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundPoliciesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FeeReturnMode string

const (
	FeeReturnModeNone         FeeReturnMode = "none"
	FeeReturnModeFull         FeeReturnMode = "full"
	FeeReturnModeProportional FeeReturnMode = "proportional"
)

type RefundPoliciesPolicyScope string

const (
	RefundPoliciesPolicyScopePlatformDefault RefundPoliciesPolicyScope = "platform_default"
	RefundPoliciesPolicyScopeMerchant        RefundPoliciesPolicyScope = "merchant"
	RefundPoliciesPolicyScopeListingType     RefundPoliciesPolicyScope = "listing_type"
	RefundPoliciesPolicyScopeCategory        RefundPoliciesPolicyScope = "category"
)

type RefundPoliciesPolicyStatus string

const (
	RefundPoliciesPolicyStatusActive   RefundPoliciesPolicyStatus = "active"
	RefundPoliciesPolicyStatusInactive RefundPoliciesPolicyStatus = "inactive"
)

type TaxReturnMode string

const (
	TaxReturnModeNone         TaxReturnMode = "none"
	TaxReturnModeFull         TaxReturnMode = "full"
	TaxReturnModeProportional TaxReturnMode = "proportional"
)

type RefundPolicies struct {
	Id                         uuid.UUID                  `db:"id"`
	PolicyCode                 string                     `db:"policy_code"`
	MerchantPartyId            nuuid.NUUID                `db:"merchant_party_id"`
	PolicyScope                RefundPoliciesPolicyScope  `db:"policy_scope"`
	AllowPartial               bool                       `db:"allow_partial"`
	AllowPostPayout            bool                       `db:"allow_post_payout"`
	FeeReturnMode              FeeReturnMode              `db:"fee_return_mode"`
	TaxReturnMode              TaxReturnMode              `db:"tax_return_mode"`
	RequiresApprovalOverAmount decimal.NullDecimal        `db:"requires_approval_over_amount"`
	PolicyStatus               RefundPoliciesPolicyStatus `db:"policy_status"`
	Metadata                   json.RawMessage            `db:"metadata"`

	shared.MetaSignature
}
type RefundPoliciesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundPolicies) ToRefundPoliciesPrimaryID() RefundPoliciesPrimaryID {
	return RefundPoliciesPrimaryID{
		Id: d.Id,
	}
}

type RefundPoliciesList []*RefundPolicies

type RefundPoliciesFilterResultList []*RefundPoliciesFilterResult

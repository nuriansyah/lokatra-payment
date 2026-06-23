package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type SettlementPoliciesDBFieldNameType string

type settlementPoliciesDBFieldName struct {
	Id              SettlementPoliciesDBFieldNameType
	PolicyCode      SettlementPoliciesDBFieldNameType
	MerchantPartyId SettlementPoliciesDBFieldNameType
	PolicyScope     SettlementPoliciesDBFieldNameType
	PolicyStatus    SettlementPoliciesDBFieldNameType
	Description     SettlementPoliciesDBFieldNameType
	Metadata        SettlementPoliciesDBFieldNameType
	MetaCreatedAt   SettlementPoliciesDBFieldNameType
	MetaCreatedBy   SettlementPoliciesDBFieldNameType
	MetaUpdatedAt   SettlementPoliciesDBFieldNameType
	MetaUpdatedBy   SettlementPoliciesDBFieldNameType
	MetaDeletedAt   SettlementPoliciesDBFieldNameType
	MetaDeletedBy   SettlementPoliciesDBFieldNameType
}

var SettlementPoliciesDBFieldName = settlementPoliciesDBFieldName{
	Id:              "id",
	PolicyCode:      "policy_code",
	MerchantPartyId: "merchant_party_id",
	PolicyScope:     "policy_scope",
	PolicyStatus:    "policy_status",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewSettlementPoliciesDBFieldNameFromStr(field string) (dbField SettlementPoliciesDBFieldNameType, found bool) {
	switch field {

	case string(SettlementPoliciesDBFieldName.Id):
		return SettlementPoliciesDBFieldName.Id, true

	case string(SettlementPoliciesDBFieldName.PolicyCode):
		return SettlementPoliciesDBFieldName.PolicyCode, true

	case string(SettlementPoliciesDBFieldName.MerchantPartyId):
		return SettlementPoliciesDBFieldName.MerchantPartyId, true

	case string(SettlementPoliciesDBFieldName.PolicyScope):
		return SettlementPoliciesDBFieldName.PolicyScope, true

	case string(SettlementPoliciesDBFieldName.PolicyStatus):
		return SettlementPoliciesDBFieldName.PolicyStatus, true

	case string(SettlementPoliciesDBFieldName.Description):
		return SettlementPoliciesDBFieldName.Description, true

	case string(SettlementPoliciesDBFieldName.Metadata):
		return SettlementPoliciesDBFieldName.Metadata, true

	case string(SettlementPoliciesDBFieldName.MetaCreatedAt):
		return SettlementPoliciesDBFieldName.MetaCreatedAt, true

	case string(SettlementPoliciesDBFieldName.MetaCreatedBy):
		return SettlementPoliciesDBFieldName.MetaCreatedBy, true

	case string(SettlementPoliciesDBFieldName.MetaUpdatedAt):
		return SettlementPoliciesDBFieldName.MetaUpdatedAt, true

	case string(SettlementPoliciesDBFieldName.MetaUpdatedBy):
		return SettlementPoliciesDBFieldName.MetaUpdatedBy, true

	case string(SettlementPoliciesDBFieldName.MetaDeletedAt):
		return SettlementPoliciesDBFieldName.MetaDeletedAt, true

	case string(SettlementPoliciesDBFieldName.MetaDeletedBy):
		return SettlementPoliciesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var SettlementPoliciesFilterJoins = map[string]JoinSpec{}

var SettlementPoliciesFilterFields = map[string]FilterFieldSpec{
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
	"policy_status": {
		SourcePath:        "policy_status",
		DefaultOutputPath: "policyStatus",
		Column:            "policy_status",
		SQLAlias:          "policy_status",
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

func NewSettlementPoliciesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = SettlementPoliciesFilterFields[field]
	return
}

type SettlementPoliciesFilterResult struct {
	SettlementPolicies
	FilterCount int `db:"count"`
}

func ValidateSettlementPoliciesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewSettlementPoliciesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewSettlementPoliciesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewSettlementPoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateSettlementPoliciesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateSettlementPoliciesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewSettlementPoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateSettlementPoliciesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type SettlementPoliciesPolicyScope string

const (
	SettlementPoliciesPolicyScopePlatformDefault SettlementPoliciesPolicyScope = "platform_default"
	SettlementPoliciesPolicyScopeMerchant        SettlementPoliciesPolicyScope = "merchant"
	SettlementPoliciesPolicyScopeListingType     SettlementPoliciesPolicyScope = "listing_type"
	SettlementPoliciesPolicyScopeCategory        SettlementPoliciesPolicyScope = "category"
)

type SettlementPoliciesPolicyStatus string

const (
	SettlementPoliciesPolicyStatusActive   SettlementPoliciesPolicyStatus = "active"
	SettlementPoliciesPolicyStatusInactive SettlementPoliciesPolicyStatus = "inactive"
)

type SettlementPolicies struct {
	Id              uuid.UUID                      `db:"id"`
	PolicyCode      string                         `db:"policy_code"`
	MerchantPartyId nuuid.NUUID                    `db:"merchant_party_id"`
	PolicyScope     SettlementPoliciesPolicyScope  `db:"policy_scope"`
	PolicyStatus    SettlementPoliciesPolicyStatus `db:"policy_status"`
	Description     null.String                    `db:"description"`
	Metadata        json.RawMessage                `db:"metadata"`

	shared.MetaSignature
}
type SettlementPoliciesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d SettlementPolicies) ToSettlementPoliciesPrimaryID() SettlementPoliciesPrimaryID {
	return SettlementPoliciesPrimaryID{
		Id: d.Id,
	}
}

type SettlementPoliciesList []*SettlementPolicies

type SettlementPoliciesFilterResultList []*SettlementPoliciesFilterResult

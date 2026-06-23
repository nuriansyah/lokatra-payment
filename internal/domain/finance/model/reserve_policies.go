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

type ReservePoliciesDBFieldNameType string

type reservePoliciesDBFieldName struct {
	Id              ReservePoliciesDBFieldNameType
	PolicyCode      ReservePoliciesDBFieldNameType
	MerchantPartyId ReservePoliciesDBFieldNameType
	PolicyScope     ReservePoliciesDBFieldNameType
	ReserveStatus   ReservePoliciesDBFieldNameType
	Description     ReservePoliciesDBFieldNameType
	Metadata        ReservePoliciesDBFieldNameType
	MetaCreatedAt   ReservePoliciesDBFieldNameType
	MetaCreatedBy   ReservePoliciesDBFieldNameType
	MetaUpdatedAt   ReservePoliciesDBFieldNameType
	MetaUpdatedBy   ReservePoliciesDBFieldNameType
	MetaDeletedAt   ReservePoliciesDBFieldNameType
	MetaDeletedBy   ReservePoliciesDBFieldNameType
}

var ReservePoliciesDBFieldName = reservePoliciesDBFieldName{
	Id:              "id",
	PolicyCode:      "policy_code",
	MerchantPartyId: "merchant_party_id",
	PolicyScope:     "policy_scope",
	ReserveStatus:   "reserve_status",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewReservePoliciesDBFieldNameFromStr(field string) (dbField ReservePoliciesDBFieldNameType, found bool) {
	switch field {

	case string(ReservePoliciesDBFieldName.Id):
		return ReservePoliciesDBFieldName.Id, true

	case string(ReservePoliciesDBFieldName.PolicyCode):
		return ReservePoliciesDBFieldName.PolicyCode, true

	case string(ReservePoliciesDBFieldName.MerchantPartyId):
		return ReservePoliciesDBFieldName.MerchantPartyId, true

	case string(ReservePoliciesDBFieldName.PolicyScope):
		return ReservePoliciesDBFieldName.PolicyScope, true

	case string(ReservePoliciesDBFieldName.ReserveStatus):
		return ReservePoliciesDBFieldName.ReserveStatus, true

	case string(ReservePoliciesDBFieldName.Description):
		return ReservePoliciesDBFieldName.Description, true

	case string(ReservePoliciesDBFieldName.Metadata):
		return ReservePoliciesDBFieldName.Metadata, true

	case string(ReservePoliciesDBFieldName.MetaCreatedAt):
		return ReservePoliciesDBFieldName.MetaCreatedAt, true

	case string(ReservePoliciesDBFieldName.MetaCreatedBy):
		return ReservePoliciesDBFieldName.MetaCreatedBy, true

	case string(ReservePoliciesDBFieldName.MetaUpdatedAt):
		return ReservePoliciesDBFieldName.MetaUpdatedAt, true

	case string(ReservePoliciesDBFieldName.MetaUpdatedBy):
		return ReservePoliciesDBFieldName.MetaUpdatedBy, true

	case string(ReservePoliciesDBFieldName.MetaDeletedAt):
		return ReservePoliciesDBFieldName.MetaDeletedAt, true

	case string(ReservePoliciesDBFieldName.MetaDeletedBy):
		return ReservePoliciesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReservePoliciesFilterJoins = map[string]JoinSpec{}

var ReservePoliciesFilterFields = map[string]FilterFieldSpec{
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
	"reserve_status": {
		SourcePath:        "reserve_status",
		DefaultOutputPath: "reserveStatus",
		Column:            "reserve_status",
		SQLAlias:          "reserve_status",
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

func NewReservePoliciesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReservePoliciesFilterFields[field]
	return
}

type ReservePoliciesFilterResult struct {
	ReservePolicies
	FilterCount int `db:"count"`
}

func ValidateReservePoliciesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReservePoliciesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReservePoliciesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReservePoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReservePoliciesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReservePoliciesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReservePoliciesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReservePoliciesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ReservePoliciesPolicyScope string

const (
	ReservePoliciesPolicyScopePlatformDefault ReservePoliciesPolicyScope = "platform_default"
	ReservePoliciesPolicyScopeMerchant        ReservePoliciesPolicyScope = "merchant"
	ReservePoliciesPolicyScopeCategory        ReservePoliciesPolicyScope = "category"
	ReservePoliciesPolicyScopeRiskTier        ReservePoliciesPolicyScope = "risk_tier"
)

type ReserveStatus string

const (
	ReserveStatusActive   ReserveStatus = "active"
	ReserveStatusInactive ReserveStatus = "inactive"
)

type ReservePolicies struct {
	Id              uuid.UUID                  `db:"id"`
	PolicyCode      string                     `db:"policy_code"`
	MerchantPartyId nuuid.NUUID                `db:"merchant_party_id"`
	PolicyScope     ReservePoliciesPolicyScope `db:"policy_scope"`
	ReserveStatus   ReserveStatus              `db:"reserve_status"`
	Description     null.String                `db:"description"`
	Metadata        json.RawMessage            `db:"metadata"`

	shared.MetaSignature
}
type ReservePoliciesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReservePolicies) ToReservePoliciesPrimaryID() ReservePoliciesPrimaryID {
	return ReservePoliciesPrimaryID{
		Id: d.Id,
	}
}

type ReservePoliciesList []*ReservePolicies

type ReservePoliciesFilterResultList []*ReservePoliciesFilterResult

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

type FeeProfilesDBFieldNameType string

type feeProfilesDBFieldName struct {
	Id              FeeProfilesDBFieldNameType
	ProfileCode     FeeProfilesDBFieldNameType
	OwnerPartyId    FeeProfilesDBFieldNameType
	ProfileScope    FeeProfilesDBFieldNameType
	EffectiveStatus FeeProfilesDBFieldNameType
	Description     FeeProfilesDBFieldNameType
	Metadata        FeeProfilesDBFieldNameType
	MetaCreatedAt   FeeProfilesDBFieldNameType
	MetaCreatedBy   FeeProfilesDBFieldNameType
	MetaUpdatedAt   FeeProfilesDBFieldNameType
	MetaUpdatedBy   FeeProfilesDBFieldNameType
	MetaDeletedAt   FeeProfilesDBFieldNameType
	MetaDeletedBy   FeeProfilesDBFieldNameType
}

var FeeProfilesDBFieldName = feeProfilesDBFieldName{
	Id:              "id",
	ProfileCode:     "profile_code",
	OwnerPartyId:    "owner_party_id",
	ProfileScope:    "profile_scope",
	EffectiveStatus: "effective_status",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewFeeProfilesDBFieldNameFromStr(field string) (dbField FeeProfilesDBFieldNameType, found bool) {
	switch field {

	case string(FeeProfilesDBFieldName.Id):
		return FeeProfilesDBFieldName.Id, true

	case string(FeeProfilesDBFieldName.ProfileCode):
		return FeeProfilesDBFieldName.ProfileCode, true

	case string(FeeProfilesDBFieldName.OwnerPartyId):
		return FeeProfilesDBFieldName.OwnerPartyId, true

	case string(FeeProfilesDBFieldName.ProfileScope):
		return FeeProfilesDBFieldName.ProfileScope, true

	case string(FeeProfilesDBFieldName.EffectiveStatus):
		return FeeProfilesDBFieldName.EffectiveStatus, true

	case string(FeeProfilesDBFieldName.Description):
		return FeeProfilesDBFieldName.Description, true

	case string(FeeProfilesDBFieldName.Metadata):
		return FeeProfilesDBFieldName.Metadata, true

	case string(FeeProfilesDBFieldName.MetaCreatedAt):
		return FeeProfilesDBFieldName.MetaCreatedAt, true

	case string(FeeProfilesDBFieldName.MetaCreatedBy):
		return FeeProfilesDBFieldName.MetaCreatedBy, true

	case string(FeeProfilesDBFieldName.MetaUpdatedAt):
		return FeeProfilesDBFieldName.MetaUpdatedAt, true

	case string(FeeProfilesDBFieldName.MetaUpdatedBy):
		return FeeProfilesDBFieldName.MetaUpdatedBy, true

	case string(FeeProfilesDBFieldName.MetaDeletedAt):
		return FeeProfilesDBFieldName.MetaDeletedAt, true

	case string(FeeProfilesDBFieldName.MetaDeletedBy):
		return FeeProfilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FeeProfilesFilterJoins = map[string]JoinSpec{}

var FeeProfilesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"profile_code": {
		SourcePath:        "profile_code",
		DefaultOutputPath: "profileCode",
		Column:            "profile_code",
		SQLAlias:          "profile_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"owner_party_id": {
		SourcePath:        "owner_party_id",
		DefaultOutputPath: "ownerPartyId",
		Column:            "owner_party_id",
		SQLAlias:          "owner_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"profile_scope": {
		SourcePath:        "profile_scope",
		DefaultOutputPath: "profileScope",
		Column:            "profile_scope",
		SQLAlias:          "profile_scope",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"effective_status": {
		SourcePath:        "effective_status",
		DefaultOutputPath: "effectiveStatus",
		Column:            "effective_status",
		SQLAlias:          "effective_status",
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

func NewFeeProfilesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FeeProfilesFilterFields[field]
	return
}

type FeeProfilesFilterResult struct {
	FeeProfiles
	FilterCount int `db:"count"`
}

func ValidateFeeProfilesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFeeProfilesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFeeProfilesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFeeProfilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFeeProfilesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFeeProfilesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFeeProfilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFeeProfilesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type EffectiveStatus string

const (
	EffectiveStatusActive   EffectiveStatus = "active"
	EffectiveStatusInactive EffectiveStatus = "inactive"
)

type ProfileScope string

const (
	ProfileScopePlatform ProfileScope = "platform"
	ProfileScopeMerchant ProfileScope = "merchant"
	ProfileScopeListing  ProfileScope = "listing"
	ProfileScopeCategory ProfileScope = "category"
	ProfileScopeCampaign ProfileScope = "campaign"
)

type FeeProfiles struct {
	Id              uuid.UUID       `db:"id"`
	ProfileCode     string          `db:"profile_code"`
	OwnerPartyId    nuuid.NUUID     `db:"owner_party_id"`
	ProfileScope    ProfileScope    `db:"profile_scope"`
	EffectiveStatus EffectiveStatus `db:"effective_status"`
	Description     null.String     `db:"description"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FeeProfilesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FeeProfiles) ToFeeProfilesPrimaryID() FeeProfilesPrimaryID {
	return FeeProfilesPrimaryID{
		Id: d.Id,
	}
}

type FeeProfilesList []*FeeProfiles

type FeeProfilesFilterResultList []*FeeProfilesFilterResult

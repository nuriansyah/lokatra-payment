package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type RoutingProfilesDBFieldNameType string

type routingProfilesDBFieldName struct {
	Id                RoutingProfilesDBFieldNameType
	MerchantId        RoutingProfilesDBFieldNameType
	Name              RoutingProfilesDBFieldNameType
	Strategy          RoutingProfilesDBFieldNameType
	IsActive          RoutingProfilesDBFieldNameType
	FallbackProfileId RoutingProfilesDBFieldNameType
	Notes             RoutingProfilesDBFieldNameType
	MetaCreatedAt     RoutingProfilesDBFieldNameType
	MetaCreatedBy     RoutingProfilesDBFieldNameType
	MetaUpdatedAt     RoutingProfilesDBFieldNameType
	MetaUpdatedBy     RoutingProfilesDBFieldNameType
	MetaDeletedAt     RoutingProfilesDBFieldNameType
	MetaDeletedBy     RoutingProfilesDBFieldNameType
}

var RoutingProfilesDBFieldName = routingProfilesDBFieldName{
	Id:                "id",
	MerchantId:        "merchant_id",
	Name:              "name",
	Strategy:          "strategy",
	IsActive:          "is_active",
	FallbackProfileId: "fallback_profile_id",
	Notes:             "notes",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewRoutingProfilesDBFieldNameFromStr(field string) (dbField RoutingProfilesDBFieldNameType, found bool) {
	switch field {

	case string(RoutingProfilesDBFieldName.Id):
		return RoutingProfilesDBFieldName.Id, true

	case string(RoutingProfilesDBFieldName.MerchantId):
		return RoutingProfilesDBFieldName.MerchantId, true

	case string(RoutingProfilesDBFieldName.Name):
		return RoutingProfilesDBFieldName.Name, true

	case string(RoutingProfilesDBFieldName.Strategy):
		return RoutingProfilesDBFieldName.Strategy, true

	case string(RoutingProfilesDBFieldName.IsActive):
		return RoutingProfilesDBFieldName.IsActive, true

	case string(RoutingProfilesDBFieldName.FallbackProfileId):
		return RoutingProfilesDBFieldName.FallbackProfileId, true

	case string(RoutingProfilesDBFieldName.Notes):
		return RoutingProfilesDBFieldName.Notes, true

	case string(RoutingProfilesDBFieldName.MetaCreatedAt):
		return RoutingProfilesDBFieldName.MetaCreatedAt, true

	case string(RoutingProfilesDBFieldName.MetaCreatedBy):
		return RoutingProfilesDBFieldName.MetaCreatedBy, true

	case string(RoutingProfilesDBFieldName.MetaUpdatedAt):
		return RoutingProfilesDBFieldName.MetaUpdatedAt, true

	case string(RoutingProfilesDBFieldName.MetaUpdatedBy):
		return RoutingProfilesDBFieldName.MetaUpdatedBy, true

	case string(RoutingProfilesDBFieldName.MetaDeletedAt):
		return RoutingProfilesDBFieldName.MetaDeletedAt, true

	case string(RoutingProfilesDBFieldName.MetaDeletedBy):
		return RoutingProfilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type RoutingProfilesFilterResult struct {
	RoutingProfiles
	FilterCount int `db:"count"`
}

func ValidateRoutingProfilesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewRoutingProfilesDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewRoutingProfilesDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewRoutingProfilesDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type RoutingProfiles struct {
	Id                uuid.UUID       `db:"id"`
	MerchantId        nuuid.NUUID     `db:"merchant_id"`
	Name              string          `db:"name"`
	Strategy          RoutingStrategy `db:"strategy"`
	IsActive          bool            `db:"is_active"`
	FallbackProfileId nuuid.NUUID     `db:"fallback_profile_id"`
	Notes             null.String     `db:"notes"`
	MetaCreatedAt     time.Time       `db:"meta_created_at"`
	MetaCreatedBy     uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt     time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy     nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt     null.Time       `db:"meta_deleted_at"`
	MetaDeletedBy     nuuid.NUUID     `db:"meta_deleted_by"`
}
type RoutingProfilesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RoutingProfiles) ToRoutingProfilesPrimaryID() RoutingProfilesPrimaryID {
	return RoutingProfilesPrimaryID{
		Id: d.Id,
	}
}

type RoutingProfilesList []*RoutingProfiles

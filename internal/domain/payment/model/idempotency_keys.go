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

type IdempotencyKeysDBFieldNameType string

type idempotencyKeysDBFieldName struct {
	Id             IdempotencyKeysDBFieldNameType
	Key            IdempotencyKeysDBFieldNameType
	ActorType      IdempotencyKeysDBFieldNameType
	ActorId        IdempotencyKeysDBFieldNameType
	RequestHash    IdempotencyKeysDBFieldNameType
	Status         IdempotencyKeysDBFieldNameType
	ResourceType   IdempotencyKeysDBFieldNameType
	ResourceId     IdempotencyKeysDBFieldNameType
	ResponseStatus IdempotencyKeysDBFieldNameType
	ResponseBody   IdempotencyKeysDBFieldNameType
	LockedUntil    IdempotencyKeysDBFieldNameType
	CompletedAt    IdempotencyKeysDBFieldNameType
	Metadata       IdempotencyKeysDBFieldNameType
	MetaCreatedAt  IdempotencyKeysDBFieldNameType
	MetaCreatedBy  IdempotencyKeysDBFieldNameType
	MetaUpdatedAt  IdempotencyKeysDBFieldNameType
	MetaUpdatedBy  IdempotencyKeysDBFieldNameType
	MetaDeletedAt  IdempotencyKeysDBFieldNameType
	MetaDeletedBy  IdempotencyKeysDBFieldNameType
}

var IdempotencyKeysDBFieldName = idempotencyKeysDBFieldName{
	Id:             "id",
	Key:            "key",
	ActorType:      "actor_type",
	ActorId:        "actor_id",
	RequestHash:    "request_hash",
	Status:         "status",
	ResourceType:   "resource_type",
	ResourceId:     "resource_id",
	ResponseStatus: "response_status",
	ResponseBody:   "response_body",
	LockedUntil:    "locked_until",
	CompletedAt:    "completed_at",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewIdempotencyKeysDBFieldNameFromStr(field string) (dbField IdempotencyKeysDBFieldNameType, found bool) {
	switch field {

	case string(IdempotencyKeysDBFieldName.Id):
		return IdempotencyKeysDBFieldName.Id, true

	case string(IdempotencyKeysDBFieldName.Key):
		return IdempotencyKeysDBFieldName.Key, true

	case string(IdempotencyKeysDBFieldName.ActorType):
		return IdempotencyKeysDBFieldName.ActorType, true

	case string(IdempotencyKeysDBFieldName.ActorId):
		return IdempotencyKeysDBFieldName.ActorId, true

	case string(IdempotencyKeysDBFieldName.RequestHash):
		return IdempotencyKeysDBFieldName.RequestHash, true

	case string(IdempotencyKeysDBFieldName.Status):
		return IdempotencyKeysDBFieldName.Status, true

	case string(IdempotencyKeysDBFieldName.ResourceType):
		return IdempotencyKeysDBFieldName.ResourceType, true

	case string(IdempotencyKeysDBFieldName.ResourceId):
		return IdempotencyKeysDBFieldName.ResourceId, true

	case string(IdempotencyKeysDBFieldName.ResponseStatus):
		return IdempotencyKeysDBFieldName.ResponseStatus, true

	case string(IdempotencyKeysDBFieldName.ResponseBody):
		return IdempotencyKeysDBFieldName.ResponseBody, true

	case string(IdempotencyKeysDBFieldName.LockedUntil):
		return IdempotencyKeysDBFieldName.LockedUntil, true

	case string(IdempotencyKeysDBFieldName.CompletedAt):
		return IdempotencyKeysDBFieldName.CompletedAt, true

	case string(IdempotencyKeysDBFieldName.Metadata):
		return IdempotencyKeysDBFieldName.Metadata, true

	case string(IdempotencyKeysDBFieldName.MetaCreatedAt):
		return IdempotencyKeysDBFieldName.MetaCreatedAt, true

	case string(IdempotencyKeysDBFieldName.MetaCreatedBy):
		return IdempotencyKeysDBFieldName.MetaCreatedBy, true

	case string(IdempotencyKeysDBFieldName.MetaUpdatedAt):
		return IdempotencyKeysDBFieldName.MetaUpdatedAt, true

	case string(IdempotencyKeysDBFieldName.MetaUpdatedBy):
		return IdempotencyKeysDBFieldName.MetaUpdatedBy, true

	case string(IdempotencyKeysDBFieldName.MetaDeletedAt):
		return IdempotencyKeysDBFieldName.MetaDeletedAt, true

	case string(IdempotencyKeysDBFieldName.MetaDeletedBy):
		return IdempotencyKeysDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var IdempotencyKeysFilterJoins = map[string]JoinSpec{}

var IdempotencyKeysFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"key": {
		SourcePath:        "key",
		DefaultOutputPath: "key",
		Column:            "key",
		SQLAlias:          "key",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"actor_type": {
		SourcePath:        "actor_type",
		DefaultOutputPath: "actorType",
		Column:            "actor_type",
		SQLAlias:          "actor_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"actor_id": {
		SourcePath:        "actor_id",
		DefaultOutputPath: "actorId",
		Column:            "actor_id",
		SQLAlias:          "actor_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"request_hash": {
		SourcePath:        "request_hash",
		DefaultOutputPath: "requestHash",
		Column:            "request_hash",
		SQLAlias:          "request_hash",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resource_type": {
		SourcePath:        "resource_type",
		DefaultOutputPath: "resourceType",
		Column:            "resource_type",
		SQLAlias:          "resource_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resource_id": {
		SourcePath:        "resource_id",
		DefaultOutputPath: "resourceId",
		Column:            "resource_id",
		SQLAlias:          "resource_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"response_status": {
		SourcePath:        "response_status",
		DefaultOutputPath: "responseStatus",
		Column:            "response_status",
		SQLAlias:          "response_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"response_body": {
		SourcePath:        "response_body",
		DefaultOutputPath: "responseBody",
		Column:            "response_body",
		SQLAlias:          "response_body",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"locked_until": {
		SourcePath:        "locked_until",
		DefaultOutputPath: "lockedUntil",
		Column:            "locked_until",
		SQLAlias:          "locked_until",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"completed_at": {
		SourcePath:        "completed_at",
		DefaultOutputPath: "completedAt",
		Column:            "completed_at",
		SQLAlias:          "completed_at",
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

func NewIdempotencyKeysFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = IdempotencyKeysFilterFields[field]
	return
}

type IdempotencyKeysFilterResult struct {
	IdempotencyKeys
	FilterCount int `db:"count"`
}

func ValidateIdempotencyKeysFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewIdempotencyKeysFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewIdempotencyKeysFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewIdempotencyKeysFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateIdempotencyKeysFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateIdempotencyKeysFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewIdempotencyKeysFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateIdempotencyKeysFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type IdempotencyStatus string

const (
	IdempotencyStatusProcessing IdempotencyStatus = "processing"
	IdempotencyStatusCompleted  IdempotencyStatus = "completed"
	IdempotencyStatusFailed     IdempotencyStatus = "failed"
)

type IdempotencyKeys struct {
	Id             uuid.UUID         `db:"id"`
	Key            string            `db:"key"`
	ActorType      null.String       `db:"actor_type"`
	ActorId        nuuid.NUUID       `db:"actor_id"`
	RequestHash    string            `db:"request_hash"`
	Status         IdempotencyStatus `db:"status"`
	ResourceType   null.String       `db:"resource_type"`
	ResourceId     nuuid.NUUID       `db:"resource_id"`
	ResponseStatus null.Int          `db:"response_status"`
	ResponseBody   json.RawMessage   `db:"response_body"`
	LockedUntil    null.Time         `db:"locked_until"`
	CompletedAt    null.Time         `db:"completed_at"`
	Metadata       json.RawMessage   `db:"metadata"`

	shared.MetaSignature
}
type IdempotencyKeysPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d IdempotencyKeys) ToIdempotencyKeysPrimaryID() IdempotencyKeysPrimaryID {
	return IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type IdempotencyKeysList []*IdempotencyKeys

type IdempotencyKeysFilterResultList []*IdempotencyKeysFilterResult

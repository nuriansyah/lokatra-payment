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

type FinanceIdempotencyKeysDBFieldNameType string

type financeIdempotencyKeysDBFieldName struct {
	Id             FinanceIdempotencyKeysDBFieldNameType
	Scope          FinanceIdempotencyKeysDBFieldNameType
	Operation      FinanceIdempotencyKeysDBFieldNameType
	IdempotencyKey FinanceIdempotencyKeysDBFieldNameType
	RequestHash    FinanceIdempotencyKeysDBFieldNameType
	ResourceType   FinanceIdempotencyKeysDBFieldNameType
	ResourceId     FinanceIdempotencyKeysDBFieldNameType
	ResponseStatus FinanceIdempotencyKeysDBFieldNameType
	ResponseBody   FinanceIdempotencyKeysDBFieldNameType
	LockedUntil    FinanceIdempotencyKeysDBFieldNameType
	CompletedAt    FinanceIdempotencyKeysDBFieldNameType
	Metadata       FinanceIdempotencyKeysDBFieldNameType
	MetaCreatedAt  FinanceIdempotencyKeysDBFieldNameType
	MetaCreatedBy  FinanceIdempotencyKeysDBFieldNameType
	MetaUpdatedAt  FinanceIdempotencyKeysDBFieldNameType
	MetaUpdatedBy  FinanceIdempotencyKeysDBFieldNameType
	MetaDeletedAt  FinanceIdempotencyKeysDBFieldNameType
	MetaDeletedBy  FinanceIdempotencyKeysDBFieldNameType
}

var FinanceIdempotencyKeysDBFieldName = financeIdempotencyKeysDBFieldName{
	Id:             "id",
	Scope:          "scope",
	Operation:      "operation",
	IdempotencyKey: "idempotency_key",
	RequestHash:    "request_hash",
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

func NewFinanceIdempotencyKeysDBFieldNameFromStr(field string) (dbField FinanceIdempotencyKeysDBFieldNameType, found bool) {
	switch field {

	case string(FinanceIdempotencyKeysDBFieldName.Id):
		return FinanceIdempotencyKeysDBFieldName.Id, true

	case string(FinanceIdempotencyKeysDBFieldName.Scope):
		return FinanceIdempotencyKeysDBFieldName.Scope, true

	case string(FinanceIdempotencyKeysDBFieldName.Operation):
		return FinanceIdempotencyKeysDBFieldName.Operation, true

	case string(FinanceIdempotencyKeysDBFieldName.IdempotencyKey):
		return FinanceIdempotencyKeysDBFieldName.IdempotencyKey, true

	case string(FinanceIdempotencyKeysDBFieldName.RequestHash):
		return FinanceIdempotencyKeysDBFieldName.RequestHash, true

	case string(FinanceIdempotencyKeysDBFieldName.ResourceType):
		return FinanceIdempotencyKeysDBFieldName.ResourceType, true

	case string(FinanceIdempotencyKeysDBFieldName.ResourceId):
		return FinanceIdempotencyKeysDBFieldName.ResourceId, true

	case string(FinanceIdempotencyKeysDBFieldName.ResponseStatus):
		return FinanceIdempotencyKeysDBFieldName.ResponseStatus, true

	case string(FinanceIdempotencyKeysDBFieldName.ResponseBody):
		return FinanceIdempotencyKeysDBFieldName.ResponseBody, true

	case string(FinanceIdempotencyKeysDBFieldName.LockedUntil):
		return FinanceIdempotencyKeysDBFieldName.LockedUntil, true

	case string(FinanceIdempotencyKeysDBFieldName.CompletedAt):
		return FinanceIdempotencyKeysDBFieldName.CompletedAt, true

	case string(FinanceIdempotencyKeysDBFieldName.Metadata):
		return FinanceIdempotencyKeysDBFieldName.Metadata, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaCreatedAt):
		return FinanceIdempotencyKeysDBFieldName.MetaCreatedAt, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaCreatedBy):
		return FinanceIdempotencyKeysDBFieldName.MetaCreatedBy, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaUpdatedAt):
		return FinanceIdempotencyKeysDBFieldName.MetaUpdatedAt, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaUpdatedBy):
		return FinanceIdempotencyKeysDBFieldName.MetaUpdatedBy, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaDeletedAt):
		return FinanceIdempotencyKeysDBFieldName.MetaDeletedAt, true

	case string(FinanceIdempotencyKeysDBFieldName.MetaDeletedBy):
		return FinanceIdempotencyKeysDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceIdempotencyKeysFilterJoins = map[string]JoinSpec{}

var FinanceIdempotencyKeysFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"scope": {
		SourcePath:        "scope",
		DefaultOutputPath: "scope",
		Column:            "scope",
		SQLAlias:          "scope",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"operation": {
		SourcePath:        "operation",
		DefaultOutputPath: "operation",
		Column:            "operation",
		SQLAlias:          "operation",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"idempotency_key": {
		SourcePath:        "idempotency_key",
		DefaultOutputPath: "idempotencyKey",
		Column:            "idempotency_key",
		SQLAlias:          "idempotency_key",
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

func NewFinanceIdempotencyKeysFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceIdempotencyKeysFilterFields[field]
	return
}

type FinanceIdempotencyKeysFilterResult struct {
	FinanceIdempotencyKeys
	FilterCount int `db:"count"`
}

func ValidateFinanceIdempotencyKeysFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceIdempotencyKeysFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceIdempotencyKeysFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceIdempotencyKeysFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceIdempotencyKeysFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceIdempotencyKeysFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceIdempotencyKeysFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceIdempotencyKeysFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FinanceIdempotencyKeys struct {
	Id             uuid.UUID       `db:"id"`
	Scope          string          `db:"scope"`
	Operation      string          `db:"operation"`
	IdempotencyKey string          `db:"idempotency_key"`
	RequestHash    string          `db:"request_hash"`
	ResourceType   null.String     `db:"resource_type"`
	ResourceId     nuuid.NUUID     `db:"resource_id"`
	ResponseStatus null.Int        `db:"response_status"`
	ResponseBody   json.RawMessage `db:"response_body"`
	LockedUntil    null.Time       `db:"locked_until"`
	CompletedAt    null.Time       `db:"completed_at"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceIdempotencyKeysPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceIdempotencyKeys) ToFinanceIdempotencyKeysPrimaryID() FinanceIdempotencyKeysPrimaryID {
	return FinanceIdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type FinanceIdempotencyKeysList []*FinanceIdempotencyKeys

type FinanceIdempotencyKeysFilterResultList []*FinanceIdempotencyKeysFilterResult

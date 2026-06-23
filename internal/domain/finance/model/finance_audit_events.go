package model

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"time"
)

type FinanceAuditEventsDBFieldNameType string

type financeAuditEventsDBFieldName struct {
	Id            FinanceAuditEventsDBFieldNameType
	AggregateType FinanceAuditEventsDBFieldNameType
	AggregateId   FinanceAuditEventsDBFieldNameType
	EventType     FinanceAuditEventsDBFieldNameType
	ActorUserId   FinanceAuditEventsDBFieldNameType
	ActorType     FinanceAuditEventsDBFieldNameType
	EventAt       FinanceAuditEventsDBFieldNameType
	OldState      FinanceAuditEventsDBFieldNameType
	NewState      FinanceAuditEventsDBFieldNameType
	CorrelationId FinanceAuditEventsDBFieldNameType
	CausationId   FinanceAuditEventsDBFieldNameType
	Metadata      FinanceAuditEventsDBFieldNameType
	MetaCreatedAt FinanceAuditEventsDBFieldNameType
	MetaCreatedBy FinanceAuditEventsDBFieldNameType
	MetaUpdatedAt FinanceAuditEventsDBFieldNameType
	MetaUpdatedBy FinanceAuditEventsDBFieldNameType
	MetaDeletedAt FinanceAuditEventsDBFieldNameType
	MetaDeletedBy FinanceAuditEventsDBFieldNameType
}

var FinanceAuditEventsDBFieldName = financeAuditEventsDBFieldName{
	Id:            "id",
	AggregateType: "aggregate_type",
	AggregateId:   "aggregate_id",
	EventType:     "event_type",
	ActorUserId:   "actor_user_id",
	ActorType:     "actor_type",
	EventAt:       "event_at",
	OldState:      "old_state",
	NewState:      "new_state",
	CorrelationId: "correlation_id",
	CausationId:   "causation_id",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewFinanceAuditEventsDBFieldNameFromStr(field string) (dbField FinanceAuditEventsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceAuditEventsDBFieldName.Id):
		return FinanceAuditEventsDBFieldName.Id, true

	case string(FinanceAuditEventsDBFieldName.AggregateType):
		return FinanceAuditEventsDBFieldName.AggregateType, true

	case string(FinanceAuditEventsDBFieldName.AggregateId):
		return FinanceAuditEventsDBFieldName.AggregateId, true

	case string(FinanceAuditEventsDBFieldName.EventType):
		return FinanceAuditEventsDBFieldName.EventType, true

	case string(FinanceAuditEventsDBFieldName.ActorUserId):
		return FinanceAuditEventsDBFieldName.ActorUserId, true

	case string(FinanceAuditEventsDBFieldName.ActorType):
		return FinanceAuditEventsDBFieldName.ActorType, true

	case string(FinanceAuditEventsDBFieldName.EventAt):
		return FinanceAuditEventsDBFieldName.EventAt, true

	case string(FinanceAuditEventsDBFieldName.OldState):
		return FinanceAuditEventsDBFieldName.OldState, true

	case string(FinanceAuditEventsDBFieldName.NewState):
		return FinanceAuditEventsDBFieldName.NewState, true

	case string(FinanceAuditEventsDBFieldName.CorrelationId):
		return FinanceAuditEventsDBFieldName.CorrelationId, true

	case string(FinanceAuditEventsDBFieldName.CausationId):
		return FinanceAuditEventsDBFieldName.CausationId, true

	case string(FinanceAuditEventsDBFieldName.Metadata):
		return FinanceAuditEventsDBFieldName.Metadata, true

	case string(FinanceAuditEventsDBFieldName.MetaCreatedAt):
		return FinanceAuditEventsDBFieldName.MetaCreatedAt, true

	case string(FinanceAuditEventsDBFieldName.MetaCreatedBy):
		return FinanceAuditEventsDBFieldName.MetaCreatedBy, true

	case string(FinanceAuditEventsDBFieldName.MetaUpdatedAt):
		return FinanceAuditEventsDBFieldName.MetaUpdatedAt, true

	case string(FinanceAuditEventsDBFieldName.MetaUpdatedBy):
		return FinanceAuditEventsDBFieldName.MetaUpdatedBy, true

	case string(FinanceAuditEventsDBFieldName.MetaDeletedAt):
		return FinanceAuditEventsDBFieldName.MetaDeletedAt, true

	case string(FinanceAuditEventsDBFieldName.MetaDeletedBy):
		return FinanceAuditEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceAuditEventsFilterJoins = map[string]JoinSpec{}

var FinanceAuditEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"aggregate_type": {
		SourcePath:        "aggregate_type",
		DefaultOutputPath: "aggregateType",
		Column:            "aggregate_type",
		SQLAlias:          "aggregate_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"aggregate_id": {
		SourcePath:        "aggregate_id",
		DefaultOutputPath: "aggregateId",
		Column:            "aggregate_id",
		SQLAlias:          "aggregate_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"event_type": {
		SourcePath:        "event_type",
		DefaultOutputPath: "eventType",
		Column:            "event_type",
		SQLAlias:          "event_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"actor_user_id": {
		SourcePath:        "actor_user_id",
		DefaultOutputPath: "actorUserId",
		Column:            "actor_user_id",
		SQLAlias:          "actor_user_id",
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
	"event_at": {
		SourcePath:        "event_at",
		DefaultOutputPath: "eventAt",
		Column:            "event_at",
		SQLAlias:          "event_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"old_state": {
		SourcePath:        "old_state",
		DefaultOutputPath: "oldState",
		Column:            "old_state",
		SQLAlias:          "old_state",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"new_state": {
		SourcePath:        "new_state",
		DefaultOutputPath: "newState",
		Column:            "new_state",
		SQLAlias:          "new_state",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"correlation_id": {
		SourcePath:        "correlation_id",
		DefaultOutputPath: "correlationId",
		Column:            "correlation_id",
		SQLAlias:          "correlation_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"causation_id": {
		SourcePath:        "causation_id",
		DefaultOutputPath: "causationId",
		Column:            "causation_id",
		SQLAlias:          "causation_id",
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

func NewFinanceAuditEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceAuditEventsFilterFields[field]
	return
}

type FinanceAuditEventsFilterResult struct {
	FinanceAuditEvents
	FilterCount int `db:"count"`
}

func ValidateFinanceAuditEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceAuditEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceAuditEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceAuditEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceAuditEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceAuditEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceAuditEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceAuditEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ActorType string

const (
	ActorTypeSystem   ActorType = "system"
	ActorTypeUser     ActorType = "user"
	ActorTypeWorker   ActorType = "worker"
	ActorTypeProvider ActorType = "provider"
)

type FinanceAuditEvents struct {
	Id            uuid.UUID       `db:"id"`
	AggregateType string          `db:"aggregate_type"`
	AggregateId   uuid.UUID       `db:"aggregate_id"`
	EventType     string          `db:"event_type"`
	ActorUserId   nuuid.NUUID     `db:"actor_user_id"`
	ActorType     ActorType       `db:"actor_type"`
	EventAt       time.Time       `db:"event_at"`
	OldState      json.RawMessage `db:"old_state"`
	NewState      json.RawMessage `db:"new_state"`
	CorrelationId null.String     `db:"correlation_id"`
	CausationId   null.String     `db:"causation_id"`
	Metadata      json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceAuditEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceAuditEvents) ToFinanceAuditEventsPrimaryID() FinanceAuditEventsPrimaryID {
	return FinanceAuditEventsPrimaryID{
		Id: d.Id,
	}
}

type FinanceAuditEventsList []*FinanceAuditEvents

type FinanceAuditEventsFilterResultList []*FinanceAuditEventsFilterResult

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

type FinanceOutboxEventsDBFieldNameType string

type financeOutboxEventsDBFieldName struct {
	Id              FinanceOutboxEventsDBFieldNameType
	AggregateType   FinanceOutboxEventsDBFieldNameType
	AggregateId     FinanceOutboxEventsDBFieldNameType
	EventType       FinanceOutboxEventsDBFieldNameType
	EventVersion    FinanceOutboxEventsDBFieldNameType
	IdempotencyKey  FinanceOutboxEventsDBFieldNameType
	CorrelationId   FinanceOutboxEventsDBFieldNameType
	Payload         FinanceOutboxEventsDBFieldNameType
	Headers         FinanceOutboxEventsDBFieldNameType
	PublishStatus   FinanceOutboxEventsDBFieldNameType
	AttemptCount    FinanceOutboxEventsDBFieldNameType
	NextAttemptAt   FinanceOutboxEventsDBFieldNameType
	PublishedAt     FinanceOutboxEventsDBFieldNameType
	LastErrorCode   FinanceOutboxEventsDBFieldNameType
	LastErrorDetail FinanceOutboxEventsDBFieldNameType
	Metadata        FinanceOutboxEventsDBFieldNameType
	MetaCreatedAt   FinanceOutboxEventsDBFieldNameType
	MetaCreatedBy   FinanceOutboxEventsDBFieldNameType
	MetaUpdatedAt   FinanceOutboxEventsDBFieldNameType
	MetaUpdatedBy   FinanceOutboxEventsDBFieldNameType
	MetaDeletedAt   FinanceOutboxEventsDBFieldNameType
	MetaDeletedBy   FinanceOutboxEventsDBFieldNameType
}

var FinanceOutboxEventsDBFieldName = financeOutboxEventsDBFieldName{
	Id:              "id",
	AggregateType:   "aggregate_type",
	AggregateId:     "aggregate_id",
	EventType:       "event_type",
	EventVersion:    "event_version",
	IdempotencyKey:  "idempotency_key",
	CorrelationId:   "correlation_id",
	Payload:         "payload",
	Headers:         "headers",
	PublishStatus:   "publish_status",
	AttemptCount:    "attempt_count",
	NextAttemptAt:   "next_attempt_at",
	PublishedAt:     "published_at",
	LastErrorCode:   "last_error_code",
	LastErrorDetail: "last_error_detail",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewFinanceOutboxEventsDBFieldNameFromStr(field string) (dbField FinanceOutboxEventsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceOutboxEventsDBFieldName.Id):
		return FinanceOutboxEventsDBFieldName.Id, true

	case string(FinanceOutboxEventsDBFieldName.AggregateType):
		return FinanceOutboxEventsDBFieldName.AggregateType, true

	case string(FinanceOutboxEventsDBFieldName.AggregateId):
		return FinanceOutboxEventsDBFieldName.AggregateId, true

	case string(FinanceOutboxEventsDBFieldName.EventType):
		return FinanceOutboxEventsDBFieldName.EventType, true

	case string(FinanceOutboxEventsDBFieldName.EventVersion):
		return FinanceOutboxEventsDBFieldName.EventVersion, true

	case string(FinanceOutboxEventsDBFieldName.IdempotencyKey):
		return FinanceOutboxEventsDBFieldName.IdempotencyKey, true

	case string(FinanceOutboxEventsDBFieldName.CorrelationId):
		return FinanceOutboxEventsDBFieldName.CorrelationId, true

	case string(FinanceOutboxEventsDBFieldName.Payload):
		return FinanceOutboxEventsDBFieldName.Payload, true

	case string(FinanceOutboxEventsDBFieldName.Headers):
		return FinanceOutboxEventsDBFieldName.Headers, true

	case string(FinanceOutboxEventsDBFieldName.PublishStatus):
		return FinanceOutboxEventsDBFieldName.PublishStatus, true

	case string(FinanceOutboxEventsDBFieldName.AttemptCount):
		return FinanceOutboxEventsDBFieldName.AttemptCount, true

	case string(FinanceOutboxEventsDBFieldName.NextAttemptAt):
		return FinanceOutboxEventsDBFieldName.NextAttemptAt, true

	case string(FinanceOutboxEventsDBFieldName.PublishedAt):
		return FinanceOutboxEventsDBFieldName.PublishedAt, true

	case string(FinanceOutboxEventsDBFieldName.LastErrorCode):
		return FinanceOutboxEventsDBFieldName.LastErrorCode, true

	case string(FinanceOutboxEventsDBFieldName.LastErrorDetail):
		return FinanceOutboxEventsDBFieldName.LastErrorDetail, true

	case string(FinanceOutboxEventsDBFieldName.Metadata):
		return FinanceOutboxEventsDBFieldName.Metadata, true

	case string(FinanceOutboxEventsDBFieldName.MetaCreatedAt):
		return FinanceOutboxEventsDBFieldName.MetaCreatedAt, true

	case string(FinanceOutboxEventsDBFieldName.MetaCreatedBy):
		return FinanceOutboxEventsDBFieldName.MetaCreatedBy, true

	case string(FinanceOutboxEventsDBFieldName.MetaUpdatedAt):
		return FinanceOutboxEventsDBFieldName.MetaUpdatedAt, true

	case string(FinanceOutboxEventsDBFieldName.MetaUpdatedBy):
		return FinanceOutboxEventsDBFieldName.MetaUpdatedBy, true

	case string(FinanceOutboxEventsDBFieldName.MetaDeletedAt):
		return FinanceOutboxEventsDBFieldName.MetaDeletedAt, true

	case string(FinanceOutboxEventsDBFieldName.MetaDeletedBy):
		return FinanceOutboxEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceOutboxEventsFilterJoins = map[string]JoinSpec{}

var FinanceOutboxEventsFilterFields = map[string]FilterFieldSpec{
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
	"event_version": {
		SourcePath:        "event_version",
		DefaultOutputPath: "eventVersion",
		Column:            "event_version",
		SQLAlias:          "event_version",
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
	"correlation_id": {
		SourcePath:        "correlation_id",
		DefaultOutputPath: "correlationId",
		Column:            "correlation_id",
		SQLAlias:          "correlation_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payload": {
		SourcePath:        "payload",
		DefaultOutputPath: "payload",
		Column:            "payload",
		SQLAlias:          "payload",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"headers": {
		SourcePath:        "headers",
		DefaultOutputPath: "headers",
		Column:            "headers",
		SQLAlias:          "headers",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"publish_status": {
		SourcePath:        "publish_status",
		DefaultOutputPath: "publishStatus",
		Column:            "publish_status",
		SQLAlias:          "publish_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"attempt_count": {
		SourcePath:        "attempt_count",
		DefaultOutputPath: "attemptCount",
		Column:            "attempt_count",
		SQLAlias:          "attempt_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"next_attempt_at": {
		SourcePath:        "next_attempt_at",
		DefaultOutputPath: "nextAttemptAt",
		Column:            "next_attempt_at",
		SQLAlias:          "next_attempt_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"published_at": {
		SourcePath:        "published_at",
		DefaultOutputPath: "publishedAt",
		Column:            "published_at",
		SQLAlias:          "published_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"last_error_code": {
		SourcePath:        "last_error_code",
		DefaultOutputPath: "lastErrorCode",
		Column:            "last_error_code",
		SQLAlias:          "last_error_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"last_error_detail": {
		SourcePath:        "last_error_detail",
		DefaultOutputPath: "lastErrorDetail",
		Column:            "last_error_detail",
		SQLAlias:          "last_error_detail",
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

func NewFinanceOutboxEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceOutboxEventsFilterFields[field]
	return
}

type FinanceOutboxEventsFilterResult struct {
	FinanceOutboxEvents
	FilterCount int `db:"count"`
}

func ValidateFinanceOutboxEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceOutboxEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceOutboxEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceOutboxEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceOutboxEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceOutboxEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceOutboxEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceOutboxEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PublishStatus string

const (
	PublishStatusPending    PublishStatus = "pending"
	PublishStatusPublishing PublishStatus = "publishing"
	PublishStatusPublished  PublishStatus = "published"
	PublishStatusFailed     PublishStatus = "failed"
	PublishStatusDead       PublishStatus = "dead"
)

type FinanceOutboxEvents struct {
	Id              uuid.UUID       `db:"id"`
	AggregateType   string          `db:"aggregate_type"`
	AggregateId     uuid.UUID       `db:"aggregate_id"`
	EventType       string          `db:"event_type"`
	EventVersion    int             `db:"event_version"`
	IdempotencyKey  string          `db:"idempotency_key"`
	CorrelationId   nuuid.NUUID     `db:"correlation_id"`
	Payload         json.RawMessage `db:"payload"`
	Headers         json.RawMessage `db:"headers"`
	PublishStatus   PublishStatus   `db:"publish_status"`
	AttemptCount    int             `db:"attempt_count"`
	NextAttemptAt   null.Time       `db:"next_attempt_at"`
	PublishedAt     null.Time       `db:"published_at"`
	LastErrorCode   null.String     `db:"last_error_code"`
	LastErrorDetail null.String     `db:"last_error_detail"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceOutboxEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceOutboxEvents) ToFinanceOutboxEventsPrimaryID() FinanceOutboxEventsPrimaryID {
	return FinanceOutboxEventsPrimaryID{
		Id: d.Id,
	}
}

type FinanceOutboxEventsList []*FinanceOutboxEvents

type FinanceOutboxEventsFilterResultList []*FinanceOutboxEventsFilterResult

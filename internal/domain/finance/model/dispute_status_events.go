package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type DisputeStatusEventsDBFieldNameType string

type disputeStatusEventsDBFieldName struct {
	Id             DisputeStatusEventsDBFieldNameType
	DisputeId      DisputeStatusEventsDBFieldNameType
	PreviousStatus DisputeStatusEventsDBFieldNameType
	NextStatus     DisputeStatusEventsDBFieldNameType
	ReasonCode     DisputeStatusEventsDBFieldNameType
	ActorId        DisputeStatusEventsDBFieldNameType
	OccurredAt     DisputeStatusEventsDBFieldNameType
	Metadata       DisputeStatusEventsDBFieldNameType
	MetaCreatedAt  DisputeStatusEventsDBFieldNameType
	MetaCreatedBy  DisputeStatusEventsDBFieldNameType
	MetaUpdatedAt  DisputeStatusEventsDBFieldNameType
	MetaUpdatedBy  DisputeStatusEventsDBFieldNameType
	MetaDeletedAt  DisputeStatusEventsDBFieldNameType
	MetaDeletedBy  DisputeStatusEventsDBFieldNameType
}

var DisputeStatusEventsDBFieldName = disputeStatusEventsDBFieldName{
	Id:             "id",
	DisputeId:      "dispute_id",
	PreviousStatus: "previous_status",
	NextStatus:     "next_status",
	ReasonCode:     "reason_code",
	ActorId:        "actor_id",
	OccurredAt:     "occurred_at",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewDisputeStatusEventsDBFieldNameFromStr(field string) (dbField DisputeStatusEventsDBFieldNameType, found bool) {
	switch field {

	case string(DisputeStatusEventsDBFieldName.Id):
		return DisputeStatusEventsDBFieldName.Id, true

	case string(DisputeStatusEventsDBFieldName.DisputeId):
		return DisputeStatusEventsDBFieldName.DisputeId, true

	case string(DisputeStatusEventsDBFieldName.PreviousStatus):
		return DisputeStatusEventsDBFieldName.PreviousStatus, true

	case string(DisputeStatusEventsDBFieldName.NextStatus):
		return DisputeStatusEventsDBFieldName.NextStatus, true

	case string(DisputeStatusEventsDBFieldName.ReasonCode):
		return DisputeStatusEventsDBFieldName.ReasonCode, true

	case string(DisputeStatusEventsDBFieldName.ActorId):
		return DisputeStatusEventsDBFieldName.ActorId, true

	case string(DisputeStatusEventsDBFieldName.OccurredAt):
		return DisputeStatusEventsDBFieldName.OccurredAt, true

	case string(DisputeStatusEventsDBFieldName.Metadata):
		return DisputeStatusEventsDBFieldName.Metadata, true

	case string(DisputeStatusEventsDBFieldName.MetaCreatedAt):
		return DisputeStatusEventsDBFieldName.MetaCreatedAt, true

	case string(DisputeStatusEventsDBFieldName.MetaCreatedBy):
		return DisputeStatusEventsDBFieldName.MetaCreatedBy, true

	case string(DisputeStatusEventsDBFieldName.MetaUpdatedAt):
		return DisputeStatusEventsDBFieldName.MetaUpdatedAt, true

	case string(DisputeStatusEventsDBFieldName.MetaUpdatedBy):
		return DisputeStatusEventsDBFieldName.MetaUpdatedBy, true

	case string(DisputeStatusEventsDBFieldName.MetaDeletedAt):
		return DisputeStatusEventsDBFieldName.MetaDeletedAt, true

	case string(DisputeStatusEventsDBFieldName.MetaDeletedBy):
		return DisputeStatusEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var DisputeStatusEventsFilterJoins = map[string]JoinSpec{}

var DisputeStatusEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dispute_id": {
		SourcePath:        "dispute_id",
		DefaultOutputPath: "disputeId",
		Column:            "dispute_id",
		SQLAlias:          "dispute_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"previous_status": {
		SourcePath:        "previous_status",
		DefaultOutputPath: "previousStatus",
		Column:            "previous_status",
		SQLAlias:          "previous_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"next_status": {
		SourcePath:        "next_status",
		DefaultOutputPath: "nextStatus",
		Column:            "next_status",
		SQLAlias:          "next_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason_code": {
		SourcePath:        "reason_code",
		DefaultOutputPath: "reasonCode",
		Column:            "reason_code",
		SQLAlias:          "reason_code",
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
	"occurred_at": {
		SourcePath:        "occurred_at",
		DefaultOutputPath: "occurredAt",
		Column:            "occurred_at",
		SQLAlias:          "occurred_at",
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

func NewDisputeStatusEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = DisputeStatusEventsFilterFields[field]
	return
}

type DisputeStatusEventsFilterResult struct {
	DisputeStatusEvents
	FilterCount int `db:"count"`
}

func ValidateDisputeStatusEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewDisputeStatusEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewDisputeStatusEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewDisputeStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateDisputeStatusEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateDisputeStatusEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewDisputeStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateDisputeStatusEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type DisputeStatusEvents struct {
	Id             uuid.UUID       `db:"id"`
	DisputeId      uuid.UUID       `db:"dispute_id"`
	PreviousStatus null.String     `db:"previous_status"`
	NextStatus     string          `db:"next_status"`
	ReasonCode     string          `db:"reason_code"`
	ActorId        uuid.UUID       `db:"actor_id"`
	OccurredAt     time.Time       `db:"occurred_at"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type DisputeStatusEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d DisputeStatusEvents) ToDisputeStatusEventsPrimaryID() DisputeStatusEventsPrimaryID {
	return DisputeStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type DisputeStatusEventsList []*DisputeStatusEvents

type DisputeStatusEventsFilterResultList []*DisputeStatusEventsFilterResult

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

type RefundStatusEventsDBFieldNameType string

type refundStatusEventsDBFieldName struct {
	Id            RefundStatusEventsDBFieldNameType
	RefundId      RefundStatusEventsDBFieldNameType
	FromStatus    RefundStatusEventsDBFieldNameType
	ToStatus      RefundStatusEventsDBFieldNameType
	ReasonCode    RefundStatusEventsDBFieldNameType
	ProviderRef   RefundStatusEventsDBFieldNameType
	EventPayload  RefundStatusEventsDBFieldNameType
	OccurredAt    RefundStatusEventsDBFieldNameType
	MetaCreatedAt RefundStatusEventsDBFieldNameType
	MetaCreatedBy RefundStatusEventsDBFieldNameType
	MetaUpdatedAt RefundStatusEventsDBFieldNameType
	MetaUpdatedBy RefundStatusEventsDBFieldNameType
	MetaDeletedAt RefundStatusEventsDBFieldNameType
	MetaDeletedBy RefundStatusEventsDBFieldNameType
}

var RefundStatusEventsDBFieldName = refundStatusEventsDBFieldName{
	Id:            "id",
	RefundId:      "refund_id",
	FromStatus:    "from_status",
	ToStatus:      "to_status",
	ReasonCode:    "reason_code",
	ProviderRef:   "provider_ref",
	EventPayload:  "event_payload",
	OccurredAt:    "occurred_at",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewRefundStatusEventsDBFieldNameFromStr(field string) (dbField RefundStatusEventsDBFieldNameType, found bool) {
	switch field {

	case string(RefundStatusEventsDBFieldName.Id):
		return RefundStatusEventsDBFieldName.Id, true

	case string(RefundStatusEventsDBFieldName.RefundId):
		return RefundStatusEventsDBFieldName.RefundId, true

	case string(RefundStatusEventsDBFieldName.FromStatus):
		return RefundStatusEventsDBFieldName.FromStatus, true

	case string(RefundStatusEventsDBFieldName.ToStatus):
		return RefundStatusEventsDBFieldName.ToStatus, true

	case string(RefundStatusEventsDBFieldName.ReasonCode):
		return RefundStatusEventsDBFieldName.ReasonCode, true

	case string(RefundStatusEventsDBFieldName.ProviderRef):
		return RefundStatusEventsDBFieldName.ProviderRef, true

	case string(RefundStatusEventsDBFieldName.EventPayload):
		return RefundStatusEventsDBFieldName.EventPayload, true

	case string(RefundStatusEventsDBFieldName.OccurredAt):
		return RefundStatusEventsDBFieldName.OccurredAt, true

	case string(RefundStatusEventsDBFieldName.MetaCreatedAt):
		return RefundStatusEventsDBFieldName.MetaCreatedAt, true

	case string(RefundStatusEventsDBFieldName.MetaCreatedBy):
		return RefundStatusEventsDBFieldName.MetaCreatedBy, true

	case string(RefundStatusEventsDBFieldName.MetaUpdatedAt):
		return RefundStatusEventsDBFieldName.MetaUpdatedAt, true

	case string(RefundStatusEventsDBFieldName.MetaUpdatedBy):
		return RefundStatusEventsDBFieldName.MetaUpdatedBy, true

	case string(RefundStatusEventsDBFieldName.MetaDeletedAt):
		return RefundStatusEventsDBFieldName.MetaDeletedAt, true

	case string(RefundStatusEventsDBFieldName.MetaDeletedBy):
		return RefundStatusEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundStatusEventsFilterJoins = map[string]JoinSpec{}

var RefundStatusEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_id": {
		SourcePath:        "refund_id",
		DefaultOutputPath: "refundId",
		Column:            "refund_id",
		SQLAlias:          "refund_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"from_status": {
		SourcePath:        "from_status",
		DefaultOutputPath: "fromStatus",
		Column:            "from_status",
		SQLAlias:          "from_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"to_status": {
		SourcePath:        "to_status",
		DefaultOutputPath: "toStatus",
		Column:            "to_status",
		SQLAlias:          "to_status",
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
	"provider_ref": {
		SourcePath:        "provider_ref",
		DefaultOutputPath: "providerRef",
		Column:            "provider_ref",
		SQLAlias:          "provider_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"event_payload": {
		SourcePath:        "event_payload",
		DefaultOutputPath: "eventPayload",
		Column:            "event_payload",
		SQLAlias:          "event_payload",
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

func NewRefundStatusEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundStatusEventsFilterFields[field]
	return
}

type RefundStatusEventsFilterResult struct {
	RefundStatusEvents
	FilterCount int `db:"count"`
}

func ValidateRefundStatusEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundStatusEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundStatusEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundStatusEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundStatusEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundStatusEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RefundStatusEvents struct {
	Id           uuid.UUID       `db:"id"`
	RefundId     uuid.UUID       `db:"refund_id"`
	FromStatus   null.String     `db:"from_status"`
	ToStatus     string          `db:"to_status"`
	ReasonCode   null.String     `db:"reason_code"`
	ProviderRef  null.String     `db:"provider_ref"`
	EventPayload json.RawMessage `db:"event_payload"`
	OccurredAt   time.Time       `db:"occurred_at"`

	shared.MetaSignature
}
type RefundStatusEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundStatusEvents) ToRefundStatusEventsPrimaryID() RefundStatusEventsPrimaryID {
	return RefundStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type RefundStatusEventsList []*RefundStatusEvents

type RefundStatusEventsFilterResultList []*RefundStatusEventsFilterResult

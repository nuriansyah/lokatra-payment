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

type PayoutStatusEventsDBFieldNameType string

type payoutStatusEventsDBFieldName struct {
	Id            PayoutStatusEventsDBFieldNameType
	PayoutId      PayoutStatusEventsDBFieldNameType
	FromStatus    PayoutStatusEventsDBFieldNameType
	ToStatus      PayoutStatusEventsDBFieldNameType
	ReasonCode    PayoutStatusEventsDBFieldNameType
	ProviderRef   PayoutStatusEventsDBFieldNameType
	EventPayload  PayoutStatusEventsDBFieldNameType
	OccurredAt    PayoutStatusEventsDBFieldNameType
	MetaCreatedAt PayoutStatusEventsDBFieldNameType
	MetaCreatedBy PayoutStatusEventsDBFieldNameType
	MetaUpdatedAt PayoutStatusEventsDBFieldNameType
	MetaUpdatedBy PayoutStatusEventsDBFieldNameType
	MetaDeletedAt PayoutStatusEventsDBFieldNameType
	MetaDeletedBy PayoutStatusEventsDBFieldNameType
}

var PayoutStatusEventsDBFieldName = payoutStatusEventsDBFieldName{
	Id:            "id",
	PayoutId:      "payout_id",
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

func NewPayoutStatusEventsDBFieldNameFromStr(field string) (dbField PayoutStatusEventsDBFieldNameType, found bool) {
	switch field {

	case string(PayoutStatusEventsDBFieldName.Id):
		return PayoutStatusEventsDBFieldName.Id, true

	case string(PayoutStatusEventsDBFieldName.PayoutId):
		return PayoutStatusEventsDBFieldName.PayoutId, true

	case string(PayoutStatusEventsDBFieldName.FromStatus):
		return PayoutStatusEventsDBFieldName.FromStatus, true

	case string(PayoutStatusEventsDBFieldName.ToStatus):
		return PayoutStatusEventsDBFieldName.ToStatus, true

	case string(PayoutStatusEventsDBFieldName.ReasonCode):
		return PayoutStatusEventsDBFieldName.ReasonCode, true

	case string(PayoutStatusEventsDBFieldName.ProviderRef):
		return PayoutStatusEventsDBFieldName.ProviderRef, true

	case string(PayoutStatusEventsDBFieldName.EventPayload):
		return PayoutStatusEventsDBFieldName.EventPayload, true

	case string(PayoutStatusEventsDBFieldName.OccurredAt):
		return PayoutStatusEventsDBFieldName.OccurredAt, true

	case string(PayoutStatusEventsDBFieldName.MetaCreatedAt):
		return PayoutStatusEventsDBFieldName.MetaCreatedAt, true

	case string(PayoutStatusEventsDBFieldName.MetaCreatedBy):
		return PayoutStatusEventsDBFieldName.MetaCreatedBy, true

	case string(PayoutStatusEventsDBFieldName.MetaUpdatedAt):
		return PayoutStatusEventsDBFieldName.MetaUpdatedAt, true

	case string(PayoutStatusEventsDBFieldName.MetaUpdatedBy):
		return PayoutStatusEventsDBFieldName.MetaUpdatedBy, true

	case string(PayoutStatusEventsDBFieldName.MetaDeletedAt):
		return PayoutStatusEventsDBFieldName.MetaDeletedAt, true

	case string(PayoutStatusEventsDBFieldName.MetaDeletedBy):
		return PayoutStatusEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutStatusEventsFilterJoins = map[string]JoinSpec{}

var PayoutStatusEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_id": {
		SourcePath:        "payout_id",
		DefaultOutputPath: "payoutId",
		Column:            "payout_id",
		SQLAlias:          "payout_id",
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

func NewPayoutStatusEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutStatusEventsFilterFields[field]
	return
}

type PayoutStatusEventsFilterResult struct {
	PayoutStatusEvents
	FilterCount int `db:"count"`
}

func ValidatePayoutStatusEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutStatusEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutStatusEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutStatusEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutStatusEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutStatusEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutStatusEvents struct {
	Id           uuid.UUID       `db:"id"`
	PayoutId     uuid.UUID       `db:"payout_id"`
	FromStatus   null.String     `db:"from_status"`
	ToStatus     string          `db:"to_status"`
	ReasonCode   null.String     `db:"reason_code"`
	ProviderRef  null.String     `db:"provider_ref"`
	EventPayload json.RawMessage `db:"event_payload"`
	OccurredAt   time.Time       `db:"occurred_at"`

	shared.MetaSignature
}
type PayoutStatusEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayoutStatusEvents) ToPayoutStatusEventsPrimaryID() PayoutStatusEventsPrimaryID {
	return PayoutStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type PayoutStatusEventsList []*PayoutStatusEvents

type PayoutStatusEventsFilterResultList []*PayoutStatusEventsFilterResult

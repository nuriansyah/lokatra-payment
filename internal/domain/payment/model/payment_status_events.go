package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PaymentStatusEventsDBFieldNameType string

type paymentStatusEventsDBFieldName struct {
	Id                     PaymentStatusEventsDBFieldNameType
	PaymentIntentId        PaymentStatusEventsDBFieldNameType
	PaymentAttemptId       PaymentStatusEventsDBFieldNameType
	ProviderWebhookEventId PaymentStatusEventsDBFieldNameType
	SourceType             PaymentStatusEventsDBFieldNameType
	EventType              PaymentStatusEventsDBFieldNameType
	OldIntentStatus        PaymentStatusEventsDBFieldNameType
	NewIntentStatus        PaymentStatusEventsDBFieldNameType
	OldAttemptStatus       PaymentStatusEventsDBFieldNameType
	NewAttemptStatus       PaymentStatusEventsDBFieldNameType
	ProviderStatus         PaymentStatusEventsDBFieldNameType
	Reason                 PaymentStatusEventsDBFieldNameType
	OccurredAt             PaymentStatusEventsDBFieldNameType
	Metadata               PaymentStatusEventsDBFieldNameType
	MetaCreatedAt          PaymentStatusEventsDBFieldNameType
	MetaCreatedBy          PaymentStatusEventsDBFieldNameType
	MetaUpdatedAt          PaymentStatusEventsDBFieldNameType
	MetaUpdatedBy          PaymentStatusEventsDBFieldNameType
	MetaDeletedAt          PaymentStatusEventsDBFieldNameType
	MetaDeletedBy          PaymentStatusEventsDBFieldNameType
}

var PaymentStatusEventsDBFieldName = paymentStatusEventsDBFieldName{
	Id:                     "id",
	PaymentIntentId:        "payment_intent_id",
	PaymentAttemptId:       "payment_attempt_id",
	ProviderWebhookEventId: "provider_webhook_event_id",
	SourceType:             "source_type",
	EventType:              "event_type",
	OldIntentStatus:        "old_intent_status",
	NewIntentStatus:        "new_intent_status",
	OldAttemptStatus:       "old_attempt_status",
	NewAttemptStatus:       "new_attempt_status",
	ProviderStatus:         "provider_status",
	Reason:                 "reason",
	OccurredAt:             "occurred_at",
	Metadata:               "metadata",
	MetaCreatedAt:          "meta_created_at",
	MetaCreatedBy:          "meta_created_by",
	MetaUpdatedAt:          "meta_updated_at",
	MetaUpdatedBy:          "meta_updated_by",
	MetaDeletedAt:          "meta_deleted_at",
	MetaDeletedBy:          "meta_deleted_by",
}

func NewPaymentStatusEventsDBFieldNameFromStr(field string) (dbField PaymentStatusEventsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentStatusEventsDBFieldName.Id):
		return PaymentStatusEventsDBFieldName.Id, true

	case string(PaymentStatusEventsDBFieldName.PaymentIntentId):
		return PaymentStatusEventsDBFieldName.PaymentIntentId, true

	case string(PaymentStatusEventsDBFieldName.PaymentAttemptId):
		return PaymentStatusEventsDBFieldName.PaymentAttemptId, true

	case string(PaymentStatusEventsDBFieldName.ProviderWebhookEventId):
		return PaymentStatusEventsDBFieldName.ProviderWebhookEventId, true

	case string(PaymentStatusEventsDBFieldName.SourceType):
		return PaymentStatusEventsDBFieldName.SourceType, true

	case string(PaymentStatusEventsDBFieldName.EventType):
		return PaymentStatusEventsDBFieldName.EventType, true

	case string(PaymentStatusEventsDBFieldName.OldIntentStatus):
		return PaymentStatusEventsDBFieldName.OldIntentStatus, true

	case string(PaymentStatusEventsDBFieldName.NewIntentStatus):
		return PaymentStatusEventsDBFieldName.NewIntentStatus, true

	case string(PaymentStatusEventsDBFieldName.OldAttemptStatus):
		return PaymentStatusEventsDBFieldName.OldAttemptStatus, true

	case string(PaymentStatusEventsDBFieldName.NewAttemptStatus):
		return PaymentStatusEventsDBFieldName.NewAttemptStatus, true

	case string(PaymentStatusEventsDBFieldName.ProviderStatus):
		return PaymentStatusEventsDBFieldName.ProviderStatus, true

	case string(PaymentStatusEventsDBFieldName.Reason):
		return PaymentStatusEventsDBFieldName.Reason, true

	case string(PaymentStatusEventsDBFieldName.OccurredAt):
		return PaymentStatusEventsDBFieldName.OccurredAt, true

	case string(PaymentStatusEventsDBFieldName.Metadata):
		return PaymentStatusEventsDBFieldName.Metadata, true

	case string(PaymentStatusEventsDBFieldName.MetaCreatedAt):
		return PaymentStatusEventsDBFieldName.MetaCreatedAt, true

	case string(PaymentStatusEventsDBFieldName.MetaCreatedBy):
		return PaymentStatusEventsDBFieldName.MetaCreatedBy, true

	case string(PaymentStatusEventsDBFieldName.MetaUpdatedAt):
		return PaymentStatusEventsDBFieldName.MetaUpdatedAt, true

	case string(PaymentStatusEventsDBFieldName.MetaUpdatedBy):
		return PaymentStatusEventsDBFieldName.MetaUpdatedBy, true

	case string(PaymentStatusEventsDBFieldName.MetaDeletedAt):
		return PaymentStatusEventsDBFieldName.MetaDeletedAt, true

	case string(PaymentStatusEventsDBFieldName.MetaDeletedBy):
		return PaymentStatusEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentStatusEventsFilterJoins = map[string]JoinSpec{}

var PaymentStatusEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_intent_id": {
		SourcePath:        "payment_intent_id",
		DefaultOutputPath: "paymentIntentId",
		Column:            "payment_intent_id",
		SQLAlias:          "payment_intent_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_webhook_event_id": {
		SourcePath:        "provider_webhook_event_id",
		DefaultOutputPath: "providerWebhookEventId",
		Column:            "provider_webhook_event_id",
		SQLAlias:          "provider_webhook_event_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
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
	"old_intent_status": {
		SourcePath:        "old_intent_status",
		DefaultOutputPath: "oldIntentStatus",
		Column:            "old_intent_status",
		SQLAlias:          "old_intent_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"new_intent_status": {
		SourcePath:        "new_intent_status",
		DefaultOutputPath: "newIntentStatus",
		Column:            "new_intent_status",
		SQLAlias:          "new_intent_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"old_attempt_status": {
		SourcePath:        "old_attempt_status",
		DefaultOutputPath: "oldAttemptStatus",
		Column:            "old_attempt_status",
		SQLAlias:          "old_attempt_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"new_attempt_status": {
		SourcePath:        "new_attempt_status",
		DefaultOutputPath: "newAttemptStatus",
		Column:            "new_attempt_status",
		SQLAlias:          "new_attempt_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_status": {
		SourcePath:        "provider_status",
		DefaultOutputPath: "providerStatus",
		Column:            "provider_status",
		SQLAlias:          "provider_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason": {
		SourcePath:        "reason",
		DefaultOutputPath: "reason",
		Column:            "reason",
		SQLAlias:          "reason",
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

func NewPaymentStatusEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentStatusEventsFilterFields[field]
	return
}

type PaymentStatusEventsFilterResult struct {
	PaymentStatusEvents
	FilterCount int `db:"count"`
}

func ValidatePaymentStatusEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentStatusEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentStatusEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentStatusEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentStatusEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentStatusEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentStatusEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentStatusEvents struct {
	Id                     uuid.UUID       `db:"id"`
	PaymentIntentId        nuuid.NUUID     `db:"payment_intent_id"`
	PaymentAttemptId       nuuid.NUUID     `db:"payment_attempt_id"`
	ProviderWebhookEventId nuuid.NUUID     `db:"provider_webhook_event_id"`
	SourceType             string          `db:"source_type"`
	EventType              string          `db:"event_type"`
	OldIntentStatus        null.String     `db:"old_intent_status"`
	NewIntentStatus        null.String     `db:"new_intent_status"`
	OldAttemptStatus       null.String     `db:"old_attempt_status"`
	NewAttemptStatus       null.String     `db:"new_attempt_status"`
	ProviderStatus         null.String     `db:"provider_status"`
	Reason                 null.String     `db:"reason"`
	OccurredAt             time.Time       `db:"occurred_at"`
	Metadata               json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PaymentStatusEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentStatusEvents) ToPaymentStatusEventsPrimaryID() PaymentStatusEventsPrimaryID {
	return PaymentStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type PaymentStatusEventsList []*PaymentStatusEvents

type PaymentStatusEventsFilterResultList []*PaymentStatusEventsFilterResult

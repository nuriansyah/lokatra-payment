package model

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
	"time"
)

type ChargebackEventsDBFieldNameType string

type chargebackEventsDBFieldName struct {
	Id               ChargebackEventsDBFieldNameType
	DisputeId        ChargebackEventsDBFieldNameType
	EventType        ChargebackEventsDBFieldNameType
	JournalEntryId   ChargebackEventsDBFieldNameType
	Amount           ChargebackEventsDBFieldNameType
	CurrencyCode     ChargebackEventsDBFieldNameType
	ProviderEventRef ChargebackEventsDBFieldNameType
	EventPayload     ChargebackEventsDBFieldNameType
	OccurredAt       ChargebackEventsDBFieldNameType
	Metadata         ChargebackEventsDBFieldNameType
	MetaCreatedAt    ChargebackEventsDBFieldNameType
	MetaCreatedBy    ChargebackEventsDBFieldNameType
	MetaUpdatedAt    ChargebackEventsDBFieldNameType
	MetaUpdatedBy    ChargebackEventsDBFieldNameType
	MetaDeletedAt    ChargebackEventsDBFieldNameType
	MetaDeletedBy    ChargebackEventsDBFieldNameType
}

var ChargebackEventsDBFieldName = chargebackEventsDBFieldName{
	Id:               "id",
	DisputeId:        "dispute_id",
	EventType:        "event_type",
	JournalEntryId:   "journal_entry_id",
	Amount:           "amount",
	CurrencyCode:     "currency_code",
	ProviderEventRef: "provider_event_ref",
	EventPayload:     "event_payload",
	OccurredAt:       "occurred_at",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewChargebackEventsDBFieldNameFromStr(field string) (dbField ChargebackEventsDBFieldNameType, found bool) {
	switch field {

	case string(ChargebackEventsDBFieldName.Id):
		return ChargebackEventsDBFieldName.Id, true

	case string(ChargebackEventsDBFieldName.DisputeId):
		return ChargebackEventsDBFieldName.DisputeId, true

	case string(ChargebackEventsDBFieldName.EventType):
		return ChargebackEventsDBFieldName.EventType, true

	case string(ChargebackEventsDBFieldName.JournalEntryId):
		return ChargebackEventsDBFieldName.JournalEntryId, true

	case string(ChargebackEventsDBFieldName.Amount):
		return ChargebackEventsDBFieldName.Amount, true

	case string(ChargebackEventsDBFieldName.CurrencyCode):
		return ChargebackEventsDBFieldName.CurrencyCode, true

	case string(ChargebackEventsDBFieldName.ProviderEventRef):
		return ChargebackEventsDBFieldName.ProviderEventRef, true

	case string(ChargebackEventsDBFieldName.EventPayload):
		return ChargebackEventsDBFieldName.EventPayload, true

	case string(ChargebackEventsDBFieldName.OccurredAt):
		return ChargebackEventsDBFieldName.OccurredAt, true

	case string(ChargebackEventsDBFieldName.Metadata):
		return ChargebackEventsDBFieldName.Metadata, true

	case string(ChargebackEventsDBFieldName.MetaCreatedAt):
		return ChargebackEventsDBFieldName.MetaCreatedAt, true

	case string(ChargebackEventsDBFieldName.MetaCreatedBy):
		return ChargebackEventsDBFieldName.MetaCreatedBy, true

	case string(ChargebackEventsDBFieldName.MetaUpdatedAt):
		return ChargebackEventsDBFieldName.MetaUpdatedAt, true

	case string(ChargebackEventsDBFieldName.MetaUpdatedBy):
		return ChargebackEventsDBFieldName.MetaUpdatedBy, true

	case string(ChargebackEventsDBFieldName.MetaDeletedAt):
		return ChargebackEventsDBFieldName.MetaDeletedAt, true

	case string(ChargebackEventsDBFieldName.MetaDeletedBy):
		return ChargebackEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ChargebackEventsFilterJoins = map[string]JoinSpec{}

var ChargebackEventsFilterFields = map[string]FilterFieldSpec{
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
	"event_type": {
		SourcePath:        "event_type",
		DefaultOutputPath: "eventType",
		Column:            "event_type",
		SQLAlias:          "event_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"journal_entry_id": {
		SourcePath:        "journal_entry_id",
		DefaultOutputPath: "journalEntryId",
		Column:            "journal_entry_id",
		SQLAlias:          "journal_entry_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_event_ref": {
		SourcePath:        "provider_event_ref",
		DefaultOutputPath: "providerEventRef",
		Column:            "provider_event_ref",
		SQLAlias:          "provider_event_ref",
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

func NewChargebackEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ChargebackEventsFilterFields[field]
	return
}

type ChargebackEventsFilterResult struct {
	ChargebackEvents
	FilterCount int `db:"count"`
}

func ValidateChargebackEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewChargebackEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewChargebackEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewChargebackEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateChargebackEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateChargebackEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewChargebackEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateChargebackEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type EventType string

const (
	EventTypeChargebackDebited      EventType = "chargeback_debited"
	EventTypeChargebackReversed     EventType = "chargeback_reversed"
	EventTypeRepresentmentSubmitted EventType = "representment_submitted"
	EventTypeWon                    EventType = "won"
	EventTypeLost                   EventType = "lost"
)

type ChargebackEvents struct {
	Id               uuid.UUID       `db:"id"`
	DisputeId        uuid.UUID       `db:"dispute_id"`
	EventType        EventType       `db:"event_type"`
	JournalEntryId   nuuid.NUUID     `db:"journal_entry_id"`
	Amount           decimal.Decimal `db:"amount"`
	CurrencyCode     string          `db:"currency_code"`
	ProviderEventRef null.String     `db:"provider_event_ref"`
	EventPayload     json.RawMessage `db:"event_payload"`
	OccurredAt       time.Time       `db:"occurred_at"`
	Metadata         json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ChargebackEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ChargebackEvents) ToChargebackEventsPrimaryID() ChargebackEventsPrimaryID {
	return ChargebackEventsPrimaryID{
		Id: d.Id,
	}
}

type ChargebackEventsList []*ChargebackEvents

type ChargebackEventsFilterResultList []*ChargebackEventsFilterResult

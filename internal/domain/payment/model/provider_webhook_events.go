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

type ProviderWebhookEventsDBFieldNameType string

type providerWebhookEventsDBFieldName struct {
	Id                 ProviderWebhookEventsDBFieldNameType
	WebhookEndpointId  ProviderWebhookEventsDBFieldNameType
	EndpointKey        ProviderWebhookEventsDBFieldNameType
	ProviderAccountId  ProviderWebhookEventsDBFieldNameType
	ProviderCode       ProviderWebhookEventsDBFieldNameType
	EventId            ProviderWebhookEventsDBFieldNameType
	EventType          ProviderWebhookEventsDBFieldNameType
	ProviderReference  ProviderWebhookEventsDBFieldNameType
	ProviderStatus     ProviderWebhookEventsDBFieldNameType
	SignatureValid     ProviderWebhookEventsDBFieldNameType
	SignatureAlgorithm ProviderWebhookEventsDBFieldNameType
	Headers            ProviderWebhookEventsDBFieldNameType
	RawBody            ProviderWebhookEventsDBFieldNameType
	RawBodySha256      ProviderWebhookEventsDBFieldNameType
	ParsedBody         ProviderWebhookEventsDBFieldNameType
	ProcessingStatus   ProviderWebhookEventsDBFieldNameType
	RetryCount         ProviderWebhookEventsDBFieldNameType
	NextRetryAt        ProviderWebhookEventsDBFieldNameType
	LockedUntil        ProviderWebhookEventsDBFieldNameType
	ReceivedAt         ProviderWebhookEventsDBFieldNameType
	ProcessedAt        ProviderWebhookEventsDBFieldNameType
	ErrorCode          ProviderWebhookEventsDBFieldNameType
	ErrorMessage       ProviderWebhookEventsDBFieldNameType
	Metadata           ProviderWebhookEventsDBFieldNameType
	MetaCreatedAt      ProviderWebhookEventsDBFieldNameType
	MetaCreatedBy      ProviderWebhookEventsDBFieldNameType
	MetaUpdatedAt      ProviderWebhookEventsDBFieldNameType
	MetaUpdatedBy      ProviderWebhookEventsDBFieldNameType
	MetaDeletedAt      ProviderWebhookEventsDBFieldNameType
	MetaDeletedBy      ProviderWebhookEventsDBFieldNameType
}

var ProviderWebhookEventsDBFieldName = providerWebhookEventsDBFieldName{
	Id:                 "id",
	WebhookEndpointId:  "webhook_endpoint_id",
	EndpointKey:        "endpoint_key",
	ProviderAccountId:  "provider_account_id",
	ProviderCode:       "provider_code",
	EventId:            "event_id",
	EventType:          "event_type",
	ProviderReference:  "provider_reference",
	ProviderStatus:     "provider_status",
	SignatureValid:     "signature_valid",
	SignatureAlgorithm: "signature_algorithm",
	Headers:            "headers",
	RawBody:            "raw_body",
	RawBodySha256:      "raw_body_sha256",
	ParsedBody:         "parsed_body",
	ProcessingStatus:   "processing_status",
	RetryCount:         "retry_count",
	NextRetryAt:        "next_retry_at",
	LockedUntil:        "locked_until",
	ReceivedAt:         "received_at",
	ProcessedAt:        "processed_at",
	ErrorCode:          "error_code",
	ErrorMessage:       "error_message",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewProviderWebhookEventsDBFieldNameFromStr(field string) (dbField ProviderWebhookEventsDBFieldNameType, found bool) {
	switch field {

	case string(ProviderWebhookEventsDBFieldName.Id):
		return ProviderWebhookEventsDBFieldName.Id, true

	case string(ProviderWebhookEventsDBFieldName.WebhookEndpointId):
		return ProviderWebhookEventsDBFieldName.WebhookEndpointId, true

	case string(ProviderWebhookEventsDBFieldName.EndpointKey):
		return ProviderWebhookEventsDBFieldName.EndpointKey, true

	case string(ProviderWebhookEventsDBFieldName.ProviderAccountId):
		return ProviderWebhookEventsDBFieldName.ProviderAccountId, true

	case string(ProviderWebhookEventsDBFieldName.ProviderCode):
		return ProviderWebhookEventsDBFieldName.ProviderCode, true

	case string(ProviderWebhookEventsDBFieldName.EventId):
		return ProviderWebhookEventsDBFieldName.EventId, true

	case string(ProviderWebhookEventsDBFieldName.EventType):
		return ProviderWebhookEventsDBFieldName.EventType, true

	case string(ProviderWebhookEventsDBFieldName.ProviderReference):
		return ProviderWebhookEventsDBFieldName.ProviderReference, true

	case string(ProviderWebhookEventsDBFieldName.ProviderStatus):
		return ProviderWebhookEventsDBFieldName.ProviderStatus, true

	case string(ProviderWebhookEventsDBFieldName.SignatureValid):
		return ProviderWebhookEventsDBFieldName.SignatureValid, true

	case string(ProviderWebhookEventsDBFieldName.SignatureAlgorithm):
		return ProviderWebhookEventsDBFieldName.SignatureAlgorithm, true

	case string(ProviderWebhookEventsDBFieldName.Headers):
		return ProviderWebhookEventsDBFieldName.Headers, true

	case string(ProviderWebhookEventsDBFieldName.RawBody):
		return ProviderWebhookEventsDBFieldName.RawBody, true

	case string(ProviderWebhookEventsDBFieldName.RawBodySha256):
		return ProviderWebhookEventsDBFieldName.RawBodySha256, true

	case string(ProviderWebhookEventsDBFieldName.ParsedBody):
		return ProviderWebhookEventsDBFieldName.ParsedBody, true

	case string(ProviderWebhookEventsDBFieldName.ProcessingStatus):
		return ProviderWebhookEventsDBFieldName.ProcessingStatus, true

	case string(ProviderWebhookEventsDBFieldName.RetryCount):
		return ProviderWebhookEventsDBFieldName.RetryCount, true

	case string(ProviderWebhookEventsDBFieldName.NextRetryAt):
		return ProviderWebhookEventsDBFieldName.NextRetryAt, true

	case string(ProviderWebhookEventsDBFieldName.LockedUntil):
		return ProviderWebhookEventsDBFieldName.LockedUntil, true

	case string(ProviderWebhookEventsDBFieldName.ReceivedAt):
		return ProviderWebhookEventsDBFieldName.ReceivedAt, true

	case string(ProviderWebhookEventsDBFieldName.ProcessedAt):
		return ProviderWebhookEventsDBFieldName.ProcessedAt, true

	case string(ProviderWebhookEventsDBFieldName.ErrorCode):
		return ProviderWebhookEventsDBFieldName.ErrorCode, true

	case string(ProviderWebhookEventsDBFieldName.ErrorMessage):
		return ProviderWebhookEventsDBFieldName.ErrorMessage, true

	case string(ProviderWebhookEventsDBFieldName.Metadata):
		return ProviderWebhookEventsDBFieldName.Metadata, true

	case string(ProviderWebhookEventsDBFieldName.MetaCreatedAt):
		return ProviderWebhookEventsDBFieldName.MetaCreatedAt, true

	case string(ProviderWebhookEventsDBFieldName.MetaCreatedBy):
		return ProviderWebhookEventsDBFieldName.MetaCreatedBy, true

	case string(ProviderWebhookEventsDBFieldName.MetaUpdatedAt):
		return ProviderWebhookEventsDBFieldName.MetaUpdatedAt, true

	case string(ProviderWebhookEventsDBFieldName.MetaUpdatedBy):
		return ProviderWebhookEventsDBFieldName.MetaUpdatedBy, true

	case string(ProviderWebhookEventsDBFieldName.MetaDeletedAt):
		return ProviderWebhookEventsDBFieldName.MetaDeletedAt, true

	case string(ProviderWebhookEventsDBFieldName.MetaDeletedBy):
		return ProviderWebhookEventsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderWebhookEventsFilterJoins = map[string]JoinSpec{}

var ProviderWebhookEventsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"webhook_endpoint_id": {
		SourcePath:        "webhook_endpoint_id",
		DefaultOutputPath: "webhookEndpointId",
		Column:            "webhook_endpoint_id",
		SQLAlias:          "webhook_endpoint_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"endpoint_key": {
		SourcePath:        "endpoint_key",
		DefaultOutputPath: "endpointKey",
		Column:            "endpoint_key",
		SQLAlias:          "endpoint_key",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_account_id": {
		SourcePath:        "provider_account_id",
		DefaultOutputPath: "providerAccountId",
		Column:            "provider_account_id",
		SQLAlias:          "provider_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_code": {
		SourcePath:        "provider_code",
		DefaultOutputPath: "providerCode",
		Column:            "provider_code",
		SQLAlias:          "provider_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"event_id": {
		SourcePath:        "event_id",
		DefaultOutputPath: "eventId",
		Column:            "event_id",
		SQLAlias:          "event_id",
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
	"provider_reference": {
		SourcePath:        "provider_reference",
		DefaultOutputPath: "providerReference",
		Column:            "provider_reference",
		SQLAlias:          "provider_reference",
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
	"signature_valid": {
		SourcePath:        "signature_valid",
		DefaultOutputPath: "signatureValid",
		Column:            "signature_valid",
		SQLAlias:          "signature_valid",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"signature_algorithm": {
		SourcePath:        "signature_algorithm",
		DefaultOutputPath: "signatureAlgorithm",
		Column:            "signature_algorithm",
		SQLAlias:          "signature_algorithm",
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
	"raw_body": {
		SourcePath:        "raw_body",
		DefaultOutputPath: "rawBody",
		Column:            "raw_body",
		SQLAlias:          "raw_body",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"raw_body_sha256": {
		SourcePath:        "raw_body_sha256",
		DefaultOutputPath: "rawBodySha256",
		Column:            "raw_body_sha256",
		SQLAlias:          "raw_body_sha256",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"parsed_body": {
		SourcePath:        "parsed_body",
		DefaultOutputPath: "parsedBody",
		Column:            "parsed_body",
		SQLAlias:          "parsed_body",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"processing_status": {
		SourcePath:        "processing_status",
		DefaultOutputPath: "processingStatus",
		Column:            "processing_status",
		SQLAlias:          "processing_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"retry_count": {
		SourcePath:        "retry_count",
		DefaultOutputPath: "retryCount",
		Column:            "retry_count",
		SQLAlias:          "retry_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"next_retry_at": {
		SourcePath:        "next_retry_at",
		DefaultOutputPath: "nextRetryAt",
		Column:            "next_retry_at",
		SQLAlias:          "next_retry_at",
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
	"received_at": {
		SourcePath:        "received_at",
		DefaultOutputPath: "receivedAt",
		Column:            "received_at",
		SQLAlias:          "received_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"processed_at": {
		SourcePath:        "processed_at",
		DefaultOutputPath: "processedAt",
		Column:            "processed_at",
		SQLAlias:          "processed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_code": {
		SourcePath:        "error_code",
		DefaultOutputPath: "errorCode",
		Column:            "error_code",
		SQLAlias:          "error_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_message": {
		SourcePath:        "error_message",
		DefaultOutputPath: "errorMessage",
		Column:            "error_message",
		SQLAlias:          "error_message",
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

func NewProviderWebhookEventsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderWebhookEventsFilterFields[field]
	return
}

type ProviderWebhookEventsFilterResult struct {
	ProviderWebhookEvents
	FilterCount int `db:"count"`
}

func ValidateProviderWebhookEventsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderWebhookEventsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderWebhookEventsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderWebhookEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderWebhookEventsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderWebhookEventsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderWebhookEventsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderWebhookEventsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type WebhookProcessingStatus string

const (
	WebhookProcessingStatusReceived   WebhookProcessingStatus = "received"
	WebhookProcessingStatusProcessing WebhookProcessingStatus = "processing"
	WebhookProcessingStatusProcessed  WebhookProcessingStatus = "processed"
	WebhookProcessingStatusFailed     WebhookProcessingStatus = "failed"
)

type ProviderWebhookEvents struct {
	Id                 uuid.UUID               `db:"id"`
	WebhookEndpointId  nuuid.NUUID             `db:"webhook_endpoint_id"`
	EndpointKey        null.String             `db:"endpoint_key"`
	ProviderAccountId  uuid.UUID               `db:"provider_account_id"`
	ProviderCode       string                  `db:"provider_code"`
	EventId            null.String             `db:"event_id"`
	EventType          null.String             `db:"event_type"`
	ProviderReference  null.String             `db:"provider_reference"`
	ProviderStatus     null.String             `db:"provider_status"`
	SignatureValid     bool                    `db:"signature_valid"`
	SignatureAlgorithm null.String             `db:"signature_algorithm"`
	Headers            json.RawMessage         `db:"headers"`
	RawBody            []byte                  `db:"raw_body"`
	RawBodySha256      string                  `db:"raw_body_sha256"`
	ParsedBody         json.RawMessage         `db:"parsed_body"`
	ProcessingStatus   WebhookProcessingStatus `db:"processing_status"`
	RetryCount         int                     `db:"retry_count"`
	NextRetryAt        null.Time               `db:"next_retry_at"`
	LockedUntil        null.Time               `db:"locked_until"`
	ReceivedAt         time.Time               `db:"received_at"`
	ProcessedAt        null.Time               `db:"processed_at"`
	ErrorCode          null.String             `db:"error_code"`
	ErrorMessage       null.String             `db:"error_message"`
	Metadata           json.RawMessage         `db:"metadata"`

	shared.MetaSignature
}
type ProviderWebhookEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderWebhookEvents) ToProviderWebhookEventsPrimaryID() ProviderWebhookEventsPrimaryID {
	return ProviderWebhookEventsPrimaryID{
		Id: d.Id,
	}
}

type ProviderWebhookEventsList []*ProviderWebhookEvents

type ProviderWebhookEventsFilterResultList []*ProviderWebhookEventsFilterResult

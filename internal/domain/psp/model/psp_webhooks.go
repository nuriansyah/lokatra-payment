package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PspWebhooksDBFieldNameType string

type pspWebhooksDBFieldName struct {
	Id                 PspWebhooksDBFieldNameType
	PspAccountId       PspWebhooksDBFieldNameType
	Psp                PspWebhooksDBFieldNameType
	PspEventId         PspWebhooksDBFieldNameType
	PspEventType       PspWebhooksDBFieldNameType
	ReceivedAt         PspWebhooksDBFieldNameType
	Headers            PspWebhooksDBFieldNameType
	RawPayload         PspWebhooksDBFieldNameType
	HmacValid          PspWebhooksDBFieldNameType
	Status             PspWebhooksDBFieldNameType
	ProcessingAttempts PspWebhooksDBFieldNameType
	LastError          PspWebhooksDBFieldNameType
	ProcessedAt        PspWebhooksDBFieldNameType
	ResolvedPaymentId  PspWebhooksDBFieldNameType
	ResolvedIntentId   PspWebhooksDBFieldNameType
	MetaCreatedAt      PspWebhooksDBFieldNameType
	MetaCreatedBy      PspWebhooksDBFieldNameType
	MetaUpdatedAt      PspWebhooksDBFieldNameType
	MetaUpdatedBy      PspWebhooksDBFieldNameType
}

var PspWebhooksDBFieldName = pspWebhooksDBFieldName{
	Id:                 "id",
	PspAccountId:       "psp_account_id",
	Psp:                "psp",
	PspEventId:         "psp_event_id",
	PspEventType:       "psp_event_type",
	ReceivedAt:         "received_at",
	Headers:            "headers",
	RawPayload:         "raw_payload",
	HmacValid:          "hmac_valid",
	Status:             "status",
	ProcessingAttempts: "processing_attempts",
	LastError:          "last_error",
	ProcessedAt:        "processed_at",
	ResolvedPaymentId:  "resolved_payment_id",
	ResolvedIntentId:   "resolved_intent_id",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
}

func NewPspWebhooksDBFieldNameFromStr(field string) (dbField PspWebhooksDBFieldNameType, found bool) {
	switch field {

	case string(PspWebhooksDBFieldName.Id):
		return PspWebhooksDBFieldName.Id, true

	case string(PspWebhooksDBFieldName.PspAccountId):
		return PspWebhooksDBFieldName.PspAccountId, true

	case string(PspWebhooksDBFieldName.Psp):
		return PspWebhooksDBFieldName.Psp, true

	case string(PspWebhooksDBFieldName.PspEventId):
		return PspWebhooksDBFieldName.PspEventId, true

	case string(PspWebhooksDBFieldName.PspEventType):
		return PspWebhooksDBFieldName.PspEventType, true

	case string(PspWebhooksDBFieldName.ReceivedAt):
		return PspWebhooksDBFieldName.ReceivedAt, true

	case string(PspWebhooksDBFieldName.Headers):
		return PspWebhooksDBFieldName.Headers, true

	case string(PspWebhooksDBFieldName.RawPayload):
		return PspWebhooksDBFieldName.RawPayload, true

	case string(PspWebhooksDBFieldName.HmacValid):
		return PspWebhooksDBFieldName.HmacValid, true

	case string(PspWebhooksDBFieldName.Status):
		return PspWebhooksDBFieldName.Status, true

	case string(PspWebhooksDBFieldName.ProcessingAttempts):
		return PspWebhooksDBFieldName.ProcessingAttempts, true

	case string(PspWebhooksDBFieldName.LastError):
		return PspWebhooksDBFieldName.LastError, true

	case string(PspWebhooksDBFieldName.ProcessedAt):
		return PspWebhooksDBFieldName.ProcessedAt, true

	case string(PspWebhooksDBFieldName.ResolvedPaymentId):
		return PspWebhooksDBFieldName.ResolvedPaymentId, true

	case string(PspWebhooksDBFieldName.ResolvedIntentId):
		return PspWebhooksDBFieldName.ResolvedIntentId, true

	case string(PspWebhooksDBFieldName.MetaCreatedAt):
		return PspWebhooksDBFieldName.MetaCreatedAt, true

	case string(PspWebhooksDBFieldName.MetaCreatedBy):
		return PspWebhooksDBFieldName.MetaCreatedBy, true

	case string(PspWebhooksDBFieldName.MetaUpdatedAt):
		return PspWebhooksDBFieldName.MetaUpdatedAt, true

	case string(PspWebhooksDBFieldName.MetaUpdatedBy):
		return PspWebhooksDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

type PspWebhooksFilterResult struct {
	PspWebhooks
	FilterCount int `db:"count"`
}

func ValidatePspWebhooksFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewPspWebhooksDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewPspWebhooksDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewPspWebhooksDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type WebhookStatus string

const (
	WebhookStatusReceived   WebhookStatus = "RECEIVED"
	WebhookStatusProcessing WebhookStatus = "PROCESSING"
	WebhookStatusProcessed  WebhookStatus = "PROCESSED"
	WebhookStatusFailed     WebhookStatus = "FAILED"
	WebhookStatusIgnored    WebhookStatus = "IGNORED"
)

type PspWebhooks struct {
	Id                 uuid.UUID       `db:"id"`
	PspAccountId       uuid.UUID       `db:"psp_account_id"`
	Psp                string          `db:"psp"`
	PspEventId         null.String     `db:"psp_event_id"`
	PspEventType       string          `db:"psp_event_type"`
	ReceivedAt         time.Time       `db:"received_at"`
	Headers            json.RawMessage `db:"headers"`
	RawPayload         json.RawMessage `db:"raw_payload"`
	HmacValid          null.Bool       `db:"hmac_valid"`
	Status             WebhookStatus   `db:"status"`
	ProcessingAttempts int             `db:"processing_attempts"`
	LastError          null.String     `db:"last_error"`
	ProcessedAt        null.Time       `db:"processed_at"`
	ResolvedPaymentId  nuuid.NUUID     `db:"resolved_payment_id"`
	ResolvedIntentId   nuuid.NUUID     `db:"resolved_intent_id"`
	MetaCreatedAt      time.Time       `db:"meta_created_at"`
	MetaCreatedBy      uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt      time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy      nuuid.NUUID     `db:"meta_updated_by"`
}
type PspWebhooksPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PspWebhooks) ToPspWebhooksPrimaryID() PspWebhooksPrimaryID {
	return PspWebhooksPrimaryID{
		Id: d.Id,
	}
}

type PspWebhooksList []*PspWebhooks

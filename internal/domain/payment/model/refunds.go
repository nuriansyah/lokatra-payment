package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type RefundsDBFieldNameType string

type refundsDBFieldName struct {
	Id               RefundsDBFieldNameType
	RefundCode       RefundsDBFieldNameType
	PaymentId        RefundsDBFieldNameType
	IntentId         RefundsDBFieldNameType
	Amount           RefundsDBFieldNameType
	Currency         RefundsDBFieldNameType
	Reason           RefundsDBFieldNameType
	ReasonDetail     RefundsDBFieldNameType
	Status           RefundsDBFieldNameType
	PspRefundId      RefundsDBFieldNameType
	PspRawResponse   RefundsDBFieldNameType
	RequestedBy      RefundsDBFieldNameType
	ReviewedBy       RefundsDBFieldNameType
	ReviewedAt       RefundsDBFieldNameType
	ReviewNotes      RefundsDBFieldNameType
	RefundedAt       RefundsDBFieldNameType
	EstimatedArrival RefundsDBFieldNameType
	FailureReason    RefundsDBFieldNameType
	IdempotencyKeyId RefundsDBFieldNameType
	Metadata         RefundsDBFieldNameType
	MetaCreatedAt    RefundsDBFieldNameType
	MetaCreatedBy    RefundsDBFieldNameType
	MetaUpdatedAt    RefundsDBFieldNameType
	MetaUpdatedBy    RefundsDBFieldNameType
	MetaDeletedAt    RefundsDBFieldNameType
}

var RefundsDBFieldName = refundsDBFieldName{
	Id:               "id",
	RefundCode:       "refund_code",
	PaymentId:        "payment_id",
	IntentId:         "intent_id",
	Amount:           "amount",
	Currency:         "currency",
	Reason:           "reason",
	ReasonDetail:     "reason_detail",
	Status:           "status",
	PspRefundId:      "psp_refund_id",
	PspRawResponse:   "psp_raw_response",
	RequestedBy:      "requested_by",
	ReviewedBy:       "reviewed_by",
	ReviewedAt:       "reviewed_at",
	ReviewNotes:      "review_notes",
	RefundedAt:       "refunded_at",
	EstimatedArrival: "estimated_arrival",
	FailureReason:    "failure_reason",
	IdempotencyKeyId: "idempotency_key_id",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
}

func NewRefundsDBFieldNameFromStr(field string) (dbField RefundsDBFieldNameType, found bool) {
	switch field {

	case string(RefundsDBFieldName.Id):
		return RefundsDBFieldName.Id, true

	case string(RefundsDBFieldName.RefundCode):
		return RefundsDBFieldName.RefundCode, true

	case string(RefundsDBFieldName.PaymentId):
		return RefundsDBFieldName.PaymentId, true

	case string(RefundsDBFieldName.IntentId):
		return RefundsDBFieldName.IntentId, true

	case string(RefundsDBFieldName.Amount):
		return RefundsDBFieldName.Amount, true

	case string(RefundsDBFieldName.Currency):
		return RefundsDBFieldName.Currency, true

	case string(RefundsDBFieldName.Reason):
		return RefundsDBFieldName.Reason, true

	case string(RefundsDBFieldName.ReasonDetail):
		return RefundsDBFieldName.ReasonDetail, true

	case string(RefundsDBFieldName.Status):
		return RefundsDBFieldName.Status, true

	case string(RefundsDBFieldName.PspRefundId):
		return RefundsDBFieldName.PspRefundId, true

	case string(RefundsDBFieldName.PspRawResponse):
		return RefundsDBFieldName.PspRawResponse, true

	case string(RefundsDBFieldName.RequestedBy):
		return RefundsDBFieldName.RequestedBy, true

	case string(RefundsDBFieldName.ReviewedBy):
		return RefundsDBFieldName.ReviewedBy, true

	case string(RefundsDBFieldName.ReviewedAt):
		return RefundsDBFieldName.ReviewedAt, true

	case string(RefundsDBFieldName.ReviewNotes):
		return RefundsDBFieldName.ReviewNotes, true

	case string(RefundsDBFieldName.RefundedAt):
		return RefundsDBFieldName.RefundedAt, true

	case string(RefundsDBFieldName.EstimatedArrival):
		return RefundsDBFieldName.EstimatedArrival, true

	case string(RefundsDBFieldName.FailureReason):
		return RefundsDBFieldName.FailureReason, true

	case string(RefundsDBFieldName.IdempotencyKeyId):
		return RefundsDBFieldName.IdempotencyKeyId, true

	case string(RefundsDBFieldName.Metadata):
		return RefundsDBFieldName.Metadata, true

	case string(RefundsDBFieldName.MetaCreatedAt):
		return RefundsDBFieldName.MetaCreatedAt, true

	case string(RefundsDBFieldName.MetaCreatedBy):
		return RefundsDBFieldName.MetaCreatedBy, true

	case string(RefundsDBFieldName.MetaUpdatedAt):
		return RefundsDBFieldName.MetaUpdatedAt, true

	case string(RefundsDBFieldName.MetaUpdatedBy):
		return RefundsDBFieldName.MetaUpdatedBy, true

	case string(RefundsDBFieldName.MetaDeletedAt):
		return RefundsDBFieldName.MetaDeletedAt, true

	}
	return "unknown", false
}

type RefundsFilterResult struct {
	Refunds
	FilterCount int `db:"count"`
}

func ValidateRefundsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewRefundsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewRefundsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewRefundsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type RefundStatus string

const (
	RefundStatusPending    RefundStatus = "PENDING"
	RefundStatusProcessing RefundStatus = "PROCESSING"
	RefundStatusSucceeded  RefundStatus = "SUCCEEDED"
	RefundStatusFailed     RefundStatus = "FAILED"
	RefundStatusCancelled  RefundStatus = "CANCELLED"
)

type Refunds struct {
	Id               uuid.UUID       `db:"id"`
	RefundCode       string          `db:"refund_code"`
	PaymentId        uuid.UUID       `db:"payment_id"`
	IntentId         uuid.UUID       `db:"intent_id"`
	Amount           decimal.Decimal `db:"amount"`
	Currency         PaymentCurrency `db:"currency"`
	Reason           string          `db:"reason"`
	ReasonDetail     null.String     `db:"reason_detail"`
	Status           RefundStatus    `db:"status"`
	PspRefundId      null.String     `db:"psp_refund_id"`
	PspRawResponse   json.RawMessage `db:"psp_raw_response"`
	RequestedBy      uuid.UUID       `db:"requested_by"`
	ReviewedBy       nuuid.NUUID     `db:"reviewed_by"`
	ReviewedAt       null.Time       `db:"reviewed_at"`
	ReviewNotes      null.String     `db:"review_notes"`
	RefundedAt       null.Time       `db:"refunded_at"`
	EstimatedArrival null.Time       `db:"estimated_arrival"`
	FailureReason    null.String     `db:"failure_reason"`
	IdempotencyKeyId nuuid.NUUID     `db:"idempotency_key_id"`
	Metadata         json.RawMessage `db:"metadata"`
	MetaCreatedAt    time.Time       `db:"meta_created_at"`
	MetaCreatedBy    uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt    time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy    nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt    null.Time       `db:"meta_deleted_at"`
}
type RefundsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Refunds) ToRefundsPrimaryID() RefundsPrimaryID {
	return RefundsPrimaryID{
		Id: d.Id,
	}
}

type RefundsList []*Refunds

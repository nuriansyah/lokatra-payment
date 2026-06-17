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
)

type PaymentIntentsDBFieldNameType string

type paymentIntentsDBFieldName struct {
	Id                  PaymentIntentsDBFieldNameType
	IntentCode          PaymentIntentsDBFieldNameType
	SourceService       PaymentIntentsDBFieldNameType
	SourceType          PaymentIntentsDBFieldNameType
	SourceId            PaymentIntentsDBFieldNameType
	MerchantId          PaymentIntentsDBFieldNameType
	CustomerId          PaymentIntentsDBFieldNameType
	Amount              PaymentIntentsDBFieldNameType
	Currency            PaymentIntentsDBFieldNameType
	Status              PaymentIntentsDBFieldNameType
	SelectedMethodCode  PaymentIntentsDBFieldNameType
	SelectedChannelCode PaymentIntentsDBFieldNameType
	Description         PaymentIntentsDBFieldNameType
	ExpiresAt           PaymentIntentsDBFieldNameType
	PaidAt              PaymentIntentsDBFieldNameType
	CanceledAt          PaymentIntentsDBFieldNameType
	CancellationReason  PaymentIntentsDBFieldNameType
	IdempotencyKey      PaymentIntentsDBFieldNameType
	SourceSnapshot      PaymentIntentsDBFieldNameType
	Metadata            PaymentIntentsDBFieldNameType
	MetaCreatedAt       PaymentIntentsDBFieldNameType
	MetaCreatedBy       PaymentIntentsDBFieldNameType
	MetaUpdatedAt       PaymentIntentsDBFieldNameType
	MetaUpdatedBy       PaymentIntentsDBFieldNameType
	MetaDeletedAt       PaymentIntentsDBFieldNameType
	MetaDeletedBy       PaymentIntentsDBFieldNameType
}

var PaymentIntentsDBFieldName = paymentIntentsDBFieldName{
	Id:                  "id",
	IntentCode:          "intent_code",
	SourceService:       "source_service",
	SourceType:          "source_type",
	SourceId:            "source_id",
	MerchantId:          "merchant_id",
	CustomerId:          "customer_id",
	Amount:              "amount",
	Currency:            "currency",
	Status:              "status",
	SelectedMethodCode:  "selected_method_code",
	SelectedChannelCode: "selected_channel_code",
	Description:         "description",
	ExpiresAt:           "expires_at",
	PaidAt:              "paid_at",
	CanceledAt:          "canceled_at",
	CancellationReason:  "cancellation_reason",
	IdempotencyKey:      "idempotency_key",
	SourceSnapshot:      "source_snapshot",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewPaymentIntentsDBFieldNameFromStr(field string) (dbField PaymentIntentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentIntentsDBFieldName.Id):
		return PaymentIntentsDBFieldName.Id, true

	case string(PaymentIntentsDBFieldName.IntentCode):
		return PaymentIntentsDBFieldName.IntentCode, true

	case string(PaymentIntentsDBFieldName.SourceService):
		return PaymentIntentsDBFieldName.SourceService, true

	case string(PaymentIntentsDBFieldName.SourceType):
		return PaymentIntentsDBFieldName.SourceType, true

	case string(PaymentIntentsDBFieldName.SourceId):
		return PaymentIntentsDBFieldName.SourceId, true

	case string(PaymentIntentsDBFieldName.MerchantId):
		return PaymentIntentsDBFieldName.MerchantId, true

	case string(PaymentIntentsDBFieldName.CustomerId):
		return PaymentIntentsDBFieldName.CustomerId, true

	case string(PaymentIntentsDBFieldName.Amount):
		return PaymentIntentsDBFieldName.Amount, true

	case string(PaymentIntentsDBFieldName.Currency):
		return PaymentIntentsDBFieldName.Currency, true

	case string(PaymentIntentsDBFieldName.Status):
		return PaymentIntentsDBFieldName.Status, true

	case string(PaymentIntentsDBFieldName.SelectedMethodCode):
		return PaymentIntentsDBFieldName.SelectedMethodCode, true

	case string(PaymentIntentsDBFieldName.SelectedChannelCode):
		return PaymentIntentsDBFieldName.SelectedChannelCode, true

	case string(PaymentIntentsDBFieldName.Description):
		return PaymentIntentsDBFieldName.Description, true

	case string(PaymentIntentsDBFieldName.ExpiresAt):
		return PaymentIntentsDBFieldName.ExpiresAt, true

	case string(PaymentIntentsDBFieldName.PaidAt):
		return PaymentIntentsDBFieldName.PaidAt, true

	case string(PaymentIntentsDBFieldName.CanceledAt):
		return PaymentIntentsDBFieldName.CanceledAt, true

	case string(PaymentIntentsDBFieldName.CancellationReason):
		return PaymentIntentsDBFieldName.CancellationReason, true

	case string(PaymentIntentsDBFieldName.IdempotencyKey):
		return PaymentIntentsDBFieldName.IdempotencyKey, true

	case string(PaymentIntentsDBFieldName.SourceSnapshot):
		return PaymentIntentsDBFieldName.SourceSnapshot, true

	case string(PaymentIntentsDBFieldName.Metadata):
		return PaymentIntentsDBFieldName.Metadata, true

	case string(PaymentIntentsDBFieldName.MetaCreatedAt):
		return PaymentIntentsDBFieldName.MetaCreatedAt, true

	case string(PaymentIntentsDBFieldName.MetaCreatedBy):
		return PaymentIntentsDBFieldName.MetaCreatedBy, true

	case string(PaymentIntentsDBFieldName.MetaUpdatedAt):
		return PaymentIntentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentIntentsDBFieldName.MetaUpdatedBy):
		return PaymentIntentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentIntentsDBFieldName.MetaDeletedAt):
		return PaymentIntentsDBFieldName.MetaDeletedAt, true

	case string(PaymentIntentsDBFieldName.MetaDeletedBy):
		return PaymentIntentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentIntentsFilterJoins = map[string]JoinSpec{}

var PaymentIntentsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"intent_code": {
		SourcePath:        "intent_code",
		DefaultOutputPath: "intentCode",
		Column:            "intent_code",
		SQLAlias:          "intent_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_service": {
		SourcePath:        "source_service",
		DefaultOutputPath: "sourceService",
		Column:            "source_service",
		SQLAlias:          "source_service",
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
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_id": {
		SourcePath:        "merchant_id",
		DefaultOutputPath: "merchantId",
		Column:            "merchant_id",
		SQLAlias:          "merchant_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"customer_id": {
		SourcePath:        "customer_id",
		DefaultOutputPath: "customerId",
		Column:            "customer_id",
		SQLAlias:          "customer_id",
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
	"currency": {
		SourcePath:        "currency",
		DefaultOutputPath: "currency",
		Column:            "currency",
		SQLAlias:          "currency",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"selected_method_code": {
		SourcePath:        "selected_method_code",
		DefaultOutputPath: "selectedMethodCode",
		Column:            "selected_method_code",
		SQLAlias:          "selected_method_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"selected_channel_code": {
		SourcePath:        "selected_channel_code",
		DefaultOutputPath: "selectedChannelCode",
		Column:            "selected_channel_code",
		SQLAlias:          "selected_channel_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"description": {
		SourcePath:        "description",
		DefaultOutputPath: "description",
		Column:            "description",
		SQLAlias:          "description",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"expires_at": {
		SourcePath:        "expires_at",
		DefaultOutputPath: "expiresAt",
		Column:            "expires_at",
		SQLAlias:          "expires_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"paid_at": {
		SourcePath:        "paid_at",
		DefaultOutputPath: "paidAt",
		Column:            "paid_at",
		SQLAlias:          "paid_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"canceled_at": {
		SourcePath:        "canceled_at",
		DefaultOutputPath: "canceledAt",
		Column:            "canceled_at",
		SQLAlias:          "canceled_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"cancellation_reason": {
		SourcePath:        "cancellation_reason",
		DefaultOutputPath: "cancellationReason",
		Column:            "cancellation_reason",
		SQLAlias:          "cancellation_reason",
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
	"source_snapshot": {
		SourcePath:        "source_snapshot",
		DefaultOutputPath: "sourceSnapshot",
		Column:            "source_snapshot",
		SQLAlias:          "source_snapshot",
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

func NewPaymentIntentsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentIntentsFilterFields[field]
	return
}

type PaymentIntentsFilterResult struct {
	PaymentIntents
	FilterCount int `db:"count"`
}

func ValidatePaymentIntentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentIntentsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentIntentsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentIntentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentIntentsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentIntentsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentIntentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentIntentsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentIntentStatus string

const (
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "requires_payment_method"
	PaymentIntentStatusRequiresConfirmation  PaymentIntentStatus = "requires_confirmation"
	PaymentIntentStatusRequiresAction        PaymentIntentStatus = "requires_action"
	PaymentIntentStatusProcessing            PaymentIntentStatus = "processing"
	PaymentIntentStatusSucceeded             PaymentIntentStatus = "succeeded"
	PaymentIntentStatusCanceled              PaymentIntentStatus = "canceled"
)

type PaymentIntents struct {
	Id                  uuid.UUID           `db:"id"`
	IntentCode          string              `db:"intent_code"`
	SourceService       string              `db:"source_service"`
	SourceType          string              `db:"source_type"`
	SourceId            uuid.UUID           `db:"source_id"`
	MerchantId          uuid.UUID           `db:"merchant_id"`
	CustomerId          nuuid.NUUID         `db:"customer_id"`
	Amount              decimal.Decimal     `db:"amount"`
	Currency            string              `db:"currency"`
	Status              PaymentIntentStatus `db:"status"`
	SelectedMethodCode  null.String         `db:"selected_method_code"`
	SelectedChannelCode null.String         `db:"selected_channel_code"`
	Description         null.String         `db:"description"`
	ExpiresAt           null.Time           `db:"expires_at"`
	PaidAt              null.Time           `db:"paid_at"`
	CanceledAt          null.Time           `db:"canceled_at"`
	CancellationReason  null.String         `db:"cancellation_reason"`
	IdempotencyKey      string              `db:"idempotency_key"`
	SourceSnapshot      json.RawMessage     `db:"source_snapshot"`
	Metadata            json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type PaymentIntentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentIntents) ToPaymentIntentsPrimaryID() PaymentIntentsPrimaryID {
	return PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentsList []*PaymentIntents

type PaymentIntentsFilterResultList []*PaymentIntentsFilterResult

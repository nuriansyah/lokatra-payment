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

type PaymentOverpaymentsDBFieldNameType string

type paymentOverpaymentsDBFieldName struct {
	Id                PaymentOverpaymentsDBFieldNameType
	PaymentIntentId   PaymentOverpaymentsDBFieldNameType
	PaidAttemptId     PaymentOverpaymentsDBFieldNameType
	OverpaidAttemptId PaymentOverpaymentsDBFieldNameType
	ExpectedAmount    PaymentOverpaymentsDBFieldNameType
	ReceivedAmount    PaymentOverpaymentsDBFieldNameType
	OverpaidAmount    PaymentOverpaymentsDBFieldNameType
	Currency          PaymentOverpaymentsDBFieldNameType
	Status            PaymentOverpaymentsDBFieldNameType
	ResolutionAction  PaymentOverpaymentsDBFieldNameType
	ResolutionNotes   PaymentOverpaymentsDBFieldNameType
	ResolvedAt        PaymentOverpaymentsDBFieldNameType
	ResolvedBy        PaymentOverpaymentsDBFieldNameType
	Metadata          PaymentOverpaymentsDBFieldNameType
	MetaCreatedAt     PaymentOverpaymentsDBFieldNameType
	MetaCreatedBy     PaymentOverpaymentsDBFieldNameType
	MetaUpdatedAt     PaymentOverpaymentsDBFieldNameType
	MetaUpdatedBy     PaymentOverpaymentsDBFieldNameType
	MetaDeletedAt     PaymentOverpaymentsDBFieldNameType
	MetaDeletedBy     PaymentOverpaymentsDBFieldNameType
}

var PaymentOverpaymentsDBFieldName = paymentOverpaymentsDBFieldName{
	Id:                "id",
	PaymentIntentId:   "payment_intent_id",
	PaidAttemptId:     "paid_attempt_id",
	OverpaidAttemptId: "overpaid_attempt_id",
	ExpectedAmount:    "expected_amount",
	ReceivedAmount:    "received_amount",
	OverpaidAmount:    "overpaid_amount",
	Currency:          "currency",
	Status:            "status",
	ResolutionAction:  "resolution_action",
	ResolutionNotes:   "resolution_notes",
	ResolvedAt:        "resolved_at",
	ResolvedBy:        "resolved_by",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPaymentOverpaymentsDBFieldNameFromStr(field string) (dbField PaymentOverpaymentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentOverpaymentsDBFieldName.Id):
		return PaymentOverpaymentsDBFieldName.Id, true

	case string(PaymentOverpaymentsDBFieldName.PaymentIntentId):
		return PaymentOverpaymentsDBFieldName.PaymentIntentId, true

	case string(PaymentOverpaymentsDBFieldName.PaidAttemptId):
		return PaymentOverpaymentsDBFieldName.PaidAttemptId, true

	case string(PaymentOverpaymentsDBFieldName.OverpaidAttemptId):
		return PaymentOverpaymentsDBFieldName.OverpaidAttemptId, true

	case string(PaymentOverpaymentsDBFieldName.ExpectedAmount):
		return PaymentOverpaymentsDBFieldName.ExpectedAmount, true

	case string(PaymentOverpaymentsDBFieldName.ReceivedAmount):
		return PaymentOverpaymentsDBFieldName.ReceivedAmount, true

	case string(PaymentOverpaymentsDBFieldName.OverpaidAmount):
		return PaymentOverpaymentsDBFieldName.OverpaidAmount, true

	case string(PaymentOverpaymentsDBFieldName.Currency):
		return PaymentOverpaymentsDBFieldName.Currency, true

	case string(PaymentOverpaymentsDBFieldName.Status):
		return PaymentOverpaymentsDBFieldName.Status, true

	case string(PaymentOverpaymentsDBFieldName.ResolutionAction):
		return PaymentOverpaymentsDBFieldName.ResolutionAction, true

	case string(PaymentOverpaymentsDBFieldName.ResolutionNotes):
		return PaymentOverpaymentsDBFieldName.ResolutionNotes, true

	case string(PaymentOverpaymentsDBFieldName.ResolvedAt):
		return PaymentOverpaymentsDBFieldName.ResolvedAt, true

	case string(PaymentOverpaymentsDBFieldName.ResolvedBy):
		return PaymentOverpaymentsDBFieldName.ResolvedBy, true

	case string(PaymentOverpaymentsDBFieldName.Metadata):
		return PaymentOverpaymentsDBFieldName.Metadata, true

	case string(PaymentOverpaymentsDBFieldName.MetaCreatedAt):
		return PaymentOverpaymentsDBFieldName.MetaCreatedAt, true

	case string(PaymentOverpaymentsDBFieldName.MetaCreatedBy):
		return PaymentOverpaymentsDBFieldName.MetaCreatedBy, true

	case string(PaymentOverpaymentsDBFieldName.MetaUpdatedAt):
		return PaymentOverpaymentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentOverpaymentsDBFieldName.MetaUpdatedBy):
		return PaymentOverpaymentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentOverpaymentsDBFieldName.MetaDeletedAt):
		return PaymentOverpaymentsDBFieldName.MetaDeletedAt, true

	case string(PaymentOverpaymentsDBFieldName.MetaDeletedBy):
		return PaymentOverpaymentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentOverpaymentsFilterJoins = map[string]JoinSpec{}

var PaymentOverpaymentsFilterFields = map[string]FilterFieldSpec{
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
	"paid_attempt_id": {
		SourcePath:        "paid_attempt_id",
		DefaultOutputPath: "paidAttemptId",
		Column:            "paid_attempt_id",
		SQLAlias:          "paid_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"overpaid_attempt_id": {
		SourcePath:        "overpaid_attempt_id",
		DefaultOutputPath: "overpaidAttemptId",
		Column:            "overpaid_attempt_id",
		SQLAlias:          "overpaid_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"expected_amount": {
		SourcePath:        "expected_amount",
		DefaultOutputPath: "expectedAmount",
		Column:            "expected_amount",
		SQLAlias:          "expected_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"received_amount": {
		SourcePath:        "received_amount",
		DefaultOutputPath: "receivedAmount",
		Column:            "received_amount",
		SQLAlias:          "received_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"overpaid_amount": {
		SourcePath:        "overpaid_amount",
		DefaultOutputPath: "overpaidAmount",
		Column:            "overpaid_amount",
		SQLAlias:          "overpaid_amount",
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
	"resolution_action": {
		SourcePath:        "resolution_action",
		DefaultOutputPath: "resolutionAction",
		Column:            "resolution_action",
		SQLAlias:          "resolution_action",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resolution_notes": {
		SourcePath:        "resolution_notes",
		DefaultOutputPath: "resolutionNotes",
		Column:            "resolution_notes",
		SQLAlias:          "resolution_notes",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resolved_at": {
		SourcePath:        "resolved_at",
		DefaultOutputPath: "resolvedAt",
		Column:            "resolved_at",
		SQLAlias:          "resolved_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resolved_by": {
		SourcePath:        "resolved_by",
		DefaultOutputPath: "resolvedBy",
		Column:            "resolved_by",
		SQLAlias:          "resolved_by",
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

func NewPaymentOverpaymentsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentOverpaymentsFilterFields[field]
	return
}

type PaymentOverpaymentsFilterResult struct {
	PaymentOverpayments
	FilterCount int `db:"count"`
}

func ValidatePaymentOverpaymentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentOverpaymentsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentOverpaymentsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentOverpaymentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentOverpaymentsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentOverpaymentsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentOverpaymentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentOverpaymentsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentOverpayments struct {
	Id                uuid.UUID       `db:"id"`
	PaymentIntentId   uuid.UUID       `db:"payment_intent_id"`
	PaidAttemptId     nuuid.NUUID     `db:"paid_attempt_id"`
	OverpaidAttemptId nuuid.NUUID     `db:"overpaid_attempt_id"`
	ExpectedAmount    decimal.Decimal `db:"expected_amount"`
	ReceivedAmount    decimal.Decimal `db:"received_amount"`
	OverpaidAmount    decimal.Decimal `db:"overpaid_amount"`
	Currency          string          `db:"currency"`
	Status            string          `db:"status"`
	ResolutionAction  null.String     `db:"resolution_action"`
	ResolutionNotes   null.String     `db:"resolution_notes"`
	ResolvedAt        null.Time       `db:"resolved_at"`
	ResolvedBy        nuuid.NUUID     `db:"resolved_by"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PaymentOverpaymentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentOverpayments) ToPaymentOverpaymentsPrimaryID() PaymentOverpaymentsPrimaryID {
	return PaymentOverpaymentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentOverpaymentsList []*PaymentOverpayments

type PaymentOverpaymentsFilterResultList []*PaymentOverpaymentsFilterResult

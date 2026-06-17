package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type PaymentPlansDBFieldNameType string

type paymentPlansDBFieldName struct {
	Id                        PaymentPlansDBFieldNameType
	PaymentIntentId           PaymentPlansDBFieldNameType
	PlanType                  PaymentPlansDBFieldNameType
	Status                    PaymentPlansDBFieldNameType
	TotalAmount               PaymentPlansDBFieldNameType
	Currency                  PaymentPlansDBFieldNameType
	InstallmentCount          PaymentPlansDBFieldNameType
	DepositAmount             PaymentPlansDBFieldNameType
	AutoCancelOnDefault       PaymentPlansDBFieldNameType
	DefaultGracePeriodSeconds PaymentPlansDBFieldNameType
	CompletedAt               PaymentPlansDBFieldNameType
	CanceledAt                PaymentPlansDBFieldNameType
	Metadata                  PaymentPlansDBFieldNameType
	MetaCreatedAt             PaymentPlansDBFieldNameType
	MetaCreatedBy             PaymentPlansDBFieldNameType
	MetaUpdatedAt             PaymentPlansDBFieldNameType
	MetaUpdatedBy             PaymentPlansDBFieldNameType
	MetaDeletedAt             PaymentPlansDBFieldNameType
	MetaDeletedBy             PaymentPlansDBFieldNameType
}

var PaymentPlansDBFieldName = paymentPlansDBFieldName{
	Id:                        "id",
	PaymentIntentId:           "payment_intent_id",
	PlanType:                  "plan_type",
	Status:                    "status",
	TotalAmount:               "total_amount",
	Currency:                  "currency",
	InstallmentCount:          "installment_count",
	DepositAmount:             "deposit_amount",
	AutoCancelOnDefault:       "auto_cancel_on_default",
	DefaultGracePeriodSeconds: "default_grace_period_seconds",
	CompletedAt:               "completed_at",
	CanceledAt:                "canceled_at",
	Metadata:                  "metadata",
	MetaCreatedAt:             "meta_created_at",
	MetaCreatedBy:             "meta_created_by",
	MetaUpdatedAt:             "meta_updated_at",
	MetaUpdatedBy:             "meta_updated_by",
	MetaDeletedAt:             "meta_deleted_at",
	MetaDeletedBy:             "meta_deleted_by",
}

func NewPaymentPlansDBFieldNameFromStr(field string) (dbField PaymentPlansDBFieldNameType, found bool) {
	switch field {

	case string(PaymentPlansDBFieldName.Id):
		return PaymentPlansDBFieldName.Id, true

	case string(PaymentPlansDBFieldName.PaymentIntentId):
		return PaymentPlansDBFieldName.PaymentIntentId, true

	case string(PaymentPlansDBFieldName.PlanType):
		return PaymentPlansDBFieldName.PlanType, true

	case string(PaymentPlansDBFieldName.Status):
		return PaymentPlansDBFieldName.Status, true

	case string(PaymentPlansDBFieldName.TotalAmount):
		return PaymentPlansDBFieldName.TotalAmount, true

	case string(PaymentPlansDBFieldName.Currency):
		return PaymentPlansDBFieldName.Currency, true

	case string(PaymentPlansDBFieldName.InstallmentCount):
		return PaymentPlansDBFieldName.InstallmentCount, true

	case string(PaymentPlansDBFieldName.DepositAmount):
		return PaymentPlansDBFieldName.DepositAmount, true

	case string(PaymentPlansDBFieldName.AutoCancelOnDefault):
		return PaymentPlansDBFieldName.AutoCancelOnDefault, true

	case string(PaymentPlansDBFieldName.DefaultGracePeriodSeconds):
		return PaymentPlansDBFieldName.DefaultGracePeriodSeconds, true

	case string(PaymentPlansDBFieldName.CompletedAt):
		return PaymentPlansDBFieldName.CompletedAt, true

	case string(PaymentPlansDBFieldName.CanceledAt):
		return PaymentPlansDBFieldName.CanceledAt, true

	case string(PaymentPlansDBFieldName.Metadata):
		return PaymentPlansDBFieldName.Metadata, true

	case string(PaymentPlansDBFieldName.MetaCreatedAt):
		return PaymentPlansDBFieldName.MetaCreatedAt, true

	case string(PaymentPlansDBFieldName.MetaCreatedBy):
		return PaymentPlansDBFieldName.MetaCreatedBy, true

	case string(PaymentPlansDBFieldName.MetaUpdatedAt):
		return PaymentPlansDBFieldName.MetaUpdatedAt, true

	case string(PaymentPlansDBFieldName.MetaUpdatedBy):
		return PaymentPlansDBFieldName.MetaUpdatedBy, true

	case string(PaymentPlansDBFieldName.MetaDeletedAt):
		return PaymentPlansDBFieldName.MetaDeletedAt, true

	case string(PaymentPlansDBFieldName.MetaDeletedBy):
		return PaymentPlansDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentPlansFilterJoins = map[string]JoinSpec{}

var PaymentPlansFilterFields = map[string]FilterFieldSpec{
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
	"plan_type": {
		SourcePath:        "plan_type",
		DefaultOutputPath: "planType",
		Column:            "plan_type",
		SQLAlias:          "plan_type",
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
	"total_amount": {
		SourcePath:        "total_amount",
		DefaultOutputPath: "totalAmount",
		Column:            "total_amount",
		SQLAlias:          "total_amount",
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
	"installment_count": {
		SourcePath:        "installment_count",
		DefaultOutputPath: "installmentCount",
		Column:            "installment_count",
		SQLAlias:          "installment_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"deposit_amount": {
		SourcePath:        "deposit_amount",
		DefaultOutputPath: "depositAmount",
		Column:            "deposit_amount",
		SQLAlias:          "deposit_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"auto_cancel_on_default": {
		SourcePath:        "auto_cancel_on_default",
		DefaultOutputPath: "autoCancelOnDefault",
		Column:            "auto_cancel_on_default",
		SQLAlias:          "auto_cancel_on_default",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"default_grace_period_seconds": {
		SourcePath:        "default_grace_period_seconds",
		DefaultOutputPath: "defaultGracePeriodSeconds",
		Column:            "default_grace_period_seconds",
		SQLAlias:          "default_grace_period_seconds",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"completed_at": {
		SourcePath:        "completed_at",
		DefaultOutputPath: "completedAt",
		Column:            "completed_at",
		SQLAlias:          "completed_at",
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

func NewPaymentPlansFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentPlansFilterFields[field]
	return
}

type PaymentPlansFilterResult struct {
	PaymentPlans
	FilterCount int `db:"count"`
}

func ValidatePaymentPlansFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentPlansFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentPlansFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentPlansFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentPlansFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentPlansFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentPlansFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentPlansFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentPlanStatus string

const (
	PaymentPlanStatusActive    PaymentPlanStatus = "active"
	PaymentPlanStatusCompleted PaymentPlanStatus = "completed"
	PaymentPlanStatusCanceled  PaymentPlanStatus = "canceled"
)

type PaymentPlanType string

const (
	PaymentPlanTypeInstallment  PaymentPlanType = "installment"
	PaymentPlanTypeSubscription PaymentPlanType = "subscription"
)

type PaymentPlans struct {
	Id                        uuid.UUID           `db:"id"`
	PaymentIntentId           uuid.UUID           `db:"payment_intent_id"`
	PlanType                  PaymentPlanType     `db:"plan_type"`
	Status                    PaymentPlanStatus   `db:"status"`
	TotalAmount               decimal.Decimal     `db:"total_amount"`
	Currency                  string              `db:"currency"`
	InstallmentCount          int                 `db:"installment_count"`
	DepositAmount             decimal.NullDecimal `db:"deposit_amount"`
	AutoCancelOnDefault       bool                `db:"auto_cancel_on_default"`
	DefaultGracePeriodSeconds int                 `db:"default_grace_period_seconds"`
	CompletedAt               null.Time           `db:"completed_at"`
	CanceledAt                null.Time           `db:"canceled_at"`
	Metadata                  json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type PaymentPlansPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentPlans) ToPaymentPlansPrimaryID() PaymentPlansPrimaryID {
	return PaymentPlansPrimaryID{
		Id: d.Id,
	}
}

type PaymentPlansList []*PaymentPlans

type PaymentPlansFilterResultList []*PaymentPlansFilterResult

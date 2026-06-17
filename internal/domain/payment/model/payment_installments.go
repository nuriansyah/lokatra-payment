package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type PaymentInstallmentsDBFieldNameType string

type paymentInstallmentsDBFieldName struct {
	Id              PaymentInstallmentsDBFieldNameType
	PaymentPlanId   PaymentInstallmentsDBFieldNameType
	PaymentIntentId PaymentInstallmentsDBFieldNameType
	InstallmentNo   PaymentInstallmentsDBFieldNameType
	DueAmount       PaymentInstallmentsDBFieldNameType
	PaidAmount      PaymentInstallmentsDBFieldNameType
	Currency        PaymentInstallmentsDBFieldNameType
	DueAt           PaymentInstallmentsDBFieldNameType
	Status          PaymentInstallmentsDBFieldNameType
	PaidAt          PaymentInstallmentsDBFieldNameType
	OverdueAt       PaymentInstallmentsDBFieldNameType
	Metadata        PaymentInstallmentsDBFieldNameType
	MetaCreatedAt   PaymentInstallmentsDBFieldNameType
	MetaCreatedBy   PaymentInstallmentsDBFieldNameType
	MetaUpdatedAt   PaymentInstallmentsDBFieldNameType
	MetaUpdatedBy   PaymentInstallmentsDBFieldNameType
	MetaDeletedAt   PaymentInstallmentsDBFieldNameType
	MetaDeletedBy   PaymentInstallmentsDBFieldNameType
}

var PaymentInstallmentsDBFieldName = paymentInstallmentsDBFieldName{
	Id:              "id",
	PaymentPlanId:   "payment_plan_id",
	PaymentIntentId: "payment_intent_id",
	InstallmentNo:   "installment_no",
	DueAmount:       "due_amount",
	PaidAmount:      "paid_amount",
	Currency:        "currency",
	DueAt:           "due_at",
	Status:          "status",
	PaidAt:          "paid_at",
	OverdueAt:       "overdue_at",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewPaymentInstallmentsDBFieldNameFromStr(field string) (dbField PaymentInstallmentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentInstallmentsDBFieldName.Id):
		return PaymentInstallmentsDBFieldName.Id, true

	case string(PaymentInstallmentsDBFieldName.PaymentPlanId):
		return PaymentInstallmentsDBFieldName.PaymentPlanId, true

	case string(PaymentInstallmentsDBFieldName.PaymentIntentId):
		return PaymentInstallmentsDBFieldName.PaymentIntentId, true

	case string(PaymentInstallmentsDBFieldName.InstallmentNo):
		return PaymentInstallmentsDBFieldName.InstallmentNo, true

	case string(PaymentInstallmentsDBFieldName.DueAmount):
		return PaymentInstallmentsDBFieldName.DueAmount, true

	case string(PaymentInstallmentsDBFieldName.PaidAmount):
		return PaymentInstallmentsDBFieldName.PaidAmount, true

	case string(PaymentInstallmentsDBFieldName.Currency):
		return PaymentInstallmentsDBFieldName.Currency, true

	case string(PaymentInstallmentsDBFieldName.DueAt):
		return PaymentInstallmentsDBFieldName.DueAt, true

	case string(PaymentInstallmentsDBFieldName.Status):
		return PaymentInstallmentsDBFieldName.Status, true

	case string(PaymentInstallmentsDBFieldName.PaidAt):
		return PaymentInstallmentsDBFieldName.PaidAt, true

	case string(PaymentInstallmentsDBFieldName.OverdueAt):
		return PaymentInstallmentsDBFieldName.OverdueAt, true

	case string(PaymentInstallmentsDBFieldName.Metadata):
		return PaymentInstallmentsDBFieldName.Metadata, true

	case string(PaymentInstallmentsDBFieldName.MetaCreatedAt):
		return PaymentInstallmentsDBFieldName.MetaCreatedAt, true

	case string(PaymentInstallmentsDBFieldName.MetaCreatedBy):
		return PaymentInstallmentsDBFieldName.MetaCreatedBy, true

	case string(PaymentInstallmentsDBFieldName.MetaUpdatedAt):
		return PaymentInstallmentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentInstallmentsDBFieldName.MetaUpdatedBy):
		return PaymentInstallmentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentInstallmentsDBFieldName.MetaDeletedAt):
		return PaymentInstallmentsDBFieldName.MetaDeletedAt, true

	case string(PaymentInstallmentsDBFieldName.MetaDeletedBy):
		return PaymentInstallmentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentInstallmentsFilterJoins = map[string]JoinSpec{}

var PaymentInstallmentsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_plan_id": {
		SourcePath:        "payment_plan_id",
		DefaultOutputPath: "paymentPlanId",
		Column:            "payment_plan_id",
		SQLAlias:          "payment_plan_id",
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
	"installment_no": {
		SourcePath:        "installment_no",
		DefaultOutputPath: "installmentNo",
		Column:            "installment_no",
		SQLAlias:          "installment_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"due_amount": {
		SourcePath:        "due_amount",
		DefaultOutputPath: "dueAmount",
		Column:            "due_amount",
		SQLAlias:          "due_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"paid_amount": {
		SourcePath:        "paid_amount",
		DefaultOutputPath: "paidAmount",
		Column:            "paid_amount",
		SQLAlias:          "paid_amount",
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
	"due_at": {
		SourcePath:        "due_at",
		DefaultOutputPath: "dueAt",
		Column:            "due_at",
		SQLAlias:          "due_at",
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
	"paid_at": {
		SourcePath:        "paid_at",
		DefaultOutputPath: "paidAt",
		Column:            "paid_at",
		SQLAlias:          "paid_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"overdue_at": {
		SourcePath:        "overdue_at",
		DefaultOutputPath: "overdueAt",
		Column:            "overdue_at",
		SQLAlias:          "overdue_at",
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

func NewPaymentInstallmentsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentInstallmentsFilterFields[field]
	return
}

type PaymentInstallmentsFilterResult struct {
	PaymentInstallments
	FilterCount int `db:"count"`
}

func ValidatePaymentInstallmentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentInstallmentsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentInstallmentsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentInstallmentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentInstallmentsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentInstallmentsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentInstallmentsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentInstallmentsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentInstallmentStatus string

const (
	PaymentInstallmentStatusPending  PaymentInstallmentStatus = "pending"
	PaymentInstallmentStatusPaid     PaymentInstallmentStatus = "paid"
	PaymentInstallmentStatusOverdue  PaymentInstallmentStatus = "overdue"
	PaymentInstallmentStatusCanceled PaymentInstallmentStatus = "canceled"
)

type PaymentInstallments struct {
	Id              uuid.UUID                `db:"id"`
	PaymentPlanId   uuid.UUID                `db:"payment_plan_id"`
	PaymentIntentId uuid.UUID                `db:"payment_intent_id"`
	InstallmentNo   int                      `db:"installment_no"`
	DueAmount       decimal.Decimal          `db:"due_amount"`
	PaidAmount      decimal.Decimal          `db:"paid_amount"`
	Currency        string                   `db:"currency"`
	DueAt           time.Time                `db:"due_at"`
	Status          PaymentInstallmentStatus `db:"status"`
	PaidAt          null.Time                `db:"paid_at"`
	OverdueAt       null.Time                `db:"overdue_at"`
	Metadata        json.RawMessage          `db:"metadata"`

	shared.MetaSignature
}
type PaymentInstallmentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentInstallments) ToPaymentInstallmentsPrimaryID() PaymentInstallmentsPrimaryID {
	return PaymentInstallmentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentInstallmentsList []*PaymentInstallments

type PaymentInstallmentsFilterResultList []*PaymentInstallmentsFilterResult

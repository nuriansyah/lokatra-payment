package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentMethodsDBFieldNameType string

type paymentMethodsDBFieldName struct {
	Id            PaymentMethodsDBFieldNameType
	Code          PaymentMethodsDBFieldNameType
	MethodType    PaymentMethodsDBFieldNameType
	Name          PaymentMethodsDBFieldNameType
	Status        PaymentMethodsDBFieldNameType
	Metadata      PaymentMethodsDBFieldNameType
	MetaCreatedAt PaymentMethodsDBFieldNameType
	MetaCreatedBy PaymentMethodsDBFieldNameType
	MetaUpdatedAt PaymentMethodsDBFieldNameType
	MetaUpdatedBy PaymentMethodsDBFieldNameType
	MetaDeletedAt PaymentMethodsDBFieldNameType
	MetaDeletedBy PaymentMethodsDBFieldNameType
}

var PaymentMethodsDBFieldName = paymentMethodsDBFieldName{
	Id:            "id",
	Code:          "code",
	MethodType:    "method_type",
	Name:          "name",
	Status:        "status",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewPaymentMethodsDBFieldNameFromStr(field string) (dbField PaymentMethodsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentMethodsDBFieldName.Id):
		return PaymentMethodsDBFieldName.Id, true

	case string(PaymentMethodsDBFieldName.Code):
		return PaymentMethodsDBFieldName.Code, true

	case string(PaymentMethodsDBFieldName.MethodType):
		return PaymentMethodsDBFieldName.MethodType, true

	case string(PaymentMethodsDBFieldName.Name):
		return PaymentMethodsDBFieldName.Name, true

	case string(PaymentMethodsDBFieldName.Status):
		return PaymentMethodsDBFieldName.Status, true

	case string(PaymentMethodsDBFieldName.Metadata):
		return PaymentMethodsDBFieldName.Metadata, true

	case string(PaymentMethodsDBFieldName.MetaCreatedAt):
		return PaymentMethodsDBFieldName.MetaCreatedAt, true

	case string(PaymentMethodsDBFieldName.MetaCreatedBy):
		return PaymentMethodsDBFieldName.MetaCreatedBy, true

	case string(PaymentMethodsDBFieldName.MetaUpdatedAt):
		return PaymentMethodsDBFieldName.MetaUpdatedAt, true

	case string(PaymentMethodsDBFieldName.MetaUpdatedBy):
		return PaymentMethodsDBFieldName.MetaUpdatedBy, true

	case string(PaymentMethodsDBFieldName.MetaDeletedAt):
		return PaymentMethodsDBFieldName.MetaDeletedAt, true

	case string(PaymentMethodsDBFieldName.MetaDeletedBy):
		return PaymentMethodsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentMethodsFilterJoins = map[string]JoinSpec{}

var PaymentMethodsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"code": {
		SourcePath:        "code",
		DefaultOutputPath: "code",
		Column:            "code",
		SQLAlias:          "code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"method_type": {
		SourcePath:        "method_type",
		DefaultOutputPath: "methodType",
		Column:            "method_type",
		SQLAlias:          "method_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"name": {
		SourcePath:        "name",
		DefaultOutputPath: "name",
		Column:            "name",
		SQLAlias:          "name",
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

func NewPaymentMethodsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentMethodsFilterFields[field]
	return
}

type PaymentMethodsFilterResult struct {
	PaymentMethods
	FilterCount int `db:"count"`
}

func ValidatePaymentMethodsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentMethodsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentMethodsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentMethodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentMethodsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentMethodsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentMethodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentMethodsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentMethodStatus string

const (
	PaymentMethodStatusActive     PaymentMethodStatus = "active"
	PaymentMethodStatusInactive   PaymentMethodStatus = "inactive"
	PaymentMethodStatusDeprecated PaymentMethodStatus = "deprecated"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCard         PaymentMethodType = "card"
	PaymentMethodTypeBankTransfer PaymentMethodType = "bank_transfer"
	PaymentMethodTypeEwallet      PaymentMethodType = "ewallet"
	PaymentMethodTypeCash         PaymentMethodType = "cash"
	PaymentMethodTypeCod          PaymentMethodType = "cod"
)

type PaymentMethods struct {
	Id         uuid.UUID           `db:"id"`
	Code       string              `db:"code"`
	MethodType PaymentMethodType   `db:"method_type"`
	Name       string              `db:"name"`
	Status     PaymentMethodStatus `db:"status"`
	Metadata   json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type PaymentMethodsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentMethods) ToPaymentMethodsPrimaryID() PaymentMethodsPrimaryID {
	return PaymentMethodsPrimaryID{
		Id: d.Id,
	}
}

type PaymentMethodsList []*PaymentMethods

type PaymentMethodsFilterResultList []*PaymentMethodsFilterResult

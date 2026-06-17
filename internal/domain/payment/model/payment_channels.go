package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentChannelsDBFieldNameType string

type paymentChannelsDBFieldName struct {
	Id            PaymentChannelsDBFieldNameType
	MethodId      PaymentChannelsDBFieldNameType
	Code          PaymentChannelsDBFieldNameType
	Name          PaymentChannelsDBFieldNameType
	CountryCode   PaymentChannelsDBFieldNameType
	Currency      PaymentChannelsDBFieldNameType
	Status        PaymentChannelsDBFieldNameType
	Metadata      PaymentChannelsDBFieldNameType
	MetaCreatedAt PaymentChannelsDBFieldNameType
	MetaCreatedBy PaymentChannelsDBFieldNameType
	MetaUpdatedAt PaymentChannelsDBFieldNameType
	MetaUpdatedBy PaymentChannelsDBFieldNameType
	MetaDeletedAt PaymentChannelsDBFieldNameType
	MetaDeletedBy PaymentChannelsDBFieldNameType
}

var PaymentChannelsDBFieldName = paymentChannelsDBFieldName{
	Id:            "id",
	MethodId:      "method_id",
	Code:          "code",
	Name:          "name",
	CountryCode:   "country_code",
	Currency:      "currency",
	Status:        "status",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewPaymentChannelsDBFieldNameFromStr(field string) (dbField PaymentChannelsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentChannelsDBFieldName.Id):
		return PaymentChannelsDBFieldName.Id, true

	case string(PaymentChannelsDBFieldName.MethodId):
		return PaymentChannelsDBFieldName.MethodId, true

	case string(PaymentChannelsDBFieldName.Code):
		return PaymentChannelsDBFieldName.Code, true

	case string(PaymentChannelsDBFieldName.Name):
		return PaymentChannelsDBFieldName.Name, true

	case string(PaymentChannelsDBFieldName.CountryCode):
		return PaymentChannelsDBFieldName.CountryCode, true

	case string(PaymentChannelsDBFieldName.Currency):
		return PaymentChannelsDBFieldName.Currency, true

	case string(PaymentChannelsDBFieldName.Status):
		return PaymentChannelsDBFieldName.Status, true

	case string(PaymentChannelsDBFieldName.Metadata):
		return PaymentChannelsDBFieldName.Metadata, true

	case string(PaymentChannelsDBFieldName.MetaCreatedAt):
		return PaymentChannelsDBFieldName.MetaCreatedAt, true

	case string(PaymentChannelsDBFieldName.MetaCreatedBy):
		return PaymentChannelsDBFieldName.MetaCreatedBy, true

	case string(PaymentChannelsDBFieldName.MetaUpdatedAt):
		return PaymentChannelsDBFieldName.MetaUpdatedAt, true

	case string(PaymentChannelsDBFieldName.MetaUpdatedBy):
		return PaymentChannelsDBFieldName.MetaUpdatedBy, true

	case string(PaymentChannelsDBFieldName.MetaDeletedAt):
		return PaymentChannelsDBFieldName.MetaDeletedAt, true

	case string(PaymentChannelsDBFieldName.MetaDeletedBy):
		return PaymentChannelsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentChannelsFilterJoins = map[string]JoinSpec{}

var PaymentChannelsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"method_id": {
		SourcePath:        "method_id",
		DefaultOutputPath: "methodId",
		Column:            "method_id",
		SQLAlias:          "method_id",
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
	"name": {
		SourcePath:        "name",
		DefaultOutputPath: "name",
		Column:            "name",
		SQLAlias:          "name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"country_code": {
		SourcePath:        "country_code",
		DefaultOutputPath: "countryCode",
		Column:            "country_code",
		SQLAlias:          "country_code",
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

func NewPaymentChannelsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentChannelsFilterFields[field]
	return
}

type PaymentChannelsFilterResult struct {
	PaymentChannels
	FilterCount int `db:"count"`
}

func ValidatePaymentChannelsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentChannelsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentChannelsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentChannelsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentChannelsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentChannelsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentChannelsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentChannelsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentChannelStatus string

const (
	PaymentChannelStatusActive     PaymentChannelStatus = "active"
	PaymentChannelStatusInactive   PaymentChannelStatus = "inactive"
	PaymentChannelStatusDeprecated PaymentChannelStatus = "deprecated"
)

type PaymentChannels struct {
	Id          uuid.UUID            `db:"id"`
	MethodId    uuid.UUID            `db:"method_id"`
	Code        string               `db:"code"`
	Name        string               `db:"name"`
	CountryCode string               `db:"country_code"`
	Currency    string               `db:"currency"`
	Status      PaymentChannelStatus `db:"status"`
	Metadata    json.RawMessage      `db:"metadata"`

	shared.MetaSignature
}
type PaymentChannelsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentChannels) ToPaymentChannelsPrimaryID() PaymentChannelsPrimaryID {
	return PaymentChannelsPrimaryID{
		Id: d.Id,
	}
}

type PaymentChannelsList []*PaymentChannels

type PaymentChannelsFilterResultList []*PaymentChannelsFilterResult

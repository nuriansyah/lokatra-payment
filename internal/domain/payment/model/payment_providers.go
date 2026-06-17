package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentProvidersDBFieldNameType string

type paymentProvidersDBFieldName struct {
	Id                    PaymentProvidersDBFieldNameType
	Code                  PaymentProvidersDBFieldNameType
	Name                  PaymentProvidersDBFieldNameType
	ProviderType          PaymentProvidersDBFieldNameType
	Status                PaymentProvidersDBFieldNameType
	SupportsRefund        PaymentProvidersDBFieldNameType
	SupportsPartialRefund PaymentProvidersDBFieldNameType
	SupportsAuthorization PaymentProvidersDBFieldNameType
	SupportsCapture       PaymentProvidersDBFieldNameType
	SupportsVoid          PaymentProvidersDBFieldNameType
	SupportsWebhook       PaymentProvidersDBFieldNameType
	Metadata              PaymentProvidersDBFieldNameType
	MetaCreatedAt         PaymentProvidersDBFieldNameType
	MetaCreatedBy         PaymentProvidersDBFieldNameType
	MetaUpdatedAt         PaymentProvidersDBFieldNameType
	MetaUpdatedBy         PaymentProvidersDBFieldNameType
	MetaDeletedAt         PaymentProvidersDBFieldNameType
	MetaDeletedBy         PaymentProvidersDBFieldNameType
}

var PaymentProvidersDBFieldName = paymentProvidersDBFieldName{
	Id:                    "id",
	Code:                  "code",
	Name:                  "name",
	ProviderType:          "provider_type",
	Status:                "status",
	SupportsRefund:        "supports_refund",
	SupportsPartialRefund: "supports_partial_refund",
	SupportsAuthorization: "supports_authorization",
	SupportsCapture:       "supports_capture",
	SupportsVoid:          "supports_void",
	SupportsWebhook:       "supports_webhook",
	Metadata:              "metadata",
	MetaCreatedAt:         "meta_created_at",
	MetaCreatedBy:         "meta_created_by",
	MetaUpdatedAt:         "meta_updated_at",
	MetaUpdatedBy:         "meta_updated_by",
	MetaDeletedAt:         "meta_deleted_at",
	MetaDeletedBy:         "meta_deleted_by",
}

func NewPaymentProvidersDBFieldNameFromStr(field string) (dbField PaymentProvidersDBFieldNameType, found bool) {
	switch field {

	case string(PaymentProvidersDBFieldName.Id):
		return PaymentProvidersDBFieldName.Id, true

	case string(PaymentProvidersDBFieldName.Code):
		return PaymentProvidersDBFieldName.Code, true

	case string(PaymentProvidersDBFieldName.Name):
		return PaymentProvidersDBFieldName.Name, true

	case string(PaymentProvidersDBFieldName.ProviderType):
		return PaymentProvidersDBFieldName.ProviderType, true

	case string(PaymentProvidersDBFieldName.Status):
		return PaymentProvidersDBFieldName.Status, true

	case string(PaymentProvidersDBFieldName.SupportsRefund):
		return PaymentProvidersDBFieldName.SupportsRefund, true

	case string(PaymentProvidersDBFieldName.SupportsPartialRefund):
		return PaymentProvidersDBFieldName.SupportsPartialRefund, true

	case string(PaymentProvidersDBFieldName.SupportsAuthorization):
		return PaymentProvidersDBFieldName.SupportsAuthorization, true

	case string(PaymentProvidersDBFieldName.SupportsCapture):
		return PaymentProvidersDBFieldName.SupportsCapture, true

	case string(PaymentProvidersDBFieldName.SupportsVoid):
		return PaymentProvidersDBFieldName.SupportsVoid, true

	case string(PaymentProvidersDBFieldName.SupportsWebhook):
		return PaymentProvidersDBFieldName.SupportsWebhook, true

	case string(PaymentProvidersDBFieldName.Metadata):
		return PaymentProvidersDBFieldName.Metadata, true

	case string(PaymentProvidersDBFieldName.MetaCreatedAt):
		return PaymentProvidersDBFieldName.MetaCreatedAt, true

	case string(PaymentProvidersDBFieldName.MetaCreatedBy):
		return PaymentProvidersDBFieldName.MetaCreatedBy, true

	case string(PaymentProvidersDBFieldName.MetaUpdatedAt):
		return PaymentProvidersDBFieldName.MetaUpdatedAt, true

	case string(PaymentProvidersDBFieldName.MetaUpdatedBy):
		return PaymentProvidersDBFieldName.MetaUpdatedBy, true

	case string(PaymentProvidersDBFieldName.MetaDeletedAt):
		return PaymentProvidersDBFieldName.MetaDeletedAt, true

	case string(PaymentProvidersDBFieldName.MetaDeletedBy):
		return PaymentProvidersDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentProvidersFilterJoins = map[string]JoinSpec{}

var PaymentProvidersFilterFields = map[string]FilterFieldSpec{
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
	"name": {
		SourcePath:        "name",
		DefaultOutputPath: "name",
		Column:            "name",
		SQLAlias:          "name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_type": {
		SourcePath:        "provider_type",
		DefaultOutputPath: "providerType",
		Column:            "provider_type",
		SQLAlias:          "provider_type",
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
	"supports_refund": {
		SourcePath:        "supports_refund",
		DefaultOutputPath: "supportsRefund",
		Column:            "supports_refund",
		SQLAlias:          "supports_refund",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"supports_partial_refund": {
		SourcePath:        "supports_partial_refund",
		DefaultOutputPath: "supportsPartialRefund",
		Column:            "supports_partial_refund",
		SQLAlias:          "supports_partial_refund",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"supports_authorization": {
		SourcePath:        "supports_authorization",
		DefaultOutputPath: "supportsAuthorization",
		Column:            "supports_authorization",
		SQLAlias:          "supports_authorization",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"supports_capture": {
		SourcePath:        "supports_capture",
		DefaultOutputPath: "supportsCapture",
		Column:            "supports_capture",
		SQLAlias:          "supports_capture",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"supports_void": {
		SourcePath:        "supports_void",
		DefaultOutputPath: "supportsVoid",
		Column:            "supports_void",
		SQLAlias:          "supports_void",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"supports_webhook": {
		SourcePath:        "supports_webhook",
		DefaultOutputPath: "supportsWebhook",
		Column:            "supports_webhook",
		SQLAlias:          "supports_webhook",
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

func NewPaymentProvidersFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentProvidersFilterFields[field]
	return
}

type PaymentProvidersFilterResult struct {
	PaymentProviders
	FilterCount int `db:"count"`
}

func ValidatePaymentProvidersFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentProvidersFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentProvidersFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentProvidersFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentProvidersFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentProvidersFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentProvidersFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentProvidersFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderStatus string

const (
	ProviderStatusActive     ProviderStatus = "active"
	ProviderStatusInactive   ProviderStatus = "inactive"
	ProviderStatusDeprecated ProviderStatus = "deprecated"
)

type ProviderType string

const (
	ProviderTypeGateway      ProviderType = "gateway"
	ProviderTypeWallet       ProviderType = "wallet"
	ProviderTypeBankTransfer ProviderType = "bank_transfer"
	ProviderTypeCash         ProviderType = "cash"
	ProviderTypeCod          ProviderType = "cod"
)

type PaymentProviders struct {
	Id                    uuid.UUID       `db:"id"`
	Code                  string          `db:"code"`
	Name                  string          `db:"name"`
	ProviderType          ProviderType    `db:"provider_type"`
	Status                ProviderStatus  `db:"status"`
	SupportsRefund        bool            `db:"supports_refund"`
	SupportsPartialRefund bool            `db:"supports_partial_refund"`
	SupportsAuthorization bool            `db:"supports_authorization"`
	SupportsCapture       bool            `db:"supports_capture"`
	SupportsVoid          bool            `db:"supports_void"`
	SupportsWebhook       bool            `db:"supports_webhook"`
	Metadata              json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PaymentProvidersPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentProviders) ToPaymentProvidersPrimaryID() PaymentProvidersPrimaryID {
	return PaymentProvidersPrimaryID{
		Id: d.Id,
	}
}

type PaymentProvidersList []*PaymentProviders

type PaymentProvidersFilterResultList []*PaymentProvidersFilterResult

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

type PaymentRouteCandidatesRuntimeDBFieldNameType string

type paymentRouteCandidatesRuntimeDBFieldName struct {
	Id                  PaymentRouteCandidatesRuntimeDBFieldNameType
	ScopeType           PaymentRouteCandidatesRuntimeDBFieldNameType
	ScopeId             PaymentRouteCandidatesRuntimeDBFieldNameType
	MerchantId          PaymentRouteCandidatesRuntimeDBFieldNameType
	MethodCode          PaymentRouteCandidatesRuntimeDBFieldNameType
	ChannelCode         PaymentRouteCandidatesRuntimeDBFieldNameType
	Currency            PaymentRouteCandidatesRuntimeDBFieldNameType
	MinAmount           PaymentRouteCandidatesRuntimeDBFieldNameType
	MaxAmount           PaymentRouteCandidatesRuntimeDBFieldNameType
	ProviderAccountId   PaymentRouteCandidatesRuntimeDBFieldNameType
	ProviderMethodCode  PaymentRouteCandidatesRuntimeDBFieldNameType
	ProviderChannelCode PaymentRouteCandidatesRuntimeDBFieldNameType
	Priority            PaymentRouteCandidatesRuntimeDBFieldNameType
	IsFallback          PaymentRouteCandidatesRuntimeDBFieldNameType
	TrafficWeight       PaymentRouteCandidatesRuntimeDBFieldNameType
	TimeoutMs           PaymentRouteCandidatesRuntimeDBFieldNameType
	MaxAttempts         PaymentRouteCandidatesRuntimeDBFieldNameType
	IsEnabled           PaymentRouteCandidatesRuntimeDBFieldNameType
	ConditionExpr       PaymentRouteCandidatesRuntimeDBFieldNameType
	Metadata            PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaCreatedAt       PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaCreatedBy       PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaUpdatedAt       PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaUpdatedBy       PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaDeletedAt       PaymentRouteCandidatesRuntimeDBFieldNameType
	MetaDeletedBy       PaymentRouteCandidatesRuntimeDBFieldNameType
}

var PaymentRouteCandidatesRuntimeDBFieldName = paymentRouteCandidatesRuntimeDBFieldName{
	Id:                  "id",
	ScopeType:           "scope_type",
	ScopeId:             "scope_id",
	MerchantId:          "merchant_id",
	MethodCode:          "method_code",
	ChannelCode:         "channel_code",
	Currency:            "currency",
	MinAmount:           "min_amount",
	MaxAmount:           "max_amount",
	ProviderAccountId:   "provider_account_id",
	ProviderMethodCode:  "provider_method_code",
	ProviderChannelCode: "provider_channel_code",
	Priority:            "priority",
	IsFallback:          "is_fallback",
	TrafficWeight:       "traffic_weight",
	TimeoutMs:           "timeout_ms",
	MaxAttempts:         "max_attempts",
	IsEnabled:           "is_enabled",
	ConditionExpr:       "condition_expr",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewPaymentRouteCandidatesRuntimeDBFieldNameFromStr(field string) (dbField PaymentRouteCandidatesRuntimeDBFieldNameType, found bool) {
	switch field {

	case string(PaymentRouteCandidatesRuntimeDBFieldName.Id):
		return PaymentRouteCandidatesRuntimeDBFieldName.Id, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ScopeType):
		return PaymentRouteCandidatesRuntimeDBFieldName.ScopeType, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ScopeId):
		return PaymentRouteCandidatesRuntimeDBFieldName.ScopeId, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MerchantId):
		return PaymentRouteCandidatesRuntimeDBFieldName.MerchantId, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MethodCode):
		return PaymentRouteCandidatesRuntimeDBFieldName.MethodCode, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ChannelCode):
		return PaymentRouteCandidatesRuntimeDBFieldName.ChannelCode, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.Currency):
		return PaymentRouteCandidatesRuntimeDBFieldName.Currency, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MinAmount):
		return PaymentRouteCandidatesRuntimeDBFieldName.MinAmount, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MaxAmount):
		return PaymentRouteCandidatesRuntimeDBFieldName.MaxAmount, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ProviderAccountId):
		return PaymentRouteCandidatesRuntimeDBFieldName.ProviderAccountId, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ProviderMethodCode):
		return PaymentRouteCandidatesRuntimeDBFieldName.ProviderMethodCode, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ProviderChannelCode):
		return PaymentRouteCandidatesRuntimeDBFieldName.ProviderChannelCode, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.Priority):
		return PaymentRouteCandidatesRuntimeDBFieldName.Priority, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.IsFallback):
		return PaymentRouteCandidatesRuntimeDBFieldName.IsFallback, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.TrafficWeight):
		return PaymentRouteCandidatesRuntimeDBFieldName.TrafficWeight, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.TimeoutMs):
		return PaymentRouteCandidatesRuntimeDBFieldName.TimeoutMs, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MaxAttempts):
		return PaymentRouteCandidatesRuntimeDBFieldName.MaxAttempts, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.IsEnabled):
		return PaymentRouteCandidatesRuntimeDBFieldName.IsEnabled, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.ConditionExpr):
		return PaymentRouteCandidatesRuntimeDBFieldName.ConditionExpr, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.Metadata):
		return PaymentRouteCandidatesRuntimeDBFieldName.Metadata, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedAt):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedAt, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedBy):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedBy, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedAt):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedAt, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedBy):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedBy, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedAt):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedAt, true

	case string(PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedBy):
		return PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentRouteCandidatesRuntimeFilterJoins = map[string]JoinSpec{}

var PaymentRouteCandidatesRuntimeFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"scope_type": {
		SourcePath:        "scope_type",
		DefaultOutputPath: "scopeType",
		Column:            "scope_type",
		SQLAlias:          "scope_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"scope_id": {
		SourcePath:        "scope_id",
		DefaultOutputPath: "scopeId",
		Column:            "scope_id",
		SQLAlias:          "scope_id",
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
	"method_code": {
		SourcePath:        "method_code",
		DefaultOutputPath: "methodCode",
		Column:            "method_code",
		SQLAlias:          "method_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"channel_code": {
		SourcePath:        "channel_code",
		DefaultOutputPath: "channelCode",
		Column:            "channel_code",
		SQLAlias:          "channel_code",
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
	"min_amount": {
		SourcePath:        "min_amount",
		DefaultOutputPath: "minAmount",
		Column:            "min_amount",
		SQLAlias:          "min_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"max_amount": {
		SourcePath:        "max_amount",
		DefaultOutputPath: "maxAmount",
		Column:            "max_amount",
		SQLAlias:          "max_amount",
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
	"provider_method_code": {
		SourcePath:        "provider_method_code",
		DefaultOutputPath: "providerMethodCode",
		Column:            "provider_method_code",
		SQLAlias:          "provider_method_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_channel_code": {
		SourcePath:        "provider_channel_code",
		DefaultOutputPath: "providerChannelCode",
		Column:            "provider_channel_code",
		SQLAlias:          "provider_channel_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"priority": {
		SourcePath:        "priority",
		DefaultOutputPath: "priority",
		Column:            "priority",
		SQLAlias:          "priority",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_fallback": {
		SourcePath:        "is_fallback",
		DefaultOutputPath: "isFallback",
		Column:            "is_fallback",
		SQLAlias:          "is_fallback",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"traffic_weight": {
		SourcePath:        "traffic_weight",
		DefaultOutputPath: "trafficWeight",
		Column:            "traffic_weight",
		SQLAlias:          "traffic_weight",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"timeout_ms": {
		SourcePath:        "timeout_ms",
		DefaultOutputPath: "timeoutMs",
		Column:            "timeout_ms",
		SQLAlias:          "timeout_ms",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"max_attempts": {
		SourcePath:        "max_attempts",
		DefaultOutputPath: "maxAttempts",
		Column:            "max_attempts",
		SQLAlias:          "max_attempts",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_enabled": {
		SourcePath:        "is_enabled",
		DefaultOutputPath: "isEnabled",
		Column:            "is_enabled",
		SQLAlias:          "is_enabled",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"condition_expr": {
		SourcePath:        "condition_expr",
		DefaultOutputPath: "conditionExpr",
		Column:            "condition_expr",
		SQLAlias:          "condition_expr",
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

func NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentRouteCandidatesRuntimeFilterFields[field]
	return
}

type PaymentRouteCandidatesRuntimeFilterResult struct {
	PaymentRouteCandidatesRuntime
	FilterCount int `db:"count"`
}

func ValidatePaymentRouteCandidatesRuntimeFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentRouteCandidatesRuntimeFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentRouteCandidatesRuntimeFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentRouteCandidatesRuntimeFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ScopeType string

const (
	ScopeTypePlatform ScopeType = "platform"
	ScopeTypeMerchant ScopeType = "merchant"
)

type PaymentRouteCandidatesRuntime struct {
	Id                  uuid.UUID           `db:"id"`
	ScopeType           ScopeType           `db:"scope_type"`
	ScopeId             nuuid.NUUID         `db:"scope_id"`
	MerchantId          nuuid.NUUID         `db:"merchant_id"`
	MethodCode          string              `db:"method_code"`
	ChannelCode         null.String         `db:"channel_code"`
	Currency            string              `db:"currency"`
	MinAmount           decimal.NullDecimal `db:"min_amount"`
	MaxAmount           decimal.NullDecimal `db:"max_amount"`
	ProviderAccountId   uuid.UUID           `db:"provider_account_id"`
	ProviderMethodCode  null.String         `db:"provider_method_code"`
	ProviderChannelCode null.String         `db:"provider_channel_code"`
	Priority            int                 `db:"priority"`
	IsFallback          bool                `db:"is_fallback"`
	TrafficWeight       int                 `db:"traffic_weight"`
	TimeoutMs           int                 `db:"timeout_ms"`
	MaxAttempts         int                 `db:"max_attempts"`
	IsEnabled           bool                `db:"is_enabled"`
	ConditionExpr       json.RawMessage     `db:"condition_expr"`
	Metadata            json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type PaymentRouteCandidatesRuntimePrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentRouteCandidatesRuntime) ToPaymentRouteCandidatesRuntimePrimaryID() PaymentRouteCandidatesRuntimePrimaryID {
	return PaymentRouteCandidatesRuntimePrimaryID{
		Id: d.Id,
	}
}

type PaymentRouteCandidatesRuntimeList []*PaymentRouteCandidatesRuntime

type PaymentRouteCandidatesRuntimeFilterResultList []*PaymentRouteCandidatesRuntimeFilterResult

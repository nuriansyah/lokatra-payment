package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PaymentRouteDecisionsDBFieldNameType string

type paymentRouteDecisionsDBFieldName struct {
	Id                        PaymentRouteDecisionsDBFieldNameType
	PaymentIntentId           PaymentRouteDecisionsDBFieldNameType
	SelectedProviderAccountId PaymentRouteDecisionsDBFieldNameType
	SelectedProviderCode      PaymentRouteDecisionsDBFieldNameType
	MethodCode                PaymentRouteDecisionsDBFieldNameType
	ChannelCode               PaymentRouteDecisionsDBFieldNameType
	Reason                    PaymentRouteDecisionsDBFieldNameType
	EvaluatedContext          PaymentRouteDecisionsDBFieldNameType
	Candidates                PaymentRouteDecisionsDBFieldNameType
	Metadata                  PaymentRouteDecisionsDBFieldNameType
	MetaCreatedAt             PaymentRouteDecisionsDBFieldNameType
	MetaCreatedBy             PaymentRouteDecisionsDBFieldNameType
	MetaUpdatedAt             PaymentRouteDecisionsDBFieldNameType
	MetaUpdatedBy             PaymentRouteDecisionsDBFieldNameType
	MetaDeletedAt             PaymentRouteDecisionsDBFieldNameType
	MetaDeletedBy             PaymentRouteDecisionsDBFieldNameType
}

var PaymentRouteDecisionsDBFieldName = paymentRouteDecisionsDBFieldName{
	Id:                        "id",
	PaymentIntentId:           "payment_intent_id",
	SelectedProviderAccountId: "selected_provider_account_id",
	SelectedProviderCode:      "selected_provider_code",
	MethodCode:                "method_code",
	ChannelCode:               "channel_code",
	Reason:                    "reason",
	EvaluatedContext:          "evaluated_context",
	Candidates:                "candidates",
	Metadata:                  "metadata",
	MetaCreatedAt:             "meta_created_at",
	MetaCreatedBy:             "meta_created_by",
	MetaUpdatedAt:             "meta_updated_at",
	MetaUpdatedBy:             "meta_updated_by",
	MetaDeletedAt:             "meta_deleted_at",
	MetaDeletedBy:             "meta_deleted_by",
}

func NewPaymentRouteDecisionsDBFieldNameFromStr(field string) (dbField PaymentRouteDecisionsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentRouteDecisionsDBFieldName.Id):
		return PaymentRouteDecisionsDBFieldName.Id, true

	case string(PaymentRouteDecisionsDBFieldName.PaymentIntentId):
		return PaymentRouteDecisionsDBFieldName.PaymentIntentId, true

	case string(PaymentRouteDecisionsDBFieldName.SelectedProviderAccountId):
		return PaymentRouteDecisionsDBFieldName.SelectedProviderAccountId, true

	case string(PaymentRouteDecisionsDBFieldName.SelectedProviderCode):
		return PaymentRouteDecisionsDBFieldName.SelectedProviderCode, true

	case string(PaymentRouteDecisionsDBFieldName.MethodCode):
		return PaymentRouteDecisionsDBFieldName.MethodCode, true

	case string(PaymentRouteDecisionsDBFieldName.ChannelCode):
		return PaymentRouteDecisionsDBFieldName.ChannelCode, true

	case string(PaymentRouteDecisionsDBFieldName.Reason):
		return PaymentRouteDecisionsDBFieldName.Reason, true

	case string(PaymentRouteDecisionsDBFieldName.EvaluatedContext):
		return PaymentRouteDecisionsDBFieldName.EvaluatedContext, true

	case string(PaymentRouteDecisionsDBFieldName.Candidates):
		return PaymentRouteDecisionsDBFieldName.Candidates, true

	case string(PaymentRouteDecisionsDBFieldName.Metadata):
		return PaymentRouteDecisionsDBFieldName.Metadata, true

	case string(PaymentRouteDecisionsDBFieldName.MetaCreatedAt):
		return PaymentRouteDecisionsDBFieldName.MetaCreatedAt, true

	case string(PaymentRouteDecisionsDBFieldName.MetaCreatedBy):
		return PaymentRouteDecisionsDBFieldName.MetaCreatedBy, true

	case string(PaymentRouteDecisionsDBFieldName.MetaUpdatedAt):
		return PaymentRouteDecisionsDBFieldName.MetaUpdatedAt, true

	case string(PaymentRouteDecisionsDBFieldName.MetaUpdatedBy):
		return PaymentRouteDecisionsDBFieldName.MetaUpdatedBy, true

	case string(PaymentRouteDecisionsDBFieldName.MetaDeletedAt):
		return PaymentRouteDecisionsDBFieldName.MetaDeletedAt, true

	case string(PaymentRouteDecisionsDBFieldName.MetaDeletedBy):
		return PaymentRouteDecisionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentRouteDecisionsFilterJoins = map[string]JoinSpec{}

var PaymentRouteDecisionsFilterFields = map[string]FilterFieldSpec{
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
	"selected_provider_account_id": {
		SourcePath:        "selected_provider_account_id",
		DefaultOutputPath: "selectedProviderAccountId",
		Column:            "selected_provider_account_id",
		SQLAlias:          "selected_provider_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"selected_provider_code": {
		SourcePath:        "selected_provider_code",
		DefaultOutputPath: "selectedProviderCode",
		Column:            "selected_provider_code",
		SQLAlias:          "selected_provider_code",
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
	"reason": {
		SourcePath:        "reason",
		DefaultOutputPath: "reason",
		Column:            "reason",
		SQLAlias:          "reason",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"evaluated_context": {
		SourcePath:        "evaluated_context",
		DefaultOutputPath: "evaluatedContext",
		Column:            "evaluated_context",
		SQLAlias:          "evaluated_context",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"candidates": {
		SourcePath:        "candidates",
		DefaultOutputPath: "candidates",
		Column:            "candidates",
		SQLAlias:          "candidates",
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

func NewPaymentRouteDecisionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentRouteDecisionsFilterFields[field]
	return
}

type PaymentRouteDecisionsFilterResult struct {
	PaymentRouteDecisions
	FilterCount int `db:"count"`
}

func ValidatePaymentRouteDecisionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentRouteDecisionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentRouteDecisionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentRouteDecisionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentRouteDecisionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentRouteDecisionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentRouteDecisionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentRouteDecisionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentRouteDecisions struct {
	Id                        uuid.UUID       `db:"id"`
	PaymentIntentId           uuid.UUID       `db:"payment_intent_id"`
	SelectedProviderAccountId nuuid.NUUID     `db:"selected_provider_account_id"`
	SelectedProviderCode      null.String     `db:"selected_provider_code"`
	MethodCode                string          `db:"method_code"`
	ChannelCode               null.String     `db:"channel_code"`
	Reason                    string          `db:"reason"`
	EvaluatedContext          json.RawMessage `db:"evaluated_context"`
	Candidates                json.RawMessage `db:"candidates"`
	Metadata                  json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PaymentRouteDecisionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentRouteDecisions) ToPaymentRouteDecisionsPrimaryID() PaymentRouteDecisionsPrimaryID {
	return PaymentRouteDecisionsPrimaryID{
		Id: d.Id,
	}
}

type PaymentRouteDecisionsList []*PaymentRouteDecisions

type PaymentRouteDecisionsFilterResultList []*PaymentRouteDecisionsFilterResult

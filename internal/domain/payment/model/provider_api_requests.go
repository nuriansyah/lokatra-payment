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

type ProviderApiRequestsDBFieldNameType string

type providerApiRequestsDBFieldName struct {
	Id                ProviderApiRequestsDBFieldNameType
	ProviderAccountId ProviderApiRequestsDBFieldNameType
	PaymentIntentId   ProviderApiRequestsDBFieldNameType
	PaymentAttemptId  ProviderApiRequestsDBFieldNameType
	Operation         ProviderApiRequestsDBFieldNameType
	IdempotencyKey    ProviderApiRequestsDBFieldNameType
	RequestMethod     ProviderApiRequestsDBFieldNameType
	RequestUrl        ProviderApiRequestsDBFieldNameType
	RequestHeaders    ProviderApiRequestsDBFieldNameType
	RequestBody       ProviderApiRequestsDBFieldNameType
	ResponseStatus    ProviderApiRequestsDBFieldNameType
	ResponseHeaders   ProviderApiRequestsDBFieldNameType
	ResponseBody      ProviderApiRequestsDBFieldNameType
	LatencyMs         ProviderApiRequestsDBFieldNameType
	Success           ProviderApiRequestsDBFieldNameType
	ErrorCode         ProviderApiRequestsDBFieldNameType
	ErrorMessage      ProviderApiRequestsDBFieldNameType
	Metadata          ProviderApiRequestsDBFieldNameType
	MetaCreatedAt     ProviderApiRequestsDBFieldNameType
	MetaCreatedBy     ProviderApiRequestsDBFieldNameType
	MetaUpdatedAt     ProviderApiRequestsDBFieldNameType
	MetaUpdatedBy     ProviderApiRequestsDBFieldNameType
	MetaDeletedAt     ProviderApiRequestsDBFieldNameType
	MetaDeletedBy     ProviderApiRequestsDBFieldNameType
}

var ProviderApiRequestsDBFieldName = providerApiRequestsDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	PaymentIntentId:   "payment_intent_id",
	PaymentAttemptId:  "payment_attempt_id",
	Operation:         "operation",
	IdempotencyKey:    "idempotency_key",
	RequestMethod:     "request_method",
	RequestUrl:        "request_url",
	RequestHeaders:    "request_headers",
	RequestBody:       "request_body",
	ResponseStatus:    "response_status",
	ResponseHeaders:   "response_headers",
	ResponseBody:      "response_body",
	LatencyMs:         "latency_ms",
	Success:           "success",
	ErrorCode:         "error_code",
	ErrorMessage:      "error_message",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewProviderApiRequestsDBFieldNameFromStr(field string) (dbField ProviderApiRequestsDBFieldNameType, found bool) {
	switch field {

	case string(ProviderApiRequestsDBFieldName.Id):
		return ProviderApiRequestsDBFieldName.Id, true

	case string(ProviderApiRequestsDBFieldName.ProviderAccountId):
		return ProviderApiRequestsDBFieldName.ProviderAccountId, true

	case string(ProviderApiRequestsDBFieldName.PaymentIntentId):
		return ProviderApiRequestsDBFieldName.PaymentIntentId, true

	case string(ProviderApiRequestsDBFieldName.PaymentAttemptId):
		return ProviderApiRequestsDBFieldName.PaymentAttemptId, true

	case string(ProviderApiRequestsDBFieldName.Operation):
		return ProviderApiRequestsDBFieldName.Operation, true

	case string(ProviderApiRequestsDBFieldName.IdempotencyKey):
		return ProviderApiRequestsDBFieldName.IdempotencyKey, true

	case string(ProviderApiRequestsDBFieldName.RequestMethod):
		return ProviderApiRequestsDBFieldName.RequestMethod, true

	case string(ProviderApiRequestsDBFieldName.RequestUrl):
		return ProviderApiRequestsDBFieldName.RequestUrl, true

	case string(ProviderApiRequestsDBFieldName.RequestHeaders):
		return ProviderApiRequestsDBFieldName.RequestHeaders, true

	case string(ProviderApiRequestsDBFieldName.RequestBody):
		return ProviderApiRequestsDBFieldName.RequestBody, true

	case string(ProviderApiRequestsDBFieldName.ResponseStatus):
		return ProviderApiRequestsDBFieldName.ResponseStatus, true

	case string(ProviderApiRequestsDBFieldName.ResponseHeaders):
		return ProviderApiRequestsDBFieldName.ResponseHeaders, true

	case string(ProviderApiRequestsDBFieldName.ResponseBody):
		return ProviderApiRequestsDBFieldName.ResponseBody, true

	case string(ProviderApiRequestsDBFieldName.LatencyMs):
		return ProviderApiRequestsDBFieldName.LatencyMs, true

	case string(ProviderApiRequestsDBFieldName.Success):
		return ProviderApiRequestsDBFieldName.Success, true

	case string(ProviderApiRequestsDBFieldName.ErrorCode):
		return ProviderApiRequestsDBFieldName.ErrorCode, true

	case string(ProviderApiRequestsDBFieldName.ErrorMessage):
		return ProviderApiRequestsDBFieldName.ErrorMessage, true

	case string(ProviderApiRequestsDBFieldName.Metadata):
		return ProviderApiRequestsDBFieldName.Metadata, true

	case string(ProviderApiRequestsDBFieldName.MetaCreatedAt):
		return ProviderApiRequestsDBFieldName.MetaCreatedAt, true

	case string(ProviderApiRequestsDBFieldName.MetaCreatedBy):
		return ProviderApiRequestsDBFieldName.MetaCreatedBy, true

	case string(ProviderApiRequestsDBFieldName.MetaUpdatedAt):
		return ProviderApiRequestsDBFieldName.MetaUpdatedAt, true

	case string(ProviderApiRequestsDBFieldName.MetaUpdatedBy):
		return ProviderApiRequestsDBFieldName.MetaUpdatedBy, true

	case string(ProviderApiRequestsDBFieldName.MetaDeletedAt):
		return ProviderApiRequestsDBFieldName.MetaDeletedAt, true

	case string(ProviderApiRequestsDBFieldName.MetaDeletedBy):
		return ProviderApiRequestsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderApiRequestsFilterJoins = map[string]JoinSpec{}

var ProviderApiRequestsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"payment_intent_id": {
		SourcePath:        "payment_intent_id",
		DefaultOutputPath: "paymentIntentId",
		Column:            "payment_intent_id",
		SQLAlias:          "payment_intent_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"operation": {
		SourcePath:        "operation",
		DefaultOutputPath: "operation",
		Column:            "operation",
		SQLAlias:          "operation",
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
	"request_method": {
		SourcePath:        "request_method",
		DefaultOutputPath: "requestMethod",
		Column:            "request_method",
		SQLAlias:          "request_method",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"request_url": {
		SourcePath:        "request_url",
		DefaultOutputPath: "requestUrl",
		Column:            "request_url",
		SQLAlias:          "request_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"request_headers": {
		SourcePath:        "request_headers",
		DefaultOutputPath: "requestHeaders",
		Column:            "request_headers",
		SQLAlias:          "request_headers",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"request_body": {
		SourcePath:        "request_body",
		DefaultOutputPath: "requestBody",
		Column:            "request_body",
		SQLAlias:          "request_body",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"response_status": {
		SourcePath:        "response_status",
		DefaultOutputPath: "responseStatus",
		Column:            "response_status",
		SQLAlias:          "response_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"response_headers": {
		SourcePath:        "response_headers",
		DefaultOutputPath: "responseHeaders",
		Column:            "response_headers",
		SQLAlias:          "response_headers",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"response_body": {
		SourcePath:        "response_body",
		DefaultOutputPath: "responseBody",
		Column:            "response_body",
		SQLAlias:          "response_body",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"latency_ms": {
		SourcePath:        "latency_ms",
		DefaultOutputPath: "latencyMs",
		Column:            "latency_ms",
		SQLAlias:          "latency_ms",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"success": {
		SourcePath:        "success",
		DefaultOutputPath: "success",
		Column:            "success",
		SQLAlias:          "success",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_code": {
		SourcePath:        "error_code",
		DefaultOutputPath: "errorCode",
		Column:            "error_code",
		SQLAlias:          "error_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_message": {
		SourcePath:        "error_message",
		DefaultOutputPath: "errorMessage",
		Column:            "error_message",
		SQLAlias:          "error_message",
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

func NewProviderApiRequestsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderApiRequestsFilterFields[field]
	return
}

type ProviderApiRequestsFilterResult struct {
	ProviderApiRequests
	FilterCount int `db:"count"`
}

func ValidateProviderApiRequestsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderApiRequestsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderApiRequestsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderApiRequestsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderApiRequestsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderApiRequestsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderApiRequestsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderApiRequestsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderApiRequests struct {
	Id                uuid.UUID       `db:"id"`
	ProviderAccountId nuuid.NUUID     `db:"provider_account_id"`
	PaymentIntentId   nuuid.NUUID     `db:"payment_intent_id"`
	PaymentAttemptId  nuuid.NUUID     `db:"payment_attempt_id"`
	Operation         string          `db:"operation"`
	IdempotencyKey    null.String     `db:"idempotency_key"`
	RequestMethod     string          `db:"request_method"`
	RequestUrl        string          `db:"request_url"`
	RequestHeaders    json.RawMessage `db:"request_headers"`
	RequestBody       json.RawMessage `db:"request_body"`
	ResponseStatus    null.Int        `db:"response_status"`
	ResponseHeaders   json.RawMessage `db:"response_headers"`
	ResponseBody      json.RawMessage `db:"response_body"`
	LatencyMs         null.Int        `db:"latency_ms"`
	Success           null.Bool       `db:"success"`
	ErrorCode         null.String     `db:"error_code"`
	ErrorMessage      null.String     `db:"error_message"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ProviderApiRequestsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderApiRequests) ToProviderApiRequestsPrimaryID() ProviderApiRequestsPrimaryID {
	return ProviderApiRequestsPrimaryID{
		Id: d.Id,
	}
}

type ProviderApiRequestsList []*ProviderApiRequests

type ProviderApiRequestsFilterResultList []*ProviderApiRequestsFilterResult

package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentAttemptsDTOFieldNameType string

type paymentAttemptsDTOFieldName struct {
	Id                    PaymentAttemptsDTOFieldNameType
	PaymentIntentId       PaymentAttemptsDTOFieldNameType
	AttemptNo             PaymentAttemptsDTOFieldNameType
	ProviderAccountId     PaymentAttemptsDTOFieldNameType
	RouteDecisionId       PaymentAttemptsDTOFieldNameType
	ProviderCode          PaymentAttemptsDTOFieldNameType
	MethodCode            PaymentAttemptsDTOFieldNameType
	ChannelCode           PaymentAttemptsDTOFieldNameType
	Amount                PaymentAttemptsDTOFieldNameType
	Currency              PaymentAttemptsDTOFieldNameType
	Status                PaymentAttemptsDTOFieldNameType
	ProviderReference     PaymentAttemptsDTOFieldNameType
	ProviderTransactionId PaymentAttemptsDTOFieldNameType
	ProviderOrderId       PaymentAttemptsDTOFieldNameType
	ProviderPaymentId     PaymentAttemptsDTOFieldNameType
	FailureCode           PaymentAttemptsDTOFieldNameType
	FailureMessage        PaymentAttemptsDTOFieldNameType
	ExpiresAt             PaymentAttemptsDTOFieldNameType
	AuthorizedAt          PaymentAttemptsDTOFieldNameType
	CapturedAt            PaymentAttemptsDTOFieldNameType
	PaidAt                PaymentAttemptsDTOFieldNameType
	FailedAt              PaymentAttemptsDTOFieldNameType
	CanceledAt            PaymentAttemptsDTOFieldNameType
	StatusSyncRequiredAt  PaymentAttemptsDTOFieldNameType
	LastStatusSyncAt      PaymentAttemptsDTOFieldNameType
	RawRequest            PaymentAttemptsDTOFieldNameType
	RawResponse           PaymentAttemptsDTOFieldNameType
	Metadata              PaymentAttemptsDTOFieldNameType
	MetaCreatedAt         PaymentAttemptsDTOFieldNameType
	MetaCreatedBy         PaymentAttemptsDTOFieldNameType
	MetaUpdatedAt         PaymentAttemptsDTOFieldNameType
	MetaUpdatedBy         PaymentAttemptsDTOFieldNameType
	MetaDeletedAt         PaymentAttemptsDTOFieldNameType
	MetaDeletedBy         PaymentAttemptsDTOFieldNameType
}

var PaymentAttemptsDTOFieldName = paymentAttemptsDTOFieldName{
	Id:                    "id",
	PaymentIntentId:       "paymentIntentId",
	AttemptNo:             "attemptNo",
	ProviderAccountId:     "providerAccountId",
	RouteDecisionId:       "routeDecisionId",
	ProviderCode:          "providerCode",
	MethodCode:            "methodCode",
	ChannelCode:           "channelCode",
	Amount:                "amount",
	Currency:              "currency",
	Status:                "status",
	ProviderReference:     "providerReference",
	ProviderTransactionId: "providerTransactionId",
	ProviderOrderId:       "providerOrderId",
	ProviderPaymentId:     "providerPaymentId",
	FailureCode:           "failureCode",
	FailureMessage:        "failureMessage",
	ExpiresAt:             "expiresAt",
	AuthorizedAt:          "authorizedAt",
	CapturedAt:            "capturedAt",
	PaidAt:                "paidAt",
	FailedAt:              "failedAt",
	CanceledAt:            "canceledAt",
	StatusSyncRequiredAt:  "statusSyncRequiredAt",
	LastStatusSyncAt:      "lastStatusSyncAt",
	RawRequest:            "rawRequest",
	RawResponse:           "rawResponse",
	Metadata:              "metadata",
	MetaCreatedAt:         "metaCreatedAt",
	MetaCreatedBy:         "metaCreatedBy",
	MetaUpdatedAt:         "metaUpdatedAt",
	MetaUpdatedBy:         "metaUpdatedBy",
	MetaDeletedAt:         "metaDeletedAt",
	MetaDeletedBy:         "metaDeletedBy",
}

func transformPaymentAttemptsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentAttemptsDTOFieldName.Id):
		return string(model.PaymentAttemptsDBFieldName.Id), true

	case string(PaymentAttemptsDTOFieldName.PaymentIntentId):
		return string(model.PaymentAttemptsDBFieldName.PaymentIntentId), true

	case string(PaymentAttemptsDTOFieldName.AttemptNo):
		return string(model.PaymentAttemptsDBFieldName.AttemptNo), true

	case string(PaymentAttemptsDTOFieldName.ProviderAccountId):
		return string(model.PaymentAttemptsDBFieldName.ProviderAccountId), true

	case string(PaymentAttemptsDTOFieldName.RouteDecisionId):
		return string(model.PaymentAttemptsDBFieldName.RouteDecisionId), true

	case string(PaymentAttemptsDTOFieldName.ProviderCode):
		return string(model.PaymentAttemptsDBFieldName.ProviderCode), true

	case string(PaymentAttemptsDTOFieldName.MethodCode):
		return string(model.PaymentAttemptsDBFieldName.MethodCode), true

	case string(PaymentAttemptsDTOFieldName.ChannelCode):
		return string(model.PaymentAttemptsDBFieldName.ChannelCode), true

	case string(PaymentAttemptsDTOFieldName.Amount):
		return string(model.PaymentAttemptsDBFieldName.Amount), true

	case string(PaymentAttemptsDTOFieldName.Currency):
		return string(model.PaymentAttemptsDBFieldName.Currency), true

	case string(PaymentAttemptsDTOFieldName.Status):
		return string(model.PaymentAttemptsDBFieldName.Status), true

	case string(PaymentAttemptsDTOFieldName.ProviderReference):
		return string(model.PaymentAttemptsDBFieldName.ProviderReference), true

	case string(PaymentAttemptsDTOFieldName.ProviderTransactionId):
		return string(model.PaymentAttemptsDBFieldName.ProviderTransactionId), true

	case string(PaymentAttemptsDTOFieldName.ProviderOrderId):
		return string(model.PaymentAttemptsDBFieldName.ProviderOrderId), true

	case string(PaymentAttemptsDTOFieldName.ProviderPaymentId):
		return string(model.PaymentAttemptsDBFieldName.ProviderPaymentId), true

	case string(PaymentAttemptsDTOFieldName.FailureCode):
		return string(model.PaymentAttemptsDBFieldName.FailureCode), true

	case string(PaymentAttemptsDTOFieldName.FailureMessage):
		return string(model.PaymentAttemptsDBFieldName.FailureMessage), true

	case string(PaymentAttemptsDTOFieldName.ExpiresAt):
		return string(model.PaymentAttemptsDBFieldName.ExpiresAt), true

	case string(PaymentAttemptsDTOFieldName.AuthorizedAt):
		return string(model.PaymentAttemptsDBFieldName.AuthorizedAt), true

	case string(PaymentAttemptsDTOFieldName.CapturedAt):
		return string(model.PaymentAttemptsDBFieldName.CapturedAt), true

	case string(PaymentAttemptsDTOFieldName.PaidAt):
		return string(model.PaymentAttemptsDBFieldName.PaidAt), true

	case string(PaymentAttemptsDTOFieldName.FailedAt):
		return string(model.PaymentAttemptsDBFieldName.FailedAt), true

	case string(PaymentAttemptsDTOFieldName.CanceledAt):
		return string(model.PaymentAttemptsDBFieldName.CanceledAt), true

	case string(PaymentAttemptsDTOFieldName.StatusSyncRequiredAt):
		return string(model.PaymentAttemptsDBFieldName.StatusSyncRequiredAt), true

	case string(PaymentAttemptsDTOFieldName.LastStatusSyncAt):
		return string(model.PaymentAttemptsDBFieldName.LastStatusSyncAt), true

	case string(PaymentAttemptsDTOFieldName.RawRequest):
		return string(model.PaymentAttemptsDBFieldName.RawRequest), true

	case string(PaymentAttemptsDTOFieldName.RawResponse):
		return string(model.PaymentAttemptsDBFieldName.RawResponse), true

	case string(PaymentAttemptsDTOFieldName.Metadata):
		return string(model.PaymentAttemptsDBFieldName.Metadata), true

	case string(PaymentAttemptsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentAttemptsDBFieldName.MetaCreatedAt), true

	case string(PaymentAttemptsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentAttemptsDBFieldName.MetaCreatedBy), true

	case string(PaymentAttemptsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentAttemptsDBFieldName.MetaUpdatedAt), true

	case string(PaymentAttemptsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentAttemptsDBFieldName.MetaUpdatedBy), true

	case string(PaymentAttemptsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentAttemptsDBFieldName.MetaDeletedAt), true

	case string(PaymentAttemptsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentAttemptsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentAttemptsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentAttemptsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentAttemptsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentAttemptsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentAttemptsProjectionOutputPath(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return failure.BadRequest(fmt.Errorf("field alias cannot be empty"))
	}
	if !strings.Contains(path, ".") {
		return nil
	}
	for _, part := range strings.Split(path, ".") {
		if strings.TrimSpace(part) == "" {
			return failure.BadRequest(fmt.Errorf("field alias %s is invalid", path))
		}
	}
	return nil
}

func transformPaymentAttemptsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentAttemptsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentAttemptsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentAttemptsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentAttemptsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentAttemptsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentAttemptsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentAttemptsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentAttemptsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentAttemptsFilter(filter *model.Filter) {
	if filter.Pagination.Strategy == "" {
		if filter.Pagination.Page > 0 {
			filter.Pagination.Strategy = model.PaginationStrategyOffset
		} else {
			filter.Pagination.Strategy = model.DefaultPaginationStrategy
		}
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = model.DefaultPageSize
	}

	if filter.Pagination.PageSize > model.MaxPageSize {
		filter.Pagination.PageSize = model.MaxPageSize
	}

	if filter.Pagination.Strategy == model.PaginationStrategyOffset && filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.Direction == "" {
		filter.Pagination.Direction = model.CursorDirectionNext
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PaymentAttemptsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentAttemptsSelectableResponse map[string]interface{}
type PaymentAttemptsSelectableListResponse []*PaymentAttemptsSelectableResponse

func assignPaymentAttemptsNestedValue(out map[string]interface{}, path string, value interface{}) {
	parts := strings.Split(path, ".")
	current := out
	for _, part := range parts[:len(parts)-1] {
		next, ok := current[part].(map[string]interface{})
		if !ok {
			next = map[string]interface{}{}
			current[part] = next
		}
		current = next
	}
	current[parts[len(parts)-1]] = value
}

func setPaymentAttemptsSelectableValue(out PaymentAttemptsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentAttemptsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentAttemptsSelectableResponse(paymentAttempts model.PaymentAttempts, filter model.Filter) PaymentAttemptsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentAttemptsDBFieldName.Id),
			string(model.PaymentAttemptsDBFieldName.PaymentIntentId),
			string(model.PaymentAttemptsDBFieldName.AttemptNo),
			string(model.PaymentAttemptsDBFieldName.ProviderAccountId),
			string(model.PaymentAttemptsDBFieldName.RouteDecisionId),
			string(model.PaymentAttemptsDBFieldName.ProviderCode),
			string(model.PaymentAttemptsDBFieldName.MethodCode),
			string(model.PaymentAttemptsDBFieldName.ChannelCode),
			string(model.PaymentAttemptsDBFieldName.Amount),
			string(model.PaymentAttemptsDBFieldName.Currency),
			string(model.PaymentAttemptsDBFieldName.Status),
			string(model.PaymentAttemptsDBFieldName.ProviderReference),
			string(model.PaymentAttemptsDBFieldName.ProviderTransactionId),
			string(model.PaymentAttemptsDBFieldName.ProviderOrderId),
			string(model.PaymentAttemptsDBFieldName.ProviderPaymentId),
			string(model.PaymentAttemptsDBFieldName.FailureCode),
			string(model.PaymentAttemptsDBFieldName.FailureMessage),
			string(model.PaymentAttemptsDBFieldName.ExpiresAt),
			string(model.PaymentAttemptsDBFieldName.AuthorizedAt),
			string(model.PaymentAttemptsDBFieldName.CapturedAt),
			string(model.PaymentAttemptsDBFieldName.PaidAt),
			string(model.PaymentAttemptsDBFieldName.FailedAt),
			string(model.PaymentAttemptsDBFieldName.CanceledAt),
			string(model.PaymentAttemptsDBFieldName.StatusSyncRequiredAt),
			string(model.PaymentAttemptsDBFieldName.LastStatusSyncAt),
			string(model.PaymentAttemptsDBFieldName.RawRequest),
			string(model.PaymentAttemptsDBFieldName.RawResponse),
			string(model.PaymentAttemptsDBFieldName.Metadata),
			string(model.PaymentAttemptsDBFieldName.MetaCreatedAt),
			string(model.PaymentAttemptsDBFieldName.MetaCreatedBy),
			string(model.PaymentAttemptsDBFieldName.MetaUpdatedAt),
			string(model.PaymentAttemptsDBFieldName.MetaUpdatedBy),
			string(model.PaymentAttemptsDBFieldName.MetaDeletedAt),
			string(model.PaymentAttemptsDBFieldName.MetaDeletedBy),
		)
	}
	paymentAttemptsSelectableResponse := PaymentAttemptsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentAttemptsDBFieldName.Id):
			key := string(PaymentAttemptsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.Id, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.PaymentIntentId):
			key := string(PaymentAttemptsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.PaymentIntentId, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.AttemptNo):
			key := string(PaymentAttemptsDTOFieldName.AttemptNo)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.AttemptNo, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderAccountId):
			key := string(PaymentAttemptsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderAccountId.UUID, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.RouteDecisionId):
			key := string(PaymentAttemptsDTOFieldName.RouteDecisionId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.RouteDecisionId.UUID, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderCode):
			key := string(PaymentAttemptsDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderCode.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MethodCode):
			key := string(PaymentAttemptsDTOFieldName.MethodCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MethodCode, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ChannelCode):
			key := string(PaymentAttemptsDTOFieldName.ChannelCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ChannelCode.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.Amount):
			key := string(PaymentAttemptsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.Amount, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.Currency):
			key := string(PaymentAttemptsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.Currency, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.Status):
			key := string(PaymentAttemptsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, model.PaymentAttemptStatus(paymentAttempts.Status), explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderReference):
			key := string(PaymentAttemptsDTOFieldName.ProviderReference)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderReference.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderTransactionId):
			key := string(PaymentAttemptsDTOFieldName.ProviderTransactionId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderTransactionId.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderOrderId):
			key := string(PaymentAttemptsDTOFieldName.ProviderOrderId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderOrderId.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ProviderPaymentId):
			key := string(PaymentAttemptsDTOFieldName.ProviderPaymentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ProviderPaymentId.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.FailureCode):
			key := string(PaymentAttemptsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.FailureCode.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.FailureMessage):
			key := string(PaymentAttemptsDTOFieldName.FailureMessage)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.FailureMessage.String, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.ExpiresAt):
			key := string(PaymentAttemptsDTOFieldName.ExpiresAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.ExpiresAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.AuthorizedAt):
			key := string(PaymentAttemptsDTOFieldName.AuthorizedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.AuthorizedAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.CapturedAt):
			key := string(PaymentAttemptsDTOFieldName.CapturedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.CapturedAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.PaidAt):
			key := string(PaymentAttemptsDTOFieldName.PaidAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.PaidAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.FailedAt):
			key := string(PaymentAttemptsDTOFieldName.FailedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.FailedAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.CanceledAt):
			key := string(PaymentAttemptsDTOFieldName.CanceledAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.CanceledAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.StatusSyncRequiredAt):
			key := string(PaymentAttemptsDTOFieldName.StatusSyncRequiredAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.StatusSyncRequiredAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.LastStatusSyncAt):
			key := string(PaymentAttemptsDTOFieldName.LastStatusSyncAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.LastStatusSyncAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.RawRequest):
			key := string(PaymentAttemptsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.RawRequest, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.RawResponse):
			key := string(PaymentAttemptsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.RawResponse, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.Metadata):
			key := string(PaymentAttemptsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.Metadata, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaCreatedAt):
			key := string(PaymentAttemptsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaCreatedAt, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaCreatedBy):
			key := string(PaymentAttemptsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaCreatedBy, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaUpdatedAt):
			key := string(PaymentAttemptsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaUpdatedBy):
			key := string(PaymentAttemptsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaDeletedAt):
			key := string(PaymentAttemptsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentAttemptsDBFieldName.MetaDeletedBy):
			key := string(PaymentAttemptsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAttemptsSelectableValue(paymentAttemptsSelectableResponse, key, paymentAttempts.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentAttemptsSelectableResponse
}

func NewPaymentAttemptsListResponseFromFilterResult(result []model.PaymentAttemptsFilterResult, filter model.Filter) PaymentAttemptsSelectableListResponse {
	dtoPaymentAttemptsListResponse := PaymentAttemptsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentAttemptsResponse := NewPaymentAttemptsSelectableResponse(row.PaymentAttempts, filter)
		dtoPaymentAttemptsListResponse = append(dtoPaymentAttemptsListResponse, &dtoPaymentAttemptsResponse)
	}
	return dtoPaymentAttemptsListResponse
}

type PaymentAttemptsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     PaymentAttemptsSelectableListResponse `json:"data"`
}

func reversePaymentAttemptsFilterResults(result []model.PaymentAttemptsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentAttemptsFilterResponse(result []model.PaymentAttemptsFilterResult, filter model.Filter) (resp PaymentAttemptsFilterResponse) {
	resp.Metadata.Strategy = filter.Pagination.Strategy
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	dataResult := result

	if filter.Pagination.IsCursorMode() && filter.Pagination.PageSize > 0 {
		if len(dataResult) > filter.Pagination.PageSize {
			resp.Metadata.HasMore = true
			if filter.Pagination.Direction == model.CursorDirectionPrev {
				resp.Metadata.HasPrev = true
			} else {
				resp.Metadata.HasNext = true
			}
			dataResult = dataResult[:filter.Pagination.PageSize]
		}
		if filter.Pagination.Direction == model.CursorDirectionPrev {
			reversePaymentAttemptsFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].Id
			resp.Metadata.PrevCursor = dataResult[0].Id
		}
		if resp.Metadata.Page <= 0 {
			resp.Metadata.Page = 1
		}
	} else {
		if len(dataResult) > 0 {
			resp.Metadata.TotalData = dataResult[0].FilterCount
			resp.Metadata.TotalPage = int(math.Ceil(float64(resp.Metadata.TotalData) / float64(filter.Pagination.PageSize)))
			resp.Metadata.HasPrev = filter.Pagination.Page > 1
			resp.Metadata.HasNext = filter.Pagination.Page < resp.Metadata.TotalPage
			resp.Metadata.HasMore = resp.Metadata.HasNext
		}
	}

	resp.Data = NewPaymentAttemptsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentAttemptsCreateRequest struct {
	PaymentIntentId       uuid.UUID                  `json:"paymentIntentId"`
	AttemptNo             int                        `json:"attemptNo"`
	ProviderAccountId     uuid.UUID                  `json:"providerAccountId"`
	RouteDecisionId       uuid.UUID                  `json:"routeDecisionId"`
	ProviderCode          string                     `json:"providerCode"`
	MethodCode            string                     `json:"methodCode"`
	ChannelCode           string                     `json:"channelCode"`
	Amount                decimal.Decimal            `json:"amount"`
	Currency              string                     `json:"currency"`
	Status                model.PaymentAttemptStatus `json:"status" example:"created" enums:"created,pending,authorized,captured,paid,failed,canceled"`
	ProviderReference     string                     `json:"providerReference"`
	ProviderTransactionId string                     `json:"providerTransactionId"`
	ProviderOrderId       string                     `json:"providerOrderId"`
	ProviderPaymentId     string                     `json:"providerPaymentId"`
	FailureCode           string                     `json:"failureCode"`
	FailureMessage        string                     `json:"failureMessage"`
	ExpiresAt             time.Time                  `json:"expiresAt"`
	AuthorizedAt          time.Time                  `json:"authorizedAt"`
	CapturedAt            time.Time                  `json:"capturedAt"`
	PaidAt                time.Time                  `json:"paidAt"`
	FailedAt              time.Time                  `json:"failedAt"`
	CanceledAt            time.Time                  `json:"canceledAt"`
	StatusSyncRequiredAt  time.Time                  `json:"statusSyncRequiredAt"`
	LastStatusSyncAt      time.Time                  `json:"lastStatusSyncAt"`
	RawRequest            json.RawMessage            `json:"rawRequest"`
	RawResponse           json.RawMessage            `json:"rawResponse"`
	Metadata              json.RawMessage            `json:"metadata"`
}

func (d *PaymentAttemptsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentAttemptsCreateRequest) ToModel() model.PaymentAttempts {
	id, _ := uuid.NewV4()
	return model.PaymentAttempts{
		Id:                    id,
		PaymentIntentId:       d.PaymentIntentId,
		AttemptNo:             d.AttemptNo,
		ProviderAccountId:     nuuid.From(d.ProviderAccountId),
		RouteDecisionId:       nuuid.From(d.RouteDecisionId),
		ProviderCode:          null.StringFrom(d.ProviderCode),
		MethodCode:            d.MethodCode,
		ChannelCode:           null.StringFrom(d.ChannelCode),
		Amount:                d.Amount,
		Currency:              d.Currency,
		Status:                d.Status,
		ProviderReference:     null.StringFrom(d.ProviderReference),
		ProviderTransactionId: null.StringFrom(d.ProviderTransactionId),
		ProviderOrderId:       null.StringFrom(d.ProviderOrderId),
		ProviderPaymentId:     null.StringFrom(d.ProviderPaymentId),
		FailureCode:           null.StringFrom(d.FailureCode),
		FailureMessage:        null.StringFrom(d.FailureMessage),
		ExpiresAt:             null.TimeFrom(d.ExpiresAt),
		AuthorizedAt:          null.TimeFrom(d.AuthorizedAt),
		CapturedAt:            null.TimeFrom(d.CapturedAt),
		PaidAt:                null.TimeFrom(d.PaidAt),
		FailedAt:              null.TimeFrom(d.FailedAt),
		CanceledAt:            null.TimeFrom(d.CanceledAt),
		StatusSyncRequiredAt:  null.TimeFrom(d.StatusSyncRequiredAt),
		LastStatusSyncAt:      null.TimeFrom(d.LastStatusSyncAt),
		RawRequest:            d.RawRequest,
		RawResponse:           d.RawResponse,
		Metadata:              d.Metadata,
	}
}

type PaymentAttemptsListCreateRequest []*PaymentAttemptsCreateRequest

func (d PaymentAttemptsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAttempts := range d {
		err = validator.Struct(paymentAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentAttemptsListCreateRequest) ToModelList() []model.PaymentAttempts {
	out := make([]model.PaymentAttempts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentAttemptsUpdateRequest struct {
	PaymentIntentId       uuid.UUID                  `json:"paymentIntentId"`
	AttemptNo             int                        `json:"attemptNo"`
	ProviderAccountId     uuid.UUID                  `json:"providerAccountId"`
	RouteDecisionId       uuid.UUID                  `json:"routeDecisionId"`
	ProviderCode          string                     `json:"providerCode"`
	MethodCode            string                     `json:"methodCode"`
	ChannelCode           string                     `json:"channelCode"`
	Amount                decimal.Decimal            `json:"amount"`
	Currency              string                     `json:"currency"`
	Status                model.PaymentAttemptStatus `json:"status" example:"created" enums:"created,pending,authorized,captured,paid,failed,canceled"`
	ProviderReference     string                     `json:"providerReference"`
	ProviderTransactionId string                     `json:"providerTransactionId"`
	ProviderOrderId       string                     `json:"providerOrderId"`
	ProviderPaymentId     string                     `json:"providerPaymentId"`
	FailureCode           string                     `json:"failureCode"`
	FailureMessage        string                     `json:"failureMessage"`
	ExpiresAt             time.Time                  `json:"expiresAt"`
	AuthorizedAt          time.Time                  `json:"authorizedAt"`
	CapturedAt            time.Time                  `json:"capturedAt"`
	PaidAt                time.Time                  `json:"paidAt"`
	FailedAt              time.Time                  `json:"failedAt"`
	CanceledAt            time.Time                  `json:"canceledAt"`
	StatusSyncRequiredAt  time.Time                  `json:"statusSyncRequiredAt"`
	LastStatusSyncAt      time.Time                  `json:"lastStatusSyncAt"`
	RawRequest            json.RawMessage            `json:"rawRequest"`
	RawResponse           json.RawMessage            `json:"rawResponse"`
	Metadata              json.RawMessage            `json:"metadata"`
}

func (d *PaymentAttemptsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentAttemptsUpdateRequest) ToModel() model.PaymentAttempts {
	return model.PaymentAttempts{
		PaymentIntentId:       d.PaymentIntentId,
		AttemptNo:             d.AttemptNo,
		ProviderAccountId:     nuuid.From(d.ProviderAccountId),
		RouteDecisionId:       nuuid.From(d.RouteDecisionId),
		ProviderCode:          null.StringFrom(d.ProviderCode),
		MethodCode:            d.MethodCode,
		ChannelCode:           null.StringFrom(d.ChannelCode),
		Amount:                d.Amount,
		Currency:              d.Currency,
		Status:                d.Status,
		ProviderReference:     null.StringFrom(d.ProviderReference),
		ProviderTransactionId: null.StringFrom(d.ProviderTransactionId),
		ProviderOrderId:       null.StringFrom(d.ProviderOrderId),
		ProviderPaymentId:     null.StringFrom(d.ProviderPaymentId),
		FailureCode:           null.StringFrom(d.FailureCode),
		FailureMessage:        null.StringFrom(d.FailureMessage),
		ExpiresAt:             null.TimeFrom(d.ExpiresAt),
		AuthorizedAt:          null.TimeFrom(d.AuthorizedAt),
		CapturedAt:            null.TimeFrom(d.CapturedAt),
		PaidAt:                null.TimeFrom(d.PaidAt),
		FailedAt:              null.TimeFrom(d.FailedAt),
		CanceledAt:            null.TimeFrom(d.CanceledAt),
		StatusSyncRequiredAt:  null.TimeFrom(d.StatusSyncRequiredAt),
		LastStatusSyncAt:      null.TimeFrom(d.LastStatusSyncAt),
		RawRequest:            d.RawRequest,
		RawResponse:           d.RawResponse,
		Metadata:              d.Metadata,
	}
}

type PaymentAttemptsBulkUpdateRequest struct {
	Id                    uuid.UUID                  `json:"id"`
	PaymentIntentId       uuid.UUID                  `json:"paymentIntentId"`
	AttemptNo             int                        `json:"attemptNo"`
	ProviderAccountId     uuid.UUID                  `json:"providerAccountId"`
	RouteDecisionId       uuid.UUID                  `json:"routeDecisionId"`
	ProviderCode          string                     `json:"providerCode"`
	MethodCode            string                     `json:"methodCode"`
	ChannelCode           string                     `json:"channelCode"`
	Amount                decimal.Decimal            `json:"amount"`
	Currency              string                     `json:"currency"`
	Status                model.PaymentAttemptStatus `json:"status" example:"created" enums:"created,pending,authorized,captured,paid,failed,canceled"`
	ProviderReference     string                     `json:"providerReference"`
	ProviderTransactionId string                     `json:"providerTransactionId"`
	ProviderOrderId       string                     `json:"providerOrderId"`
	ProviderPaymentId     string                     `json:"providerPaymentId"`
	FailureCode           string                     `json:"failureCode"`
	FailureMessage        string                     `json:"failureMessage"`
	ExpiresAt             time.Time                  `json:"expiresAt"`
	AuthorizedAt          time.Time                  `json:"authorizedAt"`
	CapturedAt            time.Time                  `json:"capturedAt"`
	PaidAt                time.Time                  `json:"paidAt"`
	FailedAt              time.Time                  `json:"failedAt"`
	CanceledAt            time.Time                  `json:"canceledAt"`
	StatusSyncRequiredAt  time.Time                  `json:"statusSyncRequiredAt"`
	LastStatusSyncAt      time.Time                  `json:"lastStatusSyncAt"`
	RawRequest            json.RawMessage            `json:"rawRequest"`
	RawResponse           json.RawMessage            `json:"rawResponse"`
	Metadata              json.RawMessage            `json:"metadata"`
}

func (d PaymentAttemptsBulkUpdateRequest) PrimaryID() PaymentAttemptsPrimaryID {
	return PaymentAttemptsPrimaryID{
		Id: d.Id,
	}
}

type PaymentAttemptsListBulkUpdateRequest []*PaymentAttemptsBulkUpdateRequest

func (d PaymentAttemptsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAttempts := range d {
		err = validator.Struct(paymentAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentAttemptsBulkUpdateRequest) ToModel() model.PaymentAttempts {
	return model.PaymentAttempts{
		Id:                    d.Id,
		PaymentIntentId:       d.PaymentIntentId,
		AttemptNo:             d.AttemptNo,
		ProviderAccountId:     nuuid.From(d.ProviderAccountId),
		RouteDecisionId:       nuuid.From(d.RouteDecisionId),
		ProviderCode:          null.StringFrom(d.ProviderCode),
		MethodCode:            d.MethodCode,
		ChannelCode:           null.StringFrom(d.ChannelCode),
		Amount:                d.Amount,
		Currency:              d.Currency,
		Status:                d.Status,
		ProviderReference:     null.StringFrom(d.ProviderReference),
		ProviderTransactionId: null.StringFrom(d.ProviderTransactionId),
		ProviderOrderId:       null.StringFrom(d.ProviderOrderId),
		ProviderPaymentId:     null.StringFrom(d.ProviderPaymentId),
		FailureCode:           null.StringFrom(d.FailureCode),
		FailureMessage:        null.StringFrom(d.FailureMessage),
		ExpiresAt:             null.TimeFrom(d.ExpiresAt),
		AuthorizedAt:          null.TimeFrom(d.AuthorizedAt),
		CapturedAt:            null.TimeFrom(d.CapturedAt),
		PaidAt:                null.TimeFrom(d.PaidAt),
		FailedAt:              null.TimeFrom(d.FailedAt),
		CanceledAt:            null.TimeFrom(d.CanceledAt),
		StatusSyncRequiredAt:  null.TimeFrom(d.StatusSyncRequiredAt),
		LastStatusSyncAt:      null.TimeFrom(d.LastStatusSyncAt),
		RawRequest:            d.RawRequest,
		RawResponse:           d.RawResponse,
		Metadata:              d.Metadata,
	}
}

type PaymentAttemptsResponse struct {
	Id                    uuid.UUID                  `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId       uuid.UUID                  `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AttemptNo             int                        `json:"attemptNo" validate:"required" example:"1"`
	ProviderAccountId     uuid.UUID                  `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RouteDecisionId       uuid.UUID                  `json:"routeDecisionId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode          string                     `json:"providerCode"`
	MethodCode            string                     `json:"methodCode" validate:"required"`
	ChannelCode           string                     `json:"channelCode"`
	Amount                decimal.Decimal            `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency              string                     `json:"currency"`
	Status                model.PaymentAttemptStatus `json:"status" validate:"oneof=created pending authorized captured paid failed canceled" enums:"created,pending,authorized,captured,paid,failed,canceled"`
	ProviderReference     string                     `json:"providerReference"`
	ProviderTransactionId string                     `json:"providerTransactionId"`
	ProviderOrderId       string                     `json:"providerOrderId"`
	ProviderPaymentId     string                     `json:"providerPaymentId"`
	FailureCode           string                     `json:"failureCode"`
	FailureMessage        string                     `json:"failureMessage"`
	ExpiresAt             time.Time                  `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	AuthorizedAt          time.Time                  `json:"authorizedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CapturedAt            time.Time                  `json:"capturedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PaidAt                time.Time                  `json:"paidAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailedAt              time.Time                  `json:"failedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CanceledAt            time.Time                  `json:"canceledAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	StatusSyncRequiredAt  time.Time                  `json:"statusSyncRequiredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LastStatusSyncAt      time.Time                  `json:"lastStatusSyncAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RawRequest            json.RawMessage            `json:"rawRequest" swaggertype:"object"`
	RawResponse           json.RawMessage            `json:"rawResponse" swaggertype:"object"`
	Metadata              json.RawMessage            `json:"metadata" swaggertype:"object"`
}

func NewPaymentAttemptsResponse(paymentAttempts model.PaymentAttempts) PaymentAttemptsResponse {
	return PaymentAttemptsResponse{
		Id:                    paymentAttempts.Id,
		PaymentIntentId:       paymentAttempts.PaymentIntentId,
		AttemptNo:             paymentAttempts.AttemptNo,
		ProviderAccountId:     paymentAttempts.ProviderAccountId.UUID,
		RouteDecisionId:       paymentAttempts.RouteDecisionId.UUID,
		ProviderCode:          paymentAttempts.ProviderCode.String,
		MethodCode:            paymentAttempts.MethodCode,
		ChannelCode:           paymentAttempts.ChannelCode.String,
		Amount:                paymentAttempts.Amount,
		Currency:              paymentAttempts.Currency,
		Status:                model.PaymentAttemptStatus(paymentAttempts.Status),
		ProviderReference:     paymentAttempts.ProviderReference.String,
		ProviderTransactionId: paymentAttempts.ProviderTransactionId.String,
		ProviderOrderId:       paymentAttempts.ProviderOrderId.String,
		ProviderPaymentId:     paymentAttempts.ProviderPaymentId.String,
		FailureCode:           paymentAttempts.FailureCode.String,
		FailureMessage:        paymentAttempts.FailureMessage.String,
		ExpiresAt:             paymentAttempts.ExpiresAt.Time,
		AuthorizedAt:          paymentAttempts.AuthorizedAt.Time,
		CapturedAt:            paymentAttempts.CapturedAt.Time,
		PaidAt:                paymentAttempts.PaidAt.Time,
		FailedAt:              paymentAttempts.FailedAt.Time,
		CanceledAt:            paymentAttempts.CanceledAt.Time,
		StatusSyncRequiredAt:  paymentAttempts.StatusSyncRequiredAt.Time,
		LastStatusSyncAt:      paymentAttempts.LastStatusSyncAt.Time,
		RawRequest:            paymentAttempts.RawRequest,
		RawResponse:           paymentAttempts.RawResponse,
		Metadata:              paymentAttempts.Metadata,
	}
}

type PaymentAttemptsListResponse []*PaymentAttemptsResponse

func NewPaymentAttemptsListResponse(paymentAttemptsList model.PaymentAttemptsList) PaymentAttemptsListResponse {
	dtoPaymentAttemptsListResponse := PaymentAttemptsListResponse{}
	for _, paymentAttempts := range paymentAttemptsList {
		dtoPaymentAttemptsResponse := NewPaymentAttemptsResponse(*paymentAttempts)
		dtoPaymentAttemptsListResponse = append(dtoPaymentAttemptsListResponse, &dtoPaymentAttemptsResponse)
	}
	return dtoPaymentAttemptsListResponse
}

type PaymentAttemptsPrimaryIDList []PaymentAttemptsPrimaryID

func (d PaymentAttemptsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAttempts := range d {
		err = validator.Struct(paymentAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentAttemptsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentAttemptsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentAttemptsPrimaryID) ToModel() model.PaymentAttemptsPrimaryID {
	return model.PaymentAttemptsPrimaryID{
		Id: d.Id,
	}
}

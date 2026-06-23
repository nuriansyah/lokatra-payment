package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentRouteCandidatesRuntimeDTOFieldNameType string

type paymentRouteCandidatesRuntimeDTOFieldName struct {
	Id                  PaymentRouteCandidatesRuntimeDTOFieldNameType
	ScopeType           PaymentRouteCandidatesRuntimeDTOFieldNameType
	ScopeId             PaymentRouteCandidatesRuntimeDTOFieldNameType
	MerchantId          PaymentRouteCandidatesRuntimeDTOFieldNameType
	MethodCode          PaymentRouteCandidatesRuntimeDTOFieldNameType
	ChannelCode         PaymentRouteCandidatesRuntimeDTOFieldNameType
	Currency            PaymentRouteCandidatesRuntimeDTOFieldNameType
	MinAmount           PaymentRouteCandidatesRuntimeDTOFieldNameType
	MaxAmount           PaymentRouteCandidatesRuntimeDTOFieldNameType
	ProviderAccountId   PaymentRouteCandidatesRuntimeDTOFieldNameType
	ProviderMethodCode  PaymentRouteCandidatesRuntimeDTOFieldNameType
	ProviderChannelCode PaymentRouteCandidatesRuntimeDTOFieldNameType
	Priority            PaymentRouteCandidatesRuntimeDTOFieldNameType
	IsFallback          PaymentRouteCandidatesRuntimeDTOFieldNameType
	TrafficWeight       PaymentRouteCandidatesRuntimeDTOFieldNameType
	TimeoutMs           PaymentRouteCandidatesRuntimeDTOFieldNameType
	MaxAttempts         PaymentRouteCandidatesRuntimeDTOFieldNameType
	IsEnabled           PaymentRouteCandidatesRuntimeDTOFieldNameType
	ConditionExpr       PaymentRouteCandidatesRuntimeDTOFieldNameType
	Metadata            PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaCreatedAt       PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaCreatedBy       PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaUpdatedAt       PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaUpdatedBy       PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaDeletedAt       PaymentRouteCandidatesRuntimeDTOFieldNameType
	MetaDeletedBy       PaymentRouteCandidatesRuntimeDTOFieldNameType
}

var PaymentRouteCandidatesRuntimeDTOFieldName = paymentRouteCandidatesRuntimeDTOFieldName{
	Id:                  "id",
	ScopeType:           "scopeType",
	ScopeId:             "scopeId",
	MerchantId:          "merchantId",
	MethodCode:          "methodCode",
	ChannelCode:         "channelCode",
	Currency:            "currency",
	MinAmount:           "minAmount",
	MaxAmount:           "maxAmount",
	ProviderAccountId:   "providerAccountId",
	ProviderMethodCode:  "providerMethodCode",
	ProviderChannelCode: "providerChannelCode",
	Priority:            "priority",
	IsFallback:          "isFallback",
	TrafficWeight:       "trafficWeight",
	TimeoutMs:           "timeoutMs",
	MaxAttempts:         "maxAttempts",
	IsEnabled:           "isEnabled",
	ConditionExpr:       "conditionExpr",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformPaymentRouteCandidatesRuntimeDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.Id):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.Id), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ScopeType):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeType), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ScopeId):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeId), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MerchantId):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MerchantId), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MethodCode):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MethodCode), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ChannelCode):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ChannelCode), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.Currency):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.Currency), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MinAmount):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MinAmount), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MaxAmount):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAmount), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderAccountId):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderAccountId), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderMethodCode):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderMethodCode), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderChannelCode):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderChannelCode), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.Priority):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.Priority), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.IsFallback):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsFallback), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.TrafficWeight):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.TrafficWeight), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.TimeoutMs):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.TimeoutMs), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MaxAttempts):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAttempts), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.IsEnabled):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsEnabled), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.ConditionExpr):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.ConditionExpr), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.Metadata):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.Metadata), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaCreatedAt):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedAt), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaCreatedBy):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedBy), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedAt), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedBy), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaDeletedAt):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedAt), true

	case string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaDeletedBy):
		return string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentRouteCandidatesRuntimeBaseFilterField(field string) bool {
	spec, found := model.NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentRouteCandidatesRuntimeProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentRouteCandidatesRuntimeProjectionOutputPath(path string) error {
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

func transformPaymentRouteCandidatesRuntimeFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentRouteCandidatesRuntimeDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentRouteCandidatesRuntimeFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentRouteCandidatesRuntimeFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentRouteCandidatesRuntimeDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentRouteCandidatesRuntimeBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentRouteCandidatesRuntimeProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentRouteCandidatesRuntimeProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentRouteCandidatesRuntimeDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentRouteCandidatesRuntimeDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentRouteCandidatesRuntimeFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentRouteCandidatesRuntimeFilter(filter *model.Filter) {
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
			Field: string(PaymentRouteCandidatesRuntimeDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentRouteCandidatesRuntimeSelectableResponse map[string]interface{}
type PaymentRouteCandidatesRuntimeSelectableListResponse []*PaymentRouteCandidatesRuntimeSelectableResponse

func assignPaymentRouteCandidatesRuntimeNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentRouteCandidatesRuntimeSelectableValue(out PaymentRouteCandidatesRuntimeSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentRouteCandidatesRuntimeNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentRouteCandidatesRuntimeSelectableResponse(paymentRouteCandidatesRuntime model.PaymentRouteCandidatesRuntime, filter model.Filter) PaymentRouteCandidatesRuntimeSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.Id),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeType),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeId),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MerchantId),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MethodCode),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ChannelCode),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.Currency),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MinAmount),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAmount),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderAccountId),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderMethodCode),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderChannelCode),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.Priority),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsFallback),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.TrafficWeight),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.TimeoutMs),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAttempts),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsEnabled),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.ConditionExpr),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.Metadata),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedAt),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedBy),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedAt),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedBy),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedAt),
			string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedBy),
		)
	}
	paymentRouteCandidatesRuntimeSelectableResponse := PaymentRouteCandidatesRuntimeSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.Id):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.Id, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeType):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ScopeType)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, model.ScopeType(paymentRouteCandidatesRuntime.ScopeType), explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ScopeId):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ScopeId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ScopeId.UUID, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MerchantId):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MerchantId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MerchantId.UUID, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MethodCode):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MethodCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MethodCode, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ChannelCode):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ChannelCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ChannelCode.String, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.Currency):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.Currency, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MinAmount):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MinAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MinAmount.Decimal, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAmount):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MaxAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MaxAmount.Decimal, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderAccountId):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ProviderAccountId, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderMethodCode):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderMethodCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ProviderMethodCode.String, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ProviderChannelCode):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ProviderChannelCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ProviderChannelCode.String, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.Priority):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.Priority)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.Priority, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsFallback):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.IsFallback)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.IsFallback, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.TrafficWeight):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.TrafficWeight)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.TrafficWeight, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.TimeoutMs):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.TimeoutMs)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.TimeoutMs, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MaxAttempts):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MaxAttempts)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MaxAttempts, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.IsEnabled):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.IsEnabled)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.IsEnabled, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.ConditionExpr):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.ConditionExpr)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.ConditionExpr, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.Metadata):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.Metadata, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedAt):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaCreatedAt, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaCreatedBy):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaCreatedBy, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedAt):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaUpdatedBy):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedAt):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentRouteCandidatesRuntimeDBFieldName.MetaDeletedBy):
			key := string(PaymentRouteCandidatesRuntimeDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteCandidatesRuntimeSelectableValue(paymentRouteCandidatesRuntimeSelectableResponse, key, paymentRouteCandidatesRuntime.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentRouteCandidatesRuntimeSelectableResponse
}

func NewPaymentRouteCandidatesRuntimeListResponseFromFilterResult(result []model.PaymentRouteCandidatesRuntimeFilterResult, filter model.Filter) PaymentRouteCandidatesRuntimeSelectableListResponse {
	dtoPaymentRouteCandidatesRuntimeListResponse := PaymentRouteCandidatesRuntimeSelectableListResponse{}
	for _, row := range result {
		dtoPaymentRouteCandidatesRuntimeResponse := NewPaymentRouteCandidatesRuntimeSelectableResponse(row.PaymentRouteCandidatesRuntime, filter)
		dtoPaymentRouteCandidatesRuntimeListResponse = append(dtoPaymentRouteCandidatesRuntimeListResponse, &dtoPaymentRouteCandidatesRuntimeResponse)
	}
	return dtoPaymentRouteCandidatesRuntimeListResponse
}

type PaymentRouteCandidatesRuntimeFilterResponse struct {
	Metadata Metadata                                            `json:"metadata"`
	Data     PaymentRouteCandidatesRuntimeSelectableListResponse `json:"data"`
}

func reversePaymentRouteCandidatesRuntimeFilterResults(result []model.PaymentRouteCandidatesRuntimeFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentRouteCandidatesRuntimeFilterResponse(result []model.PaymentRouteCandidatesRuntimeFilterResult, filter model.Filter) (resp PaymentRouteCandidatesRuntimeFilterResponse) {
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
			reversePaymentRouteCandidatesRuntimeFilterResults(dataResult)
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

	resp.Data = NewPaymentRouteCandidatesRuntimeListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentRouteCandidatesRuntimeCreateRequest struct {
	ScopeType           model.ScopeType `json:"scopeType" example:"platform" enums:"platform,merchant"`
	ScopeId             uuid.UUID       `json:"scopeId"`
	MerchantId          uuid.UUID       `json:"merchantId"`
	MethodCode          string          `json:"methodCode"`
	ChannelCode         string          `json:"channelCode"`
	Currency            string          `json:"currency"`
	MinAmount           decimal.Decimal `json:"minAmount"`
	MaxAmount           decimal.Decimal `json:"maxAmount"`
	ProviderAccountId   uuid.UUID       `json:"providerAccountId"`
	ProviderMethodCode  string          `json:"providerMethodCode"`
	ProviderChannelCode string          `json:"providerChannelCode"`
	Priority            int             `json:"priority"`
	IsFallback          bool            `json:"isFallback"`
	TrafficWeight       int             `json:"trafficWeight"`
	TimeoutMs           int             `json:"timeoutMs"`
	MaxAttempts         int             `json:"maxAttempts"`
	IsEnabled           bool            `json:"isEnabled"`
	ConditionExpr       json.RawMessage `json:"conditionExpr"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d *PaymentRouteCandidatesRuntimeCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentRouteCandidatesRuntimeCreateRequest) ToModel() model.PaymentRouteCandidatesRuntime {
	id, _ := uuid.NewV4()
	return model.PaymentRouteCandidatesRuntime{
		Id:                  id,
		ScopeType:           d.ScopeType,
		ScopeId:             nuuid.From(d.ScopeId),
		MerchantId:          nuuid.From(d.MerchantId),
		MethodCode:          d.MethodCode,
		ChannelCode:         null.StringFrom(d.ChannelCode),
		Currency:            d.Currency,
		MinAmount:           decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:           decimal.NewNullDecimal(d.MaxAmount),
		ProviderAccountId:   d.ProviderAccountId,
		ProviderMethodCode:  null.StringFrom(d.ProviderMethodCode),
		ProviderChannelCode: null.StringFrom(d.ProviderChannelCode),
		Priority:            d.Priority,
		IsFallback:          d.IsFallback,
		TrafficWeight:       d.TrafficWeight,
		TimeoutMs:           d.TimeoutMs,
		MaxAttempts:         d.MaxAttempts,
		IsEnabled:           d.IsEnabled,
		ConditionExpr:       d.ConditionExpr,
		Metadata:            d.Metadata,
	}
}

type PaymentRouteCandidatesRuntimeListCreateRequest []*PaymentRouteCandidatesRuntimeCreateRequest

func (d PaymentRouteCandidatesRuntimeListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteCandidatesRuntime := range d {
		err = validator.Struct(paymentRouteCandidatesRuntime)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRouteCandidatesRuntimeListCreateRequest) ToModelList() []model.PaymentRouteCandidatesRuntime {
	out := make([]model.PaymentRouteCandidatesRuntime, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentRouteCandidatesRuntimeUpdateRequest struct {
	ScopeType           model.ScopeType `json:"scopeType" example:"platform" enums:"platform,merchant"`
	ScopeId             uuid.UUID       `json:"scopeId"`
	MerchantId          uuid.UUID       `json:"merchantId"`
	MethodCode          string          `json:"methodCode"`
	ChannelCode         string          `json:"channelCode"`
	Currency            string          `json:"currency"`
	MinAmount           decimal.Decimal `json:"minAmount"`
	MaxAmount           decimal.Decimal `json:"maxAmount"`
	ProviderAccountId   uuid.UUID       `json:"providerAccountId"`
	ProviderMethodCode  string          `json:"providerMethodCode"`
	ProviderChannelCode string          `json:"providerChannelCode"`
	Priority            int             `json:"priority"`
	IsFallback          bool            `json:"isFallback"`
	TrafficWeight       int             `json:"trafficWeight"`
	TimeoutMs           int             `json:"timeoutMs"`
	MaxAttempts         int             `json:"maxAttempts"`
	IsEnabled           bool            `json:"isEnabled"`
	ConditionExpr       json.RawMessage `json:"conditionExpr"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d *PaymentRouteCandidatesRuntimeUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentRouteCandidatesRuntimeUpdateRequest) ToModel() model.PaymentRouteCandidatesRuntime {
	return model.PaymentRouteCandidatesRuntime{
		ScopeType:           d.ScopeType,
		ScopeId:             nuuid.From(d.ScopeId),
		MerchantId:          nuuid.From(d.MerchantId),
		MethodCode:          d.MethodCode,
		ChannelCode:         null.StringFrom(d.ChannelCode),
		Currency:            d.Currency,
		MinAmount:           decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:           decimal.NewNullDecimal(d.MaxAmount),
		ProviderAccountId:   d.ProviderAccountId,
		ProviderMethodCode:  null.StringFrom(d.ProviderMethodCode),
		ProviderChannelCode: null.StringFrom(d.ProviderChannelCode),
		Priority:            d.Priority,
		IsFallback:          d.IsFallback,
		TrafficWeight:       d.TrafficWeight,
		TimeoutMs:           d.TimeoutMs,
		MaxAttempts:         d.MaxAttempts,
		IsEnabled:           d.IsEnabled,
		ConditionExpr:       d.ConditionExpr,
		Metadata:            d.Metadata,
	}
}

type PaymentRouteCandidatesRuntimeBulkUpdateRequest struct {
	Id                  uuid.UUID       `json:"id"`
	ScopeType           model.ScopeType `json:"scopeType" example:"platform" enums:"platform,merchant"`
	ScopeId             uuid.UUID       `json:"scopeId"`
	MerchantId          uuid.UUID       `json:"merchantId"`
	MethodCode          string          `json:"methodCode"`
	ChannelCode         string          `json:"channelCode"`
	Currency            string          `json:"currency"`
	MinAmount           decimal.Decimal `json:"minAmount"`
	MaxAmount           decimal.Decimal `json:"maxAmount"`
	ProviderAccountId   uuid.UUID       `json:"providerAccountId"`
	ProviderMethodCode  string          `json:"providerMethodCode"`
	ProviderChannelCode string          `json:"providerChannelCode"`
	Priority            int             `json:"priority"`
	IsFallback          bool            `json:"isFallback"`
	TrafficWeight       int             `json:"trafficWeight"`
	TimeoutMs           int             `json:"timeoutMs"`
	MaxAttempts         int             `json:"maxAttempts"`
	IsEnabled           bool            `json:"isEnabled"`
	ConditionExpr       json.RawMessage `json:"conditionExpr"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d PaymentRouteCandidatesRuntimeBulkUpdateRequest) PrimaryID() PaymentRouteCandidatesRuntimePrimaryID {
	return PaymentRouteCandidatesRuntimePrimaryID{
		Id: d.Id,
	}
}

type PaymentRouteCandidatesRuntimeListBulkUpdateRequest []*PaymentRouteCandidatesRuntimeBulkUpdateRequest

func (d PaymentRouteCandidatesRuntimeListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteCandidatesRuntime := range d {
		err = validator.Struct(paymentRouteCandidatesRuntime)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRouteCandidatesRuntimeBulkUpdateRequest) ToModel() model.PaymentRouteCandidatesRuntime {
	return model.PaymentRouteCandidatesRuntime{
		Id:                  d.Id,
		ScopeType:           d.ScopeType,
		ScopeId:             nuuid.From(d.ScopeId),
		MerchantId:          nuuid.From(d.MerchantId),
		MethodCode:          d.MethodCode,
		ChannelCode:         null.StringFrom(d.ChannelCode),
		Currency:            d.Currency,
		MinAmount:           decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:           decimal.NewNullDecimal(d.MaxAmount),
		ProviderAccountId:   d.ProviderAccountId,
		ProviderMethodCode:  null.StringFrom(d.ProviderMethodCode),
		ProviderChannelCode: null.StringFrom(d.ProviderChannelCode),
		Priority:            d.Priority,
		IsFallback:          d.IsFallback,
		TrafficWeight:       d.TrafficWeight,
		TimeoutMs:           d.TimeoutMs,
		MaxAttempts:         d.MaxAttempts,
		IsEnabled:           d.IsEnabled,
		ConditionExpr:       d.ConditionExpr,
		Metadata:            d.Metadata,
	}
}

type PaymentRouteCandidatesRuntimeResponse struct {
	Id                  uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ScopeType           model.ScopeType `json:"scopeType" validate:"oneof=platform merchant" enums:"platform,merchant"`
	ScopeId             uuid.UUID       `json:"scopeId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantId          uuid.UUID       `json:"merchantId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodCode          string          `json:"methodCode" validate:"required"`
	ChannelCode         string          `json:"channelCode"`
	Currency            string          `json:"currency"`
	MinAmount           decimal.Decimal `json:"minAmount" format:"decimal" example:"100.50"`
	MaxAmount           decimal.Decimal `json:"maxAmount" format:"decimal" example:"100.50"`
	ProviderAccountId   uuid.UUID       `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderMethodCode  string          `json:"providerMethodCode"`
	ProviderChannelCode string          `json:"providerChannelCode"`
	Priority            int             `json:"priority" example:"1"`
	IsFallback          bool            `json:"isFallback" example:"true"`
	TrafficWeight       int             `json:"trafficWeight" example:"1"`
	TimeoutMs           int             `json:"timeoutMs" example:"1"`
	MaxAttempts         int             `json:"maxAttempts" example:"1"`
	IsEnabled           bool            `json:"isEnabled" example:"true"`
	ConditionExpr       json.RawMessage `json:"conditionExpr" swaggertype:"object"`
	Metadata            json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewPaymentRouteCandidatesRuntimeResponse(paymentRouteCandidatesRuntime model.PaymentRouteCandidatesRuntime) PaymentRouteCandidatesRuntimeResponse {
	return PaymentRouteCandidatesRuntimeResponse{
		Id:                  paymentRouteCandidatesRuntime.Id,
		ScopeType:           model.ScopeType(paymentRouteCandidatesRuntime.ScopeType),
		ScopeId:             paymentRouteCandidatesRuntime.ScopeId.UUID,
		MerchantId:          paymentRouteCandidatesRuntime.MerchantId.UUID,
		MethodCode:          paymentRouteCandidatesRuntime.MethodCode,
		ChannelCode:         paymentRouteCandidatesRuntime.ChannelCode.String,
		Currency:            paymentRouteCandidatesRuntime.Currency,
		MinAmount:           paymentRouteCandidatesRuntime.MinAmount.Decimal,
		MaxAmount:           paymentRouteCandidatesRuntime.MaxAmount.Decimal,
		ProviderAccountId:   paymentRouteCandidatesRuntime.ProviderAccountId,
		ProviderMethodCode:  paymentRouteCandidatesRuntime.ProviderMethodCode.String,
		ProviderChannelCode: paymentRouteCandidatesRuntime.ProviderChannelCode.String,
		Priority:            paymentRouteCandidatesRuntime.Priority,
		IsFallback:          paymentRouteCandidatesRuntime.IsFallback,
		TrafficWeight:       paymentRouteCandidatesRuntime.TrafficWeight,
		TimeoutMs:           paymentRouteCandidatesRuntime.TimeoutMs,
		MaxAttempts:         paymentRouteCandidatesRuntime.MaxAttempts,
		IsEnabled:           paymentRouteCandidatesRuntime.IsEnabled,
		ConditionExpr:       paymentRouteCandidatesRuntime.ConditionExpr,
		Metadata:            paymentRouteCandidatesRuntime.Metadata,
	}
}

type PaymentRouteCandidatesRuntimeListResponse []*PaymentRouteCandidatesRuntimeResponse

func NewPaymentRouteCandidatesRuntimeListResponse(paymentRouteCandidatesRuntimeList model.PaymentRouteCandidatesRuntimeList) PaymentRouteCandidatesRuntimeListResponse {
	dtoPaymentRouteCandidatesRuntimeListResponse := PaymentRouteCandidatesRuntimeListResponse{}
	for _, paymentRouteCandidatesRuntime := range paymentRouteCandidatesRuntimeList {
		dtoPaymentRouteCandidatesRuntimeResponse := NewPaymentRouteCandidatesRuntimeResponse(*paymentRouteCandidatesRuntime)
		dtoPaymentRouteCandidatesRuntimeListResponse = append(dtoPaymentRouteCandidatesRuntimeListResponse, &dtoPaymentRouteCandidatesRuntimeResponse)
	}
	return dtoPaymentRouteCandidatesRuntimeListResponse
}

type PaymentRouteCandidatesRuntimePrimaryIDList []PaymentRouteCandidatesRuntimePrimaryID

func (d PaymentRouteCandidatesRuntimePrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteCandidatesRuntime := range d {
		err = validator.Struct(paymentRouteCandidatesRuntime)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentRouteCandidatesRuntimePrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentRouteCandidatesRuntimePrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentRouteCandidatesRuntimePrimaryID) ToModel() model.PaymentRouteCandidatesRuntimePrimaryID {
	return model.PaymentRouteCandidatesRuntimePrimaryID{
		Id: d.Id,
	}
}

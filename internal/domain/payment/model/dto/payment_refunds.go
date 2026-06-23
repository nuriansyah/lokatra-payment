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

type PaymentRefundsDTOFieldNameType string

type paymentRefundsDTOFieldName struct {
	Id                PaymentRefundsDTOFieldNameType
	PaymentIntentId   PaymentRefundsDTOFieldNameType
	PaymentAttemptId  PaymentRefundsDTOFieldNameType
	RefundCode        PaymentRefundsDTOFieldNameType
	Amount            PaymentRefundsDTOFieldNameType
	Currency          PaymentRefundsDTOFieldNameType
	Reason            PaymentRefundsDTOFieldNameType
	Status            PaymentRefundsDTOFieldNameType
	ProviderRefundId  PaymentRefundsDTOFieldNameType
	ProviderReference PaymentRefundsDTOFieldNameType
	RequestedBy       PaymentRefundsDTOFieldNameType
	RequestedAt       PaymentRefundsDTOFieldNameType
	ApprovedBy        PaymentRefundsDTOFieldNameType
	ApprovedAt        PaymentRefundsDTOFieldNameType
	RejectedBy        PaymentRefundsDTOFieldNameType
	RejectedAt        PaymentRefundsDTOFieldNameType
	RejectionReason   PaymentRefundsDTOFieldNameType
	ProcessingAt      PaymentRefundsDTOFieldNameType
	SucceededAt       PaymentRefundsDTOFieldNameType
	FailedAt          PaymentRefundsDTOFieldNameType
	FailureCode       PaymentRefundsDTOFieldNameType
	FailureMessage    PaymentRefundsDTOFieldNameType
	RawRequest        PaymentRefundsDTOFieldNameType
	RawResponse       PaymentRefundsDTOFieldNameType
	Metadata          PaymentRefundsDTOFieldNameType
	MetaCreatedAt     PaymentRefundsDTOFieldNameType
	MetaCreatedBy     PaymentRefundsDTOFieldNameType
	MetaUpdatedAt     PaymentRefundsDTOFieldNameType
	MetaUpdatedBy     PaymentRefundsDTOFieldNameType
	MetaDeletedAt     PaymentRefundsDTOFieldNameType
	MetaDeletedBy     PaymentRefundsDTOFieldNameType
}

var PaymentRefundsDTOFieldName = paymentRefundsDTOFieldName{
	Id:                "id",
	PaymentIntentId:   "paymentIntentId",
	PaymentAttemptId:  "paymentAttemptId",
	RefundCode:        "refundCode",
	Amount:            "amount",
	Currency:          "currency",
	Reason:            "reason",
	Status:            "status",
	ProviderRefundId:  "providerRefundId",
	ProviderReference: "providerReference",
	RequestedBy:       "requestedBy",
	RequestedAt:       "requestedAt",
	ApprovedBy:        "approvedBy",
	ApprovedAt:        "approvedAt",
	RejectedBy:        "rejectedBy",
	RejectedAt:        "rejectedAt",
	RejectionReason:   "rejectionReason",
	ProcessingAt:      "processingAt",
	SucceededAt:       "succeededAt",
	FailedAt:          "failedAt",
	FailureCode:       "failureCode",
	FailureMessage:    "failureMessage",
	RawRequest:        "rawRequest",
	RawResponse:       "rawResponse",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformPaymentRefundsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentRefundsDTOFieldName.Id):
		return string(model.PaymentRefundsDBFieldName.Id), true

	case string(PaymentRefundsDTOFieldName.PaymentIntentId):
		return string(model.PaymentRefundsDBFieldName.PaymentIntentId), true

	case string(PaymentRefundsDTOFieldName.PaymentAttemptId):
		return string(model.PaymentRefundsDBFieldName.PaymentAttemptId), true

	case string(PaymentRefundsDTOFieldName.RefundCode):
		return string(model.PaymentRefundsDBFieldName.RefundCode), true

	case string(PaymentRefundsDTOFieldName.Amount):
		return string(model.PaymentRefundsDBFieldName.Amount), true

	case string(PaymentRefundsDTOFieldName.Currency):
		return string(model.PaymentRefundsDBFieldName.Currency), true

	case string(PaymentRefundsDTOFieldName.Reason):
		return string(model.PaymentRefundsDBFieldName.Reason), true

	case string(PaymentRefundsDTOFieldName.Status):
		return string(model.PaymentRefundsDBFieldName.Status), true

	case string(PaymentRefundsDTOFieldName.ProviderRefundId):
		return string(model.PaymentRefundsDBFieldName.ProviderRefundId), true

	case string(PaymentRefundsDTOFieldName.ProviderReference):
		return string(model.PaymentRefundsDBFieldName.ProviderReference), true

	case string(PaymentRefundsDTOFieldName.RequestedBy):
		return string(model.PaymentRefundsDBFieldName.RequestedBy), true

	case string(PaymentRefundsDTOFieldName.RequestedAt):
		return string(model.PaymentRefundsDBFieldName.RequestedAt), true

	case string(PaymentRefundsDTOFieldName.ApprovedBy):
		return string(model.PaymentRefundsDBFieldName.ApprovedBy), true

	case string(PaymentRefundsDTOFieldName.ApprovedAt):
		return string(model.PaymentRefundsDBFieldName.ApprovedAt), true

	case string(PaymentRefundsDTOFieldName.RejectedBy):
		return string(model.PaymentRefundsDBFieldName.RejectedBy), true

	case string(PaymentRefundsDTOFieldName.RejectedAt):
		return string(model.PaymentRefundsDBFieldName.RejectedAt), true

	case string(PaymentRefundsDTOFieldName.RejectionReason):
		return string(model.PaymentRefundsDBFieldName.RejectionReason), true

	case string(PaymentRefundsDTOFieldName.ProcessingAt):
		return string(model.PaymentRefundsDBFieldName.ProcessingAt), true

	case string(PaymentRefundsDTOFieldName.SucceededAt):
		return string(model.PaymentRefundsDBFieldName.SucceededAt), true

	case string(PaymentRefundsDTOFieldName.FailedAt):
		return string(model.PaymentRefundsDBFieldName.FailedAt), true

	case string(PaymentRefundsDTOFieldName.FailureCode):
		return string(model.PaymentRefundsDBFieldName.FailureCode), true

	case string(PaymentRefundsDTOFieldName.FailureMessage):
		return string(model.PaymentRefundsDBFieldName.FailureMessage), true

	case string(PaymentRefundsDTOFieldName.RawRequest):
		return string(model.PaymentRefundsDBFieldName.RawRequest), true

	case string(PaymentRefundsDTOFieldName.RawResponse):
		return string(model.PaymentRefundsDBFieldName.RawResponse), true

	case string(PaymentRefundsDTOFieldName.Metadata):
		return string(model.PaymentRefundsDBFieldName.Metadata), true

	case string(PaymentRefundsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentRefundsDBFieldName.MetaCreatedAt), true

	case string(PaymentRefundsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentRefundsDBFieldName.MetaCreatedBy), true

	case string(PaymentRefundsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentRefundsDBFieldName.MetaUpdatedAt), true

	case string(PaymentRefundsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentRefundsDBFieldName.MetaUpdatedBy), true

	case string(PaymentRefundsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentRefundsDBFieldName.MetaDeletedAt), true

	case string(PaymentRefundsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentRefundsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentRefundsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentRefundsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentRefundsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentRefundsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentRefundsProjectionOutputPath(path string) error {
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

func transformPaymentRefundsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentRefundsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentRefundsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentRefundsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentRefundsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentRefundsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentRefundsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentRefundsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentRefundsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentRefundsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentRefundsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentRefundsFilter(filter *model.Filter) {
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
			Field: string(PaymentRefundsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentRefundsSelectableResponse map[string]interface{}
type PaymentRefundsSelectableListResponse []*PaymentRefundsSelectableResponse

func assignPaymentRefundsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentRefundsSelectableValue(out PaymentRefundsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentRefundsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentRefundsSelectableResponse(paymentRefunds model.PaymentRefunds, filter model.Filter) PaymentRefundsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentRefundsDBFieldName.Id),
			string(model.PaymentRefundsDBFieldName.PaymentIntentId),
			string(model.PaymentRefundsDBFieldName.PaymentAttemptId),
			string(model.PaymentRefundsDBFieldName.RefundCode),
			string(model.PaymentRefundsDBFieldName.Amount),
			string(model.PaymentRefundsDBFieldName.Currency),
			string(model.PaymentRefundsDBFieldName.Reason),
			string(model.PaymentRefundsDBFieldName.Status),
			string(model.PaymentRefundsDBFieldName.ProviderRefundId),
			string(model.PaymentRefundsDBFieldName.ProviderReference),
			string(model.PaymentRefundsDBFieldName.RequestedBy),
			string(model.PaymentRefundsDBFieldName.RequestedAt),
			string(model.PaymentRefundsDBFieldName.ApprovedBy),
			string(model.PaymentRefundsDBFieldName.ApprovedAt),
			string(model.PaymentRefundsDBFieldName.RejectedBy),
			string(model.PaymentRefundsDBFieldName.RejectedAt),
			string(model.PaymentRefundsDBFieldName.RejectionReason),
			string(model.PaymentRefundsDBFieldName.ProcessingAt),
			string(model.PaymentRefundsDBFieldName.SucceededAt),
			string(model.PaymentRefundsDBFieldName.FailedAt),
			string(model.PaymentRefundsDBFieldName.FailureCode),
			string(model.PaymentRefundsDBFieldName.FailureMessage),
			string(model.PaymentRefundsDBFieldName.RawRequest),
			string(model.PaymentRefundsDBFieldName.RawResponse),
			string(model.PaymentRefundsDBFieldName.Metadata),
			string(model.PaymentRefundsDBFieldName.MetaCreatedAt),
			string(model.PaymentRefundsDBFieldName.MetaCreatedBy),
			string(model.PaymentRefundsDBFieldName.MetaUpdatedAt),
			string(model.PaymentRefundsDBFieldName.MetaUpdatedBy),
			string(model.PaymentRefundsDBFieldName.MetaDeletedAt),
			string(model.PaymentRefundsDBFieldName.MetaDeletedBy),
		)
	}
	paymentRefundsSelectableResponse := PaymentRefundsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentRefundsDBFieldName.Id):
			key := string(PaymentRefundsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.Id, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.PaymentIntentId):
			key := string(PaymentRefundsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.PaymentIntentId, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.PaymentAttemptId):
			key := string(PaymentRefundsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.PaymentAttemptId.UUID, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RefundCode):
			key := string(PaymentRefundsDTOFieldName.RefundCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RefundCode, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.Amount):
			key := string(PaymentRefundsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.Amount, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.Currency):
			key := string(PaymentRefundsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.Currency, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.Reason):
			key := string(PaymentRefundsDTOFieldName.Reason)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.Reason.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.Status):
			key := string(PaymentRefundsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, model.PaymentRefundStatus(paymentRefunds.Status), explicitAlias)

		case string(model.PaymentRefundsDBFieldName.ProviderRefundId):
			key := string(PaymentRefundsDTOFieldName.ProviderRefundId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.ProviderRefundId.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.ProviderReference):
			key := string(PaymentRefundsDTOFieldName.ProviderReference)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.ProviderReference.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RequestedBy):
			key := string(PaymentRefundsDTOFieldName.RequestedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RequestedBy, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RequestedAt):
			key := string(PaymentRefundsDTOFieldName.RequestedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RequestedAt, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.ApprovedBy):
			key := string(PaymentRefundsDTOFieldName.ApprovedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.ApprovedBy.UUID, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.ApprovedAt):
			key := string(PaymentRefundsDTOFieldName.ApprovedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.ApprovedAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RejectedBy):
			key := string(PaymentRefundsDTOFieldName.RejectedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RejectedBy.UUID, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RejectedAt):
			key := string(PaymentRefundsDTOFieldName.RejectedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RejectedAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RejectionReason):
			key := string(PaymentRefundsDTOFieldName.RejectionReason)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RejectionReason.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.ProcessingAt):
			key := string(PaymentRefundsDTOFieldName.ProcessingAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.ProcessingAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.SucceededAt):
			key := string(PaymentRefundsDTOFieldName.SucceededAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.SucceededAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.FailedAt):
			key := string(PaymentRefundsDTOFieldName.FailedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.FailedAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.FailureCode):
			key := string(PaymentRefundsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.FailureCode.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.FailureMessage):
			key := string(PaymentRefundsDTOFieldName.FailureMessage)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.FailureMessage.String, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RawRequest):
			key := string(PaymentRefundsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RawRequest, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.RawResponse):
			key := string(PaymentRefundsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.RawResponse, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.Metadata):
			key := string(PaymentRefundsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.Metadata, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaCreatedAt):
			key := string(PaymentRefundsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaCreatedAt, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaCreatedBy):
			key := string(PaymentRefundsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaCreatedBy, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaUpdatedAt):
			key := string(PaymentRefundsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaUpdatedBy):
			key := string(PaymentRefundsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaDeletedAt):
			key := string(PaymentRefundsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentRefundsDBFieldName.MetaDeletedBy):
			key := string(PaymentRefundsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRefundsSelectableValue(paymentRefundsSelectableResponse, key, paymentRefunds.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentRefundsSelectableResponse
}

func NewPaymentRefundsListResponseFromFilterResult(result []model.PaymentRefundsFilterResult, filter model.Filter) PaymentRefundsSelectableListResponse {
	dtoPaymentRefundsListResponse := PaymentRefundsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentRefundsResponse := NewPaymentRefundsSelectableResponse(row.PaymentRefunds, filter)
		dtoPaymentRefundsListResponse = append(dtoPaymentRefundsListResponse, &dtoPaymentRefundsResponse)
	}
	return dtoPaymentRefundsListResponse
}

type PaymentRefundsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PaymentRefundsSelectableListResponse `json:"data"`
}

func reversePaymentRefundsFilterResults(result []model.PaymentRefundsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentRefundsFilterResponse(result []model.PaymentRefundsFilterResult, filter model.Filter) (resp PaymentRefundsFilterResponse) {
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
			reversePaymentRefundsFilterResults(dataResult)
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

	resp.Data = NewPaymentRefundsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentRefundsCreateRequest struct {
	PaymentIntentId   uuid.UUID                 `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID                 `json:"paymentAttemptId"`
	RefundCode        string                    `json:"refundCode"`
	Amount            decimal.Decimal           `json:"amount"`
	Currency          string                    `json:"currency"`
	Reason            string                    `json:"reason"`
	Status            model.PaymentRefundStatus `json:"status" example:"requested" enums:"requested,approved,rejected,processing,succeeded,failed"`
	ProviderRefundId  string                    `json:"providerRefundId"`
	ProviderReference string                    `json:"providerReference"`
	RequestedBy       uuid.UUID                 `json:"requestedBy"`
	RequestedAt       time.Time                 `json:"requestedAt"`
	ApprovedBy        uuid.UUID                 `json:"approvedBy"`
	ApprovedAt        time.Time                 `json:"approvedAt"`
	RejectedBy        uuid.UUID                 `json:"rejectedBy"`
	RejectedAt        time.Time                 `json:"rejectedAt"`
	RejectionReason   string                    `json:"rejectionReason"`
	ProcessingAt      time.Time                 `json:"processingAt"`
	SucceededAt       time.Time                 `json:"succeededAt"`
	FailedAt          time.Time                 `json:"failedAt"`
	FailureCode       string                    `json:"failureCode"`
	FailureMessage    string                    `json:"failureMessage"`
	RawRequest        json.RawMessage           `json:"rawRequest"`
	RawResponse       json.RawMessage           `json:"rawResponse"`
	Metadata          json.RawMessage           `json:"metadata"`
}

func (d *PaymentRefundsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentRefundsCreateRequest) ToModel() model.PaymentRefunds {
	id, _ := uuid.NewV4()
	return model.PaymentRefunds{
		Id:                id,
		PaymentIntentId:   d.PaymentIntentId,
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		RefundCode:        d.RefundCode,
		Amount:            d.Amount,
		Currency:          d.Currency,
		Reason:            null.StringFrom(d.Reason),
		Status:            d.Status,
		ProviderRefundId:  null.StringFrom(d.ProviderRefundId),
		ProviderReference: null.StringFrom(d.ProviderReference),
		RequestedBy:       d.RequestedBy,
		RequestedAt:       d.RequestedAt,
		ApprovedBy:        nuuid.From(d.ApprovedBy),
		ApprovedAt:        null.TimeFrom(d.ApprovedAt),
		RejectedBy:        nuuid.From(d.RejectedBy),
		RejectedAt:        null.TimeFrom(d.RejectedAt),
		RejectionReason:   null.StringFrom(d.RejectionReason),
		ProcessingAt:      null.TimeFrom(d.ProcessingAt),
		SucceededAt:       null.TimeFrom(d.SucceededAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureMessage:    null.StringFrom(d.FailureMessage),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PaymentRefundsListCreateRequest []*PaymentRefundsCreateRequest

func (d PaymentRefundsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRefunds := range d {
		err = validator.Struct(paymentRefunds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRefundsListCreateRequest) ToModelList() []model.PaymentRefunds {
	out := make([]model.PaymentRefunds, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentRefundsUpdateRequest struct {
	PaymentIntentId   uuid.UUID                 `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID                 `json:"paymentAttemptId"`
	RefundCode        string                    `json:"refundCode"`
	Amount            decimal.Decimal           `json:"amount"`
	Currency          string                    `json:"currency"`
	Reason            string                    `json:"reason"`
	Status            model.PaymentRefundStatus `json:"status" example:"requested" enums:"requested,approved,rejected,processing,succeeded,failed"`
	ProviderRefundId  string                    `json:"providerRefundId"`
	ProviderReference string                    `json:"providerReference"`
	RequestedBy       uuid.UUID                 `json:"requestedBy"`
	RequestedAt       time.Time                 `json:"requestedAt"`
	ApprovedBy        uuid.UUID                 `json:"approvedBy"`
	ApprovedAt        time.Time                 `json:"approvedAt"`
	RejectedBy        uuid.UUID                 `json:"rejectedBy"`
	RejectedAt        time.Time                 `json:"rejectedAt"`
	RejectionReason   string                    `json:"rejectionReason"`
	ProcessingAt      time.Time                 `json:"processingAt"`
	SucceededAt       time.Time                 `json:"succeededAt"`
	FailedAt          time.Time                 `json:"failedAt"`
	FailureCode       string                    `json:"failureCode"`
	FailureMessage    string                    `json:"failureMessage"`
	RawRequest        json.RawMessage           `json:"rawRequest"`
	RawResponse       json.RawMessage           `json:"rawResponse"`
	Metadata          json.RawMessage           `json:"metadata"`
}

func (d *PaymentRefundsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentRefundsUpdateRequest) ToModel() model.PaymentRefunds {
	return model.PaymentRefunds{
		PaymentIntentId:   d.PaymentIntentId,
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		RefundCode:        d.RefundCode,
		Amount:            d.Amount,
		Currency:          d.Currency,
		Reason:            null.StringFrom(d.Reason),
		Status:            d.Status,
		ProviderRefundId:  null.StringFrom(d.ProviderRefundId),
		ProviderReference: null.StringFrom(d.ProviderReference),
		RequestedBy:       d.RequestedBy,
		RequestedAt:       d.RequestedAt,
		ApprovedBy:        nuuid.From(d.ApprovedBy),
		ApprovedAt:        null.TimeFrom(d.ApprovedAt),
		RejectedBy:        nuuid.From(d.RejectedBy),
		RejectedAt:        null.TimeFrom(d.RejectedAt),
		RejectionReason:   null.StringFrom(d.RejectionReason),
		ProcessingAt:      null.TimeFrom(d.ProcessingAt),
		SucceededAt:       null.TimeFrom(d.SucceededAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureMessage:    null.StringFrom(d.FailureMessage),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PaymentRefundsBulkUpdateRequest struct {
	Id                uuid.UUID                 `json:"id"`
	PaymentIntentId   uuid.UUID                 `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID                 `json:"paymentAttemptId"`
	RefundCode        string                    `json:"refundCode"`
	Amount            decimal.Decimal           `json:"amount"`
	Currency          string                    `json:"currency"`
	Reason            string                    `json:"reason"`
	Status            model.PaymentRefundStatus `json:"status" example:"requested" enums:"requested,approved,rejected,processing,succeeded,failed"`
	ProviderRefundId  string                    `json:"providerRefundId"`
	ProviderReference string                    `json:"providerReference"`
	RequestedBy       uuid.UUID                 `json:"requestedBy"`
	RequestedAt       time.Time                 `json:"requestedAt"`
	ApprovedBy        uuid.UUID                 `json:"approvedBy"`
	ApprovedAt        time.Time                 `json:"approvedAt"`
	RejectedBy        uuid.UUID                 `json:"rejectedBy"`
	RejectedAt        time.Time                 `json:"rejectedAt"`
	RejectionReason   string                    `json:"rejectionReason"`
	ProcessingAt      time.Time                 `json:"processingAt"`
	SucceededAt       time.Time                 `json:"succeededAt"`
	FailedAt          time.Time                 `json:"failedAt"`
	FailureCode       string                    `json:"failureCode"`
	FailureMessage    string                    `json:"failureMessage"`
	RawRequest        json.RawMessage           `json:"rawRequest"`
	RawResponse       json.RawMessage           `json:"rawResponse"`
	Metadata          json.RawMessage           `json:"metadata"`
}

func (d PaymentRefundsBulkUpdateRequest) PrimaryID() PaymentRefundsPrimaryID {
	return PaymentRefundsPrimaryID{
		Id: d.Id,
	}
}

type PaymentRefundsListBulkUpdateRequest []*PaymentRefundsBulkUpdateRequest

func (d PaymentRefundsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRefunds := range d {
		err = validator.Struct(paymentRefunds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRefundsBulkUpdateRequest) ToModel() model.PaymentRefunds {
	return model.PaymentRefunds{
		Id:                d.Id,
		PaymentIntentId:   d.PaymentIntentId,
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		RefundCode:        d.RefundCode,
		Amount:            d.Amount,
		Currency:          d.Currency,
		Reason:            null.StringFrom(d.Reason),
		Status:            d.Status,
		ProviderRefundId:  null.StringFrom(d.ProviderRefundId),
		ProviderReference: null.StringFrom(d.ProviderReference),
		RequestedBy:       d.RequestedBy,
		RequestedAt:       d.RequestedAt,
		ApprovedBy:        nuuid.From(d.ApprovedBy),
		ApprovedAt:        null.TimeFrom(d.ApprovedAt),
		RejectedBy:        nuuid.From(d.RejectedBy),
		RejectedAt:        null.TimeFrom(d.RejectedAt),
		RejectionReason:   null.StringFrom(d.RejectionReason),
		ProcessingAt:      null.TimeFrom(d.ProcessingAt),
		SucceededAt:       null.TimeFrom(d.SucceededAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureMessage:    null.StringFrom(d.FailureMessage),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PaymentRefundsResponse struct {
	Id                uuid.UUID                 `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId   uuid.UUID                 `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId  uuid.UUID                 `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundCode        string                    `json:"refundCode" validate:"required"`
	Amount            decimal.Decimal           `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency          string                    `json:"currency"`
	Reason            string                    `json:"reason"`
	Status            model.PaymentRefundStatus `json:"status" validate:"oneof=requested approved rejected processing succeeded failed" enums:"requested,approved,rejected,processing,succeeded,failed"`
	ProviderRefundId  string                    `json:"providerRefundId"`
	ProviderReference string                    `json:"providerReference"`
	RequestedBy       uuid.UUID                 `json:"requestedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RequestedAt       time.Time                 `json:"requestedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ApprovedBy        uuid.UUID                 `json:"approvedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ApprovedAt        time.Time                 `json:"approvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RejectedBy        uuid.UUID                 `json:"rejectedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RejectedAt        time.Time                 `json:"rejectedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RejectionReason   string                    `json:"rejectionReason"`
	ProcessingAt      time.Time                 `json:"processingAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	SucceededAt       time.Time                 `json:"succeededAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailedAt          time.Time                 `json:"failedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailureCode       string                    `json:"failureCode"`
	FailureMessage    string                    `json:"failureMessage"`
	RawRequest        json.RawMessage           `json:"rawRequest" swaggertype:"object"`
	RawResponse       json.RawMessage           `json:"rawResponse" swaggertype:"object"`
	Metadata          json.RawMessage           `json:"metadata" swaggertype:"object"`
}

func NewPaymentRefundsResponse(paymentRefunds model.PaymentRefunds) PaymentRefundsResponse {
	return PaymentRefundsResponse{
		Id:                paymentRefunds.Id,
		PaymentIntentId:   paymentRefunds.PaymentIntentId,
		PaymentAttemptId:  paymentRefunds.PaymentAttemptId.UUID,
		RefundCode:        paymentRefunds.RefundCode,
		Amount:            paymentRefunds.Amount,
		Currency:          paymentRefunds.Currency,
		Reason:            paymentRefunds.Reason.String,
		Status:            model.PaymentRefundStatus(paymentRefunds.Status),
		ProviderRefundId:  paymentRefunds.ProviderRefundId.String,
		ProviderReference: paymentRefunds.ProviderReference.String,
		RequestedBy:       paymentRefunds.RequestedBy,
		RequestedAt:       paymentRefunds.RequestedAt,
		ApprovedBy:        paymentRefunds.ApprovedBy.UUID,
		ApprovedAt:        paymentRefunds.ApprovedAt.Time,
		RejectedBy:        paymentRefunds.RejectedBy.UUID,
		RejectedAt:        paymentRefunds.RejectedAt.Time,
		RejectionReason:   paymentRefunds.RejectionReason.String,
		ProcessingAt:      paymentRefunds.ProcessingAt.Time,
		SucceededAt:       paymentRefunds.SucceededAt.Time,
		FailedAt:          paymentRefunds.FailedAt.Time,
		FailureCode:       paymentRefunds.FailureCode.String,
		FailureMessage:    paymentRefunds.FailureMessage.String,
		RawRequest:        paymentRefunds.RawRequest,
		RawResponse:       paymentRefunds.RawResponse,
		Metadata:          paymentRefunds.Metadata,
	}
}

type PaymentRefundsListResponse []*PaymentRefundsResponse

func NewPaymentRefundsListResponse(paymentRefundsList model.PaymentRefundsList) PaymentRefundsListResponse {
	dtoPaymentRefundsListResponse := PaymentRefundsListResponse{}
	for _, paymentRefunds := range paymentRefundsList {
		dtoPaymentRefundsResponse := NewPaymentRefundsResponse(*paymentRefunds)
		dtoPaymentRefundsListResponse = append(dtoPaymentRefundsListResponse, &dtoPaymentRefundsResponse)
	}
	return dtoPaymentRefundsListResponse
}

type PaymentRefundsPrimaryIDList []PaymentRefundsPrimaryID

func (d PaymentRefundsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRefunds := range d {
		err = validator.Struct(paymentRefunds)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentRefundsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentRefundsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentRefundsPrimaryID) ToModel() model.PaymentRefundsPrimaryID {
	return model.PaymentRefundsPrimaryID{
		Id: d.Id,
	}
}

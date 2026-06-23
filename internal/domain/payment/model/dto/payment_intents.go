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

type PaymentIntentsDTOFieldNameType string

type paymentIntentsDTOFieldName struct {
	Id                  PaymentIntentsDTOFieldNameType
	IntentCode          PaymentIntentsDTOFieldNameType
	SourceService       PaymentIntentsDTOFieldNameType
	SourceType          PaymentIntentsDTOFieldNameType
	SourceId            PaymentIntentsDTOFieldNameType
	MerchantId          PaymentIntentsDTOFieldNameType
	CustomerId          PaymentIntentsDTOFieldNameType
	Amount              PaymentIntentsDTOFieldNameType
	Currency            PaymentIntentsDTOFieldNameType
	Status              PaymentIntentsDTOFieldNameType
	SelectedMethodCode  PaymentIntentsDTOFieldNameType
	SelectedChannelCode PaymentIntentsDTOFieldNameType
	Description         PaymentIntentsDTOFieldNameType
	ExpiresAt           PaymentIntentsDTOFieldNameType
	PaidAt              PaymentIntentsDTOFieldNameType
	CanceledAt          PaymentIntentsDTOFieldNameType
	CancellationReason  PaymentIntentsDTOFieldNameType
	IdempotencyKey      PaymentIntentsDTOFieldNameType
	SourceSnapshot      PaymentIntentsDTOFieldNameType
	Metadata            PaymentIntentsDTOFieldNameType
	MetaCreatedAt       PaymentIntentsDTOFieldNameType
	MetaCreatedBy       PaymentIntentsDTOFieldNameType
	MetaUpdatedAt       PaymentIntentsDTOFieldNameType
	MetaUpdatedBy       PaymentIntentsDTOFieldNameType
	MetaDeletedAt       PaymentIntentsDTOFieldNameType
	MetaDeletedBy       PaymentIntentsDTOFieldNameType
}

var PaymentIntentsDTOFieldName = paymentIntentsDTOFieldName{
	Id:                  "id",
	IntentCode:          "intentCode",
	SourceService:       "sourceService",
	SourceType:          "sourceType",
	SourceId:            "sourceId",
	MerchantId:          "merchantId",
	CustomerId:          "customerId",
	Amount:              "amount",
	Currency:            "currency",
	Status:              "status",
	SelectedMethodCode:  "selectedMethodCode",
	SelectedChannelCode: "selectedChannelCode",
	Description:         "description",
	ExpiresAt:           "expiresAt",
	PaidAt:              "paidAt",
	CanceledAt:          "canceledAt",
	CancellationReason:  "cancellationReason",
	IdempotencyKey:      "idempotencyKey",
	SourceSnapshot:      "sourceSnapshot",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformPaymentIntentsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentIntentsDTOFieldName.Id):
		return string(model.PaymentIntentsDBFieldName.Id), true

	case string(PaymentIntentsDTOFieldName.IntentCode):
		return string(model.PaymentIntentsDBFieldName.IntentCode), true

	case string(PaymentIntentsDTOFieldName.SourceService):
		return string(model.PaymentIntentsDBFieldName.SourceService), true

	case string(PaymentIntentsDTOFieldName.SourceType):
		return string(model.PaymentIntentsDBFieldName.SourceType), true

	case string(PaymentIntentsDTOFieldName.SourceId):
		return string(model.PaymentIntentsDBFieldName.SourceId), true

	case string(PaymentIntentsDTOFieldName.MerchantId):
		return string(model.PaymentIntentsDBFieldName.MerchantId), true

	case string(PaymentIntentsDTOFieldName.CustomerId):
		return string(model.PaymentIntentsDBFieldName.CustomerId), true

	case string(PaymentIntentsDTOFieldName.Amount):
		return string(model.PaymentIntentsDBFieldName.Amount), true

	case string(PaymentIntentsDTOFieldName.Currency):
		return string(model.PaymentIntentsDBFieldName.Currency), true

	case string(PaymentIntentsDTOFieldName.Status):
		return string(model.PaymentIntentsDBFieldName.Status), true

	case string(PaymentIntentsDTOFieldName.SelectedMethodCode):
		return string(model.PaymentIntentsDBFieldName.SelectedMethodCode), true

	case string(PaymentIntentsDTOFieldName.SelectedChannelCode):
		return string(model.PaymentIntentsDBFieldName.SelectedChannelCode), true

	case string(PaymentIntentsDTOFieldName.Description):
		return string(model.PaymentIntentsDBFieldName.Description), true

	case string(PaymentIntentsDTOFieldName.ExpiresAt):
		return string(model.PaymentIntentsDBFieldName.ExpiresAt), true

	case string(PaymentIntentsDTOFieldName.PaidAt):
		return string(model.PaymentIntentsDBFieldName.PaidAt), true

	case string(PaymentIntentsDTOFieldName.CanceledAt):
		return string(model.PaymentIntentsDBFieldName.CanceledAt), true

	case string(PaymentIntentsDTOFieldName.CancellationReason):
		return string(model.PaymentIntentsDBFieldName.CancellationReason), true

	case string(PaymentIntentsDTOFieldName.IdempotencyKey):
		return string(model.PaymentIntentsDBFieldName.IdempotencyKey), true

	case string(PaymentIntentsDTOFieldName.SourceSnapshot):
		return string(model.PaymentIntentsDBFieldName.SourceSnapshot), true

	case string(PaymentIntentsDTOFieldName.Metadata):
		return string(model.PaymentIntentsDBFieldName.Metadata), true

	case string(PaymentIntentsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentIntentsDBFieldName.MetaCreatedAt), true

	case string(PaymentIntentsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentIntentsDBFieldName.MetaCreatedBy), true

	case string(PaymentIntentsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentIntentsDBFieldName.MetaUpdatedAt), true

	case string(PaymentIntentsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentIntentsDBFieldName.MetaUpdatedBy), true

	case string(PaymentIntentsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentIntentsDBFieldName.MetaDeletedAt), true

	case string(PaymentIntentsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentIntentsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentIntentsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentIntentsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentIntentsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentIntentsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentIntentsProjectionOutputPath(path string) error {
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

func transformPaymentIntentsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentIntentsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentIntentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentIntentsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentIntentsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentIntentsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentIntentsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentIntentsFilter(filter *model.Filter) {
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
			Field: string(PaymentIntentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentIntentsSelectableResponse map[string]interface{}
type PaymentIntentsSelectableListResponse []*PaymentIntentsSelectableResponse

func assignPaymentIntentsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentIntentsSelectableValue(out PaymentIntentsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentIntentsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentIntentsSelectableResponse(paymentIntents model.PaymentIntents, filter model.Filter) PaymentIntentsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentIntentsDBFieldName.Id),
			string(model.PaymentIntentsDBFieldName.IntentCode),
			string(model.PaymentIntentsDBFieldName.SourceService),
			string(model.PaymentIntentsDBFieldName.SourceType),
			string(model.PaymentIntentsDBFieldName.SourceId),
			string(model.PaymentIntentsDBFieldName.MerchantId),
			string(model.PaymentIntentsDBFieldName.CustomerId),
			string(model.PaymentIntentsDBFieldName.Amount),
			string(model.PaymentIntentsDBFieldName.Currency),
			string(model.PaymentIntentsDBFieldName.Status),
			string(model.PaymentIntentsDBFieldName.SelectedMethodCode),
			string(model.PaymentIntentsDBFieldName.SelectedChannelCode),
			string(model.PaymentIntentsDBFieldName.Description),
			string(model.PaymentIntentsDBFieldName.ExpiresAt),
			string(model.PaymentIntentsDBFieldName.PaidAt),
			string(model.PaymentIntentsDBFieldName.CanceledAt),
			string(model.PaymentIntentsDBFieldName.CancellationReason),
			string(model.PaymentIntentsDBFieldName.IdempotencyKey),
			string(model.PaymentIntentsDBFieldName.SourceSnapshot),
			string(model.PaymentIntentsDBFieldName.Metadata),
			string(model.PaymentIntentsDBFieldName.MetaCreatedAt),
			string(model.PaymentIntentsDBFieldName.MetaCreatedBy),
			string(model.PaymentIntentsDBFieldName.MetaUpdatedAt),
			string(model.PaymentIntentsDBFieldName.MetaUpdatedBy),
			string(model.PaymentIntentsDBFieldName.MetaDeletedAt),
			string(model.PaymentIntentsDBFieldName.MetaDeletedBy),
		)
	}
	paymentIntentsSelectableResponse := PaymentIntentsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentIntentsDBFieldName.Id):
			key := string(PaymentIntentsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.Id, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.IntentCode):
			key := string(PaymentIntentsDTOFieldName.IntentCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.IntentCode, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SourceService):
			key := string(PaymentIntentsDTOFieldName.SourceService)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SourceService, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SourceType):
			key := string(PaymentIntentsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SourceType, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SourceId):
			key := string(PaymentIntentsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SourceId, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MerchantId):
			key := string(PaymentIntentsDTOFieldName.MerchantId)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MerchantId, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.CustomerId):
			key := string(PaymentIntentsDTOFieldName.CustomerId)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.CustomerId.UUID, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.Amount):
			key := string(PaymentIntentsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.Amount, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.Currency):
			key := string(PaymentIntentsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.Currency, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.Status):
			key := string(PaymentIntentsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, model.PaymentIntentStatus(paymentIntents.Status), explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SelectedMethodCode):
			key := string(PaymentIntentsDTOFieldName.SelectedMethodCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SelectedMethodCode.String, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SelectedChannelCode):
			key := string(PaymentIntentsDTOFieldName.SelectedChannelCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SelectedChannelCode.String, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.Description):
			key := string(PaymentIntentsDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.Description.String, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.ExpiresAt):
			key := string(PaymentIntentsDTOFieldName.ExpiresAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.ExpiresAt.Time, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.PaidAt):
			key := string(PaymentIntentsDTOFieldName.PaidAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.PaidAt.Time, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.CanceledAt):
			key := string(PaymentIntentsDTOFieldName.CanceledAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.CanceledAt.Time, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.CancellationReason):
			key := string(PaymentIntentsDTOFieldName.CancellationReason)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.CancellationReason.String, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.IdempotencyKey):
			key := string(PaymentIntentsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.IdempotencyKey, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.SourceSnapshot):
			key := string(PaymentIntentsDTOFieldName.SourceSnapshot)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.SourceSnapshot, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.Metadata):
			key := string(PaymentIntentsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.Metadata, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaCreatedAt):
			key := string(PaymentIntentsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaCreatedAt, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaCreatedBy):
			key := string(PaymentIntentsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaCreatedBy, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaUpdatedAt):
			key := string(PaymentIntentsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaUpdatedBy):
			key := string(PaymentIntentsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaDeletedAt):
			key := string(PaymentIntentsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentIntentsDBFieldName.MetaDeletedBy):
			key := string(PaymentIntentsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentIntentsSelectableValue(paymentIntentsSelectableResponse, key, paymentIntents.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentIntentsSelectableResponse
}

func NewPaymentIntentsListResponseFromFilterResult(result []model.PaymentIntentsFilterResult, filter model.Filter) PaymentIntentsSelectableListResponse {
	dtoPaymentIntentsListResponse := PaymentIntentsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentIntentsResponse := NewPaymentIntentsSelectableResponse(row.PaymentIntents, filter)
		dtoPaymentIntentsListResponse = append(dtoPaymentIntentsListResponse, &dtoPaymentIntentsResponse)
	}
	return dtoPaymentIntentsListResponse
}

type PaymentIntentsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PaymentIntentsSelectableListResponse `json:"data"`
}

func reversePaymentIntentsFilterResults(result []model.PaymentIntentsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentIntentsFilterResponse(result []model.PaymentIntentsFilterResult, filter model.Filter) (resp PaymentIntentsFilterResponse) {
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
			reversePaymentIntentsFilterResults(dataResult)
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

	resp.Data = NewPaymentIntentsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentIntentsCreateRequest struct {
	IntentCode          string                    `json:"intentCode"`
	SourceService       string                    `json:"sourceService"`
	SourceType          string                    `json:"sourceType"`
	SourceId            uuid.UUID                 `json:"sourceId"`
	MerchantId          uuid.UUID                 `json:"merchantId"`
	CustomerId          uuid.UUID                 `json:"customerId"`
	Amount              decimal.Decimal           `json:"amount"`
	Currency            string                    `json:"currency"`
	Status              model.PaymentIntentStatus `json:"status" example:"requires_payment_method" enums:"requires_payment_method,requires_confirmation,requires_action,processing,succeeded,canceled"`
	SelectedMethodCode  string                    `json:"selectedMethodCode"`
	SelectedChannelCode string                    `json:"selectedChannelCode"`
	Description         string                    `json:"description"`
	ExpiresAt           time.Time                 `json:"expiresAt"`
	PaidAt              time.Time                 `json:"paidAt"`
	CanceledAt          time.Time                 `json:"canceledAt"`
	CancellationReason  string                    `json:"cancellationReason"`
	IdempotencyKey      string                    `json:"idempotencyKey"`
	SourceSnapshot      json.RawMessage           `json:"sourceSnapshot"`
	Metadata            json.RawMessage           `json:"metadata"`
}

func (d *PaymentIntentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentIntentsCreateRequest) ToModel() model.PaymentIntents {
	id, _ := uuid.NewV4()
	return model.PaymentIntents{
		Id:                  id,
		IntentCode:          d.IntentCode,
		SourceService:       d.SourceService,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		MerchantId:          d.MerchantId,
		CustomerId:          nuuid.From(d.CustomerId),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		SelectedMethodCode:  null.StringFrom(d.SelectedMethodCode),
		SelectedChannelCode: null.StringFrom(d.SelectedChannelCode),
		Description:         null.StringFrom(d.Description),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		PaidAt:              null.TimeFrom(d.PaidAt),
		CanceledAt:          null.TimeFrom(d.CanceledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		IdempotencyKey:      d.IdempotencyKey,
		SourceSnapshot:      d.SourceSnapshot,
		Metadata:            d.Metadata,
	}
}

type PaymentIntentsListCreateRequest []*PaymentIntentsCreateRequest

func (d PaymentIntentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentsListCreateRequest) ToModelList() []model.PaymentIntents {
	out := make([]model.PaymentIntents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentIntentsUpdateRequest struct {
	IntentCode          string                    `json:"intentCode"`
	SourceService       string                    `json:"sourceService"`
	SourceType          string                    `json:"sourceType"`
	SourceId            uuid.UUID                 `json:"sourceId"`
	MerchantId          uuid.UUID                 `json:"merchantId"`
	CustomerId          uuid.UUID                 `json:"customerId"`
	Amount              decimal.Decimal           `json:"amount"`
	Currency            string                    `json:"currency"`
	Status              model.PaymentIntentStatus `json:"status" example:"requires_payment_method" enums:"requires_payment_method,requires_confirmation,requires_action,processing,succeeded,canceled"`
	SelectedMethodCode  string                    `json:"selectedMethodCode"`
	SelectedChannelCode string                    `json:"selectedChannelCode"`
	Description         string                    `json:"description"`
	ExpiresAt           time.Time                 `json:"expiresAt"`
	PaidAt              time.Time                 `json:"paidAt"`
	CanceledAt          time.Time                 `json:"canceledAt"`
	CancellationReason  string                    `json:"cancellationReason"`
	IdempotencyKey      string                    `json:"idempotencyKey"`
	SourceSnapshot      json.RawMessage           `json:"sourceSnapshot"`
	Metadata            json.RawMessage           `json:"metadata"`
}

func (d *PaymentIntentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentIntentsUpdateRequest) ToModel() model.PaymentIntents {
	return model.PaymentIntents{
		IntentCode:          d.IntentCode,
		SourceService:       d.SourceService,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		MerchantId:          d.MerchantId,
		CustomerId:          nuuid.From(d.CustomerId),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		SelectedMethodCode:  null.StringFrom(d.SelectedMethodCode),
		SelectedChannelCode: null.StringFrom(d.SelectedChannelCode),
		Description:         null.StringFrom(d.Description),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		PaidAt:              null.TimeFrom(d.PaidAt),
		CanceledAt:          null.TimeFrom(d.CanceledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		IdempotencyKey:      d.IdempotencyKey,
		SourceSnapshot:      d.SourceSnapshot,
		Metadata:            d.Metadata,
	}
}

type PaymentIntentsBulkUpdateRequest struct {
	Id                  uuid.UUID                 `json:"id"`
	IntentCode          string                    `json:"intentCode"`
	SourceService       string                    `json:"sourceService"`
	SourceType          string                    `json:"sourceType"`
	SourceId            uuid.UUID                 `json:"sourceId"`
	MerchantId          uuid.UUID                 `json:"merchantId"`
	CustomerId          uuid.UUID                 `json:"customerId"`
	Amount              decimal.Decimal           `json:"amount"`
	Currency            string                    `json:"currency"`
	Status              model.PaymentIntentStatus `json:"status" example:"requires_payment_method" enums:"requires_payment_method,requires_confirmation,requires_action,processing,succeeded,canceled"`
	SelectedMethodCode  string                    `json:"selectedMethodCode"`
	SelectedChannelCode string                    `json:"selectedChannelCode"`
	Description         string                    `json:"description"`
	ExpiresAt           time.Time                 `json:"expiresAt"`
	PaidAt              time.Time                 `json:"paidAt"`
	CanceledAt          time.Time                 `json:"canceledAt"`
	CancellationReason  string                    `json:"cancellationReason"`
	IdempotencyKey      string                    `json:"idempotencyKey"`
	SourceSnapshot      json.RawMessage           `json:"sourceSnapshot"`
	Metadata            json.RawMessage           `json:"metadata"`
}

func (d PaymentIntentsBulkUpdateRequest) PrimaryID() PaymentIntentsPrimaryID {
	return PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentsListBulkUpdateRequest []*PaymentIntentsBulkUpdateRequest

func (d PaymentIntentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentsBulkUpdateRequest) ToModel() model.PaymentIntents {
	return model.PaymentIntents{
		Id:                  d.Id,
		IntentCode:          d.IntentCode,
		SourceService:       d.SourceService,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		MerchantId:          d.MerchantId,
		CustomerId:          nuuid.From(d.CustomerId),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		SelectedMethodCode:  null.StringFrom(d.SelectedMethodCode),
		SelectedChannelCode: null.StringFrom(d.SelectedChannelCode),
		Description:         null.StringFrom(d.Description),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		PaidAt:              null.TimeFrom(d.PaidAt),
		CanceledAt:          null.TimeFrom(d.CanceledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		IdempotencyKey:      d.IdempotencyKey,
		SourceSnapshot:      d.SourceSnapshot,
		Metadata:            d.Metadata,
	}
}

type PaymentIntentsResponse struct {
	Id                  uuid.UUID                 `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentCode          string                    `json:"intentCode" validate:"required"`
	SourceService       string                    `json:"sourceService" validate:"required"`
	SourceType          string                    `json:"sourceType" validate:"required"`
	SourceId            uuid.UUID                 `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantId          uuid.UUID                 `json:"merchantId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CustomerId          uuid.UUID                 `json:"customerId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount              decimal.Decimal           `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency            string                    `json:"currency"`
	Status              model.PaymentIntentStatus `json:"status" validate:"oneof=requires_payment_method requires_confirmation requires_action processing succeeded canceled" enums:"requires_payment_method,requires_confirmation,requires_action,processing,succeeded,canceled"`
	SelectedMethodCode  string                    `json:"selectedMethodCode"`
	SelectedChannelCode string                    `json:"selectedChannelCode"`
	Description         string                    `json:"description"`
	ExpiresAt           time.Time                 `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PaidAt              time.Time                 `json:"paidAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CanceledAt          time.Time                 `json:"canceledAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CancellationReason  string                    `json:"cancellationReason"`
	IdempotencyKey      string                    `json:"idempotencyKey" validate:"required"`
	SourceSnapshot      json.RawMessage           `json:"sourceSnapshot" swaggertype:"object"`
	Metadata            json.RawMessage           `json:"metadata" swaggertype:"object"`
}

func NewPaymentIntentsResponse(paymentIntents model.PaymentIntents) PaymentIntentsResponse {
	return PaymentIntentsResponse{
		Id:                  paymentIntents.Id,
		IntentCode:          paymentIntents.IntentCode,
		SourceService:       paymentIntents.SourceService,
		SourceType:          paymentIntents.SourceType,
		SourceId:            paymentIntents.SourceId,
		MerchantId:          paymentIntents.MerchantId,
		CustomerId:          paymentIntents.CustomerId.UUID,
		Amount:              paymentIntents.Amount,
		Currency:            paymentIntents.Currency,
		Status:              model.PaymentIntentStatus(paymentIntents.Status),
		SelectedMethodCode:  paymentIntents.SelectedMethodCode.String,
		SelectedChannelCode: paymentIntents.SelectedChannelCode.String,
		Description:         paymentIntents.Description.String,
		ExpiresAt:           paymentIntents.ExpiresAt.Time,
		PaidAt:              paymentIntents.PaidAt.Time,
		CanceledAt:          paymentIntents.CanceledAt.Time,
		CancellationReason:  paymentIntents.CancellationReason.String,
		IdempotencyKey:      paymentIntents.IdempotencyKey,
		SourceSnapshot:      paymentIntents.SourceSnapshot,
		Metadata:            paymentIntents.Metadata,
	}
}

type PaymentIntentsListResponse []*PaymentIntentsResponse

func NewPaymentIntentsListResponse(paymentIntentsList model.PaymentIntentsList) PaymentIntentsListResponse {
	dtoPaymentIntentsListResponse := PaymentIntentsListResponse{}
	for _, paymentIntents := range paymentIntentsList {
		dtoPaymentIntentsResponse := NewPaymentIntentsResponse(*paymentIntents)
		dtoPaymentIntentsListResponse = append(dtoPaymentIntentsListResponse, &dtoPaymentIntentsResponse)
	}
	return dtoPaymentIntentsListResponse
}

type PaymentIntentsPrimaryIDList []PaymentIntentsPrimaryID

func (d PaymentIntentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentIntentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentIntentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentIntentsPrimaryID) ToModel() model.PaymentIntentsPrimaryID {
	return model.PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

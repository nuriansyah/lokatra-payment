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

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentVoidsDTOFieldNameType string

type paymentVoidsDTOFieldName struct {
	Id                     PaymentVoidsDTOFieldNameType
	PaymentAuthorizationId PaymentVoidsDTOFieldNameType
	PaymentIntentId        PaymentVoidsDTOFieldNameType
	Amount                 PaymentVoidsDTOFieldNameType
	Currency               PaymentVoidsDTOFieldNameType
	Status                 PaymentVoidsDTOFieldNameType
	ProviderVoidId         PaymentVoidsDTOFieldNameType
	VoidedAt               PaymentVoidsDTOFieldNameType
	FailureCode            PaymentVoidsDTOFieldNameType
	FailureMessage         PaymentVoidsDTOFieldNameType
	RawRequest             PaymentVoidsDTOFieldNameType
	RawResponse            PaymentVoidsDTOFieldNameType
	Metadata               PaymentVoidsDTOFieldNameType
	MetaCreatedAt          PaymentVoidsDTOFieldNameType
	MetaCreatedBy          PaymentVoidsDTOFieldNameType
	MetaUpdatedAt          PaymentVoidsDTOFieldNameType
	MetaUpdatedBy          PaymentVoidsDTOFieldNameType
	MetaDeletedAt          PaymentVoidsDTOFieldNameType
	MetaDeletedBy          PaymentVoidsDTOFieldNameType
}

var PaymentVoidsDTOFieldName = paymentVoidsDTOFieldName{
	Id:                     "id",
	PaymentAuthorizationId: "paymentAuthorizationId",
	PaymentIntentId:        "paymentIntentId",
	Amount:                 "amount",
	Currency:               "currency",
	Status:                 "status",
	ProviderVoidId:         "providerVoidId",
	VoidedAt:               "voidedAt",
	FailureCode:            "failureCode",
	FailureMessage:         "failureMessage",
	RawRequest:             "rawRequest",
	RawResponse:            "rawResponse",
	Metadata:               "metadata",
	MetaCreatedAt:          "metaCreatedAt",
	MetaCreatedBy:          "metaCreatedBy",
	MetaUpdatedAt:          "metaUpdatedAt",
	MetaUpdatedBy:          "metaUpdatedBy",
	MetaDeletedAt:          "metaDeletedAt",
	MetaDeletedBy:          "metaDeletedBy",
}

func transformPaymentVoidsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentVoidsDTOFieldName.Id):
		return string(model.PaymentVoidsDBFieldName.Id), true

	case string(PaymentVoidsDTOFieldName.PaymentAuthorizationId):
		return string(model.PaymentVoidsDBFieldName.PaymentAuthorizationId), true

	case string(PaymentVoidsDTOFieldName.PaymentIntentId):
		return string(model.PaymentVoidsDBFieldName.PaymentIntentId), true

	case string(PaymentVoidsDTOFieldName.Amount):
		return string(model.PaymentVoidsDBFieldName.Amount), true

	case string(PaymentVoidsDTOFieldName.Currency):
		return string(model.PaymentVoidsDBFieldName.Currency), true

	case string(PaymentVoidsDTOFieldName.Status):
		return string(model.PaymentVoidsDBFieldName.Status), true

	case string(PaymentVoidsDTOFieldName.ProviderVoidId):
		return string(model.PaymentVoidsDBFieldName.ProviderVoidId), true

	case string(PaymentVoidsDTOFieldName.VoidedAt):
		return string(model.PaymentVoidsDBFieldName.VoidedAt), true

	case string(PaymentVoidsDTOFieldName.FailureCode):
		return string(model.PaymentVoidsDBFieldName.FailureCode), true

	case string(PaymentVoidsDTOFieldName.FailureMessage):
		return string(model.PaymentVoidsDBFieldName.FailureMessage), true

	case string(PaymentVoidsDTOFieldName.RawRequest):
		return string(model.PaymentVoidsDBFieldName.RawRequest), true

	case string(PaymentVoidsDTOFieldName.RawResponse):
		return string(model.PaymentVoidsDBFieldName.RawResponse), true

	case string(PaymentVoidsDTOFieldName.Metadata):
		return string(model.PaymentVoidsDBFieldName.Metadata), true

	case string(PaymentVoidsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentVoidsDBFieldName.MetaCreatedAt), true

	case string(PaymentVoidsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentVoidsDBFieldName.MetaCreatedBy), true

	case string(PaymentVoidsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentVoidsDBFieldName.MetaUpdatedAt), true

	case string(PaymentVoidsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentVoidsDBFieldName.MetaUpdatedBy), true

	case string(PaymentVoidsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentVoidsDBFieldName.MetaDeletedAt), true

	case string(PaymentVoidsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentVoidsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentVoidsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentVoidsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentVoidsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentVoidsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentVoidsProjectionOutputPath(path string) error {
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

func transformPaymentVoidsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentVoidsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentVoidsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentVoidsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentVoidsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentVoidsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentVoidsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentVoidsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentVoidsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentVoidsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentVoidsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentVoidsFilter(filter *model.Filter) {
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
			Field: string(PaymentVoidsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentVoidsSelectableResponse map[string]interface{}
type PaymentVoidsSelectableListResponse []*PaymentVoidsSelectableResponse

func assignPaymentVoidsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentVoidsSelectableValue(out PaymentVoidsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentVoidsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentVoidsSelectableResponse(paymentVoids model.PaymentVoids, filter model.Filter) PaymentVoidsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentVoidsDBFieldName.Id),
			string(model.PaymentVoidsDBFieldName.PaymentAuthorizationId),
			string(model.PaymentVoidsDBFieldName.PaymentIntentId),
			string(model.PaymentVoidsDBFieldName.Amount),
			string(model.PaymentVoidsDBFieldName.Currency),
			string(model.PaymentVoidsDBFieldName.Status),
			string(model.PaymentVoidsDBFieldName.ProviderVoidId),
			string(model.PaymentVoidsDBFieldName.VoidedAt),
			string(model.PaymentVoidsDBFieldName.FailureCode),
			string(model.PaymentVoidsDBFieldName.FailureMessage),
			string(model.PaymentVoidsDBFieldName.RawRequest),
			string(model.PaymentVoidsDBFieldName.RawResponse),
			string(model.PaymentVoidsDBFieldName.Metadata),
			string(model.PaymentVoidsDBFieldName.MetaCreatedAt),
			string(model.PaymentVoidsDBFieldName.MetaCreatedBy),
			string(model.PaymentVoidsDBFieldName.MetaUpdatedAt),
			string(model.PaymentVoidsDBFieldName.MetaUpdatedBy),
			string(model.PaymentVoidsDBFieldName.MetaDeletedAt),
			string(model.PaymentVoidsDBFieldName.MetaDeletedBy),
		)
	}
	paymentVoidsSelectableResponse := PaymentVoidsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentVoidsDBFieldName.Id):
			key := string(PaymentVoidsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.Id, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.PaymentAuthorizationId):
			key := string(PaymentVoidsDTOFieldName.PaymentAuthorizationId)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.PaymentAuthorizationId, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.PaymentIntentId):
			key := string(PaymentVoidsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.PaymentIntentId, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.Amount):
			key := string(PaymentVoidsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.Amount, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.Currency):
			key := string(PaymentVoidsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.Currency, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.Status):
			key := string(PaymentVoidsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, model.PaymentVoidStatus(paymentVoids.Status), explicitAlias)

		case string(model.PaymentVoidsDBFieldName.ProviderVoidId):
			key := string(PaymentVoidsDTOFieldName.ProviderVoidId)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.ProviderVoidId.String, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.VoidedAt):
			key := string(PaymentVoidsDTOFieldName.VoidedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.VoidedAt.Time, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.FailureCode):
			key := string(PaymentVoidsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.FailureCode.String, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.FailureMessage):
			key := string(PaymentVoidsDTOFieldName.FailureMessage)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.FailureMessage.String, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.RawRequest):
			key := string(PaymentVoidsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.RawRequest, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.RawResponse):
			key := string(PaymentVoidsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.RawResponse, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.Metadata):
			key := string(PaymentVoidsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.Metadata, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaCreatedAt):
			key := string(PaymentVoidsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaCreatedAt, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaCreatedBy):
			key := string(PaymentVoidsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaCreatedBy, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaUpdatedAt):
			key := string(PaymentVoidsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaUpdatedBy):
			key := string(PaymentVoidsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaDeletedAt):
			key := string(PaymentVoidsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentVoidsDBFieldName.MetaDeletedBy):
			key := string(PaymentVoidsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentVoidsSelectableValue(paymentVoidsSelectableResponse, key, paymentVoids.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentVoidsSelectableResponse
}

func NewPaymentVoidsListResponseFromFilterResult(result []model.PaymentVoidsFilterResult, filter model.Filter) PaymentVoidsSelectableListResponse {
	dtoPaymentVoidsListResponse := PaymentVoidsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentVoidsResponse := NewPaymentVoidsSelectableResponse(row.PaymentVoids, filter)
		dtoPaymentVoidsListResponse = append(dtoPaymentVoidsListResponse, &dtoPaymentVoidsResponse)
	}
	return dtoPaymentVoidsListResponse
}

type PaymentVoidsFilterResponse struct {
	Metadata Metadata                           `json:"metadata"`
	Data     PaymentVoidsSelectableListResponse `json:"data"`
}

func reversePaymentVoidsFilterResults(result []model.PaymentVoidsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentVoidsFilterResponse(result []model.PaymentVoidsFilterResult, filter model.Filter) (resp PaymentVoidsFilterResponse) {
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
			reversePaymentVoidsFilterResults(dataResult)
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

	resp.Data = NewPaymentVoidsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentVoidsCreateRequest struct {
	PaymentAuthorizationId uuid.UUID               `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID               `json:"paymentIntentId"`
	Amount                 decimal.Decimal         `json:"amount"`
	Currency               string                  `json:"currency"`
	Status                 model.PaymentVoidStatus `json:"status" example:"requested" enums:"requested,voided,failed"`
	ProviderVoidId         string                  `json:"providerVoidId"`
	VoidedAt               time.Time               `json:"voidedAt"`
	FailureCode            string                  `json:"failureCode"`
	FailureMessage         string                  `json:"failureMessage"`
	RawRequest             json.RawMessage         `json:"rawRequest"`
	RawResponse            json.RawMessage         `json:"rawResponse"`
	Metadata               json.RawMessage         `json:"metadata"`
}

func (d *PaymentVoidsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentVoidsCreateRequest) ToModel() model.PaymentVoids {
	id, _ := uuid.NewV4()
	return model.PaymentVoids{
		Id:                     id,
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderVoidId:         null.StringFrom(d.ProviderVoidId),
		VoidedAt:               null.TimeFrom(d.VoidedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentVoidsListCreateRequest []*PaymentVoidsCreateRequest

func (d PaymentVoidsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentVoids := range d {
		err = validator.Struct(paymentVoids)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentVoidsListCreateRequest) ToModelList() []model.PaymentVoids {
	out := make([]model.PaymentVoids, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentVoidsUpdateRequest struct {
	PaymentAuthorizationId uuid.UUID               `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID               `json:"paymentIntentId"`
	Amount                 decimal.Decimal         `json:"amount"`
	Currency               string                  `json:"currency"`
	Status                 model.PaymentVoidStatus `json:"status" example:"requested" enums:"requested,voided,failed"`
	ProviderVoidId         string                  `json:"providerVoidId"`
	VoidedAt               time.Time               `json:"voidedAt"`
	FailureCode            string                  `json:"failureCode"`
	FailureMessage         string                  `json:"failureMessage"`
	RawRequest             json.RawMessage         `json:"rawRequest"`
	RawResponse            json.RawMessage         `json:"rawResponse"`
	Metadata               json.RawMessage         `json:"metadata"`
}

func (d *PaymentVoidsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentVoidsUpdateRequest) ToModel() model.PaymentVoids {
	return model.PaymentVoids{
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderVoidId:         null.StringFrom(d.ProviderVoidId),
		VoidedAt:               null.TimeFrom(d.VoidedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentVoidsBulkUpdateRequest struct {
	Id                     uuid.UUID               `json:"id"`
	PaymentAuthorizationId uuid.UUID               `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID               `json:"paymentIntentId"`
	Amount                 decimal.Decimal         `json:"amount"`
	Currency               string                  `json:"currency"`
	Status                 model.PaymentVoidStatus `json:"status" example:"requested" enums:"requested,voided,failed"`
	ProviderVoidId         string                  `json:"providerVoidId"`
	VoidedAt               time.Time               `json:"voidedAt"`
	FailureCode            string                  `json:"failureCode"`
	FailureMessage         string                  `json:"failureMessage"`
	RawRequest             json.RawMessage         `json:"rawRequest"`
	RawResponse            json.RawMessage         `json:"rawResponse"`
	Metadata               json.RawMessage         `json:"metadata"`
}

func (d PaymentVoidsBulkUpdateRequest) PrimaryID() PaymentVoidsPrimaryID {
	return PaymentVoidsPrimaryID{
		Id: d.Id,
	}
}

type PaymentVoidsListBulkUpdateRequest []*PaymentVoidsBulkUpdateRequest

func (d PaymentVoidsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentVoids := range d {
		err = validator.Struct(paymentVoids)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentVoidsBulkUpdateRequest) ToModel() model.PaymentVoids {
	return model.PaymentVoids{
		Id:                     d.Id,
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderVoidId:         null.StringFrom(d.ProviderVoidId),
		VoidedAt:               null.TimeFrom(d.VoidedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentVoidsResponse struct {
	Id                     uuid.UUID               `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAuthorizationId uuid.UUID               `json:"paymentAuthorizationId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId        uuid.UUID               `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount                 decimal.Decimal         `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency               string                  `json:"currency"`
	Status                 model.PaymentVoidStatus `json:"status" validate:"oneof=requested voided failed" enums:"requested,voided,failed"`
	ProviderVoidId         string                  `json:"providerVoidId"`
	VoidedAt               time.Time               `json:"voidedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailureCode            string                  `json:"failureCode"`
	FailureMessage         string                  `json:"failureMessage"`
	RawRequest             json.RawMessage         `json:"rawRequest" swaggertype:"object"`
	RawResponse            json.RawMessage         `json:"rawResponse" swaggertype:"object"`
	Metadata               json.RawMessage         `json:"metadata" swaggertype:"object"`
}

func NewPaymentVoidsResponse(paymentVoids model.PaymentVoids) PaymentVoidsResponse {
	return PaymentVoidsResponse{
		Id:                     paymentVoids.Id,
		PaymentAuthorizationId: paymentVoids.PaymentAuthorizationId,
		PaymentIntentId:        paymentVoids.PaymentIntentId,
		Amount:                 paymentVoids.Amount,
		Currency:               paymentVoids.Currency,
		Status:                 model.PaymentVoidStatus(paymentVoids.Status),
		ProviderVoidId:         paymentVoids.ProviderVoidId.String,
		VoidedAt:               paymentVoids.VoidedAt.Time,
		FailureCode:            paymentVoids.FailureCode.String,
		FailureMessage:         paymentVoids.FailureMessage.String,
		RawRequest:             paymentVoids.RawRequest,
		RawResponse:            paymentVoids.RawResponse,
		Metadata:               paymentVoids.Metadata,
	}
}

type PaymentVoidsListResponse []*PaymentVoidsResponse

func NewPaymentVoidsListResponse(paymentVoidsList model.PaymentVoidsList) PaymentVoidsListResponse {
	dtoPaymentVoidsListResponse := PaymentVoidsListResponse{}
	for _, paymentVoids := range paymentVoidsList {
		dtoPaymentVoidsResponse := NewPaymentVoidsResponse(*paymentVoids)
		dtoPaymentVoidsListResponse = append(dtoPaymentVoidsListResponse, &dtoPaymentVoidsResponse)
	}
	return dtoPaymentVoidsListResponse
}

type PaymentVoidsPrimaryIDList []PaymentVoidsPrimaryID

func (d PaymentVoidsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentVoids := range d {
		err = validator.Struct(paymentVoids)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentVoidsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentVoidsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentVoidsPrimaryID) ToModel() model.PaymentVoidsPrimaryID {
	return model.PaymentVoidsPrimaryID{
		Id: d.Id,
	}
}

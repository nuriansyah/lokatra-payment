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

type PaymentCapturesDTOFieldNameType string

type paymentCapturesDTOFieldName struct {
	Id                     PaymentCapturesDTOFieldNameType
	PaymentAuthorizationId PaymentCapturesDTOFieldNameType
	PaymentIntentId        PaymentCapturesDTOFieldNameType
	Amount                 PaymentCapturesDTOFieldNameType
	Currency               PaymentCapturesDTOFieldNameType
	Status                 PaymentCapturesDTOFieldNameType
	ProviderCaptureId      PaymentCapturesDTOFieldNameType
	CapturedAt             PaymentCapturesDTOFieldNameType
	FailureCode            PaymentCapturesDTOFieldNameType
	FailureMessage         PaymentCapturesDTOFieldNameType
	RawRequest             PaymentCapturesDTOFieldNameType
	RawResponse            PaymentCapturesDTOFieldNameType
	Metadata               PaymentCapturesDTOFieldNameType
	MetaCreatedAt          PaymentCapturesDTOFieldNameType
	MetaCreatedBy          PaymentCapturesDTOFieldNameType
	MetaUpdatedAt          PaymentCapturesDTOFieldNameType
	MetaUpdatedBy          PaymentCapturesDTOFieldNameType
	MetaDeletedAt          PaymentCapturesDTOFieldNameType
	MetaDeletedBy          PaymentCapturesDTOFieldNameType
}

var PaymentCapturesDTOFieldName = paymentCapturesDTOFieldName{
	Id:                     "id",
	PaymentAuthorizationId: "paymentAuthorizationId",
	PaymentIntentId:        "paymentIntentId",
	Amount:                 "amount",
	Currency:               "currency",
	Status:                 "status",
	ProviderCaptureId:      "providerCaptureId",
	CapturedAt:             "capturedAt",
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

func transformPaymentCapturesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentCapturesDTOFieldName.Id):
		return string(model.PaymentCapturesDBFieldName.Id), true

	case string(PaymentCapturesDTOFieldName.PaymentAuthorizationId):
		return string(model.PaymentCapturesDBFieldName.PaymentAuthorizationId), true

	case string(PaymentCapturesDTOFieldName.PaymentIntentId):
		return string(model.PaymentCapturesDBFieldName.PaymentIntentId), true

	case string(PaymentCapturesDTOFieldName.Amount):
		return string(model.PaymentCapturesDBFieldName.Amount), true

	case string(PaymentCapturesDTOFieldName.Currency):
		return string(model.PaymentCapturesDBFieldName.Currency), true

	case string(PaymentCapturesDTOFieldName.Status):
		return string(model.PaymentCapturesDBFieldName.Status), true

	case string(PaymentCapturesDTOFieldName.ProviderCaptureId):
		return string(model.PaymentCapturesDBFieldName.ProviderCaptureId), true

	case string(PaymentCapturesDTOFieldName.CapturedAt):
		return string(model.PaymentCapturesDBFieldName.CapturedAt), true

	case string(PaymentCapturesDTOFieldName.FailureCode):
		return string(model.PaymentCapturesDBFieldName.FailureCode), true

	case string(PaymentCapturesDTOFieldName.FailureMessage):
		return string(model.PaymentCapturesDBFieldName.FailureMessage), true

	case string(PaymentCapturesDTOFieldName.RawRequest):
		return string(model.PaymentCapturesDBFieldName.RawRequest), true

	case string(PaymentCapturesDTOFieldName.RawResponse):
		return string(model.PaymentCapturesDBFieldName.RawResponse), true

	case string(PaymentCapturesDTOFieldName.Metadata):
		return string(model.PaymentCapturesDBFieldName.Metadata), true

	case string(PaymentCapturesDTOFieldName.MetaCreatedAt):
		return string(model.PaymentCapturesDBFieldName.MetaCreatedAt), true

	case string(PaymentCapturesDTOFieldName.MetaCreatedBy):
		return string(model.PaymentCapturesDBFieldName.MetaCreatedBy), true

	case string(PaymentCapturesDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentCapturesDBFieldName.MetaUpdatedAt), true

	case string(PaymentCapturesDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentCapturesDBFieldName.MetaUpdatedBy), true

	case string(PaymentCapturesDTOFieldName.MetaDeletedAt):
		return string(model.PaymentCapturesDBFieldName.MetaDeletedAt), true

	case string(PaymentCapturesDTOFieldName.MetaDeletedBy):
		return string(model.PaymentCapturesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentCapturesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentCapturesBaseFilterField(field string) bool {
	spec, found := model.NewPaymentCapturesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentCapturesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentCapturesProjectionOutputPath(path string) error {
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

func transformPaymentCapturesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentCapturesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentCapturesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentCapturesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentCapturesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentCapturesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentCapturesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentCapturesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentCapturesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentCapturesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentCapturesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentCapturesFilter(filter *model.Filter) {
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
			Field: string(PaymentCapturesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentCapturesSelectableResponse map[string]interface{}
type PaymentCapturesSelectableListResponse []*PaymentCapturesSelectableResponse

func assignPaymentCapturesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentCapturesSelectableValue(out PaymentCapturesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentCapturesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentCapturesSelectableResponse(paymentCaptures model.PaymentCaptures, filter model.Filter) PaymentCapturesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentCapturesDBFieldName.Id),
			string(model.PaymentCapturesDBFieldName.PaymentAuthorizationId),
			string(model.PaymentCapturesDBFieldName.PaymentIntentId),
			string(model.PaymentCapturesDBFieldName.Amount),
			string(model.PaymentCapturesDBFieldName.Currency),
			string(model.PaymentCapturesDBFieldName.Status),
			string(model.PaymentCapturesDBFieldName.ProviderCaptureId),
			string(model.PaymentCapturesDBFieldName.CapturedAt),
			string(model.PaymentCapturesDBFieldName.FailureCode),
			string(model.PaymentCapturesDBFieldName.FailureMessage),
			string(model.PaymentCapturesDBFieldName.RawRequest),
			string(model.PaymentCapturesDBFieldName.RawResponse),
			string(model.PaymentCapturesDBFieldName.Metadata),
			string(model.PaymentCapturesDBFieldName.MetaCreatedAt),
			string(model.PaymentCapturesDBFieldName.MetaCreatedBy),
			string(model.PaymentCapturesDBFieldName.MetaUpdatedAt),
			string(model.PaymentCapturesDBFieldName.MetaUpdatedBy),
			string(model.PaymentCapturesDBFieldName.MetaDeletedAt),
			string(model.PaymentCapturesDBFieldName.MetaDeletedBy),
		)
	}
	paymentCapturesSelectableResponse := PaymentCapturesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentCapturesDBFieldName.Id):
			key := string(PaymentCapturesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.Id, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.PaymentAuthorizationId):
			key := string(PaymentCapturesDTOFieldName.PaymentAuthorizationId)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.PaymentAuthorizationId, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.PaymentIntentId):
			key := string(PaymentCapturesDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.PaymentIntentId, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.Amount):
			key := string(PaymentCapturesDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.Amount, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.Currency):
			key := string(PaymentCapturesDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.Currency, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.Status):
			key := string(PaymentCapturesDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, model.PaymentCaptureStatus(paymentCaptures.Status), explicitAlias)

		case string(model.PaymentCapturesDBFieldName.ProviderCaptureId):
			key := string(PaymentCapturesDTOFieldName.ProviderCaptureId)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.ProviderCaptureId.String, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.CapturedAt):
			key := string(PaymentCapturesDTOFieldName.CapturedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.CapturedAt.Time, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.FailureCode):
			key := string(PaymentCapturesDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.FailureCode.String, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.FailureMessage):
			key := string(PaymentCapturesDTOFieldName.FailureMessage)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.FailureMessage.String, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.RawRequest):
			key := string(PaymentCapturesDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.RawRequest, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.RawResponse):
			key := string(PaymentCapturesDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.RawResponse, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.Metadata):
			key := string(PaymentCapturesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.Metadata, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaCreatedAt):
			key := string(PaymentCapturesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaCreatedAt, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaCreatedBy):
			key := string(PaymentCapturesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaCreatedBy, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaUpdatedAt):
			key := string(PaymentCapturesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaUpdatedBy):
			key := string(PaymentCapturesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaDeletedAt):
			key := string(PaymentCapturesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentCapturesDBFieldName.MetaDeletedBy):
			key := string(PaymentCapturesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentCapturesSelectableValue(paymentCapturesSelectableResponse, key, paymentCaptures.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentCapturesSelectableResponse
}

func NewPaymentCapturesListResponseFromFilterResult(result []model.PaymentCapturesFilterResult, filter model.Filter) PaymentCapturesSelectableListResponse {
	dtoPaymentCapturesListResponse := PaymentCapturesSelectableListResponse{}
	for _, row := range result {
		dtoPaymentCapturesResponse := NewPaymentCapturesSelectableResponse(row.PaymentCaptures, filter)
		dtoPaymentCapturesListResponse = append(dtoPaymentCapturesListResponse, &dtoPaymentCapturesResponse)
	}
	return dtoPaymentCapturesListResponse
}

type PaymentCapturesFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     PaymentCapturesSelectableListResponse `json:"data"`
}

func reversePaymentCapturesFilterResults(result []model.PaymentCapturesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentCapturesFilterResponse(result []model.PaymentCapturesFilterResult, filter model.Filter) (resp PaymentCapturesFilterResponse) {
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
			reversePaymentCapturesFilterResults(dataResult)
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

	resp.Data = NewPaymentCapturesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentCapturesCreateRequest struct {
	PaymentAuthorizationId uuid.UUID                  `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID                  `json:"paymentIntentId"`
	Amount                 decimal.Decimal            `json:"amount"`
	Currency               string                     `json:"currency"`
	Status                 model.PaymentCaptureStatus `json:"status" example:"requested" enums:"requested,captured,failed"`
	ProviderCaptureId      string                     `json:"providerCaptureId"`
	CapturedAt             time.Time                  `json:"capturedAt"`
	FailureCode            string                     `json:"failureCode"`
	FailureMessage         string                     `json:"failureMessage"`
	RawRequest             json.RawMessage            `json:"rawRequest"`
	RawResponse            json.RawMessage            `json:"rawResponse"`
	Metadata               json.RawMessage            `json:"metadata"`
}

func (d *PaymentCapturesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentCapturesCreateRequest) ToModel() model.PaymentCaptures {
	id, _ := uuid.NewV4()
	return model.PaymentCaptures{
		Id:                     id,
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderCaptureId:      null.StringFrom(d.ProviderCaptureId),
		CapturedAt:             null.TimeFrom(d.CapturedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentCapturesListCreateRequest []*PaymentCapturesCreateRequest

func (d PaymentCapturesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentCaptures := range d {
		err = validator.Struct(paymentCaptures)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentCapturesListCreateRequest) ToModelList() []model.PaymentCaptures {
	out := make([]model.PaymentCaptures, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentCapturesUpdateRequest struct {
	PaymentAuthorizationId uuid.UUID                  `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID                  `json:"paymentIntentId"`
	Amount                 decimal.Decimal            `json:"amount"`
	Currency               string                     `json:"currency"`
	Status                 model.PaymentCaptureStatus `json:"status" example:"requested" enums:"requested,captured,failed"`
	ProviderCaptureId      string                     `json:"providerCaptureId"`
	CapturedAt             time.Time                  `json:"capturedAt"`
	FailureCode            string                     `json:"failureCode"`
	FailureMessage         string                     `json:"failureMessage"`
	RawRequest             json.RawMessage            `json:"rawRequest"`
	RawResponse            json.RawMessage            `json:"rawResponse"`
	Metadata               json.RawMessage            `json:"metadata"`
}

func (d *PaymentCapturesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentCapturesUpdateRequest) ToModel() model.PaymentCaptures {
	return model.PaymentCaptures{
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderCaptureId:      null.StringFrom(d.ProviderCaptureId),
		CapturedAt:             null.TimeFrom(d.CapturedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentCapturesBulkUpdateRequest struct {
	Id                     uuid.UUID                  `json:"id"`
	PaymentAuthorizationId uuid.UUID                  `json:"paymentAuthorizationId"`
	PaymentIntentId        uuid.UUID                  `json:"paymentIntentId"`
	Amount                 decimal.Decimal            `json:"amount"`
	Currency               string                     `json:"currency"`
	Status                 model.PaymentCaptureStatus `json:"status" example:"requested" enums:"requested,captured,failed"`
	ProviderCaptureId      string                     `json:"providerCaptureId"`
	CapturedAt             time.Time                  `json:"capturedAt"`
	FailureCode            string                     `json:"failureCode"`
	FailureMessage         string                     `json:"failureMessage"`
	RawRequest             json.RawMessage            `json:"rawRequest"`
	RawResponse            json.RawMessage            `json:"rawResponse"`
	Metadata               json.RawMessage            `json:"metadata"`
}

func (d PaymentCapturesBulkUpdateRequest) PrimaryID() PaymentCapturesPrimaryID {
	return PaymentCapturesPrimaryID{
		Id: d.Id,
	}
}

type PaymentCapturesListBulkUpdateRequest []*PaymentCapturesBulkUpdateRequest

func (d PaymentCapturesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentCaptures := range d {
		err = validator.Struct(paymentCaptures)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentCapturesBulkUpdateRequest) ToModel() model.PaymentCaptures {
	return model.PaymentCaptures{
		Id:                     d.Id,
		PaymentAuthorizationId: d.PaymentAuthorizationId,
		PaymentIntentId:        d.PaymentIntentId,
		Amount:                 d.Amount,
		Currency:               d.Currency,
		Status:                 d.Status,
		ProviderCaptureId:      null.StringFrom(d.ProviderCaptureId),
		CapturedAt:             null.TimeFrom(d.CapturedAt),
		FailureCode:            null.StringFrom(d.FailureCode),
		FailureMessage:         null.StringFrom(d.FailureMessage),
		RawRequest:             d.RawRequest,
		RawResponse:            d.RawResponse,
		Metadata:               d.Metadata,
	}
}

type PaymentCapturesResponse struct {
	Id                     uuid.UUID                  `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAuthorizationId uuid.UUID                  `json:"paymentAuthorizationId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId        uuid.UUID                  `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount                 decimal.Decimal            `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency               string                     `json:"currency"`
	Status                 model.PaymentCaptureStatus `json:"status" validate:"oneof=requested captured failed" enums:"requested,captured,failed"`
	ProviderCaptureId      string                     `json:"providerCaptureId"`
	CapturedAt             time.Time                  `json:"capturedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailureCode            string                     `json:"failureCode"`
	FailureMessage         string                     `json:"failureMessage"`
	RawRequest             json.RawMessage            `json:"rawRequest" swaggertype:"object"`
	RawResponse            json.RawMessage            `json:"rawResponse" swaggertype:"object"`
	Metadata               json.RawMessage            `json:"metadata" swaggertype:"object"`
}

func NewPaymentCapturesResponse(paymentCaptures model.PaymentCaptures) PaymentCapturesResponse {
	return PaymentCapturesResponse{
		Id:                     paymentCaptures.Id,
		PaymentAuthorizationId: paymentCaptures.PaymentAuthorizationId,
		PaymentIntentId:        paymentCaptures.PaymentIntentId,
		Amount:                 paymentCaptures.Amount,
		Currency:               paymentCaptures.Currency,
		Status:                 model.PaymentCaptureStatus(paymentCaptures.Status),
		ProviderCaptureId:      paymentCaptures.ProviderCaptureId.String,
		CapturedAt:             paymentCaptures.CapturedAt.Time,
		FailureCode:            paymentCaptures.FailureCode.String,
		FailureMessage:         paymentCaptures.FailureMessage.String,
		RawRequest:             paymentCaptures.RawRequest,
		RawResponse:            paymentCaptures.RawResponse,
		Metadata:               paymentCaptures.Metadata,
	}
}

type PaymentCapturesListResponse []*PaymentCapturesResponse

func NewPaymentCapturesListResponse(paymentCapturesList model.PaymentCapturesList) PaymentCapturesListResponse {
	dtoPaymentCapturesListResponse := PaymentCapturesListResponse{}
	for _, paymentCaptures := range paymentCapturesList {
		dtoPaymentCapturesResponse := NewPaymentCapturesResponse(*paymentCaptures)
		dtoPaymentCapturesListResponse = append(dtoPaymentCapturesListResponse, &dtoPaymentCapturesResponse)
	}
	return dtoPaymentCapturesListResponse
}

type PaymentCapturesPrimaryIDList []PaymentCapturesPrimaryID

func (d PaymentCapturesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentCaptures := range d {
		err = validator.Struct(paymentCaptures)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentCapturesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentCapturesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentCapturesPrimaryID) ToModel() model.PaymentCapturesPrimaryID {
	return model.PaymentCapturesPrimaryID{
		Id: d.Id,
	}
}

package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderApiRequestsDTOFieldNameType string

type providerApiRequestsDTOFieldName struct {
	Id                ProviderApiRequestsDTOFieldNameType
	ProviderAccountId ProviderApiRequestsDTOFieldNameType
	PaymentIntentId   ProviderApiRequestsDTOFieldNameType
	PaymentAttemptId  ProviderApiRequestsDTOFieldNameType
	Operation         ProviderApiRequestsDTOFieldNameType
	IdempotencyKey    ProviderApiRequestsDTOFieldNameType
	RequestMethod     ProviderApiRequestsDTOFieldNameType
	RequestUrl        ProviderApiRequestsDTOFieldNameType
	RequestHeaders    ProviderApiRequestsDTOFieldNameType
	RequestBody       ProviderApiRequestsDTOFieldNameType
	ResponseStatus    ProviderApiRequestsDTOFieldNameType
	ResponseHeaders   ProviderApiRequestsDTOFieldNameType
	ResponseBody      ProviderApiRequestsDTOFieldNameType
	LatencyMs         ProviderApiRequestsDTOFieldNameType
	Success           ProviderApiRequestsDTOFieldNameType
	ErrorCode         ProviderApiRequestsDTOFieldNameType
	ErrorMessage      ProviderApiRequestsDTOFieldNameType
	Metadata          ProviderApiRequestsDTOFieldNameType
	MetaCreatedAt     ProviderApiRequestsDTOFieldNameType
	MetaCreatedBy     ProviderApiRequestsDTOFieldNameType
	MetaUpdatedAt     ProviderApiRequestsDTOFieldNameType
	MetaUpdatedBy     ProviderApiRequestsDTOFieldNameType
	MetaDeletedAt     ProviderApiRequestsDTOFieldNameType
	MetaDeletedBy     ProviderApiRequestsDTOFieldNameType
}

var ProviderApiRequestsDTOFieldName = providerApiRequestsDTOFieldName{
	Id:                "id",
	ProviderAccountId: "providerAccountId",
	PaymentIntentId:   "paymentIntentId",
	PaymentAttemptId:  "paymentAttemptId",
	Operation:         "operation",
	IdempotencyKey:    "idempotencyKey",
	RequestMethod:     "requestMethod",
	RequestUrl:        "requestUrl",
	RequestHeaders:    "requestHeaders",
	RequestBody:       "requestBody",
	ResponseStatus:    "responseStatus",
	ResponseHeaders:   "responseHeaders",
	ResponseBody:      "responseBody",
	LatencyMs:         "latencyMs",
	Success:           "success",
	ErrorCode:         "errorCode",
	ErrorMessage:      "errorMessage",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformProviderApiRequestsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderApiRequestsDTOFieldName.Id):
		return string(model.ProviderApiRequestsDBFieldName.Id), true

	case string(ProviderApiRequestsDTOFieldName.ProviderAccountId):
		return string(model.ProviderApiRequestsDBFieldName.ProviderAccountId), true

	case string(ProviderApiRequestsDTOFieldName.PaymentIntentId):
		return string(model.ProviderApiRequestsDBFieldName.PaymentIntentId), true

	case string(ProviderApiRequestsDTOFieldName.PaymentAttemptId):
		return string(model.ProviderApiRequestsDBFieldName.PaymentAttemptId), true

	case string(ProviderApiRequestsDTOFieldName.Operation):
		return string(model.ProviderApiRequestsDBFieldName.Operation), true

	case string(ProviderApiRequestsDTOFieldName.IdempotencyKey):
		return string(model.ProviderApiRequestsDBFieldName.IdempotencyKey), true

	case string(ProviderApiRequestsDTOFieldName.RequestMethod):
		return string(model.ProviderApiRequestsDBFieldName.RequestMethod), true

	case string(ProviderApiRequestsDTOFieldName.RequestUrl):
		return string(model.ProviderApiRequestsDBFieldName.RequestUrl), true

	case string(ProviderApiRequestsDTOFieldName.RequestHeaders):
		return string(model.ProviderApiRequestsDBFieldName.RequestHeaders), true

	case string(ProviderApiRequestsDTOFieldName.RequestBody):
		return string(model.ProviderApiRequestsDBFieldName.RequestBody), true

	case string(ProviderApiRequestsDTOFieldName.ResponseStatus):
		return string(model.ProviderApiRequestsDBFieldName.ResponseStatus), true

	case string(ProviderApiRequestsDTOFieldName.ResponseHeaders):
		return string(model.ProviderApiRequestsDBFieldName.ResponseHeaders), true

	case string(ProviderApiRequestsDTOFieldName.ResponseBody):
		return string(model.ProviderApiRequestsDBFieldName.ResponseBody), true

	case string(ProviderApiRequestsDTOFieldName.LatencyMs):
		return string(model.ProviderApiRequestsDBFieldName.LatencyMs), true

	case string(ProviderApiRequestsDTOFieldName.Success):
		return string(model.ProviderApiRequestsDBFieldName.Success), true

	case string(ProviderApiRequestsDTOFieldName.ErrorCode):
		return string(model.ProviderApiRequestsDBFieldName.ErrorCode), true

	case string(ProviderApiRequestsDTOFieldName.ErrorMessage):
		return string(model.ProviderApiRequestsDBFieldName.ErrorMessage), true

	case string(ProviderApiRequestsDTOFieldName.Metadata):
		return string(model.ProviderApiRequestsDBFieldName.Metadata), true

	case string(ProviderApiRequestsDTOFieldName.MetaCreatedAt):
		return string(model.ProviderApiRequestsDBFieldName.MetaCreatedAt), true

	case string(ProviderApiRequestsDTOFieldName.MetaCreatedBy):
		return string(model.ProviderApiRequestsDBFieldName.MetaCreatedBy), true

	case string(ProviderApiRequestsDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderApiRequestsDBFieldName.MetaUpdatedAt), true

	case string(ProviderApiRequestsDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderApiRequestsDBFieldName.MetaUpdatedBy), true

	case string(ProviderApiRequestsDTOFieldName.MetaDeletedAt):
		return string(model.ProviderApiRequestsDBFieldName.MetaDeletedAt), true

	case string(ProviderApiRequestsDTOFieldName.MetaDeletedBy):
		return string(model.ProviderApiRequestsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderApiRequestsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderApiRequestsBaseFilterField(field string) bool {
	spec, found := model.NewProviderApiRequestsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderApiRequestsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderApiRequestsProjectionOutputPath(path string) error {
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

func transformProviderApiRequestsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderApiRequestsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderApiRequestsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderApiRequestsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderApiRequestsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderApiRequestsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderApiRequestsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderApiRequestsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderApiRequestsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderApiRequestsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderApiRequestsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderApiRequestsFilter(filter *model.Filter) {
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
			Field: string(ProviderApiRequestsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderApiRequestsSelectableResponse map[string]interface{}
type ProviderApiRequestsSelectableListResponse []*ProviderApiRequestsSelectableResponse

func assignProviderApiRequestsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderApiRequestsSelectableValue(out ProviderApiRequestsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderApiRequestsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderApiRequestsSelectableResponse(providerApiRequests model.ProviderApiRequests, filter model.Filter) ProviderApiRequestsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderApiRequestsDBFieldName.Id),
			string(model.ProviderApiRequestsDBFieldName.ProviderAccountId),
			string(model.ProviderApiRequestsDBFieldName.PaymentIntentId),
			string(model.ProviderApiRequestsDBFieldName.PaymentAttemptId),
			string(model.ProviderApiRequestsDBFieldName.Operation),
			string(model.ProviderApiRequestsDBFieldName.IdempotencyKey),
			string(model.ProviderApiRequestsDBFieldName.RequestMethod),
			string(model.ProviderApiRequestsDBFieldName.RequestUrl),
			string(model.ProviderApiRequestsDBFieldName.RequestHeaders),
			string(model.ProviderApiRequestsDBFieldName.RequestBody),
			string(model.ProviderApiRequestsDBFieldName.ResponseStatus),
			string(model.ProviderApiRequestsDBFieldName.ResponseHeaders),
			string(model.ProviderApiRequestsDBFieldName.ResponseBody),
			string(model.ProviderApiRequestsDBFieldName.LatencyMs),
			string(model.ProviderApiRequestsDBFieldName.Success),
			string(model.ProviderApiRequestsDBFieldName.ErrorCode),
			string(model.ProviderApiRequestsDBFieldName.ErrorMessage),
			string(model.ProviderApiRequestsDBFieldName.Metadata),
			string(model.ProviderApiRequestsDBFieldName.MetaCreatedAt),
			string(model.ProviderApiRequestsDBFieldName.MetaCreatedBy),
			string(model.ProviderApiRequestsDBFieldName.MetaUpdatedAt),
			string(model.ProviderApiRequestsDBFieldName.MetaUpdatedBy),
			string(model.ProviderApiRequestsDBFieldName.MetaDeletedAt),
			string(model.ProviderApiRequestsDBFieldName.MetaDeletedBy),
		)
	}
	providerApiRequestsSelectableResponse := ProviderApiRequestsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderApiRequestsDBFieldName.Id):
			key := string(ProviderApiRequestsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.Id, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ProviderAccountId):
			key := string(ProviderApiRequestsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.ProviderAccountId.UUID, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.PaymentIntentId):
			key := string(ProviderApiRequestsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.PaymentIntentId.UUID, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.PaymentAttemptId):
			key := string(ProviderApiRequestsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.PaymentAttemptId.UUID, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.Operation):
			key := string(ProviderApiRequestsDTOFieldName.Operation)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.Operation, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.IdempotencyKey):
			key := string(ProviderApiRequestsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.IdempotencyKey.String, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.RequestMethod):
			key := string(ProviderApiRequestsDTOFieldName.RequestMethod)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.RequestMethod, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.RequestUrl):
			key := string(ProviderApiRequestsDTOFieldName.RequestUrl)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.RequestUrl, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.RequestHeaders):
			key := string(ProviderApiRequestsDTOFieldName.RequestHeaders)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.RequestHeaders, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.RequestBody):
			key := string(ProviderApiRequestsDTOFieldName.RequestBody)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.RequestBody, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ResponseStatus):
			key := string(ProviderApiRequestsDTOFieldName.ResponseStatus)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, int(providerApiRequests.ResponseStatus.ValueOrZero()), explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ResponseHeaders):
			key := string(ProviderApiRequestsDTOFieldName.ResponseHeaders)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.ResponseHeaders, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ResponseBody):
			key := string(ProviderApiRequestsDTOFieldName.ResponseBody)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.ResponseBody, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.LatencyMs):
			key := string(ProviderApiRequestsDTOFieldName.LatencyMs)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, int(providerApiRequests.LatencyMs.ValueOrZero()), explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.Success):
			key := string(ProviderApiRequestsDTOFieldName.Success)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.Success.Bool, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ErrorCode):
			key := string(ProviderApiRequestsDTOFieldName.ErrorCode)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.ErrorCode.String, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.ErrorMessage):
			key := string(ProviderApiRequestsDTOFieldName.ErrorMessage)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.ErrorMessage.String, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.Metadata):
			key := string(ProviderApiRequestsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.Metadata, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaCreatedAt):
			key := string(ProviderApiRequestsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaCreatedAt, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaCreatedBy):
			key := string(ProviderApiRequestsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaCreatedBy, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaUpdatedAt):
			key := string(ProviderApiRequestsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaUpdatedBy):
			key := string(ProviderApiRequestsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaDeletedAt):
			key := string(ProviderApiRequestsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderApiRequestsDBFieldName.MetaDeletedBy):
			key := string(ProviderApiRequestsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderApiRequestsSelectableValue(providerApiRequestsSelectableResponse, key, providerApiRequests.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return providerApiRequestsSelectableResponse
}

func NewProviderApiRequestsListResponseFromFilterResult(result []model.ProviderApiRequestsFilterResult, filter model.Filter) ProviderApiRequestsSelectableListResponse {
	dtoProviderApiRequestsListResponse := ProviderApiRequestsSelectableListResponse{}
	for _, row := range result {
		dtoProviderApiRequestsResponse := NewProviderApiRequestsSelectableResponse(row.ProviderApiRequests, filter)
		dtoProviderApiRequestsListResponse = append(dtoProviderApiRequestsListResponse, &dtoProviderApiRequestsResponse)
	}
	return dtoProviderApiRequestsListResponse
}

type ProviderApiRequestsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     ProviderApiRequestsSelectableListResponse `json:"data"`
}

func reverseProviderApiRequestsFilterResults(result []model.ProviderApiRequestsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderApiRequestsFilterResponse(result []model.ProviderApiRequestsFilterResult, filter model.Filter) (resp ProviderApiRequestsFilterResponse) {
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
			reverseProviderApiRequestsFilterResults(dataResult)
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

	resp.Data = NewProviderApiRequestsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderApiRequestsCreateRequest struct {
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID       `json:"paymentAttemptId"`
	Operation         string          `json:"operation"`
	IdempotencyKey    string          `json:"idempotencyKey"`
	RequestMethod     string          `json:"requestMethod"`
	RequestUrl        string          `json:"requestUrl"`
	RequestHeaders    json.RawMessage `json:"requestHeaders"`
	RequestBody       json.RawMessage `json:"requestBody"`
	ResponseStatus    int             `json:"responseStatus"`
	ResponseHeaders   json.RawMessage `json:"responseHeaders"`
	ResponseBody      json.RawMessage `json:"responseBody"`
	LatencyMs         int             `json:"latencyMs"`
	Success           bool            `json:"success"`
	ErrorCode         string          `json:"errorCode"`
	ErrorMessage      string          `json:"errorMessage"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *ProviderApiRequestsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderApiRequestsCreateRequest) ToModel() model.ProviderApiRequests {
	id, _ := uuid.NewV4()
	return model.ProviderApiRequests{
		Id:                id,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		PaymentIntentId:   nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		Operation:         d.Operation,
		IdempotencyKey:    null.StringFrom(d.IdempotencyKey),
		RequestMethod:     d.RequestMethod,
		RequestUrl:        d.RequestUrl,
		RequestHeaders:    d.RequestHeaders,
		RequestBody:       d.RequestBody,
		ResponseStatus:    null.IntFrom(int64(d.ResponseStatus)),
		ResponseHeaders:   d.ResponseHeaders,
		ResponseBody:      d.ResponseBody,
		LatencyMs:         null.IntFrom(int64(d.LatencyMs)),
		Success:           null.BoolFrom(d.Success),
		ErrorCode:         null.StringFrom(d.ErrorCode),
		ErrorMessage:      null.StringFrom(d.ErrorMessage),
		Metadata:          d.Metadata,
	}
}

type ProviderApiRequestsListCreateRequest []*ProviderApiRequestsCreateRequest

func (d ProviderApiRequestsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerApiRequests := range d {
		err = validator.Struct(providerApiRequests)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderApiRequestsListCreateRequest) ToModelList() []model.ProviderApiRequests {
	out := make([]model.ProviderApiRequests, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderApiRequestsUpdateRequest struct {
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID       `json:"paymentAttemptId"`
	Operation         string          `json:"operation"`
	IdempotencyKey    string          `json:"idempotencyKey"`
	RequestMethod     string          `json:"requestMethod"`
	RequestUrl        string          `json:"requestUrl"`
	RequestHeaders    json.RawMessage `json:"requestHeaders"`
	RequestBody       json.RawMessage `json:"requestBody"`
	ResponseStatus    int             `json:"responseStatus"`
	ResponseHeaders   json.RawMessage `json:"responseHeaders"`
	ResponseBody      json.RawMessage `json:"responseBody"`
	LatencyMs         int             `json:"latencyMs"`
	Success           bool            `json:"success"`
	ErrorCode         string          `json:"errorCode"`
	ErrorMessage      string          `json:"errorMessage"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *ProviderApiRequestsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderApiRequestsUpdateRequest) ToModel() model.ProviderApiRequests {
	return model.ProviderApiRequests{
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		PaymentIntentId:   nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		Operation:         d.Operation,
		IdempotencyKey:    null.StringFrom(d.IdempotencyKey),
		RequestMethod:     d.RequestMethod,
		RequestUrl:        d.RequestUrl,
		RequestHeaders:    d.RequestHeaders,
		RequestBody:       d.RequestBody,
		ResponseStatus:    null.IntFrom(int64(d.ResponseStatus)),
		ResponseHeaders:   d.ResponseHeaders,
		ResponseBody:      d.ResponseBody,
		LatencyMs:         null.IntFrom(int64(d.LatencyMs)),
		Success:           null.BoolFrom(d.Success),
		ErrorCode:         null.StringFrom(d.ErrorCode),
		ErrorMessage:      null.StringFrom(d.ErrorMessage),
		Metadata:          d.Metadata,
	}
}

type ProviderApiRequestsBulkUpdateRequest struct {
	Id                uuid.UUID       `json:"id"`
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId  uuid.UUID       `json:"paymentAttemptId"`
	Operation         string          `json:"operation"`
	IdempotencyKey    string          `json:"idempotencyKey"`
	RequestMethod     string          `json:"requestMethod"`
	RequestUrl        string          `json:"requestUrl"`
	RequestHeaders    json.RawMessage `json:"requestHeaders"`
	RequestBody       json.RawMessage `json:"requestBody"`
	ResponseStatus    int             `json:"responseStatus"`
	ResponseHeaders   json.RawMessage `json:"responseHeaders"`
	ResponseBody      json.RawMessage `json:"responseBody"`
	LatencyMs         int             `json:"latencyMs"`
	Success           bool            `json:"success"`
	ErrorCode         string          `json:"errorCode"`
	ErrorMessage      string          `json:"errorMessage"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d ProviderApiRequestsBulkUpdateRequest) PrimaryID() ProviderApiRequestsPrimaryID {
	return ProviderApiRequestsPrimaryID{
		Id: d.Id,
	}
}

type ProviderApiRequestsListBulkUpdateRequest []*ProviderApiRequestsBulkUpdateRequest

func (d ProviderApiRequestsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerApiRequests := range d {
		err = validator.Struct(providerApiRequests)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderApiRequestsBulkUpdateRequest) ToModel() model.ProviderApiRequests {
	return model.ProviderApiRequests{
		Id:                d.Id,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		PaymentIntentId:   nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:  nuuid.From(d.PaymentAttemptId),
		Operation:         d.Operation,
		IdempotencyKey:    null.StringFrom(d.IdempotencyKey),
		RequestMethod:     d.RequestMethod,
		RequestUrl:        d.RequestUrl,
		RequestHeaders:    d.RequestHeaders,
		RequestBody:       d.RequestBody,
		ResponseStatus:    null.IntFrom(int64(d.ResponseStatus)),
		ResponseHeaders:   d.ResponseHeaders,
		ResponseBody:      d.ResponseBody,
		LatencyMs:         null.IntFrom(int64(d.LatencyMs)),
		Success:           null.BoolFrom(d.Success),
		ErrorCode:         null.StringFrom(d.ErrorCode),
		ErrorMessage:      null.StringFrom(d.ErrorMessage),
		Metadata:          d.Metadata,
	}
}

type ProviderApiRequestsResponse struct {
	Id                uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID       `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId  uuid.UUID       `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Operation         string          `json:"operation" validate:"required"`
	IdempotencyKey    string          `json:"idempotencyKey"`
	RequestMethod     string          `json:"requestMethod" validate:"required"`
	RequestUrl        string          `json:"requestUrl" validate:"required,url"`
	RequestHeaders    json.RawMessage `json:"requestHeaders" swaggertype:"object"`
	RequestBody       json.RawMessage `json:"requestBody" swaggertype:"object"`
	ResponseStatus    int             `json:"responseStatus" example:"1"`
	ResponseHeaders   json.RawMessage `json:"responseHeaders" swaggertype:"object"`
	ResponseBody      json.RawMessage `json:"responseBody" swaggertype:"object"`
	LatencyMs         int             `json:"latencyMs" example:"1"`
	Success           bool            `json:"success" example:"true"`
	ErrorCode         string          `json:"errorCode"`
	ErrorMessage      string          `json:"errorMessage"`
	Metadata          json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewProviderApiRequestsResponse(providerApiRequests model.ProviderApiRequests) ProviderApiRequestsResponse {
	return ProviderApiRequestsResponse{
		Id:                providerApiRequests.Id,
		ProviderAccountId: providerApiRequests.ProviderAccountId.UUID,
		PaymentIntentId:   providerApiRequests.PaymentIntentId.UUID,
		PaymentAttemptId:  providerApiRequests.PaymentAttemptId.UUID,
		Operation:         providerApiRequests.Operation,
		IdempotencyKey:    providerApiRequests.IdempotencyKey.String,
		RequestMethod:     providerApiRequests.RequestMethod,
		RequestUrl:        providerApiRequests.RequestUrl,
		RequestHeaders:    providerApiRequests.RequestHeaders,
		RequestBody:       providerApiRequests.RequestBody,
		ResponseStatus:    int(providerApiRequests.ResponseStatus.ValueOrZero()),
		ResponseHeaders:   providerApiRequests.ResponseHeaders,
		ResponseBody:      providerApiRequests.ResponseBody,
		LatencyMs:         int(providerApiRequests.LatencyMs.ValueOrZero()),
		Success:           providerApiRequests.Success.Bool,
		ErrorCode:         providerApiRequests.ErrorCode.String,
		ErrorMessage:      providerApiRequests.ErrorMessage.String,
		Metadata:          providerApiRequests.Metadata,
	}
}

type ProviderApiRequestsListResponse []*ProviderApiRequestsResponse

func NewProviderApiRequestsListResponse(providerApiRequestsList model.ProviderApiRequestsList) ProviderApiRequestsListResponse {
	dtoProviderApiRequestsListResponse := ProviderApiRequestsListResponse{}
	for _, providerApiRequests := range providerApiRequestsList {
		dtoProviderApiRequestsResponse := NewProviderApiRequestsResponse(*providerApiRequests)
		dtoProviderApiRequestsListResponse = append(dtoProviderApiRequestsListResponse, &dtoProviderApiRequestsResponse)
	}
	return dtoProviderApiRequestsListResponse
}

type ProviderApiRequestsPrimaryIDList []ProviderApiRequestsPrimaryID

func (d ProviderApiRequestsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerApiRequests := range d {
		err = validator.Struct(providerApiRequests)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderApiRequestsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderApiRequestsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderApiRequestsPrimaryID) ToModel() model.ProviderApiRequestsPrimaryID {
	return model.ProviderApiRequestsPrimaryID{
		Id: d.Id,
	}
}

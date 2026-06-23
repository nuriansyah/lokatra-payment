package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceOutboxEventsDTOFieldNameType string

type financeOutboxEventsDTOFieldName struct {
	Id              FinanceOutboxEventsDTOFieldNameType
	AggregateType   FinanceOutboxEventsDTOFieldNameType
	AggregateId     FinanceOutboxEventsDTOFieldNameType
	EventType       FinanceOutboxEventsDTOFieldNameType
	EventVersion    FinanceOutboxEventsDTOFieldNameType
	IdempotencyKey  FinanceOutboxEventsDTOFieldNameType
	CorrelationId   FinanceOutboxEventsDTOFieldNameType
	Payload         FinanceOutboxEventsDTOFieldNameType
	Headers         FinanceOutboxEventsDTOFieldNameType
	PublishStatus   FinanceOutboxEventsDTOFieldNameType
	AttemptCount    FinanceOutboxEventsDTOFieldNameType
	NextAttemptAt   FinanceOutboxEventsDTOFieldNameType
	PublishedAt     FinanceOutboxEventsDTOFieldNameType
	LastErrorCode   FinanceOutboxEventsDTOFieldNameType
	LastErrorDetail FinanceOutboxEventsDTOFieldNameType
	Metadata        FinanceOutboxEventsDTOFieldNameType
	MetaCreatedAt   FinanceOutboxEventsDTOFieldNameType
	MetaCreatedBy   FinanceOutboxEventsDTOFieldNameType
	MetaUpdatedAt   FinanceOutboxEventsDTOFieldNameType
	MetaUpdatedBy   FinanceOutboxEventsDTOFieldNameType
	MetaDeletedAt   FinanceOutboxEventsDTOFieldNameType
	MetaDeletedBy   FinanceOutboxEventsDTOFieldNameType
}

var FinanceOutboxEventsDTOFieldName = financeOutboxEventsDTOFieldName{
	Id:              "id",
	AggregateType:   "aggregateType",
	AggregateId:     "aggregateId",
	EventType:       "eventType",
	EventVersion:    "eventVersion",
	IdempotencyKey:  "idempotencyKey",
	CorrelationId:   "correlationId",
	Payload:         "payload",
	Headers:         "headers",
	PublishStatus:   "publishStatus",
	AttemptCount:    "attemptCount",
	NextAttemptAt:   "nextAttemptAt",
	PublishedAt:     "publishedAt",
	LastErrorCode:   "lastErrorCode",
	LastErrorDetail: "lastErrorDetail",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformFinanceOutboxEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceOutboxEventsDTOFieldName.Id):
		return string(model.FinanceOutboxEventsDBFieldName.Id), true

	case string(FinanceOutboxEventsDTOFieldName.AggregateType):
		return string(model.FinanceOutboxEventsDBFieldName.AggregateType), true

	case string(FinanceOutboxEventsDTOFieldName.AggregateId):
		return string(model.FinanceOutboxEventsDBFieldName.AggregateId), true

	case string(FinanceOutboxEventsDTOFieldName.EventType):
		return string(model.FinanceOutboxEventsDBFieldName.EventType), true

	case string(FinanceOutboxEventsDTOFieldName.EventVersion):
		return string(model.FinanceOutboxEventsDBFieldName.EventVersion), true

	case string(FinanceOutboxEventsDTOFieldName.IdempotencyKey):
		return string(model.FinanceOutboxEventsDBFieldName.IdempotencyKey), true

	case string(FinanceOutboxEventsDTOFieldName.CorrelationId):
		return string(model.FinanceOutboxEventsDBFieldName.CorrelationId), true

	case string(FinanceOutboxEventsDTOFieldName.Payload):
		return string(model.FinanceOutboxEventsDBFieldName.Payload), true

	case string(FinanceOutboxEventsDTOFieldName.Headers):
		return string(model.FinanceOutboxEventsDBFieldName.Headers), true

	case string(FinanceOutboxEventsDTOFieldName.PublishStatus):
		return string(model.FinanceOutboxEventsDBFieldName.PublishStatus), true

	case string(FinanceOutboxEventsDTOFieldName.AttemptCount):
		return string(model.FinanceOutboxEventsDBFieldName.AttemptCount), true

	case string(FinanceOutboxEventsDTOFieldName.NextAttemptAt):
		return string(model.FinanceOutboxEventsDBFieldName.NextAttemptAt), true

	case string(FinanceOutboxEventsDTOFieldName.PublishedAt):
		return string(model.FinanceOutboxEventsDBFieldName.PublishedAt), true

	case string(FinanceOutboxEventsDTOFieldName.LastErrorCode):
		return string(model.FinanceOutboxEventsDBFieldName.LastErrorCode), true

	case string(FinanceOutboxEventsDTOFieldName.LastErrorDetail):
		return string(model.FinanceOutboxEventsDBFieldName.LastErrorDetail), true

	case string(FinanceOutboxEventsDTOFieldName.Metadata):
		return string(model.FinanceOutboxEventsDBFieldName.Metadata), true

	case string(FinanceOutboxEventsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceOutboxEventsDBFieldName.MetaCreatedAt), true

	case string(FinanceOutboxEventsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceOutboxEventsDBFieldName.MetaCreatedBy), true

	case string(FinanceOutboxEventsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedAt), true

	case string(FinanceOutboxEventsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedBy), true

	case string(FinanceOutboxEventsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceOutboxEventsDBFieldName.MetaDeletedAt), true

	case string(FinanceOutboxEventsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceOutboxEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceOutboxEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceOutboxEventsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceOutboxEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceOutboxEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceOutboxEventsProjectionOutputPath(path string) error {
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

func transformFinanceOutboxEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceOutboxEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceOutboxEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceOutboxEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceOutboxEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceOutboxEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceOutboxEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceOutboxEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceOutboxEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceOutboxEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceOutboxEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceOutboxEventsFilter(filter *model.Filter) {
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
			Field: string(FinanceOutboxEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceOutboxEventsSelectableResponse map[string]interface{}
type FinanceOutboxEventsSelectableListResponse []*FinanceOutboxEventsSelectableResponse

func assignFinanceOutboxEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceOutboxEventsSelectableValue(out FinanceOutboxEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceOutboxEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceOutboxEventsSelectableResponse(financeOutboxEvents model.FinanceOutboxEvents, filter model.Filter) FinanceOutboxEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceOutboxEventsDBFieldName.Id),
			string(model.FinanceOutboxEventsDBFieldName.AggregateType),
			string(model.FinanceOutboxEventsDBFieldName.AggregateId),
			string(model.FinanceOutboxEventsDBFieldName.EventType),
			string(model.FinanceOutboxEventsDBFieldName.EventVersion),
			string(model.FinanceOutboxEventsDBFieldName.IdempotencyKey),
			string(model.FinanceOutboxEventsDBFieldName.CorrelationId),
			string(model.FinanceOutboxEventsDBFieldName.Payload),
			string(model.FinanceOutboxEventsDBFieldName.Headers),
			string(model.FinanceOutboxEventsDBFieldName.PublishStatus),
			string(model.FinanceOutboxEventsDBFieldName.AttemptCount),
			string(model.FinanceOutboxEventsDBFieldName.NextAttemptAt),
			string(model.FinanceOutboxEventsDBFieldName.PublishedAt),
			string(model.FinanceOutboxEventsDBFieldName.LastErrorCode),
			string(model.FinanceOutboxEventsDBFieldName.LastErrorDetail),
			string(model.FinanceOutboxEventsDBFieldName.Metadata),
			string(model.FinanceOutboxEventsDBFieldName.MetaCreatedAt),
			string(model.FinanceOutboxEventsDBFieldName.MetaCreatedBy),
			string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedAt),
			string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedBy),
			string(model.FinanceOutboxEventsDBFieldName.MetaDeletedAt),
			string(model.FinanceOutboxEventsDBFieldName.MetaDeletedBy),
		)
	}
	financeOutboxEventsSelectableResponse := FinanceOutboxEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceOutboxEventsDBFieldName.Id):
			key := string(FinanceOutboxEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.Id, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.AggregateType):
			key := string(FinanceOutboxEventsDTOFieldName.AggregateType)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.AggregateType, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.AggregateId):
			key := string(FinanceOutboxEventsDTOFieldName.AggregateId)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.AggregateId, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.EventType):
			key := string(FinanceOutboxEventsDTOFieldName.EventType)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.EventType, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.EventVersion):
			key := string(FinanceOutboxEventsDTOFieldName.EventVersion)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.EventVersion, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.IdempotencyKey):
			key := string(FinanceOutboxEventsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.IdempotencyKey, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.CorrelationId):
			key := string(FinanceOutboxEventsDTOFieldName.CorrelationId)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.CorrelationId.UUID, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.Payload):
			key := string(FinanceOutboxEventsDTOFieldName.Payload)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.Payload, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.Headers):
			key := string(FinanceOutboxEventsDTOFieldName.Headers)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.Headers, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.PublishStatus):
			key := string(FinanceOutboxEventsDTOFieldName.PublishStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, model.PublishStatus(financeOutboxEvents.PublishStatus), explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.AttemptCount):
			key := string(FinanceOutboxEventsDTOFieldName.AttemptCount)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.AttemptCount, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.NextAttemptAt):
			key := string(FinanceOutboxEventsDTOFieldName.NextAttemptAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.NextAttemptAt.Time, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.PublishedAt):
			key := string(FinanceOutboxEventsDTOFieldName.PublishedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.PublishedAt.Time, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.LastErrorCode):
			key := string(FinanceOutboxEventsDTOFieldName.LastErrorCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.LastErrorCode.String, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.LastErrorDetail):
			key := string(FinanceOutboxEventsDTOFieldName.LastErrorDetail)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.LastErrorDetail.String, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.Metadata):
			key := string(FinanceOutboxEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.Metadata, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaCreatedAt):
			key := string(FinanceOutboxEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaCreatedAt, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaCreatedBy):
			key := string(FinanceOutboxEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaCreatedBy, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedAt):
			key := string(FinanceOutboxEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaUpdatedBy):
			key := string(FinanceOutboxEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaDeletedAt):
			key := string(FinanceOutboxEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceOutboxEventsDBFieldName.MetaDeletedBy):
			key := string(FinanceOutboxEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceOutboxEventsSelectableValue(financeOutboxEventsSelectableResponse, key, financeOutboxEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return financeOutboxEventsSelectableResponse
}

func NewFinanceOutboxEventsListResponseFromFilterResult(result []model.FinanceOutboxEventsFilterResult, filter model.Filter) FinanceOutboxEventsSelectableListResponse {
	dtoFinanceOutboxEventsListResponse := FinanceOutboxEventsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceOutboxEventsResponse := NewFinanceOutboxEventsSelectableResponse(row.FinanceOutboxEvents, filter)
		dtoFinanceOutboxEventsListResponse = append(dtoFinanceOutboxEventsListResponse, &dtoFinanceOutboxEventsResponse)
	}
	return dtoFinanceOutboxEventsListResponse
}

type FinanceOutboxEventsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     FinanceOutboxEventsSelectableListResponse `json:"data"`
}

func reverseFinanceOutboxEventsFilterResults(result []model.FinanceOutboxEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceOutboxEventsFilterResponse(result []model.FinanceOutboxEventsFilterResult, filter model.Filter) (resp FinanceOutboxEventsFilterResponse) {
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
			reverseFinanceOutboxEventsFilterResults(dataResult)
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

	resp.Data = NewFinanceOutboxEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceOutboxEventsCreateRequest struct {
	AggregateType   string              `json:"aggregateType"`
	AggregateId     uuid.UUID           `json:"aggregateId"`
	EventType       string              `json:"eventType"`
	EventVersion    int                 `json:"eventVersion"`
	IdempotencyKey  string              `json:"idempotencyKey"`
	CorrelationId   uuid.UUID           `json:"correlationId"`
	Payload         json.RawMessage     `json:"payload"`
	Headers         json.RawMessage     `json:"headers"`
	PublishStatus   model.PublishStatus `json:"publishStatus" example:"pending" enums:"pending,publishing,published,failed,dead"`
	AttemptCount    int                 `json:"attemptCount"`
	NextAttemptAt   time.Time           `json:"nextAttemptAt"`
	PublishedAt     time.Time           `json:"publishedAt"`
	LastErrorCode   string              `json:"lastErrorCode"`
	LastErrorDetail string              `json:"lastErrorDetail"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d *FinanceOutboxEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceOutboxEventsCreateRequest) ToModel() model.FinanceOutboxEvents {
	id, _ := uuid.NewV4()
	return model.FinanceOutboxEvents{
		Id:              id,
		AggregateType:   d.AggregateType,
		AggregateId:     d.AggregateId,
		EventType:       d.EventType,
		EventVersion:    d.EventVersion,
		IdempotencyKey:  d.IdempotencyKey,
		CorrelationId:   nuuid.From(d.CorrelationId),
		Payload:         d.Payload,
		Headers:         d.Headers,
		PublishStatus:   d.PublishStatus,
		AttemptCount:    d.AttemptCount,
		NextAttemptAt:   null.TimeFrom(d.NextAttemptAt),
		PublishedAt:     null.TimeFrom(d.PublishedAt),
		LastErrorCode:   null.StringFrom(d.LastErrorCode),
		LastErrorDetail: null.StringFrom(d.LastErrorDetail),
		Metadata:        d.Metadata,
	}
}

type FinanceOutboxEventsListCreateRequest []*FinanceOutboxEventsCreateRequest

func (d FinanceOutboxEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOutboxEvents := range d {
		err = validator.Struct(financeOutboxEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceOutboxEventsListCreateRequest) ToModelList() []model.FinanceOutboxEvents {
	out := make([]model.FinanceOutboxEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceOutboxEventsUpdateRequest struct {
	AggregateType   string              `json:"aggregateType"`
	AggregateId     uuid.UUID           `json:"aggregateId"`
	EventType       string              `json:"eventType"`
	EventVersion    int                 `json:"eventVersion"`
	IdempotencyKey  string              `json:"idempotencyKey"`
	CorrelationId   uuid.UUID           `json:"correlationId"`
	Payload         json.RawMessage     `json:"payload"`
	Headers         json.RawMessage     `json:"headers"`
	PublishStatus   model.PublishStatus `json:"publishStatus" example:"pending" enums:"pending,publishing,published,failed,dead"`
	AttemptCount    int                 `json:"attemptCount"`
	NextAttemptAt   time.Time           `json:"nextAttemptAt"`
	PublishedAt     time.Time           `json:"publishedAt"`
	LastErrorCode   string              `json:"lastErrorCode"`
	LastErrorDetail string              `json:"lastErrorDetail"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d *FinanceOutboxEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceOutboxEventsUpdateRequest) ToModel() model.FinanceOutboxEvents {
	return model.FinanceOutboxEvents{
		AggregateType:   d.AggregateType,
		AggregateId:     d.AggregateId,
		EventType:       d.EventType,
		EventVersion:    d.EventVersion,
		IdempotencyKey:  d.IdempotencyKey,
		CorrelationId:   nuuid.From(d.CorrelationId),
		Payload:         d.Payload,
		Headers:         d.Headers,
		PublishStatus:   d.PublishStatus,
		AttemptCount:    d.AttemptCount,
		NextAttemptAt:   null.TimeFrom(d.NextAttemptAt),
		PublishedAt:     null.TimeFrom(d.PublishedAt),
		LastErrorCode:   null.StringFrom(d.LastErrorCode),
		LastErrorDetail: null.StringFrom(d.LastErrorDetail),
		Metadata:        d.Metadata,
	}
}

type FinanceOutboxEventsBulkUpdateRequest struct {
	Id              uuid.UUID           `json:"id"`
	AggregateType   string              `json:"aggregateType"`
	AggregateId     uuid.UUID           `json:"aggregateId"`
	EventType       string              `json:"eventType"`
	EventVersion    int                 `json:"eventVersion"`
	IdempotencyKey  string              `json:"idempotencyKey"`
	CorrelationId   uuid.UUID           `json:"correlationId"`
	Payload         json.RawMessage     `json:"payload"`
	Headers         json.RawMessage     `json:"headers"`
	PublishStatus   model.PublishStatus `json:"publishStatus" example:"pending" enums:"pending,publishing,published,failed,dead"`
	AttemptCount    int                 `json:"attemptCount"`
	NextAttemptAt   time.Time           `json:"nextAttemptAt"`
	PublishedAt     time.Time           `json:"publishedAt"`
	LastErrorCode   string              `json:"lastErrorCode"`
	LastErrorDetail string              `json:"lastErrorDetail"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d FinanceOutboxEventsBulkUpdateRequest) PrimaryID() FinanceOutboxEventsPrimaryID {
	return FinanceOutboxEventsPrimaryID{
		Id: d.Id,
	}
}

type FinanceOutboxEventsListBulkUpdateRequest []*FinanceOutboxEventsBulkUpdateRequest

func (d FinanceOutboxEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOutboxEvents := range d {
		err = validator.Struct(financeOutboxEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceOutboxEventsBulkUpdateRequest) ToModel() model.FinanceOutboxEvents {
	return model.FinanceOutboxEvents{
		Id:              d.Id,
		AggregateType:   d.AggregateType,
		AggregateId:     d.AggregateId,
		EventType:       d.EventType,
		EventVersion:    d.EventVersion,
		IdempotencyKey:  d.IdempotencyKey,
		CorrelationId:   nuuid.From(d.CorrelationId),
		Payload:         d.Payload,
		Headers:         d.Headers,
		PublishStatus:   d.PublishStatus,
		AttemptCount:    d.AttemptCount,
		NextAttemptAt:   null.TimeFrom(d.NextAttemptAt),
		PublishedAt:     null.TimeFrom(d.PublishedAt),
		LastErrorCode:   null.StringFrom(d.LastErrorCode),
		LastErrorDetail: null.StringFrom(d.LastErrorDetail),
		Metadata:        d.Metadata,
	}
}

type FinanceOutboxEventsResponse struct {
	Id              uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AggregateType   string              `json:"aggregateType" validate:"required"`
	AggregateId     uuid.UUID           `json:"aggregateId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EventType       string              `json:"eventType" validate:"required"`
	EventVersion    int                 `json:"eventVersion" example:"1"`
	IdempotencyKey  string              `json:"idempotencyKey" validate:"required"`
	CorrelationId   uuid.UUID           `json:"correlationId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Payload         json.RawMessage     `json:"payload" validate:"required" swaggertype:"object"`
	Headers         json.RawMessage     `json:"headers" swaggertype:"object"`
	PublishStatus   model.PublishStatus `json:"publishStatus" validate:"oneof=pending publishing published failed dead" enums:"pending,publishing,published,failed,dead"`
	AttemptCount    int                 `json:"attemptCount" example:"1"`
	NextAttemptAt   time.Time           `json:"nextAttemptAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PublishedAt     time.Time           `json:"publishedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LastErrorCode   string              `json:"lastErrorCode"`
	LastErrorDetail string              `json:"lastErrorDetail"`
	Metadata        json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewFinanceOutboxEventsResponse(financeOutboxEvents model.FinanceOutboxEvents) FinanceOutboxEventsResponse {
	return FinanceOutboxEventsResponse{
		Id:              financeOutboxEvents.Id,
		AggregateType:   financeOutboxEvents.AggregateType,
		AggregateId:     financeOutboxEvents.AggregateId,
		EventType:       financeOutboxEvents.EventType,
		EventVersion:    financeOutboxEvents.EventVersion,
		IdempotencyKey:  financeOutboxEvents.IdempotencyKey,
		CorrelationId:   financeOutboxEvents.CorrelationId.UUID,
		Payload:         financeOutboxEvents.Payload,
		Headers:         financeOutboxEvents.Headers,
		PublishStatus:   model.PublishStatus(financeOutboxEvents.PublishStatus),
		AttemptCount:    financeOutboxEvents.AttemptCount,
		NextAttemptAt:   financeOutboxEvents.NextAttemptAt.Time,
		PublishedAt:     financeOutboxEvents.PublishedAt.Time,
		LastErrorCode:   financeOutboxEvents.LastErrorCode.String,
		LastErrorDetail: financeOutboxEvents.LastErrorDetail.String,
		Metadata:        financeOutboxEvents.Metadata,
	}
}

type FinanceOutboxEventsListResponse []*FinanceOutboxEventsResponse

func NewFinanceOutboxEventsListResponse(financeOutboxEventsList model.FinanceOutboxEventsList) FinanceOutboxEventsListResponse {
	dtoFinanceOutboxEventsListResponse := FinanceOutboxEventsListResponse{}
	for _, financeOutboxEvents := range financeOutboxEventsList {
		dtoFinanceOutboxEventsResponse := NewFinanceOutboxEventsResponse(*financeOutboxEvents)
		dtoFinanceOutboxEventsListResponse = append(dtoFinanceOutboxEventsListResponse, &dtoFinanceOutboxEventsResponse)
	}
	return dtoFinanceOutboxEventsListResponse
}

type FinanceOutboxEventsPrimaryIDList []FinanceOutboxEventsPrimaryID

func (d FinanceOutboxEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOutboxEvents := range d {
		err = validator.Struct(financeOutboxEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceOutboxEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceOutboxEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceOutboxEventsPrimaryID) ToModel() model.FinanceOutboxEventsPrimaryID {
	return model.FinanceOutboxEventsPrimaryID{
		Id: d.Id,
	}
}

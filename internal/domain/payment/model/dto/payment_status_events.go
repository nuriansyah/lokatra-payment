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

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentStatusEventsDTOFieldNameType string

type paymentStatusEventsDTOFieldName struct {
	Id                     PaymentStatusEventsDTOFieldNameType
	PaymentIntentId        PaymentStatusEventsDTOFieldNameType
	PaymentAttemptId       PaymentStatusEventsDTOFieldNameType
	ProviderWebhookEventId PaymentStatusEventsDTOFieldNameType
	SourceType             PaymentStatusEventsDTOFieldNameType
	EventType              PaymentStatusEventsDTOFieldNameType
	OldIntentStatus        PaymentStatusEventsDTOFieldNameType
	NewIntentStatus        PaymentStatusEventsDTOFieldNameType
	OldAttemptStatus       PaymentStatusEventsDTOFieldNameType
	NewAttemptStatus       PaymentStatusEventsDTOFieldNameType
	ProviderStatus         PaymentStatusEventsDTOFieldNameType
	Reason                 PaymentStatusEventsDTOFieldNameType
	OccurredAt             PaymentStatusEventsDTOFieldNameType
	Metadata               PaymentStatusEventsDTOFieldNameType
	MetaCreatedAt          PaymentStatusEventsDTOFieldNameType
	MetaCreatedBy          PaymentStatusEventsDTOFieldNameType
	MetaUpdatedAt          PaymentStatusEventsDTOFieldNameType
	MetaUpdatedBy          PaymentStatusEventsDTOFieldNameType
	MetaDeletedAt          PaymentStatusEventsDTOFieldNameType
	MetaDeletedBy          PaymentStatusEventsDTOFieldNameType
}

var PaymentStatusEventsDTOFieldName = paymentStatusEventsDTOFieldName{
	Id:                     "id",
	PaymentIntentId:        "paymentIntentId",
	PaymentAttemptId:       "paymentAttemptId",
	ProviderWebhookEventId: "providerWebhookEventId",
	SourceType:             "sourceType",
	EventType:              "eventType",
	OldIntentStatus:        "oldIntentStatus",
	NewIntentStatus:        "newIntentStatus",
	OldAttemptStatus:       "oldAttemptStatus",
	NewAttemptStatus:       "newAttemptStatus",
	ProviderStatus:         "providerStatus",
	Reason:                 "reason",
	OccurredAt:             "occurredAt",
	Metadata:               "metadata",
	MetaCreatedAt:          "metaCreatedAt",
	MetaCreatedBy:          "metaCreatedBy",
	MetaUpdatedAt:          "metaUpdatedAt",
	MetaUpdatedBy:          "metaUpdatedBy",
	MetaDeletedAt:          "metaDeletedAt",
	MetaDeletedBy:          "metaDeletedBy",
}

func transformPaymentStatusEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentStatusEventsDTOFieldName.Id):
		return string(model.PaymentStatusEventsDBFieldName.Id), true

	case string(PaymentStatusEventsDTOFieldName.PaymentIntentId):
		return string(model.PaymentStatusEventsDBFieldName.PaymentIntentId), true

	case string(PaymentStatusEventsDTOFieldName.PaymentAttemptId):
		return string(model.PaymentStatusEventsDBFieldName.PaymentAttemptId), true

	case string(PaymentStatusEventsDTOFieldName.ProviderWebhookEventId):
		return string(model.PaymentStatusEventsDBFieldName.ProviderWebhookEventId), true

	case string(PaymentStatusEventsDTOFieldName.SourceType):
		return string(model.PaymentStatusEventsDBFieldName.SourceType), true

	case string(PaymentStatusEventsDTOFieldName.EventType):
		return string(model.PaymentStatusEventsDBFieldName.EventType), true

	case string(PaymentStatusEventsDTOFieldName.OldIntentStatus):
		return string(model.PaymentStatusEventsDBFieldName.OldIntentStatus), true

	case string(PaymentStatusEventsDTOFieldName.NewIntentStatus):
		return string(model.PaymentStatusEventsDBFieldName.NewIntentStatus), true

	case string(PaymentStatusEventsDTOFieldName.OldAttemptStatus):
		return string(model.PaymentStatusEventsDBFieldName.OldAttemptStatus), true

	case string(PaymentStatusEventsDTOFieldName.NewAttemptStatus):
		return string(model.PaymentStatusEventsDBFieldName.NewAttemptStatus), true

	case string(PaymentStatusEventsDTOFieldName.ProviderStatus):
		return string(model.PaymentStatusEventsDBFieldName.ProviderStatus), true

	case string(PaymentStatusEventsDTOFieldName.Reason):
		return string(model.PaymentStatusEventsDBFieldName.Reason), true

	case string(PaymentStatusEventsDTOFieldName.OccurredAt):
		return string(model.PaymentStatusEventsDBFieldName.OccurredAt), true

	case string(PaymentStatusEventsDTOFieldName.Metadata):
		return string(model.PaymentStatusEventsDBFieldName.Metadata), true

	case string(PaymentStatusEventsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentStatusEventsDBFieldName.MetaCreatedAt), true

	case string(PaymentStatusEventsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentStatusEventsDBFieldName.MetaCreatedBy), true

	case string(PaymentStatusEventsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentStatusEventsDBFieldName.MetaUpdatedAt), true

	case string(PaymentStatusEventsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentStatusEventsDBFieldName.MetaUpdatedBy), true

	case string(PaymentStatusEventsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentStatusEventsDBFieldName.MetaDeletedAt), true

	case string(PaymentStatusEventsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentStatusEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentStatusEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentStatusEventsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentStatusEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentStatusEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentStatusEventsProjectionOutputPath(path string) error {
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

func transformPaymentStatusEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentStatusEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentStatusEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentStatusEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentStatusEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentStatusEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentStatusEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentStatusEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentStatusEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentStatusEventsFilter(filter *model.Filter) {
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
			Field: string(PaymentStatusEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentStatusEventsSelectableResponse map[string]interface{}
type PaymentStatusEventsSelectableListResponse []*PaymentStatusEventsSelectableResponse

func assignPaymentStatusEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentStatusEventsSelectableValue(out PaymentStatusEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentStatusEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentStatusEventsSelectableResponse(paymentStatusEvents model.PaymentStatusEvents, filter model.Filter) PaymentStatusEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentStatusEventsDBFieldName.Id),
			string(model.PaymentStatusEventsDBFieldName.PaymentIntentId),
			string(model.PaymentStatusEventsDBFieldName.PaymentAttemptId),
			string(model.PaymentStatusEventsDBFieldName.ProviderWebhookEventId),
			string(model.PaymentStatusEventsDBFieldName.SourceType),
			string(model.PaymentStatusEventsDBFieldName.EventType),
			string(model.PaymentStatusEventsDBFieldName.OldIntentStatus),
			string(model.PaymentStatusEventsDBFieldName.NewIntentStatus),
			string(model.PaymentStatusEventsDBFieldName.OldAttemptStatus),
			string(model.PaymentStatusEventsDBFieldName.NewAttemptStatus),
			string(model.PaymentStatusEventsDBFieldName.ProviderStatus),
			string(model.PaymentStatusEventsDBFieldName.Reason),
			string(model.PaymentStatusEventsDBFieldName.OccurredAt),
			string(model.PaymentStatusEventsDBFieldName.Metadata),
			string(model.PaymentStatusEventsDBFieldName.MetaCreatedAt),
			string(model.PaymentStatusEventsDBFieldName.MetaCreatedBy),
			string(model.PaymentStatusEventsDBFieldName.MetaUpdatedAt),
			string(model.PaymentStatusEventsDBFieldName.MetaUpdatedBy),
			string(model.PaymentStatusEventsDBFieldName.MetaDeletedAt),
			string(model.PaymentStatusEventsDBFieldName.MetaDeletedBy),
		)
	}
	paymentStatusEventsSelectableResponse := PaymentStatusEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentStatusEventsDBFieldName.Id):
			key := string(PaymentStatusEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.Id, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.PaymentIntentId):
			key := string(PaymentStatusEventsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.PaymentIntentId.UUID, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.PaymentAttemptId):
			key := string(PaymentStatusEventsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.PaymentAttemptId.UUID, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.ProviderWebhookEventId):
			key := string(PaymentStatusEventsDTOFieldName.ProviderWebhookEventId)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.ProviderWebhookEventId.UUID, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.SourceType):
			key := string(PaymentStatusEventsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.SourceType, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.EventType):
			key := string(PaymentStatusEventsDTOFieldName.EventType)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.EventType, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.OldIntentStatus):
			key := string(PaymentStatusEventsDTOFieldName.OldIntentStatus)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.OldIntentStatus.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.NewIntentStatus):
			key := string(PaymentStatusEventsDTOFieldName.NewIntentStatus)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.NewIntentStatus.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.OldAttemptStatus):
			key := string(PaymentStatusEventsDTOFieldName.OldAttemptStatus)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.OldAttemptStatus.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.NewAttemptStatus):
			key := string(PaymentStatusEventsDTOFieldName.NewAttemptStatus)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.NewAttemptStatus.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.ProviderStatus):
			key := string(PaymentStatusEventsDTOFieldName.ProviderStatus)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.ProviderStatus.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.Reason):
			key := string(PaymentStatusEventsDTOFieldName.Reason)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.Reason.String, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.OccurredAt):
			key := string(PaymentStatusEventsDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.OccurredAt, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.Metadata):
			key := string(PaymentStatusEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.Metadata, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaCreatedAt):
			key := string(PaymentStatusEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaCreatedAt, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaCreatedBy):
			key := string(PaymentStatusEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaCreatedBy, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaUpdatedAt):
			key := string(PaymentStatusEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaUpdatedBy):
			key := string(PaymentStatusEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaDeletedAt):
			key := string(PaymentStatusEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentStatusEventsDBFieldName.MetaDeletedBy):
			key := string(PaymentStatusEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentStatusEventsSelectableValue(paymentStatusEventsSelectableResponse, key, paymentStatusEvents.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentStatusEventsSelectableResponse
}

func NewPaymentStatusEventsListResponseFromFilterResult(result []model.PaymentStatusEventsFilterResult, filter model.Filter) PaymentStatusEventsSelectableListResponse {
	dtoPaymentStatusEventsListResponse := PaymentStatusEventsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentStatusEventsResponse := NewPaymentStatusEventsSelectableResponse(row.PaymentStatusEvents, filter)
		dtoPaymentStatusEventsListResponse = append(dtoPaymentStatusEventsListResponse, &dtoPaymentStatusEventsResponse)
	}
	return dtoPaymentStatusEventsListResponse
}

type PaymentStatusEventsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     PaymentStatusEventsSelectableListResponse `json:"data"`
}

func reversePaymentStatusEventsFilterResults(result []model.PaymentStatusEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentStatusEventsFilterResponse(result []model.PaymentStatusEventsFilterResult, filter model.Filter) (resp PaymentStatusEventsFilterResponse) {
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
			reversePaymentStatusEventsFilterResults(dataResult)
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

	resp.Data = NewPaymentStatusEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentStatusEventsCreateRequest struct {
	PaymentIntentId        uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId       uuid.UUID       `json:"paymentAttemptId"`
	ProviderWebhookEventId uuid.UUID       `json:"providerWebhookEventId"`
	SourceType             string          `json:"sourceType"`
	EventType              string          `json:"eventType"`
	OldIntentStatus        string          `json:"oldIntentStatus"`
	NewIntentStatus        string          `json:"newIntentStatus"`
	OldAttemptStatus       string          `json:"oldAttemptStatus"`
	NewAttemptStatus       string          `json:"newAttemptStatus"`
	ProviderStatus         string          `json:"providerStatus"`
	Reason                 string          `json:"reason"`
	OccurredAt             time.Time       `json:"occurredAt"`
	Metadata               json.RawMessage `json:"metadata"`
}

func (d *PaymentStatusEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentStatusEventsCreateRequest) ToModel() model.PaymentStatusEvents {
	id, _ := uuid.NewV4()
	return model.PaymentStatusEvents{
		Id:                     id,
		PaymentIntentId:        nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:       nuuid.From(d.PaymentAttemptId),
		ProviderWebhookEventId: nuuid.From(d.ProviderWebhookEventId),
		SourceType:             d.SourceType,
		EventType:              d.EventType,
		OldIntentStatus:        null.StringFrom(d.OldIntentStatus),
		NewIntentStatus:        null.StringFrom(d.NewIntentStatus),
		OldAttemptStatus:       null.StringFrom(d.OldAttemptStatus),
		NewAttemptStatus:       null.StringFrom(d.NewAttemptStatus),
		ProviderStatus:         null.StringFrom(d.ProviderStatus),
		Reason:                 null.StringFrom(d.Reason),
		OccurredAt:             d.OccurredAt,
		Metadata:               d.Metadata,
	}
}

type PaymentStatusEventsListCreateRequest []*PaymentStatusEventsCreateRequest

func (d PaymentStatusEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentStatusEvents := range d {
		err = validator.Struct(paymentStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentStatusEventsListCreateRequest) ToModelList() []model.PaymentStatusEvents {
	out := make([]model.PaymentStatusEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentStatusEventsUpdateRequest struct {
	PaymentIntentId        uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId       uuid.UUID       `json:"paymentAttemptId"`
	ProviderWebhookEventId uuid.UUID       `json:"providerWebhookEventId"`
	SourceType             string          `json:"sourceType"`
	EventType              string          `json:"eventType"`
	OldIntentStatus        string          `json:"oldIntentStatus"`
	NewIntentStatus        string          `json:"newIntentStatus"`
	OldAttemptStatus       string          `json:"oldAttemptStatus"`
	NewAttemptStatus       string          `json:"newAttemptStatus"`
	ProviderStatus         string          `json:"providerStatus"`
	Reason                 string          `json:"reason"`
	OccurredAt             time.Time       `json:"occurredAt"`
	Metadata               json.RawMessage `json:"metadata"`
}

func (d *PaymentStatusEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentStatusEventsUpdateRequest) ToModel() model.PaymentStatusEvents {
	return model.PaymentStatusEvents{
		PaymentIntentId:        nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:       nuuid.From(d.PaymentAttemptId),
		ProviderWebhookEventId: nuuid.From(d.ProviderWebhookEventId),
		SourceType:             d.SourceType,
		EventType:              d.EventType,
		OldIntentStatus:        null.StringFrom(d.OldIntentStatus),
		NewIntentStatus:        null.StringFrom(d.NewIntentStatus),
		OldAttemptStatus:       null.StringFrom(d.OldAttemptStatus),
		NewAttemptStatus:       null.StringFrom(d.NewAttemptStatus),
		ProviderStatus:         null.StringFrom(d.ProviderStatus),
		Reason:                 null.StringFrom(d.Reason),
		OccurredAt:             d.OccurredAt,
		Metadata:               d.Metadata,
	}
}

type PaymentStatusEventsBulkUpdateRequest struct {
	Id                     uuid.UUID       `json:"id"`
	PaymentIntentId        uuid.UUID       `json:"paymentIntentId"`
	PaymentAttemptId       uuid.UUID       `json:"paymentAttemptId"`
	ProviderWebhookEventId uuid.UUID       `json:"providerWebhookEventId"`
	SourceType             string          `json:"sourceType"`
	EventType              string          `json:"eventType"`
	OldIntentStatus        string          `json:"oldIntentStatus"`
	NewIntentStatus        string          `json:"newIntentStatus"`
	OldAttemptStatus       string          `json:"oldAttemptStatus"`
	NewAttemptStatus       string          `json:"newAttemptStatus"`
	ProviderStatus         string          `json:"providerStatus"`
	Reason                 string          `json:"reason"`
	OccurredAt             time.Time       `json:"occurredAt"`
	Metadata               json.RawMessage `json:"metadata"`
}

func (d PaymentStatusEventsBulkUpdateRequest) PrimaryID() PaymentStatusEventsPrimaryID {
	return PaymentStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type PaymentStatusEventsListBulkUpdateRequest []*PaymentStatusEventsBulkUpdateRequest

func (d PaymentStatusEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentStatusEvents := range d {
		err = validator.Struct(paymentStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentStatusEventsBulkUpdateRequest) ToModel() model.PaymentStatusEvents {
	return model.PaymentStatusEvents{
		Id:                     d.Id,
		PaymentIntentId:        nuuid.From(d.PaymentIntentId),
		PaymentAttemptId:       nuuid.From(d.PaymentAttemptId),
		ProviderWebhookEventId: nuuid.From(d.ProviderWebhookEventId),
		SourceType:             d.SourceType,
		EventType:              d.EventType,
		OldIntentStatus:        null.StringFrom(d.OldIntentStatus),
		NewIntentStatus:        null.StringFrom(d.NewIntentStatus),
		OldAttemptStatus:       null.StringFrom(d.OldAttemptStatus),
		NewAttemptStatus:       null.StringFrom(d.NewAttemptStatus),
		ProviderStatus:         null.StringFrom(d.ProviderStatus),
		Reason:                 null.StringFrom(d.Reason),
		OccurredAt:             d.OccurredAt,
		Metadata:               d.Metadata,
	}
}

type PaymentStatusEventsResponse struct {
	Id                     uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId        uuid.UUID       `json:"paymentIntentId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId       uuid.UUID       `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderWebhookEventId uuid.UUID       `json:"providerWebhookEventId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType             string          `json:"sourceType" validate:"required"`
	EventType              string          `json:"eventType" validate:"required"`
	OldIntentStatus        string          `json:"oldIntentStatus"`
	NewIntentStatus        string          `json:"newIntentStatus"`
	OldAttemptStatus       string          `json:"oldAttemptStatus"`
	NewAttemptStatus       string          `json:"newAttemptStatus"`
	ProviderStatus         string          `json:"providerStatus"`
	Reason                 string          `json:"reason"`
	OccurredAt             time.Time       `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata               json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewPaymentStatusEventsResponse(paymentStatusEvents model.PaymentStatusEvents) PaymentStatusEventsResponse {
	return PaymentStatusEventsResponse{
		Id:                     paymentStatusEvents.Id,
		PaymentIntentId:        paymentStatusEvents.PaymentIntentId.UUID,
		PaymentAttemptId:       paymentStatusEvents.PaymentAttemptId.UUID,
		ProviderWebhookEventId: paymentStatusEvents.ProviderWebhookEventId.UUID,
		SourceType:             paymentStatusEvents.SourceType,
		EventType:              paymentStatusEvents.EventType,
		OldIntentStatus:        paymentStatusEvents.OldIntentStatus.String,
		NewIntentStatus:        paymentStatusEvents.NewIntentStatus.String,
		OldAttemptStatus:       paymentStatusEvents.OldAttemptStatus.String,
		NewAttemptStatus:       paymentStatusEvents.NewAttemptStatus.String,
		ProviderStatus:         paymentStatusEvents.ProviderStatus.String,
		Reason:                 paymentStatusEvents.Reason.String,
		OccurredAt:             paymentStatusEvents.OccurredAt,
		Metadata:               paymentStatusEvents.Metadata,
	}
}

type PaymentStatusEventsListResponse []*PaymentStatusEventsResponse

func NewPaymentStatusEventsListResponse(paymentStatusEventsList model.PaymentStatusEventsList) PaymentStatusEventsListResponse {
	dtoPaymentStatusEventsListResponse := PaymentStatusEventsListResponse{}
	for _, paymentStatusEvents := range paymentStatusEventsList {
		dtoPaymentStatusEventsResponse := NewPaymentStatusEventsResponse(*paymentStatusEvents)
		dtoPaymentStatusEventsListResponse = append(dtoPaymentStatusEventsListResponse, &dtoPaymentStatusEventsResponse)
	}
	return dtoPaymentStatusEventsListResponse
}

type PaymentStatusEventsPrimaryIDList []PaymentStatusEventsPrimaryID

func (d PaymentStatusEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentStatusEvents := range d {
		err = validator.Struct(paymentStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentStatusEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentStatusEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentStatusEventsPrimaryID) ToModel() model.PaymentStatusEventsPrimaryID {
	return model.PaymentStatusEventsPrimaryID{
		Id: d.Id,
	}
}

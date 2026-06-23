package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PayoutStatusEventsDTOFieldNameType string

type payoutStatusEventsDTOFieldName struct {
	Id            PayoutStatusEventsDTOFieldNameType
	PayoutId      PayoutStatusEventsDTOFieldNameType
	FromStatus    PayoutStatusEventsDTOFieldNameType
	ToStatus      PayoutStatusEventsDTOFieldNameType
	ReasonCode    PayoutStatusEventsDTOFieldNameType
	ProviderRef   PayoutStatusEventsDTOFieldNameType
	EventPayload  PayoutStatusEventsDTOFieldNameType
	OccurredAt    PayoutStatusEventsDTOFieldNameType
	MetaCreatedAt PayoutStatusEventsDTOFieldNameType
	MetaCreatedBy PayoutStatusEventsDTOFieldNameType
	MetaUpdatedAt PayoutStatusEventsDTOFieldNameType
	MetaUpdatedBy PayoutStatusEventsDTOFieldNameType
	MetaDeletedAt PayoutStatusEventsDTOFieldNameType
	MetaDeletedBy PayoutStatusEventsDTOFieldNameType
}

var PayoutStatusEventsDTOFieldName = payoutStatusEventsDTOFieldName{
	Id:            "id",
	PayoutId:      "payoutId",
	FromStatus:    "fromStatus",
	ToStatus:      "toStatus",
	ReasonCode:    "reasonCode",
	ProviderRef:   "providerRef",
	EventPayload:  "eventPayload",
	OccurredAt:    "occurredAt",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformPayoutStatusEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutStatusEventsDTOFieldName.Id):
		return string(model.PayoutStatusEventsDBFieldName.Id), true

	case string(PayoutStatusEventsDTOFieldName.PayoutId):
		return string(model.PayoutStatusEventsDBFieldName.PayoutId), true

	case string(PayoutStatusEventsDTOFieldName.FromStatus):
		return string(model.PayoutStatusEventsDBFieldName.FromStatus), true

	case string(PayoutStatusEventsDTOFieldName.ToStatus):
		return string(model.PayoutStatusEventsDBFieldName.ToStatus), true

	case string(PayoutStatusEventsDTOFieldName.ReasonCode):
		return string(model.PayoutStatusEventsDBFieldName.ReasonCode), true

	case string(PayoutStatusEventsDTOFieldName.ProviderRef):
		return string(model.PayoutStatusEventsDBFieldName.ProviderRef), true

	case string(PayoutStatusEventsDTOFieldName.EventPayload):
		return string(model.PayoutStatusEventsDBFieldName.EventPayload), true

	case string(PayoutStatusEventsDTOFieldName.OccurredAt):
		return string(model.PayoutStatusEventsDBFieldName.OccurredAt), true

	case string(PayoutStatusEventsDTOFieldName.MetaCreatedAt):
		return string(model.PayoutStatusEventsDBFieldName.MetaCreatedAt), true

	case string(PayoutStatusEventsDTOFieldName.MetaCreatedBy):
		return string(model.PayoutStatusEventsDBFieldName.MetaCreatedBy), true

	case string(PayoutStatusEventsDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutStatusEventsDBFieldName.MetaUpdatedAt), true

	case string(PayoutStatusEventsDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutStatusEventsDBFieldName.MetaUpdatedBy), true

	case string(PayoutStatusEventsDTOFieldName.MetaDeletedAt):
		return string(model.PayoutStatusEventsDBFieldName.MetaDeletedAt), true

	case string(PayoutStatusEventsDTOFieldName.MetaDeletedBy):
		return string(model.PayoutStatusEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutStatusEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutStatusEventsBaseFilterField(field string) bool {
	spec, found := model.NewPayoutStatusEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutStatusEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutStatusEventsProjectionOutputPath(path string) error {
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

func transformPayoutStatusEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutStatusEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutStatusEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutStatusEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutStatusEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutStatusEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutStatusEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutStatusEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutStatusEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutStatusEventsFilter(filter *model.Filter) {
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
			Field: string(PayoutStatusEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutStatusEventsSelectableResponse map[string]interface{}
type PayoutStatusEventsSelectableListResponse []*PayoutStatusEventsSelectableResponse

func assignPayoutStatusEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutStatusEventsSelectableValue(out PayoutStatusEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutStatusEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutStatusEventsSelectableResponse(payoutStatusEvents model.PayoutStatusEvents, filter model.Filter) PayoutStatusEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutStatusEventsDBFieldName.Id),
			string(model.PayoutStatusEventsDBFieldName.PayoutId),
			string(model.PayoutStatusEventsDBFieldName.FromStatus),
			string(model.PayoutStatusEventsDBFieldName.ToStatus),
			string(model.PayoutStatusEventsDBFieldName.ReasonCode),
			string(model.PayoutStatusEventsDBFieldName.ProviderRef),
			string(model.PayoutStatusEventsDBFieldName.EventPayload),
			string(model.PayoutStatusEventsDBFieldName.OccurredAt),
			string(model.PayoutStatusEventsDBFieldName.MetaCreatedAt),
			string(model.PayoutStatusEventsDBFieldName.MetaCreatedBy),
			string(model.PayoutStatusEventsDBFieldName.MetaUpdatedAt),
			string(model.PayoutStatusEventsDBFieldName.MetaUpdatedBy),
			string(model.PayoutStatusEventsDBFieldName.MetaDeletedAt),
			string(model.PayoutStatusEventsDBFieldName.MetaDeletedBy),
		)
	}
	payoutStatusEventsSelectableResponse := PayoutStatusEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutStatusEventsDBFieldName.Id):
			key := string(PayoutStatusEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.Id, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.PayoutId):
			key := string(PayoutStatusEventsDTOFieldName.PayoutId)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.PayoutId, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.FromStatus):
			key := string(PayoutStatusEventsDTOFieldName.FromStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.FromStatus.String, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.ToStatus):
			key := string(PayoutStatusEventsDTOFieldName.ToStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.ToStatus, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.ReasonCode):
			key := string(PayoutStatusEventsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.ReasonCode.String, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.ProviderRef):
			key := string(PayoutStatusEventsDTOFieldName.ProviderRef)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.ProviderRef.String, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.EventPayload):
			key := string(PayoutStatusEventsDTOFieldName.EventPayload)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.EventPayload, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.OccurredAt):
			key := string(PayoutStatusEventsDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.OccurredAt, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaCreatedAt):
			key := string(PayoutStatusEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaCreatedAt, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaCreatedBy):
			key := string(PayoutStatusEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaCreatedBy, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaUpdatedAt):
			key := string(PayoutStatusEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaUpdatedBy):
			key := string(PayoutStatusEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaDeletedAt):
			key := string(PayoutStatusEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutStatusEventsDBFieldName.MetaDeletedBy):
			key := string(PayoutStatusEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutStatusEventsSelectableValue(payoutStatusEventsSelectableResponse, key, payoutStatusEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutStatusEventsSelectableResponse
}

func NewPayoutStatusEventsListResponseFromFilterResult(result []model.PayoutStatusEventsFilterResult, filter model.Filter) PayoutStatusEventsSelectableListResponse {
	dtoPayoutStatusEventsListResponse := PayoutStatusEventsSelectableListResponse{}
	for _, row := range result {
		dtoPayoutStatusEventsResponse := NewPayoutStatusEventsSelectableResponse(row.PayoutStatusEvents, filter)
		dtoPayoutStatusEventsListResponse = append(dtoPayoutStatusEventsListResponse, &dtoPayoutStatusEventsResponse)
	}
	return dtoPayoutStatusEventsListResponse
}

type PayoutStatusEventsFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     PayoutStatusEventsSelectableListResponse `json:"data"`
}

func reversePayoutStatusEventsFilterResults(result []model.PayoutStatusEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutStatusEventsFilterResponse(result []model.PayoutStatusEventsFilterResult, filter model.Filter) (resp PayoutStatusEventsFilterResponse) {
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
			reversePayoutStatusEventsFilterResults(dataResult)
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

	resp.Data = NewPayoutStatusEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutStatusEventsCreateRequest struct {
	PayoutId     uuid.UUID       `json:"payoutId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d *PayoutStatusEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutStatusEventsCreateRequest) ToModel() model.PayoutStatusEvents {
	id, _ := uuid.NewV4()
	return model.PayoutStatusEvents{
		Id:           id,
		PayoutId:     d.PayoutId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type PayoutStatusEventsListCreateRequest []*PayoutStatusEventsCreateRequest

func (d PayoutStatusEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutStatusEvents := range d {
		err = validator.Struct(payoutStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutStatusEventsListCreateRequest) ToModelList() []model.PayoutStatusEvents {
	out := make([]model.PayoutStatusEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutStatusEventsUpdateRequest struct {
	PayoutId     uuid.UUID       `json:"payoutId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d *PayoutStatusEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutStatusEventsUpdateRequest) ToModel() model.PayoutStatusEvents {
	return model.PayoutStatusEvents{
		PayoutId:     d.PayoutId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type PayoutStatusEventsBulkUpdateRequest struct {
	Id           uuid.UUID       `json:"id"`
	PayoutId     uuid.UUID       `json:"payoutId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d PayoutStatusEventsBulkUpdateRequest) PrimaryID() PayoutStatusEventsPrimaryID {
	return PayoutStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type PayoutStatusEventsListBulkUpdateRequest []*PayoutStatusEventsBulkUpdateRequest

func (d PayoutStatusEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutStatusEvents := range d {
		err = validator.Struct(payoutStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutStatusEventsBulkUpdateRequest) ToModel() model.PayoutStatusEvents {
	return model.PayoutStatusEvents{
		Id:           d.Id,
		PayoutId:     d.PayoutId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type PayoutStatusEventsResponse struct {
	Id           uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutId     uuid.UUID       `json:"payoutId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus" validate:"required"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload" swaggertype:"object"`
	OccurredAt   time.Time       `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

func NewPayoutStatusEventsResponse(payoutStatusEvents model.PayoutStatusEvents) PayoutStatusEventsResponse {
	return PayoutStatusEventsResponse{
		Id:           payoutStatusEvents.Id,
		PayoutId:     payoutStatusEvents.PayoutId,
		FromStatus:   payoutStatusEvents.FromStatus.String,
		ToStatus:     payoutStatusEvents.ToStatus,
		ReasonCode:   payoutStatusEvents.ReasonCode.String,
		ProviderRef:  payoutStatusEvents.ProviderRef.String,
		EventPayload: payoutStatusEvents.EventPayload,
		OccurredAt:   payoutStatusEvents.OccurredAt,
	}
}

type PayoutStatusEventsListResponse []*PayoutStatusEventsResponse

func NewPayoutStatusEventsListResponse(payoutStatusEventsList model.PayoutStatusEventsList) PayoutStatusEventsListResponse {
	dtoPayoutStatusEventsListResponse := PayoutStatusEventsListResponse{}
	for _, payoutStatusEvents := range payoutStatusEventsList {
		dtoPayoutStatusEventsResponse := NewPayoutStatusEventsResponse(*payoutStatusEvents)
		dtoPayoutStatusEventsListResponse = append(dtoPayoutStatusEventsListResponse, &dtoPayoutStatusEventsResponse)
	}
	return dtoPayoutStatusEventsListResponse
}

type PayoutStatusEventsPrimaryIDList []PayoutStatusEventsPrimaryID

func (d PayoutStatusEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutStatusEvents := range d {
		err = validator.Struct(payoutStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutStatusEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutStatusEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutStatusEventsPrimaryID) ToModel() model.PayoutStatusEventsPrimaryID {
	return model.PayoutStatusEventsPrimaryID{
		Id: d.Id,
	}
}

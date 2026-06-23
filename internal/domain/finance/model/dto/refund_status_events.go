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

type RefundStatusEventsDTOFieldNameType string

type refundStatusEventsDTOFieldName struct {
	Id            RefundStatusEventsDTOFieldNameType
	RefundId      RefundStatusEventsDTOFieldNameType
	FromStatus    RefundStatusEventsDTOFieldNameType
	ToStatus      RefundStatusEventsDTOFieldNameType
	ReasonCode    RefundStatusEventsDTOFieldNameType
	ProviderRef   RefundStatusEventsDTOFieldNameType
	EventPayload  RefundStatusEventsDTOFieldNameType
	OccurredAt    RefundStatusEventsDTOFieldNameType
	MetaCreatedAt RefundStatusEventsDTOFieldNameType
	MetaCreatedBy RefundStatusEventsDTOFieldNameType
	MetaUpdatedAt RefundStatusEventsDTOFieldNameType
	MetaUpdatedBy RefundStatusEventsDTOFieldNameType
	MetaDeletedAt RefundStatusEventsDTOFieldNameType
	MetaDeletedBy RefundStatusEventsDTOFieldNameType
}

var RefundStatusEventsDTOFieldName = refundStatusEventsDTOFieldName{
	Id:            "id",
	RefundId:      "refundId",
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

func transformRefundStatusEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundStatusEventsDTOFieldName.Id):
		return string(model.RefundStatusEventsDBFieldName.Id), true

	case string(RefundStatusEventsDTOFieldName.RefundId):
		return string(model.RefundStatusEventsDBFieldName.RefundId), true

	case string(RefundStatusEventsDTOFieldName.FromStatus):
		return string(model.RefundStatusEventsDBFieldName.FromStatus), true

	case string(RefundStatusEventsDTOFieldName.ToStatus):
		return string(model.RefundStatusEventsDBFieldName.ToStatus), true

	case string(RefundStatusEventsDTOFieldName.ReasonCode):
		return string(model.RefundStatusEventsDBFieldName.ReasonCode), true

	case string(RefundStatusEventsDTOFieldName.ProviderRef):
		return string(model.RefundStatusEventsDBFieldName.ProviderRef), true

	case string(RefundStatusEventsDTOFieldName.EventPayload):
		return string(model.RefundStatusEventsDBFieldName.EventPayload), true

	case string(RefundStatusEventsDTOFieldName.OccurredAt):
		return string(model.RefundStatusEventsDBFieldName.OccurredAt), true

	case string(RefundStatusEventsDTOFieldName.MetaCreatedAt):
		return string(model.RefundStatusEventsDBFieldName.MetaCreatedAt), true

	case string(RefundStatusEventsDTOFieldName.MetaCreatedBy):
		return string(model.RefundStatusEventsDBFieldName.MetaCreatedBy), true

	case string(RefundStatusEventsDTOFieldName.MetaUpdatedAt):
		return string(model.RefundStatusEventsDBFieldName.MetaUpdatedAt), true

	case string(RefundStatusEventsDTOFieldName.MetaUpdatedBy):
		return string(model.RefundStatusEventsDBFieldName.MetaUpdatedBy), true

	case string(RefundStatusEventsDTOFieldName.MetaDeletedAt):
		return string(model.RefundStatusEventsDBFieldName.MetaDeletedAt), true

	case string(RefundStatusEventsDTOFieldName.MetaDeletedBy):
		return string(model.RefundStatusEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundStatusEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundStatusEventsBaseFilterField(field string) bool {
	spec, found := model.NewRefundStatusEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundStatusEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundStatusEventsProjectionOutputPath(path string) error {
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

func transformRefundStatusEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundStatusEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundStatusEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundStatusEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundStatusEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundStatusEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundStatusEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundStatusEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundStatusEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundStatusEventsFilter(filter *model.Filter) {
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
			Field: string(RefundStatusEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundStatusEventsSelectableResponse map[string]interface{}
type RefundStatusEventsSelectableListResponse []*RefundStatusEventsSelectableResponse

func assignRefundStatusEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundStatusEventsSelectableValue(out RefundStatusEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundStatusEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundStatusEventsSelectableResponse(refundStatusEvents model.RefundStatusEvents, filter model.Filter) RefundStatusEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundStatusEventsDBFieldName.Id),
			string(model.RefundStatusEventsDBFieldName.RefundId),
			string(model.RefundStatusEventsDBFieldName.FromStatus),
			string(model.RefundStatusEventsDBFieldName.ToStatus),
			string(model.RefundStatusEventsDBFieldName.ReasonCode),
			string(model.RefundStatusEventsDBFieldName.ProviderRef),
			string(model.RefundStatusEventsDBFieldName.EventPayload),
			string(model.RefundStatusEventsDBFieldName.OccurredAt),
			string(model.RefundStatusEventsDBFieldName.MetaCreatedAt),
			string(model.RefundStatusEventsDBFieldName.MetaCreatedBy),
			string(model.RefundStatusEventsDBFieldName.MetaUpdatedAt),
			string(model.RefundStatusEventsDBFieldName.MetaUpdatedBy),
			string(model.RefundStatusEventsDBFieldName.MetaDeletedAt),
			string(model.RefundStatusEventsDBFieldName.MetaDeletedBy),
		)
	}
	refundStatusEventsSelectableResponse := RefundStatusEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundStatusEventsDBFieldName.Id):
			key := string(RefundStatusEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.Id, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.RefundId):
			key := string(RefundStatusEventsDTOFieldName.RefundId)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.RefundId, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.FromStatus):
			key := string(RefundStatusEventsDTOFieldName.FromStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.FromStatus.String, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.ToStatus):
			key := string(RefundStatusEventsDTOFieldName.ToStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.ToStatus, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.ReasonCode):
			key := string(RefundStatusEventsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.ReasonCode.String, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.ProviderRef):
			key := string(RefundStatusEventsDTOFieldName.ProviderRef)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.ProviderRef.String, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.EventPayload):
			key := string(RefundStatusEventsDTOFieldName.EventPayload)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.EventPayload, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.OccurredAt):
			key := string(RefundStatusEventsDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.OccurredAt, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaCreatedAt):
			key := string(RefundStatusEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaCreatedAt, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaCreatedBy):
			key := string(RefundStatusEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaCreatedBy, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaUpdatedAt):
			key := string(RefundStatusEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaUpdatedAt, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaUpdatedBy):
			key := string(RefundStatusEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaUpdatedBy, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaDeletedAt):
			key := string(RefundStatusEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundStatusEventsDBFieldName.MetaDeletedBy):
			key := string(RefundStatusEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundStatusEventsSelectableValue(refundStatusEventsSelectableResponse, key, refundStatusEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return refundStatusEventsSelectableResponse
}

func NewRefundStatusEventsListResponseFromFilterResult(result []model.RefundStatusEventsFilterResult, filter model.Filter) RefundStatusEventsSelectableListResponse {
	dtoRefundStatusEventsListResponse := RefundStatusEventsSelectableListResponse{}
	for _, row := range result {
		dtoRefundStatusEventsResponse := NewRefundStatusEventsSelectableResponse(row.RefundStatusEvents, filter)
		dtoRefundStatusEventsListResponse = append(dtoRefundStatusEventsListResponse, &dtoRefundStatusEventsResponse)
	}
	return dtoRefundStatusEventsListResponse
}

type RefundStatusEventsFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     RefundStatusEventsSelectableListResponse `json:"data"`
}

func reverseRefundStatusEventsFilterResults(result []model.RefundStatusEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundStatusEventsFilterResponse(result []model.RefundStatusEventsFilterResult, filter model.Filter) (resp RefundStatusEventsFilterResponse) {
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
			reverseRefundStatusEventsFilterResults(dataResult)
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

	resp.Data = NewRefundStatusEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundStatusEventsCreateRequest struct {
	RefundId     uuid.UUID       `json:"refundId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d *RefundStatusEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundStatusEventsCreateRequest) ToModel() model.RefundStatusEvents {
	id, _ := uuid.NewV4()
	return model.RefundStatusEvents{
		Id:           id,
		RefundId:     d.RefundId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type RefundStatusEventsListCreateRequest []*RefundStatusEventsCreateRequest

func (d RefundStatusEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundStatusEvents := range d {
		err = validator.Struct(refundStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundStatusEventsListCreateRequest) ToModelList() []model.RefundStatusEvents {
	out := make([]model.RefundStatusEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundStatusEventsUpdateRequest struct {
	RefundId     uuid.UUID       `json:"refundId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d *RefundStatusEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundStatusEventsUpdateRequest) ToModel() model.RefundStatusEvents {
	return model.RefundStatusEvents{
		RefundId:     d.RefundId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type RefundStatusEventsBulkUpdateRequest struct {
	Id           uuid.UUID       `json:"id"`
	RefundId     uuid.UUID       `json:"refundId"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload"`
	OccurredAt   time.Time       `json:"occurredAt"`
}

func (d RefundStatusEventsBulkUpdateRequest) PrimaryID() RefundStatusEventsPrimaryID {
	return RefundStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type RefundStatusEventsListBulkUpdateRequest []*RefundStatusEventsBulkUpdateRequest

func (d RefundStatusEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundStatusEvents := range d {
		err = validator.Struct(refundStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundStatusEventsBulkUpdateRequest) ToModel() model.RefundStatusEvents {
	return model.RefundStatusEvents{
		Id:           d.Id,
		RefundId:     d.RefundId,
		FromStatus:   null.StringFrom(d.FromStatus),
		ToStatus:     d.ToStatus,
		ReasonCode:   null.StringFrom(d.ReasonCode),
		ProviderRef:  null.StringFrom(d.ProviderRef),
		EventPayload: d.EventPayload,
		OccurredAt:   d.OccurredAt,
	}
}

type RefundStatusEventsResponse struct {
	Id           uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundId     uuid.UUID       `json:"refundId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FromStatus   string          `json:"fromStatus"`
	ToStatus     string          `json:"toStatus" validate:"required"`
	ReasonCode   string          `json:"reasonCode"`
	ProviderRef  string          `json:"providerRef"`
	EventPayload json.RawMessage `json:"eventPayload" swaggertype:"object"`
	OccurredAt   time.Time       `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

func NewRefundStatusEventsResponse(refundStatusEvents model.RefundStatusEvents) RefundStatusEventsResponse {
	return RefundStatusEventsResponse{
		Id:           refundStatusEvents.Id,
		RefundId:     refundStatusEvents.RefundId,
		FromStatus:   refundStatusEvents.FromStatus.String,
		ToStatus:     refundStatusEvents.ToStatus,
		ReasonCode:   refundStatusEvents.ReasonCode.String,
		ProviderRef:  refundStatusEvents.ProviderRef.String,
		EventPayload: refundStatusEvents.EventPayload,
		OccurredAt:   refundStatusEvents.OccurredAt,
	}
}

type RefundStatusEventsListResponse []*RefundStatusEventsResponse

func NewRefundStatusEventsListResponse(refundStatusEventsList model.RefundStatusEventsList) RefundStatusEventsListResponse {
	dtoRefundStatusEventsListResponse := RefundStatusEventsListResponse{}
	for _, refundStatusEvents := range refundStatusEventsList {
		dtoRefundStatusEventsResponse := NewRefundStatusEventsResponse(*refundStatusEvents)
		dtoRefundStatusEventsListResponse = append(dtoRefundStatusEventsListResponse, &dtoRefundStatusEventsResponse)
	}
	return dtoRefundStatusEventsListResponse
}

type RefundStatusEventsPrimaryIDList []RefundStatusEventsPrimaryID

func (d RefundStatusEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundStatusEvents := range d {
		err = validator.Struct(refundStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundStatusEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundStatusEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundStatusEventsPrimaryID) ToModel() model.RefundStatusEventsPrimaryID {
	return model.RefundStatusEventsPrimaryID{
		Id: d.Id,
	}
}

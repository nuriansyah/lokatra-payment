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

type DisputeStatusEventsDTOFieldNameType string

type disputeStatusEventsDTOFieldName struct {
	Id             DisputeStatusEventsDTOFieldNameType
	DisputeId      DisputeStatusEventsDTOFieldNameType
	PreviousStatus DisputeStatusEventsDTOFieldNameType
	NextStatus     DisputeStatusEventsDTOFieldNameType
	ReasonCode     DisputeStatusEventsDTOFieldNameType
	ActorId        DisputeStatusEventsDTOFieldNameType
	OccurredAt     DisputeStatusEventsDTOFieldNameType
	Metadata       DisputeStatusEventsDTOFieldNameType
	MetaCreatedAt  DisputeStatusEventsDTOFieldNameType
	MetaCreatedBy  DisputeStatusEventsDTOFieldNameType
	MetaUpdatedAt  DisputeStatusEventsDTOFieldNameType
	MetaUpdatedBy  DisputeStatusEventsDTOFieldNameType
	MetaDeletedAt  DisputeStatusEventsDTOFieldNameType
	MetaDeletedBy  DisputeStatusEventsDTOFieldNameType
}

var DisputeStatusEventsDTOFieldName = disputeStatusEventsDTOFieldName{
	Id:             "id",
	DisputeId:      "disputeId",
	PreviousStatus: "previousStatus",
	NextStatus:     "nextStatus",
	ReasonCode:     "reasonCode",
	ActorId:        "actorId",
	OccurredAt:     "occurredAt",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformDisputeStatusEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(DisputeStatusEventsDTOFieldName.Id):
		return string(model.DisputeStatusEventsDBFieldName.Id), true

	case string(DisputeStatusEventsDTOFieldName.DisputeId):
		return string(model.DisputeStatusEventsDBFieldName.DisputeId), true

	case string(DisputeStatusEventsDTOFieldName.PreviousStatus):
		return string(model.DisputeStatusEventsDBFieldName.PreviousStatus), true

	case string(DisputeStatusEventsDTOFieldName.NextStatus):
		return string(model.DisputeStatusEventsDBFieldName.NextStatus), true

	case string(DisputeStatusEventsDTOFieldName.ReasonCode):
		return string(model.DisputeStatusEventsDBFieldName.ReasonCode), true

	case string(DisputeStatusEventsDTOFieldName.ActorId):
		return string(model.DisputeStatusEventsDBFieldName.ActorId), true

	case string(DisputeStatusEventsDTOFieldName.OccurredAt):
		return string(model.DisputeStatusEventsDBFieldName.OccurredAt), true

	case string(DisputeStatusEventsDTOFieldName.Metadata):
		return string(model.DisputeStatusEventsDBFieldName.Metadata), true

	case string(DisputeStatusEventsDTOFieldName.MetaCreatedAt):
		return string(model.DisputeStatusEventsDBFieldName.MetaCreatedAt), true

	case string(DisputeStatusEventsDTOFieldName.MetaCreatedBy):
		return string(model.DisputeStatusEventsDBFieldName.MetaCreatedBy), true

	case string(DisputeStatusEventsDTOFieldName.MetaUpdatedAt):
		return string(model.DisputeStatusEventsDBFieldName.MetaUpdatedAt), true

	case string(DisputeStatusEventsDTOFieldName.MetaUpdatedBy):
		return string(model.DisputeStatusEventsDBFieldName.MetaUpdatedBy), true

	case string(DisputeStatusEventsDTOFieldName.MetaDeletedAt):
		return string(model.DisputeStatusEventsDBFieldName.MetaDeletedAt), true

	case string(DisputeStatusEventsDTOFieldName.MetaDeletedBy):
		return string(model.DisputeStatusEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewDisputeStatusEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isDisputeStatusEventsBaseFilterField(field string) bool {
	spec, found := model.NewDisputeStatusEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeDisputeStatusEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateDisputeStatusEventsProjectionOutputPath(path string) error {
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

func transformDisputeStatusEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformDisputeStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformDisputeStatusEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformDisputeStatusEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformDisputeStatusEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isDisputeStatusEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateDisputeStatusEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeDisputeStatusEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformDisputeStatusEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformDisputeStatusEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformDisputeStatusEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultDisputeStatusEventsFilter(filter *model.Filter) {
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
			Field: string(DisputeStatusEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type DisputeStatusEventsSelectableResponse map[string]interface{}
type DisputeStatusEventsSelectableListResponse []*DisputeStatusEventsSelectableResponse

func assignDisputeStatusEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setDisputeStatusEventsSelectableValue(out DisputeStatusEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignDisputeStatusEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewDisputeStatusEventsSelectableResponse(disputeStatusEvents model.DisputeStatusEvents, filter model.Filter) DisputeStatusEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.DisputeStatusEventsDBFieldName.Id),
			string(model.DisputeStatusEventsDBFieldName.DisputeId),
			string(model.DisputeStatusEventsDBFieldName.PreviousStatus),
			string(model.DisputeStatusEventsDBFieldName.NextStatus),
			string(model.DisputeStatusEventsDBFieldName.ReasonCode),
			string(model.DisputeStatusEventsDBFieldName.ActorId),
			string(model.DisputeStatusEventsDBFieldName.OccurredAt),
			string(model.DisputeStatusEventsDBFieldName.Metadata),
			string(model.DisputeStatusEventsDBFieldName.MetaCreatedAt),
			string(model.DisputeStatusEventsDBFieldName.MetaCreatedBy),
			string(model.DisputeStatusEventsDBFieldName.MetaUpdatedAt),
			string(model.DisputeStatusEventsDBFieldName.MetaUpdatedBy),
			string(model.DisputeStatusEventsDBFieldName.MetaDeletedAt),
			string(model.DisputeStatusEventsDBFieldName.MetaDeletedBy),
		)
	}
	disputeStatusEventsSelectableResponse := DisputeStatusEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.DisputeStatusEventsDBFieldName.Id):
			key := string(DisputeStatusEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.Id, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.DisputeId):
			key := string(DisputeStatusEventsDTOFieldName.DisputeId)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.DisputeId, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.PreviousStatus):
			key := string(DisputeStatusEventsDTOFieldName.PreviousStatus)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.PreviousStatus.String, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.NextStatus):
			key := string(DisputeStatusEventsDTOFieldName.NextStatus)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.NextStatus, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.ReasonCode):
			key := string(DisputeStatusEventsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.ReasonCode, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.ActorId):
			key := string(DisputeStatusEventsDTOFieldName.ActorId)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.ActorId, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.OccurredAt):
			key := string(DisputeStatusEventsDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.OccurredAt, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.Metadata):
			key := string(DisputeStatusEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.Metadata, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaCreatedAt):
			key := string(DisputeStatusEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaCreatedAt, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaCreatedBy):
			key := string(DisputeStatusEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaCreatedBy, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaUpdatedAt):
			key := string(DisputeStatusEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaUpdatedAt, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaUpdatedBy):
			key := string(DisputeStatusEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaUpdatedBy, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaDeletedAt):
			key := string(DisputeStatusEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.DisputeStatusEventsDBFieldName.MetaDeletedBy):
			key := string(DisputeStatusEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeStatusEventsSelectableValue(disputeStatusEventsSelectableResponse, key, disputeStatusEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return disputeStatusEventsSelectableResponse
}

func NewDisputeStatusEventsListResponseFromFilterResult(result []model.DisputeStatusEventsFilterResult, filter model.Filter) DisputeStatusEventsSelectableListResponse {
	dtoDisputeStatusEventsListResponse := DisputeStatusEventsSelectableListResponse{}
	for _, row := range result {
		dtoDisputeStatusEventsResponse := NewDisputeStatusEventsSelectableResponse(row.DisputeStatusEvents, filter)
		dtoDisputeStatusEventsListResponse = append(dtoDisputeStatusEventsListResponse, &dtoDisputeStatusEventsResponse)
	}
	return dtoDisputeStatusEventsListResponse
}

type DisputeStatusEventsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     DisputeStatusEventsSelectableListResponse `json:"data"`
}

func reverseDisputeStatusEventsFilterResults(result []model.DisputeStatusEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewDisputeStatusEventsFilterResponse(result []model.DisputeStatusEventsFilterResult, filter model.Filter) (resp DisputeStatusEventsFilterResponse) {
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
			reverseDisputeStatusEventsFilterResults(dataResult)
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

	resp.Data = NewDisputeStatusEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type DisputeStatusEventsCreateRequest struct {
	DisputeId      uuid.UUID       `json:"disputeId"`
	PreviousStatus string          `json:"previousStatus"`
	NextStatus     string          `json:"nextStatus"`
	ReasonCode     string          `json:"reasonCode"`
	ActorId        uuid.UUID       `json:"actorId"`
	OccurredAt     time.Time       `json:"occurredAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *DisputeStatusEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *DisputeStatusEventsCreateRequest) ToModel() model.DisputeStatusEvents {
	id, _ := uuid.NewV4()
	return model.DisputeStatusEvents{
		Id:             id,
		DisputeId:      d.DisputeId,
		PreviousStatus: null.StringFrom(d.PreviousStatus),
		NextStatus:     d.NextStatus,
		ReasonCode:     d.ReasonCode,
		ActorId:        d.ActorId,
		OccurredAt:     d.OccurredAt,
		Metadata:       d.Metadata,
	}
}

type DisputeStatusEventsListCreateRequest []*DisputeStatusEventsCreateRequest

func (d DisputeStatusEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeStatusEvents := range d {
		err = validator.Struct(disputeStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeStatusEventsListCreateRequest) ToModelList() []model.DisputeStatusEvents {
	out := make([]model.DisputeStatusEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type DisputeStatusEventsUpdateRequest struct {
	DisputeId      uuid.UUID       `json:"disputeId"`
	PreviousStatus string          `json:"previousStatus"`
	NextStatus     string          `json:"nextStatus"`
	ReasonCode     string          `json:"reasonCode"`
	ActorId        uuid.UUID       `json:"actorId"`
	OccurredAt     time.Time       `json:"occurredAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *DisputeStatusEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d DisputeStatusEventsUpdateRequest) ToModel() model.DisputeStatusEvents {
	return model.DisputeStatusEvents{
		DisputeId:      d.DisputeId,
		PreviousStatus: null.StringFrom(d.PreviousStatus),
		NextStatus:     d.NextStatus,
		ReasonCode:     d.ReasonCode,
		ActorId:        d.ActorId,
		OccurredAt:     d.OccurredAt,
		Metadata:       d.Metadata,
	}
}

type DisputeStatusEventsBulkUpdateRequest struct {
	Id             uuid.UUID       `json:"id"`
	DisputeId      uuid.UUID       `json:"disputeId"`
	PreviousStatus string          `json:"previousStatus"`
	NextStatus     string          `json:"nextStatus"`
	ReasonCode     string          `json:"reasonCode"`
	ActorId        uuid.UUID       `json:"actorId"`
	OccurredAt     time.Time       `json:"occurredAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d DisputeStatusEventsBulkUpdateRequest) PrimaryID() DisputeStatusEventsPrimaryID {
	return DisputeStatusEventsPrimaryID{
		Id: d.Id,
	}
}

type DisputeStatusEventsListBulkUpdateRequest []*DisputeStatusEventsBulkUpdateRequest

func (d DisputeStatusEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeStatusEvents := range d {
		err = validator.Struct(disputeStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeStatusEventsBulkUpdateRequest) ToModel() model.DisputeStatusEvents {
	return model.DisputeStatusEvents{
		Id:             d.Id,
		DisputeId:      d.DisputeId,
		PreviousStatus: null.StringFrom(d.PreviousStatus),
		NextStatus:     d.NextStatus,
		ReasonCode:     d.ReasonCode,
		ActorId:        d.ActorId,
		OccurredAt:     d.OccurredAt,
		Metadata:       d.Metadata,
	}
}

type DisputeStatusEventsResponse struct {
	Id             uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DisputeId      uuid.UUID       `json:"disputeId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PreviousStatus string          `json:"previousStatus"`
	NextStatus     string          `json:"nextStatus" validate:"required"`
	ReasonCode     string          `json:"reasonCode" validate:"required"`
	ActorId        uuid.UUID       `json:"actorId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OccurredAt     time.Time       `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewDisputeStatusEventsResponse(disputeStatusEvents model.DisputeStatusEvents) DisputeStatusEventsResponse {
	return DisputeStatusEventsResponse{
		Id:             disputeStatusEvents.Id,
		DisputeId:      disputeStatusEvents.DisputeId,
		PreviousStatus: disputeStatusEvents.PreviousStatus.String,
		NextStatus:     disputeStatusEvents.NextStatus,
		ReasonCode:     disputeStatusEvents.ReasonCode,
		ActorId:        disputeStatusEvents.ActorId,
		OccurredAt:     disputeStatusEvents.OccurredAt,
		Metadata:       disputeStatusEvents.Metadata,
	}
}

type DisputeStatusEventsListResponse []*DisputeStatusEventsResponse

func NewDisputeStatusEventsListResponse(disputeStatusEventsList model.DisputeStatusEventsList) DisputeStatusEventsListResponse {
	dtoDisputeStatusEventsListResponse := DisputeStatusEventsListResponse{}
	for _, disputeStatusEvents := range disputeStatusEventsList {
		dtoDisputeStatusEventsResponse := NewDisputeStatusEventsResponse(*disputeStatusEvents)
		dtoDisputeStatusEventsListResponse = append(dtoDisputeStatusEventsListResponse, &dtoDisputeStatusEventsResponse)
	}
	return dtoDisputeStatusEventsListResponse
}

type DisputeStatusEventsPrimaryIDList []DisputeStatusEventsPrimaryID

func (d DisputeStatusEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeStatusEvents := range d {
		err = validator.Struct(disputeStatusEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type DisputeStatusEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *DisputeStatusEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d DisputeStatusEventsPrimaryID) ToModel() model.DisputeStatusEventsPrimaryID {
	return model.DisputeStatusEventsPrimaryID{
		Id: d.Id,
	}
}

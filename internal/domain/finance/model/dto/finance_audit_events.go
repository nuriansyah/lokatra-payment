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

type FinanceAuditEventsDTOFieldNameType string

type financeAuditEventsDTOFieldName struct {
	Id            FinanceAuditEventsDTOFieldNameType
	AggregateType FinanceAuditEventsDTOFieldNameType
	AggregateId   FinanceAuditEventsDTOFieldNameType
	EventType     FinanceAuditEventsDTOFieldNameType
	ActorUserId   FinanceAuditEventsDTOFieldNameType
	ActorType     FinanceAuditEventsDTOFieldNameType
	EventAt       FinanceAuditEventsDTOFieldNameType
	OldState      FinanceAuditEventsDTOFieldNameType
	NewState      FinanceAuditEventsDTOFieldNameType
	CorrelationId FinanceAuditEventsDTOFieldNameType
	CausationId   FinanceAuditEventsDTOFieldNameType
	Metadata      FinanceAuditEventsDTOFieldNameType
	MetaCreatedAt FinanceAuditEventsDTOFieldNameType
	MetaCreatedBy FinanceAuditEventsDTOFieldNameType
	MetaUpdatedAt FinanceAuditEventsDTOFieldNameType
	MetaUpdatedBy FinanceAuditEventsDTOFieldNameType
	MetaDeletedAt FinanceAuditEventsDTOFieldNameType
	MetaDeletedBy FinanceAuditEventsDTOFieldNameType
}

var FinanceAuditEventsDTOFieldName = financeAuditEventsDTOFieldName{
	Id:            "id",
	AggregateType: "aggregateType",
	AggregateId:   "aggregateId",
	EventType:     "eventType",
	ActorUserId:   "actorUserId",
	ActorType:     "actorType",
	EventAt:       "eventAt",
	OldState:      "oldState",
	NewState:      "newState",
	CorrelationId: "correlationId",
	CausationId:   "causationId",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformFinanceAuditEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceAuditEventsDTOFieldName.Id):
		return string(model.FinanceAuditEventsDBFieldName.Id), true

	case string(FinanceAuditEventsDTOFieldName.AggregateType):
		return string(model.FinanceAuditEventsDBFieldName.AggregateType), true

	case string(FinanceAuditEventsDTOFieldName.AggregateId):
		return string(model.FinanceAuditEventsDBFieldName.AggregateId), true

	case string(FinanceAuditEventsDTOFieldName.EventType):
		return string(model.FinanceAuditEventsDBFieldName.EventType), true

	case string(FinanceAuditEventsDTOFieldName.ActorUserId):
		return string(model.FinanceAuditEventsDBFieldName.ActorUserId), true

	case string(FinanceAuditEventsDTOFieldName.ActorType):
		return string(model.FinanceAuditEventsDBFieldName.ActorType), true

	case string(FinanceAuditEventsDTOFieldName.EventAt):
		return string(model.FinanceAuditEventsDBFieldName.EventAt), true

	case string(FinanceAuditEventsDTOFieldName.OldState):
		return string(model.FinanceAuditEventsDBFieldName.OldState), true

	case string(FinanceAuditEventsDTOFieldName.NewState):
		return string(model.FinanceAuditEventsDBFieldName.NewState), true

	case string(FinanceAuditEventsDTOFieldName.CorrelationId):
		return string(model.FinanceAuditEventsDBFieldName.CorrelationId), true

	case string(FinanceAuditEventsDTOFieldName.CausationId):
		return string(model.FinanceAuditEventsDBFieldName.CausationId), true

	case string(FinanceAuditEventsDTOFieldName.Metadata):
		return string(model.FinanceAuditEventsDBFieldName.Metadata), true

	case string(FinanceAuditEventsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceAuditEventsDBFieldName.MetaCreatedAt), true

	case string(FinanceAuditEventsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceAuditEventsDBFieldName.MetaCreatedBy), true

	case string(FinanceAuditEventsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceAuditEventsDBFieldName.MetaUpdatedAt), true

	case string(FinanceAuditEventsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceAuditEventsDBFieldName.MetaUpdatedBy), true

	case string(FinanceAuditEventsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceAuditEventsDBFieldName.MetaDeletedAt), true

	case string(FinanceAuditEventsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceAuditEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceAuditEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceAuditEventsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceAuditEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceAuditEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceAuditEventsProjectionOutputPath(path string) error {
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

func transformFinanceAuditEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceAuditEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceAuditEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceAuditEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceAuditEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceAuditEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceAuditEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceAuditEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceAuditEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceAuditEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceAuditEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceAuditEventsFilter(filter *model.Filter) {
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
			Field: string(FinanceAuditEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceAuditEventsSelectableResponse map[string]interface{}
type FinanceAuditEventsSelectableListResponse []*FinanceAuditEventsSelectableResponse

func assignFinanceAuditEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceAuditEventsSelectableValue(out FinanceAuditEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceAuditEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceAuditEventsSelectableResponse(financeAuditEvents model.FinanceAuditEvents, filter model.Filter) FinanceAuditEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceAuditEventsDBFieldName.Id),
			string(model.FinanceAuditEventsDBFieldName.AggregateType),
			string(model.FinanceAuditEventsDBFieldName.AggregateId),
			string(model.FinanceAuditEventsDBFieldName.EventType),
			string(model.FinanceAuditEventsDBFieldName.ActorUserId),
			string(model.FinanceAuditEventsDBFieldName.ActorType),
			string(model.FinanceAuditEventsDBFieldName.EventAt),
			string(model.FinanceAuditEventsDBFieldName.OldState),
			string(model.FinanceAuditEventsDBFieldName.NewState),
			string(model.FinanceAuditEventsDBFieldName.CorrelationId),
			string(model.FinanceAuditEventsDBFieldName.CausationId),
			string(model.FinanceAuditEventsDBFieldName.Metadata),
			string(model.FinanceAuditEventsDBFieldName.MetaCreatedAt),
			string(model.FinanceAuditEventsDBFieldName.MetaCreatedBy),
			string(model.FinanceAuditEventsDBFieldName.MetaUpdatedAt),
			string(model.FinanceAuditEventsDBFieldName.MetaUpdatedBy),
			string(model.FinanceAuditEventsDBFieldName.MetaDeletedAt),
			string(model.FinanceAuditEventsDBFieldName.MetaDeletedBy),
		)
	}
	financeAuditEventsSelectableResponse := FinanceAuditEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceAuditEventsDBFieldName.Id):
			key := string(FinanceAuditEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.Id, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.AggregateType):
			key := string(FinanceAuditEventsDTOFieldName.AggregateType)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.AggregateType, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.AggregateId):
			key := string(FinanceAuditEventsDTOFieldName.AggregateId)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.AggregateId, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.EventType):
			key := string(FinanceAuditEventsDTOFieldName.EventType)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.EventType, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.ActorUserId):
			key := string(FinanceAuditEventsDTOFieldName.ActorUserId)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.ActorUserId.UUID, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.ActorType):
			key := string(FinanceAuditEventsDTOFieldName.ActorType)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, model.ActorType(financeAuditEvents.ActorType), explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.EventAt):
			key := string(FinanceAuditEventsDTOFieldName.EventAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.EventAt, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.OldState):
			key := string(FinanceAuditEventsDTOFieldName.OldState)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.OldState, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.NewState):
			key := string(FinanceAuditEventsDTOFieldName.NewState)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.NewState, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.CorrelationId):
			key := string(FinanceAuditEventsDTOFieldName.CorrelationId)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.CorrelationId.String, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.CausationId):
			key := string(FinanceAuditEventsDTOFieldName.CausationId)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.CausationId.String, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.Metadata):
			key := string(FinanceAuditEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.Metadata, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaCreatedAt):
			key := string(FinanceAuditEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaCreatedAt, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaCreatedBy):
			key := string(FinanceAuditEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaCreatedBy, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaUpdatedAt):
			key := string(FinanceAuditEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaUpdatedBy):
			key := string(FinanceAuditEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaDeletedAt):
			key := string(FinanceAuditEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceAuditEventsDBFieldName.MetaDeletedBy):
			key := string(FinanceAuditEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceAuditEventsSelectableValue(financeAuditEventsSelectableResponse, key, financeAuditEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return financeAuditEventsSelectableResponse
}

func NewFinanceAuditEventsListResponseFromFilterResult(result []model.FinanceAuditEventsFilterResult, filter model.Filter) FinanceAuditEventsSelectableListResponse {
	dtoFinanceAuditEventsListResponse := FinanceAuditEventsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceAuditEventsResponse := NewFinanceAuditEventsSelectableResponse(row.FinanceAuditEvents, filter)
		dtoFinanceAuditEventsListResponse = append(dtoFinanceAuditEventsListResponse, &dtoFinanceAuditEventsResponse)
	}
	return dtoFinanceAuditEventsListResponse
}

type FinanceAuditEventsFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     FinanceAuditEventsSelectableListResponse `json:"data"`
}

func reverseFinanceAuditEventsFilterResults(result []model.FinanceAuditEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceAuditEventsFilterResponse(result []model.FinanceAuditEventsFilterResult, filter model.Filter) (resp FinanceAuditEventsFilterResponse) {
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
			reverseFinanceAuditEventsFilterResults(dataResult)
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

	resp.Data = NewFinanceAuditEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceAuditEventsCreateRequest struct {
	AggregateType string          `json:"aggregateType"`
	AggregateId   uuid.UUID       `json:"aggregateId"`
	EventType     string          `json:"eventType"`
	ActorUserId   uuid.UUID       `json:"actorUserId"`
	ActorType     model.ActorType `json:"actorType" example:"system" enums:"system,user,worker,provider"`
	EventAt       time.Time       `json:"eventAt"`
	OldState      json.RawMessage `json:"oldState"`
	NewState      json.RawMessage `json:"newState"`
	CorrelationId string          `json:"correlationId"`
	CausationId   string          `json:"causationId"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *FinanceAuditEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceAuditEventsCreateRequest) ToModel() model.FinanceAuditEvents {
	id, _ := uuid.NewV4()
	return model.FinanceAuditEvents{
		Id:            id,
		AggregateType: d.AggregateType,
		AggregateId:   d.AggregateId,
		EventType:     d.EventType,
		ActorUserId:   nuuid.From(d.ActorUserId),
		ActorType:     d.ActorType,
		EventAt:       d.EventAt,
		OldState:      d.OldState,
		NewState:      d.NewState,
		CorrelationId: null.StringFrom(d.CorrelationId),
		CausationId:   null.StringFrom(d.CausationId),
		Metadata:      d.Metadata,
	}
}

type FinanceAuditEventsListCreateRequest []*FinanceAuditEventsCreateRequest

func (d FinanceAuditEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeAuditEvents := range d {
		err = validator.Struct(financeAuditEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceAuditEventsListCreateRequest) ToModelList() []model.FinanceAuditEvents {
	out := make([]model.FinanceAuditEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceAuditEventsUpdateRequest struct {
	AggregateType string          `json:"aggregateType"`
	AggregateId   uuid.UUID       `json:"aggregateId"`
	EventType     string          `json:"eventType"`
	ActorUserId   uuid.UUID       `json:"actorUserId"`
	ActorType     model.ActorType `json:"actorType" example:"system" enums:"system,user,worker,provider"`
	EventAt       time.Time       `json:"eventAt"`
	OldState      json.RawMessage `json:"oldState"`
	NewState      json.RawMessage `json:"newState"`
	CorrelationId string          `json:"correlationId"`
	CausationId   string          `json:"causationId"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *FinanceAuditEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceAuditEventsUpdateRequest) ToModel() model.FinanceAuditEvents {
	return model.FinanceAuditEvents{
		AggregateType: d.AggregateType,
		AggregateId:   d.AggregateId,
		EventType:     d.EventType,
		ActorUserId:   nuuid.From(d.ActorUserId),
		ActorType:     d.ActorType,
		EventAt:       d.EventAt,
		OldState:      d.OldState,
		NewState:      d.NewState,
		CorrelationId: null.StringFrom(d.CorrelationId),
		CausationId:   null.StringFrom(d.CausationId),
		Metadata:      d.Metadata,
	}
}

type FinanceAuditEventsBulkUpdateRequest struct {
	Id            uuid.UUID       `json:"id"`
	AggregateType string          `json:"aggregateType"`
	AggregateId   uuid.UUID       `json:"aggregateId"`
	EventType     string          `json:"eventType"`
	ActorUserId   uuid.UUID       `json:"actorUserId"`
	ActorType     model.ActorType `json:"actorType" example:"system" enums:"system,user,worker,provider"`
	EventAt       time.Time       `json:"eventAt"`
	OldState      json.RawMessage `json:"oldState"`
	NewState      json.RawMessage `json:"newState"`
	CorrelationId string          `json:"correlationId"`
	CausationId   string          `json:"causationId"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d FinanceAuditEventsBulkUpdateRequest) PrimaryID() FinanceAuditEventsPrimaryID {
	return FinanceAuditEventsPrimaryID{
		Id: d.Id,
	}
}

type FinanceAuditEventsListBulkUpdateRequest []*FinanceAuditEventsBulkUpdateRequest

func (d FinanceAuditEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeAuditEvents := range d {
		err = validator.Struct(financeAuditEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceAuditEventsBulkUpdateRequest) ToModel() model.FinanceAuditEvents {
	return model.FinanceAuditEvents{
		Id:            d.Id,
		AggregateType: d.AggregateType,
		AggregateId:   d.AggregateId,
		EventType:     d.EventType,
		ActorUserId:   nuuid.From(d.ActorUserId),
		ActorType:     d.ActorType,
		EventAt:       d.EventAt,
		OldState:      d.OldState,
		NewState:      d.NewState,
		CorrelationId: null.StringFrom(d.CorrelationId),
		CausationId:   null.StringFrom(d.CausationId),
		Metadata:      d.Metadata,
	}
}

type FinanceAuditEventsResponse struct {
	Id            uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AggregateType string          `json:"aggregateType" validate:"required"`
	AggregateId   uuid.UUID       `json:"aggregateId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EventType     string          `json:"eventType" validate:"required"`
	ActorUserId   uuid.UUID       `json:"actorUserId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ActorType     model.ActorType `json:"actorType" validate:"required,oneof=system user worker provider" enums:"system,user,worker,provider"`
	EventAt       time.Time       `json:"eventAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	OldState      json.RawMessage `json:"oldState" validate:"required" swaggertype:"object"`
	NewState      json.RawMessage `json:"newState" validate:"required" swaggertype:"object"`
	CorrelationId string          `json:"correlationId"`
	CausationId   string          `json:"causationId"`
	Metadata      json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewFinanceAuditEventsResponse(financeAuditEvents model.FinanceAuditEvents) FinanceAuditEventsResponse {
	return FinanceAuditEventsResponse{
		Id:            financeAuditEvents.Id,
		AggregateType: financeAuditEvents.AggregateType,
		AggregateId:   financeAuditEvents.AggregateId,
		EventType:     financeAuditEvents.EventType,
		ActorUserId:   financeAuditEvents.ActorUserId.UUID,
		ActorType:     model.ActorType(financeAuditEvents.ActorType),
		EventAt:       financeAuditEvents.EventAt,
		OldState:      financeAuditEvents.OldState,
		NewState:      financeAuditEvents.NewState,
		CorrelationId: financeAuditEvents.CorrelationId.String,
		CausationId:   financeAuditEvents.CausationId.String,
		Metadata:      financeAuditEvents.Metadata,
	}
}

type FinanceAuditEventsListResponse []*FinanceAuditEventsResponse

func NewFinanceAuditEventsListResponse(financeAuditEventsList model.FinanceAuditEventsList) FinanceAuditEventsListResponse {
	dtoFinanceAuditEventsListResponse := FinanceAuditEventsListResponse{}
	for _, financeAuditEvents := range financeAuditEventsList {
		dtoFinanceAuditEventsResponse := NewFinanceAuditEventsResponse(*financeAuditEvents)
		dtoFinanceAuditEventsListResponse = append(dtoFinanceAuditEventsListResponse, &dtoFinanceAuditEventsResponse)
	}
	return dtoFinanceAuditEventsListResponse
}

type FinanceAuditEventsPrimaryIDList []FinanceAuditEventsPrimaryID

func (d FinanceAuditEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeAuditEvents := range d {
		err = validator.Struct(financeAuditEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceAuditEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceAuditEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceAuditEventsPrimaryID) ToModel() model.FinanceAuditEventsPrimaryID {
	return model.FinanceAuditEventsPrimaryID{
		Id: d.Id,
	}
}

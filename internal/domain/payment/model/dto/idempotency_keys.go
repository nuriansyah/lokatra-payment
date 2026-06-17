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

type IdempotencyKeysDTOFieldNameType string

type idempotencyKeysDTOFieldName struct {
	Id             IdempotencyKeysDTOFieldNameType
	Key            IdempotencyKeysDTOFieldNameType
	ActorType      IdempotencyKeysDTOFieldNameType
	ActorId        IdempotencyKeysDTOFieldNameType
	RequestHash    IdempotencyKeysDTOFieldNameType
	Status         IdempotencyKeysDTOFieldNameType
	ResourceType   IdempotencyKeysDTOFieldNameType
	ResourceId     IdempotencyKeysDTOFieldNameType
	ResponseStatus IdempotencyKeysDTOFieldNameType
	ResponseBody   IdempotencyKeysDTOFieldNameType
	LockedUntil    IdempotencyKeysDTOFieldNameType
	CompletedAt    IdempotencyKeysDTOFieldNameType
	Metadata       IdempotencyKeysDTOFieldNameType
	MetaCreatedAt  IdempotencyKeysDTOFieldNameType
	MetaCreatedBy  IdempotencyKeysDTOFieldNameType
	MetaUpdatedAt  IdempotencyKeysDTOFieldNameType
	MetaUpdatedBy  IdempotencyKeysDTOFieldNameType
	MetaDeletedAt  IdempotencyKeysDTOFieldNameType
	MetaDeletedBy  IdempotencyKeysDTOFieldNameType
}

var IdempotencyKeysDTOFieldName = idempotencyKeysDTOFieldName{
	Id:             "id",
	Key:            "key",
	ActorType:      "actorType",
	ActorId:        "actorId",
	RequestHash:    "requestHash",
	Status:         "status",
	ResourceType:   "resourceType",
	ResourceId:     "resourceId",
	ResponseStatus: "responseStatus",
	ResponseBody:   "responseBody",
	LockedUntil:    "lockedUntil",
	CompletedAt:    "completedAt",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformIdempotencyKeysDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(IdempotencyKeysDTOFieldName.Id):
		return string(model.IdempotencyKeysDBFieldName.Id), true

	case string(IdempotencyKeysDTOFieldName.Key):
		return string(model.IdempotencyKeysDBFieldName.Key), true

	case string(IdempotencyKeysDTOFieldName.ActorType):
		return string(model.IdempotencyKeysDBFieldName.ActorType), true

	case string(IdempotencyKeysDTOFieldName.ActorId):
		return string(model.IdempotencyKeysDBFieldName.ActorId), true

	case string(IdempotencyKeysDTOFieldName.RequestHash):
		return string(model.IdempotencyKeysDBFieldName.RequestHash), true

	case string(IdempotencyKeysDTOFieldName.Status):
		return string(model.IdempotencyKeysDBFieldName.Status), true

	case string(IdempotencyKeysDTOFieldName.ResourceType):
		return string(model.IdempotencyKeysDBFieldName.ResourceType), true

	case string(IdempotencyKeysDTOFieldName.ResourceId):
		return string(model.IdempotencyKeysDBFieldName.ResourceId), true

	case string(IdempotencyKeysDTOFieldName.ResponseStatus):
		return string(model.IdempotencyKeysDBFieldName.ResponseStatus), true

	case string(IdempotencyKeysDTOFieldName.ResponseBody):
		return string(model.IdempotencyKeysDBFieldName.ResponseBody), true

	case string(IdempotencyKeysDTOFieldName.LockedUntil):
		return string(model.IdempotencyKeysDBFieldName.LockedUntil), true

	case string(IdempotencyKeysDTOFieldName.CompletedAt):
		return string(model.IdempotencyKeysDBFieldName.CompletedAt), true

	case string(IdempotencyKeysDTOFieldName.Metadata):
		return string(model.IdempotencyKeysDBFieldName.Metadata), true

	case string(IdempotencyKeysDTOFieldName.MetaCreatedAt):
		return string(model.IdempotencyKeysDBFieldName.MetaCreatedAt), true

	case string(IdempotencyKeysDTOFieldName.MetaCreatedBy):
		return string(model.IdempotencyKeysDBFieldName.MetaCreatedBy), true

	case string(IdempotencyKeysDTOFieldName.MetaUpdatedAt):
		return string(model.IdempotencyKeysDBFieldName.MetaUpdatedAt), true

	case string(IdempotencyKeysDTOFieldName.MetaUpdatedBy):
		return string(model.IdempotencyKeysDBFieldName.MetaUpdatedBy), true

	case string(IdempotencyKeysDTOFieldName.MetaDeletedAt):
		return string(model.IdempotencyKeysDBFieldName.MetaDeletedAt), true

	case string(IdempotencyKeysDTOFieldName.MetaDeletedBy):
		return string(model.IdempotencyKeysDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewIdempotencyKeysFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isIdempotencyKeysBaseFilterField(field string) bool {
	spec, found := model.NewIdempotencyKeysFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeIdempotencyKeysProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateIdempotencyKeysProjectionOutputPath(path string) error {
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

func transformIdempotencyKeysFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformIdempotencyKeysFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformIdempotencyKeysFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isIdempotencyKeysBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateIdempotencyKeysProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeIdempotencyKeysProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformIdempotencyKeysFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultIdempotencyKeysFilter(filter *model.Filter) {
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
			Field: string(IdempotencyKeysDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type IdempotencyKeysSelectableResponse map[string]interface{}
type IdempotencyKeysSelectableListResponse []*IdempotencyKeysSelectableResponse

func assignIdempotencyKeysNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setIdempotencyKeysSelectableValue(out IdempotencyKeysSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignIdempotencyKeysNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewIdempotencyKeysSelectableResponse(idempotencyKeys model.IdempotencyKeys, filter model.Filter) IdempotencyKeysSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.IdempotencyKeysDBFieldName.Id),
			string(model.IdempotencyKeysDBFieldName.Key),
			string(model.IdempotencyKeysDBFieldName.ActorType),
			string(model.IdempotencyKeysDBFieldName.ActorId),
			string(model.IdempotencyKeysDBFieldName.RequestHash),
			string(model.IdempotencyKeysDBFieldName.Status),
			string(model.IdempotencyKeysDBFieldName.ResourceType),
			string(model.IdempotencyKeysDBFieldName.ResourceId),
			string(model.IdempotencyKeysDBFieldName.ResponseStatus),
			string(model.IdempotencyKeysDBFieldName.ResponseBody),
			string(model.IdempotencyKeysDBFieldName.LockedUntil),
			string(model.IdempotencyKeysDBFieldName.CompletedAt),
			string(model.IdempotencyKeysDBFieldName.Metadata),
			string(model.IdempotencyKeysDBFieldName.MetaCreatedAt),
			string(model.IdempotencyKeysDBFieldName.MetaCreatedBy),
			string(model.IdempotencyKeysDBFieldName.MetaUpdatedAt),
			string(model.IdempotencyKeysDBFieldName.MetaUpdatedBy),
			string(model.IdempotencyKeysDBFieldName.MetaDeletedAt),
			string(model.IdempotencyKeysDBFieldName.MetaDeletedBy),
		)
	}
	idempotencyKeysSelectableResponse := IdempotencyKeysSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.IdempotencyKeysDBFieldName.Id):
			key := string(IdempotencyKeysDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.Id, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.Key):
			key := string(IdempotencyKeysDTOFieldName.Key)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.Key, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ActorType):
			key := string(IdempotencyKeysDTOFieldName.ActorType)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.ActorType.String, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ActorId):
			key := string(IdempotencyKeysDTOFieldName.ActorId)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.ActorId.UUID, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.RequestHash):
			key := string(IdempotencyKeysDTOFieldName.RequestHash)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.RequestHash, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.Status):
			key := string(IdempotencyKeysDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, model.IdempotencyStatus(idempotencyKeys.Status), explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ResourceType):
			key := string(IdempotencyKeysDTOFieldName.ResourceType)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.ResourceType.String, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ResourceId):
			key := string(IdempotencyKeysDTOFieldName.ResourceId)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.ResourceId.UUID, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ResponseStatus):
			key := string(IdempotencyKeysDTOFieldName.ResponseStatus)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, int(idempotencyKeys.ResponseStatus.ValueOrZero()), explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.ResponseBody):
			key := string(IdempotencyKeysDTOFieldName.ResponseBody)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.ResponseBody, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.LockedUntil):
			key := string(IdempotencyKeysDTOFieldName.LockedUntil)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.LockedUntil.Time, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.CompletedAt):
			key := string(IdempotencyKeysDTOFieldName.CompletedAt)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.CompletedAt.Time, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.Metadata):
			key := string(IdempotencyKeysDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.Metadata, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaCreatedAt):
			key := string(IdempotencyKeysDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaCreatedAt, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaCreatedBy):
			key := string(IdempotencyKeysDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaCreatedBy, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaUpdatedAt):
			key := string(IdempotencyKeysDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaUpdatedAt, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaUpdatedBy):
			key := string(IdempotencyKeysDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaDeletedAt):
			key := string(IdempotencyKeysDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaDeletedAt.Time, explicitAlias)

		case string(model.IdempotencyKeysDBFieldName.MetaDeletedBy):
			key := string(IdempotencyKeysDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setIdempotencyKeysSelectableValue(idempotencyKeysSelectableResponse, key, idempotencyKeys.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return idempotencyKeysSelectableResponse
}

func NewIdempotencyKeysListResponseFromFilterResult(result []model.IdempotencyKeysFilterResult, filter model.Filter) IdempotencyKeysSelectableListResponse {
	dtoIdempotencyKeysListResponse := IdempotencyKeysSelectableListResponse{}
	for _, row := range result {
		dtoIdempotencyKeysResponse := NewIdempotencyKeysSelectableResponse(row.IdempotencyKeys, filter)
		dtoIdempotencyKeysListResponse = append(dtoIdempotencyKeysListResponse, &dtoIdempotencyKeysResponse)
	}
	return dtoIdempotencyKeysListResponse
}

type IdempotencyKeysFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     IdempotencyKeysSelectableListResponse `json:"data"`
}

func reverseIdempotencyKeysFilterResults(result []model.IdempotencyKeysFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewIdempotencyKeysFilterResponse(result []model.IdempotencyKeysFilterResult, filter model.Filter) (resp IdempotencyKeysFilterResponse) {
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
			reverseIdempotencyKeysFilterResults(dataResult)
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

	resp.Data = NewIdempotencyKeysListResponseFromFilterResult(dataResult, filter)
	return resp
}

type IdempotencyKeysCreateRequest struct {
	Key            string                  `json:"key"`
	ActorType      string                  `json:"actorType"`
	ActorId        uuid.UUID               `json:"actorId"`
	RequestHash    string                  `json:"requestHash"`
	Status         model.IdempotencyStatus `json:"status" example:"processing" enums:"processing,completed,failed"`
	ResourceType   string                  `json:"resourceType"`
	ResourceId     uuid.UUID               `json:"resourceId"`
	ResponseStatus int                     `json:"responseStatus"`
	ResponseBody   json.RawMessage         `json:"responseBody"`
	LockedUntil    time.Time               `json:"lockedUntil"`
	CompletedAt    time.Time               `json:"completedAt"`
	Metadata       json.RawMessage         `json:"metadata"`
}

func (d *IdempotencyKeysCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *IdempotencyKeysCreateRequest) ToModel() model.IdempotencyKeys {
	id, _ := uuid.NewV4()
	return model.IdempotencyKeys{
		Id:             id,
		Key:            d.Key,
		ActorType:      null.StringFrom(d.ActorType),
		ActorId:        nuuid.From(d.ActorId),
		RequestHash:    d.RequestHash,
		Status:         d.Status,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type IdempotencyKeysListCreateRequest []*IdempotencyKeysCreateRequest

func (d IdempotencyKeysListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d IdempotencyKeysListCreateRequest) ToModelList() []model.IdempotencyKeys {
	out := make([]model.IdempotencyKeys, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type IdempotencyKeysUpdateRequest struct {
	Key            string                  `json:"key"`
	ActorType      string                  `json:"actorType"`
	ActorId        uuid.UUID               `json:"actorId"`
	RequestHash    string                  `json:"requestHash"`
	Status         model.IdempotencyStatus `json:"status" example:"processing" enums:"processing,completed,failed"`
	ResourceType   string                  `json:"resourceType"`
	ResourceId     uuid.UUID               `json:"resourceId"`
	ResponseStatus int                     `json:"responseStatus"`
	ResponseBody   json.RawMessage         `json:"responseBody"`
	LockedUntil    time.Time               `json:"lockedUntil"`
	CompletedAt    time.Time               `json:"completedAt"`
	Metadata       json.RawMessage         `json:"metadata"`
}

func (d *IdempotencyKeysUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d IdempotencyKeysUpdateRequest) ToModel() model.IdempotencyKeys {
	return model.IdempotencyKeys{
		Key:            d.Key,
		ActorType:      null.StringFrom(d.ActorType),
		ActorId:        nuuid.From(d.ActorId),
		RequestHash:    d.RequestHash,
		Status:         d.Status,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type IdempotencyKeysBulkUpdateRequest struct {
	Id             uuid.UUID               `json:"id"`
	Key            string                  `json:"key"`
	ActorType      string                  `json:"actorType"`
	ActorId        uuid.UUID               `json:"actorId"`
	RequestHash    string                  `json:"requestHash"`
	Status         model.IdempotencyStatus `json:"status" example:"processing" enums:"processing,completed,failed"`
	ResourceType   string                  `json:"resourceType"`
	ResourceId     uuid.UUID               `json:"resourceId"`
	ResponseStatus int                     `json:"responseStatus"`
	ResponseBody   json.RawMessage         `json:"responseBody"`
	LockedUntil    time.Time               `json:"lockedUntil"`
	CompletedAt    time.Time               `json:"completedAt"`
	Metadata       json.RawMessage         `json:"metadata"`
}

func (d IdempotencyKeysBulkUpdateRequest) PrimaryID() IdempotencyKeysPrimaryID {
	return IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type IdempotencyKeysListBulkUpdateRequest []*IdempotencyKeysBulkUpdateRequest

func (d IdempotencyKeysListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d IdempotencyKeysBulkUpdateRequest) ToModel() model.IdempotencyKeys {
	return model.IdempotencyKeys{
		Id:             d.Id,
		Key:            d.Key,
		ActorType:      null.StringFrom(d.ActorType),
		ActorId:        nuuid.From(d.ActorId),
		RequestHash:    d.RequestHash,
		Status:         d.Status,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type IdempotencyKeysResponse struct {
	Id             uuid.UUID               `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Key            string                  `json:"key" validate:"required"`
	ActorType      string                  `json:"actorType"`
	ActorId        uuid.UUID               `json:"actorId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RequestHash    string                  `json:"requestHash" validate:"required"`
	Status         model.IdempotencyStatus `json:"status" validate:"oneof=processing completed failed" enums:"processing,completed,failed"`
	ResourceType   string                  `json:"resourceType"`
	ResourceId     uuid.UUID               `json:"resourceId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ResponseStatus int                     `json:"responseStatus" example:"1"`
	ResponseBody   json.RawMessage         `json:"responseBody" validate:"required" swaggertype:"object"`
	LockedUntil    time.Time               `json:"lockedUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CompletedAt    time.Time               `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage         `json:"metadata" swaggertype:"object"`
}

func NewIdempotencyKeysResponse(idempotencyKeys model.IdempotencyKeys) IdempotencyKeysResponse {
	return IdempotencyKeysResponse{
		Id:             idempotencyKeys.Id,
		Key:            idempotencyKeys.Key,
		ActorType:      idempotencyKeys.ActorType.String,
		ActorId:        idempotencyKeys.ActorId.UUID,
		RequestHash:    idempotencyKeys.RequestHash,
		Status:         model.IdempotencyStatus(idempotencyKeys.Status),
		ResourceType:   idempotencyKeys.ResourceType.String,
		ResourceId:     idempotencyKeys.ResourceId.UUID,
		ResponseStatus: int(idempotencyKeys.ResponseStatus.ValueOrZero()),
		ResponseBody:   idempotencyKeys.ResponseBody,
		LockedUntil:    idempotencyKeys.LockedUntil.Time,
		CompletedAt:    idempotencyKeys.CompletedAt.Time,
		Metadata:       idempotencyKeys.Metadata,
	}
}

type IdempotencyKeysListResponse []*IdempotencyKeysResponse

func NewIdempotencyKeysListResponse(idempotencyKeysList model.IdempotencyKeysList) IdempotencyKeysListResponse {
	dtoIdempotencyKeysListResponse := IdempotencyKeysListResponse{}
	for _, idempotencyKeys := range idempotencyKeysList {
		dtoIdempotencyKeysResponse := NewIdempotencyKeysResponse(*idempotencyKeys)
		dtoIdempotencyKeysListResponse = append(dtoIdempotencyKeysListResponse, &dtoIdempotencyKeysResponse)
	}
	return dtoIdempotencyKeysListResponse
}

type IdempotencyKeysPrimaryIDList []IdempotencyKeysPrimaryID

func (d IdempotencyKeysPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

type IdempotencyKeysPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *IdempotencyKeysPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d IdempotencyKeysPrimaryID) ToModel() model.IdempotencyKeysPrimaryID {
	return model.IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

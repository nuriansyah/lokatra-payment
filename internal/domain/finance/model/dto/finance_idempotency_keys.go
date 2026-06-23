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

type FinanceIdempotencyKeysDTOFieldNameType string

type financeIdempotencyKeysDTOFieldName struct {
	Id             FinanceIdempotencyKeysDTOFieldNameType
	Scope          FinanceIdempotencyKeysDTOFieldNameType
	Operation      FinanceIdempotencyKeysDTOFieldNameType
	IdempotencyKey FinanceIdempotencyKeysDTOFieldNameType
	RequestHash    FinanceIdempotencyKeysDTOFieldNameType
	ResourceType   FinanceIdempotencyKeysDTOFieldNameType
	ResourceId     FinanceIdempotencyKeysDTOFieldNameType
	ResponseStatus FinanceIdempotencyKeysDTOFieldNameType
	ResponseBody   FinanceIdempotencyKeysDTOFieldNameType
	LockedUntil    FinanceIdempotencyKeysDTOFieldNameType
	CompletedAt    FinanceIdempotencyKeysDTOFieldNameType
	Metadata       FinanceIdempotencyKeysDTOFieldNameType
	MetaCreatedAt  FinanceIdempotencyKeysDTOFieldNameType
	MetaCreatedBy  FinanceIdempotencyKeysDTOFieldNameType
	MetaUpdatedAt  FinanceIdempotencyKeysDTOFieldNameType
	MetaUpdatedBy  FinanceIdempotencyKeysDTOFieldNameType
	MetaDeletedAt  FinanceIdempotencyKeysDTOFieldNameType
	MetaDeletedBy  FinanceIdempotencyKeysDTOFieldNameType
}

var FinanceIdempotencyKeysDTOFieldName = financeIdempotencyKeysDTOFieldName{
	Id:             "id",
	Scope:          "scope",
	Operation:      "operation",
	IdempotencyKey: "idempotencyKey",
	RequestHash:    "requestHash",
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

func transformFinanceIdempotencyKeysDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceIdempotencyKeysDTOFieldName.Id):
		return string(model.FinanceIdempotencyKeysDBFieldName.Id), true

	case string(FinanceIdempotencyKeysDTOFieldName.Scope):
		return string(model.FinanceIdempotencyKeysDBFieldName.Scope), true

	case string(FinanceIdempotencyKeysDTOFieldName.Operation):
		return string(model.FinanceIdempotencyKeysDBFieldName.Operation), true

	case string(FinanceIdempotencyKeysDTOFieldName.IdempotencyKey):
		return string(model.FinanceIdempotencyKeysDBFieldName.IdempotencyKey), true

	case string(FinanceIdempotencyKeysDTOFieldName.RequestHash):
		return string(model.FinanceIdempotencyKeysDBFieldName.RequestHash), true

	case string(FinanceIdempotencyKeysDTOFieldName.ResourceType):
		return string(model.FinanceIdempotencyKeysDBFieldName.ResourceType), true

	case string(FinanceIdempotencyKeysDTOFieldName.ResourceId):
		return string(model.FinanceIdempotencyKeysDBFieldName.ResourceId), true

	case string(FinanceIdempotencyKeysDTOFieldName.ResponseStatus):
		return string(model.FinanceIdempotencyKeysDBFieldName.ResponseStatus), true

	case string(FinanceIdempotencyKeysDTOFieldName.ResponseBody):
		return string(model.FinanceIdempotencyKeysDBFieldName.ResponseBody), true

	case string(FinanceIdempotencyKeysDTOFieldName.LockedUntil):
		return string(model.FinanceIdempotencyKeysDBFieldName.LockedUntil), true

	case string(FinanceIdempotencyKeysDTOFieldName.CompletedAt):
		return string(model.FinanceIdempotencyKeysDBFieldName.CompletedAt), true

	case string(FinanceIdempotencyKeysDTOFieldName.Metadata):
		return string(model.FinanceIdempotencyKeysDBFieldName.Metadata), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaCreatedAt):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedAt), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaCreatedBy):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedBy), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedAt), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedBy), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaDeletedAt):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedAt), true

	case string(FinanceIdempotencyKeysDTOFieldName.MetaDeletedBy):
		return string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceIdempotencyKeysFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceIdempotencyKeysBaseFilterField(field string) bool {
	spec, found := model.NewFinanceIdempotencyKeysFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceIdempotencyKeysProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceIdempotencyKeysProjectionOutputPath(path string) error {
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

func transformFinanceIdempotencyKeysFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceIdempotencyKeysDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceIdempotencyKeysFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceIdempotencyKeysFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceIdempotencyKeysDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceIdempotencyKeysBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceIdempotencyKeysProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceIdempotencyKeysProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceIdempotencyKeysDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceIdempotencyKeysDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceIdempotencyKeysFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceIdempotencyKeysFilter(filter *model.Filter) {
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
			Field: string(FinanceIdempotencyKeysDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceIdempotencyKeysSelectableResponse map[string]interface{}
type FinanceIdempotencyKeysSelectableListResponse []*FinanceIdempotencyKeysSelectableResponse

func assignFinanceIdempotencyKeysNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceIdempotencyKeysSelectableValue(out FinanceIdempotencyKeysSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceIdempotencyKeysNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceIdempotencyKeysSelectableResponse(financeIdempotencyKeys model.FinanceIdempotencyKeys, filter model.Filter) FinanceIdempotencyKeysSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceIdempotencyKeysDBFieldName.Id),
			string(model.FinanceIdempotencyKeysDBFieldName.Scope),
			string(model.FinanceIdempotencyKeysDBFieldName.Operation),
			string(model.FinanceIdempotencyKeysDBFieldName.IdempotencyKey),
			string(model.FinanceIdempotencyKeysDBFieldName.RequestHash),
			string(model.FinanceIdempotencyKeysDBFieldName.ResourceType),
			string(model.FinanceIdempotencyKeysDBFieldName.ResourceId),
			string(model.FinanceIdempotencyKeysDBFieldName.ResponseStatus),
			string(model.FinanceIdempotencyKeysDBFieldName.ResponseBody),
			string(model.FinanceIdempotencyKeysDBFieldName.LockedUntil),
			string(model.FinanceIdempotencyKeysDBFieldName.CompletedAt),
			string(model.FinanceIdempotencyKeysDBFieldName.Metadata),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedAt),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedBy),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedAt),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedBy),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedAt),
			string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedBy),
		)
	}
	financeIdempotencyKeysSelectableResponse := FinanceIdempotencyKeysSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceIdempotencyKeysDBFieldName.Id):
			key := string(FinanceIdempotencyKeysDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.Id, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.Scope):
			key := string(FinanceIdempotencyKeysDTOFieldName.Scope)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.Scope, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.Operation):
			key := string(FinanceIdempotencyKeysDTOFieldName.Operation)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.Operation, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.IdempotencyKey):
			key := string(FinanceIdempotencyKeysDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.IdempotencyKey, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.RequestHash):
			key := string(FinanceIdempotencyKeysDTOFieldName.RequestHash)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.RequestHash, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.ResourceType):
			key := string(FinanceIdempotencyKeysDTOFieldName.ResourceType)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.ResourceType.String, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.ResourceId):
			key := string(FinanceIdempotencyKeysDTOFieldName.ResourceId)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.ResourceId.UUID, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.ResponseStatus):
			key := string(FinanceIdempotencyKeysDTOFieldName.ResponseStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, int(financeIdempotencyKeys.ResponseStatus.ValueOrZero()), explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.ResponseBody):
			key := string(FinanceIdempotencyKeysDTOFieldName.ResponseBody)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.ResponseBody, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.LockedUntil):
			key := string(FinanceIdempotencyKeysDTOFieldName.LockedUntil)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.LockedUntil.Time, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.CompletedAt):
			key := string(FinanceIdempotencyKeysDTOFieldName.CompletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.CompletedAt.Time, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.Metadata):
			key := string(FinanceIdempotencyKeysDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.Metadata, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedAt):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaCreatedAt, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaCreatedBy):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaCreatedBy, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedAt):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaUpdatedBy):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedAt):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceIdempotencyKeysDBFieldName.MetaDeletedBy):
			key := string(FinanceIdempotencyKeysDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceIdempotencyKeysSelectableValue(financeIdempotencyKeysSelectableResponse, key, financeIdempotencyKeys.MetaDeletedBy, explicitAlias)

		}
	}
	return financeIdempotencyKeysSelectableResponse
}

func NewFinanceIdempotencyKeysListResponseFromFilterResult(result []model.FinanceIdempotencyKeysFilterResult, filter model.Filter) FinanceIdempotencyKeysSelectableListResponse {
	dtoFinanceIdempotencyKeysListResponse := FinanceIdempotencyKeysSelectableListResponse{}
	for _, row := range result {
		dtoFinanceIdempotencyKeysResponse := NewFinanceIdempotencyKeysSelectableResponse(row.FinanceIdempotencyKeys, filter)
		dtoFinanceIdempotencyKeysListResponse = append(dtoFinanceIdempotencyKeysListResponse, &dtoFinanceIdempotencyKeysResponse)
	}
	return dtoFinanceIdempotencyKeysListResponse
}

type FinanceIdempotencyKeysFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     FinanceIdempotencyKeysSelectableListResponse `json:"data"`
}

func reverseFinanceIdempotencyKeysFilterResults(result []model.FinanceIdempotencyKeysFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceIdempotencyKeysFilterResponse(result []model.FinanceIdempotencyKeysFilterResult, filter model.Filter) (resp FinanceIdempotencyKeysFilterResponse) {
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
			reverseFinanceIdempotencyKeysFilterResults(dataResult)
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

	resp.Data = NewFinanceIdempotencyKeysListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceIdempotencyKeysCreateRequest struct {
	Scope          string          `json:"scope"`
	Operation      string          `json:"operation"`
	IdempotencyKey string          `json:"idempotencyKey"`
	RequestHash    string          `json:"requestHash"`
	ResourceType   string          `json:"resourceType"`
	ResourceId     uuid.UUID       `json:"resourceId"`
	ResponseStatus int             `json:"responseStatus"`
	ResponseBody   json.RawMessage `json:"responseBody"`
	LockedUntil    time.Time       `json:"lockedUntil"`
	CompletedAt    time.Time       `json:"completedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *FinanceIdempotencyKeysCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceIdempotencyKeysCreateRequest) ToModel() model.FinanceIdempotencyKeys {
	id, _ := uuid.NewV4()
	return model.FinanceIdempotencyKeys{
		Id:             id,
		Scope:          d.Scope,
		Operation:      d.Operation,
		IdempotencyKey: d.IdempotencyKey,
		RequestHash:    d.RequestHash,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type FinanceIdempotencyKeysListCreateRequest []*FinanceIdempotencyKeysCreateRequest

func (d FinanceIdempotencyKeysListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeIdempotencyKeys := range d {
		err = validator.Struct(financeIdempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceIdempotencyKeysListCreateRequest) ToModelList() []model.FinanceIdempotencyKeys {
	out := make([]model.FinanceIdempotencyKeys, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceIdempotencyKeysUpdateRequest struct {
	Scope          string          `json:"scope"`
	Operation      string          `json:"operation"`
	IdempotencyKey string          `json:"idempotencyKey"`
	RequestHash    string          `json:"requestHash"`
	ResourceType   string          `json:"resourceType"`
	ResourceId     uuid.UUID       `json:"resourceId"`
	ResponseStatus int             `json:"responseStatus"`
	ResponseBody   json.RawMessage `json:"responseBody"`
	LockedUntil    time.Time       `json:"lockedUntil"`
	CompletedAt    time.Time       `json:"completedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *FinanceIdempotencyKeysUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceIdempotencyKeysUpdateRequest) ToModel() model.FinanceIdempotencyKeys {
	return model.FinanceIdempotencyKeys{
		Scope:          d.Scope,
		Operation:      d.Operation,
		IdempotencyKey: d.IdempotencyKey,
		RequestHash:    d.RequestHash,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type FinanceIdempotencyKeysBulkUpdateRequest struct {
	Id             uuid.UUID       `json:"id"`
	Scope          string          `json:"scope"`
	Operation      string          `json:"operation"`
	IdempotencyKey string          `json:"idempotencyKey"`
	RequestHash    string          `json:"requestHash"`
	ResourceType   string          `json:"resourceType"`
	ResourceId     uuid.UUID       `json:"resourceId"`
	ResponseStatus int             `json:"responseStatus"`
	ResponseBody   json.RawMessage `json:"responseBody"`
	LockedUntil    time.Time       `json:"lockedUntil"`
	CompletedAt    time.Time       `json:"completedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d FinanceIdempotencyKeysBulkUpdateRequest) PrimaryID() FinanceIdempotencyKeysPrimaryID {
	return FinanceIdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type FinanceIdempotencyKeysListBulkUpdateRequest []*FinanceIdempotencyKeysBulkUpdateRequest

func (d FinanceIdempotencyKeysListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeIdempotencyKeys := range d {
		err = validator.Struct(financeIdempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceIdempotencyKeysBulkUpdateRequest) ToModel() model.FinanceIdempotencyKeys {
	return model.FinanceIdempotencyKeys{
		Id:             d.Id,
		Scope:          d.Scope,
		Operation:      d.Operation,
		IdempotencyKey: d.IdempotencyKey,
		RequestHash:    d.RequestHash,
		ResourceType:   null.StringFrom(d.ResourceType),
		ResourceId:     nuuid.From(d.ResourceId),
		ResponseStatus: null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:   d.ResponseBody,
		LockedUntil:    null.TimeFrom(d.LockedUntil),
		CompletedAt:    null.TimeFrom(d.CompletedAt),
		Metadata:       d.Metadata,
	}
}

type FinanceIdempotencyKeysResponse struct {
	Id             uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Scope          string          `json:"scope" validate:"required"`
	Operation      string          `json:"operation" validate:"required"`
	IdempotencyKey string          `json:"idempotencyKey" validate:"required"`
	RequestHash    string          `json:"requestHash" validate:"required"`
	ResourceType   string          `json:"resourceType"`
	ResourceId     uuid.UUID       `json:"resourceId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ResponseStatus int             `json:"responseStatus" example:"1"`
	ResponseBody   json.RawMessage `json:"responseBody" swaggertype:"object"`
	LockedUntil    time.Time       `json:"lockedUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CompletedAt    time.Time       `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewFinanceIdempotencyKeysResponse(financeIdempotencyKeys model.FinanceIdempotencyKeys) FinanceIdempotencyKeysResponse {
	return FinanceIdempotencyKeysResponse{
		Id:             financeIdempotencyKeys.Id,
		Scope:          financeIdempotencyKeys.Scope,
		Operation:      financeIdempotencyKeys.Operation,
		IdempotencyKey: financeIdempotencyKeys.IdempotencyKey,
		RequestHash:    financeIdempotencyKeys.RequestHash,
		ResourceType:   financeIdempotencyKeys.ResourceType.String,
		ResourceId:     financeIdempotencyKeys.ResourceId.UUID,
		ResponseStatus: int(financeIdempotencyKeys.ResponseStatus.ValueOrZero()),
		ResponseBody:   financeIdempotencyKeys.ResponseBody,
		LockedUntil:    financeIdempotencyKeys.LockedUntil.Time,
		CompletedAt:    financeIdempotencyKeys.CompletedAt.Time,
		Metadata:       financeIdempotencyKeys.Metadata,
	}
}

type FinanceIdempotencyKeysListResponse []*FinanceIdempotencyKeysResponse

func NewFinanceIdempotencyKeysListResponse(financeIdempotencyKeysList model.FinanceIdempotencyKeysList) FinanceIdempotencyKeysListResponse {
	dtoFinanceIdempotencyKeysListResponse := FinanceIdempotencyKeysListResponse{}
	for _, financeIdempotencyKeys := range financeIdempotencyKeysList {
		dtoFinanceIdempotencyKeysResponse := NewFinanceIdempotencyKeysResponse(*financeIdempotencyKeys)
		dtoFinanceIdempotencyKeysListResponse = append(dtoFinanceIdempotencyKeysListResponse, &dtoFinanceIdempotencyKeysResponse)
	}
	return dtoFinanceIdempotencyKeysListResponse
}

type FinanceIdempotencyKeysPrimaryIDList []FinanceIdempotencyKeysPrimaryID

func (d FinanceIdempotencyKeysPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeIdempotencyKeys := range d {
		err = validator.Struct(financeIdempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceIdempotencyKeysPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceIdempotencyKeysPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceIdempotencyKeysPrimaryID) ToModel() model.FinanceIdempotencyKeysPrimaryID {
	return model.FinanceIdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

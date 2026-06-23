package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderCircuitBreakersDTOFieldNameType string

type providerCircuitBreakersDTOFieldName struct {
	Id                ProviderCircuitBreakersDTOFieldNameType
	ProviderAccountId ProviderCircuitBreakersDTOFieldNameType
	MethodCode        ProviderCircuitBreakersDTOFieldNameType
	ChannelCode       ProviderCircuitBreakersDTOFieldNameType
	Status            ProviderCircuitBreakersDTOFieldNameType
	FailureCount      ProviderCircuitBreakersDTOFieldNameType
	SuccessCount      ProviderCircuitBreakersDTOFieldNameType
	LastFailureAt     ProviderCircuitBreakersDTOFieldNameType
	LastSuccessAt     ProviderCircuitBreakersDTOFieldNameType
	OpenedAt          ProviderCircuitBreakersDTOFieldNameType
	OpenUntil         ProviderCircuitBreakersDTOFieldNameType
	HalfOpenAt        ProviderCircuitBreakersDTOFieldNameType
	Reason            ProviderCircuitBreakersDTOFieldNameType
	Metadata          ProviderCircuitBreakersDTOFieldNameType
	MetaCreatedAt     ProviderCircuitBreakersDTOFieldNameType
	MetaCreatedBy     ProviderCircuitBreakersDTOFieldNameType
	MetaUpdatedAt     ProviderCircuitBreakersDTOFieldNameType
	MetaUpdatedBy     ProviderCircuitBreakersDTOFieldNameType
	MetaDeletedAt     ProviderCircuitBreakersDTOFieldNameType
	MetaDeletedBy     ProviderCircuitBreakersDTOFieldNameType
}

var ProviderCircuitBreakersDTOFieldName = providerCircuitBreakersDTOFieldName{
	Id:                "id",
	ProviderAccountId: "providerAccountId",
	MethodCode:        "methodCode",
	ChannelCode:       "channelCode",
	Status:            "status",
	FailureCount:      "failureCount",
	SuccessCount:      "successCount",
	LastFailureAt:     "lastFailureAt",
	LastSuccessAt:     "lastSuccessAt",
	OpenedAt:          "openedAt",
	OpenUntil:         "openUntil",
	HalfOpenAt:        "halfOpenAt",
	Reason:            "reason",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformProviderCircuitBreakersDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderCircuitBreakersDTOFieldName.Id):
		return string(model.ProviderCircuitBreakersDBFieldName.Id), true

	case string(ProviderCircuitBreakersDTOFieldName.ProviderAccountId):
		return string(model.ProviderCircuitBreakersDBFieldName.ProviderAccountId), true

	case string(ProviderCircuitBreakersDTOFieldName.MethodCode):
		return string(model.ProviderCircuitBreakersDBFieldName.MethodCode), true

	case string(ProviderCircuitBreakersDTOFieldName.ChannelCode):
		return string(model.ProviderCircuitBreakersDBFieldName.ChannelCode), true

	case string(ProviderCircuitBreakersDTOFieldName.Status):
		return string(model.ProviderCircuitBreakersDBFieldName.Status), true

	case string(ProviderCircuitBreakersDTOFieldName.FailureCount):
		return string(model.ProviderCircuitBreakersDBFieldName.FailureCount), true

	case string(ProviderCircuitBreakersDTOFieldName.SuccessCount):
		return string(model.ProviderCircuitBreakersDBFieldName.SuccessCount), true

	case string(ProviderCircuitBreakersDTOFieldName.LastFailureAt):
		return string(model.ProviderCircuitBreakersDBFieldName.LastFailureAt), true

	case string(ProviderCircuitBreakersDTOFieldName.LastSuccessAt):
		return string(model.ProviderCircuitBreakersDBFieldName.LastSuccessAt), true

	case string(ProviderCircuitBreakersDTOFieldName.OpenedAt):
		return string(model.ProviderCircuitBreakersDBFieldName.OpenedAt), true

	case string(ProviderCircuitBreakersDTOFieldName.OpenUntil):
		return string(model.ProviderCircuitBreakersDBFieldName.OpenUntil), true

	case string(ProviderCircuitBreakersDTOFieldName.HalfOpenAt):
		return string(model.ProviderCircuitBreakersDBFieldName.HalfOpenAt), true

	case string(ProviderCircuitBreakersDTOFieldName.Reason):
		return string(model.ProviderCircuitBreakersDBFieldName.Reason), true

	case string(ProviderCircuitBreakersDTOFieldName.Metadata):
		return string(model.ProviderCircuitBreakersDBFieldName.Metadata), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaCreatedAt):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedAt), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaCreatedBy):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedBy), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedAt), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedBy), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaDeletedAt):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedAt), true

	case string(ProviderCircuitBreakersDTOFieldName.MetaDeletedBy):
		return string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderCircuitBreakersFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderCircuitBreakersBaseFilterField(field string) bool {
	spec, found := model.NewProviderCircuitBreakersFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderCircuitBreakersProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderCircuitBreakersProjectionOutputPath(path string) error {
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

func transformProviderCircuitBreakersFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderCircuitBreakersDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderCircuitBreakersFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderCircuitBreakersFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderCircuitBreakersDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderCircuitBreakersBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderCircuitBreakersProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderCircuitBreakersProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderCircuitBreakersDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderCircuitBreakersDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderCircuitBreakersFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderCircuitBreakersFilter(filter *model.Filter) {
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
			Field: string(ProviderCircuitBreakersDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderCircuitBreakersSelectableResponse map[string]interface{}
type ProviderCircuitBreakersSelectableListResponse []*ProviderCircuitBreakersSelectableResponse

func assignProviderCircuitBreakersNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderCircuitBreakersSelectableValue(out ProviderCircuitBreakersSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderCircuitBreakersNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderCircuitBreakersSelectableResponse(providerCircuitBreakers model.ProviderCircuitBreakers, filter model.Filter) ProviderCircuitBreakersSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderCircuitBreakersDBFieldName.Id),
			string(model.ProviderCircuitBreakersDBFieldName.ProviderAccountId),
			string(model.ProviderCircuitBreakersDBFieldName.MethodCode),
			string(model.ProviderCircuitBreakersDBFieldName.ChannelCode),
			string(model.ProviderCircuitBreakersDBFieldName.Status),
			string(model.ProviderCircuitBreakersDBFieldName.FailureCount),
			string(model.ProviderCircuitBreakersDBFieldName.SuccessCount),
			string(model.ProviderCircuitBreakersDBFieldName.LastFailureAt),
			string(model.ProviderCircuitBreakersDBFieldName.LastSuccessAt),
			string(model.ProviderCircuitBreakersDBFieldName.OpenedAt),
			string(model.ProviderCircuitBreakersDBFieldName.OpenUntil),
			string(model.ProviderCircuitBreakersDBFieldName.HalfOpenAt),
			string(model.ProviderCircuitBreakersDBFieldName.Reason),
			string(model.ProviderCircuitBreakersDBFieldName.Metadata),
			string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedAt),
			string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedBy),
			string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedAt),
			string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedBy),
			string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedAt),
			string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedBy),
		)
	}
	providerCircuitBreakersSelectableResponse := ProviderCircuitBreakersSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderCircuitBreakersDBFieldName.Id):
			key := string(ProviderCircuitBreakersDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.Id, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.ProviderAccountId):
			key := string(ProviderCircuitBreakersDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.ProviderAccountId, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MethodCode):
			key := string(ProviderCircuitBreakersDTOFieldName.MethodCode)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MethodCode.String, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.ChannelCode):
			key := string(ProviderCircuitBreakersDTOFieldName.ChannelCode)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.ChannelCode.String, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.Status):
			key := string(ProviderCircuitBreakersDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, model.CircuitStatus(providerCircuitBreakers.Status), explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.FailureCount):
			key := string(ProviderCircuitBreakersDTOFieldName.FailureCount)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.FailureCount, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.SuccessCount):
			key := string(ProviderCircuitBreakersDTOFieldName.SuccessCount)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.SuccessCount, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.LastFailureAt):
			key := string(ProviderCircuitBreakersDTOFieldName.LastFailureAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.LastFailureAt.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.LastSuccessAt):
			key := string(ProviderCircuitBreakersDTOFieldName.LastSuccessAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.LastSuccessAt.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.OpenedAt):
			key := string(ProviderCircuitBreakersDTOFieldName.OpenedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.OpenedAt.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.OpenUntil):
			key := string(ProviderCircuitBreakersDTOFieldName.OpenUntil)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.OpenUntil.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.HalfOpenAt):
			key := string(ProviderCircuitBreakersDTOFieldName.HalfOpenAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.HalfOpenAt.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.Reason):
			key := string(ProviderCircuitBreakersDTOFieldName.Reason)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.Reason.String, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.Metadata):
			key := string(ProviderCircuitBreakersDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.Metadata, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedAt):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaCreatedAt, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaCreatedBy):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaCreatedBy, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedAt):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaUpdatedBy):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaUpdatedBy, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedAt):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderCircuitBreakersDBFieldName.MetaDeletedBy):
			key := string(ProviderCircuitBreakersDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderCircuitBreakersSelectableValue(providerCircuitBreakersSelectableResponse, key, providerCircuitBreakers.MetaDeletedBy, explicitAlias)

		}
	}
	return providerCircuitBreakersSelectableResponse
}

func NewProviderCircuitBreakersListResponseFromFilterResult(result []model.ProviderCircuitBreakersFilterResult, filter model.Filter) ProviderCircuitBreakersSelectableListResponse {
	dtoProviderCircuitBreakersListResponse := ProviderCircuitBreakersSelectableListResponse{}
	for _, row := range result {
		dtoProviderCircuitBreakersResponse := NewProviderCircuitBreakersSelectableResponse(row.ProviderCircuitBreakers, filter)
		dtoProviderCircuitBreakersListResponse = append(dtoProviderCircuitBreakersListResponse, &dtoProviderCircuitBreakersResponse)
	}
	return dtoProviderCircuitBreakersListResponse
}

type ProviderCircuitBreakersFilterResponse struct {
	Metadata Metadata                                      `json:"metadata"`
	Data     ProviderCircuitBreakersSelectableListResponse `json:"data"`
}

func reverseProviderCircuitBreakersFilterResults(result []model.ProviderCircuitBreakersFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderCircuitBreakersFilterResponse(result []model.ProviderCircuitBreakersFilterResult, filter model.Filter) (resp ProviderCircuitBreakersFilterResponse) {
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
			reverseProviderCircuitBreakersFilterResults(dataResult)
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

	resp.Data = NewProviderCircuitBreakersListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderCircuitBreakersCreateRequest struct {
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	MethodCode        string              `json:"methodCode"`
	ChannelCode       string              `json:"channelCode"`
	Status            model.CircuitStatus `json:"status" example:"closed" enums:"closed,open,half_open"`
	FailureCount      int                 `json:"failureCount"`
	SuccessCount      int                 `json:"successCount"`
	LastFailureAt     time.Time           `json:"lastFailureAt"`
	LastSuccessAt     time.Time           `json:"lastSuccessAt"`
	OpenedAt          time.Time           `json:"openedAt"`
	OpenUntil         time.Time           `json:"openUntil"`
	HalfOpenAt        time.Time           `json:"halfOpenAt"`
	Reason            string              `json:"reason"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d *ProviderCircuitBreakersCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderCircuitBreakersCreateRequest) ToModel() model.ProviderCircuitBreakers {
	id, _ := uuid.NewV4()
	return model.ProviderCircuitBreakers{
		Id:                id,
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		Status:            d.Status,
		FailureCount:      d.FailureCount,
		SuccessCount:      d.SuccessCount,
		LastFailureAt:     null.TimeFrom(d.LastFailureAt),
		LastSuccessAt:     null.TimeFrom(d.LastSuccessAt),
		OpenedAt:          null.TimeFrom(d.OpenedAt),
		OpenUntil:         null.TimeFrom(d.OpenUntil),
		HalfOpenAt:        null.TimeFrom(d.HalfOpenAt),
		Reason:            null.StringFrom(d.Reason),
		Metadata:          d.Metadata,
	}
}

type ProviderCircuitBreakersListCreateRequest []*ProviderCircuitBreakersCreateRequest

func (d ProviderCircuitBreakersListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerCircuitBreakers := range d {
		err = validator.Struct(providerCircuitBreakers)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderCircuitBreakersListCreateRequest) ToModelList() []model.ProviderCircuitBreakers {
	out := make([]model.ProviderCircuitBreakers, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderCircuitBreakersUpdateRequest struct {
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	MethodCode        string              `json:"methodCode"`
	ChannelCode       string              `json:"channelCode"`
	Status            model.CircuitStatus `json:"status" example:"closed" enums:"closed,open,half_open"`
	FailureCount      int                 `json:"failureCount"`
	SuccessCount      int                 `json:"successCount"`
	LastFailureAt     time.Time           `json:"lastFailureAt"`
	LastSuccessAt     time.Time           `json:"lastSuccessAt"`
	OpenedAt          time.Time           `json:"openedAt"`
	OpenUntil         time.Time           `json:"openUntil"`
	HalfOpenAt        time.Time           `json:"halfOpenAt"`
	Reason            string              `json:"reason"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d *ProviderCircuitBreakersUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderCircuitBreakersUpdateRequest) ToModel() model.ProviderCircuitBreakers {
	return model.ProviderCircuitBreakers{
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		Status:            d.Status,
		FailureCount:      d.FailureCount,
		SuccessCount:      d.SuccessCount,
		LastFailureAt:     null.TimeFrom(d.LastFailureAt),
		LastSuccessAt:     null.TimeFrom(d.LastSuccessAt),
		OpenedAt:          null.TimeFrom(d.OpenedAt),
		OpenUntil:         null.TimeFrom(d.OpenUntil),
		HalfOpenAt:        null.TimeFrom(d.HalfOpenAt),
		Reason:            null.StringFrom(d.Reason),
		Metadata:          d.Metadata,
	}
}

type ProviderCircuitBreakersBulkUpdateRequest struct {
	Id                uuid.UUID           `json:"id"`
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	MethodCode        string              `json:"methodCode"`
	ChannelCode       string              `json:"channelCode"`
	Status            model.CircuitStatus `json:"status" example:"closed" enums:"closed,open,half_open"`
	FailureCount      int                 `json:"failureCount"`
	SuccessCount      int                 `json:"successCount"`
	LastFailureAt     time.Time           `json:"lastFailureAt"`
	LastSuccessAt     time.Time           `json:"lastSuccessAt"`
	OpenedAt          time.Time           `json:"openedAt"`
	OpenUntil         time.Time           `json:"openUntil"`
	HalfOpenAt        time.Time           `json:"halfOpenAt"`
	Reason            string              `json:"reason"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d ProviderCircuitBreakersBulkUpdateRequest) PrimaryID() ProviderCircuitBreakersPrimaryID {
	return ProviderCircuitBreakersPrimaryID{
		Id: d.Id,
	}
}

type ProviderCircuitBreakersListBulkUpdateRequest []*ProviderCircuitBreakersBulkUpdateRequest

func (d ProviderCircuitBreakersListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerCircuitBreakers := range d {
		err = validator.Struct(providerCircuitBreakers)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderCircuitBreakersBulkUpdateRequest) ToModel() model.ProviderCircuitBreakers {
	return model.ProviderCircuitBreakers{
		Id:                d.Id,
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		Status:            d.Status,
		FailureCount:      d.FailureCount,
		SuccessCount:      d.SuccessCount,
		LastFailureAt:     null.TimeFrom(d.LastFailureAt),
		LastSuccessAt:     null.TimeFrom(d.LastSuccessAt),
		OpenedAt:          null.TimeFrom(d.OpenedAt),
		OpenUntil:         null.TimeFrom(d.OpenUntil),
		HalfOpenAt:        null.TimeFrom(d.HalfOpenAt),
		Reason:            null.StringFrom(d.Reason),
		Metadata:          d.Metadata,
	}
}

type ProviderCircuitBreakersResponse struct {
	Id                uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID           `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodCode        string              `json:"methodCode"`
	ChannelCode       string              `json:"channelCode"`
	Status            model.CircuitStatus `json:"status" validate:"oneof=closed open half_open" enums:"closed,open,half_open"`
	FailureCount      int                 `json:"failureCount" example:"1"`
	SuccessCount      int                 `json:"successCount" example:"1"`
	LastFailureAt     time.Time           `json:"lastFailureAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LastSuccessAt     time.Time           `json:"lastSuccessAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	OpenedAt          time.Time           `json:"openedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	OpenUntil         time.Time           `json:"openUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	HalfOpenAt        time.Time           `json:"halfOpenAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Reason            string              `json:"reason"`
	Metadata          json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewProviderCircuitBreakersResponse(providerCircuitBreakers model.ProviderCircuitBreakers) ProviderCircuitBreakersResponse {
	return ProviderCircuitBreakersResponse{
		Id:                providerCircuitBreakers.Id,
		ProviderAccountId: providerCircuitBreakers.ProviderAccountId,
		MethodCode:        providerCircuitBreakers.MethodCode.String,
		ChannelCode:       providerCircuitBreakers.ChannelCode.String,
		Status:            model.CircuitStatus(providerCircuitBreakers.Status),
		FailureCount:      providerCircuitBreakers.FailureCount,
		SuccessCount:      providerCircuitBreakers.SuccessCount,
		LastFailureAt:     providerCircuitBreakers.LastFailureAt.Time,
		LastSuccessAt:     providerCircuitBreakers.LastSuccessAt.Time,
		OpenedAt:          providerCircuitBreakers.OpenedAt.Time,
		OpenUntil:         providerCircuitBreakers.OpenUntil.Time,
		HalfOpenAt:        providerCircuitBreakers.HalfOpenAt.Time,
		Reason:            providerCircuitBreakers.Reason.String,
		Metadata:          providerCircuitBreakers.Metadata,
	}
}

type ProviderCircuitBreakersListResponse []*ProviderCircuitBreakersResponse

func NewProviderCircuitBreakersListResponse(providerCircuitBreakersList model.ProviderCircuitBreakersList) ProviderCircuitBreakersListResponse {
	dtoProviderCircuitBreakersListResponse := ProviderCircuitBreakersListResponse{}
	for _, providerCircuitBreakers := range providerCircuitBreakersList {
		dtoProviderCircuitBreakersResponse := NewProviderCircuitBreakersResponse(*providerCircuitBreakers)
		dtoProviderCircuitBreakersListResponse = append(dtoProviderCircuitBreakersListResponse, &dtoProviderCircuitBreakersResponse)
	}
	return dtoProviderCircuitBreakersListResponse
}

type ProviderCircuitBreakersPrimaryIDList []ProviderCircuitBreakersPrimaryID

func (d ProviderCircuitBreakersPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerCircuitBreakers := range d {
		err = validator.Struct(providerCircuitBreakers)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderCircuitBreakersPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderCircuitBreakersPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderCircuitBreakersPrimaryID) ToModel() model.ProviderCircuitBreakersPrimaryID {
	return model.ProviderCircuitBreakersPrimaryID{
		Id: d.Id,
	}
}

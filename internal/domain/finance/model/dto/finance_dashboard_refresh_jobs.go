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

type FinanceDashboardRefreshJobsDTOFieldNameType string

type financeDashboardRefreshJobsDTOFieldName struct {
	Id             FinanceDashboardRefreshJobsDTOFieldNameType
	JobKey         FinanceDashboardRefreshJobsDTOFieldNameType
	RefreshScope   FinanceDashboardRefreshJobsDTOFieldNameType
	ScopeRef       FinanceDashboardRefreshJobsDTOFieldNameType
	CurrencyCode   FinanceDashboardRefreshJobsDTOFieldNameType
	IdempotencyKey FinanceDashboardRefreshJobsDTOFieldNameType
	RefreshStatus  FinanceDashboardRefreshJobsDTOFieldNameType
	RequestedAt    FinanceDashboardRefreshJobsDTOFieldNameType
	StartedAt      FinanceDashboardRefreshJobsDTOFieldNameType
	FinishedAt     FinanceDashboardRefreshJobsDTOFieldNameType
	ErrorCode      FinanceDashboardRefreshJobsDTOFieldNameType
	ErrorDetail    FinanceDashboardRefreshJobsDTOFieldNameType
	Metadata       FinanceDashboardRefreshJobsDTOFieldNameType
	MetaCreatedAt  FinanceDashboardRefreshJobsDTOFieldNameType
	MetaCreatedBy  FinanceDashboardRefreshJobsDTOFieldNameType
	MetaUpdatedAt  FinanceDashboardRefreshJobsDTOFieldNameType
	MetaUpdatedBy  FinanceDashboardRefreshJobsDTOFieldNameType
	MetaDeletedAt  FinanceDashboardRefreshJobsDTOFieldNameType
	MetaDeletedBy  FinanceDashboardRefreshJobsDTOFieldNameType
}

var FinanceDashboardRefreshJobsDTOFieldName = financeDashboardRefreshJobsDTOFieldName{
	Id:             "id",
	JobKey:         "jobKey",
	RefreshScope:   "refreshScope",
	ScopeRef:       "scopeRef",
	CurrencyCode:   "currencyCode",
	IdempotencyKey: "idempotencyKey",
	RefreshStatus:  "refreshStatus",
	RequestedAt:    "requestedAt",
	StartedAt:      "startedAt",
	FinishedAt:     "finishedAt",
	ErrorCode:      "errorCode",
	ErrorDetail:    "errorDetail",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformFinanceDashboardRefreshJobsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceDashboardRefreshJobsDTOFieldName.Id):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.Id), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.JobKey):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.JobKey), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.RefreshScope):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshScope), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.ScopeRef):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.ScopeRef), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.CurrencyCode):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.CurrencyCode), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.IdempotencyKey):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.IdempotencyKey), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.RefreshStatus):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshStatus), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.RequestedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.RequestedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.StartedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.StartedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.FinishedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.FinishedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.ErrorCode):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorCode), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.ErrorDetail):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorDetail), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.Metadata):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.Metadata), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedBy), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedBy), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedAt), true

	case string(FinanceDashboardRefreshJobsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceDashboardRefreshJobsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceDashboardRefreshJobsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceDashboardRefreshJobsProjectionOutputPath(path string) error {
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

func transformFinanceDashboardRefreshJobsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceDashboardRefreshJobsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceDashboardRefreshJobsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceDashboardRefreshJobsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceDashboardRefreshJobsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceDashboardRefreshJobsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceDashboardRefreshJobsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceDashboardRefreshJobsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceDashboardRefreshJobsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceDashboardRefreshJobsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceDashboardRefreshJobsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceDashboardRefreshJobsFilter(filter *model.Filter) {
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
			Field: string(FinanceDashboardRefreshJobsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceDashboardRefreshJobsSelectableResponse map[string]interface{}
type FinanceDashboardRefreshJobsSelectableListResponse []*FinanceDashboardRefreshJobsSelectableResponse

func assignFinanceDashboardRefreshJobsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceDashboardRefreshJobsSelectableValue(out FinanceDashboardRefreshJobsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceDashboardRefreshJobsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceDashboardRefreshJobsSelectableResponse(financeDashboardRefreshJobs model.FinanceDashboardRefreshJobs, filter model.Filter) FinanceDashboardRefreshJobsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceDashboardRefreshJobsDBFieldName.Id),
			string(model.FinanceDashboardRefreshJobsDBFieldName.JobKey),
			string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshScope),
			string(model.FinanceDashboardRefreshJobsDBFieldName.ScopeRef),
			string(model.FinanceDashboardRefreshJobsDBFieldName.CurrencyCode),
			string(model.FinanceDashboardRefreshJobsDBFieldName.IdempotencyKey),
			string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshStatus),
			string(model.FinanceDashboardRefreshJobsDBFieldName.RequestedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.StartedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.FinishedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorCode),
			string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorDetail),
			string(model.FinanceDashboardRefreshJobsDBFieldName.Metadata),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedBy),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedBy),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedAt),
			string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedBy),
		)
	}
	financeDashboardRefreshJobsSelectableResponse := FinanceDashboardRefreshJobsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceDashboardRefreshJobsDBFieldName.Id):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.Id, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.JobKey):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.JobKey)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.JobKey, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshScope):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.RefreshScope)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.RefreshScope, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.ScopeRef):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.ScopeRef)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.ScopeRef.UUID, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.CurrencyCode):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.CurrencyCode.String, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.IdempotencyKey):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.IdempotencyKey.String, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.RefreshStatus):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.RefreshStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, model.RefreshStatus(financeDashboardRefreshJobs.RefreshStatus), explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.RequestedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.RequestedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.RequestedAt, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.StartedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.StartedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.StartedAt.Time, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.FinishedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.FinishedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.FinishedAt.Time, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorCode):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.ErrorCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.ErrorCode.String, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.ErrorDetail):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.ErrorDetail)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.ErrorDetail.String, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.Metadata):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.Metadata, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaCreatedAt, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaCreatedBy):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaCreatedBy, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedBy):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedAt):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceDashboardRefreshJobsDBFieldName.MetaDeletedBy):
			key := string(FinanceDashboardRefreshJobsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceDashboardRefreshJobsSelectableValue(financeDashboardRefreshJobsSelectableResponse, key, financeDashboardRefreshJobs.MetaDeletedBy, explicitAlias)

		}
	}
	return financeDashboardRefreshJobsSelectableResponse
}

func NewFinanceDashboardRefreshJobsListResponseFromFilterResult(result []model.FinanceDashboardRefreshJobsFilterResult, filter model.Filter) FinanceDashboardRefreshJobsSelectableListResponse {
	dtoFinanceDashboardRefreshJobsListResponse := FinanceDashboardRefreshJobsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceDashboardRefreshJobsResponse := NewFinanceDashboardRefreshJobsSelectableResponse(row.FinanceDashboardRefreshJobs, filter)
		dtoFinanceDashboardRefreshJobsListResponse = append(dtoFinanceDashboardRefreshJobsListResponse, &dtoFinanceDashboardRefreshJobsResponse)
	}
	return dtoFinanceDashboardRefreshJobsListResponse
}

type FinanceDashboardRefreshJobsFilterResponse struct {
	Metadata Metadata                                          `json:"metadata"`
	Data     FinanceDashboardRefreshJobsSelectableListResponse `json:"data"`
}

func reverseFinanceDashboardRefreshJobsFilterResults(result []model.FinanceDashboardRefreshJobsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceDashboardRefreshJobsFilterResponse(result []model.FinanceDashboardRefreshJobsFilterResult, filter model.Filter) (resp FinanceDashboardRefreshJobsFilterResponse) {
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
			reverseFinanceDashboardRefreshJobsFilterResults(dataResult)
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

	resp.Data = NewFinanceDashboardRefreshJobsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceDashboardRefreshJobsCreateRequest struct {
	JobKey         string              `json:"jobKey"`
	RefreshScope   string              `json:"refreshScope"`
	ScopeRef       uuid.UUID           `json:"scopeRef"`
	CurrencyCode   string              `json:"currencyCode"`
	IdempotencyKey string              `json:"idempotencyKey"`
	RefreshStatus  model.RefreshStatus `json:"refreshStatus" example:"pending" enums:"pending,running,succeeded,failed"`
	RequestedAt    time.Time           `json:"requestedAt"`
	StartedAt      time.Time           `json:"startedAt"`
	FinishedAt     time.Time           `json:"finishedAt"`
	ErrorCode      string              `json:"errorCode"`
	ErrorDetail    string              `json:"errorDetail"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d *FinanceDashboardRefreshJobsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceDashboardRefreshJobsCreateRequest) ToModel() model.FinanceDashboardRefreshJobs {
	id, _ := uuid.NewV4()
	return model.FinanceDashboardRefreshJobs{
		Id:             id,
		JobKey:         d.JobKey,
		RefreshScope:   d.RefreshScope,
		ScopeRef:       nuuid.From(d.ScopeRef),
		CurrencyCode:   null.StringFrom(d.CurrencyCode),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		RefreshStatus:  d.RefreshStatus,
		RequestedAt:    d.RequestedAt,
		StartedAt:      null.TimeFrom(d.StartedAt),
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceDashboardRefreshJobsListCreateRequest []*FinanceDashboardRefreshJobsCreateRequest

func (d FinanceDashboardRefreshJobsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeDashboardRefreshJobs := range d {
		err = validator.Struct(financeDashboardRefreshJobs)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceDashboardRefreshJobsListCreateRequest) ToModelList() []model.FinanceDashboardRefreshJobs {
	out := make([]model.FinanceDashboardRefreshJobs, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceDashboardRefreshJobsUpdateRequest struct {
	JobKey         string              `json:"jobKey"`
	RefreshScope   string              `json:"refreshScope"`
	ScopeRef       uuid.UUID           `json:"scopeRef"`
	CurrencyCode   string              `json:"currencyCode"`
	IdempotencyKey string              `json:"idempotencyKey"`
	RefreshStatus  model.RefreshStatus `json:"refreshStatus" example:"pending" enums:"pending,running,succeeded,failed"`
	RequestedAt    time.Time           `json:"requestedAt"`
	StartedAt      time.Time           `json:"startedAt"`
	FinishedAt     time.Time           `json:"finishedAt"`
	ErrorCode      string              `json:"errorCode"`
	ErrorDetail    string              `json:"errorDetail"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d *FinanceDashboardRefreshJobsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceDashboardRefreshJobsUpdateRequest) ToModel() model.FinanceDashboardRefreshJobs {
	return model.FinanceDashboardRefreshJobs{
		JobKey:         d.JobKey,
		RefreshScope:   d.RefreshScope,
		ScopeRef:       nuuid.From(d.ScopeRef),
		CurrencyCode:   null.StringFrom(d.CurrencyCode),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		RefreshStatus:  d.RefreshStatus,
		RequestedAt:    d.RequestedAt,
		StartedAt:      null.TimeFrom(d.StartedAt),
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceDashboardRefreshJobsBulkUpdateRequest struct {
	Id             uuid.UUID           `json:"id"`
	JobKey         string              `json:"jobKey"`
	RefreshScope   string              `json:"refreshScope"`
	ScopeRef       uuid.UUID           `json:"scopeRef"`
	CurrencyCode   string              `json:"currencyCode"`
	IdempotencyKey string              `json:"idempotencyKey"`
	RefreshStatus  model.RefreshStatus `json:"refreshStatus" example:"pending" enums:"pending,running,succeeded,failed"`
	RequestedAt    time.Time           `json:"requestedAt"`
	StartedAt      time.Time           `json:"startedAt"`
	FinishedAt     time.Time           `json:"finishedAt"`
	ErrorCode      string              `json:"errorCode"`
	ErrorDetail    string              `json:"errorDetail"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d FinanceDashboardRefreshJobsBulkUpdateRequest) PrimaryID() FinanceDashboardRefreshJobsPrimaryID {
	return FinanceDashboardRefreshJobsPrimaryID{
		Id: d.Id,
	}
}

type FinanceDashboardRefreshJobsListBulkUpdateRequest []*FinanceDashboardRefreshJobsBulkUpdateRequest

func (d FinanceDashboardRefreshJobsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeDashboardRefreshJobs := range d {
		err = validator.Struct(financeDashboardRefreshJobs)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceDashboardRefreshJobsBulkUpdateRequest) ToModel() model.FinanceDashboardRefreshJobs {
	return model.FinanceDashboardRefreshJobs{
		Id:             d.Id,
		JobKey:         d.JobKey,
		RefreshScope:   d.RefreshScope,
		ScopeRef:       nuuid.From(d.ScopeRef),
		CurrencyCode:   null.StringFrom(d.CurrencyCode),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		RefreshStatus:  d.RefreshStatus,
		RequestedAt:    d.RequestedAt,
		StartedAt:      null.TimeFrom(d.StartedAt),
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceDashboardRefreshJobsResponse struct {
	Id             uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	JobKey         string              `json:"jobKey" validate:"required"`
	RefreshScope   string              `json:"refreshScope" validate:"required"`
	ScopeRef       uuid.UUID           `json:"scopeRef" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode   string              `json:"currencyCode"`
	IdempotencyKey string              `json:"idempotencyKey"`
	RefreshStatus  model.RefreshStatus `json:"refreshStatus" validate:"oneof=pending running succeeded failed" enums:"pending,running,succeeded,failed"`
	RequestedAt    time.Time           `json:"requestedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	StartedAt      time.Time           `json:"startedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FinishedAt     time.Time           `json:"finishedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ErrorCode      string              `json:"errorCode"`
	ErrorDetail    string              `json:"errorDetail"`
	Metadata       json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewFinanceDashboardRefreshJobsResponse(financeDashboardRefreshJobs model.FinanceDashboardRefreshJobs) FinanceDashboardRefreshJobsResponse {
	return FinanceDashboardRefreshJobsResponse{
		Id:             financeDashboardRefreshJobs.Id,
		JobKey:         financeDashboardRefreshJobs.JobKey,
		RefreshScope:   financeDashboardRefreshJobs.RefreshScope,
		ScopeRef:       financeDashboardRefreshJobs.ScopeRef.UUID,
		CurrencyCode:   financeDashboardRefreshJobs.CurrencyCode.String,
		IdempotencyKey: financeDashboardRefreshJobs.IdempotencyKey.String,
		RefreshStatus:  model.RefreshStatus(financeDashboardRefreshJobs.RefreshStatus),
		RequestedAt:    financeDashboardRefreshJobs.RequestedAt,
		StartedAt:      financeDashboardRefreshJobs.StartedAt.Time,
		FinishedAt:     financeDashboardRefreshJobs.FinishedAt.Time,
		ErrorCode:      financeDashboardRefreshJobs.ErrorCode.String,
		ErrorDetail:    financeDashboardRefreshJobs.ErrorDetail.String,
		Metadata:       financeDashboardRefreshJobs.Metadata,
	}
}

type FinanceDashboardRefreshJobsListResponse []*FinanceDashboardRefreshJobsResponse

func NewFinanceDashboardRefreshJobsListResponse(financeDashboardRefreshJobsList model.FinanceDashboardRefreshJobsList) FinanceDashboardRefreshJobsListResponse {
	dtoFinanceDashboardRefreshJobsListResponse := FinanceDashboardRefreshJobsListResponse{}
	for _, financeDashboardRefreshJobs := range financeDashboardRefreshJobsList {
		dtoFinanceDashboardRefreshJobsResponse := NewFinanceDashboardRefreshJobsResponse(*financeDashboardRefreshJobs)
		dtoFinanceDashboardRefreshJobsListResponse = append(dtoFinanceDashboardRefreshJobsListResponse, &dtoFinanceDashboardRefreshJobsResponse)
	}
	return dtoFinanceDashboardRefreshJobsListResponse
}

type FinanceDashboardRefreshJobsPrimaryIDList []FinanceDashboardRefreshJobsPrimaryID

func (d FinanceDashboardRefreshJobsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeDashboardRefreshJobs := range d {
		err = validator.Struct(financeDashboardRefreshJobs)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceDashboardRefreshJobsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceDashboardRefreshJobsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceDashboardRefreshJobsPrimaryID) ToModel() model.FinanceDashboardRefreshJobsPrimaryID {
	return model.FinanceDashboardRefreshJobsPrimaryID{
		Id: d.Id,
	}
}

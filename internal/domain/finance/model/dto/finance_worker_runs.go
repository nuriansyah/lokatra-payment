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

type FinanceWorkerRunsDTOFieldNameType string

type financeWorkerRunsDTOFieldName struct {
	Id             FinanceWorkerRunsDTOFieldNameType
	WorkerName     FinanceWorkerRunsDTOFieldNameType
	RunKey         FinanceWorkerRunsDTOFieldNameType
	RunStatus      FinanceWorkerRunsDTOFieldNameType
	StartedAt      FinanceWorkerRunsDTOFieldNameType
	FinishedAt     FinanceWorkerRunsDTOFieldNameType
	ProcessedCount FinanceWorkerRunsDTOFieldNameType
	FailedCount    FinanceWorkerRunsDTOFieldNameType
	ErrorCode      FinanceWorkerRunsDTOFieldNameType
	ErrorDetail    FinanceWorkerRunsDTOFieldNameType
	Metadata       FinanceWorkerRunsDTOFieldNameType
	MetaCreatedAt  FinanceWorkerRunsDTOFieldNameType
	MetaCreatedBy  FinanceWorkerRunsDTOFieldNameType
	MetaUpdatedAt  FinanceWorkerRunsDTOFieldNameType
	MetaUpdatedBy  FinanceWorkerRunsDTOFieldNameType
	MetaDeletedAt  FinanceWorkerRunsDTOFieldNameType
	MetaDeletedBy  FinanceWorkerRunsDTOFieldNameType
}

var FinanceWorkerRunsDTOFieldName = financeWorkerRunsDTOFieldName{
	Id:             "id",
	WorkerName:     "workerName",
	RunKey:         "runKey",
	RunStatus:      "runStatus",
	StartedAt:      "startedAt",
	FinishedAt:     "finishedAt",
	ProcessedCount: "processedCount",
	FailedCount:    "failedCount",
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

func transformFinanceWorkerRunsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceWorkerRunsDTOFieldName.Id):
		return string(model.FinanceWorkerRunsDBFieldName.Id), true

	case string(FinanceWorkerRunsDTOFieldName.WorkerName):
		return string(model.FinanceWorkerRunsDBFieldName.WorkerName), true

	case string(FinanceWorkerRunsDTOFieldName.RunKey):
		return string(model.FinanceWorkerRunsDBFieldName.RunKey), true

	case string(FinanceWorkerRunsDTOFieldName.RunStatus):
		return string(model.FinanceWorkerRunsDBFieldName.RunStatus), true

	case string(FinanceWorkerRunsDTOFieldName.StartedAt):
		return string(model.FinanceWorkerRunsDBFieldName.StartedAt), true

	case string(FinanceWorkerRunsDTOFieldName.FinishedAt):
		return string(model.FinanceWorkerRunsDBFieldName.FinishedAt), true

	case string(FinanceWorkerRunsDTOFieldName.ProcessedCount):
		return string(model.FinanceWorkerRunsDBFieldName.ProcessedCount), true

	case string(FinanceWorkerRunsDTOFieldName.FailedCount):
		return string(model.FinanceWorkerRunsDBFieldName.FailedCount), true

	case string(FinanceWorkerRunsDTOFieldName.ErrorCode):
		return string(model.FinanceWorkerRunsDBFieldName.ErrorCode), true

	case string(FinanceWorkerRunsDTOFieldName.ErrorDetail):
		return string(model.FinanceWorkerRunsDBFieldName.ErrorDetail), true

	case string(FinanceWorkerRunsDTOFieldName.Metadata):
		return string(model.FinanceWorkerRunsDBFieldName.Metadata), true

	case string(FinanceWorkerRunsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceWorkerRunsDBFieldName.MetaCreatedAt), true

	case string(FinanceWorkerRunsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceWorkerRunsDBFieldName.MetaCreatedBy), true

	case string(FinanceWorkerRunsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedAt), true

	case string(FinanceWorkerRunsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedBy), true

	case string(FinanceWorkerRunsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceWorkerRunsDBFieldName.MetaDeletedAt), true

	case string(FinanceWorkerRunsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceWorkerRunsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceWorkerRunsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceWorkerRunsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceWorkerRunsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceWorkerRunsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceWorkerRunsProjectionOutputPath(path string) error {
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

func transformFinanceWorkerRunsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceWorkerRunsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceWorkerRunsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceWorkerRunsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceWorkerRunsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceWorkerRunsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceWorkerRunsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceWorkerRunsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceWorkerRunsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceWorkerRunsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceWorkerRunsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceWorkerRunsFilter(filter *model.Filter) {
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
			Field: string(FinanceWorkerRunsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceWorkerRunsSelectableResponse map[string]interface{}
type FinanceWorkerRunsSelectableListResponse []*FinanceWorkerRunsSelectableResponse

func assignFinanceWorkerRunsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceWorkerRunsSelectableValue(out FinanceWorkerRunsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceWorkerRunsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceWorkerRunsSelectableResponse(financeWorkerRuns model.FinanceWorkerRuns, filter model.Filter) FinanceWorkerRunsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceWorkerRunsDBFieldName.Id),
			string(model.FinanceWorkerRunsDBFieldName.WorkerName),
			string(model.FinanceWorkerRunsDBFieldName.RunKey),
			string(model.FinanceWorkerRunsDBFieldName.RunStatus),
			string(model.FinanceWorkerRunsDBFieldName.StartedAt),
			string(model.FinanceWorkerRunsDBFieldName.FinishedAt),
			string(model.FinanceWorkerRunsDBFieldName.ProcessedCount),
			string(model.FinanceWorkerRunsDBFieldName.FailedCount),
			string(model.FinanceWorkerRunsDBFieldName.ErrorCode),
			string(model.FinanceWorkerRunsDBFieldName.ErrorDetail),
			string(model.FinanceWorkerRunsDBFieldName.Metadata),
			string(model.FinanceWorkerRunsDBFieldName.MetaCreatedAt),
			string(model.FinanceWorkerRunsDBFieldName.MetaCreatedBy),
			string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedAt),
			string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedBy),
			string(model.FinanceWorkerRunsDBFieldName.MetaDeletedAt),
			string(model.FinanceWorkerRunsDBFieldName.MetaDeletedBy),
		)
	}
	financeWorkerRunsSelectableResponse := FinanceWorkerRunsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceWorkerRunsDBFieldName.Id):
			key := string(FinanceWorkerRunsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.Id, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.WorkerName):
			key := string(FinanceWorkerRunsDTOFieldName.WorkerName)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.WorkerName, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.RunKey):
			key := string(FinanceWorkerRunsDTOFieldName.RunKey)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.RunKey, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.RunStatus):
			key := string(FinanceWorkerRunsDTOFieldName.RunStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, model.FinanceWorkerRunsRunStatus(financeWorkerRuns.RunStatus), explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.StartedAt):
			key := string(FinanceWorkerRunsDTOFieldName.StartedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.StartedAt, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.FinishedAt):
			key := string(FinanceWorkerRunsDTOFieldName.FinishedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.FinishedAt.Time, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.ProcessedCount):
			key := string(FinanceWorkerRunsDTOFieldName.ProcessedCount)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.ProcessedCount, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.FailedCount):
			key := string(FinanceWorkerRunsDTOFieldName.FailedCount)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.FailedCount, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.ErrorCode):
			key := string(FinanceWorkerRunsDTOFieldName.ErrorCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.ErrorCode.String, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.ErrorDetail):
			key := string(FinanceWorkerRunsDTOFieldName.ErrorDetail)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.ErrorDetail.String, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.Metadata):
			key := string(FinanceWorkerRunsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.Metadata, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaCreatedAt):
			key := string(FinanceWorkerRunsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaCreatedAt, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaCreatedBy):
			key := string(FinanceWorkerRunsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaCreatedBy, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedAt):
			key := string(FinanceWorkerRunsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaUpdatedBy):
			key := string(FinanceWorkerRunsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaDeletedAt):
			key := string(FinanceWorkerRunsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceWorkerRunsDBFieldName.MetaDeletedBy):
			key := string(FinanceWorkerRunsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceWorkerRunsSelectableValue(financeWorkerRunsSelectableResponse, key, financeWorkerRuns.MetaDeletedBy, explicitAlias)

		}
	}
	return financeWorkerRunsSelectableResponse
}

func NewFinanceWorkerRunsListResponseFromFilterResult(result []model.FinanceWorkerRunsFilterResult, filter model.Filter) FinanceWorkerRunsSelectableListResponse {
	dtoFinanceWorkerRunsListResponse := FinanceWorkerRunsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceWorkerRunsResponse := NewFinanceWorkerRunsSelectableResponse(row.FinanceWorkerRuns, filter)
		dtoFinanceWorkerRunsListResponse = append(dtoFinanceWorkerRunsListResponse, &dtoFinanceWorkerRunsResponse)
	}
	return dtoFinanceWorkerRunsListResponse
}

type FinanceWorkerRunsFilterResponse struct {
	Metadata Metadata                                `json:"metadata"`
	Data     FinanceWorkerRunsSelectableListResponse `json:"data"`
}

func reverseFinanceWorkerRunsFilterResults(result []model.FinanceWorkerRunsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceWorkerRunsFilterResponse(result []model.FinanceWorkerRunsFilterResult, filter model.Filter) (resp FinanceWorkerRunsFilterResponse) {
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
			reverseFinanceWorkerRunsFilterResults(dataResult)
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

	resp.Data = NewFinanceWorkerRunsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceWorkerRunsCreateRequest struct {
	WorkerName     string                           `json:"workerName"`
	RunKey         string                           `json:"runKey"`
	RunStatus      model.FinanceWorkerRunsRunStatus `json:"runStatus" example:"running" enums:"running,succeeded,failed"`
	StartedAt      time.Time                        `json:"startedAt"`
	FinishedAt     time.Time                        `json:"finishedAt"`
	ProcessedCount int                              `json:"processedCount"`
	FailedCount    int                              `json:"failedCount"`
	ErrorCode      string                           `json:"errorCode"`
	ErrorDetail    string                           `json:"errorDetail"`
	Metadata       json.RawMessage                  `json:"metadata"`
}

func (d *FinanceWorkerRunsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceWorkerRunsCreateRequest) ToModel() model.FinanceWorkerRuns {
	id, _ := uuid.NewV4()
	return model.FinanceWorkerRuns{
		Id:             id,
		WorkerName:     d.WorkerName,
		RunKey:         d.RunKey,
		RunStatus:      d.RunStatus,
		StartedAt:      d.StartedAt,
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ProcessedCount: d.ProcessedCount,
		FailedCount:    d.FailedCount,
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceWorkerRunsListCreateRequest []*FinanceWorkerRunsCreateRequest

func (d FinanceWorkerRunsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeWorkerRuns := range d {
		err = validator.Struct(financeWorkerRuns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceWorkerRunsListCreateRequest) ToModelList() []model.FinanceWorkerRuns {
	out := make([]model.FinanceWorkerRuns, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceWorkerRunsUpdateRequest struct {
	WorkerName     string                           `json:"workerName"`
	RunKey         string                           `json:"runKey"`
	RunStatus      model.FinanceWorkerRunsRunStatus `json:"runStatus" example:"running" enums:"running,succeeded,failed"`
	StartedAt      time.Time                        `json:"startedAt"`
	FinishedAt     time.Time                        `json:"finishedAt"`
	ProcessedCount int                              `json:"processedCount"`
	FailedCount    int                              `json:"failedCount"`
	ErrorCode      string                           `json:"errorCode"`
	ErrorDetail    string                           `json:"errorDetail"`
	Metadata       json.RawMessage                  `json:"metadata"`
}

func (d *FinanceWorkerRunsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceWorkerRunsUpdateRequest) ToModel() model.FinanceWorkerRuns {
	return model.FinanceWorkerRuns{
		WorkerName:     d.WorkerName,
		RunKey:         d.RunKey,
		RunStatus:      d.RunStatus,
		StartedAt:      d.StartedAt,
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ProcessedCount: d.ProcessedCount,
		FailedCount:    d.FailedCount,
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceWorkerRunsBulkUpdateRequest struct {
	Id             uuid.UUID                        `json:"id"`
	WorkerName     string                           `json:"workerName"`
	RunKey         string                           `json:"runKey"`
	RunStatus      model.FinanceWorkerRunsRunStatus `json:"runStatus" example:"running" enums:"running,succeeded,failed"`
	StartedAt      time.Time                        `json:"startedAt"`
	FinishedAt     time.Time                        `json:"finishedAt"`
	ProcessedCount int                              `json:"processedCount"`
	FailedCount    int                              `json:"failedCount"`
	ErrorCode      string                           `json:"errorCode"`
	ErrorDetail    string                           `json:"errorDetail"`
	Metadata       json.RawMessage                  `json:"metadata"`
}

func (d FinanceWorkerRunsBulkUpdateRequest) PrimaryID() FinanceWorkerRunsPrimaryID {
	return FinanceWorkerRunsPrimaryID{
		Id: d.Id,
	}
}

type FinanceWorkerRunsListBulkUpdateRequest []*FinanceWorkerRunsBulkUpdateRequest

func (d FinanceWorkerRunsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeWorkerRuns := range d {
		err = validator.Struct(financeWorkerRuns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceWorkerRunsBulkUpdateRequest) ToModel() model.FinanceWorkerRuns {
	return model.FinanceWorkerRuns{
		Id:             d.Id,
		WorkerName:     d.WorkerName,
		RunKey:         d.RunKey,
		RunStatus:      d.RunStatus,
		StartedAt:      d.StartedAt,
		FinishedAt:     null.TimeFrom(d.FinishedAt),
		ProcessedCount: d.ProcessedCount,
		FailedCount:    d.FailedCount,
		ErrorCode:      null.StringFrom(d.ErrorCode),
		ErrorDetail:    null.StringFrom(d.ErrorDetail),
		Metadata:       d.Metadata,
	}
}

type FinanceWorkerRunsResponse struct {
	Id             uuid.UUID                        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	WorkerName     string                           `json:"workerName" validate:"required"`
	RunKey         string                           `json:"runKey" validate:"required"`
	RunStatus      model.FinanceWorkerRunsRunStatus `json:"runStatus" validate:"oneof=running succeeded failed" enums:"running,succeeded,failed"`
	StartedAt      time.Time                        `json:"startedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FinishedAt     time.Time                        `json:"finishedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ProcessedCount int                              `json:"processedCount" example:"1"`
	FailedCount    int                              `json:"failedCount" example:"1"`
	ErrorCode      string                           `json:"errorCode"`
	ErrorDetail    string                           `json:"errorDetail"`
	Metadata       json.RawMessage                  `json:"metadata" swaggertype:"object"`
}

func NewFinanceWorkerRunsResponse(financeWorkerRuns model.FinanceWorkerRuns) FinanceWorkerRunsResponse {
	return FinanceWorkerRunsResponse{
		Id:             financeWorkerRuns.Id,
		WorkerName:     financeWorkerRuns.WorkerName,
		RunKey:         financeWorkerRuns.RunKey,
		RunStatus:      model.FinanceWorkerRunsRunStatus(financeWorkerRuns.RunStatus),
		StartedAt:      financeWorkerRuns.StartedAt,
		FinishedAt:     financeWorkerRuns.FinishedAt.Time,
		ProcessedCount: financeWorkerRuns.ProcessedCount,
		FailedCount:    financeWorkerRuns.FailedCount,
		ErrorCode:      financeWorkerRuns.ErrorCode.String,
		ErrorDetail:    financeWorkerRuns.ErrorDetail.String,
		Metadata:       financeWorkerRuns.Metadata,
	}
}

type FinanceWorkerRunsListResponse []*FinanceWorkerRunsResponse

func NewFinanceWorkerRunsListResponse(financeWorkerRunsList model.FinanceWorkerRunsList) FinanceWorkerRunsListResponse {
	dtoFinanceWorkerRunsListResponse := FinanceWorkerRunsListResponse{}
	for _, financeWorkerRuns := range financeWorkerRunsList {
		dtoFinanceWorkerRunsResponse := NewFinanceWorkerRunsResponse(*financeWorkerRuns)
		dtoFinanceWorkerRunsListResponse = append(dtoFinanceWorkerRunsListResponse, &dtoFinanceWorkerRunsResponse)
	}
	return dtoFinanceWorkerRunsListResponse
}

type FinanceWorkerRunsPrimaryIDList []FinanceWorkerRunsPrimaryID

func (d FinanceWorkerRunsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeWorkerRuns := range d {
		err = validator.Struct(financeWorkerRuns)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceWorkerRunsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceWorkerRunsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceWorkerRunsPrimaryID) ToModel() model.FinanceWorkerRunsPrimaryID {
	return model.FinanceWorkerRunsPrimaryID{
		Id: d.Id,
	}
}

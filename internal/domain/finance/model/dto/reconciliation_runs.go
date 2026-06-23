package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ReconciliationRunsDTOFieldNameType string

type reconciliationRunsDTOFieldName struct {
	Id              ReconciliationRunsDTOFieldNameType
	ReconCode       ReconciliationRunsDTOFieldNameType
	ReconType       ReconciliationRunsDTOFieldNameType
	PeriodStart     ReconciliationRunsDTOFieldNameType
	PeriodEnd       ReconciliationRunsDTOFieldNameType
	CurrencyCode    ReconciliationRunsDTOFieldNameType
	RunStatus       ReconciliationRunsDTOFieldNameType
	ToleranceAmount ReconciliationRunsDTOFieldNameType
	ToleranceDays   ReconciliationRunsDTOFieldNameType
	StartedAt       ReconciliationRunsDTOFieldNameType
	CompletedAt     ReconciliationRunsDTOFieldNameType
	Summary         ReconciliationRunsDTOFieldNameType
	MetaCreatedAt   ReconciliationRunsDTOFieldNameType
	MetaCreatedBy   ReconciliationRunsDTOFieldNameType
	MetaUpdatedAt   ReconciliationRunsDTOFieldNameType
	MetaUpdatedBy   ReconciliationRunsDTOFieldNameType
	MetaDeletedAt   ReconciliationRunsDTOFieldNameType
	MetaDeletedBy   ReconciliationRunsDTOFieldNameType
}

var ReconciliationRunsDTOFieldName = reconciliationRunsDTOFieldName{
	Id:              "id",
	ReconCode:       "reconCode",
	ReconType:       "reconType",
	PeriodStart:     "periodStart",
	PeriodEnd:       "periodEnd",
	CurrencyCode:    "currencyCode",
	RunStatus:       "runStatus",
	ToleranceAmount: "toleranceAmount",
	ToleranceDays:   "toleranceDays",
	StartedAt:       "startedAt",
	CompletedAt:     "completedAt",
	Summary:         "summary",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformReconciliationRunsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReconciliationRunsDTOFieldName.Id):
		return string(model.ReconciliationRunsDBFieldName.Id), true

	case string(ReconciliationRunsDTOFieldName.ReconCode):
		return string(model.ReconciliationRunsDBFieldName.ReconCode), true

	case string(ReconciliationRunsDTOFieldName.ReconType):
		return string(model.ReconciliationRunsDBFieldName.ReconType), true

	case string(ReconciliationRunsDTOFieldName.PeriodStart):
		return string(model.ReconciliationRunsDBFieldName.PeriodStart), true

	case string(ReconciliationRunsDTOFieldName.PeriodEnd):
		return string(model.ReconciliationRunsDBFieldName.PeriodEnd), true

	case string(ReconciliationRunsDTOFieldName.CurrencyCode):
		return string(model.ReconciliationRunsDBFieldName.CurrencyCode), true

	case string(ReconciliationRunsDTOFieldName.RunStatus):
		return string(model.ReconciliationRunsDBFieldName.RunStatus), true

	case string(ReconciliationRunsDTOFieldName.ToleranceAmount):
		return string(model.ReconciliationRunsDBFieldName.ToleranceAmount), true

	case string(ReconciliationRunsDTOFieldName.ToleranceDays):
		return string(model.ReconciliationRunsDBFieldName.ToleranceDays), true

	case string(ReconciliationRunsDTOFieldName.StartedAt):
		return string(model.ReconciliationRunsDBFieldName.StartedAt), true

	case string(ReconciliationRunsDTOFieldName.CompletedAt):
		return string(model.ReconciliationRunsDBFieldName.CompletedAt), true

	case string(ReconciliationRunsDTOFieldName.Summary):
		return string(model.ReconciliationRunsDBFieldName.Summary), true

	case string(ReconciliationRunsDTOFieldName.MetaCreatedAt):
		return string(model.ReconciliationRunsDBFieldName.MetaCreatedAt), true

	case string(ReconciliationRunsDTOFieldName.MetaCreatedBy):
		return string(model.ReconciliationRunsDBFieldName.MetaCreatedBy), true

	case string(ReconciliationRunsDTOFieldName.MetaUpdatedAt):
		return string(model.ReconciliationRunsDBFieldName.MetaUpdatedAt), true

	case string(ReconciliationRunsDTOFieldName.MetaUpdatedBy):
		return string(model.ReconciliationRunsDBFieldName.MetaUpdatedBy), true

	case string(ReconciliationRunsDTOFieldName.MetaDeletedAt):
		return string(model.ReconciliationRunsDBFieldName.MetaDeletedAt), true

	case string(ReconciliationRunsDTOFieldName.MetaDeletedBy):
		return string(model.ReconciliationRunsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReconciliationRunsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReconciliationRunsBaseFilterField(field string) bool {
	spec, found := model.NewReconciliationRunsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReconciliationRunsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReconciliationRunsProjectionOutputPath(path string) error {
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

func transformReconciliationRunsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReconciliationRunsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReconciliationRunsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReconciliationRunsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReconciliationRunsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReconciliationRunsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReconciliationRunsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReconciliationRunsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReconciliationRunsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReconciliationRunsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReconciliationRunsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReconciliationRunsFilter(filter *model.Filter) {
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
			Field: string(ReconciliationRunsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReconciliationRunsSelectableResponse map[string]interface{}
type ReconciliationRunsSelectableListResponse []*ReconciliationRunsSelectableResponse

func assignReconciliationRunsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReconciliationRunsSelectableValue(out ReconciliationRunsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReconciliationRunsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReconciliationRunsSelectableResponse(reconciliationRuns model.ReconciliationRuns, filter model.Filter) ReconciliationRunsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReconciliationRunsDBFieldName.Id),
			string(model.ReconciliationRunsDBFieldName.ReconCode),
			string(model.ReconciliationRunsDBFieldName.ReconType),
			string(model.ReconciliationRunsDBFieldName.PeriodStart),
			string(model.ReconciliationRunsDBFieldName.PeriodEnd),
			string(model.ReconciliationRunsDBFieldName.CurrencyCode),
			string(model.ReconciliationRunsDBFieldName.RunStatus),
			string(model.ReconciliationRunsDBFieldName.ToleranceAmount),
			string(model.ReconciliationRunsDBFieldName.ToleranceDays),
			string(model.ReconciliationRunsDBFieldName.StartedAt),
			string(model.ReconciliationRunsDBFieldName.CompletedAt),
			string(model.ReconciliationRunsDBFieldName.Summary),
			string(model.ReconciliationRunsDBFieldName.MetaCreatedAt),
			string(model.ReconciliationRunsDBFieldName.MetaCreatedBy),
			string(model.ReconciliationRunsDBFieldName.MetaUpdatedAt),
			string(model.ReconciliationRunsDBFieldName.MetaUpdatedBy),
			string(model.ReconciliationRunsDBFieldName.MetaDeletedAt),
			string(model.ReconciliationRunsDBFieldName.MetaDeletedBy),
		)
	}
	reconciliationRunsSelectableResponse := ReconciliationRunsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReconciliationRunsDBFieldName.Id):
			key := string(ReconciliationRunsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.Id, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.ReconCode):
			key := string(ReconciliationRunsDTOFieldName.ReconCode)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.ReconCode, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.ReconType):
			key := string(ReconciliationRunsDTOFieldName.ReconType)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, model.ReconType(reconciliationRuns.ReconType), explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.PeriodStart):
			key := string(ReconciliationRunsDTOFieldName.PeriodStart)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.PeriodStart, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.PeriodEnd):
			key := string(ReconciliationRunsDTOFieldName.PeriodEnd)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.PeriodEnd, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.CurrencyCode):
			key := string(ReconciliationRunsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.CurrencyCode.String, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.RunStatus):
			key := string(ReconciliationRunsDTOFieldName.RunStatus)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, model.ReconciliationRunsRunStatus(reconciliationRuns.RunStatus), explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.ToleranceAmount):
			key := string(ReconciliationRunsDTOFieldName.ToleranceAmount)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.ToleranceAmount, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.ToleranceDays):
			key := string(ReconciliationRunsDTOFieldName.ToleranceDays)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.ToleranceDays, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.StartedAt):
			key := string(ReconciliationRunsDTOFieldName.StartedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.StartedAt, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.CompletedAt):
			key := string(ReconciliationRunsDTOFieldName.CompletedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.CompletedAt.Time, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.Summary):
			key := string(ReconciliationRunsDTOFieldName.Summary)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.Summary, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaCreatedAt):
			key := string(ReconciliationRunsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaCreatedAt, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaCreatedBy):
			key := string(ReconciliationRunsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaCreatedBy, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaUpdatedAt):
			key := string(ReconciliationRunsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaUpdatedAt, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaUpdatedBy):
			key := string(ReconciliationRunsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaUpdatedBy, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaDeletedAt):
			key := string(ReconciliationRunsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReconciliationRunsDBFieldName.MetaDeletedBy):
			key := string(ReconciliationRunsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationRunsSelectableValue(reconciliationRunsSelectableResponse, key, reconciliationRuns.MetaDeletedBy, explicitAlias)

		}
	}
	return reconciliationRunsSelectableResponse
}

func NewReconciliationRunsListResponseFromFilterResult(result []model.ReconciliationRunsFilterResult, filter model.Filter) ReconciliationRunsSelectableListResponse {
	dtoReconciliationRunsListResponse := ReconciliationRunsSelectableListResponse{}
	for _, row := range result {
		dtoReconciliationRunsResponse := NewReconciliationRunsSelectableResponse(row.ReconciliationRuns, filter)
		dtoReconciliationRunsListResponse = append(dtoReconciliationRunsListResponse, &dtoReconciliationRunsResponse)
	}
	return dtoReconciliationRunsListResponse
}

type ReconciliationRunsFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     ReconciliationRunsSelectableListResponse `json:"data"`
}

func reverseReconciliationRunsFilterResults(result []model.ReconciliationRunsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReconciliationRunsFilterResponse(result []model.ReconciliationRunsFilterResult, filter model.Filter) (resp ReconciliationRunsFilterResponse) {
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
			reverseReconciliationRunsFilterResults(dataResult)
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

	resp.Data = NewReconciliationRunsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReconciliationRunsCreateRequest struct {
	ReconCode       string                            `json:"reconCode"`
	ReconType       model.ReconType                   `json:"reconType" example:"provider_vs_ledger" enums:"provider_vs_ledger,bank_vs_ledger,provider_vs_bank,three_way"`
	PeriodStart     time.Time                         `json:"periodStart"`
	PeriodEnd       time.Time                         `json:"periodEnd"`
	CurrencyCode    string                            `json:"currencyCode"`
	RunStatus       model.ReconciliationRunsRunStatus `json:"runStatus" example:"running" enums:"running,completed,failed,cancelled"`
	ToleranceAmount decimal.Decimal                   `json:"toleranceAmount"`
	ToleranceDays   int                               `json:"toleranceDays"`
	StartedAt       time.Time                         `json:"startedAt"`
	CompletedAt     time.Time                         `json:"completedAt"`
	Summary         json.RawMessage                   `json:"summary"`
}

func (d *ReconciliationRunsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReconciliationRunsCreateRequest) ToModel() model.ReconciliationRuns {
	id, _ := uuid.NewV4()
	return model.ReconciliationRuns{
		Id:              id,
		ReconCode:       d.ReconCode,
		ReconType:       d.ReconType,
		PeriodStart:     d.PeriodStart,
		PeriodEnd:       d.PeriodEnd,
		CurrencyCode:    null.StringFrom(d.CurrencyCode),
		RunStatus:       d.RunStatus,
		ToleranceAmount: d.ToleranceAmount,
		ToleranceDays:   d.ToleranceDays,
		StartedAt:       d.StartedAt,
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		Summary:         d.Summary,
	}
}

type ReconciliationRunsListCreateRequest []*ReconciliationRunsCreateRequest

func (d ReconciliationRunsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationRuns := range d {
		err = validator.Struct(reconciliationRuns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationRunsListCreateRequest) ToModelList() []model.ReconciliationRuns {
	out := make([]model.ReconciliationRuns, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReconciliationRunsUpdateRequest struct {
	ReconCode       string                            `json:"reconCode"`
	ReconType       model.ReconType                   `json:"reconType" example:"provider_vs_ledger" enums:"provider_vs_ledger,bank_vs_ledger,provider_vs_bank,three_way"`
	PeriodStart     time.Time                         `json:"periodStart"`
	PeriodEnd       time.Time                         `json:"periodEnd"`
	CurrencyCode    string                            `json:"currencyCode"`
	RunStatus       model.ReconciliationRunsRunStatus `json:"runStatus" example:"running" enums:"running,completed,failed,cancelled"`
	ToleranceAmount decimal.Decimal                   `json:"toleranceAmount"`
	ToleranceDays   int                               `json:"toleranceDays"`
	StartedAt       time.Time                         `json:"startedAt"`
	CompletedAt     time.Time                         `json:"completedAt"`
	Summary         json.RawMessage                   `json:"summary"`
}

func (d *ReconciliationRunsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReconciliationRunsUpdateRequest) ToModel() model.ReconciliationRuns {
	return model.ReconciliationRuns{
		ReconCode:       d.ReconCode,
		ReconType:       d.ReconType,
		PeriodStart:     d.PeriodStart,
		PeriodEnd:       d.PeriodEnd,
		CurrencyCode:    null.StringFrom(d.CurrencyCode),
		RunStatus:       d.RunStatus,
		ToleranceAmount: d.ToleranceAmount,
		ToleranceDays:   d.ToleranceDays,
		StartedAt:       d.StartedAt,
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		Summary:         d.Summary,
	}
}

type ReconciliationRunsBulkUpdateRequest struct {
	Id              uuid.UUID                         `json:"id"`
	ReconCode       string                            `json:"reconCode"`
	ReconType       model.ReconType                   `json:"reconType" example:"provider_vs_ledger" enums:"provider_vs_ledger,bank_vs_ledger,provider_vs_bank,three_way"`
	PeriodStart     time.Time                         `json:"periodStart"`
	PeriodEnd       time.Time                         `json:"periodEnd"`
	CurrencyCode    string                            `json:"currencyCode"`
	RunStatus       model.ReconciliationRunsRunStatus `json:"runStatus" example:"running" enums:"running,completed,failed,cancelled"`
	ToleranceAmount decimal.Decimal                   `json:"toleranceAmount"`
	ToleranceDays   int                               `json:"toleranceDays"`
	StartedAt       time.Time                         `json:"startedAt"`
	CompletedAt     time.Time                         `json:"completedAt"`
	Summary         json.RawMessage                   `json:"summary"`
}

func (d ReconciliationRunsBulkUpdateRequest) PrimaryID() ReconciliationRunsPrimaryID {
	return ReconciliationRunsPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationRunsListBulkUpdateRequest []*ReconciliationRunsBulkUpdateRequest

func (d ReconciliationRunsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationRuns := range d {
		err = validator.Struct(reconciliationRuns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationRunsBulkUpdateRequest) ToModel() model.ReconciliationRuns {
	return model.ReconciliationRuns{
		Id:              d.Id,
		ReconCode:       d.ReconCode,
		ReconType:       d.ReconType,
		PeriodStart:     d.PeriodStart,
		PeriodEnd:       d.PeriodEnd,
		CurrencyCode:    null.StringFrom(d.CurrencyCode),
		RunStatus:       d.RunStatus,
		ToleranceAmount: d.ToleranceAmount,
		ToleranceDays:   d.ToleranceDays,
		StartedAt:       d.StartedAt,
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		Summary:         d.Summary,
	}
}

type ReconciliationRunsResponse struct {
	Id              uuid.UUID                         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReconCode       string                            `json:"reconCode" validate:"required"`
	ReconType       model.ReconType                   `json:"reconType" validate:"required,oneof=provider_vs_ledger bank_vs_ledger provider_vs_bank three_way" enums:"provider_vs_ledger,bank_vs_ledger,provider_vs_bank,three_way"`
	PeriodStart     time.Time                         `json:"periodStart" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PeriodEnd       time.Time                         `json:"periodEnd" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CurrencyCode    string                            `json:"currencyCode"`
	RunStatus       model.ReconciliationRunsRunStatus `json:"runStatus" validate:"oneof=running completed failed cancelled" enums:"running,completed,failed,cancelled"`
	ToleranceAmount decimal.Decimal                   `json:"toleranceAmount" format:"decimal" example:"100.50"`
	ToleranceDays   int                               `json:"toleranceDays" example:"1"`
	StartedAt       time.Time                         `json:"startedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CompletedAt     time.Time                         `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Summary         json.RawMessage                   `json:"summary" swaggertype:"object"`
}

func NewReconciliationRunsResponse(reconciliationRuns model.ReconciliationRuns) ReconciliationRunsResponse {
	return ReconciliationRunsResponse{
		Id:              reconciliationRuns.Id,
		ReconCode:       reconciliationRuns.ReconCode,
		ReconType:       model.ReconType(reconciliationRuns.ReconType),
		PeriodStart:     reconciliationRuns.PeriodStart,
		PeriodEnd:       reconciliationRuns.PeriodEnd,
		CurrencyCode:    reconciliationRuns.CurrencyCode.String,
		RunStatus:       model.ReconciliationRunsRunStatus(reconciliationRuns.RunStatus),
		ToleranceAmount: reconciliationRuns.ToleranceAmount,
		ToleranceDays:   reconciliationRuns.ToleranceDays,
		StartedAt:       reconciliationRuns.StartedAt,
		CompletedAt:     reconciliationRuns.CompletedAt.Time,
		Summary:         reconciliationRuns.Summary,
	}
}

type ReconciliationRunsListResponse []*ReconciliationRunsResponse

func NewReconciliationRunsListResponse(reconciliationRunsList model.ReconciliationRunsList) ReconciliationRunsListResponse {
	dtoReconciliationRunsListResponse := ReconciliationRunsListResponse{}
	for _, reconciliationRuns := range reconciliationRunsList {
		dtoReconciliationRunsResponse := NewReconciliationRunsResponse(*reconciliationRuns)
		dtoReconciliationRunsListResponse = append(dtoReconciliationRunsListResponse, &dtoReconciliationRunsResponse)
	}
	return dtoReconciliationRunsListResponse
}

type ReconciliationRunsPrimaryIDList []ReconciliationRunsPrimaryID

func (d ReconciliationRunsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationRuns := range d {
		err = validator.Struct(reconciliationRuns)
		if err != nil {
			return
		}
	}
	return nil
}

type ReconciliationRunsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReconciliationRunsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReconciliationRunsPrimaryID) ToModel() model.ReconciliationRunsPrimaryID {
	return model.ReconciliationRunsPrimaryID{
		Id: d.Id,
	}
}

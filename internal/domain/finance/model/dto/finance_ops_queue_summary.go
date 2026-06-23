package dto

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceOpsQueueSummaryDTOFieldNameType string

type financeOpsQueueSummaryDTOFieldName struct {
	QueueName    FinanceOpsQueueSummaryDTOFieldNameType
	QueueStatus  FinanceOpsQueueSummaryDTOFieldNameType
	ItemCount    FinanceOpsQueueSummaryDTOFieldNameType
	OldestItemAt FinanceOpsQueueSummaryDTOFieldNameType
	TotalAmount  FinanceOpsQueueSummaryDTOFieldNameType
	CurrencyCode FinanceOpsQueueSummaryDTOFieldNameType
	RefreshedAt  FinanceOpsQueueSummaryDTOFieldNameType
}

var FinanceOpsQueueSummaryDTOFieldName = financeOpsQueueSummaryDTOFieldName{
	QueueName:    "queueName",
	QueueStatus:  "queueStatus",
	ItemCount:    "itemCount",
	OldestItemAt: "oldestItemAt",
	TotalAmount:  "totalAmount",
	CurrencyCode: "currencyCode",
	RefreshedAt:  "refreshedAt",
}

func transformFinanceOpsQueueSummaryDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceOpsQueueSummaryDTOFieldName.QueueName):
		return string(model.FinanceOpsQueueSummaryDBFieldName.QueueName), true

	case string(FinanceOpsQueueSummaryDTOFieldName.QueueStatus):
		return string(model.FinanceOpsQueueSummaryDBFieldName.QueueStatus), true

	case string(FinanceOpsQueueSummaryDTOFieldName.ItemCount):
		return string(model.FinanceOpsQueueSummaryDBFieldName.ItemCount), true

	case string(FinanceOpsQueueSummaryDTOFieldName.OldestItemAt):
		return string(model.FinanceOpsQueueSummaryDBFieldName.OldestItemAt), true

	case string(FinanceOpsQueueSummaryDTOFieldName.TotalAmount):
		return string(model.FinanceOpsQueueSummaryDBFieldName.TotalAmount), true

	case string(FinanceOpsQueueSummaryDTOFieldName.CurrencyCode):
		return string(model.FinanceOpsQueueSummaryDBFieldName.CurrencyCode), true

	case string(FinanceOpsQueueSummaryDTOFieldName.RefreshedAt):
		return string(model.FinanceOpsQueueSummaryDBFieldName.RefreshedAt), true

	}
	if _, found := model.NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceOpsQueueSummaryBaseFilterField(field string) bool {
	spec, found := model.NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceOpsQueueSummaryProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceOpsQueueSummaryProjectionOutputPath(path string) error {
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

func transformFinanceOpsQueueSummaryFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceOpsQueueSummaryDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceOpsQueueSummaryFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceOpsQueueSummaryFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceOpsQueueSummaryDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceOpsQueueSummaryBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceOpsQueueSummaryProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceOpsQueueSummaryProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceOpsQueueSummaryDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceOpsQueueSummaryDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceOpsQueueSummaryFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceOpsQueueSummaryFilter(filter *model.Filter) {
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
			Field: string(FinanceOpsQueueSummaryDTOFieldName.QueueName),
			Order: model.SortAsc,
		})
	}
}

type FinanceOpsQueueSummarySelectableResponse map[string]interface{}
type FinanceOpsQueueSummarySelectableListResponse []*FinanceOpsQueueSummarySelectableResponse

func assignFinanceOpsQueueSummaryNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceOpsQueueSummarySelectableValue(out FinanceOpsQueueSummarySelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceOpsQueueSummaryNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceOpsQueueSummarySelectableResponse(financeOpsQueueSummary model.FinanceOpsQueueSummary, filter model.Filter) FinanceOpsQueueSummarySelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceOpsQueueSummaryDBFieldName.QueueName),
			string(model.FinanceOpsQueueSummaryDBFieldName.QueueStatus),
			string(model.FinanceOpsQueueSummaryDBFieldName.ItemCount),
			string(model.FinanceOpsQueueSummaryDBFieldName.OldestItemAt),
			string(model.FinanceOpsQueueSummaryDBFieldName.TotalAmount),
			string(model.FinanceOpsQueueSummaryDBFieldName.CurrencyCode),
			string(model.FinanceOpsQueueSummaryDBFieldName.RefreshedAt),
		)
	}
	financeOpsQueueSummarySelectableResponse := FinanceOpsQueueSummarySelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceOpsQueueSummaryDBFieldName.QueueName):
			key := string(FinanceOpsQueueSummaryDTOFieldName.QueueName)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.QueueName, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.QueueStatus):
			key := string(FinanceOpsQueueSummaryDTOFieldName.QueueStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.QueueStatus, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.ItemCount):
			key := string(FinanceOpsQueueSummaryDTOFieldName.ItemCount)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.ItemCount, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.OldestItemAt):
			key := string(FinanceOpsQueueSummaryDTOFieldName.OldestItemAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.OldestItemAt.Time, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.TotalAmount):
			key := string(FinanceOpsQueueSummaryDTOFieldName.TotalAmount)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.TotalAmount, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.CurrencyCode):
			key := string(FinanceOpsQueueSummaryDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.CurrencyCode, explicitAlias)

		case string(model.FinanceOpsQueueSummaryDBFieldName.RefreshedAt):
			key := string(FinanceOpsQueueSummaryDTOFieldName.RefreshedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceOpsQueueSummarySelectableValue(financeOpsQueueSummarySelectableResponse, key, financeOpsQueueSummary.RefreshedAt, explicitAlias)

		}
	}
	return financeOpsQueueSummarySelectableResponse
}

func NewFinanceOpsQueueSummaryListResponseFromFilterResult(result []model.FinanceOpsQueueSummaryFilterResult, filter model.Filter) FinanceOpsQueueSummarySelectableListResponse {
	dtoFinanceOpsQueueSummaryListResponse := FinanceOpsQueueSummarySelectableListResponse{}
	for _, row := range result {
		dtoFinanceOpsQueueSummaryResponse := NewFinanceOpsQueueSummarySelectableResponse(row.FinanceOpsQueueSummary, filter)
		dtoFinanceOpsQueueSummaryListResponse = append(dtoFinanceOpsQueueSummaryListResponse, &dtoFinanceOpsQueueSummaryResponse)
	}
	return dtoFinanceOpsQueueSummaryListResponse
}

type FinanceOpsQueueSummaryFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     FinanceOpsQueueSummarySelectableListResponse `json:"data"`
}

func reverseFinanceOpsQueueSummaryFilterResults(result []model.FinanceOpsQueueSummaryFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceOpsQueueSummaryFilterResponse(result []model.FinanceOpsQueueSummaryFilterResult, filter model.Filter) (resp FinanceOpsQueueSummaryFilterResponse) {
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
			reverseFinanceOpsQueueSummaryFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].QueueName
			resp.Metadata.PrevCursor = dataResult[0].QueueName
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

	resp.Data = NewFinanceOpsQueueSummaryListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceOpsQueueSummaryCreateRequest struct {
	QueueName    string          `json:"queueName"`
	QueueStatus  string          `json:"queueStatus"`
	ItemCount    int             `json:"itemCount"`
	OldestItemAt time.Time       `json:"oldestItemAt"`
	TotalAmount  decimal.Decimal `json:"totalAmount"`
	CurrencyCode string          `json:"currencyCode"`
	RefreshedAt  time.Time       `json:"refreshedAt"`
}

func (d *FinanceOpsQueueSummaryCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceOpsQueueSummaryCreateRequest) ToModel() model.FinanceOpsQueueSummary {
	return model.FinanceOpsQueueSummary{
		QueueName:    d.QueueName,
		QueueStatus:  d.QueueStatus,
		ItemCount:    d.ItemCount,
		OldestItemAt: null.TimeFrom(d.OldestItemAt),
		TotalAmount:  d.TotalAmount,
		CurrencyCode: d.CurrencyCode,
		RefreshedAt:  d.RefreshedAt,
	}
}

type FinanceOpsQueueSummaryListCreateRequest []*FinanceOpsQueueSummaryCreateRequest

func (d FinanceOpsQueueSummaryListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOpsQueueSummary := range d {
		err = validator.Struct(financeOpsQueueSummary)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceOpsQueueSummaryListCreateRequest) ToModelList() []model.FinanceOpsQueueSummary {
	out := make([]model.FinanceOpsQueueSummary, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceOpsQueueSummaryUpdateRequest struct {
	QueueName    string          `json:"queueName"`
	QueueStatus  string          `json:"queueStatus"`
	ItemCount    int             `json:"itemCount"`
	OldestItemAt time.Time       `json:"oldestItemAt"`
	TotalAmount  decimal.Decimal `json:"totalAmount"`
	CurrencyCode string          `json:"currencyCode"`
	RefreshedAt  time.Time       `json:"refreshedAt"`
}

func (d *FinanceOpsQueueSummaryUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceOpsQueueSummaryUpdateRequest) ToModel() model.FinanceOpsQueueSummary {
	return model.FinanceOpsQueueSummary{
		QueueName:    d.QueueName,
		QueueStatus:  d.QueueStatus,
		ItemCount:    d.ItemCount,
		OldestItemAt: null.TimeFrom(d.OldestItemAt),
		TotalAmount:  d.TotalAmount,
		CurrencyCode: d.CurrencyCode,
		RefreshedAt:  d.RefreshedAt,
	}
}

type FinanceOpsQueueSummaryBulkUpdateRequest struct {
	QueueName    string          `json:"queueName"`
	QueueStatus  string          `json:"queueStatus"`
	ItemCount    int             `json:"itemCount"`
	OldestItemAt time.Time       `json:"oldestItemAt"`
	TotalAmount  decimal.Decimal `json:"totalAmount"`
	CurrencyCode string          `json:"currencyCode"`
	RefreshedAt  time.Time       `json:"refreshedAt"`
}

func (d FinanceOpsQueueSummaryBulkUpdateRequest) PrimaryID() FinanceOpsQueueSummaryPrimaryID {
	return FinanceOpsQueueSummaryPrimaryID{
		QueueName:    d.QueueName,
		QueueStatus:  d.QueueStatus,
		CurrencyCode: d.CurrencyCode,
	}
}

type FinanceOpsQueueSummaryListBulkUpdateRequest []*FinanceOpsQueueSummaryBulkUpdateRequest

func (d FinanceOpsQueueSummaryListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOpsQueueSummary := range d {
		err = validator.Struct(financeOpsQueueSummary)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceOpsQueueSummaryBulkUpdateRequest) ToModel() model.FinanceOpsQueueSummary {
	return model.FinanceOpsQueueSummary{
		QueueName:    d.QueueName,
		QueueStatus:  d.QueueStatus,
		ItemCount:    d.ItemCount,
		OldestItemAt: null.TimeFrom(d.OldestItemAt),
		TotalAmount:  d.TotalAmount,
		CurrencyCode: d.CurrencyCode,
		RefreshedAt:  d.RefreshedAt,
	}
}

type FinanceOpsQueueSummaryResponse struct {
	QueueName    string          `json:"queueName" validate:"required"`
	QueueStatus  string          `json:"queueStatus" validate:"required"`
	ItemCount    int             `json:"itemCount" example:"1"`
	OldestItemAt time.Time       `json:"oldestItemAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	TotalAmount  decimal.Decimal `json:"totalAmount" format:"decimal" example:"100.50"`
	CurrencyCode string          `json:"currencyCode"`
	RefreshedAt  time.Time       `json:"refreshedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

func NewFinanceOpsQueueSummaryResponse(financeOpsQueueSummary model.FinanceOpsQueueSummary) FinanceOpsQueueSummaryResponse {
	return FinanceOpsQueueSummaryResponse{
		QueueName:    financeOpsQueueSummary.QueueName,
		QueueStatus:  financeOpsQueueSummary.QueueStatus,
		ItemCount:    financeOpsQueueSummary.ItemCount,
		OldestItemAt: financeOpsQueueSummary.OldestItemAt.Time,
		TotalAmount:  financeOpsQueueSummary.TotalAmount,
		CurrencyCode: financeOpsQueueSummary.CurrencyCode,
		RefreshedAt:  financeOpsQueueSummary.RefreshedAt,
	}
}

type FinanceOpsQueueSummaryListResponse []*FinanceOpsQueueSummaryResponse

func NewFinanceOpsQueueSummaryListResponse(financeOpsQueueSummaryList model.FinanceOpsQueueSummaryList) FinanceOpsQueueSummaryListResponse {
	dtoFinanceOpsQueueSummaryListResponse := FinanceOpsQueueSummaryListResponse{}
	for _, financeOpsQueueSummary := range financeOpsQueueSummaryList {
		dtoFinanceOpsQueueSummaryResponse := NewFinanceOpsQueueSummaryResponse(*financeOpsQueueSummary)
		dtoFinanceOpsQueueSummaryListResponse = append(dtoFinanceOpsQueueSummaryListResponse, &dtoFinanceOpsQueueSummaryResponse)
	}
	return dtoFinanceOpsQueueSummaryListResponse
}

type FinanceOpsQueueSummaryPrimaryIDList []FinanceOpsQueueSummaryPrimaryID

func (d FinanceOpsQueueSummaryPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeOpsQueueSummary := range d {
		err = validator.Struct(financeOpsQueueSummary)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceOpsQueueSummaryPrimaryID struct {
	QueueName   string `json:"queueName" validate:"required"`
	QueueStatus string `json:"queueStatus" validate:"required"`

	CurrencyCode string `json:"currencyCode"`
}

func (d *FinanceOpsQueueSummaryPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceOpsQueueSummaryPrimaryID) ToModel() model.FinanceOpsQueueSummaryPrimaryID {
	return model.FinanceOpsQueueSummaryPrimaryID{
		QueueName:    d.QueueName,
		QueueStatus:  d.QueueStatus,
		CurrencyCode: d.CurrencyCode,
	}
}

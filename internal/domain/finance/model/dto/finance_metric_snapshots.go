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

type FinanceMetricSnapshotsDTOFieldNameType string

type financeMetricSnapshotsDTOFieldName struct {
	Id            FinanceMetricSnapshotsDTOFieldNameType
	MetricName    FinanceMetricSnapshotsDTOFieldNameType
	MetricScope   FinanceMetricSnapshotsDTOFieldNameType
	ScopeRef      FinanceMetricSnapshotsDTOFieldNameType
	PeriodStart   FinanceMetricSnapshotsDTOFieldNameType
	PeriodEnd     FinanceMetricSnapshotsDTOFieldNameType
	MetricValue   FinanceMetricSnapshotsDTOFieldNameType
	MetricUnit    FinanceMetricSnapshotsDTOFieldNameType
	Dimensions    FinanceMetricSnapshotsDTOFieldNameType
	Metadata      FinanceMetricSnapshotsDTOFieldNameType
	MetaCreatedAt FinanceMetricSnapshotsDTOFieldNameType
	MetaCreatedBy FinanceMetricSnapshotsDTOFieldNameType
	MetaUpdatedAt FinanceMetricSnapshotsDTOFieldNameType
	MetaUpdatedBy FinanceMetricSnapshotsDTOFieldNameType
	MetaDeletedAt FinanceMetricSnapshotsDTOFieldNameType
	MetaDeletedBy FinanceMetricSnapshotsDTOFieldNameType
}

var FinanceMetricSnapshotsDTOFieldName = financeMetricSnapshotsDTOFieldName{
	Id:            "id",
	MetricName:    "metricName",
	MetricScope:   "metricScope",
	ScopeRef:      "scopeRef",
	PeriodStart:   "periodStart",
	PeriodEnd:     "periodEnd",
	MetricValue:   "metricValue",
	MetricUnit:    "metricUnit",
	Dimensions:    "dimensions",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformFinanceMetricSnapshotsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceMetricSnapshotsDTOFieldName.Id):
		return string(model.FinanceMetricSnapshotsDBFieldName.Id), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetricName):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetricName), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetricScope):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetricScope), true

	case string(FinanceMetricSnapshotsDTOFieldName.ScopeRef):
		return string(model.FinanceMetricSnapshotsDBFieldName.ScopeRef), true

	case string(FinanceMetricSnapshotsDTOFieldName.PeriodStart):
		return string(model.FinanceMetricSnapshotsDBFieldName.PeriodStart), true

	case string(FinanceMetricSnapshotsDTOFieldName.PeriodEnd):
		return string(model.FinanceMetricSnapshotsDBFieldName.PeriodEnd), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetricValue):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetricValue), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetricUnit):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetricUnit), true

	case string(FinanceMetricSnapshotsDTOFieldName.Dimensions):
		return string(model.FinanceMetricSnapshotsDBFieldName.Dimensions), true

	case string(FinanceMetricSnapshotsDTOFieldName.Metadata):
		return string(model.FinanceMetricSnapshotsDBFieldName.Metadata), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedAt), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedBy), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedAt), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedBy), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedAt), true

	case string(FinanceMetricSnapshotsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceMetricSnapshotsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceMetricSnapshotsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceMetricSnapshotsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceMetricSnapshotsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceMetricSnapshotsProjectionOutputPath(path string) error {
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

func transformFinanceMetricSnapshotsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceMetricSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceMetricSnapshotsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceMetricSnapshotsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceMetricSnapshotsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceMetricSnapshotsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceMetricSnapshotsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceMetricSnapshotsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceMetricSnapshotsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceMetricSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceMetricSnapshotsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceMetricSnapshotsFilter(filter *model.Filter) {
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
			Field: string(FinanceMetricSnapshotsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceMetricSnapshotsSelectableResponse map[string]interface{}
type FinanceMetricSnapshotsSelectableListResponse []*FinanceMetricSnapshotsSelectableResponse

func assignFinanceMetricSnapshotsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceMetricSnapshotsSelectableValue(out FinanceMetricSnapshotsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceMetricSnapshotsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceMetricSnapshotsSelectableResponse(financeMetricSnapshots model.FinanceMetricSnapshots, filter model.Filter) FinanceMetricSnapshotsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceMetricSnapshotsDBFieldName.Id),
			string(model.FinanceMetricSnapshotsDBFieldName.MetricName),
			string(model.FinanceMetricSnapshotsDBFieldName.MetricScope),
			string(model.FinanceMetricSnapshotsDBFieldName.ScopeRef),
			string(model.FinanceMetricSnapshotsDBFieldName.PeriodStart),
			string(model.FinanceMetricSnapshotsDBFieldName.PeriodEnd),
			string(model.FinanceMetricSnapshotsDBFieldName.MetricValue),
			string(model.FinanceMetricSnapshotsDBFieldName.MetricUnit),
			string(model.FinanceMetricSnapshotsDBFieldName.Dimensions),
			string(model.FinanceMetricSnapshotsDBFieldName.Metadata),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedAt),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedBy),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedAt),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedBy),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedAt),
			string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedBy),
		)
	}
	financeMetricSnapshotsSelectableResponse := FinanceMetricSnapshotsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceMetricSnapshotsDBFieldName.Id):
			key := string(FinanceMetricSnapshotsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.Id, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetricName):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetricName)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetricName, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetricScope):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetricScope)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetricScope, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.ScopeRef):
			key := string(FinanceMetricSnapshotsDTOFieldName.ScopeRef)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.ScopeRef.String, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.PeriodStart):
			key := string(FinanceMetricSnapshotsDTOFieldName.PeriodStart)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.PeriodStart, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.PeriodEnd):
			key := string(FinanceMetricSnapshotsDTOFieldName.PeriodEnd)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.PeriodEnd, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetricValue):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetricValue)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetricValue, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetricUnit):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetricUnit)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetricUnit, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.Dimensions):
			key := string(FinanceMetricSnapshotsDTOFieldName.Dimensions)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.Dimensions, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.Metadata):
			key := string(FinanceMetricSnapshotsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.Metadata, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedAt):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaCreatedAt, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaCreatedBy):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaCreatedBy, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedAt):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaUpdatedBy):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedAt):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceMetricSnapshotsDBFieldName.MetaDeletedBy):
			key := string(FinanceMetricSnapshotsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceMetricSnapshotsSelectableValue(financeMetricSnapshotsSelectableResponse, key, financeMetricSnapshots.MetaDeletedBy, explicitAlias)

		}
	}
	return financeMetricSnapshotsSelectableResponse
}

func NewFinanceMetricSnapshotsListResponseFromFilterResult(result []model.FinanceMetricSnapshotsFilterResult, filter model.Filter) FinanceMetricSnapshotsSelectableListResponse {
	dtoFinanceMetricSnapshotsListResponse := FinanceMetricSnapshotsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceMetricSnapshotsResponse := NewFinanceMetricSnapshotsSelectableResponse(row.FinanceMetricSnapshots, filter)
		dtoFinanceMetricSnapshotsListResponse = append(dtoFinanceMetricSnapshotsListResponse, &dtoFinanceMetricSnapshotsResponse)
	}
	return dtoFinanceMetricSnapshotsListResponse
}

type FinanceMetricSnapshotsFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     FinanceMetricSnapshotsSelectableListResponse `json:"data"`
}

func reverseFinanceMetricSnapshotsFilterResults(result []model.FinanceMetricSnapshotsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceMetricSnapshotsFilterResponse(result []model.FinanceMetricSnapshotsFilterResult, filter model.Filter) (resp FinanceMetricSnapshotsFilterResponse) {
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
			reverseFinanceMetricSnapshotsFilterResults(dataResult)
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

	resp.Data = NewFinanceMetricSnapshotsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceMetricSnapshotsCreateRequest struct {
	MetricName  string          `json:"metricName"`
	MetricScope string          `json:"metricScope"`
	ScopeRef    string          `json:"scopeRef"`
	PeriodStart time.Time       `json:"periodStart"`
	PeriodEnd   time.Time       `json:"periodEnd"`
	MetricValue decimal.Decimal `json:"metricValue"`
	MetricUnit  string          `json:"metricUnit"`
	Dimensions  json.RawMessage `json:"dimensions"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *FinanceMetricSnapshotsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceMetricSnapshotsCreateRequest) ToModel() model.FinanceMetricSnapshots {
	id, _ := uuid.NewV4()
	return model.FinanceMetricSnapshots{
		Id:          id,
		MetricName:  d.MetricName,
		MetricScope: d.MetricScope,
		ScopeRef:    null.StringFrom(d.ScopeRef),
		PeriodStart: d.PeriodStart,
		PeriodEnd:   d.PeriodEnd,
		MetricValue: d.MetricValue,
		MetricUnit:  d.MetricUnit,
		Dimensions:  d.Dimensions,
		Metadata:    d.Metadata,
	}
}

type FinanceMetricSnapshotsListCreateRequest []*FinanceMetricSnapshotsCreateRequest

func (d FinanceMetricSnapshotsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeMetricSnapshots := range d {
		err = validator.Struct(financeMetricSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceMetricSnapshotsListCreateRequest) ToModelList() []model.FinanceMetricSnapshots {
	out := make([]model.FinanceMetricSnapshots, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceMetricSnapshotsUpdateRequest struct {
	MetricName  string          `json:"metricName"`
	MetricScope string          `json:"metricScope"`
	ScopeRef    string          `json:"scopeRef"`
	PeriodStart time.Time       `json:"periodStart"`
	PeriodEnd   time.Time       `json:"periodEnd"`
	MetricValue decimal.Decimal `json:"metricValue"`
	MetricUnit  string          `json:"metricUnit"`
	Dimensions  json.RawMessage `json:"dimensions"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *FinanceMetricSnapshotsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceMetricSnapshotsUpdateRequest) ToModel() model.FinanceMetricSnapshots {
	return model.FinanceMetricSnapshots{
		MetricName:  d.MetricName,
		MetricScope: d.MetricScope,
		ScopeRef:    null.StringFrom(d.ScopeRef),
		PeriodStart: d.PeriodStart,
		PeriodEnd:   d.PeriodEnd,
		MetricValue: d.MetricValue,
		MetricUnit:  d.MetricUnit,
		Dimensions:  d.Dimensions,
		Metadata:    d.Metadata,
	}
}

type FinanceMetricSnapshotsBulkUpdateRequest struct {
	Id          uuid.UUID       `json:"id"`
	MetricName  string          `json:"metricName"`
	MetricScope string          `json:"metricScope"`
	ScopeRef    string          `json:"scopeRef"`
	PeriodStart time.Time       `json:"periodStart"`
	PeriodEnd   time.Time       `json:"periodEnd"`
	MetricValue decimal.Decimal `json:"metricValue"`
	MetricUnit  string          `json:"metricUnit"`
	Dimensions  json.RawMessage `json:"dimensions"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d FinanceMetricSnapshotsBulkUpdateRequest) PrimaryID() FinanceMetricSnapshotsPrimaryID {
	return FinanceMetricSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type FinanceMetricSnapshotsListBulkUpdateRequest []*FinanceMetricSnapshotsBulkUpdateRequest

func (d FinanceMetricSnapshotsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeMetricSnapshots := range d {
		err = validator.Struct(financeMetricSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceMetricSnapshotsBulkUpdateRequest) ToModel() model.FinanceMetricSnapshots {
	return model.FinanceMetricSnapshots{
		Id:          d.Id,
		MetricName:  d.MetricName,
		MetricScope: d.MetricScope,
		ScopeRef:    null.StringFrom(d.ScopeRef),
		PeriodStart: d.PeriodStart,
		PeriodEnd:   d.PeriodEnd,
		MetricValue: d.MetricValue,
		MetricUnit:  d.MetricUnit,
		Dimensions:  d.Dimensions,
		Metadata:    d.Metadata,
	}
}

type FinanceMetricSnapshotsResponse struct {
	Id          uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetricName  string          `json:"metricName" validate:"required"`
	MetricScope string          `json:"metricScope" validate:"required"`
	ScopeRef    string          `json:"scopeRef"`
	PeriodStart time.Time       `json:"periodStart" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PeriodEnd   time.Time       `json:"periodEnd" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetricValue decimal.Decimal `json:"metricValue" validate:"required" format:"decimal" example:"100.50"`
	MetricUnit  string          `json:"metricUnit" validate:"required"`
	Dimensions  json.RawMessage `json:"dimensions" swaggertype:"object"`
	Metadata    json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewFinanceMetricSnapshotsResponse(financeMetricSnapshots model.FinanceMetricSnapshots) FinanceMetricSnapshotsResponse {
	return FinanceMetricSnapshotsResponse{
		Id:          financeMetricSnapshots.Id,
		MetricName:  financeMetricSnapshots.MetricName,
		MetricScope: financeMetricSnapshots.MetricScope,
		ScopeRef:    financeMetricSnapshots.ScopeRef.String,
		PeriodStart: financeMetricSnapshots.PeriodStart,
		PeriodEnd:   financeMetricSnapshots.PeriodEnd,
		MetricValue: financeMetricSnapshots.MetricValue,
		MetricUnit:  financeMetricSnapshots.MetricUnit,
		Dimensions:  financeMetricSnapshots.Dimensions,
		Metadata:    financeMetricSnapshots.Metadata,
	}
}

type FinanceMetricSnapshotsListResponse []*FinanceMetricSnapshotsResponse

func NewFinanceMetricSnapshotsListResponse(financeMetricSnapshotsList model.FinanceMetricSnapshotsList) FinanceMetricSnapshotsListResponse {
	dtoFinanceMetricSnapshotsListResponse := FinanceMetricSnapshotsListResponse{}
	for _, financeMetricSnapshots := range financeMetricSnapshotsList {
		dtoFinanceMetricSnapshotsResponse := NewFinanceMetricSnapshotsResponse(*financeMetricSnapshots)
		dtoFinanceMetricSnapshotsListResponse = append(dtoFinanceMetricSnapshotsListResponse, &dtoFinanceMetricSnapshotsResponse)
	}
	return dtoFinanceMetricSnapshotsListResponse
}

type FinanceMetricSnapshotsPrimaryIDList []FinanceMetricSnapshotsPrimaryID

func (d FinanceMetricSnapshotsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeMetricSnapshots := range d {
		err = validator.Struct(financeMetricSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceMetricSnapshotsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceMetricSnapshotsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceMetricSnapshotsPrimaryID) ToModel() model.FinanceMetricSnapshotsPrimaryID {
	return model.FinanceMetricSnapshotsPrimaryID{
		Id: d.Id,
	}
}

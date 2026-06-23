package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceCurrenciesDTOFieldNameType string

type financeCurrenciesDTOFieldName struct {
	Code          FinanceCurrenciesDTOFieldNameType
	DecimalCode   FinanceCurrenciesDTOFieldNameType
	Exponent      FinanceCurrenciesDTOFieldNameType
	IsActive      FinanceCurrenciesDTOFieldNameType
	Metadata      FinanceCurrenciesDTOFieldNameType
	MetaCreatedAt FinanceCurrenciesDTOFieldNameType
	MetaCreatedBy FinanceCurrenciesDTOFieldNameType
	MetaUpdatedAt FinanceCurrenciesDTOFieldNameType
	MetaUpdatedBy FinanceCurrenciesDTOFieldNameType
	MetaDeletedAt FinanceCurrenciesDTOFieldNameType
	MetaDeletedBy FinanceCurrenciesDTOFieldNameType
}

var FinanceCurrenciesDTOFieldName = financeCurrenciesDTOFieldName{
	Code:          "code",
	DecimalCode:   "decimalCode",
	Exponent:      "exponent",
	IsActive:      "isActive",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformFinanceCurrenciesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceCurrenciesDTOFieldName.Code):
		return string(model.FinanceCurrenciesDBFieldName.Code), true

	case string(FinanceCurrenciesDTOFieldName.DecimalCode):
		return string(model.FinanceCurrenciesDBFieldName.DecimalCode), true

	case string(FinanceCurrenciesDTOFieldName.Exponent):
		return string(model.FinanceCurrenciesDBFieldName.Exponent), true

	case string(FinanceCurrenciesDTOFieldName.IsActive):
		return string(model.FinanceCurrenciesDBFieldName.IsActive), true

	case string(FinanceCurrenciesDTOFieldName.Metadata):
		return string(model.FinanceCurrenciesDBFieldName.Metadata), true

	case string(FinanceCurrenciesDTOFieldName.MetaCreatedAt):
		return string(model.FinanceCurrenciesDBFieldName.MetaCreatedAt), true

	case string(FinanceCurrenciesDTOFieldName.MetaCreatedBy):
		return string(model.FinanceCurrenciesDBFieldName.MetaCreatedBy), true

	case string(FinanceCurrenciesDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceCurrenciesDBFieldName.MetaUpdatedAt), true

	case string(FinanceCurrenciesDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceCurrenciesDBFieldName.MetaUpdatedBy), true

	case string(FinanceCurrenciesDTOFieldName.MetaDeletedAt):
		return string(model.FinanceCurrenciesDBFieldName.MetaDeletedAt), true

	case string(FinanceCurrenciesDTOFieldName.MetaDeletedBy):
		return string(model.FinanceCurrenciesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceCurrenciesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceCurrenciesBaseFilterField(field string) bool {
	spec, found := model.NewFinanceCurrenciesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceCurrenciesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceCurrenciesProjectionOutputPath(path string) error {
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

func transformFinanceCurrenciesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceCurrenciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceCurrenciesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceCurrenciesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceCurrenciesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceCurrenciesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceCurrenciesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceCurrenciesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceCurrenciesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceCurrenciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceCurrenciesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceCurrenciesFilter(filter *model.Filter) {
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
			Field: string(FinanceCurrenciesDTOFieldName.Code),
			Order: model.SortAsc,
		})
	}
}

type FinanceCurrenciesSelectableResponse map[string]interface{}
type FinanceCurrenciesSelectableListResponse []*FinanceCurrenciesSelectableResponse

func assignFinanceCurrenciesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceCurrenciesSelectableValue(out FinanceCurrenciesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceCurrenciesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceCurrenciesSelectableResponse(financeCurrencies model.FinanceCurrencies, filter model.Filter) FinanceCurrenciesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceCurrenciesDBFieldName.Code),
			string(model.FinanceCurrenciesDBFieldName.DecimalCode),
			string(model.FinanceCurrenciesDBFieldName.Exponent),
			string(model.FinanceCurrenciesDBFieldName.IsActive),
			string(model.FinanceCurrenciesDBFieldName.Metadata),
			string(model.FinanceCurrenciesDBFieldName.MetaCreatedAt),
			string(model.FinanceCurrenciesDBFieldName.MetaCreatedBy),
			string(model.FinanceCurrenciesDBFieldName.MetaUpdatedAt),
			string(model.FinanceCurrenciesDBFieldName.MetaUpdatedBy),
			string(model.FinanceCurrenciesDBFieldName.MetaDeletedAt),
			string(model.FinanceCurrenciesDBFieldName.MetaDeletedBy),
		)
	}
	financeCurrenciesSelectableResponse := FinanceCurrenciesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceCurrenciesDBFieldName.Code):
			key := string(FinanceCurrenciesDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.Code, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.DecimalCode):
			key := string(FinanceCurrenciesDTOFieldName.DecimalCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, int(financeCurrencies.DecimalCode.ValueOrZero()), explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.Exponent):
			key := string(FinanceCurrenciesDTOFieldName.Exponent)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.Exponent, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.IsActive):
			key := string(FinanceCurrenciesDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.IsActive, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.Metadata):
			key := string(FinanceCurrenciesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.Metadata, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaCreatedAt):
			key := string(FinanceCurrenciesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaCreatedAt, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaCreatedBy):
			key := string(FinanceCurrenciesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaCreatedBy, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaUpdatedAt):
			key := string(FinanceCurrenciesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaUpdatedBy):
			key := string(FinanceCurrenciesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaDeletedAt):
			key := string(FinanceCurrenciesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceCurrenciesDBFieldName.MetaDeletedBy):
			key := string(FinanceCurrenciesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceCurrenciesSelectableValue(financeCurrenciesSelectableResponse, key, financeCurrencies.MetaDeletedBy, explicitAlias)

		}
	}
	return financeCurrenciesSelectableResponse
}

func NewFinanceCurrenciesListResponseFromFilterResult(result []model.FinanceCurrenciesFilterResult, filter model.Filter) FinanceCurrenciesSelectableListResponse {
	dtoFinanceCurrenciesListResponse := FinanceCurrenciesSelectableListResponse{}
	for _, row := range result {
		dtoFinanceCurrenciesResponse := NewFinanceCurrenciesSelectableResponse(row.FinanceCurrencies, filter)
		dtoFinanceCurrenciesListResponse = append(dtoFinanceCurrenciesListResponse, &dtoFinanceCurrenciesResponse)
	}
	return dtoFinanceCurrenciesListResponse
}

type FinanceCurrenciesFilterResponse struct {
	Metadata Metadata                                `json:"metadata"`
	Data     FinanceCurrenciesSelectableListResponse `json:"data"`
}

func reverseFinanceCurrenciesFilterResults(result []model.FinanceCurrenciesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceCurrenciesFilterResponse(result []model.FinanceCurrenciesFilterResult, filter model.Filter) (resp FinanceCurrenciesFilterResponse) {
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
			reverseFinanceCurrenciesFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].Code
			resp.Metadata.PrevCursor = dataResult[0].Code
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

	resp.Data = NewFinanceCurrenciesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceCurrenciesCreateRequest struct {
	Code        string          `json:"code"`
	DecimalCode int16           `json:"decimalCode"`
	Exponent    int16           `json:"exponent"`
	IsActive    bool            `json:"isActive"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *FinanceCurrenciesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceCurrenciesCreateRequest) ToModel() model.FinanceCurrencies {
	return model.FinanceCurrencies{
		Code:        d.Code,
		DecimalCode: null.IntFrom(int64(d.DecimalCode)),
		Exponent:    d.Exponent,
		IsActive:    d.IsActive,
		Metadata:    d.Metadata,
	}
}

type FinanceCurrenciesListCreateRequest []*FinanceCurrenciesCreateRequest

func (d FinanceCurrenciesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeCurrencies := range d {
		err = validator.Struct(financeCurrencies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceCurrenciesListCreateRequest) ToModelList() []model.FinanceCurrencies {
	out := make([]model.FinanceCurrencies, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceCurrenciesUpdateRequest struct {
	Code        string          `json:"code"`
	DecimalCode int16           `json:"decimalCode"`
	Exponent    int16           `json:"exponent"`
	IsActive    bool            `json:"isActive"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *FinanceCurrenciesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceCurrenciesUpdateRequest) ToModel() model.FinanceCurrencies {
	return model.FinanceCurrencies{
		Code:        d.Code,
		DecimalCode: null.IntFrom(int64(d.DecimalCode)),
		Exponent:    d.Exponent,
		IsActive:    d.IsActive,
		Metadata:    d.Metadata,
	}
}

type FinanceCurrenciesBulkUpdateRequest struct {
	Code        string          `json:"code"`
	DecimalCode int16           `json:"decimalCode"`
	Exponent    int16           `json:"exponent"`
	IsActive    bool            `json:"isActive"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d FinanceCurrenciesBulkUpdateRequest) PrimaryID() FinanceCurrenciesPrimaryID {
	return FinanceCurrenciesPrimaryID{
		Code: d.Code,
	}
}

type FinanceCurrenciesListBulkUpdateRequest []*FinanceCurrenciesBulkUpdateRequest

func (d FinanceCurrenciesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeCurrencies := range d {
		err = validator.Struct(financeCurrencies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceCurrenciesBulkUpdateRequest) ToModel() model.FinanceCurrencies {
	return model.FinanceCurrencies{
		Code:        d.Code,
		DecimalCode: null.IntFrom(int64(d.DecimalCode)),
		Exponent:    d.Exponent,
		IsActive:    d.IsActive,
		Metadata:    d.Metadata,
	}
}

type FinanceCurrenciesResponse struct {
	Code        string          `json:"code" validate:"required"`
	DecimalCode int16           `json:"decimalCode"`
	Exponent    int16           `json:"exponent"`
	IsActive    bool            `json:"isActive" example:"true"`
	Metadata    json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewFinanceCurrenciesResponse(financeCurrencies model.FinanceCurrencies) FinanceCurrenciesResponse {
	return FinanceCurrenciesResponse{
		Code:        financeCurrencies.Code,
		DecimalCode: int16(financeCurrencies.DecimalCode.ValueOrZero()),
		Exponent:    financeCurrencies.Exponent,
		IsActive:    financeCurrencies.IsActive,
		Metadata:    financeCurrencies.Metadata,
	}
}

type FinanceCurrenciesListResponse []*FinanceCurrenciesResponse

func NewFinanceCurrenciesListResponse(financeCurrenciesList model.FinanceCurrenciesList) FinanceCurrenciesListResponse {
	dtoFinanceCurrenciesListResponse := FinanceCurrenciesListResponse{}
	for _, financeCurrencies := range financeCurrenciesList {
		dtoFinanceCurrenciesResponse := NewFinanceCurrenciesResponse(*financeCurrencies)
		dtoFinanceCurrenciesListResponse = append(dtoFinanceCurrenciesListResponse, &dtoFinanceCurrenciesResponse)
	}
	return dtoFinanceCurrenciesListResponse
}

type FinanceCurrenciesPrimaryIDList []FinanceCurrenciesPrimaryID

func (d FinanceCurrenciesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeCurrencies := range d {
		err = validator.Struct(financeCurrencies)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceCurrenciesPrimaryID struct {
	Code string `json:"code" validate:"required"`
}

func (d *FinanceCurrenciesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceCurrenciesPrimaryID) ToModel() model.FinanceCurrenciesPrimaryID {
	return model.FinanceCurrenciesPrimaryID{
		Code: d.Code,
	}
}

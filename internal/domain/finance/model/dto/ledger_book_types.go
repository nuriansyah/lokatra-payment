package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerBookTypesDTOFieldNameType string

type ledgerBookTypesDTOFieldName struct {
	Code          LedgerBookTypesDTOFieldNameType
	Description   LedgerBookTypesDTOFieldNameType
	Metadata      LedgerBookTypesDTOFieldNameType
	MetaCreatedAt LedgerBookTypesDTOFieldNameType
	MetaCreatedBy LedgerBookTypesDTOFieldNameType
	MetaUpdatedAt LedgerBookTypesDTOFieldNameType
	MetaUpdatedBy LedgerBookTypesDTOFieldNameType
	MetaDeletedAt LedgerBookTypesDTOFieldNameType
	MetaDeletedBy LedgerBookTypesDTOFieldNameType
}

var LedgerBookTypesDTOFieldName = ledgerBookTypesDTOFieldName{
	Code:          "code",
	Description:   "description",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformLedgerBookTypesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerBookTypesDTOFieldName.Code):
		return string(model.LedgerBookTypesDBFieldName.Code), true

	case string(LedgerBookTypesDTOFieldName.Description):
		return string(model.LedgerBookTypesDBFieldName.Description), true

	case string(LedgerBookTypesDTOFieldName.Metadata):
		return string(model.LedgerBookTypesDBFieldName.Metadata), true

	case string(LedgerBookTypesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerBookTypesDBFieldName.MetaCreatedAt), true

	case string(LedgerBookTypesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerBookTypesDBFieldName.MetaCreatedBy), true

	case string(LedgerBookTypesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerBookTypesDBFieldName.MetaUpdatedAt), true

	case string(LedgerBookTypesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerBookTypesDBFieldName.MetaUpdatedBy), true

	case string(LedgerBookTypesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerBookTypesDBFieldName.MetaDeletedAt), true

	case string(LedgerBookTypesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerBookTypesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerBookTypesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerBookTypesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerBookTypesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerBookTypesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerBookTypesProjectionOutputPath(path string) error {
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

func transformLedgerBookTypesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerBookTypesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerBookTypesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerBookTypesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerBookTypesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerBookTypesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerBookTypesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerBookTypesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerBookTypesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerBookTypesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerBookTypesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerBookTypesFilter(filter *model.Filter) {
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
			Field: string(LedgerBookTypesDTOFieldName.Code),
			Order: model.SortAsc,
		})
	}
}

type LedgerBookTypesSelectableResponse map[string]interface{}
type LedgerBookTypesSelectableListResponse []*LedgerBookTypesSelectableResponse

func assignLedgerBookTypesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerBookTypesSelectableValue(out LedgerBookTypesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerBookTypesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerBookTypesSelectableResponse(ledgerBookTypes model.LedgerBookTypes, filter model.Filter) LedgerBookTypesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerBookTypesDBFieldName.Code),
			string(model.LedgerBookTypesDBFieldName.Description),
			string(model.LedgerBookTypesDBFieldName.Metadata),
			string(model.LedgerBookTypesDBFieldName.MetaCreatedAt),
			string(model.LedgerBookTypesDBFieldName.MetaCreatedBy),
			string(model.LedgerBookTypesDBFieldName.MetaUpdatedAt),
			string(model.LedgerBookTypesDBFieldName.MetaUpdatedBy),
			string(model.LedgerBookTypesDBFieldName.MetaDeletedAt),
			string(model.LedgerBookTypesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerBookTypesSelectableResponse := LedgerBookTypesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerBookTypesDBFieldName.Code):
			key := string(LedgerBookTypesDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.Code, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.Description):
			key := string(LedgerBookTypesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.Description, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.Metadata):
			key := string(LedgerBookTypesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.Metadata, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaCreatedAt):
			key := string(LedgerBookTypesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaCreatedAt, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaCreatedBy):
			key := string(LedgerBookTypesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaCreatedBy, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaUpdatedAt):
			key := string(LedgerBookTypesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaUpdatedBy):
			key := string(LedgerBookTypesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaDeletedAt):
			key := string(LedgerBookTypesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerBookTypesDBFieldName.MetaDeletedBy):
			key := string(LedgerBookTypesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBookTypesSelectableValue(ledgerBookTypesSelectableResponse, key, ledgerBookTypes.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerBookTypesSelectableResponse
}

func NewLedgerBookTypesListResponseFromFilterResult(result []model.LedgerBookTypesFilterResult, filter model.Filter) LedgerBookTypesSelectableListResponse {
	dtoLedgerBookTypesListResponse := LedgerBookTypesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerBookTypesResponse := NewLedgerBookTypesSelectableResponse(row.LedgerBookTypes, filter)
		dtoLedgerBookTypesListResponse = append(dtoLedgerBookTypesListResponse, &dtoLedgerBookTypesResponse)
	}
	return dtoLedgerBookTypesListResponse
}

type LedgerBookTypesFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     LedgerBookTypesSelectableListResponse `json:"data"`
}

func reverseLedgerBookTypesFilterResults(result []model.LedgerBookTypesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerBookTypesFilterResponse(result []model.LedgerBookTypesFilterResult, filter model.Filter) (resp LedgerBookTypesFilterResponse) {
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
			reverseLedgerBookTypesFilterResults(dataResult)
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

	resp.Data = NewLedgerBookTypesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerBookTypesCreateRequest struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *LedgerBookTypesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerBookTypesCreateRequest) ToModel() model.LedgerBookTypes {
	return model.LedgerBookTypes{
		Code:        d.Code,
		Description: d.Description,
		Metadata:    d.Metadata,
	}
}

type LedgerBookTypesListCreateRequest []*LedgerBookTypesCreateRequest

func (d LedgerBookTypesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBookTypes := range d {
		err = validator.Struct(ledgerBookTypes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerBookTypesListCreateRequest) ToModelList() []model.LedgerBookTypes {
	out := make([]model.LedgerBookTypes, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerBookTypesUpdateRequest struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d *LedgerBookTypesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerBookTypesUpdateRequest) ToModel() model.LedgerBookTypes {
	return model.LedgerBookTypes{
		Code:        d.Code,
		Description: d.Description,
		Metadata:    d.Metadata,
	}
}

type LedgerBookTypesBulkUpdateRequest struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Metadata    json.RawMessage `json:"metadata"`
}

func (d LedgerBookTypesBulkUpdateRequest) PrimaryID() LedgerBookTypesPrimaryID {
	return LedgerBookTypesPrimaryID{
		Code: d.Code,
	}
}

type LedgerBookTypesListBulkUpdateRequest []*LedgerBookTypesBulkUpdateRequest

func (d LedgerBookTypesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBookTypes := range d {
		err = validator.Struct(ledgerBookTypes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerBookTypesBulkUpdateRequest) ToModel() model.LedgerBookTypes {
	return model.LedgerBookTypes{
		Code:        d.Code,
		Description: d.Description,
		Metadata:    d.Metadata,
	}
}

type LedgerBookTypesResponse struct {
	Code        string          `json:"code" validate:"required"`
	Description string          `json:"description" validate:"required"`
	Metadata    json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewLedgerBookTypesResponse(ledgerBookTypes model.LedgerBookTypes) LedgerBookTypesResponse {
	return LedgerBookTypesResponse{
		Code:        ledgerBookTypes.Code,
		Description: ledgerBookTypes.Description,
		Metadata:    ledgerBookTypes.Metadata,
	}
}

type LedgerBookTypesListResponse []*LedgerBookTypesResponse

func NewLedgerBookTypesListResponse(ledgerBookTypesList model.LedgerBookTypesList) LedgerBookTypesListResponse {
	dtoLedgerBookTypesListResponse := LedgerBookTypesListResponse{}
	for _, ledgerBookTypes := range ledgerBookTypesList {
		dtoLedgerBookTypesResponse := NewLedgerBookTypesResponse(*ledgerBookTypes)
		dtoLedgerBookTypesListResponse = append(dtoLedgerBookTypesListResponse, &dtoLedgerBookTypesResponse)
	}
	return dtoLedgerBookTypesListResponse
}

type LedgerBookTypesPrimaryIDList []LedgerBookTypesPrimaryID

func (d LedgerBookTypesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBookTypes := range d {
		err = validator.Struct(ledgerBookTypes)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerBookTypesPrimaryID struct {
	Code string `json:"code" validate:"required"`
}

func (d *LedgerBookTypesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerBookTypesPrimaryID) ToModel() model.LedgerBookTypesPrimaryID {
	return model.LedgerBookTypesPrimaryID{
		Code: d.Code,
	}
}

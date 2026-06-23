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

type TaxDocumentSequencesDTOFieldNameType string

type taxDocumentSequencesDTOFieldName struct {
	SequenceCode  TaxDocumentSequencesDTOFieldNameType
	CurrentValue  TaxDocumentSequencesDTOFieldNameType
	Metadata      TaxDocumentSequencesDTOFieldNameType
	MetaCreatedAt TaxDocumentSequencesDTOFieldNameType
	MetaCreatedBy TaxDocumentSequencesDTOFieldNameType
	MetaUpdatedAt TaxDocumentSequencesDTOFieldNameType
	MetaUpdatedBy TaxDocumentSequencesDTOFieldNameType
	MetaDeletedAt TaxDocumentSequencesDTOFieldNameType
	MetaDeletedBy TaxDocumentSequencesDTOFieldNameType
}

var TaxDocumentSequencesDTOFieldName = taxDocumentSequencesDTOFieldName{
	SequenceCode:  "sequenceCode",
	CurrentValue:  "currentValue",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformTaxDocumentSequencesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxDocumentSequencesDTOFieldName.SequenceCode):
		return string(model.TaxDocumentSequencesDBFieldName.SequenceCode), true

	case string(TaxDocumentSequencesDTOFieldName.CurrentValue):
		return string(model.TaxDocumentSequencesDBFieldName.CurrentValue), true

	case string(TaxDocumentSequencesDTOFieldName.Metadata):
		return string(model.TaxDocumentSequencesDBFieldName.Metadata), true

	case string(TaxDocumentSequencesDTOFieldName.MetaCreatedAt):
		return string(model.TaxDocumentSequencesDBFieldName.MetaCreatedAt), true

	case string(TaxDocumentSequencesDTOFieldName.MetaCreatedBy):
		return string(model.TaxDocumentSequencesDBFieldName.MetaCreatedBy), true

	case string(TaxDocumentSequencesDTOFieldName.MetaUpdatedAt):
		return string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedAt), true

	case string(TaxDocumentSequencesDTOFieldName.MetaUpdatedBy):
		return string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedBy), true

	case string(TaxDocumentSequencesDTOFieldName.MetaDeletedAt):
		return string(model.TaxDocumentSequencesDBFieldName.MetaDeletedAt), true

	case string(TaxDocumentSequencesDTOFieldName.MetaDeletedBy):
		return string(model.TaxDocumentSequencesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxDocumentSequencesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxDocumentSequencesBaseFilterField(field string) bool {
	spec, found := model.NewTaxDocumentSequencesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxDocumentSequencesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxDocumentSequencesProjectionOutputPath(path string) error {
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

func transformTaxDocumentSequencesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxDocumentSequencesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxDocumentSequencesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxDocumentSequencesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxDocumentSequencesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxDocumentSequencesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxDocumentSequencesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxDocumentSequencesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxDocumentSequencesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxDocumentSequencesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxDocumentSequencesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxDocumentSequencesFilter(filter *model.Filter) {
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
			Field: string(TaxDocumentSequencesDTOFieldName.SequenceCode),
			Order: model.SortAsc,
		})
	}
}

type TaxDocumentSequencesSelectableResponse map[string]interface{}
type TaxDocumentSequencesSelectableListResponse []*TaxDocumentSequencesSelectableResponse

func assignTaxDocumentSequencesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxDocumentSequencesSelectableValue(out TaxDocumentSequencesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxDocumentSequencesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxDocumentSequencesSelectableResponse(taxDocumentSequences model.TaxDocumentSequences, filter model.Filter) TaxDocumentSequencesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxDocumentSequencesDBFieldName.SequenceCode),
			string(model.TaxDocumentSequencesDBFieldName.CurrentValue),
			string(model.TaxDocumentSequencesDBFieldName.Metadata),
			string(model.TaxDocumentSequencesDBFieldName.MetaCreatedAt),
			string(model.TaxDocumentSequencesDBFieldName.MetaCreatedBy),
			string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedAt),
			string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedBy),
			string(model.TaxDocumentSequencesDBFieldName.MetaDeletedAt),
			string(model.TaxDocumentSequencesDBFieldName.MetaDeletedBy),
		)
	}
	taxDocumentSequencesSelectableResponse := TaxDocumentSequencesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxDocumentSequencesDBFieldName.SequenceCode):
			key := string(TaxDocumentSequencesDTOFieldName.SequenceCode)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.SequenceCode, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.CurrentValue):
			key := string(TaxDocumentSequencesDTOFieldName.CurrentValue)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.CurrentValue, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.Metadata):
			key := string(TaxDocumentSequencesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.Metadata, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaCreatedAt):
			key := string(TaxDocumentSequencesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaCreatedAt, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaCreatedBy):
			key := string(TaxDocumentSequencesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaCreatedBy, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedAt):
			key := string(TaxDocumentSequencesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaUpdatedAt, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaUpdatedBy):
			key := string(TaxDocumentSequencesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaUpdatedBy, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaDeletedAt):
			key := string(TaxDocumentSequencesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxDocumentSequencesDBFieldName.MetaDeletedBy):
			key := string(TaxDocumentSequencesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentSequencesSelectableValue(taxDocumentSequencesSelectableResponse, key, taxDocumentSequences.MetaDeletedBy, explicitAlias)

		}
	}
	return taxDocumentSequencesSelectableResponse
}

func NewTaxDocumentSequencesListResponseFromFilterResult(result []model.TaxDocumentSequencesFilterResult, filter model.Filter) TaxDocumentSequencesSelectableListResponse {
	dtoTaxDocumentSequencesListResponse := TaxDocumentSequencesSelectableListResponse{}
	for _, row := range result {
		dtoTaxDocumentSequencesResponse := NewTaxDocumentSequencesSelectableResponse(row.TaxDocumentSequences, filter)
		dtoTaxDocumentSequencesListResponse = append(dtoTaxDocumentSequencesListResponse, &dtoTaxDocumentSequencesResponse)
	}
	return dtoTaxDocumentSequencesListResponse
}

type TaxDocumentSequencesFilterResponse struct {
	Metadata Metadata                                   `json:"metadata"`
	Data     TaxDocumentSequencesSelectableListResponse `json:"data"`
}

func reverseTaxDocumentSequencesFilterResults(result []model.TaxDocumentSequencesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxDocumentSequencesFilterResponse(result []model.TaxDocumentSequencesFilterResult, filter model.Filter) (resp TaxDocumentSequencesFilterResponse) {
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
			reverseTaxDocumentSequencesFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].SequenceCode
			resp.Metadata.PrevCursor = dataResult[0].SequenceCode
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

	resp.Data = NewTaxDocumentSequencesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxDocumentSequencesCreateRequest struct {
	SequenceCode string          `json:"sequenceCode"`
	CurrentValue int64           `json:"currentValue"`
	Metadata     json.RawMessage `json:"metadata"`
}

func (d *TaxDocumentSequencesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxDocumentSequencesCreateRequest) ToModel() model.TaxDocumentSequences {
	return model.TaxDocumentSequences{
		SequenceCode: d.SequenceCode,
		CurrentValue: d.CurrentValue,
		Metadata:     d.Metadata,
	}
}

type TaxDocumentSequencesListCreateRequest []*TaxDocumentSequencesCreateRequest

func (d TaxDocumentSequencesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentSequences := range d {
		err = validator.Struct(taxDocumentSequences)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxDocumentSequencesListCreateRequest) ToModelList() []model.TaxDocumentSequences {
	out := make([]model.TaxDocumentSequences, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxDocumentSequencesUpdateRequest struct {
	SequenceCode string          `json:"sequenceCode"`
	CurrentValue int64           `json:"currentValue"`
	Metadata     json.RawMessage `json:"metadata"`
}

func (d *TaxDocumentSequencesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxDocumentSequencesUpdateRequest) ToModel() model.TaxDocumentSequences {
	return model.TaxDocumentSequences{
		SequenceCode: d.SequenceCode,
		CurrentValue: d.CurrentValue,
		Metadata:     d.Metadata,
	}
}

type TaxDocumentSequencesBulkUpdateRequest struct {
	SequenceCode string          `json:"sequenceCode"`
	CurrentValue int64           `json:"currentValue"`
	Metadata     json.RawMessage `json:"metadata"`
}

func (d TaxDocumentSequencesBulkUpdateRequest) PrimaryID() TaxDocumentSequencesPrimaryID {
	return TaxDocumentSequencesPrimaryID{
		SequenceCode: d.SequenceCode,
	}
}

type TaxDocumentSequencesListBulkUpdateRequest []*TaxDocumentSequencesBulkUpdateRequest

func (d TaxDocumentSequencesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentSequences := range d {
		err = validator.Struct(taxDocumentSequences)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxDocumentSequencesBulkUpdateRequest) ToModel() model.TaxDocumentSequences {
	return model.TaxDocumentSequences{
		SequenceCode: d.SequenceCode,
		CurrentValue: d.CurrentValue,
		Metadata:     d.Metadata,
	}
}

type TaxDocumentSequencesResponse struct {
	SequenceCode string          `json:"sequenceCode" validate:"required"`
	CurrentValue int64           `json:"currentValue" format:"int64" example:"1"`
	Metadata     json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewTaxDocumentSequencesResponse(taxDocumentSequences model.TaxDocumentSequences) TaxDocumentSequencesResponse {
	return TaxDocumentSequencesResponse{
		SequenceCode: taxDocumentSequences.SequenceCode,
		CurrentValue: taxDocumentSequences.CurrentValue,
		Metadata:     taxDocumentSequences.Metadata,
	}
}

type TaxDocumentSequencesListResponse []*TaxDocumentSequencesResponse

func NewTaxDocumentSequencesListResponse(taxDocumentSequencesList model.TaxDocumentSequencesList) TaxDocumentSequencesListResponse {
	dtoTaxDocumentSequencesListResponse := TaxDocumentSequencesListResponse{}
	for _, taxDocumentSequences := range taxDocumentSequencesList {
		dtoTaxDocumentSequencesResponse := NewTaxDocumentSequencesResponse(*taxDocumentSequences)
		dtoTaxDocumentSequencesListResponse = append(dtoTaxDocumentSequencesListResponse, &dtoTaxDocumentSequencesResponse)
	}
	return dtoTaxDocumentSequencesListResponse
}

type TaxDocumentSequencesPrimaryIDList []TaxDocumentSequencesPrimaryID

func (d TaxDocumentSequencesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentSequences := range d {
		err = validator.Struct(taxDocumentSequences)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxDocumentSequencesPrimaryID struct {
	SequenceCode string `json:"sequenceCode" validate:"required"`
}

func (d *TaxDocumentSequencesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxDocumentSequencesPrimaryID) ToModel() model.TaxDocumentSequencesPrimaryID {
	return model.TaxDocumentSequencesPrimaryID{
		SequenceCode: d.SequenceCode,
	}
}

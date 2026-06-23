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

type LedgerAccountTypesDTOFieldNameType string

type ledgerAccountTypesDTOFieldName struct {
	Code             LedgerAccountTypesDTOFieldNameType
	NormalSide       LedgerAccountTypesDTOFieldNameType
	Category         LedgerAccountTypesDTOFieldNameType
	Description      LedgerAccountTypesDTOFieldNameType
	IsControlAccount LedgerAccountTypesDTOFieldNameType
	Metadata         LedgerAccountTypesDTOFieldNameType
	MetaCreatedAt    LedgerAccountTypesDTOFieldNameType
	MetaCreatedBy    LedgerAccountTypesDTOFieldNameType
	MetaUpdatedAt    LedgerAccountTypesDTOFieldNameType
	MetaUpdatedBy    LedgerAccountTypesDTOFieldNameType
	MetaDeletedAt    LedgerAccountTypesDTOFieldNameType
	MetaDeletedBy    LedgerAccountTypesDTOFieldNameType
}

var LedgerAccountTypesDTOFieldName = ledgerAccountTypesDTOFieldName{
	Code:             "code",
	NormalSide:       "normalSide",
	Category:         "category",
	Description:      "description",
	IsControlAccount: "isControlAccount",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformLedgerAccountTypesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerAccountTypesDTOFieldName.Code):
		return string(model.LedgerAccountTypesDBFieldName.Code), true

	case string(LedgerAccountTypesDTOFieldName.NormalSide):
		return string(model.LedgerAccountTypesDBFieldName.NormalSide), true

	case string(LedgerAccountTypesDTOFieldName.Category):
		return string(model.LedgerAccountTypesDBFieldName.Category), true

	case string(LedgerAccountTypesDTOFieldName.Description):
		return string(model.LedgerAccountTypesDBFieldName.Description), true

	case string(LedgerAccountTypesDTOFieldName.IsControlAccount):
		return string(model.LedgerAccountTypesDBFieldName.IsControlAccount), true

	case string(LedgerAccountTypesDTOFieldName.Metadata):
		return string(model.LedgerAccountTypesDBFieldName.Metadata), true

	case string(LedgerAccountTypesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerAccountTypesDBFieldName.MetaCreatedAt), true

	case string(LedgerAccountTypesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerAccountTypesDBFieldName.MetaCreatedBy), true

	case string(LedgerAccountTypesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerAccountTypesDBFieldName.MetaUpdatedAt), true

	case string(LedgerAccountTypesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerAccountTypesDBFieldName.MetaUpdatedBy), true

	case string(LedgerAccountTypesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerAccountTypesDBFieldName.MetaDeletedAt), true

	case string(LedgerAccountTypesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerAccountTypesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerAccountTypesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerAccountTypesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerAccountTypesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerAccountTypesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerAccountTypesProjectionOutputPath(path string) error {
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

func transformLedgerAccountTypesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerAccountTypesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerAccountTypesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerAccountTypesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerAccountTypesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerAccountTypesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerAccountTypesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerAccountTypesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerAccountTypesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerAccountTypesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerAccountTypesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerAccountTypesFilter(filter *model.Filter) {
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
			Field: string(LedgerAccountTypesDTOFieldName.Code),
			Order: model.SortAsc,
		})
	}
}

type LedgerAccountTypesSelectableResponse map[string]interface{}
type LedgerAccountTypesSelectableListResponse []*LedgerAccountTypesSelectableResponse

func assignLedgerAccountTypesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerAccountTypesSelectableValue(out LedgerAccountTypesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerAccountTypesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerAccountTypesSelectableResponse(ledgerAccountTypes model.LedgerAccountTypes, filter model.Filter) LedgerAccountTypesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerAccountTypesDBFieldName.Code),
			string(model.LedgerAccountTypesDBFieldName.NormalSide),
			string(model.LedgerAccountTypesDBFieldName.Category),
			string(model.LedgerAccountTypesDBFieldName.Description),
			string(model.LedgerAccountTypesDBFieldName.IsControlAccount),
			string(model.LedgerAccountTypesDBFieldName.Metadata),
			string(model.LedgerAccountTypesDBFieldName.MetaCreatedAt),
			string(model.LedgerAccountTypesDBFieldName.MetaCreatedBy),
			string(model.LedgerAccountTypesDBFieldName.MetaUpdatedAt),
			string(model.LedgerAccountTypesDBFieldName.MetaUpdatedBy),
			string(model.LedgerAccountTypesDBFieldName.MetaDeletedAt),
			string(model.LedgerAccountTypesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerAccountTypesSelectableResponse := LedgerAccountTypesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerAccountTypesDBFieldName.Code):
			key := string(LedgerAccountTypesDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.Code, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.NormalSide):
			key := string(LedgerAccountTypesDTOFieldName.NormalSide)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, model.NormalSide(ledgerAccountTypes.NormalSide), explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.Category):
			key := string(LedgerAccountTypesDTOFieldName.Category)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, model.Category(ledgerAccountTypes.Category), explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.Description):
			key := string(LedgerAccountTypesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.Description, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.IsControlAccount):
			key := string(LedgerAccountTypesDTOFieldName.IsControlAccount)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.IsControlAccount, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.Metadata):
			key := string(LedgerAccountTypesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.Metadata, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaCreatedAt):
			key := string(LedgerAccountTypesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaCreatedAt, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaCreatedBy):
			key := string(LedgerAccountTypesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaCreatedBy, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaUpdatedAt):
			key := string(LedgerAccountTypesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaUpdatedBy):
			key := string(LedgerAccountTypesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaDeletedAt):
			key := string(LedgerAccountTypesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerAccountTypesDBFieldName.MetaDeletedBy):
			key := string(LedgerAccountTypesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountTypesSelectableValue(ledgerAccountTypesSelectableResponse, key, ledgerAccountTypes.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerAccountTypesSelectableResponse
}

func NewLedgerAccountTypesListResponseFromFilterResult(result []model.LedgerAccountTypesFilterResult, filter model.Filter) LedgerAccountTypesSelectableListResponse {
	dtoLedgerAccountTypesListResponse := LedgerAccountTypesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerAccountTypesResponse := NewLedgerAccountTypesSelectableResponse(row.LedgerAccountTypes, filter)
		dtoLedgerAccountTypesListResponse = append(dtoLedgerAccountTypesListResponse, &dtoLedgerAccountTypesResponse)
	}
	return dtoLedgerAccountTypesListResponse
}

type LedgerAccountTypesFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     LedgerAccountTypesSelectableListResponse `json:"data"`
}

func reverseLedgerAccountTypesFilterResults(result []model.LedgerAccountTypesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerAccountTypesFilterResponse(result []model.LedgerAccountTypesFilterResult, filter model.Filter) (resp LedgerAccountTypesFilterResponse) {
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
			reverseLedgerAccountTypesFilterResults(dataResult)
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

	resp.Data = NewLedgerAccountTypesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerAccountTypesCreateRequest struct {
	Code             string           `json:"code"`
	NormalSide       model.NormalSide `json:"normalSide" example:"debit" enums:"debit,credit"`
	Category         model.Category   `json:"category" example:"asset" enums:"asset,liability,equity,revenue,expense,contra_asset,contra_revenue"`
	Description      string           `json:"description"`
	IsControlAccount bool             `json:"isControlAccount"`
	Metadata         json.RawMessage  `json:"metadata"`
}

func (d *LedgerAccountTypesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerAccountTypesCreateRequest) ToModel() model.LedgerAccountTypes {
	return model.LedgerAccountTypes{
		Code:             d.Code,
		NormalSide:       d.NormalSide,
		Category:         d.Category,
		Description:      d.Description,
		IsControlAccount: d.IsControlAccount,
		Metadata:         d.Metadata,
	}
}

type LedgerAccountTypesListCreateRequest []*LedgerAccountTypesCreateRequest

func (d LedgerAccountTypesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountTypes := range d {
		err = validator.Struct(ledgerAccountTypes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountTypesListCreateRequest) ToModelList() []model.LedgerAccountTypes {
	out := make([]model.LedgerAccountTypes, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerAccountTypesUpdateRequest struct {
	Code             string           `json:"code"`
	NormalSide       model.NormalSide `json:"normalSide" example:"debit" enums:"debit,credit"`
	Category         model.Category   `json:"category" example:"asset" enums:"asset,liability,equity,revenue,expense,contra_asset,contra_revenue"`
	Description      string           `json:"description"`
	IsControlAccount bool             `json:"isControlAccount"`
	Metadata         json.RawMessage  `json:"metadata"`
}

func (d *LedgerAccountTypesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerAccountTypesUpdateRequest) ToModel() model.LedgerAccountTypes {
	return model.LedgerAccountTypes{
		Code:             d.Code,
		NormalSide:       d.NormalSide,
		Category:         d.Category,
		Description:      d.Description,
		IsControlAccount: d.IsControlAccount,
		Metadata:         d.Metadata,
	}
}

type LedgerAccountTypesBulkUpdateRequest struct {
	Code             string           `json:"code"`
	NormalSide       model.NormalSide `json:"normalSide" example:"debit" enums:"debit,credit"`
	Category         model.Category   `json:"category" example:"asset" enums:"asset,liability,equity,revenue,expense,contra_asset,contra_revenue"`
	Description      string           `json:"description"`
	IsControlAccount bool             `json:"isControlAccount"`
	Metadata         json.RawMessage  `json:"metadata"`
}

func (d LedgerAccountTypesBulkUpdateRequest) PrimaryID() LedgerAccountTypesPrimaryID {
	return LedgerAccountTypesPrimaryID{
		Code: d.Code,
	}
}

type LedgerAccountTypesListBulkUpdateRequest []*LedgerAccountTypesBulkUpdateRequest

func (d LedgerAccountTypesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountTypes := range d {
		err = validator.Struct(ledgerAccountTypes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountTypesBulkUpdateRequest) ToModel() model.LedgerAccountTypes {
	return model.LedgerAccountTypes{
		Code:             d.Code,
		NormalSide:       d.NormalSide,
		Category:         d.Category,
		Description:      d.Description,
		IsControlAccount: d.IsControlAccount,
		Metadata:         d.Metadata,
	}
}

type LedgerAccountTypesResponse struct {
	Code             string           `json:"code" validate:"required"`
	NormalSide       model.NormalSide `json:"normalSide" validate:"required,oneof=debit credit" enums:"debit,credit"`
	Category         model.Category   `json:"category" validate:"required,oneof=asset liability equity revenue expense contra_asset contra_revenue" enums:"asset,liability,equity,revenue,expense,contra_asset,contra_revenue"`
	Description      string           `json:"description" validate:"required"`
	IsControlAccount bool             `json:"isControlAccount" example:"true"`
	Metadata         json.RawMessage  `json:"metadata" swaggertype:"object"`
}

func NewLedgerAccountTypesResponse(ledgerAccountTypes model.LedgerAccountTypes) LedgerAccountTypesResponse {
	return LedgerAccountTypesResponse{
		Code:             ledgerAccountTypes.Code,
		NormalSide:       model.NormalSide(ledgerAccountTypes.NormalSide),
		Category:         model.Category(ledgerAccountTypes.Category),
		Description:      ledgerAccountTypes.Description,
		IsControlAccount: ledgerAccountTypes.IsControlAccount,
		Metadata:         ledgerAccountTypes.Metadata,
	}
}

type LedgerAccountTypesListResponse []*LedgerAccountTypesResponse

func NewLedgerAccountTypesListResponse(ledgerAccountTypesList model.LedgerAccountTypesList) LedgerAccountTypesListResponse {
	dtoLedgerAccountTypesListResponse := LedgerAccountTypesListResponse{}
	for _, ledgerAccountTypes := range ledgerAccountTypesList {
		dtoLedgerAccountTypesResponse := NewLedgerAccountTypesResponse(*ledgerAccountTypes)
		dtoLedgerAccountTypesListResponse = append(dtoLedgerAccountTypesListResponse, &dtoLedgerAccountTypesResponse)
	}
	return dtoLedgerAccountTypesListResponse
}

type LedgerAccountTypesPrimaryIDList []LedgerAccountTypesPrimaryID

func (d LedgerAccountTypesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountTypes := range d {
		err = validator.Struct(ledgerAccountTypes)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerAccountTypesPrimaryID struct {
	Code string `json:"code" validate:"required"`
}

func (d *LedgerAccountTypesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerAccountTypesPrimaryID) ToModel() model.LedgerAccountTypesPrimaryID {
	return model.LedgerAccountTypesPrimaryID{
		Code: d.Code,
	}
}

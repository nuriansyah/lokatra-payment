package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerAccountMappingsDTOFieldNameType string

type ledgerAccountMappingsDTOFieldName struct {
	Id            LedgerAccountMappingsDTOFieldNameType
	MappingCode   LedgerAccountMappingsDTOFieldNameType
	BookId        LedgerAccountMappingsDTOFieldNameType
	SourceType    LedgerAccountMappingsDTOFieldNameType
	SourceSubtype LedgerAccountMappingsDTOFieldNameType
	AccountId     LedgerAccountMappingsDTOFieldNameType
	Priority      LedgerAccountMappingsDTOFieldNameType
	IsActive      LedgerAccountMappingsDTOFieldNameType
	Conditions    LedgerAccountMappingsDTOFieldNameType
	Metadata      LedgerAccountMappingsDTOFieldNameType
	MetaCreatedAt LedgerAccountMappingsDTOFieldNameType
	MetaCreatedBy LedgerAccountMappingsDTOFieldNameType
	MetaUpdatedAt LedgerAccountMappingsDTOFieldNameType
	MetaUpdatedBy LedgerAccountMappingsDTOFieldNameType
	MetaDeletedAt LedgerAccountMappingsDTOFieldNameType
	MetaDeletedBy LedgerAccountMappingsDTOFieldNameType
}

var LedgerAccountMappingsDTOFieldName = ledgerAccountMappingsDTOFieldName{
	Id:            "id",
	MappingCode:   "mappingCode",
	BookId:        "bookId",
	SourceType:    "sourceType",
	SourceSubtype: "sourceSubtype",
	AccountId:     "accountId",
	Priority:      "priority",
	IsActive:      "isActive",
	Conditions:    "conditions",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformLedgerAccountMappingsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerAccountMappingsDTOFieldName.Id):
		return string(model.LedgerAccountMappingsDBFieldName.Id), true

	case string(LedgerAccountMappingsDTOFieldName.MappingCode):
		return string(model.LedgerAccountMappingsDBFieldName.MappingCode), true

	case string(LedgerAccountMappingsDTOFieldName.BookId):
		return string(model.LedgerAccountMappingsDBFieldName.BookId), true

	case string(LedgerAccountMappingsDTOFieldName.SourceType):
		return string(model.LedgerAccountMappingsDBFieldName.SourceType), true

	case string(LedgerAccountMappingsDTOFieldName.SourceSubtype):
		return string(model.LedgerAccountMappingsDBFieldName.SourceSubtype), true

	case string(LedgerAccountMappingsDTOFieldName.AccountId):
		return string(model.LedgerAccountMappingsDBFieldName.AccountId), true

	case string(LedgerAccountMappingsDTOFieldName.Priority):
		return string(model.LedgerAccountMappingsDBFieldName.Priority), true

	case string(LedgerAccountMappingsDTOFieldName.IsActive):
		return string(model.LedgerAccountMappingsDBFieldName.IsActive), true

	case string(LedgerAccountMappingsDTOFieldName.Conditions):
		return string(model.LedgerAccountMappingsDBFieldName.Conditions), true

	case string(LedgerAccountMappingsDTOFieldName.Metadata):
		return string(model.LedgerAccountMappingsDBFieldName.Metadata), true

	case string(LedgerAccountMappingsDTOFieldName.MetaCreatedAt):
		return string(model.LedgerAccountMappingsDBFieldName.MetaCreatedAt), true

	case string(LedgerAccountMappingsDTOFieldName.MetaCreatedBy):
		return string(model.LedgerAccountMappingsDBFieldName.MetaCreatedBy), true

	case string(LedgerAccountMappingsDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedAt), true

	case string(LedgerAccountMappingsDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedBy), true

	case string(LedgerAccountMappingsDTOFieldName.MetaDeletedAt):
		return string(model.LedgerAccountMappingsDBFieldName.MetaDeletedAt), true

	case string(LedgerAccountMappingsDTOFieldName.MetaDeletedBy):
		return string(model.LedgerAccountMappingsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerAccountMappingsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerAccountMappingsBaseFilterField(field string) bool {
	spec, found := model.NewLedgerAccountMappingsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerAccountMappingsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerAccountMappingsProjectionOutputPath(path string) error {
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

func transformLedgerAccountMappingsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerAccountMappingsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerAccountMappingsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerAccountMappingsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerAccountMappingsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerAccountMappingsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerAccountMappingsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerAccountMappingsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerAccountMappingsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerAccountMappingsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerAccountMappingsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerAccountMappingsFilter(filter *model.Filter) {
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
			Field: string(LedgerAccountMappingsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerAccountMappingsSelectableResponse map[string]interface{}
type LedgerAccountMappingsSelectableListResponse []*LedgerAccountMappingsSelectableResponse

func assignLedgerAccountMappingsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerAccountMappingsSelectableValue(out LedgerAccountMappingsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerAccountMappingsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerAccountMappingsSelectableResponse(ledgerAccountMappings model.LedgerAccountMappings, filter model.Filter) LedgerAccountMappingsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerAccountMappingsDBFieldName.Id),
			string(model.LedgerAccountMappingsDBFieldName.MappingCode),
			string(model.LedgerAccountMappingsDBFieldName.BookId),
			string(model.LedgerAccountMappingsDBFieldName.SourceType),
			string(model.LedgerAccountMappingsDBFieldName.SourceSubtype),
			string(model.LedgerAccountMappingsDBFieldName.AccountId),
			string(model.LedgerAccountMappingsDBFieldName.Priority),
			string(model.LedgerAccountMappingsDBFieldName.IsActive),
			string(model.LedgerAccountMappingsDBFieldName.Conditions),
			string(model.LedgerAccountMappingsDBFieldName.Metadata),
			string(model.LedgerAccountMappingsDBFieldName.MetaCreatedAt),
			string(model.LedgerAccountMappingsDBFieldName.MetaCreatedBy),
			string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedAt),
			string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedBy),
			string(model.LedgerAccountMappingsDBFieldName.MetaDeletedAt),
			string(model.LedgerAccountMappingsDBFieldName.MetaDeletedBy),
		)
	}
	ledgerAccountMappingsSelectableResponse := LedgerAccountMappingsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerAccountMappingsDBFieldName.Id):
			key := string(LedgerAccountMappingsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.Id, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MappingCode):
			key := string(LedgerAccountMappingsDTOFieldName.MappingCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MappingCode, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.BookId):
			key := string(LedgerAccountMappingsDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.BookId, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.SourceType):
			key := string(LedgerAccountMappingsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.SourceType, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.SourceSubtype):
			key := string(LedgerAccountMappingsDTOFieldName.SourceSubtype)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.SourceSubtype.String, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.AccountId):
			key := string(LedgerAccountMappingsDTOFieldName.AccountId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.AccountId, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.Priority):
			key := string(LedgerAccountMappingsDTOFieldName.Priority)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.Priority, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.IsActive):
			key := string(LedgerAccountMappingsDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.IsActive, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.Conditions):
			key := string(LedgerAccountMappingsDTOFieldName.Conditions)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.Conditions, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.Metadata):
			key := string(LedgerAccountMappingsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.Metadata, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaCreatedAt):
			key := string(LedgerAccountMappingsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaCreatedAt, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaCreatedBy):
			key := string(LedgerAccountMappingsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaCreatedBy, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedAt):
			key := string(LedgerAccountMappingsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaUpdatedBy):
			key := string(LedgerAccountMappingsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaDeletedAt):
			key := string(LedgerAccountMappingsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerAccountMappingsDBFieldName.MetaDeletedBy):
			key := string(LedgerAccountMappingsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountMappingsSelectableValue(ledgerAccountMappingsSelectableResponse, key, ledgerAccountMappings.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerAccountMappingsSelectableResponse
}

func NewLedgerAccountMappingsListResponseFromFilterResult(result []model.LedgerAccountMappingsFilterResult, filter model.Filter) LedgerAccountMappingsSelectableListResponse {
	dtoLedgerAccountMappingsListResponse := LedgerAccountMappingsSelectableListResponse{}
	for _, row := range result {
		dtoLedgerAccountMappingsResponse := NewLedgerAccountMappingsSelectableResponse(row.LedgerAccountMappings, filter)
		dtoLedgerAccountMappingsListResponse = append(dtoLedgerAccountMappingsListResponse, &dtoLedgerAccountMappingsResponse)
	}
	return dtoLedgerAccountMappingsListResponse
}

type LedgerAccountMappingsFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     LedgerAccountMappingsSelectableListResponse `json:"data"`
}

func reverseLedgerAccountMappingsFilterResults(result []model.LedgerAccountMappingsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerAccountMappingsFilterResponse(result []model.LedgerAccountMappingsFilterResult, filter model.Filter) (resp LedgerAccountMappingsFilterResponse) {
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
			reverseLedgerAccountMappingsFilterResults(dataResult)
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

	resp.Data = NewLedgerAccountMappingsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerAccountMappingsCreateRequest struct {
	MappingCode   string          `json:"mappingCode"`
	BookId        uuid.UUID       `json:"bookId"`
	SourceType    string          `json:"sourceType"`
	SourceSubtype string          `json:"sourceSubtype"`
	AccountId     uuid.UUID       `json:"accountId"`
	Priority      int             `json:"priority"`
	IsActive      bool            `json:"isActive"`
	Conditions    json.RawMessage `json:"conditions"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *LedgerAccountMappingsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerAccountMappingsCreateRequest) ToModel() model.LedgerAccountMappings {
	id, _ := uuid.NewV4()
	return model.LedgerAccountMappings{
		Id:            id,
		MappingCode:   d.MappingCode,
		BookId:        d.BookId,
		SourceType:    d.SourceType,
		SourceSubtype: null.StringFrom(d.SourceSubtype),
		AccountId:     d.AccountId,
		Priority:      d.Priority,
		IsActive:      d.IsActive,
		Conditions:    d.Conditions,
		Metadata:      d.Metadata,
	}
}

type LedgerAccountMappingsListCreateRequest []*LedgerAccountMappingsCreateRequest

func (d LedgerAccountMappingsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountMappings := range d {
		err = validator.Struct(ledgerAccountMappings)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountMappingsListCreateRequest) ToModelList() []model.LedgerAccountMappings {
	out := make([]model.LedgerAccountMappings, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerAccountMappingsUpdateRequest struct {
	MappingCode   string          `json:"mappingCode"`
	BookId        uuid.UUID       `json:"bookId"`
	SourceType    string          `json:"sourceType"`
	SourceSubtype string          `json:"sourceSubtype"`
	AccountId     uuid.UUID       `json:"accountId"`
	Priority      int             `json:"priority"`
	IsActive      bool            `json:"isActive"`
	Conditions    json.RawMessage `json:"conditions"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *LedgerAccountMappingsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerAccountMappingsUpdateRequest) ToModel() model.LedgerAccountMappings {
	return model.LedgerAccountMappings{
		MappingCode:   d.MappingCode,
		BookId:        d.BookId,
		SourceType:    d.SourceType,
		SourceSubtype: null.StringFrom(d.SourceSubtype),
		AccountId:     d.AccountId,
		Priority:      d.Priority,
		IsActive:      d.IsActive,
		Conditions:    d.Conditions,
		Metadata:      d.Metadata,
	}
}

type LedgerAccountMappingsBulkUpdateRequest struct {
	Id            uuid.UUID       `json:"id"`
	MappingCode   string          `json:"mappingCode"`
	BookId        uuid.UUID       `json:"bookId"`
	SourceType    string          `json:"sourceType"`
	SourceSubtype string          `json:"sourceSubtype"`
	AccountId     uuid.UUID       `json:"accountId"`
	Priority      int             `json:"priority"`
	IsActive      bool            `json:"isActive"`
	Conditions    json.RawMessage `json:"conditions"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d LedgerAccountMappingsBulkUpdateRequest) PrimaryID() LedgerAccountMappingsPrimaryID {
	return LedgerAccountMappingsPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountMappingsListBulkUpdateRequest []*LedgerAccountMappingsBulkUpdateRequest

func (d LedgerAccountMappingsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountMappings := range d {
		err = validator.Struct(ledgerAccountMappings)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountMappingsBulkUpdateRequest) ToModel() model.LedgerAccountMappings {
	return model.LedgerAccountMappings{
		Id:            d.Id,
		MappingCode:   d.MappingCode,
		BookId:        d.BookId,
		SourceType:    d.SourceType,
		SourceSubtype: null.StringFrom(d.SourceSubtype),
		AccountId:     d.AccountId,
		Priority:      d.Priority,
		IsActive:      d.IsActive,
		Conditions:    d.Conditions,
		Metadata:      d.Metadata,
	}
}

type LedgerAccountMappingsResponse struct {
	Id            uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MappingCode   string          `json:"mappingCode" validate:"required"`
	BookId        uuid.UUID       `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType    string          `json:"sourceType" validate:"required"`
	SourceSubtype string          `json:"sourceSubtype"`
	AccountId     uuid.UUID       `json:"accountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Priority      int             `json:"priority" example:"1"`
	IsActive      bool            `json:"isActive" example:"true"`
	Conditions    json.RawMessage `json:"conditions" swaggertype:"object"`
	Metadata      json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewLedgerAccountMappingsResponse(ledgerAccountMappings model.LedgerAccountMappings) LedgerAccountMappingsResponse {
	return LedgerAccountMappingsResponse{
		Id:            ledgerAccountMappings.Id,
		MappingCode:   ledgerAccountMappings.MappingCode,
		BookId:        ledgerAccountMappings.BookId,
		SourceType:    ledgerAccountMappings.SourceType,
		SourceSubtype: ledgerAccountMappings.SourceSubtype.String,
		AccountId:     ledgerAccountMappings.AccountId,
		Priority:      ledgerAccountMappings.Priority,
		IsActive:      ledgerAccountMappings.IsActive,
		Conditions:    ledgerAccountMappings.Conditions,
		Metadata:      ledgerAccountMappings.Metadata,
	}
}

type LedgerAccountMappingsListResponse []*LedgerAccountMappingsResponse

func NewLedgerAccountMappingsListResponse(ledgerAccountMappingsList model.LedgerAccountMappingsList) LedgerAccountMappingsListResponse {
	dtoLedgerAccountMappingsListResponse := LedgerAccountMappingsListResponse{}
	for _, ledgerAccountMappings := range ledgerAccountMappingsList {
		dtoLedgerAccountMappingsResponse := NewLedgerAccountMappingsResponse(*ledgerAccountMappings)
		dtoLedgerAccountMappingsListResponse = append(dtoLedgerAccountMappingsListResponse, &dtoLedgerAccountMappingsResponse)
	}
	return dtoLedgerAccountMappingsListResponse
}

type LedgerAccountMappingsPrimaryIDList []LedgerAccountMappingsPrimaryID

func (d LedgerAccountMappingsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountMappings := range d {
		err = validator.Struct(ledgerAccountMappings)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerAccountMappingsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerAccountMappingsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerAccountMappingsPrimaryID) ToModel() model.LedgerAccountMappingsPrimaryID {
	return model.LedgerAccountMappingsPrimaryID{
		Id: d.Id,
	}
}

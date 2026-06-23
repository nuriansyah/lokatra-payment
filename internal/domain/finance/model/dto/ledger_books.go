package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerBooksDTOFieldNameType string

type ledgerBooksDTOFieldName struct {
	Id            LedgerBooksDTOFieldNameType
	BookCode      LedgerBooksDTOFieldNameType
	BookTypeCode  LedgerBooksDTOFieldNameType
	OwnerPartyId  LedgerBooksDTOFieldNameType
	CurrencyCode  LedgerBooksDTOFieldNameType
	BookStatus    LedgerBooksDTOFieldNameType
	CloseCutoffTz LedgerBooksDTOFieldNameType
	Metadata      LedgerBooksDTOFieldNameType
	MetaCreatedAt LedgerBooksDTOFieldNameType
	MetaCreatedBy LedgerBooksDTOFieldNameType
	MetaUpdatedAt LedgerBooksDTOFieldNameType
	MetaUpdatedBy LedgerBooksDTOFieldNameType
	MetaDeletedAt LedgerBooksDTOFieldNameType
	MetaDeletedBy LedgerBooksDTOFieldNameType
}

var LedgerBooksDTOFieldName = ledgerBooksDTOFieldName{
	Id:            "id",
	BookCode:      "bookCode",
	BookTypeCode:  "bookTypeCode",
	OwnerPartyId:  "ownerPartyId",
	CurrencyCode:  "currencyCode",
	BookStatus:    "bookStatus",
	CloseCutoffTz: "closeCutoffTz",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformLedgerBooksDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerBooksDTOFieldName.Id):
		return string(model.LedgerBooksDBFieldName.Id), true

	case string(LedgerBooksDTOFieldName.BookCode):
		return string(model.LedgerBooksDBFieldName.BookCode), true

	case string(LedgerBooksDTOFieldName.BookTypeCode):
		return string(model.LedgerBooksDBFieldName.BookTypeCode), true

	case string(LedgerBooksDTOFieldName.OwnerPartyId):
		return string(model.LedgerBooksDBFieldName.OwnerPartyId), true

	case string(LedgerBooksDTOFieldName.CurrencyCode):
		return string(model.LedgerBooksDBFieldName.CurrencyCode), true

	case string(LedgerBooksDTOFieldName.BookStatus):
		return string(model.LedgerBooksDBFieldName.BookStatus), true

	case string(LedgerBooksDTOFieldName.CloseCutoffTz):
		return string(model.LedgerBooksDBFieldName.CloseCutoffTz), true

	case string(LedgerBooksDTOFieldName.Metadata):
		return string(model.LedgerBooksDBFieldName.Metadata), true

	case string(LedgerBooksDTOFieldName.MetaCreatedAt):
		return string(model.LedgerBooksDBFieldName.MetaCreatedAt), true

	case string(LedgerBooksDTOFieldName.MetaCreatedBy):
		return string(model.LedgerBooksDBFieldName.MetaCreatedBy), true

	case string(LedgerBooksDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerBooksDBFieldName.MetaUpdatedAt), true

	case string(LedgerBooksDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerBooksDBFieldName.MetaUpdatedBy), true

	case string(LedgerBooksDTOFieldName.MetaDeletedAt):
		return string(model.LedgerBooksDBFieldName.MetaDeletedAt), true

	case string(LedgerBooksDTOFieldName.MetaDeletedBy):
		return string(model.LedgerBooksDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerBooksFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerBooksBaseFilterField(field string) bool {
	spec, found := model.NewLedgerBooksFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerBooksProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerBooksProjectionOutputPath(path string) error {
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

func transformLedgerBooksFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerBooksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerBooksFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerBooksFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerBooksDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerBooksBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerBooksProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerBooksProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerBooksDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerBooksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerBooksFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerBooksFilter(filter *model.Filter) {
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
			Field: string(LedgerBooksDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerBooksSelectableResponse map[string]interface{}
type LedgerBooksSelectableListResponse []*LedgerBooksSelectableResponse

func assignLedgerBooksNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerBooksSelectableValue(out LedgerBooksSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerBooksNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerBooksSelectableResponse(ledgerBooks model.LedgerBooks, filter model.Filter) LedgerBooksSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerBooksDBFieldName.Id),
			string(model.LedgerBooksDBFieldName.BookCode),
			string(model.LedgerBooksDBFieldName.BookTypeCode),
			string(model.LedgerBooksDBFieldName.OwnerPartyId),
			string(model.LedgerBooksDBFieldName.CurrencyCode),
			string(model.LedgerBooksDBFieldName.BookStatus),
			string(model.LedgerBooksDBFieldName.CloseCutoffTz),
			string(model.LedgerBooksDBFieldName.Metadata),
			string(model.LedgerBooksDBFieldName.MetaCreatedAt),
			string(model.LedgerBooksDBFieldName.MetaCreatedBy),
			string(model.LedgerBooksDBFieldName.MetaUpdatedAt),
			string(model.LedgerBooksDBFieldName.MetaUpdatedBy),
			string(model.LedgerBooksDBFieldName.MetaDeletedAt),
			string(model.LedgerBooksDBFieldName.MetaDeletedBy),
		)
	}
	ledgerBooksSelectableResponse := LedgerBooksSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerBooksDBFieldName.Id):
			key := string(LedgerBooksDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.Id, explicitAlias)

		case string(model.LedgerBooksDBFieldName.BookCode):
			key := string(LedgerBooksDTOFieldName.BookCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.BookCode, explicitAlias)

		case string(model.LedgerBooksDBFieldName.BookTypeCode):
			key := string(LedgerBooksDTOFieldName.BookTypeCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.BookTypeCode, explicitAlias)

		case string(model.LedgerBooksDBFieldName.OwnerPartyId):
			key := string(LedgerBooksDTOFieldName.OwnerPartyId)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.OwnerPartyId, explicitAlias)

		case string(model.LedgerBooksDBFieldName.CurrencyCode):
			key := string(LedgerBooksDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.CurrencyCode, explicitAlias)

		case string(model.LedgerBooksDBFieldName.BookStatus):
			key := string(LedgerBooksDTOFieldName.BookStatus)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, model.BookStatus(ledgerBooks.BookStatus), explicitAlias)

		case string(model.LedgerBooksDBFieldName.CloseCutoffTz):
			key := string(LedgerBooksDTOFieldName.CloseCutoffTz)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.CloseCutoffTz, explicitAlias)

		case string(model.LedgerBooksDBFieldName.Metadata):
			key := string(LedgerBooksDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.Metadata, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaCreatedAt):
			key := string(LedgerBooksDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaCreatedAt, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaCreatedBy):
			key := string(LedgerBooksDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaCreatedBy, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaUpdatedAt):
			key := string(LedgerBooksDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaUpdatedBy):
			key := string(LedgerBooksDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaDeletedAt):
			key := string(LedgerBooksDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerBooksDBFieldName.MetaDeletedBy):
			key := string(LedgerBooksDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerBooksSelectableValue(ledgerBooksSelectableResponse, key, ledgerBooks.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerBooksSelectableResponse
}

func NewLedgerBooksListResponseFromFilterResult(result []model.LedgerBooksFilterResult, filter model.Filter) LedgerBooksSelectableListResponse {
	dtoLedgerBooksListResponse := LedgerBooksSelectableListResponse{}
	for _, row := range result {
		dtoLedgerBooksResponse := NewLedgerBooksSelectableResponse(row.LedgerBooks, filter)
		dtoLedgerBooksListResponse = append(dtoLedgerBooksListResponse, &dtoLedgerBooksResponse)
	}
	return dtoLedgerBooksListResponse
}

type LedgerBooksFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     LedgerBooksSelectableListResponse `json:"data"`
}

func reverseLedgerBooksFilterResults(result []model.LedgerBooksFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerBooksFilterResponse(result []model.LedgerBooksFilterResult, filter model.Filter) (resp LedgerBooksFilterResponse) {
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
			reverseLedgerBooksFilterResults(dataResult)
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

	resp.Data = NewLedgerBooksListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerBooksCreateRequest struct {
	BookCode      string           `json:"bookCode"`
	BookTypeCode  string           `json:"bookTypeCode"`
	OwnerPartyId  uuid.UUID        `json:"ownerPartyId"`
	CurrencyCode  string           `json:"currencyCode"`
	BookStatus    model.BookStatus `json:"bookStatus" example:"active" enums:"active,locked,closed"`
	CloseCutoffTz string           `json:"closeCutoffTz"`
	Metadata      json.RawMessage  `json:"metadata"`
}

func (d *LedgerBooksCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerBooksCreateRequest) ToModel() model.LedgerBooks {
	id, _ := uuid.NewV4()
	return model.LedgerBooks{
		Id:            id,
		BookCode:      d.BookCode,
		BookTypeCode:  d.BookTypeCode,
		OwnerPartyId:  d.OwnerPartyId,
		CurrencyCode:  d.CurrencyCode,
		BookStatus:    d.BookStatus,
		CloseCutoffTz: d.CloseCutoffTz,
		Metadata:      d.Metadata,
	}
}

type LedgerBooksListCreateRequest []*LedgerBooksCreateRequest

func (d LedgerBooksListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBooks := range d {
		err = validator.Struct(ledgerBooks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerBooksListCreateRequest) ToModelList() []model.LedgerBooks {
	out := make([]model.LedgerBooks, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerBooksUpdateRequest struct {
	BookCode      string           `json:"bookCode"`
	BookTypeCode  string           `json:"bookTypeCode"`
	OwnerPartyId  uuid.UUID        `json:"ownerPartyId"`
	CurrencyCode  string           `json:"currencyCode"`
	BookStatus    model.BookStatus `json:"bookStatus" example:"active" enums:"active,locked,closed"`
	CloseCutoffTz string           `json:"closeCutoffTz"`
	Metadata      json.RawMessage  `json:"metadata"`
}

func (d *LedgerBooksUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerBooksUpdateRequest) ToModel() model.LedgerBooks {
	return model.LedgerBooks{
		BookCode:      d.BookCode,
		BookTypeCode:  d.BookTypeCode,
		OwnerPartyId:  d.OwnerPartyId,
		CurrencyCode:  d.CurrencyCode,
		BookStatus:    d.BookStatus,
		CloseCutoffTz: d.CloseCutoffTz,
		Metadata:      d.Metadata,
	}
}

type LedgerBooksBulkUpdateRequest struct {
	Id            uuid.UUID        `json:"id"`
	BookCode      string           `json:"bookCode"`
	BookTypeCode  string           `json:"bookTypeCode"`
	OwnerPartyId  uuid.UUID        `json:"ownerPartyId"`
	CurrencyCode  string           `json:"currencyCode"`
	BookStatus    model.BookStatus `json:"bookStatus" example:"active" enums:"active,locked,closed"`
	CloseCutoffTz string           `json:"closeCutoffTz"`
	Metadata      json.RawMessage  `json:"metadata"`
}

func (d LedgerBooksBulkUpdateRequest) PrimaryID() LedgerBooksPrimaryID {
	return LedgerBooksPrimaryID{
		Id: d.Id,
	}
}

type LedgerBooksListBulkUpdateRequest []*LedgerBooksBulkUpdateRequest

func (d LedgerBooksListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBooks := range d {
		err = validator.Struct(ledgerBooks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerBooksBulkUpdateRequest) ToModel() model.LedgerBooks {
	return model.LedgerBooks{
		Id:            d.Id,
		BookCode:      d.BookCode,
		BookTypeCode:  d.BookTypeCode,
		OwnerPartyId:  d.OwnerPartyId,
		CurrencyCode:  d.CurrencyCode,
		BookStatus:    d.BookStatus,
		CloseCutoffTz: d.CloseCutoffTz,
		Metadata:      d.Metadata,
	}
}

type LedgerBooksResponse struct {
	Id            uuid.UUID        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BookCode      string           `json:"bookCode" validate:"required"`
	BookTypeCode  string           `json:"bookTypeCode" validate:"required"`
	OwnerPartyId  uuid.UUID        `json:"ownerPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode  string           `json:"currencyCode" validate:"required"`
	BookStatus    model.BookStatus `json:"bookStatus" validate:"oneof=active locked closed" enums:"active,locked,closed"`
	CloseCutoffTz string           `json:"closeCutoffTz"`
	Metadata      json.RawMessage  `json:"metadata" swaggertype:"object"`
}

func NewLedgerBooksResponse(ledgerBooks model.LedgerBooks) LedgerBooksResponse {
	return LedgerBooksResponse{
		Id:            ledgerBooks.Id,
		BookCode:      ledgerBooks.BookCode,
		BookTypeCode:  ledgerBooks.BookTypeCode,
		OwnerPartyId:  ledgerBooks.OwnerPartyId,
		CurrencyCode:  ledgerBooks.CurrencyCode,
		BookStatus:    model.BookStatus(ledgerBooks.BookStatus),
		CloseCutoffTz: ledgerBooks.CloseCutoffTz,
		Metadata:      ledgerBooks.Metadata,
	}
}

type LedgerBooksListResponse []*LedgerBooksResponse

func NewLedgerBooksListResponse(ledgerBooksList model.LedgerBooksList) LedgerBooksListResponse {
	dtoLedgerBooksListResponse := LedgerBooksListResponse{}
	for _, ledgerBooks := range ledgerBooksList {
		dtoLedgerBooksResponse := NewLedgerBooksResponse(*ledgerBooks)
		dtoLedgerBooksListResponse = append(dtoLedgerBooksListResponse, &dtoLedgerBooksResponse)
	}
	return dtoLedgerBooksListResponse
}

type LedgerBooksPrimaryIDList []LedgerBooksPrimaryID

func (d LedgerBooksPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerBooks := range d {
		err = validator.Struct(ledgerBooks)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerBooksPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerBooksPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerBooksPrimaryID) ToModel() model.LedgerBooksPrimaryID {
	return model.LedgerBooksPrimaryID{
		Id: d.Id,
	}
}

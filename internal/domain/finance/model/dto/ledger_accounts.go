package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerAccountsDTOFieldNameType string

type ledgerAccountsDTOFieldName struct {
	Id                 LedgerAccountsDTOFieldNameType
	BookId             LedgerAccountsDTOFieldNameType
	AccountCode        LedgerAccountsDTOFieldNameType
	AccountName        LedgerAccountsDTOFieldNameType
	AccountTypeCode    LedgerAccountsDTOFieldNameType
	OwnerPartyId       LedgerAccountsDTOFieldNameType
	ParentAccountId    LedgerAccountsDTOFieldNameType
	CurrencyCode       LedgerAccountsDTOFieldNameType
	AllowManualPosting LedgerAccountsDTOFieldNameType
	AccountStatus      LedgerAccountsDTOFieldNameType
	Metadata           LedgerAccountsDTOFieldNameType
	MetaCreatedAt      LedgerAccountsDTOFieldNameType
	MetaCreatedBy      LedgerAccountsDTOFieldNameType
	MetaUpdatedAt      LedgerAccountsDTOFieldNameType
	MetaUpdatedBy      LedgerAccountsDTOFieldNameType
	MetaDeletedAt      LedgerAccountsDTOFieldNameType
	MetaDeletedBy      LedgerAccountsDTOFieldNameType
}

var LedgerAccountsDTOFieldName = ledgerAccountsDTOFieldName{
	Id:                 "id",
	BookId:             "bookId",
	AccountCode:        "accountCode",
	AccountName:        "accountName",
	AccountTypeCode:    "accountTypeCode",
	OwnerPartyId:       "ownerPartyId",
	ParentAccountId:    "parentAccountId",
	CurrencyCode:       "currencyCode",
	AllowManualPosting: "allowManualPosting",
	AccountStatus:      "accountStatus",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformLedgerAccountsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerAccountsDTOFieldName.Id):
		return string(model.LedgerAccountsDBFieldName.Id), true

	case string(LedgerAccountsDTOFieldName.BookId):
		return string(model.LedgerAccountsDBFieldName.BookId), true

	case string(LedgerAccountsDTOFieldName.AccountCode):
		return string(model.LedgerAccountsDBFieldName.AccountCode), true

	case string(LedgerAccountsDTOFieldName.AccountName):
		return string(model.LedgerAccountsDBFieldName.AccountName), true

	case string(LedgerAccountsDTOFieldName.AccountTypeCode):
		return string(model.LedgerAccountsDBFieldName.AccountTypeCode), true

	case string(LedgerAccountsDTOFieldName.OwnerPartyId):
		return string(model.LedgerAccountsDBFieldName.OwnerPartyId), true

	case string(LedgerAccountsDTOFieldName.ParentAccountId):
		return string(model.LedgerAccountsDBFieldName.ParentAccountId), true

	case string(LedgerAccountsDTOFieldName.CurrencyCode):
		return string(model.LedgerAccountsDBFieldName.CurrencyCode), true

	case string(LedgerAccountsDTOFieldName.AllowManualPosting):
		return string(model.LedgerAccountsDBFieldName.AllowManualPosting), true

	case string(LedgerAccountsDTOFieldName.AccountStatus):
		return string(model.LedgerAccountsDBFieldName.AccountStatus), true

	case string(LedgerAccountsDTOFieldName.Metadata):
		return string(model.LedgerAccountsDBFieldName.Metadata), true

	case string(LedgerAccountsDTOFieldName.MetaCreatedAt):
		return string(model.LedgerAccountsDBFieldName.MetaCreatedAt), true

	case string(LedgerAccountsDTOFieldName.MetaCreatedBy):
		return string(model.LedgerAccountsDBFieldName.MetaCreatedBy), true

	case string(LedgerAccountsDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerAccountsDBFieldName.MetaUpdatedAt), true

	case string(LedgerAccountsDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerAccountsDBFieldName.MetaUpdatedBy), true

	case string(LedgerAccountsDTOFieldName.MetaDeletedAt):
		return string(model.LedgerAccountsDBFieldName.MetaDeletedAt), true

	case string(LedgerAccountsDTOFieldName.MetaDeletedBy):
		return string(model.LedgerAccountsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerAccountsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerAccountsBaseFilterField(field string) bool {
	spec, found := model.NewLedgerAccountsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerAccountsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerAccountsProjectionOutputPath(path string) error {
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

func transformLedgerAccountsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerAccountsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerAccountsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerAccountsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerAccountsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerAccountsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerAccountsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerAccountsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerAccountsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerAccountsFilter(filter *model.Filter) {
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
			Field: string(LedgerAccountsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerAccountsSelectableResponse map[string]interface{}
type LedgerAccountsSelectableListResponse []*LedgerAccountsSelectableResponse

func assignLedgerAccountsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerAccountsSelectableValue(out LedgerAccountsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerAccountsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerAccountsSelectableResponse(ledgerAccounts model.LedgerAccounts, filter model.Filter) LedgerAccountsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerAccountsDBFieldName.Id),
			string(model.LedgerAccountsDBFieldName.BookId),
			string(model.LedgerAccountsDBFieldName.AccountCode),
			string(model.LedgerAccountsDBFieldName.AccountName),
			string(model.LedgerAccountsDBFieldName.AccountTypeCode),
			string(model.LedgerAccountsDBFieldName.OwnerPartyId),
			string(model.LedgerAccountsDBFieldName.ParentAccountId),
			string(model.LedgerAccountsDBFieldName.CurrencyCode),
			string(model.LedgerAccountsDBFieldName.AllowManualPosting),
			string(model.LedgerAccountsDBFieldName.AccountStatus),
			string(model.LedgerAccountsDBFieldName.Metadata),
			string(model.LedgerAccountsDBFieldName.MetaCreatedAt),
			string(model.LedgerAccountsDBFieldName.MetaCreatedBy),
			string(model.LedgerAccountsDBFieldName.MetaUpdatedAt),
			string(model.LedgerAccountsDBFieldName.MetaUpdatedBy),
			string(model.LedgerAccountsDBFieldName.MetaDeletedAt),
			string(model.LedgerAccountsDBFieldName.MetaDeletedBy),
		)
	}
	ledgerAccountsSelectableResponse := LedgerAccountsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerAccountsDBFieldName.Id):
			key := string(LedgerAccountsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.Id, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.BookId):
			key := string(LedgerAccountsDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.BookId, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.AccountCode):
			key := string(LedgerAccountsDTOFieldName.AccountCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.AccountCode, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.AccountName):
			key := string(LedgerAccountsDTOFieldName.AccountName)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.AccountName, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.AccountTypeCode):
			key := string(LedgerAccountsDTOFieldName.AccountTypeCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.AccountTypeCode, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.OwnerPartyId):
			key := string(LedgerAccountsDTOFieldName.OwnerPartyId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.OwnerPartyId.UUID, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.ParentAccountId):
			key := string(LedgerAccountsDTOFieldName.ParentAccountId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.ParentAccountId.UUID, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.CurrencyCode):
			key := string(LedgerAccountsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.CurrencyCode, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.AllowManualPosting):
			key := string(LedgerAccountsDTOFieldName.AllowManualPosting)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.AllowManualPosting, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.AccountStatus):
			key := string(LedgerAccountsDTOFieldName.AccountStatus)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, model.LedgerAccountsAccountStatus(ledgerAccounts.AccountStatus), explicitAlias)

		case string(model.LedgerAccountsDBFieldName.Metadata):
			key := string(LedgerAccountsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.Metadata, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaCreatedAt):
			key := string(LedgerAccountsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaCreatedAt, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaCreatedBy):
			key := string(LedgerAccountsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaCreatedBy, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaUpdatedAt):
			key := string(LedgerAccountsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaUpdatedBy):
			key := string(LedgerAccountsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaDeletedAt):
			key := string(LedgerAccountsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerAccountsDBFieldName.MetaDeletedBy):
			key := string(LedgerAccountsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountsSelectableValue(ledgerAccountsSelectableResponse, key, ledgerAccounts.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerAccountsSelectableResponse
}

func NewLedgerAccountsListResponseFromFilterResult(result []model.LedgerAccountsFilterResult, filter model.Filter) LedgerAccountsSelectableListResponse {
	dtoLedgerAccountsListResponse := LedgerAccountsSelectableListResponse{}
	for _, row := range result {
		dtoLedgerAccountsResponse := NewLedgerAccountsSelectableResponse(row.LedgerAccounts, filter)
		dtoLedgerAccountsListResponse = append(dtoLedgerAccountsListResponse, &dtoLedgerAccountsResponse)
	}
	return dtoLedgerAccountsListResponse
}

type LedgerAccountsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     LedgerAccountsSelectableListResponse `json:"data"`
}

func reverseLedgerAccountsFilterResults(result []model.LedgerAccountsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerAccountsFilterResponse(result []model.LedgerAccountsFilterResult, filter model.Filter) (resp LedgerAccountsFilterResponse) {
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
			reverseLedgerAccountsFilterResults(dataResult)
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

	resp.Data = NewLedgerAccountsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerAccountsCreateRequest struct {
	BookId             uuid.UUID                         `json:"bookId"`
	AccountCode        string                            `json:"accountCode"`
	AccountName        string                            `json:"accountName"`
	AccountTypeCode    string                            `json:"accountTypeCode"`
	OwnerPartyId       uuid.UUID                         `json:"ownerPartyId"`
	ParentAccountId    uuid.UUID                         `json:"parentAccountId"`
	CurrencyCode       string                            `json:"currencyCode"`
	AllowManualPosting bool                              `json:"allowManualPosting"`
	AccountStatus      model.LedgerAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive,closed"`
	Metadata           json.RawMessage                   `json:"metadata"`
}

func (d *LedgerAccountsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerAccountsCreateRequest) ToModel() model.LedgerAccounts {
	id, _ := uuid.NewV4()
	return model.LedgerAccounts{
		Id:                 id,
		BookId:             d.BookId,
		AccountCode:        d.AccountCode,
		AccountName:        d.AccountName,
		AccountTypeCode:    d.AccountTypeCode,
		OwnerPartyId:       nuuid.From(d.OwnerPartyId),
		ParentAccountId:    nuuid.From(d.ParentAccountId),
		CurrencyCode:       d.CurrencyCode,
		AllowManualPosting: d.AllowManualPosting,
		AccountStatus:      d.AccountStatus,
		Metadata:           d.Metadata,
	}
}

type LedgerAccountsListCreateRequest []*LedgerAccountsCreateRequest

func (d LedgerAccountsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccounts := range d {
		err = validator.Struct(ledgerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountsListCreateRequest) ToModelList() []model.LedgerAccounts {
	out := make([]model.LedgerAccounts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerAccountsUpdateRequest struct {
	BookId             uuid.UUID                         `json:"bookId"`
	AccountCode        string                            `json:"accountCode"`
	AccountName        string                            `json:"accountName"`
	AccountTypeCode    string                            `json:"accountTypeCode"`
	OwnerPartyId       uuid.UUID                         `json:"ownerPartyId"`
	ParentAccountId    uuid.UUID                         `json:"parentAccountId"`
	CurrencyCode       string                            `json:"currencyCode"`
	AllowManualPosting bool                              `json:"allowManualPosting"`
	AccountStatus      model.LedgerAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive,closed"`
	Metadata           json.RawMessage                   `json:"metadata"`
}

func (d *LedgerAccountsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerAccountsUpdateRequest) ToModel() model.LedgerAccounts {
	return model.LedgerAccounts{
		BookId:             d.BookId,
		AccountCode:        d.AccountCode,
		AccountName:        d.AccountName,
		AccountTypeCode:    d.AccountTypeCode,
		OwnerPartyId:       nuuid.From(d.OwnerPartyId),
		ParentAccountId:    nuuid.From(d.ParentAccountId),
		CurrencyCode:       d.CurrencyCode,
		AllowManualPosting: d.AllowManualPosting,
		AccountStatus:      d.AccountStatus,
		Metadata:           d.Metadata,
	}
}

type LedgerAccountsBulkUpdateRequest struct {
	Id                 uuid.UUID                         `json:"id"`
	BookId             uuid.UUID                         `json:"bookId"`
	AccountCode        string                            `json:"accountCode"`
	AccountName        string                            `json:"accountName"`
	AccountTypeCode    string                            `json:"accountTypeCode"`
	OwnerPartyId       uuid.UUID                         `json:"ownerPartyId"`
	ParentAccountId    uuid.UUID                         `json:"parentAccountId"`
	CurrencyCode       string                            `json:"currencyCode"`
	AllowManualPosting bool                              `json:"allowManualPosting"`
	AccountStatus      model.LedgerAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive,closed"`
	Metadata           json.RawMessage                   `json:"metadata"`
}

func (d LedgerAccountsBulkUpdateRequest) PrimaryID() LedgerAccountsPrimaryID {
	return LedgerAccountsPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountsListBulkUpdateRequest []*LedgerAccountsBulkUpdateRequest

func (d LedgerAccountsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccounts := range d {
		err = validator.Struct(ledgerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountsBulkUpdateRequest) ToModel() model.LedgerAccounts {
	return model.LedgerAccounts{
		Id:                 d.Id,
		BookId:             d.BookId,
		AccountCode:        d.AccountCode,
		AccountName:        d.AccountName,
		AccountTypeCode:    d.AccountTypeCode,
		OwnerPartyId:       nuuid.From(d.OwnerPartyId),
		ParentAccountId:    nuuid.From(d.ParentAccountId),
		CurrencyCode:       d.CurrencyCode,
		AllowManualPosting: d.AllowManualPosting,
		AccountStatus:      d.AccountStatus,
		Metadata:           d.Metadata,
	}
}

type LedgerAccountsResponse struct {
	Id                 uuid.UUID                         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BookId             uuid.UUID                         `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AccountCode        string                            `json:"accountCode" validate:"required"`
	AccountName        string                            `json:"accountName" validate:"required"`
	AccountTypeCode    string                            `json:"accountTypeCode" validate:"required"`
	OwnerPartyId       uuid.UUID                         `json:"ownerPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ParentAccountId    uuid.UUID                         `json:"parentAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode       string                            `json:"currencyCode" validate:"required"`
	AllowManualPosting bool                              `json:"allowManualPosting" example:"true"`
	AccountStatus      model.LedgerAccountsAccountStatus `json:"accountStatus" validate:"oneof=active inactive closed" enums:"active,inactive,closed"`
	Metadata           json.RawMessage                   `json:"metadata" swaggertype:"object"`
}

func NewLedgerAccountsResponse(ledgerAccounts model.LedgerAccounts) LedgerAccountsResponse {
	return LedgerAccountsResponse{
		Id:                 ledgerAccounts.Id,
		BookId:             ledgerAccounts.BookId,
		AccountCode:        ledgerAccounts.AccountCode,
		AccountName:        ledgerAccounts.AccountName,
		AccountTypeCode:    ledgerAccounts.AccountTypeCode,
		OwnerPartyId:       ledgerAccounts.OwnerPartyId.UUID,
		ParentAccountId:    ledgerAccounts.ParentAccountId.UUID,
		CurrencyCode:       ledgerAccounts.CurrencyCode,
		AllowManualPosting: ledgerAccounts.AllowManualPosting,
		AccountStatus:      model.LedgerAccountsAccountStatus(ledgerAccounts.AccountStatus),
		Metadata:           ledgerAccounts.Metadata,
	}
}

type LedgerAccountsListResponse []*LedgerAccountsResponse

func NewLedgerAccountsListResponse(ledgerAccountsList model.LedgerAccountsList) LedgerAccountsListResponse {
	dtoLedgerAccountsListResponse := LedgerAccountsListResponse{}
	for _, ledgerAccounts := range ledgerAccountsList {
		dtoLedgerAccountsResponse := NewLedgerAccountsResponse(*ledgerAccounts)
		dtoLedgerAccountsListResponse = append(dtoLedgerAccountsListResponse, &dtoLedgerAccountsResponse)
	}
	return dtoLedgerAccountsListResponse
}

type LedgerAccountsPrimaryIDList []LedgerAccountsPrimaryID

func (d LedgerAccountsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccounts := range d {
		err = validator.Struct(ledgerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerAccountsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerAccountsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerAccountsPrimaryID) ToModel() model.LedgerAccountsPrimaryID {
	return model.LedgerAccountsPrimaryID{
		Id: d.Id,
	}
}

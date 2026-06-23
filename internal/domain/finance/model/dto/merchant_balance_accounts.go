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

type MerchantBalanceAccountsDTOFieldNameType string

type merchantBalanceAccountsDTOFieldName struct {
	Id                    MerchantBalanceAccountsDTOFieldNameType
	MerchantPartyId       MerchantBalanceAccountsDTOFieldNameType
	BalanceType           MerchantBalanceAccountsDTOFieldNameType
	CurrencyCode          MerchantBalanceAccountsDTOFieldNameType
	LinkedLedgerAccountId MerchantBalanceAccountsDTOFieldNameType
	AccountStatus         MerchantBalanceAccountsDTOFieldNameType
	Metadata              MerchantBalanceAccountsDTOFieldNameType
	MetaCreatedAt         MerchantBalanceAccountsDTOFieldNameType
	MetaCreatedBy         MerchantBalanceAccountsDTOFieldNameType
	MetaUpdatedAt         MerchantBalanceAccountsDTOFieldNameType
	MetaUpdatedBy         MerchantBalanceAccountsDTOFieldNameType
	MetaDeletedAt         MerchantBalanceAccountsDTOFieldNameType
	MetaDeletedBy         MerchantBalanceAccountsDTOFieldNameType
}

var MerchantBalanceAccountsDTOFieldName = merchantBalanceAccountsDTOFieldName{
	Id:                    "id",
	MerchantPartyId:       "merchantPartyId",
	BalanceType:           "balanceType",
	CurrencyCode:          "currencyCode",
	LinkedLedgerAccountId: "linkedLedgerAccountId",
	AccountStatus:         "accountStatus",
	Metadata:              "metadata",
	MetaCreatedAt:         "metaCreatedAt",
	MetaCreatedBy:         "metaCreatedBy",
	MetaUpdatedAt:         "metaUpdatedAt",
	MetaUpdatedBy:         "metaUpdatedBy",
	MetaDeletedAt:         "metaDeletedAt",
	MetaDeletedBy:         "metaDeletedBy",
}

func transformMerchantBalanceAccountsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(MerchantBalanceAccountsDTOFieldName.Id):
		return string(model.MerchantBalanceAccountsDBFieldName.Id), true

	case string(MerchantBalanceAccountsDTOFieldName.MerchantPartyId):
		return string(model.MerchantBalanceAccountsDBFieldName.MerchantPartyId), true

	case string(MerchantBalanceAccountsDTOFieldName.BalanceType):
		return string(model.MerchantBalanceAccountsDBFieldName.BalanceType), true

	case string(MerchantBalanceAccountsDTOFieldName.CurrencyCode):
		return string(model.MerchantBalanceAccountsDBFieldName.CurrencyCode), true

	case string(MerchantBalanceAccountsDTOFieldName.LinkedLedgerAccountId):
		return string(model.MerchantBalanceAccountsDBFieldName.LinkedLedgerAccountId), true

	case string(MerchantBalanceAccountsDTOFieldName.AccountStatus):
		return string(model.MerchantBalanceAccountsDBFieldName.AccountStatus), true

	case string(MerchantBalanceAccountsDTOFieldName.Metadata):
		return string(model.MerchantBalanceAccountsDBFieldName.Metadata), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaCreatedAt):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedAt), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaCreatedBy):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedBy), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaUpdatedAt):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedAt), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaUpdatedBy):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedBy), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaDeletedAt):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedAt), true

	case string(MerchantBalanceAccountsDTOFieldName.MetaDeletedBy):
		return string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewMerchantBalanceAccountsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isMerchantBalanceAccountsBaseFilterField(field string) bool {
	spec, found := model.NewMerchantBalanceAccountsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeMerchantBalanceAccountsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateMerchantBalanceAccountsProjectionOutputPath(path string) error {
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

func transformMerchantBalanceAccountsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformMerchantBalanceAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformMerchantBalanceAccountsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformMerchantBalanceAccountsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformMerchantBalanceAccountsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isMerchantBalanceAccountsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateMerchantBalanceAccountsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeMerchantBalanceAccountsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformMerchantBalanceAccountsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformMerchantBalanceAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformMerchantBalanceAccountsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultMerchantBalanceAccountsFilter(filter *model.Filter) {
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
			Field: string(MerchantBalanceAccountsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type MerchantBalanceAccountsSelectableResponse map[string]interface{}
type MerchantBalanceAccountsSelectableListResponse []*MerchantBalanceAccountsSelectableResponse

func assignMerchantBalanceAccountsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setMerchantBalanceAccountsSelectableValue(out MerchantBalanceAccountsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignMerchantBalanceAccountsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewMerchantBalanceAccountsSelectableResponse(merchantBalanceAccounts model.MerchantBalanceAccounts, filter model.Filter) MerchantBalanceAccountsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.MerchantBalanceAccountsDBFieldName.Id),
			string(model.MerchantBalanceAccountsDBFieldName.MerchantPartyId),
			string(model.MerchantBalanceAccountsDBFieldName.BalanceType),
			string(model.MerchantBalanceAccountsDBFieldName.CurrencyCode),
			string(model.MerchantBalanceAccountsDBFieldName.LinkedLedgerAccountId),
			string(model.MerchantBalanceAccountsDBFieldName.AccountStatus),
			string(model.MerchantBalanceAccountsDBFieldName.Metadata),
			string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedAt),
			string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedBy),
			string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedAt),
			string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedBy),
			string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedAt),
			string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedBy),
		)
	}
	merchantBalanceAccountsSelectableResponse := MerchantBalanceAccountsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.MerchantBalanceAccountsDBFieldName.Id):
			key := string(MerchantBalanceAccountsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.Id, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MerchantPartyId):
			key := string(MerchantBalanceAccountsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MerchantPartyId, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.BalanceType):
			key := string(MerchantBalanceAccountsDTOFieldName.BalanceType)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, model.MerchantBalanceAccountsBalanceType(merchantBalanceAccounts.BalanceType), explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.CurrencyCode):
			key := string(MerchantBalanceAccountsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.CurrencyCode, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.LinkedLedgerAccountId):
			key := string(MerchantBalanceAccountsDTOFieldName.LinkedLedgerAccountId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.LinkedLedgerAccountId, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.AccountStatus):
			key := string(MerchantBalanceAccountsDTOFieldName.AccountStatus)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, model.MerchantBalanceAccountsAccountStatus(merchantBalanceAccounts.AccountStatus), explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.Metadata):
			key := string(MerchantBalanceAccountsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.Metadata, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedAt):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaCreatedAt, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaCreatedBy):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaCreatedBy, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedAt):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaUpdatedAt, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaUpdatedBy):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaUpdatedBy, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedAt):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaDeletedAt.Time, explicitAlias)

		case string(model.MerchantBalanceAccountsDBFieldName.MetaDeletedBy):
			key := string(MerchantBalanceAccountsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceAccountsSelectableValue(merchantBalanceAccountsSelectableResponse, key, merchantBalanceAccounts.MetaDeletedBy, explicitAlias)

		}
	}
	return merchantBalanceAccountsSelectableResponse
}

func NewMerchantBalanceAccountsListResponseFromFilterResult(result []model.MerchantBalanceAccountsFilterResult, filter model.Filter) MerchantBalanceAccountsSelectableListResponse {
	dtoMerchantBalanceAccountsListResponse := MerchantBalanceAccountsSelectableListResponse{}
	for _, row := range result {
		dtoMerchantBalanceAccountsResponse := NewMerchantBalanceAccountsSelectableResponse(row.MerchantBalanceAccounts, filter)
		dtoMerchantBalanceAccountsListResponse = append(dtoMerchantBalanceAccountsListResponse, &dtoMerchantBalanceAccountsResponse)
	}
	return dtoMerchantBalanceAccountsListResponse
}

type MerchantBalanceAccountsFilterResponse struct {
	Metadata Metadata                                      `json:"metadata"`
	Data     MerchantBalanceAccountsSelectableListResponse `json:"data"`
}

func reverseMerchantBalanceAccountsFilterResults(result []model.MerchantBalanceAccountsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewMerchantBalanceAccountsFilterResponse(result []model.MerchantBalanceAccountsFilterResult, filter model.Filter) (resp MerchantBalanceAccountsFilterResponse) {
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
			reverseMerchantBalanceAccountsFilterResults(dataResult)
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

	resp.Data = NewMerchantBalanceAccountsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type MerchantBalanceAccountsCreateRequest struct {
	MerchantPartyId       uuid.UUID                                  `json:"merchantPartyId"`
	BalanceType           model.MerchantBalanceAccountsBalanceType   `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode          string                                     `json:"currencyCode"`
	LinkedLedgerAccountId uuid.UUID                                  `json:"linkedLedgerAccountId"`
	AccountStatus         model.MerchantBalanceAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive"`
	Metadata              json.RawMessage                            `json:"metadata"`
}

func (d *MerchantBalanceAccountsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *MerchantBalanceAccountsCreateRequest) ToModel() model.MerchantBalanceAccounts {
	id, _ := uuid.NewV4()
	return model.MerchantBalanceAccounts{
		Id:                    id,
		MerchantPartyId:       d.MerchantPartyId,
		BalanceType:           d.BalanceType,
		CurrencyCode:          d.CurrencyCode,
		LinkedLedgerAccountId: d.LinkedLedgerAccountId,
		AccountStatus:         d.AccountStatus,
		Metadata:              d.Metadata,
	}
}

type MerchantBalanceAccountsListCreateRequest []*MerchantBalanceAccountsCreateRequest

func (d MerchantBalanceAccountsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceAccounts := range d {
		err = validator.Struct(merchantBalanceAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceAccountsListCreateRequest) ToModelList() []model.MerchantBalanceAccounts {
	out := make([]model.MerchantBalanceAccounts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type MerchantBalanceAccountsUpdateRequest struct {
	MerchantPartyId       uuid.UUID                                  `json:"merchantPartyId"`
	BalanceType           model.MerchantBalanceAccountsBalanceType   `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode          string                                     `json:"currencyCode"`
	LinkedLedgerAccountId uuid.UUID                                  `json:"linkedLedgerAccountId"`
	AccountStatus         model.MerchantBalanceAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive"`
	Metadata              json.RawMessage                            `json:"metadata"`
}

func (d *MerchantBalanceAccountsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d MerchantBalanceAccountsUpdateRequest) ToModel() model.MerchantBalanceAccounts {
	return model.MerchantBalanceAccounts{
		MerchantPartyId:       d.MerchantPartyId,
		BalanceType:           d.BalanceType,
		CurrencyCode:          d.CurrencyCode,
		LinkedLedgerAccountId: d.LinkedLedgerAccountId,
		AccountStatus:         d.AccountStatus,
		Metadata:              d.Metadata,
	}
}

type MerchantBalanceAccountsBulkUpdateRequest struct {
	Id                    uuid.UUID                                  `json:"id"`
	MerchantPartyId       uuid.UUID                                  `json:"merchantPartyId"`
	BalanceType           model.MerchantBalanceAccountsBalanceType   `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode          string                                     `json:"currencyCode"`
	LinkedLedgerAccountId uuid.UUID                                  `json:"linkedLedgerAccountId"`
	AccountStatus         model.MerchantBalanceAccountsAccountStatus `json:"accountStatus" example:"active" enums:"active,inactive"`
	Metadata              json.RawMessage                            `json:"metadata"`
}

func (d MerchantBalanceAccountsBulkUpdateRequest) PrimaryID() MerchantBalanceAccountsPrimaryID {
	return MerchantBalanceAccountsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceAccountsListBulkUpdateRequest []*MerchantBalanceAccountsBulkUpdateRequest

func (d MerchantBalanceAccountsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceAccounts := range d {
		err = validator.Struct(merchantBalanceAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceAccountsBulkUpdateRequest) ToModel() model.MerchantBalanceAccounts {
	return model.MerchantBalanceAccounts{
		Id:                    d.Id,
		MerchantPartyId:       d.MerchantPartyId,
		BalanceType:           d.BalanceType,
		CurrencyCode:          d.CurrencyCode,
		LinkedLedgerAccountId: d.LinkedLedgerAccountId,
		AccountStatus:         d.AccountStatus,
		Metadata:              d.Metadata,
	}
}

type MerchantBalanceAccountsResponse struct {
	Id                    uuid.UUID                                  `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId       uuid.UUID                                  `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BalanceType           model.MerchantBalanceAccountsBalanceType   `json:"balanceType" validate:"required,oneof=pending available reserved payable paid_out negative disputed refundable" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode          string                                     `json:"currencyCode" validate:"required"`
	LinkedLedgerAccountId uuid.UUID                                  `json:"linkedLedgerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AccountStatus         model.MerchantBalanceAccountsAccountStatus `json:"accountStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Metadata              json.RawMessage                            `json:"metadata" swaggertype:"object"`
}

func NewMerchantBalanceAccountsResponse(merchantBalanceAccounts model.MerchantBalanceAccounts) MerchantBalanceAccountsResponse {
	return MerchantBalanceAccountsResponse{
		Id:                    merchantBalanceAccounts.Id,
		MerchantPartyId:       merchantBalanceAccounts.MerchantPartyId,
		BalanceType:           model.MerchantBalanceAccountsBalanceType(merchantBalanceAccounts.BalanceType),
		CurrencyCode:          merchantBalanceAccounts.CurrencyCode,
		LinkedLedgerAccountId: merchantBalanceAccounts.LinkedLedgerAccountId,
		AccountStatus:         model.MerchantBalanceAccountsAccountStatus(merchantBalanceAccounts.AccountStatus),
		Metadata:              merchantBalanceAccounts.Metadata,
	}
}

type MerchantBalanceAccountsListResponse []*MerchantBalanceAccountsResponse

func NewMerchantBalanceAccountsListResponse(merchantBalanceAccountsList model.MerchantBalanceAccountsList) MerchantBalanceAccountsListResponse {
	dtoMerchantBalanceAccountsListResponse := MerchantBalanceAccountsListResponse{}
	for _, merchantBalanceAccounts := range merchantBalanceAccountsList {
		dtoMerchantBalanceAccountsResponse := NewMerchantBalanceAccountsResponse(*merchantBalanceAccounts)
		dtoMerchantBalanceAccountsListResponse = append(dtoMerchantBalanceAccountsListResponse, &dtoMerchantBalanceAccountsResponse)
	}
	return dtoMerchantBalanceAccountsListResponse
}

type MerchantBalanceAccountsPrimaryIDList []MerchantBalanceAccountsPrimaryID

func (d MerchantBalanceAccountsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceAccounts := range d {
		err = validator.Struct(merchantBalanceAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

type MerchantBalanceAccountsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *MerchantBalanceAccountsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d MerchantBalanceAccountsPrimaryID) ToModel() model.MerchantBalanceAccountsPrimaryID {
	return model.MerchantBalanceAccountsPrimaryID{
		Id: d.Id,
	}
}

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

type FinanceProviderAccountsDTOFieldNameType string

type financeProviderAccountsDTOFieldName struct {
	Id                 FinanceProviderAccountsDTOFieldNameType
	ProviderCode       FinanceProviderAccountsDTOFieldNameType
	ProviderName       FinanceProviderAccountsDTOFieldNameType
	OwnerPartyId       FinanceProviderAccountsDTOFieldNameType
	Environment        FinanceProviderAccountsDTOFieldNameType
	ApiBaseUrl         FinanceProviderAccountsDTOFieldNameType
	MerchantRef        FinanceProviderAccountsDTOFieldNameType
	SettlementCurrency FinanceProviderAccountsDTOFieldNameType
	VaultSecretRef     FinanceProviderAccountsDTOFieldNameType
	WebhookSecretRef   FinanceProviderAccountsDTOFieldNameType
	ProviderStatus     FinanceProviderAccountsDTOFieldNameType
	Capabilities       FinanceProviderAccountsDTOFieldNameType
	Metadata           FinanceProviderAccountsDTOFieldNameType
	MetaCreatedAt      FinanceProviderAccountsDTOFieldNameType
	MetaCreatedBy      FinanceProviderAccountsDTOFieldNameType
	MetaUpdatedAt      FinanceProviderAccountsDTOFieldNameType
	MetaUpdatedBy      FinanceProviderAccountsDTOFieldNameType
	MetaDeletedAt      FinanceProviderAccountsDTOFieldNameType
	MetaDeletedBy      FinanceProviderAccountsDTOFieldNameType
}

var FinanceProviderAccountsDTOFieldName = financeProviderAccountsDTOFieldName{
	Id:                 "id",
	ProviderCode:       "providerCode",
	ProviderName:       "providerName",
	OwnerPartyId:       "ownerPartyId",
	Environment:        "environment",
	ApiBaseUrl:         "apiBaseUrl",
	MerchantRef:        "merchantRef",
	SettlementCurrency: "settlementCurrency",
	VaultSecretRef:     "vaultSecretRef",
	WebhookSecretRef:   "webhookSecretRef",
	ProviderStatus:     "providerStatus",
	Capabilities:       "capabilities",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformFinanceProviderAccountsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceProviderAccountsDTOFieldName.Id):
		return string(model.FinanceProviderAccountsDBFieldName.Id), true

	case string(FinanceProviderAccountsDTOFieldName.ProviderCode):
		return string(model.FinanceProviderAccountsDBFieldName.ProviderCode), true

	case string(FinanceProviderAccountsDTOFieldName.ProviderName):
		return string(model.FinanceProviderAccountsDBFieldName.ProviderName), true

	case string(FinanceProviderAccountsDTOFieldName.OwnerPartyId):
		return string(model.FinanceProviderAccountsDBFieldName.OwnerPartyId), true

	case string(FinanceProviderAccountsDTOFieldName.Environment):
		return string(model.FinanceProviderAccountsDBFieldName.Environment), true

	case string(FinanceProviderAccountsDTOFieldName.ApiBaseUrl):
		return string(model.FinanceProviderAccountsDBFieldName.ApiBaseUrl), true

	case string(FinanceProviderAccountsDTOFieldName.MerchantRef):
		return string(model.FinanceProviderAccountsDBFieldName.MerchantRef), true

	case string(FinanceProviderAccountsDTOFieldName.SettlementCurrency):
		return string(model.FinanceProviderAccountsDBFieldName.SettlementCurrency), true

	case string(FinanceProviderAccountsDTOFieldName.VaultSecretRef):
		return string(model.FinanceProviderAccountsDBFieldName.VaultSecretRef), true

	case string(FinanceProviderAccountsDTOFieldName.WebhookSecretRef):
		return string(model.FinanceProviderAccountsDBFieldName.WebhookSecretRef), true

	case string(FinanceProviderAccountsDTOFieldName.ProviderStatus):
		return string(model.FinanceProviderAccountsDBFieldName.ProviderStatus), true

	case string(FinanceProviderAccountsDTOFieldName.Capabilities):
		return string(model.FinanceProviderAccountsDBFieldName.Capabilities), true

	case string(FinanceProviderAccountsDTOFieldName.Metadata):
		return string(model.FinanceProviderAccountsDBFieldName.Metadata), true

	case string(FinanceProviderAccountsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceProviderAccountsDBFieldName.MetaCreatedAt), true

	case string(FinanceProviderAccountsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceProviderAccountsDBFieldName.MetaCreatedBy), true

	case string(FinanceProviderAccountsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedAt), true

	case string(FinanceProviderAccountsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedBy), true

	case string(FinanceProviderAccountsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceProviderAccountsDBFieldName.MetaDeletedAt), true

	case string(FinanceProviderAccountsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceProviderAccountsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceProviderAccountsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceProviderAccountsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceProviderAccountsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceProviderAccountsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceProviderAccountsProjectionOutputPath(path string) error {
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

func transformFinanceProviderAccountsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceProviderAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceProviderAccountsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceProviderAccountsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceProviderAccountsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceProviderAccountsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceProviderAccountsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceProviderAccountsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceProviderAccountsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceProviderAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceProviderAccountsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceProviderAccountsFilter(filter *model.Filter) {
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
			Field: string(FinanceProviderAccountsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceProviderAccountsSelectableResponse map[string]interface{}
type FinanceProviderAccountsSelectableListResponse []*FinanceProviderAccountsSelectableResponse

func assignFinanceProviderAccountsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceProviderAccountsSelectableValue(out FinanceProviderAccountsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceProviderAccountsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceProviderAccountsSelectableResponse(financeProviderAccounts model.FinanceProviderAccounts, filter model.Filter) FinanceProviderAccountsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceProviderAccountsDBFieldName.Id),
			string(model.FinanceProviderAccountsDBFieldName.ProviderCode),
			string(model.FinanceProviderAccountsDBFieldName.ProviderName),
			string(model.FinanceProviderAccountsDBFieldName.OwnerPartyId),
			string(model.FinanceProviderAccountsDBFieldName.Environment),
			string(model.FinanceProviderAccountsDBFieldName.ApiBaseUrl),
			string(model.FinanceProviderAccountsDBFieldName.MerchantRef),
			string(model.FinanceProviderAccountsDBFieldName.SettlementCurrency),
			string(model.FinanceProviderAccountsDBFieldName.VaultSecretRef),
			string(model.FinanceProviderAccountsDBFieldName.WebhookSecretRef),
			string(model.FinanceProviderAccountsDBFieldName.ProviderStatus),
			string(model.FinanceProviderAccountsDBFieldName.Capabilities),
			string(model.FinanceProviderAccountsDBFieldName.Metadata),
			string(model.FinanceProviderAccountsDBFieldName.MetaCreatedAt),
			string(model.FinanceProviderAccountsDBFieldName.MetaCreatedBy),
			string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedAt),
			string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedBy),
			string(model.FinanceProviderAccountsDBFieldName.MetaDeletedAt),
			string(model.FinanceProviderAccountsDBFieldName.MetaDeletedBy),
		)
	}
	financeProviderAccountsSelectableResponse := FinanceProviderAccountsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceProviderAccountsDBFieldName.Id):
			key := string(FinanceProviderAccountsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.Id, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.ProviderCode):
			key := string(FinanceProviderAccountsDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.ProviderCode, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.ProviderName):
			key := string(FinanceProviderAccountsDTOFieldName.ProviderName)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.ProviderName, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.OwnerPartyId):
			key := string(FinanceProviderAccountsDTOFieldName.OwnerPartyId)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.OwnerPartyId, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.Environment):
			key := string(FinanceProviderAccountsDTOFieldName.Environment)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, model.Environment(financeProviderAccounts.Environment), explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.ApiBaseUrl):
			key := string(FinanceProviderAccountsDTOFieldName.ApiBaseUrl)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.ApiBaseUrl.String, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MerchantRef):
			key := string(FinanceProviderAccountsDTOFieldName.MerchantRef)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MerchantRef.String, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.SettlementCurrency):
			key := string(FinanceProviderAccountsDTOFieldName.SettlementCurrency)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.SettlementCurrency, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.VaultSecretRef):
			key := string(FinanceProviderAccountsDTOFieldName.VaultSecretRef)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.VaultSecretRef, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.WebhookSecretRef):
			key := string(FinanceProviderAccountsDTOFieldName.WebhookSecretRef)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.WebhookSecretRef.String, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.ProviderStatus):
			key := string(FinanceProviderAccountsDTOFieldName.ProviderStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, model.ProviderStatus(financeProviderAccounts.ProviderStatus), explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.Capabilities):
			key := string(FinanceProviderAccountsDTOFieldName.Capabilities)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.Capabilities, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.Metadata):
			key := string(FinanceProviderAccountsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.Metadata, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaCreatedAt):
			key := string(FinanceProviderAccountsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaCreatedAt, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaCreatedBy):
			key := string(FinanceProviderAccountsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaCreatedBy, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedAt):
			key := string(FinanceProviderAccountsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaUpdatedBy):
			key := string(FinanceProviderAccountsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaDeletedAt):
			key := string(FinanceProviderAccountsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceProviderAccountsDBFieldName.MetaDeletedBy):
			key := string(FinanceProviderAccountsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceProviderAccountsSelectableValue(financeProviderAccountsSelectableResponse, key, financeProviderAccounts.MetaDeletedBy, explicitAlias)

		}
	}
	return financeProviderAccountsSelectableResponse
}

func NewFinanceProviderAccountsListResponseFromFilterResult(result []model.FinanceProviderAccountsFilterResult, filter model.Filter) FinanceProviderAccountsSelectableListResponse {
	dtoFinanceProviderAccountsListResponse := FinanceProviderAccountsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceProviderAccountsResponse := NewFinanceProviderAccountsSelectableResponse(row.FinanceProviderAccounts, filter)
		dtoFinanceProviderAccountsListResponse = append(dtoFinanceProviderAccountsListResponse, &dtoFinanceProviderAccountsResponse)
	}
	return dtoFinanceProviderAccountsListResponse
}

type FinanceProviderAccountsFilterResponse struct {
	Metadata Metadata                                      `json:"metadata"`
	Data     FinanceProviderAccountsSelectableListResponse `json:"data"`
}

func reverseFinanceProviderAccountsFilterResults(result []model.FinanceProviderAccountsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceProviderAccountsFilterResponse(result []model.FinanceProviderAccountsFilterResult, filter model.Filter) (resp FinanceProviderAccountsFilterResponse) {
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
			reverseFinanceProviderAccountsFilterResults(dataResult)
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

	resp.Data = NewFinanceProviderAccountsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceProviderAccountsCreateRequest struct {
	ProviderCode       string               `json:"providerCode"`
	ProviderName       string               `json:"providerName"`
	OwnerPartyId       uuid.UUID            `json:"ownerPartyId"`
	Environment        model.Environment    `json:"environment" example:"sandbox" enums:"sandbox,production"`
	ApiBaseUrl         string               `json:"apiBaseUrl"`
	MerchantRef        string               `json:"merchantRef"`
	SettlementCurrency string               `json:"settlementCurrency"`
	VaultSecretRef     string               `json:"vaultSecretRef"`
	WebhookSecretRef   string               `json:"webhookSecretRef"`
	ProviderStatus     model.ProviderStatus `json:"providerStatus" example:"active" enums:"active,inactive"`
	Capabilities       json.RawMessage      `json:"capabilities"`
	Metadata           json.RawMessage      `json:"metadata"`
}

func (d *FinanceProviderAccountsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceProviderAccountsCreateRequest) ToModel() model.FinanceProviderAccounts {
	id, _ := uuid.NewV4()
	return model.FinanceProviderAccounts{
		Id:                 id,
		ProviderCode:       d.ProviderCode,
		ProviderName:       d.ProviderName,
		OwnerPartyId:       d.OwnerPartyId,
		Environment:        d.Environment,
		ApiBaseUrl:         null.StringFrom(d.ApiBaseUrl),
		MerchantRef:        null.StringFrom(d.MerchantRef),
		SettlementCurrency: d.SettlementCurrency,
		VaultSecretRef:     d.VaultSecretRef,
		WebhookSecretRef:   null.StringFrom(d.WebhookSecretRef),
		ProviderStatus:     d.ProviderStatus,
		Capabilities:       d.Capabilities,
		Metadata:           d.Metadata,
	}
}

type FinanceProviderAccountsListCreateRequest []*FinanceProviderAccountsCreateRequest

func (d FinanceProviderAccountsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeProviderAccounts := range d {
		err = validator.Struct(financeProviderAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceProviderAccountsListCreateRequest) ToModelList() []model.FinanceProviderAccounts {
	out := make([]model.FinanceProviderAccounts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceProviderAccountsUpdateRequest struct {
	ProviderCode       string               `json:"providerCode"`
	ProviderName       string               `json:"providerName"`
	OwnerPartyId       uuid.UUID            `json:"ownerPartyId"`
	Environment        model.Environment    `json:"environment" example:"sandbox" enums:"sandbox,production"`
	ApiBaseUrl         string               `json:"apiBaseUrl"`
	MerchantRef        string               `json:"merchantRef"`
	SettlementCurrency string               `json:"settlementCurrency"`
	VaultSecretRef     string               `json:"vaultSecretRef"`
	WebhookSecretRef   string               `json:"webhookSecretRef"`
	ProviderStatus     model.ProviderStatus `json:"providerStatus" example:"active" enums:"active,inactive"`
	Capabilities       json.RawMessage      `json:"capabilities"`
	Metadata           json.RawMessage      `json:"metadata"`
}

func (d *FinanceProviderAccountsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceProviderAccountsUpdateRequest) ToModel() model.FinanceProviderAccounts {
	return model.FinanceProviderAccounts{
		ProviderCode:       d.ProviderCode,
		ProviderName:       d.ProviderName,
		OwnerPartyId:       d.OwnerPartyId,
		Environment:        d.Environment,
		ApiBaseUrl:         null.StringFrom(d.ApiBaseUrl),
		MerchantRef:        null.StringFrom(d.MerchantRef),
		SettlementCurrency: d.SettlementCurrency,
		VaultSecretRef:     d.VaultSecretRef,
		WebhookSecretRef:   null.StringFrom(d.WebhookSecretRef),
		ProviderStatus:     d.ProviderStatus,
		Capabilities:       d.Capabilities,
		Metadata:           d.Metadata,
	}
}

type FinanceProviderAccountsBulkUpdateRequest struct {
	Id                 uuid.UUID            `json:"id"`
	ProviderCode       string               `json:"providerCode"`
	ProviderName       string               `json:"providerName"`
	OwnerPartyId       uuid.UUID            `json:"ownerPartyId"`
	Environment        model.Environment    `json:"environment" example:"sandbox" enums:"sandbox,production"`
	ApiBaseUrl         string               `json:"apiBaseUrl"`
	MerchantRef        string               `json:"merchantRef"`
	SettlementCurrency string               `json:"settlementCurrency"`
	VaultSecretRef     string               `json:"vaultSecretRef"`
	WebhookSecretRef   string               `json:"webhookSecretRef"`
	ProviderStatus     model.ProviderStatus `json:"providerStatus" example:"active" enums:"active,inactive"`
	Capabilities       json.RawMessage      `json:"capabilities"`
	Metadata           json.RawMessage      `json:"metadata"`
}

func (d FinanceProviderAccountsBulkUpdateRequest) PrimaryID() FinanceProviderAccountsPrimaryID {
	return FinanceProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

type FinanceProviderAccountsListBulkUpdateRequest []*FinanceProviderAccountsBulkUpdateRequest

func (d FinanceProviderAccountsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeProviderAccounts := range d {
		err = validator.Struct(financeProviderAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceProviderAccountsBulkUpdateRequest) ToModel() model.FinanceProviderAccounts {
	return model.FinanceProviderAccounts{
		Id:                 d.Id,
		ProviderCode:       d.ProviderCode,
		ProviderName:       d.ProviderName,
		OwnerPartyId:       d.OwnerPartyId,
		Environment:        d.Environment,
		ApiBaseUrl:         null.StringFrom(d.ApiBaseUrl),
		MerchantRef:        null.StringFrom(d.MerchantRef),
		SettlementCurrency: d.SettlementCurrency,
		VaultSecretRef:     d.VaultSecretRef,
		WebhookSecretRef:   null.StringFrom(d.WebhookSecretRef),
		ProviderStatus:     d.ProviderStatus,
		Capabilities:       d.Capabilities,
		Metadata:           d.Metadata,
	}
}

type FinanceProviderAccountsResponse struct {
	Id                 uuid.UUID            `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode       string               `json:"providerCode" validate:"required"`
	ProviderName       string               `json:"providerName" validate:"required"`
	OwnerPartyId       uuid.UUID            `json:"ownerPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Environment        model.Environment    `json:"environment" validate:"required,oneof=sandbox production" enums:"sandbox,production"`
	ApiBaseUrl         string               `json:"apiBaseUrl" validate:"url"`
	MerchantRef        string               `json:"merchantRef"`
	SettlementCurrency string               `json:"settlementCurrency" validate:"required"`
	VaultSecretRef     string               `json:"vaultSecretRef" validate:"required"`
	WebhookSecretRef   string               `json:"webhookSecretRef"`
	ProviderStatus     model.ProviderStatus `json:"providerStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Capabilities       json.RawMessage      `json:"capabilities" swaggertype:"object"`
	Metadata           json.RawMessage      `json:"metadata" swaggertype:"object"`
}

func NewFinanceProviderAccountsResponse(financeProviderAccounts model.FinanceProviderAccounts) FinanceProviderAccountsResponse {
	return FinanceProviderAccountsResponse{
		Id:                 financeProviderAccounts.Id,
		ProviderCode:       financeProviderAccounts.ProviderCode,
		ProviderName:       financeProviderAccounts.ProviderName,
		OwnerPartyId:       financeProviderAccounts.OwnerPartyId,
		Environment:        model.Environment(financeProviderAccounts.Environment),
		ApiBaseUrl:         financeProviderAccounts.ApiBaseUrl.String,
		MerchantRef:        financeProviderAccounts.MerchantRef.String,
		SettlementCurrency: financeProviderAccounts.SettlementCurrency,
		VaultSecretRef:     financeProviderAccounts.VaultSecretRef,
		WebhookSecretRef:   financeProviderAccounts.WebhookSecretRef.String,
		ProviderStatus:     model.ProviderStatus(financeProviderAccounts.ProviderStatus),
		Capabilities:       financeProviderAccounts.Capabilities,
		Metadata:           financeProviderAccounts.Metadata,
	}
}

type FinanceProviderAccountsListResponse []*FinanceProviderAccountsResponse

func NewFinanceProviderAccountsListResponse(financeProviderAccountsList model.FinanceProviderAccountsList) FinanceProviderAccountsListResponse {
	dtoFinanceProviderAccountsListResponse := FinanceProviderAccountsListResponse{}
	for _, financeProviderAccounts := range financeProviderAccountsList {
		dtoFinanceProviderAccountsResponse := NewFinanceProviderAccountsResponse(*financeProviderAccounts)
		dtoFinanceProviderAccountsListResponse = append(dtoFinanceProviderAccountsListResponse, &dtoFinanceProviderAccountsResponse)
	}
	return dtoFinanceProviderAccountsListResponse
}

type FinanceProviderAccountsPrimaryIDList []FinanceProviderAccountsPrimaryID

func (d FinanceProviderAccountsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeProviderAccounts := range d {
		err = validator.Struct(financeProviderAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceProviderAccountsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceProviderAccountsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceProviderAccountsPrimaryID) ToModel() model.FinanceProviderAccountsPrimaryID {
	return model.FinanceProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

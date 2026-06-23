package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerAccountResolutionRulesDTOFieldNameType string

type ledgerAccountResolutionRulesDTOFieldName struct {
	Id              LedgerAccountResolutionRulesDTOFieldNameType
	BookId          LedgerAccountResolutionRulesDTOFieldNameType
	Purpose         LedgerAccountResolutionRulesDTOFieldNameType
	MerchantPartyId LedgerAccountResolutionRulesDTOFieldNameType
	ProviderCode    LedgerAccountResolutionRulesDTOFieldNameType
	CurrencyCode    LedgerAccountResolutionRulesDTOFieldNameType
	SourceType      LedgerAccountResolutionRulesDTOFieldNameType
	SourceSubtype   LedgerAccountResolutionRulesDTOFieldNameType
	AccountId       LedgerAccountResolutionRulesDTOFieldNameType
	Priority        LedgerAccountResolutionRulesDTOFieldNameType
	IsActive        LedgerAccountResolutionRulesDTOFieldNameType
	Conditions      LedgerAccountResolutionRulesDTOFieldNameType
	Metadata        LedgerAccountResolutionRulesDTOFieldNameType
	MetaCreatedAt   LedgerAccountResolutionRulesDTOFieldNameType
	MetaCreatedBy   LedgerAccountResolutionRulesDTOFieldNameType
	MetaUpdatedAt   LedgerAccountResolutionRulesDTOFieldNameType
	MetaUpdatedBy   LedgerAccountResolutionRulesDTOFieldNameType
	MetaDeletedAt   LedgerAccountResolutionRulesDTOFieldNameType
	MetaDeletedBy   LedgerAccountResolutionRulesDTOFieldNameType
}

var LedgerAccountResolutionRulesDTOFieldName = ledgerAccountResolutionRulesDTOFieldName{
	Id:              "id",
	BookId:          "bookId",
	Purpose:         "purpose",
	MerchantPartyId: "merchantPartyId",
	ProviderCode:    "providerCode",
	CurrencyCode:    "currencyCode",
	SourceType:      "sourceType",
	SourceSubtype:   "sourceSubtype",
	AccountId:       "accountId",
	Priority:        "priority",
	IsActive:        "isActive",
	Conditions:      "conditions",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformLedgerAccountResolutionRulesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerAccountResolutionRulesDTOFieldName.Id):
		return string(model.LedgerAccountResolutionRulesDBFieldName.Id), true

	case string(LedgerAccountResolutionRulesDTOFieldName.BookId):
		return string(model.LedgerAccountResolutionRulesDBFieldName.BookId), true

	case string(LedgerAccountResolutionRulesDTOFieldName.Purpose):
		return string(model.LedgerAccountResolutionRulesDBFieldName.Purpose), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MerchantPartyId):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MerchantPartyId), true

	case string(LedgerAccountResolutionRulesDTOFieldName.ProviderCode):
		return string(model.LedgerAccountResolutionRulesDBFieldName.ProviderCode), true

	case string(LedgerAccountResolutionRulesDTOFieldName.CurrencyCode):
		return string(model.LedgerAccountResolutionRulesDBFieldName.CurrencyCode), true

	case string(LedgerAccountResolutionRulesDTOFieldName.SourceType):
		return string(model.LedgerAccountResolutionRulesDBFieldName.SourceType), true

	case string(LedgerAccountResolutionRulesDTOFieldName.SourceSubtype):
		return string(model.LedgerAccountResolutionRulesDBFieldName.SourceSubtype), true

	case string(LedgerAccountResolutionRulesDTOFieldName.AccountId):
		return string(model.LedgerAccountResolutionRulesDBFieldName.AccountId), true

	case string(LedgerAccountResolutionRulesDTOFieldName.Priority):
		return string(model.LedgerAccountResolutionRulesDBFieldName.Priority), true

	case string(LedgerAccountResolutionRulesDTOFieldName.IsActive):
		return string(model.LedgerAccountResolutionRulesDBFieldName.IsActive), true

	case string(LedgerAccountResolutionRulesDTOFieldName.Conditions):
		return string(model.LedgerAccountResolutionRulesDBFieldName.Conditions), true

	case string(LedgerAccountResolutionRulesDTOFieldName.Metadata):
		return string(model.LedgerAccountResolutionRulesDBFieldName.Metadata), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedAt), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedBy), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedAt), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedBy), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedAt), true

	case string(LedgerAccountResolutionRulesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerAccountResolutionRulesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerAccountResolutionRulesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerAccountResolutionRulesProjectionOutputPath(path string) error {
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

func transformLedgerAccountResolutionRulesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerAccountResolutionRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerAccountResolutionRulesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerAccountResolutionRulesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerAccountResolutionRulesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerAccountResolutionRulesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerAccountResolutionRulesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerAccountResolutionRulesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerAccountResolutionRulesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerAccountResolutionRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerAccountResolutionRulesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerAccountResolutionRulesFilter(filter *model.Filter) {
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
			Field: string(LedgerAccountResolutionRulesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerAccountResolutionRulesSelectableResponse map[string]interface{}
type LedgerAccountResolutionRulesSelectableListResponse []*LedgerAccountResolutionRulesSelectableResponse

func assignLedgerAccountResolutionRulesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerAccountResolutionRulesSelectableValue(out LedgerAccountResolutionRulesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerAccountResolutionRulesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerAccountResolutionRulesSelectableResponse(ledgerAccountResolutionRules model.LedgerAccountResolutionRules, filter model.Filter) LedgerAccountResolutionRulesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerAccountResolutionRulesDBFieldName.Id),
			string(model.LedgerAccountResolutionRulesDBFieldName.BookId),
			string(model.LedgerAccountResolutionRulesDBFieldName.Purpose),
			string(model.LedgerAccountResolutionRulesDBFieldName.MerchantPartyId),
			string(model.LedgerAccountResolutionRulesDBFieldName.ProviderCode),
			string(model.LedgerAccountResolutionRulesDBFieldName.CurrencyCode),
			string(model.LedgerAccountResolutionRulesDBFieldName.SourceType),
			string(model.LedgerAccountResolutionRulesDBFieldName.SourceSubtype),
			string(model.LedgerAccountResolutionRulesDBFieldName.AccountId),
			string(model.LedgerAccountResolutionRulesDBFieldName.Priority),
			string(model.LedgerAccountResolutionRulesDBFieldName.IsActive),
			string(model.LedgerAccountResolutionRulesDBFieldName.Conditions),
			string(model.LedgerAccountResolutionRulesDBFieldName.Metadata),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedAt),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedBy),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedAt),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedBy),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedAt),
			string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerAccountResolutionRulesSelectableResponse := LedgerAccountResolutionRulesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerAccountResolutionRulesDBFieldName.Id):
			key := string(LedgerAccountResolutionRulesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.Id, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.BookId):
			key := string(LedgerAccountResolutionRulesDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.BookId, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.Purpose):
			key := string(LedgerAccountResolutionRulesDTOFieldName.Purpose)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, model.Purpose(ledgerAccountResolutionRules.Purpose), explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MerchantPartyId):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MerchantPartyId.UUID, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.ProviderCode):
			key := string(LedgerAccountResolutionRulesDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.ProviderCode.String, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.CurrencyCode):
			key := string(LedgerAccountResolutionRulesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.CurrencyCode, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.SourceType):
			key := string(LedgerAccountResolutionRulesDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.SourceType.String, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.SourceSubtype):
			key := string(LedgerAccountResolutionRulesDTOFieldName.SourceSubtype)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.SourceSubtype.String, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.AccountId):
			key := string(LedgerAccountResolutionRulesDTOFieldName.AccountId)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.AccountId, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.Priority):
			key := string(LedgerAccountResolutionRulesDTOFieldName.Priority)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.Priority, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.IsActive):
			key := string(LedgerAccountResolutionRulesDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.IsActive, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.Conditions):
			key := string(LedgerAccountResolutionRulesDTOFieldName.Conditions)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.Conditions, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.Metadata):
			key := string(LedgerAccountResolutionRulesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.Metadata, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedAt):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaCreatedAt, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaCreatedBy):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaCreatedBy, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedAt):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaUpdatedBy):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedAt):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerAccountResolutionRulesDBFieldName.MetaDeletedBy):
			key := string(LedgerAccountResolutionRulesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerAccountResolutionRulesSelectableValue(ledgerAccountResolutionRulesSelectableResponse, key, ledgerAccountResolutionRules.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerAccountResolutionRulesSelectableResponse
}

func NewLedgerAccountResolutionRulesListResponseFromFilterResult(result []model.LedgerAccountResolutionRulesFilterResult, filter model.Filter) LedgerAccountResolutionRulesSelectableListResponse {
	dtoLedgerAccountResolutionRulesListResponse := LedgerAccountResolutionRulesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerAccountResolutionRulesResponse := NewLedgerAccountResolutionRulesSelectableResponse(row.LedgerAccountResolutionRules, filter)
		dtoLedgerAccountResolutionRulesListResponse = append(dtoLedgerAccountResolutionRulesListResponse, &dtoLedgerAccountResolutionRulesResponse)
	}
	return dtoLedgerAccountResolutionRulesListResponse
}

type LedgerAccountResolutionRulesFilterResponse struct {
	Metadata Metadata                                           `json:"metadata"`
	Data     LedgerAccountResolutionRulesSelectableListResponse `json:"data"`
}

func reverseLedgerAccountResolutionRulesFilterResults(result []model.LedgerAccountResolutionRulesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerAccountResolutionRulesFilterResponse(result []model.LedgerAccountResolutionRulesFilterResult, filter model.Filter) (resp LedgerAccountResolutionRulesFilterResponse) {
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
			reverseLedgerAccountResolutionRulesFilterResults(dataResult)
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

	resp.Data = NewLedgerAccountResolutionRulesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerAccountResolutionRulesCreateRequest struct {
	BookId          uuid.UUID       `json:"bookId"`
	Purpose         model.Purpose   `json:"purpose" example:"provider_clearing" enums:"provider_clearing,merchant_pending,merchant_payable,merchant_available,merchant_reserved,merchant_negative,commission_revenue,tax_payable,customer_refund_payable,payout_clearing,provider_fee_expense"`
	MerchantPartyId uuid.UUID       `json:"merchantPartyId"`
	ProviderCode    string          `json:"providerCode"`
	CurrencyCode    string          `json:"currencyCode"`
	SourceType      string          `json:"sourceType"`
	SourceSubtype   string          `json:"sourceSubtype"`
	AccountId       uuid.UUID       `json:"accountId"`
	Priority        int             `json:"priority"`
	IsActive        bool            `json:"isActive"`
	Conditions      json.RawMessage `json:"conditions"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *LedgerAccountResolutionRulesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerAccountResolutionRulesCreateRequest) ToModel() model.LedgerAccountResolutionRules {
	id, _ := uuid.NewV4()
	return model.LedgerAccountResolutionRules{
		Id:              id,
		BookId:          d.BookId,
		Purpose:         d.Purpose,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		ProviderCode:    null.StringFrom(d.ProviderCode),
		CurrencyCode:    d.CurrencyCode,
		SourceType:      null.StringFrom(d.SourceType),
		SourceSubtype:   null.StringFrom(d.SourceSubtype),
		AccountId:       d.AccountId,
		Priority:        d.Priority,
		IsActive:        d.IsActive,
		Conditions:      d.Conditions,
		Metadata:        d.Metadata,
	}
}

type LedgerAccountResolutionRulesListCreateRequest []*LedgerAccountResolutionRulesCreateRequest

func (d LedgerAccountResolutionRulesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountResolutionRules := range d {
		err = validator.Struct(ledgerAccountResolutionRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountResolutionRulesListCreateRequest) ToModelList() []model.LedgerAccountResolutionRules {
	out := make([]model.LedgerAccountResolutionRules, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerAccountResolutionRulesUpdateRequest struct {
	BookId          uuid.UUID       `json:"bookId"`
	Purpose         model.Purpose   `json:"purpose" example:"provider_clearing" enums:"provider_clearing,merchant_pending,merchant_payable,merchant_available,merchant_reserved,merchant_negative,commission_revenue,tax_payable,customer_refund_payable,payout_clearing,provider_fee_expense"`
	MerchantPartyId uuid.UUID       `json:"merchantPartyId"`
	ProviderCode    string          `json:"providerCode"`
	CurrencyCode    string          `json:"currencyCode"`
	SourceType      string          `json:"sourceType"`
	SourceSubtype   string          `json:"sourceSubtype"`
	AccountId       uuid.UUID       `json:"accountId"`
	Priority        int             `json:"priority"`
	IsActive        bool            `json:"isActive"`
	Conditions      json.RawMessage `json:"conditions"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *LedgerAccountResolutionRulesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerAccountResolutionRulesUpdateRequest) ToModel() model.LedgerAccountResolutionRules {
	return model.LedgerAccountResolutionRules{
		BookId:          d.BookId,
		Purpose:         d.Purpose,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		ProviderCode:    null.StringFrom(d.ProviderCode),
		CurrencyCode:    d.CurrencyCode,
		SourceType:      null.StringFrom(d.SourceType),
		SourceSubtype:   null.StringFrom(d.SourceSubtype),
		AccountId:       d.AccountId,
		Priority:        d.Priority,
		IsActive:        d.IsActive,
		Conditions:      d.Conditions,
		Metadata:        d.Metadata,
	}
}

type LedgerAccountResolutionRulesBulkUpdateRequest struct {
	Id              uuid.UUID       `json:"id"`
	BookId          uuid.UUID       `json:"bookId"`
	Purpose         model.Purpose   `json:"purpose" example:"provider_clearing" enums:"provider_clearing,merchant_pending,merchant_payable,merchant_available,merchant_reserved,merchant_negative,commission_revenue,tax_payable,customer_refund_payable,payout_clearing,provider_fee_expense"`
	MerchantPartyId uuid.UUID       `json:"merchantPartyId"`
	ProviderCode    string          `json:"providerCode"`
	CurrencyCode    string          `json:"currencyCode"`
	SourceType      string          `json:"sourceType"`
	SourceSubtype   string          `json:"sourceSubtype"`
	AccountId       uuid.UUID       `json:"accountId"`
	Priority        int             `json:"priority"`
	IsActive        bool            `json:"isActive"`
	Conditions      json.RawMessage `json:"conditions"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d LedgerAccountResolutionRulesBulkUpdateRequest) PrimaryID() LedgerAccountResolutionRulesPrimaryID {
	return LedgerAccountResolutionRulesPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountResolutionRulesListBulkUpdateRequest []*LedgerAccountResolutionRulesBulkUpdateRequest

func (d LedgerAccountResolutionRulesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountResolutionRules := range d {
		err = validator.Struct(ledgerAccountResolutionRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerAccountResolutionRulesBulkUpdateRequest) ToModel() model.LedgerAccountResolutionRules {
	return model.LedgerAccountResolutionRules{
		Id:              d.Id,
		BookId:          d.BookId,
		Purpose:         d.Purpose,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		ProviderCode:    null.StringFrom(d.ProviderCode),
		CurrencyCode:    d.CurrencyCode,
		SourceType:      null.StringFrom(d.SourceType),
		SourceSubtype:   null.StringFrom(d.SourceSubtype),
		AccountId:       d.AccountId,
		Priority:        d.Priority,
		IsActive:        d.IsActive,
		Conditions:      d.Conditions,
		Metadata:        d.Metadata,
	}
}

type LedgerAccountResolutionRulesResponse struct {
	Id              uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BookId          uuid.UUID       `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Purpose         model.Purpose   `json:"purpose" validate:"required,oneof=provider_clearing merchant_pending merchant_payable merchant_available merchant_reserved merchant_negative commission_revenue tax_payable customer_refund_payable payout_clearing provider_fee_expense" enums:"provider_clearing,merchant_pending,merchant_payable,merchant_available,merchant_reserved,merchant_negative,commission_revenue,tax_payable,customer_refund_payable,payout_clearing,provider_fee_expense"`
	MerchantPartyId uuid.UUID       `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode    string          `json:"providerCode"`
	CurrencyCode    string          `json:"currencyCode" validate:"required"`
	SourceType      string          `json:"sourceType"`
	SourceSubtype   string          `json:"sourceSubtype"`
	AccountId       uuid.UUID       `json:"accountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Priority        int             `json:"priority" example:"1"`
	IsActive        bool            `json:"isActive" example:"true"`
	Conditions      json.RawMessage `json:"conditions" swaggertype:"object"`
	Metadata        json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewLedgerAccountResolutionRulesResponse(ledgerAccountResolutionRules model.LedgerAccountResolutionRules) LedgerAccountResolutionRulesResponse {
	return LedgerAccountResolutionRulesResponse{
		Id:              ledgerAccountResolutionRules.Id,
		BookId:          ledgerAccountResolutionRules.BookId,
		Purpose:         model.Purpose(ledgerAccountResolutionRules.Purpose),
		MerchantPartyId: ledgerAccountResolutionRules.MerchantPartyId.UUID,
		ProviderCode:    ledgerAccountResolutionRules.ProviderCode.String,
		CurrencyCode:    ledgerAccountResolutionRules.CurrencyCode,
		SourceType:      ledgerAccountResolutionRules.SourceType.String,
		SourceSubtype:   ledgerAccountResolutionRules.SourceSubtype.String,
		AccountId:       ledgerAccountResolutionRules.AccountId,
		Priority:        ledgerAccountResolutionRules.Priority,
		IsActive:        ledgerAccountResolutionRules.IsActive,
		Conditions:      ledgerAccountResolutionRules.Conditions,
		Metadata:        ledgerAccountResolutionRules.Metadata,
	}
}

type LedgerAccountResolutionRulesListResponse []*LedgerAccountResolutionRulesResponse

func NewLedgerAccountResolutionRulesListResponse(ledgerAccountResolutionRulesList model.LedgerAccountResolutionRulesList) LedgerAccountResolutionRulesListResponse {
	dtoLedgerAccountResolutionRulesListResponse := LedgerAccountResolutionRulesListResponse{}
	for _, ledgerAccountResolutionRules := range ledgerAccountResolutionRulesList {
		dtoLedgerAccountResolutionRulesResponse := NewLedgerAccountResolutionRulesResponse(*ledgerAccountResolutionRules)
		dtoLedgerAccountResolutionRulesListResponse = append(dtoLedgerAccountResolutionRulesListResponse, &dtoLedgerAccountResolutionRulesResponse)
	}
	return dtoLedgerAccountResolutionRulesListResponse
}

type LedgerAccountResolutionRulesPrimaryIDList []LedgerAccountResolutionRulesPrimaryID

func (d LedgerAccountResolutionRulesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerAccountResolutionRules := range d {
		err = validator.Struct(ledgerAccountResolutionRules)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerAccountResolutionRulesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerAccountResolutionRulesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerAccountResolutionRulesPrimaryID) ToModel() model.LedgerAccountResolutionRulesPrimaryID {
	return model.LedgerAccountResolutionRulesPrimaryID{
		Id: d.Id,
	}
}

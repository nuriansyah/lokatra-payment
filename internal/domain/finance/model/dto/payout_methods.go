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

type PayoutMethodsDTOFieldNameType string

type payoutMethodsDTOFieldName struct {
	Id                   PayoutMethodsDTOFieldNameType
	MerchantPartyId      PayoutMethodsDTOFieldNameType
	MethodType           PayoutMethodsDTOFieldNameType
	ProviderCode         PayoutMethodsDTOFieldNameType
	CountryCode          PayoutMethodsDTOFieldNameType
	CurrencyCode         PayoutMethodsDTOFieldNameType
	AccountHolderName    PayoutMethodsDTOFieldNameType
	AccountNoLast4       PayoutMethodsDTOFieldNameType
	BankCode             PayoutMethodsDTOFieldNameType
	BankName             PayoutMethodsDTOFieldNameType
	AccountTokenRef      PayoutMethodsDTOFieldNameType
	AccountEncryptedBlob PayoutMethodsDTOFieldNameType
	VerificationStatus   PayoutMethodsDTOFieldNameType
	IsDefault            PayoutMethodsDTOFieldNameType
	MethodStatus         PayoutMethodsDTOFieldNameType
	Metadata             PayoutMethodsDTOFieldNameType
	MetaCreatedAt        PayoutMethodsDTOFieldNameType
	MetaCreatedBy        PayoutMethodsDTOFieldNameType
	MetaUpdatedAt        PayoutMethodsDTOFieldNameType
	MetaUpdatedBy        PayoutMethodsDTOFieldNameType
	MetaDeletedAt        PayoutMethodsDTOFieldNameType
	MetaDeletedBy        PayoutMethodsDTOFieldNameType
}

var PayoutMethodsDTOFieldName = payoutMethodsDTOFieldName{
	Id:                   "id",
	MerchantPartyId:      "merchantPartyId",
	MethodType:           "methodType",
	ProviderCode:         "providerCode",
	CountryCode:          "countryCode",
	CurrencyCode:         "currencyCode",
	AccountHolderName:    "accountHolderName",
	AccountNoLast4:       "accountNoLast4",
	BankCode:             "bankCode",
	BankName:             "bankName",
	AccountTokenRef:      "accountTokenRef",
	AccountEncryptedBlob: "accountEncryptedBlob",
	VerificationStatus:   "verificationStatus",
	IsDefault:            "isDefault",
	MethodStatus:         "methodStatus",
	Metadata:             "metadata",
	MetaCreatedAt:        "metaCreatedAt",
	MetaCreatedBy:        "metaCreatedBy",
	MetaUpdatedAt:        "metaUpdatedAt",
	MetaUpdatedBy:        "metaUpdatedBy",
	MetaDeletedAt:        "metaDeletedAt",
	MetaDeletedBy:        "metaDeletedBy",
}

func transformPayoutMethodsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutMethodsDTOFieldName.Id):
		return string(model.PayoutMethodsDBFieldName.Id), true

	case string(PayoutMethodsDTOFieldName.MerchantPartyId):
		return string(model.PayoutMethodsDBFieldName.MerchantPartyId), true

	case string(PayoutMethodsDTOFieldName.MethodType):
		return string(model.PayoutMethodsDBFieldName.MethodType), true

	case string(PayoutMethodsDTOFieldName.ProviderCode):
		return string(model.PayoutMethodsDBFieldName.ProviderCode), true

	case string(PayoutMethodsDTOFieldName.CountryCode):
		return string(model.PayoutMethodsDBFieldName.CountryCode), true

	case string(PayoutMethodsDTOFieldName.CurrencyCode):
		return string(model.PayoutMethodsDBFieldName.CurrencyCode), true

	case string(PayoutMethodsDTOFieldName.AccountHolderName):
		return string(model.PayoutMethodsDBFieldName.AccountHolderName), true

	case string(PayoutMethodsDTOFieldName.AccountNoLast4):
		return string(model.PayoutMethodsDBFieldName.AccountNoLast4), true

	case string(PayoutMethodsDTOFieldName.BankCode):
		return string(model.PayoutMethodsDBFieldName.BankCode), true

	case string(PayoutMethodsDTOFieldName.BankName):
		return string(model.PayoutMethodsDBFieldName.BankName), true

	case string(PayoutMethodsDTOFieldName.AccountTokenRef):
		return string(model.PayoutMethodsDBFieldName.AccountTokenRef), true

	case string(PayoutMethodsDTOFieldName.AccountEncryptedBlob):
		return string(model.PayoutMethodsDBFieldName.AccountEncryptedBlob), true

	case string(PayoutMethodsDTOFieldName.VerificationStatus):
		return string(model.PayoutMethodsDBFieldName.VerificationStatus), true

	case string(PayoutMethodsDTOFieldName.IsDefault):
		return string(model.PayoutMethodsDBFieldName.IsDefault), true

	case string(PayoutMethodsDTOFieldName.MethodStatus):
		return string(model.PayoutMethodsDBFieldName.MethodStatus), true

	case string(PayoutMethodsDTOFieldName.Metadata):
		return string(model.PayoutMethodsDBFieldName.Metadata), true

	case string(PayoutMethodsDTOFieldName.MetaCreatedAt):
		return string(model.PayoutMethodsDBFieldName.MetaCreatedAt), true

	case string(PayoutMethodsDTOFieldName.MetaCreatedBy):
		return string(model.PayoutMethodsDBFieldName.MetaCreatedBy), true

	case string(PayoutMethodsDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutMethodsDBFieldName.MetaUpdatedAt), true

	case string(PayoutMethodsDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutMethodsDBFieldName.MetaUpdatedBy), true

	case string(PayoutMethodsDTOFieldName.MetaDeletedAt):
		return string(model.PayoutMethodsDBFieldName.MetaDeletedAt), true

	case string(PayoutMethodsDTOFieldName.MetaDeletedBy):
		return string(model.PayoutMethodsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutMethodsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutMethodsBaseFilterField(field string) bool {
	spec, found := model.NewPayoutMethodsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutMethodsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutMethodsProjectionOutputPath(path string) error {
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

func transformPayoutMethodsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutMethodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutMethodsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutMethodsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutMethodsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutMethodsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutMethodsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutMethodsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutMethodsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutMethodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutMethodsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutMethodsFilter(filter *model.Filter) {
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
			Field: string(PayoutMethodsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutMethodsSelectableResponse map[string]interface{}
type PayoutMethodsSelectableListResponse []*PayoutMethodsSelectableResponse

func assignPayoutMethodsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutMethodsSelectableValue(out PayoutMethodsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutMethodsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutMethodsSelectableResponse(payoutMethods model.PayoutMethods, filter model.Filter) PayoutMethodsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutMethodsDBFieldName.Id),
			string(model.PayoutMethodsDBFieldName.MerchantPartyId),
			string(model.PayoutMethodsDBFieldName.MethodType),
			string(model.PayoutMethodsDBFieldName.ProviderCode),
			string(model.PayoutMethodsDBFieldName.CountryCode),
			string(model.PayoutMethodsDBFieldName.CurrencyCode),
			string(model.PayoutMethodsDBFieldName.AccountHolderName),
			string(model.PayoutMethodsDBFieldName.AccountNoLast4),
			string(model.PayoutMethodsDBFieldName.BankCode),
			string(model.PayoutMethodsDBFieldName.BankName),
			string(model.PayoutMethodsDBFieldName.AccountTokenRef),
			string(model.PayoutMethodsDBFieldName.AccountEncryptedBlob),
			string(model.PayoutMethodsDBFieldName.VerificationStatus),
			string(model.PayoutMethodsDBFieldName.IsDefault),
			string(model.PayoutMethodsDBFieldName.MethodStatus),
			string(model.PayoutMethodsDBFieldName.Metadata),
			string(model.PayoutMethodsDBFieldName.MetaCreatedAt),
			string(model.PayoutMethodsDBFieldName.MetaCreatedBy),
			string(model.PayoutMethodsDBFieldName.MetaUpdatedAt),
			string(model.PayoutMethodsDBFieldName.MetaUpdatedBy),
			string(model.PayoutMethodsDBFieldName.MetaDeletedAt),
			string(model.PayoutMethodsDBFieldName.MetaDeletedBy),
		)
	}
	payoutMethodsSelectableResponse := PayoutMethodsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutMethodsDBFieldName.Id):
			key := string(PayoutMethodsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.Id, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MerchantPartyId):
			key := string(PayoutMethodsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MerchantPartyId, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MethodType):
			key := string(PayoutMethodsDTOFieldName.MethodType)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, model.MethodType(payoutMethods.MethodType), explicitAlias)

		case string(model.PayoutMethodsDBFieldName.ProviderCode):
			key := string(PayoutMethodsDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.ProviderCode.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.CountryCode):
			key := string(PayoutMethodsDTOFieldName.CountryCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.CountryCode, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.CurrencyCode):
			key := string(PayoutMethodsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.CurrencyCode, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.AccountHolderName):
			key := string(PayoutMethodsDTOFieldName.AccountHolderName)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.AccountHolderName.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.AccountNoLast4):
			key := string(PayoutMethodsDTOFieldName.AccountNoLast4)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.AccountNoLast4.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.BankCode):
			key := string(PayoutMethodsDTOFieldName.BankCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.BankCode.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.BankName):
			key := string(PayoutMethodsDTOFieldName.BankName)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.BankName.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.AccountTokenRef):
			key := string(PayoutMethodsDTOFieldName.AccountTokenRef)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.AccountTokenRef.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.AccountEncryptedBlob):
			key := string(PayoutMethodsDTOFieldName.AccountEncryptedBlob)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.AccountEncryptedBlob.String, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.VerificationStatus):
			key := string(PayoutMethodsDTOFieldName.VerificationStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, model.VerificationStatus(payoutMethods.VerificationStatus), explicitAlias)

		case string(model.PayoutMethodsDBFieldName.IsDefault):
			key := string(PayoutMethodsDTOFieldName.IsDefault)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.IsDefault, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MethodStatus):
			key := string(PayoutMethodsDTOFieldName.MethodStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, model.MethodStatus(payoutMethods.MethodStatus), explicitAlias)

		case string(model.PayoutMethodsDBFieldName.Metadata):
			key := string(PayoutMethodsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.Metadata, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaCreatedAt):
			key := string(PayoutMethodsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaCreatedAt, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaCreatedBy):
			key := string(PayoutMethodsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaCreatedBy, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaUpdatedAt):
			key := string(PayoutMethodsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaUpdatedBy):
			key := string(PayoutMethodsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaDeletedAt):
			key := string(PayoutMethodsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutMethodsDBFieldName.MetaDeletedBy):
			key := string(PayoutMethodsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutMethodsSelectableValue(payoutMethodsSelectableResponse, key, payoutMethods.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutMethodsSelectableResponse
}

func NewPayoutMethodsListResponseFromFilterResult(result []model.PayoutMethodsFilterResult, filter model.Filter) PayoutMethodsSelectableListResponse {
	dtoPayoutMethodsListResponse := PayoutMethodsSelectableListResponse{}
	for _, row := range result {
		dtoPayoutMethodsResponse := NewPayoutMethodsSelectableResponse(row.PayoutMethods, filter)
		dtoPayoutMethodsListResponse = append(dtoPayoutMethodsListResponse, &dtoPayoutMethodsResponse)
	}
	return dtoPayoutMethodsListResponse
}

type PayoutMethodsFilterResponse struct {
	Metadata Metadata                            `json:"metadata"`
	Data     PayoutMethodsSelectableListResponse `json:"data"`
}

func reversePayoutMethodsFilterResults(result []model.PayoutMethodsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutMethodsFilterResponse(result []model.PayoutMethodsFilterResult, filter model.Filter) (resp PayoutMethodsFilterResponse) {
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
			reversePayoutMethodsFilterResults(dataResult)
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

	resp.Data = NewPayoutMethodsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutMethodsCreateRequest struct {
	MerchantPartyId      uuid.UUID                `json:"merchantPartyId"`
	MethodType           model.MethodType         `json:"methodType" example:"bank_account" enums:"bank_account,ewallet,manual,virtual_card"`
	ProviderCode         string                   `json:"providerCode"`
	CountryCode          string                   `json:"countryCode"`
	CurrencyCode         string                   `json:"currencyCode"`
	AccountHolderName    string                   `json:"accountHolderName"`
	AccountNoLast4       string                   `json:"accountNoLast4"`
	BankCode             string                   `json:"bankCode"`
	BankName             string                   `json:"bankName"`
	AccountTokenRef      string                   `json:"accountTokenRef"`
	AccountEncryptedBlob string                   `json:"accountEncryptedBlob"`
	VerificationStatus   model.VerificationStatus `json:"verificationStatus" example:"pending" enums:"pending,verified,failed,expired"`
	IsDefault            bool                     `json:"isDefault"`
	MethodStatus         model.MethodStatus       `json:"methodStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata             json.RawMessage          `json:"metadata"`
}

func (d *PayoutMethodsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutMethodsCreateRequest) ToModel() model.PayoutMethods {
	id, _ := uuid.NewV4()
	return model.PayoutMethods{
		Id:                   id,
		MerchantPartyId:      d.MerchantPartyId,
		MethodType:           d.MethodType,
		ProviderCode:         null.StringFrom(d.ProviderCode),
		CountryCode:          d.CountryCode,
		CurrencyCode:         d.CurrencyCode,
		AccountHolderName:    null.StringFrom(d.AccountHolderName),
		AccountNoLast4:       null.StringFrom(d.AccountNoLast4),
		BankCode:             null.StringFrom(d.BankCode),
		BankName:             null.StringFrom(d.BankName),
		AccountTokenRef:      null.StringFrom(d.AccountTokenRef),
		AccountEncryptedBlob: null.StringFrom(d.AccountEncryptedBlob),
		VerificationStatus:   d.VerificationStatus,
		IsDefault:            d.IsDefault,
		MethodStatus:         d.MethodStatus,
		Metadata:             d.Metadata,
	}
}

type PayoutMethodsListCreateRequest []*PayoutMethodsCreateRequest

func (d PayoutMethodsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutMethods := range d {
		err = validator.Struct(payoutMethods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutMethodsListCreateRequest) ToModelList() []model.PayoutMethods {
	out := make([]model.PayoutMethods, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutMethodsUpdateRequest struct {
	MerchantPartyId      uuid.UUID                `json:"merchantPartyId"`
	MethodType           model.MethodType         `json:"methodType" example:"bank_account" enums:"bank_account,ewallet,manual,virtual_card"`
	ProviderCode         string                   `json:"providerCode"`
	CountryCode          string                   `json:"countryCode"`
	CurrencyCode         string                   `json:"currencyCode"`
	AccountHolderName    string                   `json:"accountHolderName"`
	AccountNoLast4       string                   `json:"accountNoLast4"`
	BankCode             string                   `json:"bankCode"`
	BankName             string                   `json:"bankName"`
	AccountTokenRef      string                   `json:"accountTokenRef"`
	AccountEncryptedBlob string                   `json:"accountEncryptedBlob"`
	VerificationStatus   model.VerificationStatus `json:"verificationStatus" example:"pending" enums:"pending,verified,failed,expired"`
	IsDefault            bool                     `json:"isDefault"`
	MethodStatus         model.MethodStatus       `json:"methodStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata             json.RawMessage          `json:"metadata"`
}

func (d *PayoutMethodsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutMethodsUpdateRequest) ToModel() model.PayoutMethods {
	return model.PayoutMethods{
		MerchantPartyId:      d.MerchantPartyId,
		MethodType:           d.MethodType,
		ProviderCode:         null.StringFrom(d.ProviderCode),
		CountryCode:          d.CountryCode,
		CurrencyCode:         d.CurrencyCode,
		AccountHolderName:    null.StringFrom(d.AccountHolderName),
		AccountNoLast4:       null.StringFrom(d.AccountNoLast4),
		BankCode:             null.StringFrom(d.BankCode),
		BankName:             null.StringFrom(d.BankName),
		AccountTokenRef:      null.StringFrom(d.AccountTokenRef),
		AccountEncryptedBlob: null.StringFrom(d.AccountEncryptedBlob),
		VerificationStatus:   d.VerificationStatus,
		IsDefault:            d.IsDefault,
		MethodStatus:         d.MethodStatus,
		Metadata:             d.Metadata,
	}
}

type PayoutMethodsBulkUpdateRequest struct {
	Id                   uuid.UUID                `json:"id"`
	MerchantPartyId      uuid.UUID                `json:"merchantPartyId"`
	MethodType           model.MethodType         `json:"methodType" example:"bank_account" enums:"bank_account,ewallet,manual,virtual_card"`
	ProviderCode         string                   `json:"providerCode"`
	CountryCode          string                   `json:"countryCode"`
	CurrencyCode         string                   `json:"currencyCode"`
	AccountHolderName    string                   `json:"accountHolderName"`
	AccountNoLast4       string                   `json:"accountNoLast4"`
	BankCode             string                   `json:"bankCode"`
	BankName             string                   `json:"bankName"`
	AccountTokenRef      string                   `json:"accountTokenRef"`
	AccountEncryptedBlob string                   `json:"accountEncryptedBlob"`
	VerificationStatus   model.VerificationStatus `json:"verificationStatus" example:"pending" enums:"pending,verified,failed,expired"`
	IsDefault            bool                     `json:"isDefault"`
	MethodStatus         model.MethodStatus       `json:"methodStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata             json.RawMessage          `json:"metadata"`
}

func (d PayoutMethodsBulkUpdateRequest) PrimaryID() PayoutMethodsPrimaryID {
	return PayoutMethodsPrimaryID{
		Id: d.Id,
	}
}

type PayoutMethodsListBulkUpdateRequest []*PayoutMethodsBulkUpdateRequest

func (d PayoutMethodsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutMethods := range d {
		err = validator.Struct(payoutMethods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutMethodsBulkUpdateRequest) ToModel() model.PayoutMethods {
	return model.PayoutMethods{
		Id:                   d.Id,
		MerchantPartyId:      d.MerchantPartyId,
		MethodType:           d.MethodType,
		ProviderCode:         null.StringFrom(d.ProviderCode),
		CountryCode:          d.CountryCode,
		CurrencyCode:         d.CurrencyCode,
		AccountHolderName:    null.StringFrom(d.AccountHolderName),
		AccountNoLast4:       null.StringFrom(d.AccountNoLast4),
		BankCode:             null.StringFrom(d.BankCode),
		BankName:             null.StringFrom(d.BankName),
		AccountTokenRef:      null.StringFrom(d.AccountTokenRef),
		AccountEncryptedBlob: null.StringFrom(d.AccountEncryptedBlob),
		VerificationStatus:   d.VerificationStatus,
		IsDefault:            d.IsDefault,
		MethodStatus:         d.MethodStatus,
		Metadata:             d.Metadata,
	}
}

type PayoutMethodsResponse struct {
	Id                   uuid.UUID                `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId      uuid.UUID                `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodType           model.MethodType         `json:"methodType" validate:"required,oneof=bank_account ewallet manual virtual_card" enums:"bank_account,ewallet,manual,virtual_card"`
	ProviderCode         string                   `json:"providerCode"`
	CountryCode          string                   `json:"countryCode"`
	CurrencyCode         string                   `json:"currencyCode" validate:"required"`
	AccountHolderName    string                   `json:"accountHolderName"`
	AccountNoLast4       string                   `json:"accountNoLast4"`
	BankCode             string                   `json:"bankCode"`
	BankName             string                   `json:"bankName"`
	AccountTokenRef      string                   `json:"accountTokenRef"`
	AccountEncryptedBlob string                   `json:"accountEncryptedBlob"`
	VerificationStatus   model.VerificationStatus `json:"verificationStatus" validate:"oneof=pending verified failed expired" enums:"pending,verified,failed,expired"`
	IsDefault            bool                     `json:"isDefault" example:"true"`
	MethodStatus         model.MethodStatus       `json:"methodStatus" validate:"oneof=active inactive blocked" enums:"active,inactive,blocked"`
	Metadata             json.RawMessage          `json:"metadata" swaggertype:"object"`
}

func NewPayoutMethodsResponse(payoutMethods model.PayoutMethods) PayoutMethodsResponse {
	return PayoutMethodsResponse{
		Id:                   payoutMethods.Id,
		MerchantPartyId:      payoutMethods.MerchantPartyId,
		MethodType:           model.MethodType(payoutMethods.MethodType),
		ProviderCode:         payoutMethods.ProviderCode.String,
		CountryCode:          payoutMethods.CountryCode,
		CurrencyCode:         payoutMethods.CurrencyCode,
		AccountHolderName:    payoutMethods.AccountHolderName.String,
		AccountNoLast4:       payoutMethods.AccountNoLast4.String,
		BankCode:             payoutMethods.BankCode.String,
		BankName:             payoutMethods.BankName.String,
		AccountTokenRef:      payoutMethods.AccountTokenRef.String,
		AccountEncryptedBlob: payoutMethods.AccountEncryptedBlob.String,
		VerificationStatus:   model.VerificationStatus(payoutMethods.VerificationStatus),
		IsDefault:            payoutMethods.IsDefault,
		MethodStatus:         model.MethodStatus(payoutMethods.MethodStatus),
		Metadata:             payoutMethods.Metadata,
	}
}

type PayoutMethodsListResponse []*PayoutMethodsResponse

func NewPayoutMethodsListResponse(payoutMethodsList model.PayoutMethodsList) PayoutMethodsListResponse {
	dtoPayoutMethodsListResponse := PayoutMethodsListResponse{}
	for _, payoutMethods := range payoutMethodsList {
		dtoPayoutMethodsResponse := NewPayoutMethodsResponse(*payoutMethods)
		dtoPayoutMethodsListResponse = append(dtoPayoutMethodsListResponse, &dtoPayoutMethodsResponse)
	}
	return dtoPayoutMethodsListResponse
}

type PayoutMethodsPrimaryIDList []PayoutMethodsPrimaryID

func (d PayoutMethodsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutMethods := range d {
		err = validator.Struct(payoutMethods)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutMethodsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutMethodsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutMethodsPrimaryID) ToModel() model.PayoutMethodsPrimaryID {
	return model.PayoutMethodsPrimaryID{
		Id: d.Id,
	}
}

package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentInstructionsDTOFieldNameType string

type paymentInstructionsDTOFieldName struct {
	Id                  PaymentInstructionsDTOFieldNameType
	PaymentAttemptId    PaymentInstructionsDTOFieldNameType
	InstructionType     PaymentInstructionsDTOFieldNameType
	IsActive            PaymentInstructionsDTOFieldNameType
	DisplayName         PaymentInstructionsDTOFieldNameType
	AccountNumber       PaymentInstructionsDTOFieldNameType
	AccountNumberMasked PaymentInstructionsDTOFieldNameType
	AccountHolderName   PaymentInstructionsDTOFieldNameType
	BankCode            PaymentInstructionsDTOFieldNameType
	BillerCode          PaymentInstructionsDTOFieldNameType
	PaymentCode         PaymentInstructionsDTOFieldNameType
	QrString            PaymentInstructionsDTOFieldNameType
	QrImageUrl          PaymentInstructionsDTOFieldNameType
	CheckoutUrl         PaymentInstructionsDTOFieldNameType
	DeeplinkUrl         PaymentInstructionsDTOFieldNameType
	RetailOutletCode    PaymentInstructionsDTOFieldNameType
	ExpiresAt           PaymentInstructionsDTOFieldNameType
	Metadata            PaymentInstructionsDTOFieldNameType
	MetaCreatedAt       PaymentInstructionsDTOFieldNameType
	MetaCreatedBy       PaymentInstructionsDTOFieldNameType
	MetaUpdatedAt       PaymentInstructionsDTOFieldNameType
	MetaUpdatedBy       PaymentInstructionsDTOFieldNameType
	MetaDeletedAt       PaymentInstructionsDTOFieldNameType
	MetaDeletedBy       PaymentInstructionsDTOFieldNameType
}

var PaymentInstructionsDTOFieldName = paymentInstructionsDTOFieldName{
	Id:                  "id",
	PaymentAttemptId:    "paymentAttemptId",
	InstructionType:     "instructionType",
	IsActive:            "isActive",
	DisplayName:         "displayName",
	AccountNumber:       "accountNumber",
	AccountNumberMasked: "accountNumberMasked",
	AccountHolderName:   "accountHolderName",
	BankCode:            "bankCode",
	BillerCode:          "billerCode",
	PaymentCode:         "paymentCode",
	QrString:            "qrString",
	QrImageUrl:          "qrImageUrl",
	CheckoutUrl:         "checkoutUrl",
	DeeplinkUrl:         "deeplinkUrl",
	RetailOutletCode:    "retailOutletCode",
	ExpiresAt:           "expiresAt",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformPaymentInstructionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentInstructionsDTOFieldName.Id):
		return string(model.PaymentInstructionsDBFieldName.Id), true

	case string(PaymentInstructionsDTOFieldName.PaymentAttemptId):
		return string(model.PaymentInstructionsDBFieldName.PaymentAttemptId), true

	case string(PaymentInstructionsDTOFieldName.InstructionType):
		return string(model.PaymentInstructionsDBFieldName.InstructionType), true

	case string(PaymentInstructionsDTOFieldName.IsActive):
		return string(model.PaymentInstructionsDBFieldName.IsActive), true

	case string(PaymentInstructionsDTOFieldName.DisplayName):
		return string(model.PaymentInstructionsDBFieldName.DisplayName), true

	case string(PaymentInstructionsDTOFieldName.AccountNumber):
		return string(model.PaymentInstructionsDBFieldName.AccountNumber), true

	case string(PaymentInstructionsDTOFieldName.AccountNumberMasked):
		return string(model.PaymentInstructionsDBFieldName.AccountNumberMasked), true

	case string(PaymentInstructionsDTOFieldName.AccountHolderName):
		return string(model.PaymentInstructionsDBFieldName.AccountHolderName), true

	case string(PaymentInstructionsDTOFieldName.BankCode):
		return string(model.PaymentInstructionsDBFieldName.BankCode), true

	case string(PaymentInstructionsDTOFieldName.BillerCode):
		return string(model.PaymentInstructionsDBFieldName.BillerCode), true

	case string(PaymentInstructionsDTOFieldName.PaymentCode):
		return string(model.PaymentInstructionsDBFieldName.PaymentCode), true

	case string(PaymentInstructionsDTOFieldName.QrString):
		return string(model.PaymentInstructionsDBFieldName.QrString), true

	case string(PaymentInstructionsDTOFieldName.QrImageUrl):
		return string(model.PaymentInstructionsDBFieldName.QrImageUrl), true

	case string(PaymentInstructionsDTOFieldName.CheckoutUrl):
		return string(model.PaymentInstructionsDBFieldName.CheckoutUrl), true

	case string(PaymentInstructionsDTOFieldName.DeeplinkUrl):
		return string(model.PaymentInstructionsDBFieldName.DeeplinkUrl), true

	case string(PaymentInstructionsDTOFieldName.RetailOutletCode):
		return string(model.PaymentInstructionsDBFieldName.RetailOutletCode), true

	case string(PaymentInstructionsDTOFieldName.ExpiresAt):
		return string(model.PaymentInstructionsDBFieldName.ExpiresAt), true

	case string(PaymentInstructionsDTOFieldName.Metadata):
		return string(model.PaymentInstructionsDBFieldName.Metadata), true

	case string(PaymentInstructionsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentInstructionsDBFieldName.MetaCreatedAt), true

	case string(PaymentInstructionsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentInstructionsDBFieldName.MetaCreatedBy), true

	case string(PaymentInstructionsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentInstructionsDBFieldName.MetaUpdatedAt), true

	case string(PaymentInstructionsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentInstructionsDBFieldName.MetaUpdatedBy), true

	case string(PaymentInstructionsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentInstructionsDBFieldName.MetaDeletedAt), true

	case string(PaymentInstructionsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentInstructionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentInstructionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentInstructionsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentInstructionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentInstructionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentInstructionsProjectionOutputPath(path string) error {
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

func transformPaymentInstructionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentInstructionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentInstructionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentInstructionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentInstructionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentInstructionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentInstructionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentInstructionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentInstructionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentInstructionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentInstructionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentInstructionsFilter(filter *model.Filter) {
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
			Field: string(PaymentInstructionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentInstructionsSelectableResponse map[string]interface{}
type PaymentInstructionsSelectableListResponse []*PaymentInstructionsSelectableResponse

func assignPaymentInstructionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentInstructionsSelectableValue(out PaymentInstructionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentInstructionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentInstructionsSelectableResponse(paymentInstructions model.PaymentInstructions, filter model.Filter) PaymentInstructionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentInstructionsDBFieldName.Id),
			string(model.PaymentInstructionsDBFieldName.PaymentAttemptId),
			string(model.PaymentInstructionsDBFieldName.InstructionType),
			string(model.PaymentInstructionsDBFieldName.IsActive),
			string(model.PaymentInstructionsDBFieldName.DisplayName),
			string(model.PaymentInstructionsDBFieldName.AccountNumber),
			string(model.PaymentInstructionsDBFieldName.AccountNumberMasked),
			string(model.PaymentInstructionsDBFieldName.AccountHolderName),
			string(model.PaymentInstructionsDBFieldName.BankCode),
			string(model.PaymentInstructionsDBFieldName.BillerCode),
			string(model.PaymentInstructionsDBFieldName.PaymentCode),
			string(model.PaymentInstructionsDBFieldName.QrString),
			string(model.PaymentInstructionsDBFieldName.QrImageUrl),
			string(model.PaymentInstructionsDBFieldName.CheckoutUrl),
			string(model.PaymentInstructionsDBFieldName.DeeplinkUrl),
			string(model.PaymentInstructionsDBFieldName.RetailOutletCode),
			string(model.PaymentInstructionsDBFieldName.ExpiresAt),
			string(model.PaymentInstructionsDBFieldName.Metadata),
			string(model.PaymentInstructionsDBFieldName.MetaCreatedAt),
			string(model.PaymentInstructionsDBFieldName.MetaCreatedBy),
			string(model.PaymentInstructionsDBFieldName.MetaUpdatedAt),
			string(model.PaymentInstructionsDBFieldName.MetaUpdatedBy),
			string(model.PaymentInstructionsDBFieldName.MetaDeletedAt),
			string(model.PaymentInstructionsDBFieldName.MetaDeletedBy),
		)
	}
	paymentInstructionsSelectableResponse := PaymentInstructionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentInstructionsDBFieldName.Id):
			key := string(PaymentInstructionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.Id, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.PaymentAttemptId):
			key := string(PaymentInstructionsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.PaymentAttemptId, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.InstructionType):
			key := string(PaymentInstructionsDTOFieldName.InstructionType)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.InstructionType, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.IsActive):
			key := string(PaymentInstructionsDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.IsActive, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.DisplayName):
			key := string(PaymentInstructionsDTOFieldName.DisplayName)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.DisplayName.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.AccountNumber):
			key := string(PaymentInstructionsDTOFieldName.AccountNumber)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.AccountNumber.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.AccountNumberMasked):
			key := string(PaymentInstructionsDTOFieldName.AccountNumberMasked)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.AccountNumberMasked.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.AccountHolderName):
			key := string(PaymentInstructionsDTOFieldName.AccountHolderName)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.AccountHolderName.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.BankCode):
			key := string(PaymentInstructionsDTOFieldName.BankCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.BankCode.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.BillerCode):
			key := string(PaymentInstructionsDTOFieldName.BillerCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.BillerCode.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.PaymentCode):
			key := string(PaymentInstructionsDTOFieldName.PaymentCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.PaymentCode.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.QrString):
			key := string(PaymentInstructionsDTOFieldName.QrString)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.QrString.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.QrImageUrl):
			key := string(PaymentInstructionsDTOFieldName.QrImageUrl)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.QrImageUrl.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.CheckoutUrl):
			key := string(PaymentInstructionsDTOFieldName.CheckoutUrl)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.CheckoutUrl.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.DeeplinkUrl):
			key := string(PaymentInstructionsDTOFieldName.DeeplinkUrl)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.DeeplinkUrl.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.RetailOutletCode):
			key := string(PaymentInstructionsDTOFieldName.RetailOutletCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.RetailOutletCode.String, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.ExpiresAt):
			key := string(PaymentInstructionsDTOFieldName.ExpiresAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.ExpiresAt.Time, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.Metadata):
			key := string(PaymentInstructionsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.Metadata, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaCreatedAt):
			key := string(PaymentInstructionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaCreatedAt, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaCreatedBy):
			key := string(PaymentInstructionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaCreatedBy, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaUpdatedAt):
			key := string(PaymentInstructionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaUpdatedBy):
			key := string(PaymentInstructionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaDeletedAt):
			key := string(PaymentInstructionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentInstructionsDBFieldName.MetaDeletedBy):
			key := string(PaymentInstructionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstructionsSelectableValue(paymentInstructionsSelectableResponse, key, paymentInstructions.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentInstructionsSelectableResponse
}

func NewPaymentInstructionsListResponseFromFilterResult(result []model.PaymentInstructionsFilterResult, filter model.Filter) PaymentInstructionsSelectableListResponse {
	dtoPaymentInstructionsListResponse := PaymentInstructionsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentInstructionsResponse := NewPaymentInstructionsSelectableResponse(row.PaymentInstructions, filter)
		dtoPaymentInstructionsListResponse = append(dtoPaymentInstructionsListResponse, &dtoPaymentInstructionsResponse)
	}
	return dtoPaymentInstructionsListResponse
}

type PaymentInstructionsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     PaymentInstructionsSelectableListResponse `json:"data"`
}

func reversePaymentInstructionsFilterResults(result []model.PaymentInstructionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentInstructionsFilterResponse(result []model.PaymentInstructionsFilterResult, filter model.Filter) (resp PaymentInstructionsFilterResponse) {
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
			reversePaymentInstructionsFilterResults(dataResult)
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

	resp.Data = NewPaymentInstructionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentInstructionsCreateRequest struct {
	PaymentAttemptId    uuid.UUID       `json:"paymentAttemptId"`
	InstructionType     string          `json:"instructionType"`
	IsActive            bool            `json:"isActive"`
	DisplayName         string          `json:"displayName"`
	AccountNumber       string          `json:"accountNumber"`
	AccountNumberMasked string          `json:"accountNumberMasked"`
	AccountHolderName   string          `json:"accountHolderName"`
	BankCode            string          `json:"bankCode"`
	BillerCode          string          `json:"billerCode"`
	PaymentCode         string          `json:"paymentCode"`
	QrString            string          `json:"qrString"`
	QrImageUrl          string          `json:"qrImageUrl"`
	CheckoutUrl         string          `json:"checkoutUrl"`
	DeeplinkUrl         string          `json:"deeplinkUrl"`
	RetailOutletCode    string          `json:"retailOutletCode"`
	ExpiresAt           time.Time       `json:"expiresAt"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d *PaymentInstructionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentInstructionsCreateRequest) ToModel() model.PaymentInstructions {
	id, _ := uuid.NewV4()
	return model.PaymentInstructions{
		Id:                  id,
		PaymentAttemptId:    d.PaymentAttemptId,
		InstructionType:     d.InstructionType,
		IsActive:            d.IsActive,
		DisplayName:         null.StringFrom(d.DisplayName),
		AccountNumber:       null.StringFrom(d.AccountNumber),
		AccountNumberMasked: null.StringFrom(d.AccountNumberMasked),
		AccountHolderName:   null.StringFrom(d.AccountHolderName),
		BankCode:            null.StringFrom(d.BankCode),
		BillerCode:          null.StringFrom(d.BillerCode),
		PaymentCode:         null.StringFrom(d.PaymentCode),
		QrString:            null.StringFrom(d.QrString),
		QrImageUrl:          null.StringFrom(d.QrImageUrl),
		CheckoutUrl:         null.StringFrom(d.CheckoutUrl),
		DeeplinkUrl:         null.StringFrom(d.DeeplinkUrl),
		RetailOutletCode:    null.StringFrom(d.RetailOutletCode),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		Metadata:            d.Metadata,
	}
}

type PaymentInstructionsListCreateRequest []*PaymentInstructionsCreateRequest

func (d PaymentInstructionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstructions := range d {
		err = validator.Struct(paymentInstructions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentInstructionsListCreateRequest) ToModelList() []model.PaymentInstructions {
	out := make([]model.PaymentInstructions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentInstructionsUpdateRequest struct {
	PaymentAttemptId    uuid.UUID       `json:"paymentAttemptId"`
	InstructionType     string          `json:"instructionType"`
	IsActive            bool            `json:"isActive"`
	DisplayName         string          `json:"displayName"`
	AccountNumber       string          `json:"accountNumber"`
	AccountNumberMasked string          `json:"accountNumberMasked"`
	AccountHolderName   string          `json:"accountHolderName"`
	BankCode            string          `json:"bankCode"`
	BillerCode          string          `json:"billerCode"`
	PaymentCode         string          `json:"paymentCode"`
	QrString            string          `json:"qrString"`
	QrImageUrl          string          `json:"qrImageUrl"`
	CheckoutUrl         string          `json:"checkoutUrl"`
	DeeplinkUrl         string          `json:"deeplinkUrl"`
	RetailOutletCode    string          `json:"retailOutletCode"`
	ExpiresAt           time.Time       `json:"expiresAt"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d *PaymentInstructionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentInstructionsUpdateRequest) ToModel() model.PaymentInstructions {
	return model.PaymentInstructions{
		PaymentAttemptId:    d.PaymentAttemptId,
		InstructionType:     d.InstructionType,
		IsActive:            d.IsActive,
		DisplayName:         null.StringFrom(d.DisplayName),
		AccountNumber:       null.StringFrom(d.AccountNumber),
		AccountNumberMasked: null.StringFrom(d.AccountNumberMasked),
		AccountHolderName:   null.StringFrom(d.AccountHolderName),
		BankCode:            null.StringFrom(d.BankCode),
		BillerCode:          null.StringFrom(d.BillerCode),
		PaymentCode:         null.StringFrom(d.PaymentCode),
		QrString:            null.StringFrom(d.QrString),
		QrImageUrl:          null.StringFrom(d.QrImageUrl),
		CheckoutUrl:         null.StringFrom(d.CheckoutUrl),
		DeeplinkUrl:         null.StringFrom(d.DeeplinkUrl),
		RetailOutletCode:    null.StringFrom(d.RetailOutletCode),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		Metadata:            d.Metadata,
	}
}

type PaymentInstructionsBulkUpdateRequest struct {
	Id                  uuid.UUID       `json:"id"`
	PaymentAttemptId    uuid.UUID       `json:"paymentAttemptId"`
	InstructionType     string          `json:"instructionType"`
	IsActive            bool            `json:"isActive"`
	DisplayName         string          `json:"displayName"`
	AccountNumber       string          `json:"accountNumber"`
	AccountNumberMasked string          `json:"accountNumberMasked"`
	AccountHolderName   string          `json:"accountHolderName"`
	BankCode            string          `json:"bankCode"`
	BillerCode          string          `json:"billerCode"`
	PaymentCode         string          `json:"paymentCode"`
	QrString            string          `json:"qrString"`
	QrImageUrl          string          `json:"qrImageUrl"`
	CheckoutUrl         string          `json:"checkoutUrl"`
	DeeplinkUrl         string          `json:"deeplinkUrl"`
	RetailOutletCode    string          `json:"retailOutletCode"`
	ExpiresAt           time.Time       `json:"expiresAt"`
	Metadata            json.RawMessage `json:"metadata"`
}

func (d PaymentInstructionsBulkUpdateRequest) PrimaryID() PaymentInstructionsPrimaryID {
	return PaymentInstructionsPrimaryID{
		Id: d.Id,
	}
}

type PaymentInstructionsListBulkUpdateRequest []*PaymentInstructionsBulkUpdateRequest

func (d PaymentInstructionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstructions := range d {
		err = validator.Struct(paymentInstructions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentInstructionsBulkUpdateRequest) ToModel() model.PaymentInstructions {
	return model.PaymentInstructions{
		Id:                  d.Id,
		PaymentAttemptId:    d.PaymentAttemptId,
		InstructionType:     d.InstructionType,
		IsActive:            d.IsActive,
		DisplayName:         null.StringFrom(d.DisplayName),
		AccountNumber:       null.StringFrom(d.AccountNumber),
		AccountNumberMasked: null.StringFrom(d.AccountNumberMasked),
		AccountHolderName:   null.StringFrom(d.AccountHolderName),
		BankCode:            null.StringFrom(d.BankCode),
		BillerCode:          null.StringFrom(d.BillerCode),
		PaymentCode:         null.StringFrom(d.PaymentCode),
		QrString:            null.StringFrom(d.QrString),
		QrImageUrl:          null.StringFrom(d.QrImageUrl),
		CheckoutUrl:         null.StringFrom(d.CheckoutUrl),
		DeeplinkUrl:         null.StringFrom(d.DeeplinkUrl),
		RetailOutletCode:    null.StringFrom(d.RetailOutletCode),
		ExpiresAt:           null.TimeFrom(d.ExpiresAt),
		Metadata:            d.Metadata,
	}
}

type PaymentInstructionsResponse struct {
	Id                  uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId    uuid.UUID       `json:"paymentAttemptId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	InstructionType     string          `json:"instructionType" validate:"required"`
	IsActive            bool            `json:"isActive" example:"true"`
	DisplayName         string          `json:"displayName"`
	AccountNumber       string          `json:"accountNumber"`
	AccountNumberMasked string          `json:"accountNumberMasked"`
	AccountHolderName   string          `json:"accountHolderName"`
	BankCode            string          `json:"bankCode"`
	BillerCode          string          `json:"billerCode"`
	PaymentCode         string          `json:"paymentCode"`
	QrString            string          `json:"qrString"`
	QrImageUrl          string          `json:"qrImageUrl" validate:"url"`
	CheckoutUrl         string          `json:"checkoutUrl" validate:"url"`
	DeeplinkUrl         string          `json:"deeplinkUrl" validate:"url"`
	RetailOutletCode    string          `json:"retailOutletCode"`
	ExpiresAt           time.Time       `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata            json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewPaymentInstructionsResponse(paymentInstructions model.PaymentInstructions) PaymentInstructionsResponse {
	return PaymentInstructionsResponse{
		Id:                  paymentInstructions.Id,
		PaymentAttemptId:    paymentInstructions.PaymentAttemptId,
		InstructionType:     paymentInstructions.InstructionType,
		IsActive:            paymentInstructions.IsActive,
		DisplayName:         paymentInstructions.DisplayName.String,
		AccountNumber:       paymentInstructions.AccountNumber.String,
		AccountNumberMasked: paymentInstructions.AccountNumberMasked.String,
		AccountHolderName:   paymentInstructions.AccountHolderName.String,
		BankCode:            paymentInstructions.BankCode.String,
		BillerCode:          paymentInstructions.BillerCode.String,
		PaymentCode:         paymentInstructions.PaymentCode.String,
		QrString:            paymentInstructions.QrString.String,
		QrImageUrl:          paymentInstructions.QrImageUrl.String,
		CheckoutUrl:         paymentInstructions.CheckoutUrl.String,
		DeeplinkUrl:         paymentInstructions.DeeplinkUrl.String,
		RetailOutletCode:    paymentInstructions.RetailOutletCode.String,
		ExpiresAt:           paymentInstructions.ExpiresAt.Time,
		Metadata:            paymentInstructions.Metadata,
	}
}

type PaymentInstructionsListResponse []*PaymentInstructionsResponse

func NewPaymentInstructionsListResponse(paymentInstructionsList model.PaymentInstructionsList) PaymentInstructionsListResponse {
	dtoPaymentInstructionsListResponse := PaymentInstructionsListResponse{}
	for _, paymentInstructions := range paymentInstructionsList {
		dtoPaymentInstructionsResponse := NewPaymentInstructionsResponse(*paymentInstructions)
		dtoPaymentInstructionsListResponse = append(dtoPaymentInstructionsListResponse, &dtoPaymentInstructionsResponse)
	}
	return dtoPaymentInstructionsListResponse
}

type PaymentInstructionsPrimaryIDList []PaymentInstructionsPrimaryID

func (d PaymentInstructionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstructions := range d {
		err = validator.Struct(paymentInstructions)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentInstructionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentInstructionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentInstructionsPrimaryID) ToModel() model.PaymentInstructionsPrimaryID {
	return model.PaymentInstructionsPrimaryID{
		Id: d.Id,
	}
}

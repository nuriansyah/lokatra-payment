package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentAuthorizationsDTOFieldNameType string

type paymentAuthorizationsDTOFieldName struct {
	Id                      PaymentAuthorizationsDTOFieldNameType
	PaymentIntentId         PaymentAuthorizationsDTOFieldNameType
	PaymentAttemptId        PaymentAuthorizationsDTOFieldNameType
	ProviderAccountId       PaymentAuthorizationsDTOFieldNameType
	ProviderAuthorizationId PaymentAuthorizationsDTOFieldNameType
	Amount                  PaymentAuthorizationsDTOFieldNameType
	Currency                PaymentAuthorizationsDTOFieldNameType
	Status                  PaymentAuthorizationsDTOFieldNameType
	AuthorizedAt            PaymentAuthorizationsDTOFieldNameType
	ExpiresAt               PaymentAuthorizationsDTOFieldNameType
	CapturedAmount          PaymentAuthorizationsDTOFieldNameType
	FailureCode             PaymentAuthorizationsDTOFieldNameType
	FailureMessage          PaymentAuthorizationsDTOFieldNameType
	RawRequest              PaymentAuthorizationsDTOFieldNameType
	RawResponse             PaymentAuthorizationsDTOFieldNameType
	Metadata                PaymentAuthorizationsDTOFieldNameType
	MetaCreatedAt           PaymentAuthorizationsDTOFieldNameType
	MetaCreatedBy           PaymentAuthorizationsDTOFieldNameType
	MetaUpdatedAt           PaymentAuthorizationsDTOFieldNameType
	MetaUpdatedBy           PaymentAuthorizationsDTOFieldNameType
	MetaDeletedAt           PaymentAuthorizationsDTOFieldNameType
	MetaDeletedBy           PaymentAuthorizationsDTOFieldNameType
}

var PaymentAuthorizationsDTOFieldName = paymentAuthorizationsDTOFieldName{
	Id:                      "id",
	PaymentIntentId:         "paymentIntentId",
	PaymentAttemptId:        "paymentAttemptId",
	ProviderAccountId:       "providerAccountId",
	ProviderAuthorizationId: "providerAuthorizationId",
	Amount:                  "amount",
	Currency:                "currency",
	Status:                  "status",
	AuthorizedAt:            "authorizedAt",
	ExpiresAt:               "expiresAt",
	CapturedAmount:          "capturedAmount",
	FailureCode:             "failureCode",
	FailureMessage:          "failureMessage",
	RawRequest:              "rawRequest",
	RawResponse:             "rawResponse",
	Metadata:                "metadata",
	MetaCreatedAt:           "metaCreatedAt",
	MetaCreatedBy:           "metaCreatedBy",
	MetaUpdatedAt:           "metaUpdatedAt",
	MetaUpdatedBy:           "metaUpdatedBy",
	MetaDeletedAt:           "metaDeletedAt",
	MetaDeletedBy:           "metaDeletedBy",
}

func transformPaymentAuthorizationsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentAuthorizationsDTOFieldName.Id):
		return string(model.PaymentAuthorizationsDBFieldName.Id), true

	case string(PaymentAuthorizationsDTOFieldName.PaymentIntentId):
		return string(model.PaymentAuthorizationsDBFieldName.PaymentIntentId), true

	case string(PaymentAuthorizationsDTOFieldName.PaymentAttemptId):
		return string(model.PaymentAuthorizationsDBFieldName.PaymentAttemptId), true

	case string(PaymentAuthorizationsDTOFieldName.ProviderAccountId):
		return string(model.PaymentAuthorizationsDBFieldName.ProviderAccountId), true

	case string(PaymentAuthorizationsDTOFieldName.ProviderAuthorizationId):
		return string(model.PaymentAuthorizationsDBFieldName.ProviderAuthorizationId), true

	case string(PaymentAuthorizationsDTOFieldName.Amount):
		return string(model.PaymentAuthorizationsDBFieldName.Amount), true

	case string(PaymentAuthorizationsDTOFieldName.Currency):
		return string(model.PaymentAuthorizationsDBFieldName.Currency), true

	case string(PaymentAuthorizationsDTOFieldName.Status):
		return string(model.PaymentAuthorizationsDBFieldName.Status), true

	case string(PaymentAuthorizationsDTOFieldName.AuthorizedAt):
		return string(model.PaymentAuthorizationsDBFieldName.AuthorizedAt), true

	case string(PaymentAuthorizationsDTOFieldName.ExpiresAt):
		return string(model.PaymentAuthorizationsDBFieldName.ExpiresAt), true

	case string(PaymentAuthorizationsDTOFieldName.CapturedAmount):
		return string(model.PaymentAuthorizationsDBFieldName.CapturedAmount), true

	case string(PaymentAuthorizationsDTOFieldName.FailureCode):
		return string(model.PaymentAuthorizationsDBFieldName.FailureCode), true

	case string(PaymentAuthorizationsDTOFieldName.FailureMessage):
		return string(model.PaymentAuthorizationsDBFieldName.FailureMessage), true

	case string(PaymentAuthorizationsDTOFieldName.RawRequest):
		return string(model.PaymentAuthorizationsDBFieldName.RawRequest), true

	case string(PaymentAuthorizationsDTOFieldName.RawResponse):
		return string(model.PaymentAuthorizationsDBFieldName.RawResponse), true

	case string(PaymentAuthorizationsDTOFieldName.Metadata):
		return string(model.PaymentAuthorizationsDBFieldName.Metadata), true

	case string(PaymentAuthorizationsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentAuthorizationsDBFieldName.MetaCreatedAt), true

	case string(PaymentAuthorizationsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentAuthorizationsDBFieldName.MetaCreatedBy), true

	case string(PaymentAuthorizationsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedAt), true

	case string(PaymentAuthorizationsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedBy), true

	case string(PaymentAuthorizationsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentAuthorizationsDBFieldName.MetaDeletedAt), true

	case string(PaymentAuthorizationsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentAuthorizationsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentAuthorizationsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentAuthorizationsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentAuthorizationsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentAuthorizationsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentAuthorizationsProjectionOutputPath(path string) error {
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

func transformPaymentAuthorizationsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentAuthorizationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentAuthorizationsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentAuthorizationsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentAuthorizationsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentAuthorizationsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentAuthorizationsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentAuthorizationsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentAuthorizationsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentAuthorizationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentAuthorizationsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentAuthorizationsFilter(filter *model.Filter) {
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
			Field: string(PaymentAuthorizationsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentAuthorizationsSelectableResponse map[string]interface{}
type PaymentAuthorizationsSelectableListResponse []*PaymentAuthorizationsSelectableResponse

func assignPaymentAuthorizationsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentAuthorizationsSelectableValue(out PaymentAuthorizationsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentAuthorizationsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentAuthorizationsSelectableResponse(paymentAuthorizations model.PaymentAuthorizations, filter model.Filter) PaymentAuthorizationsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentAuthorizationsDBFieldName.Id),
			string(model.PaymentAuthorizationsDBFieldName.PaymentIntentId),
			string(model.PaymentAuthorizationsDBFieldName.PaymentAttemptId),
			string(model.PaymentAuthorizationsDBFieldName.ProviderAccountId),
			string(model.PaymentAuthorizationsDBFieldName.ProviderAuthorizationId),
			string(model.PaymentAuthorizationsDBFieldName.Amount),
			string(model.PaymentAuthorizationsDBFieldName.Currency),
			string(model.PaymentAuthorizationsDBFieldName.Status),
			string(model.PaymentAuthorizationsDBFieldName.AuthorizedAt),
			string(model.PaymentAuthorizationsDBFieldName.ExpiresAt),
			string(model.PaymentAuthorizationsDBFieldName.CapturedAmount),
			string(model.PaymentAuthorizationsDBFieldName.FailureCode),
			string(model.PaymentAuthorizationsDBFieldName.FailureMessage),
			string(model.PaymentAuthorizationsDBFieldName.RawRequest),
			string(model.PaymentAuthorizationsDBFieldName.RawResponse),
			string(model.PaymentAuthorizationsDBFieldName.Metadata),
			string(model.PaymentAuthorizationsDBFieldName.MetaCreatedAt),
			string(model.PaymentAuthorizationsDBFieldName.MetaCreatedBy),
			string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedAt),
			string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedBy),
			string(model.PaymentAuthorizationsDBFieldName.MetaDeletedAt),
			string(model.PaymentAuthorizationsDBFieldName.MetaDeletedBy),
		)
	}
	paymentAuthorizationsSelectableResponse := PaymentAuthorizationsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentAuthorizationsDBFieldName.Id):
			key := string(PaymentAuthorizationsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.Id, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.PaymentIntentId):
			key := string(PaymentAuthorizationsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.PaymentIntentId, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.PaymentAttemptId):
			key := string(PaymentAuthorizationsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.PaymentAttemptId.UUID, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.ProviderAccountId):
			key := string(PaymentAuthorizationsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.ProviderAccountId.UUID, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.ProviderAuthorizationId):
			key := string(PaymentAuthorizationsDTOFieldName.ProviderAuthorizationId)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.ProviderAuthorizationId.String, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.Amount):
			key := string(PaymentAuthorizationsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.Amount, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.Currency):
			key := string(PaymentAuthorizationsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.Currency, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.Status):
			key := string(PaymentAuthorizationsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, model.PaymentAuthorizationStatus(paymentAuthorizations.Status), explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.AuthorizedAt):
			key := string(PaymentAuthorizationsDTOFieldName.AuthorizedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.AuthorizedAt.Time, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.ExpiresAt):
			key := string(PaymentAuthorizationsDTOFieldName.ExpiresAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.ExpiresAt.Time, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.CapturedAmount):
			key := string(PaymentAuthorizationsDTOFieldName.CapturedAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.CapturedAmount, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.FailureCode):
			key := string(PaymentAuthorizationsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.FailureCode.String, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.FailureMessage):
			key := string(PaymentAuthorizationsDTOFieldName.FailureMessage)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.FailureMessage.String, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.RawRequest):
			key := string(PaymentAuthorizationsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.RawRequest, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.RawResponse):
			key := string(PaymentAuthorizationsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.RawResponse, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.Metadata):
			key := string(PaymentAuthorizationsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.Metadata, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaCreatedAt):
			key := string(PaymentAuthorizationsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaCreatedAt, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaCreatedBy):
			key := string(PaymentAuthorizationsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaCreatedBy, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedAt):
			key := string(PaymentAuthorizationsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaUpdatedBy):
			key := string(PaymentAuthorizationsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaDeletedAt):
			key := string(PaymentAuthorizationsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentAuthorizationsDBFieldName.MetaDeletedBy):
			key := string(PaymentAuthorizationsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentAuthorizationsSelectableValue(paymentAuthorizationsSelectableResponse, key, paymentAuthorizations.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentAuthorizationsSelectableResponse
}

func NewPaymentAuthorizationsListResponseFromFilterResult(result []model.PaymentAuthorizationsFilterResult, filter model.Filter) PaymentAuthorizationsSelectableListResponse {
	dtoPaymentAuthorizationsListResponse := PaymentAuthorizationsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentAuthorizationsResponse := NewPaymentAuthorizationsSelectableResponse(row.PaymentAuthorizations, filter)
		dtoPaymentAuthorizationsListResponse = append(dtoPaymentAuthorizationsListResponse, &dtoPaymentAuthorizationsResponse)
	}
	return dtoPaymentAuthorizationsListResponse
}

type PaymentAuthorizationsFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     PaymentAuthorizationsSelectableListResponse `json:"data"`
}

func reversePaymentAuthorizationsFilterResults(result []model.PaymentAuthorizationsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentAuthorizationsFilterResponse(result []model.PaymentAuthorizationsFilterResult, filter model.Filter) (resp PaymentAuthorizationsFilterResponse) {
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
			reversePaymentAuthorizationsFilterResults(dataResult)
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

	resp.Data = NewPaymentAuthorizationsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentAuthorizationsCreateRequest struct {
	PaymentIntentId         uuid.UUID                        `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID                        `json:"paymentAttemptId"`
	ProviderAccountId       uuid.UUID                        `json:"providerAccountId"`
	ProviderAuthorizationId string                           `json:"providerAuthorizationId"`
	Amount                  decimal.Decimal                  `json:"amount"`
	Currency                string                           `json:"currency"`
	Status                  model.PaymentAuthorizationStatus `json:"status" example:"requested" enums:"requested,authorized,captured,voided,failed"`
	AuthorizedAt            time.Time                        `json:"authorizedAt"`
	ExpiresAt               time.Time                        `json:"expiresAt"`
	CapturedAmount          decimal.Decimal                  `json:"capturedAmount"`
	FailureCode             string                           `json:"failureCode"`
	FailureMessage          string                           `json:"failureMessage"`
	RawRequest              json.RawMessage                  `json:"rawRequest"`
	RawResponse             json.RawMessage                  `json:"rawResponse"`
	Metadata                json.RawMessage                  `json:"metadata"`
}

func (d *PaymentAuthorizationsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentAuthorizationsCreateRequest) ToModel() model.PaymentAuthorizations {
	id, _ := uuid.NewV4()
	return model.PaymentAuthorizations{
		Id:                      id,
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		ProviderAccountId:       nuuid.From(d.ProviderAccountId),
		ProviderAuthorizationId: null.StringFrom(d.ProviderAuthorizationId),
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		AuthorizedAt:            null.TimeFrom(d.AuthorizedAt),
		ExpiresAt:               null.TimeFrom(d.ExpiresAt),
		CapturedAmount:          d.CapturedAmount,
		FailureCode:             null.StringFrom(d.FailureCode),
		FailureMessage:          null.StringFrom(d.FailureMessage),
		RawRequest:              d.RawRequest,
		RawResponse:             d.RawResponse,
		Metadata:                d.Metadata,
	}
}

type PaymentAuthorizationsListCreateRequest []*PaymentAuthorizationsCreateRequest

func (d PaymentAuthorizationsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAuthorizations := range d {
		err = validator.Struct(paymentAuthorizations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentAuthorizationsListCreateRequest) ToModelList() []model.PaymentAuthorizations {
	out := make([]model.PaymentAuthorizations, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentAuthorizationsUpdateRequest struct {
	PaymentIntentId         uuid.UUID                        `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID                        `json:"paymentAttemptId"`
	ProviderAccountId       uuid.UUID                        `json:"providerAccountId"`
	ProviderAuthorizationId string                           `json:"providerAuthorizationId"`
	Amount                  decimal.Decimal                  `json:"amount"`
	Currency                string                           `json:"currency"`
	Status                  model.PaymentAuthorizationStatus `json:"status" example:"requested" enums:"requested,authorized,captured,voided,failed"`
	AuthorizedAt            time.Time                        `json:"authorizedAt"`
	ExpiresAt               time.Time                        `json:"expiresAt"`
	CapturedAmount          decimal.Decimal                  `json:"capturedAmount"`
	FailureCode             string                           `json:"failureCode"`
	FailureMessage          string                           `json:"failureMessage"`
	RawRequest              json.RawMessage                  `json:"rawRequest"`
	RawResponse             json.RawMessage                  `json:"rawResponse"`
	Metadata                json.RawMessage                  `json:"metadata"`
}

func (d *PaymentAuthorizationsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentAuthorizationsUpdateRequest) ToModel() model.PaymentAuthorizations {
	return model.PaymentAuthorizations{
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		ProviderAccountId:       nuuid.From(d.ProviderAccountId),
		ProviderAuthorizationId: null.StringFrom(d.ProviderAuthorizationId),
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		AuthorizedAt:            null.TimeFrom(d.AuthorizedAt),
		ExpiresAt:               null.TimeFrom(d.ExpiresAt),
		CapturedAmount:          d.CapturedAmount,
		FailureCode:             null.StringFrom(d.FailureCode),
		FailureMessage:          null.StringFrom(d.FailureMessage),
		RawRequest:              d.RawRequest,
		RawResponse:             d.RawResponse,
		Metadata:                d.Metadata,
	}
}

type PaymentAuthorizationsBulkUpdateRequest struct {
	Id                      uuid.UUID                        `json:"id"`
	PaymentIntentId         uuid.UUID                        `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID                        `json:"paymentAttemptId"`
	ProviderAccountId       uuid.UUID                        `json:"providerAccountId"`
	ProviderAuthorizationId string                           `json:"providerAuthorizationId"`
	Amount                  decimal.Decimal                  `json:"amount"`
	Currency                string                           `json:"currency"`
	Status                  model.PaymentAuthorizationStatus `json:"status" example:"requested" enums:"requested,authorized,captured,voided,failed"`
	AuthorizedAt            time.Time                        `json:"authorizedAt"`
	ExpiresAt               time.Time                        `json:"expiresAt"`
	CapturedAmount          decimal.Decimal                  `json:"capturedAmount"`
	FailureCode             string                           `json:"failureCode"`
	FailureMessage          string                           `json:"failureMessage"`
	RawRequest              json.RawMessage                  `json:"rawRequest"`
	RawResponse             json.RawMessage                  `json:"rawResponse"`
	Metadata                json.RawMessage                  `json:"metadata"`
}

func (d PaymentAuthorizationsBulkUpdateRequest) PrimaryID() PaymentAuthorizationsPrimaryID {
	return PaymentAuthorizationsPrimaryID{
		Id: d.Id,
	}
}

type PaymentAuthorizationsListBulkUpdateRequest []*PaymentAuthorizationsBulkUpdateRequest

func (d PaymentAuthorizationsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAuthorizations := range d {
		err = validator.Struct(paymentAuthorizations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentAuthorizationsBulkUpdateRequest) ToModel() model.PaymentAuthorizations {
	return model.PaymentAuthorizations{
		Id:                      d.Id,
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		ProviderAccountId:       nuuid.From(d.ProviderAccountId),
		ProviderAuthorizationId: null.StringFrom(d.ProviderAuthorizationId),
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		AuthorizedAt:            null.TimeFrom(d.AuthorizedAt),
		ExpiresAt:               null.TimeFrom(d.ExpiresAt),
		CapturedAmount:          d.CapturedAmount,
		FailureCode:             null.StringFrom(d.FailureCode),
		FailureMessage:          null.StringFrom(d.FailureMessage),
		RawRequest:              d.RawRequest,
		RawResponse:             d.RawResponse,
		Metadata:                d.Metadata,
	}
}

type PaymentAuthorizationsResponse struct {
	Id                      uuid.UUID                        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId         uuid.UUID                        `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId        uuid.UUID                        `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId       uuid.UUID                        `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAuthorizationId string                           `json:"providerAuthorizationId"`
	Amount                  decimal.Decimal                  `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency                string                           `json:"currency"`
	Status                  model.PaymentAuthorizationStatus `json:"status" validate:"oneof=requested authorized captured voided failed" enums:"requested,authorized,captured,voided,failed"`
	AuthorizedAt            time.Time                        `json:"authorizedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ExpiresAt               time.Time                        `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CapturedAmount          decimal.Decimal                  `json:"capturedAmount" format:"decimal" example:"100.50"`
	FailureCode             string                           `json:"failureCode"`
	FailureMessage          string                           `json:"failureMessage"`
	RawRequest              json.RawMessage                  `json:"rawRequest" swaggertype:"object"`
	RawResponse             json.RawMessage                  `json:"rawResponse" swaggertype:"object"`
	Metadata                json.RawMessage                  `json:"metadata" swaggertype:"object"`
}

func NewPaymentAuthorizationsResponse(paymentAuthorizations model.PaymentAuthorizations) PaymentAuthorizationsResponse {
	return PaymentAuthorizationsResponse{
		Id:                      paymentAuthorizations.Id,
		PaymentIntentId:         paymentAuthorizations.PaymentIntentId,
		PaymentAttemptId:        paymentAuthorizations.PaymentAttemptId.UUID,
		ProviderAccountId:       paymentAuthorizations.ProviderAccountId.UUID,
		ProviderAuthorizationId: paymentAuthorizations.ProviderAuthorizationId.String,
		Amount:                  paymentAuthorizations.Amount,
		Currency:                paymentAuthorizations.Currency,
		Status:                  model.PaymentAuthorizationStatus(paymentAuthorizations.Status),
		AuthorizedAt:            paymentAuthorizations.AuthorizedAt.Time,
		ExpiresAt:               paymentAuthorizations.ExpiresAt.Time,
		CapturedAmount:          paymentAuthorizations.CapturedAmount,
		FailureCode:             paymentAuthorizations.FailureCode.String,
		FailureMessage:          paymentAuthorizations.FailureMessage.String,
		RawRequest:              paymentAuthorizations.RawRequest,
		RawResponse:             paymentAuthorizations.RawResponse,
		Metadata:                paymentAuthorizations.Metadata,
	}
}

type PaymentAuthorizationsListResponse []*PaymentAuthorizationsResponse

func NewPaymentAuthorizationsListResponse(paymentAuthorizationsList model.PaymentAuthorizationsList) PaymentAuthorizationsListResponse {
	dtoPaymentAuthorizationsListResponse := PaymentAuthorizationsListResponse{}
	for _, paymentAuthorizations := range paymentAuthorizationsList {
		dtoPaymentAuthorizationsResponse := NewPaymentAuthorizationsResponse(*paymentAuthorizations)
		dtoPaymentAuthorizationsListResponse = append(dtoPaymentAuthorizationsListResponse, &dtoPaymentAuthorizationsResponse)
	}
	return dtoPaymentAuthorizationsListResponse
}

type PaymentAuthorizationsPrimaryIDList []PaymentAuthorizationsPrimaryID

func (d PaymentAuthorizationsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentAuthorizations := range d {
		err = validator.Struct(paymentAuthorizations)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentAuthorizationsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentAuthorizationsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentAuthorizationsPrimaryID) ToModel() model.PaymentAuthorizationsPrimaryID {
	return model.PaymentAuthorizationsPrimaryID{
		Id: d.Id,
	}
}

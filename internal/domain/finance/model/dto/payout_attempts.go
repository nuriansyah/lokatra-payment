package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PayoutAttemptsDTOFieldNameType string

type payoutAttemptsDTOFieldName struct {
	Id                PayoutAttemptsDTOFieldNameType
	PayoutId          PayoutAttemptsDTOFieldNameType
	AttemptNo         PayoutAttemptsDTOFieldNameType
	AttemptType       PayoutAttemptsDTOFieldNameType
	ProviderAccountId PayoutAttemptsDTOFieldNameType
	Amount            PayoutAttemptsDTOFieldNameType
	CurrencyCode      PayoutAttemptsDTOFieldNameType
	AttemptStatus     PayoutAttemptsDTOFieldNameType
	ProviderPayoutRef PayoutAttemptsDTOFieldNameType
	FailureCode       PayoutAttemptsDTOFieldNameType
	FailureReason     PayoutAttemptsDTOFieldNameType
	RawRequest        PayoutAttemptsDTOFieldNameType
	RawResponse       PayoutAttemptsDTOFieldNameType
	Metadata          PayoutAttemptsDTOFieldNameType
	MetaCreatedAt     PayoutAttemptsDTOFieldNameType
	MetaCreatedBy     PayoutAttemptsDTOFieldNameType
	MetaUpdatedAt     PayoutAttemptsDTOFieldNameType
	MetaUpdatedBy     PayoutAttemptsDTOFieldNameType
	MetaDeletedAt     PayoutAttemptsDTOFieldNameType
	MetaDeletedBy     PayoutAttemptsDTOFieldNameType
}

var PayoutAttemptsDTOFieldName = payoutAttemptsDTOFieldName{
	Id:                "id",
	PayoutId:          "payoutId",
	AttemptNo:         "attemptNo",
	AttemptType:       "attemptType",
	ProviderAccountId: "providerAccountId",
	Amount:            "amount",
	CurrencyCode:      "currencyCode",
	AttemptStatus:     "attemptStatus",
	ProviderPayoutRef: "providerPayoutRef",
	FailureCode:       "failureCode",
	FailureReason:     "failureReason",
	RawRequest:        "rawRequest",
	RawResponse:       "rawResponse",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformPayoutAttemptsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutAttemptsDTOFieldName.Id):
		return string(model.PayoutAttemptsDBFieldName.Id), true

	case string(PayoutAttemptsDTOFieldName.PayoutId):
		return string(model.PayoutAttemptsDBFieldName.PayoutId), true

	case string(PayoutAttemptsDTOFieldName.AttemptNo):
		return string(model.PayoutAttemptsDBFieldName.AttemptNo), true

	case string(PayoutAttemptsDTOFieldName.AttemptType):
		return string(model.PayoutAttemptsDBFieldName.AttemptType), true

	case string(PayoutAttemptsDTOFieldName.ProviderAccountId):
		return string(model.PayoutAttemptsDBFieldName.ProviderAccountId), true

	case string(PayoutAttemptsDTOFieldName.Amount):
		return string(model.PayoutAttemptsDBFieldName.Amount), true

	case string(PayoutAttemptsDTOFieldName.CurrencyCode):
		return string(model.PayoutAttemptsDBFieldName.CurrencyCode), true

	case string(PayoutAttemptsDTOFieldName.AttemptStatus):
		return string(model.PayoutAttemptsDBFieldName.AttemptStatus), true

	case string(PayoutAttemptsDTOFieldName.ProviderPayoutRef):
		return string(model.PayoutAttemptsDBFieldName.ProviderPayoutRef), true

	case string(PayoutAttemptsDTOFieldName.FailureCode):
		return string(model.PayoutAttemptsDBFieldName.FailureCode), true

	case string(PayoutAttemptsDTOFieldName.FailureReason):
		return string(model.PayoutAttemptsDBFieldName.FailureReason), true

	case string(PayoutAttemptsDTOFieldName.RawRequest):
		return string(model.PayoutAttemptsDBFieldName.RawRequest), true

	case string(PayoutAttemptsDTOFieldName.RawResponse):
		return string(model.PayoutAttemptsDBFieldName.RawResponse), true

	case string(PayoutAttemptsDTOFieldName.Metadata):
		return string(model.PayoutAttemptsDBFieldName.Metadata), true

	case string(PayoutAttemptsDTOFieldName.MetaCreatedAt):
		return string(model.PayoutAttemptsDBFieldName.MetaCreatedAt), true

	case string(PayoutAttemptsDTOFieldName.MetaCreatedBy):
		return string(model.PayoutAttemptsDBFieldName.MetaCreatedBy), true

	case string(PayoutAttemptsDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutAttemptsDBFieldName.MetaUpdatedAt), true

	case string(PayoutAttemptsDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutAttemptsDBFieldName.MetaUpdatedBy), true

	case string(PayoutAttemptsDTOFieldName.MetaDeletedAt):
		return string(model.PayoutAttemptsDBFieldName.MetaDeletedAt), true

	case string(PayoutAttemptsDTOFieldName.MetaDeletedBy):
		return string(model.PayoutAttemptsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutAttemptsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutAttemptsBaseFilterField(field string) bool {
	spec, found := model.NewPayoutAttemptsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutAttemptsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutAttemptsProjectionOutputPath(path string) error {
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

func transformPayoutAttemptsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutAttemptsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutAttemptsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutAttemptsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutAttemptsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutAttemptsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutAttemptsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutAttemptsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutAttemptsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutAttemptsFilter(filter *model.Filter) {
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
			Field: string(PayoutAttemptsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutAttemptsSelectableResponse map[string]interface{}
type PayoutAttemptsSelectableListResponse []*PayoutAttemptsSelectableResponse

func assignPayoutAttemptsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutAttemptsSelectableValue(out PayoutAttemptsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutAttemptsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutAttemptsSelectableResponse(payoutAttempts model.PayoutAttempts, filter model.Filter) PayoutAttemptsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutAttemptsDBFieldName.Id),
			string(model.PayoutAttemptsDBFieldName.PayoutId),
			string(model.PayoutAttemptsDBFieldName.AttemptNo),
			string(model.PayoutAttemptsDBFieldName.AttemptType),
			string(model.PayoutAttemptsDBFieldName.ProviderAccountId),
			string(model.PayoutAttemptsDBFieldName.Amount),
			string(model.PayoutAttemptsDBFieldName.CurrencyCode),
			string(model.PayoutAttemptsDBFieldName.AttemptStatus),
			string(model.PayoutAttemptsDBFieldName.ProviderPayoutRef),
			string(model.PayoutAttemptsDBFieldName.FailureCode),
			string(model.PayoutAttemptsDBFieldName.FailureReason),
			string(model.PayoutAttemptsDBFieldName.RawRequest),
			string(model.PayoutAttemptsDBFieldName.RawResponse),
			string(model.PayoutAttemptsDBFieldName.Metadata),
			string(model.PayoutAttemptsDBFieldName.MetaCreatedAt),
			string(model.PayoutAttemptsDBFieldName.MetaCreatedBy),
			string(model.PayoutAttemptsDBFieldName.MetaUpdatedAt),
			string(model.PayoutAttemptsDBFieldName.MetaUpdatedBy),
			string(model.PayoutAttemptsDBFieldName.MetaDeletedAt),
			string(model.PayoutAttemptsDBFieldName.MetaDeletedBy),
		)
	}
	payoutAttemptsSelectableResponse := PayoutAttemptsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutAttemptsDBFieldName.Id):
			key := string(PayoutAttemptsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.Id, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.PayoutId):
			key := string(PayoutAttemptsDTOFieldName.PayoutId)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.PayoutId, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.AttemptNo):
			key := string(PayoutAttemptsDTOFieldName.AttemptNo)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.AttemptNo, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.AttemptType):
			key := string(PayoutAttemptsDTOFieldName.AttemptType)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, model.PayoutAttemptsAttemptType(payoutAttempts.AttemptType), explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.ProviderAccountId):
			key := string(PayoutAttemptsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.ProviderAccountId.UUID, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.Amount):
			key := string(PayoutAttemptsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.Amount, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.CurrencyCode):
			key := string(PayoutAttemptsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.CurrencyCode, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.AttemptStatus):
			key := string(PayoutAttemptsDTOFieldName.AttemptStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, model.PayoutAttemptsAttemptStatus(payoutAttempts.AttemptStatus), explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.ProviderPayoutRef):
			key := string(PayoutAttemptsDTOFieldName.ProviderPayoutRef)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.ProviderPayoutRef.String, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.FailureCode):
			key := string(PayoutAttemptsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.FailureCode.String, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.FailureReason):
			key := string(PayoutAttemptsDTOFieldName.FailureReason)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.FailureReason.String, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.RawRequest):
			key := string(PayoutAttemptsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.RawRequest, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.RawResponse):
			key := string(PayoutAttemptsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.RawResponse, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.Metadata):
			key := string(PayoutAttemptsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.Metadata, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaCreatedAt):
			key := string(PayoutAttemptsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaCreatedAt, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaCreatedBy):
			key := string(PayoutAttemptsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaCreatedBy, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaUpdatedAt):
			key := string(PayoutAttemptsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaUpdatedBy):
			key := string(PayoutAttemptsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaDeletedAt):
			key := string(PayoutAttemptsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutAttemptsDBFieldName.MetaDeletedBy):
			key := string(PayoutAttemptsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutAttemptsSelectableValue(payoutAttemptsSelectableResponse, key, payoutAttempts.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutAttemptsSelectableResponse
}

func NewPayoutAttemptsListResponseFromFilterResult(result []model.PayoutAttemptsFilterResult, filter model.Filter) PayoutAttemptsSelectableListResponse {
	dtoPayoutAttemptsListResponse := PayoutAttemptsSelectableListResponse{}
	for _, row := range result {
		dtoPayoutAttemptsResponse := NewPayoutAttemptsSelectableResponse(row.PayoutAttempts, filter)
		dtoPayoutAttemptsListResponse = append(dtoPayoutAttemptsListResponse, &dtoPayoutAttemptsResponse)
	}
	return dtoPayoutAttemptsListResponse
}

type PayoutAttemptsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PayoutAttemptsSelectableListResponse `json:"data"`
}

func reversePayoutAttemptsFilterResults(result []model.PayoutAttemptsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutAttemptsFilterResponse(result []model.PayoutAttemptsFilterResult, filter model.Filter) (resp PayoutAttemptsFilterResponse) {
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
			reversePayoutAttemptsFilterResults(dataResult)
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

	resp.Data = NewPayoutAttemptsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutAttemptsCreateRequest struct {
	PayoutId          uuid.UUID                         `json:"payoutId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.PayoutAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.PayoutAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderPayoutRef string                            `json:"providerPayoutRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d *PayoutAttemptsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutAttemptsCreateRequest) ToModel() model.PayoutAttempts {
	id, _ := uuid.NewV4()
	return model.PayoutAttempts{
		Id:                id,
		PayoutId:          d.PayoutId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PayoutAttemptsListCreateRequest []*PayoutAttemptsCreateRequest

func (d PayoutAttemptsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutAttempts := range d {
		err = validator.Struct(payoutAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutAttemptsListCreateRequest) ToModelList() []model.PayoutAttempts {
	out := make([]model.PayoutAttempts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutAttemptsUpdateRequest struct {
	PayoutId          uuid.UUID                         `json:"payoutId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.PayoutAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.PayoutAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderPayoutRef string                            `json:"providerPayoutRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d *PayoutAttemptsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutAttemptsUpdateRequest) ToModel() model.PayoutAttempts {
	return model.PayoutAttempts{
		PayoutId:          d.PayoutId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PayoutAttemptsBulkUpdateRequest struct {
	Id                uuid.UUID                         `json:"id"`
	PayoutId          uuid.UUID                         `json:"payoutId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.PayoutAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.PayoutAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderPayoutRef string                            `json:"providerPayoutRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d PayoutAttemptsBulkUpdateRequest) PrimaryID() PayoutAttemptsPrimaryID {
	return PayoutAttemptsPrimaryID{
		Id: d.Id,
	}
}

type PayoutAttemptsListBulkUpdateRequest []*PayoutAttemptsBulkUpdateRequest

func (d PayoutAttemptsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutAttempts := range d {
		err = validator.Struct(payoutAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutAttemptsBulkUpdateRequest) ToModel() model.PayoutAttempts {
	return model.PayoutAttempts{
		Id:                d.Id,
		PayoutId:          d.PayoutId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type PayoutAttemptsResponse struct {
	Id                uuid.UUID                         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutId          uuid.UUID                         `json:"payoutId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AttemptNo         int                               `json:"attemptNo" validate:"required" example:"1"`
	AttemptType       model.PayoutAttemptsAttemptType   `json:"attemptType" validate:"required,oneof=provider manual bank_transfer" enums:"provider,manual,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount            decimal.Decimal                   `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	CurrencyCode      string                            `json:"currencyCode" validate:"required"`
	AttemptStatus     model.PayoutAttemptsAttemptStatus `json:"attemptStatus" validate:"required,oneof=created processing succeeded failed cancelled" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderPayoutRef string                            `json:"providerPayoutRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest" swaggertype:"object"`
	RawResponse       json.RawMessage                   `json:"rawResponse" swaggertype:"object"`
	Metadata          json.RawMessage                   `json:"metadata" swaggertype:"object"`
}

func NewPayoutAttemptsResponse(payoutAttempts model.PayoutAttempts) PayoutAttemptsResponse {
	return PayoutAttemptsResponse{
		Id:                payoutAttempts.Id,
		PayoutId:          payoutAttempts.PayoutId,
		AttemptNo:         payoutAttempts.AttemptNo,
		AttemptType:       model.PayoutAttemptsAttemptType(payoutAttempts.AttemptType),
		ProviderAccountId: payoutAttempts.ProviderAccountId.UUID,
		Amount:            payoutAttempts.Amount,
		CurrencyCode:      payoutAttempts.CurrencyCode,
		AttemptStatus:     model.PayoutAttemptsAttemptStatus(payoutAttempts.AttemptStatus),
		ProviderPayoutRef: payoutAttempts.ProviderPayoutRef.String,
		FailureCode:       payoutAttempts.FailureCode.String,
		FailureReason:     payoutAttempts.FailureReason.String,
		RawRequest:        payoutAttempts.RawRequest,
		RawResponse:       payoutAttempts.RawResponse,
		Metadata:          payoutAttempts.Metadata,
	}
}

type PayoutAttemptsListResponse []*PayoutAttemptsResponse

func NewPayoutAttemptsListResponse(payoutAttemptsList model.PayoutAttemptsList) PayoutAttemptsListResponse {
	dtoPayoutAttemptsListResponse := PayoutAttemptsListResponse{}
	for _, payoutAttempts := range payoutAttemptsList {
		dtoPayoutAttemptsResponse := NewPayoutAttemptsResponse(*payoutAttempts)
		dtoPayoutAttemptsListResponse = append(dtoPayoutAttemptsListResponse, &dtoPayoutAttemptsResponse)
	}
	return dtoPayoutAttemptsListResponse
}

type PayoutAttemptsPrimaryIDList []PayoutAttemptsPrimaryID

func (d PayoutAttemptsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutAttempts := range d {
		err = validator.Struct(payoutAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutAttemptsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutAttemptsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutAttemptsPrimaryID) ToModel() model.PayoutAttemptsPrimaryID {
	return model.PayoutAttemptsPrimaryID{
		Id: d.Id,
	}
}

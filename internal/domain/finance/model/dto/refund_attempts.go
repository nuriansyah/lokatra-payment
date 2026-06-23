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

type RefundAttemptsDTOFieldNameType string

type refundAttemptsDTOFieldName struct {
	Id                RefundAttemptsDTOFieldNameType
	RefundId          RefundAttemptsDTOFieldNameType
	AttemptNo         RefundAttemptsDTOFieldNameType
	AttemptType       RefundAttemptsDTOFieldNameType
	ProviderAccountId RefundAttemptsDTOFieldNameType
	Amount            RefundAttemptsDTOFieldNameType
	CurrencyCode      RefundAttemptsDTOFieldNameType
	AttemptStatus     RefundAttemptsDTOFieldNameType
	ProviderRefundRef RefundAttemptsDTOFieldNameType
	FailureCode       RefundAttemptsDTOFieldNameType
	FailureReason     RefundAttemptsDTOFieldNameType
	RawRequest        RefundAttemptsDTOFieldNameType
	RawResponse       RefundAttemptsDTOFieldNameType
	Metadata          RefundAttemptsDTOFieldNameType
	MetaCreatedAt     RefundAttemptsDTOFieldNameType
	MetaCreatedBy     RefundAttemptsDTOFieldNameType
	MetaUpdatedAt     RefundAttemptsDTOFieldNameType
	MetaUpdatedBy     RefundAttemptsDTOFieldNameType
	MetaDeletedAt     RefundAttemptsDTOFieldNameType
	MetaDeletedBy     RefundAttemptsDTOFieldNameType
}

var RefundAttemptsDTOFieldName = refundAttemptsDTOFieldName{
	Id:                "id",
	RefundId:          "refundId",
	AttemptNo:         "attemptNo",
	AttemptType:       "attemptType",
	ProviderAccountId: "providerAccountId",
	Amount:            "amount",
	CurrencyCode:      "currencyCode",
	AttemptStatus:     "attemptStatus",
	ProviderRefundRef: "providerRefundRef",
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

func transformRefundAttemptsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundAttemptsDTOFieldName.Id):
		return string(model.RefundAttemptsDBFieldName.Id), true

	case string(RefundAttemptsDTOFieldName.RefundId):
		return string(model.RefundAttemptsDBFieldName.RefundId), true

	case string(RefundAttemptsDTOFieldName.AttemptNo):
		return string(model.RefundAttemptsDBFieldName.AttemptNo), true

	case string(RefundAttemptsDTOFieldName.AttemptType):
		return string(model.RefundAttemptsDBFieldName.AttemptType), true

	case string(RefundAttemptsDTOFieldName.ProviderAccountId):
		return string(model.RefundAttemptsDBFieldName.ProviderAccountId), true

	case string(RefundAttemptsDTOFieldName.Amount):
		return string(model.RefundAttemptsDBFieldName.Amount), true

	case string(RefundAttemptsDTOFieldName.CurrencyCode):
		return string(model.RefundAttemptsDBFieldName.CurrencyCode), true

	case string(RefundAttemptsDTOFieldName.AttemptStatus):
		return string(model.RefundAttemptsDBFieldName.AttemptStatus), true

	case string(RefundAttemptsDTOFieldName.ProviderRefundRef):
		return string(model.RefundAttemptsDBFieldName.ProviderRefundRef), true

	case string(RefundAttemptsDTOFieldName.FailureCode):
		return string(model.RefundAttemptsDBFieldName.FailureCode), true

	case string(RefundAttemptsDTOFieldName.FailureReason):
		return string(model.RefundAttemptsDBFieldName.FailureReason), true

	case string(RefundAttemptsDTOFieldName.RawRequest):
		return string(model.RefundAttemptsDBFieldName.RawRequest), true

	case string(RefundAttemptsDTOFieldName.RawResponse):
		return string(model.RefundAttemptsDBFieldName.RawResponse), true

	case string(RefundAttemptsDTOFieldName.Metadata):
		return string(model.RefundAttemptsDBFieldName.Metadata), true

	case string(RefundAttemptsDTOFieldName.MetaCreatedAt):
		return string(model.RefundAttemptsDBFieldName.MetaCreatedAt), true

	case string(RefundAttemptsDTOFieldName.MetaCreatedBy):
		return string(model.RefundAttemptsDBFieldName.MetaCreatedBy), true

	case string(RefundAttemptsDTOFieldName.MetaUpdatedAt):
		return string(model.RefundAttemptsDBFieldName.MetaUpdatedAt), true

	case string(RefundAttemptsDTOFieldName.MetaUpdatedBy):
		return string(model.RefundAttemptsDBFieldName.MetaUpdatedBy), true

	case string(RefundAttemptsDTOFieldName.MetaDeletedAt):
		return string(model.RefundAttemptsDBFieldName.MetaDeletedAt), true

	case string(RefundAttemptsDTOFieldName.MetaDeletedBy):
		return string(model.RefundAttemptsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundAttemptsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundAttemptsBaseFilterField(field string) bool {
	spec, found := model.NewRefundAttemptsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundAttemptsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundAttemptsProjectionOutputPath(path string) error {
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

func transformRefundAttemptsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundAttemptsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundAttemptsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundAttemptsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundAttemptsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundAttemptsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundAttemptsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundAttemptsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundAttemptsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundAttemptsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundAttemptsFilter(filter *model.Filter) {
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
			Field: string(RefundAttemptsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundAttemptsSelectableResponse map[string]interface{}
type RefundAttemptsSelectableListResponse []*RefundAttemptsSelectableResponse

func assignRefundAttemptsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundAttemptsSelectableValue(out RefundAttemptsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundAttemptsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundAttemptsSelectableResponse(refundAttempts model.RefundAttempts, filter model.Filter) RefundAttemptsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundAttemptsDBFieldName.Id),
			string(model.RefundAttemptsDBFieldName.RefundId),
			string(model.RefundAttemptsDBFieldName.AttemptNo),
			string(model.RefundAttemptsDBFieldName.AttemptType),
			string(model.RefundAttemptsDBFieldName.ProviderAccountId),
			string(model.RefundAttemptsDBFieldName.Amount),
			string(model.RefundAttemptsDBFieldName.CurrencyCode),
			string(model.RefundAttemptsDBFieldName.AttemptStatus),
			string(model.RefundAttemptsDBFieldName.ProviderRefundRef),
			string(model.RefundAttemptsDBFieldName.FailureCode),
			string(model.RefundAttemptsDBFieldName.FailureReason),
			string(model.RefundAttemptsDBFieldName.RawRequest),
			string(model.RefundAttemptsDBFieldName.RawResponse),
			string(model.RefundAttemptsDBFieldName.Metadata),
			string(model.RefundAttemptsDBFieldName.MetaCreatedAt),
			string(model.RefundAttemptsDBFieldName.MetaCreatedBy),
			string(model.RefundAttemptsDBFieldName.MetaUpdatedAt),
			string(model.RefundAttemptsDBFieldName.MetaUpdatedBy),
			string(model.RefundAttemptsDBFieldName.MetaDeletedAt),
			string(model.RefundAttemptsDBFieldName.MetaDeletedBy),
		)
	}
	refundAttemptsSelectableResponse := RefundAttemptsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundAttemptsDBFieldName.Id):
			key := string(RefundAttemptsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.Id, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.RefundId):
			key := string(RefundAttemptsDTOFieldName.RefundId)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.RefundId, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.AttemptNo):
			key := string(RefundAttemptsDTOFieldName.AttemptNo)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.AttemptNo, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.AttemptType):
			key := string(RefundAttemptsDTOFieldName.AttemptType)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, model.RefundAttemptsAttemptType(refundAttempts.AttemptType), explicitAlias)

		case string(model.RefundAttemptsDBFieldName.ProviderAccountId):
			key := string(RefundAttemptsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.ProviderAccountId.UUID, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.Amount):
			key := string(RefundAttemptsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.Amount, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.CurrencyCode):
			key := string(RefundAttemptsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.CurrencyCode, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.AttemptStatus):
			key := string(RefundAttemptsDTOFieldName.AttemptStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, model.RefundAttemptsAttemptStatus(refundAttempts.AttemptStatus), explicitAlias)

		case string(model.RefundAttemptsDBFieldName.ProviderRefundRef):
			key := string(RefundAttemptsDTOFieldName.ProviderRefundRef)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.ProviderRefundRef.String, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.FailureCode):
			key := string(RefundAttemptsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.FailureCode.String, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.FailureReason):
			key := string(RefundAttemptsDTOFieldName.FailureReason)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.FailureReason.String, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.RawRequest):
			key := string(RefundAttemptsDTOFieldName.RawRequest)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.RawRequest, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.RawResponse):
			key := string(RefundAttemptsDTOFieldName.RawResponse)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.RawResponse, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.Metadata):
			key := string(RefundAttemptsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.Metadata, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaCreatedAt):
			key := string(RefundAttemptsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaCreatedAt, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaCreatedBy):
			key := string(RefundAttemptsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaCreatedBy, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaUpdatedAt):
			key := string(RefundAttemptsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaUpdatedAt, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaUpdatedBy):
			key := string(RefundAttemptsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaUpdatedBy, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaDeletedAt):
			key := string(RefundAttemptsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundAttemptsDBFieldName.MetaDeletedBy):
			key := string(RefundAttemptsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAttemptsSelectableValue(refundAttemptsSelectableResponse, key, refundAttempts.MetaDeletedBy, explicitAlias)

		}
	}
	return refundAttemptsSelectableResponse
}

func NewRefundAttemptsListResponseFromFilterResult(result []model.RefundAttemptsFilterResult, filter model.Filter) RefundAttemptsSelectableListResponse {
	dtoRefundAttemptsListResponse := RefundAttemptsSelectableListResponse{}
	for _, row := range result {
		dtoRefundAttemptsResponse := NewRefundAttemptsSelectableResponse(row.RefundAttempts, filter)
		dtoRefundAttemptsListResponse = append(dtoRefundAttemptsListResponse, &dtoRefundAttemptsResponse)
	}
	return dtoRefundAttemptsListResponse
}

type RefundAttemptsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     RefundAttemptsSelectableListResponse `json:"data"`
}

func reverseRefundAttemptsFilterResults(result []model.RefundAttemptsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundAttemptsFilterResponse(result []model.RefundAttemptsFilterResult, filter model.Filter) (resp RefundAttemptsFilterResponse) {
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
			reverseRefundAttemptsFilterResults(dataResult)
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

	resp.Data = NewRefundAttemptsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundAttemptsCreateRequest struct {
	RefundId          uuid.UUID                         `json:"refundId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.RefundAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,wallet_credit,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.RefundAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderRefundRef string                            `json:"providerRefundRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d *RefundAttemptsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundAttemptsCreateRequest) ToModel() model.RefundAttempts {
	id, _ := uuid.NewV4()
	return model.RefundAttempts{
		Id:                id,
		RefundId:          d.RefundId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderRefundRef: null.StringFrom(d.ProviderRefundRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type RefundAttemptsListCreateRequest []*RefundAttemptsCreateRequest

func (d RefundAttemptsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAttempts := range d {
		err = validator.Struct(refundAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundAttemptsListCreateRequest) ToModelList() []model.RefundAttempts {
	out := make([]model.RefundAttempts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundAttemptsUpdateRequest struct {
	RefundId          uuid.UUID                         `json:"refundId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.RefundAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,wallet_credit,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.RefundAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderRefundRef string                            `json:"providerRefundRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d *RefundAttemptsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundAttemptsUpdateRequest) ToModel() model.RefundAttempts {
	return model.RefundAttempts{
		RefundId:          d.RefundId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderRefundRef: null.StringFrom(d.ProviderRefundRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type RefundAttemptsBulkUpdateRequest struct {
	Id                uuid.UUID                         `json:"id"`
	RefundId          uuid.UUID                         `json:"refundId"`
	AttemptNo         int                               `json:"attemptNo"`
	AttemptType       model.RefundAttemptsAttemptType   `json:"attemptType" example:"provider" enums:"provider,manual,wallet_credit,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId"`
	Amount            decimal.Decimal                   `json:"amount"`
	CurrencyCode      string                            `json:"currencyCode"`
	AttemptStatus     model.RefundAttemptsAttemptStatus `json:"attemptStatus" example:"created" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderRefundRef string                            `json:"providerRefundRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest"`
	RawResponse       json.RawMessage                   `json:"rawResponse"`
	Metadata          json.RawMessage                   `json:"metadata"`
}

func (d RefundAttemptsBulkUpdateRequest) PrimaryID() RefundAttemptsPrimaryID {
	return RefundAttemptsPrimaryID{
		Id: d.Id,
	}
}

type RefundAttemptsListBulkUpdateRequest []*RefundAttemptsBulkUpdateRequest

func (d RefundAttemptsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAttempts := range d {
		err = validator.Struct(refundAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundAttemptsBulkUpdateRequest) ToModel() model.RefundAttempts {
	return model.RefundAttempts{
		Id:                d.Id,
		RefundId:          d.RefundId,
		AttemptNo:         d.AttemptNo,
		AttemptType:       d.AttemptType,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		Amount:            d.Amount,
		CurrencyCode:      d.CurrencyCode,
		AttemptStatus:     d.AttemptStatus,
		ProviderRefundRef: null.StringFrom(d.ProviderRefundRef),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		RawRequest:        d.RawRequest,
		RawResponse:       d.RawResponse,
		Metadata:          d.Metadata,
	}
}

type RefundAttemptsResponse struct {
	Id                uuid.UUID                         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundId          uuid.UUID                         `json:"refundId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AttemptNo         int                               `json:"attemptNo" validate:"required" example:"1"`
	AttemptType       model.RefundAttemptsAttemptType   `json:"attemptType" validate:"required,oneof=provider manual wallet_credit bank_transfer" enums:"provider,manual,wallet_credit,bank_transfer"`
	ProviderAccountId uuid.UUID                         `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount            decimal.Decimal                   `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	CurrencyCode      string                            `json:"currencyCode" validate:"required"`
	AttemptStatus     model.RefundAttemptsAttemptStatus `json:"attemptStatus" validate:"required,oneof=created processing succeeded failed cancelled" enums:"created,processing,succeeded,failed,cancelled"`
	ProviderRefundRef string                            `json:"providerRefundRef"`
	FailureCode       string                            `json:"failureCode"`
	FailureReason     string                            `json:"failureReason"`
	RawRequest        json.RawMessage                   `json:"rawRequest" swaggertype:"object"`
	RawResponse       json.RawMessage                   `json:"rawResponse" swaggertype:"object"`
	Metadata          json.RawMessage                   `json:"metadata" swaggertype:"object"`
}

func NewRefundAttemptsResponse(refundAttempts model.RefundAttempts) RefundAttemptsResponse {
	return RefundAttemptsResponse{
		Id:                refundAttempts.Id,
		RefundId:          refundAttempts.RefundId,
		AttemptNo:         refundAttempts.AttemptNo,
		AttemptType:       model.RefundAttemptsAttemptType(refundAttempts.AttemptType),
		ProviderAccountId: refundAttempts.ProviderAccountId.UUID,
		Amount:            refundAttempts.Amount,
		CurrencyCode:      refundAttempts.CurrencyCode,
		AttemptStatus:     model.RefundAttemptsAttemptStatus(refundAttempts.AttemptStatus),
		ProviderRefundRef: refundAttempts.ProviderRefundRef.String,
		FailureCode:       refundAttempts.FailureCode.String,
		FailureReason:     refundAttempts.FailureReason.String,
		RawRequest:        refundAttempts.RawRequest,
		RawResponse:       refundAttempts.RawResponse,
		Metadata:          refundAttempts.Metadata,
	}
}

type RefundAttemptsListResponse []*RefundAttemptsResponse

func NewRefundAttemptsListResponse(refundAttemptsList model.RefundAttemptsList) RefundAttemptsListResponse {
	dtoRefundAttemptsListResponse := RefundAttemptsListResponse{}
	for _, refundAttempts := range refundAttemptsList {
		dtoRefundAttemptsResponse := NewRefundAttemptsResponse(*refundAttempts)
		dtoRefundAttemptsListResponse = append(dtoRefundAttemptsListResponse, &dtoRefundAttemptsResponse)
	}
	return dtoRefundAttemptsListResponse
}

type RefundAttemptsPrimaryIDList []RefundAttemptsPrimaryID

func (d RefundAttemptsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAttempts := range d {
		err = validator.Struct(refundAttempts)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundAttemptsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundAttemptsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundAttemptsPrimaryID) ToModel() model.RefundAttemptsPrimaryID {
	return model.RefundAttemptsPrimaryID{
		Id: d.Id,
	}
}

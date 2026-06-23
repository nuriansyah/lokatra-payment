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

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RefundRequestsDTOFieldNameType string

type refundRequestsDTOFieldName struct {
	Id                   RefundRequestsDTOFieldNameType
	RefundCode           RefundRequestsDTOFieldNameType
	PaymentRefId         RefundRequestsDTOFieldNameType
	MerchantPartyId      RefundRequestsDTOFieldNameType
	CustomerPartyId      RefundRequestsDTOFieldNameType
	RefundPolicyId       RefundRequestsDTOFieldNameType
	CurrencyCode         RefundRequestsDTOFieldNameType
	RequestedAmount      RefundRequestsDTOFieldNameType
	ApprovedAmount       RefundRequestsDTOFieldNameType
	RefundReasonCode     RefundRequestsDTOFieldNameType
	RefundStatus         RefundRequestsDTOFieldNameType
	RequestedAt          RefundRequestsDTOFieldNameType
	ApprovedAt           RefundRequestsDTOFieldNameType
	SettledFinanciallyAt RefundRequestsDTOFieldNameType
	Metadata             RefundRequestsDTOFieldNameType
	MetaCreatedAt        RefundRequestsDTOFieldNameType
	MetaCreatedBy        RefundRequestsDTOFieldNameType
	MetaUpdatedAt        RefundRequestsDTOFieldNameType
	MetaUpdatedBy        RefundRequestsDTOFieldNameType
	MetaDeletedAt        RefundRequestsDTOFieldNameType
	MetaDeletedBy        RefundRequestsDTOFieldNameType
}

var RefundRequestsDTOFieldName = refundRequestsDTOFieldName{
	Id:                   "id",
	RefundCode:           "refundCode",
	PaymentRefId:         "paymentRefId",
	MerchantPartyId:      "merchantPartyId",
	CustomerPartyId:      "customerPartyId",
	RefundPolicyId:       "refundPolicyId",
	CurrencyCode:         "currencyCode",
	RequestedAmount:      "requestedAmount",
	ApprovedAmount:       "approvedAmount",
	RefundReasonCode:     "refundReasonCode",
	RefundStatus:         "refundStatus",
	RequestedAt:          "requestedAt",
	ApprovedAt:           "approvedAt",
	SettledFinanciallyAt: "settledFinanciallyAt",
	Metadata:             "metadata",
	MetaCreatedAt:        "metaCreatedAt",
	MetaCreatedBy:        "metaCreatedBy",
	MetaUpdatedAt:        "metaUpdatedAt",
	MetaUpdatedBy:        "metaUpdatedBy",
	MetaDeletedAt:        "metaDeletedAt",
	MetaDeletedBy:        "metaDeletedBy",
}

func transformRefundRequestsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundRequestsDTOFieldName.Id):
		return string(model.RefundRequestsDBFieldName.Id), true

	case string(RefundRequestsDTOFieldName.RefundCode):
		return string(model.RefundRequestsDBFieldName.RefundCode), true

	case string(RefundRequestsDTOFieldName.PaymentRefId):
		return string(model.RefundRequestsDBFieldName.PaymentRefId), true

	case string(RefundRequestsDTOFieldName.MerchantPartyId):
		return string(model.RefundRequestsDBFieldName.MerchantPartyId), true

	case string(RefundRequestsDTOFieldName.CustomerPartyId):
		return string(model.RefundRequestsDBFieldName.CustomerPartyId), true

	case string(RefundRequestsDTOFieldName.RefundPolicyId):
		return string(model.RefundRequestsDBFieldName.RefundPolicyId), true

	case string(RefundRequestsDTOFieldName.CurrencyCode):
		return string(model.RefundRequestsDBFieldName.CurrencyCode), true

	case string(RefundRequestsDTOFieldName.RequestedAmount):
		return string(model.RefundRequestsDBFieldName.RequestedAmount), true

	case string(RefundRequestsDTOFieldName.ApprovedAmount):
		return string(model.RefundRequestsDBFieldName.ApprovedAmount), true

	case string(RefundRequestsDTOFieldName.RefundReasonCode):
		return string(model.RefundRequestsDBFieldName.RefundReasonCode), true

	case string(RefundRequestsDTOFieldName.RefundStatus):
		return string(model.RefundRequestsDBFieldName.RefundStatus), true

	case string(RefundRequestsDTOFieldName.RequestedAt):
		return string(model.RefundRequestsDBFieldName.RequestedAt), true

	case string(RefundRequestsDTOFieldName.ApprovedAt):
		return string(model.RefundRequestsDBFieldName.ApprovedAt), true

	case string(RefundRequestsDTOFieldName.SettledFinanciallyAt):
		return string(model.RefundRequestsDBFieldName.SettledFinanciallyAt), true

	case string(RefundRequestsDTOFieldName.Metadata):
		return string(model.RefundRequestsDBFieldName.Metadata), true

	case string(RefundRequestsDTOFieldName.MetaCreatedAt):
		return string(model.RefundRequestsDBFieldName.MetaCreatedAt), true

	case string(RefundRequestsDTOFieldName.MetaCreatedBy):
		return string(model.RefundRequestsDBFieldName.MetaCreatedBy), true

	case string(RefundRequestsDTOFieldName.MetaUpdatedAt):
		return string(model.RefundRequestsDBFieldName.MetaUpdatedAt), true

	case string(RefundRequestsDTOFieldName.MetaUpdatedBy):
		return string(model.RefundRequestsDBFieldName.MetaUpdatedBy), true

	case string(RefundRequestsDTOFieldName.MetaDeletedAt):
		return string(model.RefundRequestsDBFieldName.MetaDeletedAt), true

	case string(RefundRequestsDTOFieldName.MetaDeletedBy):
		return string(model.RefundRequestsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundRequestsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundRequestsBaseFilterField(field string) bool {
	spec, found := model.NewRefundRequestsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundRequestsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundRequestsProjectionOutputPath(path string) error {
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

func transformRefundRequestsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundRequestsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundRequestsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundRequestsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundRequestsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundRequestsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundRequestsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundRequestsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundRequestsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundRequestsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundRequestsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundRequestsFilter(filter *model.Filter) {
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
			Field: string(RefundRequestsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundRequestsSelectableResponse map[string]interface{}
type RefundRequestsSelectableListResponse []*RefundRequestsSelectableResponse

func assignRefundRequestsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundRequestsSelectableValue(out RefundRequestsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundRequestsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundRequestsSelectableResponse(refundRequests model.RefundRequests, filter model.Filter) RefundRequestsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundRequestsDBFieldName.Id),
			string(model.RefundRequestsDBFieldName.RefundCode),
			string(model.RefundRequestsDBFieldName.PaymentRefId),
			string(model.RefundRequestsDBFieldName.MerchantPartyId),
			string(model.RefundRequestsDBFieldName.CustomerPartyId),
			string(model.RefundRequestsDBFieldName.RefundPolicyId),
			string(model.RefundRequestsDBFieldName.CurrencyCode),
			string(model.RefundRequestsDBFieldName.RequestedAmount),
			string(model.RefundRequestsDBFieldName.ApprovedAmount),
			string(model.RefundRequestsDBFieldName.RefundReasonCode),
			string(model.RefundRequestsDBFieldName.RefundStatus),
			string(model.RefundRequestsDBFieldName.RequestedAt),
			string(model.RefundRequestsDBFieldName.ApprovedAt),
			string(model.RefundRequestsDBFieldName.SettledFinanciallyAt),
			string(model.RefundRequestsDBFieldName.Metadata),
			string(model.RefundRequestsDBFieldName.MetaCreatedAt),
			string(model.RefundRequestsDBFieldName.MetaCreatedBy),
			string(model.RefundRequestsDBFieldName.MetaUpdatedAt),
			string(model.RefundRequestsDBFieldName.MetaUpdatedBy),
			string(model.RefundRequestsDBFieldName.MetaDeletedAt),
			string(model.RefundRequestsDBFieldName.MetaDeletedBy),
		)
	}
	refundRequestsSelectableResponse := RefundRequestsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundRequestsDBFieldName.Id):
			key := string(RefundRequestsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.Id, explicitAlias)

		case string(model.RefundRequestsDBFieldName.RefundCode):
			key := string(RefundRequestsDTOFieldName.RefundCode)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.RefundCode, explicitAlias)

		case string(model.RefundRequestsDBFieldName.PaymentRefId):
			key := string(RefundRequestsDTOFieldName.PaymentRefId)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.PaymentRefId, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MerchantPartyId):
			key := string(RefundRequestsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MerchantPartyId.UUID, explicitAlias)

		case string(model.RefundRequestsDBFieldName.CustomerPartyId):
			key := string(RefundRequestsDTOFieldName.CustomerPartyId)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.CustomerPartyId.UUID, explicitAlias)

		case string(model.RefundRequestsDBFieldName.RefundPolicyId):
			key := string(RefundRequestsDTOFieldName.RefundPolicyId)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.RefundPolicyId.UUID, explicitAlias)

		case string(model.RefundRequestsDBFieldName.CurrencyCode):
			key := string(RefundRequestsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.CurrencyCode, explicitAlias)

		case string(model.RefundRequestsDBFieldName.RequestedAmount):
			key := string(RefundRequestsDTOFieldName.RequestedAmount)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.RequestedAmount, explicitAlias)

		case string(model.RefundRequestsDBFieldName.ApprovedAmount):
			key := string(RefundRequestsDTOFieldName.ApprovedAmount)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.ApprovedAmount.Decimal, explicitAlias)

		case string(model.RefundRequestsDBFieldName.RefundReasonCode):
			key := string(RefundRequestsDTOFieldName.RefundReasonCode)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.RefundReasonCode, explicitAlias)

		case string(model.RefundRequestsDBFieldName.RefundStatus):
			key := string(RefundRequestsDTOFieldName.RefundStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, model.RefundStatus(refundRequests.RefundStatus), explicitAlias)

		case string(model.RefundRequestsDBFieldName.RequestedAt):
			key := string(RefundRequestsDTOFieldName.RequestedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.RequestedAt, explicitAlias)

		case string(model.RefundRequestsDBFieldName.ApprovedAt):
			key := string(RefundRequestsDTOFieldName.ApprovedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.ApprovedAt.Time, explicitAlias)

		case string(model.RefundRequestsDBFieldName.SettledFinanciallyAt):
			key := string(RefundRequestsDTOFieldName.SettledFinanciallyAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.SettledFinanciallyAt.Time, explicitAlias)

		case string(model.RefundRequestsDBFieldName.Metadata):
			key := string(RefundRequestsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.Metadata, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaCreatedAt):
			key := string(RefundRequestsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaCreatedAt, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaCreatedBy):
			key := string(RefundRequestsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaCreatedBy, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaUpdatedAt):
			key := string(RefundRequestsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaUpdatedAt, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaUpdatedBy):
			key := string(RefundRequestsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaUpdatedBy, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaDeletedAt):
			key := string(RefundRequestsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundRequestsDBFieldName.MetaDeletedBy):
			key := string(RefundRequestsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundRequestsSelectableValue(refundRequestsSelectableResponse, key, refundRequests.MetaDeletedBy, explicitAlias)

		}
	}
	return refundRequestsSelectableResponse
}

func NewRefundRequestsListResponseFromFilterResult(result []model.RefundRequestsFilterResult, filter model.Filter) RefundRequestsSelectableListResponse {
	dtoRefundRequestsListResponse := RefundRequestsSelectableListResponse{}
	for _, row := range result {
		dtoRefundRequestsResponse := NewRefundRequestsSelectableResponse(row.RefundRequests, filter)
		dtoRefundRequestsListResponse = append(dtoRefundRequestsListResponse, &dtoRefundRequestsResponse)
	}
	return dtoRefundRequestsListResponse
}

type RefundRequestsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     RefundRequestsSelectableListResponse `json:"data"`
}

func reverseRefundRequestsFilterResults(result []model.RefundRequestsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundRequestsFilterResponse(result []model.RefundRequestsFilterResult, filter model.Filter) (resp RefundRequestsFilterResponse) {
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
			reverseRefundRequestsFilterResults(dataResult)
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

	resp.Data = NewRefundRequestsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundRequestsCreateRequest struct {
	RefundCode           string             `json:"refundCode"`
	PaymentRefId         uuid.UUID          `json:"paymentRefId"`
	MerchantPartyId      uuid.UUID          `json:"merchantPartyId"`
	CustomerPartyId      uuid.UUID          `json:"customerPartyId"`
	RefundPolicyId       uuid.UUID          `json:"refundPolicyId"`
	CurrencyCode         string             `json:"currencyCode"`
	RequestedAmount      decimal.Decimal    `json:"requestedAmount"`
	ApprovedAmount       decimal.Decimal    `json:"approvedAmount"`
	RefundReasonCode     string             `json:"refundReasonCode"`
	RefundStatus         model.RefundStatus `json:"refundStatus" example:"requested" enums:"requested,approved,processing,succeeded,failed,rejected,cancelled"`
	RequestedAt          time.Time          `json:"requestedAt"`
	ApprovedAt           time.Time          `json:"approvedAt"`
	SettledFinanciallyAt time.Time          `json:"settledFinanciallyAt"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d *RefundRequestsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundRequestsCreateRequest) ToModel() model.RefundRequests {
	id, _ := uuid.NewV4()
	return model.RefundRequests{
		Id:                   id,
		RefundCode:           d.RefundCode,
		PaymentRefId:         d.PaymentRefId,
		MerchantPartyId:      nuuid.From(d.MerchantPartyId),
		CustomerPartyId:      nuuid.From(d.CustomerPartyId),
		RefundPolicyId:       nuuid.From(d.RefundPolicyId),
		CurrencyCode:         d.CurrencyCode,
		RequestedAmount:      d.RequestedAmount,
		ApprovedAmount:       decimal.NewNullDecimal(d.ApprovedAmount),
		RefundReasonCode:     d.RefundReasonCode,
		RefundStatus:         d.RefundStatus,
		RequestedAt:          d.RequestedAt,
		ApprovedAt:           null.TimeFrom(d.ApprovedAt),
		SettledFinanciallyAt: null.TimeFrom(d.SettledFinanciallyAt),
		Metadata:             d.Metadata,
	}
}

type RefundRequestsListCreateRequest []*RefundRequestsCreateRequest

func (d RefundRequestsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundRequests := range d {
		err = validator.Struct(refundRequests)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundRequestsListCreateRequest) ToModelList() []model.RefundRequests {
	out := make([]model.RefundRequests, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundRequestsUpdateRequest struct {
	RefundCode           string             `json:"refundCode"`
	PaymentRefId         uuid.UUID          `json:"paymentRefId"`
	MerchantPartyId      uuid.UUID          `json:"merchantPartyId"`
	CustomerPartyId      uuid.UUID          `json:"customerPartyId"`
	RefundPolicyId       uuid.UUID          `json:"refundPolicyId"`
	CurrencyCode         string             `json:"currencyCode"`
	RequestedAmount      decimal.Decimal    `json:"requestedAmount"`
	ApprovedAmount       decimal.Decimal    `json:"approvedAmount"`
	RefundReasonCode     string             `json:"refundReasonCode"`
	RefundStatus         model.RefundStatus `json:"refundStatus" example:"requested" enums:"requested,approved,processing,succeeded,failed,rejected,cancelled"`
	RequestedAt          time.Time          `json:"requestedAt"`
	ApprovedAt           time.Time          `json:"approvedAt"`
	SettledFinanciallyAt time.Time          `json:"settledFinanciallyAt"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d *RefundRequestsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundRequestsUpdateRequest) ToModel() model.RefundRequests {
	return model.RefundRequests{
		RefundCode:           d.RefundCode,
		PaymentRefId:         d.PaymentRefId,
		MerchantPartyId:      nuuid.From(d.MerchantPartyId),
		CustomerPartyId:      nuuid.From(d.CustomerPartyId),
		RefundPolicyId:       nuuid.From(d.RefundPolicyId),
		CurrencyCode:         d.CurrencyCode,
		RequestedAmount:      d.RequestedAmount,
		ApprovedAmount:       decimal.NewNullDecimal(d.ApprovedAmount),
		RefundReasonCode:     d.RefundReasonCode,
		RefundStatus:         d.RefundStatus,
		RequestedAt:          d.RequestedAt,
		ApprovedAt:           null.TimeFrom(d.ApprovedAt),
		SettledFinanciallyAt: null.TimeFrom(d.SettledFinanciallyAt),
		Metadata:             d.Metadata,
	}
}

type RefundRequestsBulkUpdateRequest struct {
	Id                   uuid.UUID          `json:"id"`
	RefundCode           string             `json:"refundCode"`
	PaymentRefId         uuid.UUID          `json:"paymentRefId"`
	MerchantPartyId      uuid.UUID          `json:"merchantPartyId"`
	CustomerPartyId      uuid.UUID          `json:"customerPartyId"`
	RefundPolicyId       uuid.UUID          `json:"refundPolicyId"`
	CurrencyCode         string             `json:"currencyCode"`
	RequestedAmount      decimal.Decimal    `json:"requestedAmount"`
	ApprovedAmount       decimal.Decimal    `json:"approvedAmount"`
	RefundReasonCode     string             `json:"refundReasonCode"`
	RefundStatus         model.RefundStatus `json:"refundStatus" example:"requested" enums:"requested,approved,processing,succeeded,failed,rejected,cancelled"`
	RequestedAt          time.Time          `json:"requestedAt"`
	ApprovedAt           time.Time          `json:"approvedAt"`
	SettledFinanciallyAt time.Time          `json:"settledFinanciallyAt"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d RefundRequestsBulkUpdateRequest) PrimaryID() RefundRequestsPrimaryID {
	return RefundRequestsPrimaryID{
		Id: d.Id,
	}
}

type RefundRequestsListBulkUpdateRequest []*RefundRequestsBulkUpdateRequest

func (d RefundRequestsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundRequests := range d {
		err = validator.Struct(refundRequests)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundRequestsBulkUpdateRequest) ToModel() model.RefundRequests {
	return model.RefundRequests{
		Id:                   d.Id,
		RefundCode:           d.RefundCode,
		PaymentRefId:         d.PaymentRefId,
		MerchantPartyId:      nuuid.From(d.MerchantPartyId),
		CustomerPartyId:      nuuid.From(d.CustomerPartyId),
		RefundPolicyId:       nuuid.From(d.RefundPolicyId),
		CurrencyCode:         d.CurrencyCode,
		RequestedAmount:      d.RequestedAmount,
		ApprovedAmount:       decimal.NewNullDecimal(d.ApprovedAmount),
		RefundReasonCode:     d.RefundReasonCode,
		RefundStatus:         d.RefundStatus,
		RequestedAt:          d.RequestedAt,
		ApprovedAt:           null.TimeFrom(d.ApprovedAt),
		SettledFinanciallyAt: null.TimeFrom(d.SettledFinanciallyAt),
		Metadata:             d.Metadata,
	}
}

type RefundRequestsResponse struct {
	Id                   uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundCode           string             `json:"refundCode" validate:"required"`
	PaymentRefId         uuid.UUID          `json:"paymentRefId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId      uuid.UUID          `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CustomerPartyId      uuid.UUID          `json:"customerPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundPolicyId       uuid.UUID          `json:"refundPolicyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode         string             `json:"currencyCode" validate:"required"`
	RequestedAmount      decimal.Decimal    `json:"requestedAmount" validate:"required" format:"decimal" example:"100.50"`
	ApprovedAmount       decimal.Decimal    `json:"approvedAmount" format:"decimal" example:"100.50"`
	RefundReasonCode     string             `json:"refundReasonCode" validate:"required"`
	RefundStatus         model.RefundStatus `json:"refundStatus" validate:"oneof=requested approved processing succeeded failed rejected cancelled" enums:"requested,approved,processing,succeeded,failed,rejected,cancelled"`
	RequestedAt          time.Time          `json:"requestedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ApprovedAt           time.Time          `json:"approvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	SettledFinanciallyAt time.Time          `json:"settledFinanciallyAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata             json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewRefundRequestsResponse(refundRequests model.RefundRequests) RefundRequestsResponse {
	return RefundRequestsResponse{
		Id:                   refundRequests.Id,
		RefundCode:           refundRequests.RefundCode,
		PaymentRefId:         refundRequests.PaymentRefId,
		MerchantPartyId:      refundRequests.MerchantPartyId.UUID,
		CustomerPartyId:      refundRequests.CustomerPartyId.UUID,
		RefundPolicyId:       refundRequests.RefundPolicyId.UUID,
		CurrencyCode:         refundRequests.CurrencyCode,
		RequestedAmount:      refundRequests.RequestedAmount,
		ApprovedAmount:       refundRequests.ApprovedAmount.Decimal,
		RefundReasonCode:     refundRequests.RefundReasonCode,
		RefundStatus:         model.RefundStatus(refundRequests.RefundStatus),
		RequestedAt:          refundRequests.RequestedAt,
		ApprovedAt:           refundRequests.ApprovedAt.Time,
		SettledFinanciallyAt: refundRequests.SettledFinanciallyAt.Time,
		Metadata:             refundRequests.Metadata,
	}
}

type RefundRequestsListResponse []*RefundRequestsResponse

func NewRefundRequestsListResponse(refundRequestsList model.RefundRequestsList) RefundRequestsListResponse {
	dtoRefundRequestsListResponse := RefundRequestsListResponse{}
	for _, refundRequests := range refundRequestsList {
		dtoRefundRequestsResponse := NewRefundRequestsResponse(*refundRequests)
		dtoRefundRequestsListResponse = append(dtoRefundRequestsListResponse, &dtoRefundRequestsResponse)
	}
	return dtoRefundRequestsListResponse
}

type RefundRequestsPrimaryIDList []RefundRequestsPrimaryID

func (d RefundRequestsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundRequests := range d {
		err = validator.Struct(refundRequests)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundRequestsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundRequestsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundRequestsPrimaryID) ToModel() model.RefundRequestsPrimaryID {
	return model.RefundRequestsPrimaryID{
		Id: d.Id,
	}
}

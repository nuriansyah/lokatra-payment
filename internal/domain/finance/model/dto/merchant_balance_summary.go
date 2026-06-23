package dto

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type MerchantBalanceSummaryDTOFieldNameType string

type merchantBalanceSummaryDTOFieldName struct {
	MerchantPartyId  MerchantBalanceSummaryDTOFieldNameType
	CurrencyCode     MerchantBalanceSummaryDTOFieldNameType
	PendingAmount    MerchantBalanceSummaryDTOFieldNameType
	AvailableAmount  MerchantBalanceSummaryDTOFieldNameType
	ReservedAmount   MerchantBalanceSummaryDTOFieldNameType
	PayableAmount    MerchantBalanceSummaryDTOFieldNameType
	PaidOutAmount    MerchantBalanceSummaryDTOFieldNameType
	NegativeAmount   MerchantBalanceSummaryDTOFieldNameType
	RefundableAmount MerchantBalanceSummaryDTOFieldNameType
	RefreshedAt      MerchantBalanceSummaryDTOFieldNameType
}

var MerchantBalanceSummaryDTOFieldName = merchantBalanceSummaryDTOFieldName{
	MerchantPartyId:  "merchantPartyId",
	CurrencyCode:     "currencyCode",
	PendingAmount:    "pendingAmount",
	AvailableAmount:  "availableAmount",
	ReservedAmount:   "reservedAmount",
	PayableAmount:    "payableAmount",
	PaidOutAmount:    "paidOutAmount",
	NegativeAmount:   "negativeAmount",
	RefundableAmount: "refundableAmount",
	RefreshedAt:      "refreshedAt",
}

func transformMerchantBalanceSummaryDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(MerchantBalanceSummaryDTOFieldName.MerchantPartyId):
		return string(model.MerchantBalanceSummaryDBFieldName.MerchantPartyId), true

	case string(MerchantBalanceSummaryDTOFieldName.CurrencyCode):
		return string(model.MerchantBalanceSummaryDBFieldName.CurrencyCode), true

	case string(MerchantBalanceSummaryDTOFieldName.PendingAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.PendingAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.AvailableAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.AvailableAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.ReservedAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.ReservedAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.PayableAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.PayableAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.PaidOutAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.PaidOutAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.NegativeAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.NegativeAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.RefundableAmount):
		return string(model.MerchantBalanceSummaryDBFieldName.RefundableAmount), true

	case string(MerchantBalanceSummaryDTOFieldName.RefreshedAt):
		return string(model.MerchantBalanceSummaryDBFieldName.RefreshedAt), true

	}
	if _, found := model.NewMerchantBalanceSummaryFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isMerchantBalanceSummaryBaseFilterField(field string) bool {
	spec, found := model.NewMerchantBalanceSummaryFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeMerchantBalanceSummaryProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateMerchantBalanceSummaryProjectionOutputPath(path string) error {
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

func transformMerchantBalanceSummaryFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformMerchantBalanceSummaryDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformMerchantBalanceSummaryFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformMerchantBalanceSummaryFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformMerchantBalanceSummaryDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isMerchantBalanceSummaryBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateMerchantBalanceSummaryProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeMerchantBalanceSummaryProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformMerchantBalanceSummaryDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformMerchantBalanceSummaryDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformMerchantBalanceSummaryFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultMerchantBalanceSummaryFilter(filter *model.Filter) {
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
			Field: string(MerchantBalanceSummaryDTOFieldName.MerchantPartyId),
			Order: model.SortAsc,
		})
	}
}

type MerchantBalanceSummarySelectableResponse map[string]interface{}
type MerchantBalanceSummarySelectableListResponse []*MerchantBalanceSummarySelectableResponse

func assignMerchantBalanceSummaryNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setMerchantBalanceSummarySelectableValue(out MerchantBalanceSummarySelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignMerchantBalanceSummaryNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewMerchantBalanceSummarySelectableResponse(merchantBalanceSummary model.MerchantBalanceSummary, filter model.Filter) MerchantBalanceSummarySelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.MerchantBalanceSummaryDBFieldName.MerchantPartyId),
			string(model.MerchantBalanceSummaryDBFieldName.CurrencyCode),
			string(model.MerchantBalanceSummaryDBFieldName.PendingAmount),
			string(model.MerchantBalanceSummaryDBFieldName.AvailableAmount),
			string(model.MerchantBalanceSummaryDBFieldName.ReservedAmount),
			string(model.MerchantBalanceSummaryDBFieldName.PayableAmount),
			string(model.MerchantBalanceSummaryDBFieldName.PaidOutAmount),
			string(model.MerchantBalanceSummaryDBFieldName.NegativeAmount),
			string(model.MerchantBalanceSummaryDBFieldName.RefundableAmount),
			string(model.MerchantBalanceSummaryDBFieldName.RefreshedAt),
		)
	}
	merchantBalanceSummarySelectableResponse := MerchantBalanceSummarySelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.MerchantBalanceSummaryDBFieldName.MerchantPartyId):
			key := string(MerchantBalanceSummaryDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.MerchantPartyId, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.CurrencyCode):
			key := string(MerchantBalanceSummaryDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.CurrencyCode, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.PendingAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.PendingAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.PendingAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.AvailableAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.AvailableAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.AvailableAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.ReservedAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.ReservedAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.ReservedAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.PayableAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.PayableAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.PayableAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.PaidOutAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.PaidOutAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.PaidOutAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.NegativeAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.NegativeAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.NegativeAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.RefundableAmount):
			key := string(MerchantBalanceSummaryDTOFieldName.RefundableAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.RefundableAmount, explicitAlias)

		case string(model.MerchantBalanceSummaryDBFieldName.RefreshedAt):
			key := string(MerchantBalanceSummaryDTOFieldName.RefreshedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSummarySelectableValue(merchantBalanceSummarySelectableResponse, key, merchantBalanceSummary.RefreshedAt, explicitAlias)

		}
	}
	return merchantBalanceSummarySelectableResponse
}

func NewMerchantBalanceSummaryListResponseFromFilterResult(result []model.MerchantBalanceSummaryFilterResult, filter model.Filter) MerchantBalanceSummarySelectableListResponse {
	dtoMerchantBalanceSummaryListResponse := MerchantBalanceSummarySelectableListResponse{}
	for _, row := range result {
		dtoMerchantBalanceSummaryResponse := NewMerchantBalanceSummarySelectableResponse(row.MerchantBalanceSummary, filter)
		dtoMerchantBalanceSummaryListResponse = append(dtoMerchantBalanceSummaryListResponse, &dtoMerchantBalanceSummaryResponse)
	}
	return dtoMerchantBalanceSummaryListResponse
}

type MerchantBalanceSummaryFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     MerchantBalanceSummarySelectableListResponse `json:"data"`
}

func reverseMerchantBalanceSummaryFilterResults(result []model.MerchantBalanceSummaryFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewMerchantBalanceSummaryFilterResponse(result []model.MerchantBalanceSummaryFilterResult, filter model.Filter) (resp MerchantBalanceSummaryFilterResponse) {
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
			reverseMerchantBalanceSummaryFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].MerchantPartyId
			resp.Metadata.PrevCursor = dataResult[0].MerchantPartyId
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

	resp.Data = NewMerchantBalanceSummaryListResponseFromFilterResult(dataResult, filter)
	return resp
}

type MerchantBalanceSummaryCreateRequest struct {
	CurrencyCode     string          `json:"currencyCode"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	PayableAmount    decimal.Decimal `json:"payableAmount"`
	PaidOutAmount    decimal.Decimal `json:"paidOutAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	RefundableAmount decimal.Decimal `json:"refundableAmount"`
	RefreshedAt      time.Time       `json:"refreshedAt"`
}

func (d *MerchantBalanceSummaryCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *MerchantBalanceSummaryCreateRequest) ToModel() model.MerchantBalanceSummary {
	id, _ := uuid.NewV4()
	return model.MerchantBalanceSummary{
		MerchantPartyId:  id,
		CurrencyCode:     d.CurrencyCode,
		PendingAmount:    d.PendingAmount,
		AvailableAmount:  d.AvailableAmount,
		ReservedAmount:   d.ReservedAmount,
		PayableAmount:    d.PayableAmount,
		PaidOutAmount:    d.PaidOutAmount,
		NegativeAmount:   d.NegativeAmount,
		RefundableAmount: d.RefundableAmount,
		RefreshedAt:      d.RefreshedAt,
	}
}

type MerchantBalanceSummaryListCreateRequest []*MerchantBalanceSummaryCreateRequest

func (d MerchantBalanceSummaryListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSummary := range d {
		err = validator.Struct(merchantBalanceSummary)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceSummaryListCreateRequest) ToModelList() []model.MerchantBalanceSummary {
	out := make([]model.MerchantBalanceSummary, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type MerchantBalanceSummaryUpdateRequest struct {
	CurrencyCode     string          `json:"currencyCode"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	PayableAmount    decimal.Decimal `json:"payableAmount"`
	PaidOutAmount    decimal.Decimal `json:"paidOutAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	RefundableAmount decimal.Decimal `json:"refundableAmount"`
	RefreshedAt      time.Time       `json:"refreshedAt"`
}

func (d *MerchantBalanceSummaryUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d MerchantBalanceSummaryUpdateRequest) ToModel() model.MerchantBalanceSummary {
	return model.MerchantBalanceSummary{
		CurrencyCode:     d.CurrencyCode,
		PendingAmount:    d.PendingAmount,
		AvailableAmount:  d.AvailableAmount,
		ReservedAmount:   d.ReservedAmount,
		PayableAmount:    d.PayableAmount,
		PaidOutAmount:    d.PaidOutAmount,
		NegativeAmount:   d.NegativeAmount,
		RefundableAmount: d.RefundableAmount,
		RefreshedAt:      d.RefreshedAt,
	}
}

type MerchantBalanceSummaryBulkUpdateRequest struct {
	MerchantPartyId  uuid.UUID       `json:"merchantPartyId"`
	CurrencyCode     string          `json:"currencyCode"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	PayableAmount    decimal.Decimal `json:"payableAmount"`
	PaidOutAmount    decimal.Decimal `json:"paidOutAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	RefundableAmount decimal.Decimal `json:"refundableAmount"`
	RefreshedAt      time.Time       `json:"refreshedAt"`
}

func (d MerchantBalanceSummaryBulkUpdateRequest) PrimaryID() MerchantBalanceSummaryPrimaryID {
	return MerchantBalanceSummaryPrimaryID{
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
	}
}

type MerchantBalanceSummaryListBulkUpdateRequest []*MerchantBalanceSummaryBulkUpdateRequest

func (d MerchantBalanceSummaryListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSummary := range d {
		err = validator.Struct(merchantBalanceSummary)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceSummaryBulkUpdateRequest) ToModel() model.MerchantBalanceSummary {
	return model.MerchantBalanceSummary{
		MerchantPartyId:  d.MerchantPartyId,
		CurrencyCode:     d.CurrencyCode,
		PendingAmount:    d.PendingAmount,
		AvailableAmount:  d.AvailableAmount,
		ReservedAmount:   d.ReservedAmount,
		PayableAmount:    d.PayableAmount,
		PaidOutAmount:    d.PaidOutAmount,
		NegativeAmount:   d.NegativeAmount,
		RefundableAmount: d.RefundableAmount,
		RefreshedAt:      d.RefreshedAt,
	}
}

type MerchantBalanceSummaryResponse struct {
	MerchantPartyId  uuid.UUID       `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode     string          `json:"currencyCode" validate:"required"`
	PendingAmount    decimal.Decimal `json:"pendingAmount" format:"decimal" example:"100.50"`
	AvailableAmount  decimal.Decimal `json:"availableAmount" format:"decimal" example:"100.50"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount" format:"decimal" example:"100.50"`
	PayableAmount    decimal.Decimal `json:"payableAmount" format:"decimal" example:"100.50"`
	PaidOutAmount    decimal.Decimal `json:"paidOutAmount" format:"decimal" example:"100.50"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount" format:"decimal" example:"100.50"`
	RefundableAmount decimal.Decimal `json:"refundableAmount" format:"decimal" example:"100.50"`
	RefreshedAt      time.Time       `json:"refreshedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

func NewMerchantBalanceSummaryResponse(merchantBalanceSummary model.MerchantBalanceSummary) MerchantBalanceSummaryResponse {
	return MerchantBalanceSummaryResponse{
		MerchantPartyId:  merchantBalanceSummary.MerchantPartyId,
		CurrencyCode:     merchantBalanceSummary.CurrencyCode,
		PendingAmount:    merchantBalanceSummary.PendingAmount,
		AvailableAmount:  merchantBalanceSummary.AvailableAmount,
		ReservedAmount:   merchantBalanceSummary.ReservedAmount,
		PayableAmount:    merchantBalanceSummary.PayableAmount,
		PaidOutAmount:    merchantBalanceSummary.PaidOutAmount,
		NegativeAmount:   merchantBalanceSummary.NegativeAmount,
		RefundableAmount: merchantBalanceSummary.RefundableAmount,
		RefreshedAt:      merchantBalanceSummary.RefreshedAt,
	}
}

type MerchantBalanceSummaryListResponse []*MerchantBalanceSummaryResponse

func NewMerchantBalanceSummaryListResponse(merchantBalanceSummaryList model.MerchantBalanceSummaryList) MerchantBalanceSummaryListResponse {
	dtoMerchantBalanceSummaryListResponse := MerchantBalanceSummaryListResponse{}
	for _, merchantBalanceSummary := range merchantBalanceSummaryList {
		dtoMerchantBalanceSummaryResponse := NewMerchantBalanceSummaryResponse(*merchantBalanceSummary)
		dtoMerchantBalanceSummaryListResponse = append(dtoMerchantBalanceSummaryListResponse, &dtoMerchantBalanceSummaryResponse)
	}
	return dtoMerchantBalanceSummaryListResponse
}

type MerchantBalanceSummaryPrimaryIDList []MerchantBalanceSummaryPrimaryID

func (d MerchantBalanceSummaryPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSummary := range d {
		err = validator.Struct(merchantBalanceSummary)
		if err != nil {
			return
		}
	}
	return nil
}

type MerchantBalanceSummaryPrimaryID struct {
	MerchantPartyId uuid.UUID `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode    string    `json:"currencyCode" validate:"required"`
}

func (d *MerchantBalanceSummaryPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d MerchantBalanceSummaryPrimaryID) ToModel() model.MerchantBalanceSummaryPrimaryID {
	return model.MerchantBalanceSummaryPrimaryID{
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
	}
}

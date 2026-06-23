package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RefundAllocationsDTOFieldNameType string

type refundAllocationsDTOFieldName struct {
	Id                 RefundAllocationsDTOFieldNameType
	RefundId           RefundAllocationsDTOFieldNameType
	AllocationType     RefundAllocationsDTOFieldNameType
	ResponsiblePartyId RefundAllocationsDTOFieldNameType
	Amount             RefundAllocationsDTOFieldNameType
	CurrencyCode       RefundAllocationsDTOFieldNameType
	AllocationStatus   RefundAllocationsDTOFieldNameType
	Metadata           RefundAllocationsDTOFieldNameType
	MetaCreatedAt      RefundAllocationsDTOFieldNameType
	MetaCreatedBy      RefundAllocationsDTOFieldNameType
	MetaUpdatedAt      RefundAllocationsDTOFieldNameType
	MetaUpdatedBy      RefundAllocationsDTOFieldNameType
	MetaDeletedAt      RefundAllocationsDTOFieldNameType
	MetaDeletedBy      RefundAllocationsDTOFieldNameType
}

var RefundAllocationsDTOFieldName = refundAllocationsDTOFieldName{
	Id:                 "id",
	RefundId:           "refundId",
	AllocationType:     "allocationType",
	ResponsiblePartyId: "responsiblePartyId",
	Amount:             "amount",
	CurrencyCode:       "currencyCode",
	AllocationStatus:   "allocationStatus",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformRefundAllocationsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundAllocationsDTOFieldName.Id):
		return string(model.RefundAllocationsDBFieldName.Id), true

	case string(RefundAllocationsDTOFieldName.RefundId):
		return string(model.RefundAllocationsDBFieldName.RefundId), true

	case string(RefundAllocationsDTOFieldName.AllocationType):
		return string(model.RefundAllocationsDBFieldName.AllocationType), true

	case string(RefundAllocationsDTOFieldName.ResponsiblePartyId):
		return string(model.RefundAllocationsDBFieldName.ResponsiblePartyId), true

	case string(RefundAllocationsDTOFieldName.Amount):
		return string(model.RefundAllocationsDBFieldName.Amount), true

	case string(RefundAllocationsDTOFieldName.CurrencyCode):
		return string(model.RefundAllocationsDBFieldName.CurrencyCode), true

	case string(RefundAllocationsDTOFieldName.AllocationStatus):
		return string(model.RefundAllocationsDBFieldName.AllocationStatus), true

	case string(RefundAllocationsDTOFieldName.Metadata):
		return string(model.RefundAllocationsDBFieldName.Metadata), true

	case string(RefundAllocationsDTOFieldName.MetaCreatedAt):
		return string(model.RefundAllocationsDBFieldName.MetaCreatedAt), true

	case string(RefundAllocationsDTOFieldName.MetaCreatedBy):
		return string(model.RefundAllocationsDBFieldName.MetaCreatedBy), true

	case string(RefundAllocationsDTOFieldName.MetaUpdatedAt):
		return string(model.RefundAllocationsDBFieldName.MetaUpdatedAt), true

	case string(RefundAllocationsDTOFieldName.MetaUpdatedBy):
		return string(model.RefundAllocationsDBFieldName.MetaUpdatedBy), true

	case string(RefundAllocationsDTOFieldName.MetaDeletedAt):
		return string(model.RefundAllocationsDBFieldName.MetaDeletedAt), true

	case string(RefundAllocationsDTOFieldName.MetaDeletedBy):
		return string(model.RefundAllocationsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundAllocationsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundAllocationsBaseFilterField(field string) bool {
	spec, found := model.NewRefundAllocationsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundAllocationsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundAllocationsProjectionOutputPath(path string) error {
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

func transformRefundAllocationsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundAllocationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundAllocationsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundAllocationsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundAllocationsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundAllocationsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundAllocationsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundAllocationsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundAllocationsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundAllocationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundAllocationsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundAllocationsFilter(filter *model.Filter) {
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
			Field: string(RefundAllocationsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundAllocationsSelectableResponse map[string]interface{}
type RefundAllocationsSelectableListResponse []*RefundAllocationsSelectableResponse

func assignRefundAllocationsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundAllocationsSelectableValue(out RefundAllocationsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundAllocationsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundAllocationsSelectableResponse(refundAllocations model.RefundAllocations, filter model.Filter) RefundAllocationsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundAllocationsDBFieldName.Id),
			string(model.RefundAllocationsDBFieldName.RefundId),
			string(model.RefundAllocationsDBFieldName.AllocationType),
			string(model.RefundAllocationsDBFieldName.ResponsiblePartyId),
			string(model.RefundAllocationsDBFieldName.Amount),
			string(model.RefundAllocationsDBFieldName.CurrencyCode),
			string(model.RefundAllocationsDBFieldName.AllocationStatus),
			string(model.RefundAllocationsDBFieldName.Metadata),
			string(model.RefundAllocationsDBFieldName.MetaCreatedAt),
			string(model.RefundAllocationsDBFieldName.MetaCreatedBy),
			string(model.RefundAllocationsDBFieldName.MetaUpdatedAt),
			string(model.RefundAllocationsDBFieldName.MetaUpdatedBy),
			string(model.RefundAllocationsDBFieldName.MetaDeletedAt),
			string(model.RefundAllocationsDBFieldName.MetaDeletedBy),
		)
	}
	refundAllocationsSelectableResponse := RefundAllocationsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundAllocationsDBFieldName.Id):
			key := string(RefundAllocationsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.Id, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.RefundId):
			key := string(RefundAllocationsDTOFieldName.RefundId)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.RefundId, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.AllocationType):
			key := string(RefundAllocationsDTOFieldName.AllocationType)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, model.AllocationType(refundAllocations.AllocationType), explicitAlias)

		case string(model.RefundAllocationsDBFieldName.ResponsiblePartyId):
			key := string(RefundAllocationsDTOFieldName.ResponsiblePartyId)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.ResponsiblePartyId.UUID, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.Amount):
			key := string(RefundAllocationsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.Amount, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.CurrencyCode):
			key := string(RefundAllocationsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.CurrencyCode, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.AllocationStatus):
			key := string(RefundAllocationsDTOFieldName.AllocationStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, model.AllocationStatus(refundAllocations.AllocationStatus), explicitAlias)

		case string(model.RefundAllocationsDBFieldName.Metadata):
			key := string(RefundAllocationsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.Metadata, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaCreatedAt):
			key := string(RefundAllocationsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaCreatedAt, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaCreatedBy):
			key := string(RefundAllocationsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaCreatedBy, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaUpdatedAt):
			key := string(RefundAllocationsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaUpdatedAt, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaUpdatedBy):
			key := string(RefundAllocationsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaUpdatedBy, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaDeletedAt):
			key := string(RefundAllocationsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundAllocationsDBFieldName.MetaDeletedBy):
			key := string(RefundAllocationsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundAllocationsSelectableValue(refundAllocationsSelectableResponse, key, refundAllocations.MetaDeletedBy, explicitAlias)

		}
	}
	return refundAllocationsSelectableResponse
}

func NewRefundAllocationsListResponseFromFilterResult(result []model.RefundAllocationsFilterResult, filter model.Filter) RefundAllocationsSelectableListResponse {
	dtoRefundAllocationsListResponse := RefundAllocationsSelectableListResponse{}
	for _, row := range result {
		dtoRefundAllocationsResponse := NewRefundAllocationsSelectableResponse(row.RefundAllocations, filter)
		dtoRefundAllocationsListResponse = append(dtoRefundAllocationsListResponse, &dtoRefundAllocationsResponse)
	}
	return dtoRefundAllocationsListResponse
}

type RefundAllocationsFilterResponse struct {
	Metadata Metadata                                `json:"metadata"`
	Data     RefundAllocationsSelectableListResponse `json:"data"`
}

func reverseRefundAllocationsFilterResults(result []model.RefundAllocationsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundAllocationsFilterResponse(result []model.RefundAllocationsFilterResult, filter model.Filter) (resp RefundAllocationsFilterResponse) {
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
			reverseRefundAllocationsFilterResults(dataResult)
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

	resp.Data = NewRefundAllocationsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundAllocationsCreateRequest struct {
	RefundId           uuid.UUID              `json:"refundId"`
	AllocationType     model.AllocationType   `json:"allocationType" example:"merchant_recovery" enums:"merchant_recovery,platform_fee_reversal,tax_reversal,provider_fee_loss,customer_credit"`
	ResponsiblePartyId uuid.UUID              `json:"responsiblePartyId"`
	Amount             decimal.Decimal        `json:"amount"`
	CurrencyCode       string                 `json:"currencyCode"`
	AllocationStatus   model.AllocationStatus `json:"allocationStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata           json.RawMessage        `json:"metadata"`
}

func (d *RefundAllocationsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundAllocationsCreateRequest) ToModel() model.RefundAllocations {
	id, _ := uuid.NewV4()
	return model.RefundAllocations{
		Id:                 id,
		RefundId:           d.RefundId,
		AllocationType:     d.AllocationType,
		ResponsiblePartyId: nuuid.From(d.ResponsiblePartyId),
		Amount:             d.Amount,
		CurrencyCode:       d.CurrencyCode,
		AllocationStatus:   d.AllocationStatus,
		Metadata:           d.Metadata,
	}
}

type RefundAllocationsListCreateRequest []*RefundAllocationsCreateRequest

func (d RefundAllocationsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAllocations := range d {
		err = validator.Struct(refundAllocations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundAllocationsListCreateRequest) ToModelList() []model.RefundAllocations {
	out := make([]model.RefundAllocations, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundAllocationsUpdateRequest struct {
	RefundId           uuid.UUID              `json:"refundId"`
	AllocationType     model.AllocationType   `json:"allocationType" example:"merchant_recovery" enums:"merchant_recovery,platform_fee_reversal,tax_reversal,provider_fee_loss,customer_credit"`
	ResponsiblePartyId uuid.UUID              `json:"responsiblePartyId"`
	Amount             decimal.Decimal        `json:"amount"`
	CurrencyCode       string                 `json:"currencyCode"`
	AllocationStatus   model.AllocationStatus `json:"allocationStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata           json.RawMessage        `json:"metadata"`
}

func (d *RefundAllocationsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundAllocationsUpdateRequest) ToModel() model.RefundAllocations {
	return model.RefundAllocations{
		RefundId:           d.RefundId,
		AllocationType:     d.AllocationType,
		ResponsiblePartyId: nuuid.From(d.ResponsiblePartyId),
		Amount:             d.Amount,
		CurrencyCode:       d.CurrencyCode,
		AllocationStatus:   d.AllocationStatus,
		Metadata:           d.Metadata,
	}
}

type RefundAllocationsBulkUpdateRequest struct {
	Id                 uuid.UUID              `json:"id"`
	RefundId           uuid.UUID              `json:"refundId"`
	AllocationType     model.AllocationType   `json:"allocationType" example:"merchant_recovery" enums:"merchant_recovery,platform_fee_reversal,tax_reversal,provider_fee_loss,customer_credit"`
	ResponsiblePartyId uuid.UUID              `json:"responsiblePartyId"`
	Amount             decimal.Decimal        `json:"amount"`
	CurrencyCode       string                 `json:"currencyCode"`
	AllocationStatus   model.AllocationStatus `json:"allocationStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata           json.RawMessage        `json:"metadata"`
}

func (d RefundAllocationsBulkUpdateRequest) PrimaryID() RefundAllocationsPrimaryID {
	return RefundAllocationsPrimaryID{
		Id: d.Id,
	}
}

type RefundAllocationsListBulkUpdateRequest []*RefundAllocationsBulkUpdateRequest

func (d RefundAllocationsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAllocations := range d {
		err = validator.Struct(refundAllocations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundAllocationsBulkUpdateRequest) ToModel() model.RefundAllocations {
	return model.RefundAllocations{
		Id:                 d.Id,
		RefundId:           d.RefundId,
		AllocationType:     d.AllocationType,
		ResponsiblePartyId: nuuid.From(d.ResponsiblePartyId),
		Amount:             d.Amount,
		CurrencyCode:       d.CurrencyCode,
		AllocationStatus:   d.AllocationStatus,
		Metadata:           d.Metadata,
	}
}

type RefundAllocationsResponse struct {
	Id                 uuid.UUID              `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundId           uuid.UUID              `json:"refundId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AllocationType     model.AllocationType   `json:"allocationType" validate:"required,oneof=merchant_recovery platform_fee_reversal tax_reversal provider_fee_loss customer_credit" enums:"merchant_recovery,platform_fee_reversal,tax_reversal,provider_fee_loss,customer_credit"`
	ResponsiblePartyId uuid.UUID              `json:"responsiblePartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount             decimal.Decimal        `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	CurrencyCode       string                 `json:"currencyCode" validate:"required"`
	AllocationStatus   model.AllocationStatus `json:"allocationStatus" validate:"oneof=computed posted reversed" enums:"computed,posted,reversed"`
	Metadata           json.RawMessage        `json:"metadata" swaggertype:"object"`
}

func NewRefundAllocationsResponse(refundAllocations model.RefundAllocations) RefundAllocationsResponse {
	return RefundAllocationsResponse{
		Id:                 refundAllocations.Id,
		RefundId:           refundAllocations.RefundId,
		AllocationType:     model.AllocationType(refundAllocations.AllocationType),
		ResponsiblePartyId: refundAllocations.ResponsiblePartyId.UUID,
		Amount:             refundAllocations.Amount,
		CurrencyCode:       refundAllocations.CurrencyCode,
		AllocationStatus:   model.AllocationStatus(refundAllocations.AllocationStatus),
		Metadata:           refundAllocations.Metadata,
	}
}

type RefundAllocationsListResponse []*RefundAllocationsResponse

func NewRefundAllocationsListResponse(refundAllocationsList model.RefundAllocationsList) RefundAllocationsListResponse {
	dtoRefundAllocationsListResponse := RefundAllocationsListResponse{}
	for _, refundAllocations := range refundAllocationsList {
		dtoRefundAllocationsResponse := NewRefundAllocationsResponse(*refundAllocations)
		dtoRefundAllocationsListResponse = append(dtoRefundAllocationsListResponse, &dtoRefundAllocationsResponse)
	}
	return dtoRefundAllocationsListResponse
}

type RefundAllocationsPrimaryIDList []RefundAllocationsPrimaryID

func (d RefundAllocationsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundAllocations := range d {
		err = validator.Struct(refundAllocations)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundAllocationsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundAllocationsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundAllocationsPrimaryID) ToModel() model.RefundAllocationsPrimaryID {
	return model.RefundAllocationsPrimaryID{
		Id: d.Id,
	}
}

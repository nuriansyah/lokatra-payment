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

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RefundApprovalsDTOFieldNameType string

type refundApprovalsDTOFieldName struct {
	Id             RefundApprovalsDTOFieldNameType
	RefundId       RefundApprovalsDTOFieldNameType
	ApprovalStatus RefundApprovalsDTOFieldNameType
	ApprovedAmount RefundApprovalsDTOFieldNameType
	ReasonCode     RefundApprovalsDTOFieldNameType
	ReasonDetail   RefundApprovalsDTOFieldNameType
	ApprovedBy     RefundApprovalsDTOFieldNameType
	ApprovedAt     RefundApprovalsDTOFieldNameType
	Metadata       RefundApprovalsDTOFieldNameType
	MetaCreatedAt  RefundApprovalsDTOFieldNameType
	MetaCreatedBy  RefundApprovalsDTOFieldNameType
	MetaUpdatedAt  RefundApprovalsDTOFieldNameType
	MetaUpdatedBy  RefundApprovalsDTOFieldNameType
	MetaDeletedAt  RefundApprovalsDTOFieldNameType
	MetaDeletedBy  RefundApprovalsDTOFieldNameType
}

var RefundApprovalsDTOFieldName = refundApprovalsDTOFieldName{
	Id:             "id",
	RefundId:       "refundId",
	ApprovalStatus: "approvalStatus",
	ApprovedAmount: "approvedAmount",
	ReasonCode:     "reasonCode",
	ReasonDetail:   "reasonDetail",
	ApprovedBy:     "approvedBy",
	ApprovedAt:     "approvedAt",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformRefundApprovalsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundApprovalsDTOFieldName.Id):
		return string(model.RefundApprovalsDBFieldName.Id), true

	case string(RefundApprovalsDTOFieldName.RefundId):
		return string(model.RefundApprovalsDBFieldName.RefundId), true

	case string(RefundApprovalsDTOFieldName.ApprovalStatus):
		return string(model.RefundApprovalsDBFieldName.ApprovalStatus), true

	case string(RefundApprovalsDTOFieldName.ApprovedAmount):
		return string(model.RefundApprovalsDBFieldName.ApprovedAmount), true

	case string(RefundApprovalsDTOFieldName.ReasonCode):
		return string(model.RefundApprovalsDBFieldName.ReasonCode), true

	case string(RefundApprovalsDTOFieldName.ReasonDetail):
		return string(model.RefundApprovalsDBFieldName.ReasonDetail), true

	case string(RefundApprovalsDTOFieldName.ApprovedBy):
		return string(model.RefundApprovalsDBFieldName.ApprovedBy), true

	case string(RefundApprovalsDTOFieldName.ApprovedAt):
		return string(model.RefundApprovalsDBFieldName.ApprovedAt), true

	case string(RefundApprovalsDTOFieldName.Metadata):
		return string(model.RefundApprovalsDBFieldName.Metadata), true

	case string(RefundApprovalsDTOFieldName.MetaCreatedAt):
		return string(model.RefundApprovalsDBFieldName.MetaCreatedAt), true

	case string(RefundApprovalsDTOFieldName.MetaCreatedBy):
		return string(model.RefundApprovalsDBFieldName.MetaCreatedBy), true

	case string(RefundApprovalsDTOFieldName.MetaUpdatedAt):
		return string(model.RefundApprovalsDBFieldName.MetaUpdatedAt), true

	case string(RefundApprovalsDTOFieldName.MetaUpdatedBy):
		return string(model.RefundApprovalsDBFieldName.MetaUpdatedBy), true

	case string(RefundApprovalsDTOFieldName.MetaDeletedAt):
		return string(model.RefundApprovalsDBFieldName.MetaDeletedAt), true

	case string(RefundApprovalsDTOFieldName.MetaDeletedBy):
		return string(model.RefundApprovalsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundApprovalsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundApprovalsBaseFilterField(field string) bool {
	spec, found := model.NewRefundApprovalsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundApprovalsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundApprovalsProjectionOutputPath(path string) error {
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

func transformRefundApprovalsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundApprovalsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundApprovalsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundApprovalsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundApprovalsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundApprovalsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundApprovalsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundApprovalsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundApprovalsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundApprovalsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundApprovalsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundApprovalsFilter(filter *model.Filter) {
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
			Field: string(RefundApprovalsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundApprovalsSelectableResponse map[string]interface{}
type RefundApprovalsSelectableListResponse []*RefundApprovalsSelectableResponse

func assignRefundApprovalsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundApprovalsSelectableValue(out RefundApprovalsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundApprovalsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundApprovalsSelectableResponse(refundApprovals model.RefundApprovals, filter model.Filter) RefundApprovalsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundApprovalsDBFieldName.Id),
			string(model.RefundApprovalsDBFieldName.RefundId),
			string(model.RefundApprovalsDBFieldName.ApprovalStatus),
			string(model.RefundApprovalsDBFieldName.ApprovedAmount),
			string(model.RefundApprovalsDBFieldName.ReasonCode),
			string(model.RefundApprovalsDBFieldName.ReasonDetail),
			string(model.RefundApprovalsDBFieldName.ApprovedBy),
			string(model.RefundApprovalsDBFieldName.ApprovedAt),
			string(model.RefundApprovalsDBFieldName.Metadata),
			string(model.RefundApprovalsDBFieldName.MetaCreatedAt),
			string(model.RefundApprovalsDBFieldName.MetaCreatedBy),
			string(model.RefundApprovalsDBFieldName.MetaUpdatedAt),
			string(model.RefundApprovalsDBFieldName.MetaUpdatedBy),
			string(model.RefundApprovalsDBFieldName.MetaDeletedAt),
			string(model.RefundApprovalsDBFieldName.MetaDeletedBy),
		)
	}
	refundApprovalsSelectableResponse := RefundApprovalsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundApprovalsDBFieldName.Id):
			key := string(RefundApprovalsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.Id, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.RefundId):
			key := string(RefundApprovalsDTOFieldName.RefundId)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.RefundId, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ApprovalStatus):
			key := string(RefundApprovalsDTOFieldName.ApprovalStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, model.RefundApprovalsApprovalStatus(refundApprovals.ApprovalStatus), explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ApprovedAmount):
			key := string(RefundApprovalsDTOFieldName.ApprovedAmount)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.ApprovedAmount, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ReasonCode):
			key := string(RefundApprovalsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.ReasonCode, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ReasonDetail):
			key := string(RefundApprovalsDTOFieldName.ReasonDetail)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.ReasonDetail.String, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ApprovedBy):
			key := string(RefundApprovalsDTOFieldName.ApprovedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.ApprovedBy, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.ApprovedAt):
			key := string(RefundApprovalsDTOFieldName.ApprovedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.ApprovedAt, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.Metadata):
			key := string(RefundApprovalsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.Metadata, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaCreatedAt):
			key := string(RefundApprovalsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaCreatedAt, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaCreatedBy):
			key := string(RefundApprovalsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaCreatedBy, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaUpdatedAt):
			key := string(RefundApprovalsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaUpdatedAt, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaUpdatedBy):
			key := string(RefundApprovalsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaUpdatedBy, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaDeletedAt):
			key := string(RefundApprovalsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundApprovalsDBFieldName.MetaDeletedBy):
			key := string(RefundApprovalsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundApprovalsSelectableValue(refundApprovalsSelectableResponse, key, refundApprovals.MetaDeletedBy, explicitAlias)

		}
	}
	return refundApprovalsSelectableResponse
}

func NewRefundApprovalsListResponseFromFilterResult(result []model.RefundApprovalsFilterResult, filter model.Filter) RefundApprovalsSelectableListResponse {
	dtoRefundApprovalsListResponse := RefundApprovalsSelectableListResponse{}
	for _, row := range result {
		dtoRefundApprovalsResponse := NewRefundApprovalsSelectableResponse(row.RefundApprovals, filter)
		dtoRefundApprovalsListResponse = append(dtoRefundApprovalsListResponse, &dtoRefundApprovalsResponse)
	}
	return dtoRefundApprovalsListResponse
}

type RefundApprovalsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     RefundApprovalsSelectableListResponse `json:"data"`
}

func reverseRefundApprovalsFilterResults(result []model.RefundApprovalsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundApprovalsFilterResponse(result []model.RefundApprovalsFilterResult, filter model.Filter) (resp RefundApprovalsFilterResponse) {
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
			reverseRefundApprovalsFilterResults(dataResult)
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

	resp.Data = NewRefundApprovalsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundApprovalsCreateRequest struct {
	RefundId       uuid.UUID                           `json:"refundId"`
	ApprovalStatus model.RefundApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ApprovedAmount decimal.Decimal                     `json:"approvedAmount"`
	ReasonCode     string                              `json:"reasonCode"`
	ReasonDetail   string                              `json:"reasonDetail"`
	ApprovedBy     uuid.UUID                           `json:"approvedBy"`
	ApprovedAt     time.Time                           `json:"approvedAt"`
	Metadata       json.RawMessage                     `json:"metadata"`
}

func (d *RefundApprovalsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundApprovalsCreateRequest) ToModel() model.RefundApprovals {
	id, _ := uuid.NewV4()
	return model.RefundApprovals{
		Id:             id,
		RefundId:       d.RefundId,
		ApprovalStatus: d.ApprovalStatus,
		ApprovedAmount: d.ApprovedAmount,
		ReasonCode:     d.ReasonCode,
		ReasonDetail:   null.StringFrom(d.ReasonDetail),
		ApprovedBy:     d.ApprovedBy,
		ApprovedAt:     d.ApprovedAt,
		Metadata:       d.Metadata,
	}
}

type RefundApprovalsListCreateRequest []*RefundApprovalsCreateRequest

func (d RefundApprovalsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundApprovals := range d {
		err = validator.Struct(refundApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundApprovalsListCreateRequest) ToModelList() []model.RefundApprovals {
	out := make([]model.RefundApprovals, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundApprovalsUpdateRequest struct {
	RefundId       uuid.UUID                           `json:"refundId"`
	ApprovalStatus model.RefundApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ApprovedAmount decimal.Decimal                     `json:"approvedAmount"`
	ReasonCode     string                              `json:"reasonCode"`
	ReasonDetail   string                              `json:"reasonDetail"`
	ApprovedBy     uuid.UUID                           `json:"approvedBy"`
	ApprovedAt     time.Time                           `json:"approvedAt"`
	Metadata       json.RawMessage                     `json:"metadata"`
}

func (d *RefundApprovalsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundApprovalsUpdateRequest) ToModel() model.RefundApprovals {
	return model.RefundApprovals{
		RefundId:       d.RefundId,
		ApprovalStatus: d.ApprovalStatus,
		ApprovedAmount: d.ApprovedAmount,
		ReasonCode:     d.ReasonCode,
		ReasonDetail:   null.StringFrom(d.ReasonDetail),
		ApprovedBy:     d.ApprovedBy,
		ApprovedAt:     d.ApprovedAt,
		Metadata:       d.Metadata,
	}
}

type RefundApprovalsBulkUpdateRequest struct {
	Id             uuid.UUID                           `json:"id"`
	RefundId       uuid.UUID                           `json:"refundId"`
	ApprovalStatus model.RefundApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ApprovedAmount decimal.Decimal                     `json:"approvedAmount"`
	ReasonCode     string                              `json:"reasonCode"`
	ReasonDetail   string                              `json:"reasonDetail"`
	ApprovedBy     uuid.UUID                           `json:"approvedBy"`
	ApprovedAt     time.Time                           `json:"approvedAt"`
	Metadata       json.RawMessage                     `json:"metadata"`
}

func (d RefundApprovalsBulkUpdateRequest) PrimaryID() RefundApprovalsPrimaryID {
	return RefundApprovalsPrimaryID{
		Id: d.Id,
	}
}

type RefundApprovalsListBulkUpdateRequest []*RefundApprovalsBulkUpdateRequest

func (d RefundApprovalsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundApprovals := range d {
		err = validator.Struct(refundApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundApprovalsBulkUpdateRequest) ToModel() model.RefundApprovals {
	return model.RefundApprovals{
		Id:             d.Id,
		RefundId:       d.RefundId,
		ApprovalStatus: d.ApprovalStatus,
		ApprovedAmount: d.ApprovedAmount,
		ReasonCode:     d.ReasonCode,
		ReasonDetail:   null.StringFrom(d.ReasonDetail),
		ApprovedBy:     d.ApprovedBy,
		ApprovedAt:     d.ApprovedAt,
		Metadata:       d.Metadata,
	}
}

type RefundApprovalsResponse struct {
	Id             uuid.UUID                           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundId       uuid.UUID                           `json:"refundId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ApprovalStatus model.RefundApprovalsApprovalStatus `json:"approvalStatus" validate:"required,oneof=approved rejected" enums:"approved,rejected"`
	ApprovedAmount decimal.Decimal                     `json:"approvedAmount" validate:"required" format:"decimal" example:"100.50"`
	ReasonCode     string                              `json:"reasonCode" validate:"required"`
	ReasonDetail   string                              `json:"reasonDetail"`
	ApprovedBy     uuid.UUID                           `json:"approvedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ApprovedAt     time.Time                           `json:"approvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage                     `json:"metadata" swaggertype:"object"`
}

func NewRefundApprovalsResponse(refundApprovals model.RefundApprovals) RefundApprovalsResponse {
	return RefundApprovalsResponse{
		Id:             refundApprovals.Id,
		RefundId:       refundApprovals.RefundId,
		ApprovalStatus: model.RefundApprovalsApprovalStatus(refundApprovals.ApprovalStatus),
		ApprovedAmount: refundApprovals.ApprovedAmount,
		ReasonCode:     refundApprovals.ReasonCode,
		ReasonDetail:   refundApprovals.ReasonDetail.String,
		ApprovedBy:     refundApprovals.ApprovedBy,
		ApprovedAt:     refundApprovals.ApprovedAt,
		Metadata:       refundApprovals.Metadata,
	}
}

type RefundApprovalsListResponse []*RefundApprovalsResponse

func NewRefundApprovalsListResponse(refundApprovalsList model.RefundApprovalsList) RefundApprovalsListResponse {
	dtoRefundApprovalsListResponse := RefundApprovalsListResponse{}
	for _, refundApprovals := range refundApprovalsList {
		dtoRefundApprovalsResponse := NewRefundApprovalsResponse(*refundApprovals)
		dtoRefundApprovalsListResponse = append(dtoRefundApprovalsListResponse, &dtoRefundApprovalsResponse)
	}
	return dtoRefundApprovalsListResponse
}

type RefundApprovalsPrimaryIDList []RefundApprovalsPrimaryID

func (d RefundApprovalsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundApprovals := range d {
		err = validator.Struct(refundApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundApprovalsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundApprovalsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundApprovalsPrimaryID) ToModel() model.RefundApprovalsPrimaryID {
	return model.RefundApprovalsPrimaryID{
		Id: d.Id,
	}
}

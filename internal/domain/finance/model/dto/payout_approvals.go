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

type PayoutApprovalsDTOFieldNameType string

type payoutApprovalsDTOFieldName struct {
	Id                     PayoutApprovalsDTOFieldNameType
	PayoutId               PayoutApprovalsDTOFieldNameType
	ApprovalStatus         PayoutApprovalsDTOFieldNameType
	ReasonCode             PayoutApprovalsDTOFieldNameType
	ReasonDetail           PayoutApprovalsDTOFieldNameType
	ApprovedBy             PayoutApprovalsDTOFieldNameType
	ApprovedAt             PayoutApprovalsDTOFieldNameType
	Metadata               PayoutApprovalsDTOFieldNameType
	ApprovedAmountSnapshot PayoutApprovalsDTOFieldNameType
	CurrencyCodeSnapshot   PayoutApprovalsDTOFieldNameType
	PayoutRevisionHash     PayoutApprovalsDTOFieldNameType
	MetaCreatedAt          PayoutApprovalsDTOFieldNameType
	MetaCreatedBy          PayoutApprovalsDTOFieldNameType
	MetaUpdatedAt          PayoutApprovalsDTOFieldNameType
	MetaUpdatedBy          PayoutApprovalsDTOFieldNameType
	MetaDeletedAt          PayoutApprovalsDTOFieldNameType
	MetaDeletedBy          PayoutApprovalsDTOFieldNameType
}

var PayoutApprovalsDTOFieldName = payoutApprovalsDTOFieldName{
	Id:                     "id",
	PayoutId:               "payoutId",
	ApprovalStatus:         "approvalStatus",
	ReasonCode:             "reasonCode",
	ReasonDetail:           "reasonDetail",
	ApprovedBy:             "approvedBy",
	ApprovedAt:             "approvedAt",
	Metadata:               "metadata",
	ApprovedAmountSnapshot: "approvedAmountSnapshot",
	CurrencyCodeSnapshot:   "currencyCodeSnapshot",
	PayoutRevisionHash:     "payoutRevisionHash",
	MetaCreatedAt:          "metaCreatedAt",
	MetaCreatedBy:          "metaCreatedBy",
	MetaUpdatedAt:          "metaUpdatedAt",
	MetaUpdatedBy:          "metaUpdatedBy",
	MetaDeletedAt:          "metaDeletedAt",
	MetaDeletedBy:          "metaDeletedBy",
}

func transformPayoutApprovalsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutApprovalsDTOFieldName.Id):
		return string(model.PayoutApprovalsDBFieldName.Id), true

	case string(PayoutApprovalsDTOFieldName.PayoutId):
		return string(model.PayoutApprovalsDBFieldName.PayoutId), true

	case string(PayoutApprovalsDTOFieldName.ApprovalStatus):
		return string(model.PayoutApprovalsDBFieldName.ApprovalStatus), true

	case string(PayoutApprovalsDTOFieldName.ReasonCode):
		return string(model.PayoutApprovalsDBFieldName.ReasonCode), true

	case string(PayoutApprovalsDTOFieldName.ReasonDetail):
		return string(model.PayoutApprovalsDBFieldName.ReasonDetail), true

	case string(PayoutApprovalsDTOFieldName.ApprovedBy):
		return string(model.PayoutApprovalsDBFieldName.ApprovedBy), true

	case string(PayoutApprovalsDTOFieldName.ApprovedAt):
		return string(model.PayoutApprovalsDBFieldName.ApprovedAt), true

	case string(PayoutApprovalsDTOFieldName.Metadata):
		return string(model.PayoutApprovalsDBFieldName.Metadata), true

	case string(PayoutApprovalsDTOFieldName.ApprovedAmountSnapshot):
		return string(model.PayoutApprovalsDBFieldName.ApprovedAmountSnapshot), true

	case string(PayoutApprovalsDTOFieldName.CurrencyCodeSnapshot):
		return string(model.PayoutApprovalsDBFieldName.CurrencyCodeSnapshot), true

	case string(PayoutApprovalsDTOFieldName.PayoutRevisionHash):
		return string(model.PayoutApprovalsDBFieldName.PayoutRevisionHash), true

	case string(PayoutApprovalsDTOFieldName.MetaCreatedAt):
		return string(model.PayoutApprovalsDBFieldName.MetaCreatedAt), true

	case string(PayoutApprovalsDTOFieldName.MetaCreatedBy):
		return string(model.PayoutApprovalsDBFieldName.MetaCreatedBy), true

	case string(PayoutApprovalsDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutApprovalsDBFieldName.MetaUpdatedAt), true

	case string(PayoutApprovalsDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutApprovalsDBFieldName.MetaUpdatedBy), true

	case string(PayoutApprovalsDTOFieldName.MetaDeletedAt):
		return string(model.PayoutApprovalsDBFieldName.MetaDeletedAt), true

	case string(PayoutApprovalsDTOFieldName.MetaDeletedBy):
		return string(model.PayoutApprovalsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutApprovalsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutApprovalsBaseFilterField(field string) bool {
	spec, found := model.NewPayoutApprovalsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutApprovalsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutApprovalsProjectionOutputPath(path string) error {
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

func transformPayoutApprovalsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutApprovalsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutApprovalsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutApprovalsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutApprovalsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutApprovalsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutApprovalsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutApprovalsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutApprovalsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutApprovalsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutApprovalsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutApprovalsFilter(filter *model.Filter) {
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
			Field: string(PayoutApprovalsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutApprovalsSelectableResponse map[string]interface{}
type PayoutApprovalsSelectableListResponse []*PayoutApprovalsSelectableResponse

func assignPayoutApprovalsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutApprovalsSelectableValue(out PayoutApprovalsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutApprovalsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutApprovalsSelectableResponse(payoutApprovals model.PayoutApprovals, filter model.Filter) PayoutApprovalsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutApprovalsDBFieldName.Id),
			string(model.PayoutApprovalsDBFieldName.PayoutId),
			string(model.PayoutApprovalsDBFieldName.ApprovalStatus),
			string(model.PayoutApprovalsDBFieldName.ReasonCode),
			string(model.PayoutApprovalsDBFieldName.ReasonDetail),
			string(model.PayoutApprovalsDBFieldName.ApprovedBy),
			string(model.PayoutApprovalsDBFieldName.ApprovedAt),
			string(model.PayoutApprovalsDBFieldName.Metadata),
			string(model.PayoutApprovalsDBFieldName.ApprovedAmountSnapshot),
			string(model.PayoutApprovalsDBFieldName.CurrencyCodeSnapshot),
			string(model.PayoutApprovalsDBFieldName.PayoutRevisionHash),
			string(model.PayoutApprovalsDBFieldName.MetaCreatedAt),
			string(model.PayoutApprovalsDBFieldName.MetaCreatedBy),
			string(model.PayoutApprovalsDBFieldName.MetaUpdatedAt),
			string(model.PayoutApprovalsDBFieldName.MetaUpdatedBy),
			string(model.PayoutApprovalsDBFieldName.MetaDeletedAt),
			string(model.PayoutApprovalsDBFieldName.MetaDeletedBy),
		)
	}
	payoutApprovalsSelectableResponse := PayoutApprovalsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutApprovalsDBFieldName.Id):
			key := string(PayoutApprovalsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.Id, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.PayoutId):
			key := string(PayoutApprovalsDTOFieldName.PayoutId)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.PayoutId, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ApprovalStatus):
			key := string(PayoutApprovalsDTOFieldName.ApprovalStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, model.PayoutApprovalsApprovalStatus(payoutApprovals.ApprovalStatus), explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ReasonCode):
			key := string(PayoutApprovalsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.ReasonCode, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ReasonDetail):
			key := string(PayoutApprovalsDTOFieldName.ReasonDetail)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.ReasonDetail.String, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ApprovedBy):
			key := string(PayoutApprovalsDTOFieldName.ApprovedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.ApprovedBy, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ApprovedAt):
			key := string(PayoutApprovalsDTOFieldName.ApprovedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.ApprovedAt, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.Metadata):
			key := string(PayoutApprovalsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.Metadata, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.ApprovedAmountSnapshot):
			key := string(PayoutApprovalsDTOFieldName.ApprovedAmountSnapshot)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.ApprovedAmountSnapshot.Decimal, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.CurrencyCodeSnapshot):
			key := string(PayoutApprovalsDTOFieldName.CurrencyCodeSnapshot)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.CurrencyCodeSnapshot.String, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.PayoutRevisionHash):
			key := string(PayoutApprovalsDTOFieldName.PayoutRevisionHash)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.PayoutRevisionHash.String, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaCreatedAt):
			key := string(PayoutApprovalsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaCreatedAt, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaCreatedBy):
			key := string(PayoutApprovalsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaCreatedBy, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaUpdatedAt):
			key := string(PayoutApprovalsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaUpdatedBy):
			key := string(PayoutApprovalsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaDeletedAt):
			key := string(PayoutApprovalsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutApprovalsDBFieldName.MetaDeletedBy):
			key := string(PayoutApprovalsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutApprovalsSelectableValue(payoutApprovalsSelectableResponse, key, payoutApprovals.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutApprovalsSelectableResponse
}

func NewPayoutApprovalsListResponseFromFilterResult(result []model.PayoutApprovalsFilterResult, filter model.Filter) PayoutApprovalsSelectableListResponse {
	dtoPayoutApprovalsListResponse := PayoutApprovalsSelectableListResponse{}
	for _, row := range result {
		dtoPayoutApprovalsResponse := NewPayoutApprovalsSelectableResponse(row.PayoutApprovals, filter)
		dtoPayoutApprovalsListResponse = append(dtoPayoutApprovalsListResponse, &dtoPayoutApprovalsResponse)
	}
	return dtoPayoutApprovalsListResponse
}

type PayoutApprovalsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     PayoutApprovalsSelectableListResponse `json:"data"`
}

func reversePayoutApprovalsFilterResults(result []model.PayoutApprovalsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutApprovalsFilterResponse(result []model.PayoutApprovalsFilterResult, filter model.Filter) (resp PayoutApprovalsFilterResponse) {
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
			reversePayoutApprovalsFilterResults(dataResult)
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

	resp.Data = NewPayoutApprovalsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutApprovalsCreateRequest struct {
	PayoutId               uuid.UUID                           `json:"payoutId"`
	ApprovalStatus         model.PayoutApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ReasonCode             string                              `json:"reasonCode"`
	ReasonDetail           string                              `json:"reasonDetail"`
	ApprovedBy             uuid.UUID                           `json:"approvedBy"`
	ApprovedAt             time.Time                           `json:"approvedAt"`
	Metadata               json.RawMessage                     `json:"metadata"`
	ApprovedAmountSnapshot decimal.Decimal                     `json:"approvedAmountSnapshot"`
	CurrencyCodeSnapshot   string                              `json:"currencyCodeSnapshot"`
	PayoutRevisionHash     string                              `json:"payoutRevisionHash"`
}

func (d *PayoutApprovalsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutApprovalsCreateRequest) ToModel() model.PayoutApprovals {
	id, _ := uuid.NewV4()
	return model.PayoutApprovals{
		Id:                     id,
		PayoutId:               d.PayoutId,
		ApprovalStatus:         d.ApprovalStatus,
		ReasonCode:             d.ReasonCode,
		ReasonDetail:           null.StringFrom(d.ReasonDetail),
		ApprovedBy:             d.ApprovedBy,
		ApprovedAt:             d.ApprovedAt,
		Metadata:               d.Metadata,
		ApprovedAmountSnapshot: decimal.NewNullDecimal(d.ApprovedAmountSnapshot),
		CurrencyCodeSnapshot:   null.StringFrom(d.CurrencyCodeSnapshot),
		PayoutRevisionHash:     null.StringFrom(d.PayoutRevisionHash),
	}
}

type PayoutApprovalsListCreateRequest []*PayoutApprovalsCreateRequest

func (d PayoutApprovalsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutApprovals := range d {
		err = validator.Struct(payoutApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutApprovalsListCreateRequest) ToModelList() []model.PayoutApprovals {
	out := make([]model.PayoutApprovals, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutApprovalsUpdateRequest struct {
	PayoutId               uuid.UUID                           `json:"payoutId"`
	ApprovalStatus         model.PayoutApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ReasonCode             string                              `json:"reasonCode"`
	ReasonDetail           string                              `json:"reasonDetail"`
	ApprovedBy             uuid.UUID                           `json:"approvedBy"`
	ApprovedAt             time.Time                           `json:"approvedAt"`
	Metadata               json.RawMessage                     `json:"metadata"`
	ApprovedAmountSnapshot decimal.Decimal                     `json:"approvedAmountSnapshot"`
	CurrencyCodeSnapshot   string                              `json:"currencyCodeSnapshot"`
	PayoutRevisionHash     string                              `json:"payoutRevisionHash"`
}

func (d *PayoutApprovalsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutApprovalsUpdateRequest) ToModel() model.PayoutApprovals {
	return model.PayoutApprovals{
		PayoutId:               d.PayoutId,
		ApprovalStatus:         d.ApprovalStatus,
		ReasonCode:             d.ReasonCode,
		ReasonDetail:           null.StringFrom(d.ReasonDetail),
		ApprovedBy:             d.ApprovedBy,
		ApprovedAt:             d.ApprovedAt,
		Metadata:               d.Metadata,
		ApprovedAmountSnapshot: decimal.NewNullDecimal(d.ApprovedAmountSnapshot),
		CurrencyCodeSnapshot:   null.StringFrom(d.CurrencyCodeSnapshot),
		PayoutRevisionHash:     null.StringFrom(d.PayoutRevisionHash),
	}
}

type PayoutApprovalsBulkUpdateRequest struct {
	Id                     uuid.UUID                           `json:"id"`
	PayoutId               uuid.UUID                           `json:"payoutId"`
	ApprovalStatus         model.PayoutApprovalsApprovalStatus `json:"approvalStatus" example:"approved" enums:"approved,rejected"`
	ReasonCode             string                              `json:"reasonCode"`
	ReasonDetail           string                              `json:"reasonDetail"`
	ApprovedBy             uuid.UUID                           `json:"approvedBy"`
	ApprovedAt             time.Time                           `json:"approvedAt"`
	Metadata               json.RawMessage                     `json:"metadata"`
	ApprovedAmountSnapshot decimal.Decimal                     `json:"approvedAmountSnapshot"`
	CurrencyCodeSnapshot   string                              `json:"currencyCodeSnapshot"`
	PayoutRevisionHash     string                              `json:"payoutRevisionHash"`
}

func (d PayoutApprovalsBulkUpdateRequest) PrimaryID() PayoutApprovalsPrimaryID {
	return PayoutApprovalsPrimaryID{
		Id: d.Id,
	}
}

type PayoutApprovalsListBulkUpdateRequest []*PayoutApprovalsBulkUpdateRequest

func (d PayoutApprovalsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutApprovals := range d {
		err = validator.Struct(payoutApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutApprovalsBulkUpdateRequest) ToModel() model.PayoutApprovals {
	return model.PayoutApprovals{
		Id:                     d.Id,
		PayoutId:               d.PayoutId,
		ApprovalStatus:         d.ApprovalStatus,
		ReasonCode:             d.ReasonCode,
		ReasonDetail:           null.StringFrom(d.ReasonDetail),
		ApprovedBy:             d.ApprovedBy,
		ApprovedAt:             d.ApprovedAt,
		Metadata:               d.Metadata,
		ApprovedAmountSnapshot: decimal.NewNullDecimal(d.ApprovedAmountSnapshot),
		CurrencyCodeSnapshot:   null.StringFrom(d.CurrencyCodeSnapshot),
		PayoutRevisionHash:     null.StringFrom(d.PayoutRevisionHash),
	}
}

type PayoutApprovalsResponse struct {
	Id                     uuid.UUID                           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutId               uuid.UUID                           `json:"payoutId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ApprovalStatus         model.PayoutApprovalsApprovalStatus `json:"approvalStatus" validate:"required,oneof=approved rejected" enums:"approved,rejected"`
	ReasonCode             string                              `json:"reasonCode" validate:"required"`
	ReasonDetail           string                              `json:"reasonDetail"`
	ApprovedBy             uuid.UUID                           `json:"approvedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ApprovedAt             time.Time                           `json:"approvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata               json.RawMessage                     `json:"metadata" swaggertype:"object"`
	ApprovedAmountSnapshot decimal.Decimal                     `json:"approvedAmountSnapshot" format:"decimal" example:"100.50"`
	CurrencyCodeSnapshot   string                              `json:"currencyCodeSnapshot"`
	PayoutRevisionHash     string                              `json:"payoutRevisionHash"`
}

func NewPayoutApprovalsResponse(payoutApprovals model.PayoutApprovals) PayoutApprovalsResponse {
	return PayoutApprovalsResponse{
		Id:                     payoutApprovals.Id,
		PayoutId:               payoutApprovals.PayoutId,
		ApprovalStatus:         model.PayoutApprovalsApprovalStatus(payoutApprovals.ApprovalStatus),
		ReasonCode:             payoutApprovals.ReasonCode,
		ReasonDetail:           payoutApprovals.ReasonDetail.String,
		ApprovedBy:             payoutApprovals.ApprovedBy,
		ApprovedAt:             payoutApprovals.ApprovedAt,
		Metadata:               payoutApprovals.Metadata,
		ApprovedAmountSnapshot: payoutApprovals.ApprovedAmountSnapshot.Decimal,
		CurrencyCodeSnapshot:   payoutApprovals.CurrencyCodeSnapshot.String,
		PayoutRevisionHash:     payoutApprovals.PayoutRevisionHash.String,
	}
}

type PayoutApprovalsListResponse []*PayoutApprovalsResponse

func NewPayoutApprovalsListResponse(payoutApprovalsList model.PayoutApprovalsList) PayoutApprovalsListResponse {
	dtoPayoutApprovalsListResponse := PayoutApprovalsListResponse{}
	for _, payoutApprovals := range payoutApprovalsList {
		dtoPayoutApprovalsResponse := NewPayoutApprovalsResponse(*payoutApprovals)
		dtoPayoutApprovalsListResponse = append(dtoPayoutApprovalsListResponse, &dtoPayoutApprovalsResponse)
	}
	return dtoPayoutApprovalsListResponse
}

type PayoutApprovalsPrimaryIDList []PayoutApprovalsPrimaryID

func (d PayoutApprovalsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutApprovals := range d {
		err = validator.Struct(payoutApprovals)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutApprovalsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutApprovalsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutApprovalsPrimaryID) ToModel() model.PayoutApprovalsPrimaryID {
	return model.PayoutApprovalsPrimaryID{
		Id: d.Id,
	}
}

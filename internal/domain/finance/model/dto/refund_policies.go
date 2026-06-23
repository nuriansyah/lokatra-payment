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

type RefundPoliciesDTOFieldNameType string

type refundPoliciesDTOFieldName struct {
	Id                         RefundPoliciesDTOFieldNameType
	PolicyCode                 RefundPoliciesDTOFieldNameType
	MerchantPartyId            RefundPoliciesDTOFieldNameType
	PolicyScope                RefundPoliciesDTOFieldNameType
	AllowPartial               RefundPoliciesDTOFieldNameType
	AllowPostPayout            RefundPoliciesDTOFieldNameType
	FeeReturnMode              RefundPoliciesDTOFieldNameType
	TaxReturnMode              RefundPoliciesDTOFieldNameType
	RequiresApprovalOverAmount RefundPoliciesDTOFieldNameType
	PolicyStatus               RefundPoliciesDTOFieldNameType
	Metadata                   RefundPoliciesDTOFieldNameType
	MetaCreatedAt              RefundPoliciesDTOFieldNameType
	MetaCreatedBy              RefundPoliciesDTOFieldNameType
	MetaUpdatedAt              RefundPoliciesDTOFieldNameType
	MetaUpdatedBy              RefundPoliciesDTOFieldNameType
	MetaDeletedAt              RefundPoliciesDTOFieldNameType
	MetaDeletedBy              RefundPoliciesDTOFieldNameType
}

var RefundPoliciesDTOFieldName = refundPoliciesDTOFieldName{
	Id:                         "id",
	PolicyCode:                 "policyCode",
	MerchantPartyId:            "merchantPartyId",
	PolicyScope:                "policyScope",
	AllowPartial:               "allowPartial",
	AllowPostPayout:            "allowPostPayout",
	FeeReturnMode:              "feeReturnMode",
	TaxReturnMode:              "taxReturnMode",
	RequiresApprovalOverAmount: "requiresApprovalOverAmount",
	PolicyStatus:               "policyStatus",
	Metadata:                   "metadata",
	MetaCreatedAt:              "metaCreatedAt",
	MetaCreatedBy:              "metaCreatedBy",
	MetaUpdatedAt:              "metaUpdatedAt",
	MetaUpdatedBy:              "metaUpdatedBy",
	MetaDeletedAt:              "metaDeletedAt",
	MetaDeletedBy:              "metaDeletedBy",
}

func transformRefundPoliciesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(RefundPoliciesDTOFieldName.Id):
		return string(model.RefundPoliciesDBFieldName.Id), true

	case string(RefundPoliciesDTOFieldName.PolicyCode):
		return string(model.RefundPoliciesDBFieldName.PolicyCode), true

	case string(RefundPoliciesDTOFieldName.MerchantPartyId):
		return string(model.RefundPoliciesDBFieldName.MerchantPartyId), true

	case string(RefundPoliciesDTOFieldName.PolicyScope):
		return string(model.RefundPoliciesDBFieldName.PolicyScope), true

	case string(RefundPoliciesDTOFieldName.AllowPartial):
		return string(model.RefundPoliciesDBFieldName.AllowPartial), true

	case string(RefundPoliciesDTOFieldName.AllowPostPayout):
		return string(model.RefundPoliciesDBFieldName.AllowPostPayout), true

	case string(RefundPoliciesDTOFieldName.FeeReturnMode):
		return string(model.RefundPoliciesDBFieldName.FeeReturnMode), true

	case string(RefundPoliciesDTOFieldName.TaxReturnMode):
		return string(model.RefundPoliciesDBFieldName.TaxReturnMode), true

	case string(RefundPoliciesDTOFieldName.RequiresApprovalOverAmount):
		return string(model.RefundPoliciesDBFieldName.RequiresApprovalOverAmount), true

	case string(RefundPoliciesDTOFieldName.PolicyStatus):
		return string(model.RefundPoliciesDBFieldName.PolicyStatus), true

	case string(RefundPoliciesDTOFieldName.Metadata):
		return string(model.RefundPoliciesDBFieldName.Metadata), true

	case string(RefundPoliciesDTOFieldName.MetaCreatedAt):
		return string(model.RefundPoliciesDBFieldName.MetaCreatedAt), true

	case string(RefundPoliciesDTOFieldName.MetaCreatedBy):
		return string(model.RefundPoliciesDBFieldName.MetaCreatedBy), true

	case string(RefundPoliciesDTOFieldName.MetaUpdatedAt):
		return string(model.RefundPoliciesDBFieldName.MetaUpdatedAt), true

	case string(RefundPoliciesDTOFieldName.MetaUpdatedBy):
		return string(model.RefundPoliciesDBFieldName.MetaUpdatedBy), true

	case string(RefundPoliciesDTOFieldName.MetaDeletedAt):
		return string(model.RefundPoliciesDBFieldName.MetaDeletedAt), true

	case string(RefundPoliciesDTOFieldName.MetaDeletedBy):
		return string(model.RefundPoliciesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewRefundPoliciesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isRefundPoliciesBaseFilterField(field string) bool {
	spec, found := model.NewRefundPoliciesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeRefundPoliciesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateRefundPoliciesProjectionOutputPath(path string) error {
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

func transformRefundPoliciesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformRefundPoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformRefundPoliciesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformRefundPoliciesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformRefundPoliciesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isRefundPoliciesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateRefundPoliciesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeRefundPoliciesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundPoliciesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundPoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformRefundPoliciesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultRefundPoliciesFilter(filter *model.Filter) {
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
			Field: string(RefundPoliciesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundPoliciesSelectableResponse map[string]interface{}
type RefundPoliciesSelectableListResponse []*RefundPoliciesSelectableResponse

func assignRefundPoliciesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setRefundPoliciesSelectableValue(out RefundPoliciesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignRefundPoliciesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewRefundPoliciesSelectableResponse(refundPolicies model.RefundPolicies, filter model.Filter) RefundPoliciesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundPoliciesDBFieldName.Id),
			string(model.RefundPoliciesDBFieldName.PolicyCode),
			string(model.RefundPoliciesDBFieldName.MerchantPartyId),
			string(model.RefundPoliciesDBFieldName.PolicyScope),
			string(model.RefundPoliciesDBFieldName.AllowPartial),
			string(model.RefundPoliciesDBFieldName.AllowPostPayout),
			string(model.RefundPoliciesDBFieldName.FeeReturnMode),
			string(model.RefundPoliciesDBFieldName.TaxReturnMode),
			string(model.RefundPoliciesDBFieldName.RequiresApprovalOverAmount),
			string(model.RefundPoliciesDBFieldName.PolicyStatus),
			string(model.RefundPoliciesDBFieldName.Metadata),
			string(model.RefundPoliciesDBFieldName.MetaCreatedAt),
			string(model.RefundPoliciesDBFieldName.MetaCreatedBy),
			string(model.RefundPoliciesDBFieldName.MetaUpdatedAt),
			string(model.RefundPoliciesDBFieldName.MetaUpdatedBy),
			string(model.RefundPoliciesDBFieldName.MetaDeletedAt),
			string(model.RefundPoliciesDBFieldName.MetaDeletedBy),
		)
	}
	refundPoliciesSelectableResponse := RefundPoliciesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.RefundPoliciesDBFieldName.Id):
			key := string(RefundPoliciesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.Id, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.PolicyCode):
			key := string(RefundPoliciesDTOFieldName.PolicyCode)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.PolicyCode, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MerchantPartyId):
			key := string(RefundPoliciesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MerchantPartyId.UUID, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.PolicyScope):
			key := string(RefundPoliciesDTOFieldName.PolicyScope)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, model.RefundPoliciesPolicyScope(refundPolicies.PolicyScope), explicitAlias)

		case string(model.RefundPoliciesDBFieldName.AllowPartial):
			key := string(RefundPoliciesDTOFieldName.AllowPartial)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.AllowPartial, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.AllowPostPayout):
			key := string(RefundPoliciesDTOFieldName.AllowPostPayout)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.AllowPostPayout, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.FeeReturnMode):
			key := string(RefundPoliciesDTOFieldName.FeeReturnMode)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, model.FeeReturnMode(refundPolicies.FeeReturnMode), explicitAlias)

		case string(model.RefundPoliciesDBFieldName.TaxReturnMode):
			key := string(RefundPoliciesDTOFieldName.TaxReturnMode)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, model.TaxReturnMode(refundPolicies.TaxReturnMode), explicitAlias)

		case string(model.RefundPoliciesDBFieldName.RequiresApprovalOverAmount):
			key := string(RefundPoliciesDTOFieldName.RequiresApprovalOverAmount)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.RequiresApprovalOverAmount.Decimal, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.PolicyStatus):
			key := string(RefundPoliciesDTOFieldName.PolicyStatus)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, model.RefundPoliciesPolicyStatus(refundPolicies.PolicyStatus), explicitAlias)

		case string(model.RefundPoliciesDBFieldName.Metadata):
			key := string(RefundPoliciesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.Metadata, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaCreatedAt):
			key := string(RefundPoliciesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaCreatedAt, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaCreatedBy):
			key := string(RefundPoliciesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaCreatedBy, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaUpdatedAt):
			key := string(RefundPoliciesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaUpdatedAt, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaUpdatedBy):
			key := string(RefundPoliciesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaUpdatedBy, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaDeletedAt):
			key := string(RefundPoliciesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaDeletedAt.Time, explicitAlias)

		case string(model.RefundPoliciesDBFieldName.MetaDeletedBy):
			key := string(RefundPoliciesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setRefundPoliciesSelectableValue(refundPoliciesSelectableResponse, key, refundPolicies.MetaDeletedBy, explicitAlias)

		}
	}
	return refundPoliciesSelectableResponse
}

func NewRefundPoliciesListResponseFromFilterResult(result []model.RefundPoliciesFilterResult, filter model.Filter) RefundPoliciesSelectableListResponse {
	dtoRefundPoliciesListResponse := RefundPoliciesSelectableListResponse{}
	for _, row := range result {
		dtoRefundPoliciesResponse := NewRefundPoliciesSelectableResponse(row.RefundPolicies, filter)
		dtoRefundPoliciesListResponse = append(dtoRefundPoliciesListResponse, &dtoRefundPoliciesResponse)
	}
	return dtoRefundPoliciesListResponse
}

type RefundPoliciesFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     RefundPoliciesSelectableListResponse `json:"data"`
}

func reverseRefundPoliciesFilterResults(result []model.RefundPoliciesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewRefundPoliciesFilterResponse(result []model.RefundPoliciesFilterResult, filter model.Filter) (resp RefundPoliciesFilterResponse) {
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
			reverseRefundPoliciesFilterResults(dataResult)
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

	resp.Data = NewRefundPoliciesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type RefundPoliciesCreateRequest struct {
	PolicyCode                 string                           `json:"policyCode"`
	MerchantPartyId            uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope                model.RefundPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	AllowPartial               bool                             `json:"allowPartial"`
	AllowPostPayout            bool                             `json:"allowPostPayout"`
	FeeReturnMode              model.FeeReturnMode              `json:"feeReturnMode" example:"none" enums:"none,full,proportional"`
	TaxReturnMode              model.TaxReturnMode              `json:"taxReturnMode" example:"none" enums:"none,full,proportional"`
	RequiresApprovalOverAmount decimal.Decimal                  `json:"requiresApprovalOverAmount"`
	PolicyStatus               model.RefundPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Metadata                   json.RawMessage                  `json:"metadata"`
}

func (d *RefundPoliciesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundPoliciesCreateRequest) ToModel() model.RefundPolicies {
	id, _ := uuid.NewV4()
	return model.RefundPolicies{
		Id:                         id,
		PolicyCode:                 d.PolicyCode,
		MerchantPartyId:            nuuid.From(d.MerchantPartyId),
		PolicyScope:                d.PolicyScope,
		AllowPartial:               d.AllowPartial,
		AllowPostPayout:            d.AllowPostPayout,
		FeeReturnMode:              d.FeeReturnMode,
		TaxReturnMode:              d.TaxReturnMode,
		RequiresApprovalOverAmount: decimal.NewNullDecimal(d.RequiresApprovalOverAmount),
		PolicyStatus:               d.PolicyStatus,
		Metadata:                   d.Metadata,
	}
}

type RefundPoliciesListCreateRequest []*RefundPoliciesCreateRequest

func (d RefundPoliciesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundPolicies := range d {
		err = validator.Struct(refundPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundPoliciesListCreateRequest) ToModelList() []model.RefundPolicies {
	out := make([]model.RefundPolicies, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundPoliciesUpdateRequest struct {
	PolicyCode                 string                           `json:"policyCode"`
	MerchantPartyId            uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope                model.RefundPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	AllowPartial               bool                             `json:"allowPartial"`
	AllowPostPayout            bool                             `json:"allowPostPayout"`
	FeeReturnMode              model.FeeReturnMode              `json:"feeReturnMode" example:"none" enums:"none,full,proportional"`
	TaxReturnMode              model.TaxReturnMode              `json:"taxReturnMode" example:"none" enums:"none,full,proportional"`
	RequiresApprovalOverAmount decimal.Decimal                  `json:"requiresApprovalOverAmount"`
	PolicyStatus               model.RefundPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Metadata                   json.RawMessage                  `json:"metadata"`
}

func (d *RefundPoliciesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundPoliciesUpdateRequest) ToModel() model.RefundPolicies {
	return model.RefundPolicies{
		PolicyCode:                 d.PolicyCode,
		MerchantPartyId:            nuuid.From(d.MerchantPartyId),
		PolicyScope:                d.PolicyScope,
		AllowPartial:               d.AllowPartial,
		AllowPostPayout:            d.AllowPostPayout,
		FeeReturnMode:              d.FeeReturnMode,
		TaxReturnMode:              d.TaxReturnMode,
		RequiresApprovalOverAmount: decimal.NewNullDecimal(d.RequiresApprovalOverAmount),
		PolicyStatus:               d.PolicyStatus,
		Metadata:                   d.Metadata,
	}
}

type RefundPoliciesBulkUpdateRequest struct {
	Id                         uuid.UUID                        `json:"id"`
	PolicyCode                 string                           `json:"policyCode"`
	MerchantPartyId            uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope                model.RefundPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	AllowPartial               bool                             `json:"allowPartial"`
	AllowPostPayout            bool                             `json:"allowPostPayout"`
	FeeReturnMode              model.FeeReturnMode              `json:"feeReturnMode" example:"none" enums:"none,full,proportional"`
	TaxReturnMode              model.TaxReturnMode              `json:"taxReturnMode" example:"none" enums:"none,full,proportional"`
	RequiresApprovalOverAmount decimal.Decimal                  `json:"requiresApprovalOverAmount"`
	PolicyStatus               model.RefundPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Metadata                   json.RawMessage                  `json:"metadata"`
}

func (d RefundPoliciesBulkUpdateRequest) PrimaryID() RefundPoliciesPrimaryID {
	return RefundPoliciesPrimaryID{
		Id: d.Id,
	}
}

type RefundPoliciesListBulkUpdateRequest []*RefundPoliciesBulkUpdateRequest

func (d RefundPoliciesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundPolicies := range d {
		err = validator.Struct(refundPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundPoliciesBulkUpdateRequest) ToModel() model.RefundPolicies {
	return model.RefundPolicies{
		Id:                         d.Id,
		PolicyCode:                 d.PolicyCode,
		MerchantPartyId:            nuuid.From(d.MerchantPartyId),
		PolicyScope:                d.PolicyScope,
		AllowPartial:               d.AllowPartial,
		AllowPostPayout:            d.AllowPostPayout,
		FeeReturnMode:              d.FeeReturnMode,
		TaxReturnMode:              d.TaxReturnMode,
		RequiresApprovalOverAmount: decimal.NewNullDecimal(d.RequiresApprovalOverAmount),
		PolicyStatus:               d.PolicyStatus,
		Metadata:                   d.Metadata,
	}
}

type RefundPoliciesResponse struct {
	Id                         uuid.UUID                        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyCode                 string                           `json:"policyCode" validate:"required"`
	MerchantPartyId            uuid.UUID                        `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyScope                model.RefundPoliciesPolicyScope  `json:"policyScope" validate:"required,oneof=platform_default merchant listing_type category" enums:"platform_default,merchant,listing_type,category"`
	AllowPartial               bool                             `json:"allowPartial" example:"true"`
	AllowPostPayout            bool                             `json:"allowPostPayout" example:"true"`
	FeeReturnMode              model.FeeReturnMode              `json:"feeReturnMode" validate:"required,oneof=none full proportional" enums:"none,full,proportional"`
	TaxReturnMode              model.TaxReturnMode              `json:"taxReturnMode" validate:"required,oneof=none full proportional" enums:"none,full,proportional"`
	RequiresApprovalOverAmount decimal.Decimal                  `json:"requiresApprovalOverAmount" format:"decimal" example:"100.50"`
	PolicyStatus               model.RefundPoliciesPolicyStatus `json:"policyStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Metadata                   json.RawMessage                  `json:"metadata" swaggertype:"object"`
}

func NewRefundPoliciesResponse(refundPolicies model.RefundPolicies) RefundPoliciesResponse {
	return RefundPoliciesResponse{
		Id:                         refundPolicies.Id,
		PolicyCode:                 refundPolicies.PolicyCode,
		MerchantPartyId:            refundPolicies.MerchantPartyId.UUID,
		PolicyScope:                model.RefundPoliciesPolicyScope(refundPolicies.PolicyScope),
		AllowPartial:               refundPolicies.AllowPartial,
		AllowPostPayout:            refundPolicies.AllowPostPayout,
		FeeReturnMode:              model.FeeReturnMode(refundPolicies.FeeReturnMode),
		TaxReturnMode:              model.TaxReturnMode(refundPolicies.TaxReturnMode),
		RequiresApprovalOverAmount: refundPolicies.RequiresApprovalOverAmount.Decimal,
		PolicyStatus:               model.RefundPoliciesPolicyStatus(refundPolicies.PolicyStatus),
		Metadata:                   refundPolicies.Metadata,
	}
}

type RefundPoliciesListResponse []*RefundPoliciesResponse

func NewRefundPoliciesListResponse(refundPoliciesList model.RefundPoliciesList) RefundPoliciesListResponse {
	dtoRefundPoliciesListResponse := RefundPoliciesListResponse{}
	for _, refundPolicies := range refundPoliciesList {
		dtoRefundPoliciesResponse := NewRefundPoliciesResponse(*refundPolicies)
		dtoRefundPoliciesListResponse = append(dtoRefundPoliciesListResponse, &dtoRefundPoliciesResponse)
	}
	return dtoRefundPoliciesListResponse
}

type RefundPoliciesPrimaryIDList []RefundPoliciesPrimaryID

func (d RefundPoliciesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refundPolicies := range d {
		err = validator.Struct(refundPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundPoliciesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundPoliciesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundPoliciesPrimaryID) ToModel() model.RefundPoliciesPrimaryID {
	return model.RefundPoliciesPrimaryID{
		Id: d.Id,
	}
}

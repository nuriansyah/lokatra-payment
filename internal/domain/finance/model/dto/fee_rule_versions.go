package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FeeRuleVersionsDTOFieldNameType string

type feeRuleVersionsDTOFieldName struct {
	Id            FeeRuleVersionsDTOFieldNameType
	RuleSetId     FeeRuleVersionsDTOFieldNameType
	VersionNo     FeeRuleVersionsDTOFieldNameType
	FormulaType   FeeRuleVersionsDTOFieldNameType
	AppliesTo     FeeRuleVersionsDTOFieldNameType
	PayerType     FeeRuleVersionsDTOFieldNameType
	RecipientType FeeRuleVersionsDTOFieldNameType
	Conditions    FeeRuleVersionsDTOFieldNameType
	IsCurrent     FeeRuleVersionsDTOFieldNameType
	MetaCreatedAt FeeRuleVersionsDTOFieldNameType
	MetaCreatedBy FeeRuleVersionsDTOFieldNameType
	MetaUpdatedAt FeeRuleVersionsDTOFieldNameType
	MetaUpdatedBy FeeRuleVersionsDTOFieldNameType
	MetaDeletedAt FeeRuleVersionsDTOFieldNameType
	MetaDeletedBy FeeRuleVersionsDTOFieldNameType
}

var FeeRuleVersionsDTOFieldName = feeRuleVersionsDTOFieldName{
	Id:            "id",
	RuleSetId:     "ruleSetId",
	VersionNo:     "versionNo",
	FormulaType:   "formulaType",
	AppliesTo:     "appliesTo",
	PayerType:     "payerType",
	RecipientType: "recipientType",
	Conditions:    "conditions",
	IsCurrent:     "isCurrent",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformFeeRuleVersionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FeeRuleVersionsDTOFieldName.Id):
		return string(model.FeeRuleVersionsDBFieldName.Id), true

	case string(FeeRuleVersionsDTOFieldName.RuleSetId):
		return string(model.FeeRuleVersionsDBFieldName.RuleSetId), true

	case string(FeeRuleVersionsDTOFieldName.VersionNo):
		return string(model.FeeRuleVersionsDBFieldName.VersionNo), true

	case string(FeeRuleVersionsDTOFieldName.FormulaType):
		return string(model.FeeRuleVersionsDBFieldName.FormulaType), true

	case string(FeeRuleVersionsDTOFieldName.AppliesTo):
		return string(model.FeeRuleVersionsDBFieldName.AppliesTo), true

	case string(FeeRuleVersionsDTOFieldName.PayerType):
		return string(model.FeeRuleVersionsDBFieldName.PayerType), true

	case string(FeeRuleVersionsDTOFieldName.RecipientType):
		return string(model.FeeRuleVersionsDBFieldName.RecipientType), true

	case string(FeeRuleVersionsDTOFieldName.Conditions):
		return string(model.FeeRuleVersionsDBFieldName.Conditions), true

	case string(FeeRuleVersionsDTOFieldName.IsCurrent):
		return string(model.FeeRuleVersionsDBFieldName.IsCurrent), true

	case string(FeeRuleVersionsDTOFieldName.MetaCreatedAt):
		return string(model.FeeRuleVersionsDBFieldName.MetaCreatedAt), true

	case string(FeeRuleVersionsDTOFieldName.MetaCreatedBy):
		return string(model.FeeRuleVersionsDBFieldName.MetaCreatedBy), true

	case string(FeeRuleVersionsDTOFieldName.MetaUpdatedAt):
		return string(model.FeeRuleVersionsDBFieldName.MetaUpdatedAt), true

	case string(FeeRuleVersionsDTOFieldName.MetaUpdatedBy):
		return string(model.FeeRuleVersionsDBFieldName.MetaUpdatedBy), true

	case string(FeeRuleVersionsDTOFieldName.MetaDeletedAt):
		return string(model.FeeRuleVersionsDBFieldName.MetaDeletedAt), true

	case string(FeeRuleVersionsDTOFieldName.MetaDeletedBy):
		return string(model.FeeRuleVersionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFeeRuleVersionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFeeRuleVersionsBaseFilterField(field string) bool {
	spec, found := model.NewFeeRuleVersionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFeeRuleVersionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFeeRuleVersionsProjectionOutputPath(path string) error {
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

func transformFeeRuleVersionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFeeRuleVersionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFeeRuleVersionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFeeRuleVersionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFeeRuleVersionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFeeRuleVersionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFeeRuleVersionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFeeRuleVersionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFeeRuleVersionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFeeRuleVersionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFeeRuleVersionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFeeRuleVersionsFilter(filter *model.Filter) {
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
			Field: string(FeeRuleVersionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FeeRuleVersionsSelectableResponse map[string]interface{}
type FeeRuleVersionsSelectableListResponse []*FeeRuleVersionsSelectableResponse

func assignFeeRuleVersionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFeeRuleVersionsSelectableValue(out FeeRuleVersionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFeeRuleVersionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFeeRuleVersionsSelectableResponse(feeRuleVersions model.FeeRuleVersions, filter model.Filter) FeeRuleVersionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FeeRuleVersionsDBFieldName.Id),
			string(model.FeeRuleVersionsDBFieldName.RuleSetId),
			string(model.FeeRuleVersionsDBFieldName.VersionNo),
			string(model.FeeRuleVersionsDBFieldName.FormulaType),
			string(model.FeeRuleVersionsDBFieldName.AppliesTo),
			string(model.FeeRuleVersionsDBFieldName.PayerType),
			string(model.FeeRuleVersionsDBFieldName.RecipientType),
			string(model.FeeRuleVersionsDBFieldName.Conditions),
			string(model.FeeRuleVersionsDBFieldName.IsCurrent),
			string(model.FeeRuleVersionsDBFieldName.MetaCreatedAt),
			string(model.FeeRuleVersionsDBFieldName.MetaCreatedBy),
			string(model.FeeRuleVersionsDBFieldName.MetaUpdatedAt),
			string(model.FeeRuleVersionsDBFieldName.MetaUpdatedBy),
			string(model.FeeRuleVersionsDBFieldName.MetaDeletedAt),
			string(model.FeeRuleVersionsDBFieldName.MetaDeletedBy),
		)
	}
	feeRuleVersionsSelectableResponse := FeeRuleVersionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FeeRuleVersionsDBFieldName.Id):
			key := string(FeeRuleVersionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.Id, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.RuleSetId):
			key := string(FeeRuleVersionsDTOFieldName.RuleSetId)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.RuleSetId, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.VersionNo):
			key := string(FeeRuleVersionsDTOFieldName.VersionNo)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.VersionNo, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.FormulaType):
			key := string(FeeRuleVersionsDTOFieldName.FormulaType)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, model.FormulaType(feeRuleVersions.FormulaType), explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.AppliesTo):
			key := string(FeeRuleVersionsDTOFieldName.AppliesTo)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, model.AppliesTo(feeRuleVersions.AppliesTo), explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.PayerType):
			key := string(FeeRuleVersionsDTOFieldName.PayerType)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, model.PayerType(feeRuleVersions.PayerType), explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.RecipientType):
			key := string(FeeRuleVersionsDTOFieldName.RecipientType)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, model.RecipientType(feeRuleVersions.RecipientType), explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.Conditions):
			key := string(FeeRuleVersionsDTOFieldName.Conditions)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.Conditions, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.IsCurrent):
			key := string(FeeRuleVersionsDTOFieldName.IsCurrent)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.IsCurrent, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaCreatedAt):
			key := string(FeeRuleVersionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaCreatedAt, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaCreatedBy):
			key := string(FeeRuleVersionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaCreatedBy, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaUpdatedAt):
			key := string(FeeRuleVersionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaUpdatedAt, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaUpdatedBy):
			key := string(FeeRuleVersionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaUpdatedBy, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaDeletedAt):
			key := string(FeeRuleVersionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaDeletedAt.Time, explicitAlias)

		case string(model.FeeRuleVersionsDBFieldName.MetaDeletedBy):
			key := string(FeeRuleVersionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleVersionsSelectableValue(feeRuleVersionsSelectableResponse, key, feeRuleVersions.MetaDeletedBy, explicitAlias)

		}
	}
	return feeRuleVersionsSelectableResponse
}

func NewFeeRuleVersionsListResponseFromFilterResult(result []model.FeeRuleVersionsFilterResult, filter model.Filter) FeeRuleVersionsSelectableListResponse {
	dtoFeeRuleVersionsListResponse := FeeRuleVersionsSelectableListResponse{}
	for _, row := range result {
		dtoFeeRuleVersionsResponse := NewFeeRuleVersionsSelectableResponse(row.FeeRuleVersions, filter)
		dtoFeeRuleVersionsListResponse = append(dtoFeeRuleVersionsListResponse, &dtoFeeRuleVersionsResponse)
	}
	return dtoFeeRuleVersionsListResponse
}

type FeeRuleVersionsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     FeeRuleVersionsSelectableListResponse `json:"data"`
}

func reverseFeeRuleVersionsFilterResults(result []model.FeeRuleVersionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFeeRuleVersionsFilterResponse(result []model.FeeRuleVersionsFilterResult, filter model.Filter) (resp FeeRuleVersionsFilterResponse) {
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
			reverseFeeRuleVersionsFilterResults(dataResult)
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

	resp.Data = NewFeeRuleVersionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FeeRuleVersionsCreateRequest struct {
	RuleSetId     uuid.UUID           `json:"ruleSetId"`
	VersionNo     int                 `json:"versionNo"`
	FormulaType   model.FormulaType   `json:"formulaType" example:"flat" enums:"flat,percentage,tiered,hybrid"`
	AppliesTo     model.AppliesTo     `json:"appliesTo" example:"gmv" enums:"gmv,net,commissionable_amount,payout,refund"`
	PayerType     model.PayerType     `json:"payerType" example:"customer" enums:"customer,merchant,platform"`
	RecipientType model.RecipientType `json:"recipientType" example:"platform" enums:"platform,provider,tax_authority,affiliate"`
	Conditions    json.RawMessage     `json:"conditions"`
	IsCurrent     bool                `json:"isCurrent"`
}

func (d *FeeRuleVersionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FeeRuleVersionsCreateRequest) ToModel() model.FeeRuleVersions {
	id, _ := uuid.NewV4()
	return model.FeeRuleVersions{
		Id:            id,
		RuleSetId:     d.RuleSetId,
		VersionNo:     d.VersionNo,
		FormulaType:   d.FormulaType,
		AppliesTo:     d.AppliesTo,
		PayerType:     d.PayerType,
		RecipientType: d.RecipientType,
		Conditions:    d.Conditions,
		IsCurrent:     d.IsCurrent,
	}
}

type FeeRuleVersionsListCreateRequest []*FeeRuleVersionsCreateRequest

func (d FeeRuleVersionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleVersions := range d {
		err = validator.Struct(feeRuleVersions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRuleVersionsListCreateRequest) ToModelList() []model.FeeRuleVersions {
	out := make([]model.FeeRuleVersions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FeeRuleVersionsUpdateRequest struct {
	RuleSetId     uuid.UUID           `json:"ruleSetId"`
	VersionNo     int                 `json:"versionNo"`
	FormulaType   model.FormulaType   `json:"formulaType" example:"flat" enums:"flat,percentage,tiered,hybrid"`
	AppliesTo     model.AppliesTo     `json:"appliesTo" example:"gmv" enums:"gmv,net,commissionable_amount,payout,refund"`
	PayerType     model.PayerType     `json:"payerType" example:"customer" enums:"customer,merchant,platform"`
	RecipientType model.RecipientType `json:"recipientType" example:"platform" enums:"platform,provider,tax_authority,affiliate"`
	Conditions    json.RawMessage     `json:"conditions"`
	IsCurrent     bool                `json:"isCurrent"`
}

func (d *FeeRuleVersionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FeeRuleVersionsUpdateRequest) ToModel() model.FeeRuleVersions {
	return model.FeeRuleVersions{
		RuleSetId:     d.RuleSetId,
		VersionNo:     d.VersionNo,
		FormulaType:   d.FormulaType,
		AppliesTo:     d.AppliesTo,
		PayerType:     d.PayerType,
		RecipientType: d.RecipientType,
		Conditions:    d.Conditions,
		IsCurrent:     d.IsCurrent,
	}
}

type FeeRuleVersionsBulkUpdateRequest struct {
	Id            uuid.UUID           `json:"id"`
	RuleSetId     uuid.UUID           `json:"ruleSetId"`
	VersionNo     int                 `json:"versionNo"`
	FormulaType   model.FormulaType   `json:"formulaType" example:"flat" enums:"flat,percentage,tiered,hybrid"`
	AppliesTo     model.AppliesTo     `json:"appliesTo" example:"gmv" enums:"gmv,net,commissionable_amount,payout,refund"`
	PayerType     model.PayerType     `json:"payerType" example:"customer" enums:"customer,merchant,platform"`
	RecipientType model.RecipientType `json:"recipientType" example:"platform" enums:"platform,provider,tax_authority,affiliate"`
	Conditions    json.RawMessage     `json:"conditions"`
	IsCurrent     bool                `json:"isCurrent"`
}

func (d FeeRuleVersionsBulkUpdateRequest) PrimaryID() FeeRuleVersionsPrimaryID {
	return FeeRuleVersionsPrimaryID{
		Id: d.Id,
	}
}

type FeeRuleVersionsListBulkUpdateRequest []*FeeRuleVersionsBulkUpdateRequest

func (d FeeRuleVersionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleVersions := range d {
		err = validator.Struct(feeRuleVersions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRuleVersionsBulkUpdateRequest) ToModel() model.FeeRuleVersions {
	return model.FeeRuleVersions{
		Id:            d.Id,
		RuleSetId:     d.RuleSetId,
		VersionNo:     d.VersionNo,
		FormulaType:   d.FormulaType,
		AppliesTo:     d.AppliesTo,
		PayerType:     d.PayerType,
		RecipientType: d.RecipientType,
		Conditions:    d.Conditions,
		IsCurrent:     d.IsCurrent,
	}
}

type FeeRuleVersionsResponse struct {
	Id            uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RuleSetId     uuid.UUID           `json:"ruleSetId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	VersionNo     int                 `json:"versionNo" validate:"required" example:"1"`
	FormulaType   model.FormulaType   `json:"formulaType" validate:"required,oneof=flat percentage tiered hybrid" enums:"flat,percentage,tiered,hybrid"`
	AppliesTo     model.AppliesTo     `json:"appliesTo" validate:"required,oneof=gmv net commissionable_amount payout refund" enums:"gmv,net,commissionable_amount,payout,refund"`
	PayerType     model.PayerType     `json:"payerType" validate:"required,oneof=customer merchant platform" enums:"customer,merchant,platform"`
	RecipientType model.RecipientType `json:"recipientType" validate:"required,oneof=platform provider tax_authority affiliate" enums:"platform,provider,tax_authority,affiliate"`
	Conditions    json.RawMessage     `json:"conditions" swaggertype:"object"`
	IsCurrent     bool                `json:"isCurrent" example:"true"`
}

func NewFeeRuleVersionsResponse(feeRuleVersions model.FeeRuleVersions) FeeRuleVersionsResponse {
	return FeeRuleVersionsResponse{
		Id:            feeRuleVersions.Id,
		RuleSetId:     feeRuleVersions.RuleSetId,
		VersionNo:     feeRuleVersions.VersionNo,
		FormulaType:   model.FormulaType(feeRuleVersions.FormulaType),
		AppliesTo:     model.AppliesTo(feeRuleVersions.AppliesTo),
		PayerType:     model.PayerType(feeRuleVersions.PayerType),
		RecipientType: model.RecipientType(feeRuleVersions.RecipientType),
		Conditions:    feeRuleVersions.Conditions,
		IsCurrent:     feeRuleVersions.IsCurrent,
	}
}

type FeeRuleVersionsListResponse []*FeeRuleVersionsResponse

func NewFeeRuleVersionsListResponse(feeRuleVersionsList model.FeeRuleVersionsList) FeeRuleVersionsListResponse {
	dtoFeeRuleVersionsListResponse := FeeRuleVersionsListResponse{}
	for _, feeRuleVersions := range feeRuleVersionsList {
		dtoFeeRuleVersionsResponse := NewFeeRuleVersionsResponse(*feeRuleVersions)
		dtoFeeRuleVersionsListResponse = append(dtoFeeRuleVersionsListResponse, &dtoFeeRuleVersionsResponse)
	}
	return dtoFeeRuleVersionsListResponse
}

type FeeRuleVersionsPrimaryIDList []FeeRuleVersionsPrimaryID

func (d FeeRuleVersionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleVersions := range d {
		err = validator.Struct(feeRuleVersions)
		if err != nil {
			return
		}
	}
	return nil
}

type FeeRuleVersionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FeeRuleVersionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FeeRuleVersionsPrimaryID) ToModel() model.FeeRuleVersionsPrimaryID {
	return model.FeeRuleVersionsPrimaryID{
		Id: d.Id,
	}
}

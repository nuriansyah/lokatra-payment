package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FeeRulesDTOFieldNameType string

type feeRulesDTOFieldName struct {
	Id               FeeRulesDTOFieldNameType
	FeeRuleVersionId FeeRulesDTOFieldNameType
	RuleName         FeeRulesDTOFieldNameType
	MinAmount        FeeRulesDTOFieldNameType
	MaxAmount        FeeRulesDTOFieldNameType
	PercentageRate   FeeRulesDTOFieldNameType
	FlatAmount       FeeRulesDTOFieldNameType
	CapAmount        FeeRulesDTOFieldNameType
	FloorAmount      FeeRulesDTOFieldNameType
	TaxInclusive     FeeRulesDTOFieldNameType
	SortOrder        FeeRulesDTOFieldNameType
	Metadata         FeeRulesDTOFieldNameType
	MetaCreatedAt    FeeRulesDTOFieldNameType
	MetaCreatedBy    FeeRulesDTOFieldNameType
	MetaUpdatedAt    FeeRulesDTOFieldNameType
	MetaUpdatedBy    FeeRulesDTOFieldNameType
	MetaDeletedAt    FeeRulesDTOFieldNameType
	MetaDeletedBy    FeeRulesDTOFieldNameType
}

var FeeRulesDTOFieldName = feeRulesDTOFieldName{
	Id:               "id",
	FeeRuleVersionId: "feeRuleVersionId",
	RuleName:         "ruleName",
	MinAmount:        "minAmount",
	MaxAmount:        "maxAmount",
	PercentageRate:   "percentageRate",
	FlatAmount:       "flatAmount",
	CapAmount:        "capAmount",
	FloorAmount:      "floorAmount",
	TaxInclusive:     "taxInclusive",
	SortOrder:        "sortOrder",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformFeeRulesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FeeRulesDTOFieldName.Id):
		return string(model.FeeRulesDBFieldName.Id), true

	case string(FeeRulesDTOFieldName.FeeRuleVersionId):
		return string(model.FeeRulesDBFieldName.FeeRuleVersionId), true

	case string(FeeRulesDTOFieldName.RuleName):
		return string(model.FeeRulesDBFieldName.RuleName), true

	case string(FeeRulesDTOFieldName.MinAmount):
		return string(model.FeeRulesDBFieldName.MinAmount), true

	case string(FeeRulesDTOFieldName.MaxAmount):
		return string(model.FeeRulesDBFieldName.MaxAmount), true

	case string(FeeRulesDTOFieldName.PercentageRate):
		return string(model.FeeRulesDBFieldName.PercentageRate), true

	case string(FeeRulesDTOFieldName.FlatAmount):
		return string(model.FeeRulesDBFieldName.FlatAmount), true

	case string(FeeRulesDTOFieldName.CapAmount):
		return string(model.FeeRulesDBFieldName.CapAmount), true

	case string(FeeRulesDTOFieldName.FloorAmount):
		return string(model.FeeRulesDBFieldName.FloorAmount), true

	case string(FeeRulesDTOFieldName.TaxInclusive):
		return string(model.FeeRulesDBFieldName.TaxInclusive), true

	case string(FeeRulesDTOFieldName.SortOrder):
		return string(model.FeeRulesDBFieldName.SortOrder), true

	case string(FeeRulesDTOFieldName.Metadata):
		return string(model.FeeRulesDBFieldName.Metadata), true

	case string(FeeRulesDTOFieldName.MetaCreatedAt):
		return string(model.FeeRulesDBFieldName.MetaCreatedAt), true

	case string(FeeRulesDTOFieldName.MetaCreatedBy):
		return string(model.FeeRulesDBFieldName.MetaCreatedBy), true

	case string(FeeRulesDTOFieldName.MetaUpdatedAt):
		return string(model.FeeRulesDBFieldName.MetaUpdatedAt), true

	case string(FeeRulesDTOFieldName.MetaUpdatedBy):
		return string(model.FeeRulesDBFieldName.MetaUpdatedBy), true

	case string(FeeRulesDTOFieldName.MetaDeletedAt):
		return string(model.FeeRulesDBFieldName.MetaDeletedAt), true

	case string(FeeRulesDTOFieldName.MetaDeletedBy):
		return string(model.FeeRulesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFeeRulesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFeeRulesBaseFilterField(field string) bool {
	spec, found := model.NewFeeRulesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFeeRulesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFeeRulesProjectionOutputPath(path string) error {
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

func transformFeeRulesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFeeRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFeeRulesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFeeRulesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFeeRulesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFeeRulesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFeeRulesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFeeRulesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFeeRulesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFeeRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFeeRulesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFeeRulesFilter(filter *model.Filter) {
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
			Field: string(FeeRulesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FeeRulesSelectableResponse map[string]interface{}
type FeeRulesSelectableListResponse []*FeeRulesSelectableResponse

func assignFeeRulesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFeeRulesSelectableValue(out FeeRulesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFeeRulesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFeeRulesSelectableResponse(feeRules model.FeeRules, filter model.Filter) FeeRulesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FeeRulesDBFieldName.Id),
			string(model.FeeRulesDBFieldName.FeeRuleVersionId),
			string(model.FeeRulesDBFieldName.RuleName),
			string(model.FeeRulesDBFieldName.MinAmount),
			string(model.FeeRulesDBFieldName.MaxAmount),
			string(model.FeeRulesDBFieldName.PercentageRate),
			string(model.FeeRulesDBFieldName.FlatAmount),
			string(model.FeeRulesDBFieldName.CapAmount),
			string(model.FeeRulesDBFieldName.FloorAmount),
			string(model.FeeRulesDBFieldName.TaxInclusive),
			string(model.FeeRulesDBFieldName.SortOrder),
			string(model.FeeRulesDBFieldName.Metadata),
			string(model.FeeRulesDBFieldName.MetaCreatedAt),
			string(model.FeeRulesDBFieldName.MetaCreatedBy),
			string(model.FeeRulesDBFieldName.MetaUpdatedAt),
			string(model.FeeRulesDBFieldName.MetaUpdatedBy),
			string(model.FeeRulesDBFieldName.MetaDeletedAt),
			string(model.FeeRulesDBFieldName.MetaDeletedBy),
		)
	}
	feeRulesSelectableResponse := FeeRulesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FeeRulesDBFieldName.Id):
			key := string(FeeRulesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.Id, explicitAlias)

		case string(model.FeeRulesDBFieldName.FeeRuleVersionId):
			key := string(FeeRulesDTOFieldName.FeeRuleVersionId)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.FeeRuleVersionId, explicitAlias)

		case string(model.FeeRulesDBFieldName.RuleName):
			key := string(FeeRulesDTOFieldName.RuleName)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.RuleName, explicitAlias)

		case string(model.FeeRulesDBFieldName.MinAmount):
			key := string(FeeRulesDTOFieldName.MinAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MinAmount.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.MaxAmount):
			key := string(FeeRulesDTOFieldName.MaxAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MaxAmount.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.PercentageRate):
			key := string(FeeRulesDTOFieldName.PercentageRate)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.PercentageRate.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.FlatAmount):
			key := string(FeeRulesDTOFieldName.FlatAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.FlatAmount.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.CapAmount):
			key := string(FeeRulesDTOFieldName.CapAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.CapAmount.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.FloorAmount):
			key := string(FeeRulesDTOFieldName.FloorAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.FloorAmount.Decimal, explicitAlias)

		case string(model.FeeRulesDBFieldName.TaxInclusive):
			key := string(FeeRulesDTOFieldName.TaxInclusive)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.TaxInclusive, explicitAlias)

		case string(model.FeeRulesDBFieldName.SortOrder):
			key := string(FeeRulesDTOFieldName.SortOrder)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.SortOrder, explicitAlias)

		case string(model.FeeRulesDBFieldName.Metadata):
			key := string(FeeRulesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.Metadata, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaCreatedAt):
			key := string(FeeRulesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaCreatedAt, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaCreatedBy):
			key := string(FeeRulesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaCreatedBy, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaUpdatedAt):
			key := string(FeeRulesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaUpdatedAt, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaUpdatedBy):
			key := string(FeeRulesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaUpdatedBy, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaDeletedAt):
			key := string(FeeRulesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaDeletedAt.Time, explicitAlias)

		case string(model.FeeRulesDBFieldName.MetaDeletedBy):
			key := string(FeeRulesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRulesSelectableValue(feeRulesSelectableResponse, key, feeRules.MetaDeletedBy, explicitAlias)

		}
	}
	return feeRulesSelectableResponse
}

func NewFeeRulesListResponseFromFilterResult(result []model.FeeRulesFilterResult, filter model.Filter) FeeRulesSelectableListResponse {
	dtoFeeRulesListResponse := FeeRulesSelectableListResponse{}
	for _, row := range result {
		dtoFeeRulesResponse := NewFeeRulesSelectableResponse(row.FeeRules, filter)
		dtoFeeRulesListResponse = append(dtoFeeRulesListResponse, &dtoFeeRulesResponse)
	}
	return dtoFeeRulesListResponse
}

type FeeRulesFilterResponse struct {
	Metadata Metadata                       `json:"metadata"`
	Data     FeeRulesSelectableListResponse `json:"data"`
}

func reverseFeeRulesFilterResults(result []model.FeeRulesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFeeRulesFilterResponse(result []model.FeeRulesFilterResult, filter model.Filter) (resp FeeRulesFilterResponse) {
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
			reverseFeeRulesFilterResults(dataResult)
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

	resp.Data = NewFeeRulesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FeeRulesCreateRequest struct {
	FeeRuleVersionId uuid.UUID       `json:"feeRuleVersionId"`
	RuleName         string          `json:"ruleName"`
	MinAmount        decimal.Decimal `json:"minAmount"`
	MaxAmount        decimal.Decimal `json:"maxAmount"`
	PercentageRate   decimal.Decimal `json:"percentageRate"`
	FlatAmount       decimal.Decimal `json:"flatAmount"`
	CapAmount        decimal.Decimal `json:"capAmount"`
	FloorAmount      decimal.Decimal `json:"floorAmount"`
	TaxInclusive     bool            `json:"taxInclusive"`
	SortOrder        int             `json:"sortOrder"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *FeeRulesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FeeRulesCreateRequest) ToModel() model.FeeRules {
	id, _ := uuid.NewV4()
	return model.FeeRules{
		Id:               id,
		FeeRuleVersionId: d.FeeRuleVersionId,
		RuleName:         d.RuleName,
		MinAmount:        decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:        decimal.NewNullDecimal(d.MaxAmount),
		PercentageRate:   decimal.NewNullDecimal(d.PercentageRate),
		FlatAmount:       decimal.NewNullDecimal(d.FlatAmount),
		CapAmount:        decimal.NewNullDecimal(d.CapAmount),
		FloorAmount:      decimal.NewNullDecimal(d.FloorAmount),
		TaxInclusive:     d.TaxInclusive,
		SortOrder:        d.SortOrder,
		Metadata:         d.Metadata,
	}
}

type FeeRulesListCreateRequest []*FeeRulesCreateRequest

func (d FeeRulesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRules := range d {
		err = validator.Struct(feeRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRulesListCreateRequest) ToModelList() []model.FeeRules {
	out := make([]model.FeeRules, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FeeRulesUpdateRequest struct {
	FeeRuleVersionId uuid.UUID       `json:"feeRuleVersionId"`
	RuleName         string          `json:"ruleName"`
	MinAmount        decimal.Decimal `json:"minAmount"`
	MaxAmount        decimal.Decimal `json:"maxAmount"`
	PercentageRate   decimal.Decimal `json:"percentageRate"`
	FlatAmount       decimal.Decimal `json:"flatAmount"`
	CapAmount        decimal.Decimal `json:"capAmount"`
	FloorAmount      decimal.Decimal `json:"floorAmount"`
	TaxInclusive     bool            `json:"taxInclusive"`
	SortOrder        int             `json:"sortOrder"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *FeeRulesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FeeRulesUpdateRequest) ToModel() model.FeeRules {
	return model.FeeRules{
		FeeRuleVersionId: d.FeeRuleVersionId,
		RuleName:         d.RuleName,
		MinAmount:        decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:        decimal.NewNullDecimal(d.MaxAmount),
		PercentageRate:   decimal.NewNullDecimal(d.PercentageRate),
		FlatAmount:       decimal.NewNullDecimal(d.FlatAmount),
		CapAmount:        decimal.NewNullDecimal(d.CapAmount),
		FloorAmount:      decimal.NewNullDecimal(d.FloorAmount),
		TaxInclusive:     d.TaxInclusive,
		SortOrder:        d.SortOrder,
		Metadata:         d.Metadata,
	}
}

type FeeRulesBulkUpdateRequest struct {
	Id               uuid.UUID       `json:"id"`
	FeeRuleVersionId uuid.UUID       `json:"feeRuleVersionId"`
	RuleName         string          `json:"ruleName"`
	MinAmount        decimal.Decimal `json:"minAmount"`
	MaxAmount        decimal.Decimal `json:"maxAmount"`
	PercentageRate   decimal.Decimal `json:"percentageRate"`
	FlatAmount       decimal.Decimal `json:"flatAmount"`
	CapAmount        decimal.Decimal `json:"capAmount"`
	FloorAmount      decimal.Decimal `json:"floorAmount"`
	TaxInclusive     bool            `json:"taxInclusive"`
	SortOrder        int             `json:"sortOrder"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d FeeRulesBulkUpdateRequest) PrimaryID() FeeRulesPrimaryID {
	return FeeRulesPrimaryID{
		Id: d.Id,
	}
}

type FeeRulesListBulkUpdateRequest []*FeeRulesBulkUpdateRequest

func (d FeeRulesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRules := range d {
		err = validator.Struct(feeRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRulesBulkUpdateRequest) ToModel() model.FeeRules {
	return model.FeeRules{
		Id:               d.Id,
		FeeRuleVersionId: d.FeeRuleVersionId,
		RuleName:         d.RuleName,
		MinAmount:        decimal.NewNullDecimal(d.MinAmount),
		MaxAmount:        decimal.NewNullDecimal(d.MaxAmount),
		PercentageRate:   decimal.NewNullDecimal(d.PercentageRate),
		FlatAmount:       decimal.NewNullDecimal(d.FlatAmount),
		CapAmount:        decimal.NewNullDecimal(d.CapAmount),
		FloorAmount:      decimal.NewNullDecimal(d.FloorAmount),
		TaxInclusive:     d.TaxInclusive,
		SortOrder:        d.SortOrder,
		Metadata:         d.Metadata,
	}
}

type FeeRulesResponse struct {
	Id               uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FeeRuleVersionId uuid.UUID       `json:"feeRuleVersionId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RuleName         string          `json:"ruleName" validate:"required"`
	MinAmount        decimal.Decimal `json:"minAmount" format:"decimal" example:"100.50"`
	MaxAmount        decimal.Decimal `json:"maxAmount" format:"decimal" example:"100.50"`
	PercentageRate   decimal.Decimal `json:"percentageRate" format:"decimal" example:"100.50"`
	FlatAmount       decimal.Decimal `json:"flatAmount" format:"decimal" example:"100.50"`
	CapAmount        decimal.Decimal `json:"capAmount" format:"decimal" example:"100.50"`
	FloorAmount      decimal.Decimal `json:"floorAmount" format:"decimal" example:"100.50"`
	TaxInclusive     bool            `json:"taxInclusive" example:"true"`
	SortOrder        int             `json:"sortOrder" example:"1"`
	Metadata         json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewFeeRulesResponse(feeRules model.FeeRules) FeeRulesResponse {
	return FeeRulesResponse{
		Id:               feeRules.Id,
		FeeRuleVersionId: feeRules.FeeRuleVersionId,
		RuleName:         feeRules.RuleName,
		MinAmount:        feeRules.MinAmount.Decimal,
		MaxAmount:        feeRules.MaxAmount.Decimal,
		PercentageRate:   feeRules.PercentageRate.Decimal,
		FlatAmount:       feeRules.FlatAmount.Decimal,
		CapAmount:        feeRules.CapAmount.Decimal,
		FloorAmount:      feeRules.FloorAmount.Decimal,
		TaxInclusive:     feeRules.TaxInclusive,
		SortOrder:        feeRules.SortOrder,
		Metadata:         feeRules.Metadata,
	}
}

type FeeRulesListResponse []*FeeRulesResponse

func NewFeeRulesListResponse(feeRulesList model.FeeRulesList) FeeRulesListResponse {
	dtoFeeRulesListResponse := FeeRulesListResponse{}
	for _, feeRules := range feeRulesList {
		dtoFeeRulesResponse := NewFeeRulesResponse(*feeRules)
		dtoFeeRulesListResponse = append(dtoFeeRulesListResponse, &dtoFeeRulesResponse)
	}
	return dtoFeeRulesListResponse
}

type FeeRulesPrimaryIDList []FeeRulesPrimaryID

func (d FeeRulesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRules := range d {
		err = validator.Struct(feeRules)
		if err != nil {
			return
		}
	}
	return nil
}

type FeeRulesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FeeRulesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FeeRulesPrimaryID) ToModel() model.FeeRulesPrimaryID {
	return model.FeeRulesPrimaryID{
		Id: d.Id,
	}
}

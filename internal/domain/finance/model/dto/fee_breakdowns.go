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

type FeeBreakdownsDTOFieldNameType string

type feeBreakdownsDTOFieldName struct {
	Id               FeeBreakdownsDTOFieldNameType
	SourceType       FeeBreakdownsDTOFieldNameType
	SourceId         FeeBreakdownsDTOFieldNameType
	FeeProfileId     FeeBreakdownsDTOFieldNameType
	FeeRuleId        FeeBreakdownsDTOFieldNameType
	CurrencyCode     FeeBreakdownsDTOFieldNameType
	BaseAmount       FeeBreakdownsDTOFieldNameType
	FeeAmount        FeeBreakdownsDTOFieldNameType
	TaxAmount        FeeBreakdownsDTOFieldNameType
	RecipientPartyId FeeBreakdownsDTOFieldNameType
	BreakdownStatus  FeeBreakdownsDTOFieldNameType
	Metadata         FeeBreakdownsDTOFieldNameType
	MetaCreatedAt    FeeBreakdownsDTOFieldNameType
	MetaCreatedBy    FeeBreakdownsDTOFieldNameType
	MetaUpdatedAt    FeeBreakdownsDTOFieldNameType
	MetaUpdatedBy    FeeBreakdownsDTOFieldNameType
	MetaDeletedAt    FeeBreakdownsDTOFieldNameType
	MetaDeletedBy    FeeBreakdownsDTOFieldNameType
}

var FeeBreakdownsDTOFieldName = feeBreakdownsDTOFieldName{
	Id:               "id",
	SourceType:       "sourceType",
	SourceId:         "sourceId",
	FeeProfileId:     "feeProfileId",
	FeeRuleId:        "feeRuleId",
	CurrencyCode:     "currencyCode",
	BaseAmount:       "baseAmount",
	FeeAmount:        "feeAmount",
	TaxAmount:        "taxAmount",
	RecipientPartyId: "recipientPartyId",
	BreakdownStatus:  "breakdownStatus",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformFeeBreakdownsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FeeBreakdownsDTOFieldName.Id):
		return string(model.FeeBreakdownsDBFieldName.Id), true

	case string(FeeBreakdownsDTOFieldName.SourceType):
		return string(model.FeeBreakdownsDBFieldName.SourceType), true

	case string(FeeBreakdownsDTOFieldName.SourceId):
		return string(model.FeeBreakdownsDBFieldName.SourceId), true

	case string(FeeBreakdownsDTOFieldName.FeeProfileId):
		return string(model.FeeBreakdownsDBFieldName.FeeProfileId), true

	case string(FeeBreakdownsDTOFieldName.FeeRuleId):
		return string(model.FeeBreakdownsDBFieldName.FeeRuleId), true

	case string(FeeBreakdownsDTOFieldName.CurrencyCode):
		return string(model.FeeBreakdownsDBFieldName.CurrencyCode), true

	case string(FeeBreakdownsDTOFieldName.BaseAmount):
		return string(model.FeeBreakdownsDBFieldName.BaseAmount), true

	case string(FeeBreakdownsDTOFieldName.FeeAmount):
		return string(model.FeeBreakdownsDBFieldName.FeeAmount), true

	case string(FeeBreakdownsDTOFieldName.TaxAmount):
		return string(model.FeeBreakdownsDBFieldName.TaxAmount), true

	case string(FeeBreakdownsDTOFieldName.RecipientPartyId):
		return string(model.FeeBreakdownsDBFieldName.RecipientPartyId), true

	case string(FeeBreakdownsDTOFieldName.BreakdownStatus):
		return string(model.FeeBreakdownsDBFieldName.BreakdownStatus), true

	case string(FeeBreakdownsDTOFieldName.Metadata):
		return string(model.FeeBreakdownsDBFieldName.Metadata), true

	case string(FeeBreakdownsDTOFieldName.MetaCreatedAt):
		return string(model.FeeBreakdownsDBFieldName.MetaCreatedAt), true

	case string(FeeBreakdownsDTOFieldName.MetaCreatedBy):
		return string(model.FeeBreakdownsDBFieldName.MetaCreatedBy), true

	case string(FeeBreakdownsDTOFieldName.MetaUpdatedAt):
		return string(model.FeeBreakdownsDBFieldName.MetaUpdatedAt), true

	case string(FeeBreakdownsDTOFieldName.MetaUpdatedBy):
		return string(model.FeeBreakdownsDBFieldName.MetaUpdatedBy), true

	case string(FeeBreakdownsDTOFieldName.MetaDeletedAt):
		return string(model.FeeBreakdownsDBFieldName.MetaDeletedAt), true

	case string(FeeBreakdownsDTOFieldName.MetaDeletedBy):
		return string(model.FeeBreakdownsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFeeBreakdownsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFeeBreakdownsBaseFilterField(field string) bool {
	spec, found := model.NewFeeBreakdownsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFeeBreakdownsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFeeBreakdownsProjectionOutputPath(path string) error {
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

func transformFeeBreakdownsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFeeBreakdownsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFeeBreakdownsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFeeBreakdownsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFeeBreakdownsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFeeBreakdownsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFeeBreakdownsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFeeBreakdownsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFeeBreakdownsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFeeBreakdownsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFeeBreakdownsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFeeBreakdownsFilter(filter *model.Filter) {
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
			Field: string(FeeBreakdownsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FeeBreakdownsSelectableResponse map[string]interface{}
type FeeBreakdownsSelectableListResponse []*FeeBreakdownsSelectableResponse

func assignFeeBreakdownsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFeeBreakdownsSelectableValue(out FeeBreakdownsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFeeBreakdownsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFeeBreakdownsSelectableResponse(feeBreakdowns model.FeeBreakdowns, filter model.Filter) FeeBreakdownsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FeeBreakdownsDBFieldName.Id),
			string(model.FeeBreakdownsDBFieldName.SourceType),
			string(model.FeeBreakdownsDBFieldName.SourceId),
			string(model.FeeBreakdownsDBFieldName.FeeProfileId),
			string(model.FeeBreakdownsDBFieldName.FeeRuleId),
			string(model.FeeBreakdownsDBFieldName.CurrencyCode),
			string(model.FeeBreakdownsDBFieldName.BaseAmount),
			string(model.FeeBreakdownsDBFieldName.FeeAmount),
			string(model.FeeBreakdownsDBFieldName.TaxAmount),
			string(model.FeeBreakdownsDBFieldName.RecipientPartyId),
			string(model.FeeBreakdownsDBFieldName.BreakdownStatus),
			string(model.FeeBreakdownsDBFieldName.Metadata),
			string(model.FeeBreakdownsDBFieldName.MetaCreatedAt),
			string(model.FeeBreakdownsDBFieldName.MetaCreatedBy),
			string(model.FeeBreakdownsDBFieldName.MetaUpdatedAt),
			string(model.FeeBreakdownsDBFieldName.MetaUpdatedBy),
			string(model.FeeBreakdownsDBFieldName.MetaDeletedAt),
			string(model.FeeBreakdownsDBFieldName.MetaDeletedBy),
		)
	}
	feeBreakdownsSelectableResponse := FeeBreakdownsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FeeBreakdownsDBFieldName.Id):
			key := string(FeeBreakdownsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.Id, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.SourceType):
			key := string(FeeBreakdownsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.SourceType, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.SourceId):
			key := string(FeeBreakdownsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.SourceId, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.FeeProfileId):
			key := string(FeeBreakdownsDTOFieldName.FeeProfileId)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.FeeProfileId, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.FeeRuleId):
			key := string(FeeBreakdownsDTOFieldName.FeeRuleId)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.FeeRuleId, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.CurrencyCode):
			key := string(FeeBreakdownsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.CurrencyCode, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.BaseAmount):
			key := string(FeeBreakdownsDTOFieldName.BaseAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.BaseAmount, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.FeeAmount):
			key := string(FeeBreakdownsDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.FeeAmount, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.TaxAmount):
			key := string(FeeBreakdownsDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.TaxAmount, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.RecipientPartyId):
			key := string(FeeBreakdownsDTOFieldName.RecipientPartyId)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.RecipientPartyId.UUID, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.BreakdownStatus):
			key := string(FeeBreakdownsDTOFieldName.BreakdownStatus)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, model.FeeBreakdownsBreakdownStatus(feeBreakdowns.BreakdownStatus), explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.Metadata):
			key := string(FeeBreakdownsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.Metadata, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaCreatedAt):
			key := string(FeeBreakdownsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaCreatedAt, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaCreatedBy):
			key := string(FeeBreakdownsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaCreatedBy, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaUpdatedAt):
			key := string(FeeBreakdownsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaUpdatedAt, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaUpdatedBy):
			key := string(FeeBreakdownsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaUpdatedBy, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaDeletedAt):
			key := string(FeeBreakdownsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaDeletedAt.Time, explicitAlias)

		case string(model.FeeBreakdownsDBFieldName.MetaDeletedBy):
			key := string(FeeBreakdownsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeBreakdownsSelectableValue(feeBreakdownsSelectableResponse, key, feeBreakdowns.MetaDeletedBy, explicitAlias)

		}
	}
	return feeBreakdownsSelectableResponse
}

func NewFeeBreakdownsListResponseFromFilterResult(result []model.FeeBreakdownsFilterResult, filter model.Filter) FeeBreakdownsSelectableListResponse {
	dtoFeeBreakdownsListResponse := FeeBreakdownsSelectableListResponse{}
	for _, row := range result {
		dtoFeeBreakdownsResponse := NewFeeBreakdownsSelectableResponse(row.FeeBreakdowns, filter)
		dtoFeeBreakdownsListResponse = append(dtoFeeBreakdownsListResponse, &dtoFeeBreakdownsResponse)
	}
	return dtoFeeBreakdownsListResponse
}

type FeeBreakdownsFilterResponse struct {
	Metadata Metadata                            `json:"metadata"`
	Data     FeeBreakdownsSelectableListResponse `json:"data"`
}

func reverseFeeBreakdownsFilterResults(result []model.FeeBreakdownsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFeeBreakdownsFilterResponse(result []model.FeeBreakdownsFilterResult, filter model.Filter) (resp FeeBreakdownsFilterResponse) {
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
			reverseFeeBreakdownsFilterResults(dataResult)
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

	resp.Data = NewFeeBreakdownsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FeeBreakdownsCreateRequest struct {
	SourceType       string                             `json:"sourceType"`
	SourceId         uuid.UUID                          `json:"sourceId"`
	FeeProfileId     uuid.UUID                          `json:"feeProfileId"`
	FeeRuleId        uuid.UUID                          `json:"feeRuleId"`
	CurrencyCode     string                             `json:"currencyCode"`
	BaseAmount       decimal.Decimal                    `json:"baseAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	RecipientPartyId uuid.UUID                          `json:"recipientPartyId"`
	BreakdownStatus  model.FeeBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d *FeeBreakdownsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FeeBreakdownsCreateRequest) ToModel() model.FeeBreakdowns {
	id, _ := uuid.NewV4()
	return model.FeeBreakdowns{
		Id:               id,
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		FeeProfileId:     d.FeeProfileId,
		FeeRuleId:        d.FeeRuleId,
		CurrencyCode:     d.CurrencyCode,
		BaseAmount:       d.BaseAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		RecipientPartyId: nuuid.From(d.RecipientPartyId),
		BreakdownStatus:  d.BreakdownStatus,
		Metadata:         d.Metadata,
	}
}

type FeeBreakdownsListCreateRequest []*FeeBreakdownsCreateRequest

func (d FeeBreakdownsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeBreakdowns := range d {
		err = validator.Struct(feeBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeBreakdownsListCreateRequest) ToModelList() []model.FeeBreakdowns {
	out := make([]model.FeeBreakdowns, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FeeBreakdownsUpdateRequest struct {
	SourceType       string                             `json:"sourceType"`
	SourceId         uuid.UUID                          `json:"sourceId"`
	FeeProfileId     uuid.UUID                          `json:"feeProfileId"`
	FeeRuleId        uuid.UUID                          `json:"feeRuleId"`
	CurrencyCode     string                             `json:"currencyCode"`
	BaseAmount       decimal.Decimal                    `json:"baseAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	RecipientPartyId uuid.UUID                          `json:"recipientPartyId"`
	BreakdownStatus  model.FeeBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d *FeeBreakdownsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FeeBreakdownsUpdateRequest) ToModel() model.FeeBreakdowns {
	return model.FeeBreakdowns{
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		FeeProfileId:     d.FeeProfileId,
		FeeRuleId:        d.FeeRuleId,
		CurrencyCode:     d.CurrencyCode,
		BaseAmount:       d.BaseAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		RecipientPartyId: nuuid.From(d.RecipientPartyId),
		BreakdownStatus:  d.BreakdownStatus,
		Metadata:         d.Metadata,
	}
}

type FeeBreakdownsBulkUpdateRequest struct {
	Id               uuid.UUID                          `json:"id"`
	SourceType       string                             `json:"sourceType"`
	SourceId         uuid.UUID                          `json:"sourceId"`
	FeeProfileId     uuid.UUID                          `json:"feeProfileId"`
	FeeRuleId        uuid.UUID                          `json:"feeRuleId"`
	CurrencyCode     string                             `json:"currencyCode"`
	BaseAmount       decimal.Decimal                    `json:"baseAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	RecipientPartyId uuid.UUID                          `json:"recipientPartyId"`
	BreakdownStatus  model.FeeBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,reversed"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d FeeBreakdownsBulkUpdateRequest) PrimaryID() FeeBreakdownsPrimaryID {
	return FeeBreakdownsPrimaryID{
		Id: d.Id,
	}
}

type FeeBreakdownsListBulkUpdateRequest []*FeeBreakdownsBulkUpdateRequest

func (d FeeBreakdownsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeBreakdowns := range d {
		err = validator.Struct(feeBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeBreakdownsBulkUpdateRequest) ToModel() model.FeeBreakdowns {
	return model.FeeBreakdowns{
		Id:               d.Id,
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		FeeProfileId:     d.FeeProfileId,
		FeeRuleId:        d.FeeRuleId,
		CurrencyCode:     d.CurrencyCode,
		BaseAmount:       d.BaseAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		RecipientPartyId: nuuid.From(d.RecipientPartyId),
		BreakdownStatus:  d.BreakdownStatus,
		Metadata:         d.Metadata,
	}
}

type FeeBreakdownsResponse struct {
	Id               uuid.UUID                          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType       string                             `json:"sourceType" validate:"required"`
	SourceId         uuid.UUID                          `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FeeProfileId     uuid.UUID                          `json:"feeProfileId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FeeRuleId        uuid.UUID                          `json:"feeRuleId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode     string                             `json:"currencyCode" validate:"required"`
	BaseAmount       decimal.Decimal                    `json:"baseAmount" validate:"required" format:"decimal" example:"100.50"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount" validate:"required" format:"decimal" example:"100.50"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount" format:"decimal" example:"100.50"`
	RecipientPartyId uuid.UUID                          `json:"recipientPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BreakdownStatus  model.FeeBreakdownsBreakdownStatus `json:"breakdownStatus" validate:"oneof=computed posted reversed" enums:"computed,posted,reversed"`
	Metadata         json.RawMessage                    `json:"metadata" swaggertype:"object"`
}

func NewFeeBreakdownsResponse(feeBreakdowns model.FeeBreakdowns) FeeBreakdownsResponse {
	return FeeBreakdownsResponse{
		Id:               feeBreakdowns.Id,
		SourceType:       feeBreakdowns.SourceType,
		SourceId:         feeBreakdowns.SourceId,
		FeeProfileId:     feeBreakdowns.FeeProfileId,
		FeeRuleId:        feeBreakdowns.FeeRuleId,
		CurrencyCode:     feeBreakdowns.CurrencyCode,
		BaseAmount:       feeBreakdowns.BaseAmount,
		FeeAmount:        feeBreakdowns.FeeAmount,
		TaxAmount:        feeBreakdowns.TaxAmount,
		RecipientPartyId: feeBreakdowns.RecipientPartyId.UUID,
		BreakdownStatus:  model.FeeBreakdownsBreakdownStatus(feeBreakdowns.BreakdownStatus),
		Metadata:         feeBreakdowns.Metadata,
	}
}

type FeeBreakdownsListResponse []*FeeBreakdownsResponse

func NewFeeBreakdownsListResponse(feeBreakdownsList model.FeeBreakdownsList) FeeBreakdownsListResponse {
	dtoFeeBreakdownsListResponse := FeeBreakdownsListResponse{}
	for _, feeBreakdowns := range feeBreakdownsList {
		dtoFeeBreakdownsResponse := NewFeeBreakdownsResponse(*feeBreakdowns)
		dtoFeeBreakdownsListResponse = append(dtoFeeBreakdownsListResponse, &dtoFeeBreakdownsResponse)
	}
	return dtoFeeBreakdownsListResponse
}

type FeeBreakdownsPrimaryIDList []FeeBreakdownsPrimaryID

func (d FeeBreakdownsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeBreakdowns := range d {
		err = validator.Struct(feeBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

type FeeBreakdownsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FeeBreakdownsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FeeBreakdownsPrimaryID) ToModel() model.FeeBreakdownsPrimaryID {
	return model.FeeBreakdownsPrimaryID{
		Id: d.Id,
	}
}

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

type TaxBreakdownsDTOFieldNameType string

type taxBreakdownsDTOFieldName struct {
	Id                 TaxBreakdownsDTOFieldNameType
	SourceType         TaxBreakdownsDTOFieldNameType
	SourceId           TaxBreakdownsDTOFieldNameType
	TaxRuleId          TaxBreakdownsDTOFieldNameType
	CurrencyCode       TaxBreakdownsDTOFieldNameType
	TaxableAmount      TaxBreakdownsDTOFieldNameType
	TaxAmount          TaxBreakdownsDTOFieldNameType
	LiabilityPartyId   TaxBreakdownsDTOFieldNameType
	BeneficiaryPartyId TaxBreakdownsDTOFieldNameType
	BreakdownStatus    TaxBreakdownsDTOFieldNameType
	Metadata           TaxBreakdownsDTOFieldNameType
	MetaCreatedAt      TaxBreakdownsDTOFieldNameType
	MetaCreatedBy      TaxBreakdownsDTOFieldNameType
	MetaUpdatedAt      TaxBreakdownsDTOFieldNameType
	MetaUpdatedBy      TaxBreakdownsDTOFieldNameType
	MetaDeletedAt      TaxBreakdownsDTOFieldNameType
	MetaDeletedBy      TaxBreakdownsDTOFieldNameType
}

var TaxBreakdownsDTOFieldName = taxBreakdownsDTOFieldName{
	Id:                 "id",
	SourceType:         "sourceType",
	SourceId:           "sourceId",
	TaxRuleId:          "taxRuleId",
	CurrencyCode:       "currencyCode",
	TaxableAmount:      "taxableAmount",
	TaxAmount:          "taxAmount",
	LiabilityPartyId:   "liabilityPartyId",
	BeneficiaryPartyId: "beneficiaryPartyId",
	BreakdownStatus:    "breakdownStatus",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformTaxBreakdownsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxBreakdownsDTOFieldName.Id):
		return string(model.TaxBreakdownsDBFieldName.Id), true

	case string(TaxBreakdownsDTOFieldName.SourceType):
		return string(model.TaxBreakdownsDBFieldName.SourceType), true

	case string(TaxBreakdownsDTOFieldName.SourceId):
		return string(model.TaxBreakdownsDBFieldName.SourceId), true

	case string(TaxBreakdownsDTOFieldName.TaxRuleId):
		return string(model.TaxBreakdownsDBFieldName.TaxRuleId), true

	case string(TaxBreakdownsDTOFieldName.CurrencyCode):
		return string(model.TaxBreakdownsDBFieldName.CurrencyCode), true

	case string(TaxBreakdownsDTOFieldName.TaxableAmount):
		return string(model.TaxBreakdownsDBFieldName.TaxableAmount), true

	case string(TaxBreakdownsDTOFieldName.TaxAmount):
		return string(model.TaxBreakdownsDBFieldName.TaxAmount), true

	case string(TaxBreakdownsDTOFieldName.LiabilityPartyId):
		return string(model.TaxBreakdownsDBFieldName.LiabilityPartyId), true

	case string(TaxBreakdownsDTOFieldName.BeneficiaryPartyId):
		return string(model.TaxBreakdownsDBFieldName.BeneficiaryPartyId), true

	case string(TaxBreakdownsDTOFieldName.BreakdownStatus):
		return string(model.TaxBreakdownsDBFieldName.BreakdownStatus), true

	case string(TaxBreakdownsDTOFieldName.Metadata):
		return string(model.TaxBreakdownsDBFieldName.Metadata), true

	case string(TaxBreakdownsDTOFieldName.MetaCreatedAt):
		return string(model.TaxBreakdownsDBFieldName.MetaCreatedAt), true

	case string(TaxBreakdownsDTOFieldName.MetaCreatedBy):
		return string(model.TaxBreakdownsDBFieldName.MetaCreatedBy), true

	case string(TaxBreakdownsDTOFieldName.MetaUpdatedAt):
		return string(model.TaxBreakdownsDBFieldName.MetaUpdatedAt), true

	case string(TaxBreakdownsDTOFieldName.MetaUpdatedBy):
		return string(model.TaxBreakdownsDBFieldName.MetaUpdatedBy), true

	case string(TaxBreakdownsDTOFieldName.MetaDeletedAt):
		return string(model.TaxBreakdownsDBFieldName.MetaDeletedAt), true

	case string(TaxBreakdownsDTOFieldName.MetaDeletedBy):
		return string(model.TaxBreakdownsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxBreakdownsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxBreakdownsBaseFilterField(field string) bool {
	spec, found := model.NewTaxBreakdownsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxBreakdownsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxBreakdownsProjectionOutputPath(path string) error {
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

func transformTaxBreakdownsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxBreakdownsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxBreakdownsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxBreakdownsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxBreakdownsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxBreakdownsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxBreakdownsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxBreakdownsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxBreakdownsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxBreakdownsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxBreakdownsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxBreakdownsFilter(filter *model.Filter) {
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
			Field: string(TaxBreakdownsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type TaxBreakdownsSelectableResponse map[string]interface{}
type TaxBreakdownsSelectableListResponse []*TaxBreakdownsSelectableResponse

func assignTaxBreakdownsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxBreakdownsSelectableValue(out TaxBreakdownsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxBreakdownsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxBreakdownsSelectableResponse(taxBreakdowns model.TaxBreakdowns, filter model.Filter) TaxBreakdownsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxBreakdownsDBFieldName.Id),
			string(model.TaxBreakdownsDBFieldName.SourceType),
			string(model.TaxBreakdownsDBFieldName.SourceId),
			string(model.TaxBreakdownsDBFieldName.TaxRuleId),
			string(model.TaxBreakdownsDBFieldName.CurrencyCode),
			string(model.TaxBreakdownsDBFieldName.TaxableAmount),
			string(model.TaxBreakdownsDBFieldName.TaxAmount),
			string(model.TaxBreakdownsDBFieldName.LiabilityPartyId),
			string(model.TaxBreakdownsDBFieldName.BeneficiaryPartyId),
			string(model.TaxBreakdownsDBFieldName.BreakdownStatus),
			string(model.TaxBreakdownsDBFieldName.Metadata),
			string(model.TaxBreakdownsDBFieldName.MetaCreatedAt),
			string(model.TaxBreakdownsDBFieldName.MetaCreatedBy),
			string(model.TaxBreakdownsDBFieldName.MetaUpdatedAt),
			string(model.TaxBreakdownsDBFieldName.MetaUpdatedBy),
			string(model.TaxBreakdownsDBFieldName.MetaDeletedAt),
			string(model.TaxBreakdownsDBFieldName.MetaDeletedBy),
		)
	}
	taxBreakdownsSelectableResponse := TaxBreakdownsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxBreakdownsDBFieldName.Id):
			key := string(TaxBreakdownsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.Id, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.SourceType):
			key := string(TaxBreakdownsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.SourceType, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.SourceId):
			key := string(TaxBreakdownsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.SourceId, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.TaxRuleId):
			key := string(TaxBreakdownsDTOFieldName.TaxRuleId)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.TaxRuleId.UUID, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.CurrencyCode):
			key := string(TaxBreakdownsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.CurrencyCode, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.TaxableAmount):
			key := string(TaxBreakdownsDTOFieldName.TaxableAmount)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.TaxableAmount, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.TaxAmount):
			key := string(TaxBreakdownsDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.TaxAmount, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.LiabilityPartyId):
			key := string(TaxBreakdownsDTOFieldName.LiabilityPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.LiabilityPartyId.UUID, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.BeneficiaryPartyId):
			key := string(TaxBreakdownsDTOFieldName.BeneficiaryPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.BeneficiaryPartyId.UUID, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.BreakdownStatus):
			key := string(TaxBreakdownsDTOFieldName.BreakdownStatus)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, model.TaxBreakdownsBreakdownStatus(taxBreakdowns.BreakdownStatus), explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.Metadata):
			key := string(TaxBreakdownsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.Metadata, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaCreatedAt):
			key := string(TaxBreakdownsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaCreatedAt, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaCreatedBy):
			key := string(TaxBreakdownsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaCreatedBy, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaUpdatedAt):
			key := string(TaxBreakdownsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaUpdatedAt, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaUpdatedBy):
			key := string(TaxBreakdownsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaUpdatedBy, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaDeletedAt):
			key := string(TaxBreakdownsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxBreakdownsDBFieldName.MetaDeletedBy):
			key := string(TaxBreakdownsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxBreakdownsSelectableValue(taxBreakdownsSelectableResponse, key, taxBreakdowns.MetaDeletedBy, explicitAlias)

		}
	}
	return taxBreakdownsSelectableResponse
}

func NewTaxBreakdownsListResponseFromFilterResult(result []model.TaxBreakdownsFilterResult, filter model.Filter) TaxBreakdownsSelectableListResponse {
	dtoTaxBreakdownsListResponse := TaxBreakdownsSelectableListResponse{}
	for _, row := range result {
		dtoTaxBreakdownsResponse := NewTaxBreakdownsSelectableResponse(row.TaxBreakdowns, filter)
		dtoTaxBreakdownsListResponse = append(dtoTaxBreakdownsListResponse, &dtoTaxBreakdownsResponse)
	}
	return dtoTaxBreakdownsListResponse
}

type TaxBreakdownsFilterResponse struct {
	Metadata Metadata                            `json:"metadata"`
	Data     TaxBreakdownsSelectableListResponse `json:"data"`
}

func reverseTaxBreakdownsFilterResults(result []model.TaxBreakdownsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxBreakdownsFilterResponse(result []model.TaxBreakdownsFilterResult, filter model.Filter) (resp TaxBreakdownsFilterResponse) {
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
			reverseTaxBreakdownsFilterResults(dataResult)
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

	resp.Data = NewTaxBreakdownsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxBreakdownsCreateRequest struct {
	SourceType         string                             `json:"sourceType"`
	SourceId           uuid.UUID                          `json:"sourceId"`
	TaxRuleId          uuid.UUID                          `json:"taxRuleId"`
	CurrencyCode       string                             `json:"currencyCode"`
	TaxableAmount      decimal.Decimal                    `json:"taxableAmount"`
	TaxAmount          decimal.Decimal                    `json:"taxAmount"`
	LiabilityPartyId   uuid.UUID                          `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID                          `json:"beneficiaryPartyId"`
	BreakdownStatus    model.TaxBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,void"`
	Metadata           json.RawMessage                    `json:"metadata"`
}

func (d *TaxBreakdownsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxBreakdownsCreateRequest) ToModel() model.TaxBreakdowns {
	id, _ := uuid.NewV4()
	return model.TaxBreakdowns{
		Id:                 id,
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		TaxRuleId:          nuuid.From(d.TaxRuleId),
		CurrencyCode:       d.CurrencyCode,
		TaxableAmount:      d.TaxableAmount,
		TaxAmount:          d.TaxAmount,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		BreakdownStatus:    d.BreakdownStatus,
		Metadata:           d.Metadata,
	}
}

type TaxBreakdownsListCreateRequest []*TaxBreakdownsCreateRequest

func (d TaxBreakdownsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxBreakdowns := range d {
		err = validator.Struct(taxBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxBreakdownsListCreateRequest) ToModelList() []model.TaxBreakdowns {
	out := make([]model.TaxBreakdowns, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxBreakdownsUpdateRequest struct {
	SourceType         string                             `json:"sourceType"`
	SourceId           uuid.UUID                          `json:"sourceId"`
	TaxRuleId          uuid.UUID                          `json:"taxRuleId"`
	CurrencyCode       string                             `json:"currencyCode"`
	TaxableAmount      decimal.Decimal                    `json:"taxableAmount"`
	TaxAmount          decimal.Decimal                    `json:"taxAmount"`
	LiabilityPartyId   uuid.UUID                          `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID                          `json:"beneficiaryPartyId"`
	BreakdownStatus    model.TaxBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,void"`
	Metadata           json.RawMessage                    `json:"metadata"`
}

func (d *TaxBreakdownsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxBreakdownsUpdateRequest) ToModel() model.TaxBreakdowns {
	return model.TaxBreakdowns{
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		TaxRuleId:          nuuid.From(d.TaxRuleId),
		CurrencyCode:       d.CurrencyCode,
		TaxableAmount:      d.TaxableAmount,
		TaxAmount:          d.TaxAmount,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		BreakdownStatus:    d.BreakdownStatus,
		Metadata:           d.Metadata,
	}
}

type TaxBreakdownsBulkUpdateRequest struct {
	Id                 uuid.UUID                          `json:"id"`
	SourceType         string                             `json:"sourceType"`
	SourceId           uuid.UUID                          `json:"sourceId"`
	TaxRuleId          uuid.UUID                          `json:"taxRuleId"`
	CurrencyCode       string                             `json:"currencyCode"`
	TaxableAmount      decimal.Decimal                    `json:"taxableAmount"`
	TaxAmount          decimal.Decimal                    `json:"taxAmount"`
	LiabilityPartyId   uuid.UUID                          `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID                          `json:"beneficiaryPartyId"`
	BreakdownStatus    model.TaxBreakdownsBreakdownStatus `json:"breakdownStatus" example:"computed" enums:"computed,posted,void"`
	Metadata           json.RawMessage                    `json:"metadata"`
}

func (d TaxBreakdownsBulkUpdateRequest) PrimaryID() TaxBreakdownsPrimaryID {
	return TaxBreakdownsPrimaryID{
		Id: d.Id,
	}
}

type TaxBreakdownsListBulkUpdateRequest []*TaxBreakdownsBulkUpdateRequest

func (d TaxBreakdownsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxBreakdowns := range d {
		err = validator.Struct(taxBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxBreakdownsBulkUpdateRequest) ToModel() model.TaxBreakdowns {
	return model.TaxBreakdowns{
		Id:                 d.Id,
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		TaxRuleId:          nuuid.From(d.TaxRuleId),
		CurrencyCode:       d.CurrencyCode,
		TaxableAmount:      d.TaxableAmount,
		TaxAmount:          d.TaxAmount,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		BreakdownStatus:    d.BreakdownStatus,
		Metadata:           d.Metadata,
	}
}

type TaxBreakdownsResponse struct {
	Id                 uuid.UUID                          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType         string                             `json:"sourceType" validate:"required"`
	SourceId           uuid.UUID                          `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	TaxRuleId          uuid.UUID                          `json:"taxRuleId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode       string                             `json:"currencyCode" validate:"required"`
	TaxableAmount      decimal.Decimal                    `json:"taxableAmount" validate:"required" format:"decimal" example:"100.50"`
	TaxAmount          decimal.Decimal                    `json:"taxAmount" validate:"required" format:"decimal" example:"100.50"`
	LiabilityPartyId   uuid.UUID                          `json:"liabilityPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BeneficiaryPartyId uuid.UUID                          `json:"beneficiaryPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BreakdownStatus    model.TaxBreakdownsBreakdownStatus `json:"breakdownStatus" validate:"oneof=computed posted void" enums:"computed,posted,void"`
	Metadata           json.RawMessage                    `json:"metadata" swaggertype:"object"`
}

func NewTaxBreakdownsResponse(taxBreakdowns model.TaxBreakdowns) TaxBreakdownsResponse {
	return TaxBreakdownsResponse{
		Id:                 taxBreakdowns.Id,
		SourceType:         taxBreakdowns.SourceType,
		SourceId:           taxBreakdowns.SourceId,
		TaxRuleId:          taxBreakdowns.TaxRuleId.UUID,
		CurrencyCode:       taxBreakdowns.CurrencyCode,
		TaxableAmount:      taxBreakdowns.TaxableAmount,
		TaxAmount:          taxBreakdowns.TaxAmount,
		LiabilityPartyId:   taxBreakdowns.LiabilityPartyId.UUID,
		BeneficiaryPartyId: taxBreakdowns.BeneficiaryPartyId.UUID,
		BreakdownStatus:    model.TaxBreakdownsBreakdownStatus(taxBreakdowns.BreakdownStatus),
		Metadata:           taxBreakdowns.Metadata,
	}
}

type TaxBreakdownsListResponse []*TaxBreakdownsResponse

func NewTaxBreakdownsListResponse(taxBreakdownsList model.TaxBreakdownsList) TaxBreakdownsListResponse {
	dtoTaxBreakdownsListResponse := TaxBreakdownsListResponse{}
	for _, taxBreakdowns := range taxBreakdownsList {
		dtoTaxBreakdownsResponse := NewTaxBreakdownsResponse(*taxBreakdowns)
		dtoTaxBreakdownsListResponse = append(dtoTaxBreakdownsListResponse, &dtoTaxBreakdownsResponse)
	}
	return dtoTaxBreakdownsListResponse
}

type TaxBreakdownsPrimaryIDList []TaxBreakdownsPrimaryID

func (d TaxBreakdownsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxBreakdowns := range d {
		err = validator.Struct(taxBreakdowns)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxBreakdownsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *TaxBreakdownsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxBreakdownsPrimaryID) ToModel() model.TaxBreakdownsPrimaryID {
	return model.TaxBreakdownsPrimaryID{
		Id: d.Id,
	}
}

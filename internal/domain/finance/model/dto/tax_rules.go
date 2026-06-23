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

type TaxRulesDTOFieldNameType string

type taxRulesDTOFieldName struct {
	Id                 TaxRulesDTOFieldNameType
	RuleCode           TaxRulesDTOFieldNameType
	CountryCode        TaxRulesDTOFieldNameType
	TaxType            TaxRulesDTOFieldNameType
	Rate               TaxRulesDTOFieldNameType
	TaxInclusive       TaxRulesDTOFieldNameType
	LiabilityPartyId   TaxRulesDTOFieldNameType
	BeneficiaryPartyId TaxRulesDTOFieldNameType
	Priority           TaxRulesDTOFieldNameType
	IsActive           TaxRulesDTOFieldNameType
	ValidFrom          TaxRulesDTOFieldNameType
	ValidUntil         TaxRulesDTOFieldNameType
	Metadata           TaxRulesDTOFieldNameType
	MetaCreatedAt      TaxRulesDTOFieldNameType
	MetaCreatedBy      TaxRulesDTOFieldNameType
	MetaUpdatedAt      TaxRulesDTOFieldNameType
	MetaUpdatedBy      TaxRulesDTOFieldNameType
	MetaDeletedAt      TaxRulesDTOFieldNameType
	MetaDeletedBy      TaxRulesDTOFieldNameType
}

var TaxRulesDTOFieldName = taxRulesDTOFieldName{
	Id:                 "id",
	RuleCode:           "ruleCode",
	CountryCode:        "countryCode",
	TaxType:            "taxType",
	Rate:               "rate",
	TaxInclusive:       "taxInclusive",
	LiabilityPartyId:   "liabilityPartyId",
	BeneficiaryPartyId: "beneficiaryPartyId",
	Priority:           "priority",
	IsActive:           "isActive",
	ValidFrom:          "validFrom",
	ValidUntil:         "validUntil",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformTaxRulesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxRulesDTOFieldName.Id):
		return string(model.TaxRulesDBFieldName.Id), true

	case string(TaxRulesDTOFieldName.RuleCode):
		return string(model.TaxRulesDBFieldName.RuleCode), true

	case string(TaxRulesDTOFieldName.CountryCode):
		return string(model.TaxRulesDBFieldName.CountryCode), true

	case string(TaxRulesDTOFieldName.TaxType):
		return string(model.TaxRulesDBFieldName.TaxType), true

	case string(TaxRulesDTOFieldName.Rate):
		return string(model.TaxRulesDBFieldName.Rate), true

	case string(TaxRulesDTOFieldName.TaxInclusive):
		return string(model.TaxRulesDBFieldName.TaxInclusive), true

	case string(TaxRulesDTOFieldName.LiabilityPartyId):
		return string(model.TaxRulesDBFieldName.LiabilityPartyId), true

	case string(TaxRulesDTOFieldName.BeneficiaryPartyId):
		return string(model.TaxRulesDBFieldName.BeneficiaryPartyId), true

	case string(TaxRulesDTOFieldName.Priority):
		return string(model.TaxRulesDBFieldName.Priority), true

	case string(TaxRulesDTOFieldName.IsActive):
		return string(model.TaxRulesDBFieldName.IsActive), true

	case string(TaxRulesDTOFieldName.ValidFrom):
		return string(model.TaxRulesDBFieldName.ValidFrom), true

	case string(TaxRulesDTOFieldName.ValidUntil):
		return string(model.TaxRulesDBFieldName.ValidUntil), true

	case string(TaxRulesDTOFieldName.Metadata):
		return string(model.TaxRulesDBFieldName.Metadata), true

	case string(TaxRulesDTOFieldName.MetaCreatedAt):
		return string(model.TaxRulesDBFieldName.MetaCreatedAt), true

	case string(TaxRulesDTOFieldName.MetaCreatedBy):
		return string(model.TaxRulesDBFieldName.MetaCreatedBy), true

	case string(TaxRulesDTOFieldName.MetaUpdatedAt):
		return string(model.TaxRulesDBFieldName.MetaUpdatedAt), true

	case string(TaxRulesDTOFieldName.MetaUpdatedBy):
		return string(model.TaxRulesDBFieldName.MetaUpdatedBy), true

	case string(TaxRulesDTOFieldName.MetaDeletedAt):
		return string(model.TaxRulesDBFieldName.MetaDeletedAt), true

	case string(TaxRulesDTOFieldName.MetaDeletedBy):
		return string(model.TaxRulesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxRulesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxRulesBaseFilterField(field string) bool {
	spec, found := model.NewTaxRulesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxRulesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxRulesProjectionOutputPath(path string) error {
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

func transformTaxRulesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxRulesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxRulesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxRulesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxRulesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxRulesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxRulesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxRulesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxRulesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxRulesFilter(filter *model.Filter) {
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
			Field: string(TaxRulesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type TaxRulesSelectableResponse map[string]interface{}
type TaxRulesSelectableListResponse []*TaxRulesSelectableResponse

func assignTaxRulesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxRulesSelectableValue(out TaxRulesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxRulesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxRulesSelectableResponse(taxRules model.TaxRules, filter model.Filter) TaxRulesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxRulesDBFieldName.Id),
			string(model.TaxRulesDBFieldName.RuleCode),
			string(model.TaxRulesDBFieldName.CountryCode),
			string(model.TaxRulesDBFieldName.TaxType),
			string(model.TaxRulesDBFieldName.Rate),
			string(model.TaxRulesDBFieldName.TaxInclusive),
			string(model.TaxRulesDBFieldName.LiabilityPartyId),
			string(model.TaxRulesDBFieldName.BeneficiaryPartyId),
			string(model.TaxRulesDBFieldName.Priority),
			string(model.TaxRulesDBFieldName.IsActive),
			string(model.TaxRulesDBFieldName.ValidFrom),
			string(model.TaxRulesDBFieldName.ValidUntil),
			string(model.TaxRulesDBFieldName.Metadata),
			string(model.TaxRulesDBFieldName.MetaCreatedAt),
			string(model.TaxRulesDBFieldName.MetaCreatedBy),
			string(model.TaxRulesDBFieldName.MetaUpdatedAt),
			string(model.TaxRulesDBFieldName.MetaUpdatedBy),
			string(model.TaxRulesDBFieldName.MetaDeletedAt),
			string(model.TaxRulesDBFieldName.MetaDeletedBy),
		)
	}
	taxRulesSelectableResponse := TaxRulesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxRulesDBFieldName.Id):
			key := string(TaxRulesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.Id, explicitAlias)

		case string(model.TaxRulesDBFieldName.RuleCode):
			key := string(TaxRulesDTOFieldName.RuleCode)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.RuleCode, explicitAlias)

		case string(model.TaxRulesDBFieldName.CountryCode):
			key := string(TaxRulesDTOFieldName.CountryCode)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.CountryCode, explicitAlias)

		case string(model.TaxRulesDBFieldName.TaxType):
			key := string(TaxRulesDTOFieldName.TaxType)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.TaxType, explicitAlias)

		case string(model.TaxRulesDBFieldName.Rate):
			key := string(TaxRulesDTOFieldName.Rate)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.Rate, explicitAlias)

		case string(model.TaxRulesDBFieldName.TaxInclusive):
			key := string(TaxRulesDTOFieldName.TaxInclusive)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.TaxInclusive, explicitAlias)

		case string(model.TaxRulesDBFieldName.LiabilityPartyId):
			key := string(TaxRulesDTOFieldName.LiabilityPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.LiabilityPartyId.UUID, explicitAlias)

		case string(model.TaxRulesDBFieldName.BeneficiaryPartyId):
			key := string(TaxRulesDTOFieldName.BeneficiaryPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.BeneficiaryPartyId.UUID, explicitAlias)

		case string(model.TaxRulesDBFieldName.Priority):
			key := string(TaxRulesDTOFieldName.Priority)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.Priority, explicitAlias)

		case string(model.TaxRulesDBFieldName.IsActive):
			key := string(TaxRulesDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.IsActive, explicitAlias)

		case string(model.TaxRulesDBFieldName.ValidFrom):
			key := string(TaxRulesDTOFieldName.ValidFrom)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.ValidFrom, explicitAlias)

		case string(model.TaxRulesDBFieldName.ValidUntil):
			key := string(TaxRulesDTOFieldName.ValidUntil)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.ValidUntil.Time, explicitAlias)

		case string(model.TaxRulesDBFieldName.Metadata):
			key := string(TaxRulesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.Metadata, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaCreatedAt):
			key := string(TaxRulesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaCreatedAt, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaCreatedBy):
			key := string(TaxRulesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaCreatedBy, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaUpdatedAt):
			key := string(TaxRulesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaUpdatedAt, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaUpdatedBy):
			key := string(TaxRulesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaUpdatedBy, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaDeletedAt):
			key := string(TaxRulesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxRulesDBFieldName.MetaDeletedBy):
			key := string(TaxRulesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxRulesSelectableValue(taxRulesSelectableResponse, key, taxRules.MetaDeletedBy, explicitAlias)

		}
	}
	return taxRulesSelectableResponse
}

func NewTaxRulesListResponseFromFilterResult(result []model.TaxRulesFilterResult, filter model.Filter) TaxRulesSelectableListResponse {
	dtoTaxRulesListResponse := TaxRulesSelectableListResponse{}
	for _, row := range result {
		dtoTaxRulesResponse := NewTaxRulesSelectableResponse(row.TaxRules, filter)
		dtoTaxRulesListResponse = append(dtoTaxRulesListResponse, &dtoTaxRulesResponse)
	}
	return dtoTaxRulesListResponse
}

type TaxRulesFilterResponse struct {
	Metadata Metadata                       `json:"metadata"`
	Data     TaxRulesSelectableListResponse `json:"data"`
}

func reverseTaxRulesFilterResults(result []model.TaxRulesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxRulesFilterResponse(result []model.TaxRulesFilterResult, filter model.Filter) (resp TaxRulesFilterResponse) {
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
			reverseTaxRulesFilterResults(dataResult)
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

	resp.Data = NewTaxRulesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxRulesCreateRequest struct {
	RuleCode           string          `json:"ruleCode"`
	CountryCode        string          `json:"countryCode"`
	TaxType            string          `json:"taxType"`
	Rate               decimal.Decimal `json:"rate"`
	TaxInclusive       bool            `json:"taxInclusive"`
	LiabilityPartyId   uuid.UUID       `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID       `json:"beneficiaryPartyId"`
	Priority           int             `json:"priority"`
	IsActive           bool            `json:"isActive"`
	ValidFrom          time.Time       `json:"validFrom"`
	ValidUntil         time.Time       `json:"validUntil"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d *TaxRulesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxRulesCreateRequest) ToModel() model.TaxRules {
	id, _ := uuid.NewV4()
	return model.TaxRules{
		Id:                 id,
		RuleCode:           d.RuleCode,
		CountryCode:        d.CountryCode,
		TaxType:            d.TaxType,
		Rate:               d.Rate,
		TaxInclusive:       d.TaxInclusive,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		Priority:           d.Priority,
		IsActive:           d.IsActive,
		ValidFrom:          d.ValidFrom,
		ValidUntil:         null.TimeFrom(d.ValidUntil),
		Metadata:           d.Metadata,
	}
}

type TaxRulesListCreateRequest []*TaxRulesCreateRequest

func (d TaxRulesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxRules := range d {
		err = validator.Struct(taxRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxRulesListCreateRequest) ToModelList() []model.TaxRules {
	out := make([]model.TaxRules, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxRulesUpdateRequest struct {
	RuleCode           string          `json:"ruleCode"`
	CountryCode        string          `json:"countryCode"`
	TaxType            string          `json:"taxType"`
	Rate               decimal.Decimal `json:"rate"`
	TaxInclusive       bool            `json:"taxInclusive"`
	LiabilityPartyId   uuid.UUID       `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID       `json:"beneficiaryPartyId"`
	Priority           int             `json:"priority"`
	IsActive           bool            `json:"isActive"`
	ValidFrom          time.Time       `json:"validFrom"`
	ValidUntil         time.Time       `json:"validUntil"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d *TaxRulesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxRulesUpdateRequest) ToModel() model.TaxRules {
	return model.TaxRules{
		RuleCode:           d.RuleCode,
		CountryCode:        d.CountryCode,
		TaxType:            d.TaxType,
		Rate:               d.Rate,
		TaxInclusive:       d.TaxInclusive,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		Priority:           d.Priority,
		IsActive:           d.IsActive,
		ValidFrom:          d.ValidFrom,
		ValidUntil:         null.TimeFrom(d.ValidUntil),
		Metadata:           d.Metadata,
	}
}

type TaxRulesBulkUpdateRequest struct {
	Id                 uuid.UUID       `json:"id"`
	RuleCode           string          `json:"ruleCode"`
	CountryCode        string          `json:"countryCode"`
	TaxType            string          `json:"taxType"`
	Rate               decimal.Decimal `json:"rate"`
	TaxInclusive       bool            `json:"taxInclusive"`
	LiabilityPartyId   uuid.UUID       `json:"liabilityPartyId"`
	BeneficiaryPartyId uuid.UUID       `json:"beneficiaryPartyId"`
	Priority           int             `json:"priority"`
	IsActive           bool            `json:"isActive"`
	ValidFrom          time.Time       `json:"validFrom"`
	ValidUntil         time.Time       `json:"validUntil"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d TaxRulesBulkUpdateRequest) PrimaryID() TaxRulesPrimaryID {
	return TaxRulesPrimaryID{
		Id: d.Id,
	}
}

type TaxRulesListBulkUpdateRequest []*TaxRulesBulkUpdateRequest

func (d TaxRulesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxRules := range d {
		err = validator.Struct(taxRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxRulesBulkUpdateRequest) ToModel() model.TaxRules {
	return model.TaxRules{
		Id:                 d.Id,
		RuleCode:           d.RuleCode,
		CountryCode:        d.CountryCode,
		TaxType:            d.TaxType,
		Rate:               d.Rate,
		TaxInclusive:       d.TaxInclusive,
		LiabilityPartyId:   nuuid.From(d.LiabilityPartyId),
		BeneficiaryPartyId: nuuid.From(d.BeneficiaryPartyId),
		Priority:           d.Priority,
		IsActive:           d.IsActive,
		ValidFrom:          d.ValidFrom,
		ValidUntil:         null.TimeFrom(d.ValidUntil),
		Metadata:           d.Metadata,
	}
}

type TaxRulesResponse struct {
	Id                 uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RuleCode           string          `json:"ruleCode" validate:"required"`
	CountryCode        string          `json:"countryCode"`
	TaxType            string          `json:"taxType" validate:"required"`
	Rate               decimal.Decimal `json:"rate" format:"decimal" example:"100.50"`
	TaxInclusive       bool            `json:"taxInclusive" example:"true"`
	LiabilityPartyId   uuid.UUID       `json:"liabilityPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BeneficiaryPartyId uuid.UUID       `json:"beneficiaryPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Priority           int             `json:"priority" example:"1"`
	IsActive           bool            `json:"isActive" example:"true"`
	ValidFrom          time.Time       `json:"validFrom" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ValidUntil         time.Time       `json:"validUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata           json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewTaxRulesResponse(taxRules model.TaxRules) TaxRulesResponse {
	return TaxRulesResponse{
		Id:                 taxRules.Id,
		RuleCode:           taxRules.RuleCode,
		CountryCode:        taxRules.CountryCode,
		TaxType:            taxRules.TaxType,
		Rate:               taxRules.Rate,
		TaxInclusive:       taxRules.TaxInclusive,
		LiabilityPartyId:   taxRules.LiabilityPartyId.UUID,
		BeneficiaryPartyId: taxRules.BeneficiaryPartyId.UUID,
		Priority:           taxRules.Priority,
		IsActive:           taxRules.IsActive,
		ValidFrom:          taxRules.ValidFrom,
		ValidUntil:         taxRules.ValidUntil.Time,
		Metadata:           taxRules.Metadata,
	}
}

type TaxRulesListResponse []*TaxRulesResponse

func NewTaxRulesListResponse(taxRulesList model.TaxRulesList) TaxRulesListResponse {
	dtoTaxRulesListResponse := TaxRulesListResponse{}
	for _, taxRules := range taxRulesList {
		dtoTaxRulesResponse := NewTaxRulesResponse(*taxRules)
		dtoTaxRulesListResponse = append(dtoTaxRulesListResponse, &dtoTaxRulesResponse)
	}
	return dtoTaxRulesListResponse
}

type TaxRulesPrimaryIDList []TaxRulesPrimaryID

func (d TaxRulesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxRules := range d {
		err = validator.Struct(taxRules)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxRulesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *TaxRulesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxRulesPrimaryID) ToModel() model.TaxRulesPrimaryID {
	return model.TaxRulesPrimaryID{
		Id: d.Id,
	}
}

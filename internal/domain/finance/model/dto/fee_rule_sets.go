package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FeeRuleSetsDTOFieldNameType string

type feeRuleSetsDTOFieldName struct {
	Id             FeeRuleSetsDTOFieldNameType
	FeeProfileId   FeeRuleSetsDTOFieldNameType
	RuleSetCode    FeeRuleSetsDTOFieldNameType
	Precedence     FeeRuleSetsDTOFieldNameType
	EffectiveFrom  FeeRuleSetsDTOFieldNameType
	EffectiveUntil FeeRuleSetsDTOFieldNameType
	RuleSetStatus  FeeRuleSetsDTOFieldNameType
	Metadata       FeeRuleSetsDTOFieldNameType
	MetaCreatedAt  FeeRuleSetsDTOFieldNameType
	MetaCreatedBy  FeeRuleSetsDTOFieldNameType
	MetaUpdatedAt  FeeRuleSetsDTOFieldNameType
	MetaUpdatedBy  FeeRuleSetsDTOFieldNameType
	MetaDeletedAt  FeeRuleSetsDTOFieldNameType
	MetaDeletedBy  FeeRuleSetsDTOFieldNameType
}

var FeeRuleSetsDTOFieldName = feeRuleSetsDTOFieldName{
	Id:             "id",
	FeeProfileId:   "feeProfileId",
	RuleSetCode:    "ruleSetCode",
	Precedence:     "precedence",
	EffectiveFrom:  "effectiveFrom",
	EffectiveUntil: "effectiveUntil",
	RuleSetStatus:  "ruleSetStatus",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformFeeRuleSetsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FeeRuleSetsDTOFieldName.Id):
		return string(model.FeeRuleSetsDBFieldName.Id), true

	case string(FeeRuleSetsDTOFieldName.FeeProfileId):
		return string(model.FeeRuleSetsDBFieldName.FeeProfileId), true

	case string(FeeRuleSetsDTOFieldName.RuleSetCode):
		return string(model.FeeRuleSetsDBFieldName.RuleSetCode), true

	case string(FeeRuleSetsDTOFieldName.Precedence):
		return string(model.FeeRuleSetsDBFieldName.Precedence), true

	case string(FeeRuleSetsDTOFieldName.EffectiveFrom):
		return string(model.FeeRuleSetsDBFieldName.EffectiveFrom), true

	case string(FeeRuleSetsDTOFieldName.EffectiveUntil):
		return string(model.FeeRuleSetsDBFieldName.EffectiveUntil), true

	case string(FeeRuleSetsDTOFieldName.RuleSetStatus):
		return string(model.FeeRuleSetsDBFieldName.RuleSetStatus), true

	case string(FeeRuleSetsDTOFieldName.Metadata):
		return string(model.FeeRuleSetsDBFieldName.Metadata), true

	case string(FeeRuleSetsDTOFieldName.MetaCreatedAt):
		return string(model.FeeRuleSetsDBFieldName.MetaCreatedAt), true

	case string(FeeRuleSetsDTOFieldName.MetaCreatedBy):
		return string(model.FeeRuleSetsDBFieldName.MetaCreatedBy), true

	case string(FeeRuleSetsDTOFieldName.MetaUpdatedAt):
		return string(model.FeeRuleSetsDBFieldName.MetaUpdatedAt), true

	case string(FeeRuleSetsDTOFieldName.MetaUpdatedBy):
		return string(model.FeeRuleSetsDBFieldName.MetaUpdatedBy), true

	case string(FeeRuleSetsDTOFieldName.MetaDeletedAt):
		return string(model.FeeRuleSetsDBFieldName.MetaDeletedAt), true

	case string(FeeRuleSetsDTOFieldName.MetaDeletedBy):
		return string(model.FeeRuleSetsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFeeRuleSetsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFeeRuleSetsBaseFilterField(field string) bool {
	spec, found := model.NewFeeRuleSetsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFeeRuleSetsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFeeRuleSetsProjectionOutputPath(path string) error {
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

func transformFeeRuleSetsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFeeRuleSetsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFeeRuleSetsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFeeRuleSetsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFeeRuleSetsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFeeRuleSetsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFeeRuleSetsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFeeRuleSetsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFeeRuleSetsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFeeRuleSetsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFeeRuleSetsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFeeRuleSetsFilter(filter *model.Filter) {
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
			Field: string(FeeRuleSetsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FeeRuleSetsSelectableResponse map[string]interface{}
type FeeRuleSetsSelectableListResponse []*FeeRuleSetsSelectableResponse

func assignFeeRuleSetsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFeeRuleSetsSelectableValue(out FeeRuleSetsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFeeRuleSetsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFeeRuleSetsSelectableResponse(feeRuleSets model.FeeRuleSets, filter model.Filter) FeeRuleSetsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FeeRuleSetsDBFieldName.Id),
			string(model.FeeRuleSetsDBFieldName.FeeProfileId),
			string(model.FeeRuleSetsDBFieldName.RuleSetCode),
			string(model.FeeRuleSetsDBFieldName.Precedence),
			string(model.FeeRuleSetsDBFieldName.EffectiveFrom),
			string(model.FeeRuleSetsDBFieldName.EffectiveUntil),
			string(model.FeeRuleSetsDBFieldName.RuleSetStatus),
			string(model.FeeRuleSetsDBFieldName.Metadata),
			string(model.FeeRuleSetsDBFieldName.MetaCreatedAt),
			string(model.FeeRuleSetsDBFieldName.MetaCreatedBy),
			string(model.FeeRuleSetsDBFieldName.MetaUpdatedAt),
			string(model.FeeRuleSetsDBFieldName.MetaUpdatedBy),
			string(model.FeeRuleSetsDBFieldName.MetaDeletedAt),
			string(model.FeeRuleSetsDBFieldName.MetaDeletedBy),
		)
	}
	feeRuleSetsSelectableResponse := FeeRuleSetsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FeeRuleSetsDBFieldName.Id):
			key := string(FeeRuleSetsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.Id, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.FeeProfileId):
			key := string(FeeRuleSetsDTOFieldName.FeeProfileId)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.FeeProfileId, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.RuleSetCode):
			key := string(FeeRuleSetsDTOFieldName.RuleSetCode)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.RuleSetCode, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.Precedence):
			key := string(FeeRuleSetsDTOFieldName.Precedence)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.Precedence, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.EffectiveFrom):
			key := string(FeeRuleSetsDTOFieldName.EffectiveFrom)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.EffectiveFrom, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.EffectiveUntil):
			key := string(FeeRuleSetsDTOFieldName.EffectiveUntil)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.EffectiveUntil.Time, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.RuleSetStatus):
			key := string(FeeRuleSetsDTOFieldName.RuleSetStatus)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, model.RuleSetStatus(feeRuleSets.RuleSetStatus), explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.Metadata):
			key := string(FeeRuleSetsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.Metadata, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaCreatedAt):
			key := string(FeeRuleSetsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaCreatedAt, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaCreatedBy):
			key := string(FeeRuleSetsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaCreatedBy, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaUpdatedAt):
			key := string(FeeRuleSetsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaUpdatedAt, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaUpdatedBy):
			key := string(FeeRuleSetsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaUpdatedBy, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaDeletedAt):
			key := string(FeeRuleSetsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaDeletedAt.Time, explicitAlias)

		case string(model.FeeRuleSetsDBFieldName.MetaDeletedBy):
			key := string(FeeRuleSetsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeRuleSetsSelectableValue(feeRuleSetsSelectableResponse, key, feeRuleSets.MetaDeletedBy, explicitAlias)

		}
	}
	return feeRuleSetsSelectableResponse
}

func NewFeeRuleSetsListResponseFromFilterResult(result []model.FeeRuleSetsFilterResult, filter model.Filter) FeeRuleSetsSelectableListResponse {
	dtoFeeRuleSetsListResponse := FeeRuleSetsSelectableListResponse{}
	for _, row := range result {
		dtoFeeRuleSetsResponse := NewFeeRuleSetsSelectableResponse(row.FeeRuleSets, filter)
		dtoFeeRuleSetsListResponse = append(dtoFeeRuleSetsListResponse, &dtoFeeRuleSetsResponse)
	}
	return dtoFeeRuleSetsListResponse
}

type FeeRuleSetsFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     FeeRuleSetsSelectableListResponse `json:"data"`
}

func reverseFeeRuleSetsFilterResults(result []model.FeeRuleSetsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFeeRuleSetsFilterResponse(result []model.FeeRuleSetsFilterResult, filter model.Filter) (resp FeeRuleSetsFilterResponse) {
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
			reverseFeeRuleSetsFilterResults(dataResult)
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

	resp.Data = NewFeeRuleSetsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FeeRuleSetsCreateRequest struct {
	FeeProfileId   uuid.UUID           `json:"feeProfileId"`
	RuleSetCode    string              `json:"ruleSetCode"`
	Precedence     int                 `json:"precedence"`
	EffectiveFrom  time.Time           `json:"effectiveFrom"`
	EffectiveUntil time.Time           `json:"effectiveUntil"`
	RuleSetStatus  model.RuleSetStatus `json:"ruleSetStatus" example:"active" enums:"active,inactive"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d *FeeRuleSetsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FeeRuleSetsCreateRequest) ToModel() model.FeeRuleSets {
	id, _ := uuid.NewV4()
	return model.FeeRuleSets{
		Id:             id,
		FeeProfileId:   d.FeeProfileId,
		RuleSetCode:    d.RuleSetCode,
		Precedence:     d.Precedence,
		EffectiveFrom:  d.EffectiveFrom,
		EffectiveUntil: null.TimeFrom(d.EffectiveUntil),
		RuleSetStatus:  d.RuleSetStatus,
		Metadata:       d.Metadata,
	}
}

type FeeRuleSetsListCreateRequest []*FeeRuleSetsCreateRequest

func (d FeeRuleSetsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleSets := range d {
		err = validator.Struct(feeRuleSets)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRuleSetsListCreateRequest) ToModelList() []model.FeeRuleSets {
	out := make([]model.FeeRuleSets, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FeeRuleSetsUpdateRequest struct {
	FeeProfileId   uuid.UUID           `json:"feeProfileId"`
	RuleSetCode    string              `json:"ruleSetCode"`
	Precedence     int                 `json:"precedence"`
	EffectiveFrom  time.Time           `json:"effectiveFrom"`
	EffectiveUntil time.Time           `json:"effectiveUntil"`
	RuleSetStatus  model.RuleSetStatus `json:"ruleSetStatus" example:"active" enums:"active,inactive"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d *FeeRuleSetsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FeeRuleSetsUpdateRequest) ToModel() model.FeeRuleSets {
	return model.FeeRuleSets{
		FeeProfileId:   d.FeeProfileId,
		RuleSetCode:    d.RuleSetCode,
		Precedence:     d.Precedence,
		EffectiveFrom:  d.EffectiveFrom,
		EffectiveUntil: null.TimeFrom(d.EffectiveUntil),
		RuleSetStatus:  d.RuleSetStatus,
		Metadata:       d.Metadata,
	}
}

type FeeRuleSetsBulkUpdateRequest struct {
	Id             uuid.UUID           `json:"id"`
	FeeProfileId   uuid.UUID           `json:"feeProfileId"`
	RuleSetCode    string              `json:"ruleSetCode"`
	Precedence     int                 `json:"precedence"`
	EffectiveFrom  time.Time           `json:"effectiveFrom"`
	EffectiveUntil time.Time           `json:"effectiveUntil"`
	RuleSetStatus  model.RuleSetStatus `json:"ruleSetStatus" example:"active" enums:"active,inactive"`
	Metadata       json.RawMessage     `json:"metadata"`
}

func (d FeeRuleSetsBulkUpdateRequest) PrimaryID() FeeRuleSetsPrimaryID {
	return FeeRuleSetsPrimaryID{
		Id: d.Id,
	}
}

type FeeRuleSetsListBulkUpdateRequest []*FeeRuleSetsBulkUpdateRequest

func (d FeeRuleSetsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleSets := range d {
		err = validator.Struct(feeRuleSets)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeRuleSetsBulkUpdateRequest) ToModel() model.FeeRuleSets {
	return model.FeeRuleSets{
		Id:             d.Id,
		FeeProfileId:   d.FeeProfileId,
		RuleSetCode:    d.RuleSetCode,
		Precedence:     d.Precedence,
		EffectiveFrom:  d.EffectiveFrom,
		EffectiveUntil: null.TimeFrom(d.EffectiveUntil),
		RuleSetStatus:  d.RuleSetStatus,
		Metadata:       d.Metadata,
	}
}

type FeeRuleSetsResponse struct {
	Id             uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	FeeProfileId   uuid.UUID           `json:"feeProfileId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RuleSetCode    string              `json:"ruleSetCode" validate:"required"`
	Precedence     int                 `json:"precedence" example:"1"`
	EffectiveFrom  time.Time           `json:"effectiveFrom" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	EffectiveUntil time.Time           `json:"effectiveUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RuleSetStatus  model.RuleSetStatus `json:"ruleSetStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Metadata       json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewFeeRuleSetsResponse(feeRuleSets model.FeeRuleSets) FeeRuleSetsResponse {
	return FeeRuleSetsResponse{
		Id:             feeRuleSets.Id,
		FeeProfileId:   feeRuleSets.FeeProfileId,
		RuleSetCode:    feeRuleSets.RuleSetCode,
		Precedence:     feeRuleSets.Precedence,
		EffectiveFrom:  feeRuleSets.EffectiveFrom,
		EffectiveUntil: feeRuleSets.EffectiveUntil.Time,
		RuleSetStatus:  model.RuleSetStatus(feeRuleSets.RuleSetStatus),
		Metadata:       feeRuleSets.Metadata,
	}
}

type FeeRuleSetsListResponse []*FeeRuleSetsResponse

func NewFeeRuleSetsListResponse(feeRuleSetsList model.FeeRuleSetsList) FeeRuleSetsListResponse {
	dtoFeeRuleSetsListResponse := FeeRuleSetsListResponse{}
	for _, feeRuleSets := range feeRuleSetsList {
		dtoFeeRuleSetsResponse := NewFeeRuleSetsResponse(*feeRuleSets)
		dtoFeeRuleSetsListResponse = append(dtoFeeRuleSetsListResponse, &dtoFeeRuleSetsResponse)
	}
	return dtoFeeRuleSetsListResponse
}

type FeeRuleSetsPrimaryIDList []FeeRuleSetsPrimaryID

func (d FeeRuleSetsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeRuleSets := range d {
		err = validator.Struct(feeRuleSets)
		if err != nil {
			return
		}
	}
	return nil
}

type FeeRuleSetsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FeeRuleSetsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FeeRuleSetsPrimaryID) ToModel() model.FeeRuleSetsPrimaryID {
	return model.FeeRuleSetsPrimaryID{
		Id: d.Id,
	}
}

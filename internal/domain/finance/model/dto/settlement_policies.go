package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type SettlementPoliciesDTOFieldNameType string

type settlementPoliciesDTOFieldName struct {
	Id              SettlementPoliciesDTOFieldNameType
	PolicyCode      SettlementPoliciesDTOFieldNameType
	MerchantPartyId SettlementPoliciesDTOFieldNameType
	PolicyScope     SettlementPoliciesDTOFieldNameType
	PolicyStatus    SettlementPoliciesDTOFieldNameType
	Description     SettlementPoliciesDTOFieldNameType
	Metadata        SettlementPoliciesDTOFieldNameType
	MetaCreatedAt   SettlementPoliciesDTOFieldNameType
	MetaCreatedBy   SettlementPoliciesDTOFieldNameType
	MetaUpdatedAt   SettlementPoliciesDTOFieldNameType
	MetaUpdatedBy   SettlementPoliciesDTOFieldNameType
	MetaDeletedAt   SettlementPoliciesDTOFieldNameType
	MetaDeletedBy   SettlementPoliciesDTOFieldNameType
}

var SettlementPoliciesDTOFieldName = settlementPoliciesDTOFieldName{
	Id:              "id",
	PolicyCode:      "policyCode",
	MerchantPartyId: "merchantPartyId",
	PolicyScope:     "policyScope",
	PolicyStatus:    "policyStatus",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformSettlementPoliciesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(SettlementPoliciesDTOFieldName.Id):
		return string(model.SettlementPoliciesDBFieldName.Id), true

	case string(SettlementPoliciesDTOFieldName.PolicyCode):
		return string(model.SettlementPoliciesDBFieldName.PolicyCode), true

	case string(SettlementPoliciesDTOFieldName.MerchantPartyId):
		return string(model.SettlementPoliciesDBFieldName.MerchantPartyId), true

	case string(SettlementPoliciesDTOFieldName.PolicyScope):
		return string(model.SettlementPoliciesDBFieldName.PolicyScope), true

	case string(SettlementPoliciesDTOFieldName.PolicyStatus):
		return string(model.SettlementPoliciesDBFieldName.PolicyStatus), true

	case string(SettlementPoliciesDTOFieldName.Description):
		return string(model.SettlementPoliciesDBFieldName.Description), true

	case string(SettlementPoliciesDTOFieldName.Metadata):
		return string(model.SettlementPoliciesDBFieldName.Metadata), true

	case string(SettlementPoliciesDTOFieldName.MetaCreatedAt):
		return string(model.SettlementPoliciesDBFieldName.MetaCreatedAt), true

	case string(SettlementPoliciesDTOFieldName.MetaCreatedBy):
		return string(model.SettlementPoliciesDBFieldName.MetaCreatedBy), true

	case string(SettlementPoliciesDTOFieldName.MetaUpdatedAt):
		return string(model.SettlementPoliciesDBFieldName.MetaUpdatedAt), true

	case string(SettlementPoliciesDTOFieldName.MetaUpdatedBy):
		return string(model.SettlementPoliciesDBFieldName.MetaUpdatedBy), true

	case string(SettlementPoliciesDTOFieldName.MetaDeletedAt):
		return string(model.SettlementPoliciesDBFieldName.MetaDeletedAt), true

	case string(SettlementPoliciesDTOFieldName.MetaDeletedBy):
		return string(model.SettlementPoliciesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewSettlementPoliciesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isSettlementPoliciesBaseFilterField(field string) bool {
	spec, found := model.NewSettlementPoliciesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeSettlementPoliciesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateSettlementPoliciesProjectionOutputPath(path string) error {
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

func transformSettlementPoliciesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformSettlementPoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformSettlementPoliciesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformSettlementPoliciesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformSettlementPoliciesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isSettlementPoliciesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateSettlementPoliciesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeSettlementPoliciesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformSettlementPoliciesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformSettlementPoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformSettlementPoliciesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultSettlementPoliciesFilter(filter *model.Filter) {
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
			Field: string(SettlementPoliciesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type SettlementPoliciesSelectableResponse map[string]interface{}
type SettlementPoliciesSelectableListResponse []*SettlementPoliciesSelectableResponse

func assignSettlementPoliciesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setSettlementPoliciesSelectableValue(out SettlementPoliciesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignSettlementPoliciesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewSettlementPoliciesSelectableResponse(settlementPolicies model.SettlementPolicies, filter model.Filter) SettlementPoliciesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.SettlementPoliciesDBFieldName.Id),
			string(model.SettlementPoliciesDBFieldName.PolicyCode),
			string(model.SettlementPoliciesDBFieldName.MerchantPartyId),
			string(model.SettlementPoliciesDBFieldName.PolicyScope),
			string(model.SettlementPoliciesDBFieldName.PolicyStatus),
			string(model.SettlementPoliciesDBFieldName.Description),
			string(model.SettlementPoliciesDBFieldName.Metadata),
			string(model.SettlementPoliciesDBFieldName.MetaCreatedAt),
			string(model.SettlementPoliciesDBFieldName.MetaCreatedBy),
			string(model.SettlementPoliciesDBFieldName.MetaUpdatedAt),
			string(model.SettlementPoliciesDBFieldName.MetaUpdatedBy),
			string(model.SettlementPoliciesDBFieldName.MetaDeletedAt),
			string(model.SettlementPoliciesDBFieldName.MetaDeletedBy),
		)
	}
	settlementPoliciesSelectableResponse := SettlementPoliciesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.SettlementPoliciesDBFieldName.Id):
			key := string(SettlementPoliciesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.Id, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.PolicyCode):
			key := string(SettlementPoliciesDTOFieldName.PolicyCode)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.PolicyCode, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MerchantPartyId):
			key := string(SettlementPoliciesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MerchantPartyId.UUID, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.PolicyScope):
			key := string(SettlementPoliciesDTOFieldName.PolicyScope)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, model.SettlementPoliciesPolicyScope(settlementPolicies.PolicyScope), explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.PolicyStatus):
			key := string(SettlementPoliciesDTOFieldName.PolicyStatus)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, model.SettlementPoliciesPolicyStatus(settlementPolicies.PolicyStatus), explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.Description):
			key := string(SettlementPoliciesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.Description.String, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.Metadata):
			key := string(SettlementPoliciesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.Metadata, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaCreatedAt):
			key := string(SettlementPoliciesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaCreatedAt, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaCreatedBy):
			key := string(SettlementPoliciesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaCreatedBy, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaUpdatedAt):
			key := string(SettlementPoliciesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaUpdatedAt, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaUpdatedBy):
			key := string(SettlementPoliciesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaUpdatedBy, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaDeletedAt):
			key := string(SettlementPoliciesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaDeletedAt.Time, explicitAlias)

		case string(model.SettlementPoliciesDBFieldName.MetaDeletedBy):
			key := string(SettlementPoliciesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPoliciesSelectableValue(settlementPoliciesSelectableResponse, key, settlementPolicies.MetaDeletedBy, explicitAlias)

		}
	}
	return settlementPoliciesSelectableResponse
}

func NewSettlementPoliciesListResponseFromFilterResult(result []model.SettlementPoliciesFilterResult, filter model.Filter) SettlementPoliciesSelectableListResponse {
	dtoSettlementPoliciesListResponse := SettlementPoliciesSelectableListResponse{}
	for _, row := range result {
		dtoSettlementPoliciesResponse := NewSettlementPoliciesSelectableResponse(row.SettlementPolicies, filter)
		dtoSettlementPoliciesListResponse = append(dtoSettlementPoliciesListResponse, &dtoSettlementPoliciesResponse)
	}
	return dtoSettlementPoliciesListResponse
}

type SettlementPoliciesFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     SettlementPoliciesSelectableListResponse `json:"data"`
}

func reverseSettlementPoliciesFilterResults(result []model.SettlementPoliciesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewSettlementPoliciesFilterResponse(result []model.SettlementPoliciesFilterResult, filter model.Filter) (resp SettlementPoliciesFilterResponse) {
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
			reverseSettlementPoliciesFilterResults(dataResult)
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

	resp.Data = NewSettlementPoliciesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type SettlementPoliciesCreateRequest struct {
	PolicyCode      string                               `json:"policyCode"`
	MerchantPartyId uuid.UUID                            `json:"merchantPartyId"`
	PolicyScope     model.SettlementPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	PolicyStatus    model.SettlementPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Description     string                               `json:"description"`
	Metadata        json.RawMessage                      `json:"metadata"`
}

func (d *SettlementPoliciesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *SettlementPoliciesCreateRequest) ToModel() model.SettlementPolicies {
	id, _ := uuid.NewV4()
	return model.SettlementPolicies{
		Id:              id,
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		PolicyStatus:    d.PolicyStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type SettlementPoliciesListCreateRequest []*SettlementPoliciesCreateRequest

func (d SettlementPoliciesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicies := range d {
		err = validator.Struct(settlementPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementPoliciesListCreateRequest) ToModelList() []model.SettlementPolicies {
	out := make([]model.SettlementPolicies, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type SettlementPoliciesUpdateRequest struct {
	PolicyCode      string                               `json:"policyCode"`
	MerchantPartyId uuid.UUID                            `json:"merchantPartyId"`
	PolicyScope     model.SettlementPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	PolicyStatus    model.SettlementPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Description     string                               `json:"description"`
	Metadata        json.RawMessage                      `json:"metadata"`
}

func (d *SettlementPoliciesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d SettlementPoliciesUpdateRequest) ToModel() model.SettlementPolicies {
	return model.SettlementPolicies{
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		PolicyStatus:    d.PolicyStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type SettlementPoliciesBulkUpdateRequest struct {
	Id              uuid.UUID                            `json:"id"`
	PolicyCode      string                               `json:"policyCode"`
	MerchantPartyId uuid.UUID                            `json:"merchantPartyId"`
	PolicyScope     model.SettlementPoliciesPolicyScope  `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,listing_type,category"`
	PolicyStatus    model.SettlementPoliciesPolicyStatus `json:"policyStatus" example:"active" enums:"active,inactive"`
	Description     string                               `json:"description"`
	Metadata        json.RawMessage                      `json:"metadata"`
}

func (d SettlementPoliciesBulkUpdateRequest) PrimaryID() SettlementPoliciesPrimaryID {
	return SettlementPoliciesPrimaryID{
		Id: d.Id,
	}
}

type SettlementPoliciesListBulkUpdateRequest []*SettlementPoliciesBulkUpdateRequest

func (d SettlementPoliciesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicies := range d {
		err = validator.Struct(settlementPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementPoliciesBulkUpdateRequest) ToModel() model.SettlementPolicies {
	return model.SettlementPolicies{
		Id:              d.Id,
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		PolicyStatus:    d.PolicyStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type SettlementPoliciesResponse struct {
	Id              uuid.UUID                            `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyCode      string                               `json:"policyCode" validate:"required"`
	MerchantPartyId uuid.UUID                            `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyScope     model.SettlementPoliciesPolicyScope  `json:"policyScope" validate:"required,oneof=platform_default merchant listing_type category" enums:"platform_default,merchant,listing_type,category"`
	PolicyStatus    model.SettlementPoliciesPolicyStatus `json:"policyStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Description     string                               `json:"description"`
	Metadata        json.RawMessage                      `json:"metadata" swaggertype:"object"`
}

func NewSettlementPoliciesResponse(settlementPolicies model.SettlementPolicies) SettlementPoliciesResponse {
	return SettlementPoliciesResponse{
		Id:              settlementPolicies.Id,
		PolicyCode:      settlementPolicies.PolicyCode,
		MerchantPartyId: settlementPolicies.MerchantPartyId.UUID,
		PolicyScope:     model.SettlementPoliciesPolicyScope(settlementPolicies.PolicyScope),
		PolicyStatus:    model.SettlementPoliciesPolicyStatus(settlementPolicies.PolicyStatus),
		Description:     settlementPolicies.Description.String,
		Metadata:        settlementPolicies.Metadata,
	}
}

type SettlementPoliciesListResponse []*SettlementPoliciesResponse

func NewSettlementPoliciesListResponse(settlementPoliciesList model.SettlementPoliciesList) SettlementPoliciesListResponse {
	dtoSettlementPoliciesListResponse := SettlementPoliciesListResponse{}
	for _, settlementPolicies := range settlementPoliciesList {
		dtoSettlementPoliciesResponse := NewSettlementPoliciesResponse(*settlementPolicies)
		dtoSettlementPoliciesListResponse = append(dtoSettlementPoliciesListResponse, &dtoSettlementPoliciesResponse)
	}
	return dtoSettlementPoliciesListResponse
}

type SettlementPoliciesPrimaryIDList []SettlementPoliciesPrimaryID

func (d SettlementPoliciesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicies := range d {
		err = validator.Struct(settlementPolicies)
		if err != nil {
			return
		}
	}
	return nil
}

type SettlementPoliciesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *SettlementPoliciesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d SettlementPoliciesPrimaryID) ToModel() model.SettlementPoliciesPrimaryID {
	return model.SettlementPoliciesPrimaryID{
		Id: d.Id,
	}
}

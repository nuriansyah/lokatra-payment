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

type ReservePoliciesDTOFieldNameType string

type reservePoliciesDTOFieldName struct {
	Id              ReservePoliciesDTOFieldNameType
	PolicyCode      ReservePoliciesDTOFieldNameType
	MerchantPartyId ReservePoliciesDTOFieldNameType
	PolicyScope     ReservePoliciesDTOFieldNameType
	ReserveStatus   ReservePoliciesDTOFieldNameType
	Description     ReservePoliciesDTOFieldNameType
	Metadata        ReservePoliciesDTOFieldNameType
	MetaCreatedAt   ReservePoliciesDTOFieldNameType
	MetaCreatedBy   ReservePoliciesDTOFieldNameType
	MetaUpdatedAt   ReservePoliciesDTOFieldNameType
	MetaUpdatedBy   ReservePoliciesDTOFieldNameType
	MetaDeletedAt   ReservePoliciesDTOFieldNameType
	MetaDeletedBy   ReservePoliciesDTOFieldNameType
}

var ReservePoliciesDTOFieldName = reservePoliciesDTOFieldName{
	Id:              "id",
	PolicyCode:      "policyCode",
	MerchantPartyId: "merchantPartyId",
	PolicyScope:     "policyScope",
	ReserveStatus:   "reserveStatus",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformReservePoliciesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReservePoliciesDTOFieldName.Id):
		return string(model.ReservePoliciesDBFieldName.Id), true

	case string(ReservePoliciesDTOFieldName.PolicyCode):
		return string(model.ReservePoliciesDBFieldName.PolicyCode), true

	case string(ReservePoliciesDTOFieldName.MerchantPartyId):
		return string(model.ReservePoliciesDBFieldName.MerchantPartyId), true

	case string(ReservePoliciesDTOFieldName.PolicyScope):
		return string(model.ReservePoliciesDBFieldName.PolicyScope), true

	case string(ReservePoliciesDTOFieldName.ReserveStatus):
		return string(model.ReservePoliciesDBFieldName.ReserveStatus), true

	case string(ReservePoliciesDTOFieldName.Description):
		return string(model.ReservePoliciesDBFieldName.Description), true

	case string(ReservePoliciesDTOFieldName.Metadata):
		return string(model.ReservePoliciesDBFieldName.Metadata), true

	case string(ReservePoliciesDTOFieldName.MetaCreatedAt):
		return string(model.ReservePoliciesDBFieldName.MetaCreatedAt), true

	case string(ReservePoliciesDTOFieldName.MetaCreatedBy):
		return string(model.ReservePoliciesDBFieldName.MetaCreatedBy), true

	case string(ReservePoliciesDTOFieldName.MetaUpdatedAt):
		return string(model.ReservePoliciesDBFieldName.MetaUpdatedAt), true

	case string(ReservePoliciesDTOFieldName.MetaUpdatedBy):
		return string(model.ReservePoliciesDBFieldName.MetaUpdatedBy), true

	case string(ReservePoliciesDTOFieldName.MetaDeletedAt):
		return string(model.ReservePoliciesDBFieldName.MetaDeletedAt), true

	case string(ReservePoliciesDTOFieldName.MetaDeletedBy):
		return string(model.ReservePoliciesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReservePoliciesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReservePoliciesBaseFilterField(field string) bool {
	spec, found := model.NewReservePoliciesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReservePoliciesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReservePoliciesProjectionOutputPath(path string) error {
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

func transformReservePoliciesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReservePoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReservePoliciesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReservePoliciesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReservePoliciesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReservePoliciesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReservePoliciesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReservePoliciesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReservePoliciesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReservePoliciesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReservePoliciesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReservePoliciesFilter(filter *model.Filter) {
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
			Field: string(ReservePoliciesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReservePoliciesSelectableResponse map[string]interface{}
type ReservePoliciesSelectableListResponse []*ReservePoliciesSelectableResponse

func assignReservePoliciesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReservePoliciesSelectableValue(out ReservePoliciesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReservePoliciesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReservePoliciesSelectableResponse(reservePolicies model.ReservePolicies, filter model.Filter) ReservePoliciesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReservePoliciesDBFieldName.Id),
			string(model.ReservePoliciesDBFieldName.PolicyCode),
			string(model.ReservePoliciesDBFieldName.MerchantPartyId),
			string(model.ReservePoliciesDBFieldName.PolicyScope),
			string(model.ReservePoliciesDBFieldName.ReserveStatus),
			string(model.ReservePoliciesDBFieldName.Description),
			string(model.ReservePoliciesDBFieldName.Metadata),
			string(model.ReservePoliciesDBFieldName.MetaCreatedAt),
			string(model.ReservePoliciesDBFieldName.MetaCreatedBy),
			string(model.ReservePoliciesDBFieldName.MetaUpdatedAt),
			string(model.ReservePoliciesDBFieldName.MetaUpdatedBy),
			string(model.ReservePoliciesDBFieldName.MetaDeletedAt),
			string(model.ReservePoliciesDBFieldName.MetaDeletedBy),
		)
	}
	reservePoliciesSelectableResponse := ReservePoliciesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReservePoliciesDBFieldName.Id):
			key := string(ReservePoliciesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.Id, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.PolicyCode):
			key := string(ReservePoliciesDTOFieldName.PolicyCode)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.PolicyCode, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MerchantPartyId):
			key := string(ReservePoliciesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MerchantPartyId.UUID, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.PolicyScope):
			key := string(ReservePoliciesDTOFieldName.PolicyScope)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, model.ReservePoliciesPolicyScope(reservePolicies.PolicyScope), explicitAlias)

		case string(model.ReservePoliciesDBFieldName.ReserveStatus):
			key := string(ReservePoliciesDTOFieldName.ReserveStatus)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, model.ReserveStatus(reservePolicies.ReserveStatus), explicitAlias)

		case string(model.ReservePoliciesDBFieldName.Description):
			key := string(ReservePoliciesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.Description.String, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.Metadata):
			key := string(ReservePoliciesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.Metadata, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaCreatedAt):
			key := string(ReservePoliciesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaCreatedAt, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaCreatedBy):
			key := string(ReservePoliciesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaCreatedBy, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaUpdatedAt):
			key := string(ReservePoliciesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaUpdatedAt, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaUpdatedBy):
			key := string(ReservePoliciesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaUpdatedBy, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaDeletedAt):
			key := string(ReservePoliciesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReservePoliciesDBFieldName.MetaDeletedBy):
			key := string(ReservePoliciesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReservePoliciesSelectableValue(reservePoliciesSelectableResponse, key, reservePolicies.MetaDeletedBy, explicitAlias)

		}
	}
	return reservePoliciesSelectableResponse
}

func NewReservePoliciesListResponseFromFilterResult(result []model.ReservePoliciesFilterResult, filter model.Filter) ReservePoliciesSelectableListResponse {
	dtoReservePoliciesListResponse := ReservePoliciesSelectableListResponse{}
	for _, row := range result {
		dtoReservePoliciesResponse := NewReservePoliciesSelectableResponse(row.ReservePolicies, filter)
		dtoReservePoliciesListResponse = append(dtoReservePoliciesListResponse, &dtoReservePoliciesResponse)
	}
	return dtoReservePoliciesListResponse
}

type ReservePoliciesFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     ReservePoliciesSelectableListResponse `json:"data"`
}

func reverseReservePoliciesFilterResults(result []model.ReservePoliciesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReservePoliciesFilterResponse(result []model.ReservePoliciesFilterResult, filter model.Filter) (resp ReservePoliciesFilterResponse) {
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
			reverseReservePoliciesFilterResults(dataResult)
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

	resp.Data = NewReservePoliciesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReservePoliciesCreateRequest struct {
	PolicyCode      string                           `json:"policyCode"`
	MerchantPartyId uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope     model.ReservePoliciesPolicyScope `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,category,risk_tier"`
	ReserveStatus   model.ReserveStatus              `json:"reserveStatus" example:"active" enums:"active,inactive"`
	Description     string                           `json:"description"`
	Metadata        json.RawMessage                  `json:"metadata"`
}

func (d *ReservePoliciesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReservePoliciesCreateRequest) ToModel() model.ReservePolicies {
	id, _ := uuid.NewV4()
	return model.ReservePolicies{
		Id:              id,
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		ReserveStatus:   d.ReserveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type ReservePoliciesListCreateRequest []*ReservePoliciesCreateRequest

func (d ReservePoliciesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reservePolicies := range d {
		err = validator.Struct(reservePolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReservePoliciesListCreateRequest) ToModelList() []model.ReservePolicies {
	out := make([]model.ReservePolicies, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReservePoliciesUpdateRequest struct {
	PolicyCode      string                           `json:"policyCode"`
	MerchantPartyId uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope     model.ReservePoliciesPolicyScope `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,category,risk_tier"`
	ReserveStatus   model.ReserveStatus              `json:"reserveStatus" example:"active" enums:"active,inactive"`
	Description     string                           `json:"description"`
	Metadata        json.RawMessage                  `json:"metadata"`
}

func (d *ReservePoliciesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReservePoliciesUpdateRequest) ToModel() model.ReservePolicies {
	return model.ReservePolicies{
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		ReserveStatus:   d.ReserveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type ReservePoliciesBulkUpdateRequest struct {
	Id              uuid.UUID                        `json:"id"`
	PolicyCode      string                           `json:"policyCode"`
	MerchantPartyId uuid.UUID                        `json:"merchantPartyId"`
	PolicyScope     model.ReservePoliciesPolicyScope `json:"policyScope" example:"platform_default" enums:"platform_default,merchant,category,risk_tier"`
	ReserveStatus   model.ReserveStatus              `json:"reserveStatus" example:"active" enums:"active,inactive"`
	Description     string                           `json:"description"`
	Metadata        json.RawMessage                  `json:"metadata"`
}

func (d ReservePoliciesBulkUpdateRequest) PrimaryID() ReservePoliciesPrimaryID {
	return ReservePoliciesPrimaryID{
		Id: d.Id,
	}
}

type ReservePoliciesListBulkUpdateRequest []*ReservePoliciesBulkUpdateRequest

func (d ReservePoliciesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reservePolicies := range d {
		err = validator.Struct(reservePolicies)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReservePoliciesBulkUpdateRequest) ToModel() model.ReservePolicies {
	return model.ReservePolicies{
		Id:              d.Id,
		PolicyCode:      d.PolicyCode,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		PolicyScope:     d.PolicyScope,
		ReserveStatus:   d.ReserveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type ReservePoliciesResponse struct {
	Id              uuid.UUID                        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyCode      string                           `json:"policyCode" validate:"required"`
	MerchantPartyId uuid.UUID                        `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PolicyScope     model.ReservePoliciesPolicyScope `json:"policyScope" validate:"required,oneof=platform_default merchant category risk_tier" enums:"platform_default,merchant,category,risk_tier"`
	ReserveStatus   model.ReserveStatus              `json:"reserveStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Description     string                           `json:"description"`
	Metadata        json.RawMessage                  `json:"metadata" swaggertype:"object"`
}

func NewReservePoliciesResponse(reservePolicies model.ReservePolicies) ReservePoliciesResponse {
	return ReservePoliciesResponse{
		Id:              reservePolicies.Id,
		PolicyCode:      reservePolicies.PolicyCode,
		MerchantPartyId: reservePolicies.MerchantPartyId.UUID,
		PolicyScope:     model.ReservePoliciesPolicyScope(reservePolicies.PolicyScope),
		ReserveStatus:   model.ReserveStatus(reservePolicies.ReserveStatus),
		Description:     reservePolicies.Description.String,
		Metadata:        reservePolicies.Metadata,
	}
}

type ReservePoliciesListResponse []*ReservePoliciesResponse

func NewReservePoliciesListResponse(reservePoliciesList model.ReservePoliciesList) ReservePoliciesListResponse {
	dtoReservePoliciesListResponse := ReservePoliciesListResponse{}
	for _, reservePolicies := range reservePoliciesList {
		dtoReservePoliciesResponse := NewReservePoliciesResponse(*reservePolicies)
		dtoReservePoliciesListResponse = append(dtoReservePoliciesListResponse, &dtoReservePoliciesResponse)
	}
	return dtoReservePoliciesListResponse
}

type ReservePoliciesPrimaryIDList []ReservePoliciesPrimaryID

func (d ReservePoliciesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reservePolicies := range d {
		err = validator.Struct(reservePolicies)
		if err != nil {
			return
		}
	}
	return nil
}

type ReservePoliciesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReservePoliciesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReservePoliciesPrimaryID) ToModel() model.ReservePoliciesPrimaryID {
	return model.ReservePoliciesPrimaryID{
		Id: d.Id,
	}
}

package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderWebhookEndpointsDTOFieldNameType string

type providerWebhookEndpointsDTOFieldName struct {
	Id                 ProviderWebhookEndpointsDTOFieldNameType
	ProviderAccountId  ProviderWebhookEndpointsDTOFieldNameType
	ProviderCode       ProviderWebhookEndpointsDTOFieldNameType
	EndpointKey        ProviderWebhookEndpointsDTOFieldNameType
	Environment        ProviderWebhookEndpointsDTOFieldNameType
	SecretRef          ProviderWebhookEndpointsDTOFieldNameType
	SignatureAlgorithm ProviderWebhookEndpointsDTOFieldNameType
	IsActive           ProviderWebhookEndpointsDTOFieldNameType
	Metadata           ProviderWebhookEndpointsDTOFieldNameType
	MetaCreatedAt      ProviderWebhookEndpointsDTOFieldNameType
	MetaCreatedBy      ProviderWebhookEndpointsDTOFieldNameType
	MetaUpdatedAt      ProviderWebhookEndpointsDTOFieldNameType
	MetaUpdatedBy      ProviderWebhookEndpointsDTOFieldNameType
	MetaDeletedAt      ProviderWebhookEndpointsDTOFieldNameType
	MetaDeletedBy      ProviderWebhookEndpointsDTOFieldNameType
}

var ProviderWebhookEndpointsDTOFieldName = providerWebhookEndpointsDTOFieldName{
	Id:                 "id",
	ProviderAccountId:  "providerAccountId",
	ProviderCode:       "providerCode",
	EndpointKey:        "endpointKey",
	Environment:        "environment",
	SecretRef:          "secretRef",
	SignatureAlgorithm: "signatureAlgorithm",
	IsActive:           "isActive",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformProviderWebhookEndpointsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderWebhookEndpointsDTOFieldName.Id):
		return string(model.ProviderWebhookEndpointsDBFieldName.Id), true

	case string(ProviderWebhookEndpointsDTOFieldName.ProviderAccountId):
		return string(model.ProviderWebhookEndpointsDBFieldName.ProviderAccountId), true

	case string(ProviderWebhookEndpointsDTOFieldName.ProviderCode):
		return string(model.ProviderWebhookEndpointsDBFieldName.ProviderCode), true

	case string(ProviderWebhookEndpointsDTOFieldName.EndpointKey):
		return string(model.ProviderWebhookEndpointsDBFieldName.EndpointKey), true

	case string(ProviderWebhookEndpointsDTOFieldName.Environment):
		return string(model.ProviderWebhookEndpointsDBFieldName.Environment), true

	case string(ProviderWebhookEndpointsDTOFieldName.SecretRef):
		return string(model.ProviderWebhookEndpointsDBFieldName.SecretRef), true

	case string(ProviderWebhookEndpointsDTOFieldName.SignatureAlgorithm):
		return string(model.ProviderWebhookEndpointsDBFieldName.SignatureAlgorithm), true

	case string(ProviderWebhookEndpointsDTOFieldName.IsActive):
		return string(model.ProviderWebhookEndpointsDBFieldName.IsActive), true

	case string(ProviderWebhookEndpointsDTOFieldName.Metadata):
		return string(model.ProviderWebhookEndpointsDBFieldName.Metadata), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaCreatedAt):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedAt), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaCreatedBy):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedBy), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedAt), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedBy), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaDeletedAt):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedAt), true

	case string(ProviderWebhookEndpointsDTOFieldName.MetaDeletedBy):
		return string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderWebhookEndpointsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderWebhookEndpointsBaseFilterField(field string) bool {
	spec, found := model.NewProviderWebhookEndpointsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderWebhookEndpointsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderWebhookEndpointsProjectionOutputPath(path string) error {
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

func transformProviderWebhookEndpointsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderWebhookEndpointsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderWebhookEndpointsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderWebhookEndpointsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderWebhookEndpointsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderWebhookEndpointsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderWebhookEndpointsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderWebhookEndpointsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderWebhookEndpointsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderWebhookEndpointsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderWebhookEndpointsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderWebhookEndpointsFilter(filter *model.Filter) {
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
			Field: string(ProviderWebhookEndpointsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderWebhookEndpointsSelectableResponse map[string]interface{}
type ProviderWebhookEndpointsSelectableListResponse []*ProviderWebhookEndpointsSelectableResponse

func assignProviderWebhookEndpointsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderWebhookEndpointsSelectableValue(out ProviderWebhookEndpointsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderWebhookEndpointsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderWebhookEndpointsSelectableResponse(providerWebhookEndpoints model.ProviderWebhookEndpoints, filter model.Filter) ProviderWebhookEndpointsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderWebhookEndpointsDBFieldName.Id),
			string(model.ProviderWebhookEndpointsDBFieldName.ProviderAccountId),
			string(model.ProviderWebhookEndpointsDBFieldName.ProviderCode),
			string(model.ProviderWebhookEndpointsDBFieldName.EndpointKey),
			string(model.ProviderWebhookEndpointsDBFieldName.Environment),
			string(model.ProviderWebhookEndpointsDBFieldName.SecretRef),
			string(model.ProviderWebhookEndpointsDBFieldName.SignatureAlgorithm),
			string(model.ProviderWebhookEndpointsDBFieldName.IsActive),
			string(model.ProviderWebhookEndpointsDBFieldName.Metadata),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedAt),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedBy),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedAt),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedBy),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedAt),
			string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedBy),
		)
	}
	providerWebhookEndpointsSelectableResponse := ProviderWebhookEndpointsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderWebhookEndpointsDBFieldName.Id):
			key := string(ProviderWebhookEndpointsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.Id, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.ProviderAccountId):
			key := string(ProviderWebhookEndpointsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.ProviderAccountId, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.ProviderCode):
			key := string(ProviderWebhookEndpointsDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.ProviderCode, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.EndpointKey):
			key := string(ProviderWebhookEndpointsDTOFieldName.EndpointKey)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.EndpointKey, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.Environment):
			key := string(ProviderWebhookEndpointsDTOFieldName.Environment)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.Environment, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.SecretRef):
			key := string(ProviderWebhookEndpointsDTOFieldName.SecretRef)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.SecretRef, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.SignatureAlgorithm):
			key := string(ProviderWebhookEndpointsDTOFieldName.SignatureAlgorithm)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.SignatureAlgorithm, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.IsActive):
			key := string(ProviderWebhookEndpointsDTOFieldName.IsActive)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.IsActive, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.Metadata):
			key := string(ProviderWebhookEndpointsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.Metadata, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedAt):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaCreatedAt, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaCreatedBy):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaCreatedBy, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedAt):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaUpdatedBy):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaUpdatedBy, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedAt):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderWebhookEndpointsDBFieldName.MetaDeletedBy):
			key := string(ProviderWebhookEndpointsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEndpointsSelectableValue(providerWebhookEndpointsSelectableResponse, key, providerWebhookEndpoints.MetaDeletedBy, explicitAlias)

		}
	}
	return providerWebhookEndpointsSelectableResponse
}

func NewProviderWebhookEndpointsListResponseFromFilterResult(result []model.ProviderWebhookEndpointsFilterResult, filter model.Filter) ProviderWebhookEndpointsSelectableListResponse {
	dtoProviderWebhookEndpointsListResponse := ProviderWebhookEndpointsSelectableListResponse{}
	for _, row := range result {
		dtoProviderWebhookEndpointsResponse := NewProviderWebhookEndpointsSelectableResponse(row.ProviderWebhookEndpoints, filter)
		dtoProviderWebhookEndpointsListResponse = append(dtoProviderWebhookEndpointsListResponse, &dtoProviderWebhookEndpointsResponse)
	}
	return dtoProviderWebhookEndpointsListResponse
}

type ProviderWebhookEndpointsFilterResponse struct {
	Metadata Metadata                                       `json:"metadata"`
	Data     ProviderWebhookEndpointsSelectableListResponse `json:"data"`
}

func reverseProviderWebhookEndpointsFilterResults(result []model.ProviderWebhookEndpointsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderWebhookEndpointsFilterResponse(result []model.ProviderWebhookEndpointsFilterResult, filter model.Filter) (resp ProviderWebhookEndpointsFilterResponse) {
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
			reverseProviderWebhookEndpointsFilterResults(dataResult)
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

	resp.Data = NewProviderWebhookEndpointsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderWebhookEndpointsCreateRequest struct {
	ProviderAccountId  uuid.UUID       `json:"providerAccountId"`
	ProviderCode       string          `json:"providerCode"`
	EndpointKey        string          `json:"endpointKey"`
	Environment        string          `json:"environment"`
	SecretRef          string          `json:"secretRef"`
	SignatureAlgorithm string          `json:"signatureAlgorithm"`
	IsActive           bool            `json:"isActive"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d *ProviderWebhookEndpointsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderWebhookEndpointsCreateRequest) ToModel() model.ProviderWebhookEndpoints {
	id, _ := uuid.NewV4()
	return model.ProviderWebhookEndpoints{
		Id:                 id,
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EndpointKey:        d.EndpointKey,
		Environment:        d.Environment,
		SecretRef:          d.SecretRef,
		SignatureAlgorithm: d.SignatureAlgorithm,
		IsActive:           d.IsActive,
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEndpointsListCreateRequest []*ProviderWebhookEndpointsCreateRequest

func (d ProviderWebhookEndpointsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEndpoints := range d {
		err = validator.Struct(providerWebhookEndpoints)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderWebhookEndpointsListCreateRequest) ToModelList() []model.ProviderWebhookEndpoints {
	out := make([]model.ProviderWebhookEndpoints, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderWebhookEndpointsUpdateRequest struct {
	ProviderAccountId  uuid.UUID       `json:"providerAccountId"`
	ProviderCode       string          `json:"providerCode"`
	EndpointKey        string          `json:"endpointKey"`
	Environment        string          `json:"environment"`
	SecretRef          string          `json:"secretRef"`
	SignatureAlgorithm string          `json:"signatureAlgorithm"`
	IsActive           bool            `json:"isActive"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d *ProviderWebhookEndpointsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderWebhookEndpointsUpdateRequest) ToModel() model.ProviderWebhookEndpoints {
	return model.ProviderWebhookEndpoints{
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EndpointKey:        d.EndpointKey,
		Environment:        d.Environment,
		SecretRef:          d.SecretRef,
		SignatureAlgorithm: d.SignatureAlgorithm,
		IsActive:           d.IsActive,
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEndpointsBulkUpdateRequest struct {
	Id                 uuid.UUID       `json:"id"`
	ProviderAccountId  uuid.UUID       `json:"providerAccountId"`
	ProviderCode       string          `json:"providerCode"`
	EndpointKey        string          `json:"endpointKey"`
	Environment        string          `json:"environment"`
	SecretRef          string          `json:"secretRef"`
	SignatureAlgorithm string          `json:"signatureAlgorithm"`
	IsActive           bool            `json:"isActive"`
	Metadata           json.RawMessage `json:"metadata"`
}

func (d ProviderWebhookEndpointsBulkUpdateRequest) PrimaryID() ProviderWebhookEndpointsPrimaryID {
	return ProviderWebhookEndpointsPrimaryID{
		Id: d.Id,
	}
}

type ProviderWebhookEndpointsListBulkUpdateRequest []*ProviderWebhookEndpointsBulkUpdateRequest

func (d ProviderWebhookEndpointsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEndpoints := range d {
		err = validator.Struct(providerWebhookEndpoints)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderWebhookEndpointsBulkUpdateRequest) ToModel() model.ProviderWebhookEndpoints {
	return model.ProviderWebhookEndpoints{
		Id:                 d.Id,
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EndpointKey:        d.EndpointKey,
		Environment:        d.Environment,
		SecretRef:          d.SecretRef,
		SignatureAlgorithm: d.SignatureAlgorithm,
		IsActive:           d.IsActive,
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEndpointsResponse struct {
	Id                 uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId  uuid.UUID       `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode       string          `json:"providerCode" validate:"required"`
	EndpointKey        string          `json:"endpointKey" validate:"required"`
	Environment        string          `json:"environment"`
	SecretRef          string          `json:"secretRef" validate:"required"`
	SignatureAlgorithm string          `json:"signatureAlgorithm" validate:"required"`
	IsActive           bool            `json:"isActive" example:"true"`
	Metadata           json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewProviderWebhookEndpointsResponse(providerWebhookEndpoints model.ProviderWebhookEndpoints) ProviderWebhookEndpointsResponse {
	return ProviderWebhookEndpointsResponse{
		Id:                 providerWebhookEndpoints.Id,
		ProviderAccountId:  providerWebhookEndpoints.ProviderAccountId,
		ProviderCode:       providerWebhookEndpoints.ProviderCode,
		EndpointKey:        providerWebhookEndpoints.EndpointKey,
		Environment:        providerWebhookEndpoints.Environment,
		SecretRef:          providerWebhookEndpoints.SecretRef,
		SignatureAlgorithm: providerWebhookEndpoints.SignatureAlgorithm,
		IsActive:           providerWebhookEndpoints.IsActive,
		Metadata:           providerWebhookEndpoints.Metadata,
	}
}

type ProviderWebhookEndpointsListResponse []*ProviderWebhookEndpointsResponse

func NewProviderWebhookEndpointsListResponse(providerWebhookEndpointsList model.ProviderWebhookEndpointsList) ProviderWebhookEndpointsListResponse {
	dtoProviderWebhookEndpointsListResponse := ProviderWebhookEndpointsListResponse{}
	for _, providerWebhookEndpoints := range providerWebhookEndpointsList {
		dtoProviderWebhookEndpointsResponse := NewProviderWebhookEndpointsResponse(*providerWebhookEndpoints)
		dtoProviderWebhookEndpointsListResponse = append(dtoProviderWebhookEndpointsListResponse, &dtoProviderWebhookEndpointsResponse)
	}
	return dtoProviderWebhookEndpointsListResponse
}

type ProviderWebhookEndpointsPrimaryIDList []ProviderWebhookEndpointsPrimaryID

func (d ProviderWebhookEndpointsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEndpoints := range d {
		err = validator.Struct(providerWebhookEndpoints)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderWebhookEndpointsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderWebhookEndpointsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderWebhookEndpointsPrimaryID) ToModel() model.ProviderWebhookEndpointsPrimaryID {
	return model.ProviderWebhookEndpointsPrimaryID{
		Id: d.Id,
	}
}

package dto

import (
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RoutingProfilesDTOFieldNameType string

type routingProfilesDTOFieldName struct {
	Id                RoutingProfilesDTOFieldNameType
	MerchantId        RoutingProfilesDTOFieldNameType
	Name              RoutingProfilesDTOFieldNameType
	Strategy          RoutingProfilesDTOFieldNameType
	IsActive          RoutingProfilesDTOFieldNameType
	FallbackProfileId RoutingProfilesDTOFieldNameType
	Notes             RoutingProfilesDTOFieldNameType
	MetaCreatedAt     RoutingProfilesDTOFieldNameType
	MetaCreatedBy     RoutingProfilesDTOFieldNameType
	MetaUpdatedAt     RoutingProfilesDTOFieldNameType
	MetaUpdatedBy     RoutingProfilesDTOFieldNameType
	MetaDeletedAt     RoutingProfilesDTOFieldNameType
	MetaDeletedBy     RoutingProfilesDTOFieldNameType
}

var RoutingProfilesDTOFieldName = routingProfilesDTOFieldName{
	Id:                "id",
	MerchantId:        "merchantId",
	Name:              "name",
	Strategy:          "strategy",
	IsActive:          "isActive",
	FallbackProfileId: "fallbackProfileId",
	Notes:             "notes",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func NewRoutingProfilesListResponseFromFilterResult(result []model.RoutingProfilesFilterResult, filter model.Filter) RoutingProfilesSelectableListResponse {
	dtoRoutingProfilesListResponse := RoutingProfilesSelectableListResponse{}
	for _, routingProfiles := range result {
		dtoRoutingProfilesResponse := NewRoutingProfilesSelectableResponse(routingProfiles.RoutingProfiles, filter)
		dtoRoutingProfilesListResponse = append(dtoRoutingProfilesListResponse, &dtoRoutingProfilesResponse)
	}
	return dtoRoutingProfilesListResponse
}

func transformRoutingProfilesDTOFieldNameFromStr(field string) (dbField model.RoutingProfilesDBFieldNameType, found bool) {
	switch field {

	case string(RoutingProfilesDTOFieldName.Id):
		return model.RoutingProfilesDBFieldName.Id, true

	case string(RoutingProfilesDTOFieldName.MerchantId):
		return model.RoutingProfilesDBFieldName.MerchantId, true

	case string(RoutingProfilesDTOFieldName.Name):
		return model.RoutingProfilesDBFieldName.Name, true

	case string(RoutingProfilesDTOFieldName.Strategy):
		return model.RoutingProfilesDBFieldName.Strategy, true

	case string(RoutingProfilesDTOFieldName.IsActive):
		return model.RoutingProfilesDBFieldName.IsActive, true

	case string(RoutingProfilesDTOFieldName.FallbackProfileId):
		return model.RoutingProfilesDBFieldName.FallbackProfileId, true

	case string(RoutingProfilesDTOFieldName.Notes):
		return model.RoutingProfilesDBFieldName.Notes, true

	case string(RoutingProfilesDTOFieldName.MetaCreatedAt):
		return model.RoutingProfilesDBFieldName.MetaCreatedAt, true

	case string(RoutingProfilesDTOFieldName.MetaCreatedBy):
		return model.RoutingProfilesDBFieldName.MetaCreatedBy, true

	case string(RoutingProfilesDTOFieldName.MetaUpdatedAt):
		return model.RoutingProfilesDBFieldName.MetaUpdatedAt, true

	case string(RoutingProfilesDTOFieldName.MetaUpdatedBy):
		return model.RoutingProfilesDBFieldName.MetaUpdatedBy, true

	case string(RoutingProfilesDTOFieldName.MetaDeletedAt):
		return model.RoutingProfilesDBFieldName.MetaDeletedAt, true

	case string(RoutingProfilesDTOFieldName.MetaDeletedBy):
		return model.RoutingProfilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformRoutingProfilesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformRoutingProfilesDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRoutingProfilesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRoutingProfilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultRoutingProfilesFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(RoutingProfilesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RoutingProfilesSelectableResponse map[string]interface{}
type RoutingProfilesSelectableListResponse []*RoutingProfilesSelectableResponse

func NewRoutingProfilesSelectableResponse(routingProfiles model.RoutingProfiles, filter model.Filter) RoutingProfilesSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RoutingProfilesDBFieldName.Id),
			string(model.RoutingProfilesDBFieldName.MerchantId),
			string(model.RoutingProfilesDBFieldName.Name),
			string(model.RoutingProfilesDBFieldName.Strategy),
			string(model.RoutingProfilesDBFieldName.IsActive),
			string(model.RoutingProfilesDBFieldName.FallbackProfileId),
			string(model.RoutingProfilesDBFieldName.Notes),
			string(model.RoutingProfilesDBFieldName.MetaCreatedAt),
			string(model.RoutingProfilesDBFieldName.MetaCreatedBy),
			string(model.RoutingProfilesDBFieldName.MetaUpdatedAt),
			string(model.RoutingProfilesDBFieldName.MetaUpdatedBy),
			string(model.RoutingProfilesDBFieldName.MetaDeletedAt),
			string(model.RoutingProfilesDBFieldName.MetaDeletedBy),
		)
	}
	routingProfilesSelectableResponse := RoutingProfilesSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.RoutingProfilesDBFieldName.Id):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.Id)] = routingProfiles.Id

		case string(model.RoutingProfilesDBFieldName.MerchantId):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MerchantId)] = routingProfiles.MerchantId

		case string(model.RoutingProfilesDBFieldName.Name):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.Name)] = routingProfiles.Name

		case string(model.RoutingProfilesDBFieldName.Strategy):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.Strategy)] = routingProfiles.Strategy

		case string(model.RoutingProfilesDBFieldName.IsActive):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.IsActive)] = routingProfiles.IsActive

		case string(model.RoutingProfilesDBFieldName.FallbackProfileId):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.FallbackProfileId)] = routingProfiles.FallbackProfileId

		case string(model.RoutingProfilesDBFieldName.Notes):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.Notes)] = routingProfiles.Notes

		case string(model.RoutingProfilesDBFieldName.MetaCreatedAt):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaCreatedAt)] = routingProfiles.MetaCreatedAt

		case string(model.RoutingProfilesDBFieldName.MetaCreatedBy):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaCreatedBy)] = routingProfiles.MetaCreatedBy

		case string(model.RoutingProfilesDBFieldName.MetaUpdatedAt):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaUpdatedAt)] = routingProfiles.MetaUpdatedAt

		case string(model.RoutingProfilesDBFieldName.MetaUpdatedBy):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaUpdatedBy)] = routingProfiles.MetaUpdatedBy

		case string(model.RoutingProfilesDBFieldName.MetaDeletedAt):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaDeletedAt)] = routingProfiles.MetaDeletedAt

		case string(model.RoutingProfilesDBFieldName.MetaDeletedBy):
			routingProfilesSelectableResponse[string(RoutingProfilesDTOFieldName.MetaDeletedBy)] = routingProfiles.MetaDeletedBy

		}
	}
	return routingProfilesSelectableResponse
}

type RoutingProfilesFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     RoutingProfilesSelectableListResponse `json:"data"`
}

func NewRoutingProfilesFilterResponse(result []model.RoutingProfilesFilterResult, filter model.Filter) (resp RoutingProfilesFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewRoutingProfilesListResponseFromFilterResult(result, filter)
	return resp
}

type RoutingProfilesCreateRequest struct {
	MerchantId        uuid.UUID             `json:"merchantId"`
	Name              string                `json:"name"`
	Strategy          model.RoutingStrategy `json:"strategy" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	IsActive          bool                  `json:"isActive"`
	FallbackProfileId uuid.UUID             `json:"fallbackProfileId"`
	Notes             string                `json:"notes"`
	MetaCreatedAt     time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy     uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt     time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy     uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt     time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy     uuid.UUID             `json:"metaDeletedBy"`
}

func (d *RoutingProfilesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RoutingProfilesCreateRequest) ToModel() model.RoutingProfiles {
	id, _ := uuid.NewV4()
	return model.RoutingProfiles{
		Id:                id,
		MerchantId:        nuuid.From(d.MerchantId),
		Name:              d.Name,
		Strategy:          d.Strategy,
		IsActive:          d.IsActive,
		FallbackProfileId: nuuid.From(d.FallbackProfileId),
		Notes:             null.StringFrom(d.Notes),
		MetaCreatedAt:     d.MetaCreatedAt,
		MetaCreatedBy:     d.MetaCreatedBy,
		MetaUpdatedAt:     d.MetaUpdatedAt,
		MetaUpdatedBy:     nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:     null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:     nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingProfilesListCreateRequest []*RoutingProfilesCreateRequest

func (d RoutingProfilesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingProfiles := range d {
		err = validator.Struct(routingProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingProfilesListCreateRequest) ToModelList() []model.RoutingProfiles {
	out := make([]model.RoutingProfiles, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RoutingProfilesUpdateRequest struct {
	MerchantId        uuid.UUID             `json:"merchantId"`
	Name              string                `json:"name"`
	Strategy          model.RoutingStrategy `json:"strategy" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	IsActive          bool                  `json:"isActive"`
	FallbackProfileId uuid.UUID             `json:"fallbackProfileId"`
	Notes             string                `json:"notes"`
	MetaCreatedAt     time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy     uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt     time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy     uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt     time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy     uuid.UUID             `json:"metaDeletedBy"`
}

func (d *RoutingProfilesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RoutingProfilesUpdateRequest) ToModel() model.RoutingProfiles {
	return model.RoutingProfiles{
		MerchantId:        nuuid.From(d.MerchantId),
		Name:              d.Name,
		Strategy:          d.Strategy,
		IsActive:          d.IsActive,
		FallbackProfileId: nuuid.From(d.FallbackProfileId),
		Notes:             null.StringFrom(d.Notes),
		MetaCreatedAt:     d.MetaCreatedAt,
		MetaCreatedBy:     d.MetaCreatedBy,
		MetaUpdatedAt:     d.MetaUpdatedAt,
		MetaUpdatedBy:     nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:     null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:     nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingProfilesBulkUpdateRequest struct {
	Id                uuid.UUID             `json:"id"`
	MerchantId        uuid.UUID             `json:"merchantId"`
	Name              string                `json:"name"`
	Strategy          model.RoutingStrategy `json:"strategy" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	IsActive          bool                  `json:"isActive"`
	FallbackProfileId uuid.UUID             `json:"fallbackProfileId"`
	Notes             string                `json:"notes"`
	MetaCreatedAt     time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy     uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt     time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy     uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt     time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy     uuid.UUID             `json:"metaDeletedBy"`
}

func (d RoutingProfilesBulkUpdateRequest) PrimaryID() RoutingProfilesPrimaryID {
	return RoutingProfilesPrimaryID{
		Id: d.Id,
	}
}

type RoutingProfilesListBulkUpdateRequest []*RoutingProfilesBulkUpdateRequest

func (d RoutingProfilesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingProfiles := range d {
		err = validator.Struct(routingProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingProfilesBulkUpdateRequest) ToModel() model.RoutingProfiles {
	return model.RoutingProfiles{
		Id:                d.Id,
		MerchantId:        nuuid.From(d.MerchantId),
		Name:              d.Name,
		Strategy:          d.Strategy,
		IsActive:          d.IsActive,
		FallbackProfileId: nuuid.From(d.FallbackProfileId),
		Notes:             null.StringFrom(d.Notes),
		MetaCreatedAt:     d.MetaCreatedAt,
		MetaCreatedBy:     d.MetaCreatedBy,
		MetaUpdatedAt:     d.MetaUpdatedAt,
		MetaUpdatedBy:     nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:     null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:     nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingProfilesResponse struct {
	Id                uuid.UUID             `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantId        uuid.UUID             `json:"merchantId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name              string                `json:"name" validate:"required"`
	Strategy          model.RoutingStrategy `json:"strategy" validate:"oneof=LOWEST_COST HIGHEST_SUCCESS_RATE ROUND_ROBIN GEO_PREFERRED MANUAL WATERFALL" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	IsActive          bool                  `json:"isActive" example:"true"`
	FallbackProfileId uuid.UUID             `json:"fallbackProfileId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Notes             string                `json:"notes"`
	MetaCreatedAt     time.Time             `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy     uuid.UUID             `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt     time.Time             `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy     uuid.UUID             `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt     time.Time             `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy     uuid.UUID             `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewRoutingProfilesResponse(routingProfiles model.RoutingProfiles) RoutingProfilesResponse {
	return RoutingProfilesResponse{
		Id:                routingProfiles.Id,
		MerchantId:        routingProfiles.MerchantId.UUID,
		Name:              routingProfiles.Name,
		Strategy:          model.RoutingStrategy(routingProfiles.Strategy),
		IsActive:          routingProfiles.IsActive,
		FallbackProfileId: routingProfiles.FallbackProfileId.UUID,
		Notes:             routingProfiles.Notes.String,
		MetaCreatedAt:     routingProfiles.MetaCreatedAt,
		MetaCreatedBy:     routingProfiles.MetaCreatedBy,
		MetaUpdatedAt:     routingProfiles.MetaUpdatedAt,
		MetaUpdatedBy:     routingProfiles.MetaUpdatedBy.UUID,
		MetaDeletedAt:     routingProfiles.MetaDeletedAt.Time,
		MetaDeletedBy:     routingProfiles.MetaDeletedBy.UUID,
	}
}

type RoutingProfilesListResponse []*RoutingProfilesResponse

func NewRoutingProfilesListResponse(routingProfilesList model.RoutingProfilesList) RoutingProfilesListResponse {
	dtoRoutingProfilesListResponse := RoutingProfilesListResponse{}
	for _, routingProfiles := range routingProfilesList {
		dtoRoutingProfilesResponse := NewRoutingProfilesResponse(*routingProfiles)
		dtoRoutingProfilesListResponse = append(dtoRoutingProfilesListResponse, &dtoRoutingProfilesResponse)
	}
	return dtoRoutingProfilesListResponse
}

type RoutingProfilesPrimaryIDList []RoutingProfilesPrimaryID

func (d RoutingProfilesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingProfiles := range d {
		err = validator.Struct(routingProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

type RoutingProfilesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RoutingProfilesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RoutingProfilesPrimaryID) ToModel() model.RoutingProfilesPrimaryID {
	return model.RoutingProfilesPrimaryID{
		Id: d.Id,
	}
}

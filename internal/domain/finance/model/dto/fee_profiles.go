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

type FeeProfilesDTOFieldNameType string

type feeProfilesDTOFieldName struct {
	Id              FeeProfilesDTOFieldNameType
	ProfileCode     FeeProfilesDTOFieldNameType
	OwnerPartyId    FeeProfilesDTOFieldNameType
	ProfileScope    FeeProfilesDTOFieldNameType
	EffectiveStatus FeeProfilesDTOFieldNameType
	Description     FeeProfilesDTOFieldNameType
	Metadata        FeeProfilesDTOFieldNameType
	MetaCreatedAt   FeeProfilesDTOFieldNameType
	MetaCreatedBy   FeeProfilesDTOFieldNameType
	MetaUpdatedAt   FeeProfilesDTOFieldNameType
	MetaUpdatedBy   FeeProfilesDTOFieldNameType
	MetaDeletedAt   FeeProfilesDTOFieldNameType
	MetaDeletedBy   FeeProfilesDTOFieldNameType
}

var FeeProfilesDTOFieldName = feeProfilesDTOFieldName{
	Id:              "id",
	ProfileCode:     "profileCode",
	OwnerPartyId:    "ownerPartyId",
	ProfileScope:    "profileScope",
	EffectiveStatus: "effectiveStatus",
	Description:     "description",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformFeeProfilesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FeeProfilesDTOFieldName.Id):
		return string(model.FeeProfilesDBFieldName.Id), true

	case string(FeeProfilesDTOFieldName.ProfileCode):
		return string(model.FeeProfilesDBFieldName.ProfileCode), true

	case string(FeeProfilesDTOFieldName.OwnerPartyId):
		return string(model.FeeProfilesDBFieldName.OwnerPartyId), true

	case string(FeeProfilesDTOFieldName.ProfileScope):
		return string(model.FeeProfilesDBFieldName.ProfileScope), true

	case string(FeeProfilesDTOFieldName.EffectiveStatus):
		return string(model.FeeProfilesDBFieldName.EffectiveStatus), true

	case string(FeeProfilesDTOFieldName.Description):
		return string(model.FeeProfilesDBFieldName.Description), true

	case string(FeeProfilesDTOFieldName.Metadata):
		return string(model.FeeProfilesDBFieldName.Metadata), true

	case string(FeeProfilesDTOFieldName.MetaCreatedAt):
		return string(model.FeeProfilesDBFieldName.MetaCreatedAt), true

	case string(FeeProfilesDTOFieldName.MetaCreatedBy):
		return string(model.FeeProfilesDBFieldName.MetaCreatedBy), true

	case string(FeeProfilesDTOFieldName.MetaUpdatedAt):
		return string(model.FeeProfilesDBFieldName.MetaUpdatedAt), true

	case string(FeeProfilesDTOFieldName.MetaUpdatedBy):
		return string(model.FeeProfilesDBFieldName.MetaUpdatedBy), true

	case string(FeeProfilesDTOFieldName.MetaDeletedAt):
		return string(model.FeeProfilesDBFieldName.MetaDeletedAt), true

	case string(FeeProfilesDTOFieldName.MetaDeletedBy):
		return string(model.FeeProfilesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFeeProfilesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFeeProfilesBaseFilterField(field string) bool {
	spec, found := model.NewFeeProfilesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFeeProfilesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFeeProfilesProjectionOutputPath(path string) error {
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

func transformFeeProfilesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFeeProfilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFeeProfilesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFeeProfilesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFeeProfilesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFeeProfilesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFeeProfilesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFeeProfilesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFeeProfilesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFeeProfilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFeeProfilesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFeeProfilesFilter(filter *model.Filter) {
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
			Field: string(FeeProfilesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FeeProfilesSelectableResponse map[string]interface{}
type FeeProfilesSelectableListResponse []*FeeProfilesSelectableResponse

func assignFeeProfilesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFeeProfilesSelectableValue(out FeeProfilesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFeeProfilesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFeeProfilesSelectableResponse(feeProfiles model.FeeProfiles, filter model.Filter) FeeProfilesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FeeProfilesDBFieldName.Id),
			string(model.FeeProfilesDBFieldName.ProfileCode),
			string(model.FeeProfilesDBFieldName.OwnerPartyId),
			string(model.FeeProfilesDBFieldName.ProfileScope),
			string(model.FeeProfilesDBFieldName.EffectiveStatus),
			string(model.FeeProfilesDBFieldName.Description),
			string(model.FeeProfilesDBFieldName.Metadata),
			string(model.FeeProfilesDBFieldName.MetaCreatedAt),
			string(model.FeeProfilesDBFieldName.MetaCreatedBy),
			string(model.FeeProfilesDBFieldName.MetaUpdatedAt),
			string(model.FeeProfilesDBFieldName.MetaUpdatedBy),
			string(model.FeeProfilesDBFieldName.MetaDeletedAt),
			string(model.FeeProfilesDBFieldName.MetaDeletedBy),
		)
	}
	feeProfilesSelectableResponse := FeeProfilesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FeeProfilesDBFieldName.Id):
			key := string(FeeProfilesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.Id, explicitAlias)

		case string(model.FeeProfilesDBFieldName.ProfileCode):
			key := string(FeeProfilesDTOFieldName.ProfileCode)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.ProfileCode, explicitAlias)

		case string(model.FeeProfilesDBFieldName.OwnerPartyId):
			key := string(FeeProfilesDTOFieldName.OwnerPartyId)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.OwnerPartyId.UUID, explicitAlias)

		case string(model.FeeProfilesDBFieldName.ProfileScope):
			key := string(FeeProfilesDTOFieldName.ProfileScope)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, model.ProfileScope(feeProfiles.ProfileScope), explicitAlias)

		case string(model.FeeProfilesDBFieldName.EffectiveStatus):
			key := string(FeeProfilesDTOFieldName.EffectiveStatus)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, model.EffectiveStatus(feeProfiles.EffectiveStatus), explicitAlias)

		case string(model.FeeProfilesDBFieldName.Description):
			key := string(FeeProfilesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.Description.String, explicitAlias)

		case string(model.FeeProfilesDBFieldName.Metadata):
			key := string(FeeProfilesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.Metadata, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaCreatedAt):
			key := string(FeeProfilesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaCreatedAt, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaCreatedBy):
			key := string(FeeProfilesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaCreatedBy, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaUpdatedAt):
			key := string(FeeProfilesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaUpdatedAt, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaUpdatedBy):
			key := string(FeeProfilesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaUpdatedBy, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaDeletedAt):
			key := string(FeeProfilesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaDeletedAt.Time, explicitAlias)

		case string(model.FeeProfilesDBFieldName.MetaDeletedBy):
			key := string(FeeProfilesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFeeProfilesSelectableValue(feeProfilesSelectableResponse, key, feeProfiles.MetaDeletedBy, explicitAlias)

		}
	}
	return feeProfilesSelectableResponse
}

func NewFeeProfilesListResponseFromFilterResult(result []model.FeeProfilesFilterResult, filter model.Filter) FeeProfilesSelectableListResponse {
	dtoFeeProfilesListResponse := FeeProfilesSelectableListResponse{}
	for _, row := range result {
		dtoFeeProfilesResponse := NewFeeProfilesSelectableResponse(row.FeeProfiles, filter)
		dtoFeeProfilesListResponse = append(dtoFeeProfilesListResponse, &dtoFeeProfilesResponse)
	}
	return dtoFeeProfilesListResponse
}

type FeeProfilesFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     FeeProfilesSelectableListResponse `json:"data"`
}

func reverseFeeProfilesFilterResults(result []model.FeeProfilesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFeeProfilesFilterResponse(result []model.FeeProfilesFilterResult, filter model.Filter) (resp FeeProfilesFilterResponse) {
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
			reverseFeeProfilesFilterResults(dataResult)
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

	resp.Data = NewFeeProfilesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FeeProfilesCreateRequest struct {
	ProfileCode     string                `json:"profileCode"`
	OwnerPartyId    uuid.UUID             `json:"ownerPartyId"`
	ProfileScope    model.ProfileScope    `json:"profileScope" example:"platform" enums:"platform,merchant,listing,category,campaign"`
	EffectiveStatus model.EffectiveStatus `json:"effectiveStatus" example:"active" enums:"active,inactive"`
	Description     string                `json:"description"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d *FeeProfilesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FeeProfilesCreateRequest) ToModel() model.FeeProfiles {
	id, _ := uuid.NewV4()
	return model.FeeProfiles{
		Id:              id,
		ProfileCode:     d.ProfileCode,
		OwnerPartyId:    nuuid.From(d.OwnerPartyId),
		ProfileScope:    d.ProfileScope,
		EffectiveStatus: d.EffectiveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type FeeProfilesListCreateRequest []*FeeProfilesCreateRequest

func (d FeeProfilesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeProfiles := range d {
		err = validator.Struct(feeProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeProfilesListCreateRequest) ToModelList() []model.FeeProfiles {
	out := make([]model.FeeProfiles, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FeeProfilesUpdateRequest struct {
	ProfileCode     string                `json:"profileCode"`
	OwnerPartyId    uuid.UUID             `json:"ownerPartyId"`
	ProfileScope    model.ProfileScope    `json:"profileScope" example:"platform" enums:"platform,merchant,listing,category,campaign"`
	EffectiveStatus model.EffectiveStatus `json:"effectiveStatus" example:"active" enums:"active,inactive"`
	Description     string                `json:"description"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d *FeeProfilesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FeeProfilesUpdateRequest) ToModel() model.FeeProfiles {
	return model.FeeProfiles{
		ProfileCode:     d.ProfileCode,
		OwnerPartyId:    nuuid.From(d.OwnerPartyId),
		ProfileScope:    d.ProfileScope,
		EffectiveStatus: d.EffectiveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type FeeProfilesBulkUpdateRequest struct {
	Id              uuid.UUID             `json:"id"`
	ProfileCode     string                `json:"profileCode"`
	OwnerPartyId    uuid.UUID             `json:"ownerPartyId"`
	ProfileScope    model.ProfileScope    `json:"profileScope" example:"platform" enums:"platform,merchant,listing,category,campaign"`
	EffectiveStatus model.EffectiveStatus `json:"effectiveStatus" example:"active" enums:"active,inactive"`
	Description     string                `json:"description"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d FeeProfilesBulkUpdateRequest) PrimaryID() FeeProfilesPrimaryID {
	return FeeProfilesPrimaryID{
		Id: d.Id,
	}
}

type FeeProfilesListBulkUpdateRequest []*FeeProfilesBulkUpdateRequest

func (d FeeProfilesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeProfiles := range d {
		err = validator.Struct(feeProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FeeProfilesBulkUpdateRequest) ToModel() model.FeeProfiles {
	return model.FeeProfiles{
		Id:              d.Id,
		ProfileCode:     d.ProfileCode,
		OwnerPartyId:    nuuid.From(d.OwnerPartyId),
		ProfileScope:    d.ProfileScope,
		EffectiveStatus: d.EffectiveStatus,
		Description:     null.StringFrom(d.Description),
		Metadata:        d.Metadata,
	}
}

type FeeProfilesResponse struct {
	Id              uuid.UUID             `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProfileCode     string                `json:"profileCode" validate:"required"`
	OwnerPartyId    uuid.UUID             `json:"ownerPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProfileScope    model.ProfileScope    `json:"profileScope" validate:"required,oneof=platform merchant listing category campaign" enums:"platform,merchant,listing,category,campaign"`
	EffectiveStatus model.EffectiveStatus `json:"effectiveStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Description     string                `json:"description"`
	Metadata        json.RawMessage       `json:"metadata" swaggertype:"object"`
}

func NewFeeProfilesResponse(feeProfiles model.FeeProfiles) FeeProfilesResponse {
	return FeeProfilesResponse{
		Id:              feeProfiles.Id,
		ProfileCode:     feeProfiles.ProfileCode,
		OwnerPartyId:    feeProfiles.OwnerPartyId.UUID,
		ProfileScope:    model.ProfileScope(feeProfiles.ProfileScope),
		EffectiveStatus: model.EffectiveStatus(feeProfiles.EffectiveStatus),
		Description:     feeProfiles.Description.String,
		Metadata:        feeProfiles.Metadata,
	}
}

type FeeProfilesListResponse []*FeeProfilesResponse

func NewFeeProfilesListResponse(feeProfilesList model.FeeProfilesList) FeeProfilesListResponse {
	dtoFeeProfilesListResponse := FeeProfilesListResponse{}
	for _, feeProfiles := range feeProfilesList {
		dtoFeeProfilesResponse := NewFeeProfilesResponse(*feeProfiles)
		dtoFeeProfilesListResponse = append(dtoFeeProfilesListResponse, &dtoFeeProfilesResponse)
	}
	return dtoFeeProfilesListResponse
}

type FeeProfilesPrimaryIDList []FeeProfilesPrimaryID

func (d FeeProfilesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, feeProfiles := range d {
		err = validator.Struct(feeProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

type FeeProfilesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FeeProfilesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FeeProfilesPrimaryID) ToModel() model.FeeProfilesPrimaryID {
	return model.FeeProfilesPrimaryID{
		Id: d.Id,
	}
}

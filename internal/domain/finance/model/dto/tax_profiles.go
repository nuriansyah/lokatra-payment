package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type TaxProfilesDTOFieldNameType string

type taxProfilesDTOFieldName struct {
	Id                   TaxProfilesDTOFieldNameType
	OwnerPartyId         TaxProfilesDTOFieldNameType
	CountryCode          TaxProfilesDTOFieldNameType
	TaxResidencyCountry  TaxProfilesDTOFieldNameType
	TaxIdMasked          TaxProfilesDTOFieldNameType
	TaxEntityType        TaxProfilesDTOFieldNameType
	IsVatRegistered      TaxProfilesDTOFieldNameType
	IsWithholdingSubject TaxProfilesDTOFieldNameType
	ProfileStatus        TaxProfilesDTOFieldNameType
	Metadata             TaxProfilesDTOFieldNameType
	MetaCreatedAt        TaxProfilesDTOFieldNameType
	MetaCreatedBy        TaxProfilesDTOFieldNameType
	MetaUpdatedAt        TaxProfilesDTOFieldNameType
	MetaUpdatedBy        TaxProfilesDTOFieldNameType
	MetaDeletedAt        TaxProfilesDTOFieldNameType
	MetaDeletedBy        TaxProfilesDTOFieldNameType
}

var TaxProfilesDTOFieldName = taxProfilesDTOFieldName{
	Id:                   "id",
	OwnerPartyId:         "ownerPartyId",
	CountryCode:          "countryCode",
	TaxResidencyCountry:  "taxResidencyCountry",
	TaxIdMasked:          "taxIdMasked",
	TaxEntityType:        "taxEntityType",
	IsVatRegistered:      "isVatRegistered",
	IsWithholdingSubject: "isWithholdingSubject",
	ProfileStatus:        "profileStatus",
	Metadata:             "metadata",
	MetaCreatedAt:        "metaCreatedAt",
	MetaCreatedBy:        "metaCreatedBy",
	MetaUpdatedAt:        "metaUpdatedAt",
	MetaUpdatedBy:        "metaUpdatedBy",
	MetaDeletedAt:        "metaDeletedAt",
	MetaDeletedBy:        "metaDeletedBy",
}

func transformTaxProfilesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxProfilesDTOFieldName.Id):
		return string(model.TaxProfilesDBFieldName.Id), true

	case string(TaxProfilesDTOFieldName.OwnerPartyId):
		return string(model.TaxProfilesDBFieldName.OwnerPartyId), true

	case string(TaxProfilesDTOFieldName.CountryCode):
		return string(model.TaxProfilesDBFieldName.CountryCode), true

	case string(TaxProfilesDTOFieldName.TaxResidencyCountry):
		return string(model.TaxProfilesDBFieldName.TaxResidencyCountry), true

	case string(TaxProfilesDTOFieldName.TaxIdMasked):
		return string(model.TaxProfilesDBFieldName.TaxIdMasked), true

	case string(TaxProfilesDTOFieldName.TaxEntityType):
		return string(model.TaxProfilesDBFieldName.TaxEntityType), true

	case string(TaxProfilesDTOFieldName.IsVatRegistered):
		return string(model.TaxProfilesDBFieldName.IsVatRegistered), true

	case string(TaxProfilesDTOFieldName.IsWithholdingSubject):
		return string(model.TaxProfilesDBFieldName.IsWithholdingSubject), true

	case string(TaxProfilesDTOFieldName.ProfileStatus):
		return string(model.TaxProfilesDBFieldName.ProfileStatus), true

	case string(TaxProfilesDTOFieldName.Metadata):
		return string(model.TaxProfilesDBFieldName.Metadata), true

	case string(TaxProfilesDTOFieldName.MetaCreatedAt):
		return string(model.TaxProfilesDBFieldName.MetaCreatedAt), true

	case string(TaxProfilesDTOFieldName.MetaCreatedBy):
		return string(model.TaxProfilesDBFieldName.MetaCreatedBy), true

	case string(TaxProfilesDTOFieldName.MetaUpdatedAt):
		return string(model.TaxProfilesDBFieldName.MetaUpdatedAt), true

	case string(TaxProfilesDTOFieldName.MetaUpdatedBy):
		return string(model.TaxProfilesDBFieldName.MetaUpdatedBy), true

	case string(TaxProfilesDTOFieldName.MetaDeletedAt):
		return string(model.TaxProfilesDBFieldName.MetaDeletedAt), true

	case string(TaxProfilesDTOFieldName.MetaDeletedBy):
		return string(model.TaxProfilesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxProfilesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxProfilesBaseFilterField(field string) bool {
	spec, found := model.NewTaxProfilesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxProfilesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxProfilesProjectionOutputPath(path string) error {
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

func transformTaxProfilesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxProfilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxProfilesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxProfilesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxProfilesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxProfilesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxProfilesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxProfilesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxProfilesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxProfilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxProfilesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxProfilesFilter(filter *model.Filter) {
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
			Field: string(TaxProfilesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type TaxProfilesSelectableResponse map[string]interface{}
type TaxProfilesSelectableListResponse []*TaxProfilesSelectableResponse

func assignTaxProfilesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxProfilesSelectableValue(out TaxProfilesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxProfilesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxProfilesSelectableResponse(taxProfiles model.TaxProfiles, filter model.Filter) TaxProfilesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxProfilesDBFieldName.Id),
			string(model.TaxProfilesDBFieldName.OwnerPartyId),
			string(model.TaxProfilesDBFieldName.CountryCode),
			string(model.TaxProfilesDBFieldName.TaxResidencyCountry),
			string(model.TaxProfilesDBFieldName.TaxIdMasked),
			string(model.TaxProfilesDBFieldName.TaxEntityType),
			string(model.TaxProfilesDBFieldName.IsVatRegistered),
			string(model.TaxProfilesDBFieldName.IsWithholdingSubject),
			string(model.TaxProfilesDBFieldName.ProfileStatus),
			string(model.TaxProfilesDBFieldName.Metadata),
			string(model.TaxProfilesDBFieldName.MetaCreatedAt),
			string(model.TaxProfilesDBFieldName.MetaCreatedBy),
			string(model.TaxProfilesDBFieldName.MetaUpdatedAt),
			string(model.TaxProfilesDBFieldName.MetaUpdatedBy),
			string(model.TaxProfilesDBFieldName.MetaDeletedAt),
			string(model.TaxProfilesDBFieldName.MetaDeletedBy),
		)
	}
	taxProfilesSelectableResponse := TaxProfilesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxProfilesDBFieldName.Id):
			key := string(TaxProfilesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.Id, explicitAlias)

		case string(model.TaxProfilesDBFieldName.OwnerPartyId):
			key := string(TaxProfilesDTOFieldName.OwnerPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.OwnerPartyId, explicitAlias)

		case string(model.TaxProfilesDBFieldName.CountryCode):
			key := string(TaxProfilesDTOFieldName.CountryCode)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.CountryCode, explicitAlias)

		case string(model.TaxProfilesDBFieldName.TaxResidencyCountry):
			key := string(TaxProfilesDTOFieldName.TaxResidencyCountry)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.TaxResidencyCountry.String, explicitAlias)

		case string(model.TaxProfilesDBFieldName.TaxIdMasked):
			key := string(TaxProfilesDTOFieldName.TaxIdMasked)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.TaxIdMasked.String, explicitAlias)

		case string(model.TaxProfilesDBFieldName.TaxEntityType):
			key := string(TaxProfilesDTOFieldName.TaxEntityType)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.TaxEntityType.String, explicitAlias)

		case string(model.TaxProfilesDBFieldName.IsVatRegistered):
			key := string(TaxProfilesDTOFieldName.IsVatRegistered)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.IsVatRegistered, explicitAlias)

		case string(model.TaxProfilesDBFieldName.IsWithholdingSubject):
			key := string(TaxProfilesDTOFieldName.IsWithholdingSubject)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.IsWithholdingSubject, explicitAlias)

		case string(model.TaxProfilesDBFieldName.ProfileStatus):
			key := string(TaxProfilesDTOFieldName.ProfileStatus)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, model.ProfileStatus(taxProfiles.ProfileStatus), explicitAlias)

		case string(model.TaxProfilesDBFieldName.Metadata):
			key := string(TaxProfilesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.Metadata, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaCreatedAt):
			key := string(TaxProfilesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaCreatedAt, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaCreatedBy):
			key := string(TaxProfilesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaCreatedBy, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaUpdatedAt):
			key := string(TaxProfilesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaUpdatedAt, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaUpdatedBy):
			key := string(TaxProfilesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaUpdatedBy, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaDeletedAt):
			key := string(TaxProfilesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxProfilesDBFieldName.MetaDeletedBy):
			key := string(TaxProfilesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxProfilesSelectableValue(taxProfilesSelectableResponse, key, taxProfiles.MetaDeletedBy, explicitAlias)

		}
	}
	return taxProfilesSelectableResponse
}

func NewTaxProfilesListResponseFromFilterResult(result []model.TaxProfilesFilterResult, filter model.Filter) TaxProfilesSelectableListResponse {
	dtoTaxProfilesListResponse := TaxProfilesSelectableListResponse{}
	for _, row := range result {
		dtoTaxProfilesResponse := NewTaxProfilesSelectableResponse(row.TaxProfiles, filter)
		dtoTaxProfilesListResponse = append(dtoTaxProfilesListResponse, &dtoTaxProfilesResponse)
	}
	return dtoTaxProfilesListResponse
}

type TaxProfilesFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     TaxProfilesSelectableListResponse `json:"data"`
}

func reverseTaxProfilesFilterResults(result []model.TaxProfilesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxProfilesFilterResponse(result []model.TaxProfilesFilterResult, filter model.Filter) (resp TaxProfilesFilterResponse) {
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
			reverseTaxProfilesFilterResults(dataResult)
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

	resp.Data = NewTaxProfilesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxProfilesCreateRequest struct {
	OwnerPartyId         uuid.UUID           `json:"ownerPartyId"`
	CountryCode          string              `json:"countryCode"`
	TaxResidencyCountry  string              `json:"taxResidencyCountry"`
	TaxIdMasked          string              `json:"taxIdMasked"`
	TaxEntityType        string              `json:"taxEntityType"`
	IsVatRegistered      bool                `json:"isVatRegistered"`
	IsWithholdingSubject bool                `json:"isWithholdingSubject"`
	ProfileStatus        model.ProfileStatus `json:"profileStatus" example:"active" enums:"active,inactive"`
	Metadata             json.RawMessage     `json:"metadata"`
}

func (d *TaxProfilesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxProfilesCreateRequest) ToModel() model.TaxProfiles {
	id, _ := uuid.NewV4()
	return model.TaxProfiles{
		Id:                   id,
		OwnerPartyId:         d.OwnerPartyId,
		CountryCode:          d.CountryCode,
		TaxResidencyCountry:  null.StringFrom(d.TaxResidencyCountry),
		TaxIdMasked:          null.StringFrom(d.TaxIdMasked),
		TaxEntityType:        null.StringFrom(d.TaxEntityType),
		IsVatRegistered:      d.IsVatRegistered,
		IsWithholdingSubject: d.IsWithholdingSubject,
		ProfileStatus:        d.ProfileStatus,
		Metadata:             d.Metadata,
	}
}

type TaxProfilesListCreateRequest []*TaxProfilesCreateRequest

func (d TaxProfilesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxProfiles := range d {
		err = validator.Struct(taxProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxProfilesListCreateRequest) ToModelList() []model.TaxProfiles {
	out := make([]model.TaxProfiles, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxProfilesUpdateRequest struct {
	OwnerPartyId         uuid.UUID           `json:"ownerPartyId"`
	CountryCode          string              `json:"countryCode"`
	TaxResidencyCountry  string              `json:"taxResidencyCountry"`
	TaxIdMasked          string              `json:"taxIdMasked"`
	TaxEntityType        string              `json:"taxEntityType"`
	IsVatRegistered      bool                `json:"isVatRegistered"`
	IsWithholdingSubject bool                `json:"isWithholdingSubject"`
	ProfileStatus        model.ProfileStatus `json:"profileStatus" example:"active" enums:"active,inactive"`
	Metadata             json.RawMessage     `json:"metadata"`
}

func (d *TaxProfilesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxProfilesUpdateRequest) ToModel() model.TaxProfiles {
	return model.TaxProfiles{
		OwnerPartyId:         d.OwnerPartyId,
		CountryCode:          d.CountryCode,
		TaxResidencyCountry:  null.StringFrom(d.TaxResidencyCountry),
		TaxIdMasked:          null.StringFrom(d.TaxIdMasked),
		TaxEntityType:        null.StringFrom(d.TaxEntityType),
		IsVatRegistered:      d.IsVatRegistered,
		IsWithholdingSubject: d.IsWithholdingSubject,
		ProfileStatus:        d.ProfileStatus,
		Metadata:             d.Metadata,
	}
}

type TaxProfilesBulkUpdateRequest struct {
	Id                   uuid.UUID           `json:"id"`
	OwnerPartyId         uuid.UUID           `json:"ownerPartyId"`
	CountryCode          string              `json:"countryCode"`
	TaxResidencyCountry  string              `json:"taxResidencyCountry"`
	TaxIdMasked          string              `json:"taxIdMasked"`
	TaxEntityType        string              `json:"taxEntityType"`
	IsVatRegistered      bool                `json:"isVatRegistered"`
	IsWithholdingSubject bool                `json:"isWithholdingSubject"`
	ProfileStatus        model.ProfileStatus `json:"profileStatus" example:"active" enums:"active,inactive"`
	Metadata             json.RawMessage     `json:"metadata"`
}

func (d TaxProfilesBulkUpdateRequest) PrimaryID() TaxProfilesPrimaryID {
	return TaxProfilesPrimaryID{
		Id: d.Id,
	}
}

type TaxProfilesListBulkUpdateRequest []*TaxProfilesBulkUpdateRequest

func (d TaxProfilesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxProfiles := range d {
		err = validator.Struct(taxProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxProfilesBulkUpdateRequest) ToModel() model.TaxProfiles {
	return model.TaxProfiles{
		Id:                   d.Id,
		OwnerPartyId:         d.OwnerPartyId,
		CountryCode:          d.CountryCode,
		TaxResidencyCountry:  null.StringFrom(d.TaxResidencyCountry),
		TaxIdMasked:          null.StringFrom(d.TaxIdMasked),
		TaxEntityType:        null.StringFrom(d.TaxEntityType),
		IsVatRegistered:      d.IsVatRegistered,
		IsWithholdingSubject: d.IsWithholdingSubject,
		ProfileStatus:        d.ProfileStatus,
		Metadata:             d.Metadata,
	}
}

type TaxProfilesResponse struct {
	Id                   uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OwnerPartyId         uuid.UUID           `json:"ownerPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CountryCode          string              `json:"countryCode"`
	TaxResidencyCountry  string              `json:"taxResidencyCountry"`
	TaxIdMasked          string              `json:"taxIdMasked"`
	TaxEntityType        string              `json:"taxEntityType"`
	IsVatRegistered      bool                `json:"isVatRegistered" example:"true"`
	IsWithholdingSubject bool                `json:"isWithholdingSubject" example:"true"`
	ProfileStatus        model.ProfileStatus `json:"profileStatus" validate:"oneof=active inactive" enums:"active,inactive"`
	Metadata             json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewTaxProfilesResponse(taxProfiles model.TaxProfiles) TaxProfilesResponse {
	return TaxProfilesResponse{
		Id:                   taxProfiles.Id,
		OwnerPartyId:         taxProfiles.OwnerPartyId,
		CountryCode:          taxProfiles.CountryCode,
		TaxResidencyCountry:  taxProfiles.TaxResidencyCountry.String,
		TaxIdMasked:          taxProfiles.TaxIdMasked.String,
		TaxEntityType:        taxProfiles.TaxEntityType.String,
		IsVatRegistered:      taxProfiles.IsVatRegistered,
		IsWithholdingSubject: taxProfiles.IsWithholdingSubject,
		ProfileStatus:        model.ProfileStatus(taxProfiles.ProfileStatus),
		Metadata:             taxProfiles.Metadata,
	}
}

type TaxProfilesListResponse []*TaxProfilesResponse

func NewTaxProfilesListResponse(taxProfilesList model.TaxProfilesList) TaxProfilesListResponse {
	dtoTaxProfilesListResponse := TaxProfilesListResponse{}
	for _, taxProfiles := range taxProfilesList {
		dtoTaxProfilesResponse := NewTaxProfilesResponse(*taxProfiles)
		dtoTaxProfilesListResponse = append(dtoTaxProfilesListResponse, &dtoTaxProfilesResponse)
	}
	return dtoTaxProfilesListResponse
}

type TaxProfilesPrimaryIDList []TaxProfilesPrimaryID

func (d TaxProfilesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxProfiles := range d {
		err = validator.Struct(taxProfiles)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxProfilesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *TaxProfilesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxProfilesPrimaryID) ToModel() model.TaxProfilesPrimaryID {
	return model.TaxProfilesPrimaryID{
		Id: d.Id,
	}
}

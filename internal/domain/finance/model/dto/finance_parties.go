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

type FinancePartiesDTOFieldNameType string

type financePartiesDTOFieldName struct {
	Id            FinancePartiesDTOFieldNameType
	PartyType     FinancePartiesDTOFieldNameType
	ExternalRef   FinancePartiesDTOFieldNameType
	LegalName     FinancePartiesDTOFieldNameType
	DisplayName   FinancePartiesDTOFieldNameType
	CountryCode   FinancePartiesDTOFieldNameType
	Email         FinancePartiesDTOFieldNameType
	Phone         FinancePartiesDTOFieldNameType
	PartyStatus   FinancePartiesDTOFieldNameType
	Metadata      FinancePartiesDTOFieldNameType
	MetaCreatedAt FinancePartiesDTOFieldNameType
	MetaCreatedBy FinancePartiesDTOFieldNameType
	MetaUpdatedAt FinancePartiesDTOFieldNameType
	MetaUpdatedBy FinancePartiesDTOFieldNameType
	MetaDeletedAt FinancePartiesDTOFieldNameType
	MetaDeletedBy FinancePartiesDTOFieldNameType
}

var FinancePartiesDTOFieldName = financePartiesDTOFieldName{
	Id:            "id",
	PartyType:     "partyType",
	ExternalRef:   "externalRef",
	LegalName:     "legalName",
	DisplayName:   "displayName",
	CountryCode:   "countryCode",
	Email:         "email",
	Phone:         "phone",
	PartyStatus:   "partyStatus",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformFinancePartiesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinancePartiesDTOFieldName.Id):
		return string(model.FinancePartiesDBFieldName.Id), true

	case string(FinancePartiesDTOFieldName.PartyType):
		return string(model.FinancePartiesDBFieldName.PartyType), true

	case string(FinancePartiesDTOFieldName.ExternalRef):
		return string(model.FinancePartiesDBFieldName.ExternalRef), true

	case string(FinancePartiesDTOFieldName.LegalName):
		return string(model.FinancePartiesDBFieldName.LegalName), true

	case string(FinancePartiesDTOFieldName.DisplayName):
		return string(model.FinancePartiesDBFieldName.DisplayName), true

	case string(FinancePartiesDTOFieldName.CountryCode):
		return string(model.FinancePartiesDBFieldName.CountryCode), true

	case string(FinancePartiesDTOFieldName.Email):
		return string(model.FinancePartiesDBFieldName.Email), true

	case string(FinancePartiesDTOFieldName.Phone):
		return string(model.FinancePartiesDBFieldName.Phone), true

	case string(FinancePartiesDTOFieldName.PartyStatus):
		return string(model.FinancePartiesDBFieldName.PartyStatus), true

	case string(FinancePartiesDTOFieldName.Metadata):
		return string(model.FinancePartiesDBFieldName.Metadata), true

	case string(FinancePartiesDTOFieldName.MetaCreatedAt):
		return string(model.FinancePartiesDBFieldName.MetaCreatedAt), true

	case string(FinancePartiesDTOFieldName.MetaCreatedBy):
		return string(model.FinancePartiesDBFieldName.MetaCreatedBy), true

	case string(FinancePartiesDTOFieldName.MetaUpdatedAt):
		return string(model.FinancePartiesDBFieldName.MetaUpdatedAt), true

	case string(FinancePartiesDTOFieldName.MetaUpdatedBy):
		return string(model.FinancePartiesDBFieldName.MetaUpdatedBy), true

	case string(FinancePartiesDTOFieldName.MetaDeletedAt):
		return string(model.FinancePartiesDBFieldName.MetaDeletedAt), true

	case string(FinancePartiesDTOFieldName.MetaDeletedBy):
		return string(model.FinancePartiesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinancePartiesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinancePartiesBaseFilterField(field string) bool {
	spec, found := model.NewFinancePartiesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinancePartiesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinancePartiesProjectionOutputPath(path string) error {
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

func transformFinancePartiesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinancePartiesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinancePartiesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinancePartiesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinancePartiesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinancePartiesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinancePartiesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinancePartiesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinancePartiesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinancePartiesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinancePartiesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinancePartiesFilter(filter *model.Filter) {
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
			Field: string(FinancePartiesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinancePartiesSelectableResponse map[string]interface{}
type FinancePartiesSelectableListResponse []*FinancePartiesSelectableResponse

func assignFinancePartiesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinancePartiesSelectableValue(out FinancePartiesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinancePartiesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinancePartiesSelectableResponse(financeParties model.FinanceParties, filter model.Filter) FinancePartiesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinancePartiesDBFieldName.Id),
			string(model.FinancePartiesDBFieldName.PartyType),
			string(model.FinancePartiesDBFieldName.ExternalRef),
			string(model.FinancePartiesDBFieldName.LegalName),
			string(model.FinancePartiesDBFieldName.DisplayName),
			string(model.FinancePartiesDBFieldName.CountryCode),
			string(model.FinancePartiesDBFieldName.Email),
			string(model.FinancePartiesDBFieldName.Phone),
			string(model.FinancePartiesDBFieldName.PartyStatus),
			string(model.FinancePartiesDBFieldName.Metadata),
			string(model.FinancePartiesDBFieldName.MetaCreatedAt),
			string(model.FinancePartiesDBFieldName.MetaCreatedBy),
			string(model.FinancePartiesDBFieldName.MetaUpdatedAt),
			string(model.FinancePartiesDBFieldName.MetaUpdatedBy),
			string(model.FinancePartiesDBFieldName.MetaDeletedAt),
			string(model.FinancePartiesDBFieldName.MetaDeletedBy),
		)
	}
	financePartiesSelectableResponse := FinancePartiesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinancePartiesDBFieldName.Id):
			key := string(FinancePartiesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.Id, explicitAlias)

		case string(model.FinancePartiesDBFieldName.PartyType):
			key := string(FinancePartiesDTOFieldName.PartyType)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, model.PartyType(financeParties.PartyType), explicitAlias)

		case string(model.FinancePartiesDBFieldName.ExternalRef):
			key := string(FinancePartiesDTOFieldName.ExternalRef)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.ExternalRef.String, explicitAlias)

		case string(model.FinancePartiesDBFieldName.LegalName):
			key := string(FinancePartiesDTOFieldName.LegalName)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.LegalName, explicitAlias)

		case string(model.FinancePartiesDBFieldName.DisplayName):
			key := string(FinancePartiesDTOFieldName.DisplayName)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.DisplayName.String, explicitAlias)

		case string(model.FinancePartiesDBFieldName.CountryCode):
			key := string(FinancePartiesDTOFieldName.CountryCode)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.CountryCode, explicitAlias)

		case string(model.FinancePartiesDBFieldName.Email):
			key := string(FinancePartiesDTOFieldName.Email)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.Email.String, explicitAlias)

		case string(model.FinancePartiesDBFieldName.Phone):
			key := string(FinancePartiesDTOFieldName.Phone)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.Phone.String, explicitAlias)

		case string(model.FinancePartiesDBFieldName.PartyStatus):
			key := string(FinancePartiesDTOFieldName.PartyStatus)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, model.PartyStatus(financeParties.PartyStatus), explicitAlias)

		case string(model.FinancePartiesDBFieldName.Metadata):
			key := string(FinancePartiesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.Metadata, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaCreatedAt):
			key := string(FinancePartiesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaCreatedAt, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaCreatedBy):
			key := string(FinancePartiesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaCreatedBy, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaUpdatedAt):
			key := string(FinancePartiesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaUpdatedAt, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaUpdatedBy):
			key := string(FinancePartiesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaUpdatedBy, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaDeletedAt):
			key := string(FinancePartiesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinancePartiesDBFieldName.MetaDeletedBy):
			key := string(FinancePartiesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinancePartiesSelectableValue(financePartiesSelectableResponse, key, financeParties.MetaDeletedBy, explicitAlias)

		}
	}
	return financePartiesSelectableResponse
}

func NewFinancePartiesListResponseFromFilterResult(result []model.FinancePartiesFilterResult, filter model.Filter) FinancePartiesSelectableListResponse {
	dtoFinancePartiesListResponse := FinancePartiesSelectableListResponse{}
	for _, row := range result {
		dtoFinancePartiesResponse := NewFinancePartiesSelectableResponse(row.FinanceParties, filter)
		dtoFinancePartiesListResponse = append(dtoFinancePartiesListResponse, &dtoFinancePartiesResponse)
	}
	return dtoFinancePartiesListResponse
}

type FinancePartiesFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     FinancePartiesSelectableListResponse `json:"data"`
}

func reverseFinancePartiesFilterResults(result []model.FinancePartiesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinancePartiesFilterResponse(result []model.FinancePartiesFilterResult, filter model.Filter) (resp FinancePartiesFilterResponse) {
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
			reverseFinancePartiesFilterResults(dataResult)
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

	resp.Data = NewFinancePartiesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinancePartiesCreateRequest struct {
	PartyType   model.PartyType   `json:"partyType" example:"platform" enums:"platform,merchant,customer,provider,bank,tax_authority,affiliate,guide,driver,manual_ops"`
	ExternalRef string            `json:"externalRef"`
	LegalName   string            `json:"legalName"`
	DisplayName string            `json:"displayName"`
	CountryCode string            `json:"countryCode"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	PartyStatus model.PartyStatus `json:"partyStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata    json.RawMessage   `json:"metadata"`
}

func (d *FinancePartiesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinancePartiesCreateRequest) ToModel() model.FinanceParties {
	id, _ := uuid.NewV4()
	return model.FinanceParties{
		Id:          id,
		PartyType:   d.PartyType,
		ExternalRef: null.StringFrom(d.ExternalRef),
		LegalName:   d.LegalName,
		DisplayName: null.StringFrom(d.DisplayName),
		CountryCode: d.CountryCode,
		Email:       null.StringFrom(d.Email),
		Phone:       null.StringFrom(d.Phone),
		PartyStatus: d.PartyStatus,
		Metadata:    d.Metadata,
	}
}

type FinancePartiesListCreateRequest []*FinancePartiesCreateRequest

func (d FinancePartiesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeParties := range d {
		err = validator.Struct(financeParties)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinancePartiesListCreateRequest) ToModelList() []model.FinanceParties {
	out := make([]model.FinanceParties, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinancePartiesUpdateRequest struct {
	PartyType   model.PartyType   `json:"partyType" example:"platform" enums:"platform,merchant,customer,provider,bank,tax_authority,affiliate,guide,driver,manual_ops"`
	ExternalRef string            `json:"externalRef"`
	LegalName   string            `json:"legalName"`
	DisplayName string            `json:"displayName"`
	CountryCode string            `json:"countryCode"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	PartyStatus model.PartyStatus `json:"partyStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata    json.RawMessage   `json:"metadata"`
}

func (d *FinancePartiesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinancePartiesUpdateRequest) ToModel() model.FinanceParties {
	return model.FinanceParties{
		PartyType:   d.PartyType,
		ExternalRef: null.StringFrom(d.ExternalRef),
		LegalName:   d.LegalName,
		DisplayName: null.StringFrom(d.DisplayName),
		CountryCode: d.CountryCode,
		Email:       null.StringFrom(d.Email),
		Phone:       null.StringFrom(d.Phone),
		PartyStatus: d.PartyStatus,
		Metadata:    d.Metadata,
	}
}

type FinancePartiesBulkUpdateRequest struct {
	Id          uuid.UUID         `json:"id"`
	PartyType   model.PartyType   `json:"partyType" example:"platform" enums:"platform,merchant,customer,provider,bank,tax_authority,affiliate,guide,driver,manual_ops"`
	ExternalRef string            `json:"externalRef"`
	LegalName   string            `json:"legalName"`
	DisplayName string            `json:"displayName"`
	CountryCode string            `json:"countryCode"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	PartyStatus model.PartyStatus `json:"partyStatus" example:"active" enums:"active,inactive,blocked"`
	Metadata    json.RawMessage   `json:"metadata"`
}

func (d FinancePartiesBulkUpdateRequest) PrimaryID() FinancePartiesPrimaryID {
	return FinancePartiesPrimaryID{
		Id: d.Id,
	}
}

type FinancePartiesListBulkUpdateRequest []*FinancePartiesBulkUpdateRequest

func (d FinancePartiesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeParties := range d {
		err = validator.Struct(financeParties)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinancePartiesBulkUpdateRequest) ToModel() model.FinanceParties {
	return model.FinanceParties{
		Id:          d.Id,
		PartyType:   d.PartyType,
		ExternalRef: null.StringFrom(d.ExternalRef),
		LegalName:   d.LegalName,
		DisplayName: null.StringFrom(d.DisplayName),
		CountryCode: d.CountryCode,
		Email:       null.StringFrom(d.Email),
		Phone:       null.StringFrom(d.Phone),
		PartyStatus: d.PartyStatus,
		Metadata:    d.Metadata,
	}
}

type FinancePartiesResponse struct {
	Id          uuid.UUID         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PartyType   model.PartyType   `json:"partyType" validate:"required,oneof=platform merchant customer provider bank tax_authority affiliate guide driver manual_ops" enums:"platform,merchant,customer,provider,bank,tax_authority,affiliate,guide,driver,manual_ops"`
	ExternalRef string            `json:"externalRef"`
	LegalName   string            `json:"legalName" validate:"required"`
	DisplayName string            `json:"displayName"`
	CountryCode string            `json:"countryCode"`
	Email       string            `json:"email" validate:"email"`
	Phone       string            `json:"phone"`
	PartyStatus model.PartyStatus `json:"partyStatus" validate:"oneof=active inactive blocked" enums:"active,inactive,blocked"`
	Metadata    json.RawMessage   `json:"metadata" swaggertype:"object"`
}

func NewFinancePartiesResponse(financeParties model.FinanceParties) FinancePartiesResponse {
	return FinancePartiesResponse{
		Id:          financeParties.Id,
		PartyType:   model.PartyType(financeParties.PartyType),
		ExternalRef: financeParties.ExternalRef.String,
		LegalName:   financeParties.LegalName,
		DisplayName: financeParties.DisplayName.String,
		CountryCode: financeParties.CountryCode,
		Email:       financeParties.Email.String,
		Phone:       financeParties.Phone.String,
		PartyStatus: model.PartyStatus(financeParties.PartyStatus),
		Metadata:    financeParties.Metadata,
	}
}

type FinancePartiesListResponse []*FinancePartiesResponse

func NewFinancePartiesListResponse(financePartiesList model.FinancePartiesList) FinancePartiesListResponse {
	dtoFinancePartiesListResponse := FinancePartiesListResponse{}
	for _, financeParties := range financePartiesList {
		dtoFinancePartiesResponse := NewFinancePartiesResponse(*financeParties)
		dtoFinancePartiesListResponse = append(dtoFinancePartiesListResponse, &dtoFinancePartiesResponse)
	}
	return dtoFinancePartiesListResponse
}

type FinancePartiesPrimaryIDList []FinancePartiesPrimaryID

func (d FinancePartiesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeParties := range d {
		err = validator.Struct(financeParties)
		if err != nil {
			return
		}
	}
	return nil
}

type FinancePartiesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinancePartiesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinancePartiesPrimaryID) ToModel() model.FinancePartiesPrimaryID {
	return model.FinancePartiesPrimaryID{
		Id: d.Id,
	}
}

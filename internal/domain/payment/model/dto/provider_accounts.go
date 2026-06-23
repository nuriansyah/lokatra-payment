package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderAccountsDTOFieldNameType string

type providerAccountsDTOFieldName struct {
	Id                  ProviderAccountsDTOFieldNameType
	ProviderId          ProviderAccountsDTOFieldNameType
	AccountName         ProviderAccountsDTOFieldNameType
	Environment         ProviderAccountsDTOFieldNameType
	OwnerType           ProviderAccountsDTOFieldNameType
	OwnerId             ProviderAccountsDTOFieldNameType
	MerchantRef         ProviderAccountsDTOFieldNameType
	CredentialSecretRef ProviderAccountsDTOFieldNameType
	WebhookSecretRef    ProviderAccountsDTOFieldNameType
	PublicKeyRef        ProviderAccountsDTOFieldNameType
	Status              ProviderAccountsDTOFieldNameType
	Config              ProviderAccountsDTOFieldNameType
	Metadata            ProviderAccountsDTOFieldNameType
	MetaCreatedAt       ProviderAccountsDTOFieldNameType
	MetaCreatedBy       ProviderAccountsDTOFieldNameType
	MetaUpdatedAt       ProviderAccountsDTOFieldNameType
	MetaUpdatedBy       ProviderAccountsDTOFieldNameType
	MetaDeletedAt       ProviderAccountsDTOFieldNameType
	MetaDeletedBy       ProviderAccountsDTOFieldNameType
}

var ProviderAccountsDTOFieldName = providerAccountsDTOFieldName{
	Id:                  "id",
	ProviderId:          "providerId",
	AccountName:         "accountName",
	Environment:         "environment",
	OwnerType:           "ownerType",
	OwnerId:             "ownerId",
	MerchantRef:         "merchantRef",
	CredentialSecretRef: "credentialSecretRef",
	WebhookSecretRef:    "webhookSecretRef",
	PublicKeyRef:        "publicKeyRef",
	Status:              "status",
	Config:              "config",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformProviderAccountsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderAccountsDTOFieldName.Id):
		return string(model.ProviderAccountsDBFieldName.Id), true

	case string(ProviderAccountsDTOFieldName.ProviderId):
		return string(model.ProviderAccountsDBFieldName.ProviderId), true

	case string(ProviderAccountsDTOFieldName.AccountName):
		return string(model.ProviderAccountsDBFieldName.AccountName), true

	case string(ProviderAccountsDTOFieldName.Environment):
		return string(model.ProviderAccountsDBFieldName.Environment), true

	case string(ProviderAccountsDTOFieldName.OwnerType):
		return string(model.ProviderAccountsDBFieldName.OwnerType), true

	case string(ProviderAccountsDTOFieldName.OwnerId):
		return string(model.ProviderAccountsDBFieldName.OwnerId), true

	case string(ProviderAccountsDTOFieldName.MerchantRef):
		return string(model.ProviderAccountsDBFieldName.MerchantRef), true

	case string(ProviderAccountsDTOFieldName.CredentialSecretRef):
		return string(model.ProviderAccountsDBFieldName.CredentialSecretRef), true

	case string(ProviderAccountsDTOFieldName.WebhookSecretRef):
		return string(model.ProviderAccountsDBFieldName.WebhookSecretRef), true

	case string(ProviderAccountsDTOFieldName.PublicKeyRef):
		return string(model.ProviderAccountsDBFieldName.PublicKeyRef), true

	case string(ProviderAccountsDTOFieldName.Status):
		return string(model.ProviderAccountsDBFieldName.Status), true

	case string(ProviderAccountsDTOFieldName.Config):
		return string(model.ProviderAccountsDBFieldName.Config), true

	case string(ProviderAccountsDTOFieldName.Metadata):
		return string(model.ProviderAccountsDBFieldName.Metadata), true

	case string(ProviderAccountsDTOFieldName.MetaCreatedAt):
		return string(model.ProviderAccountsDBFieldName.MetaCreatedAt), true

	case string(ProviderAccountsDTOFieldName.MetaCreatedBy):
		return string(model.ProviderAccountsDBFieldName.MetaCreatedBy), true

	case string(ProviderAccountsDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderAccountsDBFieldName.MetaUpdatedAt), true

	case string(ProviderAccountsDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderAccountsDBFieldName.MetaUpdatedBy), true

	case string(ProviderAccountsDTOFieldName.MetaDeletedAt):
		return string(model.ProviderAccountsDBFieldName.MetaDeletedAt), true

	case string(ProviderAccountsDTOFieldName.MetaDeletedBy):
		return string(model.ProviderAccountsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderAccountsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderAccountsBaseFilterField(field string) bool {
	spec, found := model.NewProviderAccountsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderAccountsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderAccountsProjectionOutputPath(path string) error {
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

func transformProviderAccountsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderAccountsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderAccountsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderAccountsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderAccountsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderAccountsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderAccountsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderAccountsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderAccountsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderAccountsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderAccountsFilter(filter *model.Filter) {
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
			Field: string(ProviderAccountsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderAccountsSelectableResponse map[string]interface{}
type ProviderAccountsSelectableListResponse []*ProviderAccountsSelectableResponse

func assignProviderAccountsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderAccountsSelectableValue(out ProviderAccountsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderAccountsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderAccountsSelectableResponse(providerAccounts model.ProviderAccounts, filter model.Filter) ProviderAccountsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderAccountsDBFieldName.Id),
			string(model.ProviderAccountsDBFieldName.ProviderId),
			string(model.ProviderAccountsDBFieldName.AccountName),
			string(model.ProviderAccountsDBFieldName.Environment),
			string(model.ProviderAccountsDBFieldName.OwnerType),
			string(model.ProviderAccountsDBFieldName.OwnerId),
			string(model.ProviderAccountsDBFieldName.MerchantRef),
			string(model.ProviderAccountsDBFieldName.CredentialSecretRef),
			string(model.ProviderAccountsDBFieldName.WebhookSecretRef),
			string(model.ProviderAccountsDBFieldName.PublicKeyRef),
			string(model.ProviderAccountsDBFieldName.Status),
			string(model.ProviderAccountsDBFieldName.Config),
			string(model.ProviderAccountsDBFieldName.Metadata),
			string(model.ProviderAccountsDBFieldName.MetaCreatedAt),
			string(model.ProviderAccountsDBFieldName.MetaCreatedBy),
			string(model.ProviderAccountsDBFieldName.MetaUpdatedAt),
			string(model.ProviderAccountsDBFieldName.MetaUpdatedBy),
			string(model.ProviderAccountsDBFieldName.MetaDeletedAt),
			string(model.ProviderAccountsDBFieldName.MetaDeletedBy),
		)
	}
	providerAccountsSelectableResponse := ProviderAccountsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderAccountsDBFieldName.Id):
			key := string(ProviderAccountsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.Id, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.ProviderId):
			key := string(ProviderAccountsDTOFieldName.ProviderId)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.ProviderId, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.AccountName):
			key := string(ProviderAccountsDTOFieldName.AccountName)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.AccountName, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.Environment):
			key := string(ProviderAccountsDTOFieldName.Environment)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.Environment, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.OwnerType):
			key := string(ProviderAccountsDTOFieldName.OwnerType)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.OwnerType, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.OwnerId):
			key := string(ProviderAccountsDTOFieldName.OwnerId)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.OwnerId.UUID, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MerchantRef):
			key := string(ProviderAccountsDTOFieldName.MerchantRef)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MerchantRef.String, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.CredentialSecretRef):
			key := string(ProviderAccountsDTOFieldName.CredentialSecretRef)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.CredentialSecretRef, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.WebhookSecretRef):
			key := string(ProviderAccountsDTOFieldName.WebhookSecretRef)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.WebhookSecretRef.String, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.PublicKeyRef):
			key := string(ProviderAccountsDTOFieldName.PublicKeyRef)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.PublicKeyRef.String, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.Status):
			key := string(ProviderAccountsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, model.ProviderAccountStatus(providerAccounts.Status), explicitAlias)

		case string(model.ProviderAccountsDBFieldName.Config):
			key := string(ProviderAccountsDTOFieldName.Config)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.Config, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.Metadata):
			key := string(ProviderAccountsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.Metadata, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaCreatedAt):
			key := string(ProviderAccountsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaCreatedAt, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaCreatedBy):
			key := string(ProviderAccountsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaCreatedBy, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaUpdatedAt):
			key := string(ProviderAccountsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaUpdatedBy):
			key := string(ProviderAccountsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaUpdatedBy, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaDeletedAt):
			key := string(ProviderAccountsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderAccountsDBFieldName.MetaDeletedBy):
			key := string(ProviderAccountsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderAccountsSelectableValue(providerAccountsSelectableResponse, key, providerAccounts.MetaDeletedBy, explicitAlias)

		}
	}
	return providerAccountsSelectableResponse
}

func NewProviderAccountsListResponseFromFilterResult(result []model.ProviderAccountsFilterResult, filter model.Filter) ProviderAccountsSelectableListResponse {
	dtoProviderAccountsListResponse := ProviderAccountsSelectableListResponse{}
	for _, row := range result {
		dtoProviderAccountsResponse := NewProviderAccountsSelectableResponse(row.ProviderAccounts, filter)
		dtoProviderAccountsListResponse = append(dtoProviderAccountsListResponse, &dtoProviderAccountsResponse)
	}
	return dtoProviderAccountsListResponse
}

type ProviderAccountsFilterResponse struct {
	Metadata Metadata                               `json:"metadata"`
	Data     ProviderAccountsSelectableListResponse `json:"data"`
}

func reverseProviderAccountsFilterResults(result []model.ProviderAccountsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderAccountsFilterResponse(result []model.ProviderAccountsFilterResult, filter model.Filter) (resp ProviderAccountsFilterResponse) {
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
			reverseProviderAccountsFilterResults(dataResult)
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

	resp.Data = NewProviderAccountsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderAccountsCreateRequest struct {
	ProviderId          uuid.UUID                   `json:"providerId"`
	AccountName         string                      `json:"accountName"`
	Environment         string                      `json:"environment"`
	OwnerType           string                      `json:"ownerType"`
	OwnerId             uuid.UUID                   `json:"ownerId"`
	MerchantRef         string                      `json:"merchantRef"`
	CredentialSecretRef string                      `json:"credentialSecretRef"`
	WebhookSecretRef    string                      `json:"webhookSecretRef"`
	PublicKeyRef        string                      `json:"publicKeyRef"`
	Status              model.ProviderAccountStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Config              json.RawMessage             `json:"config"`
	Metadata            json.RawMessage             `json:"metadata"`
}

func (d *ProviderAccountsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderAccountsCreateRequest) ToModel() model.ProviderAccounts {
	id, _ := uuid.NewV4()
	return model.ProviderAccounts{
		Id:                  id,
		ProviderId:          d.ProviderId,
		AccountName:         d.AccountName,
		Environment:         d.Environment,
		OwnerType:           d.OwnerType,
		OwnerId:             nuuid.From(d.OwnerId),
		MerchantRef:         null.StringFrom(d.MerchantRef),
		CredentialSecretRef: d.CredentialSecretRef,
		WebhookSecretRef:    null.StringFrom(d.WebhookSecretRef),
		PublicKeyRef:        null.StringFrom(d.PublicKeyRef),
		Status:              d.Status,
		Config:              d.Config,
		Metadata:            d.Metadata,
	}
}

type ProviderAccountsListCreateRequest []*ProviderAccountsCreateRequest

func (d ProviderAccountsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerAccounts := range d {
		err = validator.Struct(providerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderAccountsListCreateRequest) ToModelList() []model.ProviderAccounts {
	out := make([]model.ProviderAccounts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderAccountsUpdateRequest struct {
	ProviderId          uuid.UUID                   `json:"providerId"`
	AccountName         string                      `json:"accountName"`
	Environment         string                      `json:"environment"`
	OwnerType           string                      `json:"ownerType"`
	OwnerId             uuid.UUID                   `json:"ownerId"`
	MerchantRef         string                      `json:"merchantRef"`
	CredentialSecretRef string                      `json:"credentialSecretRef"`
	WebhookSecretRef    string                      `json:"webhookSecretRef"`
	PublicKeyRef        string                      `json:"publicKeyRef"`
	Status              model.ProviderAccountStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Config              json.RawMessage             `json:"config"`
	Metadata            json.RawMessage             `json:"metadata"`
}

func (d *ProviderAccountsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderAccountsUpdateRequest) ToModel() model.ProviderAccounts {
	return model.ProviderAccounts{
		ProviderId:          d.ProviderId,
		AccountName:         d.AccountName,
		Environment:         d.Environment,
		OwnerType:           d.OwnerType,
		OwnerId:             nuuid.From(d.OwnerId),
		MerchantRef:         null.StringFrom(d.MerchantRef),
		CredentialSecretRef: d.CredentialSecretRef,
		WebhookSecretRef:    null.StringFrom(d.WebhookSecretRef),
		PublicKeyRef:        null.StringFrom(d.PublicKeyRef),
		Status:              d.Status,
		Config:              d.Config,
		Metadata:            d.Metadata,
	}
}

type ProviderAccountsBulkUpdateRequest struct {
	Id                  uuid.UUID                   `json:"id"`
	ProviderId          uuid.UUID                   `json:"providerId"`
	AccountName         string                      `json:"accountName"`
	Environment         string                      `json:"environment"`
	OwnerType           string                      `json:"ownerType"`
	OwnerId             uuid.UUID                   `json:"ownerId"`
	MerchantRef         string                      `json:"merchantRef"`
	CredentialSecretRef string                      `json:"credentialSecretRef"`
	WebhookSecretRef    string                      `json:"webhookSecretRef"`
	PublicKeyRef        string                      `json:"publicKeyRef"`
	Status              model.ProviderAccountStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Config              json.RawMessage             `json:"config"`
	Metadata            json.RawMessage             `json:"metadata"`
}

func (d ProviderAccountsBulkUpdateRequest) PrimaryID() ProviderAccountsPrimaryID {
	return ProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

type ProviderAccountsListBulkUpdateRequest []*ProviderAccountsBulkUpdateRequest

func (d ProviderAccountsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerAccounts := range d {
		err = validator.Struct(providerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderAccountsBulkUpdateRequest) ToModel() model.ProviderAccounts {
	return model.ProviderAccounts{
		Id:                  d.Id,
		ProviderId:          d.ProviderId,
		AccountName:         d.AccountName,
		Environment:         d.Environment,
		OwnerType:           d.OwnerType,
		OwnerId:             nuuid.From(d.OwnerId),
		MerchantRef:         null.StringFrom(d.MerchantRef),
		CredentialSecretRef: d.CredentialSecretRef,
		WebhookSecretRef:    null.StringFrom(d.WebhookSecretRef),
		PublicKeyRef:        null.StringFrom(d.PublicKeyRef),
		Status:              d.Status,
		Config:              d.Config,
		Metadata:            d.Metadata,
	}
}

type ProviderAccountsResponse struct {
	Id                  uuid.UUID                   `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderId          uuid.UUID                   `json:"providerId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AccountName         string                      `json:"accountName" validate:"required"`
	Environment         string                      `json:"environment"`
	OwnerType           string                      `json:"ownerType"`
	OwnerId             uuid.UUID                   `json:"ownerId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantRef         string                      `json:"merchantRef"`
	CredentialSecretRef string                      `json:"credentialSecretRef" validate:"required"`
	WebhookSecretRef    string                      `json:"webhookSecretRef"`
	PublicKeyRef        string                      `json:"publicKeyRef"`
	Status              model.ProviderAccountStatus `json:"status" validate:"oneof=active inactive deprecated" enums:"active,inactive,deprecated"`
	Config              json.RawMessage             `json:"config" swaggertype:"object"`
	Metadata            json.RawMessage             `json:"metadata" swaggertype:"object"`
}

func NewProviderAccountsResponse(providerAccounts model.ProviderAccounts) ProviderAccountsResponse {
	return ProviderAccountsResponse{
		Id:                  providerAccounts.Id,
		ProviderId:          providerAccounts.ProviderId,
		AccountName:         providerAccounts.AccountName,
		Environment:         providerAccounts.Environment,
		OwnerType:           providerAccounts.OwnerType,
		OwnerId:             providerAccounts.OwnerId.UUID,
		MerchantRef:         providerAccounts.MerchantRef.String,
		CredentialSecretRef: providerAccounts.CredentialSecretRef,
		WebhookSecretRef:    providerAccounts.WebhookSecretRef.String,
		PublicKeyRef:        providerAccounts.PublicKeyRef.String,
		Status:              model.ProviderAccountStatus(providerAccounts.Status),
		Config:              providerAccounts.Config,
		Metadata:            providerAccounts.Metadata,
	}
}

type ProviderAccountsListResponse []*ProviderAccountsResponse

func NewProviderAccountsListResponse(providerAccountsList model.ProviderAccountsList) ProviderAccountsListResponse {
	dtoProviderAccountsListResponse := ProviderAccountsListResponse{}
	for _, providerAccounts := range providerAccountsList {
		dtoProviderAccountsResponse := NewProviderAccountsResponse(*providerAccounts)
		dtoProviderAccountsListResponse = append(dtoProviderAccountsListResponse, &dtoProviderAccountsResponse)
	}
	return dtoProviderAccountsListResponse
}

type ProviderAccountsPrimaryIDList []ProviderAccountsPrimaryID

func (d ProviderAccountsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerAccounts := range d {
		err = validator.Struct(providerAccounts)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderAccountsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderAccountsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderAccountsPrimaryID) ToModel() model.ProviderAccountsPrimaryID {
	return model.ProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

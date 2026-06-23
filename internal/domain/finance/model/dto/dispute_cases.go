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

type DisputeCasesDTOFieldNameType string

type disputeCasesDTOFieldName struct {
	Id                 DisputeCasesDTOFieldNameType
	DisputeCode        DisputeCasesDTOFieldNameType
	BookId             DisputeCasesDTOFieldNameType
	SourceType         DisputeCasesDTOFieldNameType
	SourceId           DisputeCasesDTOFieldNameType
	MerchantPartyId    DisputeCasesDTOFieldNameType
	CustomerPartyId    DisputeCasesDTOFieldNameType
	ProviderCode       DisputeCasesDTOFieldNameType
	ProviderDisputeRef DisputeCasesDTOFieldNameType
	IdempotencyKey     DisputeCasesDTOFieldNameType
	CurrencyCode       DisputeCasesDTOFieldNameType
	DisputedAmount     DisputeCasesDTOFieldNameType
	DisputeReasonCode  DisputeCasesDTOFieldNameType
	DisputeStatus      DisputeCasesDTOFieldNameType
	OpenedAt           DisputeCasesDTOFieldNameType
	DueAt              DisputeCasesDTOFieldNameType
	ClosedAt           DisputeCasesDTOFieldNameType
	Metadata           DisputeCasesDTOFieldNameType
	MetaCreatedAt      DisputeCasesDTOFieldNameType
	MetaCreatedBy      DisputeCasesDTOFieldNameType
	MetaUpdatedAt      DisputeCasesDTOFieldNameType
	MetaUpdatedBy      DisputeCasesDTOFieldNameType
	MetaDeletedAt      DisputeCasesDTOFieldNameType
	MetaDeletedBy      DisputeCasesDTOFieldNameType
}

var DisputeCasesDTOFieldName = disputeCasesDTOFieldName{
	Id:                 "id",
	DisputeCode:        "disputeCode",
	BookId:             "bookId",
	SourceType:         "sourceType",
	SourceId:           "sourceId",
	MerchantPartyId:    "merchantPartyId",
	CustomerPartyId:    "customerPartyId",
	ProviderCode:       "providerCode",
	ProviderDisputeRef: "providerDisputeRef",
	IdempotencyKey:     "idempotencyKey",
	CurrencyCode:       "currencyCode",
	DisputedAmount:     "disputedAmount",
	DisputeReasonCode:  "disputeReasonCode",
	DisputeStatus:      "disputeStatus",
	OpenedAt:           "openedAt",
	DueAt:              "dueAt",
	ClosedAt:           "closedAt",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformDisputeCasesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(DisputeCasesDTOFieldName.Id):
		return string(model.DisputeCasesDBFieldName.Id), true

	case string(DisputeCasesDTOFieldName.DisputeCode):
		return string(model.DisputeCasesDBFieldName.DisputeCode), true

	case string(DisputeCasesDTOFieldName.BookId):
		return string(model.DisputeCasesDBFieldName.BookId), true

	case string(DisputeCasesDTOFieldName.SourceType):
		return string(model.DisputeCasesDBFieldName.SourceType), true

	case string(DisputeCasesDTOFieldName.SourceId):
		return string(model.DisputeCasesDBFieldName.SourceId), true

	case string(DisputeCasesDTOFieldName.MerchantPartyId):
		return string(model.DisputeCasesDBFieldName.MerchantPartyId), true

	case string(DisputeCasesDTOFieldName.CustomerPartyId):
		return string(model.DisputeCasesDBFieldName.CustomerPartyId), true

	case string(DisputeCasesDTOFieldName.ProviderCode):
		return string(model.DisputeCasesDBFieldName.ProviderCode), true

	case string(DisputeCasesDTOFieldName.ProviderDisputeRef):
		return string(model.DisputeCasesDBFieldName.ProviderDisputeRef), true

	case string(DisputeCasesDTOFieldName.IdempotencyKey):
		return string(model.DisputeCasesDBFieldName.IdempotencyKey), true

	case string(DisputeCasesDTOFieldName.CurrencyCode):
		return string(model.DisputeCasesDBFieldName.CurrencyCode), true

	case string(DisputeCasesDTOFieldName.DisputedAmount):
		return string(model.DisputeCasesDBFieldName.DisputedAmount), true

	case string(DisputeCasesDTOFieldName.DisputeReasonCode):
		return string(model.DisputeCasesDBFieldName.DisputeReasonCode), true

	case string(DisputeCasesDTOFieldName.DisputeStatus):
		return string(model.DisputeCasesDBFieldName.DisputeStatus), true

	case string(DisputeCasesDTOFieldName.OpenedAt):
		return string(model.DisputeCasesDBFieldName.OpenedAt), true

	case string(DisputeCasesDTOFieldName.DueAt):
		return string(model.DisputeCasesDBFieldName.DueAt), true

	case string(DisputeCasesDTOFieldName.ClosedAt):
		return string(model.DisputeCasesDBFieldName.ClosedAt), true

	case string(DisputeCasesDTOFieldName.Metadata):
		return string(model.DisputeCasesDBFieldName.Metadata), true

	case string(DisputeCasesDTOFieldName.MetaCreatedAt):
		return string(model.DisputeCasesDBFieldName.MetaCreatedAt), true

	case string(DisputeCasesDTOFieldName.MetaCreatedBy):
		return string(model.DisputeCasesDBFieldName.MetaCreatedBy), true

	case string(DisputeCasesDTOFieldName.MetaUpdatedAt):
		return string(model.DisputeCasesDBFieldName.MetaUpdatedAt), true

	case string(DisputeCasesDTOFieldName.MetaUpdatedBy):
		return string(model.DisputeCasesDBFieldName.MetaUpdatedBy), true

	case string(DisputeCasesDTOFieldName.MetaDeletedAt):
		return string(model.DisputeCasesDBFieldName.MetaDeletedAt), true

	case string(DisputeCasesDTOFieldName.MetaDeletedBy):
		return string(model.DisputeCasesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewDisputeCasesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isDisputeCasesBaseFilterField(field string) bool {
	spec, found := model.NewDisputeCasesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeDisputeCasesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateDisputeCasesProjectionOutputPath(path string) error {
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

func transformDisputeCasesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformDisputeCasesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformDisputeCasesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformDisputeCasesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformDisputeCasesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isDisputeCasesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateDisputeCasesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeDisputeCasesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformDisputeCasesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformDisputeCasesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformDisputeCasesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultDisputeCasesFilter(filter *model.Filter) {
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
			Field: string(DisputeCasesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type DisputeCasesSelectableResponse map[string]interface{}
type DisputeCasesSelectableListResponse []*DisputeCasesSelectableResponse

func assignDisputeCasesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setDisputeCasesSelectableValue(out DisputeCasesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignDisputeCasesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewDisputeCasesSelectableResponse(disputeCases model.DisputeCases, filter model.Filter) DisputeCasesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.DisputeCasesDBFieldName.Id),
			string(model.DisputeCasesDBFieldName.DisputeCode),
			string(model.DisputeCasesDBFieldName.BookId),
			string(model.DisputeCasesDBFieldName.SourceType),
			string(model.DisputeCasesDBFieldName.SourceId),
			string(model.DisputeCasesDBFieldName.MerchantPartyId),
			string(model.DisputeCasesDBFieldName.CustomerPartyId),
			string(model.DisputeCasesDBFieldName.ProviderCode),
			string(model.DisputeCasesDBFieldName.ProviderDisputeRef),
			string(model.DisputeCasesDBFieldName.IdempotencyKey),
			string(model.DisputeCasesDBFieldName.CurrencyCode),
			string(model.DisputeCasesDBFieldName.DisputedAmount),
			string(model.DisputeCasesDBFieldName.DisputeReasonCode),
			string(model.DisputeCasesDBFieldName.DisputeStatus),
			string(model.DisputeCasesDBFieldName.OpenedAt),
			string(model.DisputeCasesDBFieldName.DueAt),
			string(model.DisputeCasesDBFieldName.ClosedAt),
			string(model.DisputeCasesDBFieldName.Metadata),
			string(model.DisputeCasesDBFieldName.MetaCreatedAt),
			string(model.DisputeCasesDBFieldName.MetaCreatedBy),
			string(model.DisputeCasesDBFieldName.MetaUpdatedAt),
			string(model.DisputeCasesDBFieldName.MetaUpdatedBy),
			string(model.DisputeCasesDBFieldName.MetaDeletedAt),
			string(model.DisputeCasesDBFieldName.MetaDeletedBy),
		)
	}
	disputeCasesSelectableResponse := DisputeCasesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.DisputeCasesDBFieldName.Id):
			key := string(DisputeCasesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.Id, explicitAlias)

		case string(model.DisputeCasesDBFieldName.DisputeCode):
			key := string(DisputeCasesDTOFieldName.DisputeCode)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.DisputeCode, explicitAlias)

		case string(model.DisputeCasesDBFieldName.BookId):
			key := string(DisputeCasesDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.BookId, explicitAlias)

		case string(model.DisputeCasesDBFieldName.SourceType):
			key := string(DisputeCasesDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.SourceType, explicitAlias)

		case string(model.DisputeCasesDBFieldName.SourceId):
			key := string(DisputeCasesDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.SourceId, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MerchantPartyId):
			key := string(DisputeCasesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MerchantPartyId.UUID, explicitAlias)

		case string(model.DisputeCasesDBFieldName.CustomerPartyId):
			key := string(DisputeCasesDTOFieldName.CustomerPartyId)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.CustomerPartyId.UUID, explicitAlias)

		case string(model.DisputeCasesDBFieldName.ProviderCode):
			key := string(DisputeCasesDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.ProviderCode.String, explicitAlias)

		case string(model.DisputeCasesDBFieldName.ProviderDisputeRef):
			key := string(DisputeCasesDTOFieldName.ProviderDisputeRef)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.ProviderDisputeRef.String, explicitAlias)

		case string(model.DisputeCasesDBFieldName.IdempotencyKey):
			key := string(DisputeCasesDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.IdempotencyKey.String, explicitAlias)

		case string(model.DisputeCasesDBFieldName.CurrencyCode):
			key := string(DisputeCasesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.CurrencyCode, explicitAlias)

		case string(model.DisputeCasesDBFieldName.DisputedAmount):
			key := string(DisputeCasesDTOFieldName.DisputedAmount)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.DisputedAmount, explicitAlias)

		case string(model.DisputeCasesDBFieldName.DisputeReasonCode):
			key := string(DisputeCasesDTOFieldName.DisputeReasonCode)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.DisputeReasonCode, explicitAlias)

		case string(model.DisputeCasesDBFieldName.DisputeStatus):
			key := string(DisputeCasesDTOFieldName.DisputeStatus)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, model.DisputeStatus(disputeCases.DisputeStatus), explicitAlias)

		case string(model.DisputeCasesDBFieldName.OpenedAt):
			key := string(DisputeCasesDTOFieldName.OpenedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.OpenedAt, explicitAlias)

		case string(model.DisputeCasesDBFieldName.DueAt):
			key := string(DisputeCasesDTOFieldName.DueAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.DueAt.Time, explicitAlias)

		case string(model.DisputeCasesDBFieldName.ClosedAt):
			key := string(DisputeCasesDTOFieldName.ClosedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.ClosedAt.Time, explicitAlias)

		case string(model.DisputeCasesDBFieldName.Metadata):
			key := string(DisputeCasesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.Metadata, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaCreatedAt):
			key := string(DisputeCasesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaCreatedAt, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaCreatedBy):
			key := string(DisputeCasesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaCreatedBy, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaUpdatedAt):
			key := string(DisputeCasesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaUpdatedAt, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaUpdatedBy):
			key := string(DisputeCasesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaUpdatedBy, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaDeletedAt):
			key := string(DisputeCasesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaDeletedAt.Time, explicitAlias)

		case string(model.DisputeCasesDBFieldName.MetaDeletedBy):
			key := string(DisputeCasesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeCasesSelectableValue(disputeCasesSelectableResponse, key, disputeCases.MetaDeletedBy, explicitAlias)

		}
	}
	return disputeCasesSelectableResponse
}

func NewDisputeCasesListResponseFromFilterResult(result []model.DisputeCasesFilterResult, filter model.Filter) DisputeCasesSelectableListResponse {
	dtoDisputeCasesListResponse := DisputeCasesSelectableListResponse{}
	for _, row := range result {
		dtoDisputeCasesResponse := NewDisputeCasesSelectableResponse(row.DisputeCases, filter)
		dtoDisputeCasesListResponse = append(dtoDisputeCasesListResponse, &dtoDisputeCasesResponse)
	}
	return dtoDisputeCasesListResponse
}

type DisputeCasesFilterResponse struct {
	Metadata Metadata                           `json:"metadata"`
	Data     DisputeCasesSelectableListResponse `json:"data"`
}

func reverseDisputeCasesFilterResults(result []model.DisputeCasesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewDisputeCasesFilterResponse(result []model.DisputeCasesFilterResult, filter model.Filter) (resp DisputeCasesFilterResponse) {
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
			reverseDisputeCasesFilterResults(dataResult)
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

	resp.Data = NewDisputeCasesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type DisputeCasesCreateRequest struct {
	DisputeCode        string              `json:"disputeCode"`
	BookId             uuid.UUID           `json:"bookId"`
	SourceType         string              `json:"sourceType"`
	SourceId           uuid.UUID           `json:"sourceId"`
	MerchantPartyId    uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId    uuid.UUID           `json:"customerPartyId"`
	ProviderCode       string              `json:"providerCode"`
	ProviderDisputeRef string              `json:"providerDisputeRef"`
	IdempotencyKey     string              `json:"idempotencyKey"`
	CurrencyCode       string              `json:"currencyCode"`
	DisputedAmount     decimal.Decimal     `json:"disputedAmount"`
	DisputeReasonCode  string              `json:"disputeReasonCode"`
	DisputeStatus      model.DisputeStatus `json:"disputeStatus" example:"opened" enums:"opened,evidence_required,evidence_submitted,under_review,won,lost,accepted,closed,canceled"`
	OpenedAt           time.Time           `json:"openedAt"`
	DueAt              time.Time           `json:"dueAt"`
	ClosedAt           time.Time           `json:"closedAt"`
	Metadata           json.RawMessage     `json:"metadata"`
}

func (d *DisputeCasesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *DisputeCasesCreateRequest) ToModel() model.DisputeCases {
	id, _ := uuid.NewV4()
	return model.DisputeCases{
		Id:                 id,
		DisputeCode:        d.DisputeCode,
		BookId:             d.BookId,
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		MerchantPartyId:    nuuid.From(d.MerchantPartyId),
		CustomerPartyId:    nuuid.From(d.CustomerPartyId),
		ProviderCode:       null.StringFrom(d.ProviderCode),
		ProviderDisputeRef: null.StringFrom(d.ProviderDisputeRef),
		IdempotencyKey:     null.StringFrom(d.IdempotencyKey),
		CurrencyCode:       d.CurrencyCode,
		DisputedAmount:     d.DisputedAmount,
		DisputeReasonCode:  d.DisputeReasonCode,
		DisputeStatus:      d.DisputeStatus,
		OpenedAt:           d.OpenedAt,
		DueAt:              null.TimeFrom(d.DueAt),
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Metadata:           d.Metadata,
	}
}

type DisputeCasesListCreateRequest []*DisputeCasesCreateRequest

func (d DisputeCasesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeCases := range d {
		err = validator.Struct(disputeCases)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeCasesListCreateRequest) ToModelList() []model.DisputeCases {
	out := make([]model.DisputeCases, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type DisputeCasesUpdateRequest struct {
	DisputeCode        string              `json:"disputeCode"`
	BookId             uuid.UUID           `json:"bookId"`
	SourceType         string              `json:"sourceType"`
	SourceId           uuid.UUID           `json:"sourceId"`
	MerchantPartyId    uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId    uuid.UUID           `json:"customerPartyId"`
	ProviderCode       string              `json:"providerCode"`
	ProviderDisputeRef string              `json:"providerDisputeRef"`
	IdempotencyKey     string              `json:"idempotencyKey"`
	CurrencyCode       string              `json:"currencyCode"`
	DisputedAmount     decimal.Decimal     `json:"disputedAmount"`
	DisputeReasonCode  string              `json:"disputeReasonCode"`
	DisputeStatus      model.DisputeStatus `json:"disputeStatus" example:"opened" enums:"opened,evidence_required,evidence_submitted,under_review,won,lost,accepted,closed,canceled"`
	OpenedAt           time.Time           `json:"openedAt"`
	DueAt              time.Time           `json:"dueAt"`
	ClosedAt           time.Time           `json:"closedAt"`
	Metadata           json.RawMessage     `json:"metadata"`
}

func (d *DisputeCasesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d DisputeCasesUpdateRequest) ToModel() model.DisputeCases {
	return model.DisputeCases{
		DisputeCode:        d.DisputeCode,
		BookId:             d.BookId,
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		MerchantPartyId:    nuuid.From(d.MerchantPartyId),
		CustomerPartyId:    nuuid.From(d.CustomerPartyId),
		ProviderCode:       null.StringFrom(d.ProviderCode),
		ProviderDisputeRef: null.StringFrom(d.ProviderDisputeRef),
		IdempotencyKey:     null.StringFrom(d.IdempotencyKey),
		CurrencyCode:       d.CurrencyCode,
		DisputedAmount:     d.DisputedAmount,
		DisputeReasonCode:  d.DisputeReasonCode,
		DisputeStatus:      d.DisputeStatus,
		OpenedAt:           d.OpenedAt,
		DueAt:              null.TimeFrom(d.DueAt),
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Metadata:           d.Metadata,
	}
}

type DisputeCasesBulkUpdateRequest struct {
	Id                 uuid.UUID           `json:"id"`
	DisputeCode        string              `json:"disputeCode"`
	BookId             uuid.UUID           `json:"bookId"`
	SourceType         string              `json:"sourceType"`
	SourceId           uuid.UUID           `json:"sourceId"`
	MerchantPartyId    uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId    uuid.UUID           `json:"customerPartyId"`
	ProviderCode       string              `json:"providerCode"`
	ProviderDisputeRef string              `json:"providerDisputeRef"`
	IdempotencyKey     string              `json:"idempotencyKey"`
	CurrencyCode       string              `json:"currencyCode"`
	DisputedAmount     decimal.Decimal     `json:"disputedAmount"`
	DisputeReasonCode  string              `json:"disputeReasonCode"`
	DisputeStatus      model.DisputeStatus `json:"disputeStatus" example:"opened" enums:"opened,evidence_required,evidence_submitted,under_review,won,lost,accepted,closed,canceled"`
	OpenedAt           time.Time           `json:"openedAt"`
	DueAt              time.Time           `json:"dueAt"`
	ClosedAt           time.Time           `json:"closedAt"`
	Metadata           json.RawMessage     `json:"metadata"`
}

func (d DisputeCasesBulkUpdateRequest) PrimaryID() DisputeCasesPrimaryID {
	return DisputeCasesPrimaryID{
		Id: d.Id,
	}
}

type DisputeCasesListBulkUpdateRequest []*DisputeCasesBulkUpdateRequest

func (d DisputeCasesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeCases := range d {
		err = validator.Struct(disputeCases)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeCasesBulkUpdateRequest) ToModel() model.DisputeCases {
	return model.DisputeCases{
		Id:                 d.Id,
		DisputeCode:        d.DisputeCode,
		BookId:             d.BookId,
		SourceType:         d.SourceType,
		SourceId:           d.SourceId,
		MerchantPartyId:    nuuid.From(d.MerchantPartyId),
		CustomerPartyId:    nuuid.From(d.CustomerPartyId),
		ProviderCode:       null.StringFrom(d.ProviderCode),
		ProviderDisputeRef: null.StringFrom(d.ProviderDisputeRef),
		IdempotencyKey:     null.StringFrom(d.IdempotencyKey),
		CurrencyCode:       d.CurrencyCode,
		DisputedAmount:     d.DisputedAmount,
		DisputeReasonCode:  d.DisputeReasonCode,
		DisputeStatus:      d.DisputeStatus,
		OpenedAt:           d.OpenedAt,
		DueAt:              null.TimeFrom(d.DueAt),
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Metadata:           d.Metadata,
	}
}

type DisputeCasesResponse struct {
	Id                 uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DisputeCode        string              `json:"disputeCode" validate:"required"`
	BookId             uuid.UUID           `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType         string              `json:"sourceType" validate:"required"`
	SourceId           uuid.UUID           `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId    uuid.UUID           `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CustomerPartyId    uuid.UUID           `json:"customerPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode       string              `json:"providerCode"`
	ProviderDisputeRef string              `json:"providerDisputeRef"`
	IdempotencyKey     string              `json:"idempotencyKey"`
	CurrencyCode       string              `json:"currencyCode" validate:"required"`
	DisputedAmount     decimal.Decimal     `json:"disputedAmount" validate:"required" format:"decimal" example:"100.50"`
	DisputeReasonCode  string              `json:"disputeReasonCode" validate:"required"`
	DisputeStatus      model.DisputeStatus `json:"disputeStatus" validate:"oneof=opened evidence_required evidence_submitted under_review won lost accepted closed canceled" enums:"opened,evidence_required,evidence_submitted,under_review,won,lost,accepted,closed,canceled"`
	OpenedAt           time.Time           `json:"openedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	DueAt              time.Time           `json:"dueAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ClosedAt           time.Time           `json:"closedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata           json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewDisputeCasesResponse(disputeCases model.DisputeCases) DisputeCasesResponse {
	return DisputeCasesResponse{
		Id:                 disputeCases.Id,
		DisputeCode:        disputeCases.DisputeCode,
		BookId:             disputeCases.BookId,
		SourceType:         disputeCases.SourceType,
		SourceId:           disputeCases.SourceId,
		MerchantPartyId:    disputeCases.MerchantPartyId.UUID,
		CustomerPartyId:    disputeCases.CustomerPartyId.UUID,
		ProviderCode:       disputeCases.ProviderCode.String,
		ProviderDisputeRef: disputeCases.ProviderDisputeRef.String,
		IdempotencyKey:     disputeCases.IdempotencyKey.String,
		CurrencyCode:       disputeCases.CurrencyCode,
		DisputedAmount:     disputeCases.DisputedAmount,
		DisputeReasonCode:  disputeCases.DisputeReasonCode,
		DisputeStatus:      model.DisputeStatus(disputeCases.DisputeStatus),
		OpenedAt:           disputeCases.OpenedAt,
		DueAt:              disputeCases.DueAt.Time,
		ClosedAt:           disputeCases.ClosedAt.Time,
		Metadata:           disputeCases.Metadata,
	}
}

type DisputeCasesListResponse []*DisputeCasesResponse

func NewDisputeCasesListResponse(disputeCasesList model.DisputeCasesList) DisputeCasesListResponse {
	dtoDisputeCasesListResponse := DisputeCasesListResponse{}
	for _, disputeCases := range disputeCasesList {
		dtoDisputeCasesResponse := NewDisputeCasesResponse(*disputeCases)
		dtoDisputeCasesListResponse = append(dtoDisputeCasesListResponse, &dtoDisputeCasesResponse)
	}
	return dtoDisputeCasesListResponse
}

type DisputeCasesPrimaryIDList []DisputeCasesPrimaryID

func (d DisputeCasesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeCases := range d {
		err = validator.Struct(disputeCases)
		if err != nil {
			return
		}
	}
	return nil
}

type DisputeCasesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *DisputeCasesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d DisputeCasesPrimaryID) ToModel() model.DisputeCasesPrimaryID {
	return model.DisputeCasesPrimaryID{
		Id: d.Id,
	}
}

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

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type CreditNotesDTOFieldNameType string

type creditNotesDTOFieldName struct {
	Id            CreditNotesDTOFieldNameType
	CreditNoteNo  CreditNotesDTOFieldNameType
	TaxInvoiceId  CreditNotesDTOFieldNameType
	SourceType    CreditNotesDTOFieldNameType
	SourceId      CreditNotesDTOFieldNameType
	CurrencyCode  CreditNotesDTOFieldNameType
	TaxableAmount CreditNotesDTOFieldNameType
	TaxAmount     CreditNotesDTOFieldNameType
	TotalAmount   CreditNotesDTOFieldNameType
	ReasonCode    CreditNotesDTOFieldNameType
	ReasonDetail  CreditNotesDTOFieldNameType
	IssuedAt      CreditNotesDTOFieldNameType
	Metadata      CreditNotesDTOFieldNameType
	MetaCreatedAt CreditNotesDTOFieldNameType
	MetaCreatedBy CreditNotesDTOFieldNameType
	MetaUpdatedAt CreditNotesDTOFieldNameType
	MetaUpdatedBy CreditNotesDTOFieldNameType
	MetaDeletedAt CreditNotesDTOFieldNameType
	MetaDeletedBy CreditNotesDTOFieldNameType
}

var CreditNotesDTOFieldName = creditNotesDTOFieldName{
	Id:            "id",
	CreditNoteNo:  "creditNoteNo",
	TaxInvoiceId:  "taxInvoiceId",
	SourceType:    "sourceType",
	SourceId:      "sourceId",
	CurrencyCode:  "currencyCode",
	TaxableAmount: "taxableAmount",
	TaxAmount:     "taxAmount",
	TotalAmount:   "totalAmount",
	ReasonCode:    "reasonCode",
	ReasonDetail:  "reasonDetail",
	IssuedAt:      "issuedAt",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformCreditNotesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(CreditNotesDTOFieldName.Id):
		return string(model.CreditNotesDBFieldName.Id), true

	case string(CreditNotesDTOFieldName.CreditNoteNo):
		return string(model.CreditNotesDBFieldName.CreditNoteNo), true

	case string(CreditNotesDTOFieldName.TaxInvoiceId):
		return string(model.CreditNotesDBFieldName.TaxInvoiceId), true

	case string(CreditNotesDTOFieldName.SourceType):
		return string(model.CreditNotesDBFieldName.SourceType), true

	case string(CreditNotesDTOFieldName.SourceId):
		return string(model.CreditNotesDBFieldName.SourceId), true

	case string(CreditNotesDTOFieldName.CurrencyCode):
		return string(model.CreditNotesDBFieldName.CurrencyCode), true

	case string(CreditNotesDTOFieldName.TaxableAmount):
		return string(model.CreditNotesDBFieldName.TaxableAmount), true

	case string(CreditNotesDTOFieldName.TaxAmount):
		return string(model.CreditNotesDBFieldName.TaxAmount), true

	case string(CreditNotesDTOFieldName.TotalAmount):
		return string(model.CreditNotesDBFieldName.TotalAmount), true

	case string(CreditNotesDTOFieldName.ReasonCode):
		return string(model.CreditNotesDBFieldName.ReasonCode), true

	case string(CreditNotesDTOFieldName.ReasonDetail):
		return string(model.CreditNotesDBFieldName.ReasonDetail), true

	case string(CreditNotesDTOFieldName.IssuedAt):
		return string(model.CreditNotesDBFieldName.IssuedAt), true

	case string(CreditNotesDTOFieldName.Metadata):
		return string(model.CreditNotesDBFieldName.Metadata), true

	case string(CreditNotesDTOFieldName.MetaCreatedAt):
		return string(model.CreditNotesDBFieldName.MetaCreatedAt), true

	case string(CreditNotesDTOFieldName.MetaCreatedBy):
		return string(model.CreditNotesDBFieldName.MetaCreatedBy), true

	case string(CreditNotesDTOFieldName.MetaUpdatedAt):
		return string(model.CreditNotesDBFieldName.MetaUpdatedAt), true

	case string(CreditNotesDTOFieldName.MetaUpdatedBy):
		return string(model.CreditNotesDBFieldName.MetaUpdatedBy), true

	case string(CreditNotesDTOFieldName.MetaDeletedAt):
		return string(model.CreditNotesDBFieldName.MetaDeletedAt), true

	case string(CreditNotesDTOFieldName.MetaDeletedBy):
		return string(model.CreditNotesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewCreditNotesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isCreditNotesBaseFilterField(field string) bool {
	spec, found := model.NewCreditNotesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeCreditNotesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateCreditNotesProjectionOutputPath(path string) error {
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

func transformCreditNotesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformCreditNotesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformCreditNotesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformCreditNotesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformCreditNotesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isCreditNotesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateCreditNotesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeCreditNotesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformCreditNotesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformCreditNotesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformCreditNotesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultCreditNotesFilter(filter *model.Filter) {
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
			Field: string(CreditNotesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type CreditNotesSelectableResponse map[string]interface{}
type CreditNotesSelectableListResponse []*CreditNotesSelectableResponse

func assignCreditNotesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setCreditNotesSelectableValue(out CreditNotesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignCreditNotesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewCreditNotesSelectableResponse(creditNotes model.CreditNotes, filter model.Filter) CreditNotesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.CreditNotesDBFieldName.Id),
			string(model.CreditNotesDBFieldName.CreditNoteNo),
			string(model.CreditNotesDBFieldName.TaxInvoiceId),
			string(model.CreditNotesDBFieldName.SourceType),
			string(model.CreditNotesDBFieldName.SourceId),
			string(model.CreditNotesDBFieldName.CurrencyCode),
			string(model.CreditNotesDBFieldName.TaxableAmount),
			string(model.CreditNotesDBFieldName.TaxAmount),
			string(model.CreditNotesDBFieldName.TotalAmount),
			string(model.CreditNotesDBFieldName.ReasonCode),
			string(model.CreditNotesDBFieldName.ReasonDetail),
			string(model.CreditNotesDBFieldName.IssuedAt),
			string(model.CreditNotesDBFieldName.Metadata),
			string(model.CreditNotesDBFieldName.MetaCreatedAt),
			string(model.CreditNotesDBFieldName.MetaCreatedBy),
			string(model.CreditNotesDBFieldName.MetaUpdatedAt),
			string(model.CreditNotesDBFieldName.MetaUpdatedBy),
			string(model.CreditNotesDBFieldName.MetaDeletedAt),
			string(model.CreditNotesDBFieldName.MetaDeletedBy),
		)
	}
	creditNotesSelectableResponse := CreditNotesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.CreditNotesDBFieldName.Id):
			key := string(CreditNotesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.Id, explicitAlias)

		case string(model.CreditNotesDBFieldName.CreditNoteNo):
			key := string(CreditNotesDTOFieldName.CreditNoteNo)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.CreditNoteNo, explicitAlias)

		case string(model.CreditNotesDBFieldName.TaxInvoiceId):
			key := string(CreditNotesDTOFieldName.TaxInvoiceId)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.TaxInvoiceId, explicitAlias)

		case string(model.CreditNotesDBFieldName.SourceType):
			key := string(CreditNotesDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.SourceType, explicitAlias)

		case string(model.CreditNotesDBFieldName.SourceId):
			key := string(CreditNotesDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.SourceId, explicitAlias)

		case string(model.CreditNotesDBFieldName.CurrencyCode):
			key := string(CreditNotesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.CurrencyCode, explicitAlias)

		case string(model.CreditNotesDBFieldName.TaxableAmount):
			key := string(CreditNotesDTOFieldName.TaxableAmount)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.TaxableAmount, explicitAlias)

		case string(model.CreditNotesDBFieldName.TaxAmount):
			key := string(CreditNotesDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.TaxAmount, explicitAlias)

		case string(model.CreditNotesDBFieldName.TotalAmount):
			key := string(CreditNotesDTOFieldName.TotalAmount)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.TotalAmount, explicitAlias)

		case string(model.CreditNotesDBFieldName.ReasonCode):
			key := string(CreditNotesDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.ReasonCode, explicitAlias)

		case string(model.CreditNotesDBFieldName.ReasonDetail):
			key := string(CreditNotesDTOFieldName.ReasonDetail)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.ReasonDetail.String, explicitAlias)

		case string(model.CreditNotesDBFieldName.IssuedAt):
			key := string(CreditNotesDTOFieldName.IssuedAt)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.IssuedAt, explicitAlias)

		case string(model.CreditNotesDBFieldName.Metadata):
			key := string(CreditNotesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.Metadata, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaCreatedAt):
			key := string(CreditNotesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaCreatedAt, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaCreatedBy):
			key := string(CreditNotesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaCreatedBy, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaUpdatedAt):
			key := string(CreditNotesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaUpdatedAt, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaUpdatedBy):
			key := string(CreditNotesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaUpdatedBy, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaDeletedAt):
			key := string(CreditNotesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaDeletedAt.Time, explicitAlias)

		case string(model.CreditNotesDBFieldName.MetaDeletedBy):
			key := string(CreditNotesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setCreditNotesSelectableValue(creditNotesSelectableResponse, key, creditNotes.MetaDeletedBy, explicitAlias)

		}
	}
	return creditNotesSelectableResponse
}

func NewCreditNotesListResponseFromFilterResult(result []model.CreditNotesFilterResult, filter model.Filter) CreditNotesSelectableListResponse {
	dtoCreditNotesListResponse := CreditNotesSelectableListResponse{}
	for _, row := range result {
		dtoCreditNotesResponse := NewCreditNotesSelectableResponse(row.CreditNotes, filter)
		dtoCreditNotesListResponse = append(dtoCreditNotesListResponse, &dtoCreditNotesResponse)
	}
	return dtoCreditNotesListResponse
}

type CreditNotesFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     CreditNotesSelectableListResponse `json:"data"`
}

func reverseCreditNotesFilterResults(result []model.CreditNotesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewCreditNotesFilterResponse(result []model.CreditNotesFilterResult, filter model.Filter) (resp CreditNotesFilterResponse) {
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
			reverseCreditNotesFilterResults(dataResult)
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

	resp.Data = NewCreditNotesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type CreditNotesCreateRequest struct {
	CreditNoteNo  string          `json:"creditNoteNo"`
	TaxInvoiceId  uuid.UUID       `json:"taxInvoiceId"`
	SourceType    string          `json:"sourceType"`
	SourceId      uuid.UUID       `json:"sourceId"`
	CurrencyCode  string          `json:"currencyCode"`
	TaxableAmount decimal.Decimal `json:"taxableAmount"`
	TaxAmount     decimal.Decimal `json:"taxAmount"`
	TotalAmount   decimal.Decimal `json:"totalAmount"`
	ReasonCode    string          `json:"reasonCode"`
	ReasonDetail  string          `json:"reasonDetail"`
	IssuedAt      time.Time       `json:"issuedAt"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *CreditNotesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *CreditNotesCreateRequest) ToModel() model.CreditNotes {
	id, _ := uuid.NewV4()
	return model.CreditNotes{
		Id:            id,
		CreditNoteNo:  d.CreditNoteNo,
		TaxInvoiceId:  d.TaxInvoiceId,
		SourceType:    d.SourceType,
		SourceId:      d.SourceId,
		CurrencyCode:  d.CurrencyCode,
		TaxableAmount: d.TaxableAmount,
		TaxAmount:     d.TaxAmount,
		TotalAmount:   d.TotalAmount,
		ReasonCode:    d.ReasonCode,
		ReasonDetail:  null.StringFrom(d.ReasonDetail),
		IssuedAt:      d.IssuedAt,
		Metadata:      d.Metadata,
	}
}

type CreditNotesListCreateRequest []*CreditNotesCreateRequest

func (d CreditNotesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, creditNotes := range d {
		err = validator.Struct(creditNotes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CreditNotesListCreateRequest) ToModelList() []model.CreditNotes {
	out := make([]model.CreditNotes, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type CreditNotesUpdateRequest struct {
	CreditNoteNo  string          `json:"creditNoteNo"`
	TaxInvoiceId  uuid.UUID       `json:"taxInvoiceId"`
	SourceType    string          `json:"sourceType"`
	SourceId      uuid.UUID       `json:"sourceId"`
	CurrencyCode  string          `json:"currencyCode"`
	TaxableAmount decimal.Decimal `json:"taxableAmount"`
	TaxAmount     decimal.Decimal `json:"taxAmount"`
	TotalAmount   decimal.Decimal `json:"totalAmount"`
	ReasonCode    string          `json:"reasonCode"`
	ReasonDetail  string          `json:"reasonDetail"`
	IssuedAt      time.Time       `json:"issuedAt"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d *CreditNotesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d CreditNotesUpdateRequest) ToModel() model.CreditNotes {
	return model.CreditNotes{
		CreditNoteNo:  d.CreditNoteNo,
		TaxInvoiceId:  d.TaxInvoiceId,
		SourceType:    d.SourceType,
		SourceId:      d.SourceId,
		CurrencyCode:  d.CurrencyCode,
		TaxableAmount: d.TaxableAmount,
		TaxAmount:     d.TaxAmount,
		TotalAmount:   d.TotalAmount,
		ReasonCode:    d.ReasonCode,
		ReasonDetail:  null.StringFrom(d.ReasonDetail),
		IssuedAt:      d.IssuedAt,
		Metadata:      d.Metadata,
	}
}

type CreditNotesBulkUpdateRequest struct {
	Id            uuid.UUID       `json:"id"`
	CreditNoteNo  string          `json:"creditNoteNo"`
	TaxInvoiceId  uuid.UUID       `json:"taxInvoiceId"`
	SourceType    string          `json:"sourceType"`
	SourceId      uuid.UUID       `json:"sourceId"`
	CurrencyCode  string          `json:"currencyCode"`
	TaxableAmount decimal.Decimal `json:"taxableAmount"`
	TaxAmount     decimal.Decimal `json:"taxAmount"`
	TotalAmount   decimal.Decimal `json:"totalAmount"`
	ReasonCode    string          `json:"reasonCode"`
	ReasonDetail  string          `json:"reasonDetail"`
	IssuedAt      time.Time       `json:"issuedAt"`
	Metadata      json.RawMessage `json:"metadata"`
}

func (d CreditNotesBulkUpdateRequest) PrimaryID() CreditNotesPrimaryID {
	return CreditNotesPrimaryID{
		Id: d.Id,
	}
}

type CreditNotesListBulkUpdateRequest []*CreditNotesBulkUpdateRequest

func (d CreditNotesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, creditNotes := range d {
		err = validator.Struct(creditNotes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CreditNotesBulkUpdateRequest) ToModel() model.CreditNotes {
	return model.CreditNotes{
		Id:            d.Id,
		CreditNoteNo:  d.CreditNoteNo,
		TaxInvoiceId:  d.TaxInvoiceId,
		SourceType:    d.SourceType,
		SourceId:      d.SourceId,
		CurrencyCode:  d.CurrencyCode,
		TaxableAmount: d.TaxableAmount,
		TaxAmount:     d.TaxAmount,
		TotalAmount:   d.TotalAmount,
		ReasonCode:    d.ReasonCode,
		ReasonDetail:  null.StringFrom(d.ReasonDetail),
		IssuedAt:      d.IssuedAt,
		Metadata:      d.Metadata,
	}
}

type CreditNotesResponse struct {
	Id            uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CreditNoteNo  string          `json:"creditNoteNo" validate:"required"`
	TaxInvoiceId  uuid.UUID       `json:"taxInvoiceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType    string          `json:"sourceType" validate:"required"`
	SourceId      uuid.UUID       `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode  string          `json:"currencyCode" validate:"required"`
	TaxableAmount decimal.Decimal `json:"taxableAmount" validate:"required" format:"decimal" example:"100.50"`
	TaxAmount     decimal.Decimal `json:"taxAmount" validate:"required" format:"decimal" example:"100.50"`
	TotalAmount   decimal.Decimal `json:"totalAmount" validate:"required" format:"decimal" example:"100.50"`
	ReasonCode    string          `json:"reasonCode" validate:"required"`
	ReasonDetail  string          `json:"reasonDetail"`
	IssuedAt      time.Time       `json:"issuedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata      json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewCreditNotesResponse(creditNotes model.CreditNotes) CreditNotesResponse {
	return CreditNotesResponse{
		Id:            creditNotes.Id,
		CreditNoteNo:  creditNotes.CreditNoteNo,
		TaxInvoiceId:  creditNotes.TaxInvoiceId,
		SourceType:    creditNotes.SourceType,
		SourceId:      creditNotes.SourceId,
		CurrencyCode:  creditNotes.CurrencyCode,
		TaxableAmount: creditNotes.TaxableAmount,
		TaxAmount:     creditNotes.TaxAmount,
		TotalAmount:   creditNotes.TotalAmount,
		ReasonCode:    creditNotes.ReasonCode,
		ReasonDetail:  creditNotes.ReasonDetail.String,
		IssuedAt:      creditNotes.IssuedAt,
		Metadata:      creditNotes.Metadata,
	}
}

type CreditNotesListResponse []*CreditNotesResponse

func NewCreditNotesListResponse(creditNotesList model.CreditNotesList) CreditNotesListResponse {
	dtoCreditNotesListResponse := CreditNotesListResponse{}
	for _, creditNotes := range creditNotesList {
		dtoCreditNotesResponse := NewCreditNotesResponse(*creditNotes)
		dtoCreditNotesListResponse = append(dtoCreditNotesListResponse, &dtoCreditNotesResponse)
	}
	return dtoCreditNotesListResponse
}

type CreditNotesPrimaryIDList []CreditNotesPrimaryID

func (d CreditNotesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, creditNotes := range d {
		err = validator.Struct(creditNotes)
		if err != nil {
			return
		}
	}
	return nil
}

type CreditNotesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *CreditNotesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d CreditNotesPrimaryID) ToModel() model.CreditNotesPrimaryID {
	return model.CreditNotesPrimaryID{
		Id: d.Id,
	}
}

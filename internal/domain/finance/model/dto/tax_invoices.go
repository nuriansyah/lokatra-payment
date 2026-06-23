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

type TaxInvoicesDTOFieldNameType string

type taxInvoicesDTOFieldName struct {
	Id              TaxInvoicesDTOFieldNameType
	InvoiceNo       TaxInvoicesDTOFieldNameType
	SourceType      TaxInvoicesDTOFieldNameType
	SourceId        TaxInvoicesDTOFieldNameType
	MerchantPartyId TaxInvoicesDTOFieldNameType
	CustomerPartyId TaxInvoicesDTOFieldNameType
	CurrencyCode    TaxInvoicesDTOFieldNameType
	TaxableAmount   TaxInvoicesDTOFieldNameType
	TaxAmount       TaxInvoicesDTOFieldNameType
	TotalAmount     TaxInvoicesDTOFieldNameType
	InvoiceStatus   TaxInvoicesDTOFieldNameType
	IssuedAt        TaxInvoicesDTOFieldNameType
	Metadata        TaxInvoicesDTOFieldNameType
	MetaCreatedAt   TaxInvoicesDTOFieldNameType
	MetaCreatedBy   TaxInvoicesDTOFieldNameType
	MetaUpdatedAt   TaxInvoicesDTOFieldNameType
	MetaUpdatedBy   TaxInvoicesDTOFieldNameType
	MetaDeletedAt   TaxInvoicesDTOFieldNameType
	MetaDeletedBy   TaxInvoicesDTOFieldNameType
}

var TaxInvoicesDTOFieldName = taxInvoicesDTOFieldName{
	Id:              "id",
	InvoiceNo:       "invoiceNo",
	SourceType:      "sourceType",
	SourceId:        "sourceId",
	MerchantPartyId: "merchantPartyId",
	CustomerPartyId: "customerPartyId",
	CurrencyCode:    "currencyCode",
	TaxableAmount:   "taxableAmount",
	TaxAmount:       "taxAmount",
	TotalAmount:     "totalAmount",
	InvoiceStatus:   "invoiceStatus",
	IssuedAt:        "issuedAt",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformTaxInvoicesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxInvoicesDTOFieldName.Id):
		return string(model.TaxInvoicesDBFieldName.Id), true

	case string(TaxInvoicesDTOFieldName.InvoiceNo):
		return string(model.TaxInvoicesDBFieldName.InvoiceNo), true

	case string(TaxInvoicesDTOFieldName.SourceType):
		return string(model.TaxInvoicesDBFieldName.SourceType), true

	case string(TaxInvoicesDTOFieldName.SourceId):
		return string(model.TaxInvoicesDBFieldName.SourceId), true

	case string(TaxInvoicesDTOFieldName.MerchantPartyId):
		return string(model.TaxInvoicesDBFieldName.MerchantPartyId), true

	case string(TaxInvoicesDTOFieldName.CustomerPartyId):
		return string(model.TaxInvoicesDBFieldName.CustomerPartyId), true

	case string(TaxInvoicesDTOFieldName.CurrencyCode):
		return string(model.TaxInvoicesDBFieldName.CurrencyCode), true

	case string(TaxInvoicesDTOFieldName.TaxableAmount):
		return string(model.TaxInvoicesDBFieldName.TaxableAmount), true

	case string(TaxInvoicesDTOFieldName.TaxAmount):
		return string(model.TaxInvoicesDBFieldName.TaxAmount), true

	case string(TaxInvoicesDTOFieldName.TotalAmount):
		return string(model.TaxInvoicesDBFieldName.TotalAmount), true

	case string(TaxInvoicesDTOFieldName.InvoiceStatus):
		return string(model.TaxInvoicesDBFieldName.InvoiceStatus), true

	case string(TaxInvoicesDTOFieldName.IssuedAt):
		return string(model.TaxInvoicesDBFieldName.IssuedAt), true

	case string(TaxInvoicesDTOFieldName.Metadata):
		return string(model.TaxInvoicesDBFieldName.Metadata), true

	case string(TaxInvoicesDTOFieldName.MetaCreatedAt):
		return string(model.TaxInvoicesDBFieldName.MetaCreatedAt), true

	case string(TaxInvoicesDTOFieldName.MetaCreatedBy):
		return string(model.TaxInvoicesDBFieldName.MetaCreatedBy), true

	case string(TaxInvoicesDTOFieldName.MetaUpdatedAt):
		return string(model.TaxInvoicesDBFieldName.MetaUpdatedAt), true

	case string(TaxInvoicesDTOFieldName.MetaUpdatedBy):
		return string(model.TaxInvoicesDBFieldName.MetaUpdatedBy), true

	case string(TaxInvoicesDTOFieldName.MetaDeletedAt):
		return string(model.TaxInvoicesDBFieldName.MetaDeletedAt), true

	case string(TaxInvoicesDTOFieldName.MetaDeletedBy):
		return string(model.TaxInvoicesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxInvoicesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxInvoicesBaseFilterField(field string) bool {
	spec, found := model.NewTaxInvoicesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxInvoicesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxInvoicesProjectionOutputPath(path string) error {
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

func transformTaxInvoicesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxInvoicesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxInvoicesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxInvoicesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxInvoicesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxInvoicesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxInvoicesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxInvoicesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxInvoicesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxInvoicesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxInvoicesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxInvoicesFilter(filter *model.Filter) {
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
			Field: string(TaxInvoicesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type TaxInvoicesSelectableResponse map[string]interface{}
type TaxInvoicesSelectableListResponse []*TaxInvoicesSelectableResponse

func assignTaxInvoicesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxInvoicesSelectableValue(out TaxInvoicesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxInvoicesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxInvoicesSelectableResponse(taxInvoices model.TaxInvoices, filter model.Filter) TaxInvoicesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxInvoicesDBFieldName.Id),
			string(model.TaxInvoicesDBFieldName.InvoiceNo),
			string(model.TaxInvoicesDBFieldName.SourceType),
			string(model.TaxInvoicesDBFieldName.SourceId),
			string(model.TaxInvoicesDBFieldName.MerchantPartyId),
			string(model.TaxInvoicesDBFieldName.CustomerPartyId),
			string(model.TaxInvoicesDBFieldName.CurrencyCode),
			string(model.TaxInvoicesDBFieldName.TaxableAmount),
			string(model.TaxInvoicesDBFieldName.TaxAmount),
			string(model.TaxInvoicesDBFieldName.TotalAmount),
			string(model.TaxInvoicesDBFieldName.InvoiceStatus),
			string(model.TaxInvoicesDBFieldName.IssuedAt),
			string(model.TaxInvoicesDBFieldName.Metadata),
			string(model.TaxInvoicesDBFieldName.MetaCreatedAt),
			string(model.TaxInvoicesDBFieldName.MetaCreatedBy),
			string(model.TaxInvoicesDBFieldName.MetaUpdatedAt),
			string(model.TaxInvoicesDBFieldName.MetaUpdatedBy),
			string(model.TaxInvoicesDBFieldName.MetaDeletedAt),
			string(model.TaxInvoicesDBFieldName.MetaDeletedBy),
		)
	}
	taxInvoicesSelectableResponse := TaxInvoicesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxInvoicesDBFieldName.Id):
			key := string(TaxInvoicesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.Id, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.InvoiceNo):
			key := string(TaxInvoicesDTOFieldName.InvoiceNo)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.InvoiceNo, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.SourceType):
			key := string(TaxInvoicesDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.SourceType, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.SourceId):
			key := string(TaxInvoicesDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.SourceId, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MerchantPartyId):
			key := string(TaxInvoicesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MerchantPartyId.UUID, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.CustomerPartyId):
			key := string(TaxInvoicesDTOFieldName.CustomerPartyId)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.CustomerPartyId.UUID, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.CurrencyCode):
			key := string(TaxInvoicesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.CurrencyCode, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.TaxableAmount):
			key := string(TaxInvoicesDTOFieldName.TaxableAmount)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.TaxableAmount, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.TaxAmount):
			key := string(TaxInvoicesDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.TaxAmount, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.TotalAmount):
			key := string(TaxInvoicesDTOFieldName.TotalAmount)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.TotalAmount, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.InvoiceStatus):
			key := string(TaxInvoicesDTOFieldName.InvoiceStatus)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, model.InvoiceStatus(taxInvoices.InvoiceStatus), explicitAlias)

		case string(model.TaxInvoicesDBFieldName.IssuedAt):
			key := string(TaxInvoicesDTOFieldName.IssuedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.IssuedAt.Time, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.Metadata):
			key := string(TaxInvoicesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.Metadata, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaCreatedAt):
			key := string(TaxInvoicesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaCreatedAt, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaCreatedBy):
			key := string(TaxInvoicesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaCreatedBy, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaUpdatedAt):
			key := string(TaxInvoicesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaUpdatedAt, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaUpdatedBy):
			key := string(TaxInvoicesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaUpdatedBy, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaDeletedAt):
			key := string(TaxInvoicesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxInvoicesDBFieldName.MetaDeletedBy):
			key := string(TaxInvoicesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxInvoicesSelectableValue(taxInvoicesSelectableResponse, key, taxInvoices.MetaDeletedBy, explicitAlias)

		}
	}
	return taxInvoicesSelectableResponse
}

func NewTaxInvoicesListResponseFromFilterResult(result []model.TaxInvoicesFilterResult, filter model.Filter) TaxInvoicesSelectableListResponse {
	dtoTaxInvoicesListResponse := TaxInvoicesSelectableListResponse{}
	for _, row := range result {
		dtoTaxInvoicesResponse := NewTaxInvoicesSelectableResponse(row.TaxInvoices, filter)
		dtoTaxInvoicesListResponse = append(dtoTaxInvoicesListResponse, &dtoTaxInvoicesResponse)
	}
	return dtoTaxInvoicesListResponse
}

type TaxInvoicesFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     TaxInvoicesSelectableListResponse `json:"data"`
}

func reverseTaxInvoicesFilterResults(result []model.TaxInvoicesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxInvoicesFilterResponse(result []model.TaxInvoicesFilterResult, filter model.Filter) (resp TaxInvoicesFilterResponse) {
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
			reverseTaxInvoicesFilterResults(dataResult)
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

	resp.Data = NewTaxInvoicesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxInvoicesCreateRequest struct {
	InvoiceNo       string              `json:"invoiceNo"`
	SourceType      string              `json:"sourceType"`
	SourceId        uuid.UUID           `json:"sourceId"`
	MerchantPartyId uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId uuid.UUID           `json:"customerPartyId"`
	CurrencyCode    string              `json:"currencyCode"`
	TaxableAmount   decimal.Decimal     `json:"taxableAmount"`
	TaxAmount       decimal.Decimal     `json:"taxAmount"`
	TotalAmount     decimal.Decimal     `json:"totalAmount"`
	InvoiceStatus   model.InvoiceStatus `json:"invoiceStatus" example:"draft" enums:"draft,issued,void"`
	IssuedAt        time.Time           `json:"issuedAt"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d *TaxInvoicesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxInvoicesCreateRequest) ToModel() model.TaxInvoices {
	id, _ := uuid.NewV4()
	return model.TaxInvoices{
		Id:              id,
		InvoiceNo:       d.InvoiceNo,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		CustomerPartyId: nuuid.From(d.CustomerPartyId),
		CurrencyCode:    d.CurrencyCode,
		TaxableAmount:   d.TaxableAmount,
		TaxAmount:       d.TaxAmount,
		TotalAmount:     d.TotalAmount,
		InvoiceStatus:   d.InvoiceStatus,
		IssuedAt:        null.TimeFrom(d.IssuedAt),
		Metadata:        d.Metadata,
	}
}

type TaxInvoicesListCreateRequest []*TaxInvoicesCreateRequest

func (d TaxInvoicesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxInvoices := range d {
		err = validator.Struct(taxInvoices)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxInvoicesListCreateRequest) ToModelList() []model.TaxInvoices {
	out := make([]model.TaxInvoices, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxInvoicesUpdateRequest struct {
	InvoiceNo       string              `json:"invoiceNo"`
	SourceType      string              `json:"sourceType"`
	SourceId        uuid.UUID           `json:"sourceId"`
	MerchantPartyId uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId uuid.UUID           `json:"customerPartyId"`
	CurrencyCode    string              `json:"currencyCode"`
	TaxableAmount   decimal.Decimal     `json:"taxableAmount"`
	TaxAmount       decimal.Decimal     `json:"taxAmount"`
	TotalAmount     decimal.Decimal     `json:"totalAmount"`
	InvoiceStatus   model.InvoiceStatus `json:"invoiceStatus" example:"draft" enums:"draft,issued,void"`
	IssuedAt        time.Time           `json:"issuedAt"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d *TaxInvoicesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxInvoicesUpdateRequest) ToModel() model.TaxInvoices {
	return model.TaxInvoices{
		InvoiceNo:       d.InvoiceNo,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		CustomerPartyId: nuuid.From(d.CustomerPartyId),
		CurrencyCode:    d.CurrencyCode,
		TaxableAmount:   d.TaxableAmount,
		TaxAmount:       d.TaxAmount,
		TotalAmount:     d.TotalAmount,
		InvoiceStatus:   d.InvoiceStatus,
		IssuedAt:        null.TimeFrom(d.IssuedAt),
		Metadata:        d.Metadata,
	}
}

type TaxInvoicesBulkUpdateRequest struct {
	Id              uuid.UUID           `json:"id"`
	InvoiceNo       string              `json:"invoiceNo"`
	SourceType      string              `json:"sourceType"`
	SourceId        uuid.UUID           `json:"sourceId"`
	MerchantPartyId uuid.UUID           `json:"merchantPartyId"`
	CustomerPartyId uuid.UUID           `json:"customerPartyId"`
	CurrencyCode    string              `json:"currencyCode"`
	TaxableAmount   decimal.Decimal     `json:"taxableAmount"`
	TaxAmount       decimal.Decimal     `json:"taxAmount"`
	TotalAmount     decimal.Decimal     `json:"totalAmount"`
	InvoiceStatus   model.InvoiceStatus `json:"invoiceStatus" example:"draft" enums:"draft,issued,void"`
	IssuedAt        time.Time           `json:"issuedAt"`
	Metadata        json.RawMessage     `json:"metadata"`
}

func (d TaxInvoicesBulkUpdateRequest) PrimaryID() TaxInvoicesPrimaryID {
	return TaxInvoicesPrimaryID{
		Id: d.Id,
	}
}

type TaxInvoicesListBulkUpdateRequest []*TaxInvoicesBulkUpdateRequest

func (d TaxInvoicesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxInvoices := range d {
		err = validator.Struct(taxInvoices)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxInvoicesBulkUpdateRequest) ToModel() model.TaxInvoices {
	return model.TaxInvoices{
		Id:              d.Id,
		InvoiceNo:       d.InvoiceNo,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		MerchantPartyId: nuuid.From(d.MerchantPartyId),
		CustomerPartyId: nuuid.From(d.CustomerPartyId),
		CurrencyCode:    d.CurrencyCode,
		TaxableAmount:   d.TaxableAmount,
		TaxAmount:       d.TaxAmount,
		TotalAmount:     d.TotalAmount,
		InvoiceStatus:   d.InvoiceStatus,
		IssuedAt:        null.TimeFrom(d.IssuedAt),
		Metadata:        d.Metadata,
	}
}

type TaxInvoicesResponse struct {
	Id              uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	InvoiceNo       string              `json:"invoiceNo" validate:"required"`
	SourceType      string              `json:"sourceType" validate:"required"`
	SourceId        uuid.UUID           `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId uuid.UUID           `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CustomerPartyId uuid.UUID           `json:"customerPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode    string              `json:"currencyCode" validate:"required"`
	TaxableAmount   decimal.Decimal     `json:"taxableAmount" validate:"required" format:"decimal" example:"100.50"`
	TaxAmount       decimal.Decimal     `json:"taxAmount" validate:"required" format:"decimal" example:"100.50"`
	TotalAmount     decimal.Decimal     `json:"totalAmount" validate:"required" format:"decimal" example:"100.50"`
	InvoiceStatus   model.InvoiceStatus `json:"invoiceStatus" validate:"oneof=draft issued void" enums:"draft,issued,void"`
	IssuedAt        time.Time           `json:"issuedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata        json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewTaxInvoicesResponse(taxInvoices model.TaxInvoices) TaxInvoicesResponse {
	return TaxInvoicesResponse{
		Id:              taxInvoices.Id,
		InvoiceNo:       taxInvoices.InvoiceNo,
		SourceType:      taxInvoices.SourceType,
		SourceId:        taxInvoices.SourceId,
		MerchantPartyId: taxInvoices.MerchantPartyId.UUID,
		CustomerPartyId: taxInvoices.CustomerPartyId.UUID,
		CurrencyCode:    taxInvoices.CurrencyCode,
		TaxableAmount:   taxInvoices.TaxableAmount,
		TaxAmount:       taxInvoices.TaxAmount,
		TotalAmount:     taxInvoices.TotalAmount,
		InvoiceStatus:   model.InvoiceStatus(taxInvoices.InvoiceStatus),
		IssuedAt:        taxInvoices.IssuedAt.Time,
		Metadata:        taxInvoices.Metadata,
	}
}

type TaxInvoicesListResponse []*TaxInvoicesResponse

func NewTaxInvoicesListResponse(taxInvoicesList model.TaxInvoicesList) TaxInvoicesListResponse {
	dtoTaxInvoicesListResponse := TaxInvoicesListResponse{}
	for _, taxInvoices := range taxInvoicesList {
		dtoTaxInvoicesResponse := NewTaxInvoicesResponse(*taxInvoices)
		dtoTaxInvoicesListResponse = append(dtoTaxInvoicesListResponse, &dtoTaxInvoicesResponse)
	}
	return dtoTaxInvoicesListResponse
}

type TaxInvoicesPrimaryIDList []TaxInvoicesPrimaryID

func (d TaxInvoicesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxInvoices := range d {
		err = validator.Struct(taxInvoices)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxInvoicesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *TaxInvoicesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxInvoicesPrimaryID) ToModel() model.TaxInvoicesPrimaryID {
	return model.TaxInvoicesPrimaryID{
		Id: d.Id,
	}
}

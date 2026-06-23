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

type ProviderStatementLinesDTOFieldNameType string

type providerStatementLinesDTOFieldName struct {
	Id              ProviderStatementLinesDTOFieldNameType
	StatementFileId ProviderStatementLinesDTOFieldNameType
	LineNo          ProviderStatementLinesDTOFieldNameType
	ProviderRef     ProviderStatementLinesDTOFieldNameType
	TransactionType ProviderStatementLinesDTOFieldNameType
	BookingDate     ProviderStatementLinesDTOFieldNameType
	ValueDate       ProviderStatementLinesDTOFieldNameType
	CurrencyCode    ProviderStatementLinesDTOFieldNameType
	GrossAmount     ProviderStatementLinesDTOFieldNameType
	FeeAmount       ProviderStatementLinesDTOFieldNameType
	NetAmount       ProviderStatementLinesDTOFieldNameType
	RawLine         ProviderStatementLinesDTOFieldNameType
	LineHash        ProviderStatementLinesDTOFieldNameType
	Metadata        ProviderStatementLinesDTOFieldNameType
	MetaCreatedAt   ProviderStatementLinesDTOFieldNameType
	MetaCreatedBy   ProviderStatementLinesDTOFieldNameType
	MetaUpdatedAt   ProviderStatementLinesDTOFieldNameType
	MetaUpdatedBy   ProviderStatementLinesDTOFieldNameType
	MetaDeletedAt   ProviderStatementLinesDTOFieldNameType
	MetaDeletedBy   ProviderStatementLinesDTOFieldNameType
}

var ProviderStatementLinesDTOFieldName = providerStatementLinesDTOFieldName{
	Id:              "id",
	StatementFileId: "statementFileId",
	LineNo:          "lineNo",
	ProviderRef:     "providerRef",
	TransactionType: "transactionType",
	BookingDate:     "bookingDate",
	ValueDate:       "valueDate",
	CurrencyCode:    "currencyCode",
	GrossAmount:     "grossAmount",
	FeeAmount:       "feeAmount",
	NetAmount:       "netAmount",
	RawLine:         "rawLine",
	LineHash:        "lineHash",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformProviderStatementLinesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderStatementLinesDTOFieldName.Id):
		return string(model.ProviderStatementLinesDBFieldName.Id), true

	case string(ProviderStatementLinesDTOFieldName.StatementFileId):
		return string(model.ProviderStatementLinesDBFieldName.StatementFileId), true

	case string(ProviderStatementLinesDTOFieldName.LineNo):
		return string(model.ProviderStatementLinesDBFieldName.LineNo), true

	case string(ProviderStatementLinesDTOFieldName.ProviderRef):
		return string(model.ProviderStatementLinesDBFieldName.ProviderRef), true

	case string(ProviderStatementLinesDTOFieldName.TransactionType):
		return string(model.ProviderStatementLinesDBFieldName.TransactionType), true

	case string(ProviderStatementLinesDTOFieldName.BookingDate):
		return string(model.ProviderStatementLinesDBFieldName.BookingDate), true

	case string(ProviderStatementLinesDTOFieldName.ValueDate):
		return string(model.ProviderStatementLinesDBFieldName.ValueDate), true

	case string(ProviderStatementLinesDTOFieldName.CurrencyCode):
		return string(model.ProviderStatementLinesDBFieldName.CurrencyCode), true

	case string(ProviderStatementLinesDTOFieldName.GrossAmount):
		return string(model.ProviderStatementLinesDBFieldName.GrossAmount), true

	case string(ProviderStatementLinesDTOFieldName.FeeAmount):
		return string(model.ProviderStatementLinesDBFieldName.FeeAmount), true

	case string(ProviderStatementLinesDTOFieldName.NetAmount):
		return string(model.ProviderStatementLinesDBFieldName.NetAmount), true

	case string(ProviderStatementLinesDTOFieldName.RawLine):
		return string(model.ProviderStatementLinesDBFieldName.RawLine), true

	case string(ProviderStatementLinesDTOFieldName.LineHash):
		return string(model.ProviderStatementLinesDBFieldName.LineHash), true

	case string(ProviderStatementLinesDTOFieldName.Metadata):
		return string(model.ProviderStatementLinesDBFieldName.Metadata), true

	case string(ProviderStatementLinesDTOFieldName.MetaCreatedAt):
		return string(model.ProviderStatementLinesDBFieldName.MetaCreatedAt), true

	case string(ProviderStatementLinesDTOFieldName.MetaCreatedBy):
		return string(model.ProviderStatementLinesDBFieldName.MetaCreatedBy), true

	case string(ProviderStatementLinesDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderStatementLinesDBFieldName.MetaUpdatedAt), true

	case string(ProviderStatementLinesDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderStatementLinesDBFieldName.MetaUpdatedBy), true

	case string(ProviderStatementLinesDTOFieldName.MetaDeletedAt):
		return string(model.ProviderStatementLinesDBFieldName.MetaDeletedAt), true

	case string(ProviderStatementLinesDTOFieldName.MetaDeletedBy):
		return string(model.ProviderStatementLinesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderStatementLinesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderStatementLinesBaseFilterField(field string) bool {
	spec, found := model.NewProviderStatementLinesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderStatementLinesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderStatementLinesProjectionOutputPath(path string) error {
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

func transformProviderStatementLinesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderStatementLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderStatementLinesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderStatementLinesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderStatementLinesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderStatementLinesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderStatementLinesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderStatementLinesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderStatementLinesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderStatementLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderStatementLinesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderStatementLinesFilter(filter *model.Filter) {
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
			Field: string(ProviderStatementLinesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderStatementLinesSelectableResponse map[string]interface{}
type ProviderStatementLinesSelectableListResponse []*ProviderStatementLinesSelectableResponse

func assignProviderStatementLinesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderStatementLinesSelectableValue(out ProviderStatementLinesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderStatementLinesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderStatementLinesSelectableResponse(providerStatementLines model.ProviderStatementLines, filter model.Filter) ProviderStatementLinesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderStatementLinesDBFieldName.Id),
			string(model.ProviderStatementLinesDBFieldName.StatementFileId),
			string(model.ProviderStatementLinesDBFieldName.LineNo),
			string(model.ProviderStatementLinesDBFieldName.ProviderRef),
			string(model.ProviderStatementLinesDBFieldName.TransactionType),
			string(model.ProviderStatementLinesDBFieldName.BookingDate),
			string(model.ProviderStatementLinesDBFieldName.ValueDate),
			string(model.ProviderStatementLinesDBFieldName.CurrencyCode),
			string(model.ProviderStatementLinesDBFieldName.GrossAmount),
			string(model.ProviderStatementLinesDBFieldName.FeeAmount),
			string(model.ProviderStatementLinesDBFieldName.NetAmount),
			string(model.ProviderStatementLinesDBFieldName.RawLine),
			string(model.ProviderStatementLinesDBFieldName.LineHash),
			string(model.ProviderStatementLinesDBFieldName.Metadata),
			string(model.ProviderStatementLinesDBFieldName.MetaCreatedAt),
			string(model.ProviderStatementLinesDBFieldName.MetaCreatedBy),
			string(model.ProviderStatementLinesDBFieldName.MetaUpdatedAt),
			string(model.ProviderStatementLinesDBFieldName.MetaUpdatedBy),
			string(model.ProviderStatementLinesDBFieldName.MetaDeletedAt),
			string(model.ProviderStatementLinesDBFieldName.MetaDeletedBy),
		)
	}
	providerStatementLinesSelectableResponse := ProviderStatementLinesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderStatementLinesDBFieldName.Id):
			key := string(ProviderStatementLinesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.Id, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.StatementFileId):
			key := string(ProviderStatementLinesDTOFieldName.StatementFileId)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.StatementFileId, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.LineNo):
			key := string(ProviderStatementLinesDTOFieldName.LineNo)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.LineNo, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.ProviderRef):
			key := string(ProviderStatementLinesDTOFieldName.ProviderRef)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.ProviderRef.String, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.TransactionType):
			key := string(ProviderStatementLinesDTOFieldName.TransactionType)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.TransactionType, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.BookingDate):
			key := string(ProviderStatementLinesDTOFieldName.BookingDate)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.BookingDate.Time, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.ValueDate):
			key := string(ProviderStatementLinesDTOFieldName.ValueDate)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.ValueDate.Time, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.CurrencyCode):
			key := string(ProviderStatementLinesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.CurrencyCode, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.GrossAmount):
			key := string(ProviderStatementLinesDTOFieldName.GrossAmount)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.GrossAmount, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.FeeAmount):
			key := string(ProviderStatementLinesDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.FeeAmount, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.NetAmount):
			key := string(ProviderStatementLinesDTOFieldName.NetAmount)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.NetAmount, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.RawLine):
			key := string(ProviderStatementLinesDTOFieldName.RawLine)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.RawLine, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.LineHash):
			key := string(ProviderStatementLinesDTOFieldName.LineHash)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.LineHash.String, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.Metadata):
			key := string(ProviderStatementLinesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.Metadata, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaCreatedAt):
			key := string(ProviderStatementLinesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaCreatedAt, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaCreatedBy):
			key := string(ProviderStatementLinesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaCreatedBy, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaUpdatedAt):
			key := string(ProviderStatementLinesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaUpdatedBy):
			key := string(ProviderStatementLinesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaUpdatedBy, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaDeletedAt):
			key := string(ProviderStatementLinesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderStatementLinesDBFieldName.MetaDeletedBy):
			key := string(ProviderStatementLinesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementLinesSelectableValue(providerStatementLinesSelectableResponse, key, providerStatementLines.MetaDeletedBy, explicitAlias)

		}
	}
	return providerStatementLinesSelectableResponse
}

func NewProviderStatementLinesListResponseFromFilterResult(result []model.ProviderStatementLinesFilterResult, filter model.Filter) ProviderStatementLinesSelectableListResponse {
	dtoProviderStatementLinesListResponse := ProviderStatementLinesSelectableListResponse{}
	for _, row := range result {
		dtoProviderStatementLinesResponse := NewProviderStatementLinesSelectableResponse(row.ProviderStatementLines, filter)
		dtoProviderStatementLinesListResponse = append(dtoProviderStatementLinesListResponse, &dtoProviderStatementLinesResponse)
	}
	return dtoProviderStatementLinesListResponse
}

type ProviderStatementLinesFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     ProviderStatementLinesSelectableListResponse `json:"data"`
}

func reverseProviderStatementLinesFilterResults(result []model.ProviderStatementLinesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderStatementLinesFilterResponse(result []model.ProviderStatementLinesFilterResult, filter model.Filter) (resp ProviderStatementLinesFilterResponse) {
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
			reverseProviderStatementLinesFilterResults(dataResult)
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

	resp.Data = NewProviderStatementLinesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderStatementLinesCreateRequest struct {
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	ProviderRef     string          `json:"providerRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	GrossAmount     decimal.Decimal `json:"grossAmount"`
	FeeAmount       decimal.Decimal `json:"feeAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *ProviderStatementLinesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderStatementLinesCreateRequest) ToModel() model.ProviderStatementLines {
	id, _ := uuid.NewV4()
	return model.ProviderStatementLines{
		Id:              id,
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		ProviderRef:     null.StringFrom(d.ProviderRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		GrossAmount:     d.GrossAmount,
		FeeAmount:       d.FeeAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type ProviderStatementLinesListCreateRequest []*ProviderStatementLinesCreateRequest

func (d ProviderStatementLinesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementLines := range d {
		err = validator.Struct(providerStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderStatementLinesListCreateRequest) ToModelList() []model.ProviderStatementLines {
	out := make([]model.ProviderStatementLines, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderStatementLinesUpdateRequest struct {
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	ProviderRef     string          `json:"providerRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	GrossAmount     decimal.Decimal `json:"grossAmount"`
	FeeAmount       decimal.Decimal `json:"feeAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *ProviderStatementLinesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderStatementLinesUpdateRequest) ToModel() model.ProviderStatementLines {
	return model.ProviderStatementLines{
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		ProviderRef:     null.StringFrom(d.ProviderRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		GrossAmount:     d.GrossAmount,
		FeeAmount:       d.FeeAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type ProviderStatementLinesBulkUpdateRequest struct {
	Id              uuid.UUID       `json:"id"`
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	ProviderRef     string          `json:"providerRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	GrossAmount     decimal.Decimal `json:"grossAmount"`
	FeeAmount       decimal.Decimal `json:"feeAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d ProviderStatementLinesBulkUpdateRequest) PrimaryID() ProviderStatementLinesPrimaryID {
	return ProviderStatementLinesPrimaryID{
		Id: d.Id,
	}
}

type ProviderStatementLinesListBulkUpdateRequest []*ProviderStatementLinesBulkUpdateRequest

func (d ProviderStatementLinesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementLines := range d {
		err = validator.Struct(providerStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderStatementLinesBulkUpdateRequest) ToModel() model.ProviderStatementLines {
	return model.ProviderStatementLines{
		Id:              d.Id,
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		ProviderRef:     null.StringFrom(d.ProviderRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		GrossAmount:     d.GrossAmount,
		FeeAmount:       d.FeeAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type ProviderStatementLinesResponse struct {
	Id              uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	StatementFileId uuid.UUID       `json:"statementFileId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LineNo          int             `json:"lineNo" validate:"required" example:"1"`
	ProviderRef     string          `json:"providerRef"`
	TransactionType string          `json:"transactionType" validate:"required"`
	BookingDate     time.Time       `json:"bookingDate" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ValueDate       time.Time       `json:"valueDate" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CurrencyCode    string          `json:"currencyCode" validate:"required"`
	GrossAmount     decimal.Decimal `json:"grossAmount" format:"decimal" example:"100.50"`
	FeeAmount       decimal.Decimal `json:"feeAmount" format:"decimal" example:"100.50"`
	NetAmount       decimal.Decimal `json:"netAmount" format:"decimal" example:"100.50"`
	RawLine         json.RawMessage `json:"rawLine" swaggertype:"object"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewProviderStatementLinesResponse(providerStatementLines model.ProviderStatementLines) ProviderStatementLinesResponse {
	return ProviderStatementLinesResponse{
		Id:              providerStatementLines.Id,
		StatementFileId: providerStatementLines.StatementFileId,
		LineNo:          providerStatementLines.LineNo,
		ProviderRef:     providerStatementLines.ProviderRef.String,
		TransactionType: providerStatementLines.TransactionType,
		BookingDate:     providerStatementLines.BookingDate.Time,
		ValueDate:       providerStatementLines.ValueDate.Time,
		CurrencyCode:    providerStatementLines.CurrencyCode,
		GrossAmount:     providerStatementLines.GrossAmount,
		FeeAmount:       providerStatementLines.FeeAmount,
		NetAmount:       providerStatementLines.NetAmount,
		RawLine:         providerStatementLines.RawLine,
		LineHash:        providerStatementLines.LineHash.String,
		Metadata:        providerStatementLines.Metadata,
	}
}

type ProviderStatementLinesListResponse []*ProviderStatementLinesResponse

func NewProviderStatementLinesListResponse(providerStatementLinesList model.ProviderStatementLinesList) ProviderStatementLinesListResponse {
	dtoProviderStatementLinesListResponse := ProviderStatementLinesListResponse{}
	for _, providerStatementLines := range providerStatementLinesList {
		dtoProviderStatementLinesResponse := NewProviderStatementLinesResponse(*providerStatementLines)
		dtoProviderStatementLinesListResponse = append(dtoProviderStatementLinesListResponse, &dtoProviderStatementLinesResponse)
	}
	return dtoProviderStatementLinesListResponse
}

type ProviderStatementLinesPrimaryIDList []ProviderStatementLinesPrimaryID

func (d ProviderStatementLinesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementLines := range d {
		err = validator.Struct(providerStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderStatementLinesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderStatementLinesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderStatementLinesPrimaryID) ToModel() model.ProviderStatementLinesPrimaryID {
	return model.ProviderStatementLinesPrimaryID{
		Id: d.Id,
	}
}

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

type BankStatementLinesDTOFieldNameType string

type bankStatementLinesDTOFieldName struct {
	Id              BankStatementLinesDTOFieldNameType
	StatementFileId BankStatementLinesDTOFieldNameType
	LineNo          BankStatementLinesDTOFieldNameType
	BankRef         BankStatementLinesDTOFieldNameType
	TransactionType BankStatementLinesDTOFieldNameType
	BookingDate     BankStatementLinesDTOFieldNameType
	ValueDate       BankStatementLinesDTOFieldNameType
	CurrencyCode    BankStatementLinesDTOFieldNameType
	DebitAmount     BankStatementLinesDTOFieldNameType
	CreditAmount    BankStatementLinesDTOFieldNameType
	NetAmount       BankStatementLinesDTOFieldNameType
	RawLine         BankStatementLinesDTOFieldNameType
	LineHash        BankStatementLinesDTOFieldNameType
	Metadata        BankStatementLinesDTOFieldNameType
	MetaCreatedAt   BankStatementLinesDTOFieldNameType
	MetaCreatedBy   BankStatementLinesDTOFieldNameType
	MetaUpdatedAt   BankStatementLinesDTOFieldNameType
	MetaUpdatedBy   BankStatementLinesDTOFieldNameType
	MetaDeletedAt   BankStatementLinesDTOFieldNameType
	MetaDeletedBy   BankStatementLinesDTOFieldNameType
}

var BankStatementLinesDTOFieldName = bankStatementLinesDTOFieldName{
	Id:              "id",
	StatementFileId: "statementFileId",
	LineNo:          "lineNo",
	BankRef:         "bankRef",
	TransactionType: "transactionType",
	BookingDate:     "bookingDate",
	ValueDate:       "valueDate",
	CurrencyCode:    "currencyCode",
	DebitAmount:     "debitAmount",
	CreditAmount:    "creditAmount",
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

func transformBankStatementLinesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(BankStatementLinesDTOFieldName.Id):
		return string(model.BankStatementLinesDBFieldName.Id), true

	case string(BankStatementLinesDTOFieldName.StatementFileId):
		return string(model.BankStatementLinesDBFieldName.StatementFileId), true

	case string(BankStatementLinesDTOFieldName.LineNo):
		return string(model.BankStatementLinesDBFieldName.LineNo), true

	case string(BankStatementLinesDTOFieldName.BankRef):
		return string(model.BankStatementLinesDBFieldName.BankRef), true

	case string(BankStatementLinesDTOFieldName.TransactionType):
		return string(model.BankStatementLinesDBFieldName.TransactionType), true

	case string(BankStatementLinesDTOFieldName.BookingDate):
		return string(model.BankStatementLinesDBFieldName.BookingDate), true

	case string(BankStatementLinesDTOFieldName.ValueDate):
		return string(model.BankStatementLinesDBFieldName.ValueDate), true

	case string(BankStatementLinesDTOFieldName.CurrencyCode):
		return string(model.BankStatementLinesDBFieldName.CurrencyCode), true

	case string(BankStatementLinesDTOFieldName.DebitAmount):
		return string(model.BankStatementLinesDBFieldName.DebitAmount), true

	case string(BankStatementLinesDTOFieldName.CreditAmount):
		return string(model.BankStatementLinesDBFieldName.CreditAmount), true

	case string(BankStatementLinesDTOFieldName.NetAmount):
		return string(model.BankStatementLinesDBFieldName.NetAmount), true

	case string(BankStatementLinesDTOFieldName.RawLine):
		return string(model.BankStatementLinesDBFieldName.RawLine), true

	case string(BankStatementLinesDTOFieldName.LineHash):
		return string(model.BankStatementLinesDBFieldName.LineHash), true

	case string(BankStatementLinesDTOFieldName.Metadata):
		return string(model.BankStatementLinesDBFieldName.Metadata), true

	case string(BankStatementLinesDTOFieldName.MetaCreatedAt):
		return string(model.BankStatementLinesDBFieldName.MetaCreatedAt), true

	case string(BankStatementLinesDTOFieldName.MetaCreatedBy):
		return string(model.BankStatementLinesDBFieldName.MetaCreatedBy), true

	case string(BankStatementLinesDTOFieldName.MetaUpdatedAt):
		return string(model.BankStatementLinesDBFieldName.MetaUpdatedAt), true

	case string(BankStatementLinesDTOFieldName.MetaUpdatedBy):
		return string(model.BankStatementLinesDBFieldName.MetaUpdatedBy), true

	case string(BankStatementLinesDTOFieldName.MetaDeletedAt):
		return string(model.BankStatementLinesDBFieldName.MetaDeletedAt), true

	case string(BankStatementLinesDTOFieldName.MetaDeletedBy):
		return string(model.BankStatementLinesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewBankStatementLinesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isBankStatementLinesBaseFilterField(field string) bool {
	spec, found := model.NewBankStatementLinesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeBankStatementLinesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateBankStatementLinesProjectionOutputPath(path string) error {
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

func transformBankStatementLinesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformBankStatementLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformBankStatementLinesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformBankStatementLinesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformBankStatementLinesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isBankStatementLinesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateBankStatementLinesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeBankStatementLinesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformBankStatementLinesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformBankStatementLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformBankStatementLinesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultBankStatementLinesFilter(filter *model.Filter) {
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
			Field: string(BankStatementLinesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type BankStatementLinesSelectableResponse map[string]interface{}
type BankStatementLinesSelectableListResponse []*BankStatementLinesSelectableResponse

func assignBankStatementLinesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setBankStatementLinesSelectableValue(out BankStatementLinesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignBankStatementLinesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewBankStatementLinesSelectableResponse(bankStatementLines model.BankStatementLines, filter model.Filter) BankStatementLinesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.BankStatementLinesDBFieldName.Id),
			string(model.BankStatementLinesDBFieldName.StatementFileId),
			string(model.BankStatementLinesDBFieldName.LineNo),
			string(model.BankStatementLinesDBFieldName.BankRef),
			string(model.BankStatementLinesDBFieldName.TransactionType),
			string(model.BankStatementLinesDBFieldName.BookingDate),
			string(model.BankStatementLinesDBFieldName.ValueDate),
			string(model.BankStatementLinesDBFieldName.CurrencyCode),
			string(model.BankStatementLinesDBFieldName.DebitAmount),
			string(model.BankStatementLinesDBFieldName.CreditAmount),
			string(model.BankStatementLinesDBFieldName.NetAmount),
			string(model.BankStatementLinesDBFieldName.RawLine),
			string(model.BankStatementLinesDBFieldName.LineHash),
			string(model.BankStatementLinesDBFieldName.Metadata),
			string(model.BankStatementLinesDBFieldName.MetaCreatedAt),
			string(model.BankStatementLinesDBFieldName.MetaCreatedBy),
			string(model.BankStatementLinesDBFieldName.MetaUpdatedAt),
			string(model.BankStatementLinesDBFieldName.MetaUpdatedBy),
			string(model.BankStatementLinesDBFieldName.MetaDeletedAt),
			string(model.BankStatementLinesDBFieldName.MetaDeletedBy),
		)
	}
	bankStatementLinesSelectableResponse := BankStatementLinesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.BankStatementLinesDBFieldName.Id):
			key := string(BankStatementLinesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.Id, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.StatementFileId):
			key := string(BankStatementLinesDTOFieldName.StatementFileId)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.StatementFileId, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.LineNo):
			key := string(BankStatementLinesDTOFieldName.LineNo)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.LineNo, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.BankRef):
			key := string(BankStatementLinesDTOFieldName.BankRef)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.BankRef.String, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.TransactionType):
			key := string(BankStatementLinesDTOFieldName.TransactionType)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.TransactionType, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.BookingDate):
			key := string(BankStatementLinesDTOFieldName.BookingDate)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.BookingDate.Time, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.ValueDate):
			key := string(BankStatementLinesDTOFieldName.ValueDate)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.ValueDate.Time, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.CurrencyCode):
			key := string(BankStatementLinesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.CurrencyCode, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.DebitAmount):
			key := string(BankStatementLinesDTOFieldName.DebitAmount)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.DebitAmount, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.CreditAmount):
			key := string(BankStatementLinesDTOFieldName.CreditAmount)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.CreditAmount, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.NetAmount):
			key := string(BankStatementLinesDTOFieldName.NetAmount)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.NetAmount, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.RawLine):
			key := string(BankStatementLinesDTOFieldName.RawLine)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.RawLine, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.LineHash):
			key := string(BankStatementLinesDTOFieldName.LineHash)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.LineHash.String, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.Metadata):
			key := string(BankStatementLinesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.Metadata, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaCreatedAt):
			key := string(BankStatementLinesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaCreatedAt, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaCreatedBy):
			key := string(BankStatementLinesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaCreatedBy, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaUpdatedAt):
			key := string(BankStatementLinesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaUpdatedAt, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaUpdatedBy):
			key := string(BankStatementLinesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaUpdatedBy, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaDeletedAt):
			key := string(BankStatementLinesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaDeletedAt.Time, explicitAlias)

		case string(model.BankStatementLinesDBFieldName.MetaDeletedBy):
			key := string(BankStatementLinesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementLinesSelectableValue(bankStatementLinesSelectableResponse, key, bankStatementLines.MetaDeletedBy, explicitAlias)

		}
	}
	return bankStatementLinesSelectableResponse
}

func NewBankStatementLinesListResponseFromFilterResult(result []model.BankStatementLinesFilterResult, filter model.Filter) BankStatementLinesSelectableListResponse {
	dtoBankStatementLinesListResponse := BankStatementLinesSelectableListResponse{}
	for _, row := range result {
		dtoBankStatementLinesResponse := NewBankStatementLinesSelectableResponse(row.BankStatementLines, filter)
		dtoBankStatementLinesListResponse = append(dtoBankStatementLinesListResponse, &dtoBankStatementLinesResponse)
	}
	return dtoBankStatementLinesListResponse
}

type BankStatementLinesFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     BankStatementLinesSelectableListResponse `json:"data"`
}

func reverseBankStatementLinesFilterResults(result []model.BankStatementLinesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewBankStatementLinesFilterResponse(result []model.BankStatementLinesFilterResult, filter model.Filter) (resp BankStatementLinesFilterResponse) {
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
			reverseBankStatementLinesFilterResults(dataResult)
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

	resp.Data = NewBankStatementLinesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type BankStatementLinesCreateRequest struct {
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	BankRef         string          `json:"bankRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	DebitAmount     decimal.Decimal `json:"debitAmount"`
	CreditAmount    decimal.Decimal `json:"creditAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *BankStatementLinesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *BankStatementLinesCreateRequest) ToModel() model.BankStatementLines {
	id, _ := uuid.NewV4()
	return model.BankStatementLines{
		Id:              id,
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		BankRef:         null.StringFrom(d.BankRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		DebitAmount:     d.DebitAmount,
		CreditAmount:    d.CreditAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type BankStatementLinesListCreateRequest []*BankStatementLinesCreateRequest

func (d BankStatementLinesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementLines := range d {
		err = validator.Struct(bankStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d BankStatementLinesListCreateRequest) ToModelList() []model.BankStatementLines {
	out := make([]model.BankStatementLines, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type BankStatementLinesUpdateRequest struct {
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	BankRef         string          `json:"bankRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	DebitAmount     decimal.Decimal `json:"debitAmount"`
	CreditAmount    decimal.Decimal `json:"creditAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d *BankStatementLinesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d BankStatementLinesUpdateRequest) ToModel() model.BankStatementLines {
	return model.BankStatementLines{
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		BankRef:         null.StringFrom(d.BankRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		DebitAmount:     d.DebitAmount,
		CreditAmount:    d.CreditAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type BankStatementLinesBulkUpdateRequest struct {
	Id              uuid.UUID       `json:"id"`
	StatementFileId uuid.UUID       `json:"statementFileId"`
	LineNo          int             `json:"lineNo"`
	BankRef         string          `json:"bankRef"`
	TransactionType string          `json:"transactionType"`
	BookingDate     time.Time       `json:"bookingDate"`
	ValueDate       time.Time       `json:"valueDate"`
	CurrencyCode    string          `json:"currencyCode"`
	DebitAmount     decimal.Decimal `json:"debitAmount"`
	CreditAmount    decimal.Decimal `json:"creditAmount"`
	NetAmount       decimal.Decimal `json:"netAmount"`
	RawLine         json.RawMessage `json:"rawLine"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata"`
}

func (d BankStatementLinesBulkUpdateRequest) PrimaryID() BankStatementLinesPrimaryID {
	return BankStatementLinesPrimaryID{
		Id: d.Id,
	}
}

type BankStatementLinesListBulkUpdateRequest []*BankStatementLinesBulkUpdateRequest

func (d BankStatementLinesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementLines := range d {
		err = validator.Struct(bankStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d BankStatementLinesBulkUpdateRequest) ToModel() model.BankStatementLines {
	return model.BankStatementLines{
		Id:              d.Id,
		StatementFileId: d.StatementFileId,
		LineNo:          d.LineNo,
		BankRef:         null.StringFrom(d.BankRef),
		TransactionType: d.TransactionType,
		BookingDate:     null.TimeFrom(d.BookingDate),
		ValueDate:       null.TimeFrom(d.ValueDate),
		CurrencyCode:    d.CurrencyCode,
		DebitAmount:     d.DebitAmount,
		CreditAmount:    d.CreditAmount,
		NetAmount:       d.NetAmount,
		RawLine:         d.RawLine,
		LineHash:        null.StringFrom(d.LineHash),
		Metadata:        d.Metadata,
	}
}

type BankStatementLinesResponse struct {
	Id              uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	StatementFileId uuid.UUID       `json:"statementFileId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LineNo          int             `json:"lineNo" validate:"required" example:"1"`
	BankRef         string          `json:"bankRef"`
	TransactionType string          `json:"transactionType" validate:"required"`
	BookingDate     time.Time       `json:"bookingDate" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ValueDate       time.Time       `json:"valueDate" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CurrencyCode    string          `json:"currencyCode" validate:"required"`
	DebitAmount     decimal.Decimal `json:"debitAmount" format:"decimal" example:"100.50"`
	CreditAmount    decimal.Decimal `json:"creditAmount" format:"decimal" example:"100.50"`
	NetAmount       decimal.Decimal `json:"netAmount" format:"decimal" example:"100.50"`
	RawLine         json.RawMessage `json:"rawLine" swaggertype:"object"`
	LineHash        string          `json:"lineHash"`
	Metadata        json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewBankStatementLinesResponse(bankStatementLines model.BankStatementLines) BankStatementLinesResponse {
	return BankStatementLinesResponse{
		Id:              bankStatementLines.Id,
		StatementFileId: bankStatementLines.StatementFileId,
		LineNo:          bankStatementLines.LineNo,
		BankRef:         bankStatementLines.BankRef.String,
		TransactionType: bankStatementLines.TransactionType,
		BookingDate:     bankStatementLines.BookingDate.Time,
		ValueDate:       bankStatementLines.ValueDate.Time,
		CurrencyCode:    bankStatementLines.CurrencyCode,
		DebitAmount:     bankStatementLines.DebitAmount,
		CreditAmount:    bankStatementLines.CreditAmount,
		NetAmount:       bankStatementLines.NetAmount,
		RawLine:         bankStatementLines.RawLine,
		LineHash:        bankStatementLines.LineHash.String,
		Metadata:        bankStatementLines.Metadata,
	}
}

type BankStatementLinesListResponse []*BankStatementLinesResponse

func NewBankStatementLinesListResponse(bankStatementLinesList model.BankStatementLinesList) BankStatementLinesListResponse {
	dtoBankStatementLinesListResponse := BankStatementLinesListResponse{}
	for _, bankStatementLines := range bankStatementLinesList {
		dtoBankStatementLinesResponse := NewBankStatementLinesResponse(*bankStatementLines)
		dtoBankStatementLinesListResponse = append(dtoBankStatementLinesListResponse, &dtoBankStatementLinesResponse)
	}
	return dtoBankStatementLinesListResponse
}

type BankStatementLinesPrimaryIDList []BankStatementLinesPrimaryID

func (d BankStatementLinesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementLines := range d {
		err = validator.Struct(bankStatementLines)
		if err != nil {
			return
		}
	}
	return nil
}

type BankStatementLinesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *BankStatementLinesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d BankStatementLinesPrimaryID) ToModel() model.BankStatementLinesPrimaryID {
	return model.BankStatementLinesPrimaryID{
		Id: d.Id,
	}
}

package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerJournalLinesDTOFieldNameType string

type ledgerJournalLinesDTOFieldName struct {
	Id               LedgerJournalLinesDTOFieldNameType
	JournalEntryId   LedgerJournalLinesDTOFieldNameType
	LineNo           LedgerJournalLinesDTOFieldNameType
	AccountId        LedgerJournalLinesDTOFieldNameType
	LineSide         LedgerJournalLinesDTOFieldNameType
	Amount           LedgerJournalLinesDTOFieldNameType
	CurrencyCode     LedgerJournalLinesDTOFieldNameType
	FxRateLockId     LedgerJournalLinesDTOFieldNameType
	AmountReporting  LedgerJournalLinesDTOFieldNameType
	ReferencePartyId LedgerJournalLinesDTOFieldNameType
	Dimensions       LedgerJournalLinesDTOFieldNameType
	Metadata         LedgerJournalLinesDTOFieldNameType
	MetaCreatedAt    LedgerJournalLinesDTOFieldNameType
	MetaCreatedBy    LedgerJournalLinesDTOFieldNameType
	MetaUpdatedAt    LedgerJournalLinesDTOFieldNameType
	MetaUpdatedBy    LedgerJournalLinesDTOFieldNameType
	MetaDeletedAt    LedgerJournalLinesDTOFieldNameType
	MetaDeletedBy    LedgerJournalLinesDTOFieldNameType
}

var LedgerJournalLinesDTOFieldName = ledgerJournalLinesDTOFieldName{
	Id:               "id",
	JournalEntryId:   "journalEntryId",
	LineNo:           "lineNo",
	AccountId:        "accountId",
	LineSide:         "lineSide",
	Amount:           "amount",
	CurrencyCode:     "currencyCode",
	FxRateLockId:     "fxRateLockId",
	AmountReporting:  "amountReporting",
	ReferencePartyId: "referencePartyId",
	Dimensions:       "dimensions",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformLedgerJournalLinesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerJournalLinesDTOFieldName.Id):
		return string(model.LedgerJournalLinesDBFieldName.Id), true

	case string(LedgerJournalLinesDTOFieldName.JournalEntryId):
		return string(model.LedgerJournalLinesDBFieldName.JournalEntryId), true

	case string(LedgerJournalLinesDTOFieldName.LineNo):
		return string(model.LedgerJournalLinesDBFieldName.LineNo), true

	case string(LedgerJournalLinesDTOFieldName.AccountId):
		return string(model.LedgerJournalLinesDBFieldName.AccountId), true

	case string(LedgerJournalLinesDTOFieldName.LineSide):
		return string(model.LedgerJournalLinesDBFieldName.LineSide), true

	case string(LedgerJournalLinesDTOFieldName.Amount):
		return string(model.LedgerJournalLinesDBFieldName.Amount), true

	case string(LedgerJournalLinesDTOFieldName.CurrencyCode):
		return string(model.LedgerJournalLinesDBFieldName.CurrencyCode), true

	case string(LedgerJournalLinesDTOFieldName.FxRateLockId):
		return string(model.LedgerJournalLinesDBFieldName.FxRateLockId), true

	case string(LedgerJournalLinesDTOFieldName.AmountReporting):
		return string(model.LedgerJournalLinesDBFieldName.AmountReporting), true

	case string(LedgerJournalLinesDTOFieldName.ReferencePartyId):
		return string(model.LedgerJournalLinesDBFieldName.ReferencePartyId), true

	case string(LedgerJournalLinesDTOFieldName.Dimensions):
		return string(model.LedgerJournalLinesDBFieldName.Dimensions), true

	case string(LedgerJournalLinesDTOFieldName.Metadata):
		return string(model.LedgerJournalLinesDBFieldName.Metadata), true

	case string(LedgerJournalLinesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerJournalLinesDBFieldName.MetaCreatedAt), true

	case string(LedgerJournalLinesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerJournalLinesDBFieldName.MetaCreatedBy), true

	case string(LedgerJournalLinesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerJournalLinesDBFieldName.MetaUpdatedAt), true

	case string(LedgerJournalLinesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerJournalLinesDBFieldName.MetaUpdatedBy), true

	case string(LedgerJournalLinesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerJournalLinesDBFieldName.MetaDeletedAt), true

	case string(LedgerJournalLinesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerJournalLinesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerJournalLinesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerJournalLinesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerJournalLinesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerJournalLinesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerJournalLinesProjectionOutputPath(path string) error {
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

func transformLedgerJournalLinesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerJournalLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerJournalLinesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerJournalLinesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerJournalLinesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerJournalLinesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerJournalLinesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerJournalLinesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerJournalLinesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerJournalLinesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerJournalLinesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerJournalLinesFilter(filter *model.Filter) {
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
			Field: string(LedgerJournalLinesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerJournalLinesSelectableResponse map[string]interface{}
type LedgerJournalLinesSelectableListResponse []*LedgerJournalLinesSelectableResponse

func assignLedgerJournalLinesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerJournalLinesSelectableValue(out LedgerJournalLinesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerJournalLinesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerJournalLinesSelectableResponse(ledgerJournalLines model.LedgerJournalLines, filter model.Filter) LedgerJournalLinesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerJournalLinesDBFieldName.Id),
			string(model.LedgerJournalLinesDBFieldName.JournalEntryId),
			string(model.LedgerJournalLinesDBFieldName.LineNo),
			string(model.LedgerJournalLinesDBFieldName.AccountId),
			string(model.LedgerJournalLinesDBFieldName.LineSide),
			string(model.LedgerJournalLinesDBFieldName.Amount),
			string(model.LedgerJournalLinesDBFieldName.CurrencyCode),
			string(model.LedgerJournalLinesDBFieldName.FxRateLockId),
			string(model.LedgerJournalLinesDBFieldName.AmountReporting),
			string(model.LedgerJournalLinesDBFieldName.ReferencePartyId),
			string(model.LedgerJournalLinesDBFieldName.Dimensions),
			string(model.LedgerJournalLinesDBFieldName.Metadata),
			string(model.LedgerJournalLinesDBFieldName.MetaCreatedAt),
			string(model.LedgerJournalLinesDBFieldName.MetaCreatedBy),
			string(model.LedgerJournalLinesDBFieldName.MetaUpdatedAt),
			string(model.LedgerJournalLinesDBFieldName.MetaUpdatedBy),
			string(model.LedgerJournalLinesDBFieldName.MetaDeletedAt),
			string(model.LedgerJournalLinesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerJournalLinesSelectableResponse := LedgerJournalLinesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerJournalLinesDBFieldName.Id):
			key := string(LedgerJournalLinesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.Id, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.JournalEntryId):
			key := string(LedgerJournalLinesDTOFieldName.JournalEntryId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.JournalEntryId, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.LineNo):
			key := string(LedgerJournalLinesDTOFieldName.LineNo)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.LineNo, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.AccountId):
			key := string(LedgerJournalLinesDTOFieldName.AccountId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.AccountId, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.LineSide):
			key := string(LedgerJournalLinesDTOFieldName.LineSide)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, model.LineSide(ledgerJournalLines.LineSide), explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.Amount):
			key := string(LedgerJournalLinesDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.Amount, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.CurrencyCode):
			key := string(LedgerJournalLinesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.CurrencyCode, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.FxRateLockId):
			key := string(LedgerJournalLinesDTOFieldName.FxRateLockId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.FxRateLockId.UUID, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.AmountReporting):
			key := string(LedgerJournalLinesDTOFieldName.AmountReporting)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.AmountReporting.Decimal, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.ReferencePartyId):
			key := string(LedgerJournalLinesDTOFieldName.ReferencePartyId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.ReferencePartyId.UUID, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.Dimensions):
			key := string(LedgerJournalLinesDTOFieldName.Dimensions)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.Dimensions, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.Metadata):
			key := string(LedgerJournalLinesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.Metadata, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaCreatedAt):
			key := string(LedgerJournalLinesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaCreatedAt, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaCreatedBy):
			key := string(LedgerJournalLinesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaCreatedBy, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaUpdatedAt):
			key := string(LedgerJournalLinesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaUpdatedBy):
			key := string(LedgerJournalLinesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaDeletedAt):
			key := string(LedgerJournalLinesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerJournalLinesDBFieldName.MetaDeletedBy):
			key := string(LedgerJournalLinesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalLinesSelectableValue(ledgerJournalLinesSelectableResponse, key, ledgerJournalLines.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerJournalLinesSelectableResponse
}

func NewLedgerJournalLinesListResponseFromFilterResult(result []model.LedgerJournalLinesFilterResult, filter model.Filter) LedgerJournalLinesSelectableListResponse {
	dtoLedgerJournalLinesListResponse := LedgerJournalLinesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerJournalLinesResponse := NewLedgerJournalLinesSelectableResponse(row.LedgerJournalLines, filter)
		dtoLedgerJournalLinesListResponse = append(dtoLedgerJournalLinesListResponse, &dtoLedgerJournalLinesResponse)
	}
	return dtoLedgerJournalLinesListResponse
}

type LedgerJournalLinesFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     LedgerJournalLinesSelectableListResponse `json:"data"`
}

func reverseLedgerJournalLinesFilterResults(result []model.LedgerJournalLinesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerJournalLinesFilterResponse(result []model.LedgerJournalLinesFilterResult, filter model.Filter) (resp LedgerJournalLinesFilterResponse) {
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
			reverseLedgerJournalLinesFilterResults(dataResult)
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

	resp.Data = NewLedgerJournalLinesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerJournalLinesCreateRequest struct {
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	LineNo           int             `json:"lineNo"`
	AccountId        uuid.UUID       `json:"accountId"`
	LineSide         model.LineSide  `json:"lineSide" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	FxRateLockId     uuid.UUID       `json:"fxRateLockId"`
	AmountReporting  decimal.Decimal `json:"amountReporting"`
	ReferencePartyId uuid.UUID       `json:"referencePartyId"`
	Dimensions       json.RawMessage `json:"dimensions"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *LedgerJournalLinesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerJournalLinesCreateRequest) ToModel() model.LedgerJournalLines {
	id, _ := uuid.NewV4()
	return model.LedgerJournalLines{
		Id:               id,
		JournalEntryId:   d.JournalEntryId,
		LineNo:           d.LineNo,
		AccountId:        d.AccountId,
		LineSide:         d.LineSide,
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		FxRateLockId:     nuuid.From(d.FxRateLockId),
		AmountReporting:  decimal.NewNullDecimal(d.AmountReporting),
		ReferencePartyId: nuuid.From(d.ReferencePartyId),
		Dimensions:       d.Dimensions,
		Metadata:         d.Metadata,
	}
}

type LedgerJournalLinesListCreateRequest []*LedgerJournalLinesCreateRequest

func (d LedgerJournalLinesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalLines := range d {
		err = validator.Struct(ledgerJournalLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalLinesListCreateRequest) ToModelList() []model.LedgerJournalLines {
	out := make([]model.LedgerJournalLines, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerJournalLinesUpdateRequest struct {
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	LineNo           int             `json:"lineNo"`
	AccountId        uuid.UUID       `json:"accountId"`
	LineSide         model.LineSide  `json:"lineSide" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	FxRateLockId     uuid.UUID       `json:"fxRateLockId"`
	AmountReporting  decimal.Decimal `json:"amountReporting"`
	ReferencePartyId uuid.UUID       `json:"referencePartyId"`
	Dimensions       json.RawMessage `json:"dimensions"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *LedgerJournalLinesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerJournalLinesUpdateRequest) ToModel() model.LedgerJournalLines {
	return model.LedgerJournalLines{
		JournalEntryId:   d.JournalEntryId,
		LineNo:           d.LineNo,
		AccountId:        d.AccountId,
		LineSide:         d.LineSide,
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		FxRateLockId:     nuuid.From(d.FxRateLockId),
		AmountReporting:  decimal.NewNullDecimal(d.AmountReporting),
		ReferencePartyId: nuuid.From(d.ReferencePartyId),
		Dimensions:       d.Dimensions,
		Metadata:         d.Metadata,
	}
}

type LedgerJournalLinesBulkUpdateRequest struct {
	Id               uuid.UUID       `json:"id"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	LineNo           int             `json:"lineNo"`
	AccountId        uuid.UUID       `json:"accountId"`
	LineSide         model.LineSide  `json:"lineSide" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	FxRateLockId     uuid.UUID       `json:"fxRateLockId"`
	AmountReporting  decimal.Decimal `json:"amountReporting"`
	ReferencePartyId uuid.UUID       `json:"referencePartyId"`
	Dimensions       json.RawMessage `json:"dimensions"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d LedgerJournalLinesBulkUpdateRequest) PrimaryID() LedgerJournalLinesPrimaryID {
	return LedgerJournalLinesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalLinesListBulkUpdateRequest []*LedgerJournalLinesBulkUpdateRequest

func (d LedgerJournalLinesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalLines := range d {
		err = validator.Struct(ledgerJournalLines)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalLinesBulkUpdateRequest) ToModel() model.LedgerJournalLines {
	return model.LedgerJournalLines{
		Id:               d.Id,
		JournalEntryId:   d.JournalEntryId,
		LineNo:           d.LineNo,
		AccountId:        d.AccountId,
		LineSide:         d.LineSide,
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		FxRateLockId:     nuuid.From(d.FxRateLockId),
		AmountReporting:  decimal.NewNullDecimal(d.AmountReporting),
		ReferencePartyId: nuuid.From(d.ReferencePartyId),
		Dimensions:       d.Dimensions,
		Metadata:         d.Metadata,
	}
}

type LedgerJournalLinesResponse struct {
	Id               uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LineNo           int             `json:"lineNo" validate:"required" example:"1"`
	AccountId        uuid.UUID       `json:"accountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LineSide         model.LineSide  `json:"lineSide" validate:"required,oneof=debit credit" enums:"debit,credit"`
	Amount           decimal.Decimal `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	CurrencyCode     string          `json:"currencyCode" validate:"required"`
	FxRateLockId     uuid.UUID       `json:"fxRateLockId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AmountReporting  decimal.Decimal `json:"amountReporting" format:"decimal" example:"100.50"`
	ReferencePartyId uuid.UUID       `json:"referencePartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Dimensions       json.RawMessage `json:"dimensions" swaggertype:"object"`
	Metadata         json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewLedgerJournalLinesResponse(ledgerJournalLines model.LedgerJournalLines) LedgerJournalLinesResponse {
	return LedgerJournalLinesResponse{
		Id:               ledgerJournalLines.Id,
		JournalEntryId:   ledgerJournalLines.JournalEntryId,
		LineNo:           ledgerJournalLines.LineNo,
		AccountId:        ledgerJournalLines.AccountId,
		LineSide:         model.LineSide(ledgerJournalLines.LineSide),
		Amount:           ledgerJournalLines.Amount,
		CurrencyCode:     ledgerJournalLines.CurrencyCode,
		FxRateLockId:     ledgerJournalLines.FxRateLockId.UUID,
		AmountReporting:  ledgerJournalLines.AmountReporting.Decimal,
		ReferencePartyId: ledgerJournalLines.ReferencePartyId.UUID,
		Dimensions:       ledgerJournalLines.Dimensions,
		Metadata:         ledgerJournalLines.Metadata,
	}
}

type LedgerJournalLinesListResponse []*LedgerJournalLinesResponse

func NewLedgerJournalLinesListResponse(ledgerJournalLinesList model.LedgerJournalLinesList) LedgerJournalLinesListResponse {
	dtoLedgerJournalLinesListResponse := LedgerJournalLinesListResponse{}
	for _, ledgerJournalLines := range ledgerJournalLinesList {
		dtoLedgerJournalLinesResponse := NewLedgerJournalLinesResponse(*ledgerJournalLines)
		dtoLedgerJournalLinesListResponse = append(dtoLedgerJournalLinesListResponse, &dtoLedgerJournalLinesResponse)
	}
	return dtoLedgerJournalLinesListResponse
}

type LedgerJournalLinesPrimaryIDList []LedgerJournalLinesPrimaryID

func (d LedgerJournalLinesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalLines := range d {
		err = validator.Struct(ledgerJournalLines)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerJournalLinesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerJournalLinesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerJournalLinesPrimaryID) ToModel() model.LedgerJournalLinesPrimaryID {
	return model.LedgerJournalLinesPrimaryID{
		Id: d.Id,
	}
}

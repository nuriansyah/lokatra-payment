package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type AccountingPeriodsDTOFieldNameType string

type accountingPeriodsDTOFieldName struct {
	Id            AccountingPeriodsDTOFieldNameType
	BookId        AccountingPeriodsDTOFieldNameType
	PeriodCode    AccountingPeriodsDTOFieldNameType
	PeriodType    AccountingPeriodsDTOFieldNameType
	PeriodStart   AccountingPeriodsDTOFieldNameType
	PeriodEnd     AccountingPeriodsDTOFieldNameType
	PeriodStatus  AccountingPeriodsDTOFieldNameType
	ClosedAt      AccountingPeriodsDTOFieldNameType
	ClosedBy      AccountingPeriodsDTOFieldNameType
	LockReason    AccountingPeriodsDTOFieldNameType
	ClosingHash   AccountingPeriodsDTOFieldNameType
	Metadata      AccountingPeriodsDTOFieldNameType
	MetaCreatedAt AccountingPeriodsDTOFieldNameType
	MetaCreatedBy AccountingPeriodsDTOFieldNameType
	MetaUpdatedAt AccountingPeriodsDTOFieldNameType
	MetaUpdatedBy AccountingPeriodsDTOFieldNameType
	MetaDeletedAt AccountingPeriodsDTOFieldNameType
	MetaDeletedBy AccountingPeriodsDTOFieldNameType
}

var AccountingPeriodsDTOFieldName = accountingPeriodsDTOFieldName{
	Id:            "id",
	BookId:        "bookId",
	PeriodCode:    "periodCode",
	PeriodType:    "periodType",
	PeriodStart:   "periodStart",
	PeriodEnd:     "periodEnd",
	PeriodStatus:  "periodStatus",
	ClosedAt:      "closedAt",
	ClosedBy:      "closedBy",
	LockReason:    "lockReason",
	ClosingHash:   "closingHash",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformAccountingPeriodsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(AccountingPeriodsDTOFieldName.Id):
		return string(model.AccountingPeriodsDBFieldName.Id), true

	case string(AccountingPeriodsDTOFieldName.BookId):
		return string(model.AccountingPeriodsDBFieldName.BookId), true

	case string(AccountingPeriodsDTOFieldName.PeriodCode):
		return string(model.AccountingPeriodsDBFieldName.PeriodCode), true

	case string(AccountingPeriodsDTOFieldName.PeriodType):
		return string(model.AccountingPeriodsDBFieldName.PeriodType), true

	case string(AccountingPeriodsDTOFieldName.PeriodStart):
		return string(model.AccountingPeriodsDBFieldName.PeriodStart), true

	case string(AccountingPeriodsDTOFieldName.PeriodEnd):
		return string(model.AccountingPeriodsDBFieldName.PeriodEnd), true

	case string(AccountingPeriodsDTOFieldName.PeriodStatus):
		return string(model.AccountingPeriodsDBFieldName.PeriodStatus), true

	case string(AccountingPeriodsDTOFieldName.ClosedAt):
		return string(model.AccountingPeriodsDBFieldName.ClosedAt), true

	case string(AccountingPeriodsDTOFieldName.ClosedBy):
		return string(model.AccountingPeriodsDBFieldName.ClosedBy), true

	case string(AccountingPeriodsDTOFieldName.LockReason):
		return string(model.AccountingPeriodsDBFieldName.LockReason), true

	case string(AccountingPeriodsDTOFieldName.ClosingHash):
		return string(model.AccountingPeriodsDBFieldName.ClosingHash), true

	case string(AccountingPeriodsDTOFieldName.Metadata):
		return string(model.AccountingPeriodsDBFieldName.Metadata), true

	case string(AccountingPeriodsDTOFieldName.MetaCreatedAt):
		return string(model.AccountingPeriodsDBFieldName.MetaCreatedAt), true

	case string(AccountingPeriodsDTOFieldName.MetaCreatedBy):
		return string(model.AccountingPeriodsDBFieldName.MetaCreatedBy), true

	case string(AccountingPeriodsDTOFieldName.MetaUpdatedAt):
		return string(model.AccountingPeriodsDBFieldName.MetaUpdatedAt), true

	case string(AccountingPeriodsDTOFieldName.MetaUpdatedBy):
		return string(model.AccountingPeriodsDBFieldName.MetaUpdatedBy), true

	case string(AccountingPeriodsDTOFieldName.MetaDeletedAt):
		return string(model.AccountingPeriodsDBFieldName.MetaDeletedAt), true

	case string(AccountingPeriodsDTOFieldName.MetaDeletedBy):
		return string(model.AccountingPeriodsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewAccountingPeriodsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isAccountingPeriodsBaseFilterField(field string) bool {
	spec, found := model.NewAccountingPeriodsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeAccountingPeriodsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateAccountingPeriodsProjectionOutputPath(path string) error {
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

func transformAccountingPeriodsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformAccountingPeriodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformAccountingPeriodsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformAccountingPeriodsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformAccountingPeriodsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isAccountingPeriodsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateAccountingPeriodsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeAccountingPeriodsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformAccountingPeriodsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformAccountingPeriodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformAccountingPeriodsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultAccountingPeriodsFilter(filter *model.Filter) {
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
			Field: string(AccountingPeriodsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type AccountingPeriodsSelectableResponse map[string]interface{}
type AccountingPeriodsSelectableListResponse []*AccountingPeriodsSelectableResponse

func assignAccountingPeriodsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setAccountingPeriodsSelectableValue(out AccountingPeriodsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignAccountingPeriodsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewAccountingPeriodsSelectableResponse(accountingPeriods model.AccountingPeriods, filter model.Filter) AccountingPeriodsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.AccountingPeriodsDBFieldName.Id),
			string(model.AccountingPeriodsDBFieldName.BookId),
			string(model.AccountingPeriodsDBFieldName.PeriodCode),
			string(model.AccountingPeriodsDBFieldName.PeriodType),
			string(model.AccountingPeriodsDBFieldName.PeriodStart),
			string(model.AccountingPeriodsDBFieldName.PeriodEnd),
			string(model.AccountingPeriodsDBFieldName.PeriodStatus),
			string(model.AccountingPeriodsDBFieldName.ClosedAt),
			string(model.AccountingPeriodsDBFieldName.ClosedBy),
			string(model.AccountingPeriodsDBFieldName.LockReason),
			string(model.AccountingPeriodsDBFieldName.ClosingHash),
			string(model.AccountingPeriodsDBFieldName.Metadata),
			string(model.AccountingPeriodsDBFieldName.MetaCreatedAt),
			string(model.AccountingPeriodsDBFieldName.MetaCreatedBy),
			string(model.AccountingPeriodsDBFieldName.MetaUpdatedAt),
			string(model.AccountingPeriodsDBFieldName.MetaUpdatedBy),
			string(model.AccountingPeriodsDBFieldName.MetaDeletedAt),
			string(model.AccountingPeriodsDBFieldName.MetaDeletedBy),
		)
	}
	accountingPeriodsSelectableResponse := AccountingPeriodsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.AccountingPeriodsDBFieldName.Id):
			key := string(AccountingPeriodsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.Id, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.BookId):
			key := string(AccountingPeriodsDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.BookId, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.PeriodCode):
			key := string(AccountingPeriodsDTOFieldName.PeriodCode)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.PeriodCode, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.PeriodType):
			key := string(AccountingPeriodsDTOFieldName.PeriodType)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.PeriodType, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.PeriodStart):
			key := string(AccountingPeriodsDTOFieldName.PeriodStart)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.PeriodStart, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.PeriodEnd):
			key := string(AccountingPeriodsDTOFieldName.PeriodEnd)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.PeriodEnd, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.PeriodStatus):
			key := string(AccountingPeriodsDTOFieldName.PeriodStatus)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, model.PeriodStatus(accountingPeriods.PeriodStatus), explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.ClosedAt):
			key := string(AccountingPeriodsDTOFieldName.ClosedAt)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.ClosedAt.Time, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.ClosedBy):
			key := string(AccountingPeriodsDTOFieldName.ClosedBy)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.ClosedBy.UUID, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.LockReason):
			key := string(AccountingPeriodsDTOFieldName.LockReason)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.LockReason.String, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.ClosingHash):
			key := string(AccountingPeriodsDTOFieldName.ClosingHash)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.ClosingHash.String, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.Metadata):
			key := string(AccountingPeriodsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.Metadata, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaCreatedAt):
			key := string(AccountingPeriodsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaCreatedAt, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaCreatedBy):
			key := string(AccountingPeriodsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaCreatedBy, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaUpdatedAt):
			key := string(AccountingPeriodsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaUpdatedAt, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaUpdatedBy):
			key := string(AccountingPeriodsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaUpdatedBy, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaDeletedAt):
			key := string(AccountingPeriodsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaDeletedAt.Time, explicitAlias)

		case string(model.AccountingPeriodsDBFieldName.MetaDeletedBy):
			key := string(AccountingPeriodsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setAccountingPeriodsSelectableValue(accountingPeriodsSelectableResponse, key, accountingPeriods.MetaDeletedBy, explicitAlias)

		}
	}
	return accountingPeriodsSelectableResponse
}

func NewAccountingPeriodsListResponseFromFilterResult(result []model.AccountingPeriodsFilterResult, filter model.Filter) AccountingPeriodsSelectableListResponse {
	dtoAccountingPeriodsListResponse := AccountingPeriodsSelectableListResponse{}
	for _, row := range result {
		dtoAccountingPeriodsResponse := NewAccountingPeriodsSelectableResponse(row.AccountingPeriods, filter)
		dtoAccountingPeriodsListResponse = append(dtoAccountingPeriodsListResponse, &dtoAccountingPeriodsResponse)
	}
	return dtoAccountingPeriodsListResponse
}

type AccountingPeriodsFilterResponse struct {
	Metadata Metadata                                `json:"metadata"`
	Data     AccountingPeriodsSelectableListResponse `json:"data"`
}

func reverseAccountingPeriodsFilterResults(result []model.AccountingPeriodsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewAccountingPeriodsFilterResponse(result []model.AccountingPeriodsFilterResult, filter model.Filter) (resp AccountingPeriodsFilterResponse) {
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
			reverseAccountingPeriodsFilterResults(dataResult)
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

	resp.Data = NewAccountingPeriodsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type AccountingPeriodsCreateRequest struct {
	BookId       uuid.UUID          `json:"bookId"`
	PeriodCode   string             `json:"periodCode"`
	PeriodType   string             `json:"periodType"`
	PeriodStart  time.Time          `json:"periodStart"`
	PeriodEnd    time.Time          `json:"periodEnd"`
	PeriodStatus model.PeriodStatus `json:"periodStatus" example:"open" enums:"open,closed,locked"`
	ClosedAt     time.Time          `json:"closedAt"`
	ClosedBy     uuid.UUID          `json:"closedBy"`
	LockReason   string             `json:"lockReason"`
	ClosingHash  string             `json:"closingHash"`
	Metadata     json.RawMessage    `json:"metadata"`
}

func (d *AccountingPeriodsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *AccountingPeriodsCreateRequest) ToModel() model.AccountingPeriods {
	id, _ := uuid.NewV4()
	return model.AccountingPeriods{
		Id:           id,
		BookId:       d.BookId,
		PeriodCode:   d.PeriodCode,
		PeriodType:   d.PeriodType,
		PeriodStart:  d.PeriodStart,
		PeriodEnd:    d.PeriodEnd,
		PeriodStatus: d.PeriodStatus,
		ClosedAt:     null.TimeFrom(d.ClosedAt),
		ClosedBy:     nuuid.From(d.ClosedBy),
		LockReason:   null.StringFrom(d.LockReason),
		ClosingHash:  null.StringFrom(d.ClosingHash),
		Metadata:     d.Metadata,
	}
}

type AccountingPeriodsListCreateRequest []*AccountingPeriodsCreateRequest

func (d AccountingPeriodsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, accountingPeriods := range d {
		err = validator.Struct(accountingPeriods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d AccountingPeriodsListCreateRequest) ToModelList() []model.AccountingPeriods {
	out := make([]model.AccountingPeriods, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type AccountingPeriodsUpdateRequest struct {
	BookId       uuid.UUID          `json:"bookId"`
	PeriodCode   string             `json:"periodCode"`
	PeriodType   string             `json:"periodType"`
	PeriodStart  time.Time          `json:"periodStart"`
	PeriodEnd    time.Time          `json:"periodEnd"`
	PeriodStatus model.PeriodStatus `json:"periodStatus" example:"open" enums:"open,closed,locked"`
	ClosedAt     time.Time          `json:"closedAt"`
	ClosedBy     uuid.UUID          `json:"closedBy"`
	LockReason   string             `json:"lockReason"`
	ClosingHash  string             `json:"closingHash"`
	Metadata     json.RawMessage    `json:"metadata"`
}

func (d *AccountingPeriodsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d AccountingPeriodsUpdateRequest) ToModel() model.AccountingPeriods {
	return model.AccountingPeriods{
		BookId:       d.BookId,
		PeriodCode:   d.PeriodCode,
		PeriodType:   d.PeriodType,
		PeriodStart:  d.PeriodStart,
		PeriodEnd:    d.PeriodEnd,
		PeriodStatus: d.PeriodStatus,
		ClosedAt:     null.TimeFrom(d.ClosedAt),
		ClosedBy:     nuuid.From(d.ClosedBy),
		LockReason:   null.StringFrom(d.LockReason),
		ClosingHash:  null.StringFrom(d.ClosingHash),
		Metadata:     d.Metadata,
	}
}

type AccountingPeriodsBulkUpdateRequest struct {
	Id           uuid.UUID          `json:"id"`
	BookId       uuid.UUID          `json:"bookId"`
	PeriodCode   string             `json:"periodCode"`
	PeriodType   string             `json:"periodType"`
	PeriodStart  time.Time          `json:"periodStart"`
	PeriodEnd    time.Time          `json:"periodEnd"`
	PeriodStatus model.PeriodStatus `json:"periodStatus" example:"open" enums:"open,closed,locked"`
	ClosedAt     time.Time          `json:"closedAt"`
	ClosedBy     uuid.UUID          `json:"closedBy"`
	LockReason   string             `json:"lockReason"`
	ClosingHash  string             `json:"closingHash"`
	Metadata     json.RawMessage    `json:"metadata"`
}

func (d AccountingPeriodsBulkUpdateRequest) PrimaryID() AccountingPeriodsPrimaryID {
	return AccountingPeriodsPrimaryID{
		Id: d.Id,
	}
}

type AccountingPeriodsListBulkUpdateRequest []*AccountingPeriodsBulkUpdateRequest

func (d AccountingPeriodsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, accountingPeriods := range d {
		err = validator.Struct(accountingPeriods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d AccountingPeriodsBulkUpdateRequest) ToModel() model.AccountingPeriods {
	return model.AccountingPeriods{
		Id:           d.Id,
		BookId:       d.BookId,
		PeriodCode:   d.PeriodCode,
		PeriodType:   d.PeriodType,
		PeriodStart:  d.PeriodStart,
		PeriodEnd:    d.PeriodEnd,
		PeriodStatus: d.PeriodStatus,
		ClosedAt:     null.TimeFrom(d.ClosedAt),
		ClosedBy:     nuuid.From(d.ClosedBy),
		LockReason:   null.StringFrom(d.LockReason),
		ClosingHash:  null.StringFrom(d.ClosingHash),
		Metadata:     d.Metadata,
	}
}

type AccountingPeriodsResponse struct {
	Id           uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BookId       uuid.UUID          `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PeriodCode   string             `json:"periodCode" validate:"required"`
	PeriodType   string             `json:"periodType" validate:"required"`
	PeriodStart  time.Time          `json:"periodStart" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PeriodEnd    time.Time          `json:"periodEnd" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PeriodStatus model.PeriodStatus `json:"periodStatus" validate:"oneof=open closed locked" enums:"open,closed,locked"`
	ClosedAt     time.Time          `json:"closedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ClosedBy     uuid.UUID          `json:"closedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LockReason   string             `json:"lockReason"`
	ClosingHash  string             `json:"closingHash"`
	Metadata     json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewAccountingPeriodsResponse(accountingPeriods model.AccountingPeriods) AccountingPeriodsResponse {
	return AccountingPeriodsResponse{
		Id:           accountingPeriods.Id,
		BookId:       accountingPeriods.BookId,
		PeriodCode:   accountingPeriods.PeriodCode,
		PeriodType:   accountingPeriods.PeriodType,
		PeriodStart:  accountingPeriods.PeriodStart,
		PeriodEnd:    accountingPeriods.PeriodEnd,
		PeriodStatus: model.PeriodStatus(accountingPeriods.PeriodStatus),
		ClosedAt:     accountingPeriods.ClosedAt.Time,
		ClosedBy:     accountingPeriods.ClosedBy.UUID,
		LockReason:   accountingPeriods.LockReason.String,
		ClosingHash:  accountingPeriods.ClosingHash.String,
		Metadata:     accountingPeriods.Metadata,
	}
}

type AccountingPeriodsListResponse []*AccountingPeriodsResponse

func NewAccountingPeriodsListResponse(accountingPeriodsList model.AccountingPeriodsList) AccountingPeriodsListResponse {
	dtoAccountingPeriodsListResponse := AccountingPeriodsListResponse{}
	for _, accountingPeriods := range accountingPeriodsList {
		dtoAccountingPeriodsResponse := NewAccountingPeriodsResponse(*accountingPeriods)
		dtoAccountingPeriodsListResponse = append(dtoAccountingPeriodsListResponse, &dtoAccountingPeriodsResponse)
	}
	return dtoAccountingPeriodsListResponse
}

type AccountingPeriodsPrimaryIDList []AccountingPeriodsPrimaryID

func (d AccountingPeriodsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, accountingPeriods := range d {
		err = validator.Struct(accountingPeriods)
		if err != nil {
			return
		}
	}
	return nil
}

type AccountingPeriodsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *AccountingPeriodsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d AccountingPeriodsPrimaryID) ToModel() model.AccountingPeriodsPrimaryID {
	return model.AccountingPeriodsPrimaryID{
		Id: d.Id,
	}
}

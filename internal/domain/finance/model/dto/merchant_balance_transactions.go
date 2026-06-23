package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type MerchantBalanceTransactionsDTOFieldNameType string

type merchantBalanceTransactionsDTOFieldName struct {
	Id               MerchantBalanceTransactionsDTOFieldNameType
	BalanceAccountId MerchantBalanceTransactionsDTOFieldNameType
	MerchantPartyId  MerchantBalanceTransactionsDTOFieldNameType
	BalanceType      MerchantBalanceTransactionsDTOFieldNameType
	CurrencyCode     MerchantBalanceTransactionsDTOFieldNameType
	SourceType       MerchantBalanceTransactionsDTOFieldNameType
	SourceId         MerchantBalanceTransactionsDTOFieldNameType
	Direction        MerchantBalanceTransactionsDTOFieldNameType
	Amount           MerchantBalanceTransactionsDTOFieldNameType
	BalanceBefore    MerchantBalanceTransactionsDTOFieldNameType
	BalanceAfter     MerchantBalanceTransactionsDTOFieldNameType
	ReasonCode       MerchantBalanceTransactionsDTOFieldNameType
	Metadata         MerchantBalanceTransactionsDTOFieldNameType
	MetaCreatedAt    MerchantBalanceTransactionsDTOFieldNameType
	MetaCreatedBy    MerchantBalanceTransactionsDTOFieldNameType
	MetaUpdatedAt    MerchantBalanceTransactionsDTOFieldNameType
	MetaUpdatedBy    MerchantBalanceTransactionsDTOFieldNameType
	MetaDeletedAt    MerchantBalanceTransactionsDTOFieldNameType
	MetaDeletedBy    MerchantBalanceTransactionsDTOFieldNameType
}

var MerchantBalanceTransactionsDTOFieldName = merchantBalanceTransactionsDTOFieldName{
	Id:               "id",
	BalanceAccountId: "balanceAccountId",
	MerchantPartyId:  "merchantPartyId",
	BalanceType:      "balanceType",
	CurrencyCode:     "currencyCode",
	SourceType:       "sourceType",
	SourceId:         "sourceId",
	Direction:        "direction",
	Amount:           "amount",
	BalanceBefore:    "balanceBefore",
	BalanceAfter:     "balanceAfter",
	ReasonCode:       "reasonCode",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformMerchantBalanceTransactionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(MerchantBalanceTransactionsDTOFieldName.Id):
		return string(model.MerchantBalanceTransactionsDBFieldName.Id), true

	case string(MerchantBalanceTransactionsDTOFieldName.BalanceAccountId):
		return string(model.MerchantBalanceTransactionsDBFieldName.BalanceAccountId), true

	case string(MerchantBalanceTransactionsDTOFieldName.MerchantPartyId):
		return string(model.MerchantBalanceTransactionsDBFieldName.MerchantPartyId), true

	case string(MerchantBalanceTransactionsDTOFieldName.BalanceType):
		return string(model.MerchantBalanceTransactionsDBFieldName.BalanceType), true

	case string(MerchantBalanceTransactionsDTOFieldName.CurrencyCode):
		return string(model.MerchantBalanceTransactionsDBFieldName.CurrencyCode), true

	case string(MerchantBalanceTransactionsDTOFieldName.SourceType):
		return string(model.MerchantBalanceTransactionsDBFieldName.SourceType), true

	case string(MerchantBalanceTransactionsDTOFieldName.SourceId):
		return string(model.MerchantBalanceTransactionsDBFieldName.SourceId), true

	case string(MerchantBalanceTransactionsDTOFieldName.Direction):
		return string(model.MerchantBalanceTransactionsDBFieldName.Direction), true

	case string(MerchantBalanceTransactionsDTOFieldName.Amount):
		return string(model.MerchantBalanceTransactionsDBFieldName.Amount), true

	case string(MerchantBalanceTransactionsDTOFieldName.BalanceBefore):
		return string(model.MerchantBalanceTransactionsDBFieldName.BalanceBefore), true

	case string(MerchantBalanceTransactionsDTOFieldName.BalanceAfter):
		return string(model.MerchantBalanceTransactionsDBFieldName.BalanceAfter), true

	case string(MerchantBalanceTransactionsDTOFieldName.ReasonCode):
		return string(model.MerchantBalanceTransactionsDBFieldName.ReasonCode), true

	case string(MerchantBalanceTransactionsDTOFieldName.Metadata):
		return string(model.MerchantBalanceTransactionsDBFieldName.Metadata), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaCreatedAt):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedAt), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaCreatedBy):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedBy), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaUpdatedAt):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedAt), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaUpdatedBy):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedBy), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaDeletedAt):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedAt), true

	case string(MerchantBalanceTransactionsDTOFieldName.MetaDeletedBy):
		return string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewMerchantBalanceTransactionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isMerchantBalanceTransactionsBaseFilterField(field string) bool {
	spec, found := model.NewMerchantBalanceTransactionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeMerchantBalanceTransactionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateMerchantBalanceTransactionsProjectionOutputPath(path string) error {
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

func transformMerchantBalanceTransactionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformMerchantBalanceTransactionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformMerchantBalanceTransactionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformMerchantBalanceTransactionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformMerchantBalanceTransactionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isMerchantBalanceTransactionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateMerchantBalanceTransactionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeMerchantBalanceTransactionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformMerchantBalanceTransactionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformMerchantBalanceTransactionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformMerchantBalanceTransactionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultMerchantBalanceTransactionsFilter(filter *model.Filter) {
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
			Field: string(MerchantBalanceTransactionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type MerchantBalanceTransactionsSelectableResponse map[string]interface{}
type MerchantBalanceTransactionsSelectableListResponse []*MerchantBalanceTransactionsSelectableResponse

func assignMerchantBalanceTransactionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setMerchantBalanceTransactionsSelectableValue(out MerchantBalanceTransactionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignMerchantBalanceTransactionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewMerchantBalanceTransactionsSelectableResponse(merchantBalanceTransactions model.MerchantBalanceTransactions, filter model.Filter) MerchantBalanceTransactionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.MerchantBalanceTransactionsDBFieldName.Id),
			string(model.MerchantBalanceTransactionsDBFieldName.BalanceAccountId),
			string(model.MerchantBalanceTransactionsDBFieldName.MerchantPartyId),
			string(model.MerchantBalanceTransactionsDBFieldName.BalanceType),
			string(model.MerchantBalanceTransactionsDBFieldName.CurrencyCode),
			string(model.MerchantBalanceTransactionsDBFieldName.SourceType),
			string(model.MerchantBalanceTransactionsDBFieldName.SourceId),
			string(model.MerchantBalanceTransactionsDBFieldName.Direction),
			string(model.MerchantBalanceTransactionsDBFieldName.Amount),
			string(model.MerchantBalanceTransactionsDBFieldName.BalanceBefore),
			string(model.MerchantBalanceTransactionsDBFieldName.BalanceAfter),
			string(model.MerchantBalanceTransactionsDBFieldName.ReasonCode),
			string(model.MerchantBalanceTransactionsDBFieldName.Metadata),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedAt),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedBy),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedAt),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedBy),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedAt),
			string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedBy),
		)
	}
	merchantBalanceTransactionsSelectableResponse := MerchantBalanceTransactionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.MerchantBalanceTransactionsDBFieldName.Id):
			key := string(MerchantBalanceTransactionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.Id, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.BalanceAccountId):
			key := string(MerchantBalanceTransactionsDTOFieldName.BalanceAccountId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.BalanceAccountId, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MerchantPartyId):
			key := string(MerchantBalanceTransactionsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MerchantPartyId, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.BalanceType):
			key := string(MerchantBalanceTransactionsDTOFieldName.BalanceType)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, model.MerchantBalanceTransactionsBalanceType(merchantBalanceTransactions.BalanceType), explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.CurrencyCode):
			key := string(MerchantBalanceTransactionsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.CurrencyCode, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.SourceType):
			key := string(MerchantBalanceTransactionsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.SourceType, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.SourceId):
			key := string(MerchantBalanceTransactionsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.SourceId, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.Direction):
			key := string(MerchantBalanceTransactionsDTOFieldName.Direction)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, model.Direction(merchantBalanceTransactions.Direction), explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.Amount):
			key := string(MerchantBalanceTransactionsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.Amount, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.BalanceBefore):
			key := string(MerchantBalanceTransactionsDTOFieldName.BalanceBefore)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.BalanceBefore, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.BalanceAfter):
			key := string(MerchantBalanceTransactionsDTOFieldName.BalanceAfter)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.BalanceAfter, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.ReasonCode):
			key := string(MerchantBalanceTransactionsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.ReasonCode, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.Metadata):
			key := string(MerchantBalanceTransactionsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.Metadata, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedAt):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaCreatedAt, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaCreatedBy):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaCreatedBy, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedAt):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaUpdatedAt, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaUpdatedBy):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaUpdatedBy, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedAt):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaDeletedAt.Time, explicitAlias)

		case string(model.MerchantBalanceTransactionsDBFieldName.MetaDeletedBy):
			key := string(MerchantBalanceTransactionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceTransactionsSelectableValue(merchantBalanceTransactionsSelectableResponse, key, merchantBalanceTransactions.MetaDeletedBy, explicitAlias)

		}
	}
	return merchantBalanceTransactionsSelectableResponse
}

func NewMerchantBalanceTransactionsListResponseFromFilterResult(result []model.MerchantBalanceTransactionsFilterResult, filter model.Filter) MerchantBalanceTransactionsSelectableListResponse {
	dtoMerchantBalanceTransactionsListResponse := MerchantBalanceTransactionsSelectableListResponse{}
	for _, row := range result {
		dtoMerchantBalanceTransactionsResponse := NewMerchantBalanceTransactionsSelectableResponse(row.MerchantBalanceTransactions, filter)
		dtoMerchantBalanceTransactionsListResponse = append(dtoMerchantBalanceTransactionsListResponse, &dtoMerchantBalanceTransactionsResponse)
	}
	return dtoMerchantBalanceTransactionsListResponse
}

type MerchantBalanceTransactionsFilterResponse struct {
	Metadata Metadata                                          `json:"metadata"`
	Data     MerchantBalanceTransactionsSelectableListResponse `json:"data"`
}

func reverseMerchantBalanceTransactionsFilterResults(result []model.MerchantBalanceTransactionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewMerchantBalanceTransactionsFilterResponse(result []model.MerchantBalanceTransactionsFilterResult, filter model.Filter) (resp MerchantBalanceTransactionsFilterResponse) {
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
			reverseMerchantBalanceTransactionsFilterResults(dataResult)
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

	resp.Data = NewMerchantBalanceTransactionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type MerchantBalanceTransactionsCreateRequest struct {
	BalanceAccountId uuid.UUID                                    `json:"balanceAccountId"`
	MerchantPartyId  uuid.UUID                                    `json:"merchantPartyId"`
	BalanceType      model.MerchantBalanceTransactionsBalanceType `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode     string                                       `json:"currencyCode"`
	SourceType       string                                       `json:"sourceType"`
	SourceId         uuid.UUID                                    `json:"sourceId"`
	Direction        model.Direction                              `json:"direction" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal                              `json:"amount"`
	BalanceBefore    decimal.Decimal                              `json:"balanceBefore"`
	BalanceAfter     decimal.Decimal                              `json:"balanceAfter"`
	ReasonCode       string                                       `json:"reasonCode"`
	Metadata         json.RawMessage                              `json:"metadata"`
}

func (d *MerchantBalanceTransactionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *MerchantBalanceTransactionsCreateRequest) ToModel() model.MerchantBalanceTransactions {
	id, _ := uuid.NewV4()
	return model.MerchantBalanceTransactions{
		Id:               id,
		BalanceAccountId: d.BalanceAccountId,
		MerchantPartyId:  d.MerchantPartyId,
		BalanceType:      d.BalanceType,
		CurrencyCode:     d.CurrencyCode,
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		Direction:        d.Direction,
		Amount:           d.Amount,
		BalanceBefore:    d.BalanceBefore,
		BalanceAfter:     d.BalanceAfter,
		ReasonCode:       d.ReasonCode,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceTransactionsListCreateRequest []*MerchantBalanceTransactionsCreateRequest

func (d MerchantBalanceTransactionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceTransactions := range d {
		err = validator.Struct(merchantBalanceTransactions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceTransactionsListCreateRequest) ToModelList() []model.MerchantBalanceTransactions {
	out := make([]model.MerchantBalanceTransactions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type MerchantBalanceTransactionsUpdateRequest struct {
	BalanceAccountId uuid.UUID                                    `json:"balanceAccountId"`
	MerchantPartyId  uuid.UUID                                    `json:"merchantPartyId"`
	BalanceType      model.MerchantBalanceTransactionsBalanceType `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode     string                                       `json:"currencyCode"`
	SourceType       string                                       `json:"sourceType"`
	SourceId         uuid.UUID                                    `json:"sourceId"`
	Direction        model.Direction                              `json:"direction" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal                              `json:"amount"`
	BalanceBefore    decimal.Decimal                              `json:"balanceBefore"`
	BalanceAfter     decimal.Decimal                              `json:"balanceAfter"`
	ReasonCode       string                                       `json:"reasonCode"`
	Metadata         json.RawMessage                              `json:"metadata"`
}

func (d *MerchantBalanceTransactionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d MerchantBalanceTransactionsUpdateRequest) ToModel() model.MerchantBalanceTransactions {
	return model.MerchantBalanceTransactions{
		BalanceAccountId: d.BalanceAccountId,
		MerchantPartyId:  d.MerchantPartyId,
		BalanceType:      d.BalanceType,
		CurrencyCode:     d.CurrencyCode,
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		Direction:        d.Direction,
		Amount:           d.Amount,
		BalanceBefore:    d.BalanceBefore,
		BalanceAfter:     d.BalanceAfter,
		ReasonCode:       d.ReasonCode,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceTransactionsBulkUpdateRequest struct {
	Id               uuid.UUID                                    `json:"id"`
	BalanceAccountId uuid.UUID                                    `json:"balanceAccountId"`
	MerchantPartyId  uuid.UUID                                    `json:"merchantPartyId"`
	BalanceType      model.MerchantBalanceTransactionsBalanceType `json:"balanceType" example:"pending" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode     string                                       `json:"currencyCode"`
	SourceType       string                                       `json:"sourceType"`
	SourceId         uuid.UUID                                    `json:"sourceId"`
	Direction        model.Direction                              `json:"direction" example:"debit" enums:"debit,credit"`
	Amount           decimal.Decimal                              `json:"amount"`
	BalanceBefore    decimal.Decimal                              `json:"balanceBefore"`
	BalanceAfter     decimal.Decimal                              `json:"balanceAfter"`
	ReasonCode       string                                       `json:"reasonCode"`
	Metadata         json.RawMessage                              `json:"metadata"`
}

func (d MerchantBalanceTransactionsBulkUpdateRequest) PrimaryID() MerchantBalanceTransactionsPrimaryID {
	return MerchantBalanceTransactionsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceTransactionsListBulkUpdateRequest []*MerchantBalanceTransactionsBulkUpdateRequest

func (d MerchantBalanceTransactionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceTransactions := range d {
		err = validator.Struct(merchantBalanceTransactions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceTransactionsBulkUpdateRequest) ToModel() model.MerchantBalanceTransactions {
	return model.MerchantBalanceTransactions{
		Id:               d.Id,
		BalanceAccountId: d.BalanceAccountId,
		MerchantPartyId:  d.MerchantPartyId,
		BalanceType:      d.BalanceType,
		CurrencyCode:     d.CurrencyCode,
		SourceType:       d.SourceType,
		SourceId:         d.SourceId,
		Direction:        d.Direction,
		Amount:           d.Amount,
		BalanceBefore:    d.BalanceBefore,
		BalanceAfter:     d.BalanceAfter,
		ReasonCode:       d.ReasonCode,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceTransactionsResponse struct {
	Id               uuid.UUID                                    `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BalanceAccountId uuid.UUID                                    `json:"balanceAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId  uuid.UUID                                    `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BalanceType      model.MerchantBalanceTransactionsBalanceType `json:"balanceType" validate:"required,oneof=pending available reserved payable paid_out negative disputed refundable" enums:"pending,available,reserved,payable,paid_out,negative,disputed,refundable"`
	CurrencyCode     string                                       `json:"currencyCode" validate:"required"`
	SourceType       string                                       `json:"sourceType" validate:"required"`
	SourceId         uuid.UUID                                    `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Direction        model.Direction                              `json:"direction" validate:"required,oneof=debit credit" enums:"debit,credit"`
	Amount           decimal.Decimal                              `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	BalanceBefore    decimal.Decimal                              `json:"balanceBefore" validate:"required" format:"decimal" example:"100.50"`
	BalanceAfter     decimal.Decimal                              `json:"balanceAfter" validate:"required" format:"decimal" example:"100.50"`
	ReasonCode       string                                       `json:"reasonCode" validate:"required"`
	Metadata         json.RawMessage                              `json:"metadata" swaggertype:"object"`
}

func NewMerchantBalanceTransactionsResponse(merchantBalanceTransactions model.MerchantBalanceTransactions) MerchantBalanceTransactionsResponse {
	return MerchantBalanceTransactionsResponse{
		Id:               merchantBalanceTransactions.Id,
		BalanceAccountId: merchantBalanceTransactions.BalanceAccountId,
		MerchantPartyId:  merchantBalanceTransactions.MerchantPartyId,
		BalanceType:      model.MerchantBalanceTransactionsBalanceType(merchantBalanceTransactions.BalanceType),
		CurrencyCode:     merchantBalanceTransactions.CurrencyCode,
		SourceType:       merchantBalanceTransactions.SourceType,
		SourceId:         merchantBalanceTransactions.SourceId,
		Direction:        model.Direction(merchantBalanceTransactions.Direction),
		Amount:           merchantBalanceTransactions.Amount,
		BalanceBefore:    merchantBalanceTransactions.BalanceBefore,
		BalanceAfter:     merchantBalanceTransactions.BalanceAfter,
		ReasonCode:       merchantBalanceTransactions.ReasonCode,
		Metadata:         merchantBalanceTransactions.Metadata,
	}
}

type MerchantBalanceTransactionsListResponse []*MerchantBalanceTransactionsResponse

func NewMerchantBalanceTransactionsListResponse(merchantBalanceTransactionsList model.MerchantBalanceTransactionsList) MerchantBalanceTransactionsListResponse {
	dtoMerchantBalanceTransactionsListResponse := MerchantBalanceTransactionsListResponse{}
	for _, merchantBalanceTransactions := range merchantBalanceTransactionsList {
		dtoMerchantBalanceTransactionsResponse := NewMerchantBalanceTransactionsResponse(*merchantBalanceTransactions)
		dtoMerchantBalanceTransactionsListResponse = append(dtoMerchantBalanceTransactionsListResponse, &dtoMerchantBalanceTransactionsResponse)
	}
	return dtoMerchantBalanceTransactionsListResponse
}

type MerchantBalanceTransactionsPrimaryIDList []MerchantBalanceTransactionsPrimaryID

func (d MerchantBalanceTransactionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceTransactions := range d {
		err = validator.Struct(merchantBalanceTransactions)
		if err != nil {
			return
		}
	}
	return nil
}

type MerchantBalanceTransactionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *MerchantBalanceTransactionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d MerchantBalanceTransactionsPrimaryID) ToModel() model.MerchantBalanceTransactionsPrimaryID {
	return model.MerchantBalanceTransactionsPrimaryID{
		Id: d.Id,
	}
}

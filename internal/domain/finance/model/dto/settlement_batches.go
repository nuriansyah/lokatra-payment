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

type SettlementBatchesDTOFieldNameType string

type settlementBatchesDTOFieldName struct {
	Id               SettlementBatchesDTOFieldNameType
	BatchCode        SettlementBatchesDTOFieldNameType
	MerchantPartyId  SettlementBatchesDTOFieldNameType
	CurrencyCode     SettlementBatchesDTOFieldNameType
	PeriodStart      SettlementBatchesDTOFieldNameType
	PeriodEnd        SettlementBatchesDTOFieldNameType
	GrossAmount      SettlementBatchesDTOFieldNameType
	FeeAmount        SettlementBatchesDTOFieldNameType
	TaxAmount        SettlementBatchesDTOFieldNameType
	ReserveAmount    SettlementBatchesDTOFieldNameType
	AdjustmentAmount SettlementBatchesDTOFieldNameType
	NetAmount        SettlementBatchesDTOFieldNameType
	BatchStatus      SettlementBatchesDTOFieldNameType
	ApprovedAt       SettlementBatchesDTOFieldNameType
	LockedAt         SettlementBatchesDTOFieldNameType
	Metadata         SettlementBatchesDTOFieldNameType
	MetaCreatedAt    SettlementBatchesDTOFieldNameType
	MetaCreatedBy    SettlementBatchesDTOFieldNameType
	MetaUpdatedAt    SettlementBatchesDTOFieldNameType
	MetaUpdatedBy    SettlementBatchesDTOFieldNameType
	MetaDeletedAt    SettlementBatchesDTOFieldNameType
	MetaDeletedBy    SettlementBatchesDTOFieldNameType
}

var SettlementBatchesDTOFieldName = settlementBatchesDTOFieldName{
	Id:               "id",
	BatchCode:        "batchCode",
	MerchantPartyId:  "merchantPartyId",
	CurrencyCode:     "currencyCode",
	PeriodStart:      "periodStart",
	PeriodEnd:        "periodEnd",
	GrossAmount:      "grossAmount",
	FeeAmount:        "feeAmount",
	TaxAmount:        "taxAmount",
	ReserveAmount:    "reserveAmount",
	AdjustmentAmount: "adjustmentAmount",
	NetAmount:        "netAmount",
	BatchStatus:      "batchStatus",
	ApprovedAt:       "approvedAt",
	LockedAt:         "lockedAt",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformSettlementBatchesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(SettlementBatchesDTOFieldName.Id):
		return string(model.SettlementBatchesDBFieldName.Id), true

	case string(SettlementBatchesDTOFieldName.BatchCode):
		return string(model.SettlementBatchesDBFieldName.BatchCode), true

	case string(SettlementBatchesDTOFieldName.MerchantPartyId):
		return string(model.SettlementBatchesDBFieldName.MerchantPartyId), true

	case string(SettlementBatchesDTOFieldName.CurrencyCode):
		return string(model.SettlementBatchesDBFieldName.CurrencyCode), true

	case string(SettlementBatchesDTOFieldName.PeriodStart):
		return string(model.SettlementBatchesDBFieldName.PeriodStart), true

	case string(SettlementBatchesDTOFieldName.PeriodEnd):
		return string(model.SettlementBatchesDBFieldName.PeriodEnd), true

	case string(SettlementBatchesDTOFieldName.GrossAmount):
		return string(model.SettlementBatchesDBFieldName.GrossAmount), true

	case string(SettlementBatchesDTOFieldName.FeeAmount):
		return string(model.SettlementBatchesDBFieldName.FeeAmount), true

	case string(SettlementBatchesDTOFieldName.TaxAmount):
		return string(model.SettlementBatchesDBFieldName.TaxAmount), true

	case string(SettlementBatchesDTOFieldName.ReserveAmount):
		return string(model.SettlementBatchesDBFieldName.ReserveAmount), true

	case string(SettlementBatchesDTOFieldName.AdjustmentAmount):
		return string(model.SettlementBatchesDBFieldName.AdjustmentAmount), true

	case string(SettlementBatchesDTOFieldName.NetAmount):
		return string(model.SettlementBatchesDBFieldName.NetAmount), true

	case string(SettlementBatchesDTOFieldName.BatchStatus):
		return string(model.SettlementBatchesDBFieldName.BatchStatus), true

	case string(SettlementBatchesDTOFieldName.ApprovedAt):
		return string(model.SettlementBatchesDBFieldName.ApprovedAt), true

	case string(SettlementBatchesDTOFieldName.LockedAt):
		return string(model.SettlementBatchesDBFieldName.LockedAt), true

	case string(SettlementBatchesDTOFieldName.Metadata):
		return string(model.SettlementBatchesDBFieldName.Metadata), true

	case string(SettlementBatchesDTOFieldName.MetaCreatedAt):
		return string(model.SettlementBatchesDBFieldName.MetaCreatedAt), true

	case string(SettlementBatchesDTOFieldName.MetaCreatedBy):
		return string(model.SettlementBatchesDBFieldName.MetaCreatedBy), true

	case string(SettlementBatchesDTOFieldName.MetaUpdatedAt):
		return string(model.SettlementBatchesDBFieldName.MetaUpdatedAt), true

	case string(SettlementBatchesDTOFieldName.MetaUpdatedBy):
		return string(model.SettlementBatchesDBFieldName.MetaUpdatedBy), true

	case string(SettlementBatchesDTOFieldName.MetaDeletedAt):
		return string(model.SettlementBatchesDBFieldName.MetaDeletedAt), true

	case string(SettlementBatchesDTOFieldName.MetaDeletedBy):
		return string(model.SettlementBatchesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewSettlementBatchesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isSettlementBatchesBaseFilterField(field string) bool {
	spec, found := model.NewSettlementBatchesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeSettlementBatchesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateSettlementBatchesProjectionOutputPath(path string) error {
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

func transformSettlementBatchesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformSettlementBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformSettlementBatchesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformSettlementBatchesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformSettlementBatchesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isSettlementBatchesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateSettlementBatchesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeSettlementBatchesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformSettlementBatchesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformSettlementBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformSettlementBatchesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultSettlementBatchesFilter(filter *model.Filter) {
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
			Field: string(SettlementBatchesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type SettlementBatchesSelectableResponse map[string]interface{}
type SettlementBatchesSelectableListResponse []*SettlementBatchesSelectableResponse

func assignSettlementBatchesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setSettlementBatchesSelectableValue(out SettlementBatchesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignSettlementBatchesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewSettlementBatchesSelectableResponse(settlementBatches model.SettlementBatches, filter model.Filter) SettlementBatchesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.SettlementBatchesDBFieldName.Id),
			string(model.SettlementBatchesDBFieldName.BatchCode),
			string(model.SettlementBatchesDBFieldName.MerchantPartyId),
			string(model.SettlementBatchesDBFieldName.CurrencyCode),
			string(model.SettlementBatchesDBFieldName.PeriodStart),
			string(model.SettlementBatchesDBFieldName.PeriodEnd),
			string(model.SettlementBatchesDBFieldName.GrossAmount),
			string(model.SettlementBatchesDBFieldName.FeeAmount),
			string(model.SettlementBatchesDBFieldName.TaxAmount),
			string(model.SettlementBatchesDBFieldName.ReserveAmount),
			string(model.SettlementBatchesDBFieldName.AdjustmentAmount),
			string(model.SettlementBatchesDBFieldName.NetAmount),
			string(model.SettlementBatchesDBFieldName.BatchStatus),
			string(model.SettlementBatchesDBFieldName.ApprovedAt),
			string(model.SettlementBatchesDBFieldName.LockedAt),
			string(model.SettlementBatchesDBFieldName.Metadata),
			string(model.SettlementBatchesDBFieldName.MetaCreatedAt),
			string(model.SettlementBatchesDBFieldName.MetaCreatedBy),
			string(model.SettlementBatchesDBFieldName.MetaUpdatedAt),
			string(model.SettlementBatchesDBFieldName.MetaUpdatedBy),
			string(model.SettlementBatchesDBFieldName.MetaDeletedAt),
			string(model.SettlementBatchesDBFieldName.MetaDeletedBy),
		)
	}
	settlementBatchesSelectableResponse := SettlementBatchesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.SettlementBatchesDBFieldName.Id):
			key := string(SettlementBatchesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.Id, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.BatchCode):
			key := string(SettlementBatchesDTOFieldName.BatchCode)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.BatchCode, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MerchantPartyId):
			key := string(SettlementBatchesDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MerchantPartyId, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.CurrencyCode):
			key := string(SettlementBatchesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.CurrencyCode, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.PeriodStart):
			key := string(SettlementBatchesDTOFieldName.PeriodStart)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.PeriodStart, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.PeriodEnd):
			key := string(SettlementBatchesDTOFieldName.PeriodEnd)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.PeriodEnd, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.GrossAmount):
			key := string(SettlementBatchesDTOFieldName.GrossAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.GrossAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.FeeAmount):
			key := string(SettlementBatchesDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.FeeAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.TaxAmount):
			key := string(SettlementBatchesDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.TaxAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.ReserveAmount):
			key := string(SettlementBatchesDTOFieldName.ReserveAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.ReserveAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.AdjustmentAmount):
			key := string(SettlementBatchesDTOFieldName.AdjustmentAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.AdjustmentAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.NetAmount):
			key := string(SettlementBatchesDTOFieldName.NetAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.NetAmount, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.BatchStatus):
			key := string(SettlementBatchesDTOFieldName.BatchStatus)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, model.SettlementBatchesBatchStatus(settlementBatches.BatchStatus), explicitAlias)

		case string(model.SettlementBatchesDBFieldName.ApprovedAt):
			key := string(SettlementBatchesDTOFieldName.ApprovedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.ApprovedAt.Time, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.LockedAt):
			key := string(SettlementBatchesDTOFieldName.LockedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.LockedAt.Time, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.Metadata):
			key := string(SettlementBatchesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.Metadata, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaCreatedAt):
			key := string(SettlementBatchesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaCreatedAt, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaCreatedBy):
			key := string(SettlementBatchesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaCreatedBy, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaUpdatedAt):
			key := string(SettlementBatchesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaUpdatedAt, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaUpdatedBy):
			key := string(SettlementBatchesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaUpdatedBy, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaDeletedAt):
			key := string(SettlementBatchesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaDeletedAt.Time, explicitAlias)

		case string(model.SettlementBatchesDBFieldName.MetaDeletedBy):
			key := string(SettlementBatchesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchesSelectableValue(settlementBatchesSelectableResponse, key, settlementBatches.MetaDeletedBy, explicitAlias)

		}
	}
	return settlementBatchesSelectableResponse
}

func NewSettlementBatchesListResponseFromFilterResult(result []model.SettlementBatchesFilterResult, filter model.Filter) SettlementBatchesSelectableListResponse {
	dtoSettlementBatchesListResponse := SettlementBatchesSelectableListResponse{}
	for _, row := range result {
		dtoSettlementBatchesResponse := NewSettlementBatchesSelectableResponse(row.SettlementBatches, filter)
		dtoSettlementBatchesListResponse = append(dtoSettlementBatchesListResponse, &dtoSettlementBatchesResponse)
	}
	return dtoSettlementBatchesListResponse
}

type SettlementBatchesFilterResponse struct {
	Metadata Metadata                                `json:"metadata"`
	Data     SettlementBatchesSelectableListResponse `json:"data"`
}

func reverseSettlementBatchesFilterResults(result []model.SettlementBatchesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewSettlementBatchesFilterResponse(result []model.SettlementBatchesFilterResult, filter model.Filter) (resp SettlementBatchesFilterResponse) {
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
			reverseSettlementBatchesFilterResults(dataResult)
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

	resp.Data = NewSettlementBatchesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type SettlementBatchesCreateRequest struct {
	BatchCode        string                             `json:"batchCode"`
	MerchantPartyId  uuid.UUID                          `json:"merchantPartyId"`
	CurrencyCode     string                             `json:"currencyCode"`
	PeriodStart      time.Time                          `json:"periodStart"`
	PeriodEnd        time.Time                          `json:"periodEnd"`
	GrossAmount      decimal.Decimal                    `json:"grossAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	ReserveAmount    decimal.Decimal                    `json:"reserveAmount"`
	AdjustmentAmount decimal.Decimal                    `json:"adjustmentAmount"`
	NetAmount        decimal.Decimal                    `json:"netAmount"`
	BatchStatus      model.SettlementBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,ready,funding,funded,payouting,paid,partially_paid,failed,cancelled"`
	ApprovedAt       time.Time                          `json:"approvedAt"`
	LockedAt         time.Time                          `json:"lockedAt"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d *SettlementBatchesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *SettlementBatchesCreateRequest) ToModel() model.SettlementBatches {
	id, _ := uuid.NewV4()
	return model.SettlementBatches{
		Id:               id,
		BatchCode:        d.BatchCode,
		MerchantPartyId:  d.MerchantPartyId,
		CurrencyCode:     d.CurrencyCode,
		PeriodStart:      d.PeriodStart,
		PeriodEnd:        d.PeriodEnd,
		GrossAmount:      d.GrossAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		ReserveAmount:    d.ReserveAmount,
		AdjustmentAmount: d.AdjustmentAmount,
		NetAmount:        d.NetAmount,
		BatchStatus:      d.BatchStatus,
		ApprovedAt:       null.TimeFrom(d.ApprovedAt),
		LockedAt:         null.TimeFrom(d.LockedAt),
		Metadata:         d.Metadata,
	}
}

type SettlementBatchesListCreateRequest []*SettlementBatchesCreateRequest

func (d SettlementBatchesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatches := range d {
		err = validator.Struct(settlementBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementBatchesListCreateRequest) ToModelList() []model.SettlementBatches {
	out := make([]model.SettlementBatches, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type SettlementBatchesUpdateRequest struct {
	BatchCode        string                             `json:"batchCode"`
	MerchantPartyId  uuid.UUID                          `json:"merchantPartyId"`
	CurrencyCode     string                             `json:"currencyCode"`
	PeriodStart      time.Time                          `json:"periodStart"`
	PeriodEnd        time.Time                          `json:"periodEnd"`
	GrossAmount      decimal.Decimal                    `json:"grossAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	ReserveAmount    decimal.Decimal                    `json:"reserveAmount"`
	AdjustmentAmount decimal.Decimal                    `json:"adjustmentAmount"`
	NetAmount        decimal.Decimal                    `json:"netAmount"`
	BatchStatus      model.SettlementBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,ready,funding,funded,payouting,paid,partially_paid,failed,cancelled"`
	ApprovedAt       time.Time                          `json:"approvedAt"`
	LockedAt         time.Time                          `json:"lockedAt"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d *SettlementBatchesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d SettlementBatchesUpdateRequest) ToModel() model.SettlementBatches {
	return model.SettlementBatches{
		BatchCode:        d.BatchCode,
		MerchantPartyId:  d.MerchantPartyId,
		CurrencyCode:     d.CurrencyCode,
		PeriodStart:      d.PeriodStart,
		PeriodEnd:        d.PeriodEnd,
		GrossAmount:      d.GrossAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		ReserveAmount:    d.ReserveAmount,
		AdjustmentAmount: d.AdjustmentAmount,
		NetAmount:        d.NetAmount,
		BatchStatus:      d.BatchStatus,
		ApprovedAt:       null.TimeFrom(d.ApprovedAt),
		LockedAt:         null.TimeFrom(d.LockedAt),
		Metadata:         d.Metadata,
	}
}

type SettlementBatchesBulkUpdateRequest struct {
	Id               uuid.UUID                          `json:"id"`
	BatchCode        string                             `json:"batchCode"`
	MerchantPartyId  uuid.UUID                          `json:"merchantPartyId"`
	CurrencyCode     string                             `json:"currencyCode"`
	PeriodStart      time.Time                          `json:"periodStart"`
	PeriodEnd        time.Time                          `json:"periodEnd"`
	GrossAmount      decimal.Decimal                    `json:"grossAmount"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount"`
	ReserveAmount    decimal.Decimal                    `json:"reserveAmount"`
	AdjustmentAmount decimal.Decimal                    `json:"adjustmentAmount"`
	NetAmount        decimal.Decimal                    `json:"netAmount"`
	BatchStatus      model.SettlementBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,ready,funding,funded,payouting,paid,partially_paid,failed,cancelled"`
	ApprovedAt       time.Time                          `json:"approvedAt"`
	LockedAt         time.Time                          `json:"lockedAt"`
	Metadata         json.RawMessage                    `json:"metadata"`
}

func (d SettlementBatchesBulkUpdateRequest) PrimaryID() SettlementBatchesPrimaryID {
	return SettlementBatchesPrimaryID{
		Id: d.Id,
	}
}

type SettlementBatchesListBulkUpdateRequest []*SettlementBatchesBulkUpdateRequest

func (d SettlementBatchesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatches := range d {
		err = validator.Struct(settlementBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementBatchesBulkUpdateRequest) ToModel() model.SettlementBatches {
	return model.SettlementBatches{
		Id:               d.Id,
		BatchCode:        d.BatchCode,
		MerchantPartyId:  d.MerchantPartyId,
		CurrencyCode:     d.CurrencyCode,
		PeriodStart:      d.PeriodStart,
		PeriodEnd:        d.PeriodEnd,
		GrossAmount:      d.GrossAmount,
		FeeAmount:        d.FeeAmount,
		TaxAmount:        d.TaxAmount,
		ReserveAmount:    d.ReserveAmount,
		AdjustmentAmount: d.AdjustmentAmount,
		NetAmount:        d.NetAmount,
		BatchStatus:      d.BatchStatus,
		ApprovedAt:       null.TimeFrom(d.ApprovedAt),
		LockedAt:         null.TimeFrom(d.LockedAt),
		Metadata:         d.Metadata,
	}
}

type SettlementBatchesResponse struct {
	Id               uuid.UUID                          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BatchCode        string                             `json:"batchCode" validate:"required"`
	MerchantPartyId  uuid.UUID                          `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode     string                             `json:"currencyCode" validate:"required"`
	PeriodStart      time.Time                          `json:"periodStart" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PeriodEnd        time.Time                          `json:"periodEnd" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	GrossAmount      decimal.Decimal                    `json:"grossAmount" format:"decimal" example:"100.50"`
	FeeAmount        decimal.Decimal                    `json:"feeAmount" format:"decimal" example:"100.50"`
	TaxAmount        decimal.Decimal                    `json:"taxAmount" format:"decimal" example:"100.50"`
	ReserveAmount    decimal.Decimal                    `json:"reserveAmount" format:"decimal" example:"100.50"`
	AdjustmentAmount decimal.Decimal                    `json:"adjustmentAmount" format:"decimal" example:"100.50"`
	NetAmount        decimal.Decimal                    `json:"netAmount" format:"decimal" example:"100.50"`
	BatchStatus      model.SettlementBatchesBatchStatus `json:"batchStatus" validate:"oneof=draft ready funding funded payouting paid partially_paid failed cancelled" enums:"draft,ready,funding,funded,payouting,paid,partially_paid,failed,cancelled"`
	ApprovedAt       time.Time                          `json:"approvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LockedAt         time.Time                          `json:"lockedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata         json.RawMessage                    `json:"metadata" swaggertype:"object"`
}

func NewSettlementBatchesResponse(settlementBatches model.SettlementBatches) SettlementBatchesResponse {
	return SettlementBatchesResponse{
		Id:               settlementBatches.Id,
		BatchCode:        settlementBatches.BatchCode,
		MerchantPartyId:  settlementBatches.MerchantPartyId,
		CurrencyCode:     settlementBatches.CurrencyCode,
		PeriodStart:      settlementBatches.PeriodStart,
		PeriodEnd:        settlementBatches.PeriodEnd,
		GrossAmount:      settlementBatches.GrossAmount,
		FeeAmount:        settlementBatches.FeeAmount,
		TaxAmount:        settlementBatches.TaxAmount,
		ReserveAmount:    settlementBatches.ReserveAmount,
		AdjustmentAmount: settlementBatches.AdjustmentAmount,
		NetAmount:        settlementBatches.NetAmount,
		BatchStatus:      model.SettlementBatchesBatchStatus(settlementBatches.BatchStatus),
		ApprovedAt:       settlementBatches.ApprovedAt.Time,
		LockedAt:         settlementBatches.LockedAt.Time,
		Metadata:         settlementBatches.Metadata,
	}
}

type SettlementBatchesListResponse []*SettlementBatchesResponse

func NewSettlementBatchesListResponse(settlementBatchesList model.SettlementBatchesList) SettlementBatchesListResponse {
	dtoSettlementBatchesListResponse := SettlementBatchesListResponse{}
	for _, settlementBatches := range settlementBatchesList {
		dtoSettlementBatchesResponse := NewSettlementBatchesResponse(*settlementBatches)
		dtoSettlementBatchesListResponse = append(dtoSettlementBatchesListResponse, &dtoSettlementBatchesResponse)
	}
	return dtoSettlementBatchesListResponse
}

type SettlementBatchesPrimaryIDList []SettlementBatchesPrimaryID

func (d SettlementBatchesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatches := range d {
		err = validator.Struct(settlementBatches)
		if err != nil {
			return
		}
	}
	return nil
}

type SettlementBatchesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *SettlementBatchesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d SettlementBatchesPrimaryID) ToModel() model.SettlementBatchesPrimaryID {
	return model.SettlementBatchesPrimaryID{
		Id: d.Id,
	}
}

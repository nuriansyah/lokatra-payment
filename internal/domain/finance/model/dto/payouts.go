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

type PayoutsDTOFieldNameType string

type payoutsDTOFieldName struct {
	Id                PayoutsDTOFieldNameType
	PayoutCode        PayoutsDTOFieldNameType
	PayoutBatchId     PayoutsDTOFieldNameType
	SettlementBatchId PayoutsDTOFieldNameType
	MerchantPartyId   PayoutsDTOFieldNameType
	PayoutMethodId    PayoutsDTOFieldNameType
	ProviderAccountId PayoutsDTOFieldNameType
	CurrencyCode      PayoutsDTOFieldNameType
	Amount            PayoutsDTOFieldNameType
	FeeAmount         PayoutsDTOFieldNameType
	NetSentAmount     PayoutsDTOFieldNameType
	IdempotencyKey    PayoutsDTOFieldNameType
	ProviderPayoutRef PayoutsDTOFieldNameType
	PayoutStatus      PayoutsDTOFieldNameType
	HoldReasonCode    PayoutsDTOFieldNameType
	InitiatedAt       PayoutsDTOFieldNameType
	CompletedAt       PayoutsDTOFieldNameType
	FailedAt          PayoutsDTOFieldNameType
	FailureCode       PayoutsDTOFieldNameType
	FailureReason     PayoutsDTOFieldNameType
	Metadata          PayoutsDTOFieldNameType
	MetaCreatedAt     PayoutsDTOFieldNameType
	MetaCreatedBy     PayoutsDTOFieldNameType
	MetaUpdatedAt     PayoutsDTOFieldNameType
	MetaUpdatedBy     PayoutsDTOFieldNameType
	MetaDeletedAt     PayoutsDTOFieldNameType
	MetaDeletedBy     PayoutsDTOFieldNameType
}

var PayoutsDTOFieldName = payoutsDTOFieldName{
	Id:                "id",
	PayoutCode:        "payoutCode",
	PayoutBatchId:     "payoutBatchId",
	SettlementBatchId: "settlementBatchId",
	MerchantPartyId:   "merchantPartyId",
	PayoutMethodId:    "payoutMethodId",
	ProviderAccountId: "providerAccountId",
	CurrencyCode:      "currencyCode",
	Amount:            "amount",
	FeeAmount:         "feeAmount",
	NetSentAmount:     "netSentAmount",
	IdempotencyKey:    "idempotencyKey",
	ProviderPayoutRef: "providerPayoutRef",
	PayoutStatus:      "payoutStatus",
	HoldReasonCode:    "holdReasonCode",
	InitiatedAt:       "initiatedAt",
	CompletedAt:       "completedAt",
	FailedAt:          "failedAt",
	FailureCode:       "failureCode",
	FailureReason:     "failureReason",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformPayoutsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutsDTOFieldName.Id):
		return string(model.PayoutsDBFieldName.Id), true

	case string(PayoutsDTOFieldName.PayoutCode):
		return string(model.PayoutsDBFieldName.PayoutCode), true

	case string(PayoutsDTOFieldName.PayoutBatchId):
		return string(model.PayoutsDBFieldName.PayoutBatchId), true

	case string(PayoutsDTOFieldName.SettlementBatchId):
		return string(model.PayoutsDBFieldName.SettlementBatchId), true

	case string(PayoutsDTOFieldName.MerchantPartyId):
		return string(model.PayoutsDBFieldName.MerchantPartyId), true

	case string(PayoutsDTOFieldName.PayoutMethodId):
		return string(model.PayoutsDBFieldName.PayoutMethodId), true

	case string(PayoutsDTOFieldName.ProviderAccountId):
		return string(model.PayoutsDBFieldName.ProviderAccountId), true

	case string(PayoutsDTOFieldName.CurrencyCode):
		return string(model.PayoutsDBFieldName.CurrencyCode), true

	case string(PayoutsDTOFieldName.Amount):
		return string(model.PayoutsDBFieldName.Amount), true

	case string(PayoutsDTOFieldName.FeeAmount):
		return string(model.PayoutsDBFieldName.FeeAmount), true

	case string(PayoutsDTOFieldName.NetSentAmount):
		return string(model.PayoutsDBFieldName.NetSentAmount), true

	case string(PayoutsDTOFieldName.IdempotencyKey):
		return string(model.PayoutsDBFieldName.IdempotencyKey), true

	case string(PayoutsDTOFieldName.ProviderPayoutRef):
		return string(model.PayoutsDBFieldName.ProviderPayoutRef), true

	case string(PayoutsDTOFieldName.PayoutStatus):
		return string(model.PayoutsDBFieldName.PayoutStatus), true

	case string(PayoutsDTOFieldName.HoldReasonCode):
		return string(model.PayoutsDBFieldName.HoldReasonCode), true

	case string(PayoutsDTOFieldName.InitiatedAt):
		return string(model.PayoutsDBFieldName.InitiatedAt), true

	case string(PayoutsDTOFieldName.CompletedAt):
		return string(model.PayoutsDBFieldName.CompletedAt), true

	case string(PayoutsDTOFieldName.FailedAt):
		return string(model.PayoutsDBFieldName.FailedAt), true

	case string(PayoutsDTOFieldName.FailureCode):
		return string(model.PayoutsDBFieldName.FailureCode), true

	case string(PayoutsDTOFieldName.FailureReason):
		return string(model.PayoutsDBFieldName.FailureReason), true

	case string(PayoutsDTOFieldName.Metadata):
		return string(model.PayoutsDBFieldName.Metadata), true

	case string(PayoutsDTOFieldName.MetaCreatedAt):
		return string(model.PayoutsDBFieldName.MetaCreatedAt), true

	case string(PayoutsDTOFieldName.MetaCreatedBy):
		return string(model.PayoutsDBFieldName.MetaCreatedBy), true

	case string(PayoutsDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutsDBFieldName.MetaUpdatedAt), true

	case string(PayoutsDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutsDBFieldName.MetaUpdatedBy), true

	case string(PayoutsDTOFieldName.MetaDeletedAt):
		return string(model.PayoutsDBFieldName.MetaDeletedAt), true

	case string(PayoutsDTOFieldName.MetaDeletedBy):
		return string(model.PayoutsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutsBaseFilterField(field string) bool {
	spec, found := model.NewPayoutsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutsProjectionOutputPath(path string) error {
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

func transformPayoutsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutsFilter(filter *model.Filter) {
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
			Field: string(PayoutsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutsSelectableResponse map[string]interface{}
type PayoutsSelectableListResponse []*PayoutsSelectableResponse

func assignPayoutsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutsSelectableValue(out PayoutsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutsSelectableResponse(payouts model.Payouts, filter model.Filter) PayoutsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutsDBFieldName.Id),
			string(model.PayoutsDBFieldName.PayoutCode),
			string(model.PayoutsDBFieldName.PayoutBatchId),
			string(model.PayoutsDBFieldName.SettlementBatchId),
			string(model.PayoutsDBFieldName.MerchantPartyId),
			string(model.PayoutsDBFieldName.PayoutMethodId),
			string(model.PayoutsDBFieldName.ProviderAccountId),
			string(model.PayoutsDBFieldName.CurrencyCode),
			string(model.PayoutsDBFieldName.Amount),
			string(model.PayoutsDBFieldName.FeeAmount),
			string(model.PayoutsDBFieldName.NetSentAmount),
			string(model.PayoutsDBFieldName.IdempotencyKey),
			string(model.PayoutsDBFieldName.ProviderPayoutRef),
			string(model.PayoutsDBFieldName.PayoutStatus),
			string(model.PayoutsDBFieldName.HoldReasonCode),
			string(model.PayoutsDBFieldName.InitiatedAt),
			string(model.PayoutsDBFieldName.CompletedAt),
			string(model.PayoutsDBFieldName.FailedAt),
			string(model.PayoutsDBFieldName.FailureCode),
			string(model.PayoutsDBFieldName.FailureReason),
			string(model.PayoutsDBFieldName.Metadata),
			string(model.PayoutsDBFieldName.MetaCreatedAt),
			string(model.PayoutsDBFieldName.MetaCreatedBy),
			string(model.PayoutsDBFieldName.MetaUpdatedAt),
			string(model.PayoutsDBFieldName.MetaUpdatedBy),
			string(model.PayoutsDBFieldName.MetaDeletedAt),
			string(model.PayoutsDBFieldName.MetaDeletedBy),
		)
	}
	payoutsSelectableResponse := PayoutsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutsDBFieldName.Id):
			key := string(PayoutsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.Id, explicitAlias)

		case string(model.PayoutsDBFieldName.PayoutCode):
			key := string(PayoutsDTOFieldName.PayoutCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.PayoutCode, explicitAlias)

		case string(model.PayoutsDBFieldName.PayoutBatchId):
			key := string(PayoutsDTOFieldName.PayoutBatchId)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.PayoutBatchId.UUID, explicitAlias)

		case string(model.PayoutsDBFieldName.SettlementBatchId):
			key := string(PayoutsDTOFieldName.SettlementBatchId)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.SettlementBatchId.UUID, explicitAlias)

		case string(model.PayoutsDBFieldName.MerchantPartyId):
			key := string(PayoutsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MerchantPartyId, explicitAlias)

		case string(model.PayoutsDBFieldName.PayoutMethodId):
			key := string(PayoutsDTOFieldName.PayoutMethodId)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.PayoutMethodId, explicitAlias)

		case string(model.PayoutsDBFieldName.ProviderAccountId):
			key := string(PayoutsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.ProviderAccountId.UUID, explicitAlias)

		case string(model.PayoutsDBFieldName.CurrencyCode):
			key := string(PayoutsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.CurrencyCode, explicitAlias)

		case string(model.PayoutsDBFieldName.Amount):
			key := string(PayoutsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.Amount, explicitAlias)

		case string(model.PayoutsDBFieldName.FeeAmount):
			key := string(PayoutsDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.FeeAmount, explicitAlias)

		case string(model.PayoutsDBFieldName.NetSentAmount):
			key := string(PayoutsDTOFieldName.NetSentAmount)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.NetSentAmount, explicitAlias)

		case string(model.PayoutsDBFieldName.IdempotencyKey):
			key := string(PayoutsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.IdempotencyKey, explicitAlias)

		case string(model.PayoutsDBFieldName.ProviderPayoutRef):
			key := string(PayoutsDTOFieldName.ProviderPayoutRef)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.ProviderPayoutRef.String, explicitAlias)

		case string(model.PayoutsDBFieldName.PayoutStatus):
			key := string(PayoutsDTOFieldName.PayoutStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, model.PayoutStatus(payouts.PayoutStatus), explicitAlias)

		case string(model.PayoutsDBFieldName.HoldReasonCode):
			key := string(PayoutsDTOFieldName.HoldReasonCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.HoldReasonCode.String, explicitAlias)

		case string(model.PayoutsDBFieldName.InitiatedAt):
			key := string(PayoutsDTOFieldName.InitiatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.InitiatedAt.Time, explicitAlias)

		case string(model.PayoutsDBFieldName.CompletedAt):
			key := string(PayoutsDTOFieldName.CompletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.CompletedAt.Time, explicitAlias)

		case string(model.PayoutsDBFieldName.FailedAt):
			key := string(PayoutsDTOFieldName.FailedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.FailedAt.Time, explicitAlias)

		case string(model.PayoutsDBFieldName.FailureCode):
			key := string(PayoutsDTOFieldName.FailureCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.FailureCode.String, explicitAlias)

		case string(model.PayoutsDBFieldName.FailureReason):
			key := string(PayoutsDTOFieldName.FailureReason)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.FailureReason.String, explicitAlias)

		case string(model.PayoutsDBFieldName.Metadata):
			key := string(PayoutsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.Metadata, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaCreatedAt):
			key := string(PayoutsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaCreatedAt, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaCreatedBy):
			key := string(PayoutsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaCreatedBy, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaUpdatedAt):
			key := string(PayoutsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaUpdatedBy):
			key := string(PayoutsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaDeletedAt):
			key := string(PayoutsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutsDBFieldName.MetaDeletedBy):
			key := string(PayoutsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutsSelectableValue(payoutsSelectableResponse, key, payouts.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutsSelectableResponse
}

func NewPayoutsListResponseFromFilterResult(result []model.PayoutsFilterResult, filter model.Filter) PayoutsSelectableListResponse {
	dtoPayoutsListResponse := PayoutsSelectableListResponse{}
	for _, row := range result {
		dtoPayoutsResponse := NewPayoutsSelectableResponse(row.Payouts, filter)
		dtoPayoutsListResponse = append(dtoPayoutsListResponse, &dtoPayoutsResponse)
	}
	return dtoPayoutsListResponse
}

type PayoutsFilterResponse struct {
	Metadata Metadata                      `json:"metadata"`
	Data     PayoutsSelectableListResponse `json:"data"`
}

func reversePayoutsFilterResults(result []model.PayoutsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutsFilterResponse(result []model.PayoutsFilterResult, filter model.Filter) (resp PayoutsFilterResponse) {
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
			reversePayoutsFilterResults(dataResult)
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

	resp.Data = NewPayoutsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutsCreateRequest struct {
	PayoutCode        string             `json:"payoutCode"`
	PayoutBatchId     uuid.UUID          `json:"payoutBatchId"`
	SettlementBatchId uuid.UUID          `json:"settlementBatchId"`
	MerchantPartyId   uuid.UUID          `json:"merchantPartyId"`
	PayoutMethodId    uuid.UUID          `json:"payoutMethodId"`
	ProviderAccountId uuid.UUID          `json:"providerAccountId"`
	CurrencyCode      string             `json:"currencyCode"`
	Amount            decimal.Decimal    `json:"amount"`
	FeeAmount         decimal.Decimal    `json:"feeAmount"`
	NetSentAmount     decimal.Decimal    `json:"netSentAmount"`
	IdempotencyKey    string             `json:"idempotencyKey"`
	ProviderPayoutRef string             `json:"providerPayoutRef"`
	PayoutStatus      model.PayoutStatus `json:"payoutStatus" example:"created" enums:"created,queued,processing,succeeded,failed,reversed,cancelled,held"`
	HoldReasonCode    string             `json:"holdReasonCode"`
	InitiatedAt       time.Time          `json:"initiatedAt"`
	CompletedAt       time.Time          `json:"completedAt"`
	FailedAt          time.Time          `json:"failedAt"`
	FailureCode       string             `json:"failureCode"`
	FailureReason     string             `json:"failureReason"`
	Metadata          json.RawMessage    `json:"metadata"`
}

func (d *PayoutsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutsCreateRequest) ToModel() model.Payouts {
	id, _ := uuid.NewV4()
	return model.Payouts{
		Id:                id,
		PayoutCode:        d.PayoutCode,
		PayoutBatchId:     nuuid.From(d.PayoutBatchId),
		SettlementBatchId: nuuid.From(d.SettlementBatchId),
		MerchantPartyId:   d.MerchantPartyId,
		PayoutMethodId:    d.PayoutMethodId,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		Amount:            d.Amount,
		FeeAmount:         d.FeeAmount,
		NetSentAmount:     d.NetSentAmount,
		IdempotencyKey:    d.IdempotencyKey,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		PayoutStatus:      d.PayoutStatus,
		HoldReasonCode:    null.StringFrom(d.HoldReasonCode),
		InitiatedAt:       null.TimeFrom(d.InitiatedAt),
		CompletedAt:       null.TimeFrom(d.CompletedAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		Metadata:          d.Metadata,
	}
}

type PayoutsListCreateRequest []*PayoutsCreateRequest

func (d PayoutsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payouts := range d {
		err = validator.Struct(payouts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutsListCreateRequest) ToModelList() []model.Payouts {
	out := make([]model.Payouts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutsUpdateRequest struct {
	PayoutCode        string             `json:"payoutCode"`
	PayoutBatchId     uuid.UUID          `json:"payoutBatchId"`
	SettlementBatchId uuid.UUID          `json:"settlementBatchId"`
	MerchantPartyId   uuid.UUID          `json:"merchantPartyId"`
	PayoutMethodId    uuid.UUID          `json:"payoutMethodId"`
	ProviderAccountId uuid.UUID          `json:"providerAccountId"`
	CurrencyCode      string             `json:"currencyCode"`
	Amount            decimal.Decimal    `json:"amount"`
	FeeAmount         decimal.Decimal    `json:"feeAmount"`
	NetSentAmount     decimal.Decimal    `json:"netSentAmount"`
	IdempotencyKey    string             `json:"idempotencyKey"`
	ProviderPayoutRef string             `json:"providerPayoutRef"`
	PayoutStatus      model.PayoutStatus `json:"payoutStatus" example:"created" enums:"created,queued,processing,succeeded,failed,reversed,cancelled,held"`
	HoldReasonCode    string             `json:"holdReasonCode"`
	InitiatedAt       time.Time          `json:"initiatedAt"`
	CompletedAt       time.Time          `json:"completedAt"`
	FailedAt          time.Time          `json:"failedAt"`
	FailureCode       string             `json:"failureCode"`
	FailureReason     string             `json:"failureReason"`
	Metadata          json.RawMessage    `json:"metadata"`
}

func (d *PayoutsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutsUpdateRequest) ToModel() model.Payouts {
	return model.Payouts{
		PayoutCode:        d.PayoutCode,
		PayoutBatchId:     nuuid.From(d.PayoutBatchId),
		SettlementBatchId: nuuid.From(d.SettlementBatchId),
		MerchantPartyId:   d.MerchantPartyId,
		PayoutMethodId:    d.PayoutMethodId,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		Amount:            d.Amount,
		FeeAmount:         d.FeeAmount,
		NetSentAmount:     d.NetSentAmount,
		IdempotencyKey:    d.IdempotencyKey,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		PayoutStatus:      d.PayoutStatus,
		HoldReasonCode:    null.StringFrom(d.HoldReasonCode),
		InitiatedAt:       null.TimeFrom(d.InitiatedAt),
		CompletedAt:       null.TimeFrom(d.CompletedAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		Metadata:          d.Metadata,
	}
}

type PayoutsBulkUpdateRequest struct {
	Id                uuid.UUID          `json:"id"`
	PayoutCode        string             `json:"payoutCode"`
	PayoutBatchId     uuid.UUID          `json:"payoutBatchId"`
	SettlementBatchId uuid.UUID          `json:"settlementBatchId"`
	MerchantPartyId   uuid.UUID          `json:"merchantPartyId"`
	PayoutMethodId    uuid.UUID          `json:"payoutMethodId"`
	ProviderAccountId uuid.UUID          `json:"providerAccountId"`
	CurrencyCode      string             `json:"currencyCode"`
	Amount            decimal.Decimal    `json:"amount"`
	FeeAmount         decimal.Decimal    `json:"feeAmount"`
	NetSentAmount     decimal.Decimal    `json:"netSentAmount"`
	IdempotencyKey    string             `json:"idempotencyKey"`
	ProviderPayoutRef string             `json:"providerPayoutRef"`
	PayoutStatus      model.PayoutStatus `json:"payoutStatus" example:"created" enums:"created,queued,processing,succeeded,failed,reversed,cancelled,held"`
	HoldReasonCode    string             `json:"holdReasonCode"`
	InitiatedAt       time.Time          `json:"initiatedAt"`
	CompletedAt       time.Time          `json:"completedAt"`
	FailedAt          time.Time          `json:"failedAt"`
	FailureCode       string             `json:"failureCode"`
	FailureReason     string             `json:"failureReason"`
	Metadata          json.RawMessage    `json:"metadata"`
}

func (d PayoutsBulkUpdateRequest) PrimaryID() PayoutsPrimaryID {
	return PayoutsPrimaryID{
		Id: d.Id,
	}
}

type PayoutsListBulkUpdateRequest []*PayoutsBulkUpdateRequest

func (d PayoutsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payouts := range d {
		err = validator.Struct(payouts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutsBulkUpdateRequest) ToModel() model.Payouts {
	return model.Payouts{
		Id:                d.Id,
		PayoutCode:        d.PayoutCode,
		PayoutBatchId:     nuuid.From(d.PayoutBatchId),
		SettlementBatchId: nuuid.From(d.SettlementBatchId),
		MerchantPartyId:   d.MerchantPartyId,
		PayoutMethodId:    d.PayoutMethodId,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		Amount:            d.Amount,
		FeeAmount:         d.FeeAmount,
		NetSentAmount:     d.NetSentAmount,
		IdempotencyKey:    d.IdempotencyKey,
		ProviderPayoutRef: null.StringFrom(d.ProviderPayoutRef),
		PayoutStatus:      d.PayoutStatus,
		HoldReasonCode:    null.StringFrom(d.HoldReasonCode),
		InitiatedAt:       null.TimeFrom(d.InitiatedAt),
		CompletedAt:       null.TimeFrom(d.CompletedAt),
		FailedAt:          null.TimeFrom(d.FailedAt),
		FailureCode:       null.StringFrom(d.FailureCode),
		FailureReason:     null.StringFrom(d.FailureReason),
		Metadata:          d.Metadata,
	}
}

type PayoutsResponse struct {
	Id                uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutCode        string             `json:"payoutCode" validate:"required"`
	PayoutBatchId     uuid.UUID          `json:"payoutBatchId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SettlementBatchId uuid.UUID          `json:"settlementBatchId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId   uuid.UUID          `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutMethodId    uuid.UUID          `json:"payoutMethodId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID          `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode      string             `json:"currencyCode" validate:"required"`
	Amount            decimal.Decimal    `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	FeeAmount         decimal.Decimal    `json:"feeAmount" format:"decimal" example:"100.50"`
	NetSentAmount     decimal.Decimal    `json:"netSentAmount" validate:"required" format:"decimal" example:"100.50"`
	IdempotencyKey    string             `json:"idempotencyKey" validate:"required"`
	ProviderPayoutRef string             `json:"providerPayoutRef"`
	PayoutStatus      model.PayoutStatus `json:"payoutStatus" validate:"oneof=created queued processing succeeded failed reversed cancelled held" enums:"created,queued,processing,succeeded,failed,reversed,cancelled,held"`
	HoldReasonCode    string             `json:"holdReasonCode"`
	InitiatedAt       time.Time          `json:"initiatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CompletedAt       time.Time          `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailedAt          time.Time          `json:"failedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailureCode       string             `json:"failureCode"`
	FailureReason     string             `json:"failureReason"`
	Metadata          json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewPayoutsResponse(payouts model.Payouts) PayoutsResponse {
	return PayoutsResponse{
		Id:                payouts.Id,
		PayoutCode:        payouts.PayoutCode,
		PayoutBatchId:     payouts.PayoutBatchId.UUID,
		SettlementBatchId: payouts.SettlementBatchId.UUID,
		MerchantPartyId:   payouts.MerchantPartyId,
		PayoutMethodId:    payouts.PayoutMethodId,
		ProviderAccountId: payouts.ProviderAccountId.UUID,
		CurrencyCode:      payouts.CurrencyCode,
		Amount:            payouts.Amount,
		FeeAmount:         payouts.FeeAmount,
		NetSentAmount:     payouts.NetSentAmount,
		IdempotencyKey:    payouts.IdempotencyKey,
		ProviderPayoutRef: payouts.ProviderPayoutRef.String,
		PayoutStatus:      model.PayoutStatus(payouts.PayoutStatus),
		HoldReasonCode:    payouts.HoldReasonCode.String,
		InitiatedAt:       payouts.InitiatedAt.Time,
		CompletedAt:       payouts.CompletedAt.Time,
		FailedAt:          payouts.FailedAt.Time,
		FailureCode:       payouts.FailureCode.String,
		FailureReason:     payouts.FailureReason.String,
		Metadata:          payouts.Metadata,
	}
}

type PayoutsListResponse []*PayoutsResponse

func NewPayoutsListResponse(payoutsList model.PayoutsList) PayoutsListResponse {
	dtoPayoutsListResponse := PayoutsListResponse{}
	for _, payouts := range payoutsList {
		dtoPayoutsResponse := NewPayoutsResponse(*payouts)
		dtoPayoutsListResponse = append(dtoPayoutsListResponse, &dtoPayoutsResponse)
	}
	return dtoPayoutsListResponse
}

type PayoutsPrimaryIDList []PayoutsPrimaryID

func (d PayoutsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payouts := range d {
		err = validator.Struct(payouts)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutsPrimaryID) ToModel() model.PayoutsPrimaryID {
	return model.PayoutsPrimaryID{
		Id: d.Id,
	}
}

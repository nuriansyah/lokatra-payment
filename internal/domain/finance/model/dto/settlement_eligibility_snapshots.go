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

type SettlementEligibilitySnapshotsDTOFieldNameType string

type settlementEligibilitySnapshotsDTOFieldName struct {
	Id                        SettlementEligibilitySnapshotsDTOFieldNameType
	SourceType                SettlementEligibilitySnapshotsDTOFieldNameType
	SourceId                  SettlementEligibilitySnapshotsDTOFieldNameType
	MerchantPartyId           SettlementEligibilitySnapshotsDTOFieldNameType
	SettlementPolicyVersionId SettlementEligibilitySnapshotsDTOFieldNameType
	CurrencyCode              SettlementEligibilitySnapshotsDTOFieldNameType
	GrossAmount               SettlementEligibilitySnapshotsDTOFieldNameType
	FeeAmount                 SettlementEligibilitySnapshotsDTOFieldNameType
	TaxAmount                 SettlementEligibilitySnapshotsDTOFieldNameType
	ReserveAmount             SettlementEligibilitySnapshotsDTOFieldNameType
	NetSettleableAmount       SettlementEligibilitySnapshotsDTOFieldNameType
	EligibilityStatus         SettlementEligibilitySnapshotsDTOFieldNameType
	EligibleAt                SettlementEligibilitySnapshotsDTOFieldNameType
	SnapshotPayload           SettlementEligibilitySnapshotsDTOFieldNameType
	MetaCreatedAt             SettlementEligibilitySnapshotsDTOFieldNameType
	MetaCreatedBy             SettlementEligibilitySnapshotsDTOFieldNameType
	MetaUpdatedAt             SettlementEligibilitySnapshotsDTOFieldNameType
	MetaUpdatedBy             SettlementEligibilitySnapshotsDTOFieldNameType
	MetaDeletedAt             SettlementEligibilitySnapshotsDTOFieldNameType
	MetaDeletedBy             SettlementEligibilitySnapshotsDTOFieldNameType
}

var SettlementEligibilitySnapshotsDTOFieldName = settlementEligibilitySnapshotsDTOFieldName{
	Id:                        "id",
	SourceType:                "sourceType",
	SourceId:                  "sourceId",
	MerchantPartyId:           "merchantPartyId",
	SettlementPolicyVersionId: "settlementPolicyVersionId",
	CurrencyCode:              "currencyCode",
	GrossAmount:               "grossAmount",
	FeeAmount:                 "feeAmount",
	TaxAmount:                 "taxAmount",
	ReserveAmount:             "reserveAmount",
	NetSettleableAmount:       "netSettleableAmount",
	EligibilityStatus:         "eligibilityStatus",
	EligibleAt:                "eligibleAt",
	SnapshotPayload:           "snapshotPayload",
	MetaCreatedAt:             "metaCreatedAt",
	MetaCreatedBy:             "metaCreatedBy",
	MetaUpdatedAt:             "metaUpdatedAt",
	MetaUpdatedBy:             "metaUpdatedBy",
	MetaDeletedAt:             "metaDeletedAt",
	MetaDeletedBy:             "metaDeletedBy",
}

func transformSettlementEligibilitySnapshotsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(SettlementEligibilitySnapshotsDTOFieldName.Id):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.Id), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.SourceType):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.SourceType), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.SourceId):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.SourceId), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MerchantPartyId):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MerchantPartyId), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.SettlementPolicyVersionId):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.SettlementPolicyVersionId), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.CurrencyCode):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.CurrencyCode), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.GrossAmount):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.GrossAmount), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.FeeAmount):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.FeeAmount), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.TaxAmount):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.TaxAmount), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.ReserveAmount):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.ReserveAmount), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.NetSettleableAmount):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.NetSettleableAmount), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.EligibilityStatus):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.EligibilityStatus), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.EligibleAt):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.EligibleAt), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.SnapshotPayload):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.SnapshotPayload), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaCreatedAt):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedAt), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaCreatedBy):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedBy), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaUpdatedAt):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedAt), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaUpdatedBy):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedBy), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaDeletedAt):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedAt), true

	case string(SettlementEligibilitySnapshotsDTOFieldName.MetaDeletedBy):
		return string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isSettlementEligibilitySnapshotsBaseFilterField(field string) bool {
	spec, found := model.NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeSettlementEligibilitySnapshotsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateSettlementEligibilitySnapshotsProjectionOutputPath(path string) error {
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

func transformSettlementEligibilitySnapshotsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformSettlementEligibilitySnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformSettlementEligibilitySnapshotsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformSettlementEligibilitySnapshotsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformSettlementEligibilitySnapshotsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isSettlementEligibilitySnapshotsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateSettlementEligibilitySnapshotsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeSettlementEligibilitySnapshotsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformSettlementEligibilitySnapshotsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformSettlementEligibilitySnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformSettlementEligibilitySnapshotsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultSettlementEligibilitySnapshotsFilter(filter *model.Filter) {
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
			Field: string(SettlementEligibilitySnapshotsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type SettlementEligibilitySnapshotsSelectableResponse map[string]interface{}
type SettlementEligibilitySnapshotsSelectableListResponse []*SettlementEligibilitySnapshotsSelectableResponse

func assignSettlementEligibilitySnapshotsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setSettlementEligibilitySnapshotsSelectableValue(out SettlementEligibilitySnapshotsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignSettlementEligibilitySnapshotsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewSettlementEligibilitySnapshotsSelectableResponse(settlementEligibilitySnapshots model.SettlementEligibilitySnapshots, filter model.Filter) SettlementEligibilitySnapshotsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.SettlementEligibilitySnapshotsDBFieldName.Id),
			string(model.SettlementEligibilitySnapshotsDBFieldName.SourceType),
			string(model.SettlementEligibilitySnapshotsDBFieldName.SourceId),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MerchantPartyId),
			string(model.SettlementEligibilitySnapshotsDBFieldName.SettlementPolicyVersionId),
			string(model.SettlementEligibilitySnapshotsDBFieldName.CurrencyCode),
			string(model.SettlementEligibilitySnapshotsDBFieldName.GrossAmount),
			string(model.SettlementEligibilitySnapshotsDBFieldName.FeeAmount),
			string(model.SettlementEligibilitySnapshotsDBFieldName.TaxAmount),
			string(model.SettlementEligibilitySnapshotsDBFieldName.ReserveAmount),
			string(model.SettlementEligibilitySnapshotsDBFieldName.NetSettleableAmount),
			string(model.SettlementEligibilitySnapshotsDBFieldName.EligibilityStatus),
			string(model.SettlementEligibilitySnapshotsDBFieldName.EligibleAt),
			string(model.SettlementEligibilitySnapshotsDBFieldName.SnapshotPayload),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedAt),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedBy),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedAt),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedBy),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedAt),
			string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedBy),
		)
	}
	settlementEligibilitySnapshotsSelectableResponse := SettlementEligibilitySnapshotsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.SettlementEligibilitySnapshotsDBFieldName.Id):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.Id, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.SourceType):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.SourceType, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.SourceId):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.SourceId, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MerchantPartyId):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MerchantPartyId, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.SettlementPolicyVersionId):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.SettlementPolicyVersionId)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.SettlementPolicyVersionId, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.CurrencyCode):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.CurrencyCode, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.GrossAmount):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.GrossAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.GrossAmount, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.FeeAmount):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.FeeAmount, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.TaxAmount):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.TaxAmount, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.ReserveAmount):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.ReserveAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.ReserveAmount, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.NetSettleableAmount):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.NetSettleableAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.NetSettleableAmount, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.EligibilityStatus):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.EligibilityStatus)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, model.EligibilityStatus(settlementEligibilitySnapshots.EligibilityStatus), explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.EligibleAt):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.EligibleAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.EligibleAt.Time, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.SnapshotPayload):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.SnapshotPayload)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.SnapshotPayload, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedAt):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaCreatedAt, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaCreatedBy):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaCreatedBy, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedAt):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaUpdatedAt, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedBy):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaUpdatedBy, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedAt):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaDeletedAt.Time, explicitAlias)

		case string(model.SettlementEligibilitySnapshotsDBFieldName.MetaDeletedBy):
			key := string(SettlementEligibilitySnapshotsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementEligibilitySnapshotsSelectableValue(settlementEligibilitySnapshotsSelectableResponse, key, settlementEligibilitySnapshots.MetaDeletedBy, explicitAlias)

		}
	}
	return settlementEligibilitySnapshotsSelectableResponse
}

func NewSettlementEligibilitySnapshotsListResponseFromFilterResult(result []model.SettlementEligibilitySnapshotsFilterResult, filter model.Filter) SettlementEligibilitySnapshotsSelectableListResponse {
	dtoSettlementEligibilitySnapshotsListResponse := SettlementEligibilitySnapshotsSelectableListResponse{}
	for _, row := range result {
		dtoSettlementEligibilitySnapshotsResponse := NewSettlementEligibilitySnapshotsSelectableResponse(row.SettlementEligibilitySnapshots, filter)
		dtoSettlementEligibilitySnapshotsListResponse = append(dtoSettlementEligibilitySnapshotsListResponse, &dtoSettlementEligibilitySnapshotsResponse)
	}
	return dtoSettlementEligibilitySnapshotsListResponse
}

type SettlementEligibilitySnapshotsFilterResponse struct {
	Metadata Metadata                                             `json:"metadata"`
	Data     SettlementEligibilitySnapshotsSelectableListResponse `json:"data"`
}

func reverseSettlementEligibilitySnapshotsFilterResults(result []model.SettlementEligibilitySnapshotsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewSettlementEligibilitySnapshotsFilterResponse(result []model.SettlementEligibilitySnapshotsFilterResult, filter model.Filter) (resp SettlementEligibilitySnapshotsFilterResponse) {
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
			reverseSettlementEligibilitySnapshotsFilterResults(dataResult)
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

	resp.Data = NewSettlementEligibilitySnapshotsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type SettlementEligibilitySnapshotsCreateRequest struct {
	SourceType                string                  `json:"sourceType"`
	SourceId                  uuid.UUID               `json:"sourceId"`
	MerchantPartyId           uuid.UUID               `json:"merchantPartyId"`
	SettlementPolicyVersionId uuid.UUID               `json:"settlementPolicyVersionId"`
	CurrencyCode              string                  `json:"currencyCode"`
	GrossAmount               decimal.Decimal         `json:"grossAmount"`
	FeeAmount                 decimal.Decimal         `json:"feeAmount"`
	TaxAmount                 decimal.Decimal         `json:"taxAmount"`
	ReserveAmount             decimal.Decimal         `json:"reserveAmount"`
	NetSettleableAmount       decimal.Decimal         `json:"netSettleableAmount"`
	EligibilityStatus         model.EligibilityStatus `json:"eligibilityStatus" example:"pending" enums:"pending,eligible,held,blocked,settled"`
	EligibleAt                time.Time               `json:"eligibleAt"`
	SnapshotPayload           json.RawMessage         `json:"snapshotPayload"`
}

func (d *SettlementEligibilitySnapshotsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *SettlementEligibilitySnapshotsCreateRequest) ToModel() model.SettlementEligibilitySnapshots {
	id, _ := uuid.NewV4()
	return model.SettlementEligibilitySnapshots{
		Id:                        id,
		SourceType:                d.SourceType,
		SourceId:                  d.SourceId,
		MerchantPartyId:           d.MerchantPartyId,
		SettlementPolicyVersionId: d.SettlementPolicyVersionId,
		CurrencyCode:              d.CurrencyCode,
		GrossAmount:               d.GrossAmount,
		FeeAmount:                 d.FeeAmount,
		TaxAmount:                 d.TaxAmount,
		ReserveAmount:             d.ReserveAmount,
		NetSettleableAmount:       d.NetSettleableAmount,
		EligibilityStatus:         d.EligibilityStatus,
		EligibleAt:                null.TimeFrom(d.EligibleAt),
		SnapshotPayload:           d.SnapshotPayload,
	}
}

type SettlementEligibilitySnapshotsListCreateRequest []*SettlementEligibilitySnapshotsCreateRequest

func (d SettlementEligibilitySnapshotsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementEligibilitySnapshots := range d {
		err = validator.Struct(settlementEligibilitySnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementEligibilitySnapshotsListCreateRequest) ToModelList() []model.SettlementEligibilitySnapshots {
	out := make([]model.SettlementEligibilitySnapshots, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type SettlementEligibilitySnapshotsUpdateRequest struct {
	SourceType                string                  `json:"sourceType"`
	SourceId                  uuid.UUID               `json:"sourceId"`
	MerchantPartyId           uuid.UUID               `json:"merchantPartyId"`
	SettlementPolicyVersionId uuid.UUID               `json:"settlementPolicyVersionId"`
	CurrencyCode              string                  `json:"currencyCode"`
	GrossAmount               decimal.Decimal         `json:"grossAmount"`
	FeeAmount                 decimal.Decimal         `json:"feeAmount"`
	TaxAmount                 decimal.Decimal         `json:"taxAmount"`
	ReserveAmount             decimal.Decimal         `json:"reserveAmount"`
	NetSettleableAmount       decimal.Decimal         `json:"netSettleableAmount"`
	EligibilityStatus         model.EligibilityStatus `json:"eligibilityStatus" example:"pending" enums:"pending,eligible,held,blocked,settled"`
	EligibleAt                time.Time               `json:"eligibleAt"`
	SnapshotPayload           json.RawMessage         `json:"snapshotPayload"`
}

func (d *SettlementEligibilitySnapshotsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d SettlementEligibilitySnapshotsUpdateRequest) ToModel() model.SettlementEligibilitySnapshots {
	return model.SettlementEligibilitySnapshots{
		SourceType:                d.SourceType,
		SourceId:                  d.SourceId,
		MerchantPartyId:           d.MerchantPartyId,
		SettlementPolicyVersionId: d.SettlementPolicyVersionId,
		CurrencyCode:              d.CurrencyCode,
		GrossAmount:               d.GrossAmount,
		FeeAmount:                 d.FeeAmount,
		TaxAmount:                 d.TaxAmount,
		ReserveAmount:             d.ReserveAmount,
		NetSettleableAmount:       d.NetSettleableAmount,
		EligibilityStatus:         d.EligibilityStatus,
		EligibleAt:                null.TimeFrom(d.EligibleAt),
		SnapshotPayload:           d.SnapshotPayload,
	}
}

type SettlementEligibilitySnapshotsBulkUpdateRequest struct {
	Id                        uuid.UUID               `json:"id"`
	SourceType                string                  `json:"sourceType"`
	SourceId                  uuid.UUID               `json:"sourceId"`
	MerchantPartyId           uuid.UUID               `json:"merchantPartyId"`
	SettlementPolicyVersionId uuid.UUID               `json:"settlementPolicyVersionId"`
	CurrencyCode              string                  `json:"currencyCode"`
	GrossAmount               decimal.Decimal         `json:"grossAmount"`
	FeeAmount                 decimal.Decimal         `json:"feeAmount"`
	TaxAmount                 decimal.Decimal         `json:"taxAmount"`
	ReserveAmount             decimal.Decimal         `json:"reserveAmount"`
	NetSettleableAmount       decimal.Decimal         `json:"netSettleableAmount"`
	EligibilityStatus         model.EligibilityStatus `json:"eligibilityStatus" example:"pending" enums:"pending,eligible,held,blocked,settled"`
	EligibleAt                time.Time               `json:"eligibleAt"`
	SnapshotPayload           json.RawMessage         `json:"snapshotPayload"`
}

func (d SettlementEligibilitySnapshotsBulkUpdateRequest) PrimaryID() SettlementEligibilitySnapshotsPrimaryID {
	return SettlementEligibilitySnapshotsPrimaryID{
		Id: d.Id,
	}
}

type SettlementEligibilitySnapshotsListBulkUpdateRequest []*SettlementEligibilitySnapshotsBulkUpdateRequest

func (d SettlementEligibilitySnapshotsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementEligibilitySnapshots := range d {
		err = validator.Struct(settlementEligibilitySnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementEligibilitySnapshotsBulkUpdateRequest) ToModel() model.SettlementEligibilitySnapshots {
	return model.SettlementEligibilitySnapshots{
		Id:                        d.Id,
		SourceType:                d.SourceType,
		SourceId:                  d.SourceId,
		MerchantPartyId:           d.MerchantPartyId,
		SettlementPolicyVersionId: d.SettlementPolicyVersionId,
		CurrencyCode:              d.CurrencyCode,
		GrossAmount:               d.GrossAmount,
		FeeAmount:                 d.FeeAmount,
		TaxAmount:                 d.TaxAmount,
		ReserveAmount:             d.ReserveAmount,
		NetSettleableAmount:       d.NetSettleableAmount,
		EligibilityStatus:         d.EligibilityStatus,
		EligibleAt:                null.TimeFrom(d.EligibleAt),
		SnapshotPayload:           d.SnapshotPayload,
	}
}

type SettlementEligibilitySnapshotsResponse struct {
	Id                        uuid.UUID               `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType                string                  `json:"sourceType" validate:"required"`
	SourceId                  uuid.UUID               `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId           uuid.UUID               `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SettlementPolicyVersionId uuid.UUID               `json:"settlementPolicyVersionId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode              string                  `json:"currencyCode" validate:"required"`
	GrossAmount               decimal.Decimal         `json:"grossAmount" validate:"required" format:"decimal" example:"100.50"`
	FeeAmount                 decimal.Decimal         `json:"feeAmount" format:"decimal" example:"100.50"`
	TaxAmount                 decimal.Decimal         `json:"taxAmount" format:"decimal" example:"100.50"`
	ReserveAmount             decimal.Decimal         `json:"reserveAmount" format:"decimal" example:"100.50"`
	NetSettleableAmount       decimal.Decimal         `json:"netSettleableAmount" validate:"required" format:"decimal" example:"100.50"`
	EligibilityStatus         model.EligibilityStatus `json:"eligibilityStatus" validate:"required,oneof=pending eligible held blocked settled" enums:"pending,eligible,held,blocked,settled"`
	EligibleAt                time.Time               `json:"eligibleAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	SnapshotPayload           json.RawMessage         `json:"snapshotPayload" swaggertype:"object"`
}

func NewSettlementEligibilitySnapshotsResponse(settlementEligibilitySnapshots model.SettlementEligibilitySnapshots) SettlementEligibilitySnapshotsResponse {
	return SettlementEligibilitySnapshotsResponse{
		Id:                        settlementEligibilitySnapshots.Id,
		SourceType:                settlementEligibilitySnapshots.SourceType,
		SourceId:                  settlementEligibilitySnapshots.SourceId,
		MerchantPartyId:           settlementEligibilitySnapshots.MerchantPartyId,
		SettlementPolicyVersionId: settlementEligibilitySnapshots.SettlementPolicyVersionId,
		CurrencyCode:              settlementEligibilitySnapshots.CurrencyCode,
		GrossAmount:               settlementEligibilitySnapshots.GrossAmount,
		FeeAmount:                 settlementEligibilitySnapshots.FeeAmount,
		TaxAmount:                 settlementEligibilitySnapshots.TaxAmount,
		ReserveAmount:             settlementEligibilitySnapshots.ReserveAmount,
		NetSettleableAmount:       settlementEligibilitySnapshots.NetSettleableAmount,
		EligibilityStatus:         model.EligibilityStatus(settlementEligibilitySnapshots.EligibilityStatus),
		EligibleAt:                settlementEligibilitySnapshots.EligibleAt.Time,
		SnapshotPayload:           settlementEligibilitySnapshots.SnapshotPayload,
	}
}

type SettlementEligibilitySnapshotsListResponse []*SettlementEligibilitySnapshotsResponse

func NewSettlementEligibilitySnapshotsListResponse(settlementEligibilitySnapshotsList model.SettlementEligibilitySnapshotsList) SettlementEligibilitySnapshotsListResponse {
	dtoSettlementEligibilitySnapshotsListResponse := SettlementEligibilitySnapshotsListResponse{}
	for _, settlementEligibilitySnapshots := range settlementEligibilitySnapshotsList {
		dtoSettlementEligibilitySnapshotsResponse := NewSettlementEligibilitySnapshotsResponse(*settlementEligibilitySnapshots)
		dtoSettlementEligibilitySnapshotsListResponse = append(dtoSettlementEligibilitySnapshotsListResponse, &dtoSettlementEligibilitySnapshotsResponse)
	}
	return dtoSettlementEligibilitySnapshotsListResponse
}

type SettlementEligibilitySnapshotsPrimaryIDList []SettlementEligibilitySnapshotsPrimaryID

func (d SettlementEligibilitySnapshotsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementEligibilitySnapshots := range d {
		err = validator.Struct(settlementEligibilitySnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

type SettlementEligibilitySnapshotsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *SettlementEligibilitySnapshotsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d SettlementEligibilitySnapshotsPrimaryID) ToModel() model.SettlementEligibilitySnapshotsPrimaryID {
	return model.SettlementEligibilitySnapshotsPrimaryID{
		Id: d.Id,
	}
}

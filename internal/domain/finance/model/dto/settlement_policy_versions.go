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

type SettlementPolicyVersionsDTOFieldNameType string

type settlementPolicyVersionsDTOFieldName struct {
	Id                 SettlementPolicyVersionsDTOFieldNameType
	SettlementPolicyId SettlementPolicyVersionsDTOFieldNameType
	VersionNo          SettlementPolicyVersionsDTOFieldNameType
	TriggerType        SettlementPolicyVersionsDTOFieldNameType
	DelayDays          SettlementPolicyVersionsDTOFieldNameType
	MinPayoutAmount    SettlementPolicyVersionsDTOFieldNameType
	PayoutFrequency    SettlementPolicyVersionsDTOFieldNameType
	ReservePolicyId    SettlementPolicyVersionsDTOFieldNameType
	AutoRelease        SettlementPolicyVersionsDTOFieldNameType
	Conditions         SettlementPolicyVersionsDTOFieldNameType
	IsCurrent          SettlementPolicyVersionsDTOFieldNameType
	MetaCreatedAt      SettlementPolicyVersionsDTOFieldNameType
	MetaCreatedBy      SettlementPolicyVersionsDTOFieldNameType
	MetaUpdatedAt      SettlementPolicyVersionsDTOFieldNameType
	MetaUpdatedBy      SettlementPolicyVersionsDTOFieldNameType
	MetaDeletedAt      SettlementPolicyVersionsDTOFieldNameType
	MetaDeletedBy      SettlementPolicyVersionsDTOFieldNameType
}

var SettlementPolicyVersionsDTOFieldName = settlementPolicyVersionsDTOFieldName{
	Id:                 "id",
	SettlementPolicyId: "settlementPolicyId",
	VersionNo:          "versionNo",
	TriggerType:        "triggerType",
	DelayDays:          "delayDays",
	MinPayoutAmount:    "minPayoutAmount",
	PayoutFrequency:    "payoutFrequency",
	ReservePolicyId:    "reservePolicyId",
	AutoRelease:        "autoRelease",
	Conditions:         "conditions",
	IsCurrent:          "isCurrent",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformSettlementPolicyVersionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(SettlementPolicyVersionsDTOFieldName.Id):
		return string(model.SettlementPolicyVersionsDBFieldName.Id), true

	case string(SettlementPolicyVersionsDTOFieldName.SettlementPolicyId):
		return string(model.SettlementPolicyVersionsDBFieldName.SettlementPolicyId), true

	case string(SettlementPolicyVersionsDTOFieldName.VersionNo):
		return string(model.SettlementPolicyVersionsDBFieldName.VersionNo), true

	case string(SettlementPolicyVersionsDTOFieldName.TriggerType):
		return string(model.SettlementPolicyVersionsDBFieldName.TriggerType), true

	case string(SettlementPolicyVersionsDTOFieldName.DelayDays):
		return string(model.SettlementPolicyVersionsDBFieldName.DelayDays), true

	case string(SettlementPolicyVersionsDTOFieldName.MinPayoutAmount):
		return string(model.SettlementPolicyVersionsDBFieldName.MinPayoutAmount), true

	case string(SettlementPolicyVersionsDTOFieldName.PayoutFrequency):
		return string(model.SettlementPolicyVersionsDBFieldName.PayoutFrequency), true

	case string(SettlementPolicyVersionsDTOFieldName.ReservePolicyId):
		return string(model.SettlementPolicyVersionsDBFieldName.ReservePolicyId), true

	case string(SettlementPolicyVersionsDTOFieldName.AutoRelease):
		return string(model.SettlementPolicyVersionsDBFieldName.AutoRelease), true

	case string(SettlementPolicyVersionsDTOFieldName.Conditions):
		return string(model.SettlementPolicyVersionsDBFieldName.Conditions), true

	case string(SettlementPolicyVersionsDTOFieldName.IsCurrent):
		return string(model.SettlementPolicyVersionsDBFieldName.IsCurrent), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaCreatedAt):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedAt), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaCreatedBy):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedBy), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaUpdatedAt):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedAt), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaUpdatedBy):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedBy), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaDeletedAt):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedAt), true

	case string(SettlementPolicyVersionsDTOFieldName.MetaDeletedBy):
		return string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewSettlementPolicyVersionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isSettlementPolicyVersionsBaseFilterField(field string) bool {
	spec, found := model.NewSettlementPolicyVersionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeSettlementPolicyVersionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateSettlementPolicyVersionsProjectionOutputPath(path string) error {
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

func transformSettlementPolicyVersionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformSettlementPolicyVersionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformSettlementPolicyVersionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformSettlementPolicyVersionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformSettlementPolicyVersionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isSettlementPolicyVersionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateSettlementPolicyVersionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeSettlementPolicyVersionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformSettlementPolicyVersionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformSettlementPolicyVersionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformSettlementPolicyVersionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultSettlementPolicyVersionsFilter(filter *model.Filter) {
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
			Field: string(SettlementPolicyVersionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type SettlementPolicyVersionsSelectableResponse map[string]interface{}
type SettlementPolicyVersionsSelectableListResponse []*SettlementPolicyVersionsSelectableResponse

func assignSettlementPolicyVersionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setSettlementPolicyVersionsSelectableValue(out SettlementPolicyVersionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignSettlementPolicyVersionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewSettlementPolicyVersionsSelectableResponse(settlementPolicyVersions model.SettlementPolicyVersions, filter model.Filter) SettlementPolicyVersionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.SettlementPolicyVersionsDBFieldName.Id),
			string(model.SettlementPolicyVersionsDBFieldName.SettlementPolicyId),
			string(model.SettlementPolicyVersionsDBFieldName.VersionNo),
			string(model.SettlementPolicyVersionsDBFieldName.TriggerType),
			string(model.SettlementPolicyVersionsDBFieldName.DelayDays),
			string(model.SettlementPolicyVersionsDBFieldName.MinPayoutAmount),
			string(model.SettlementPolicyVersionsDBFieldName.PayoutFrequency),
			string(model.SettlementPolicyVersionsDBFieldName.ReservePolicyId),
			string(model.SettlementPolicyVersionsDBFieldName.AutoRelease),
			string(model.SettlementPolicyVersionsDBFieldName.Conditions),
			string(model.SettlementPolicyVersionsDBFieldName.IsCurrent),
			string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedAt),
			string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedBy),
			string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedAt),
			string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedBy),
			string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedAt),
			string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedBy),
		)
	}
	settlementPolicyVersionsSelectableResponse := SettlementPolicyVersionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.SettlementPolicyVersionsDBFieldName.Id):
			key := string(SettlementPolicyVersionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.Id, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.SettlementPolicyId):
			key := string(SettlementPolicyVersionsDTOFieldName.SettlementPolicyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.SettlementPolicyId, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.VersionNo):
			key := string(SettlementPolicyVersionsDTOFieldName.VersionNo)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.VersionNo, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.TriggerType):
			key := string(SettlementPolicyVersionsDTOFieldName.TriggerType)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, model.TriggerType(settlementPolicyVersions.TriggerType), explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.DelayDays):
			key := string(SettlementPolicyVersionsDTOFieldName.DelayDays)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.DelayDays, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MinPayoutAmount):
			key := string(SettlementPolicyVersionsDTOFieldName.MinPayoutAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MinPayoutAmount, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.PayoutFrequency):
			key := string(SettlementPolicyVersionsDTOFieldName.PayoutFrequency)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, model.PayoutFrequency(settlementPolicyVersions.PayoutFrequency), explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.ReservePolicyId):
			key := string(SettlementPolicyVersionsDTOFieldName.ReservePolicyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.ReservePolicyId.UUID, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.AutoRelease):
			key := string(SettlementPolicyVersionsDTOFieldName.AutoRelease)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.AutoRelease, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.Conditions):
			key := string(SettlementPolicyVersionsDTOFieldName.Conditions)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.Conditions, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.IsCurrent):
			key := string(SettlementPolicyVersionsDTOFieldName.IsCurrent)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.IsCurrent, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedAt):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaCreatedAt, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaCreatedBy):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaCreatedBy, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedAt):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaUpdatedAt, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaUpdatedBy):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaUpdatedBy, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedAt):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaDeletedAt.Time, explicitAlias)

		case string(model.SettlementPolicyVersionsDBFieldName.MetaDeletedBy):
			key := string(SettlementPolicyVersionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementPolicyVersionsSelectableValue(settlementPolicyVersionsSelectableResponse, key, settlementPolicyVersions.MetaDeletedBy, explicitAlias)

		}
	}
	return settlementPolicyVersionsSelectableResponse
}

func NewSettlementPolicyVersionsListResponseFromFilterResult(result []model.SettlementPolicyVersionsFilterResult, filter model.Filter) SettlementPolicyVersionsSelectableListResponse {
	dtoSettlementPolicyVersionsListResponse := SettlementPolicyVersionsSelectableListResponse{}
	for _, row := range result {
		dtoSettlementPolicyVersionsResponse := NewSettlementPolicyVersionsSelectableResponse(row.SettlementPolicyVersions, filter)
		dtoSettlementPolicyVersionsListResponse = append(dtoSettlementPolicyVersionsListResponse, &dtoSettlementPolicyVersionsResponse)
	}
	return dtoSettlementPolicyVersionsListResponse
}

type SettlementPolicyVersionsFilterResponse struct {
	Metadata Metadata                                       `json:"metadata"`
	Data     SettlementPolicyVersionsSelectableListResponse `json:"data"`
}

func reverseSettlementPolicyVersionsFilterResults(result []model.SettlementPolicyVersionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewSettlementPolicyVersionsFilterResponse(result []model.SettlementPolicyVersionsFilterResult, filter model.Filter) (resp SettlementPolicyVersionsFilterResponse) {
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
			reverseSettlementPolicyVersionsFilterResults(dataResult)
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

	resp.Data = NewSettlementPolicyVersionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type SettlementPolicyVersionsCreateRequest struct {
	SettlementPolicyId uuid.UUID             `json:"settlementPolicyId"`
	VersionNo          int                   `json:"versionNo"`
	TriggerType        model.TriggerType     `json:"triggerType" example:"payment_settled" enums:"payment_settled,visit_completed,service_started,service_completed,manual"`
	DelayDays          int                   `json:"delayDays"`
	MinPayoutAmount    decimal.Decimal       `json:"minPayoutAmount"`
	PayoutFrequency    model.PayoutFrequency `json:"payoutFrequency" example:"daily" enums:"daily,weekly,biweekly,monthly,manual"`
	ReservePolicyId    uuid.UUID             `json:"reservePolicyId"`
	AutoRelease        bool                  `json:"autoRelease"`
	Conditions         json.RawMessage       `json:"conditions"`
	IsCurrent          bool                  `json:"isCurrent"`
}

func (d *SettlementPolicyVersionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *SettlementPolicyVersionsCreateRequest) ToModel() model.SettlementPolicyVersions {
	id, _ := uuid.NewV4()
	return model.SettlementPolicyVersions{
		Id:                 id,
		SettlementPolicyId: d.SettlementPolicyId,
		VersionNo:          d.VersionNo,
		TriggerType:        d.TriggerType,
		DelayDays:          d.DelayDays,
		MinPayoutAmount:    d.MinPayoutAmount,
		PayoutFrequency:    d.PayoutFrequency,
		ReservePolicyId:    nuuid.From(d.ReservePolicyId),
		AutoRelease:        d.AutoRelease,
		Conditions:         d.Conditions,
		IsCurrent:          d.IsCurrent,
	}
}

type SettlementPolicyVersionsListCreateRequest []*SettlementPolicyVersionsCreateRequest

func (d SettlementPolicyVersionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicyVersions := range d {
		err = validator.Struct(settlementPolicyVersions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementPolicyVersionsListCreateRequest) ToModelList() []model.SettlementPolicyVersions {
	out := make([]model.SettlementPolicyVersions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type SettlementPolicyVersionsUpdateRequest struct {
	SettlementPolicyId uuid.UUID             `json:"settlementPolicyId"`
	VersionNo          int                   `json:"versionNo"`
	TriggerType        model.TriggerType     `json:"triggerType" example:"payment_settled" enums:"payment_settled,visit_completed,service_started,service_completed,manual"`
	DelayDays          int                   `json:"delayDays"`
	MinPayoutAmount    decimal.Decimal       `json:"minPayoutAmount"`
	PayoutFrequency    model.PayoutFrequency `json:"payoutFrequency" example:"daily" enums:"daily,weekly,biweekly,monthly,manual"`
	ReservePolicyId    uuid.UUID             `json:"reservePolicyId"`
	AutoRelease        bool                  `json:"autoRelease"`
	Conditions         json.RawMessage       `json:"conditions"`
	IsCurrent          bool                  `json:"isCurrent"`
}

func (d *SettlementPolicyVersionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d SettlementPolicyVersionsUpdateRequest) ToModel() model.SettlementPolicyVersions {
	return model.SettlementPolicyVersions{
		SettlementPolicyId: d.SettlementPolicyId,
		VersionNo:          d.VersionNo,
		TriggerType:        d.TriggerType,
		DelayDays:          d.DelayDays,
		MinPayoutAmount:    d.MinPayoutAmount,
		PayoutFrequency:    d.PayoutFrequency,
		ReservePolicyId:    nuuid.From(d.ReservePolicyId),
		AutoRelease:        d.AutoRelease,
		Conditions:         d.Conditions,
		IsCurrent:          d.IsCurrent,
	}
}

type SettlementPolicyVersionsBulkUpdateRequest struct {
	Id                 uuid.UUID             `json:"id"`
	SettlementPolicyId uuid.UUID             `json:"settlementPolicyId"`
	VersionNo          int                   `json:"versionNo"`
	TriggerType        model.TriggerType     `json:"triggerType" example:"payment_settled" enums:"payment_settled,visit_completed,service_started,service_completed,manual"`
	DelayDays          int                   `json:"delayDays"`
	MinPayoutAmount    decimal.Decimal       `json:"minPayoutAmount"`
	PayoutFrequency    model.PayoutFrequency `json:"payoutFrequency" example:"daily" enums:"daily,weekly,biweekly,monthly,manual"`
	ReservePolicyId    uuid.UUID             `json:"reservePolicyId"`
	AutoRelease        bool                  `json:"autoRelease"`
	Conditions         json.RawMessage       `json:"conditions"`
	IsCurrent          bool                  `json:"isCurrent"`
}

func (d SettlementPolicyVersionsBulkUpdateRequest) PrimaryID() SettlementPolicyVersionsPrimaryID {
	return SettlementPolicyVersionsPrimaryID{
		Id: d.Id,
	}
}

type SettlementPolicyVersionsListBulkUpdateRequest []*SettlementPolicyVersionsBulkUpdateRequest

func (d SettlementPolicyVersionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicyVersions := range d {
		err = validator.Struct(settlementPolicyVersions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementPolicyVersionsBulkUpdateRequest) ToModel() model.SettlementPolicyVersions {
	return model.SettlementPolicyVersions{
		Id:                 d.Id,
		SettlementPolicyId: d.SettlementPolicyId,
		VersionNo:          d.VersionNo,
		TriggerType:        d.TriggerType,
		DelayDays:          d.DelayDays,
		MinPayoutAmount:    d.MinPayoutAmount,
		PayoutFrequency:    d.PayoutFrequency,
		ReservePolicyId:    nuuid.From(d.ReservePolicyId),
		AutoRelease:        d.AutoRelease,
		Conditions:         d.Conditions,
		IsCurrent:          d.IsCurrent,
	}
}

type SettlementPolicyVersionsResponse struct {
	Id                 uuid.UUID             `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SettlementPolicyId uuid.UUID             `json:"settlementPolicyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	VersionNo          int                   `json:"versionNo" validate:"required" example:"1"`
	TriggerType        model.TriggerType     `json:"triggerType" validate:"required,oneof=payment_settled visit_completed service_started service_completed manual" enums:"payment_settled,visit_completed,service_started,service_completed,manual"`
	DelayDays          int                   `json:"delayDays" example:"1"`
	MinPayoutAmount    decimal.Decimal       `json:"minPayoutAmount" format:"decimal" example:"100.50"`
	PayoutFrequency    model.PayoutFrequency `json:"payoutFrequency" validate:"required,oneof=daily weekly biweekly monthly manual" enums:"daily,weekly,biweekly,monthly,manual"`
	ReservePolicyId    uuid.UUID             `json:"reservePolicyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AutoRelease        bool                  `json:"autoRelease" example:"true"`
	Conditions         json.RawMessage       `json:"conditions" swaggertype:"object"`
	IsCurrent          bool                  `json:"isCurrent" example:"true"`
}

func NewSettlementPolicyVersionsResponse(settlementPolicyVersions model.SettlementPolicyVersions) SettlementPolicyVersionsResponse {
	return SettlementPolicyVersionsResponse{
		Id:                 settlementPolicyVersions.Id,
		SettlementPolicyId: settlementPolicyVersions.SettlementPolicyId,
		VersionNo:          settlementPolicyVersions.VersionNo,
		TriggerType:        model.TriggerType(settlementPolicyVersions.TriggerType),
		DelayDays:          settlementPolicyVersions.DelayDays,
		MinPayoutAmount:    settlementPolicyVersions.MinPayoutAmount,
		PayoutFrequency:    model.PayoutFrequency(settlementPolicyVersions.PayoutFrequency),
		ReservePolicyId:    settlementPolicyVersions.ReservePolicyId.UUID,
		AutoRelease:        settlementPolicyVersions.AutoRelease,
		Conditions:         settlementPolicyVersions.Conditions,
		IsCurrent:          settlementPolicyVersions.IsCurrent,
	}
}

type SettlementPolicyVersionsListResponse []*SettlementPolicyVersionsResponse

func NewSettlementPolicyVersionsListResponse(settlementPolicyVersionsList model.SettlementPolicyVersionsList) SettlementPolicyVersionsListResponse {
	dtoSettlementPolicyVersionsListResponse := SettlementPolicyVersionsListResponse{}
	for _, settlementPolicyVersions := range settlementPolicyVersionsList {
		dtoSettlementPolicyVersionsResponse := NewSettlementPolicyVersionsResponse(*settlementPolicyVersions)
		dtoSettlementPolicyVersionsListResponse = append(dtoSettlementPolicyVersionsListResponse, &dtoSettlementPolicyVersionsResponse)
	}
	return dtoSettlementPolicyVersionsListResponse
}

type SettlementPolicyVersionsPrimaryIDList []SettlementPolicyVersionsPrimaryID

func (d SettlementPolicyVersionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementPolicyVersions := range d {
		err = validator.Struct(settlementPolicyVersions)
		if err != nil {
			return
		}
	}
	return nil
}

type SettlementPolicyVersionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *SettlementPolicyVersionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d SettlementPolicyVersionsPrimaryID) ToModel() model.SettlementPolicyVersionsPrimaryID {
	return model.SettlementPolicyVersionsPrimaryID{
		Id: d.Id,
	}
}

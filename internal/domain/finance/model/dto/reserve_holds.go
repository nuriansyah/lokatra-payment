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

type ReserveHoldsDTOFieldNameType string

type reserveHoldsDTOFieldName struct {
	Id              ReserveHoldsDTOFieldNameType
	MerchantPartyId ReserveHoldsDTOFieldNameType
	CurrencyCode    ReserveHoldsDTOFieldNameType
	SourceType      ReserveHoldsDTOFieldNameType
	SourceId        ReserveHoldsDTOFieldNameType
	HoldType        ReserveHoldsDTOFieldNameType
	HoldStatus      ReserveHoldsDTOFieldNameType
	HoldAmount      ReserveHoldsDTOFieldNameType
	ReleasedAmount  ReserveHoldsDTOFieldNameType
	EffectiveAt     ReserveHoldsDTOFieldNameType
	ReleaseAt       ReserveHoldsDTOFieldNameType
	ReleasedAt      ReserveHoldsDTOFieldNameType
	ReasonCode      ReserveHoldsDTOFieldNameType
	ReasonDetail    ReserveHoldsDTOFieldNameType
	Metadata        ReserveHoldsDTOFieldNameType
	MetaCreatedAt   ReserveHoldsDTOFieldNameType
	MetaCreatedBy   ReserveHoldsDTOFieldNameType
	MetaUpdatedAt   ReserveHoldsDTOFieldNameType
	MetaUpdatedBy   ReserveHoldsDTOFieldNameType
	MetaDeletedAt   ReserveHoldsDTOFieldNameType
	MetaDeletedBy   ReserveHoldsDTOFieldNameType
}

var ReserveHoldsDTOFieldName = reserveHoldsDTOFieldName{
	Id:              "id",
	MerchantPartyId: "merchantPartyId",
	CurrencyCode:    "currencyCode",
	SourceType:      "sourceType",
	SourceId:        "sourceId",
	HoldType:        "holdType",
	HoldStatus:      "holdStatus",
	HoldAmount:      "holdAmount",
	ReleasedAmount:  "releasedAmount",
	EffectiveAt:     "effectiveAt",
	ReleaseAt:       "releaseAt",
	ReleasedAt:      "releasedAt",
	ReasonCode:      "reasonCode",
	ReasonDetail:    "reasonDetail",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformReserveHoldsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReserveHoldsDTOFieldName.Id):
		return string(model.ReserveHoldsDBFieldName.Id), true

	case string(ReserveHoldsDTOFieldName.MerchantPartyId):
		return string(model.ReserveHoldsDBFieldName.MerchantPartyId), true

	case string(ReserveHoldsDTOFieldName.CurrencyCode):
		return string(model.ReserveHoldsDBFieldName.CurrencyCode), true

	case string(ReserveHoldsDTOFieldName.SourceType):
		return string(model.ReserveHoldsDBFieldName.SourceType), true

	case string(ReserveHoldsDTOFieldName.SourceId):
		return string(model.ReserveHoldsDBFieldName.SourceId), true

	case string(ReserveHoldsDTOFieldName.HoldType):
		return string(model.ReserveHoldsDBFieldName.HoldType), true

	case string(ReserveHoldsDTOFieldName.HoldStatus):
		return string(model.ReserveHoldsDBFieldName.HoldStatus), true

	case string(ReserveHoldsDTOFieldName.HoldAmount):
		return string(model.ReserveHoldsDBFieldName.HoldAmount), true

	case string(ReserveHoldsDTOFieldName.ReleasedAmount):
		return string(model.ReserveHoldsDBFieldName.ReleasedAmount), true

	case string(ReserveHoldsDTOFieldName.EffectiveAt):
		return string(model.ReserveHoldsDBFieldName.EffectiveAt), true

	case string(ReserveHoldsDTOFieldName.ReleaseAt):
		return string(model.ReserveHoldsDBFieldName.ReleaseAt), true

	case string(ReserveHoldsDTOFieldName.ReleasedAt):
		return string(model.ReserveHoldsDBFieldName.ReleasedAt), true

	case string(ReserveHoldsDTOFieldName.ReasonCode):
		return string(model.ReserveHoldsDBFieldName.ReasonCode), true

	case string(ReserveHoldsDTOFieldName.ReasonDetail):
		return string(model.ReserveHoldsDBFieldName.ReasonDetail), true

	case string(ReserveHoldsDTOFieldName.Metadata):
		return string(model.ReserveHoldsDBFieldName.Metadata), true

	case string(ReserveHoldsDTOFieldName.MetaCreatedAt):
		return string(model.ReserveHoldsDBFieldName.MetaCreatedAt), true

	case string(ReserveHoldsDTOFieldName.MetaCreatedBy):
		return string(model.ReserveHoldsDBFieldName.MetaCreatedBy), true

	case string(ReserveHoldsDTOFieldName.MetaUpdatedAt):
		return string(model.ReserveHoldsDBFieldName.MetaUpdatedAt), true

	case string(ReserveHoldsDTOFieldName.MetaUpdatedBy):
		return string(model.ReserveHoldsDBFieldName.MetaUpdatedBy), true

	case string(ReserveHoldsDTOFieldName.MetaDeletedAt):
		return string(model.ReserveHoldsDBFieldName.MetaDeletedAt), true

	case string(ReserveHoldsDTOFieldName.MetaDeletedBy):
		return string(model.ReserveHoldsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReserveHoldsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReserveHoldsBaseFilterField(field string) bool {
	spec, found := model.NewReserveHoldsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReserveHoldsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReserveHoldsProjectionOutputPath(path string) error {
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

func transformReserveHoldsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReserveHoldsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReserveHoldsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReserveHoldsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReserveHoldsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReserveHoldsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReserveHoldsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReserveHoldsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReserveHoldsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReserveHoldsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReserveHoldsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReserveHoldsFilter(filter *model.Filter) {
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
			Field: string(ReserveHoldsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReserveHoldsSelectableResponse map[string]interface{}
type ReserveHoldsSelectableListResponse []*ReserveHoldsSelectableResponse

func assignReserveHoldsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReserveHoldsSelectableValue(out ReserveHoldsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReserveHoldsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReserveHoldsSelectableResponse(reserveHolds model.ReserveHolds, filter model.Filter) ReserveHoldsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReserveHoldsDBFieldName.Id),
			string(model.ReserveHoldsDBFieldName.MerchantPartyId),
			string(model.ReserveHoldsDBFieldName.CurrencyCode),
			string(model.ReserveHoldsDBFieldName.SourceType),
			string(model.ReserveHoldsDBFieldName.SourceId),
			string(model.ReserveHoldsDBFieldName.HoldType),
			string(model.ReserveHoldsDBFieldName.HoldStatus),
			string(model.ReserveHoldsDBFieldName.HoldAmount),
			string(model.ReserveHoldsDBFieldName.ReleasedAmount),
			string(model.ReserveHoldsDBFieldName.EffectiveAt),
			string(model.ReserveHoldsDBFieldName.ReleaseAt),
			string(model.ReserveHoldsDBFieldName.ReleasedAt),
			string(model.ReserveHoldsDBFieldName.ReasonCode),
			string(model.ReserveHoldsDBFieldName.ReasonDetail),
			string(model.ReserveHoldsDBFieldName.Metadata),
			string(model.ReserveHoldsDBFieldName.MetaCreatedAt),
			string(model.ReserveHoldsDBFieldName.MetaCreatedBy),
			string(model.ReserveHoldsDBFieldName.MetaUpdatedAt),
			string(model.ReserveHoldsDBFieldName.MetaUpdatedBy),
			string(model.ReserveHoldsDBFieldName.MetaDeletedAt),
			string(model.ReserveHoldsDBFieldName.MetaDeletedBy),
		)
	}
	reserveHoldsSelectableResponse := ReserveHoldsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReserveHoldsDBFieldName.Id):
			key := string(ReserveHoldsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.Id, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MerchantPartyId):
			key := string(ReserveHoldsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MerchantPartyId, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.CurrencyCode):
			key := string(ReserveHoldsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.CurrencyCode, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.SourceType):
			key := string(ReserveHoldsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.SourceType, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.SourceId):
			key := string(ReserveHoldsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.SourceId, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.HoldType):
			key := string(ReserveHoldsDTOFieldName.HoldType)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.HoldType, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.HoldStatus):
			key := string(ReserveHoldsDTOFieldName.HoldStatus)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, model.HoldStatus(reserveHolds.HoldStatus), explicitAlias)

		case string(model.ReserveHoldsDBFieldName.HoldAmount):
			key := string(ReserveHoldsDTOFieldName.HoldAmount)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.HoldAmount, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.ReleasedAmount):
			key := string(ReserveHoldsDTOFieldName.ReleasedAmount)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.ReleasedAmount, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.EffectiveAt):
			key := string(ReserveHoldsDTOFieldName.EffectiveAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.EffectiveAt, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.ReleaseAt):
			key := string(ReserveHoldsDTOFieldName.ReleaseAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.ReleaseAt.Time, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.ReleasedAt):
			key := string(ReserveHoldsDTOFieldName.ReleasedAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.ReleasedAt.Time, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.ReasonCode):
			key := string(ReserveHoldsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.ReasonCode, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.ReasonDetail):
			key := string(ReserveHoldsDTOFieldName.ReasonDetail)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.ReasonDetail.String, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.Metadata):
			key := string(ReserveHoldsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.Metadata, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaCreatedAt):
			key := string(ReserveHoldsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaCreatedAt, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaCreatedBy):
			key := string(ReserveHoldsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaCreatedBy, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaUpdatedAt):
			key := string(ReserveHoldsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaUpdatedAt, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaUpdatedBy):
			key := string(ReserveHoldsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaUpdatedBy, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaDeletedAt):
			key := string(ReserveHoldsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReserveHoldsDBFieldName.MetaDeletedBy):
			key := string(ReserveHoldsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReserveHoldsSelectableValue(reserveHoldsSelectableResponse, key, reserveHolds.MetaDeletedBy, explicitAlias)

		}
	}
	return reserveHoldsSelectableResponse
}

func NewReserveHoldsListResponseFromFilterResult(result []model.ReserveHoldsFilterResult, filter model.Filter) ReserveHoldsSelectableListResponse {
	dtoReserveHoldsListResponse := ReserveHoldsSelectableListResponse{}
	for _, row := range result {
		dtoReserveHoldsResponse := NewReserveHoldsSelectableResponse(row.ReserveHolds, filter)
		dtoReserveHoldsListResponse = append(dtoReserveHoldsListResponse, &dtoReserveHoldsResponse)
	}
	return dtoReserveHoldsListResponse
}

type ReserveHoldsFilterResponse struct {
	Metadata Metadata                           `json:"metadata"`
	Data     ReserveHoldsSelectableListResponse `json:"data"`
}

func reverseReserveHoldsFilterResults(result []model.ReserveHoldsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReserveHoldsFilterResponse(result []model.ReserveHoldsFilterResult, filter model.Filter) (resp ReserveHoldsFilterResponse) {
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
			reverseReserveHoldsFilterResults(dataResult)
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

	resp.Data = NewReserveHoldsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReserveHoldsCreateRequest struct {
	MerchantPartyId uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode    string           `json:"currencyCode"`
	SourceType      string           `json:"sourceType"`
	SourceId        uuid.UUID        `json:"sourceId"`
	HoldType        string           `json:"holdType"`
	HoldStatus      model.HoldStatus `json:"holdStatus" example:"active" enums:"active,released,consumed,canceled"`
	HoldAmount      decimal.Decimal  `json:"holdAmount"`
	ReleasedAmount  decimal.Decimal  `json:"releasedAmount"`
	EffectiveAt     time.Time        `json:"effectiveAt"`
	ReleaseAt       time.Time        `json:"releaseAt"`
	ReleasedAt      time.Time        `json:"releasedAt"`
	ReasonCode      string           `json:"reasonCode"`
	ReasonDetail    string           `json:"reasonDetail"`
	Metadata        json.RawMessage  `json:"metadata"`
}

func (d *ReserveHoldsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReserveHoldsCreateRequest) ToModel() model.ReserveHolds {
	id, _ := uuid.NewV4()
	return model.ReserveHolds{
		Id:              id,
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		HoldType:        d.HoldType,
		HoldStatus:      d.HoldStatus,
		HoldAmount:      d.HoldAmount,
		ReleasedAmount:  d.ReleasedAmount,
		EffectiveAt:     d.EffectiveAt,
		ReleaseAt:       null.TimeFrom(d.ReleaseAt),
		ReleasedAt:      null.TimeFrom(d.ReleasedAt),
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Metadata:        d.Metadata,
	}
}

type ReserveHoldsListCreateRequest []*ReserveHoldsCreateRequest

func (d ReserveHoldsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reserveHolds := range d {
		err = validator.Struct(reserveHolds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReserveHoldsListCreateRequest) ToModelList() []model.ReserveHolds {
	out := make([]model.ReserveHolds, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReserveHoldsUpdateRequest struct {
	MerchantPartyId uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode    string           `json:"currencyCode"`
	SourceType      string           `json:"sourceType"`
	SourceId        uuid.UUID        `json:"sourceId"`
	HoldType        string           `json:"holdType"`
	HoldStatus      model.HoldStatus `json:"holdStatus" example:"active" enums:"active,released,consumed,canceled"`
	HoldAmount      decimal.Decimal  `json:"holdAmount"`
	ReleasedAmount  decimal.Decimal  `json:"releasedAmount"`
	EffectiveAt     time.Time        `json:"effectiveAt"`
	ReleaseAt       time.Time        `json:"releaseAt"`
	ReleasedAt      time.Time        `json:"releasedAt"`
	ReasonCode      string           `json:"reasonCode"`
	ReasonDetail    string           `json:"reasonDetail"`
	Metadata        json.RawMessage  `json:"metadata"`
}

func (d *ReserveHoldsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReserveHoldsUpdateRequest) ToModel() model.ReserveHolds {
	return model.ReserveHolds{
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		HoldType:        d.HoldType,
		HoldStatus:      d.HoldStatus,
		HoldAmount:      d.HoldAmount,
		ReleasedAmount:  d.ReleasedAmount,
		EffectiveAt:     d.EffectiveAt,
		ReleaseAt:       null.TimeFrom(d.ReleaseAt),
		ReleasedAt:      null.TimeFrom(d.ReleasedAt),
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Metadata:        d.Metadata,
	}
}

type ReserveHoldsBulkUpdateRequest struct {
	Id              uuid.UUID        `json:"id"`
	MerchantPartyId uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode    string           `json:"currencyCode"`
	SourceType      string           `json:"sourceType"`
	SourceId        uuid.UUID        `json:"sourceId"`
	HoldType        string           `json:"holdType"`
	HoldStatus      model.HoldStatus `json:"holdStatus" example:"active" enums:"active,released,consumed,canceled"`
	HoldAmount      decimal.Decimal  `json:"holdAmount"`
	ReleasedAmount  decimal.Decimal  `json:"releasedAmount"`
	EffectiveAt     time.Time        `json:"effectiveAt"`
	ReleaseAt       time.Time        `json:"releaseAt"`
	ReleasedAt      time.Time        `json:"releasedAt"`
	ReasonCode      string           `json:"reasonCode"`
	ReasonDetail    string           `json:"reasonDetail"`
	Metadata        json.RawMessage  `json:"metadata"`
}

func (d ReserveHoldsBulkUpdateRequest) PrimaryID() ReserveHoldsPrimaryID {
	return ReserveHoldsPrimaryID{
		Id: d.Id,
	}
}

type ReserveHoldsListBulkUpdateRequest []*ReserveHoldsBulkUpdateRequest

func (d ReserveHoldsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reserveHolds := range d {
		err = validator.Struct(reserveHolds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReserveHoldsBulkUpdateRequest) ToModel() model.ReserveHolds {
	return model.ReserveHolds{
		Id:              d.Id,
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
		SourceType:      d.SourceType,
		SourceId:        d.SourceId,
		HoldType:        d.HoldType,
		HoldStatus:      d.HoldStatus,
		HoldAmount:      d.HoldAmount,
		ReleasedAmount:  d.ReleasedAmount,
		EffectiveAt:     d.EffectiveAt,
		ReleaseAt:       null.TimeFrom(d.ReleaseAt),
		ReleasedAt:      null.TimeFrom(d.ReleasedAt),
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Metadata:        d.Metadata,
	}
}

type ReserveHoldsResponse struct {
	Id              uuid.UUID        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId uuid.UUID        `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode    string           `json:"currencyCode" validate:"required"`
	SourceType      string           `json:"sourceType" validate:"required"`
	SourceId        uuid.UUID        `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	HoldType        string           `json:"holdType" validate:"required"`
	HoldStatus      model.HoldStatus `json:"holdStatus" validate:"oneof=active released consumed canceled" enums:"active,released,consumed,canceled"`
	HoldAmount      decimal.Decimal  `json:"holdAmount" validate:"required" format:"decimal" example:"100.50"`
	ReleasedAmount  decimal.Decimal  `json:"releasedAmount" format:"decimal" example:"100.50"`
	EffectiveAt     time.Time        `json:"effectiveAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ReleaseAt       time.Time        `json:"releaseAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ReleasedAt      time.Time        `json:"releasedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ReasonCode      string           `json:"reasonCode" validate:"required"`
	ReasonDetail    string           `json:"reasonDetail"`
	Metadata        json.RawMessage  `json:"metadata" swaggertype:"object"`
}

func NewReserveHoldsResponse(reserveHolds model.ReserveHolds) ReserveHoldsResponse {
	return ReserveHoldsResponse{
		Id:              reserveHolds.Id,
		MerchantPartyId: reserveHolds.MerchantPartyId,
		CurrencyCode:    reserveHolds.CurrencyCode,
		SourceType:      reserveHolds.SourceType,
		SourceId:        reserveHolds.SourceId,
		HoldType:        reserveHolds.HoldType,
		HoldStatus:      model.HoldStatus(reserveHolds.HoldStatus),
		HoldAmount:      reserveHolds.HoldAmount,
		ReleasedAmount:  reserveHolds.ReleasedAmount,
		EffectiveAt:     reserveHolds.EffectiveAt,
		ReleaseAt:       reserveHolds.ReleaseAt.Time,
		ReleasedAt:      reserveHolds.ReleasedAt.Time,
		ReasonCode:      reserveHolds.ReasonCode,
		ReasonDetail:    reserveHolds.ReasonDetail.String,
		Metadata:        reserveHolds.Metadata,
	}
}

type ReserveHoldsListResponse []*ReserveHoldsResponse

func NewReserveHoldsListResponse(reserveHoldsList model.ReserveHoldsList) ReserveHoldsListResponse {
	dtoReserveHoldsListResponse := ReserveHoldsListResponse{}
	for _, reserveHolds := range reserveHoldsList {
		dtoReserveHoldsResponse := NewReserveHoldsResponse(*reserveHolds)
		dtoReserveHoldsListResponse = append(dtoReserveHoldsListResponse, &dtoReserveHoldsResponse)
	}
	return dtoReserveHoldsListResponse
}

type ReserveHoldsPrimaryIDList []ReserveHoldsPrimaryID

func (d ReserveHoldsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reserveHolds := range d {
		err = validator.Struct(reserveHolds)
		if err != nil {
			return
		}
	}
	return nil
}

type ReserveHoldsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReserveHoldsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReserveHoldsPrimaryID) ToModel() model.ReserveHoldsPrimaryID {
	return model.ReserveHoldsPrimaryID{
		Id: d.Id,
	}
}

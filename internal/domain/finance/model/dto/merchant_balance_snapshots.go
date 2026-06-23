package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type MerchantBalanceSnapshotsDTOFieldNameType string

type merchantBalanceSnapshotsDTOFieldName struct {
	Id               MerchantBalanceSnapshotsDTOFieldNameType
	BalanceAccountId MerchantBalanceSnapshotsDTOFieldNameType
	SnapshotAt       MerchantBalanceSnapshotsDTOFieldNameType
	AvailableAmount  MerchantBalanceSnapshotsDTOFieldNameType
	PendingAmount    MerchantBalanceSnapshotsDTOFieldNameType
	ReservedAmount   MerchantBalanceSnapshotsDTOFieldNameType
	DisputedAmount   MerchantBalanceSnapshotsDTOFieldNameType
	NegativeAmount   MerchantBalanceSnapshotsDTOFieldNameType
	Metadata         MerchantBalanceSnapshotsDTOFieldNameType
	MetaCreatedAt    MerchantBalanceSnapshotsDTOFieldNameType
	MetaCreatedBy    MerchantBalanceSnapshotsDTOFieldNameType
	MetaUpdatedAt    MerchantBalanceSnapshotsDTOFieldNameType
	MetaUpdatedBy    MerchantBalanceSnapshotsDTOFieldNameType
	MetaDeletedAt    MerchantBalanceSnapshotsDTOFieldNameType
	MetaDeletedBy    MerchantBalanceSnapshotsDTOFieldNameType
}

var MerchantBalanceSnapshotsDTOFieldName = merchantBalanceSnapshotsDTOFieldName{
	Id:               "id",
	BalanceAccountId: "balanceAccountId",
	SnapshotAt:       "snapshotAt",
	AvailableAmount:  "availableAmount",
	PendingAmount:    "pendingAmount",
	ReservedAmount:   "reservedAmount",
	DisputedAmount:   "disputedAmount",
	NegativeAmount:   "negativeAmount",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformMerchantBalanceSnapshotsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(MerchantBalanceSnapshotsDTOFieldName.Id):
		return string(model.MerchantBalanceSnapshotsDBFieldName.Id), true

	case string(MerchantBalanceSnapshotsDTOFieldName.BalanceAccountId):
		return string(model.MerchantBalanceSnapshotsDBFieldName.BalanceAccountId), true

	case string(MerchantBalanceSnapshotsDTOFieldName.SnapshotAt):
		return string(model.MerchantBalanceSnapshotsDBFieldName.SnapshotAt), true

	case string(MerchantBalanceSnapshotsDTOFieldName.AvailableAmount):
		return string(model.MerchantBalanceSnapshotsDBFieldName.AvailableAmount), true

	case string(MerchantBalanceSnapshotsDTOFieldName.PendingAmount):
		return string(model.MerchantBalanceSnapshotsDBFieldName.PendingAmount), true

	case string(MerchantBalanceSnapshotsDTOFieldName.ReservedAmount):
		return string(model.MerchantBalanceSnapshotsDBFieldName.ReservedAmount), true

	case string(MerchantBalanceSnapshotsDTOFieldName.DisputedAmount):
		return string(model.MerchantBalanceSnapshotsDBFieldName.DisputedAmount), true

	case string(MerchantBalanceSnapshotsDTOFieldName.NegativeAmount):
		return string(model.MerchantBalanceSnapshotsDBFieldName.NegativeAmount), true

	case string(MerchantBalanceSnapshotsDTOFieldName.Metadata):
		return string(model.MerchantBalanceSnapshotsDBFieldName.Metadata), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaCreatedAt):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedAt), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaCreatedBy):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedBy), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaUpdatedAt):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedAt), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaUpdatedBy):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedBy), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaDeletedAt):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedAt), true

	case string(MerchantBalanceSnapshotsDTOFieldName.MetaDeletedBy):
		return string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isMerchantBalanceSnapshotsBaseFilterField(field string) bool {
	spec, found := model.NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeMerchantBalanceSnapshotsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateMerchantBalanceSnapshotsProjectionOutputPath(path string) error {
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

func transformMerchantBalanceSnapshotsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformMerchantBalanceSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformMerchantBalanceSnapshotsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformMerchantBalanceSnapshotsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformMerchantBalanceSnapshotsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isMerchantBalanceSnapshotsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateMerchantBalanceSnapshotsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeMerchantBalanceSnapshotsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformMerchantBalanceSnapshotsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformMerchantBalanceSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformMerchantBalanceSnapshotsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultMerchantBalanceSnapshotsFilter(filter *model.Filter) {
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
			Field: string(MerchantBalanceSnapshotsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type MerchantBalanceSnapshotsSelectableResponse map[string]interface{}
type MerchantBalanceSnapshotsSelectableListResponse []*MerchantBalanceSnapshotsSelectableResponse

func assignMerchantBalanceSnapshotsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setMerchantBalanceSnapshotsSelectableValue(out MerchantBalanceSnapshotsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignMerchantBalanceSnapshotsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewMerchantBalanceSnapshotsSelectableResponse(merchantBalanceSnapshots model.MerchantBalanceSnapshots, filter model.Filter) MerchantBalanceSnapshotsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.MerchantBalanceSnapshotsDBFieldName.Id),
			string(model.MerchantBalanceSnapshotsDBFieldName.BalanceAccountId),
			string(model.MerchantBalanceSnapshotsDBFieldName.SnapshotAt),
			string(model.MerchantBalanceSnapshotsDBFieldName.AvailableAmount),
			string(model.MerchantBalanceSnapshotsDBFieldName.PendingAmount),
			string(model.MerchantBalanceSnapshotsDBFieldName.ReservedAmount),
			string(model.MerchantBalanceSnapshotsDBFieldName.DisputedAmount),
			string(model.MerchantBalanceSnapshotsDBFieldName.NegativeAmount),
			string(model.MerchantBalanceSnapshotsDBFieldName.Metadata),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedAt),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedBy),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedAt),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedBy),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedAt),
			string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedBy),
		)
	}
	merchantBalanceSnapshotsSelectableResponse := MerchantBalanceSnapshotsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.MerchantBalanceSnapshotsDBFieldName.Id):
			key := string(MerchantBalanceSnapshotsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.Id, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.BalanceAccountId):
			key := string(MerchantBalanceSnapshotsDTOFieldName.BalanceAccountId)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.BalanceAccountId, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.SnapshotAt):
			key := string(MerchantBalanceSnapshotsDTOFieldName.SnapshotAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.SnapshotAt, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.AvailableAmount):
			key := string(MerchantBalanceSnapshotsDTOFieldName.AvailableAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.AvailableAmount, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.PendingAmount):
			key := string(MerchantBalanceSnapshotsDTOFieldName.PendingAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.PendingAmount, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.ReservedAmount):
			key := string(MerchantBalanceSnapshotsDTOFieldName.ReservedAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.ReservedAmount, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.DisputedAmount):
			key := string(MerchantBalanceSnapshotsDTOFieldName.DisputedAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.DisputedAmount, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.NegativeAmount):
			key := string(MerchantBalanceSnapshotsDTOFieldName.NegativeAmount)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.NegativeAmount, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.Metadata):
			key := string(MerchantBalanceSnapshotsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.Metadata, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedAt):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaCreatedAt, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaCreatedBy):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaCreatedBy, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedAt):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaUpdatedAt, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaUpdatedBy):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaUpdatedBy, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedAt):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaDeletedAt.Time, explicitAlias)

		case string(model.MerchantBalanceSnapshotsDBFieldName.MetaDeletedBy):
			key := string(MerchantBalanceSnapshotsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setMerchantBalanceSnapshotsSelectableValue(merchantBalanceSnapshotsSelectableResponse, key, merchantBalanceSnapshots.MetaDeletedBy, explicitAlias)

		}
	}
	return merchantBalanceSnapshotsSelectableResponse
}

func NewMerchantBalanceSnapshotsListResponseFromFilterResult(result []model.MerchantBalanceSnapshotsFilterResult, filter model.Filter) MerchantBalanceSnapshotsSelectableListResponse {
	dtoMerchantBalanceSnapshotsListResponse := MerchantBalanceSnapshotsSelectableListResponse{}
	for _, row := range result {
		dtoMerchantBalanceSnapshotsResponse := NewMerchantBalanceSnapshotsSelectableResponse(row.MerchantBalanceSnapshots, filter)
		dtoMerchantBalanceSnapshotsListResponse = append(dtoMerchantBalanceSnapshotsListResponse, &dtoMerchantBalanceSnapshotsResponse)
	}
	return dtoMerchantBalanceSnapshotsListResponse
}

type MerchantBalanceSnapshotsFilterResponse struct {
	Metadata Metadata                                       `json:"metadata"`
	Data     MerchantBalanceSnapshotsSelectableListResponse `json:"data"`
}

func reverseMerchantBalanceSnapshotsFilterResults(result []model.MerchantBalanceSnapshotsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewMerchantBalanceSnapshotsFilterResponse(result []model.MerchantBalanceSnapshotsFilterResult, filter model.Filter) (resp MerchantBalanceSnapshotsFilterResponse) {
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
			reverseMerchantBalanceSnapshotsFilterResults(dataResult)
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

	resp.Data = NewMerchantBalanceSnapshotsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type MerchantBalanceSnapshotsCreateRequest struct {
	BalanceAccountId uuid.UUID       `json:"balanceAccountId"`
	SnapshotAt       time.Time       `json:"snapshotAt"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	DisputedAmount   decimal.Decimal `json:"disputedAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *MerchantBalanceSnapshotsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *MerchantBalanceSnapshotsCreateRequest) ToModel() model.MerchantBalanceSnapshots {
	id, _ := uuid.NewV4()
	return model.MerchantBalanceSnapshots{
		Id:               id,
		BalanceAccountId: d.BalanceAccountId,
		SnapshotAt:       d.SnapshotAt,
		AvailableAmount:  d.AvailableAmount,
		PendingAmount:    d.PendingAmount,
		ReservedAmount:   d.ReservedAmount,
		DisputedAmount:   d.DisputedAmount,
		NegativeAmount:   d.NegativeAmount,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceSnapshotsListCreateRequest []*MerchantBalanceSnapshotsCreateRequest

func (d MerchantBalanceSnapshotsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSnapshots := range d {
		err = validator.Struct(merchantBalanceSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceSnapshotsListCreateRequest) ToModelList() []model.MerchantBalanceSnapshots {
	out := make([]model.MerchantBalanceSnapshots, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type MerchantBalanceSnapshotsUpdateRequest struct {
	BalanceAccountId uuid.UUID       `json:"balanceAccountId"`
	SnapshotAt       time.Time       `json:"snapshotAt"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	DisputedAmount   decimal.Decimal `json:"disputedAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *MerchantBalanceSnapshotsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d MerchantBalanceSnapshotsUpdateRequest) ToModel() model.MerchantBalanceSnapshots {
	return model.MerchantBalanceSnapshots{
		BalanceAccountId: d.BalanceAccountId,
		SnapshotAt:       d.SnapshotAt,
		AvailableAmount:  d.AvailableAmount,
		PendingAmount:    d.PendingAmount,
		ReservedAmount:   d.ReservedAmount,
		DisputedAmount:   d.DisputedAmount,
		NegativeAmount:   d.NegativeAmount,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceSnapshotsBulkUpdateRequest struct {
	Id               uuid.UUID       `json:"id"`
	BalanceAccountId uuid.UUID       `json:"balanceAccountId"`
	SnapshotAt       time.Time       `json:"snapshotAt"`
	AvailableAmount  decimal.Decimal `json:"availableAmount"`
	PendingAmount    decimal.Decimal `json:"pendingAmount"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount"`
	DisputedAmount   decimal.Decimal `json:"disputedAmount"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d MerchantBalanceSnapshotsBulkUpdateRequest) PrimaryID() MerchantBalanceSnapshotsPrimaryID {
	return MerchantBalanceSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceSnapshotsListBulkUpdateRequest []*MerchantBalanceSnapshotsBulkUpdateRequest

func (d MerchantBalanceSnapshotsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSnapshots := range d {
		err = validator.Struct(merchantBalanceSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d MerchantBalanceSnapshotsBulkUpdateRequest) ToModel() model.MerchantBalanceSnapshots {
	return model.MerchantBalanceSnapshots{
		Id:               d.Id,
		BalanceAccountId: d.BalanceAccountId,
		SnapshotAt:       d.SnapshotAt,
		AvailableAmount:  d.AvailableAmount,
		PendingAmount:    d.PendingAmount,
		ReservedAmount:   d.ReservedAmount,
		DisputedAmount:   d.DisputedAmount,
		NegativeAmount:   d.NegativeAmount,
		Metadata:         d.Metadata,
	}
}

type MerchantBalanceSnapshotsResponse struct {
	Id               uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BalanceAccountId uuid.UUID       `json:"balanceAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SnapshotAt       time.Time       `json:"snapshotAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	AvailableAmount  decimal.Decimal `json:"availableAmount" format:"decimal" example:"100.50"`
	PendingAmount    decimal.Decimal `json:"pendingAmount" format:"decimal" example:"100.50"`
	ReservedAmount   decimal.Decimal `json:"reservedAmount" format:"decimal" example:"100.50"`
	DisputedAmount   decimal.Decimal `json:"disputedAmount" format:"decimal" example:"100.50"`
	NegativeAmount   decimal.Decimal `json:"negativeAmount" format:"decimal" example:"100.50"`
	Metadata         json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewMerchantBalanceSnapshotsResponse(merchantBalanceSnapshots model.MerchantBalanceSnapshots) MerchantBalanceSnapshotsResponse {
	return MerchantBalanceSnapshotsResponse{
		Id:               merchantBalanceSnapshots.Id,
		BalanceAccountId: merchantBalanceSnapshots.BalanceAccountId,
		SnapshotAt:       merchantBalanceSnapshots.SnapshotAt,
		AvailableAmount:  merchantBalanceSnapshots.AvailableAmount,
		PendingAmount:    merchantBalanceSnapshots.PendingAmount,
		ReservedAmount:   merchantBalanceSnapshots.ReservedAmount,
		DisputedAmount:   merchantBalanceSnapshots.DisputedAmount,
		NegativeAmount:   merchantBalanceSnapshots.NegativeAmount,
		Metadata:         merchantBalanceSnapshots.Metadata,
	}
}

type MerchantBalanceSnapshotsListResponse []*MerchantBalanceSnapshotsResponse

func NewMerchantBalanceSnapshotsListResponse(merchantBalanceSnapshotsList model.MerchantBalanceSnapshotsList) MerchantBalanceSnapshotsListResponse {
	dtoMerchantBalanceSnapshotsListResponse := MerchantBalanceSnapshotsListResponse{}
	for _, merchantBalanceSnapshots := range merchantBalanceSnapshotsList {
		dtoMerchantBalanceSnapshotsResponse := NewMerchantBalanceSnapshotsResponse(*merchantBalanceSnapshots)
		dtoMerchantBalanceSnapshotsListResponse = append(dtoMerchantBalanceSnapshotsListResponse, &dtoMerchantBalanceSnapshotsResponse)
	}
	return dtoMerchantBalanceSnapshotsListResponse
}

type MerchantBalanceSnapshotsPrimaryIDList []MerchantBalanceSnapshotsPrimaryID

func (d MerchantBalanceSnapshotsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, merchantBalanceSnapshots := range d {
		err = validator.Struct(merchantBalanceSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

type MerchantBalanceSnapshotsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *MerchantBalanceSnapshotsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d MerchantBalanceSnapshotsPrimaryID) ToModel() model.MerchantBalanceSnapshotsPrimaryID {
	return model.MerchantBalanceSnapshotsPrimaryID{
		Id: d.Id,
	}
}

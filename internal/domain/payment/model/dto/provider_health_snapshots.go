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

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderHealthSnapshotsDTOFieldNameType string

type providerHealthSnapshotsDTOFieldName struct {
	Id                ProviderHealthSnapshotsDTOFieldNameType
	ProviderAccountId ProviderHealthSnapshotsDTOFieldNameType
	MethodCode        ProviderHealthSnapshotsDTOFieldNameType
	ChannelCode       ProviderHealthSnapshotsDTOFieldNameType
	HealthScore       ProviderHealthSnapshotsDTOFieldNameType
	SuccessRate       ProviderHealthSnapshotsDTOFieldNameType
	TimeoutRate       ProviderHealthSnapshotsDTOFieldNameType
	ErrorRate         ProviderHealthSnapshotsDTOFieldNameType
	P95LatencyMs      ProviderHealthSnapshotsDTOFieldNameType
	SampleSize        ProviderHealthSnapshotsDTOFieldNameType
	WindowStartedAt   ProviderHealthSnapshotsDTOFieldNameType
	WindowEndedAt     ProviderHealthSnapshotsDTOFieldNameType
	Metadata          ProviderHealthSnapshotsDTOFieldNameType
	MetaCreatedAt     ProviderHealthSnapshotsDTOFieldNameType
	MetaCreatedBy     ProviderHealthSnapshotsDTOFieldNameType
	MetaUpdatedAt     ProviderHealthSnapshotsDTOFieldNameType
	MetaUpdatedBy     ProviderHealthSnapshotsDTOFieldNameType
	MetaDeletedAt     ProviderHealthSnapshotsDTOFieldNameType
	MetaDeletedBy     ProviderHealthSnapshotsDTOFieldNameType
}

var ProviderHealthSnapshotsDTOFieldName = providerHealthSnapshotsDTOFieldName{
	Id:                "id",
	ProviderAccountId: "providerAccountId",
	MethodCode:        "methodCode",
	ChannelCode:       "channelCode",
	HealthScore:       "healthScore",
	SuccessRate:       "successRate",
	TimeoutRate:       "timeoutRate",
	ErrorRate:         "errorRate",
	P95LatencyMs:      "p95LatencyMs",
	SampleSize:        "sampleSize",
	WindowStartedAt:   "windowStartedAt",
	WindowEndedAt:     "windowEndedAt",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformProviderHealthSnapshotsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderHealthSnapshotsDTOFieldName.Id):
		return string(model.ProviderHealthSnapshotsDBFieldName.Id), true

	case string(ProviderHealthSnapshotsDTOFieldName.ProviderAccountId):
		return string(model.ProviderHealthSnapshotsDBFieldName.ProviderAccountId), true

	case string(ProviderHealthSnapshotsDTOFieldName.MethodCode):
		return string(model.ProviderHealthSnapshotsDBFieldName.MethodCode), true

	case string(ProviderHealthSnapshotsDTOFieldName.ChannelCode):
		return string(model.ProviderHealthSnapshotsDBFieldName.ChannelCode), true

	case string(ProviderHealthSnapshotsDTOFieldName.HealthScore):
		return string(model.ProviderHealthSnapshotsDBFieldName.HealthScore), true

	case string(ProviderHealthSnapshotsDTOFieldName.SuccessRate):
		return string(model.ProviderHealthSnapshotsDBFieldName.SuccessRate), true

	case string(ProviderHealthSnapshotsDTOFieldName.TimeoutRate):
		return string(model.ProviderHealthSnapshotsDBFieldName.TimeoutRate), true

	case string(ProviderHealthSnapshotsDTOFieldName.ErrorRate):
		return string(model.ProviderHealthSnapshotsDBFieldName.ErrorRate), true

	case string(ProviderHealthSnapshotsDTOFieldName.P95LatencyMs):
		return string(model.ProviderHealthSnapshotsDBFieldName.P95LatencyMs), true

	case string(ProviderHealthSnapshotsDTOFieldName.SampleSize):
		return string(model.ProviderHealthSnapshotsDBFieldName.SampleSize), true

	case string(ProviderHealthSnapshotsDTOFieldName.WindowStartedAt):
		return string(model.ProviderHealthSnapshotsDBFieldName.WindowStartedAt), true

	case string(ProviderHealthSnapshotsDTOFieldName.WindowEndedAt):
		return string(model.ProviderHealthSnapshotsDBFieldName.WindowEndedAt), true

	case string(ProviderHealthSnapshotsDTOFieldName.Metadata):
		return string(model.ProviderHealthSnapshotsDBFieldName.Metadata), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaCreatedAt):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedAt), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaCreatedBy):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedBy), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedAt), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedBy), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaDeletedAt):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedAt), true

	case string(ProviderHealthSnapshotsDTOFieldName.MetaDeletedBy):
		return string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderHealthSnapshotsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderHealthSnapshotsBaseFilterField(field string) bool {
	spec, found := model.NewProviderHealthSnapshotsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderHealthSnapshotsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderHealthSnapshotsProjectionOutputPath(path string) error {
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

func transformProviderHealthSnapshotsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderHealthSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderHealthSnapshotsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderHealthSnapshotsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderHealthSnapshotsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderHealthSnapshotsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderHealthSnapshotsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderHealthSnapshotsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderHealthSnapshotsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderHealthSnapshotsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderHealthSnapshotsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderHealthSnapshotsFilter(filter *model.Filter) {
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
			Field: string(ProviderHealthSnapshotsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderHealthSnapshotsSelectableResponse map[string]interface{}
type ProviderHealthSnapshotsSelectableListResponse []*ProviderHealthSnapshotsSelectableResponse

func assignProviderHealthSnapshotsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderHealthSnapshotsSelectableValue(out ProviderHealthSnapshotsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderHealthSnapshotsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderHealthSnapshotsSelectableResponse(providerHealthSnapshots model.ProviderHealthSnapshots, filter model.Filter) ProviderHealthSnapshotsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderHealthSnapshotsDBFieldName.Id),
			string(model.ProviderHealthSnapshotsDBFieldName.ProviderAccountId),
			string(model.ProviderHealthSnapshotsDBFieldName.MethodCode),
			string(model.ProviderHealthSnapshotsDBFieldName.ChannelCode),
			string(model.ProviderHealthSnapshotsDBFieldName.HealthScore),
			string(model.ProviderHealthSnapshotsDBFieldName.SuccessRate),
			string(model.ProviderHealthSnapshotsDBFieldName.TimeoutRate),
			string(model.ProviderHealthSnapshotsDBFieldName.ErrorRate),
			string(model.ProviderHealthSnapshotsDBFieldName.P95LatencyMs),
			string(model.ProviderHealthSnapshotsDBFieldName.SampleSize),
			string(model.ProviderHealthSnapshotsDBFieldName.WindowStartedAt),
			string(model.ProviderHealthSnapshotsDBFieldName.WindowEndedAt),
			string(model.ProviderHealthSnapshotsDBFieldName.Metadata),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedAt),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedBy),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedAt),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedBy),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedAt),
			string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedBy),
		)
	}
	providerHealthSnapshotsSelectableResponse := ProviderHealthSnapshotsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderHealthSnapshotsDBFieldName.Id):
			key := string(ProviderHealthSnapshotsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.Id, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.ProviderAccountId):
			key := string(ProviderHealthSnapshotsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.ProviderAccountId, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MethodCode):
			key := string(ProviderHealthSnapshotsDTOFieldName.MethodCode)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MethodCode.String, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.ChannelCode):
			key := string(ProviderHealthSnapshotsDTOFieldName.ChannelCode)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.ChannelCode.String, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.HealthScore):
			key := string(ProviderHealthSnapshotsDTOFieldName.HealthScore)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.HealthScore, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.SuccessRate):
			key := string(ProviderHealthSnapshotsDTOFieldName.SuccessRate)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.SuccessRate.Decimal, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.TimeoutRate):
			key := string(ProviderHealthSnapshotsDTOFieldName.TimeoutRate)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.TimeoutRate.Decimal, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.ErrorRate):
			key := string(ProviderHealthSnapshotsDTOFieldName.ErrorRate)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.ErrorRate.Decimal, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.P95LatencyMs):
			key := string(ProviderHealthSnapshotsDTOFieldName.P95LatencyMs)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, int(providerHealthSnapshots.P95LatencyMs.ValueOrZero()), explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.SampleSize):
			key := string(ProviderHealthSnapshotsDTOFieldName.SampleSize)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.SampleSize, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.WindowStartedAt):
			key := string(ProviderHealthSnapshotsDTOFieldName.WindowStartedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.WindowStartedAt, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.WindowEndedAt):
			key := string(ProviderHealthSnapshotsDTOFieldName.WindowEndedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.WindowEndedAt, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.Metadata):
			key := string(ProviderHealthSnapshotsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.Metadata, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedAt):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaCreatedAt, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaCreatedBy):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaCreatedBy, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedAt):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaUpdatedBy):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedAt):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderHealthSnapshotsDBFieldName.MetaDeletedBy):
			key := string(ProviderHealthSnapshotsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderHealthSnapshotsSelectableValue(providerHealthSnapshotsSelectableResponse, key, providerHealthSnapshots.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return providerHealthSnapshotsSelectableResponse
}

func NewProviderHealthSnapshotsListResponseFromFilterResult(result []model.ProviderHealthSnapshotsFilterResult, filter model.Filter) ProviderHealthSnapshotsSelectableListResponse {
	dtoProviderHealthSnapshotsListResponse := ProviderHealthSnapshotsSelectableListResponse{}
	for _, row := range result {
		dtoProviderHealthSnapshotsResponse := NewProviderHealthSnapshotsSelectableResponse(row.ProviderHealthSnapshots, filter)
		dtoProviderHealthSnapshotsListResponse = append(dtoProviderHealthSnapshotsListResponse, &dtoProviderHealthSnapshotsResponse)
	}
	return dtoProviderHealthSnapshotsListResponse
}

type ProviderHealthSnapshotsFilterResponse struct {
	Metadata Metadata                                      `json:"metadata"`
	Data     ProviderHealthSnapshotsSelectableListResponse `json:"data"`
}

func reverseProviderHealthSnapshotsFilterResults(result []model.ProviderHealthSnapshotsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderHealthSnapshotsFilterResponse(result []model.ProviderHealthSnapshotsFilterResult, filter model.Filter) (resp ProviderHealthSnapshotsFilterResponse) {
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
			reverseProviderHealthSnapshotsFilterResults(dataResult)
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

	resp.Data = NewProviderHealthSnapshotsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderHealthSnapshotsCreateRequest struct {
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	MethodCode        string          `json:"methodCode"`
	ChannelCode       string          `json:"channelCode"`
	HealthScore       int             `json:"healthScore"`
	SuccessRate       decimal.Decimal `json:"successRate"`
	TimeoutRate       decimal.Decimal `json:"timeoutRate"`
	ErrorRate         decimal.Decimal `json:"errorRate"`
	P95LatencyMs      int             `json:"p95LatencyMs"`
	SampleSize        int             `json:"sampleSize"`
	WindowStartedAt   time.Time       `json:"windowStartedAt"`
	WindowEndedAt     time.Time       `json:"windowEndedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *ProviderHealthSnapshotsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderHealthSnapshotsCreateRequest) ToModel() model.ProviderHealthSnapshots {
	id, _ := uuid.NewV4()
	return model.ProviderHealthSnapshots{
		Id:                id,
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		HealthScore:       d.HealthScore,
		SuccessRate:       decimal.NewNullDecimal(d.SuccessRate),
		TimeoutRate:       decimal.NewNullDecimal(d.TimeoutRate),
		ErrorRate:         decimal.NewNullDecimal(d.ErrorRate),
		P95LatencyMs:      null.IntFrom(int64(d.P95LatencyMs)),
		SampleSize:        d.SampleSize,
		WindowStartedAt:   d.WindowStartedAt,
		WindowEndedAt:     d.WindowEndedAt,
		Metadata:          d.Metadata,
	}
}

type ProviderHealthSnapshotsListCreateRequest []*ProviderHealthSnapshotsCreateRequest

func (d ProviderHealthSnapshotsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerHealthSnapshots := range d {
		err = validator.Struct(providerHealthSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderHealthSnapshotsListCreateRequest) ToModelList() []model.ProviderHealthSnapshots {
	out := make([]model.ProviderHealthSnapshots, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderHealthSnapshotsUpdateRequest struct {
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	MethodCode        string          `json:"methodCode"`
	ChannelCode       string          `json:"channelCode"`
	HealthScore       int             `json:"healthScore"`
	SuccessRate       decimal.Decimal `json:"successRate"`
	TimeoutRate       decimal.Decimal `json:"timeoutRate"`
	ErrorRate         decimal.Decimal `json:"errorRate"`
	P95LatencyMs      int             `json:"p95LatencyMs"`
	SampleSize        int             `json:"sampleSize"`
	WindowStartedAt   time.Time       `json:"windowStartedAt"`
	WindowEndedAt     time.Time       `json:"windowEndedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *ProviderHealthSnapshotsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderHealthSnapshotsUpdateRequest) ToModel() model.ProviderHealthSnapshots {
	return model.ProviderHealthSnapshots{
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		HealthScore:       d.HealthScore,
		SuccessRate:       decimal.NewNullDecimal(d.SuccessRate),
		TimeoutRate:       decimal.NewNullDecimal(d.TimeoutRate),
		ErrorRate:         decimal.NewNullDecimal(d.ErrorRate),
		P95LatencyMs:      null.IntFrom(int64(d.P95LatencyMs)),
		SampleSize:        d.SampleSize,
		WindowStartedAt:   d.WindowStartedAt,
		WindowEndedAt:     d.WindowEndedAt,
		Metadata:          d.Metadata,
	}
}

type ProviderHealthSnapshotsBulkUpdateRequest struct {
	Id                uuid.UUID       `json:"id"`
	ProviderAccountId uuid.UUID       `json:"providerAccountId"`
	MethodCode        string          `json:"methodCode"`
	ChannelCode       string          `json:"channelCode"`
	HealthScore       int             `json:"healthScore"`
	SuccessRate       decimal.Decimal `json:"successRate"`
	TimeoutRate       decimal.Decimal `json:"timeoutRate"`
	ErrorRate         decimal.Decimal `json:"errorRate"`
	P95LatencyMs      int             `json:"p95LatencyMs"`
	SampleSize        int             `json:"sampleSize"`
	WindowStartedAt   time.Time       `json:"windowStartedAt"`
	WindowEndedAt     time.Time       `json:"windowEndedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d ProviderHealthSnapshotsBulkUpdateRequest) PrimaryID() ProviderHealthSnapshotsPrimaryID {
	return ProviderHealthSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type ProviderHealthSnapshotsListBulkUpdateRequest []*ProviderHealthSnapshotsBulkUpdateRequest

func (d ProviderHealthSnapshotsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerHealthSnapshots := range d {
		err = validator.Struct(providerHealthSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderHealthSnapshotsBulkUpdateRequest) ToModel() model.ProviderHealthSnapshots {
	return model.ProviderHealthSnapshots{
		Id:                d.Id,
		ProviderAccountId: d.ProviderAccountId,
		MethodCode:        null.StringFrom(d.MethodCode),
		ChannelCode:       null.StringFrom(d.ChannelCode),
		HealthScore:       d.HealthScore,
		SuccessRate:       decimal.NewNullDecimal(d.SuccessRate),
		TimeoutRate:       decimal.NewNullDecimal(d.TimeoutRate),
		ErrorRate:         decimal.NewNullDecimal(d.ErrorRate),
		P95LatencyMs:      null.IntFrom(int64(d.P95LatencyMs)),
		SampleSize:        d.SampleSize,
		WindowStartedAt:   d.WindowStartedAt,
		WindowEndedAt:     d.WindowEndedAt,
		Metadata:          d.Metadata,
	}
}

type ProviderHealthSnapshotsResponse struct {
	Id                uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID       `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodCode        string          `json:"methodCode"`
	ChannelCode       string          `json:"channelCode"`
	HealthScore       int             `json:"healthScore" example:"1"`
	SuccessRate       decimal.Decimal `json:"successRate" format:"decimal" example:"100.50"`
	TimeoutRate       decimal.Decimal `json:"timeoutRate" format:"decimal" example:"100.50"`
	ErrorRate         decimal.Decimal `json:"errorRate" format:"decimal" example:"100.50"`
	P95LatencyMs      int             `json:"p95LatencyMs" example:"1"`
	SampleSize        int             `json:"sampleSize" example:"1"`
	WindowStartedAt   time.Time       `json:"windowStartedAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	WindowEndedAt     time.Time       `json:"windowEndedAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata          json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewProviderHealthSnapshotsResponse(providerHealthSnapshots model.ProviderHealthSnapshots) ProviderHealthSnapshotsResponse {
	return ProviderHealthSnapshotsResponse{
		Id:                providerHealthSnapshots.Id,
		ProviderAccountId: providerHealthSnapshots.ProviderAccountId,
		MethodCode:        providerHealthSnapshots.MethodCode.String,
		ChannelCode:       providerHealthSnapshots.ChannelCode.String,
		HealthScore:       providerHealthSnapshots.HealthScore,
		SuccessRate:       providerHealthSnapshots.SuccessRate.Decimal,
		TimeoutRate:       providerHealthSnapshots.TimeoutRate.Decimal,
		ErrorRate:         providerHealthSnapshots.ErrorRate.Decimal,
		P95LatencyMs:      int(providerHealthSnapshots.P95LatencyMs.ValueOrZero()),
		SampleSize:        providerHealthSnapshots.SampleSize,
		WindowStartedAt:   providerHealthSnapshots.WindowStartedAt,
		WindowEndedAt:     providerHealthSnapshots.WindowEndedAt,
		Metadata:          providerHealthSnapshots.Metadata,
	}
}

type ProviderHealthSnapshotsListResponse []*ProviderHealthSnapshotsResponse

func NewProviderHealthSnapshotsListResponse(providerHealthSnapshotsList model.ProviderHealthSnapshotsList) ProviderHealthSnapshotsListResponse {
	dtoProviderHealthSnapshotsListResponse := ProviderHealthSnapshotsListResponse{}
	for _, providerHealthSnapshots := range providerHealthSnapshotsList {
		dtoProviderHealthSnapshotsResponse := NewProviderHealthSnapshotsResponse(*providerHealthSnapshots)
		dtoProviderHealthSnapshotsListResponse = append(dtoProviderHealthSnapshotsListResponse, &dtoProviderHealthSnapshotsResponse)
	}
	return dtoProviderHealthSnapshotsListResponse
}

type ProviderHealthSnapshotsPrimaryIDList []ProviderHealthSnapshotsPrimaryID

func (d ProviderHealthSnapshotsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerHealthSnapshots := range d {
		err = validator.Struct(providerHealthSnapshots)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderHealthSnapshotsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderHealthSnapshotsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderHealthSnapshotsPrimaryID) ToModel() model.ProviderHealthSnapshotsPrimaryID {
	return model.ProviderHealthSnapshotsPrimaryID{
		Id: d.Id,
	}
}

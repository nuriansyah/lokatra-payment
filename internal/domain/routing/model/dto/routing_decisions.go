package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RoutingDecisionsDTOFieldNameType string

type routingDecisionsDTOFieldName struct {
	Id              RoutingDecisionsDTOFieldNameType
	PaymentIntentId RoutingDecisionsDTOFieldNameType
	ProfileId       RoutingDecisionsDTOFieldNameType
	RuleId          RoutingDecisionsDTOFieldNameType
	StrategyUsed    RoutingDecisionsDTOFieldNameType
	CandidatePsps   RoutingDecisionsDTOFieldNameType
	DecisionReason  RoutingDecisionsDTOFieldNameType
	DecidedAt       RoutingDecisionsDTOFieldNameType
	MetaCreatedAt   RoutingDecisionsDTOFieldNameType
	MetaCreatedBy   RoutingDecisionsDTOFieldNameType
	MetaUpdatedAt   RoutingDecisionsDTOFieldNameType
	MetaUpdatedBy   RoutingDecisionsDTOFieldNameType
	MetaDeletedAt   RoutingDecisionsDTOFieldNameType
	MetaDeletedBy   RoutingDecisionsDTOFieldNameType
}

var RoutingDecisionsDTOFieldName = routingDecisionsDTOFieldName{
	Id:              "id",
	PaymentIntentId: "paymentIntentId",
	ProfileId:       "profileId",
	RuleId:          "ruleId",
	StrategyUsed:    "strategyUsed",
	CandidatePsps:   "candidatePsps",
	DecisionReason:  "decisionReason",
	DecidedAt:       "decidedAt",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func NewRoutingDecisionsListResponseFromFilterResult(result []model.RoutingDecisionsFilterResult, filter model.Filter) RoutingDecisionsSelectableListResponse {
	dtoRoutingDecisionsListResponse := RoutingDecisionsSelectableListResponse{}
	for _, routingDecisions := range result {
		dtoRoutingDecisionsResponse := NewRoutingDecisionsSelectableResponse(routingDecisions.RoutingDecisions, filter)
		dtoRoutingDecisionsListResponse = append(dtoRoutingDecisionsListResponse, &dtoRoutingDecisionsResponse)
	}
	return dtoRoutingDecisionsListResponse
}

func transformRoutingDecisionsDTOFieldNameFromStr(field string) (dbField model.RoutingDecisionsDBFieldNameType, found bool) {
	switch field {

	case string(RoutingDecisionsDTOFieldName.Id):
		return model.RoutingDecisionsDBFieldName.Id, true

	case string(RoutingDecisionsDTOFieldName.PaymentIntentId):
		return model.RoutingDecisionsDBFieldName.PaymentIntentId, true

	case string(RoutingDecisionsDTOFieldName.ProfileId):
		return model.RoutingDecisionsDBFieldName.ProfileId, true

	case string(RoutingDecisionsDTOFieldName.RuleId):
		return model.RoutingDecisionsDBFieldName.RuleId, true

	case string(RoutingDecisionsDTOFieldName.StrategyUsed):
		return model.RoutingDecisionsDBFieldName.StrategyUsed, true

	case string(RoutingDecisionsDTOFieldName.CandidatePsps):
		return model.RoutingDecisionsDBFieldName.CandidatePsps, true

	case string(RoutingDecisionsDTOFieldName.DecisionReason):
		return model.RoutingDecisionsDBFieldName.DecisionReason, true

	case string(RoutingDecisionsDTOFieldName.DecidedAt):
		return model.RoutingDecisionsDBFieldName.DecidedAt, true

	case string(RoutingDecisionsDTOFieldName.MetaCreatedAt):
		return model.RoutingDecisionsDBFieldName.MetaCreatedAt, true

	case string(RoutingDecisionsDTOFieldName.MetaCreatedBy):
		return model.RoutingDecisionsDBFieldName.MetaCreatedBy, true

	case string(RoutingDecisionsDTOFieldName.MetaUpdatedAt):
		return model.RoutingDecisionsDBFieldName.MetaUpdatedAt, true

	case string(RoutingDecisionsDTOFieldName.MetaUpdatedBy):
		return model.RoutingDecisionsDBFieldName.MetaUpdatedBy, true

	case string(RoutingDecisionsDTOFieldName.MetaDeletedAt):
		return model.RoutingDecisionsDBFieldName.MetaDeletedAt, true

	case string(RoutingDecisionsDTOFieldName.MetaDeletedBy):
		return model.RoutingDecisionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformRoutingDecisionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformRoutingDecisionsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRoutingDecisionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRoutingDecisionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultRoutingDecisionsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(RoutingDecisionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RoutingDecisionsSelectableResponse map[string]interface{}
type RoutingDecisionsSelectableListResponse []*RoutingDecisionsSelectableResponse

func NewRoutingDecisionsSelectableResponse(routingDecisions model.RoutingDecisions, filter model.Filter) RoutingDecisionsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RoutingDecisionsDBFieldName.Id),
			string(model.RoutingDecisionsDBFieldName.PaymentIntentId),
			string(model.RoutingDecisionsDBFieldName.ProfileId),
			string(model.RoutingDecisionsDBFieldName.RuleId),
			string(model.RoutingDecisionsDBFieldName.StrategyUsed),
			string(model.RoutingDecisionsDBFieldName.CandidatePsps),
			string(model.RoutingDecisionsDBFieldName.DecisionReason),
			string(model.RoutingDecisionsDBFieldName.DecidedAt),
			string(model.RoutingDecisionsDBFieldName.MetaCreatedAt),
			string(model.RoutingDecisionsDBFieldName.MetaCreatedBy),
			string(model.RoutingDecisionsDBFieldName.MetaUpdatedAt),
			string(model.RoutingDecisionsDBFieldName.MetaUpdatedBy),
			string(model.RoutingDecisionsDBFieldName.MetaDeletedAt),
			string(model.RoutingDecisionsDBFieldName.MetaDeletedBy),
		)
	}
	routingDecisionsSelectableResponse := RoutingDecisionsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.RoutingDecisionsDBFieldName.Id):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.Id)] = routingDecisions.Id

		case string(model.RoutingDecisionsDBFieldName.PaymentIntentId):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.PaymentIntentId)] = routingDecisions.PaymentIntentId

		case string(model.RoutingDecisionsDBFieldName.ProfileId):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.ProfileId)] = routingDecisions.ProfileId

		case string(model.RoutingDecisionsDBFieldName.RuleId):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.RuleId)] = routingDecisions.RuleId

		case string(model.RoutingDecisionsDBFieldName.StrategyUsed):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.StrategyUsed)] = routingDecisions.StrategyUsed

		case string(model.RoutingDecisionsDBFieldName.CandidatePsps):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.CandidatePsps)] = routingDecisions.CandidatePsps

		case string(model.RoutingDecisionsDBFieldName.DecisionReason):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.DecisionReason)] = routingDecisions.DecisionReason

		case string(model.RoutingDecisionsDBFieldName.DecidedAt):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.DecidedAt)] = routingDecisions.DecidedAt

		case string(model.RoutingDecisionsDBFieldName.MetaCreatedAt):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaCreatedAt)] = routingDecisions.MetaCreatedAt

		case string(model.RoutingDecisionsDBFieldName.MetaCreatedBy):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaCreatedBy)] = routingDecisions.MetaCreatedBy

		case string(model.RoutingDecisionsDBFieldName.MetaUpdatedAt):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaUpdatedAt)] = routingDecisions.MetaUpdatedAt

		case string(model.RoutingDecisionsDBFieldName.MetaUpdatedBy):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaUpdatedBy)] = routingDecisions.MetaUpdatedBy

		case string(model.RoutingDecisionsDBFieldName.MetaDeletedAt):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaDeletedAt)] = routingDecisions.MetaDeletedAt

		case string(model.RoutingDecisionsDBFieldName.MetaDeletedBy):
			routingDecisionsSelectableResponse[string(RoutingDecisionsDTOFieldName.MetaDeletedBy)] = routingDecisions.MetaDeletedBy

		}
	}
	return routingDecisionsSelectableResponse
}

type RoutingDecisionsFilterResponse struct {
	Metadata Metadata                               `json:"metadata"`
	Data     RoutingDecisionsSelectableListResponse `json:"data"`
}

func NewRoutingDecisionsFilterResponse(result []model.RoutingDecisionsFilterResult, filter model.Filter) (resp RoutingDecisionsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewRoutingDecisionsListResponseFromFilterResult(result, filter)
	return resp
}

type RoutingDecisionsCreateRequest struct {
	PaymentIntentId uuid.UUID             `json:"paymentIntentId"`
	ProfileId       uuid.UUID             `json:"profileId"`
	RuleId          uuid.UUID             `json:"ruleId"`
	StrategyUsed    model.RoutingStrategy `json:"strategyUsed" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	CandidatePsps   json.RawMessage       `json:"candidatePsps"`
	DecisionReason  string                `json:"decisionReason"`
	DecidedAt       time.Time             `json:"decidedAt"`
	MetaCreatedAt   time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID             `json:"metaDeletedBy"`
}

func (d *RoutingDecisionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RoutingDecisionsCreateRequest) ToModel() model.RoutingDecisions {
	id, _ := uuid.NewV4()
	return model.RoutingDecisions{
		Id:              id,
		PaymentIntentId: d.PaymentIntentId,
		ProfileId:       nuuid.From(d.ProfileId),
		RuleId:          nuuid.From(d.RuleId),
		StrategyUsed:    d.StrategyUsed,
		CandidatePsps:   d.CandidatePsps,
		DecisionReason:  null.StringFrom(d.DecisionReason),
		DecidedAt:       d.DecidedAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingDecisionsListCreateRequest []*RoutingDecisionsCreateRequest

func (d RoutingDecisionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingDecisions := range d {
		err = validator.Struct(routingDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingDecisionsListCreateRequest) ToModelList() []model.RoutingDecisions {
	out := make([]model.RoutingDecisions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RoutingDecisionsUpdateRequest struct {
	PaymentIntentId uuid.UUID             `json:"paymentIntentId"`
	ProfileId       uuid.UUID             `json:"profileId"`
	RuleId          uuid.UUID             `json:"ruleId"`
	StrategyUsed    model.RoutingStrategy `json:"strategyUsed" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	CandidatePsps   json.RawMessage       `json:"candidatePsps"`
	DecisionReason  string                `json:"decisionReason"`
	DecidedAt       time.Time             `json:"decidedAt"`
	MetaCreatedAt   time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID             `json:"metaDeletedBy"`
}

func (d *RoutingDecisionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RoutingDecisionsUpdateRequest) ToModel() model.RoutingDecisions {
	return model.RoutingDecisions{
		PaymentIntentId: d.PaymentIntentId,
		ProfileId:       nuuid.From(d.ProfileId),
		RuleId:          nuuid.From(d.RuleId),
		StrategyUsed:    d.StrategyUsed,
		CandidatePsps:   d.CandidatePsps,
		DecisionReason:  null.StringFrom(d.DecisionReason),
		DecidedAt:       d.DecidedAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingDecisionsBulkUpdateRequest struct {
	Id              uuid.UUID             `json:"id"`
	PaymentIntentId uuid.UUID             `json:"paymentIntentId"`
	ProfileId       uuid.UUID             `json:"profileId"`
	RuleId          uuid.UUID             `json:"ruleId"`
	StrategyUsed    model.RoutingStrategy `json:"strategyUsed" example:"LOWEST_COST" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	CandidatePsps   json.RawMessage       `json:"candidatePsps"`
	DecisionReason  string                `json:"decisionReason"`
	DecidedAt       time.Time             `json:"decidedAt"`
	MetaCreatedAt   time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID             `json:"metaDeletedBy"`
}

func (d RoutingDecisionsBulkUpdateRequest) PrimaryID() RoutingDecisionsPrimaryID {
	return RoutingDecisionsPrimaryID{
		Id: d.Id,
	}
}

type RoutingDecisionsListBulkUpdateRequest []*RoutingDecisionsBulkUpdateRequest

func (d RoutingDecisionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingDecisions := range d {
		err = validator.Struct(routingDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingDecisionsBulkUpdateRequest) ToModel() model.RoutingDecisions {
	return model.RoutingDecisions{
		Id:              d.Id,
		PaymentIntentId: d.PaymentIntentId,
		ProfileId:       nuuid.From(d.ProfileId),
		RuleId:          nuuid.From(d.RuleId),
		StrategyUsed:    d.StrategyUsed,
		CandidatePsps:   d.CandidatePsps,
		DecisionReason:  null.StringFrom(d.DecisionReason),
		DecidedAt:       d.DecidedAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingDecisionsResponse struct {
	Id              uuid.UUID             `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId uuid.UUID             `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProfileId       uuid.UUID             `json:"profileId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RuleId          uuid.UUID             `json:"ruleId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	StrategyUsed    model.RoutingStrategy `json:"strategyUsed" validate:"required,oneof=LOWEST_COST HIGHEST_SUCCESS_RATE ROUND_ROBIN GEO_PREFERRED MANUAL WATERFALL" enums:"LOWEST_COST,HIGHEST_SUCCESS_RATE,ROUND_ROBIN,GEO_PREFERRED,MANUAL,WATERFALL"`
	CandidatePsps   json.RawMessage       `json:"candidatePsps" swaggertype:"object"`
	DecisionReason  string                `json:"decisionReason"`
	DecidedAt       time.Time             `json:"decidedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedAt   time.Time             `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy   uuid.UUID             `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt   time.Time             `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy   uuid.UUID             `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt   time.Time             `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy   uuid.UUID             `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewRoutingDecisionsResponse(routingDecisions model.RoutingDecisions) RoutingDecisionsResponse {
	return RoutingDecisionsResponse{
		Id:              routingDecisions.Id,
		PaymentIntentId: routingDecisions.PaymentIntentId,
		ProfileId:       routingDecisions.ProfileId.UUID,
		RuleId:          routingDecisions.RuleId.UUID,
		StrategyUsed:    model.RoutingStrategy(routingDecisions.StrategyUsed),
		CandidatePsps:   routingDecisions.CandidatePsps,
		DecisionReason:  routingDecisions.DecisionReason.String,
		DecidedAt:       routingDecisions.DecidedAt,
		MetaCreatedAt:   routingDecisions.MetaCreatedAt,
		MetaCreatedBy:   routingDecisions.MetaCreatedBy,
		MetaUpdatedAt:   routingDecisions.MetaUpdatedAt,
		MetaUpdatedBy:   routingDecisions.MetaUpdatedBy.UUID,
		MetaDeletedAt:   routingDecisions.MetaDeletedAt.Time,
		MetaDeletedBy:   routingDecisions.MetaDeletedBy.UUID,
	}
}

type RoutingDecisionsListResponse []*RoutingDecisionsResponse

func NewRoutingDecisionsListResponse(routingDecisionsList model.RoutingDecisionsList) RoutingDecisionsListResponse {
	dtoRoutingDecisionsListResponse := RoutingDecisionsListResponse{}
	for _, routingDecisions := range routingDecisionsList {
		dtoRoutingDecisionsResponse := NewRoutingDecisionsResponse(*routingDecisions)
		dtoRoutingDecisionsListResponse = append(dtoRoutingDecisionsListResponse, &dtoRoutingDecisionsResponse)
	}
	return dtoRoutingDecisionsListResponse
}

type RoutingDecisionsPrimaryIDList []RoutingDecisionsPrimaryID

func (d RoutingDecisionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingDecisions := range d {
		err = validator.Struct(routingDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

type RoutingDecisionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RoutingDecisionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RoutingDecisionsPrimaryID) ToModel() model.RoutingDecisionsPrimaryID {
	return model.RoutingDecisionsPrimaryID{
		Id: d.Id,
	}
}

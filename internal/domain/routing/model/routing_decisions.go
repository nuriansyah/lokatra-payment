package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type RoutingDecisionsDBFieldNameType string

type routingDecisionsDBFieldName struct {
	Id              RoutingDecisionsDBFieldNameType
	PaymentIntentId RoutingDecisionsDBFieldNameType
	ProfileId       RoutingDecisionsDBFieldNameType
	RuleId          RoutingDecisionsDBFieldNameType
	StrategyUsed    RoutingDecisionsDBFieldNameType
	CandidatePsps   RoutingDecisionsDBFieldNameType
	DecisionReason  RoutingDecisionsDBFieldNameType
	DecidedAt       RoutingDecisionsDBFieldNameType
	MetaCreatedAt   RoutingDecisionsDBFieldNameType
	MetaCreatedBy   RoutingDecisionsDBFieldNameType
	MetaUpdatedAt   RoutingDecisionsDBFieldNameType
	MetaUpdatedBy   RoutingDecisionsDBFieldNameType
	MetaDeletedAt   RoutingDecisionsDBFieldNameType
	MetaDeletedBy   RoutingDecisionsDBFieldNameType
}

var RoutingDecisionsDBFieldName = routingDecisionsDBFieldName{
	Id:              "id",
	PaymentIntentId: "payment_intent_id",
	ProfileId:       "profile_id",
	RuleId:          "rule_id",
	StrategyUsed:    "strategy_used",
	CandidatePsps:   "candidate_psps",
	DecisionReason:  "decision_reason",
	DecidedAt:       "decided_at",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewRoutingDecisionsDBFieldNameFromStr(field string) (dbField RoutingDecisionsDBFieldNameType, found bool) {
	switch field {

	case string(RoutingDecisionsDBFieldName.Id):
		return RoutingDecisionsDBFieldName.Id, true

	case string(RoutingDecisionsDBFieldName.PaymentIntentId):
		return RoutingDecisionsDBFieldName.PaymentIntentId, true

	case string(RoutingDecisionsDBFieldName.ProfileId):
		return RoutingDecisionsDBFieldName.ProfileId, true

	case string(RoutingDecisionsDBFieldName.RuleId):
		return RoutingDecisionsDBFieldName.RuleId, true

	case string(RoutingDecisionsDBFieldName.StrategyUsed):
		return RoutingDecisionsDBFieldName.StrategyUsed, true

	case string(RoutingDecisionsDBFieldName.CandidatePsps):
		return RoutingDecisionsDBFieldName.CandidatePsps, true

	case string(RoutingDecisionsDBFieldName.DecisionReason):
		return RoutingDecisionsDBFieldName.DecisionReason, true

	case string(RoutingDecisionsDBFieldName.DecidedAt):
		return RoutingDecisionsDBFieldName.DecidedAt, true

	case string(RoutingDecisionsDBFieldName.MetaCreatedAt):
		return RoutingDecisionsDBFieldName.MetaCreatedAt, true

	case string(RoutingDecisionsDBFieldName.MetaCreatedBy):
		return RoutingDecisionsDBFieldName.MetaCreatedBy, true

	case string(RoutingDecisionsDBFieldName.MetaUpdatedAt):
		return RoutingDecisionsDBFieldName.MetaUpdatedAt, true

	case string(RoutingDecisionsDBFieldName.MetaUpdatedBy):
		return RoutingDecisionsDBFieldName.MetaUpdatedBy, true

	case string(RoutingDecisionsDBFieldName.MetaDeletedAt):
		return RoutingDecisionsDBFieldName.MetaDeletedAt, true

	case string(RoutingDecisionsDBFieldName.MetaDeletedBy):
		return RoutingDecisionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type RoutingDecisionsFilterResult struct {
	RoutingDecisions
	FilterCount int `db:"count"`
}

func ValidateRoutingDecisionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewRoutingDecisionsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewRoutingDecisionsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewRoutingDecisionsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type RoutingStrategy string

const (
	RoutingStrategyLowestCost         RoutingStrategy = "LOWEST_COST"
	RoutingStrategyHighestSuccessRate RoutingStrategy = "HIGHEST_SUCCESS_RATE"
	RoutingStrategyRoundRobin         RoutingStrategy = "ROUND_ROBIN"
	RoutingStrategyGeoPreferred       RoutingStrategy = "GEO_PREFERRED"
	RoutingStrategyManual             RoutingStrategy = "MANUAL"
	RoutingStrategyWaterfall          RoutingStrategy = "WATERFALL"
)

type RoutingDecisions struct {
	Id              uuid.UUID       `db:"id"`
	PaymentIntentId uuid.UUID       `db:"payment_intent_id"`
	ProfileId       nuuid.NUUID     `db:"profile_id"`
	RuleId          nuuid.NUUID     `db:"rule_id"`
	StrategyUsed    RoutingStrategy `db:"strategy_used"`
	CandidatePsps   json.RawMessage `db:"candidate_psps"`
	DecisionReason  null.String     `db:"decision_reason"`
	DecidedAt       time.Time       `db:"decided_at"`
	MetaCreatedAt   time.Time       `db:"meta_created_at"`
	MetaCreatedBy   uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt   time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy   nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt   null.Time       `db:"meta_deleted_at"`
	MetaDeletedBy   nuuid.NUUID     `db:"meta_deleted_by"`
}
type RoutingDecisionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RoutingDecisions) ToRoutingDecisionsPrimaryID() RoutingDecisionsPrimaryID {
	return RoutingDecisionsPrimaryID{
		Id: d.Id,
	}
}

type RoutingDecisionsList []*RoutingDecisions

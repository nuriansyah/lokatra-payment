package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// Enums
// ---------------------------------------------------------------------------

type PayrouteRoutingStrategy string

const (
	PayrouteStrategyPriority   PayrouteRoutingStrategy = "priority"
	PayrouteStrategySuccessRate PayrouteRoutingStrategy = "success_rate"
	PayrouteStrategyCostAware  PayrouteRoutingStrategy = "cost_aware"
	PayrouteStrategyWeighted   PayrouteRoutingStrategy = "weighted"
	PayrouteStrategyCombined   PayrouteRoutingStrategy = "combined"
)

type PayrouteRuleStatus string

const (
	PayrouteStatusDraft      PayrouteRuleStatus = "draft"
	PayrouteStatusActive     PayrouteRuleStatus = "active"
	PayrouteStatusRollingOut PayrouteRuleStatus = "rolling_out"
	PayrouteStatusDisabled   PayrouteRuleStatus = "disabled"
	PayrouteStatusArchived   PayrouteRuleStatus = "archived"
)

type PayrouteRuleScope string

const (
	PayrouteScopeGlobal         PayrouteRuleScope = "global"
	PayrouteScopeMethod         PayrouteRuleScope = "method"
	PayrouteScopeMethodChannel  PayrouteRuleScope = "method_channel"
	PayrouteScopeMethodAmount   PayrouteRuleScope = "method_amount_band"
)

// ---------------------------------------------------------------------------
// Scope Config
// ---------------------------------------------------------------------------

type RoutingScopeConfig struct {
	MethodCode  string `json:"method_code,omitempty"`
	ChannelCode string `json:"channel_code,omitempty"`
	Currency    string `json:"currency,omitempty"`
	MinAmount   string `json:"min_amount,omitempty"`
	MaxAmount   string `json:"max_amount,omitempty"`
}

// ---------------------------------------------------------------------------
// PSP List Item
// ---------------------------------------------------------------------------

type RoutingPSPItem struct {
	ProviderAccountID   uuid.UUID `json:"provider_account_id"`
	Priority            int       `json:"priority"`
	TrafficWeight       int       `json:"traffic_weight"`
	IsFallback          bool      `json:"is_fallback"`
	MaxAttempts         int       `json:"max_attempts"`
	TimeoutMs           int       `json:"timeout_ms"`
	ProviderMethodCode  string    `json:"provider_method_code,omitempty"`
	ProviderChannelCode string    `json:"provider_channel_code,omitempty"`
}

// ---------------------------------------------------------------------------
// Strategy Config
// ---------------------------------------------------------------------------

type RoutingStrategyConfig struct {
	WindowSeconds      int             `json:"window_seconds,omitempty"`
	MinSampleSize      int             `json:"min_sample_size,omitempty"`
	MdrWeight          decimal.Decimal `json:"mdr_weight,omitempty"`
	SuccessRateWeight  decimal.Decimal `json:"success_rate_weight,omitempty"`
	ShadowPercentage   int             `json:"shadow_percentage,omitempty"`
}

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type PayrouteRoutingRules struct {
	Id                    uuid.UUID               `db:"id"`
	Name                  string                  `db:"name"`
	Description           null.String             `db:"description"`
	Scope                 PayrouteRuleScope       `db:"scope"`
	ScopeConfig           json.RawMessage         `db:"scope_config"`
	Strategy              PayrouteRoutingStrategy  `db:"strategy"`
	StrategyConfig        json.RawMessage         `db:"strategy_config"`
	PspList               json.RawMessage         `db:"sp_list"`
	Version               int                     `db:"version"`
	Status                PayrouteRuleStatus      `db:"status"`
	RolloutPercentage     int                     `db:"rollout_percentage"`
	RolloutStartedAt      null.Time               `db:"rollout_started_at"`
	RolloutRollbackVersion null.Int               `db:"rollout_rollback_version"`
	CreatedBy             uuid.UUID               `db:"created_by"`
	ApprovedBy            *uuid.UUID              `db:"approved_by"`

	shared.MetaSignature
}

type PayrouteRoutingRulesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayrouteRoutingRules) ToPayrouteRoutingRulesPrimaryID() PayrouteRoutingRulesPrimaryID {
	return PayrouteRoutingRulesPrimaryID{Id: d.Id}
}

// ParsePspList deserializes the JSONB psp_list into typed structs.
func (d PayrouteRoutingRules) ParsePspList() ([]RoutingPSPItem, error) {
	var items []RoutingPSPItem
	if len(d.PspList) == 0 {
		return items, nil
	}
	if err := json.Unmarshal(d.PspList, &items); err != nil {
		return nil, fmt.Errorf("invalid psp_list: %w", err)
	}
	return items, nil
}

// ParseScopeConfig deserializes the JSONB scope_config.
func (d PayrouteRoutingRules) ParseScopeConfig() (RoutingScopeConfig, error) {
	var cfg RoutingScopeConfig
	if len(d.ScopeConfig) == 0 {
		return cfg, nil
	}
	if err := json.Unmarshal(d.ScopeConfig, &cfg); err != nil {
		return cfg, fmt.Errorf("invalid scope_config: %w", err)
	}
	return cfg, nil
}

// ParseStrategyConfig deserializes the JSONB strategy_config.
func (d PayrouteRoutingRules) ParseStrategyConfig() (RoutingStrategyConfig, error) {
	var cfg RoutingStrategyConfig
	if len(d.StrategyConfig) == 0 {
		return cfg, nil
	}
	if err := json.Unmarshal(d.StrategyConfig, &cfg); err != nil {
		return cfg, fmt.Errorf("invalid strategy_config: %w", err)
	}
	return cfg, nil
}

// ---------------------------------------------------------------------------
// DB Field Names — compile-time safe column references
// ---------------------------------------------------------------------------

type PayrouteRoutingRulesDBFieldNameType string

type payrouteRoutingRulesDBFieldName struct {
	Id                    PayrouteRoutingRulesDBFieldNameType
	Name                  PayrouteRoutingRulesDBFieldNameType
	Description           PayrouteRoutingRulesDBFieldNameType
	Scope                 PayrouteRoutingRulesDBFieldNameType
	ScopeConfig           PayrouteRoutingRulesDBFieldNameType
	Strategy              PayrouteRoutingRulesDBFieldNameType
	StrategyConfig        PayrouteRoutingRulesDBFieldNameType
	PspList               PayrouteRoutingRulesDBFieldNameType
	Version               PayrouteRoutingRulesDBFieldNameType
	Status                PayrouteRoutingRulesDBFieldNameType
	RolloutPercentage     PayrouteRoutingRulesDBFieldNameType
	RolloutStartedAt      PayrouteRoutingRulesDBFieldNameType
	RolloutRollbackVersion PayrouteRoutingRulesDBFieldNameType
	CreatedBy             PayrouteRoutingRulesDBFieldNameType
	ApprovedBy            PayrouteRoutingRulesDBFieldNameType
	MetaCreatedAt         PayrouteRoutingRulesDBFieldNameType
	MetaCreatedBy         PayrouteRoutingRulesDBFieldNameType
	MetaUpdatedAt         PayrouteRoutingRulesDBFieldNameType
	MetaUpdatedBy         PayrouteRoutingRulesDBFieldNameType
	MetaDeletedAt         PayrouteRoutingRulesDBFieldNameType
	MetaDeletedBy         PayrouteRoutingRulesDBFieldNameType
}

var PayrouteRoutingRulesDBFieldName = payrouteRoutingRulesDBFieldName{
	Id:                    "id",
	Name:                  "name",
	Description:           "description",
	Scope:                 "scope",
	ScopeConfig:           "scope_config",
	Strategy:              "strategy",
	StrategyConfig:        "strategy_config",
	PspList:               "psp_list",
	Version:               "version",
	Status:                "status",
	RolloutPercentage:     "rollout_percentage",
	RolloutStartedAt:      "rollout_started_at",
	RolloutRollbackVersion: "rollout_rollback_version",
	CreatedBy:             "created_by",
	ApprovedBy:            "approved_by",
	MetaCreatedAt:         "meta_created_at",
	MetaCreatedBy:         "meta_created_by",
	MetaUpdatedAt:         "meta_updated_at",
	MetaUpdatedBy:         "meta_updated_by",
	MetaDeletedAt:         "meta_deleted_at",
	MetaDeletedBy:         "meta_deleted_by",
}

func NewPayrouteRoutingRulesDBFieldNameFromStr(field string) (dbField PayrouteRoutingRulesDBFieldNameType, found bool) {
	switch field {
	case string(PayrouteRoutingRulesDBFieldName.Id):
		return PayrouteRoutingRulesDBFieldName.Id, true
	case string(PayrouteRoutingRulesDBFieldName.Name):
		return PayrouteRoutingRulesDBFieldName.Name, true
	case string(PayrouteRoutingRulesDBFieldName.Scope):
		return PayrouteRoutingRulesDBFieldName.Scope, true
	case string(PayrouteRoutingRulesDBFieldName.Strategy):
		return PayrouteRoutingRulesDBFieldName.Strategy, true
	case string(PayrouteRoutingRulesDBFieldName.Status):
		return PayrouteRoutingRulesDBFieldName.Status, true
	case string(PayrouteRoutingRulesDBFieldName.Version):
		return PayrouteRoutingRulesDBFieldName.Version, true
	case string(PayrouteRoutingRulesDBFieldName.MetaCreatedAt):
		return PayrouteRoutingRulesDBFieldName.MetaCreatedAt, true
	case string(PayrouteRoutingRulesDBFieldName.MetaUpdatedAt):
		return PayrouteRoutingRulesDBFieldName.MetaUpdatedAt, true
	}
	return "unknown", false
}

// ---------------------------------------------------------------------------
// Filter & Query Support
// ---------------------------------------------------------------------------

var PayrouteRoutingRulesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath: "id", DefaultOutputPath: "id", Column: "id",
		SQLAlias: "id", Selectable: true, Filterable: true, Sortable: true,
	},
	"name": {
		SourcePath: "name", DefaultOutputPath: "name", Column: "name",
		SQLAlias: "name", Selectable: true, Filterable: true, Sortable: true,
	},
	"scope": {
		SourcePath: "scope", DefaultOutputPath: "scope", Column: "scope",
		SQLAlias: "scope", Selectable: true, Filterable: true, Sortable: true,
	},
	"strategy": {
		SourcePath: "strategy", DefaultOutputPath: "strategy", Column: "strategy",
		SQLAlias: "strategy", Selectable: true, Filterable: true, Sortable: true,
	},
	"status": {
		SourcePath: "status", DefaultOutputPath: "status", Column: "status",
		SQLAlias: "status", Selectable: true, Filterable: true, Sortable: true,
	},
	"version": {
		SourcePath: "version", DefaultOutputPath: "version", Column: "version",
		SQLAlias: "version", Selectable: true, Filterable: true, Sortable: true,
	},
	"created_by": {
		SourcePath: "created_by", DefaultOutputPath: "createdBy", Column: "created_by",
		SQLAlias: "created_by", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_created_at": {
		SourcePath: "meta_created_at", DefaultOutputPath: "metaCreatedAt", Column: "meta_created_at",
		SQLAlias: "meta_created_at", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_updated_at": {
		SourcePath: "meta_updated_at", DefaultOutputPath: "metaUpdatedAt", Column: "meta_updated_at",
		SQLAlias: "meta_updated_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayrouteRoutingRulesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayrouteRoutingRulesFilterFields[field]
	return
}

func ValidatePayrouteRoutingRulesFieldNameFilter(filter Filter) error {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayrouteRoutingRulesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			return failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayrouteRoutingRulesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			return failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayrouteRoutingRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// List types
// ---------------------------------------------------------------------------

type PayrouteRoutingRulesList []*PayrouteRoutingRules

type PayrouteRoutingRulesFilterResult struct {
	PayrouteRoutingRules
	FilterCount int `db:"count"`
}

type PayrouteRoutingRulesFilterResultList []*PayrouteRoutingRulesFilterResult

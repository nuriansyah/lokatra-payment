package dto

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// PayRoute Routing Rules DTOs
// ---------------------------------------------------------------------------

type CreatePayrouteRoutingRuleRequest struct {
	Name           string          `json:"name" validate:"required"`
	Description    string          `json:"description,omitempty"`
	Scope          string          `json:"scope" validate:"required,oneof=global method method_channel method_amount_band"`
	ScopeConfig    json.RawMessage `json:"scope_config,omitempty"`
	Strategy       string          `json:"strategy" validate:"required,oneof=priority success_rate cost_aware weighted combined"`
	StrategyConfig json.RawMessage `json:"strategy_config,omitempty"`
	PspList        json.RawMessage `json:"psp_list" validate:"required"`
}

type UpdatePayrouteRoutingRuleRequest struct {
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	Scope          string          `json:"scope,omitempty"`
	ScopeConfig    json.RawMessage `json:"scope_config,omitempty"`
	Strategy       string          `json:"strategy,omitempty"`
	StrategyConfig json.RawMessage `json:"strategy_config,omitempty"`
	PspList        json.RawMessage `json:"psp_list,omitempty"`
}

type PayrouteRoutingRuleResponse struct {
	Id                uuid.UUID       `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description,omitempty"`
	Scope             string          `json:"scope"`
	ScopeConfig       json.RawMessage `json:"scope_config"`
	Strategy          string          `json:"strategy"`
	StrategyConfig    json.RawMessage `json:"strategy_config"`
	PspList           json.RawMessage `json:"psp_list"`
	Version           int             `json:"version"`
	Status            string          `json:"status"`
	RolloutPercentage int             `json:"rollout_percentage"`
	CreatedBy         uuid.UUID       `json:"created_by"`
	MetaCreatedAt     time.Time       `json:"meta_created_at"`
	MetaUpdatedAt     time.Time       `json:"meta_updated_at"`
}

// ---------------------------------------------------------------------------
// MDR Rate DTOs
// ---------------------------------------------------------------------------

type UpsertMDRRateRequest struct {
	ProviderAccountID string          `json:"provider_account_id" validate:"required"`
	PaymentMethodID   string          `json:"payment_method_id" validate:"required"`
	PaymentChannelID  string          `json:"payment_channel_id,omitempty"`
	Percentage        decimal.Decimal `json:"percentage" validate:"required"`
	FixedFee          decimal.Decimal `json:"fixed_fee" validate:"required"`
	Currency          string          `json:"currency" validate:"required"`
}

type MDRRateResponse struct {
	Id                uuid.UUID       `json:"id"`
	ProviderAccountID uuid.UUID       `json:"provider_account_id"`
	PaymentMethodID   uuid.UUID       `json:"payment_method_id"`
	PaymentChannelID  *uuid.UUID      `json:"payment_channel_id"`
	Percentage        decimal.Decimal `json:"percentage"`
	FixedFee          decimal.Decimal `json:"fixed_fee"`
	Currency          string          `json:"currency"`
	EffectiveFrom     time.Time       `json:"effective_from"`
	EffectiveTo       null.Time       `json:"effective_to"`
}

// ---------------------------------------------------------------------------
// PSP Health DTOs
// ---------------------------------------------------------------------------

type PSPHealthResponse struct {
	ProviderAccountID uuid.UUID       `json:"provider_account_id"`
	MethodCode        string          `json:"method_code"`
	ChannelCode       string          `json:"channel_code"`
	Status            string          `json:"status"`
	SuccessRate       decimal.Decimal `json:"success_rate"`
	AvgLatencyMs      int             `json:"avg_latency_ms"`
	SampleSize        int             `json:"sample_size"`
	WindowStart       time.Time       `json:"window_start"`
	WindowEnd         time.Time       `json:"window_end"`
}

type KillSwitchRequest struct {
	Reason string `json:"reason" validate:"required"`
}

// ---------------------------------------------------------------------------
// Audit Log DTOs
// ---------------------------------------------------------------------------

type AuditLogResponse struct {
	Id           uuid.UUID       `json:"id"`
	ActorUserId  uuid.UUID       `json:"actor_user_id"`
	ActorEmail   string          `json:"actor_email,omitempty"`
	Action       string          `json:"action"`
	TargetEntity string          `json:"target_entity"`
	TargetId     uuid.UUID       `json:"target_id"`
	BeforeValue  json.RawMessage `json:"before_value"`
	AfterValue   json.RawMessage `json:"after_value"`
	Reason       string          `json:"reason,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Config DTOs
// ---------------------------------------------------------------------------

type UpdateConfigRequest struct {
	Value       json.RawMessage `json:"value" validate:"required"`
	Description string          `json:"description,omitempty"`
}

type ConfigResponse struct {
	Key           string          `json:"key"`
	Value         json.RawMessage `json:"value"`
	Description   string          `json:"description,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// ---------------------------------------------------------------------------
// Interface
// ---------------------------------------------------------------------------

type PayrouteRoutingRulesRepository interface {
	CreatePayrouteRoutingRules(ctx context.Context, item *model.PayrouteRoutingRules) error
	ResolvePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID) (model.PayrouteRoutingRules, error)
	ResolvePayrouteRoutingRulesByFilter(ctx context.Context, filter model.Filter) ([]model.PayrouteRoutingRulesFilterResult, error)
	UpdatePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID, item *model.PayrouteRoutingRules, fields ...PayrouteRoutingRulesUpdateField) error
	DeletePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID) error
	ResolveActiveRuleByScope(ctx context.Context, methodCode, channelCode, currency string) (*model.PayrouteRoutingRules, error)
	ResolveConflictingRulesByScope(ctx context.Context, scope string, scopeConfig map[string]interface{}, excludeID interface{}) (int, error)
}

// ---------------------------------------------------------------------------
// Field Types
// ---------------------------------------------------------------------------

type PayrouteRoutingRulesField string

type PayrouteRoutingRulesUpdateField struct {
	field PayrouteRoutingRulesField
	value interface{}
}

type PayrouteRoutingRulesUpdateFieldList []PayrouteRoutingRulesUpdateField

func NewPayrouteRoutingRulesUpdateField(field PayrouteRoutingRulesField, val interface{}) PayrouteRoutingRulesUpdateField {
	return PayrouteRoutingRulesUpdateField{field: field, value: val}
}

// ---------------------------------------------------------------------------
// Field Constants
// ---------------------------------------------------------------------------

const (
	PayrouteRoutingRulesFieldId                    PayrouteRoutingRulesField = "id"
	PayrouteRoutingRulesFieldName                  PayrouteRoutingRulesField = "name"
	PayrouteRoutingRulesFieldDescription           PayrouteRoutingRulesField = "description"
	PayrouteRoutingRulesFieldScope                 PayrouteRoutingRulesField = "scope"
	PayrouteRoutingRulesFieldScopeConfig           PayrouteRoutingRulesField = "scope_config"
	PayrouteRoutingRulesFieldStrategy              PayrouteRoutingRulesField = "strategy"
	PayrouteRoutingRulesFieldStrategyConfig        PayrouteRoutingRulesField = "strategy_config"
	PayrouteRoutingRulesFieldPspList               PayrouteRoutingRulesField = "psp_list"
	PayrouteRoutingRulesFieldVersion               PayrouteRoutingRulesField = "version"
	PayrouteRoutingRulesFieldStatus                PayrouteRoutingRulesField = "status"
	PayrouteRoutingRulesFieldRolloutPercentage     PayrouteRoutingRulesField = "rollout_percentage"
	PayrouteRoutingRulesFieldRolloutStartedAt      PayrouteRoutingRulesField = "rollout_started_at"
	PayrouteRoutingRulesFieldRolloutRollbackVersion PayrouteRoutingRulesField = "rollout_rollback_version"
	PayrouteRoutingRulesFieldCreatedBy             PayrouteRoutingRulesField = "created_by"
	PayrouteRoutingRulesFieldApprovedBy            PayrouteRoutingRulesField = "approved_by"
	PayrouteRoutingRulesFieldMetaCreatedAt         PayrouteRoutingRulesField = "meta_created_at"
	PayrouteRoutingRulesFieldMetaCreatedBy         PayrouteRoutingRulesField = "meta_created_by"
	PayrouteRoutingRulesFieldMetaUpdatedAt         PayrouteRoutingRulesField = "meta_updated_at"
	PayrouteRoutingRulesFieldMetaUpdatedBy         PayrouteRoutingRulesField = "meta_updated_by"
	PayrouteRoutingRulesFieldMetaDeletedAt         PayrouteRoutingRulesField = "meta_deleted_at"
	PayrouteRoutingRulesFieldMetaDeletedBy         PayrouteRoutingRulesField = "meta_deleted_by"
)

// Default insert fields
var payrouteRoutingRulesDefaultInsertFields = []PayrouteRoutingRulesField{
	PayrouteRoutingRulesFieldId,
	PayrouteRoutingRulesFieldName,
	PayrouteRoutingRulesFieldDescription,
	PayrouteRoutingRulesFieldScope,
	PayrouteRoutingRulesFieldScopeConfig,
	PayrouteRoutingRulesFieldStrategy,
	PayrouteRoutingRulesFieldStrategyConfig,
	PayrouteRoutingRulesFieldPspList,
	PayrouteRoutingRulesFieldVersion,
	PayrouteRoutingRulesFieldStatus,
	PayrouteRoutingRulesFieldRolloutPercentage,
	PayrouteRoutingRulesFieldRolloutStartedAt,
	PayrouteRoutingRulesFieldRolloutRollbackVersion,
	PayrouteRoutingRulesFieldCreatedBy,
	PayrouteRoutingRulesFieldApprovedBy,
	PayrouteRoutingRulesFieldMetaCreatedAt,
	PayrouteRoutingRulesFieldMetaCreatedBy,
	PayrouteRoutingRulesFieldMetaUpdatedAt,
	PayrouteRoutingRulesFieldMetaUpdatedBy,
	PayrouteRoutingRulesFieldMetaDeletedAt,
	PayrouteRoutingRulesFieldMetaDeletedBy,
}

// ---------------------------------------------------------------------------
// SQL Helpers
// ---------------------------------------------------------------------------

func defaultPayrouteRoutingRulesSelectFields() string {
	return `"id","name","description","scope","scope_config","strategy","strategy_config","psp_list",
		"version","status","rollout_percentage","rollout_started_at","rollout_rollback_version",
		"created_by","approved_by",
		"meta_created_at","meta_created_by","meta_updated_at","meta_updated_by","meta_deleted_at","meta_deleted_by"`
}

func composePayrouteRoutingRulesInsertFieldsAndValues(items []*model.PayrouteRoutingRules) (fieldStr string, valueListStr []string, args []interface{}) {
	fields := []string{
		`"id"`, `"name"`, `"description"`, `"scope"`, `"scope_config"`,
		`"strategy"`, `"strategy_config"`, `"psp_list"`, `"version"`, `"status"`,
		`"rollout_percentage"`, `"rollout_started_at"`, `"rollout_rollback_version"`,
		`"created_by"`, `"approved_by"`,
		`"meta_created_at"`, `"meta_created_by"`, `"meta_updated_at"`, `"meta_updated_by"`,
		`"meta_deleted_at"`, `"meta_deleted_by"`,
	}
	index := 0
	for _, item := range items {
		var values []string
		for range fields {
			index++
			values = append(values, fmt.Sprintf("$%d", index))
		}
		args = append(args,
			item.Id, item.Name, item.Description, item.Scope, item.ScopeConfig,
			item.Strategy, item.StrategyConfig, item.PspList, item.Version, item.Status,
			item.RolloutPercentage, item.RolloutStartedAt, item.RolloutRollbackVersion,
			item.CreatedBy, item.ApprovedBy,
			item.MetaCreatedAt, item.MetaCreatedBy, item.MetaUpdatedAt, item.MetaUpdatedBy,
			item.MetaDeletedAt, item.MetaDeletedBy,
		)
		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
	return
}

// ---------------------------------------------------------------------------
// Implementation
// ---------------------------------------------------------------------------

// CreatePayrouteRoutingRules inserts a new routing rule.
func (r *RepositoryImpl) CreatePayrouteRoutingRules(ctx context.Context, item *model.PayrouteRoutingRules) error {
	fields, values, args := composePayrouteRoutingRulesInsertFieldsAndValues([]*model.PayrouteRoutingRules{item})
	query := fmt.Sprintf(`INSERT INTO "payroute_routing_rules" %s VALUES %s`, fields, strings.Join(values, ","))
	if _, err := r.exec(ctx, query, args); err != nil {
		log.Error().Err(err).Msg("[CreatePayrouteRoutingRules] failed")
		return failure.InternalError(err)
	}
	return nil
}

// ResolvePayrouteRoutingRulesByID retrieves a routing rule by primary key.
func (r *RepositoryImpl) ResolvePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID) (model.PayrouteRoutingRules, error) {
	var item model.PayrouteRoutingRules
	query := fmt.Sprintf(`SELECT %s FROM "payroute_routing_rules" WHERE "id" = $1 AND "meta_deleted_at" IS NULL`, defaultPayrouteRoutingRulesSelectFields())
	if err := r.db.Read.GetContext(ctx, &item, query, id.Id); err != nil {
		if err == sql.ErrNoRows {
			return item, failure.NotFound("payroute_routing_rules")
		}
		log.Error().Err(err).Msg("[ResolvePayrouteRoutingRulesByID] failed")
		return item, failure.InternalError(err)
	}
	return item, nil
}

// ResolvePayrouteRoutingRulesByFilter retrieves routing rules with filtering.
func (r *RepositoryImpl) ResolvePayrouteRoutingRulesByFilter(ctx context.Context, filter model.Filter) (result []model.PayrouteRoutingRulesFilterResult, err error) {
	var args []interface{}
	whereClauses := []string{`"meta_deleted_at" IS NULL`}

	for _, f := range filter.FilterFields {
		switch f.Field {
		case "status":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"status" = $%d`, len(args)))
		case "scope":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"scope" = $%d`, len(args)))
		case "strategy":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"strategy" = $%d`, len(args)))
		case "name":
			args = append(args, "%"+fmt.Sprintf("%v", f.Value)+"%")
			whereClauses = append(whereClauses, fmt.Sprintf(`"name" LIKE $%d`, len(args)))
		}
	}

	whereStr := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM "payroute_routing_rules" WHERE %s`, whereStr)
	var total int
	if err = r.db.Read.GetContext(ctx, &total, countQuery, args...); err != nil {
		log.Error().Err(err).Msg("[ResolvePayrouteRoutingRulesByFilter] count failed")
		return nil, failure.InternalError(err)
	}

	orderBy := `"meta_created_at" DESC`
	if len(filter.Sorts) > 0 {
		s := filter.Sorts[0]
		dir := "ASC"
		if strings.EqualFold(string(s.Order), "desc") {
			dir = "DESC"
		}
		orderBy = fmt.Sprintf(`"%s" %s`, s.Field, dir)
	}

	offset := 0
	if filter.Pagination.Page > 1 {
		offset = (filter.Pagination.Page - 1) * filter.Pagination.PageSize
	}
	limit := filter.Pagination.PageSize
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit, offset)
	query := fmt.Sprintf(`SELECT %s FROM "payroute_routing_rules" WHERE %s ORDER BY %s LIMIT $%d OFFSET $%d`,
		defaultPayrouteRoutingRulesSelectFields(), whereStr, orderBy, len(args)-1, len(args))

	var items []model.PayrouteRoutingRules
	if err = r.db.Read.SelectContext(ctx, &items, query, args...); err != nil {
		log.Error().Err(err).Msg("[ResolvePayrouteRoutingRulesByFilter] query failed")
		return nil, failure.InternalError(err)
	}

	result = make([]model.PayrouteRoutingRulesFilterResult, len(items))
	for i, item := range items {
		result[i] = model.PayrouteRoutingRulesFilterResult{
			PayrouteRoutingRules: item,
			FilterCount:          total,
		}
	}
	return result, nil
}

// ResolveActiveRuleByScope finds the active rule matching a payment scope.
func (r *RepositoryImpl) ResolveActiveRuleByScope(ctx context.Context, methodCode, channelCode, currency string) (*model.PayrouteRoutingRules, error) {
	query := fmt.Sprintf(`SELECT %s FROM "payroute_routing_rules"
		WHERE "status" IN ('active','rolling_out') AND "meta_deleted_at" IS NULL
		AND ("scope_config"->>'method_code' = $1 OR "scope" = 'global')
		ORDER BY 
			CASE WHEN "scope" = 'method_channel' AND "scope_config"->>'channel_code' = $2 THEN 1
			     WHEN "scope" = 'method' THEN 2
			     WHEN "scope" = 'method_amount_band' THEN 3
			     WHEN "scope" = 'global' THEN 4
			     ELSE 5 END,
			"version" DESC
		LIMIT 1`, defaultPayrouteRoutingRulesSelectFields())

	var item model.PayrouteRoutingRules
	if err := r.db.Read.GetContext(ctx, &item, query, methodCode, channelCode); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Err(err).Msg("[ResolveActiveRuleByScope] failed")
		return nil, failure.InternalError(err)
	}
	return &item, nil
}

// ResolveConflictingRulesByScope checks for scope conflicts (ERR-ROUTE-04).
func (r *RepositoryImpl) ResolveConflictingRulesByScope(ctx context.Context, scope string, scopeConfig map[string]interface{}, excludeID interface{}) (int, error) {
	var args []interface{}
	whereClauses := []string{`"status" IN ('active','rolling_out')`, `"meta_deleted_at" IS NULL`}

	if excludeID != nil {
		args = append(args, excludeID)
		whereClauses = append(whereClauses, fmt.Sprintf(`"id" != $%d`, len(args)))
	}

	args = append(args, scope)
	whereClauses = append(whereClauses, fmt.Sprintf(`"scope" = $%d`, len(args)))

	if methodCode, ok := scopeConfig["method_code"].(string); ok {
		args = append(args, methodCode)
		whereClauses = append(whereClauses, fmt.Sprintf(`("scope_config"->>'method_code' = $%d)`, len(args)))
	}

	query := fmt.Sprintf(`SELECT COUNT(*) FROM "payroute_routing_rules" WHERE %s`, strings.Join(whereClauses, " AND "))
	var count int
	if err := r.db.Read.GetContext(ctx, &count, query, args...); err != nil {
		return 0, failure.InternalError(err)
	}
	return count, nil
}
func (r *RepositoryImpl) UpdatePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID, _ *model.PayrouteRoutingRules, fields ...PayrouteRoutingRulesUpdateField) error {
	if len(fields) == 0 {
		return nil
	}
	var setClauses []string
	var args []interface{}
	for i, f := range fields {
		args = append(args, f.value)
		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, f.field, i+1))
	}
	args = append(args, id.Id)
	query := fmt.Sprintf(`UPDATE "payroute_routing_rules" SET %s WHERE "id" = $%d AND "meta_deleted_at" IS NULL`,
		strings.Join(setClauses, ", "), len(args))
	result, err := r.exec(ctx, query, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayrouteRoutingRulesByID] failed")
		return failure.InternalError(err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return failure.WithCode(shared.ErrRouteOptimisticLockConflict, "routing rule was modified by another user")
	}
	return nil
}

// DeletePayrouteRoutingRulesByID soft-deletes a routing rule.
func (r *RepositoryImpl) DeletePayrouteRoutingRulesByID(ctx context.Context, id model.PayrouteRoutingRulesPrimaryID) error {
	query := `UPDATE "payroute_routing_rules" SET "meta_deleted_at" = now() WHERE "id" = $1 AND "meta_deleted_at" IS NULL`
	if _, err := r.exec(ctx, query, []interface{}{id.Id}); err != nil {
		log.Error().Err(err).Msg("[DeletePayrouteRoutingRulesByID] failed")
		return failure.InternalError(err)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Rule Versions
// ---------------------------------------------------------------------------

type PayrouteRoutingRuleVersionsRepository interface {
	CreatePayrouteRoutingRuleVersions(ctx context.Context, item *model.PayrouteRoutingRuleVersions) error
	ResolvePayrouteRoutingRuleVersionsByRuleID(ctx context.Context, ruleID interface{}) ([]model.PayrouteRoutingRuleVersions, error)
	ResolvePayrouteRoutingRuleVersionsByVersion(ctx context.Context, ruleID interface{}, version int) (*model.PayrouteRoutingRuleVersions, error)
}

func (r *RepositoryImpl) CreatePayrouteRoutingRuleVersions(ctx context.Context, item *model.PayrouteRoutingRuleVersions) error {
	query := `INSERT INTO "payroute_routing_rule_versions" 
		("id","rule_id","version","snapshot","changed_by","change_reason","meta_created_at")
		VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.exec(ctx, query, []interface{}{
		item.Id, item.RuleId, item.Version, item.Snapshot, item.ChangedBy, item.ChangeReason, item.MetaCreatedAt,
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayrouteRoutingRuleVersions] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolvePayrouteRoutingRuleVersionsByRuleID(ctx context.Context, ruleID interface{}) ([]model.PayrouteRoutingRuleVersions, error) {
	var items []model.PayrouteRoutingRuleVersions
	query := `SELECT "id","rule_id","version","snapshot","changed_by","change_reason","meta_created_at"
		FROM "payroute_routing_rule_versions" WHERE "rule_id" = $1 ORDER BY "version" DESC`
	if err := r.db.Read.SelectContext(ctx, &items, query, ruleID); err != nil {
		log.Error().Err(err).Msg("[ResolvePayrouteRoutingRuleVersionsByRuleID] failed")
		return nil, failure.InternalError(err)
	}
	return items, nil
}

func (r *RepositoryImpl) ResolvePayrouteRoutingRuleVersionsByVersion(ctx context.Context, ruleID interface{}, version int) (*model.PayrouteRoutingRuleVersions, error) {
	var item model.PayrouteRoutingRuleVersions
	query := `SELECT "id","rule_id","version","snapshot","changed_by","change_reason","meta_created_at"
		FROM "payroute_routing_rule_versions" WHERE "rule_id" = $1 AND "version" = $2`
	if err := r.db.Read.GetContext(ctx, &item, query, ruleID, version); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Err(err).Msg("[ResolvePayrouteRoutingRuleVersionsByVersion] failed")
		return nil, failure.InternalError(err)
	}
	return &item, nil
}

// ---------------------------------------------------------------------------
// MDR Rates
// ---------------------------------------------------------------------------

type PayrouteMdrRatesRepository interface {
	CreatePayrouteMdrRates(ctx context.Context, item *model.PayrouteMdrRates) error
	ResolvePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID) (model.PayrouteMdrRates, error)
	ResolveActiveMDRRate(ctx context.Context, providerAccountID, paymentMethodID interface{}, amount string) (*model.PayrouteMdrRates, error)
	ResolvePayrouteMdrRatesByFilter(ctx context.Context, filter model.Filter) ([]model.PayrouteMdrRatesFilterResult, error)
	UpdatePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID, item *model.PayrouteMdrRates, fields ...PayrouteMdrRatesUpdateField) error
	DeletePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID) error
}

type PayrouteMdrRatesField string
type PayrouteMdrRatesUpdateField struct {
	field PayrouteMdrRatesField
	value interface{}
}

const (
	PayrouteMdrRatesFieldId              PayrouteMdrRatesField = "id"
	PayrouteMdrRatesFieldProviderAccount PayrouteMdrRatesField = "provider_account_id"
	PayrouteMdrRatesFieldPaymentMethod   PayrouteMdrRatesField = "payment_method_id"
	PayrouteMdrRatesFieldPaymentChannel  PayrouteMdrRatesField = "payment_channel_id"
	PayrouteMdrRatesFieldPercentage      PayrouteMdrRatesField = "percentage"
	PayrouteMdrRatesFieldFixedFee        PayrouteMdrRatesField = "fixed_fee"
	PayrouteMdrRatesFieldMinAmount       PayrouteMdrRatesField = "min_amount"
	PayrouteMdrRatesFieldMaxAmount       PayrouteMdrRatesField = "max_amount"
	PayrouteMdrRatesFieldCurrency        PayrouteMdrRatesField = "currency"
	PayrouteMdrRatesFieldEffectiveFrom   PayrouteMdrRatesField = "effective_from"
	PayrouteMdrRatesFieldEffectiveTo     PayrouteMdrRatesField = "effective_to"
	PayrouteMdrRatesFieldMetaCreatedAt   PayrouteMdrRatesField = "meta_created_at"
	PayrouteMdrRatesFieldMetaUpdatedAt   PayrouteMdrRatesField = "meta_updated_at"
	PayrouteMdrRatesFieldMetaDeletedAt   PayrouteMdrRatesField = "meta_deleted_at"
)

func NewPayrouteMdrRatesUpdateField(field PayrouteMdrRatesField, val interface{}) PayrouteMdrRatesUpdateField {
	return PayrouteMdrRatesUpdateField{field: field, value: val}
}

func defaultPayrouteMdrRatesSelectFields() string {
	return `"id","provider_account_id","payment_method_id","payment_channel_id","percentage","fixed_fee",
		"min_amount","max_amount","currency","effective_from","effective_to",
		"meta_created_at","meta_updated_at","meta_deleted_at"`
}

func (r *RepositoryImpl) CreatePayrouteMdrRates(ctx context.Context, item *model.PayrouteMdrRates) error {
	query := fmt.Sprintf(`INSERT INTO "payroute_mdr_rates" 
		("id","provider_account_id","payment_method_id","payment_channel_id","percentage","fixed_fee",
		 "min_amount","max_amount","currency","effective_from","effective_to",
		 "meta_created_at","meta_updated_at")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`)
	_, err := r.exec(ctx, query, []interface{}{
		item.Id, item.ProviderAccountId, item.PaymentMethodId, item.PaymentChannelId,
		item.Percentage, item.FixedFee, item.MinAmount, item.MaxAmount,
		item.Currency, item.EffectiveFrom, item.EffectiveTo,
		item.MetaCreatedAt, item.MetaUpdatedAt,
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayrouteMdrRates] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolvePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID) (model.PayrouteMdrRates, error) {
	var item model.PayrouteMdrRates
	query := fmt.Sprintf(`SELECT %s FROM "payroute_mdr_rates" WHERE "id" = $1 AND "meta_deleted_at" IS NULL`, defaultPayrouteMdrRatesSelectFields())
	if err := r.db.Read.GetContext(ctx, &item, query, id.Id); err != nil {
		if err == sql.ErrNoRows {
			return item, failure.NotFound("payroute_mdr_rates")
		}
		return item, failure.InternalError(err)
	}
	return item, nil
}

func (r *RepositoryImpl) ResolveActiveMDRRate(ctx context.Context, providerAccountID, paymentMethodID interface{}, _ string) (*model.PayrouteMdrRates, error) {
	query := fmt.Sprintf(`SELECT %s FROM "payroute_mdr_rates"
		WHERE "provider_account_id" = $1 AND "payment_method_id" = $2
		AND "meta_deleted_at" IS NULL
		AND ("effective_to" IS NULL OR "effective_to" > now())
		AND "effective_from" <= now()
		ORDER BY "effective_from" DESC LIMIT 1`, defaultPayrouteMdrRatesSelectFields())
	var item model.PayrouteMdrRates
	if err := r.db.Read.GetContext(ctx, &item, query, providerAccountID, paymentMethodID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, failure.InternalError(err)
	}
	return &item, nil
}

func (r *RepositoryImpl) ResolvePayrouteMdrRatesByFilter(ctx context.Context, filter model.Filter) (result []model.PayrouteMdrRatesFilterResult, err error) {
	var args []interface{}
	whereClauses := []string{`"meta_deleted_at" IS NULL`}

	for _, f := range filter.FilterFields {
		switch f.Field {
		case "provider_account_id":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"provider_account_id" = $%d`, len(args)))
		case "payment_method_id":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"payment_method_id" = $%d`, len(args)))
		}
	}

	whereStr := strings.Join(whereClauses, " AND ")
	offset := 0
	if filter.Pagination.Page > 1 {
		offset = (filter.Pagination.Page - 1) * filter.Pagination.PageSize
	}
	limit := filter.Pagination.PageSize
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit, offset)
	query := fmt.Sprintf(`SELECT %s FROM "payroute_mdr_rates" WHERE %s ORDER BY "percentage" ASC LIMIT $%d OFFSET $%d`,
		defaultPayrouteMdrRatesSelectFields(), whereStr, len(args)-1, len(args))

	var items []model.PayrouteMdrRates
	if err = r.db.Read.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, failure.InternalError(err)
	}

	result = make([]model.PayrouteMdrRatesFilterResult, len(items))
	for i, item := range items {
		result[i] = model.PayrouteMdrRatesFilterResult{PayrouteMdrRates: item}
	}
	return result, nil
}

func (r *RepositoryImpl) UpdatePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID, _ *model.PayrouteMdrRates, fields ...PayrouteMdrRatesUpdateField) error {
	if len(fields) == 0 {
		return nil
	}
	var setClauses []string
	var args []interface{}
	for i, f := range fields {
		args = append(args, f.value)
		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, f.field, i+1))
	}
	args = append(args, id.Id)
	query := fmt.Sprintf(`UPDATE "payroute_mdr_rates" SET %s WHERE "id" = $%d AND "meta_deleted_at" IS NULL`,
		strings.Join(setClauses, ", "), len(args))
	if _, err := r.exec(ctx, query, args); err != nil {
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) DeletePayrouteMdrRatesByID(ctx context.Context, id model.PayrouteMdrRatesPrimaryID) error {
	query := `UPDATE "payroute_mdr_rates" SET "meta_deleted_at" = now() WHERE "id" = $1 AND "meta_deleted_at" IS NULL`
	if _, err := r.exec(ctx, query, []interface{}{id.Id}); err != nil {
		return failure.InternalError(err)
	}
	return nil
}

// ---------------------------------------------------------------------------
// PSP Health
// ---------------------------------------------------------------------------

type PayroutePspHealthRepository interface {
	UpsertPayroutePspHealth(ctx context.Context, item *model.PayroutePspHealth) error
	ResolvePayroutePspHealth(ctx context.Context, providerAccountID interface{}, methodCode string) ([]model.PayroutePspHealth, error)
	ResolveLatestPSPHealth(ctx context.Context, providerAccountID interface{}, methodCode, channelCode string) (*model.PayroutePspHealth, error)
	ComputeAndStoreHealthAggregates(ctx context.Context, since interface{}) error
}

func (r *RepositoryImpl) UpsertPayroutePspHealth(ctx context.Context, item *model.PayroutePspHealth) error {
	query := `INSERT INTO "payroute_psp_health"
		("id","provider_account_id","method_code","channel_code","currency",
		 "window_start","window_end","total_attempts","successful","failed_system","failed_decline",
		 "success_rate","avg_latency_ms","p95_latency_ms","sample_size","meta_created_at")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
		ON CONFLICT ("provider_account_id","method_code","window_start") DO UPDATE SET
			"window_end" = EXCLUDED."window_end",
			"total_attempts" = EXCLUDED."total_attempts",
			"successful" = EXCLUDED."successful",
			"failed_system" = EXCLUDED."failed_system",
			"failed_decline" = EXCLUDED."failed_decline",
			"success_rate" = EXCLUDED."success_rate",
			"avg_latency_ms" = EXCLUDED."avg_latency_ms",
			"p95_latency_ms" = EXCLUDED."p95_latency_ms",
			"sample_size" = EXCLUDED."sample_size"`
	_, err := r.exec(ctx, query, []interface{}{
		item.Id, item.ProviderAccountId, item.MethodCode, item.ChannelCode, item.Currency,
		item.WindowStart, item.WindowEnd, item.TotalAttempts, item.Successful,
		item.FailedSystem, item.FailedDecline, item.SuccessRate,
		item.AvgLatencyMs, item.P95LatencyMs, item.SampleSize, item.MetaCreatedAt,
	})
	if err != nil {
		log.Error().Err(err).Msg("[UpsertPayroutePspHealth] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolvePayroutePspHealth(ctx context.Context, providerAccountID interface{}, methodCode string) ([]model.PayroutePspHealth, error) {
	var items []model.PayroutePspHealth
	query := fmt.Sprintf(`SELECT %s FROM "payroute_psp_health"
		WHERE "provider_account_id" = $1 AND "method_code" = $2
		ORDER BY "window_start" DESC LIMIT 24`, defaultPayroutePspHealthSelectFields())
	if err := r.db.Read.SelectContext(ctx, &items, query, providerAccountID, methodCode); err != nil {
		return nil, failure.InternalError(err)
	}
	return items, nil
}

func (r *RepositoryImpl) ResolveLatestPSPHealth(ctx context.Context, providerAccountID interface{}, methodCode, channelCode string) (*model.PayroutePspHealth, error) {
	var item model.PayroutePspHealth
	query := fmt.Sprintf(`SELECT %s FROM "payroute_psp_health"
		WHERE "provider_account_id" = $1 AND "method_code" = $2 AND ("channel_code" = $3 OR "channel_code" IS NULL)
		ORDER BY "window_start" DESC LIMIT 1`, defaultPayroutePspHealthSelectFields())
	if err := r.db.Read.GetContext(ctx, &item, query, providerAccountID, methodCode, channelCode); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, failure.InternalError(err)
	}
	return &item, nil
}

func defaultPayroutePspHealthSelectFields() string {
	return `"id","provider_account_id","method_code","channel_code","currency",
		"window_start","window_end","total_attempts","successful","failed_system","failed_decline",
		"success_rate","avg_latency_ms","p95_latency_ms","sample_size","meta_created_at"`
}

func (r *RepositoryImpl) ComputeAndStoreHealthAggregates(ctx context.Context, since interface{}) error {
	// This is a background task - compute health from payment_attempts
	query := `
		INSERT INTO "payroute_psp_health" 
			("id","provider_account_id","method_code","channel_code","currency",
			 "window_start","window_end","total_attempts","successful","failed_system","failed_decline",
			 "success_rate","avg_latency_ms","sample_size","meta_created_at")
		SELECT 
			gen_random_uuid(),
			pa."provider_account_id",
			pa."method_code",
			pa."channel_code",
			pa."currency",
			date_trunc('hour', $1::timestamptz) as window_start,
			date_trunc('hour', $1::timestamptz) + interval '1 hour' as window_end,
			COUNT(*) as total_attempts,
			COUNT(*) FILTER (WHERE pa."status" IN ('paid','captured')) as successful,
			COUNT(*) FILTER (WHERE pa."failure_code" IS NOT NULL AND pa."failure_code" != 'hard_decline') as failed_system,
			COUNT(*) FILTER (WHERE pa."failure_code" = 'hard_decline') as failed_decline,
			ROUND(COUNT(*) FILTER (WHERE pa."status" IN ('paid','captured'))::decimal / NULLIF(COUNT(*),0), 4) as success_rate,
			AVG(EXTRACT(EPOCH FROM (pa."meta_updated_at" - pa."meta_created_at")) * 1000)::int as avg_latency_ms,
			COUNT(*) as sample_size,
			now()
		FROM "payment_attempts" pa
		WHERE pa."meta_created_at" >= $1
		GROUP BY pa."provider_account_id", pa."method_code", pa."channel_code", pa."currency"
		HAVING COUNT(*) >= 1
		ON CONFLICT ("provider_account_id","method_code","window_start") DO UPDATE SET
			"window_end" = EXCLUDED."window_end",
			"total_attempts" = EXCLUDED."total_attempts",
			"successful" = EXCLUDED."successful",
			"failed_system" = EXCLUDED."failed_system",
			"failed_decline" = EXCLUDED."failed_decline",
			"success_rate" = EXCLUDED."success_rate",
			"avg_latency_ms" = EXCLUDED."avg_latency_ms",
			"sample_size" = EXCLUDED."sample_size"`
	if _, err := r.exec(ctx, query, []interface{}{since}); err != nil {
		log.Error().Err(err).Msg("[ComputeAndStoreHealthAggregates] failed")
		return failure.InternalError(err)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Audit Log
// ---------------------------------------------------------------------------

type PayrouteAuditLogRepository interface {
	CreatePayrouteAuditLog(ctx context.Context, item *model.PayrouteAuditLog) error
	ResolvePayrouteAuditLogByFilter(ctx context.Context, filter model.Filter) ([]model.PayrouteAuditLogFilterResult, error)
}

func (r *RepositoryImpl) CreatePayrouteAuditLog(ctx context.Context, item *model.PayrouteAuditLog) error {
	query := `INSERT INTO "payroute_audit_log"
		("id","actor_user_id","actor_email","actor_role","action","target_entity","target_id",
		 "before_value","after_value","reason","ip_address","user_agent","meta_created_at")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	_, err := r.exec(ctx, query, []interface{}{
		item.Id, item.ActorUserId, item.ActorEmail, item.ActorRole,
		item.Action, item.TargetEntity, item.TargetId,
		item.BeforeValue, item.AfterValue, item.Reason,
		item.IpAddress, item.UserAgent, item.MetaCreatedAt,
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayrouteAuditLog] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolvePayrouteAuditLogByFilter(ctx context.Context, filter model.Filter) (result []model.PayrouteAuditLogFilterResult, err error) {
	var args []interface{}
	whereClauses := []string{}

	for _, f := range filter.FilterFields {
		switch f.Field {
		case "target_entity":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"target_entity" = $%d`, len(args)))
		case "actor_user_id":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"actor_user_id" = $%d`, len(args)))
		case "action":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"action" = $%d`, len(args)))
		}
	}

	whereStr := ""
	if len(whereClauses) > 0 {
		whereStr = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	offset := 0
	if filter.Pagination.Page > 1 {
		offset = (filter.Pagination.Page - 1) * filter.Pagination.PageSize
	}
	limit := filter.Pagination.PageSize
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit, offset)
	query := fmt.Sprintf(`SELECT "id","actor_user_id","actor_email","actor_role","action","target_entity","target_id",
		"before_value","after_value","reason","ip_address","user_agent","meta_created_at"
		FROM "payroute_audit_log" %s ORDER BY "meta_created_at" DESC LIMIT $%d OFFSET $%d`,
		whereStr, len(args)-1, len(args))

	var items []model.PayrouteAuditLog
	if err = r.db.Read.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, failure.InternalError(err)
	}

	result = make([]model.PayrouteAuditLogFilterResult, len(items))
	for i, item := range items {
		result[i] = model.PayrouteAuditLogFilterResult{PayrouteAuditLog: item}
	}
	return result, nil
}

// ---------------------------------------------------------------------------
// Config
// ---------------------------------------------------------------------------

type PayrouteConfigRepository interface {
	GetPayrouteConfig(ctx context.Context, key string) (*model.PayrouteConfig, error)
	UpsertPayrouteConfig(ctx context.Context, item *model.PayrouteConfig) error
	ResolvePayrouteConfigByFilter(ctx context.Context, filter model.Filter) ([]model.PayrouteConfigFilterResult, error)
}

func (r *RepositoryImpl) GetPayrouteConfig(ctx context.Context, key string) (*model.PayrouteConfig, error) {
	var item model.PayrouteConfig
	query := `SELECT "key","value","description","updated_by","meta_updated_at"
		FROM "payroute_config" WHERE "key" = $1`
	if err := r.db.Read.GetContext(ctx, &item, query, key); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, failure.InternalError(err)
	}
	return &item, nil
}

func (r *RepositoryImpl) UpsertPayrouteConfig(ctx context.Context, item *model.PayrouteConfig) error {
	query := `INSERT INTO "payroute_config" ("key","value","description","updated_by","meta_updated_at")
		VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT ("key") DO UPDATE SET
			"value" = EXCLUDED."value",
			"description" = EXCLUDED."description",
			"updated_by" = EXCLUDED."updated_by",
			"meta_updated_at" = EXCLUDED."meta_updated_at"`
	_, err := r.exec(ctx, query, []interface{}{
		item.Key, item.Value, item.Description, item.UpdatedBy, item.MetaUpdatedAt,
	})
	if err != nil {
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolvePayrouteConfigByFilter(ctx context.Context, filter model.Filter) ([]model.PayrouteConfigFilterResult, error) {
	query := `SELECT "key","value","description","updated_by","meta_updated_at" FROM "payroute_config" ORDER BY "key"`
	var items []model.PayrouteConfig
	if err := r.db.Read.SelectContext(ctx, &items, query); err != nil {
		return nil, failure.InternalError(err)
	}
	result := make([]model.PayrouteConfigFilterResult, len(items))
	for i, item := range items {
		result[i] = model.PayrouteConfigFilterResult{PayrouteConfig: item}
	}
	return result, nil
}

// ---------------------------------------------------------------------------
// Compile-time interface check
// ---------------------------------------------------------------------------

var _ PayrouteRoutingRulesRepository = (*RepositoryImpl)(nil)
var _ PayrouteRoutingRuleVersionsRepository = (*RepositoryImpl)(nil)
var _ PayrouteMdrRatesRepository = (*RepositoryImpl)(nil)
var _ PayroutePspHealthRepository = (*RepositoryImpl)(nil)
var _ PayrouteAuditLogRepository = (*RepositoryImpl)(nil)
var _ PayrouteConfigRepository = (*RepositoryImpl)(nil)

// Helper: jsonBytes converts a value to json.RawMessage for JSONB columns
func jsonBytes(v interface{}) json.RawMessage {
	if v == nil {
		return json.RawMessage(`{}`)
	}
	b, err := json.Marshal(v)
	if err != nil {
		return json.RawMessage(`{}`)
	}
	return b
}

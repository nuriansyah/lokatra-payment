package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/configs"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// PayRouteAdminHandler — Dashboard API for routing rules, health, kill-switch
// ---------------------------------------------------------------------------

type PayRouteAdminHandler struct {
	config     *configs.Config
	repo       repository.Repository
	routingEng *service.RoutingEngine
	analytics  service.AnalyticsPublisher
}

func ProvidePayRouteAdminHandler(
	config *configs.Config,
	repo repository.Repository,
	routingEng *service.RoutingEngine,
	analytics service.AnalyticsPublisher,
) *PayRouteAdminHandler {
	if analytics == nil {
		analytics = service.NewLogAnalyticsPublisher()
	}
	return &PayRouteAdminHandler{
		config:     config,
		repo:       repo,
		routingEng: routingEng,
		analytics:  analytics,
	}
}

func (h *PayRouteAdminHandler) Router(r chi.Router) {
	r.Route("/payroute", func(r chi.Router) {
		r.Route("/routing-rules", func(r chi.Router) {
			r.Get("/", h.ListRoutingRules)
			r.Post("/", h.CreateRoutingRule)
			r.Get("/{ruleID}", h.GetRoutingRule)
			r.Put("/{ruleID}", h.UpdateRoutingRule)
			r.Post("/{ruleID}/activate", h.ActivateRule)
			r.Post("/{ruleID}/disable", h.DisableRule)
			r.Post("/{ruleID}/archive", h.ArchiveRule)
			r.Get("/{ruleID}/versions", h.ListRuleVersions)
			r.Post("/{ruleID}/rollback/{version}", h.RollbackRule)
		})

		r.Route("/psp-health", func(r chi.Router) {
			r.Get("/", h.ListPSPHealth)
			r.Get("/{accountID}", h.GetPSPHealth)
			r.Post("/{accountID}/kill-switch", h.KillSwitchPSP)
			r.Post("/{accountID}/enable", h.EnablePSP)
		})

		r.Route("/mdr-rates", func(r chi.Router) {
			r.Get("/", h.ListMDRRates)
			r.Post("/", h.UpsertMDRRate)
		})

		r.Get("/audit-log", h.ListAuditLog)
		r.Get("/config", h.ListConfig)
		r.Put("/config/{key}", h.UpdateConfig)
	})
}

// ---------------------------------------------------------------------------
// Routing Rules CRUD
// ---------------------------------------------------------------------------

type createRoutingRuleRequest struct {
	Name           string          `json:"name"`
	Description    string          `json:"description,omitempty"`
	Scope          string          `json:"scope"`
	ScopeConfig    json.RawMessage `json:"scope_config"`
	Strategy       string          `json:"strategy"`
	StrategyConfig json.RawMessage `json:"strategy_config"`
	PspList        json.RawMessage `json:"psp_list"`
}

func (h *PayRouteAdminHandler) ListRoutingRules(w http.ResponseWriter, r *http.Request) {
	filter := model.Filter{
		Pagination: model.Pagination{Page: 1, PageSize: 50},
	}
	if status := r.URL.Query().Get("status"); status != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field: "status", Operator: model.OperatorEqual, Value: status,
		})
	}
	if scope := r.URL.Query().Get("scope"); scope != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field: "scope", Operator: model.OperatorEqual, Value: scope,
		})
	}

	result, err := h.repo.ResolvePayrouteRoutingRulesByFilter(r.Context(), filter)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"data":  result,
		"total": len(result),
	})
}

func (h *PayRouteAdminHandler) CreateRoutingRule(w http.ResponseWriter, r *http.Request) {
	var req createRoutingRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		writeErrorResponse(w, failure.BadRequestFromString("name is required"))
		return
	}

	actorID := extractActorID(r)
	now := time.Now().UTC()

	item := model.PayrouteRoutingRules{
		Id:                uuid.Must(uuid.NewV7()),
		Name:              req.Name,
		Description:       null.StringFrom(req.Description),
		Scope:             model.PayrouteRuleScope(req.Scope),
		ScopeConfig:       req.ScopeConfig,
		Strategy:          model.PayrouteRoutingStrategy(req.Strategy),
		StrategyConfig:    req.StrategyConfig,
		PspList:           req.PspList,
		Version:           1,
		Status:            model.PayrouteStatusDraft,
		RolloutPercentage: 100,
		CreatedBy:         actorID,
		MetaSignature: shared.MetaSignature{
			MetaCreatedAt: now,
			MetaCreatedBy: actorID,
			MetaUpdatedAt: null.TimeFrom(now),
			MetaUpdatedBy: &actorID,
		},
	}

	if err := h.repo.CreatePayrouteRoutingRules(r.Context(), &item); err != nil {
		writeErrorResponse(w, err)
		return
	}

	h.createVersionSnapshot(r.Context(), item, actorID, "initial creation")
	h.auditLog(r.Context(), actorID, "routing_rule.create", "payroute_routing_rules", item.Id, nil, item)
	service.PublishRoutingRuleChanged(r.Context(), h.analytics, item.Id.String(), actorID.String(), nil, item)

	writeJSONResponse(w, http.StatusCreated, item)
}

func (h *PayRouteAdminHandler) GetRoutingRule(w http.ResponseWriter, r *http.Request) {
	ruleID, err := uuid.FromString(chi.URLParam(r, "ruleID"))
	if err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	item, err := h.repo.ResolvePayrouteRoutingRulesByID(r.Context(), model.PayrouteRoutingRulesPrimaryID{Id: ruleID})
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, item)
}

func (h *PayRouteAdminHandler) UpdateRoutingRule(w http.ResponseWriter, r *http.Request) {
	ruleID, err := uuid.FromString(chi.URLParam(r, "ruleID"))
	if err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	var req createRoutingRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	actorID := extractActorID(r)
	existing, err := h.repo.ResolvePayrouteRoutingRulesByID(r.Context(), model.PayrouteRoutingRulesPrimaryID{Id: ruleID})
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	beforeSnapshot, _ := json.Marshal(existing)

	fields := []repository.PayrouteRoutingRulesUpdateField{
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldName, req.Name),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldDescription, req.Description),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldScope, req.Scope),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldScopeConfig, req.ScopeConfig),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldStrategy, req.Strategy),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldStrategyConfig, req.StrategyConfig),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldPspList, req.PspList),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldVersion, existing.Version+1),
	}
	if err := h.repo.UpdatePayrouteRoutingRulesByID(r.Context(), model.PayrouteRoutingRulesPrimaryID{Id: ruleID}, nil, fields...); err != nil {
		writeErrorResponse(w, err)
		return
	}

	updated := existing
	updated.Name = req.Name
	updated.Scope = model.PayrouteRuleScope(req.Scope)
	updated.ScopeConfig = req.ScopeConfig
	updated.Strategy = model.PayrouteRoutingStrategy(req.Strategy)
	updated.StrategyConfig = req.StrategyConfig
	updated.PspList = req.PspList
	updated.Version = existing.Version + 1
	h.createVersionSnapshot(r.Context(), updated, actorID, "rule updated")

	h.routingEng.RuleCache().InvalidateAll()

	afterSnapshot, _ := json.Marshal(updated)
	h.auditLog(r.Context(), actorID, "routing_rule.update", "payroute_routing_rules", ruleID, beforeSnapshot, afterSnapshot)
	service.PublishRoutingRuleChanged(r.Context(), h.analytics, ruleID.String(), actorID.String(), beforeSnapshot, afterSnapshot)

	writeJSONResponse(w, http.StatusOK, updated)
}

func (h *PayRouteAdminHandler) ActivateRule(w http.ResponseWriter, r *http.Request) {
	h.setRuleStatus(w, r, model.PayrouteStatusActive, "routing_rule.activate")
}

func (h *PayRouteAdminHandler) DisableRule(w http.ResponseWriter, r *http.Request) {
	h.setRuleStatus(w, r, model.PayrouteStatusDisabled, "routing_rule.disable")
}

func (h *PayRouteAdminHandler) ArchiveRule(w http.ResponseWriter, r *http.Request) {
	h.setRuleStatus(w, r, model.PayrouteStatusArchived, "routing_rule.archive")
}

func (h *PayRouteAdminHandler) setRuleStatus(w http.ResponseWriter, r *http.Request, status model.PayrouteRuleStatus, action string) {
	ruleID, err := uuid.FromString(chi.URLParam(r, "ruleID"))
	if err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	actorID := extractActorID(r)
	fields := []repository.PayrouteRoutingRulesUpdateField{
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldStatus, status),
	}
	if err := h.repo.UpdatePayrouteRoutingRulesByID(r.Context(), model.PayrouteRoutingRulesPrimaryID{Id: ruleID}, nil, fields...); err != nil {
		writeErrorResponse(w, err)
		return
	}

	h.routingEng.RuleCache().InvalidateAll()
	h.auditLog(r.Context(), actorID, action, "payroute_routing_rules", ruleID, nil, map[string]string{"status": string(status)})
	writeJSONResponse(w, http.StatusOK, map[string]string{"status": string(status)})
}

func (h *PayRouteAdminHandler) ListRuleVersions(w http.ResponseWriter, r *http.Request) {
	ruleID, err := uuid.FromString(chi.URLParam(r, "ruleID"))
	if err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	versions, err := h.repo.ResolvePayrouteRoutingRuleVersionsByRuleID(r.Context(), ruleID)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, versions)
}

func (h *PayRouteAdminHandler) RollbackRule(w http.ResponseWriter, r *http.Request) {
	ruleID, err := uuid.FromString(chi.URLParam(r, "ruleID"))
	if err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}
	version := 0
	if v := chi.URLParam(r, "version"); v != "" {
		if vi, parseErr := strconv.Atoi(v); parseErr == nil {
			version = vi
		}
	}
	if version <= 0 {
		writeErrorResponse(w, failure.BadRequestFromString("version must be positive"))
		return
	}

	actorID := extractActorID(r)
	snapshot, err := h.repo.ResolvePayrouteRoutingRuleVersionsByVersion(r.Context(), ruleID, version)
	if err != nil || snapshot == nil {
		writeErrorResponse(w, failure.NotFound("version not found"))
		return
	}

	var snapData model.PayrouteRoutingRules
	if err := json.Unmarshal(snapshot.Snapshot, &snapData); err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	fields := []repository.PayrouteRoutingRulesUpdateField{
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldName, snapData.Name),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldScope, snapData.Scope),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldScopeConfig, snapData.ScopeConfig),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldStrategy, snapData.Strategy),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldStrategyConfig, snapData.StrategyConfig),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldPspList, snapData.PspList),
		repository.NewPayrouteRoutingRulesUpdateField(repository.PayrouteRoutingRulesFieldVersion, snapData.Version+1),
	}
	if err := h.repo.UpdatePayrouteRoutingRulesByID(r.Context(), model.PayrouteRoutingRulesPrimaryID{Id: ruleID}, nil, fields...); err != nil {
		writeErrorResponse(w, err)
		return
	}

	h.routingEng.RuleCache().InvalidateAll()
	h.auditLog(r.Context(), actorID, "routing_rule.rollback", "payroute_routing_rules", ruleID, nil, map[string]interface{}{
		"rollback_to_version": version,
	})
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"status": "rolled_back", "version": version})
}

// ---------------------------------------------------------------------------
// PSP Health & Kill-Switch
// ---------------------------------------------------------------------------

func (h *PayRouteAdminHandler) ListPSPHealth(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"data": []interface{}{}, "total": 0})
}

func (h *PayRouteAdminHandler) GetPSPHealth(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"data": []interface{}{}})
}

func (h *PayRouteAdminHandler) KillSwitchPSP(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")
	reason := r.FormValue("reason")
	actorID := extractActorID(r)

	if strings.TrimSpace(accountID) == "" {
		writeErrorResponse(w, failure.BadRequestFromString("accountID is required"))
		return
	}
	if strings.TrimSpace(reason) == "" {
		writeErrorResponse(w, failure.BadRequestFromString("reason is required for kill-switch"))
		return
	}

	h.routingEng.KillSwitch().Disable(accountID)

	service.PublishKillSwitchActivated(r.Context(), h.analytics, accountID, extractActorID(r).String(), reason)

	targetUUID, _ := uuid.FromString(accountID)
	h.auditLog(r.Context(), actorID, "psp.kill_switch", "provider_accounts", targetUUID, nil, map[string]string{
		"action": "disabled",
		"reason": reason,
	})
	writeJSONResponse(w, http.StatusOK, map[string]string{
		"status":     "disabled",
		"account_id": accountID,
	})
}

func (h *PayRouteAdminHandler) EnablePSP(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")
	actorID := extractActorID(r)

	h.routingEng.KillSwitch().Enable(accountID)
	targetUUID, _ := uuid.FromString(accountID)
	h.auditLog(r.Context(), actorID, "psp.enable", "provider_accounts", targetUUID, nil, map[string]string{"action": "enabled"})
	writeJSONResponse(w, http.StatusOK, map[string]string{
		"status":     "enabled",
		"account_id": accountID,
	})
}

// ---------------------------------------------------------------------------
// MDR Rates
// ---------------------------------------------------------------------------

type upsertMDRRateRequest struct {
	ProviderAccountID string          `json:"provider_account_id"`
	PaymentMethodID   string          `json:"payment_method_id"`
	PaymentChannelID  string          `json:"payment_channel_id,omitempty"`
	Percentage        decimal.Decimal `json:"percentage"`
	FixedFee          decimal.Decimal `json:"fixed_fee"`
	Currency          string          `json:"currency"`
}

func (h *PayRouteAdminHandler) ListMDRRates(w http.ResponseWriter, r *http.Request) {
	filter := model.Filter{Pagination: model.Pagination{Page: 1, PageSize: 50}}
	if providerID := r.URL.Query().Get("provider_account_id"); providerID != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field: "provider_account_id", Operator: model.OperatorEqual, Value: providerID,
		})
	}
	result, err := h.repo.ResolvePayrouteMdrRatesByFilter(r.Context(), filter)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *PayRouteAdminHandler) UpsertMDRRate(w http.ResponseWriter, r *http.Request) {
	var req upsertMDRRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	actorID := extractActorID(r)
	now := time.Now().UTC()

	item := model.PayrouteMdrRates{
		Id:                uuid.Must(uuid.NewV7()),
		ProviderAccountId: uuid.Must(uuid.FromString(req.ProviderAccountID)),
		PaymentMethodId:   uuid.Must(uuid.FromString(req.PaymentMethodID)),
		Percentage:        req.Percentage,
		FixedFee:          req.FixedFee,
		Currency:          req.Currency,
		MetaSignature: shared.MetaSignature{
			MetaCreatedAt: now,
			MetaCreatedBy: actorID,
			MetaUpdatedAt: null.TimeFrom(now),
			MetaUpdatedBy: &actorID,
		},
	}

	if req.PaymentChannelID != "" {
		chID := uuid.Must(uuid.FromString(req.PaymentChannelID))
		item.PaymentChannelId = &chID
	}

	if err := h.repo.CreatePayrouteMdrRates(r.Context(), &item); err != nil {
		writeErrorResponse(w, err)
		return
	}
	h.auditLog(r.Context(), actorID, "mdr_rate.upsert", "payroute_mdr_rates", item.Id, nil, item)
	writeJSONResponse(w, http.StatusCreated, item)
}

// ---------------------------------------------------------------------------
// Audit Log
// ---------------------------------------------------------------------------

func (h *PayRouteAdminHandler) ListAuditLog(w http.ResponseWriter, r *http.Request) {
	filter := model.Filter{Pagination: model.Pagination{Page: 1, PageSize: 50}}
	if entity := r.URL.Query().Get("target_entity"); entity != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field: "target_entity", Operator: model.OperatorEqual, Value: entity,
		})
	}
	if action := r.URL.Query().Get("action"); action != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field: "action", Operator: model.OperatorEqual, Value: action,
		})
	}
	result, err := h.repo.ResolvePayrouteAuditLogByFilter(r.Context(), filter)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"data": result})
}

// ---------------------------------------------------------------------------
// Config
// ---------------------------------------------------------------------------

func (h *PayRouteAdminHandler) ListConfig(w http.ResponseWriter, r *http.Request) {
	result, err := h.repo.ResolvePayrouteConfigByFilter(r.Context(), model.Filter{})
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{"data": result})
}

type updateConfigRequest struct {
	Value       json.RawMessage `json:"value"`
	Description string          `json:"description,omitempty"`
}

func (h *PayRouteAdminHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	var req updateConfigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, failure.BadRequest(err))
		return
	}

	actorID := extractActorID(r)
	now := time.Now().UTC()
	item := model.PayrouteConfig{
		Key:           key,
		Value:         req.Value,
		Description:   null.StringFrom(req.Description),
		UpdatedBy:     &actorID,
		MetaUpdatedAt: null.TimeFrom(now),
	}
	if err := h.repo.UpsertPayrouteConfig(r.Context(), &item); err != nil {
		writeErrorResponse(w, err)
		return
	}
	h.auditLog(r.Context(), actorID, "config.update", "payroute_config", uuid.Nil, nil, item)
	writeJSONResponse(w, http.StatusOK, item)
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func (h *PayRouteAdminHandler) createVersionSnapshot(ctx context.Context, rule model.PayrouteRoutingRules, actorID uuid.UUID, reason string) {
	snapshot, _ := json.Marshal(rule)
	version := model.PayrouteRoutingRuleVersions{
		Id:           uuid.Must(uuid.NewV7()),
		RuleId:       rule.Id,
		Version:      rule.Version,
		Snapshot:     snapshot,
		ChangedBy:    actorID,
		ChangeReason: null.StringFrom(reason),
		MetaSignature: shared.MetaSignature{
			MetaCreatedAt: time.Now().UTC(),
			MetaCreatedBy: actorID,
		},
	}
	_ = h.repo.CreatePayrouteRoutingRuleVersions(ctx, &version)
}

func (h *PayRouteAdminHandler) auditLog(ctx context.Context, actorID uuid.UUID, action, target string, targetID uuid.UUID, before, after interface{}) {
	beforeJSON, _ := json.Marshal(before)
	afterJSON, _ := json.Marshal(after)
	entry := model.PayrouteAuditLog{
		Id:           uuid.Must(uuid.NewV7()),
		ActorUserId:  actorID,
		Action:       action,
		TargetEntity: target,
		TargetId:     &targetID,
		BeforeValue:  beforeJSON,
		AfterValue:   afterJSON,
		MetaCreatedAt: null.TimeFrom(time.Now().UTC()),
	}
	_ = h.repo.CreatePayrouteAuditLog(ctx, &entry)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	if e, ok := err.(*failure.Failure); ok {
		code = e.Code
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":  err.Error(),
		"status": code,
	})
}

func extractActorID(r *http.Request) uuid.UUID {
	if id := r.Header.Get("X-Actor-ID"); id != "" {
		if parsed, err := uuid.FromString(id); err == nil {
			return parsed
		}
	}
	return uuid.Nil
}

// Ensure imports are used
var _ = pg.ProviderXendit

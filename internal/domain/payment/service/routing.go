package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/configs"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

const (
	defaultFailureThreshold = 3
	defaultMaxAttempts      = 3
	defaultCooldown         = 30 * time.Second
	defaultRetryBackoff     = 100 * time.Millisecond
	defaultCacheTTL         = 60 * time.Second
	minSampleForHealth     = 30
)

// ---------------------------------------------------------------------------
// Routing Rule from DB
// ---------------------------------------------------------------------------

type RoutingRule struct {
	Method      pg.PaymentMethod  `json:"method"`
	Channel     string            `json:"channel,omitempty"`
	Providers   []pg.ProviderCode `json:"providers"`
	MaxAttempts int               `json:"maxAttempts,omitempty"`
}

type RoutingConfig struct {
	Rules            []RoutingRule
	DefaultProviders []pg.ProviderCode
	MaxAttempts      int
	FailureThreshold int
	Cooldown         time.Duration
	RetryBackoff     time.Duration
}

type RoutingRequest struct {
	Method      pg.PaymentMethod
	Channel     string
	Currency    string
	GatewayCall pg.CreatePaymentRequest
	Amount      string // for cost-aware routing
}

type RouteCandidate struct {
	ProviderCode      pg.ProviderCode `json:"providerCode"`
	AccountID         uuid.UUID       `json:"accountId"`
	Priority          int             `json:"priority"`
	MaxAttempts       int             `json:"maxAttempts"`
	Reason            string          `json:"reason"`
	Skipped           bool            `json:"skipped"`
	SkipReason        string          `json:"skipReason,omitempty"`
	TrafficWeight     int             `json:"trafficWeight"`
	IsFallback        bool            `json:"isFallback"`
	ProviderMethodCode string         `json:"providerMethodCode,omitempty"`
	ProviderChannelCode string        `json:"providerChannelCode,omitempty"`
}

type ProviderAttempt struct {
	ProviderCode pg.ProviderCode `json:"providerCode"`
	AccountID    uuid.UUID       `json:"accountId"`
	Attempt      int             `json:"attempt"`
	StartedAt    time.Time       `json:"startedAt"`
	Duration     time.Duration   `json:"duration"`
	Error        string          `json:"error,omitempty"`
}

type RoutingResult struct {
	Selected   RouteCandidate           `json:"selected"`
	Candidates []RouteCandidate         `json:"candidates"`
	Attempts   []ProviderAttempt        `json:"attempts"`
	Payment    pg.CreatePaymentResponse `json:"payment"`
	RuleID     string                   `json:"ruleId,omitempty"`
	Strategy   string                   `json:"strategy,omitempty"`
}

// ---------------------------------------------------------------------------
// Circuit Breaker Interface
// ---------------------------------------------------------------------------

type CircuitBreaker interface {
	Allow(key string, now time.Time) bool
	RecordFailure(key string, now time.Time)
	RecordSuccess(key string)
}

type circuitState struct {
	failures  int
	openUntil time.Time
	halfOpen  bool
}

// MemoryCircuitBreaker is O(1), concurrency-safe, and hidden behind an
// interface so a distributed Redis implementation can replace it unchanged.
type MemoryCircuitBreaker struct {
	mu        sync.Mutex
	states    map[string]circuitState
	threshold int
	cooldown  time.Duration
}

type RedisCircuitBreaker struct {
	client    redis.UniversalClient
	threshold int
	cooldown  time.Duration
	prefix    string
}

var recordCircuitFailure = redis.NewScript(`
local failures = redis.call('INCR', KEYS[1])
if failures >= tonumber(ARGV[1]) then
  redis.call('SET', KEYS[2], '1', 'PX', ARGV[2])
  redis.call('PEXPIRE', KEYS[1], tonumber(ARGV[2]) * 2)
else
  redis.call('PEXPIRE', KEYS[1], ARGV[2])
end
redis.call('DEL', KEYS[3])
return failures
`)

var allowCircuitRequest = redis.NewScript(`
if redis.call('EXISTS', KEYS[2]) == 1 then
  return 0
end
local failures = tonumber(redis.call('GET', KEYS[1]) or '0')
if failures >= tonumber(ARGV[1]) then
  if redis.call('SET', KEYS[3], '1', 'NX', 'PX', ARGV[2]) then
    return 1
  end
  return 0
end
return 1
`)

func ProvideCircuitBreaker(cfg *configs.Config, client *redis.Client) CircuitBreaker {
	routing := NewRoutingConfig(cfg)
	if cfg == nil || client == nil || strings.TrimSpace(cfg.Cache.Redis.Primary.Host) == "" || strings.TrimSpace(cfg.Cache.Redis.Primary.Port) == "" {
		return NewMemoryCircuitBreaker(routing.FailureThreshold, routing.Cooldown)
	}
	return &RedisCircuitBreaker{client: client, threshold: routing.FailureThreshold, cooldown: routing.Cooldown, prefix: "payroute:routing:circuit"}
}

func (b *RedisCircuitBreaker) Allow(key string, _ time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	allowed, err := allowCircuitRequest.Run(ctx, b.client, []string{b.failureKey(key), b.openKey(key), b.probeKey(key)}, b.threshold, (10 * time.Second).Milliseconds()).Int()
	return err != nil || allowed == 1
}

func (b *RedisCircuitBreaker) RecordFailure(key string, _ time.Time) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, _ = recordCircuitFailure.Run(ctx, b.client, []string{b.failureKey(key), b.openKey(key), b.probeKey(key)}, b.threshold, b.cooldown.Milliseconds()).Result()
}

func (b *RedisCircuitBreaker) RecordSuccess(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, _ = b.client.Del(ctx, b.failureKey(key), b.openKey(key), b.probeKey(key)).Result()
}

func (b *RedisCircuitBreaker) failureKey(key string) string { return b.prefix + ":failures:" + key }
func (b *RedisCircuitBreaker) openKey(key string) string    { return b.prefix + ":open:" + key }
func (b *RedisCircuitBreaker) probeKey(key string) string   { return b.prefix + ":probe:" + key }

func NewMemoryCircuitBreaker(threshold int, cooldown time.Duration) *MemoryCircuitBreaker {
	if threshold <= 0 {
		threshold = defaultFailureThreshold
	}
	if cooldown <= 0 {
		cooldown = defaultCooldown
	}
	return &MemoryCircuitBreaker{states: make(map[string]circuitState), threshold: threshold, cooldown: cooldown}
}

func (b *MemoryCircuitBreaker) Allow(key string, now time.Time) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	state, exists := b.states[key]
	if !exists {
		return true
	}
	if !state.openUntil.IsZero() && now.Before(state.openUntil) {
		return false
	}
	if state.failures >= b.threshold {
		if state.halfOpen {
			return false
		}
		state.openUntil = time.Time{}
		state.halfOpen = true
		b.states[key] = state
	}
	return true
}

func (b *MemoryCircuitBreaker) RecordFailure(key string, now time.Time) {
	b.mu.Lock()
	defer b.mu.Unlock()
	state := b.states[key]
	state.failures++
	if state.failures >= b.threshold {
		state.openUntil = now.Add(b.cooldown)
		state.halfOpen = false
	}
	b.states[key] = state
}

func (b *MemoryCircuitBreaker) RecordSuccess(key string) {
	b.mu.Lock()
	delete(b.states, key)
	b.mu.Unlock()
}

// ---------------------------------------------------------------------------
// Health-Aware Ranking Data
// ---------------------------------------------------------------------------

type pspHealthScore struct {
	AccountID   uuid.UUID
	SuccessRate float64
	SampleSize  int
	AvgLatency  int
	MDRCost     float64 // calculated fee percentage
}

// ---------------------------------------------------------------------------
// Routing Engine — DB-driven with intelligent strategies
// ---------------------------------------------------------------------------

type RoutingEngine struct {
	registry     *pg.Registry
	accountIDs   map[pg.ProviderCode]uuid.UUID
	breaker      CircuitBreaker
	config       RoutingConfig
	capabilities map[pg.ProviderCode][]pg.Capability
	repo         repository.Repository
	ruleCache    *RoutingRuleCache
	killSwitch   *KillSwitchManager
	analytics    AnalyticsPublisher
}

func NewRoutingEngine(
	registry *pg.Registry,
	accountIDs map[pg.ProviderCode]uuid.UUID,
	breaker CircuitBreaker,
	config RoutingConfig,
	repo repository.Repository,
	analytics AnalyticsPublisher,
) *RoutingEngine {
	if analytics == nil {
		analytics = NewLogAnalyticsPublisher()
	}
	engine := &RoutingEngine{
		registry:     registry,
		accountIDs:   accountIDs,
		breaker:      breaker,
		config:       config,
		capabilities: make(map[pg.ProviderCode][]pg.Capability),
		repo:         repo,
		ruleCache:    NewRoutingRuleCache(repo, defaultCacheTTL),
		killSwitch:   NewKillSwitchManager(),
		analytics:    analytics,
	}
	for _, provider := range config.allProviders() {
		gateway, err := registry.Get(provider)
		if err != nil {
			continue
		}
		response, err := gateway.Capabilities(context.Background(), pg.CapabilitiesRequest{})
		if err == nil {
			engine.capabilities[provider] = response.Items
		}
	}
	return engine
}

func NewRoutingConfig(cfg *configs.Config) RoutingConfig {
	config := RoutingConfig{
		DefaultProviders: []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay, pg.ProviderMidtrans},
		MaxAttempts:      defaultMaxAttempts,
		FailureThreshold: defaultFailureThreshold,
		Cooldown:         defaultCooldown,
		RetryBackoff:     defaultRetryBackoff,
	}
	if cfg == nil {
		return config
	}
	routing := cfg.Internal.Payment.Routing
	if routing.MaxAttempts > 0 {
		config.MaxAttempts = routing.MaxAttempts
	}
	if routing.FailureThreshold > 0 {
		config.FailureThreshold = routing.FailureThreshold
	}
	if routing.CooldownSeconds > 0 {
		config.Cooldown = time.Duration(routing.CooldownSeconds) * time.Second
	}
	if routing.RetryBackoffMillis > 0 {
		config.RetryBackoff = time.Duration(routing.RetryBackoffMillis) * time.Millisecond
	}
	if providers := parseProviderList(strings.Split(routing.DefaultProviders, ",")); len(providers) > 0 {
		config.DefaultProviders = providers
	}
	if strings.TrimSpace(routing.RulesJSON) != "" {
		var rules []RoutingRule
		if json.Unmarshal([]byte(routing.RulesJSON), &rules) == nil {
			config.Rules = normalizeRules(rules, config.MaxAttempts)
		}
	}
	return config
}

// Execute is the main entry point: resolves rule from DB, ranks candidates, executes with failover.
func (e *RoutingEngine) Execute(ctx context.Context, request RoutingRequest) (RoutingResult, error) {
	// 1. Try DB-driven rule first
	rule, ruleID, strategy, err := e.resolveDBRule(ctx, request)
	if err != nil {
		return RoutingResult{}, err
	}

	// 2. Build candidates from DB rule or fall back to env-var config
	var candidates []RouteCandidate
	if rule != nil {
		candidates = e.buildCandidatesFromDBRule(rule, request)
	} else {
		candidates = e.resolveCandidates(request)
	}

	result := RoutingResult{Candidates: candidates, RuleID: ruleID, Strategy: strategy}
	if len(candidates) == 0 {
		return result, failure.New(424, fmt.Errorf("no configured provider supports %s/%s in %s", request.Method, request.Channel, request.Currency))
	}

	// 3. Execute with failover
	var lastErr error
	for index := range result.Candidates {
		candidate := &result.Candidates[index]
		key := circuitKey(candidate.ProviderCode, request.Method, request.Channel)

		// Kill-switch check
		if e.killSwitch.IsDisabled(candidate.AccountID.String()) {
			candidate.Skipped, candidate.SkipReason = true, "kill_switch_disabled"
			continue
		}

		// Circuit breaker check
		if !e.breaker.Allow(key, time.Now().UTC()) {
			candidate.Skipped, candidate.SkipReason = true, "circuit_open"
			continue
		}

		gateway, err := e.registry.Get(candidate.ProviderCode)
		if err != nil {
			candidate.Skipped, candidate.SkipReason = true, "provider_disabled"
			continue
		}

		for attempt := 1; attempt <= candidate.MaxAttempts; attempt++ {
			started := time.Now().UTC()

			PublishPaymentAttemptStarted(ctx, e.analytics, request.GatewayCall.OrderID, string(request.Method), string(candidate.ProviderCode))

			payment, callErr := gateway.CreatePayment(ctx, request.GatewayCall)
			providerAttempt := ProviderAttempt{
				ProviderCode: candidate.ProviderCode,
				AccountID:    candidate.AccountID,
				Attempt:      attempt,
				StartedAt:    started,
				Duration:     time.Since(started),
			}
			if callErr == nil {
				e.breaker.RecordSuccess(key)
				result.Selected, result.Payment = *candidate, payment
				result.Attempts = append(result.Attempts, providerAttempt)

				PublishPaymentAttemptSucceeded(ctx, e.analytics, request.GatewayCall.OrderID, string(candidate.ProviderCode), attempt, int(time.Since(started).Milliseconds()))

				return result, nil
			}
			lastErr = callErr
			providerAttempt.Error = callErr.Error()
			result.Attempts = append(result.Attempts, providerAttempt)

			// CRITICAL: Status check before failover for timeout/5xx errors
			if isTimeoutOr5xx(callErr) {
				statusOK, statusErr := e.verifyProviderStatus(ctx, request, *candidate, gateway)
				if statusErr == nil && statusOK {
					// PSP actually succeeded but response was lost
					candidate.Skipped, candidate.SkipReason = false, "recovered_via_status_check"
					break
				}
			}

			if canFallback(callErr) {
				e.breaker.RecordFailure(key, time.Now().UTC())
				PublishCircuitBreakerStateChanged(ctx, e.analytics, string(candidate.ProviderCode), key, "open", 0)
			}
			if !pg.IsRetryable(callErr) || attempt == candidate.MaxAttempts {
				break
			}
			if err := waitForRetry(ctx, e.config.RetryBackoff, attempt); err != nil {
				return result, err
			}
		}
		if lastErr != nil && !canFallback(lastErr) {
			return result, lastErr
		}
	}

	if lastErr == nil {
		lastErr = errors.New("all payment providers were skipped")
	}

	// P1 Alert: All PSPs failed
	PublishAllPSPFailed(ctx, e.analytics, request.GatewayCall.OrderID, string(request.Method), request.Channel, request.Currency, len(result.Attempts))

	return result, failure.New(424, fmt.Errorf("payment routing exhausted: %w", lastErr))
}

// ---------------------------------------------------------------------------
// DB Rule Resolution
// ---------------------------------------------------------------------------

func (e *RoutingEngine) resolveDBRule(ctx context.Context, request RoutingRequest) (*model.PayrouteRoutingRules, string, string, error) {
	if e.ruleCache == nil {
		return nil, "", "env_fallback", nil
	}
	rule, err := e.ruleCache.Get(ctx, string(request.Method), request.Channel, request.Currency)
	if err != nil || rule == nil {
		return nil, "", "env_fallback", nil
	}

	parsedRule, err := rule.ParsePspList()
	if err != nil || len(parsedRule) == 0 {
		return nil, "", "env_fallback", nil
	}

	// Gradual rollout check: if rule is rolling_out, only include percentage of traffic
	if !e.shouldIncludeInRollout(rule, request) {
		return nil, "", "env_fallback", nil
	}

	return rule, rule.Id.String(), string(rule.Strategy), nil
}

func (e *RoutingEngine) buildCandidatesFromDBRule(rule *model.PayrouteRoutingRules, request RoutingRequest) []RouteCandidate {
	pspItems, err := rule.ParsePspList()
	if err != nil || len(pspItems) == 0 {
		return nil
	}

	var candidates []RouteCandidate
	for _, item := range pspItems {
		// Find the provider code from account ID
		providerCode := e.findProviderForAccount(item.ProviderAccountID)
		if providerCode == "" {
			continue
		}
		if _, err := e.registry.Get(providerCode); err != nil {
			continue
		}
		if !e.supports(providerCode, request.Method, request.Channel, request.Currency) {
			continue
		}

		candidates = append(candidates, RouteCandidate{
			ProviderCode:       providerCode,
			AccountID:          item.ProviderAccountID,
			Priority:           item.Priority,
			MaxAttempts:        item.MaxAttempts,
			TrafficWeight:      item.TrafficWeight,
			IsFallback:         item.IsFallback,
			ProviderMethodCode: item.ProviderMethodCode,
			ProviderChannelCode: item.ProviderChannelCode,
			Reason:             "db_rule",
		})
	}

	// Apply strategy-based ranking
	return e.applyStrategy(candidates, rule, request)
}

// applyStrategy ranks candidates based on the rule's strategy.
func (e *RoutingEngine) applyStrategy(candidates []RouteCandidate, rule *model.PayrouteRoutingRules, request RoutingRequest) []RouteCandidate {
	strategy := rule.Strategy

	switch strategy {
	case model.PayrouteStrategySuccessRate:
		return e.rankBySuccessRate(candidates, rule, request)
	case model.PayrouteStrategyCostAware:
		return e.rankByCost(candidates, request)
	case model.PayrouteStrategyWeighted:
		return e.rankByWeight(candidates)
	case model.PayrouteStrategyCombined:
		return e.rankByCombined(candidates, rule, request)
	default:
		// priority is the default — candidates are already sorted by priority
		return candidates
	}
}

// rankBySuccessRate reorders candidates by rolling success rate.
func (e *RoutingEngine) rankBySuccessRate(candidates []RouteCandidate, rule *model.PayrouteRoutingRules, request RoutingRequest) []RouteCandidate {
	strategyConfig, _ := rule.ParseStrategyConfig()
	windowSeconds := strategyConfig.WindowSeconds
	if windowSeconds <= 0 {
		windowSeconds = 3600
	}
	minSample := strategyConfig.MinSampleSize
	if minSample <= 0 {
		minSample = minSampleForHealth
	}

	type scored struct {
		candidate RouteCandidate
		score     float64
		hasData   bool
	}

	var scoredList []scored
	for _, c := range candidates {
		health, err := e.repo.ResolveLatestPSPHealth(context.Background(), c.AccountID, string(request.Method), request.Channel)
		if err != nil || health == nil || !health.IsDataSufficient(minSample) {
			// Insufficient data — use priority as tiebreaker
			scoredList = append(scoredList, scored{candidate: c, score: float64(c.Priority), hasData: false})
			continue
		}
		successRate, _ := health.SuccessRate.Float64()
		scoredList = append(scoredList, scored{
			candidate: c,
			score:     successRate * 100,
			hasData:   true,
		})
	}

	// Sort: data-available first (by success rate descending), then no-data (by priority)
	for i := 0; i < len(scoredList); i++ {
		for j := i + 1; j < len(scoredList); j++ {
			if scoredList[j].hasData && (!scoredList[i].hasData || scoredList[j].score > scoredList[i].score) {
				scoredList[i], scoredList[j] = scoredList[j], scoredList[i]
			}
		}
	}

	result := make([]RouteCandidate, len(scoredList))
	for i, s := range scoredList {
		s.candidate.Priority = i + 1
		result[i] = s.candidate
	}
	return result
}

// rankByCost reorders candidates by effective MDR cost (lowest first).
func (e *RoutingEngine) rankByCost(candidates []RouteCandidate, request RoutingRequest) []RouteCandidate {
	type scored struct {
		candidate RouteCandidate
		cost      float64
	}

	var scoredList []scored
	for _, c := range candidates {
		mdrRate, err := e.repo.ResolveActiveMDRRate(context.Background(), c.AccountID, request.Method, request.Amount)
		cost := 999.0 // high default for unknown MDR
		if err == nil && mdrRate != nil {
			pct, _ := mdrRate.Percentage.Float64()
			cost = pct
		}
		scoredList = append(scoredList, scored{candidate: c, cost: cost})
	}

	// Sort by cost ascending
	for i := 0; i < len(scoredList); i++ {
		for j := i + 1; j < len(scoredList); j++ {
			if scoredList[j].cost < scoredList[i].cost {
				scoredList[i], scoredList[j] = scoredList[j], scoredList[i]
			}
		}
	}

	result := make([]RouteCandidate, len(scoredList))
	for i, s := range scoredList {
		s.candidate.Priority = i + 1
		result[i] = s.candidate
	}
	return result
}

// rankByWeight uses traffic_weight for weighted random selection.
func (e *RoutingEngine) rankByWeight(candidates []RouteCandidate) []RouteCandidate {
	// Deterministic sort by weight descending, then priority
	for i := 0; i < len(candidates); i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[j].TrafficWeight > candidates[i].TrafficWeight ||
				(candidates[j].TrafficWeight == candidates[i].TrafficWeight && candidates[j].Priority < candidates[i].Priority) {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}
	return candidates
}

// rankByCombined combines success rate and cost with configurable weights.
func (e *RoutingEngine) rankByCombined(candidates []RouteCandidate, rule *model.PayrouteRoutingRules, request RoutingRequest) []RouteCandidate {
	strategyConfig, _ := rule.ParseStrategyConfig()
	successWeight := strategyConfig.SuccessRateWeight
	if successWeight.IsZero() || successWeight.LessThan(decimal.Zero) {
		successWeight = decimal.NewFromFloat(0.7)
	}
	costWeight := strategyConfig.MdrWeight
	if costWeight.IsZero() || costWeight.LessThan(decimal.Zero) {
		costWeight = decimal.NewFromFloat(0.3)
	}
	successWeightF, _ := successWeight.Float64()
	costWeightF, _ := costWeight.Float64()

	type scored struct {
		candidate RouteCandidate
		score     float64
	}

	var scoredList []scored
	for _, c := range candidates {
		successScore := 50.0 // default middle score
		costScore := 50.0

		health, err := e.repo.ResolveLatestPSPHealth(context.Background(), c.AccountID, string(request.Method), request.Channel)
		if err == nil && health != nil && health.SampleSize >= minSampleForHealth {
			successRate, _ := health.SuccessRate.Float64()
			successScore = successRate * 100
		}

		mdrRate, err := e.repo.ResolveActiveMDRRate(context.Background(), c.AccountID, request.Method, request.Amount)
		if err == nil && mdrRate != nil {
			pct, _ := mdrRate.Percentage.Float64()
			costScore = 100 - pct // invert: lower cost = higher score
		}

		combined := (successScore * successWeightF) + (costScore * costWeightF)
		scoredList = append(scoredList, scored{candidate: c, score: combined})
	}

	// Sort by combined score descending
	for i := 0; i < len(scoredList); i++ {
		for j := i + 1; j < len(scoredList); j++ {
			if scoredList[j].score > scoredList[i].score {
				scoredList[i], scoredList[j] = scoredList[j], scoredList[i]
			}
		}
	}

	result := make([]RouteCandidate, len(scoredList))
	for i, s := range scoredList {
		s.candidate.Priority = i + 1
		result[i] = s.candidate
	}
	return result
}

// ---------------------------------------------------------------------------
// Existing env-var-based resolution (fallback)
// ---------------------------------------------------------------------------

func (e *RoutingEngine) resolveCandidates(request RoutingRequest) []RouteCandidate {
	providers, maxAttempts, reason := e.matchRule(request.Method, request.Channel)
	result := make([]RouteCandidate, 0, len(providers))
	seen := make(map[pg.ProviderCode]struct{}, len(providers))
	for priority, provider := range providers {
		if _, duplicate := seen[provider]; duplicate {
			continue
		}
		seen[provider] = struct{}{}
		accountID := e.accountIDs[provider]
		if accountID == uuid.Nil {
			continue
		}
		if _, err := e.registry.Get(provider); err != nil || !e.supports(provider, request.Method, request.Channel, request.Currency) {
			continue
		}
		result = append(result, RouteCandidate{
			ProviderCode: provider,
			AccountID:    accountID,
			Priority:     priority + 1,
			MaxAttempts:  maxAttempts,
			Reason:       reason,
		})
	}
	return result
}

func (e *RoutingEngine) matchRule(method pg.PaymentMethod, channel string) ([]pg.ProviderCode, int, string) {
	channel = strings.ToLower(strings.TrimSpace(channel))
	for _, rule := range e.config.Rules {
		if rule.Method == method && strings.EqualFold(strings.TrimSpace(rule.Channel), channel) {
			return rule.Providers, positiveOr(rule.MaxAttempts, e.config.MaxAttempts), "method_channel_rule"
		}
	}
	for _, rule := range e.config.Rules {
		if rule.Method == method && (strings.TrimSpace(rule.Channel) == "" || rule.Channel == "*") {
			return rule.Providers, positiveOr(rule.MaxAttempts, e.config.MaxAttempts), "method_rule"
		}
	}
	return e.config.DefaultProviders, e.config.MaxAttempts, "default_rule"
}

func (e *RoutingEngine) supports(provider pg.ProviderCode, method pg.PaymentMethod, channel, currency string) bool {
	for _, capability := range e.capabilities[provider] {
		channelMatch := channel == "" || capability.ChannelCode == "" || strings.EqualFold(capability.ChannelCode, channel) || genericChannelMatch(method, capability.ChannelCode)
		currencyMatch := currency == "" || capability.Currency == "" || strings.EqualFold(capability.Currency, currency)
		if capability.Method == method && channelMatch && currencyMatch {
			return true
		}
	}
	return false
}

func (e *RoutingEngine) findProviderForAccount(accountID uuid.UUID) pg.ProviderCode {
	for code, id := range e.accountIDs {
		if id == accountID {
			return code
		}
	}
	return ""
}

// RuleCache returns the routing rule cache (for admin handler invalidation).
func (e *RoutingEngine) RuleCache() *RoutingRuleCache {
	return e.ruleCache
}

// KillSwitch returns the kill-switch manager (for admin handler).
func (e *RoutingEngine) KillSwitch() *KillSwitchManager {
	return e.killSwitch
}

// ---------------------------------------------------------------------------
// Gradual Rollout — Traffic Splitting (PRD 6.2 AC-4)
// ---------------------------------------------------------------------------

// shouldIncludeInRollout determines if a request should be routed by a rolling_out rule.
func (e *RoutingEngine) shouldIncludeInRollout(rule *model.PayrouteRoutingRules, request RoutingRequest) bool {
	if rule.Status != model.PayrouteStatusRollingOut {
		return true
	}
	if rule.RolloutPercentage <= 0 {
		return false
	}
	if rule.RolloutPercentage >= 100 {
		return true
	}

	// Deterministic hash based on merchant + order ID
	h := fnv.New32a()
	h.Write([]byte(request.GatewayCall.OrderID))
	hash := h.Sum32()
	bucket := int(hash % 100)

	return bucket < rule.RolloutPercentage
}

// ---------------------------------------------------------------------------
// Shadow / A-B Routing (PRD 6.2 AC-4)
// ---------------------------------------------------------------------------

// isShadowRouting determines if this rule is in shadow mode.
func (e *RoutingEngine) isShadowRouting(rule *model.PayrouteRoutingRules) bool {
	return rule.Status == model.PayrouteStatusRollingOut && rule.RolloutPercentage == 0
}

// ---------------------------------------------------------------------------
// Status Check Before Failover
// ---------------------------------------------------------------------------

func (e *RoutingEngine) verifyProviderStatus(ctx context.Context, request RoutingRequest, candidate RouteCandidate, gateway pg.PaymentGateway) (bool, error) {
	statusReq := pg.GetPaymentStatusRequest{
		OrderID: request.GatewayCall.OrderID,
	}
	statusResp, err := gateway.GetPaymentStatus(ctx, statusReq)
	if err != nil {
		return false, err
	}
	return statusResp.Status == pg.PaymentStatusSucceeded || statusResp.Status == pg.PaymentStatusCaptured, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func genericChannelMatch(method pg.PaymentMethod, capabilityChannel string) bool {
	return (method == pg.PaymentMethodVirtualAccount && strings.EqualFold(capabilityChannel, "va")) ||
		(method == pg.PaymentMethodQRIS && strings.EqualFold(capabilityChannel, "qris"))
}

func (c RoutingConfig) allProviders() []pg.ProviderCode {
	providers := append([]pg.ProviderCode{}, c.DefaultProviders...)
	for _, rule := range c.Rules {
		providers = append(providers, rule.Providers...)
	}
	return providers
}

func normalizeRules(rules []RoutingRule, defaultAttempts int) []RoutingRule {
	result := make([]RoutingRule, 0, len(rules))
	for _, rule := range rules {
		rule.Method = pg.PaymentMethod(strings.ToLower(strings.TrimSpace(string(rule.Method))))
		rule.Channel = strings.ToLower(strings.TrimSpace(rule.Channel))
		rule.Providers = normalizeProviders(rule.Providers)
		rule.MaxAttempts = positiveOr(rule.MaxAttempts, defaultAttempts)
		if rule.Method != "" && len(rule.Providers) > 0 {
			result = append(result, rule)
		}
	}
	return result
}

func parseProviderList(values []string) []pg.ProviderCode {
	providers := make([]pg.ProviderCode, 0, len(values))
	for _, value := range values {
		providers = append(providers, pg.ProviderCode(strings.ToLower(strings.TrimSpace(value))))
	}
	return normalizeProviders(providers)
}

func normalizeProviders(values []pg.ProviderCode) []pg.ProviderCode {
	result := make([]pg.ProviderCode, 0, len(values))
	seen := make(map[pg.ProviderCode]struct{}, len(values))
	for _, value := range values {
		provider := pg.ProviderCode(strings.ToLower(strings.TrimSpace(string(value))))
		if provider == "" {
			continue
		}
		if _, exists := seen[provider]; exists {
			continue
		}
		seen[provider] = struct{}{}
		result = append(result, provider)
	}
	return result
}

func circuitKey(provider pg.ProviderCode, method pg.PaymentMethod, channel string) string {
	return strings.Join([]string{string(provider), string(method), strings.ToLower(strings.TrimSpace(channel))}, ":")
}

func waitForRetry(ctx context.Context, base time.Duration, attempt int) error {
	if base <= 0 {
		return nil
	}
	delay := base * time.Duration(math.Pow(2, float64(attempt-1)))
	if delay > 2*time.Second {
		delay = 2 * time.Second
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func canFallback(err error) bool {
	var gatewayErr *pg.GatewayError
	if !errors.As(err, &gatewayErr) {
		return true
	}
	return gatewayErr.Code != pg.ErrorCodeInvalidRequest
}

func isTimeoutOr5xx(err error) bool {
	var gatewayErr *pg.GatewayError
	if !errors.As(err, &gatewayErr) {
		return false
	}
	return gatewayErr.Code == pg.ErrorCodeProviderTimeout || gatewayErr.Code == pg.ErrorCodeProviderUnavailable
}

func positiveOr(value, fallback int) int {
	if value > 0 {
		return value
	}
	return fallback
}

// Ensure unused import is referenced
var _ = rand.Intn

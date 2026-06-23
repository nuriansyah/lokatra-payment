package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/configs"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

const (
	defaultFailureThreshold = 3
	defaultMaxAttempts      = 3
	defaultCooldown         = 30 * time.Second
	defaultRetryBackoff     = 100 * time.Millisecond
)

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
}

type RouteCandidate struct {
	ProviderCode pg.ProviderCode `json:"providerCode"`
	AccountID    uuid.UUID       `json:"accountId"`
	Priority     int             `json:"priority"`
	MaxAttempts  int             `json:"maxAttempts"`
	Reason       string          `json:"reason"`
	Skipped      bool            `json:"skipped"`
	SkipReason   string          `json:"skipReason,omitempty"`
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
}

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
	return &RedisCircuitBreaker{client: client, threshold: routing.FailureThreshold, cooldown: routing.Cooldown, prefix: "payment:routing:circuit"}
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

type RoutingEngine struct {
	registry     *pg.Registry
	accountIDs   map[pg.ProviderCode]uuid.UUID
	breaker      CircuitBreaker
	config       RoutingConfig
	capabilities map[pg.ProviderCode][]pg.Capability
}

func NewRoutingEngine(registry *pg.Registry, accountIDs map[pg.ProviderCode]uuid.UUID, breaker CircuitBreaker, config RoutingConfig) *RoutingEngine {
	engine := &RoutingEngine{registry: registry, accountIDs: accountIDs, breaker: breaker, config: config, capabilities: make(map[pg.ProviderCode][]pg.Capability)}
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

func (e *RoutingEngine) Execute(ctx context.Context, request RoutingRequest) (RoutingResult, error) {
	candidates := e.resolveCandidates(request)
	result := RoutingResult{Candidates: candidates}
	if len(candidates) == 0 {
		return result, failure.New(424, fmt.Errorf("no configured provider supports %s/%s in %s", request.Method, request.Channel, request.Currency))
	}
	var lastErr error
	for index := range result.Candidates {
		candidate := &result.Candidates[index]
		key := circuitKey(candidate.ProviderCode, request.Method, request.Channel)
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
			payment, callErr := gateway.CreatePayment(ctx, request.GatewayCall)
			providerAttempt := ProviderAttempt{ProviderCode: candidate.ProviderCode, AccountID: candidate.AccountID, Attempt: attempt, StartedAt: started, Duration: time.Since(started)}
			if callErr == nil {
				e.breaker.RecordSuccess(key)
				result.Selected, result.Payment = *candidate, payment
				result.Attempts = append(result.Attempts, providerAttempt)
				return result, nil
			}
			lastErr = callErr
			providerAttempt.Error = callErr.Error()
			result.Attempts = append(result.Attempts, providerAttempt)
			if canFallback(callErr) {
				e.breaker.RecordFailure(key, time.Now().UTC())
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
	return result, failure.New(424, fmt.Errorf("payment routing exhausted: %w", lastErr))
}

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
		result = append(result, RouteCandidate{ProviderCode: provider, AccountID: accountID, Priority: priority + 1, MaxAttempts: maxAttempts, Reason: reason})
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

func genericChannelMatch(method pg.PaymentMethod, capabilityChannel string) bool {
	return (method == pg.PaymentMethodVirtualAccount && strings.EqualFold(capabilityChannel, "va")) || (method == pg.PaymentMethodQRIS && strings.EqualFold(capabilityChannel, "qris"))
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
	delay := base * time.Duration(1<<(attempt-1))
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

func positiveOr(value, fallback int) int {
	if value > 0 {
		return value
	}
	return fallback
}

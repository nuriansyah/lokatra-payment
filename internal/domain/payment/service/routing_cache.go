package service

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
)

// ---------------------------------------------------------------------------
// Routing Rule Cache — TTL-based in-memory cache for DB-driven rules
// ---------------------------------------------------------------------------

type routingRuleCacheEntry struct {
	rule    *model.PayrouteRoutingRules
	loaded  time.Time
}

type RoutingRuleCache struct {
	mu      sync.RWMutex
	entries map[string]*routingRuleCacheEntry
	ttl     time.Duration
	repo    repository.Repository
}

// NewRoutingRuleCache creates a cache that resolves rules from DB with TTL.
func NewRoutingRuleCache(repo repository.Repository, ttl time.Duration) *RoutingRuleCache {
	if ttl <= 0 {
		ttl = 60 * time.Second
	}
	return &RoutingRuleCache{
		entries: make(map[string]*routingRuleCacheEntry),
		ttl:     ttl,
		repo:    repo,
	}
}

// cacheKey builds a deterministic key from routing scope.
func cacheKey(methodCode, channelCode, currency string) string {
	return strings.ToLower(strings.TrimSpace(methodCode)) + ":" +
		strings.ToLower(strings.TrimSpace(channelCode)) + ":" +
		strings.ToUpper(strings.TrimSpace(currency))
}

// Get returns a cached rule or resolves from DB.
func (c *RoutingRuleCache) Get(ctx context.Context, methodCode, channelCode, currency string) (*model.PayrouteRoutingRules, error) {
	if c.repo == nil {
		return nil, nil
	}
	key := cacheKey(methodCode, channelCode, currency)

	// Fast path: check cache with read lock
	c.mu.RLock()
	if entry, ok := c.entries[key]; ok {
		if time.Since(entry.loaded) < c.ttl {
			c.mu.RUnlock()
			return entry.rule, nil
		}
	}
	c.mu.RUnlock()

	// Slow path: resolve from DB
	rule, err := c.repo.ResolveActiveRuleByScope(ctx, methodCode, channelCode, currency)
	if err != nil {
		return nil, err
	}

	// Cache the result (even if nil to avoid repeated DB lookups for missing rules)
	c.mu.Lock()
	c.entries[key] = &routingRuleCacheEntry{
		rule:   rule,
		loaded: time.Now(),
	}
	c.mu.Unlock()

	return rule, nil
}

// Invalidate removes a specific cache entry (called when rule changes).
func (c *RoutingRuleCache) Invalidate(methodCode, channelCode, currency string) {
	key := cacheKey(methodCode, channelCode, currency)
	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
}

// InvalidateAll clears the entire cache (called on kill-switch or bulk changes).
func (c *RoutingRuleCache) InvalidateAll() {
	c.mu.Lock()
	c.entries = make(map[string]*routingRuleCacheEntry)
	c.mu.Unlock()
}

// Size returns the number of cached entries (for observability).
func (c *RoutingRuleCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}

// ---------------------------------------------------------------------------
// Kill-Switch Manager — in-memory set of disabled PSPs
// ---------------------------------------------------------------------------

type KillSwitchManager struct {
	mu           sync.RWMutex
	disabledPSPs map[string]bool // key: provider_account_id string
}

func NewKillSwitchManager() *KillSwitchManager {
	return &KillSwitchManager{
		disabledPSPs: make(map[string]bool),
	}
}

// IsDisabled returns true if the PSP is globally disabled.
func (k *KillSwitchManager) IsDisabled(accountID string) bool {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.disabledPSPs[accountID]
}

// Disable adds a PSP to the disabled set.
func (k *KillSwitchManager) Disable(accountID string) {
	k.mu.Lock()
	k.disabledPSPs[accountID] = true
	k.mu.Unlock()
}

// Enable removes a PSP from the disabled set.
func (k *KillSwitchManager) Enable(accountID string) {
	k.mu.Lock()
	delete(k.disabledPSPs, accountID)
	k.mu.Unlock()
}

// DisabledList returns all currently disabled PSP IDs.
func (k *KillSwitchManager) DisabledList() []string {
	k.mu.RLock()
	defer k.mu.RUnlock()
	result := make([]string, 0, len(k.disabledPSPs))
	for id := range k.disabledPSPs {
		result = append(result, id)
	}
	return result
}

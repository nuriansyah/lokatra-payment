package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

// ---------------------------------------------------------------------------
// Analytics Publisher — Routing event instrumentation
// ---------------------------------------------------------------------------

type AnalyticsEvent struct {
	EventType  string                 `json:"event_type"`
	Properties map[string]interface{} `json:"properties"`
	Timestamp  time.Time              `json:"timestamp"`
}

type AnalyticsPublisher interface {
	Publish(ctx context.Context, event AnalyticsEvent) error
}

// ProvideAnalyticsPublisher creates the default analytics publisher (log-based).
func ProvideAnalyticsPublisher() AnalyticsPublisher {
	return NewLogAnalyticsPublisher()
}

// LogAnalyticsPublisher writes events to structured logs (can be replaced with
// Kafka/PubSub/etc. in production).
type LogAnalyticsPublisher struct{}

func NewLogAnalyticsPublisher() *LogAnalyticsPublisher {
	return &LogAnalyticsPublisher{}
}

func (p *LogAnalyticsPublisher) Publish(ctx context.Context, event AnalyticsEvent) error {
	data, _ := json.Marshal(event.Properties)
	log.Info().
		Str("event_type", event.EventType).
		RawJSON("properties", data).
		Time("timestamp", event.Timestamp).
		Msg("[Analytics]")
	return nil
}

// ---------------------------------------------------------------------------
// Routing Event Helpers — Standardized event factories
// ---------------------------------------------------------------------------

func PublishPaymentAttemptStarted(ctx context.Context, pub AnalyticsPublisher, transactionID, method, psp string) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "payment_attempt_started",
		Properties: map[string]interface{}{
			"transaction_id": transactionID,
			"method":         method,
			"initial_psp":    psp,
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishPaymentAttemptSucceeded(ctx context.Context, pub AnalyticsPublisher, transactionID, psp string, attempt int, latencyMs int) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "payment_attempt_succeeded",
		Properties: map[string]interface{}{
			"transaction_id": transactionID,
			"psp_used":       psp,
			"attempt_number": attempt,
			"latency_ms":     latencyMs,
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishFailoverTriggered(ctx context.Context, pub AnalyticsPublisher, transactionID, failedPSP, reason, fallbackPSP string) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "payment_failover_triggered",
		Properties: map[string]interface{}{
			"transaction_id": transactionID,
			"failed_psp":     failedPSP,
			"reason":         reason,
			"fallback_psp":   fallbackPSP,
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishRoutingRuleChanged(ctx context.Context, pub AnalyticsPublisher, ruleID, actor string, before, after interface{}) {
	beforeJSON, _ := json.Marshal(before)
	afterJSON, _ := json.Marshal(after)
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "routing_rule_changed",
		Properties: map[string]interface{}{
			"rule_id":     ruleID,
			"actor":       actor,
			"before":      string(beforeJSON),
			"after":       string(afterJSON),
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishCircuitBreakerStateChanged(ctx context.Context, pub AnalyticsPublisher, pspID, scope, newState string, metricValue float64) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "circuit_breaker_state_changed",
		Properties: map[string]interface{}{
			"psp_id":          pspID,
			"scope":           scope,
			"new_state":       newState,
			"trigger_value":   metricValue,
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishKillSwitchActivated(ctx context.Context, pub AnalyticsPublisher, pspID, actor, reason string) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "kill_switch_activated",
		Properties: map[string]interface{}{
			"psp_id": pspID,
			"actor":  actor,
			"reason": reason,
		},
		Timestamp: time.Now().UTC(),
	})
}

func PublishAllPSPFailed(ctx context.Context, pub AnalyticsPublisher, transactionID, method, channel, currency string, attempts int) {
	pub.Publish(ctx, AnalyticsEvent{
		EventType: "p1_alert_all_psp_failed",
		Properties: map[string]interface{}{
			"severity":       "P1",
			"transaction_id": transactionID,
			"method":         method,
			"channel":        channel,
			"currency":       currency,
			"total_attempts": attempts,
		},
		Timestamp: time.Now().UTC(),
	})
}

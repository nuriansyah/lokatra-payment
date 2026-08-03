package service

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
)

// ---------------------------------------------------------------------------
// Health Monitor — Background service that computes PSP health aggregates
// ---------------------------------------------------------------------------

type HealthMonitor struct {
	repo     repository.Repository
	interval time.Duration
	window   time.Duration
	stopCh   chan struct{}
}

// ProvideHealthMonitor creates a HealthMonitor with default config (60s interval, 1h window).
func ProvideHealthMonitor(repo repository.Repository) *HealthMonitor {
	return NewHealthMonitor(repo, 60*time.Second, 1*time.Hour)
}

// NewHealthMonitor creates a health monitor that periodically computes aggregates.
func NewHealthMonitor(repo repository.Repository, interval, window time.Duration) *HealthMonitor {
	if interval <= 0 {
		interval = 60 * time.Second
	}
	if window <= 0 {
		window = 1 * time.Hour
	}
	return &HealthMonitor{
		repo:     repo,
		interval: interval,
		window:   window,
		stopCh:   make(chan struct{}),
	}
}

// Start begins the background health computation loop.
func (m *HealthMonitor) Start(ctx context.Context) {
	log.Info().Dur("interval", m.interval).Dur("window", m.window).Msg("[HealthMonitor] starting")
	ticker := time.NewTicker(m.interval)
	defer ticker.Stop()

	// Run immediately on start
	m.compute(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("[HealthMonitor] stopped due to context cancellation")
			return
		case <-m.stopCh:
			log.Info().Msg("[HealthMonitor] stopped")
			return
		case <-ticker.C:
			m.compute(ctx)
		}
	}
}

// Stop signals the monitor to stop.
func (m *HealthMonitor) Stop() {
	close(m.stopCh)
}

func (m *HealthMonitor) compute(ctx context.Context) {
	since := time.Now().UTC().Add(-m.window)
	if err := m.repo.ComputeAndStoreHealthAggregates(ctx, since); err != nil {
		log.Error().Err(err).Msg("[HealthMonitor] failed to compute health aggregates")
		return
	}
	log.Debug().Time("since", since).Msg("[HealthMonitor] health aggregates computed successfully")
}

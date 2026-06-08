/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// Package monitoring provides alerting for PhoenixOS.
//
// ROLE: Observability Layer
// PURPOSE: Detect anomalies and alert operators
// DEPENDS ON: MetricsCollector, Logger
// DEPENDED BY: PhoenixDashboard, PhoenixMind
//
// ARCHITECTURE NOTE:
// This package implements the alerting strategy that was identified as
// HIGH priority in the adversarial audit (Q32). Without this,
// security incidents go undetected.
//
// AGENT INSTRUCTIONS:
// 1. Define AlertManager interface
// 2. Implement rule-based alerting
// 3. Add alert channels (email, Slack, PagerDuty)
// 4. Add alert suppression and deduplication
// 5. Add alert history and acknowledgment
//
// TODO ITEMS:
// - [ ] Define AlertManager interface
// - [ ] Define AlertRule struct
// - [ ] Implement RuleBasedAlertManager
// - [ ] Implement EmailAlerter
// - [ ] Implement SlackAlerter
// - [ ] Implement PagerDutyAlerter
// - [ ] Add alert suppression
// - [ ] Add alert deduplication
// - [ ] Add alert history
// - [ ] Add alert acknowledgment
// - [ ] Write unit tests for alerting logic
// - [ ] Write integration tests for alert channels
//
// SECURITY NOTES:
// - Alerts must be authenticated
// - Alert channels must be encrypted
// - Alert history must be tamper-evident
//
// REFERENCES:
// - PHASE_5_FORMAL_VERIFICATION_AND_SECURITY_ARCHITECTURE.md (Section 1.6: PhoenixGuard)
package monitoring

import (
	"context"
	"sync"
	"time"
)

// AlertManager handles the lifecycle and distribution of system alerts.
type AlertManager interface {
	CreateAlert(ctx context.Context, alert Alert) error
	AcknowledgeAlert(ctx context.Context, alertID string, userID string) error
	GetAlerts(ctx context.Context, filter AlertFilter) ([]Alert, error)
	Subscribe(ctx context.Context, channel string) (<-chan Alert, error)
}

// Severity defines the priority level of an alert.
type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityError    Severity = "error"
	SeverityCritical Severity = "critical"
)

// Alert represents a single notification of a system event or anomaly.
type Alert struct {
	ID             string
	Severity       Severity
	Title          string
	Description    string
	Source         string
	CreatedAt      time.Time
	AcknowledgedAt *time.Time
	AcknowledgedBy *string
	Labels         map[string]string
}

// AlertFilter provides criteria for querying active or historical alerts.
type AlertFilter struct {
	Severity Severity
	Source   string
	Since    time.Time
}

// Alerter defines the interface for specific alert delivery channels.
type Alerter interface {
	Send(ctx context.Context, alert Alert) error
}

// AlertRule defines a condition that triggers an alert based on system state.
type AlertRule struct {
	Name     string
	Severity Severity
	Message  string
	Cooldown time.Duration
}

// RuleBasedAlertManager implements AlertManager using a set of evaluation rules.
type RuleBasedAlertManager struct {
	rules    []AlertRule
	alerts   map[string]Alert
	alerters []Alerter
	channels map[string]chan Alert
	mu       sync.RWMutex
}

// CreateAlert records a new alert and broadcasts it to subscribers and alerters.
func (m *RuleBasedAlertManager) CreateAlert(ctx context.Context, alert Alert) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if alert.ID == "" {
		alert.ID = "ALRT-" + time.Now().Format("20060102-150405.000")
	}
	alert.CreatedAt = time.Now().UTC()

	m.alerts[alert.ID] = alert

	// Broadcast to subscribers
	for _, ch := range m.channels {
		select {
		case ch <- alert:
		default:
			// Non-blocking broadcast
		}
	}

	// Send to external alerters
	for _, alerter := range m.alerters {
		go func(a Alerter, alt Alert) {
			_ = a.Send(ctx, alt)
		}(alerter, alert)
	}

	return nil
}

// AcknowledgeAlert marks an alert as acknowledged by a user.
func (m *RuleBasedAlertManager) AcknowledgeAlert(ctx context.Context, alertID string, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	alert, ok := m.alerts[alertID]
	if !ok {
		return nil
	}

	now := time.Now().UTC()
	alert.AcknowledgedAt = &now
	alert.AcknowledgedBy = &userID
	m.alerts[alertID] = alert

	return nil
}

// GetAlerts returns a filtered list of alerts.
func (m *RuleBasedAlertManager) GetAlerts(ctx context.Context, filter AlertFilter) ([]Alert, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var results []Alert
	for _, alert := range m.alerts {
		if filter.Severity != "" && alert.Severity != filter.Severity {
			continue
		}
		if filter.Source != "" && alert.Source != filter.Source {
			continue
		}
		if !filter.Since.IsZero() && alert.CreatedAt.Before(filter.Since) {
			continue
		}
		results = append(results, alert)
	}
	return results, nil
}

// Subscribe returns a channel that receives new alerts.
func (m *RuleBasedAlertManager) Subscribe(ctx context.Context, channelID string) (<-chan Alert, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ch := make(chan Alert, 100)
	m.channels[channelID] = ch

	return ch, nil
}

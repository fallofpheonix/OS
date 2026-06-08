// Package monitor implements sensory analysis for PhoenixOS telemetry.
// Domain Logic: Transforms raw telemetry events into quantified drift scores using Kalman filters and EWMA baselines.
// Responsibility: Acts as the first analytical layer to detect anomalies in real-time telemetry streams.
package monitor

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/math/kalman"
)

// DriftScore quantifies how anomalous an event is relative to baseline behavior.
// Concurrency: Read-only instances are thread-safe.
// State Management: Encapsulates statistical metrics (Z-score, Drift, Frequency) for a specific event sequence.
type DriftScore struct {
	EventID       int64   `json:"seq_id"`
	PID           int     `json:"pid"`
	UID           int     `json:"uid"`
	EventType     string  `json:"event_type"`
	OriginalScore float64 `json:"original_score"`
	SmoothedScore float64 `json:"smoothed_score"`
	Baseline      float64 `json:"baseline"`
	ZScore        float64 `json:"z_score"`
	DriftScore    float64 `json:"drift_score"`
	Severity      float64 `json:"severity"`  // GS (Raw Severity)
	Frequency     float64 `json:"frequency"` // FA (Historical Frequency)
	WallTimeUnix  int64   `json:"wall_time_unix"`
}

// MonitorService is the sensory analysis engine.
// Concurrency: Not thread-safe for concurrent Process calls; intended for single-goroutine ingestion or external synchronization.
// State Management: Maintains Kalman filter state, EWMA baselines, and a frequency map for historical context.
type MonitorService struct {
	busCh   chan bus.TelemetryEvent
	outBus  *bus.Bus
	kalman  *kalman.KalmanFilter
	ewma    float64
	alpha   float64
	varEWMA float64
	// Frequency tracking: map[context_key]count
	freqMap map[string]float64
	totalEv int64
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewMonitorService creates a new monitor service instance.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewMonitorService(inCh chan bus.TelemetryEvent, outBus *bus.Bus) *MonitorService {
	return &MonitorService{
		busCh:   inCh,
		outBus:  outBus,
		kalman:  kalman.NewKalmanFilter(0.01, 0.1, 1.0, 0.0),
		alpha:   0.05,
		freqMap: make(map[string]float64),
	}
}

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// Start launches the background goroutine that continuously processes events.
// I/O: Reads from busCh, writes to outBus via Process.
// Side Effects: Starts a long-running goroutine.
// Complexity: O(1) to start, O(N) over N events.
func (m *MonitorService) Start() {
	go func() {
		for event := range m.busCh {
			m.Process(event)
		}
	}()
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// Process transforms a raw TelemetryEvent into a DriftScore.
// I/O: Publishes scored event to outBus.
// Side Effects: Updates internal statistical state (Kalman, EWMA, Frequency Map).
// Complexity: O(1) constant time operations.
func (m *MonitorService) Process(event bus.TelemetryEvent) DriftScore {
	raw := event.Severity.Float64()
	drift := m.kalman.CheckDrift(raw)
	smoothed := m.kalman.Update(raw)

	diff := smoothed - m.ewma
	m.ewma += m.alpha * diff
	m.varEWMA = (1 - m.alpha) * (m.varEWMA + m.alpha*diff*diff)

	stddev := math.Sqrt(m.varEWMA)
	var zscore float64
	if stddev > 0.001 {
		zscore = (smoothed - m.ewma) / stddev
	}

	// Track Frequency (FA)
	m.totalEv++
	ctxKey := fmt.Sprintf("%s:%d", event.EventType, event.UID)
	m.freqMap[ctxKey]++
	fa := m.freqMap[ctxKey] / float64(m.totalEv)

	score := DriftScore{
		EventID:       event.SeqID,
		PID:           event.PID,
		UID:           event.UID,
		EventType:     event.EventType,
		OriginalScore: raw,
		SmoothedScore: smoothed,
		Baseline:      m.ewma,
		ZScore:        zscore,
		DriftScore:    drift,
		Severity:      raw,
		Frequency:     fa,
		WallTimeUnix:  event.WallTimeUnix,
	}

	payloadBytes, _ := json.Marshal(score)

	m.outBus.Publish("telemetry.scored", bus.TelemetryEvent{
		SeqID:        event.SeqID,
		WallTimeUnix: event.WallTimeUnix,
		Source:       "phoenix.monitor",
		EventType:    "monitor.score",
		Severity:     event.Severity,
		Payload:      payloadBytes,
	})

	return score
}

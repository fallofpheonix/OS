package monitor

import (
	"encoding/json"
	"fmt"
	"math"

	"phoenix/bus"
	"phoenix/common/math/kalman"
)

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

func NewMonitorService(inCh chan bus.TelemetryEvent, outBus *bus.Bus) *MonitorService {
	return &MonitorService{
		busCh:   inCh,
		outBus:  outBus,
		kalman:  kalman.NewKalmanFilter(0.01, 0.1, 1.0, 0.0),
		alpha:   0.05,
		freqMap: make(map[string]float64),
	}
}

func (m *MonitorService) Start() {
	go func() {
		for event := range m.busCh {
			m.Process(event)
		}
	}()
}

func (m *MonitorService) Process(event bus.TelemetryEvent) DriftScore {
	raw := event.Severity
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
		Severity:     raw,
		Payload:      payloadBytes,
	})

	return score
}

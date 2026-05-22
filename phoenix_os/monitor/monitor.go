package monitor

import (
	"encoding/json"
	"math"

	"phoenix/bus"
	"phoenix/common/math/kalman"
)

type DriftScore struct {
	EventID         int64   `json:"seq_id"`
	OriginalScore   float64 `json:"original_score"`
	SmoothedScore   float64 `json:"smoothed_score"`
	Baseline        float64 `json:"baseline"`
	ZScore          float64 `json:"z_score"`
	DriftScore      float64 `json:"drift_score"`
	ImportanceScore float64 `json:"importance_score"` // SI (Multiplicative Factor)
	WallTimeUnix    int64   `json:"wall_time_unix"`
}

type MonitorService struct {
	busCh   chan bus.TelemetryEvent
	outBus  *bus.Bus
	kalman  *kalman.KalmanFilter
	ewma    float64
	alpha   float64
	varEWMA float64
}

func NewMonitorService(inCh chan bus.TelemetryEvent, outBus *bus.Bus) *MonitorService {
	return &MonitorService{
		busCh:  inCh,
		outBus: outBus,
		kalman: kalman.NewKalmanFilter(0.01, 0.1, 1.0, 0.0),
		alpha:  0.05,
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

	// SI Calculation: Combine UID, PID, and EventType criticality
	si := 1.0
	if event.UID < 100 { // System users
		si *= 1.5
	}
	if event.PID == 1 || event.PID < 500 { // Core system processes
		si *= 1.2
	}
	if event.EventType == "execve" || event.EventType == "ptrace" {
		si *= 1.3
	}

	score := DriftScore{
		EventID:         event.SeqID,
		OriginalScore:   raw,
		SmoothedScore:   smoothed,
		Baseline:        m.ewma,
		ZScore:          zscore,
		DriftScore:      drift,
		ImportanceScore: si,
		WallTimeUnix:    event.WallTimeUnix,
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

package monitor

import (
	"encoding/json"
	"math"

	"github.com/fallofpheonix/phoenix_os/bus"
	"github.com/fallofpheonix/PheonixOS/phoenix_os/agents/internal/telemetry/math"
)

// DriftScore represents the anomaly evaluation of an event
type DriftScore struct {
	EventID       int64   `json:"seq_id"`
	OriginalScore float64 `json:"original_score"`
	SmoothedScore float64 `json:"smoothed_score"`
	Baseline      float64 `json:"baseline"`
	ZScore        float64 `json:"z_score"`
	DriftScore    float64 `json:"drift_score"` // Kalman innovation score
}

// MonitorService observes the bus and applies EWMA + Kalman
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
			raw := event.Severity
			
			// 1. Calculate drift score before update
			drift := m.kalman.CheckDrift(raw)
			
			// 2. Update filter
			smoothed := m.kalman.Update(raw)

			// 3. EWMA
			diff := smoothed - m.ewma
			m.ewma += m.alpha * diff
			m.varEWMA = (1 - m.alpha) * (m.varEWMA + m.alpha*diff*diff)

			stddev := math.Sqrt(m.varEWMA)
			var zscore float64
			if stddev > 0.001 {
				zscore = (smoothed - m.ewma) / stddev
			}

			score := DriftScore{
				EventID:       event.SeqID,
				OriginalScore: raw,
				SmoothedScore: smoothed,
				Baseline:      m.ewma,
				ZScore:        zscore,
				DriftScore:    drift,
			}

			payloadBytes, _ := json.Marshal(score)

			m.outBus.Publish("telemetry.scored", bus.TelemetryEvent{
				SeqID:     event.SeqID,
				Source:    "phoenix.monitor",
				EventType: "monitor.score",
				Severity:  raw,
				Payload:   payloadBytes,
			})
		}
	}()
}

package monitor

import (
	"encoding/json"
	"math"

	"github.com/fallofpheonix/phoenix_os/bus"
)

// DriftScore represents the anomaly evaluation of an event
type DriftScore struct {
	EventID       int64   `json:"seq_id"`
	OriginalScore float64 `json:"original_score"`
	SmoothedScore float64 `json:"smoothed_score"`
	Baseline      float64 `json:"baseline"`
	ZScore        float64 `json:"z_score"`
	DriftScore    float64 `json:"drift_score"` // Kalman innovation score
	WallTimeUnix  int64   `json:"wall_time_unix"`
}

// KalmanFilter implements a basic one-dimensional Kalman filter for state estimation and anomaly detection.
type KalmanFilter struct {
	q float64 // process noise covariance
	r float64 // measurement noise covariance
	p float64 // estimation error covariance
	k float64 // kalman gain
	x float64 // state estimate
}

// NewKalmanFilter creates a filter instance.
func NewKalmanFilter(q, r, p, initialValue float64) *KalmanFilter {
	return &KalmanFilter{q: q, r: r, p: p, x: initialValue}
}

// Predict performs the time-update step (prediction).
func (kf *KalmanFilter) Predict() (float64, float64) {
	return kf.x, kf.p + kf.q
}

// Update incorporates a new measurement into the filter state.
func (kf *KalmanFilter) Update(measurement float64) float64 {
	kf.p = kf.p + kf.q
	denominator := kf.p + kf.r
	if denominator == 0 {
		return kf.x
	}
	kf.k = kf.p / denominator
	kf.x = kf.x + kf.k*(measurement-kf.x)
	kf.p = (1 - kf.k) * kf.p
	return kf.x
}

// CheckDrift compares a measurement against the predicted state.
func (kf *KalmanFilter) CheckDrift(measurement float64) float64 {
	predX, predP := kf.Predict()
	innovation := measurement - predX
	innovationCov := predP + kf.r
	if innovationCov <= 0 {
		return 0
	}
	return math.Abs(innovation) / math.Sqrt(innovationCov)
}

// MonitorService observes the bus and applies EWMA + Kalman
type MonitorService struct {
	busCh   chan bus.TelemetryEvent
	outBus  *bus.Bus
	kalman  *KalmanFilter
	ewma    float64
	alpha   float64
	varEWMA float64
}

func NewMonitorService(inCh chan bus.TelemetryEvent, outBus *bus.Bus) *MonitorService {
	return &MonitorService{
		busCh:  inCh,
		outBus: outBus,
		kalman: NewKalmanFilter(0.01, 0.1, 1.0, 0.0),
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
		}
	}()
}

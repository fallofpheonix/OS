/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — KALMAN FILTER FOR DRIFT DETECTION
//
// The KalmanFilter provides optimal state estimation for noisy telemetry.
// It's used by the MonitorService to smooth severity scores and detect
// anomalous drift from expected behavior.
//
// WORKFLOW:
//   Monitor.Process(event) → kalman.CheckDrift(raw) → prediction error
//   Monitor.Process(event) → kalman.Update(raw) → smoothed value
//   Monitor.RefineDriftDetection(entropy) → kalman.SetGain(newGain)
//
// ALGORITHM: Standard Kalman filter with:
//   q: process noise covariance (0.01 — low, assumes stable system)
//   r: measurement noise covariance (0.1 — moderate, tolerates some noise)
//   p: estimation error covariance (initialized to 1.0)
//   k: Kalman gain (computed from p, r)
//   x: estimated state (the smoothed value)
//
// COMPLEXITY: O(1) per Update/CheckDrift call.
// No allocations after initialization.
//
// SECURITY NOTE: SetGain() allows external override of the Kalman gain.
// This breaks the mathematical optimality of the filter. Use with caution.
// =========================================================================
package kalman

import "math"

type KalmanFilter struct {
	q float64
	r float64
	p float64
	k float64
	x float64
}

func NewKalmanFilter(q, r, p, initialValue float64) *KalmanFilter {
	return &KalmanFilter{q: q, r: r, p: p, x: initialValue}
}

func (kf *KalmanFilter) Predict() (x, p float64) {
	return kf.x, kf.p + kf.q
}

func (kf *KalmanFilter) Update(measurement float64) float64 {
	kf.p += kf.q
	denominator := kf.p + kf.r
	if denominator == 0 {
		return kf.x
	}
	kf.k = kf.p / denominator
	kf.x += kf.k * (measurement - kf.x)
	kf.p = (1 - kf.k) * kf.p
	return kf.x
}

func (kf *KalmanFilter) CheckDrift(measurement float64) float64 {
	predX, predP := kf.Predict()
	innovation := measurement - predX
	innovationCov := predP + kf.r
	if innovationCov <= 0 {
		return 0
	}
	return math.Abs(innovation) / math.Sqrt(innovationCov)
}

func (kf *KalmanFilter) SetGain(k float64) {
	kf.k = k
}

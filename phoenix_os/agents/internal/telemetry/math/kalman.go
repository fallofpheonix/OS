package kalman

import "math"

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
// It returns the predicted state and the prediction error covariance.
func (kf *KalmanFilter) Predict() (float64, float64) {
	// For this constant-state model: X_pred = X, P_pred = P + Q
	return kf.x, kf.p + kf.q
}

// Update incorporates a new measurement into the filter state.
// It returns the updated state estimate.
func (kf *KalmanFilter) Update(measurement float64) float64 {
	// Prediction step
	kf.p = kf.p + kf.q

	// Measurement update step
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
// It returns a drift score (z-score like) representing how many standard deviations
// the measurement is from the prediction.
func (kf *KalmanFilter) CheckDrift(measurement float64) float64 {
	predX, predP := kf.Predict()
	innovation := measurement - predX
	// Innovation covariance S = P_pred + R
	innovationCov := predP + kf.r
	
	if innovationCov <= 0 {
		return 0
	}
	
	// Return normalized innovation (z-score)
	return math.Abs(innovation) / math.Sqrt(innovationCov)
}

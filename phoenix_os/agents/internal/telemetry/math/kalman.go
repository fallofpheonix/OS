package math

// KalmanFilter implements a basic one-dimensional Kalman filter for state estimation.
type KalmanFilter struct {
	Q float64 // process noise covariance
	R float64 // measurement noise covariance
	P float64 // estimation error covariance
	K float64 // kalman gain
	X float64 // value
}

// NewKalmanFilter creates a filter instance.
func NewKalmanFilter(q, r, p, initialValue float64) *KalmanFilter {
	return &KalmanFilter{Q: q, R: r, P: p, X: initialValue}
}

// Update incorporates a new measurement into the filter state.
func (kf *KalmanFilter) Update(measurement float64) float64 {
	// prediction
	kf.P = kf.P + kf.Q

	// measurement update
	kf.K = kf.P / (kf.P + kf.R)
	kf.X = kf.X + kf.K*(measurement-kf.X)
	kf.P = (1 - kf.K) * kf.P

	return kf.X
}

package physics

// KalmanFilter implements a simple 1D Kalman filter for anomaly detection
type KalmanFilter struct {
	Q float64 // process noise covariance
	R float64 // measurement noise covariance
	P float64 // estimation error covariance
	X float64 // value
	K float64 // kalman gain
}

func NewKalmanFilter(q, r, p, initialValue float64) *KalmanFilter {
	return &KalmanFilter{
		Q: q,
		R: r,
		P: p,
		X: initialValue,
	}
}

func (kf *KalmanFilter) Update(measurement float64) float64 {
	// Prediction update
	// X = X
	// P = P + Q
	kf.P = kf.P + kf.Q

	// Measurement update
	// K = P / (P + R)
	kf.K = kf.P / (kf.P + kf.R)
	// X = X + K * (measurement - X)
	kf.X = kf.X + kf.K*(measurement-kf.X)
	// P = (1 - K) * P
	kf.P = (1 - kf.K) * kf.P

	return kf.X
}

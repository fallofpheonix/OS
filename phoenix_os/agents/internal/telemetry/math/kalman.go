package kalman

// KalmanFilter implements a basic one-dimensional Kalman filter for state estimation.
type KalmanFilter struct {
        q float64 // process noise covariance
        r float64 // measurement noise covariance
        p float64 // estimation error covariance
        k float64 // kalman gain
        x float64 // value
}

// NewKalmanFilter creates a filter instance.
func NewKalmanFilter(q, r, p, initialValue float64) *KalmanFilter {
        return &KalmanFilter{q: q, r: r, p: p, x: initialValue}
}

// Update incorporates a new measurement into the filter state.
func (kf *KalmanFilter) Update(measurement float64) float64 {
        // prediction
        kf.p = kf.p + kf.q

        // measurement update
        denominator := kf.p + kf.r
        if denominator == 0 {
                return kf.x
        }
        kf.k = kf.p / denominator
        kf.x = kf.x + kf.k*(measurement-kf.x)
        kf.p = (1 - kf.k) * kf.p

        return kf.x
}

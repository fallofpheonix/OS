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

func (kf *KalmanFilter) Predict() (float64, float64) {
	return kf.x, kf.p + kf.q
}

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

func (kf *KalmanFilter) CheckDrift(measurement float64) float64 {
	predX, predP := kf.Predict()
	innovation := measurement - predX
	innovationCov := predP + kf.r
	if innovationCov <= 0 {
		return 0
	}
	return math.Abs(innovation) / math.Sqrt(innovationCov)
}

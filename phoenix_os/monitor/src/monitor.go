package main

import (
	"math"
)

type Result struct {
	Entropy      float64 `json:"entropy"`
	KLDivergence float64 `json:"kl_divergence"`
	IsAnomaly    bool    `json:"is_anomaly"`
}

func CalculateEntropy(data []byte) Result {
	if len(data) == 0 {
		return Result{}
	}
	counts := make([]int, 256)
	for _, b := range data {
		counts[b]++
	}
	var entropy float64
	n := float64(len(data))
	for i := 0; i < 256; i++ {
		p := float64(counts[i]) / n
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return Result{
		Entropy:   entropy,
		IsAnomaly: entropy > 7.5,
	}
}

// KalmanFilter for signal smoothing
type KalmanFilter struct {
	Q, R, P, X, K float64
}

func NewKalmanFilter(q, r, p, initialX float64) *KalmanFilter {
	return &KalmanFilter{Q: q, R: r, P: p, X: initialX}
}

func (f *KalmanFilter) Filter(measurement float64) float64 {
	f.P = f.P + f.Q
	f.K = f.P / (f.P + f.R)
	f.X = f.X + f.K*(measurement-f.X)
	f.P = (1 - f.K) * f.P
	return f.X
}

func main() {
	// Boot check
	CalculateEntropy([]byte("PHOENIX"))
}

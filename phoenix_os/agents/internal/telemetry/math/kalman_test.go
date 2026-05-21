package kalman

import "testing"

func TestKalmanFilter(t *testing.T) {
	kf := NewKalmanFilter(0.01, 0.1, 1.0, 0.0)
	
	// Simulate gradual drift
	measurements := []float64{1.0, 1.1, 1.2, 1.3}
	
	last := 0.0
	for _, m := range measurements {
		last = kf.Update(m)
	}
	
	if last < 1.0 {
		t.Errorf("Filter failed to track drift, got %f", last)
	}
}

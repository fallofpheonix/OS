package security

import (
	"testing"
)

func TestPhysicsState(t *testing.T) {
	counts := map[string]float64{
		"process":    50,
		"filesystem": 50,
	}
	
	state := ComputeState(counts, 0.5)
	
	// ln(2) approx 0.693
	if state.Entropy < 0.69 || state.Entropy > 0.70 {
		t.Errorf("Expected entropy around 0.693, got %f", state.Entropy)
	}
	
	if state.Temperature != 50.0 {
		t.Errorf("Expected temperature 50.0, got %f", state.Temperature)
	}
}

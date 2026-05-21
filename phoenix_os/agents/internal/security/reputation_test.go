package security

import "testing"

func TestReputationDeduction(t *testing.T) {
	rm := NewReputationManager()
	rm.Reputation["node-1"] = 1.0

	rm.Deduct("node-1", 0.2)
	score := rm.GetScore("node-1")
	if score != 0.8 {
		t.Errorf("Expected 0.8, got %f", score)
	}
}

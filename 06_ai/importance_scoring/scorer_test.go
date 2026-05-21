package importance_scoring

import "testing"

func TestScorer(t *testing.T) {
	s := NewScorer()
	score := s.Calculate(1.0, 0.5, 0.0)
	expected := 0.5*1.0 + 0.3*0.5 + 0.2*0.0
	if score != expected {
		t.Errorf("Expected %f, got %f", expected, score)
	}
}

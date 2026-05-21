package main

import (
	"math"
	"testing"
)

func TestCalculateSDI(t *testing.T) {
	t.Run("PureState", func(t *testing.T) {
		states := StateVector{1, 1, 1, 1}
		sdi := CalculateSDI(states)
		if sdi != 0 {
			t.Errorf("Expected 0 SDI for pure state, got %f", sdi)
		}
	})

	t.Run("MixedState", func(t *testing.T) {
		states := StateVector{1, -1}
		sdi := CalculateSDI(states)
		expected := - (0.5 * math.Log(0.5) + 0.5 * math.Log(0.5))
		if math.Abs(sdi-expected) > 0.0001 {
			t.Errorf("Expected %f, got %f", expected, sdi)
		}
	})
}

func BenchmarkCalculateSDI(b *testing.B) {
	states := make(StateVector, 100)
	for i := range states {
		if i%2 == 0 {
			states[i] = 1
		} else {
			states[i] = -1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateSDI(states)
	}
}

package main

import "testing"

func TestSDI(t *testing.T) {
	states := StateVector{1, 1, 1, 1}
	sdi := CalculateSDI(states)
	if sdi != 0 {
		t.Errorf("Expected 0 SDI for pure state, got %f", sdi)
	}
}

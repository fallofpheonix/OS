package main

import "testing"

func TestPID(t *testing.T) {
	pid := NewPID(1.0, 0.0, 0.0)
	out := pid.Update(10.0, 5.0, 1.0)
	if out != 5.0 {
		t.Errorf("Expected 5.0, got %f", out)
	}
}

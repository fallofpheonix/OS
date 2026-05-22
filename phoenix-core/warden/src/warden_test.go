package main

import (
	"os"
	"testing"
)

func TestWardenFSM(t *testing.T) {
	// Create a temporary config file for the test
	configContent := `{
		"thresholds": {
			"safe": 0.3,
			"watch": 0.5,
			"suspicious": 0.7,
			"critical": 0.9
		}
	}`
	configFile := "test_warden.json"
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(configFile)

	w, err := NewWarden(configFile)
	if err != nil {
		t.Fatalf("Failed to create Warden with config: %v", err)
	}

	// Initial State
	if w.CurrentState != StateSafe {
		t.Errorf("Expected initial state Safe, got %v", w.CurrentState)
	}

	// Transition to WATCH
	w.EvaluateSDI(0.4)
	if w.CurrentState != StateWatch {
		t.Errorf("Expected Watch state for SDI 0.4, got %v", w.CurrentState)
	}
	if w.Throttling != 0.1 {
		t.Errorf("Expected 0.1 throttling for Watch state, got %f", w.Throttling)
	}

	// Transition to SUSPICIOUS
	w.EvaluateSDI(0.6)
	if w.CurrentState != StateSuspicious {
		t.Errorf("Expected Suspicious state for SDI 0.6, got %v", w.CurrentState)
	}
	if w.Throttling != 0.5 {
		t.Errorf("Expected 0.5 throttling for Suspicious state, got %f", w.Throttling)
	}

	// Transition to CRITICAL
	w.EvaluateSDI(0.8)
	if w.CurrentState != StateCritical {
		t.Errorf("Expected Critical state for SDI 0.8, got %v", w.CurrentState)
	}
	if w.Throttling != 0.9 {
		t.Errorf("Expected 0.9 throttling for Critical state, got %f", w.Throttling)
	}

	// Transition to COMPROMISED
	w.EvaluateSDI(0.95)
	if w.CurrentState != StateCompromised {
		t.Errorf("Expected Compromised state for SDI 0.95, got %v", w.CurrentState)
	}
	if w.Throttling != 1.0 {
		t.Errorf("Expected 1.0 throttling for Compromised state, got %f", w.Throttling)
	}

	// Transition BACK to SAFE (Stability Check)
	w.EvaluateSDI(0.1)
	if w.CurrentState != StateSafe {
		t.Errorf("Expected Safe state for SDI 0.1, got %v", w.CurrentState)
	}
	if w.Throttling != 0.0 {
		t.Errorf("Expected 0.0 throttling for Safe state, got %f", w.Throttling)
	}
}

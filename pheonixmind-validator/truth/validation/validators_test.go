package validation

import (
	"testing"
)

func TestRuntimeValidator(t *testing.T) {
	v := &RuntimeValidator{}
	res := v.Validate("RUNNING")
	if !res.Valid {
		t.Errorf("Expected valid result for RUNNING")
	}
}

func TestDependencyValidator(t *testing.T) {
	v := &DependencyValidator{}
	res := v.Validate("CLEAN")
	if !res.Valid {
		t.Errorf("Expected valid result for CLEAN")
	}
}

func TestObservationValidator(t *testing.T) {
	v := &ObservationValidator{}
	res := v.Validate("0.1")
	if !res.Valid {
		t.Errorf("Expected valid result for low drift")
	}
}

func TestSecurityValidator(t *testing.T) {
	v := &SecurityValidator{}
	res := v.Validate("NONE")
	if !res.Valid {
		t.Errorf("Expected valid result for NONE threat")
	}
}

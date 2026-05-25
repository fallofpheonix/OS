package state

import (
	"testing"
)

func TestFSMTransitions(t *testing.T) {
	registry := NewRegistry()

	// Safe -> Watch
	if err := registry.Set(Watch, "test transition", 1); err != nil {
		t.Fatalf("Failed to transition Safe -> Watch: %v", err)
	}
	if registry.Get() != Watch {
		t.Errorf("Expected state Watch, got %s", registry.Get())
	}

	// Watch -> Alert
	if err := registry.Set(Alert, "test transition", 2); err != nil {
		t.Fatalf("Failed to transition Watch -> Alert: %v", err)
	}

	// Illegal Transition: Alert -> Safe
	if err := registry.Set(Safe, "illegal transition", 3); err == nil {
		t.Error("Expected error for illegal transition Alert -> Safe, got nil")
	}

	violations := registry.Detector.GetViolations()
	if len(violations) != 1 {
		t.Errorf("Expected 1 violation recorded, got %d", len(violations))
	}
	if violations[0].From != Alert || violations[0].To != Safe {
		t.Errorf("Recorded wrong violation: %s -> %s", violations[0].From, violations[0].To)
	}
}

func TestRollback(t *testing.T) {
	registry := NewRegistry()

	registry.Set(Watch, "step 1", 1)
	registry.Set(Alert, "step 2", 2)

	if err := registry.Rollback(); err != nil {
		t.Fatalf("Rollback failed: %v", err)
	}
	if registry.Get() != Watch {
		t.Errorf("Expected state Watch after rollback, got %s", registry.Get())
	}

	if err := registry.Rollback(); err != nil {
		t.Fatalf("Rollback failed: %v", err)
	}
	if registry.Get() != Safe {
		t.Errorf("Expected state Safe after second rollback, got %s", registry.Get())
	}
}

func TestCompatRegistry(t *testing.T) {
	compat := NewCompatRegistry()

	target, err := compat.Lookup("SUSPICIOUS")
	if err != nil {
		t.Fatalf("Lookup failed: %v", err)
	}
	if target != Watch {
		t.Errorf("Expected Watch, got %s", target)
	}

	target, err = compat.Lookup("SAFE")
	if err != nil {
		t.Fatalf("Lookup failed: %v", err)
	}
	if target != Safe {
		t.Errorf("Expected Safe, got %s", target)
	}
}

func TestReplay(t *testing.T) {
	// Replay scenario: same events -> same state
	events := []struct {
		target RuntimeState
		tick   int64
	}{
		{Watch, 1},
		{Alert, 2},
		{Contain, 3},
		{Recovery, 4},
		{Safe, 5},
	}

	registry := NewRegistry()
	for _, e := range events {
		if err := registry.Set(e.target, "replay", e.tick); err != nil {
			t.Fatalf("Replay failed at tick %d: %v", e.tick, err)
		}
	}

	if registry.Get() != Safe {
		t.Errorf("Replay ended in wrong state: %s", registry.Get())
	}
}

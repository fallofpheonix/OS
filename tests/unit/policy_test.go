package unit

import "testing"


func TestIsolationLifecycle(t *testing.T) {
	e := NewIsolationEngine(StateObserve)

	// Valid sequence: Observe -> Watch -> Throttle -> Isolate -> Recover -> Observe
	transitions := []struct {
		next IsolationState
	}{
		{StateWatch},
		{StateThrottle},
		{StateIsolate},
		{StateRecover},
		{StateObserve},
	}

	for _, tr := range transitions {
		if err := e.Transition(tr.next, "ev-1", "dec-1"); err != nil {
			t.Errorf("transition to %s failed: %v", tr.next, err)
		}
	}
}

func TestIllegalContainmentTransition(t *testing.T) {
	e := NewIsolationEngine(StateObserve)

	// Illegal jump: Observe -> Isolate
	if err := e.Transition(StateIsolate, "ev-1", "dec-1"); err == nil {
		t.Error("expected error for illegal containment transition")
	}
}

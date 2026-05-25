package unit

import "testing"

func TestStateTransition(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	
	if err := r.Transition(StateWatch, "test", "ev-1", "dec-1"); err != nil {
		t.Errorf("expected transition to Watch, got %v", err)
	}
	
	if err := r.Transition(StateContain, "bad", "ev-2", "dec-2"); err == nil {
		t.Error("expected illegal transition error")
	}

	if len(r.History) != 2 {
		t.Errorf("expected history length 2, got %d", len(r.History))
	}
}

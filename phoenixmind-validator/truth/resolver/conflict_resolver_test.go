package resolver

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenixmind-validator/truth/evidence"
)

func TestConflictResolver(t *testing.T) {
	// Higher priority wins
	if state := ResolveConflict(evidence.VALIDATED, evidence.WARNING); state != evidence.WARNING {
		t.Errorf("Expected WARNING to win, got %s", state)
	}

	// Lower priority loses
	if state := ResolveConflict(evidence.BLOCKED, evidence.OBSERVED); state != evidence.BLOCKED {
		t.Errorf("Expected BLOCKED to win, got %s", state)
	}
}

func TestMergeTruth(t *testing.T) {
	evidenceSet := []evidence.Evidence{
		{State: evidence.OBSERVED},
		{State: evidence.VALIDATED},
		{State: evidence.WARNING},
	}
	
	if finalState := MergeTruth(evidenceSet); finalState != evidence.WARNING {
		t.Errorf("Expected final merged state to be WARNING, got %s", finalState)
	}
}

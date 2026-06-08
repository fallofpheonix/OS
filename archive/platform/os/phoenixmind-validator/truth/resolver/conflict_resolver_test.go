/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package resolver

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
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

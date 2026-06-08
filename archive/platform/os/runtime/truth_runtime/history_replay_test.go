/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package truth_runtime

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestHistoryReplay(t *testing.T) {
	runtime := NewRuntime()

	history := []evidence.Evidence{
		{EntityID: "replay-test", State: evidence.OBSERVED},
		{EntityID: "replay-test", State: evidence.VALIDATED},
		{EntityID: "replay-test", State: evidence.WARNING},
	}

	for _, ev := range history {
		runtime.IngestEvidence(ev)
	}

	entity, ok := runtime.TruthRegistry.GetEntity("replay-test")
	if !ok {
		t.Fatalf("Entity 'replay-test' not found in registry")
	}

	if entity.CurrentState != evidence.WARNING {
		t.Errorf("Expected final state to be WARNING after replay, got %s", entity.CurrentState)
	}

	// Test the history validation function
	if !runtime.ValidateHistory(entity.History) {
		t.Errorf("Expected history to be considered valid")
	}
}

func TestConflictChain(t *testing.T) {
	runtime := NewRuntime()
	entityID := "conflict-chain"

	// Chain: OBSERVED -> VALIDATED -> WARNING -> ESCALATED -> BLOCKED
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.OBSERVED})
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.VALIDATED})
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.WARNING})
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.ESCALATED})
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.BLOCKED})
	
	entity, _ := runtime.TruthRegistry.GetEntity(entityID)
	if entity.CurrentState != evidence.BLOCKED {
		t.Errorf("Expected final state to be BLOCKED, got %s", entity.CurrentState)
	}

	// Add a lower-priority event; state should not change
	runtime.IngestEvidence(evidence.Evidence{EntityID: entityID, State: evidence.VALIDATED})
	entity, _ = runtime.TruthRegistry.GetEntity(entityID)
	if entity.CurrentState != evidence.BLOCKED {
		t.Errorf("Expected state to remain BLOCKED, got %s", entity.CurrentState)
	}
}

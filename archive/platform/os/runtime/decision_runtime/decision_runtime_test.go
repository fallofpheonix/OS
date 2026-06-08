/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package decision_runtime

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestMakeDecision(t *testing.T) {
	entityID := "test-entity-decision"

	// Test BLOCKED state
	decision := MakeDecision(entityID, evidence.BLOCKED, 0.9)
	if decision.Action != "CONTAIN" {
		t.Errorf("Expected action CONTAIN for BLOCKED, got %s", decision.Action)
	}

	// Test ESCALATED state
	decision = MakeDecision(entityID, evidence.ESCALATED, 0.7)
	if decision.Action != "ESCALATE" {
		t.Errorf("Expected action ESCALATE for ESCALATED, got %s", decision.Action)
	}

	// Test WARNING state
	decision = MakeDecision(entityID, evidence.WARNING, 0.5)
	if decision.Action != "LOG_AND_MONITOR" {
		t.Errorf("Expected action LOG_AND_MONITOR for WARNING, got %s", decision.Action)
	}

	// Test VALIDATED state
	decision = MakeDecision(entityID, evidence.VALIDATED, 0.99)
	if decision.Action != "ALLOW" {
		t.Errorf("Expected action ALLOW for VALIDATED, got %s", decision.Action)
	}

	// Test UNKNOWN state
	decision = MakeDecision(entityID, evidence.UNKNOWN, 0.1)
	if decision.Action != "INVESTIGATE" {
		t.Errorf("Expected action INVESTIGATE for UNKNOWN, got %s", decision.Action)
	}
}

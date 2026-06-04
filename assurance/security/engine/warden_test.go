/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package engine

import (
	"testing"
)

func TestWarden_Transition(t *testing.T) {
	w := NewWarden()

	// Valid ladder: SAFE -> WATCH
	if err := w.Transition(StateWatch); err != nil {
		t.Errorf("Failed valid transition SAFE -> WATCH: %v", err)
	}

	// Invalid skip: WATCH -> COMPROMISED (must pass through SUSPICIOUS and CRITICAL)
	if err := w.Transition(StateCompromised); err == nil {
		t.Error("Allowed invalid skip WATCH -> COMPROMISED")
	}

	// Valid ladder: WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED
	steps := []SystemState{StateSuspicious, StateCritical, StateCompromised}
	for _, state := range steps {
		if err := w.Transition(state); err != nil {
			t.Errorf("Failed valid transition to %s: %v", state, err)
		}
	}

	// Lockout test
	w.Lock()
	if err := w.Transition(StateCritical); err == nil {
		t.Error("Allowed transition after FSM lock")
	}
}

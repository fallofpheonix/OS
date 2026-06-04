/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import "testing"

func TestAlertToSafeReject(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	// We want to simulate a situation where the previous state in history does NOT match the rules.
	// Current: ALERT. History: [INIT: SAFE->SAFE, WATCH->ALERT]
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	r.Transition(StateAlert, "spike", "ev-2", "dec-2")

	// Manually inject a history record that would lead to an illegal rollback
	// If current is ALERT, it must rollback to WATCH.
	// To test rejection, we need current to be ALERT, but previous to be SAFE.
	r.History = append(r.History, StateRecord{
		Previous: StateSafe, // Invalid for ALERT
		Current:  StateAlert,
	})
	r.CurrentState = StateAlert

	if err := r.Rollback("illegal-path", "ev-3", "dec-3"); err == nil {
		t.Error("expected error for illegal rollback path ALERT->SAFE, but got nil")
	}
}

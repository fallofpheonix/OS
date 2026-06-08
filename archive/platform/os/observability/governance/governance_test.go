/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package governance

import "testing"

func TestReadinessScore(t *testing.T) {
	// READY case
	readyInput := ReadinessInput{
		ObservationCycles: 15,
		RuntimeHealth:     0.95,
		DriftValue:        0.05,
		CriticalAlerts:    0,
	}
	if score := CalculateReadinessScore(readyInput); score != "READY" {
		t.Errorf("Expected READY, got %s", score)
	}

	// REVIEW case
	reviewInput := ReadinessInput{
		ObservationCycles: 8,
		RuntimeHealth:     0.90,
		DriftValue:        0.2,
		CriticalAlerts:    0,
	}
	if score := CalculateReadinessScore(reviewInput); score != "REVIEW" {
		t.Errorf("Expected REVIEW, got %s", score)
	}

	// BLOCKED case
	blockedInput := ReadinessInput{
		CriticalAlerts: 1,
	}
	if score := CalculateReadinessScore(blockedInput); score != "BLOCKED" {
		t.Errorf("Expected BLOCKED, got %s", score)
	}
}

func TestTransitionGuard(t *testing.T) {
	// Forbidden case
	if ok, _ := GuardTransition("F1", "Training"); ok {
		t.Errorf("Expected transition F1->Training to be forbidden")
	}

	// Allowed case
	if ok, _ := GuardTransition("F1", "F1-Hardening"); !ok {
		t.Errorf("Expected transition F1->F1-Hardening to be allowed")
	}
}

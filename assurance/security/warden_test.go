/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package warden

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

func TestWardenStateLadder(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	// Register Evidence Weight Invariant mapping to TLA+ spec
	w.RegisterInvariant(&EvidenceWeightInvariant{
		StateThresholds: map[SystemState]float64{
			StateSuspicious:  0.50,
			StateCritical:    0.80,
			StateCompromised: 0.95,
		},
	})

	// Valid sequence: SAFE -> WATCH -> SUSPICIOUS
	req1 := AuthorityEscalationRequest{TargetState: StateWatch, ActuationClass: ClassLog, EvidenceWeight: 1.0}
	if !w.ActuateRequest(req1, 1, 1) {
		t.Errorf("Failed valid transition SAFE -> WATCH")
	}

	req2 := AuthorityEscalationRequest{TargetState: StateSuspicious, ActuationClass: ClassLog, EvidenceWeight: 1.0}
	if !w.ActuateRequest(req2, 2, 2) {
		t.Errorf("Failed valid transition WATCH -> SUSPICIOUS")
	}

	// Invalid Evidence Weight: SUSPICIOUS -> CRITICAL with low weight
	reqLow := AuthorityEscalationRequest{TargetState: StateCritical, ActuationClass: ClassLog, EvidenceWeight: 0.70}
	if w.ActuateRequest(reqLow, 3, 3) {
		t.Errorf("Allowed transition SUSPICIOUS -> CRITICAL with insufficient evidence")
	}

	// Note: Since Invariant violation triggers emergencyHalt, state becomes COMPROMISED
	// Let's reset the Warden for further tests.
	w = NewWarden(b)
	w.RegisterInvariant(&EvidenceWeightInvariant{
		StateThresholds: map[SystemState]float64{
			StateSuspicious:  0.50,
			StateCritical:    0.80,
			StateCompromised: 0.95,
		},
	})
	w.ActuateRequest(req1, 1, 4)
	w.ActuateRequest(req2, 2, 5)

	// Invalid sequence: SUSPICIOUS -> SAFE (illegal jump, must go back through WATCH)
	req3 := AuthorityEscalationRequest{TargetState: StateSafe, ActuationClass: ClassLog, EvidenceWeight: 1.0}
	if w.ActuateRequest(req3, 3, 6) {
		t.Errorf("Allowed illegal transition SUSPICIOUS -> SAFE")
	}

	// Valid recovery: SUSPICIOUS -> WATCH -> SAFE
	req4 := AuthorityEscalationRequest{TargetState: StateWatch, ActuationClass: ClassLog, EvidenceWeight: 1.0}
	if !w.ActuateRequest(req4, 4, 7) {
		t.Errorf("Failed valid recovery transition SUSPICIOUS -> WATCH")
	}

	req5 := AuthorityEscalationRequest{TargetState: StateSafe, ActuationClass: ClassLog, EvidenceWeight: 1.0}
	if !w.ActuateRequest(req5, 5, 8) {
		t.Errorf("Failed valid recovery transition WATCH -> SAFE")
	}
}

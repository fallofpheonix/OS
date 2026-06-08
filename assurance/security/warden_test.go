/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
package security

import (
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"testing"
)

func TestWardenStateLadder(t *testing.T) {
	b := bus.NewBus()
	// Dummy ledger to satisfy NewWarden
	w := NewWarden(b, &ledger.Ledger{})
	w.shadowMode.Store(false)

	// Register Evidence Weight Invariant mapping to TLA+ spec
	w.RegisterInvariant(&EvidenceWeightInvariant{
		StateThresholds: map[SystemState]phxmath.FixedPoint{
			StateSuspicious:  phxmath.NewFixedPointRaw(500000),
			StateCritical:    phxmath.NewFixedPointRaw(800000),
			StateCompromised: phxmath.NewFixedPointRaw(950000),
		},
	})

	h := computePolicyHash(w.policies.Load())

	// Valid sequence: SAFE -> WATCH -> SUSPICIOUS
	req1 := AuthorityEscalationRequest{TargetState: StateWatch, ActuationClass: ClassLog, EvidenceWeight: phxmath.NewFixedPoint(1), PolicyHash: h}
	if ok, _ := w.ActuateRequest(req1, 1, 1); !ok {
		t.Errorf("Failed valid transition SAFE -> WATCH")
	}

	req2 := AuthorityEscalationRequest{TargetState: StateSuspicious, ActuationClass: ClassLog, EvidenceWeight: phxmath.NewFixedPoint(1), PolicyHash: h}
	if ok, _ := w.ActuateRequest(req2, 2, 2); !ok {
		t.Errorf("Failed valid transition WATCH -> SUSPICIOUS")
	}

	// Invalid Evidence Weight: SUSPICIOUS -> CRITICAL with low weight
	reqLow := AuthorityEscalationRequest{TargetState: StateCritical, ActuationClass: ClassLog, EvidenceWeight: phxmath.NewFixedPointRaw(700000), PolicyHash: h}
	if ok, _ := w.ActuateRequest(reqLow, 3, 3); ok {
		t.Errorf("Allowed transition SUSPICIOUS -> CRITICAL with insufficient evidence")
	}
}

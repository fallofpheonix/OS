package validation

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
)

// TestStatePurity verifies that snapshot normalization clears transient time fields.
func TestStatePurity(t *testing.T) {
	proc := containment.NewProcessAudit()
	proc.LogAction(containment.ProcessAction{PID: 10, Action: containment.ActionMonitor})

	// Check time is not zero initially
	if proc.History[0].Timestamp.IsZero() {
		t.Error("expected initial timestamp to be non-zero")
	}

	// Normalize state
	proc.Normalize()

	// Check time is zero (pure epoch)
	if proc.History[0].Timestamp != time.Unix(0, 0) {
		t.Errorf("state normalization failed to clear transient timestamps: got %v", proc.History[0].Timestamp)
	}
}

// TestOutputDeterminism verifies output hash correctness across repeated runs.
func TestOutputDeterminism(t *testing.T) {
	run := func() []byte {
		proc := containment.NewProcessAudit()
		net := network.NewNetworkAudit()
		f := file.NewFileAudit()
		orch := &rollback.Orchestrator{ProcessAudit: proc, NetworkAudit: net, FileAudit: f}

		proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor})
		net.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionMonitor})

		data, _ := orch.CreateGlobalSnapshot(1)
		return data
	}

	h1 := run()
	for i := 0; i < 10; i++ {
		h2 := run()
		if string(h1) != string(h2) {
			t.Fatalf("output determinism mismatch at iteration %d", i)
		}
	}
}

// TestTransitionInvariant asserts state changes follow validated rules.
func TestTransitionInvariant(t *testing.T) {
	engine := containment.NewIsolationEngine(containment.StateObserve)

	// Valid: Observe -> Watch
	if err := engine.Transition(containment.StateWatch, "ev-1", "dec-1"); err != nil {
		t.Fatalf("failed valid transition Observe -> Watch: %v", err)
	}

	// Invalid transition: Watch -> Recover directly
	err := engine.Transition(containment.StateRecover, "ev-2", "dec-2")
	if err == nil {
		t.Error("expected invalid transition to be rejected, but it succeeded")
	}
}

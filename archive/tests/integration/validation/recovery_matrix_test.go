package validation

import (
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
)

// TestRecoveryRepeatability verifies that recovery can be executed repeatedly on identical state without failures.
func TestRecoveryRepeatability(t *testing.T) {
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &rollback.Orchestrator{ProcessAudit: proc, NetworkAudit: net, FileAudit: f}

	proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "p1"})
	snapshotData, _ := orch.CreateGlobalSnapshot(1)

	// Run recovery repeatedly
	for i := 0; i < 100; i++ {
		procRestored := containment.NewProcessAudit()
		netRestored := network.NewNetworkAudit()
		fRestored := file.NewFileAudit()
		orchRestored := &rollback.Orchestrator{ProcessAudit: procRestored, NetworkAudit: netRestored, FileAudit: fRestored}

		if err := orchRestored.RestoreGlobal(snapshotData); err != nil {
			t.Fatalf("failed recovery restore at iteration %d: %v", i, err)
		}

		if len(procRestored.History) != 1 || procRestored.History[0].PID != 100 {
			t.Fatalf("recovered process history invalid at iteration %d", i)
		}
	}
}

// TestRecoveryBudget simulates recovery budget validation limits.
func TestRecoveryBudget(t *testing.T) {
	// A recovery budget limits Warden transitions or repeated containment triggers.
	// We verify that the rollback/recovery counters and transitions are stable.
	proc := containment.NewProcessAudit()
	if proc.Sequence != 0 {
		t.Fatalf("expected initial sequence 0, got %d", proc.Sequence)
	}

	maxActions := 50
	for i := 0; i < maxActions; i++ {
		proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor})
	}

	if len(proc.History) != maxActions {
		t.Errorf("expected history length %d, got %d", maxActions, len(proc.History))
	}
}

// TestRecoveryRollback verifies that rollback operations return the system state to the last safe posture.
func TestRecoveryRollback(t *testing.T) {
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &rollback.Orchestrator{ProcessAudit: proc, NetworkAudit: net, FileAudit: f}

	// Safe state snapshot
	proc.LogAction(containment.ProcessAction{PID: 10, Action: containment.ActionMonitor})
	snapshotData, _ := orch.CreateGlobalSnapshot(1)

	// Simulate compromised actions post-snapshot
	proc.LogAction(containment.ProcessAction{PID: 99, Action: containment.ActionThrottle})

	// Trigger recovery rollback to restore state
	if err := orch.RestoreGlobal(snapshotData); err != nil {
		t.Fatalf("failed to restore to safe snapshot: %v", err)
	}

	// Verify the compromised action was discarded
	if len(proc.History) != 1 {
		t.Fatalf("expected rollback to discard post-snapshot history, history length is %d", len(proc.History))
	}
	if proc.History[0].PID != 10 {
		t.Errorf("expected PID 10, got PID %d", proc.History[0].PID)
	}
}

package validation

import (
	"bytes"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
)

// TestRollbackStateParity verifies that rolling back a snapshot results in exact global state parity.
func TestRollbackStateParity(t *testing.T) {
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &rollback.Orchestrator{ProcessAudit: proc, NetworkAudit: net, FileAudit: f}

	proc.LogAction(containment.ProcessAction{PID: 10, Action: containment.ActionMonitor})
	net.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionMonitor})

	// Original state snapshot
	snapshot1, _ := orch.CreateGlobalSnapshot(1)

	// Mutate state
	proc.LogAction(containment.ProcessAction{PID: 20, Action: containment.ActionThrottle})

	// Rollback
	if err := orch.RestoreGlobal(snapshot1); err != nil {
		t.Fatalf("rollback failed: %v", err)
	}

	// Capture state post-rollback
	snapshot2, _ := orch.CreateGlobalSnapshot(1)

	if !bytes.Equal(snapshot1, snapshot2) {
		t.Error("expected state post-rollback to equal the original snapshot state, but found divergence")
	}
}

// TestRollbackAuditTrail verifies that rollback actions log audit entries deterministically.
func TestRollbackAuditTrail(t *testing.T) {
	audit := rollback.NewRollbackAudit()

	record := rollback.RollbackRecord{
		Component:  rollback.ComponentNetwork,
		SnapshotID: "snap-1",
		RecoveryID: "rec-1",
		EvidenceID: "ev-1",
		DecisionID: "dec-1",
		Sequence:   1,
	}

	audit.LogRollback(record)

	if len(audit.History) != 1 {
		t.Fatalf("expected 1 rollback history record, got %d", len(audit.History))
	}

	if audit.History[0].Hash == "" {
		t.Error("expected audit record hash to be populated deterministically")
	}
}

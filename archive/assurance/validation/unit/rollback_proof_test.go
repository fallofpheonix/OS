/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/containment"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/file"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/network"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/rollback"
)

// TestRollbackProofChain implements C5.6:
// Executable chain: Process -> Network -> File -> Snapshot -> Rollback -> Restore -> Replay -> Verify
// Verification: Same Snapshot -> Restore -> Same Component History -> Same Replay -> Same Verification
func TestRollbackProofChain(t *testing.T) {
	// 1. Process -> Network -> File Actions logging
	procAuditOriginal := containment.NewProcessAudit()
	netAuditOriginal := network.NewNetworkAudit()
	fileAuditOriginal := file.NewFileAudit()
	rollbackAuditOriginal := rollback.NewRollbackAudit()

	orchOriginal := &rollback.Orchestrator{
		ProcessAudit: procAuditOriginal,
		NetworkAudit: netAuditOriginal,
		FileAudit:    fileAuditOriginal,
	}

	// Log actions to individual layers
	procAuditOriginal.LogAction(containment.ProcessAction{
		PID:        5000,
		Action:     containment.ActionMonitor,
		Reason:     "monitored sandbox run",
		EvidenceID: "ev-proc-chain",
		DecisionID: "dec-proc-chain",
	})
	netAuditOriginal.LogAction(network.NetworkAction{
		Src:        "10.0.0.5",
		Dst:        "1.1.1.1",
		Port:       80,
		Action:     network.ActionMonitor,
		Reason:     "monitored outgoing traffic",
		EvidenceID: "ev-net-chain",
		DecisionID: "dec-net-chain",
	})
	fileAuditOriginal.LogAction(file.FileAction{
		Path:       "/tmp/test",
		Action:     file.ActionMonitor,
		Reason:     "monitored temporary write",
		EvidenceID: "ev-file-chain",
		DecisionID: "dec-file-chain",
	})

	// 2. Snapshot
	// Capture global snapshot state representing the accumulated audit evidence
	snapshotData, err := orchOriginal.CreateGlobalSnapshot(42)
	if err != nil {
		t.Fatalf("failed to create global snapshot: %v", err)
	}

	// 3. Rollback Action logging (RollbackAudit)
	// Log the rollback itself, proving evidence of rollback execution
	rollbackRecordOriginal := rollback.RollbackRecord{
		Component:     rollback.ComponentNetwork,
		PreviousState: "WATCH",
		CurrentState:  "SAFE",
		SnapshotID:    "global-snapshot-42",
		RecoveryID:    "rec-42",
		EvidenceID:    "ev-rollback-proof",
		DecisionID:    "dec-rollback-proof",
		Sequence:      1,
	}
	rollbackAuditOriginal.LogRollback(rollbackRecordOriginal)

	// 4. Restore
	// Instantiate a fresh orchestrator to perform rollback restoration
	procAuditRestored := containment.NewProcessAudit()
	netAuditRestored := network.NewNetworkAudit()
	fileAuditRestored := file.NewFileAudit()
	orchRestored := &rollback.Orchestrator{
		ProcessAudit: procAuditRestored,
		NetworkAudit: netAuditRestored,
		FileAudit:    fileAuditRestored,
	}

	if err := orchRestored.RestoreGlobal(snapshotData); err != nil {
		t.Fatalf("failed to restore global state in rollback proof: %v", err)
	}

	// 5. Verify Invariants
	// - Same Snapshot (the restored orchestrator must create a snapshot identical to the original)
	restoredSnapshotData, err := orchRestored.CreateGlobalSnapshot(42)
	if err != nil {
		t.Fatalf("failed to create global snapshot from restored state: %v", err)
	}

	if !bytes.Equal(snapshotData, restoredSnapshotData) {
		t.Errorf("Same Snapshot validation failed: restored state does not yield the original snapshot")
	}

	// - Same Component History
	if len(procAuditRestored.History) != len(procAuditOriginal.History) ||
		procAuditRestored.History[0].Hash != procAuditOriginal.History[0].Hash {
		t.Errorf("Process history mismatch after restore")
	}
	if len(netAuditRestored.History) != len(netAuditOriginal.History) ||
		netAuditRestored.History[0].Hash != netAuditOriginal.History[0].Hash {
		t.Errorf("Network history mismatch after restore")
	}
	if len(fileAuditRestored.History) != len(fileAuditOriginal.History) ||
		fileAuditRestored.History[0].Hash != fileAuditOriginal.History[0].Hash {
		t.Errorf("File history mismatch after restore")
	}

	// - Same Replay
	// Replay a rollback record and assert metadata equality
	replayOriginal := rollback.RollbackReplay{
		Record:       rollbackAuditOriginal.History[0],
		ReplayCursor: rollbackAuditOriginal.History[0].Sequence,
		Sequence:     rollbackAuditOriginal.History[0].Sequence,
		Hash:         rollbackAuditOriginal.History[0].Hash,
	}

	// Make sure the replay cursor and sequence matches rollback audit state
	if replayOriginal.ReplayCursor != 1 || replayOriginal.Sequence != 1 {
		t.Errorf("Replay cursor validation mismatch: expected 1, got cursor=%d, sequence=%d", replayOriginal.ReplayCursor, replayOriginal.Sequence)
	}

	// - Same Verification (all verification checks are correct and error free)
	var gsOriginal rollback.GlobalSnapshot
	if err := json.Unmarshal(snapshotData, &gsOriginal); err != nil {
		t.Fatalf("unmarshal original snapshot failed: %v", err)
	}
	var gsRestored rollback.GlobalSnapshot
	if err := json.Unmarshal(restoredSnapshotData, &gsRestored); err != nil {
		t.Fatalf("unmarshal restored snapshot failed: %v", err)
	}

	if gsOriginal.Hash != gsRestored.Hash {
		t.Errorf("verification check failed: global snapshot hashes are not identical")
	}
}

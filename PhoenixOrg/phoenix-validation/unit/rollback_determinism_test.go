package unit

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
)

// TestRollbackDeterminism validates C5.5:
// 1. Snapshot hash equality
// 2. Restore equality
// 3. Sequence equality
// 4. Replay equality
// 5. Provider equality
func TestRollbackDeterminism(t *testing.T) {
	// Setup the initial state stream
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &Orchestrator{
		ProcessAudit: proc,
		NetworkAudit: net,
		FileAudit:    f,
	}

	// Log actions to establish state
	proc.LogAction(containment.ProcessAction{
		PID:        42,
		Action:     containment.ActionMonitor,
		Reason:     "process initialization",
		EvidenceID: "ev-proc-1",
		DecisionID: "dec-proc-1",
	})
	proc.LogAction(containment.ProcessAction{
		PID:        42,
		Action:     containment.ActionThrottle,
		Reason:     "high CPU usage limit",
		EvidenceID: "ev-proc-2",
		DecisionID: "dec-proc-2",
	})

	net.LogAction(network.NetworkAction{
		Src:        "192.168.1.50",
		Dst:        "8.8.8.8",
		Port:       53,
		Action:     network.ActionMonitor,
		Reason:     "DNS query monitor",
		EvidenceID: "ev-net-1",
		DecisionID: "dec-net-1",
	})
	net.LogAction(network.NetworkAction{
		Src:        "192.168.1.50",
		Dst:        "10.0.0.1",
		Port:       443,
		Action:     network.ActionThrottle,
		Reason:     "bandwidth throttle",
		EvidenceID: "ev-net-2",
		DecisionID: "dec-net-2",
	})

	f.LogAction(file.FileAction{
		Path:       "/etc/shadow",
		Action:     file.ActionMonitor,
		Reason:     "sensitive file access",
		EvidenceID: "ev-file-1",
		DecisionID: "dec-file-1",
	})
	f.LogAction(file.FileAction{
		Path:       "/var/log/auth.log",
		Action:     file.ActionFreeze,
		Reason:     "unauthorized audit edit",
		EvidenceID: "ev-file-2",
		DecisionID: "dec-file-2",
	})

	// 1. Snapshot Hash Equality
	// Creating the snapshot twice from the same orchestrator state must yield exact byte and hash equality.
	snapBytes1, err := orch.CreateGlobalSnapshot(100)
	if err != nil {
		t.Fatalf("failed to create global snapshot 1: %v", err)
	}

	snapBytes2, err := orch.CreateGlobalSnapshot(100)
	if err != nil {
		t.Fatalf("failed to create global snapshot 2: %v", err)
	}

	if !bytes.Equal(snapBytes1, snapBytes2) {
		t.Errorf("snapshot hash/payload equality failed: multiple snapshot invocations must yield deterministic byte-for-byte identical output")
	}

	// Verify JSON matches and internal hashes are valid
	var gs1, gs2 GlobalSnapshot
	if err := json.Unmarshal(snapBytes1, &gs1); err != nil {
		t.Fatalf("failed to unmarshal snapshot 1: %v", err)
	}
	if err := json.Unmarshal(snapBytes2, &gs2); err != nil {
		t.Fatalf("failed to unmarshal snapshot 2: %v", err)
	}

	if gs1.Hash != gs2.Hash {
		t.Errorf("expected global snapshot hashes to be equal, got gs1=%s, gs2=%s", gs1.Hash, gs2.Hash)
	}

	// 2. Restore Equality & Provider Equality
	// Restore into a fresh orchestrator and assert exact structural equality of the history
	procRestored := containment.NewProcessAudit()
	netRestored := network.NewNetworkAudit()
	fRestored := file.NewFileAudit()
	orchRestored := &Orchestrator{
		ProcessAudit: procRestored,
		NetworkAudit: netRestored,
		FileAudit:    fRestored,
	}

	if err := orchRestored.RestoreGlobal(snapBytes1); err != nil {
		t.Fatalf("failed to restore global state: %v", err)
	}

	// Verify deep equality of the process provider state
	if len(proc.History) != len(procRestored.History) {
		t.Errorf("process history length mismatch: expected %d, got %d", len(proc.History), len(procRestored.History))
	} else {
		for i := range proc.History {
			p1, p2 := proc.History[i], procRestored.History[i]
			if p1.PID != p2.PID || p1.Action != p2.Action || p1.Reason != p2.Reason ||
				p1.EvidenceID != p2.EvidenceID || p1.DecisionID != p2.DecisionID ||
				p1.Sequence != p2.Sequence || p1.Hash != p2.Hash {
				t.Errorf("process history item %d mismatch: %+v vs %+v", i, p1, p2)
			}
		}
	}

	// Verify deep equality of the network provider state
	if len(net.History) != len(netRestored.History) {
		t.Errorf("network history length mismatch: expected %d, got %d", len(net.History), len(netRestored.History))
	} else {
		for i := range net.History {
			n1, n2 := net.History[i], netRestored.History[i]
			if n1.Src != n2.Src || n1.Dst != n2.Dst || n1.Port != n2.Port ||
				n1.Action != n2.Action || n1.Reason != n2.Reason ||
				n1.EvidenceID != n2.EvidenceID || n1.DecisionID != n2.DecisionID ||
				n1.Sequence != n2.Sequence || n1.Hash != n2.Hash {
				t.Errorf("network history item %d mismatch: %+v vs %+v", i, n1, n2)
			}
		}
	}

	// Verify deep equality of the file provider state
	if len(f.History) != len(fRestored.History) {
		t.Errorf("file history length mismatch: expected %d, got %d", len(f.History), len(fRestored.History))
	} else {
		for i := range f.History {
			f1, f2 := f.History[i], fRestored.History[i]
			if f1.Path != f2.Path || f1.Action != f2.Action || f1.Reason != f2.Reason ||
				f1.EvidenceID != f2.EvidenceID || f1.DecisionID != f2.DecisionID ||
				f1.Sequence != f2.Sequence || f1.Hash != f2.Hash {
				t.Errorf("file history item %d mismatch: %+v vs %+v", i, f1, f2)
			}
		}
	}

	// 3. Sequence Equality
	if gs1.Sequence != 100 {
		t.Errorf("global snapshot sequence mismatch: expected 100, got %d", gs1.Sequence)
	}
	if procRestored.Sequence != proc.Sequence {
		t.Errorf("process audit sequence mismatch: expected %d, got %d", proc.Sequence, procRestored.Sequence)
	}
	if netRestored.Sequence != net.Sequence {
		t.Errorf("network audit sequence mismatch: expected %d, got %d", net.Sequence, netRestored.Sequence)
	}
	if fRestored.Sequence != f.Sequence {
		t.Errorf("file audit sequence mismatch: expected %d, got %d", f.Sequence, fRestored.Sequence)
	}

	// 4. Replay Equality
	// Replaying actions from both restored and original providers must produce identical replay objects.
	for i := range proc.History {
		actionOriginal := proc.History[i]
		actionRestored := procRestored.History[i]

		// Check process replay properties deterministically
		if actionOriginal.Hash != actionRestored.Hash {
			t.Errorf("process replay mismatch: original hash %q, restored hash %q", actionOriginal.Hash, actionRestored.Hash)
		}
	}

	for i := range net.History {
		actionOriginal := net.History[i]
		actionRestored := netRestored.History[i]

		// Simulate NetworkReplay construction and assert equality
		replayOriginal := network.NetworkReplay{
			ConnectionID:  "conn-test",
			Action:        actionOriginal,
			PreviousState: "NONE",
			CurrentState:  string(actionOriginal.Action),
			ReplayCursor:  actionOriginal.Sequence,
			Sequence:      actionOriginal.Sequence,
			Hash:          actionOriginal.Hash,
		}

		replayRestored := network.NetworkReplay{
			ConnectionID:  "conn-test",
			Action:        actionRestored,
			PreviousState: "NONE",
			CurrentState:  string(actionRestored.Action),
			ReplayCursor:  actionRestored.Sequence,
			Sequence:      actionRestored.Sequence,
			Hash:          actionRestored.Hash,
		}

		if replayOriginal.ReplayCursor != replayRestored.ReplayCursor ||
			replayOriginal.Hash != replayRestored.Hash ||
			replayOriginal.Action.Hash != replayRestored.Action.Hash {
			t.Errorf("network replay equality failed at step %d", i)
		}
	}

	for i := range f.History {
		actionOriginal := f.History[i]
		actionRestored := fRestored.History[i]

		// Simulate FileReplay construction and assert equality
		replayOriginal := file.FileReplay{
			Path:          actionOriginal.Path,
			Action:        actionOriginal,
			PreviousState: "NONE",
			CurrentState:  string(actionOriginal.Action),
			ReplayCursor:  actionOriginal.Sequence,
			Sequence:      actionOriginal.Sequence,
			Hash:          actionOriginal.Hash,
		}

		replayRestored := file.FileReplay{
			Path:          actionRestored.Path,
			Action:        actionRestored,
			PreviousState: "NONE",
			CurrentState:  string(actionRestored.Action),
			ReplayCursor:  actionRestored.Sequence,
			Sequence:      actionRestored.Sequence,
			Hash:          actionRestored.Hash,
		}

		if replayOriginal.ReplayCursor != replayRestored.ReplayCursor ||
			replayOriginal.Hash != replayRestored.Hash ||
			replayOriginal.Action.Hash != replayRestored.Action.Hash {
			t.Errorf("file replay equality failed at step %d", i)
		}
	}
}

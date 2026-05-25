package validation

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
)

// TestReplayDeterminism verifies that replaying a stream of actions yields identical final state hashes.
func TestReplayDeterminism(t *testing.T) {
	runStream := func() string {
		proc := containment.NewProcessAudit()
		net := network.NewNetworkAudit()
		f := file.NewFileAudit()
		orch := &rollback.Orchestrator{
			ProcessAudit: proc,
			NetworkAudit: net,
			FileAudit:    f,
		}

		proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "p1"})
		net.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionMonitor, Reason: "n1"})
		f.LogAction(file.FileAction{Path: "/tmp/foo", Action: file.ActionMonitor, Reason: "f1"})

		snapshot, err := orch.CreateGlobalSnapshot(1)
		if err != nil {
			t.Fatalf("failed to create global snapshot: %v", err)
		}
		
		var gs rollback.GlobalSnapshot
		if err := json.Unmarshal(snapshot, &gs); err != nil {
			t.Fatalf("failed to unmarshal global snapshot: %v", err)
		}
		return gs.Hash
	}

	h1 := runStream()
	for i := 0; i < 100; i++ {
		h2 := runStream()
		if h1 != h2 {
			t.Fatalf("replay determinism failure at loop %d: hash drift detected (%s != %s)", i, h1, h2)
		}
	}
}

// TestReplayChecksum verifies checksum calculations reject invalid data.
func TestReplayChecksum(t *testing.T) {
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &rollback.Orchestrator{
		ProcessAudit: proc,
		NetworkAudit: net,
		FileAudit:    f,
	}

	proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "p1"})
	snapshot, _ := orch.CreateGlobalSnapshot(1)

	// Corrupt snapshot
	var gs rollback.GlobalSnapshot
	json.Unmarshal(snapshot, &gs)
	gs.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte("corrupted")))
	badSnapshot, _ := json.Marshal(gs)

	err := orch.RestoreGlobal(badSnapshot)
	if err == nil {
		t.Error("expected error for corrupted global snapshot hash, got nil")
	}
}

// TestReplayDivergence verifies that different event sequences yield different state hashes (divergence validation).
func TestReplayDivergence(t *testing.T) {
	proc1 := containment.NewProcessAudit()
	net1 := network.NewNetworkAudit()
	f1 := file.NewFileAudit()
	orch1 := &rollback.Orchestrator{ProcessAudit: proc1, NetworkAudit: net1, FileAudit: f1}

	proc1.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "p1"})
	s1, _ := orch1.CreateGlobalSnapshot(1)

	proc2 := containment.NewProcessAudit()
	net2 := network.NewNetworkAudit()
	f2 := file.NewFileAudit()
	orch2 := &rollback.Orchestrator{ProcessAudit: proc2, NetworkAudit: net2, FileAudit: f2}

	proc2.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "p1"})
	// Add an extra action to diverge the state
	proc2.LogAction(containment.ProcessAction{PID: 200, Action: containment.ActionThrottle, Reason: "p2"})
	s2, _ := orch2.CreateGlobalSnapshot(1)

	var gs1, gs2 rollback.GlobalSnapshot
	json.Unmarshal(s1, &gs1)
	json.Unmarshal(s2, &gs2)

	if gs1.Hash == gs2.Hash {
		t.Error("expected distinct states to yield different hashes, but got duplicate hashes")
	}
}


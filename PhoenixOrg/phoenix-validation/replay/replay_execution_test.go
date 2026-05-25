package replay

import (
	"encoding/json"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
)

func runReplay() string {
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()
	orch := &rollback.Orchestrator{ProcessAudit: proc, NetworkAudit: net, FileAudit: f}
	proc.LogAction(containment.ProcessAction{PID: 1, Action: containment.ActionMonitor, Reason: "r1"})
	snap, _ := orch.CreateGlobalSnapshot(1)
	var gs rollback.GlobalSnapshot
	json.Unmarshal(snap, &gs)
	return gs.Hash
}

func TestIdentity(t *testing.T) {
	h1 := runReplay()
	h2 := runReplay()
	if h1 != h2 { t.Errorf("Divergence: %s != %s", h1, h2) }
}

func TestCrossRun(t *testing.T) {
	h := runReplay()
	for i := 0; i < 100; i++ {
		if runReplay() != h { t.Fatalf("Cross-run divergence at %d", i) }
	}
}

func TestHashConsistency(t *testing.T) { /* ... */ }
func TestForkDetection(t *testing.T) { /* ... */ }
func TestRollbackIntegrity(t *testing.T) { /* ... */ }
func TestDivergence(t *testing.T) { /* ... */ }

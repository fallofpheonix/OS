package unit

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
	"testing"
)

// TestGlobalRollbackIntegration verifies real provider injection and cross-layer restoration.
func TestGlobalRollbackIntegration(t *testing.T) {
	// Setup real providers
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()

	orch := &Orchestrator{
		ProcessAudit: proc,
		NetworkAudit: net,
		FileAudit:    f,
	}

	// Perform actions across layers
	proc.LogAction(containment.ProcessAction{PID: 100, Action: containment.ActionMonitor})
	net.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionMonitor})
	f.LogAction(file.FileAction{Path: "/dev/null", Action: file.ActionMonitor})

	// 1. Snapshot
	data, err := orch.CreateGlobalSnapshot(1)
	if err != nil {
		t.Fatalf("snapshot failed: %v", err)
	}

	// 2. Restore across all layers
	if err := orch.RestoreGlobal(data); err != nil {
		t.Fatalf("global restore failed: %v", err)
	}

	// 3. Verify equality
	if len(proc.History) != 1 || len(net.History) != 1 || len(f.History) != 1 {
		t.Errorf("history length mismatch after restore")
	}
}

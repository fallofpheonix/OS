/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/file"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/network"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/rollback"
	"testing"
)

// TestGlobalRollbackIntegration verifies real provider injection and cross-layer restoration.
func TestGlobalRollbackIntegration(t *testing.T) {
	// Setup real providers
	proc := containment.NewProcessAudit()
	net := network.NewNetworkAudit()
	f := file.NewFileAudit()

	orch := &rollback.Orchestrator{
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

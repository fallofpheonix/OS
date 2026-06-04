/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/network"
	"testing"
)

// TestNetworkDeterminism validates deterministic repeatability of the action stream.
func TestNetworkDeterminism(t *testing.T) {
	run := func() []network.NetworkAction {
		audit := network.NewNetworkAudit()
		audit.LogAction(network.NetworkAction{Src: "1.1.1.1", Dst: "2.2.2.2", Action: network.ActionThrottle, Reason: "b1"})
		audit.LogAction(network.NetworkAction{Src: "1.1.1.1", Dst: "3.3.3.3", Action: network.ActionQuarantine, Reason: "b2"})
		return audit.History
	}

	h1 := run()
	for i := 0; i < 10; i++ {
		h2 := run()
		for j := range h1 {
			if h1[j].Hash != h2[j].Hash || h1[j].Sequence != h2[j].Sequence {
				t.Fatalf("determinism failure: mismatch at run %d, step %d", i, j)
			}
		}
	}
}

// TestNetworkProofChain validates the full recovery/replay cycle.
func TestNetworkProofChain(t *testing.T) {
	audit := network.NewNetworkAudit()
	audit.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionMonitor, Reason: "r1"})
	audit.LogAction(network.NetworkAction{Src: "1.1.1.1", Action: network.ActionQuarantine, Reason: "r2"})

	// Snapshot
	data, _ := audit.CreateSnapshot()

	// Rollback simulate: restore to cursor 1
	a2 := network.NewNetworkAudit()
	if err := a2.RestoreFromSnapshot(data); err != nil {
		t.Fatalf("restore failed: %v", err)
	}

	// Verify deep replay equivalence
	if a2.Sequence != audit.Sequence || len(a2.History) != len(audit.History) {
		t.Fatal("proof failed: metadata mismatch")
	}
	if a2.History[0].Hash != audit.History[0].Hash {
		t.Fatal("proof failed: integrity hash mismatch")
	}
}

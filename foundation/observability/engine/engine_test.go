/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package engine

import (
	"testing"
)

func TestGraphEngine_Lineage(t *testing.T) {
	e := NewGraphEngine()

	// Pre-register root/init to prevent re-parenting of first nodes
	e.nodes[1] = &ProcessNode{PID: 1, Status: "ACTIVE"}

	// 1. Test Fork & Re-parenting
	e.HandleFork(100, 1)   // Normal
	e.HandleFork(200, 100) // Child
	e.HandleFork(300, 999) // Orphan -> 0xDEADBEEF

	n100, _ := e.GetNode(100)
	if n100.PPID != 1 {
		t.Errorf("Expected PID 100 PPID 1, got %d", n100.PPID)
	}

	n300, _ := e.GetNode(300)
	if n300.PPID != 0xDEADBEEF {
		t.Errorf("Expected PID 300 to be re-parented to 0xDEADBEEF, got %d", n300.PPID)
	}

	// 2. Test Virtual Fork (Isolation)
	e.HandleIsolation(200)
	n200, _ := e.GetNode(200)
	if n200.Status != "ISOLATED" {
		t.Errorf("Expected PID 200 status ISOLATED, got %s", n200.Status)
	}
	if len(n200.Children) == 0 {
		t.Error("Expected virtual fork child for PID 200")
	}

	// 3. Test Causal Chain Retrieval
	chain, _ := e.GetCausalChain(200)
	if len(chain) != 3 { // 200 -> 100 -> 1
		t.Errorf("Expected chain length 3, got %d", len(chain))
	}

	orphanChain, _ := e.GetCausalChain(300)
	if len(orphanChain) != 2 || orphanChain[1].PID != 0xDEADBEEF {
		t.Error("Orphan chain failed to include 0xDEADBEEF")
	}
}

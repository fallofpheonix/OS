/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package replication

import (
	"testing"
)

func TestReplicationEngine_ConflictResolution(t *testing.T) {
	engine := NewReplicationEngine()

	entry1 := SyncEntry{
		Index:  10,
		NodeID: "node-alpha",
		Weight: 0.5,
		Hash:   "hash-alpha",
	}

	if err := engine.Replicate(entry1); err != nil {
		t.Fatalf("Failed to replicate entry 1: %v", err)
	}

	// Conflict: Same index, higher weight node
	entry2 := SyncEntry{
		Index:  10,
		NodeID: "node-beta",
		Weight: 0.9,
		Hash:   "hash-beta",
	}

	if err := engine.Replicate(entry2); err != nil {
		t.Fatalf("Failed to resolve conflict with higher weight: %v", err)
	}

	if engine.localChain[10].NodeID != "node-beta" {
		t.Error("Replication failed to overwrite with higher-weight entry")
	}

	// Conflict: Same index, lower weight node (should be rejected)
	entry3 := SyncEntry{
		Index:  10,
		NodeID: "node-gamma",
		Weight: 0.1,
		Hash:   "hash-gamma",
	}

	if err := engine.Replicate(entry3); err == nil {
		t.Error("Allowed replication of lower-weight entry over existing higher-weight")
	}
}

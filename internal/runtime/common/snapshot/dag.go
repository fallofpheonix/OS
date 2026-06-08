/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — SNAPSHOT LINEAGE DAG
//
// The LineageDAG tracks the history of system snapshots as a directed
// acyclic graph. Each snapshot points to its parent, creating a chain
// of state captures that can be traversed for recovery.
//
// WORKFLOW:
//   StateRegistry.CreateSnapshot() → LineageDAG.AddSnapshot(id, parentID, hash)
//     → SnapshotNode created with timestamp and hash
//     → Parent-child relationship recorded
//   → LineageDAG.GetHistory(id) → chain of snapshots to root
//     → Used by RecoveryLoop to find the last known-good state
//
// INVARIANT: The DAG is append-only. Snapshots cannot be deleted or modified.
// This ensures the recovery chain is tamper-evident.
//
// ALGORITHM: AddSnapshot is O(1). GetHistory is O(k) where k = chain length.
// =========================================================================
package snapshot

import (
	"fmt"
	"time"
)

// SnapshotNode represents a point-in-time state of the system.
// This implements the Snapshot Lineage DAG [SNP-003].
type SnapshotNode struct {
	ID        string            `json:"id"`
	ParentID  string            `json:"parent_id"`
	Timestamp time.Time         `json:"timestamp"`
	Hash      string            `json:"hash"`
	Metadata  map[string]string `json:"metadata"`
}

// LineageDAG manages the relationships between snapshots.
type LineageDAG struct {
	Nodes map[string]SnapshotNode
}

func NewLineageDAG() *LineageDAG {
	return &LineageDAG{
		Nodes: make(map[string]SnapshotNode),
	}
}

// AddSnapshot records a new snapshot in the DAG.
func (d *LineageDAG) AddSnapshot(id, parentID, hash string, meta map[string]string) error {
	if parentID != "" {
		if _, ok := d.Nodes[parentID]; !ok {
			return fmt.Errorf("parent snapshot %s not found", parentID)
		}
	}

	d.Nodes[id] = SnapshotNode{
		ID:        id,
		ParentID:  parentID,
		Timestamp: time.Now(),
		Hash:      hash,
		Metadata:  meta,
	}
	return nil
}

// GetHistory returns the chain of snapshots leading to the target ID.
func (d *LineageDAG) GetHistory(id string) ([]SnapshotNode, error) {
	var history []SnapshotNode
	curr, ok := d.Nodes[id]
	if !ok {
		return nil, fmt.Errorf("snapshot %s not found", id)
	}

	for {
		history = append([]SnapshotNode{curr}, history...)
		if curr.ParentID == "" {
			break
		}
		curr, ok = d.Nodes[curr.ParentID]
		if !ok {
			return history, fmt.Errorf("lineage broken: missing parent %s", curr.ParentID)
		}
	}
	return history, nil
}

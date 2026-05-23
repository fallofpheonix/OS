package snapshot

import (
	"fmt"
	"time"
)

// SnapshotNode represents a point-in-time state of the system.
// This implements the Snapshot Lineage DAG [SNP-003].
type SnapshotNode struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	Timestamp time.Time `json:"timestamp"`
	Hash      string    `json:"hash"`
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

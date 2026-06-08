// Package replication implements synchronization logic for the Merkle DAG across nodes.
// Domain Logic: Manages consistent distribution of the ledger using weight-based conflict resolution.
// Responsibility: Ensures all nodes in the distributed nexus converge on the same monotonic chain.
package replication

import (
	"fmt"
	"log"
	"sync"
)

// SyncEntry represents an entry being replicated across the distributed cluster.
// Concurrency: Read-only instances are thread-safe.
// State Management: Encapsulates index, node identity, weight, and payload for a single replication unit.
type SyncEntry struct {
	Index   uint64
	NodeID  string
	Weight  float64
	Hash    string
	Payload []byte
}

// ReplicationEngine manages the consistent distribution of the ledger.
// Concurrency: Thread-safe via sync.RWMutex.
// State Management: Maintains a local map of the chain and tracks the maximum verified index.
type ReplicationEngine struct {
	mu         sync.RWMutex
	localChain map[uint64]SyncEntry
	maxIndex   uint64
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewReplicationEngine initializes a new replication engine instance.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewReplicationEngine() *ReplicationEngine {
	return &ReplicationEngine{
		localChain: make(map[uint64]SyncEntry),
	}
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// Replicate ingests a sync entry and resolves potential conflicts using higher-weight-wins logic.
// I/O: None.
// Side Effects: Modifies localChain map and updates maxIndex. Logs conflict resolutions.
// Complexity: O(1) average case for map operations.
func (e *ReplicationEngine) Replicate(entry SyncEntry) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	existing, exists := e.localChain[entry.Index]
	if !exists {
		e.localChain[entry.Index] = entry
		if entry.Index > e.maxIndex {
			e.maxIndex = entry.Index
		}
		return nil
	}

	if entry.Weight > existing.Weight {
		log.Printf("[REPLICATION] Conflict Resolved: Overwriting index %d with higher-weight node %s", entry.Index, entry.NodeID)
		e.localChain[entry.Index] = entry
	} else {
		return fmt.Errorf("REPLICATION_REJECTED: node %s weight (%.2f) <= local weight (%.2f)", entry.NodeID, entry.Weight, existing.Weight)
	}

	return nil
}

// LABEL: [READ_ONLY] [UNCONSTRAINED] [STABLE]
// GetHead returns the latest verified index in the replicated chain.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func (e *ReplicationEngine) GetHead() uint64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.maxIndex
}

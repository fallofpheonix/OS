/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: snapshot.go
 * PATH: Phoenix.Nucleus/ledger/snapshot.go
 *
 * PURPOSE:
 * Implements the Snapshot Engine to address the "Replay Explosion" risk (Q818).
 * Provides bit-for-bit deterministic state capture and verification.
 *
 * SUBSYSTEM:
 * Nucleus / Ledger Cycle / Recovery
 *
 * DEPENDENCIES:
 * crypto/sha256, encoding/json, fmt
 *
 * DEPENDENTS:
 * Phoenix.Nucleus/recovery
 *
 * SECURITY:
 * Snapshots are hash-verified to prevent state corruption during cold-boot.
 *
 * PERFORMANCE:
 * Reduces recovery time from O(LedgerLength) to O(EventsSinceSnapshot).
 */

package ledger

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// BEGINNER EXPLANATION:
// This file is like "Saving your game." Instead of reading every page of the 
// journal every time we start, we take a photo of the current state. 
// If the photo matches the journal's history, we can start from the photo.

// INTERMEDIATE EXPLANATION:
// The SnapshotEngine serializes system state (Authority, Memory, etc.) at 
// specific Ledger sequences. It addresses the scalability limit of linear replay.

// EXPERT EXPLANATION:
// Implements Axiom Q821. Snapshots provide deterministic state anchors. 
// The system can reconstruct its entire cognitive web by loading the 
// nearest Snapshot and replaying subsequent Ledger events, ensuring 
// bit-level parity with the original live state.

/**
 * Snapshot
 *
 * Represents a serialized state at a specific ledger sequence.
 */
type Snapshot struct {
	Sequence uint64 `json:"sequence"`
	State    []byte `json:"state"`
	Hash     string `json:"hash"`
}

/**
 * SnapshotEngine
 *
 * Manages state captures to address Replay Explosion.
 *
 * Responsibilities:
 * - State serialization and hashing.
 * - Integrity verification.
 */
type SnapshotEngine struct {
	Snapshots map[uint64]*Snapshot
}

/**
 * NewSnapshotEngine
 *
 * Initializes a new snapshot engine.
 */
func NewSnapshotEngine() *SnapshotEngine {
	return &SnapshotEngine{
		Snapshots: make(map[uint64]*Snapshot),
	}
}

/**
 * Take
 *
 * Captures the current state and records it as a snapshot.
 *
 * Input:
 * - seq: Current ledger sequence.
 * - state: Interface to be serialized (usually Registry or Memory).
 */
func (se *SnapshotEngine) Take(seq uint64, state interface{}) (*Snapshot, error) {
	data, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)
	s := &Snapshot{
		Sequence: seq,
		State:    data,
		Hash:     fmt.Sprintf("%x", hash),
	}

	se.Snapshots[seq] = s
	return s, nil
}

/**
 * Verify
 *
 * Checks the snapshot for bit-for-bit integrity.
 */
func (se *SnapshotEngine) Verify(seq uint64) bool {
	s, ok := se.Snapshots[seq]
	if !ok {
		return false
	}
	
	hash := sha256.Sum256(s.State)
	return fmt.Sprintf("%x", hash) == s.Hash
}

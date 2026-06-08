/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package state

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Snapshot represents a point-in-time capture of the StateRegistry.
type Snapshot struct {
	CurrentState SystemState   `json:"current_state"`
	History      []StateRecord `json:"history"`
	Version      string        `json:"version"`
	Sequence     int           `json:"sequence"`
	Hash         string        `json:"hash"`
}

const CurrentVersion = "1.0.0"

// CalculateHash computes a deterministic hash of the full snapshot content.
func CalculateHash(s Snapshot) string {
	b, _ := json.Marshal(struct {
		State    SystemState
		History  []StateRecord
		Version  string
		Sequence int
	}{
		State:    s.CurrentState,
		History:  s.History,
		Version:  s.Version,
		Sequence: s.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// CreateSnapshot serializes the current registry state with integrity metadata.
func (r *StateRegistry) CreateSnapshot() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s := Snapshot{
		CurrentState: r.CurrentState,
		History:      r.History,
		Version:      CurrentVersion,
		Sequence:     len(r.History),
	}
	s.Hash = CalculateHash(s)

	GlobalMetrics.IncSnapshotCreates()
	return json.Marshal(s)
}

// RestoreFromSnapshot deserializes and validates a snapshot.
func (r *StateRegistry) RestoreFromSnapshot(data []byte) error {
	var s Snapshot
	if err := json.Unmarshal(data, &s); err != nil {
		GlobalMetrics.IncReplayMismatch()
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	// Integrity validation
	if s.Version != CurrentVersion {
		GlobalMetrics.IncReplayMismatch()
		return fmt.Errorf("incompatible version: %s", s.Version)
	}
	if CalculateHash(s) != s.Hash {
		GlobalMetrics.IncReplayMismatch()
		return fmt.Errorf("snapshot integrity check failed")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.CurrentState = s.CurrentState
	r.History = s.History
	GlobalMetrics.IncSnapshotRestores()
	return nil
}

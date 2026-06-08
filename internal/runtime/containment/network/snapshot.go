/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package network

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// NetworkSnapshot captures the audit state for network containment.
type NetworkSnapshot struct {
	History      []NetworkAction `json:"history"`
	Sequence     int             `json:"sequence"`
	ReplayCursor int             `json:"replay_cursor"`
	Version      string          `json:"version"`
	Hash         string          `json:"hash"`
}

const NetworkVersion = "1.0.0"

func calculateSnapshotHash(s NetworkSnapshot) string {
	b, _ := json.Marshal(struct {
		History      []NetworkAction
		Sequence     int
		ReplayCursor int
		Version      string
	}{
		History:      s.History,
		Sequence:     s.Sequence,
		ReplayCursor: s.ReplayCursor,
		Version:      s.Version,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// Normalize clears timestamps for deterministic snapshotting.
func (a *NetworkAudit) Normalize() {
	for i := range a.History {
		a.History[i].Timestamp = time.Unix(0, 0)
	}
}

// CreateSnapshot serializes the network audit state with integrity metadata.
func (a *NetworkAudit) CreateSnapshot() ([]byte, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Normalize()

	s := NetworkSnapshot{
		History:      a.History,
		Sequence:     a.Sequence,
		ReplayCursor: a.ReplayCursor,
		Version:      NetworkVersion,
	}
	s.Hash = calculateSnapshotHash(s)

	return json.Marshal(s)
}

// RestoreFromSnapshot deserializes and validates the network audit state.
func (a *NetworkAudit) RestoreFromSnapshot(data []byte) error {
	var s NetworkSnapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	if s.Version != NetworkVersion {
		return fmt.Errorf("incompatible version: %s", s.Version)
	}
	if calculateSnapshotHash(s) != s.Hash {
		return fmt.Errorf("snapshot integrity check failed")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.History = s.History
	a.Sequence = s.Sequence
	a.ReplayCursor = s.ReplayCursor
	return nil
}

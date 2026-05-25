package containment

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Snapshot represents a point-in-time capture of the IsolationEngine.
type Snapshot struct {
	CurrentState IsolationState    `json:"current_state"`
	History      []IsolationRecord `json:"history"`
	Version      string            `json:"version"`
	Sequence     int               `json:"sequence"`
	Hash         string            `json:"hash"`
}

const CurrentVersion = "1.0.0"

// calculateHash computes a deterministic hash of the full snapshot content.
func calculateHash(s Snapshot) string {
	b, _ := json.Marshal(struct {
		State    IsolationState
		History  []IsolationRecord
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

// CreateSnapshot serializes the current engine state with integrity metadata.
func (e *IsolationEngine) CreateSnapshot() ([]byte, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	s := Snapshot{
		CurrentState: e.CurrentState,
		History:      e.History,
		Version:      CurrentVersion,
		Sequence:     len(e.History),
	}
	s.Hash = calculateHash(s)

	return json.Marshal(s)
}

// RestoreFromSnapshot deserializes and validates a snapshot.
func (e *IsolationEngine) RestoreFromSnapshot(data []byte) error {
	var s Snapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	// Integrity validation
	if s.Version != CurrentVersion {
		return fmt.Errorf("incompatible version: %s", s.Version)
	}
	if calculateHash(s) != s.Hash {
		return fmt.Errorf("snapshot integrity check failed")
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.CurrentState = s.CurrentState
	e.History = s.History
	return nil
}

package containment

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// ProcessSnapshot captures the entire audit state for process containment.
type ProcessSnapshot struct {
	History  []ProcessAction `json:"history"`
	Sequence int             `json:"sequence"`
	Version  string          `json:"version"`
	Hash     string          `json:"hash"`
}

const ProcessVersion = "1.0.0"

func calculateSnapshotHash(s ProcessSnapshot) string {
	b, _ := json.Marshal(struct {
		History  []ProcessAction
		Sequence int
		Version  string
	}{
		History:  s.History,
		Sequence: s.Sequence,
		Version:  s.Version,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// Normalize clears timestamps for deterministic snapshotting.
func (a *ProcessAudit) Normalize() {
	for i := range a.History {
		a.History[i].Timestamp = time.Unix(0, 0)
	}
}

// CreateSnapshot serializes the audit log with integrity metadata.
func (a *ProcessAudit) CreateSnapshot() ([]byte, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Normalize()

	s := ProcessSnapshot{
		History:  a.History,
		Sequence: a.Sequence,
		Version:  ProcessVersion,
	}
	s.Hash = calculateSnapshotHash(s)

	return json.Marshal(s)
}

// RestoreFromSnapshot deserializes and validates the audit log.
func (a *ProcessAudit) RestoreFromSnapshot(data []byte) error {
	var s ProcessSnapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	if s.Version != ProcessVersion {
		return fmt.Errorf("incompatible version: %s", s.Version)
	}
	if calculateSnapshotHash(s) != s.Hash {
		return fmt.Errorf("snapshot integrity check failed")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.History = s.History
	a.Sequence = s.Sequence
	return nil
}

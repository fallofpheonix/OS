package file

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// FileSnapshot captures the audit state for file containment.
type FileSnapshot struct {
	History      []FileAction `json:"history"`
	Sequence     int          `json:"sequence"`
	ReplayCursor int          `json:"replay_cursor"`
	Version      string       `json:"version"`
	Hash         string       `json:"hash"`
}

const FileVersion = "1.0.0"

func calculateSnapshotHash(s FileSnapshot) string {
	b, _ := json.Marshal(struct {
		History      []FileAction
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
func (a *FileAudit) Normalize() {
	for i := range a.History {
		a.History[i].Timestamp = time.Unix(0, 0)
	}
}

// CreateSnapshot serializes the file audit state with integrity metadata.
func (a *FileAudit) CreateSnapshot() ([]byte, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Normalize()

	s := FileSnapshot{
		History:      a.History,
		Sequence:     a.Sequence,
		ReplayCursor: a.ReplayCursor,
		Version:      FileVersion,
	}
	s.Hash = calculateSnapshotHash(s)

	return json.Marshal(s)
}

// RestoreFromSnapshot deserializes and validates the file audit state.
func (a *FileAudit) RestoreFromSnapshot(data []byte) error {
	var s FileSnapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	if s.Version != FileVersion {
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

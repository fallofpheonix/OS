package truth

import (
	"fmt"
	"time"
)

// Snapshot (B2) represents a point-in-time state of the TruthLedger.
type Snapshot struct {
	Timestamp time.Time
	LastIndex int
	LastHash  string
}

// ReplaySnapshotEngine (B2) manages the creation and retrieval of snapshots.
type ReplaySnapshotEngine struct {
	Snapshots []Snapshot
}

func NewReplaySnapshotEngine() *ReplaySnapshotEngine {
	return &ReplaySnapshotEngine{
		Snapshots: make([]Snapshot, 0),
	}
}

// CreateSnapshot creates a new snapshot from the current ledger state.
func (e *ReplaySnapshotEngine) CreateSnapshot(l *TruthLedger) (*Snapshot, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	snap := Snapshot{
		Timestamp: time.Now(),
		LastIndex: len(l.Entries) - 1,
		LastHash:  l.LastHash,
	}
	e.Snapshots = append(e.Snapshots, snap)
	return &snap, nil
}

// ConsistencyChecker (B6) validates that the ledger hasn't diverged from a snapshot.
type ConsistencyChecker struct{}

func (c *ConsistencyChecker) CheckConsistency(l *TruthLedger, snap *Snapshot) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if snap.LastIndex >= len(l.Entries) {
		return fmt.Errorf("snapshot index %d out of range", snap.LastIndex)
	}

	currentHashAtIndex := l.Entries[snap.LastIndex].Hash
	if currentHashAtIndex != snap.LastHash {
		return fmt.Errorf("ledger hash at index %d does not match snapshot", snap.LastIndex)
	}

	return nil
}

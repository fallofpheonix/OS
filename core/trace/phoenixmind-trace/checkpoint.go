package trace

import (
	"time"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

// Checkpoint represents a snapshot of the system's truth state at a specific point in time.
type Checkpoint struct {
	ID        string
	Timestamp time.Time
	RootHash  string                       // A hash of the entire truth graph state at this point.
	EntityStates map[string]evidence.TruthState // Snapshot of key entity states.
	Metadata  map[string]interface{}
}

// CheckpointManager handles the creation and retrieval of system checkpoints.
type CheckpointManager struct {
	Checkpoints []Checkpoint
}

// CreateCheckpoint generates a new checkpoint from the current system state.
func (cm *CheckpointManager) CreateCheckpoint(rootHash string, entityStates map[string]evidence.TruthState) Checkpoint {
	cp := Checkpoint{
		ID:        "cp-" + time.Now().Format("20060102150405"), // Simple ID generation
		Timestamp: time.Now(),
		RootHash:  rootHash,
		EntityStates: entityStates,
	}
	cm.Checkpoints = append(cm.Checkpoints, cp)
	return cp
}

// GetCheckpoint retrieves a checkpoint by its ID.
func (cm *CheckpointManager) GetCheckpoint(id string) (Checkpoint, bool) {
	for _, cp := range cm.Checkpoints {
		if cp.ID == id {
			return cp, true
		}
	}
	return Checkpoint{}, false
}

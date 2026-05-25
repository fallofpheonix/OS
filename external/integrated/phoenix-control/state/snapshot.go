package state

import (
	"encoding/json"
	"fmt"
)

// StateSnapshot captures the state of the registry at a point in time.
type StateSnapshot struct {
	CurrentState RuntimeState `json:"current_state"`
	Tick         int64        `json:"tick"`
	Checksum     string       `json:"checksum"` // SHA-256 placeholder
}

// StateSerializer handles the conversion of state snapshots to/from bytes.
type StateSerializer struct{}

func NewStateSerializer() *StateSerializer {
	return &StateSerializer{}
}

func (s *StateSerializer) Serialize(snapshot StateSnapshot) ([]byte, error) {
	return json.Marshal(snapshot)
}

func (s *StateSerializer) Deserialize(data []byte) (StateSnapshot, error) {
	var snapshot StateSnapshot
	err := json.Unmarshal(data, &snapshot)
	if err != nil {
		return StateSnapshot{}, fmt.Errorf("deserialization failed: %w", err)
	}
	return snapshot, nil
}

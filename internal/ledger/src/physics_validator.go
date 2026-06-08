package ledger

import (
	"encoding/json"
	"errors"
	"fmt"
)

// PhysicsValidator implements the SemanticValidator contract for L0 physical invariants.
type PhysicsValidator struct {
	lastTick uint64
}

func NewPhysicsValidator() *PhysicsValidator {
	return &PhysicsValidator{}
}

// Validate checks basic range and monotonicity constraints.
// CONTRACT: Coherence and Entropy must be in [0, 1] range.
//
//	Ticks must be monotonic.
func (v *PhysicsValidator) Validate(eventID, causeID string, tick uint64, payload []byte) error {
	if eventID == "" {
		return errors.New("physics violation: empty EventID")
	}

	if tick < v.lastTick {
		return fmt.Errorf("physics violation: temporal paradox (tick %d < last %d)", tick, v.lastTick)
	}

	// Range checks for physical metrics
	var wrapper struct {
		Raw json.RawMessage `json:"raw"`
	}

	if err := json.Unmarshal(payload, &wrapper); err == nil {
		var data struct {
			Coherence int64 `json:"coherence"` // Scaled by 1,000,000
			Entropy   int64 `json:"entropy"`   // Scaled by 1,000,000
		}
		if err := json.Unmarshal(wrapper.Raw, &data); err == nil {
			if data.Coherence < 0 || data.Coherence > 1000000 {
				return fmt.Errorf("physics violation: coherence out of range (%d)", data.Coherence)
			}
			if data.Entropy < 0 || data.Entropy > 1000000 {
				return fmt.Errorf("physics violation: entropy out of range (%d)", data.Entropy)
			}
		}
	}

	v.lastTick = tick
	return nil
}

package engine

import (
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// VerificationRule defines the interface for deterministic game-rule validation.
type VerificationRule interface {
	Verify(ws *WorldState, entityID string) (bool, string)
}

// ProximityRule ensures an entity has reached its target position.
type ProximityRule struct {
	TargetPos phxmath.FixedPoint
	Threshold phxmath.FixedPoint
}

func (r *ProximityRule) Verify(ws *WorldState, entityID string) (bool, string) {
	e, ok := ws.GetEntity(entityID)
	if !ok {
		return false, "entity not found"
	}

	// Simple 1D distance: abs(e.Pos - r.TargetPos) <= r.Threshold
	diff := e.Pos.SaturatingSub(r.TargetPos)
	if diff.V < 0 {
		diff.V = -diff.V
	}

	if diff.V <= r.Threshold.V {
		return true, "REACHED_TARGET"
	}

	return false, "TOO_FAR"
}

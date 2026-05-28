package truth_runtime

import (
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
	"github.com/fallofpheonix/phoenixmind-validator/truth/registry"
	"github.com/fallofpheonix/phoenixmind-validator/truth/resolver"
	"github.com/fallofpheonix/phoenixmind-validator/truth/scoring"
)

// Runtime orchestrates the truth validation process.
type Runtime struct {
	TruthRegistry *registry.TruthRegistry
}

func NewRuntime() *Runtime {
	return &Runtime{
		TruthRegistry: registry.NewTruthRegistry(),
	}
}

// IngestEvidence scores, resolves, and updates the truth state for a new piece of evidence.
func (r *Runtime) IngestEvidence(rawEvidence evidence.Evidence) (evidence.TruthState, error) {
	// 1. Score the evidence
	score := scoring.CalculateTruthScore(rawEvidence)
	_ = score // Score will be used later for more complex logic

	// 2. Update the entity in the registry
	r.TruthRegistry.UpdateEntity(rawEvidence)

	// 3. Retrieve the updated entity to get its final state
	entity, ok := r.TruthRegistry.GetEntity(rawEvidence.EntityID)
	if !ok {
		// This should not happen if UpdateEntity works correctly
		return evidence.UNKNOWN, nil
	}

	return entity.CurrentState, nil
}

// ReplayEvidence validates a sequence of historical events.
func (r *Runtime) ValidateHistory(history []evidence.Evidence) bool {
	// Placeholder for more complex replay validation
	// For now, just ensures resolution is consistent.
	finalState := resolver.MergeTruth(history)
	return finalState != evidence.UNKNOWN
}

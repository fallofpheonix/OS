package validators

import (
	"fmt"
	"github.com/fallofpheonix/PheonixTruth/src"
	"github.com/fallofpheonix/PheonixGuard"
)

// TransitionValidator enforces the Warden's Finite State Machine (FSM) rules.
type TransitionValidator struct {
	allowedTransitions map[warden.SystemState][]warden.SystemState
}

// NewTransitionValidator creates a new TransitionValidator with the default Warden FSM rules.
func NewTransitionValidator() *TransitionValidator {
	return &TransitionValidator{
		allowedTransitions: map[warden.SystemState][]warden.SystemState{
			warden.StateSafe:        {warden.StateWatch},
			warden.StateWatch:       {warden.StateSafe, warden.StateSuspicious},
			warden.StateSuspicious:  {warden.StateWatch, warden.StateCritical},
			warden.StateCritical:    {warden.StateSuspicious, warden.StateCompromised},
			warden.StateCompromised: {warden.StateCritical}, // Recovery path
		},
	}
}

// Name returns the validator name.
func (v *TransitionValidator) Name() string {
	return "TransitionValidator"
}

// Validate checks if the state transition in the ledger entry is valid according to the FSM.
func (v *TransitionValidator) Validate(entry *ledger.LedgerEntry) ValidationResult {
	before := warden.SystemState(entry.StateBefore)
	after := warden.SystemState(entry.StateAfter)

	if before == "" || after == "" {
		// If states are not provided, we might be dealing with a non-state-changing event.
		// For now, we allow it if both are empty, or flag it if only one is empty.
		if before == after {
			return ValidationResult{Valid: true}
		}
		return ValidationResult{
			Valid:  false,
			Reason: fmt.Sprintf("Missing state information: %s -> %s", entry.StateBefore, entry.StateAfter),
		}
	}

	if before == after {
		return ValidationResult{Valid: true}
	}

	allowed, ok := v.allowedTransitions[before]
	if !ok {
		return ValidationResult{
			Valid:  false,
			Reason: fmt.Sprintf("Unknown source state: %s", before),
		}
	}

	for _, a := range allowed {
		if a == after {
			return ValidationResult{Valid: true}
		}
	}

	return ValidationResult{
		Valid:  false,
		Reason: fmt.Sprintf("Illegal state transition: %s -> %s", before, after),
	}
}

// Reset is a no-op for the stateless TransitionValidator.
func (v *TransitionValidator) Reset() {}

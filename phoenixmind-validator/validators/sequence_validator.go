package validators

import (
	"fmt"
	"github.com/fallofpheonix/PheonixTruth/src"
)

// SequenceValidator ensures LogicalTick continuity in the execution trace.
type SequenceValidator struct {
	lastTick uint64
}

// NewSequenceValidator creates a new SequenceValidator.
func NewSequenceValidator() *SequenceValidator {
	return &SequenceValidator{
		lastTick: 0,
	}
}

// Name returns the validator name.
func (v *SequenceValidator) Name() string {
	return "SequenceValidator"
}

// Validate checks if the current entry's LogicalTick is strictly greater than the last one.
func (v *SequenceValidator) Validate(entry *ledger.LedgerEntry) ValidationResult {
	if v.lastTick != 0 && entry.LogicalTick <= v.lastTick {
		return ValidationResult{
			Valid:  false,
			Reason: fmt.Sprintf("Sequence regression: current tick %d <= last tick %d", entry.LogicalTick, v.lastTick),
		}
	}
	v.lastTick = entry.LogicalTick
	return ValidationResult{Valid: true}
}

// Reset clears the validator state.
func (v *SequenceValidator) Reset() {
	v.lastTick = 0
}

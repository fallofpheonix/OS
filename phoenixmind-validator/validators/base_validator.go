package validators

import (
	"github.com/fallofpheonix/PheonixTruth/src"
)

// ValidationResult represents the outcome of a validation check.
type ValidationResult struct {
	Valid  bool
	Reason string
	Error  error
}

// Validator is the core interface for all security rule-checkers.
type Validator interface {
	// Validate checks a ledger entry against specific invariants.
	Validate(entry *ledger.LedgerEntry) ValidationResult
	// Name returns the unique identifier for the validator.
	Name() string
	// Reset resets the internal state of the validator (if any).
	Reset()
}

package validators

import (
	"github.com/fallofpheonix/PheonixTruth/src"
)

// ValidatorRegistry is a central hub for managing and executing validators.
type ValidatorRegistry struct {
	validators []Validator
}

// NewValidatorRegistry creates a new ValidatorRegistry.
func NewValidatorRegistry() *ValidatorRegistry {
	return &ValidatorRegistry{
		validators: make([]Validator, 0),
	}
}

// Register adds a validator to the registry.
func (r *ValidatorRegistry) Register(v Validator) {
	r.validators = append(r.validators, v)
}

// ValidateAll executes all registered validators against a ledger entry.
// It returns a slice of failed validation results.
func (r *ValidatorRegistry) ValidateAll(entry *ledger.LedgerEntry) []ValidationResult {
	var failures []ValidationResult
	for _, v := range r.validators {
		res := v.Validate(entry)
		if !res.Valid {
			failures = append(failures, res)
		}
	}
	return failures
}

// GetValidators returns the list of registered validators.
func (r *ValidatorRegistry) GetValidators() []Validator {
	return r.validators
}

// ResetAll resets the state of all registered validators.
func (r *ValidatorRegistry) ResetAll() {
	for _, v := range r.validators {
		v.Reset()
	}
}

/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package validators

import (
	"github.com/fallofpheonix/phoenix/governance/truth/engine"
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

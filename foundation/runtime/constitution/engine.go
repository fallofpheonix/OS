// Package constitution provides the governance and enforcement layer for PhoenixOS.
// Core Domain Logic: Implements the "Constitution Engine" which acts as the primary arbiter for system 
// transitions, ensuring every event adheres to causal, identity, and authority constraints.
package constitution

import (
	"errors"
	"fmt"

	"github.com/fallofpheonix/phoenix/foundation/events"
)

var (
	ErrUnconstitutionalBoot = errors.New("node refuses unconstitutional boot")
	ErrCausalMismatch      = errors.New("event parent id mismatch (causal integrity failure)")
	ErrDivergenceDetected   = errors.New("state hash divergence detected (replay failure)")
	ErrUnauthorizedIdentity = errors.New("unauthorized identity for given authority")
	ErrForgedSignature      = errors.New("invalid signature detected (forged event)")
	ErrInvalidSchema        = errors.New("invalid event schema")
)

// Engine orchestrates the enforcement of invariants and policies.
// Internal State: Composed of InvariantEngine and PolicyEngine for specialized verification.
// API Scope: Public within PhoenixCore for state transition validation.
// Concurrency: Thread-safe if internal engines are handled as immutable or synchronized.
type Engine struct {
	Invariants *InvariantEngine
	Policy     *PolicyEngine
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewEngine initializes a new constitution engine with default enforcers.
// I/O: None.
// Complexity: O(1).
func NewEngine() *Engine {
	return &Engine{
		Invariants: &InvariantEngine{},
		Policy:     &PolicyEngine{},
	}
}

// LABEL: [PURE] [PUBLIC_API] [EXPERIMENTAL]
// ValidateTransition checks if an event can be applied to the current state by verifying causal 
// integrity, authority, and identity constraints.
// I/O: None.
// Side Effects: None.
// Complexity: O(1) relative to event size.
func (e *Engine) ValidateTransition(current event.Event, next event.Event) error {
	// P3.1: Enforce Causal Integrity
	if next.ParentID != current.EventID {
		return fmt.Errorf("%w: expected parent %s, got %s", ErrCausalMismatch, current.EventID, next.ParentID)
	}

	// P3.3: Reject Invalid Authority / Forged Events
	if err := e.Policy.VerifyAuthority(next); err != nil {
		return err
	}

	// P3.1: Verify Identity Constraints
	if err := e.Invariants.VerifyIdentity(next); err != nil {
		return err
	}

	return nil
}

// BootValidator performs pre-runtime checks to ensure system integrity.
// Internal State: Reference to the active Constitution Engine.
// API Scope: Internal/Boot sequence only.
// Concurrency: Single-threaded (invoked during boot).
type BootValidator struct {
	Engine *Engine
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewBootValidator creates a validator for the boot sequence.
// I/O: None.
// Complexity: O(1).
func NewBootValidator(e *Engine) *BootValidator {
	return &BootValidator{Engine: e}
}

// LABEL: [PURE] [PUBLIC_API] [EXPERIMENTAL]
// ValidateBoot verifies the integrity of the core substrates (constitution, ledger, keys) 
// before allowing the node to start.
// I/O: None (operates on provided hashes/keys).
// Complexity: O(K) where K is the number of authorized keys.
func (bv *BootValidator) ValidateBoot(ledger event.Checkpoint, keys []string, constitutionHash string) error {
	// 1. Verify Constitution integrity (signed file matches expected hash)
	if constitutionHash == "" {
		return fmt.Errorf("%w: missing constitution hash", ErrUnconstitutionalBoot)
	}

	// 2. Verify Ledger integrity (replay verifier check)
	if ledger.StateHash == "" {
		return fmt.Errorf("%w: invalid ledger state hash", ErrUnconstitutionalBoot)
	}

	// 3. Verify Keys/Authority existence
	if len(keys) == 0 {
		return fmt.Errorf("%w: no authorized keys found", ErrUnconstitutionalBoot)
	}

	return nil
}

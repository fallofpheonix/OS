package runtime

import (
	"fmt"

	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
	"github.com/fallofpheonix/phoenix/internal/state"
)

// Dispatcher acts as the "Authority Air-Gap" between the Ledger and Physics layers.
// It verifies block integrity and applies state transitions.
type Dispatcher struct {
	Guard *state.StateGuard
	Token state.AuthorityToken
}

// NewDispatcher initializes a new dispatcher.
func NewDispatcher(guard *state.StateGuard, token state.AuthorityToken) *Dispatcher {
	return &Dispatcher{
		Guard: guard,
		Token: token,
	}
}

// ProcessBlock validates a finalized block and applies its events to the state.
// PROTOCOL-021: Authority-Blind Execution Pipeline.
func (d *Dispatcher) ProcessBlock(b contracts.FinalizedBlock) error {
	// 1. Structural and Merkle Validation (Authority Layer check)
	if err := protocol.ValidateBlock(b); err != nil {
		return fmt.Errorf("dispatcher: block validation failed: %w", err)
	}

	// 2. Physics Layer Application
	for i, ev := range b.Events {
		applied := contracts.AppliedEvent{
			Height: b.Height,
			Epoch:  b.Epoch,
			Event:  ev,
		}

		if err := d.Guard.Apply(d.Token, applied); err != nil {
			return fmt.Errorf("dispatcher: failed to apply event %d at height %d: %w", i, b.Height, err)
		}
	}

	// 3. Post-Execution Integrity Check
	// The local StateRoot MUST match the block's claimed StateRoot.
	localRoot := d.Guard.CalculateHash()
	if localRoot != b.StateRoot {
		return fmt.Errorf("dispatcher: state divergence at height %d. local=%s, block=%s (ERR_STATE)",
			b.Height, localRoot, b.StateRoot)
	}

	return nil
}

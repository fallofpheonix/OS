package replay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

// ReplayEngine handles consensus-grade event ingestion and state projection.
type ReplayEngine struct {
	mu         sync.RWMutex
	State      *engine.WorldState
	Validators [][]byte
	Rules      map[string]engine.VerificationRule
}

// NewReplayEngine creates a new replay engine.
func NewReplayEngine(state *engine.WorldState) *ReplayEngine {
	return &ReplayEngine{
		State: state,
		Rules: make(map[string]engine.VerificationRule),
	}
}

// AddAuthorizedValidator adds a public key to the set of authorized signers.
func (r *ReplayEngine) AddAuthorizedValidator(pubKey []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Validators = append(r.Validators, pubKey)
	r.State.Validators = r.Validators
}

// ProcessEnvelope validates a signed envelope and applies the inner event.
func (r *ReplayEngine) ProcessEnvelope(env *contracts.SignedEnvelope) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 1. Verify Envelope Signature
	if !consensus.VerifyEnvelope(env) {
		return fmt.Errorf("envelope signature verification failed")
	}

	// 2. Verify Validator Membership
	isAuthorized := false
	for _, v := range r.Validators {
		if bytes.Equal(v, env.Validator[:]) {
			isAuthorized = true
			break
		}
	}
	if !isAuthorized {
		return fmt.Errorf("unauthorized validator: %x", env.Validator)
	}

	// 3. Unmarshal Event
	var ev contracts.Event
	if err := json.Unmarshal(env.Payload, &ev); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// 4. Apply Event (Physics Gate)
	applied := contracts.AppliedEvent{
		Height: env.Sequence,
		Epoch:  env.Epoch,
		Event:  ev,
	}
	if err := r.State.ApplyEvent(applied, r.Rules); err != nil {
		return fmt.Errorf("failed to apply event at height %d: %w", env.Sequence, err)
	}

	// 5. Update Sequence (Forensic Gate)
	r.State.ApplyEnvelope(env)

	return nil
}

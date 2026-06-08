package state

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
)

// AuthorityToken is an uninstantiable token that grants mutation rights.
type AuthorityToken struct {
	_ struct{}
}

var (
	once           sync.Once
	authorityToken AuthorityToken
)

// RequestAuthorityToken retrieves the one-time authority token.
func RequestAuthorityToken() AuthorityToken {
	once.Do(func() {
		authorityToken = AuthorityToken{}
	})
	return authorityToken
}

// StateGuard is the thread-safe authority boundary for worldState.
type StateGuard struct {
	mu sync.RWMutex
	ws *worldState
}

// NewStateGuard initializes a new guard with a blank state.
func NewStateGuard(seed int64) *StateGuard {
	return &StateGuard{
		ws: newWorldState(seed),
	}
}

// Apply attempts to mutate the state using an enriched AppliedEvent.
// INV-004: Sole path for mutation.
// PROTOCOL-021: Context (Height, Epoch) is provided by the dispatcher.
func (g *StateGuard) Apply(token AuthorityToken, applied contracts.AppliedEvent) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	event := applied.Event

	// 1. Epoch Transition
	if applied.Epoch > g.ws.Epoch {
		g.processEpochTransition(applied.Epoch)
	}

	// 2. Ledger Invariants
	// NOTE: Index is now handled by the ledger/runtime sequence.
	// But we still track EventCount for internal consistency.

	// 3. Physics / Semantic Logic
	switch event.Type {
	case contracts.EventSpawn:
		var p struct {
			ID  string             `json:"id"`
			Pos phxmath.FixedPoint `json:"pos"`
		}
		if err := json.Unmarshal(event.Payload, &p); err != nil {
			return fmt.Errorf("state: invalid spawn payload: %w", err)
		}
		g.ws.Entities[p.ID] = &Entity{
			ID:     p.ID,
			Pos:    p.Pos,
			Status: "SPAWNED",
		}
	case contracts.EventMove:
		var p struct {
			ID  string             `json:"id"`
			Pos phxmath.FixedPoint `json:"pos"`
		}
		if err := json.Unmarshal(event.Payload, &p); err != nil {
			return fmt.Errorf("state: invalid move payload: %w", err)
		}
		entity, ok := g.ws.Entities[p.ID]
		if !ok {
			return fmt.Errorf("state: entity %s not found", p.ID)
		}
		entity.Pos = p.Pos
	case contracts.EventUpdateValidator:
		var p struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		}
		if err := json.Unmarshal(event.Payload, &p); err != nil {
			return fmt.Errorf("state: invalid validator payload: %w", err)
		}
		vKey, _ := hex.DecodeString(p.ID)
		if len(g.ws.PendingValidators) == 0 && len(g.ws.Validators) > 0 {
			g.ws.PendingValidators = make([][]byte, len(g.ws.Validators))
			copy(g.ws.PendingValidators, g.ws.Validators)
		}
		if p.Status == "ADD" {
			g.ws.PendingValidators = append(g.ws.PendingValidators, vKey)
		} else if p.Status == "REMOVE" {
			for i, v := range g.ws.PendingValidators {
				if bytes.Equal(v, vKey) {
					g.ws.PendingValidators = append(g.ws.PendingValidators[:i], g.ws.PendingValidators[i+1:]...)
					break
				}
			}
		}
	}

	// 4. Finalize
	g.ws.Tick = applied.Height
	g.ws.EventCount++

	d, err := protocol.DigestEvent(event)
	if err != nil {
		return fmt.Errorf("state: failed to digest event: %w", err)
	}
	g.ws.LastEventHash = d
	g.ws.StateHash = g.ws.calculateHash()

	return nil
}

func (g *StateGuard) processEpochTransition(epoch uint64) {
	if len(g.ws.PendingValidators) > 0 {
		g.ws.Validators = make([][]byte, len(g.ws.PendingValidators))
		for i, v := range g.ws.PendingValidators {
			g.ws.Validators[i] = make([]byte, len(v))
			copy(g.ws.Validators[i], v)
		}
		g.ws.PendingValidators = nil
	}
	g.ws.Epoch = epoch
}

// CalculateHash returns the current state commitment (StateRoot).
func (g *StateGuard) CalculateHash() contracts.Hash {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.ws.StateHash
}

// Snapshot produces a durable commitment.
func (g *StateGuard) Snapshot() ([]byte, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return json.Marshal(g.ws)
}

// Restore performs an atomic installation.
func (g *StateGuard) Restore(data []byte, expectedHash contracts.Hash) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var newState worldState
	if err := json.Unmarshal(data, &newState); err != nil {
		return err
	}

	actualHash := newState.calculateHash()
	if actualHash != expectedHash {
		return fmt.Errorf("restore failed: hash mismatch. actual=%s, expected=%s", actualHash, expectedHash)
	}

	g.ws = &newState
	return nil
}

// ApplyEnvelope updates forensic metadata.
func (g *StateGuard) ApplyEnvelope(nodeID string, sequence uint64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.ws.LastSeenSequences == nil {
		g.ws.LastSeenSequences = make(map[string]uint64)
	}
	g.ws.LastSeenSequences[nodeID] = sequence
	g.ws.StateHash = g.ws.calculateHash()
}

// GetLastSequence retrieves forensic metadata.
func (g *StateGuard) GetLastSequence(nodeID string) uint64 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.ws.LastSeenSequences[nodeID]
}

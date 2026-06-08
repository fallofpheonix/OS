/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package knowledge implements the deterministic world model for PhoenixOS,
// encompassing causal graph mapping and high-confidence semantic belief management.
package knowledge

import (
	"encoding/json"
	"sync"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// BeliefState defines the lifecycle state of a semantic conclusion.
type BeliefState string

const (
	// BeliefActive indicates the belief is currently authoritative.
	BeliefActive BeliefState = "ACTIVE"
	// BeliefContested indicates the belief is under re-evaluation due to reality drift.
	BeliefContested BeliefState = "CONTESTED"
	// BeliefSuperseded indicates a newer version of this belief has been committed.
	BeliefSuperseded BeliefState = "SUPERSEDED"
	// BeliefExpired indicates the belief is no longer relevant to the current system state.
	BeliefExpired BeliefState = "EXPIRED"
)

// Belief represents a versioned, high-confidence semantic conclusion.
// It is grounded in specific FactIDs and carries a confidence score.
type Belief struct {
	ID         string                 `json:"id"`
	Version    uint32                 `json:"version"`
	State      BeliefState            `json:"state"`
	FactIDs    []string               `json:"fact_ids"`
	Confidence ledger.ConfidenceScore `json:"confidence"`
	Statement  string                 `json:"statement"`
	Utility    float64                `json:"utility"` // Prediction Utility metric
}

// BeliefEngine manages the lifecycle, versioning, and deterministic replay of beliefs.
// Internal state is protected by a RWMutex to support concurrent reasoning cycles.
type BeliefEngine struct {
	mu      sync.RWMutex
	Beliefs map[string][]*Belief // Map of BeliefID to chronological version slice
}

// NewBeliefEngine initializes an empty belief engine with an allocated map.
func NewBeliefEngine() *BeliefEngine {
	return &BeliefEngine{
		Beliefs: make(map[string][]*Belief),
	}
}

// Commit records a new version of a belief and updates historical states.
// If b.Version is 0, it is automatically incremented based on existing version count.
// Side Effects: Marks the previous latest version as SUPERSEDED.
// Complexity: O(1) time / O(1) space.
func (be *BeliefEngine) Commit(b *Belief) {
	be.mu.Lock()
	defer be.mu.Unlock()

	versions := be.Beliefs[b.ID]
	if b.Version == 0 {
		b.Version = uint32(len(versions) + 1)
	}

	if len(versions) > 0 {
		versions[len(versions)-1].State = BeliefSuperseded
	}
	be.Beliefs[b.ID] = append(be.Beliefs[b.ID], b)
}

// GetLatest retrieves the most recent version of a belief by ID.
// Returns (*Belief, true) if found, (nil, false) otherwise.
// Complexity: O(1) time / O(1) space.
func (be *BeliefEngine) GetLatest(id string) (*Belief, bool) {
	be.mu.RLock()
	defer be.mu.RUnlock()

	versions, ok := be.Beliefs[id]
	if !ok || len(versions) == 0 {
		return nil, false
	}
	return versions[len(versions)-1], true
}

// ReconstructFromLedger re-synchronizes the internal belief state from a ledger event stream.
// It enforces bit-for-bit identical state reconstruction required for deterministic replay.
// Inputs: events ([]*ledger.Event) - Slice of chronological ledger events.
// Returns: error if JSON unmarshaling fails.
// Complexity: O(N) where N is the number of belief events.
func (be *BeliefEngine) ReconstructFromLedger(events []*ledger.Event) error {
	for _, e := range events {
		if e.Type != ledger.EventBelief {
			continue
		}
		var p ledger.BeliefPayload
		if err := json.Unmarshal(e.Payload, &p); err != nil {
			return err
		}

		version := uint32(1)
		if existing, ok := be.Beliefs[p.BeliefID]; ok {
			version = uint32(len(existing) + 1)
		}

		be.Commit(&Belief{
			ID:         p.BeliefID,
			Version:    version,
			State:      BeliefActive,
			FactIDs:    p.FactIDs,
			Confidence: p.Confidence,
			Statement:  p.Statement,
		})
	}
	return nil
}

/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package memory implements the tiered storage and versioned fact substrate for PhoenixOS.
package memory

import (
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// FactState defines the lifecycle state of a verified fact node.
type FactState string

const (
	// StateActive indicates the fact is current and authoritative.
	StateActive FactState = "ACTIVE"
	// StateSuperseded indicates a newer version of this fact exists.
	StateSuperseded FactState = "SUPERSEDED"
	// StateInvalidated indicates the fact has been proven false by reality drift auditing.
	StateInvalidated FactState = "INVALIDATED"
	// StateArchived indicates the fact is stored for historical audit only.
	StateArchived FactState = "ARCHIVED"
)

// Fact represents the smallest unit of verified historical truth in PhoenixOS.
// It is immutable once stored; subsequent updates create new versioned instances.
type Fact struct {
	ID         string                 `json:"id"`
	Version    uint32                 `json:"version"`
	State      FactState              `json:"state"`
	Confidence ledger.ConfidenceScore `json:"confidence"`
	Timestamp  int64                  `json:"timestamp"` // Unix epoch
	Data       []byte                 `json:"data"`      // Raw telemetry payload
}

// Memory implements a multi-versioned key-value store for Facts.
// It ensures that historical world models can be reconstructed bit-for-bit.
type Memory struct {
	Facts map[string][]*Fact // FactID to chronological version slice mapping
}

// NewMemory initializes an empty fact store with an allocated map.
func NewMemory() *Memory {
	return &Memory{
		Facts: make(map[string][]*Fact),
	}
}

// Store adds a new version of a fact to memory and manages state transitions.
// Side Effects: Automatically transitions the previous latest version to SUPERSEDED.
// Complexity: O(1) time / O(1) space.
func (m *Memory) Store(f *Fact) {
	versions := m.Facts[f.ID]
	if len(versions) > 0 {
		versions[len(versions)-1].State = StateSuperseded
	}
	m.Facts[f.ID] = append(m.Facts[f.ID], f)
}

// Recall retrieves the most recent active version of a fact by ID.
// Returns (*Fact, true) if found, (nil, false) otherwise.
// Complexity: O(1) time / O(1) space.
func (m *Memory) Recall(id string) (*Fact, bool) {
	versions, ok := m.Facts[id]
	if !ok || len(versions) == 0 {
		return nil, false
	}
	return versions[len(versions)-1], true
}

// RecallVersion retrieves a specific historical version of a fact.
// Inputs: id (string), version (uint32).
// Returns (*Fact, true) if exact version found, (nil, false) otherwise.
// Complexity: O(V) where V is the number of versions for the ID.
func (m *Memory) RecallVersion(id string, version uint32) (*Fact, bool) {
	versions, ok := m.Facts[id]
	if !ok {
		return nil, false
	}
	for _, f := range versions {
		if f.Version == version {
			return f, true
		}
	}
	return nil, false
}

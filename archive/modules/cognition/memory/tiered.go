/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package memory implements the tiered storage and versioned fact substrate for PhoenixOS.
package memory

import (
	"fmt"
	"sync"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// Tier defines the priority and retention level of a memory segment.
type Tier string

const (
	// TierWorking is for volatile, high-frequency facts undergoing active evaluation.
	TierWorking Tier = "WORKING"
	// TierEpisodic is for chronological event sequences (The "What Happened" log).
	TierEpisodic Tier = "EPISODIC"
	// TierSemantic is for long-term knowledge and causal relationships (The "Why" log).
	TierSemantic Tier = "SEMANTIC"
	// TierProcedural is for learned reaction patterns and system policies.
	TierProcedural Tier = "PROCEDURAL"
)

// TieredMemory implements a hierarchical cognitive storage system.
// It facilitates the transition of facts from volatile short-term context to
// permanent, searchable persistence.
type TieredMemory struct {
	mu sync.RWMutex

	Working    map[string]*Fact
	Episodic   []*Fact
	Semantic   map[string]*Fact
	Procedural map[string]*Fact

	// Store is the optional persistence layer (SQLite-Vec).
	Store *VectorStore
}

// NewTieredMemory initializes a new tiered memory store with allocated segments.
func NewTieredMemory() *TieredMemory {
	return &TieredMemory{
		Working:    make(map[string]*Fact),
		Episodic:   make([]*Fact, 0),
		Semantic:   make(map[string]*Fact),
		Procedural: make(map[string]*Fact),
	}
}

// Ingest places a new fact into the volatile Working memory tier.
// Complexity: O(1) time / O(1) space.
func (tm *TieredMemory) Ingest(f *Fact) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.Working[f.ID] = f
}

// Consolidate migrates a fact from Working memory to a permanent tier.
// Inputs: id (Fact ID), target (Target Tier).
// Side Effects: Deletes from Working, appends/stores in target. Triggers async SQLite persistence if Store is attached.
// Complexity: O(1) time (amortized) / O(1) space.
func (tm *TieredMemory) Consolidate(id string, target Tier) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	f, ok := tm.Working[id]
	if !ok {
		return
	}

	delete(tm.Working, id)

	switch target {
	case TierEpisodic:
		tm.Episodic = append(tm.Episodic, f)
	case TierSemantic:
		tm.Semantic[id] = f
	case TierProcedural:
		tm.Procedural[id] = f
	}

	// Async persistence to SQLite substrate (Phase 4B).
	if tm.Store != nil {
		go func(fact *Fact) {
			if err := tm.Store.PersistFact(target, fact, nil); err != nil {
				fmt.Printf("[ERROR] Memory Persistence failed for %s: %v\n", fact.ID, err)
			}
		}(f)
	}
}

// Decay garbage-collects facts from Working memory that fall below a confidence threshold.
// Inputs: threshold (ledger.ConfidenceScore).
// Complexity: O(W) where W is the number of facts in Working memory.
func (tm *TieredMemory) Decay(threshold float64) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for id, f := range tm.Working {
		if phxmath.FixedPoint(f.Confidence).Float64() < threshold {
			delete(tm.Working, id)
		}
	}
}

// Search performs a prioritized lookup across all memory tiers (Working -> Semantic -> Procedural -> Episodic).
// Returns (*Fact, true) if found, (nil, false) otherwise.
// Complexity: O(1) for mapped tiers, O(E) for linear search of Episodic tier.
func (tm *TieredMemory) Search(id string) (*Fact, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if f, ok := tm.Working[id]; ok {
		return f, true
	}
	if f, ok := tm.Semantic[id]; ok {
		return f, true
	}
	if f, ok := tm.Procedural[id]; ok {
		return f, true
	}

	// Scan episodic (expensive, should use index in production)
	for _, f := range tm.Episodic {
		if f.ID == id {
			return f, true
		}
	}

	return nil, false
}

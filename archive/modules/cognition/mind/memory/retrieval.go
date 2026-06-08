/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 * [LABELS]: cognitive-memory, tiered-storage
 */

/*
 * FILE HEADER:
 * Purpose: Provides contextual grounding for the advisory intelligence layer using formal tiered memory.
 * Subsystem: Memory Retrieval
 * Dependencies: github.com/fallofpheonix/phoenix/foundation/runtime/bus, github.com/fallofpheonix/phoenix/cognition/memory
 * Dependents: Advisory Intelligence Layer, AIOrchestrator
 */

package memory

import (
	"fmt"

	rootMemory "github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

// MemoryBridge defines the contract for contextual grounding from episodic memory.
type MemoryBridge interface {
	RetrieveContext(event bus.TelemetryEvent) (string, error)
}

// CognitiveMemoryBridge implements the MemoryBridge interface using the formal TieredMemory.
type CognitiveMemoryBridge struct {
	Store *rootMemory.TieredMemory
}

func NewCognitiveMemoryBridge(store *rootMemory.TieredMemory) *CognitiveMemoryBridge {
	return &CognitiveMemoryBridge{
		Store: store,
	}
}

// RetrieveContext retrieves context from the formal tiered memory system.
func (m *CognitiveMemoryBridge) RetrieveContext(event bus.TelemetryEvent) (string, error) {
	// 1. Search Working Memory for immediate context
	id := fmt.Sprintf("evt-%d", event.SeqID)
	if fact, ok := m.Store.Search(id); ok {
		return fmt.Sprintf("Cognitive Context (ID: %s): Fact version %d, State: %s, Confidence: %.2f",
			fact.ID, fact.Version, fact.State, fact.Confidence), nil
	}

	// 2. Fallback logic for semantic search (to be expanded with HDF5/vector optimizations)
	return fmt.Sprintf("Cognitive Context: No formal fact recorded for %s yet. Awaiting consolidation.", event.EventType), nil
}

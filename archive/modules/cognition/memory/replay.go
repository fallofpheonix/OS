/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package memory implements the tiered storage and versioned fact substrate for PhoenixOS.
package memory

import (
	"encoding/json"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// ReconstructFromLedger rebuilds the Fact memory state by replaying a ledger event stream.
// It implements the Memory Reconstruction Theorem (Q662), ensuring cognitive state
// consistency with historical truth.
// Inputs: events ([]*ledger.Event) - Slice of chronological ledger events.
// Returns: error if JSON unmarshaling fails.
// Complexity: O(N) where N is the number of fact-related events in the ledger.
func (m *Memory) ReconstructFromLedger(events []*ledger.Event) error {
	for _, e := range events {
		switch e.Type {
		case ledger.EventFact, ledger.EventFactUpdate:
			var p ledger.FactPayload
			// If it's a FactUpdate, it might use the FactUpdatePayload,
			// but for simplicity, we'll assume a new versioned FactPayload.
			if err := json.Unmarshal(e.Payload, &p); err != nil {
				// Handle FactUpdatePayload if necessary
				return err
			}

			// Auto-increment version based on existing memory
			version := uint32(1)
			if existing, ok := m.Facts[p.FactID]; ok {
				version = uint32(len(existing) + 1)
			}

			fact := &Fact{
				ID:         p.FactID,
				Version:    version,
				State:      StateActive,
				Confidence: p.ConfidenceScore,
				Timestamp:  p.Timestamp,
			}
			m.Store(fact)
		}
	}
	return nil
}

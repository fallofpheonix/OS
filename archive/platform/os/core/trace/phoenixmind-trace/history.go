/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package trace

import (
	"time"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

// HistoryEntry records a state or event at a point in time for a specific entity.
type HistoryEntry struct {
	EntityID  string
	Timestamp time.Time
	Truth     evidence.TruthState
	Metadata  map[string]interface{}
}

// History acts as a persistent store for all historical states and events.
type History struct {
	Entries []HistoryEntry
}

// AddEntry adds a new historical record.
func (h *History) AddEntry(entityID string, state evidence.TruthState, metadata map[string]interface{}) {
	h.Entries = append(h.Entries, HistoryEntry{
		EntityID:  entityID,
		Timestamp: time.Now(),
		Truth:     state,
		Metadata:  metadata,
	})
}

// GetHistoryForEntity retrieves all historical entries for a specific entity.
func (h *History) GetHistoryForEntity(entityID string) []HistoryEntry {
	var entityHistory []HistoryEntry
	for _, entry := range h.Entries {
		if entry.EntityID == entityID {
			entityHistory = append(entityHistory, entry)
		}
	}
	return entityHistory
}

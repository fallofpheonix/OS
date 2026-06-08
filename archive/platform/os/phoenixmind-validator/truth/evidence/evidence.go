/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package evidence

import (
	"sync"
)

// EvidenceRegistry manages active system evidence.
type EvidenceRegistry struct {
	mu       sync.RWMutex
	registry map[string]*Evidence
}

// NewEvidenceRegistry creates a new registry.
func NewEvidenceRegistry() *EvidenceRegistry {
	return &EvidenceRegistry{
		registry: make(map[string]*Evidence),
	}
}

// Add adds a piece of evidence to the registry.
func (er *EvidenceRegistry) Add(e *Evidence) {
	er.mu.Lock()
	defer er.mu.Unlock()
	er.registry[e.EntityID] = e
}

// Get retrieves a piece of evidence by entity ID.
func (er *EvidenceRegistry) Get(entityID string) (*Evidence, bool) {
	er.mu.RLock()
	defer er.mu.RUnlock()
	e, ok := er.registry[entityID]
	return e, ok
}

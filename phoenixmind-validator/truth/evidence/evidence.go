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

package authority

import (
	"errors"
	"sync"
)

// CapabilityToken represents a scoped authority token.
type CapabilityToken struct {
	ID        string
	Issuer    string
	Subject   string
	Scope     []string
	TTL       int64
	Signature []byte
}

// Manager governs token lifecycle and validation.
type Manager struct {
	mu         sync.RWMutex
	tokens     map[string]*CapabilityToken
	revokedIDs map[string]bool
}

func NewManager() *Manager {
	return &Manager{
		tokens:     make(map[string]*CapabilityToken),
		revokedIDs: make(map[string]bool),
	}
}

// RegisterToken stores a valid capability token.
func (m *Manager) RegisterToken(t *CapabilityToken) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[t.ID] = t
}

// RevokeToken marks a token as invalid.
func (m *Manager) RevokeToken(tokenID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.revokedIDs[tokenID] = true
}

// Validate checks if a token is active and within scope.
func (m *Manager) Validate(tokenID string, requiredScope string) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.revokedIDs[tokenID] {
		return errors.New("token revoked")
	}

	t, ok := m.tokens[tokenID]
	if !ok {
		return errors.New("token unknown")
	}

	for _, s := range t.Scope {
		if s == requiredScope {
			return nil
		}
	}

	return errors.New("insufficient scope")
}

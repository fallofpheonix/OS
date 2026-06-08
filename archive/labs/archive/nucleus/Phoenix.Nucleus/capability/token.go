/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package capability

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
	"github.com/fallofpheonix/Phoenix.Nucleus/ledger"
)

// Token represents a bearer capability that manifests authority.
// Capabilities are NOT copied; they are attenuated (split) or transferred.
type Token struct {
	ID             string                  `json:"id"`
	ParentID       string                  `json:"parent_id"`
	AuthorityID    string                  `json:"authority_id"`
	Atoms          ledger.AuthorityAtom `json:"atoms"`
	Scope          []string                `json:"scope"`
	ExpiresAt      time.Time               `json:"expires_at"`
	Signature      string                  `json:"signature"` // Proof of issuance
}

// Attenuate creates a new child token from this token, reducing this token's authority.
func (t *Token) Attenuate(childID string, atoms ledger.AuthorityAtom, scope []string) (*Token, error) {
	if atoms > t.Atoms {
		return nil, fmt.Errorf("insufficient authority: requested %v, available %v", atoms, t.Atoms)
	}

	child := &Token{
		ID:          childID,
		ParentID:    t.ID,
		AuthorityID: t.AuthorityID,
		Atoms:       atoms,
		Scope:       scope,
		ExpiresAt:   t.ExpiresAt,
	}

	t.Atoms -= atoms
	child.Signature = child.CalculateSignature()
	return child, nil
}

// CalculateSignature generates a deterministic signature for the token.
func (t *Token) CalculateSignature() string {
	data, _ := json.Marshal(t)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

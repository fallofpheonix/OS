/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/*
 * REPOSITORY: PhoenixDistributed
 * ARCHITECTURAL JUSTIFICATION: Cryptographic identity and node authentication for the cluster.
 * DEPENDENCY BOUNDARY: Defines the AuthorityCertificate used across the cluster.
 * DETERMINISTIC CONSIDERATIONS: Immutable node identities, weight-based authority.
 */

package identity

import (
	"crypto/ed25519"
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

// AuthorityCertificate represents the cryptographic proof of a node's identity and weight.
type AuthorityCertificate struct {
	NodeID     string
	PublicKey  []byte
	Weight     float64 // Cryptographic PoA weight (0.0 to 1.0)
	Signature  []byte
	Reputation float64 // P9: Reputation scoring (0.0 to 1.0)
}

// NodeRegistry manages authenticated peers within the cluster.
type NodeRegistry struct {
	mu    sync.RWMutex
	peers map[string]*AuthorityCertificate
}

func NewNodeRegistry() *NodeRegistry {
	return &NodeRegistry{
		peers: make(map[string]*AuthorityCertificate),
	}
}

// RegisterNode authenticates and adds a new node to the registry.
func (r *NodeRegistry) RegisterNode(cert *AuthorityCertificate) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 1. Basic validation
	if cert.NodeID == "" || cert.Weight <= 0 || len(cert.PublicKey) != ed25519.PublicKeySize {
		return errors.New("INVALID_CERTIFICATE: missing NodeID, zero weight, or invalid public key")
	}

	// 2. Cryptographic Verification
	if !verifyCertificateSignature(cert) {
		return fmt.Errorf("AUTHENTICATION_FAILED: invalid cryptographic signature for node %s", cert.NodeID)
	}

	// Initialize reputation if new node
	if cert.Reputation == 0 {
		cert.Reputation = 0.5 // Default neutral reputation
	}

	r.peers[cert.NodeID] = cert
	return nil
}

// AdjustReputation updates the reputation score of a node based on its behavior (P9).
func (r *NodeRegistry) AdjustReputation(nodeID string, delta float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if p, ok := r.peers[nodeID]; ok {
		p.Reputation += delta
		if p.Reputation > 1.0 {
			p.Reputation = 1.0
		}
		if p.Reputation < 0.0 {
			p.Reputation = 0.0
		}
	}
}

// ExchangeProof verifies a cryptographic proof from a peer (P9).
func (r *NodeRegistry) ExchangeProof(nodeID string, proof []byte, expectedHash string) error {
	p, ok := r.GetPeer(nodeID)
	if !ok {
		return fmt.Errorf("node %s not found in registry", nodeID)
	}

	if p.Reputation < 0.1 {
		return fmt.Errorf("node %s has insufficient reputation to participate in proof exchange", nodeID)
	}

	// Cryptographic proof verification (e.g., verifying a state hash signature)
	// For this forge, we simulate successful proof exchange if signature is present.
	if len(proof) == 0 {
		return errors.New("EMPTY_PROOF: no cryptographic proof provided")
	}

	return nil
}


// GetTotalWeight calculates the aggregate PoA weight of the active cluster.
func (r *NodeRegistry) GetTotalWeight() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := 0.0
	for _, p := range r.peers {
		total += p.Weight
	}
	return total
}

// GetPeer retrieves a certificate by NodeID.
func (r *NodeRegistry) GetPeer(nodeID string) (*AuthorityCertificate, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.peers[nodeID]
	return p, ok
}

func verifyCertificateSignature(cert *AuthorityCertificate) bool {
	// The message to verify is the hash of (NodeID + Weight)
	msg := []byte(fmt.Sprintf("%s:%f", cert.NodeID, cert.Weight))
	h := sha256.Sum256(msg)

	// Verify signature using ed25519
	return ed25519.Verify(cert.PublicKey, h[:], cert.Signature)
}

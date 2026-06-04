package proofs

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/distributed/identity"
)

// STATUS: EXPERIMENTAL
// Proof 4: Federation (Multi-node legitimacy)
func TestFederationProof(t *testing.T) {
	registry := identity.NewNodeRegistry()

	// 1. Generate a mock node certificate
	pub, priv, _ := ed25519.GenerateKey(nil)
	nodeID := "node-beta"
	weight := 0.5
	msg := []byte(fmt.Sprintf("%s:%f", nodeID, weight))
	h := sha256.Sum256(msg)
	sig := ed25519.Sign(priv, h[:])

	cert := &identity.AuthorityCertificate{
		NodeID:    nodeID,
		PublicKey: pub,
		Weight:    weight,
		Signature: sig,
	}

	// 2. Admission: Register the node
	err := registry.RegisterNode(cert)
	if err != nil {
		t.Fatalf("Node admission failed: %v", err)
	}

	// 3. Proof Exchange: Verify a mock state proof from the node
	err = registry.ExchangeProof(nodeID, []byte("valid-state-signature"), "0xabcd")
	if err != nil {
		t.Fatalf("Proof exchange failed: %v", err)
	}

	// 4. Reputation & Revocation: Lower reputation and verify isolation
	registry.AdjustReputation(nodeID, -0.45) // Reputation drops from 0.5 to 0.05
	
	err = registry.ExchangeProof(nodeID, []byte("valid-state-signature"), "0xabcd")
	if err == nil {
		t.Error("Expected proof exchange to fail for low-reputation node")
	}
}

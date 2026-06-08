/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package identity

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestNodeRegistry_RegisterNode(t *testing.T) {
	registry := NewNodeRegistry()

	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	nodeID := "node-alpha"
	weight := 1.0
	msg := []byte(fmt.Sprintf("%s:%f", nodeID, weight))
	h := sha256.Sum256(msg)
	sig := ed25519.Sign(priv, h[:])

	cert := &AuthorityCertificate{
		NodeID:    nodeID,
		PublicKey: pub,
		Weight:    weight,
		Signature: sig,
	}

	if err := registry.RegisterNode(cert); err != nil {
		t.Fatalf("Failed to register valid node: %v", err)
	}

	if registry.GetTotalWeight() != 1.0 {
		t.Errorf("Expected total weight 1.0, got %.2f", registry.GetTotalWeight())
	}

	// Test duplicate/invalid
	invalidCert := &AuthorityCertificate{
		NodeID: "",
		Weight: 0,
	}
	if err := registry.RegisterNode(invalidCert); err == nil {
		t.Error("Allowed registration of invalid certificate")
	}
}

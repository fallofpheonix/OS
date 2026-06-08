/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	"testing"
)

func TestChain_Genesis(t *testing.T) {
	chain := NewChain()
	genesis := NewGenesisBlock()

	payload := []byte("genesis-payload")
	event := NewEvent(0, EventGenesis, payload, "", genesis.RootAuthority)

	if err := chain.Append(event); err != nil {
		t.Fatalf("failed to append genesis event: %v", err)
	}

	if chain.GetHead().Hash != event.Hash {
		t.Errorf("head hash mismatch")
	}
}

func TestChain_TamperDetection(t *testing.T) {
	chain := NewChain()
	genesis := NewGenesisBlock()
	event0 := NewEvent(0, EventGenesis, []byte("genesis"), "", genesis.RootAuthority)
	chain.Append(event0)

	event1 := NewEvent(1, EventAuthorityIssue, []byte("authority"), event0.Hash, genesis.RootAuthority)
	chain.Append(event1)

	// Tamper with event0
	event0.Payload = []byte("tampered")

	if err := chain.VerifyChain(); err == nil {
		t.Errorf("failed to detect tampering")
	}
}

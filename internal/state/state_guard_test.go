package state

import (
	"encoding/json"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestStateGuard_Authority(t *testing.T) {
	// SG-009: Mutation Invariant Verification

	g := NewStateGuard(1234)
	token := RequestAuthorityToken()

	// 1. Authorized Mutation
	ev1 := contracts.Event{
		Version: 1,
		Type:    contracts.EventSpawn,
	}
	payload, _ := json.Marshal(map[string]interface{}{"id": "agent_01", "pos": phxmath.NewFixedPoint(10)})
	ev1.Payload = payload

	applied1 := contracts.AppliedEvent{
		Height: 1,
		Epoch:  0,
		Event:  ev1,
	}

	err := g.Apply(token, applied1)
	if err != nil {
		t.Fatalf("Authorized Apply failed: %v", err)
	}

	hash1 := g.CalculateHash()
	if hash1 == (contracts.Hash{}) {
		t.Error("StateHash was not updated after Apply")
	}

	// 2. Verified Determinism
	g2 := NewStateGuard(1234)
	g2.Apply(token, applied1)
	if hash1 != g2.CalculateHash() {
		t.Error("Non-deterministic state hash across different guard instances")
	}
}

func TestStateGuard_RestoreIntegrity(t *testing.T) {
	// RESTORE-003: Proof of Hash Verification

	g := NewStateGuard(1234)
	token := RequestAuthorityToken()

	ev1 := contracts.Event{Version: 1, Type: contracts.EventSpawn}
	payload, _ := json.Marshal(map[string]interface{}{"id": "a", "pos": phxmath.NewFixedPoint(0)})
	ev1.Payload = payload

	applied1 := contracts.AppliedEvent{Height: 1, Epoch: 0, Event: ev1}
	if err := g.Apply(token, applied1); err != nil {
		t.Fatalf("Apply failed: %v", err)
	}

	originalHash := g.CalculateHash()
	data, _ := g.Snapshot()

	// 1. Valid Restore
	g2 := NewStateGuard(1234)
	err := g2.Restore(data, originalHash)
	if err != nil {
		t.Fatalf("Valid Restore failed: %v", err)
	}

	// 2. Corrupt Restore (Wrong Expected Hash)
	wrongHash := contracts.Hash{0xDE, 0xAD}
	err = g2.Restore(data, wrongHash)
	if err == nil {
		t.Error("Restore accepted a snapshot with a mismatched expected hash")
	}

	// 3. Corrupt Data
	data[len(data)-5] ^= 0xFF
	err = g2.Restore(data, originalHash)
	if err == nil {
		t.Error("Restore accepted corrupted JSON data")
	}
}

func TestStateGuard_ForensicAnchor(t *testing.T) {
	// SEQ-202: Sequence Participation in StateHash

	g := NewStateGuard(1234)
	h1 := g.CalculateHash()

	g.ApplyEnvelope("validator_a", 100)
	h2 := g.CalculateHash()

	if h1 == h2 {
		t.Error("ApplyEnvelope did not alter the StateHash (Forensic Anchor failure)")
	}

	if g.GetLastSequence("validator_a") != 100 {
		t.Error("GetLastSequence returned incorrect value")
	}
}

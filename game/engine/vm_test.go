package engine

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestVM_ProximityProbe(t *testing.T) {
	code := []Instruction{
		{Op: OpPush, Args: []interface{}{"agent_01"}},
		{Op: OpPush, Args: []interface{}{phxmath.NewFixedPoint(100)}},
		{Op: OpMove},
		{Op: OpPush, Args: []interface{}{"agent_01"}},
		{Op: OpVerify},
	}

	vm := NewVM(code)
	vm.Rules["agent_01"] = &ProximityRule{TargetPos: phxmath.NewFixedPoint(100), Threshold: phxmath.NewFixedPoint(10)}

	spawnEvent := contracts.Event{Version: 1, Type: contracts.EventSpawn}
	spawnEvent.Payload, _ = json.Marshal(map[string]interface{}{"id": "agent_01", "pos": phxmath.NewFixedPoint(0)})

	vm.State.ApplyEvent(contracts.AppliedEvent{Height: 1, Epoch: 0, Event: spawnEvent}, vm.Rules)

	err := vm.Run()
	if err != nil {
		t.Fatalf("VM run failed: %v", err)
	}

	entity, ok := vm.State.GetEntity("agent_01")
	if !ok || entity.Pos.V != phxmath.NewFixedPoint(100).V {
		t.Errorf("Expected pos %v, got %v", phxmath.NewFixedPoint(100).V, entity.Pos.V)
	}

	if entity.Status != "VERIFIED" {
		t.Errorf("Expected status VERIFIED, got %s", entity.Status)
	}
}

func TestVM_DeterminismInvariant(t *testing.T) {
	ws3 := NewWorldState(0)
	ws3.Tick = 1
	ws3.entities["z"] = &Entity{ID: "z", Pos: phxmath.NewFixedPoint(1)}
	ws3.entities["a"] = &Entity{ID: "a", Pos: phxmath.NewFixedPoint(2)}
	h3 := ws3.CalculateHash()

	ws4 := NewWorldState(0)
	ws4.Tick = 1
	ws4.entities["a"] = &Entity{ID: "a", Pos: phxmath.NewFixedPoint(2)}
	ws4.entities["z"] = &Entity{ID: "z", Pos: phxmath.NewFixedPoint(1)}
	h4 := ws4.CalculateHash()

	if h3 != h4 {
		t.Errorf("Map Iteration Order Determinism Broken!\n3: %s\n4: %s", h3, h4)
	}
}

func TestConsensus_MinimumValidators(t *testing.T) {
	pub, _, _ := ed25519.GenerateKey(nil)
	ws := NewWorldState(0)
	ws.Validators = [][]byte{pub}

	_, err := ws.CheckQuorum([]contracts.SignatureEntry{}, contracts.Hash{})
	if err == nil || err.Error() != "insufficient validators for BFT: 1 (min 4)" {
		t.Errorf("Expected insufficient validators error, got: %v", err)
	}
}

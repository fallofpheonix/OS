package replay

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestVM_ReplayInvariant(t *testing.T) {
	// 1. First Execution
	v1, spawns := setupWithKey()
	v1.Run()
	hash1 := v1.State.StateHash

	// 2. Replay from Events
	totalEvents := append(spawns, v1.Events...)
	applied := make([]contracts.AppliedEvent, len(totalEvents))
	for i, ev := range totalEvents {
		applied[i] = contracts.AppliedEvent{
			Height: uint64(i + 1),
			Epoch:  0,
			Event:  ev,
		}
	}

	// Start with identical genesis state
	genesis := engine.NewWorldState(0)
	genesis.Validators = [][]byte{v1.PrivateKey.Public().(ed25519.PublicKey)}

	v2_state, err := engine.Replay(genesis, applied, v1.Rules)
	if err != nil {
		t.Fatalf("Replay failed: %v", err)
	}
	hash2 := v2_state.StateHash

	if hash1 != hash2 {
		t.Logf("v1 Tick: %d, EventCount: %d, LastEventHash: %s", v1.State.Tick, v1.State.EventCount, v1.State.LastEventHash)
		t.Logf("v2 Tick: %d, EventCount: %d, LastEventHash: %s", v2_state.Tick, v2_state.EventCount, v2_state.LastEventHash)
		t.Errorf("Determinism Invariant Broken!\n1: %s\n2: %s", hash1, hash2)
	}
}

func setupWithKey() (*engine.VM, []contracts.Event) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	code := []engine.Instruction{
		{Op: engine.OpPush, Args: []interface{}{"agent_a"}},
		{Op: engine.OpPush, Args: []interface{}{phxmath.NewFixedPoint(10)}},
		{Op: engine.OpMove},
		{Op: engine.OpPush, Args: []interface{}{"agent_b"}},
		{Op: engine.OpPush, Args: []interface{}{phxmath.NewFixedPoint(20)}},
		{Op: engine.OpMove},
		{Op: engine.OpPush, Args: []interface{}{"agent_a"}},
		{Op: engine.OpVerify},
		{Op: engine.OpPush, Args: []interface{}{"agent_b"}},
		{Op: engine.OpVerify},
	}

	v := engine.NewVM(code)
	v.PrivateKey = priv
	v.Rules["agent_a"] = &engine.ProximityRule{TargetPos: phxmath.NewFixedPoint(10), Threshold: phxmath.NewFixedPoint(0)}
	v.Rules["agent_b"] = &engine.ProximityRule{TargetPos: phxmath.NewFixedPoint(20), Threshold: phxmath.NewFixedPoint(0)}

	ws := engine.NewWorldState(0)
	ws.Validators = [][]byte{pub}

	// Initial spawns
	e1 := contracts.Event{Version: 1, Type: contracts.EventSpawn}
	e1.Payload, _ = json.Marshal(map[string]interface{}{"id": "agent_a", "pos": 0})
	ws.ApplyEvent(contracts.AppliedEvent{Height: 1, Epoch: 0, Event: e1}, v.Rules)

	e2 := contracts.Event{Version: 1, Type: contracts.EventSpawn}
	e2.Payload, _ = json.Marshal(map[string]interface{}{"id": "agent_b", "pos": 0})
	ws.ApplyEvent(contracts.AppliedEvent{Height: 2, Epoch: 0, Event: e2}, v.Rules)

	v.State = *ws
	return v, []contracts.Event{e1, e2}
}

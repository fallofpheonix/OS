package replay

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestReplay_CanonicalProof(t *testing.T) {
	vm := engine.NewVM([]engine.Instruction{})

	spawnEvent := contracts.Event{
		Version: 1,
		Type:    contracts.EventSpawn,
	}
	spawnEvent.Payload, _ = json.Marshal(map[string]interface{}{"id": "agent_x", "pos": phxmath.NewFixedPoint(0)})

	applied := contracts.AppliedEvent{
		Height: 1,
		Epoch:  0,
		Event:  spawnEvent,
	}

	err := vm.State.ApplyEvent(applied, vm.Rules)
	if err != nil {
		t.Fatalf("ApplyEvent failed: %v", err)
	}

	stateHashA := vm.State.StateHash

	// Replay
	replayedState, err := engine.Replay(nil, []contracts.AppliedEvent{applied}, vm.Rules)
	if err != nil {
		t.Fatalf("Replay failed: %v", err)
	}

	if stateHashA != replayedState.StateHash {
		t.Errorf("Replay DIVERGENCE!\nOriginal: %s\nReplayed: %s", stateHashA, replayedState.StateHash)
	}
}

func TestReplay_Randomized(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	entities := []string{"alpha", "beta", "gamma", "delta"}
	rules := make(map[string]engine.VerificationRule)
	for _, id := range entities {
		rules[id] = &engine.ProximityRule{
			TargetPos: phxmath.NewFixedPoint(int64(r.Intn(100))),
			Threshold: phxmath.NewFixedPoint(10),
		}
	}

	var appliedEvents []contracts.AppliedEvent
	tick := uint64(1)

	// Spawns
	for _, id := range entities {
		e := contracts.Event{Version: 1, Type: contracts.EventSpawn}
		e.Payload, _ = json.Marshal(map[string]interface{}{"id": id, "pos": phxmath.NewFixedPoint(0)})

		a := contracts.AppliedEvent{Height: tick, Epoch: 0, Event: e}
		appliedEvents = append(appliedEvents, a)
		tick++
	}

	// 100 random events
	for i := 0; i < 100; i++ {
		entityID := entities[r.Intn(len(entities))]
		var eType contracts.EventType
		var payload []byte

		switch r.Intn(2) {
		case 0:
			eType = contracts.EventMove
			pos := phxmath.NewFixedPoint(int64(r.Intn(100)))
			payload, _ = json.Marshal(map[string]interface{}{"id": entityID, "pos": pos})
		case 1:
			eType = contracts.EventVerify
			payload, _ = json.Marshal(map[string]interface{}{"id": entityID})
		}

		e := contracts.Event{Version: 1, Type: eType, Payload: payload}
		a := contracts.AppliedEvent{Height: tick, Epoch: 0, Event: e}
		appliedEvents = append(appliedEvents, a)
		tick++
	}

	// Execution 1
	state1, _ := engine.Replay(nil, appliedEvents, rules)
	// Execution 2
	state2, _ := engine.Replay(nil, appliedEvents, rules)

	if state1.StateHash != state2.StateHash {
		t.Errorf("Randomized Replay DIVERGENCE!\n1: %s\n2: %s", state1.StateHash, state2.StateHash)
	}
}

func TestValidator_Lifecycle(t *testing.T) {
	_, priv1, _ := ed25519.GenerateKey(nil)
	_, priv2, _ := ed25519.GenerateKey(nil)
	pub1 := priv1.Public().(ed25519.PublicKey)
	pub2 := priv2.Public().(ed25519.PublicKey)

	ws := engine.NewWorldState(0)
	ws.Validators = [][]byte{pub1}
	rules := make(map[string]engine.VerificationRule)

	// 1. Authorized event (Add pub2)
	payload, _ := json.Marshal(map[string]interface{}{"id": hex.EncodeToString(pub2), "status": "ADD"})
	e1 := contracts.Event{
		Version: 1,
		Type:    contracts.EventUpdateValidator,
		Payload: payload,
	}

	applied1 := contracts.AppliedEvent{Height: 1, Epoch: 0, Event: e1}
	if err := ws.ApplyEvent(applied1, rules); err != nil {
		t.Fatalf("ApplyEvent failed: %v", err)
	}

	if len(ws.Validators) != 1 {
		t.Errorf("Validator set should not change until epoch boundary")
	}

	// 2. Trigger Epoch Transition (Move to Tick 101)
	e_epoch := contracts.Event{
		Version: 1,
		Type:    contracts.EventMove,
	}
	e_epoch.Payload, _ = json.Marshal(map[string]interface{}{"id": "none", "pos": 0})

	applied2 := contracts.AppliedEvent{Height: 101, Epoch: 1, Event: e_epoch}
	ws.ApplyEvent(applied2, rules)

	if len(ws.Validators) != 2 {
		t.Errorf("Expected 2 validators after epoch transition, got %d", len(ws.Validators))
	}
}

func TestState_SnapshotRestore(t *testing.T) {
	ws := engine.NewWorldState(0)
	ws.Tick = 10
	engine.InjectEntity(ws, &engine.Entity{ID: "a", Pos: phxmath.NewFixedPoint(100), Status: "ACTIVE"})
	ws.StateHash = ws.CalculateHash()

	snapshot, err := ws.Snapshot()
	if err != nil {
		t.Fatalf("Snapshot failed: %v", err)
	}

	ws2 := engine.NewWorldState(0)
	err = ws2.Restore(snapshot)
	if err != nil {
		t.Fatalf("Restore failed: %v", err)
	}

	if ws2.Tick != ws.Tick || ws2.StateHash != ws.StateHash {
		t.Error("Restore failed to reproduce state")
	}
}

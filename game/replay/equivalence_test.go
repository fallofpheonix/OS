package replay

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestSnapshotReplayEquivalence(t *testing.T) {
	// 1. Setup genesis state and events
	pub, priv, _ := ed25519.GenerateKey(nil)
	ws := engine.NewWorldState(1234)
	ws.Validators = [][]byte{pub}

	replayer := NewReplayEngine(ws)
	replayer.AddAuthorizedValidator(pub)

	events := generateTestEvents(priv, 10)

	// 2. Full replay from genesis
	for i, ev := range events {
		payload, _ := json.Marshal(ev)
		env := &contracts.SignedEnvelope{
			Version:   1,
			Type:      0, // MsgEvent
			Payload:   payload,
			Sequence:  uint64(i) + 1,
			Timestamp: 123456789,
			Validator: contracts.NodeID(pub),
		}
		consensus.SignEnvelope(env, priv)

		if err := replayer.ProcessEnvelope(env); err != nil {
			t.Fatalf("Failed to process event %d: %v", i, err)
		}
	}
	finalGenesisHash := ws.CalculateHash()

	// 3. Take snapshot at event 5
	wsMid := engine.NewWorldState(1234)
	wsMid.Validators = [][]byte{pub}
	replayerMid := NewReplayEngine(wsMid)
	replayerMid.AddAuthorizedValidator(pub)

	for i := 0; i < 5; i++ {
		ev := events[i]
		payload, _ := json.Marshal(ev)
		env := &contracts.SignedEnvelope{
			Version:   1,
			Type:      0, // MsgEvent
			Payload:   payload,
			Sequence:  uint64(i) + 1,
			Timestamp: 123456789,
			Validator: contracts.NodeID(pub),
		}
		consensus.SignEnvelope(env, priv)
		replayerMid.ProcessEnvelope(env)
	}

	snapshot, err := wsMid.Snapshot()
	if err != nil {
		t.Fatalf("Failed to take snapshot: %v", err)
	}

	// 4. Restore from snapshot and replay remaining events
	wsRestored := engine.NewWorldState(0)
	if err := wsRestored.Restore(snapshot); err != nil {
		t.Fatalf("Failed to restore snapshot: %v", err)
	}

	replayerRestored := NewReplayEngine(wsRestored)
	replayerRestored.AddAuthorizedValidator(pub)

	for i := 5; i < 10; i++ {
		ev := events[i]
		payload, _ := json.Marshal(ev)
		env := &contracts.SignedEnvelope{
			Version:   1,
			Type:      0, // MsgEvent
			Payload:   payload,
			Sequence:  uint64(i) + 1,
			Timestamp: 123456789,
			Validator: contracts.NodeID(pub),
		}
		consensus.SignEnvelope(env, priv)

		if err := replayerRestored.ProcessEnvelope(env); err != nil {
			t.Fatalf("Failed to process event %d after restore: %v", i, err)
		}
	}
	finalRestoredHash := wsRestored.CalculateHash()

	// 5. Assert Equivalence
	if finalGenesisHash != finalRestoredHash {
		t.Errorf("Equivalence Failure!\nGenesis Hash:  %s\nRestored Hash: %s",
			finalGenesisHash, finalRestoredHash)
	} else {
		t.Logf("Snapshot + Replay Equivalence PROVEN: %s", finalGenesisHash)
	}
}

func generateTestEvents(priv ed25519.PrivateKey, count int) []contracts.Event {
	events := make([]contracts.Event, count)

	for i := 0; i < count; i++ {
		ev := contracts.Event{
			Version: 1,
			Type:    contracts.EventSpawn,
		}
		ev.Payload, _ = json.Marshal(map[string]interface{}{"id": "test", "pos": 0})
		events[i] = ev
	}
	return events
}

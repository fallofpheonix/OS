package runtime

import (
	"encoding/json"
	"testing"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
	"github.com/fallofpheonix/phoenix/internal/state"
)

func TestDispatcher_ProcessBlock(t *testing.T) {
	// 1. Setup StateGuard
	guard := state.NewStateGuard(1234)
	token := state.RequestAuthorityToken()
	dispatcher := NewDispatcher(guard, token)

	// 2. Construct a valid block
	ev1 := contracts.Event{
		Version: 1,
		Type:    contracts.EventSpawn,
	}
	ev1.Payload, _ = json.Marshal(map[string]interface{}{"id": "agent_0", "pos": phxmath.NewFixedPoint(0)})

	events := []contracts.Event{ev1}

	// Pre-calculate StateRoot by applying manually
	tempGuard := state.NewStateGuard(1234)
	tempGuard.Apply(token, contracts.AppliedEvent{Height: 1, Epoch: 0, Event: ev1})
	expectedStateRoot := tempGuard.CalculateHash()

	merkleRoot, _ := protocol.CalculateMerkleRoot(events)

	block := contracts.FinalizedBlock{
		Version:    1,
		Height:     1,
		Epoch:      0,
		Round:      1,
		MerkleRoot: merkleRoot,
		Events:     events,
		StateRoot:  expectedStateRoot,
	}

	// 3. Process Block
	if err := dispatcher.ProcessBlock(block); err != nil {
		t.Fatalf("Dispatcher failed to process valid block: %v", err)
	}

	// 4. Verify Side Effects
	if guard.CalculateHash() != expectedStateRoot {
		t.Errorf("StateRoot mismatch after dispatch")
	}
}

func TestDispatcher_StateDivergence(t *testing.T) {
	guard := state.NewStateGuard(1234)
	token := state.RequestAuthorityToken()
	dispatcher := NewDispatcher(guard, token)

	ev1 := contracts.Event{Version: 1, Type: contracts.EventSpawn}
	ev1.Payload, _ = json.Marshal(map[string]interface{}{"id": "a", "pos": 0})

	merkleRoot, _ := protocol.CalculateMerkleRoot([]contracts.Event{ev1})

	block := contracts.FinalizedBlock{
		Version:    1,
		Height:     1,
		MerkleRoot: merkleRoot,
		Events:     []contracts.Event{ev1},
		StateRoot:  contracts.Hash{0xDE, 0xAD}, // MALICIOUS / WRONG
	}

	err := dispatcher.ProcessBlock(block)
	if err == nil {
		t.Error("Expected error due to state divergence, got nil")
	}
}

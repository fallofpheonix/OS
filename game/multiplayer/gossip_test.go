package multiplayer

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestGossipPropagation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Create Node A
	nodeA, err := NewGossipNode(ctx, 0)
	if err != nil {
		t.Fatalf("Failed to create node A: %v", err)
	}
	defer nodeA.Host.Close()

	// 2. Create Node B
	nodeB, err := NewGossipNode(ctx, 0)
	if err != nil {
		t.Fatalf("Failed to create node B: %v", err)
	}
	defer nodeB.Host.Close()

	// 3. Connect Node A to Node B
	addrA := nodeA.Host.Addrs()[0].String() + "/p2p/" + nodeA.Host.ID().String()
	err = nodeB.AddPeer(ctx, addrA)
	if err != nil {
		t.Fatalf("Failed to connect B to A: %v", err)
	}

	// Wait for connection to stabilize
	time.Sleep(1 * time.Second)

	// 4. Node A broadcasts an envelope
	event := contracts.Event{
		Version: 1,
		Type:    contracts.EventMove,
		Payload: []byte("TEST_GOSSIP"),
	}

	_, privA, _ := ed25519.GenerateKey(nil)

	payload, _ := json.Marshal(event)
	env := consensus.SignedEnvelope{
		Version: 1,
		Type:    consensus.MsgEvent,
		Payload: payload,
	}
	if err := consensus.SignEnvelope(&env, privA); err != nil {
		t.Fatalf("Failed to sign envelope: %v", err)
	}

	err = nodeA.Broadcast(ctx, env)
	if err != nil {
		t.Fatalf("Failed to broadcast: %v", err)
	}

	// 5. Node B should receive the envelope
	select {
	case receivedEnv := <-nodeB.Envelopes:
		if !consensus.VerifyEnvelope(&receivedEnv) {
			t.Error("Envelope signature verification failed")
		}

		var receivedEvent contracts.Event
		if err := json.Unmarshal(receivedEnv.Payload, &receivedEvent); err != nil {
			t.Fatalf("Failed to unmarshal event from envelope: %v", err)
		}

		if string(receivedEvent.Payload) != "TEST_GOSSIP" {
			t.Errorf("Mismatch in received event! Expected TEST_GOSSIP, got %s", string(receivedEvent.Payload))
		}
	case <-ctx.Done():
		t.Error("Timed out waiting for envelope propagation")
	}
}

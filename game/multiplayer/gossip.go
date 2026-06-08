package multiplayer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fallofpheonix/phoenix/internal/consensus"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// GossipNode handles P2P propagation of signed messages.
type GossipNode struct {
	Host      host.Host
	PubSub    *pubsub.PubSub
	Topic     *pubsub.Topic
	Sub       *pubsub.Subscription
	Envelopes chan consensus.SignedEnvelope
}

// NewGossipNode creates a new libp2p gossip node.
func NewGossipNode(ctx context.Context, listenPort int) (*GossipNode, error) {
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
	)
	if err != nil {
		return nil, err
	}

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, err
	}

	topic, err := ps.Join("phoenix-state-gossip")
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	return &GossipNode{
		Host:      h,
		PubSub:    ps,
		Topic:     topic,
		Sub:       sub,
		Envelopes: make(chan consensus.SignedEnvelope, 100),
	}, nil
}

// Broadcast sends an envelope to the P2P network.
func (gn *GossipNode) Broadcast(ctx context.Context, env consensus.SignedEnvelope) error {
	data, err := json.Marshal(env)
	if err != nil {
		return err
	}
	return gn.Topic.Publish(ctx, data)
}

// Listen starts a loop that receives envelopes from the P2P network.
func (gn *GossipNode) Listen(ctx context.Context) {
	for {
		msg, err := gn.Sub.Next(ctx)
		if err != nil {
			close(gn.Envelopes)
			return
		}
		// Only process messages from other peers
		if msg.ReceivedFrom == gn.Host.ID() {
			continue
		}

		var env consensus.SignedEnvelope
		if err := json.Unmarshal(msg.Data, &env); err != nil {
			fmt.Printf("Gossip: failed to unmarshal envelope: %v\n", err)
			continue
		}

		// CONSENSUS-014: Mandatory Envelope Verification
		if !consensus.VerifyEnvelope(&env) {
			fmt.Printf("Gossip: rejected envelope with invalid signature from %s\n", msg.ReceivedFrom)
			continue
		}

		gn.Envelopes <- env
	}
}

// AddPeer connects to another node.
func (gn *GossipNode) AddPeer(ctx context.Context, addr string) error {
	info, err := peer.AddrInfoFromString(addr)
	if err != nil {
		return err
	}
	return gn.Host.Connect(ctx, *info)
}

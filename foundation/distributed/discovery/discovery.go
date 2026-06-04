/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package discovery

import (
	"context"
)

// Peer represents a discovered node in the Phoenix network.
type Peer struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	PublicKey []byte `json:"public_key"`
	Status    string `json:"status"`
}

// PeerDiscovery defines the interface for P2P node discovery (The Phoenix Beacon).
type PeerDiscovery interface {
	// Start begins the discovery process (mDNS, UDP Broadcast, or DHT).
	Start(ctx context.Context) error

	// Stop halts the discovery process.
	Stop() error

	// Peers returns the current list of verified peers in the network.
	Peers() []Peer

	// Watch returns a channel that signals changes in peer membership.
	Watch(ctx context.Context) (<-chan []Peer, error)

	// Register advertises the local node to the network.
	Register(ctx context.Context, localPeer Peer) error
}

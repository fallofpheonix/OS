/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package discovery

import (
	"sync"
	"time"
)

// PeerRegistry maintains the active list of verified nodes and their metadata.
type PeerRegistry struct {
	peers map[string]*PeerInfo
	mu    sync.RWMutex
}

// PeerInfo contains detailed information about a cluster participant.
type PeerInfo struct {
	Identity      string
	Endpoints     []string
	Reputation    float64
	LastSeen      time.Time
	Authenticated bool
}

// NewPeerRegistry creates a new instance of the PeerRegistry.
func NewPeerRegistry() *PeerRegistry {
	return &PeerRegistry{
		peers: make(map[string]*PeerInfo),
	}
}

// Register adds or updates a peer in the registry.
func (r *PeerRegistry) Register(peer Peer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if info, exists := r.peers[peer.ID]; exists {
		info.LastSeen = time.Now()

		found := false
		for _, addr := range info.Endpoints {
			if addr == peer.Address {
				found = true
				break
			}
		}
		if !found {
			info.Endpoints = append(info.Endpoints, peer.Address)
		}
	} else {
		r.peers[peer.ID] = &PeerInfo{
			Identity:      peer.ID,
			Endpoints:     []string{peer.Address},
			Reputation:    0.5, // Default reputation
			LastSeen:      time.Now(),
			Authenticated: false,
		}
	}
}

// GetAuthenticatedPeers returns a list of peers that have passed cryptographic verification.
func (r *PeerRegistry) GetAuthenticatedPeers() []*PeerInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var authenticated []*PeerInfo
	for _, info := range r.peers {
		if info.Authenticated {
			authenticated = append(authenticated, info)
		}
	}
	return authenticated
}

// AdjustReputation modifies a peer's reputation score.
func (r *PeerRegistry) AdjustReputation(identity string, delta float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if info, exists := r.peers[identity]; exists {
		info.Reputation += delta
		if info.Reputation < 0 {
			info.Reputation = 0
		}
		if info.Reputation > 1.0 {
			info.Reputation = 1.0
		}
	}
}

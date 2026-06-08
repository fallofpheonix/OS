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
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// BeaconTransport implements PeerDiscovery using mDNS-like UDP multicasting.
type BeaconTransport struct {
	port     int
	identity string
	peers    map[string]Peer
	mu       sync.RWMutex
	cancel   context.CancelFunc
	watchers []chan []Peer
}

// NewBeaconTransport creates a new UDP multicast discovery transport.
func NewBeaconTransport(port int, identity string) *BeaconTransport {
	return &BeaconTransport{
		port:     port,
		identity: identity,
		peers:    make(map[string]Peer),
	}
}

func (t *BeaconTransport) Start(ctx context.Context) error {
	ctx, t.cancel = context.WithCancel(ctx)

	// Start listener
	go t.listen(ctx)

	// Start broadcaster
	go t.broadcast(ctx)

	return nil
}

func (t *BeaconTransport) Stop() error {
	if t.cancel != nil {
		t.cancel()
	}
	return nil
}

func (t *BeaconTransport) Peers() []Peer {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var res []Peer
	for _, p := range t.peers {
		res = append(res, p)
	}
	return res
}

func (t *BeaconTransport) Watch(ctx context.Context) (<-chan []Peer, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	ch := make(chan []Peer, 1)
	t.watchers = append(t.watchers, ch)
	return ch, nil
}

func (t *BeaconTransport) Register(ctx context.Context, localPeer Peer) error {
	// In Beacon mode, we just update our identity
	t.mu.Lock()
	t.identity = localPeer.ID
	t.mu.Unlock()
	return nil
}

func (t *BeaconTransport) listen(ctx context.Context) {
	addr := &net.UDPAddr{
		Port: t.port,
		IP:   net.ParseIP("224.0.0.1"), // Standard Local Multicast
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Printf("[Beacon Listener] Error: %v", err)
		return
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("[Beacon Listener] Error closing connection: %v", closeErr)
		}
	}()

	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
				log.Printf("[Beacon Listener] Error setting read deadline: %v", err)
				return
			}
			n, src, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}

			payload := string(buf[:n])
			// Simple Beacon Format: "PHOENIX_BEACON:<identity>"
			if len(payload) > 15 && payload[:15] == "PHOENIX_BEACON:" {
				peerIdentity := payload[15:]
				if peerIdentity != t.identity {
					t.mu.Lock()
					if _, exists := t.peers[peerIdentity]; !exists {
						t.peers[peerIdentity] = Peer{
							ID:      peerIdentity,
							Address: src.String(),
							Status:  "discovered",
						}
						t.notifyWatchers()
					}
					t.mu.Unlock()
				}
			}
		}
	}
}

func (t *BeaconTransport) notifyWatchers() {
	var res []Peer
	for _, p := range t.peers {
		res = append(res, p)
	}
	for _, ch := range t.watchers {
		select {
		case ch <- res:
		default:
		}
	}
}

func (t *BeaconTransport) broadcast(ctx context.Context) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("224.0.0.1:%d", t.port))
	if err != nil {
		log.Printf("[Beacon Broadcaster] Error: %v", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Printf("[Beacon Broadcaster] Error: %v", err)
		return
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Printf("[Beacon Broadcaster] Error closing connection: %v", closeErr)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		t.mu.RLock()
		beacon := []byte(fmt.Sprintf("PHOENIX_BEACON:%s", t.identity))
		t.mu.RUnlock()

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, err := conn.Write(beacon)
			if err != nil {
				log.Printf("[Beacon Broadcaster] Send Error: %v", err)
			}
		}
	}
}

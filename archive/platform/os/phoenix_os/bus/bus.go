/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: bus.go
 *
 * Purpose:
 * High-performance, thread-safe, and cryptographically hardened event bus. 
 * Acts as the nervous system of PhoenixOS, routing telemetry between subsystems.
 *
 * Subsystem:
 * Nucleus-Bus (Terminus implementation)
 *
 * Dependencies:
 * - standard go crypto (ed25519, sha256)
 *
 * Security:
 * - [CRITICAL]: Cryptographic event authentication (ed25519).
 * - [AUDIT]: Causal hashing (prev_hash) for immutable event chains.
 *
 * Performance:
 * - [BOUNDED]: Queue capacity limits memory explosion (PERF-003).
 * - [CONCURRENT]: RWMutex and atomic operations for high-throughput.
 * - [ADAPTIVE]: Load-shedding based on severity and watermarks.
 *
 * @labels bus, telemetry, message-queue, crypto, phase-2-complete
 */
package bus

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	QueueCapacity     = 4096 // Bounded to prevent memory explosion (Ref: PERF-003)
	HighWatermark     = 0.85
	CriticalWatermark = 0.95
)

/*
 * @struct TelemetryEvent
 * @description The fundamental unit of data in PhoenixOS. Matching EVENT_SCHEMA.md.
 */
type TelemetryEvent struct {
	SeqID        int64           `json:"seq_id"`
	LamportClock uint64          `json:"lamport_clock"`
	LogicalTick  uint64          `json:"logical_tick"`
	MonotonicNs  int64           `json:"monotonic_ns"`
	WallTimeUnix int64           `json:"wall_time_unix"`
	Source       string          `json:"source"`
	HostID       string          `json:"host_id"`
	PID          int             `json:"pid"`
	TID          int             `json:"tid"`
	UID          int             `json:"uid"`
	GID          int             `json:"gid"`
	EventType    string          `json:"event_type"`
	Severity     float64         `json:"severity"`
	Payload      json.RawMessage `json:"payload"`
	PrevHash     string          `json:"prev_hash"`
	Hash         string          `json:"hash"`

	// Cryptographic Hardening
	Signature []byte `json:"signature"`
	PublicKey []byte `json:"public_key"`

	// Hardening Fields
	EventID     string `json:"event_id"`
	CausalID    string `json:"causal_id"`
	SequenceNo  int64  `json:"sequence_no"`
	SourceEpoch int64  `json:"source_epoch"`
	Tgid        uint32 `json:"tgid"`
	Nsproxy     uint32 `json:"nsproxy"`
}

/**
 * Authenticate verifies the cryptographic signature of the event.
 * @return bool True if signature is valid.
 * @complexity O(1) - Constant time crypto verification.
 */
func (e *TelemetryEvent) Authenticate() bool {
	if len(e.Signature) == 0 || len(e.PublicKey) != ed25519.PublicKeySize {
		return false
	}
	h := sha256.Sum256([]byte(e.Hash))
	return ed25519.Verify(e.PublicKey, h[:], e.Signature)
}

type topicState struct {
	mu           sync.RWMutex
	subscribers  []chan TelemetryEvent
	overflowSeen bool
	lastHash     string
}

/*
 * @class Bus
 * @description Thread-safe pub/sub engine with load shedding and causal hashing.
 */
type Bus struct {
	topicsMu     sync.RWMutex
	topics       map[string]*topicState
	
	Dropped      int64 // Atomic
	Sampled      int64 // Atomic
	
	OnOverflow   func(topic string, pressure float64, event TelemetryEvent)
	lamportClock uint64 // Atomic
}

/**
 * NewBus creates and initializes a new event bus.
 * @return *Bus
 */
func NewBus() *Bus {
	return &Bus{
		topics: make(map[string]*topicState),
	}
}

/**
 * getTopicState retrieves or initializes a topic state.
 * @param topic The topic name.
 * @return *topicState
 */
func (b *Bus) getTopicState(topic string) *topicState {
	b.topicsMu.RLock()
	ts, ok := b.topics[topic]
	b.topicsMu.RUnlock()
	if ok {
		return ts
	}

	b.topicsMu.Lock()
	defer b.topicsMu.Unlock()
	if ts, ok = b.topics[topic]; ok {
		return ts
	}
	ts = &topicState{
		subscribers: make([]chan TelemetryEvent, 0),
	}
	b.topics[topic] = ts
	return ts
}

/**
 * Subscribe creates a new channel for receiving events from a topic.
 * @param topic The topic name.
 * @return chan TelemetryEvent
 */
func (b *Bus) Subscribe(topic string) chan TelemetryEvent {
	ts := b.getTopicState(topic)
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	ch := make(chan TelemetryEvent, QueueCapacity)
	ts.subscribers = append(ts.subscribers, ch)
	return ch
}

/**
 * Publish broadcasts an event to all topic subscribers.
 * @param topic The topic name.
 * @param event The event to broadcast.
 * @complexity O(S) where S is the number of subscribers.
 */
func (b *Bus) Publish(topic string, event TelemetryEvent) {
	// Increment Global Logical Clock atomically
	event.LamportClock = atomic.AddUint64(&b.lamportClock, 1)

	// Pre-calculate payload hash outside critical section
	pHash := sha256.Sum256(event.Payload)

	ts := b.getTopicState(topic)
	
	ts.mu.Lock()
	event.PrevHash = ts.lastHash
	
	// Minimal hashing inside critical section
	h := sha256.New()
	h.Write([]byte(event.EventID))
	h.Write([]byte(event.PrevHash))
	h.Write([]byte(strconv.FormatUint(event.LamportClock, 10)))
	h.Write(pHash[:])
	event.Hash = hex.EncodeToString(h.Sum(nil))
	
	ts.lastHash = event.Hash
	
	// Snapshot subscribers to allow concurrent mutations (Subscribe)
	// without holding the topic lock during delivery.
	subs := make([]chan TelemetryEvent, len(ts.subscribers))
	copy(subs, ts.subscribers)
	ts.mu.Unlock()

	for _, ch := range subs {
		fillRatio := float64(len(ch)) / float64(QueueCapacity)

		if fillRatio >= HighWatermark {
			if event.Severity < 0.8 {
				atomic.AddInt64(&b.Dropped, 1)
				continue
			}
		}

		if fillRatio >= CriticalWatermark {
			ts.mu.Lock()
			if !ts.overflowSeen && b.OnOverflow != nil {
				ts.overflowSeen = true
				b.OnOverflow(topic, fillRatio, event)
			}
			ts.mu.Unlock()

			if event.Severity < 0.9 {
				atomic.AddInt64(&b.Dropped, 1)
				continue
			}
		} else {
			ts.mu.Lock()
			ts.overflowSeen = false
			ts.mu.Unlock()
		}

		select {
		case ch <- event:
		default:
			if event.Severity >= 0.8 {
				select {
				case <-ch:
					ch <- event
				default:
					atomic.AddInt64(&b.Dropped, 1)
				}
			} else {
				atomic.AddInt64(&b.Dropped, 1)
			}
		}
	}
}

/**
 * QueuePressure returns the highest fill ratio among all topic subscribers.
 * @param topic The topic name.
 * @return float64 The max fill ratio (0.0 to 1.0).
 */
func (b *Bus) QueuePressure(topic string) float64 {
	b.topicsMu.RLock()
	ts, ok := b.topics[topic]
	b.topicsMu.RUnlock()

	if !ok {
		return 0.0
	}

	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	if len(ts.subscribers) == 0 {
		return 0.0
	}

	var maxPressure float64
	for _, ch := range ts.subscribers {
		pressure := float64(len(ch)) / float64(QueueCapacity)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}
	return maxPressure
}


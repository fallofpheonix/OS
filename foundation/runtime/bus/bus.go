// Package bus provides the central event distribution backbone for PhoenixOS.
// Core Domain Logic: Implements a high-performance, topic-based pub/sub system with 
// adaptive overflow protection and priority-based shedding to ensure system-wide telemetry delivery.
package bus

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
)

const (
	QueueCapacity     = 65536
	HighWatermark     = 0.85
	CriticalWatermark = 0.95
)

// TelemetryEvent represents the canonical unit of data flowing through the bus.
// Internal State: Encapsulates temporal, identity, and causal metadata alongside a generic payload.
// API Scope: Public; primary data structure for system-wide observability.
// Concurrency: Thread-safe (read-only after creation).
type TelemetryEvent struct {
	SeqID        int64           `json:"seq_id"`
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

	// Legacy & Compatibility Fields
	LamportClock uint64 `json:"lamport_clock"`
	Nsproxy      uint32 `json:"nsproxy"`
	Tgid         uint32 `json:"tgid"`

	// Hardening Fields
	EventID     string `json:"event_id"`
	CausalID    string `json:"causal_id"`
	SequenceNo  int64  `json:"sequence_no"`
	SourceEpoch int64  `json:"source_epoch"`
}

// Bus is the central nervous system for event distribution.
// Internal State: Manages a map of subscribers and tracks dropped events and overflow conditions.
// API Scope: Public; critical substrate for all inter-component communication.
// Concurrency: Thread-safe via sync.RWMutex.
type Bus struct {
	mu           sync.RWMutex
	subscribers  map[string][]chan TelemetryEvent
	Dropped      int64
	Sampled      int64
	OnOverflow   func(topic string, pressure float64, event TelemetryEvent)
	overflowSeen map[string]bool
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewBus initializes the event distribution backbone.
// I/O: None.
// Complexity: O(1).
func NewBus() *Bus {
	return &Bus{
		subscribers:  make(map[string][]chan TelemetryEvent),
		overflowSeen: make(map[string]bool),
	}
}

// LABEL: [MUTATES_STATE] [PUBLIC_API] [STABLE]
// Subscribe registers a new consumer channel for a given topic.
// I/O: None.
// Side Effects: Allocates a new buffered channel and updates the subscribers map.
// Complexity: O(1).
func (b *Bus) Subscribe(topic string) chan TelemetryEvent {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan TelemetryEvent, QueueCapacity)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

// LABEL: [MUTATES_STATE] [PUBLIC_API] [STABLE]
// Publish distributes a TelemetryEvent to all subscribers of the given topic with adaptive shedding.
// I/O: None (operates on channels).
// Side Effects: Increments Dropped count on overflow; triggers OnOverflow callback if pressure is high.
// Complexity: O(S) where S is the number of subscribers for the given topic.
func (b *Bus) Publish(topic string, event TelemetryEvent) {
	// [HARDENING]: Ensure event integrity.
	if event.Hash == "" {
		sig := ComputeEventSignature(event, nil)
		event.Hash = hex.EncodeToString(sig)
	}

	b.mu.RLock()
	subs, ok := b.subscribers[topic]
	b.mu.RUnlock()

	if !ok {
		return
	}

	for _, ch := range subs {
		fillRatio := float64(len(ch)) / float64(QueueCapacity)

		if fillRatio >= HighWatermark {
			if event.Severity < 0.8 {
				b.Dropped++
				continue
			}
		}

		if fillRatio >= CriticalWatermark {
			b.mu.Lock()
			seen := b.overflowSeen[topic]
			if !seen && b.OnOverflow != nil {
				b.overflowSeen[topic] = true
				go b.OnOverflow(topic, fillRatio, event)
			}
			b.mu.Unlock()

			if event.Severity < 0.9 {
				b.Dropped++
				continue
			}
		} else {
			b.mu.Lock()
			if b.overflowSeen[topic] {
				b.overflowSeen[topic] = false
			}
			b.mu.Unlock()
		}

		select {
		case ch <- event:
		default:
			if event.Severity >= 0.8 {
				select {
				case <-ch: // Drop oldest
					ch <- event // Insert new
					fmt.Printf("[BUS SHIELD] Priority Pre-emption: dropped oldest to make room for seq %d\n", event.SeqID)
				default:
					b.Dropped++
				}
			} else {
				b.Dropped++
				fmt.Printf("[BUS EMERGENCY] Queue at 100%%. Dropping seq %d\n", event.SeqID)
			}
		}
	}
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// QueuePressure returns the maximum fill ratio across all subscriber channels for a given topic.
// I/O: None.
// Complexity: O(S) where S is the number of subscribers for the topic.
func (b *Bus) QueuePressure(topic string) float64 {
	b.mu.RLock()
	subs, ok := b.subscribers[topic]
	b.mu.RUnlock()

	if !ok || len(subs) == 0 {
		return 0.0
	}

	var maxPressure float64
	for _, ch := range subs {
		pressure := float64(len(ch)) / float64(QueueCapacity)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}
	return maxPressure
}

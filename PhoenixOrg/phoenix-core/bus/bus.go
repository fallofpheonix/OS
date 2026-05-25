package bus

import (
	"encoding/json"
	"fmt"
	"sync"
)

const (
	QueueCapacity     = 65536
	HighWatermark     = 0.85
	CriticalWatermark = 0.95
)

// TelemetryEvent structure matching EVENT_SCHEMA.md
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

	// Hardening Fields
	EventID     string `json:"event_id"`
	CausalID    string `json:"causal_id"`
	SequenceNo  int64  `json:"sequence_no"`
	SourceEpoch int64  `json:"source_epoch"`
}

type Bus struct {
	mu           sync.RWMutex
	subscribers  map[string][]chan TelemetryEvent
	Dropped      int64
	Sampled      int64
	OnOverflow   func(topic string, pressure float64, event TelemetryEvent)
	overflowSeen map[string]bool
}

func NewBus() *Bus {
	return &Bus{
		subscribers:  make(map[string][]chan TelemetryEvent),
		overflowSeen: make(map[string]bool),
	}
}

func (b *Bus) Subscribe(topic string) chan TelemetryEvent {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan TelemetryEvent, QueueCapacity)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

func (b *Bus) Publish(topic string, event TelemetryEvent) {
	b.mu.RLock()
	subs, ok := b.subscribers[topic]
	b.mu.RUnlock()

	if !ok {
		return
	}

	for _, ch := range subs {
		fillRatio := float64(len(ch)) / float64(QueueCapacity)

		// 1. Evidence Reserve / Critical Event Lane
		// If queue is filled past HighWatermark (85%), block non-critical events.
		// Reserved exclusively for critical events (Severity >= 0.8).
		if fillRatio >= HighWatermark {
			if event.Severity < 0.8 {
				b.Dropped++
				continue
			}
		}

		// 2. Critical Watermark & Overflow Snapshot (95%)
		if fillRatio >= CriticalWatermark {
			b.mu.Lock()
			seen := b.overflowSeen[topic]
			if !seen && b.OnOverflow != nil {
				b.overflowSeen[topic] = true
				// Trigger overflow callback synchronously to ensure deterministic ledger ordering
				b.OnOverflow(topic, fillRatio, event)
			}
			b.mu.Unlock()

			// Elevate requirements further in the critical zone
			if event.Severity < 0.9 {
				b.Dropped++
				continue
			}
		} else {
			// Reset overflow trigger when pressure drops back down
			b.mu.Lock()
			if b.overflowSeen[topic] {
				b.overflowSeen[topic] = false
			}
			b.mu.Unlock()
		}

		select {
		case ch <- event:
		default:
			// RED TEAM MITIGATION: Pre-emption Shield
			// If queue is 100% full, but event is high-priority, drop the OLDEST
			// event to make room for the new one.
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

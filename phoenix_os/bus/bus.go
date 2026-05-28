package bus

import (
	"crypto/sha256"
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
	LamportClock uint64          `json:"lamport_clock"`
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
	Tgid        uint32 `json:"tgid"`
	Nsproxy     uint32 `json:"nsproxy"`
}

type Bus struct {
	mu           sync.RWMutex
	subscribers  map[string][]chan TelemetryEvent
	Dropped      int64
	Sampled      int64
	OnOverflow   func(topic string, pressure float64, event TelemetryEvent)
	overflowSeen map[string]bool
	topicHashes  map[string]string
	lamportClock uint64
}

func NewBus() *Bus {
	return &Bus{
		subscribers:  make(map[string][]chan TelemetryEvent),
		overflowSeen: make(map[string]bool),
		topicHashes:  make(map[string]string),
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
	b.mu.Lock()
	// Increment Logical Clock
	b.lamportClock++
	event.LamportClock = b.lamportClock

	// Chaining for Ordering Proofs
	event.PrevHash = b.topicHashes[topic]
	
	h := sha256.New()
	h.Write([]byte(event.EventID))
	h.Write([]byte(event.PrevHash))
	h.Write([]byte(fmt.Sprintf("%d", event.LamportClock)))
	h.Write(event.Payload)
	event.Hash = fmt.Sprintf("%x", h.Sum(nil))
	
	b.topicHashes[topic] = event.Hash
	
	subs, ok := b.subscribers[topic]
	b.mu.Unlock()

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
				b.OnOverflow(topic, fillRatio, event)
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
				case <-ch:
					ch <- event
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

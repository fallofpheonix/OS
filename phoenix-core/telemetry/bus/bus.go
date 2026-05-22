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
}

type Bus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan TelemetryEvent
	Dropped     int64
	Sampled     int64
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]chan TelemetryEvent),
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

		if fillRatio >= CriticalWatermark {
			if event.Severity < 0.8 {
				b.Dropped++
				continue
			}
		}

		if fillRatio >= HighWatermark && fillRatio < CriticalWatermark {
			if event.Severity < 0.5 && event.SeqID%2 == 0 {
				b.Sampled++
				continue
			}
		}

		select {
		case ch <- event:
		default:
			b.Dropped++
			fmt.Printf("[BUS EMERGENCY] Queue at 100%%. Dropping seq %d\n", event.SeqID)
		}
	}
}

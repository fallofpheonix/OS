package agent_bus

import (
	"sync"
)

// Event represents a message passed through the agent bus.
type Event struct {
	Type    string      `json:"type"`
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
}

// Bus is an in-memory pub/sub broker for agents.
type Bus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Event
}

// NewBus initializes a new Agent Bus.
func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe returns a channel that receives events for a specific topic.
func (b *Bus) Subscribe(topic string) chan Event {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan Event, 100)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

// Publish broadcasts an event to all subscribers of a topic.
func (b *Bus) Publish(topic string, event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if subs, ok := b.subscribers[topic]; ok {
		for _, ch := range subs {
			select {
			case ch <- event:
			default:
				// Subscriber too slow, skip to avoid blocking the bus
			}
		}
	}
}

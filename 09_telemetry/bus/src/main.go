package main

import (
	"sync"
)

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan interface{}
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (b *EventBus) Subscribe(topic string) chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan interface{}, 1024)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

func (b *EventBus) Publish(topic string, event interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if subscribers, ok := b.subscribers[topic]; ok {
		for _, ch := range subscribers {
			select {
			case ch <- event:
			default:
				// Buffer full, drop or handle accordingly
			}
		}
	}
}

func main() {
	// Simple simulation
	bus := NewEventBus()
	ch := bus.Subscribe("telemetry")
	
	go func() {
		bus.Publish("telemetry", "test-event")
	}()
	
	<-ch
}

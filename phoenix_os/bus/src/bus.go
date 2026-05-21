package main

import (
	"sync"
	"time"
)

type Message struct {
	Topic string      `json:"topic"`
	ID    string      `json:"id"`
	TS    int64       `json:"ts"`
	Data  interface{} `json:"data"`
}

type Bus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Message
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]chan Message),
	}
}

func (b *Bus) Subscribe(topic string) chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 1024)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

func (b *Bus) Publish(topic string, data interface{}) {
	b.mu.RLock()
	subs, ok := b.subscribers[topic]
	b.mu.RUnlock()

	if !ok {
		return
	}

	msg := Message{
		Topic: topic,
		TS:    time.Now().UnixNano(),
		Data:  data,
	}

	for _, ch := range subs {
		select {
		case ch <- msg:
		default:
			// Buffer full, drop event to prevent blocking system-critical collectors
		}
	}
}

func main() {
	// Simple boot check
	bus := NewBus()
	ch := bus.Subscribe("phoenix.sys.boot")
	
	go bus.Publish("phoenix.sys.boot", "Bus is Online")
	<-ch
}

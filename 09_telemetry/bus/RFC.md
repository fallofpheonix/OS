# RFC: Event Bus (L3)

## 1. Design
The Event Bus uses Go channels and a pub/sub pattern for low-latency distribution.

## 2. Interface
```go
type Bus struct { ... }
func (b *Bus) Publish(topic string, event interface{})
func (b *Bus) Subscribe(topic string) <-chan interface{}
```

## 3. Reliability
- Non-blocking publish to prevent collector stalling.
- Drop policy for slow subscribers.

---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# ADR-003: Event Bus Architecture

## Status
Accepted

## Date
2026-05-31

## Context
The system needs a central event distribution mechanism that can handle high throughput, priority events, and overflow conditions without blocking.

## Decision
Implement a topic-based pub/sub event bus with:
- 65536 capacity per subscriber channel
- HighWatermark (85%) for critical event lanes
- CriticalWatermark (95%) for extreme pressure
- Pre-emption shield for high-severity events at 100%

## Consequences

### Easier
- Decoupled producers and consumers
- Priority-based event handling
- Overflow protection prevents system lockup

### More Difficult
- No guaranteed delivery (drop-on-full)
- No event ordering across subscribers
- No replay capability (events are transient)

## Alternatives Considered
1. **Kafka-style log** — Rejected: too complex for single-node
2. **In-memory queue** — Rejected: no overflow protection
3. **Channel-per-subscriber** — Rejected: doesn't support fan-out

## References
- [PhoenixCore README](../../PhoenixCore/README.md)
- [bus/bus.go](../../PhoenixCore/bus/bus.go)

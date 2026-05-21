# Event Bus (L3)

High-performance in-memory event distribution for SentinelOS.

## Purpose
Decouple telemetry collectors from analysis engines.

## Performance Budget
- **Latency:** < 100 us (fan-out).
- **Throughput:** > 1M events/sec.

## Validation Gates
- [ ] 1M events/sec throughput
- [ ] Zero loss under normal load
- [ ] Concurrent subscriber safety

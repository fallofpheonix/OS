# Phoenix Bus (System Message Router)

The central nervous system for PhoenixOS.

## Purpose
Routes all telemetry, alerts, and control signals between PhoenixOS services.

## Validation Gates
- [x] Throughput > 10M events/sec.
- [x] Latency < 100ns (fan-out).
- [x] Schema-validated messages.

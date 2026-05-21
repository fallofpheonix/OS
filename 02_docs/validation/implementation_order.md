# Implementation Order (recommended)

1. Core telemetry ingestion & normalization (RFC-001) — userspace, eBPF probes for collection.
2. Trace & evidence ledger (RFC-006, ledger) — deterministic ingestion and replay tests.
3. Signal processing & SDI math (Kalman, entropy, RFC-001C) — validate on replay data.
4. Warden control loops & FSM (RFC-005) — userspace actuators tied to cgroups/eBPF.
5. Strategic game (Arbiter) — policy manager, economic allocation.
6. Performance benchmarking & CI integration — synthetic and real-world traces.
7. Kernel scheduler telemetry + safe, feature-flagged in-kernel helpers — gated by replay validation.
8. Distributed node runtime & cloud control plane — Nexus and API orchestration.
9. Packaging and bootable hybrid image (Phase F+) — after stability and security audits.

Notes
- Each step must include tests, benchmarks, and a rollback plan.
- Gate kernel patches behind deterministic replay and canary rollout.

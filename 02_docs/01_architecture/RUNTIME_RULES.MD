# RUNTIME_RULES.md

## 1. Flow Invariants
The Phoenix Matrix enforces a unidirectional, deterministic flow:
Telemetry → Replay → Truth → Arbiter → Warden → Containment → Recovery

## 2. Purity Rules
- No non-deterministic primitives in the fast-path.
- All state transitions must be HMAC-signed.
- AI advice is asynchronous and non-blocking.

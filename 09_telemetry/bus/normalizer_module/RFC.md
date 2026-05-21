# RFC: Event Normalizer Module

Status: DRAFT

Purpose
- Define the normalization pipeline for raw telemetry events into the SentinelOS canonical event schema.

Scope
- Userspace only. No eBPF or kernel probes executed by this module. It is a validation and development harness used to prove latency and correctness targets before promotion.

Validation Gates
- Acceptance: per-event normalization latency <= 5 microseconds (Python prototype target: <10μs; final: Go/Rust <5μs)
- Correctness: schema validation pass rate 100% for provided sample inputs

Replay Requirements
- Must be able to accept JSON event streams and replay deterministically 3x producing identical outputs.

Threat Model
- Accepts untrusted event inputs; must validate and sanitize `path`, `comm`, and numeric fields. Not privileged.

Failure Modes
- Schema drift: missing fields or type mismatches
- Performance regression under bursty inputs

# Design: Guard Runtime

## Summary
The Guard Runtime is a userspace daemon that connects fast-detection logic to the Kernel Actuator. It provides a low-latency, auditable execution path for emergency mitigations while preserving system safety and replay determinism.

## Goals
- Provide a predictable, low-latency path (<100ms median) from detection → mitigative actuation.
- Maintain an auditable evidence ledger entry for each action (trace_hash, sdi, policy, action, result, time, confidence, replay, experiment).
- Operate within bounded privileges and follow least-privilege principles.
- Support deterministic replay for validation and forensics.

## Non-Goals
- Replace higher-level strategic agents or long-horizon planning.
- Perform heavy-weight machine learning inference (offload to separate processes/services).

## Architecture Overview
- Components:
  - Guard Daemon: userspace process responsible for receiving detection events and coordinating actuation.
  - Fast Detector Interface: in-process or IPC interface used by detection components to submit events.
  - Kernel Actuator Client: minimal, well-audited interface used to request kernel-level mitigations.
  - Evidence Logger: appends cryptographic evidence tuples to the Evidence Ledger.
  - Replay Adapter: records inputs/outputs necessary for deterministic replay.

## Interfaces
- IPC: use Unix domain sockets with a compact, versioned message schema (msgpack or CBOR).
- CLI: `phoenix-guardctl status|replay|simulate` for operators.
- Kernel Actuator: small syscall shim or netlink-like interface; strictly typed commands.

### Message Schema (proposal)
{
  "id": "uuid",
  "time": "iso8601",
  "source": "detector-id",
  "confidence": 0.0-1.0,
  "policy": "policy-id",
  "action": { "type": "throttle|kill|isolate|quarantine", "params": {...} }
}

## Security & Privilege Model
- Run Guard Daemon as a dedicated service account with minimal filesystem access.
- Kernel Actuator requires explicit approval via policy; all requests are logged and signed.
- Use process sandboxing (seccomp, capability drops) where available.

## Failure Modes
- Missed deadlines (latency spikes): fall back to higher-latency strategic pipeline and record incident.
- Kernel Actuator rejects action: log full evidence and escalate to operator channel.
- Replay divergence: mark record as suspect and trigger forensic collection.

## Validation & Testing
- Unit tests for message parsing, policy checking, and evidence logging.
- Integration tests with a Kernel Actuator test harness (simulated kernel responses).
- Determinism tests: run controlled replay inputs and assert same outputs and ledger entries.
- Performance benchmarks: measure p50/p95/p99 latency for the full detection→actuation path; target p95 < 100ms.

## Milestones
1. RFC + message schema (this doc)
2. Prototype Guard Daemon with stub Kernel Actuator + tests
3. Integrate Evidence Ledger logging and deterministic replay adapter
4. Hardening (seccomp, capabilities), benchmarks, and CI gates

## Open Questions
- Preferred IPC encoding: msgpack vs CBOR vs protobuf for compactness and schema evolution?
- How to version kernel actuator commands across distributions?

## Commands / Local Run
Use the following to run the prototype (example):
```
# start guard daemon (dev)
PYTHONPATH=$PWD .venv/bin/python -m phoenix.guard.daemon

# operator CLI
.venv/bin/phoenix-guardctl status
```

## Authors
Maintainers: Phoenix Security / Kernel teams

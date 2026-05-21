# PhoenixOS Telemetry Event Schema

## Philosophy
This document freezes the telemetry schema contract. The Phoenix Bus, Trace, and Monitor layers rely exclusively on this deterministic schema. The schema must never be evolved ad hoc.

## JSON Schema Definition

Every event emitted by `phoenix-guard` (or the Replay Harness acting as the Guard Adapter) must conform to this schema.

```json
{
  "seq_id": 12421,
  "monotonic_ns": 18273645123,
  "wall_time_unix": 1710000000,
  "source": "guard.exec",
  "host_id": "node-01",
  "pid": 812,
  "tid": 812,
  "uid": 1000,
  "gid": 1000,
  "event_type": "process.exec",
  "severity": 0.42,
  "payload": {},
  "prev_hash": "...",
  "hash": "..."
}
```

## Field Definitions
- `seq_id` (Integer): Global monotonic sequence ID injected by the Guard adapter.
- `monotonic_ns` (Integer): The trusted monotonic clock time of the event (used strictly for replay ordering).
- `wall_time_unix` (Integer): Wall-clock time of the event (untrusted, informational only).
- `source` (String): The origin of the event (e.g., `guard.exec`, `guard.fs`).
- `host_id` (String): Identity of the host machine.
- `pid` (Integer): Process ID.
- `tid` (Integer): Thread ID.
- `uid` (Integer): User ID.
- `gid` (Integer): Group ID.
- `event_type` (String): Bounded event type (e.g., `process.exec`, `fs.read`, `net.connect`).
- `severity` (Float): Initial pre-calculated heuristic severity score (0.0 to 1.0).
- `payload` (Object): Context-specific properties (e.g., paths, IP addresses).
- `prev_hash` (String): Cryptographic link to the preceding event hash.
- `hash` (String): SHA-256 hash of this event's fields + `prev_hash`.

## Isomorphic Guarantees
Whether an event originates from native eBPF probes or the `JSONL` Replay harness, it must undergo mapping into this exact structure before entering `phoenix-bus`. The Bus and Trace layers do not know the underlying origin type.

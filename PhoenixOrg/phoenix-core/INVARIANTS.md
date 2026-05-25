# PhoenixOS: Cybernetic State & Time Invariants

This document specifies the formal mathematical and systems invariants required to guarantee 100% deterministic replayability and state stability in the PhoenixOS Phase 1 runtime.

---

## 1. Core Invariants

### I1: Pure Input-Output Determinism
$$\text{Replay}(S_0, E_{1..N}) \implies S_N$$
Given an identical initial system state $S_0$ and an identical sequence of telemetry events $E_{1..N}$, the final state $S_N$ must be byte-for-byte identical across runs.

### I2: Reproducible Cryptographic Ledger
$$\text{Hash}(\text{Ledger}_A) = \text{Hash}(\text{Ledger}_B)$$
Two executions replaying the same input events must produce an identical SHA-256 hash-chain verification root in the Evidence Ledger. Any variance in ordering, content, or count violates this invariant.

### I3: Deterministic FSM Transitions
$$\delta(S_k, e) \to S_{k+1}$$
The Warden Finite-State Machine state transitions must be a pure function of the logical event sequence and current state. Transition times must be evaluated solely against logical event timestamps, bypassing host scheduling jitters.

### I4: Zero Wall-Clock Dependency
$$\frac{\partial S}{\partial t_{\text{system}}} = 0$$
No component in the active execution path may call `time.Now()`, `time.NewTicker()`, or any wall-clock timer. All logical time progresses strictly inline as a data field of the telemetry stream (`EventTimeUnix` / `WallTimeUnix`).

### I5: Ordered Serialization
All persisted states and payload messages must use strict, stable serialization. Struct-to-JSON serialization must guarantee field ordering. Map iteration (`map[string]interface{}`) must never be serialized or hashed without canonicalization.

### I6: Deterministic Iteration
No state update or hash generation may rely on Go's non-deterministic map iteration. Collections requiring iteration must be stored in slices or sorted keys map wrappers.

### I7: Bounded Queues
All internal event channels must enforce backpressure or deterministic drop/sampling strategies. Queue sizes are bounded to prevent tail latency inflation or memory exhaustion.

### I8: Bounded Memory Space
The active working set size of the sliding window and FSM cache must remain bounded:
$$\lim_{t \to \infty} \text{Memory}(t) \le M_{\text{limit}}$$
Historical states and events older than the defined sliding window duration (60s) must be pruned deterministically inline.

### I9: Restart-Safe Reproducibility
The runtime must produce the same ledger hash-chain regardless of whether it is run continuously or halted and resumed (via deterministic transaction replay from the SQLite WAL storage).

---

## 2. Invariant Verification Matrix

| Invariant | Subsystem | Enforcement Mechanism | Verification Method |
| :--- | :--- | :--- | :--- |
| **I1 / I2** | Global / Ledger | Flat-struct hashing, sequential validation | Multi-run hash matching against `test_events.jsonl` |
| **I3 / I4** | Warden FSM | Logical Unix timestamps passed in event data | Hysteresis verification via simulated event timeline |
| **I5 / I6** | Ledger / Serialization | Struct-based parsing, no raw map-hashing | Compilation constraints & code review audits |
| **I7 / I8** | Event Bus / TCS | Bounded channels, inline window pruning | Out-of-memory checks, event throughput stress tests |
| **I9** | Trace / WAL | SQLite journal_mode=WAL, unique SeqID constraint | Database integrity checks and replay validation |

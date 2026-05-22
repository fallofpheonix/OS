# PHOENIXOS MASTER ISSUE BACKLOG

## P0 — DETERMINISM FOUNDATION

### 1. Deterministic Monotonic Clock Abstraction

Labels:
`determinism`, `core`, `critical`, `architecture`

Tasks:
* Create logical monotonic runtime clock
* Remove wall-clock dependency
* Replay-safe timestamp generation
* Sequence-linked event timing

Acceptance:
* Replay outputs identical timing state
* No `std::chrono::system_clock`
* Replay checksum stable

---

### 2. Canonical Event Serialization Format

Labels:
`serialization`, `determinism-critical`

Tasks:
* Define canonical byte ordering
* Eliminate architecture-specific padding
* Versioned schema encoding
* Stable hash generation

Acceptance:
* Identical hashes across machines
* Stable replay deserialization

---

### 3. Deterministic Hash Chain Ledger

Labels:
`ledger`, `critical`, `replay`

Tasks:
* Append-only ledger
* SHA-256 chained records
* Sequence-indexed blocks
* Replay verification hooks

Acceptance:
* Tamper detection operational
* Ledger rewind reproducible

---

### 4. Replay Integrity Validator

Labels:
`verification`, `replay`

Tasks:
* Replay recorded event streams
* Compare final state hashes
* Detect divergence points

Acceptance:
* Divergence detection under 1 event offset

---

### 5. Event Ordering Enforcement Layer

Labels:
`scheduler`, `determinism`

Tasks:
* Remove nondeterministic queue scheduling
* Stable ordering guarantees
* FIFO replay enforcement

Acceptance:
* Event order identical in replay

---

### 6. Floating Point Determinism Audit

Labels:
`math`, `high-risk`

Tasks:
* Audit float usage
* Replace unstable operations
* Fixed-point evaluation where needed

Acceptance:
* Identical replay across CPU vendors

---

### 7. Replay Snapshot State System

Labels:
`snapshot`, `replay`

Tasks:
* Create checkpoint snapshots
* Delta replay support
* Snapshot hash verification

Acceptance:
* Fast replay rewind functional

---

### 8. Deterministic UUID Replacement

Labels:
`determinism`, `infra`

Tasks:
* Remove random UUID generation
* Replace with sequence-derived IDs

Acceptance:
* Stable entity IDs in replay

---

### 9. Replay-Aware Memory Allocator

Labels:
`memory`, `critical`

Tasks:
* Stable allocation ordering
* Deterministic arena allocator
* Replay-safe reuse

Acceptance:
* Memory layout stable under replay

---

### 10. Deterministic Randomness Engine

Labels:
`rng`, `core`

Tasks:
* Seed-controlled PRNG
* Replay seed persistence
* Remove nondeterministic entropy

Acceptance:
* Replay outputs identical

---

## P1 — TELEMETRY + TCS (L3 Phoenix Monitor)

### 11. Telemetry Collection Pipeline

Tasks:
* Structured event ingestion
* Stable schema contracts
* Sequence-number enforcement

---

### 12. Telemetry Compression Layer

Tasks:
* Delta compression
* Replay-safe decompression

---

### 13. Bounded Telemetry Queue

Tasks:
* Ring-buffer queue
* Overflow handling policy
* Saturation metrics

---

### 14. Telemetry Corruption Detector

Tasks:
* CRC verification
* Ledger mismatch detection

---

### 15. Telemetry Schema Versioning

Tasks:
* Backward-compatible schema evolution
* Replay compatibility guarantees

---

### 16. TCS Confidence Scoring Engine

Labels:
`math`, `telemetry`, `scoring`

Tasks:
* Compute entropy metrics for incoming data streams
* Implement Kalman filter for state estimation and noise reduction
* Establish sensor fusion weighted scoring

Acceptance:
* Validated noise rejection
* Accurate tracking of latent variables

---

### 17. Signal-to-Noise Processor

Labels:
`math`, `signal-processing`

Tasks:
* Implement high-pass and low-pass event filters
* Extract baseline deviations from telemetry
* Normalize raw eBPF metrics for the scoring engine

Acceptance:
* Deterministic output on given telemetry inputs

---

## P2 — KERNEL & INTEGRITY (L1 Phoenix Guard & L2 Phoenix Kernel)

### 18. eBPF Probe Lifecycle Manager

Labels:
`ebpf`, `kernel`, `lifecycle`

Tasks:
* Safe attachment and detachment of eBPF programs
* Fallback handling on load failure
* Monitor eBPF maps for saturation

Acceptance:
* Zero kernel panics on probe lifecycle events
* Probe state explicitly tracked

---

### 19. XDP Fast-Path Enforcement Module

Labels:
`xdp`, `network`, `fast-path`

Tasks:
* <100ms packet dropping for malicious flows
* High-confidence (Entropy > 7.9) heuristic bypass integration
* Deterministic packet logging into ledger

Acceptance:
* Sub-100ms latency for drop actuation
* Fast-path decisions bypass L5.5 strategy layers

---

### 20. Replay-Safe System Call Interceptor

Labels:
`kernel`, `syscall`, `determinism`

Tasks:
* Hook critical syscalls deterministically
* Record syscall arguments into telemetry queue
* Ensure hook ordering does not introduce race conditions

Acceptance:
* Syscall traces match 1:1 on replay instances

---

## P3 — GRAPH INTELLIGENCE (L4 Phoenix Trace)

### 21. Causal DAG Builder

Labels:
`graph`, `lineage`, `causality`

Tasks:
* Build Directed Acyclic Graphs of process/network events
* Link child processes to parent context
* Track causal chain of file modifications

Acceptance:
* DAG visualization matches actual causal chains
* No orphaned nodes in standard operation

---

### 22. 3-Tier Storage Lifecycle Manager (Hot/Warm/Cold)

Labels:
`storage`, `architecture`, `data-lifecycle`

Tasks:
* Define policies for promoting/demoting DAG nodes
* HOT: RAM cache (active analysis)
* WARM: SSD storage (recent context)
* COLD: Compressed archive (historical replay)

Acceptance:
* Memory boundaries respected for HOT tier
* Deterministic eviction of nodes

---

### 23. Critical Node Pruning Protections

Labels:
`graph`, `safety`

Tasks:
* Protect `init`, `auth`, `kernel`, and `systemd` nodes from being pruned
* Implement reference counting for protected lineages

Acceptance:
* Base system nodes exist perpetually in WARM/HOT storage

---

## P4 — ACTUATION & CONTROL (L5 Phoenix Warden)

### 24. Finite-State Control Machine

Labels:
`control`, `state-machine`, `core`

Tasks:
* Implement strictly defined states: SAFE, WATCH, SUSPICIOUS, CRITICAL, COMPROMISED
* Define allowed state transitions
* Prevent direct SDI mapping to control gains

Acceptance:
* Invalid state transitions rejected and logged
* Control decisions strictly governed by current FSM state

---

### 25. Rate-Limited Actuator Engine

Labels:
`control`, `safety`, `actuation`

Tasks:
* Implement token-bucket or leaky-bucket rate limiting for system actions
* Bound the maximum frequency of kill/isolate commands
* Record all actuation attempts to Phoenix Ledger

Acceptance:
* Actuator unable to exceed predefined rate limits
* Attempted violations trapped

---

### 26. State Reversal & Rollback Mechanism

Labels:
`control`, `recovery`

Tasks:
* Support undoing isolation/throttling actions
* Track "action impact" to reverse specific controls safely

Acceptance:
* Rollbacks return the system exactly to the prior constraint state

---

## P5 — STRATEGIC POLICY (L5.5 Phoenix Arbiter)

### 27. Game-Theoretic Stackelberg Solver

Labels:
`math`, `policy`, `ai`

Tasks:
* Model attacker-defender dynamics as a Stackelberg game
* Compute optimal defensive strategy (Leader) against anticipated attacks (Follower)
* Keep AI output as advisory, never direct actuation

Acceptance:
* Solver returns deterministic policy recommendations
* AI output explicitly logged as non-authoritative

---

### 28. Cost-Benefit Action Optimizer

Labels:
`policy`, `optimization`

Tasks:
* Assign "cost" to system disruption and "benefit" to risk reduction
* Evaluate potential actions against the cost-benefit matrix
* Feed optimal recommendations to Phoenix Warden

Acceptance:
* High-cost low-benefit actions are consistently suppressed

---

## P6 — SYSTEM PHYSICS (L6 Phoenix Sentinel)

### 29. Thermodynamic SDI Monitor

Labels:
`physics`, `entropy`

Tasks:
* Model system state as a thermodynamic state
* Continuously monitor the System Disorder Index (SDI)
* Identify state "heating" (disorder increasing)

Acceptance:
* SDI accurately reflects system stress and anomaly levels

---

## P7 — SWARM COORDINATION (L7 Phoenix Nexus)

### 30. Proof-of-Authority Consensus Protocol

Labels:
`swarm`, `consensus`, `network`

Tasks:
* Implement PoA among trusted Swarm nodes
* Maintain deterministic ledger state across the cluster
* Ensure single-node stability is foundational before scaling

Acceptance:
* Multi-node cluster achieves consensus deterministically
* Network partitions result in safe degradation

---

### 31. Peer Reputation Scoring Engine

Labels:
`swarm`, `security`

Tasks:
* Assign reputation scores to nodes based on verified ledgers
* Quarantine nodes falling below reputation thresholds

Acceptance:
* Corrupt or divergent nodes are successfully quarantined from consensus

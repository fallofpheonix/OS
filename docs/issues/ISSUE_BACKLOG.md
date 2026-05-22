# PHOENIXOS ISSUE BACKLOG

This backlog follows the stage-based hierarchy and the deterministic control paradigm.

## M1-Deterministic-Replay (P0 - Priority)

### GROUP A: Deterministic Event Core
- [ ] ISSUE-001: Canonical Event Schema
- [ ] ISSUE-002: Logical Event Time (Remove wall-clock)
- [ ] ISSUE-003: Deterministic Sequence Ordering
- [ ] ISSUE-004: Canonical JSON Encoder
- [ ] ISSUE-005: Map Iteration Audit

### GROUP B: Replay Runtime
- [ ] ISSUE-006: Replay Engine Core
- [ ] ISSUE-007: Replay Hash Validation
- [ ] ISSUE-008: Replay Rewind
- [ ] ISSUE-009: Replay Snapshotting
- [ ] ISSUE-010: Replay Compression

### GROUP C: Ledger
- [ ] ISSUE-011: Async Ledger Worker
- [ ] ISSUE-012: Hash Chain Verification
- [ ] ISSUE-013: Ledger Recovery
- [ ] ISSUE-014: Ledger Flush Semantics

### GROUP D: TCS + Monitor
- [ ] ISSUE-015: Sliding Window TCS
- [ ] ISSUE-016: Inline Pruning
- [ ] ISSUE-017: Degradation Monitor
- [ ] ISSUE-018: Telemetry Confidence Score
- [ ] ISSUE-019: Oscillation Protection

### GROUP E: FSM
- [ ] ISSUE-020: Warden FSM Core
- [ ] ISSUE-021: Hysteresis Logic
- [ ] ISSUE-022: Cooldown Semantics
- [ ] ISSUE-023: Actuation Classes
- [ ] ISSUE-024: Replay-safe Mode (Block actuation during replay)

### GROUP F: Verification Harnesses
- [ ] ISSUE-025: Determinism Harness (P0 Blockers)
- [ ] ISSUE-026: Fault Injection Harness
- [ ] ISSUE-027: Saturation Harness
- [ ] ISSUE-028: Oscillation Harness
- [ ] ISSUE-029: Bounded Memory Verification
- [ ] ISSUE-030: Bounded CPU Verification

## M2-Real-Telemetry (Stage 2)
(Blocked until M1 passes)
- [ ] ISSUE-031: Ringbuf Adapter
- [ ] ISSUE-032: Tracepoint Collector
- [ ] ISSUE-033: Kprobe Collector
- [ ] ISSUE-034: XDP Ingress
- [ ] ISSUE-035: Kernel Event Normalization
- [ ] ISSUE-036: Telemetry Rate Limiting
- [ ] ISSUE-037: Ringbuf Overflow Handling

## M3-Immutable-Runtime (Stage 3)
(Blocked until M2 passes)
- [ ] ISSUE-038: LinuxKit Builder
- [ ] ISSUE-039: QEMU Boot Runtime
- [ ] ISSUE-040: PID1 Warden
- [ ] ISSUE-041: Immutable Initrd
- [ ] ISSUE-042: BPF Filesystem Mount
- [ ] ISSUE-043: Minimal Boot Verification

## M4-Constrained-Control (Stage 4)
- [ ] ISSUE-044-049: Socket Blocking, Process Isolation, Budgeting, Rollback, Recovery, Control Loop Stabilization.

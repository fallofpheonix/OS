# PhoenixOS Architecture Invariants

This file defines the non-negotiable architectural invariants and safety properties that must be preserved by all systems within the PhoenixOS substrate.

## Formal Safety Invariants

These invariants must be proved statically (via TLA+) and validated dynamically (via `PhoenixValidation` replay):

1. **History Immutability (Ledger)**: The cryptographic history of the ledger is append-only. Once an entry is committed and verified by quorum, it can never be deleted or modified.
2. **Deterministic Replayability (State)**: Given the exact same sequence of input events and initial state, the system must reconstruct the identical state transitions and outputs bit-for-bit.
3. **Evidence Immutability**: All evidence records, once committed to the truth ledger, are cryptographically sealed and cannot be modified.
4. **Advisory Containment (AI)**: The advisory layer (`PhoenixMind`) has zero direct actuation capabilities. It can only emit structured recommendations (`AdvisoryEnvelope`) which require substrate approval.
5. **Enforcement Authorization**: The low-latency enforcement path (`PhoenixGuard` / `Warden`) is forbidden from executing any action that is not explicitly permitted by a signed policy or valid FSM transition.
6. **Consensus Validity**: The distributed ledger cannot commit state transitions without an authenticated quorum proof from authorized nodes.
7. **Kernel Isolation**: The kernel space (`PhoenixKernel`) must never accept directives or query instructions directly from the AI advisory layer (`PhoenixMind`).
8. **Acyclic Lineage (Trace)**: Telemetry correlation graphs and process execution lineage paths must remain directed acyclic graphs (DAGs). Cycles are forbidden.
9. **Deterministic Truth Engine**: The evaluation of evidence trust scores and confidence assessments must remain a pure function of input evidence properties, calibrated models, and weights.
10. **Memory Snapshot Reproducibility**: Retrieval vectors, episodic memories, and semantic snapshots must remain identical when reconstructed from a given epoch seed.

## Formal Liveness Guarantees

1. **Event Processing**: All events published to valid topics on the event bus must eventually be processed or dead-lettered with lineage preserved.
2. **Ledger Commits**: Evidence records submitted to the distributed ledger must eventually commit or raise a consensus/network divergence alarm.
3. **Consensus Resolution**: Under normal network conditions (fewer than $F$ Byzantine failures), the node consensus term must eventually resolve and progress.
4. **Replay Termination**: Any deterministic state replay operation must terminate within a bounded timeframe to avoid denial-of-service on verification routines.
5. **Validation Completion**: Fuzz and chaos testing tasks must eventually complete and produce structured validation reports.
6. **Trace Convergence**: Process and causality trees must eventually resolve to their terminal roots.
7. **Recovery Bounds**: Systems entering the `RECOVERING` state of the FSM must eventually transition to a stable run-time state (`IDLE` or `SAFE`) or trigger a hard shutdown.
8. **Synchronization Stability**: Diverged nodes must eventually synchronize with the canonical ledger or be evicted from the validator quorum.

## Hard Architectural Rules

1. **AI Containment**: Never place AI components inside critical enforcement loops.
2. **Kernel Boundary**: Never allow kernel-facing probes or hooks outside the `PhoenixKernel` repository.
3. **No Side Effects**: State mutations are strictly prohibited during event replay validation.
4. **Transparency**: Every enforcement action taken by the system must be fully explainable via its lineage graph, evidence markers, and associated FSM log.
5. **Stability First**: Never scale or distribute a cluster running in an unstable state.
6. **Action Boundaries**: Every enforcement action requires a corresponding rollback plan, timeout, scope limit, and deterministic state transition path.
7. **Consensus Resilience**: Distributed operations must remain resilient to network partitions, node failures, and message replays.

# PhoenixOS TLA+ Formal Verification Specifications

This directory contains TLA+ (Temporal Logic of Actions) and PlusCal formal specifications for the critical safety and liveness properties of the PhoenixOS substrate.

## Formal Verification Map

The following models are verified using the TLC Model Checker to ensure mathematical correctness of the system's core behaviors.

### 1. Ledger Safety (`LedgerSafety.tla`)
- **Variables**: `ledger`, `history`, `checkpoints`, `tamper_flag`
- **Safety Invariant**: `LedgerAppendOnly == \A t1, t2 \in Domain(history): t1 <= t2 => history[t1] = SubSequence(history[t2], 1, Len(history[t1]))` (History is never mutated).
- **Liveness Property**: `LedgerProgress == []<>(Len(ledger) > prev_len)`
- **Counterexample Expectation**: Any attempt to modify a committed ledger entry triggers a tamper state.

### 2. Event Ordering (`EventOrdering.tla`)
- **Variables**: `event_bus_queue`, `monotonic_logical_clock`, `causal_graph`
- **Safety Invariant**: `CausalOrdering == \A e1, e2 \in events: e2 \in e1.causal_chain => e2.monotonic_time < e1.monotonic_time`
- **Liveness Property**: `EventDelivery == \A e \in events: e.status = "published" ~> e.status = "processed"`

### 3. Replay Correctness (`ReplayCorrectness.tla`)
- **Variables**: `state`, `inputs`, `replay_pointer`, `replay_states`
- **Safety Invariant**: `DeterministicState == (replay_pointer = Len(inputs)) => (state = expected_final_state)`
- **Properties**: Strong determinism validation. Identical input sequences guarantee bit-for-bit identical state transitions.

### 4. Guard State Machine (`GuardFSM.tla`)
- **Variables**: `fsm_state`, `policy_approval`, `action_scope`, `rollback_triggered`
- **Safety Invariant**: `NoUnapprovedAction == (fsm_state = "ISOLATING" \/ fsm_state = "RESPONDING") => policy_approval = TRUE`
- **Forbidden Transitions**: `IDLE` -> `ISOLATING` without `DETECTING` trigger; any transition skipping the `RESPONDING` rollback setup.

### 5. Consensus Safety (`ConsensusSafety.tla`)
- **Variables**: `nodes`, `leader`, `term`, `votes`, `quorum_proofs`
- **Safety Invariant**: `SingleLeaderPerTerm == \A n1, n2 \in nodes: (n1.leader = n2.leader) \/ (n1.term /= n2.term)`
- **Liveness Property**: `ConsensusEventualResolution == []<>(consensus_reached = TRUE)`

### 6. Memory Consistency (`MemoryConsistency.tla`)
- **Variables**: `memory_records`, `semantic_vectors`, `vector_indexes`, `snapshots`
- **Safety Invariant**: `SnapshotReproducibility == \A s1, s2 \in snapshots: s1.epoch = s2.epoch => s1.records_hash = s2.records_hash`

### 7. Truth Consistency (`TruthConsistency.tla`)
- **Variables**: `evidence_set`, `contradictions`, `confidence_score`, `assessments`
- **Safety Invariant**: `DeterministicAssessment == \A a1, a2 \in assessments: a1.evidence_hash = a2.evidence_hash => a1.conclusion = a2.conclusion`

## Running Model Checks

TLA+ specifications are checked using the TLC Command Line tool:

```bash
java -cp tla2tools.jar tlc2.TLC -workers auto LedgerSafety.tla
```

Detailed model definitions, constants, and `.cfg` files live alongside each `.tla` spec in this folder.

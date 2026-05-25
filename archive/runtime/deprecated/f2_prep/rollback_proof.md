# Rollback & Consistency Proofs

## 1. No Oscillation Proof
The Warden FSM implements a `DwellTicks` (Hysteresis) barrier. Any transition from a higher-severity state to a lower-severity state is blocked unless `current_tick >= last_transition_tick + DwellTicks`. This ensures that the system cannot flap between states rapidly, preventing unstable actuation loops.

## 2. Rollback Consistency
Rollback is deterministic because it relies on the immutable `Truth Ledger`. A rollback to `Tick T` involves:
1. Identifying the last valid snapshot at `Tick S <= T`.
2. Replaying the event log from `S` to `T`.
Since the event log and snapshots are hash-chained and immutable, the resulting state is guaranteed to be identical across re-runs.

## 3. Hash Chain Integrity
Every entry in the `Phoenix Ledger` contains the SHA-256 hash of its predecessor. 
`Hash(E_n) = SHA256(Payload_n || Sequence_n || Hash(E_{n-1}))`
Any mutation of `Payload_i` (where `i < n`) will invalidate the entire chain from `i` to `n`, making tampering immediately detectable by the verifier.

## 4. Decision Reproducibility
The `Arbiter` is a pure function of:
`Evaluate(Score, Policy, TCS) -> (State, Class, Authorized)`
Given the same input vector and the same policy configuration, the Arbiter will always produce the same output, ensuring deterministic security policy enforcement.

# PHOENIX MATRIX: ARCHITECTURAL BASELINE

**Status:** AUTHORITATIVE
**Last Updated:** 2026-06-05
**Compliance:** MANDATORY

This document defines the authoritative design decisions for the Phoenix Matrix OS substrate. All future implementation MUST adhere to these selected baselines.

## 1. Ledger Architecture
*   **LEDGER-001 (Source of Truth):** **C. Event Ledger.** The state is a deterministic projection of the immutable event log.
*   **LEDGER-002 (State Recovery):** **C. Ledger Replay.** Replay is the authoritative mechanism for reconstructing state.
*   **LEDGER-003 (Ordering):** **C. Ledger Index.** Monotonic indices determine absolute temporal causality.
*   **LEDGER-004 (Schema):** **C. Canonical Schema.** Events use bit-perfect typed structures. Max size: 1MB (LEDGER-SPEC-001).
*   **LEDGER-005 (Hashing):** **D. Canonical Binary Encoding.** Fixed BigEndian byte streams ensure cross-architecture compatibility.

## 2. Replay Engine
*   **REPLAY-001 (Integrity Gaps):** **C. Halt.** Any missing event index is a fatal integrity violation.
*   **REPLAY-002 (Duplicate Indices):** **C. Halt.** Indicates ledger corruption or illegal state mutation.
*   **REPLAY-003 (Validation Scope):** **D. All of the above.** Replay validates state hashes, chain links, and event signatures.
*   **REPLAY-004 (Tick Semantics):** **D. Round-based Ticks.** Ticks represent deterministic simulation steps and increment per consensus round (LEDGER-SPEC-001).

## 3. Consensus (BFT-SPEC-002)
*   **CONSENSUS-001 (Final Authority):** **C. Quorum Certificate.** Proof of agreement signed by `2f + 1` authorized validators via a **HotStuff 3-chain pipeline**.
*   **CONSENSUS-002 (Membership):** **C. Signed Governance Event.** Validator set updates are deterministic state transitions authored via `UPDATE_VALIDATOR`.
*   **CONSENSUS-003 (Signature Target):** **D. Explicit Protocol Object.** Validators sign `SignedEnvelope` digests which bind payload, epoch, and sequence.
*   **CONSENSUS-004 (Fork Resolution):** **C. Highest QC.** The chain with the most advanced verified finality proof wins.
*   **CONSENSUS-005 (Safety):** **B. Honest nodes never diverge after finalization.** Pipelined agreement ensures zero-fork finality.

## 4. Persistence
*   **STORAGE-001 (Crash Safety):** **C. Detectable Recovery State.** Partial lines are skipped during recovery (reverts to last complete JSONL).
*   **STORAGE-002 (Recovery Strategy):** **B. Snapshot + Ledger Tail.** Snapshots optimize startup; replay validates the gap.
*   **STORAGE-003 (Snapshot Role):** **B. Cache.** Snapshots are performance artifacts, never the primary source of truth.
*   **STORAGE-004 (Snapshot Verification):** **C. State Hash.** Restoring requires recalculating and verifying `SnapshotHeader.StateHash`.

## 5. VM Design
*   **VM-001 (State Mutation):** **C. Only through ApplyEvent.** Gameplay/simulation cannot modify state without producing a signed event.
*   **VM-002 (AI Autonomy):** **B. Through events.** AI agents propose transitions which are processed via deterministic paths.
*   **VM-003 (Opcode Execution):** **B. Events.** VM instructions produce event emissions, isolating logic from side-effects.
*   **VM-004 (Determinism):** **C. Always.** Determinism is a SACRED axiom.

## 6. Security
*   **SECURITY-001 (Key Storage):** **C. Encrypted Keystore.** Private keys protected by **Argon2id + XChaCha20-Poly1305**.
*   **SECURITY-002 (Invalid Signatures):** **C. Rejected.** Forged messages discarded at gossip layer.
*   **SECURITY-003 (Replay Protection):** **C. Nonce/Sequence Validation.** Enforced via `LastSeenSequences` in state and `SnapshotHeader`.
*   **SECURITY-004 (Removal):** **D. No.** Removed validators lose authority at Round 0 of the following Epoch.

## 7. Networking
*   **NETWORK-001 (Node Sync):** **C. Snapshot + Proof.** Nodes bootstrap via finalized QC and snapshot.
*   **NETWORK-002 (Peer Trust):** **C. Untrusted.** Local state machines independently verify every network byte.
*   **NETWORK-003 (Message Lifecycle):** **C. Validation.** Messages must pass signature and schema validation before ingestion.

---
*Authorized by Phoenix Sovereign Governance*

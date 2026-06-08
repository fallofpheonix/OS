# Phoenix Matrix OS: Foundational Mandates

This file establishes the immutable axioms and architectural mandates for the Phoenix monorepo. All development MUST adhere to these rules.

## 1. Six Immutable Axioms
1. **Determinism is sacred.** No reliance on non-deterministic primitives (unordered maps, race-dependent ordering, float64 in state paths).
2. **Replay is authoritative.** Replay governs semantics; logs, metrics, and AI-outputs are secondary.
3. **AI is collaborative with guarded autonomy.** AI can propose high-impact actuations but requires explicit user/governance permission.
4. **Control must remain bounded.** Actuation is rate-limited, isolated, and reversible.
5. **Telemetry correctness > AI sophistication.** Precise, replayable telemetry is the foundation.
6. **Never scale instability.** Single-node stability must be achieved before distributed scaling.

## 2. Engineering Standards
*   **Contract-First:** Subsystems MUST communicate via versioned contracts in `foundation/contracts`.
*   **Zero Float64:** `float64` is prohibited in all security, physics, and state-transition paths. Use `phxmath.FixedPoint`.
*   **Atomic Persistence:** The ledger is the only source of truth. All state changes MUST be traceable to a signed ledger event.
*   **BFT Readiness:** Every node is untrusted. All network messages MUST be validated cryptographically and semantically before ingestion.

## 4. Current Phase: Stage B (Formal Invariants)
*   **Status:** Stage A (Hardening) COMPLETE. Transitioning to Stage B.
*   **Recent Milestone:** STEP 6.6 LEDGER HARDENING COMPLETED.
    *   Proved Snapshot + Replay Equivalence (LEDGER-007).
    *   Implemented Corruption Detection Suite (LEDGER-008).
    *   Standardized Canonical Binary Encoding across all layers (LEDGER-010).
    *   Established Golden Hash benchmarks for cross-platform consistency (LEDGER-011).

## 5. Immediate Roadmap
1.  **STEP 7: NETWORKING** - P2P gossip integration and state synchronization.
2.  **STEP 8: CONSENSUS** - Quorum Certificate (QC) logic and Weighted-BFT implementation.
3.  **Formal Proofs:** TLA+ modeling of the Warden FSM and transactional event bus guarantees.

---
*Authorized by Phoenix Sovereign Governance*

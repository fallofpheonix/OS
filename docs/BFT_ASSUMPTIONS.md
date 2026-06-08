# PHOENIX MATRIX: BFT FAULT MODEL & ASSUMPTIONS

**Status:** AUTHORITATIVE
**Document ID:** BFT-SPEC-001
**Maturity:** Stage 6.6 (Specification)

## 1. Network Assumptions
*   **Model:** **Partially Synchronous.** The system assumes that after some unknown Global Stabilization Time (GST), messages between honest nodes are delivered within a known bound `Δ`.
*   **Transport:** GossipSub over libp2p. Reliability is assumed via re-broadcast; safety does not depend on perfect liveness.
*   **Latency Target:** `Δ < 2000ms` for nominal operation.

## 2. Validator Model
*   **Validator Count (N):** Target `N < 100`. Initial deployment `N = 4` to `N = 10`.
*   **Fault Tolerance (f):** The system tolerates up to `f` Byzantine nodes where `N >= 3f + 1`.
*   **Weighting:** Initial implementation uses **Uniform Weighting** (1 vote per public key). 
*   **Admission:** **Permissioned.** New validators must be added via a Governance QC (`UPDATE_VALIDATOR`).

## 3. Byzantine Adversary Capabilities
We assume an adversary can:
*   **Equivocate:** Sign conflicting events or votes at the same height/round.
*   **Delay:** Withhold messages to honest nodes up to the `Δ` bound.
*   **Collude:** Up to `f` nodes can coordinate to prevent finality or attempt forks.
*   **Replay:** Capture and rebroadcast signed envelopes (mitigated by sequence/epoch tracking).

We assume an adversary CANNOT:
*   **Forge Signatures:** Breakthrough Ed25519 security.
*   **Break Hashing:** Produce SHA-256 collisions.
*   **Invert VRF:** (If implemented later) Predict random seeds.

## 4. Safety & Liveness
*   **Safety (Axiomatic):** Honest nodes will never commit conflicting states at the same index. Safety is maintained under any number of faults up to `f`.
*   **Liveness:** The system will continue to make progress (commit new events) after GST, provided at least `2f + 1` nodes are honest and communicative.

---
*Authorized by Phoenix Sovereign Governance*

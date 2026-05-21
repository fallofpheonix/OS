# PhoenixOS: Task List & Technical Strategy

## Priority 0 (P0): Critical Foundation
1. **[Phoenix Ledger]** Evidence Ledger: Implement cryptographic tuple `(trace_hash, sdi, policy, action, result, time, confidence, replay, experiment)`.
2. **[Phoenix Guard]** Fast Path Enforcement: Bypassing strategic layers for <100ms mitigation of heuristics (entropy bursts, crypto writes).
3. **[Phoenix Trace]** 3-Tier DAG Storage: Implement HOT (active), WARM (compressed), COLD (skeleton) storage to prevent memory explosion.

## Priority 1 (P1): Core Intelligence
4. **[Phoenix Sentinel]** Finite-State Controller: Replace continuous SDI-to-PID gain with discrete state progression (SAFE -> WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED).
5. **[Phoenix Monitor]** Process Importance Score: Replace simple PageRank with multi-factor scoring (Centrality + Criticality + Entropy + Spread + Depth).
6. **[Phoenix Guard]** Guard Runtime: Userspace daemon connecting fast-detector logic directly to the Kernel Actuator.

## Priority 2 (P2): Swarm Reliability
7. **[Phoenix Arbiter]** Byzantine Swarm Protection: Implement Proof-of-Authority + Node Reputation Score + Weighted Quorum consensus.
8. **[Phoenix Arbiter]** Suspicion Counter: Track lying nodes and deduct reputation.

## Priority 3 (P3): Advanced Control
9.  **[Phoenix Arbiter]** MARL Stability: Implement Action Debt, Cooldown Timers, and Maximum Containment Rates before Lyapunov analysis.
10. **[Phoenix Kernel]** Global Energy Budget: System-wide control of action limits.

---
## Phase 1: Core Communications & Monitoring (Legacy Plan)
*Tasks 1-3 moved to updated priorities above or completed in workflow repo.*

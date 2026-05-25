# PhoenixOS Release Gates

This document defines the strict cryptographic, deterministic, and functional gating criteria required to transition PhoenixOS between phases of development.

---

## Phase F0 Exit Criteria (Foundation Stabilization)
Gates the transition of Contracts, State, and Containment from **YELLOW** to **GREEN** (Hardened).

| Gate Requirement | Target Metric | Status |
| :--- | :--- | :--- |
| **Replay Stress** | $\ge 1000$ unique execution traces executed | `PENDING` |
| **Hash Mismatch** | Exactly 0 drift or mismatch alerts across all runs | `PENDING` |
| **Snapshot Corruption** | Exactly 0 corrupted payloads accepted; 100% rejection rate of bad hashes/versions | `PASS` (Verified in C1.5/C2.4/C3.6/C4.6) |
| **Restore Repeatability** | 100% byte-for-byte state recovery across 100 consecutive runs | `PASS` (Verified in C5.5/C5.6) |
| **Chaos Suite** | 100% pass rate under telemetry jitter and out-of-order logs | `PENDING` |
| **Fuzz Suite** | 100% pass rate on systematic ledger and FSM payload mutations | `PENDING` |
| **Containment State** | Overall status: **GREEN** | `PENDING` |

---

## Phase F1 Exit Criteria (Runtime Completion)
Gates the transition from the foundation substrate to the live active research platform.

| Gate Requirement | Target Metric | Status |
| :--- | :--- | :--- |
| **eBPF Live Loader** | Production loader successfully injected into host kernel | `PENDING` |
| **Telemetry Ingestion** | Event ring buffer matches host execution events with zero drops | `PENDING` |
| **Causal Trace** | Temporal graph updates active and process lineage DAGs fully searchable | `PENDING` |
| **Warden Control** | FSM active on host process boundaries via systemd and cgroups | `PENDING` |
| **Arbiter Policy** | Cost-benefit policy evaluation returns correct mitigation scores under attack | `PENDING` |
| **Recovery Loop** | Snapshots and rollbacks trigger automatically on security policy alerts | `PENDING` |

---

## Phase F2 Exit Criteria (Recovery Platform)
Gates the transition of the active recovery loop to automated production capability.

| Gate Requirement | Target Metric | Status |
| :--- | :--- | :--- |
| **Cross Rollback** | Multicomponent (Process + Network + File) rollbacks complete in $< 100$ms | `PENDING` |
| **Recovery Replay** | Verification of post-restore replay parity with no state divergence | `PENDING` |
| **Determinism Check** | Replay of recovery actions remains 100% deterministic across varying hosts | `PENDING` |
| **Verification Proof** | Formal verification proof chains pass successfully on every rollback | `PENDING` |

Risk Register Update — Phoenix
================================

Seeded risks detected during repo audit:

- Premature kernel work
  - Likelihood: High
  - Impact: High
  - Mitigation: Freeze kernel prototypes until experiments validated

- Missing telemetry schema enforcement
  - Likelihood: High
  - Impact: High
  - Mitigation: Prioritize RFC-001 and normalizer implementation

- Unconnected theory islands
  - Likelihood: Medium
  - Impact: Medium
  - Mitigation: Map each theory to telemetry and experiments (THEORY_TO_OS_MAP.md)
# Risk Register Update (Architecture Constraints)

| ID | Risk | Severity | Mitigation |
| :--- | :--- | :--- | :--- |
| **ARCH-01** | Mathematical computations (Entropy/Games) take >1ms, crashing the event pipeline. | CRITICAL | Move compute off the main eBPF ring-buffer reader thread. Implement C/Rust mathematical backend instead of Python. |
| **ARCH-02** | PID Controller overshoot starves benign system processes. | HIGH | Implement strict bounds (Max Throttle) and fail-open policies on PID actuators. Validate via R031. |
| **ARCH-03** | AI layers implemented before deterministic physics layers are proven. | CRITICAL | Enforce dependency lock: AI subsystems (`06_ai`) cannot be integrated into runtime until L4 (Graphs) and L6 (Physics) pass validation gates. |
| **ARCH-04** | Replay divergence from event reorder, queue pressure, or burst loss. | CRITICAL | Add deterministic sequence allocation, bounded reorder windows, and replay equivalence tests. |
| **ARCH-05** | Ring-buffer starvation drops critical evidence under telemetry flood. | CRITICAL | Reserve critical lanes, add overflow snapshots, and prioritize evidence traffic. |
| **ARCH-06** | Warden oscillation causes self-DoS and unstable recovery. | HIGH | Add hysteresis, cooldown windows, state dwell limits, and recovery budgets. |
| **ARCH-07** | Ledger transitions are too weak to prove state-before/state-after integrity. | HIGH | Add validation hashes, policy versioning, and replay hashes to each transition. |
| **ARCH-08** | Recovery or shadow-mode paths become broader attack surfaces than the live path. | HIGH | Split identities for recovery vs runtime, require just-in-time approval, and log every fallback invocation. |
| **ARCH-09** | Policy/config drift causes nodes to enforce different security thresholds. | HIGH | Pin signed policy snapshots, enforce rollout quorum, and alarm on checksum mismatch or stale config. |
| **ARCH-10** | Sensor or telemetry poisoning corrupts the evidence chain before scoring. | MEDIUM | Authenticate sensor sources, cross-check with independent signals, and make ingestion tamper-evident. |
| **ARCH-11** | Release-chain compromise bypasses validated runtime guarantees. | HIGH | Require hermetic builds, provenance attestation, and signed deployment artifacts. |

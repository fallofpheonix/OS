# RFC-010: Arbiter-Warden Control Interface

**Status:** Approved / Implemented
**Review Date:** 2026-05-23
**Author:** Documentation Evolution Agent
**Target Subsystem:** L5/L5.5 Control Loop

## 1. Objective
Define a rigid, state-aware, and failure-safe interface between the L5.5 Strategic Policy Arbiter and the L5 Actuation Warden. The system must adhere to the principle: **Every autonomous action must be observable, explainable, reversible, and bounded.**

## 2. Interface Formalization
The interface shall be modeled as a restricted Finite State Machine with mandatory policy constraints:

```text
Arbiter (Goal) -> Policy Constraint Layer -> Warden (Action)
```

### 2.1 Actuation Classes (Risk Tiers)
All actions must be categorized:
*   **CLASS 0 (Observe):** Data gathering only.
*   **CLASS 1 (Log/Alert):** Informational only.
*   **CLASS 2 (Throttle):** Resource constraints (e.g., process limit).
*   **CLASS 3 (Local Isolate):** Local process/container isolation.
*   **CLASS 4 (Cluster Isolate):** Distributed node isolation (Requires consensus).
*   **CLASS 5 (Kernel Emergency):** Direct kernel intervention (Requires high-confidence quorum).

### 2.2 Safety Guardrails & Telemetry Confidence
*   **Telemetry Confidence Score (TCS):** Arbiter actions are gated by a calculated TCS (based on packet loss, drift, queue pressure, etc.). High-risk classes (3+) require TCS > threshold.
*   **Actuation Budgeting:** Enforce hard limits on CPU, Network, Isolation, and AI inference budgets consumed by security actions to prevent self-DoS.
*   **Simulation-before-Execution:** High-risk actions (Class 4+) must pass through a shadow simulation to predict impact before real-world execution.

### 2.3 Warden Stabilization Constraints
To defend the state machine against rapid oscillation and containment storms, the Warden implements a strict multi-tier stabilization framework:
*   **Actuation Cooldown:** Locks further transitions for 10 logical ticks post-actuation. Critical overrides (Class 3+) bypass this lock.
*   **State Dwell Limits:** Minimum time of 30 logical ticks that the Warden must stay in an elevated state (e.g. SUSPICIOUS, CONTAINED) before being eligible for automatic de-escalation.
*   **Recovery Budgeting:** Limits automatic recoveries (de-escalations) to a maximum of 3 per session, after which the Warden will lock in the higher safety state until manual operator reset.

## 3. Failure Modes, Recovery & Distributed Coordination
*   **Distributed Consensus:** Cross-node actions (Class 4+) require leader-validated quorum to prevent fragmentation and split-brain scenarios.
*   **Recovery Semantics:** Reintegration requires validated evidence of safety (e.g., threat cleared, state stable) before moving back to `SAFE` state.
*   **Communication Loss:** Fail-closed (maintain current state).

## 4. Implementation Sequencing
1.  Define Protocol Buffer definition (Class + Action + Constraints).
2.  Implement Safety Layer (Validation logic + TCS check).
3.  Implement Budgeting and Rate-limiting.
4.  Integrate with Warden FSM (starting with SAFE -> WATCH -> SUSPICIOUS).

## 5. Risks
- Undefined failure modes leading to cluster-wide self-DoS.
- Latency added to the Fast-Path by safety checks.

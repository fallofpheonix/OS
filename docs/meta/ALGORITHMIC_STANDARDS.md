# PhoenixOS: Algorithmic Standards (v1.0)

This document defines the **binding mathematical foundations** for all PhoenixOS subsystems. No subsystem may implement core logic using algorithms outside this whitelist without a formal RFC approval.

## 1. Determinism & Evidence (Stages 1, 5)
*   **Canonical Serialization:** MUST use Canonical JSON (sorted keys, consistent float representation) or Bencode for all wire and ledger formats.
*   **Hash-Chaining:** Merkle DAG structures required for all evidence chains.
*   **Causal Ordering:** Lamport Logical Clocks (must increment on internal event, $\max(local, recv) + 1$ on communication).

## 2. Telemetry & Physics (Stages 2, 9, 10)
*   **Running Statistics:** Welford’s Online Algorithm for variance/deviation ($O(1)$ memory, numerically stable).
*   **Smoothing:** Exponentially Weighted Moving Average (EWMA) with fixed-point arithmetic.
*   **Anomalies:** Shannon Entropy ($H(X) = -\sum P(x) \log_2 P(x)$).
*   **State Estimation:** Kalman Filters (fixed-point iteration).

## 3. Constrained Control (Stages 4, 11)
*   **Actuation Gating:** PID Control (Proportional-Integral-Derivative) with mandatory clamping/anti-windup logic.
*   **Traffic Shaping:** Token Bucket / Leaky Bucket (deterministic rate-limiting).

## 4. Graph & Causality (Stage 6)
*   **Traversal:** Time-Respecting Paths (only follow edges where $t_{edge\_start} < t_{edge\_end}$).
*   **Risk Propagation:** Personalized PageRank on process/cgroup graphs.

## 5. Strategic Defense (Stage 11)
*   **Attacker Profiling:** Bayesian Inference ($P(A|B) = \frac{P(B|A) \cdot P(A)}{P(B)}$).
*   **Defense Strategy:** Stackelberg Security Games (Leader-Follower optimality).

---
**Violation Note:** Implementation of floating-point-based "naive" algorithms (e.g., standard variance) in the Replay path is a P0 Security Vulnerability due to floating-point drift.

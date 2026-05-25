# PhoenixOS: Algorithm & Model Map

This document maps the project's evolution stages to specific algorithms and mathematical models.

## 1. Determinism & Evidence (Stages 1 & 5)
*   **Canonical Serialization:** Sorted key ordering and stable float representation (e.g., Canonical JSON) to ensure identical hashes for identical data.
*   **Merkle DAGs:** Structured evidence ledger allowing parallel event branching with a single root hash for system state verification.
*   **Lamport Logical Clocks:** Causal event ordering replacing non-deterministic wall-clock time in replay traces.

## 2. Signal Processing & Physics (Stage 10)
*   **Welford’s Online Algorithm:** $O(1)$ memory, numerically stable running variance and standard deviation calculation for telemetry streams.
*   **EWMA (Exponentially Weighted Moving Average):** Bounded, predictable smoothing of noisy telemetry signals.
*   **Shannon Entropy:** $H(X) = -\sum P(x) \log_2 P(x)$ to detect packed malware or encrypted exfiltration.
*   **Kalman Filters:** Real-time system-state estimation to isolate signals from background OS noise and detect telemetry spoofing.

## 3. Constrained Control & Actuation (Stages 4 & 12)
*   **PID Control (Proportional-Integral-Derivative):** Smooth, non-oscillatory throttling of system resources based on anomaly magnitude and persistence.
*   **Token/Leaky Bucket:** Strict mathematical bounds on data egress and process execution rates during suspicious states.

## 4. Graph & Causality Runtime (Stage 6)
*   **Temporal Graph Traversal:** Time-respecting pathfinding ($t_1 < t_2$) to reconstruct valid attack chains.
*   **Personalized PageRank:** Propagating risk scores through the process lineage and network DAG to identify blast radius.

## 5. Strategic Defense & Game Theory (Stage 11)
*   **Bayesian Inference:** Updating attacker profiles and intent probabilities as new telemetry evidence is replayed.
*   **Stackelberg Security Games:** Solving for optimal defensive allocations (probes, decoys, isolation) against a rational, observing adversary.

## Stage-Algorithm Alignment Matrix

| Stage | Key Algorithms / Models |
| :--- | :--- |
| **1: Det. Replay** | Lamport Clocks, Canonical Serialization, Stable Hashing, Bounded FIFO |
| **2: Real Telemetry** | eBPF-Ringbuf, EWMA Rate-Estimators, Little’s Law Queue Design |
| **4: Actuation** | PID Control, Token Bucket, Warden FSM |
| **5: Evidence** | Hash Chains, Merkle-DAGs, LTL Runtime Verification |
| **6: Graphs** | DAG Traversal, BFS/DFS, PageRank-Centrality, Cycle Detection |
| **7: Distributed** | Vector Clocks, CRDTs, Gossip/Consensus-lite |
| **9: Advisory AI** | Isolation Forest, Autoencoders (Anomaly), Graph Embeddings |
| **10: Physics** | Kalman Filters, Shannon Entropy, Welford's Algorithm |
| **11: Game Theory** | Stackelberg Games, Bayesian Games, LP/IP Optimization |
| **12: Cybernetic** | Closed-loop Bounded Policy, SMT-based Model Checking |

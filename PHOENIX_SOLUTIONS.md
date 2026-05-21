# PhoenixOS: Expert-Aligned Solutions

This document defines the refined architectural bridges for the PhoenixOS Matrix.

## 1. Finite-State Control (L6 -> L5)
**Solution:** Replace linear/exponential gain with a **State-Aware Controller**.
- **Regime:** Map SDI $\in [0, 1]$ to discrete states:
  - `0.0 - 0.3`: `SAFE` (Collect telemetry)
  - `0.3 - 0.5`: `WATCH` (Enable L3 math filters)
  - `0.5 - 0.7`: `SUSPICIOUS` (Throttle CPU to 50%)
  - `0.7 - 0.9`: `CRITICAL` (Process Suspension)
  - `0.9 - 1.0`: `COMPROMISED` (Isolation + Termination)

## 2. Kernel Guard Fast-Path (P0)
**Solution:** **[PLANNED]**
- **Detector:** Phoenix Monitor executes L3 heuristics (Entropy/Rename) and signals the **Guard Runtime**.
- **Actuator:** `phoenix_guard.c` (eBPF) monitors the policy map and blocks syscalls in <100ms.

## 3. Importance Scoring ($S_I$)
**Solution:** Use a multi-dimensional importance vector for process valuation.
- $S_I = w_1 \cdot \text{Rank} + w_2 \cdot \text{Criticality} + w_3 \cdot \text{Entropy} + w_4 \cdot \text{Depth}$
- This prevents "Mimicry" by focusing on what a process **does** and its position in the **Phoenix Trace** DAG, not just its name.

## 4. Reputable Consensus
**Solution:** Implement **Proof-of-Anomaly (PoA)** with weighted voting.
- Every node maintains a **Reputation Score**.
- Anomalies from "High Reputation" nodes require a smaller quorum to trigger swarm-wide containment.
- "Lying Nodes" (False Positives) lose reputation and are eventually isolated.

## 5. 3-Tier Trace Storage
**Solution:** Solve memory explosion via **Lifecycle Compression**.
- **HOT (In-Memory):** Active DAG.
- **WARM (SSD/Disk):** Protobuf-compressed lineage of processes exited in the last 24h.
- **COLD (Archive):** Minified "Skeleton" chain (ID, Parent, Entry, Exit) for long-term forensics.

## 6. Swarm Debt & Cooldown
**Solution:** Control swarm oscillation using **Action Debt**.
- Every actuation "costs" energy. Nodes have a recovery rate.
- This naturally prevents "Containment Storms" where nodes enter a loop of mutual re-throttling.

## 7. Cryptographic Evidence Chain (P0)
**Solution:** **[PLANNED]**
- Every actuation is permanently recorded in the **Phoenix Ledger**.
- Each entry is hashed and chained, providing a verifiable "Reasoning Chain" for audits and automated replay.

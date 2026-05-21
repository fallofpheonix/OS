# PhoenixOS: Gap Analysis & Expert Assessment

This document tracks critical gaps between PhoenixOS theory and implementation, aligned with expert architectural assessments.

## 1. The Transfer Function Gap (L6 -> L5)
**Status:** **[REFINED]**
**Original Problem:** Direct SDI-to-Gain mapping risk. High SDI could lead to massive $K_p$, causing system-wide oscillation and "containment storms."
**Expert Assessment:** Move from non-linear gain to a **Finite-State Controller**.
**States:**
- `SAFE`: Observation only.
- `WATCH`: Increase telemetry sampling rates.
- `SUSPICIOUS`: Trigger process throttling (PID).
- `CRITICAL`: Atomic isolation (STOP signal).
- `COMPROMISED`: Forensic snapshot + termination.

## 2. Enforcement Lag
**Status:** **[SOLVED in P0]**
**Problem:** Path from detection (userspace) to enforcement (kernel) creates an exploit window.
**Expert Assessment:** Implement a **Kernel Guard (Fast Path)**.
**Implementation:** Phoenix Guard (<100ms) bypasses strategic layers for high-confidence heuristics (Entropy bursts, crypto-write signatures).

## 3. Strategic Mimicry
**Status:** **[REFINED]**
**Problem:** Attackers mimicking low-priority processes to avoid game-theoretic monitoring.
**Expert Assessment:** PageRank is insufficient. Need an **Importance Score** ($S_I$).
**Formula:** $S_I = \text{Centrality} + \text{File Criticality} + \text{Entropy Contribution} + \text{Network Spread} + \text{Lineage Depth}$.

## 4. Byzantine Swarm Poisoning (L7)
**Status:** **[IN PROGRESS]**
**Problem:** Single node compromise can trigger a cluster-wide self-DoS.
**Expert Assessment:** Merkle proofs (PoA) alone are not enough. Need **Weighted Quorum** based on **Node Reputation**.
**Flow:** `Node Claim` -> `Evidence Verification` -> `Peer Quorum` -> `Acceptance`.

## 5. Memory Exhaustion in Lineage (L4)
**Status:** **[P0 BLOCKER]**
**Problem:** DAG explosion from thousands of short-lived processes.
**Expert Assessment:** Implement **3-Tier Storage**.
- **HOT:** Active graph nodes.
- **WARM:** Compressed lineage for recently exited processes.
- **COLD:** Skeleton chain (Forensic integrity).
- **Retention:** `init`, `db`, `auth`, `kernel` nodes are immutable and never pruned.

## 6. Stability of the MARL Swarm (L7)
**Status:** **[P3 PRIORITY]**
**Problem:** Oscillatory containment (Nodes mutual-throttling).
**Expert Assessment:** Implement **Action Debt** and **Cooldown Timers**.
**Rule:** A node throttled cannot be re-throttled or trigger a counter-throttle within the cooldown window.

## 7. Evidence Layer vs. AI Autonomy
**Status:** **[SOLVED in P0]**
**Problem:** Proving "Why" an autonomous action was taken.
**Expert Assessment:** **Content-Addressable Evidence Ledger**.
**Implementation:** Phoenix Ledger (P0) implemented with SHA-256 hash chaining of `(trace_hash, sdi, policy, action, result, time, confidence, replay, experiment)` tuples.

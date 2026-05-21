# RFC-001D: Strategic Containment via Game Theory

## Status
Proposed

## Context
Using Game Theory to move beyond static rules into dynamic, strategic response.

## Game Model: OS vs. Ransomware
### 1. Players
- **Defender ($D$):** Pheonix Containment Engine.
- **Attacker ($A$):** Suspected Ransomware process.

### 2. Actions
- **$A$:** [Idle, Write, Rename, Network Call]
- **$D$:** [Monitor, Throttle, Snapshot, Isolate, Terminate]

### 3. Utility Function ($U$)
$$U_D = \text{System Stability} - \text{Data Loss} - \text{Detection Latency}$$
$$U_A = \text{Data Encrypted} + \text{Privilege Gain} - \text{Detection Probability}$$

## Strategy: Minimax Containment
The OS chooses a containment action that minimizes the Attacker's maximum possible gain.
- **If Suspicion is Low:** Action = Throttle (Low cost, reversible).
- **If Suspicion is High:** Action = Isolate + Snapshot (High cost, high protection).

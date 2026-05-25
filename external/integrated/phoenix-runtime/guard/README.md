# Phoenix Guard (Fast-Path Enforcement)

The low-latency security "short-circuit" for PhoenixOS.

## Purpose
Detects and stops high-confidence threats (e.g., ransomware encryption bursts) in <100ms by bypassing the slow strategic decision layers (L5.5/L7).

## Architecture: The Two Paths
1.  **Fast Path:** Entropy bursts, bulk renames, crypto writes. (Detected in L2/L3, Enforced in L1).
2.  **Slow Path:** Complex lineage anomalies, strategic games, MARL swarm logic. (L4-L7).

## Performance Budget
- **End-to-End Latency:** < 100 ms.
- **Actuation:** Atomic process suspension/kill via BPF or CGroups.

## Validation Gates
- [ ] Fast-path trigger latency < 50ms.
- [ ] Zero false-positives on standard OS boot/update tasks.

# RFC: Phoenix Guard

## 1. Description
Phoenix Guard implements a "Guard Runtime" that sits between L2 (Kernel Telemetry) and L1 (Actuation). It uses heuristics rather than deep models to ensure speed.

## 2. Fast-Path Heuristics
- **HEU-01:** File Write Entropy > 7.9 for > 5 consecutive 4KB blocks.
- **HEU-02:** Rename rate > 50 files/sec in a single process tree.
- **HEU-03:** Read-Encrypt-Write pattern (detected via buffer similarity mismatch).

## 3. Actuation Path
Monitor (Fast Detector) -> BPF Map (Command) -> Kernel Hook (Enforcement).
Latency target: < 100ms.

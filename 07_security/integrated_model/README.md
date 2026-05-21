# Sentinel Integrated Model (L3-L6)

This module integrates the individual mathematical primitives into a unified "Idea Matrix" runtime.

## Purpose
Demonstrate the end-to-end flow:
`Telemetry (L2)` -> `Entropy (L3)` -> `Graph (L4)` -> `Physics/SDI (L6)`

## State Flow
1. **Collector:** Ingests raw byte streams.
2. **L3 Engine:** Computes Shannon Entropy to detect data randomness.
3. **L4 Engine:** Maps events to a process lineage DAG.
4. **L6 Engine:** Uses the entropy scores and graph connectivity to calculate the global **Security Disorder Index (SDI)**.

## Performance Budget
- **End-to-End Latency:** < 2 ms per event batch.

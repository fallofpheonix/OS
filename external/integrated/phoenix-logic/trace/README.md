# Phoenix Trace (Process Causality Graph)

Causal lineage tracking for PhoenixOS.

## Purpose
Maintains a real-time directed acyclic graph (DAG) of all system processes and their resource interactions.

## Validation Gates
- [ ] BFS lineage query < 1ms.
- [ ] Zero cycles (DAG invariant).
- [ ] Memory pruning for stale processes.

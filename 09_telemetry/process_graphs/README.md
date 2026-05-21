# Process Graph (L4)

Causal lineage DAG for Pheonix telemetry.

## Purpose
Convert raw timeline events into directed process-resource graphs.

## Performance Budget
- **Insertion:** < 10 us.
- **Traversal (15 levels):** < 1 ms.
- **Memory:** ~1KB per node.

## Validation Gates
- [ ] BFS/DFS traversal correctness
- [ ] Loop detection (ensure DAG property)
- [ ] Memory pruning for exited processes

# PhoenixTrace — Causal Lineage Layer

> Process lineage DAGs, causal graphs, and forensic graph intelligence for the PhoenixOS ecosystem.

## Overview

PhoenixTrace provides the causal lineage tracking for PhoenixOS. It maintains directed acyclic graphs (DAGs) of process relationships, enabling forensic analysis, threat detection, and causal chain reconstruction.

**All process relationships must be tracked in the causal DAG.**

## Repository Structure

```
PhoenixTrace/
├── engine/
│   ├── mapper.go           # ProcessNode, GraphEngine core structures
│   ├── lineage.go          # HandleFork, HandleIsolation, GetCausalChain
│   ├── engine_test.go      # Engine tests
│   ├── process_graphs/
│   │   ├── process_graphs.go  # Graph data structure
│   │   └── schema.json
│   └── process_lineage/
│       └── lineage.go      # Process lineage tracking
├── go.mod
└── go.sum
```

## Core Principles

1. **Append-Only**: Nodes and edges are never removed
2. **Orphan Handling**: Parentless processes re-parent to 0xDEADBEEF
3. **Virtual Fork**: Isolated processes create virtual child nodes
4. **Causal Chain**: Backward traversal from any process to root

## Build & Test

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...
```

## Dependencies

- **Depends on**: PhoenixCore (contracts)
- **Depended by**: PhoenixMind (graph feature), PhoenixValidation (lineage tests)

## Invariants

- DAG must be acyclic
- All processes must have a lineage chain
- Orphan processes must be re-parented
- Causal chains must be deterministic

## License

PhoenixTrace is part of the PhoenixOS ecosystem.

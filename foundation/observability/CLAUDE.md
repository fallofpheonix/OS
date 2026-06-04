# PhoenixTrace — Causal Lineage Layer

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
go test ./...
go test -race ./...
```

## Architecture
PhoenixTrace maintains directed acyclic graphs (DAGs) of process relationships for forensic analysis and threat detection.

## Key Components
- **engine/mapper.go** — ProcessNode, GraphEngine core structures
- **engine/lineage.go** — HandleFork, HandleIsolation, GetCausalChain
- **engine/process_graphs/** — Graph data structure

## Invariants
- DAG must be acyclic
- All processes must have a lineage chain
- Orphan processes must be re-parented to 0xDEADBEEF
- Causal chains must be deterministic

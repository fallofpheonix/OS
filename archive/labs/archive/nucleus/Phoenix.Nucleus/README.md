---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# Phoenix.Nucleus — Core Engine Layer

> The foundation of PhoenixOS: canonical contracts, event bus, FSM enforcement, kernel interface, and distributed coordination.

## Overview

Phoenix.Nucleus contains the core engine components that form the backbone of the PhoenixOS ecosystem. All other directories depend on Nucleus for contracts, types, and core functionality.

**Invariant: All cross-boundary types MUST be defined in PhoenixCore. No other package may export cross-boundary types.**

## Sub-Packages

| Package | Purpose | Key Files |
|---------|---------|-----------|
| **PhoenixCore** | Canonical contracts, event bus, ledger, state management | bus.go, ledger.go, state/ |
| **PhoenixGuard** | Warden FSM, kill switch, bounded execution, trust matrix | warden.go, killswitch.go, executor.go |
| **PhoenixKernel** | eBPF probes, syscall tracing, LSM hooks, namespace isolation | ebpf_loader.go, phoenix_exec.c |
| **PhoenixDistributed** | PoA consensus, peer discovery, replicated ledger | poa.go, beacon.go, node.go |
| **PhoenixFormal** | TLA+ specifications, architecture rules, dependency policies | GuardFSM.tla, ARCHITECTURE_RULES.md |
| **PhoenixValidation** | Replay validation, fuzzing, chaos testing, invariant tests | engine.go, fuzz_test.go |
| **PhoenixTrace** | Causal DAG lineage, process graph intelligence | mapper.go, lineage.go |
| **PhoenixTruth** | Evidence evaluation, contradiction detection | evaluator.go, contradiction.go |

## Dependency Graph

```
PhoenixCore (contracts)
    ↑
    ├── PhoenixGuard (depends on Core)
    ├── PhoenixKernel (depends on Core)
    ├── PhoenixDistributed (depends on Core)
    ├── PhoenixFormal (depends on Core)
    ├── PhoenixValidation (depends on Core, Guard, Kernel)
    ├── PhoenixTrace (depends on Core)
    └── PhoenixTruth (depends on Core)
```

## Build & Test

```bash
# Build all Nucleus packages
go build ./Phoenix.Nucleus/...

# Test all Nucleus packages
go test ./Phoenix.Nucleus/...

# Test with race detector
go test -race ./Phoenix.Nucleus/...
```

## Invariants

1. **PhoenixCore is canonical** — All cross-boundary types defined here
2. **Determinism is sacred** — No non-deterministic primitives
3. **Replay is authoritative** — State reconstructable from ledger
4. **FSM ladder is strict** — No state skipping allowed
5. **All actuations have rollback** — Bounded execution required
6. **Kernel trust is minimal** — eBPF programs must be verified

## Related Documents

- [PhoenixCore README](PhoenixCore/README.md)
- [INVARIANTS.md](PhoenixCore/INVARIANTS.md)
- [ARCHITECTURE_RULES.md](PhoenixFormal/ARCHITECTURE_RULES.md)
- [WORKING_MODEL.md](../WORKING_MODEL.md)

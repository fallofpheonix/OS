# PhoenixValidation — Deterministic Testing Layer

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
go test -fuzz=. ./chaos/...
```

## Architecture
PhoenixValidation provides deterministic replay validation, fuzzing, chaos engineering, and invariant testing for the PhoenixOS ecosystem.

## Key Components
- **replay/engine.go** — State reconstruction from event streams
- **replay/authority.go** — Hash comparison for drift detection
- **determinism/** — Determinism verification tests (stubs)
- **evidence/** — Evidence chain verification (stubs)
- **formal/** — Formal invariant tests (stubs)
- **proofs/** — Formal proofs (stubs)
- **kernel/** — Kernel-level tests
- **security/** — Security tests
- **chaos/** — Chaos engineering
- **soak/** — Long-running tests

## Invariants
- Replay must produce identical hash chains
- No deterministic drift allowed
- FSM transitions must follow the ladder
- Ledger must be append-only

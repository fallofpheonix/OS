# PhoenixGuard — Security Enforcement Layer

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
python -m pytest tests/
```

## Architecture
PhoenixGuard implements the Warden FSM (SAFE→WATCH→SUSPICIOUS→CRITICAL→COMPROMISED) with bounded execution harnesses. All enforcement actions MUST go through PhoenixGuard.

## Key Components
- **engine/warden.go** — Core FSM with strict state ladder
- **emergency/killswitch.go** — One-way emergency halt
- **actuation/executor.go** — Timeout + rollback execution
- **actuation/sandbox.go** — Process isolation actuator
- **policies/trust_matrix.go** — Cross-domain access control
- **runtime/** — Python runtime (filesystem, orchestration, shell)

## Invariants
- FSM transitions must follow the strict state ladder (no skipping)
- All actuations must have a rollback plan
- Kill switch is irreversible within a process lifetime
- No AI system can directly trigger state transitions

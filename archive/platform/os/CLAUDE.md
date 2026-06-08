# PhoenixOS — Top-Level Orchestration

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
go build ./...
go test ./...
```

## Architecture
PhoenixOS is the top-level orchestration layer that integrates all subsystems and provides the main binary (phoenixd).

## Key Components
- **cmd/phoenixd/** — Main binary entry point
- **phoenix_os/** — Active runtime core
- **core/** — Core utilities
- **cognition/** — AI integration
- **runtime/** — Runtime systems
- **sandbox/** — Sandbox execution

## Invariants
- All subsystems must be initialized before use
- Graceful shutdown must release all resources
- Health checks must pass before accepting traffic

# PhoenixCore — Canonical Contract System

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
PhoenixCore is the canonical contract source for all cross-boundary types. All other packages MUST import contracts from here. No other package may export cross-boundary types.

## Key Components
- **bus/** — Event bus with 65534 capacity, overflow protection, priority lanes
- **ledger/** — Append-only Merkle ledger with hash chain verification
- **state/** — System state registry with audit trail
- **tcs/** — Telemetry Confidence Score calculation
- **monitor/** — Drift detection with Kalman filter and Z-score
- **arbiter/** — Policy validation and strategic decision-making
- **containment/** — Process, file, and network containment primitives
- **common/** — Shared utilities (serialization, allocation, clock, Kalman)

## Invariants
- All messages include schema_version, created_at, source_repo, replay_sequence, validation_hash
- Event bus uses BigEndian for binary serialization
- Ledger is append-only with SHA-256 hash chain
- No non-deterministic primitives (unordered maps, race-dependent ordering)

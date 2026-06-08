# PhoenixDistributed — Distributed Coordination Layer

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
PhoenixDistributed provides PoA consensus, UDP multicast peer discovery, and a consensus-backed distributed ledger for multi-node operation.

## Key Components
- **consensus/poa.go** — Quorum voting with Ed25519 signatures
- **discovery/beacon.go** — UDP multicast beacon transport
- **discovery/registry.go** — Peer registry with reputation
- **identity/node.go** — AuthorityCertificate, NodeRegistry
- **ledger/** — ConsensusLedger interface and implementations

## Invariants
- Quorum requires ≥2/3 of total cluster weight
- All votes must be cryptographically verified
- No single node can authorize high-risk transitions
- Consensus is deterministic given the same input

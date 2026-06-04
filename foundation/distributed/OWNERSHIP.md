# PhoenixDistributed — Repository Ownership

## Owner
**Primary:** Phoenix.Nucleus Team
**Contact:** [team@phoenixos.dev]

## Purpose
PoA consensus, peer discovery, replication, node identity.

## Dependencies
- **Internal:** PhoenixCore (contracts, event bus, ledger)
- **External:** Go standard library, network stack

## Consumed By
- PhoenixOS (top-level orchestration)
- PhoenixGuard (consensus-based decisions)

## Invariants
- Only certified nodes may participate
- Quorum must be explicitly proven
- Divergence must be detected before commit
- Replay and ledger state must stay consistent

## Criticality: P1
Required for distributed operation, but single-node mode works without it.

## Status: ACTIVE

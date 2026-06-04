# PhoenixCore — Repository Ownership

## Owner
**Primary:** Phoenix.Nucleus Team
**Contact:** [team@phoenixos.dev]

## Purpose
Canonical contract source for all cross-boundary types. Event bus, ledger, state management, monitoring.

## Dependencies
- **Internal:** None (this is the foundation)
- **External:** Go standard library only

## Consumed By
- PhoenixGuard (FSM, enforcement)
- PhoenixKernel (eBPF, telemetry)
- PhoenixDistributed (consensus, replication)
- PhoenixTrace (lineage, DAG)
- PhoenixTruth (evidence, truth scoring)
- PhoenixValidation (replay, determinism)
- PhoenixMind (AI orchestrator)
- PhoenixOS (top-level orchestration)

## Invariants
- All messages include schema_version, created_at, source_repo, replay_sequence, validation_hash
- Event bus uses BigEndian for binary serialization
- Ledger is append-only with SHA-256 hash chain
- No non-deterministic primitives (unordered maps, race-dependent ordering)

## Criticality: P0
This is the foundation. If PhoenixCore fails, everything fails.

## Status: ACTIVE

# PhoenixGuard — Repository Ownership

## Owner
**Primary:** Phoenix.Nucleus Team
**Contact:** [team@phoenixos.dev]

## Purpose
Security enforcement layer. Warden FSM, bounded execution, kill switch, process isolation.

## Dependencies
- **Internal:** PhoenixCore (contracts, event bus, ledger)
- **External:** Go standard library, Linux kernel (eBPF, namespaces)

## Consumed By
- PhoenixMind (advisory only, no direct control)
- PhoenixOS (top-level orchestration)
- PhoenixDashboard (read-only monitoring)

## Invariants
- FSM transitions must follow the strict state ladder (no skipping)
- All actuations must have a rollback plan
- Kill switch is irreversible within a process lifetime
- No AI system can directly trigger state transitions

## Criticality: P0
This is the security boundary. If PhoenixGuard fails, the system is compromised.

## Status: ACTIVE

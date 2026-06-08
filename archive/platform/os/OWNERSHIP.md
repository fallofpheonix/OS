# PhoenixOS — Repository Ownership

## Owner
**Primary:** Phoenix.Terminus Team
**Contact:** [team@phoenixos.dev]

## Purpose
Top-level orchestration, CLI, API, deployment, health checks.

## Dependencies
- **Internal:** PhoenixCore, PhoenixGuard, PhoenixKernel, PhoenixDistributed, PhoenixTrace, PhoenixTruth, PhoenixValidation, PhoenixMind
- **External:** Go standard library, Docker, systemd

## Consumed By
- Operators (CLI, API)
- Load balancers (health checks)
- Orchestrators (Kubernetes, Docker)

## Invariants
- All components must be authenticated
- All components must be authorized
- All components must be monitored
- All components must be recoverable

## Criticality: P0
This is the entry point. If PhoenixOS fails, the system cannot start.

## Status: ACTIVE

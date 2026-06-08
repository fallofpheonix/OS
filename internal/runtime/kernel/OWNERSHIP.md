# PhoenixKernel — Repository Ownership

## Owner
**Primary:** Phoenix.Nucleus Team
**Contact:** [team@phoenixos.dev]

## Purpose
eBPF probes, syscall tracing, LSM hooks, namespace operations.

## Dependencies
- **Internal:** PhoenixCore (contracts, event bus)
- **External:** Linux kernel, libbpf, iproute2

## Consumed By
- PhoenixGuard (enforcement)
- PhoenixOS (top-level orchestration)

## Invariants
- eBPF probes must be signed
- Namespace operations must be verified
- Telemetry must be deterministic
- No AI dependencies

## Criticality: P0
This is the kernel boundary. If PhoenixKernel fails, the system cannot monitor or enforce.

## Status: ACTIVE

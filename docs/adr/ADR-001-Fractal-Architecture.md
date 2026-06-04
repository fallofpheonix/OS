---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# ADR-001: Contract-First Architecture Model

## Status
Accepted

## Date
2026-06-04

## Context
Executable evidence showed that Phoenix's primary architectural failure mode is contract ownership drift, not repository count.

Observed failures concentrated at boundary contracts:
- EventEnvelope versus Event
- Replay versus Reconstruct
- Actuator versus Kill

## Decision
PhoenixOS adopts a contract-first architecture:

1. Contracts are the only source of truth.
2. Only contract packages define public interfaces.
3. Implementations must adapt to contracts, never define competing public APIs.
4. Repository extraction is blocked until contract compatibility suites pass.

Internal boundary packages:
- contracts/events
- contracts/replay
- contracts/security

Governance rules:
- Validation may not import runtime internals.
- Guard may not import implementation internals.
- Tests may not access private engine state.
- Contracts may not import implementation packages.

## Consequences

### Easier
- Stable ownership of public interfaces
- Mechanical later extraction
- CI-enforced boundary checks

### More Difficult
- Up-front work to define compatibility suites
- Temporary duplication while adapters are introduced
- More explicit versioning discipline

## Alternatives Considered
1. Repository-first decomposition — rejected because contracts were not yet stable.
2. Microservices — rejected because it would increase coordination overhead before boundary stabilization.
3. Fractal self-similar directories — rejected because it did not resolve ownership drift.

## References
- [FINAL_ARCHITECTURE.md](../architecture/FINAL_ARCHITECTURE.md)
- [ARCHITECTURE_DEBT.md](../architecture/ARCHITECTURE_DEBT.md)

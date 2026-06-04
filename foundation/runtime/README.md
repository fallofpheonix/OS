# PhoenixCore — Canonical Contract System

> The deterministic core runtime containing all canonical contracts, protocol definitions, and specifications for the PhoenixOS ecosystem.

## Overview

PhoenixCore is the single source of truth for all cross-boundary types, message schemas, state machine definitions, and protocol contracts in the PhoenixOS ecosystem.

**ALL repos depend on PhoenixCore contracts. NO other repo may export cross-boundary types.**

This invariant is absolute and non-negotiable. Every component, agent, and service in the PhoenixOS ecosystem must import its contract types from this repository. Any type that crosses a process, service, or repository boundary MUST be defined here.

## Repository Structure

```
PhoenixCore/
├── contracts/          # Canonical contract definitions (YAML)
│   ├── event-bus.yaml  # Event bus topics, routing, priorities
│   ├── ledger.yaml     # Append-only ledger protocols
│   └── fsm/            # Finite State Machine contracts
│       ├── guard.yaml
│       ├── warden.yaml
│       ├── arbiter.yaml
│       ├── distributed.yaml
│       └── replay-engine.yaml
├── proto/              # Protocol Buffer definitions
│   └── v1/             # Version 1 schemas
│       ├── event.proto
│       ├── ledger.proto
│       ├── advisory.proto
│       ├── enforcement.proto
│       ├── trace.proto
│       ├── truth.proto
│       ├── validation.proto
│       ├── distributed.proto
│       ├── memory.proto
│       └── simulation.proto
├── schemas/            # JSON / Avro / other schema formats
│   └── v1/
├── tla/                # TLA+ formal specifications
├── versioning/         # Versioning rules and compatibility matrix
│   ├── versioning.md
│   └── compatibility-matrix.md
├── invariants/         # Architecture invariants
│   └── architecture-invariants.md
└── examples/           # Usage examples and reference implementations
```

## Core Principles

1. **Determinism** — Every state transition is deterministic and reproducible.
2. **Evidence-Based** — All decisions require cryptographic evidence chains.
3. **Append-Only Ledger** — State history is immutable and auditable.
4. **Zero-Trust** — Every component must prove its claims via evidence.
5. **Replay-Safe** — All operations can be replayed from any checkpoint.

## Contract Guarantees

- All messages include `schema_version`, `created_at`, `updated_at`, `source_repo`, `replay_sequence`, and `validation_hash` mandatory fields.
- Proto definitions follow `proto3` syntax with explicit package namespacing under `phoenix.*`.
- FSM contracts define complete state graphs with guards, rollback paths, and forbidden transitions.
- The event bus contract specifies topic priorities, routing, dead-letter, and retention rules.

## Versioning

See [versioning/versioning.md](versioning/versioning.md) for the full versioning policy.
See [versioning/compatibility-matrix.md](versioning/compatibility-matrix.md) for the compatibility matrix.

## Invariants

See [invariants/architecture-invariants.md](invariants/architecture-invariants.md) for the 12 non-negotiable architecture invariants.

## License

PhoenixCore is part of the PhoenixOS ecosystem. All contracts defined here are canonical and authoritative.

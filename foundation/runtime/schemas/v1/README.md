# PhoenixOS Canonical Schemas (v1)

This directory contains the canonical JSON and YAML schema definitions used for run-time validation of event payloads, configuration files, and state serialization formats across the PhoenixOS ecosystem.

## Relationship to Protobuf

While wire-format serialization utilizes the Protobuf structures defined in `proto/v1/`, runtime validation of configuration files (e.g., policy definitions, FSM mappings, and routing topologies) is performed against JSON Schema draft-07 compatible JSON/YAML schemas.

All JSON/YAML schemas are generated deterministically from the canonical Protobuf structures to ensure zero divergence.

## Directory Structure

```
schemas/v1/
├── README.md             # This document
├── event-envelope.json   # Event payload schema
├── fsm-policy.json       # Bounded FSM configuration schema
├── routing-topology.json # Event-bus routing map schema
└── ledger-checkpoint.json# Ledger verification checkpoint schema
```

## Mandatory Fields for Schema Payloads

To maintain compatibility with `proto/v1/`, all JSON/YAML structures must implement:
- `schema_version`: Strict SemVer string (e.g., `"1.0.0"`).
- `created_at`: RFC3339 formatted UTC timestamp string.
- `updated_at`: RFC3339 formatted UTC timestamp string.
- `source_repo`: Absolute identifier of the originating repository.
- `replay_sequence`: Monotonic unsigned 64-bit sequence integer.
- `validation_hash`: Cryptographic checksum (SHA-256) of the validated payload.

## Validation Strategy

1. **Static Validation**: Schemas are checked during compile time in `PhoenixValidation`.
2. **In-Flight Validation**: The `Warden` and `Guard` enforce schema conformity on all event envelopes prior to state machine transitions.
3. **Replay Validation**: Replay validation enforces that historical trace events match their schema definitions.

# PhoenixOS Compatibility Matrix

This document tracks the compatibility status of all shared contracts, APIs, and FSM definitions across the 18 repositories of the PhoenixOS ecosystem.

## Canonical Compatibility Matrix

| Contract / Schema | Primary Producer | Primary Consumer(s) | Target Version | Compatibility Level | Replay Compatibility |
|:---|:---|:---|:---|:---|:---|
| `event.proto` | All Subsystems | `PhoenixTrace`, `PhoenixValidation` | `v1.0.0` | Backward Compatible | Mandatory |
| `ledger.proto` | `PhoenixDistributed` | `PhoenixGuard`, `PhoenixValidation` | `v1.0.0` | Fully Immutable | Mandatory |
| `advisory.proto` | `PhoenixMind` | `PhoenixGuard`, `PhoenixDashboard` | `v1.0.0` | Backward Compatible | Optional |
| `enforcement.proto` | `PhoenixGuard` | `PhoenixKernel`, `PhoenixTrace` | `v1.0.0` | Strict Lockstep | Mandatory |
| `trace.proto` | `PhoenixTrace` | `PhoenixMind`, `PhoenixValidation` | `v1.0.0` | Backward Compatible | Mandatory |
| `truth.proto` | `PhoenixTruth` | `PhoenixMind`, `PhoenixGuard` | `v1.0.0` | Backward Compatible | Mandatory |
| `validation.proto` | `PhoenixValidation` | `PhoenixFormal`, `PhoenixResearch` | `v1.0.0` | Backward Compatible | Mandatory |
| `distributed.proto` | `PhoenixDistributed` | `PhoenixCore`, `PhoenixGuard` | `v1.0.0` | Backward Compatible | Mandatory |
| `memory.proto` | `PhoenixMemoryLab` | `PhoenixMind` | `v1.0.0` | Backward Compatible | Optional |
| `simulation.proto` | `PhoenixStimulation` | `ParticleStimulator`, `PhoenixRedteam` | `v1.0.0` | Backward Compatible | Mandatory |

## Compatibility Class Definitions

- **Fully Immutable**: This structure cannot be changed in any minor or patch release. Any change necessitates a major version fork (e.g. `v2`).
- **Strict Lockstep**: The producer and consumer must be updated simultaneously. Mismatches will trigger an immediate system halt.
- **Backward Compatible**: Consumers can process older versions of the message without modification or data loss. Minor extensions are allowed.

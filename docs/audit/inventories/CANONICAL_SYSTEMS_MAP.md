# Canonical Systems Map

> **Authority Phase**: 4B.2 Canonical Ownership & Extraction Blockers
> **Status**: AUTHORITATIVE
> **Last Updated**: 2026-06-04

This document defines the single source of truth for every major subsystem implementation. All shadow implementations are secondary and should not be imported by production code.

## 1. Subsystem Authority

| Subsystem | Canonical Implementation | Shadow (Simulation) | Archive (Legacy) |
| :--- | :--- | :--- | :--- |
| **Ledger** | `foundation/ledger/` | `platform/crucible/PhoenixStimulation/phoenix_os/ledger/` | `labs/archive/nucleus/Phoenix.Nucleus/ledger/` |
| **Authority** | `foundation/runtime/authority/` | (None) | `labs/archive/nucleus/Phoenix.Nucleus/authority/` |
| **Kernel** | `foundation/runtime/kernel/` | `platform/crucible/PhoenixStimulation/phoenix_os/kernel/` | (None) |
| **Security** | `assurance/security/` | `platform/crucible/PhoenixStimulation/phoenix_os/security/` | (None) |
| **Simulation** | `platform/crucible/simulation/` | `platform/crucible/game/simulation/` | `labs/archive/terminus/` |
| **Mind** | `cognition/mind/` | `platform/os/core/pheonixmind-core/` | `platform/os/_legacy_archive/` |

## 2. Implementation Truth Rules

1. **Production Imports**: Production code (`foundation`, `assurance`, `governance`) MUST only import from **Canonical** implementations.
2. **Shadow Policy**: Shadow implementations exist solely for adversarial testing and high-entropy simulation. They MUST NOT be exposed as public APIs for the substrate.
3. **Typo Resolution**: `pheonixmind-core` is identified as a typo-drift of `phoenixmind`. The correct spelling `phoenix` is mandatory for all new modules.

---
**Strategic Mandate**: Converge all active development on Canonical paths. Shadow systems are frozen for production use.

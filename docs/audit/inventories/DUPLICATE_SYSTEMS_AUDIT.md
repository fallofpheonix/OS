# Duplicate Systems Audit

> **Authority Phase**: 4A.11 Duplicate Systems Audit
> **Status**: AUDIT COMPLETE
> **Last Updated**: 2026-06-04

This document inventories duplicate or shadow implementations across the PhoenixOS repository.

## 1. Top-Level Duplicates

| System A | System B | Analysis | Recommendation |
| :--- | :--- | :--- | :--- |
| `platform/crucible/PhoenixVirtualization` | `platform/crucible/PhoenixVirtualizer` | `Virtualizer` is a refactored/flattened version. `Virtualization` contains older shell scripts and logs. | **CONSOLIDATE**: Migrate unique scripts from `Virtualization` to `Virtualizer` and archive the former. |
| `foundation/ledger` | `platform/crucible/PhoenixStimulation/phoenix_os/ledger` | `foundation` is the production ledger. `PhoenixStimulation` contains a legacy `truth_ledger` module. | **SHADOW**: Identify if simulation requires the legacy ledger; if not, point simulation to `foundation/ledger`. |
| `foundation/runtime/authority` | `labs/archive/nucleus/Phoenix.Nucleus/authority` | Identical filenames. `foundation` is the current production version. | **LEGACY**: The archived version is a snapshot of the production version. |

## 2. Simulation Redundancy

Crucible contains multiple "Simulation" engines with overlapping responsibilities:

- **ParticleStimulator**: Legacy simulator focused on "Particle" (atomic event) stimulation.
- **PhoenixSimulation**: General purpose simulation framework.
- **game/simulation**: Specific to the WARDEN.EXE gameplay loop.

**Conflict**: There is no "Unified Simulation Contract". Each sub-system implements its own Event Bus and State Reconciliation.

## 3. Implementation Shadows

| Concept | Production | Shadow |
| :--- | :--- | :--- |
| **Kernel Probes** | `foundation/runtime/kernel` | `platform/crucible/PhoenixStimulation/phoenix_os/kernel` |
| **Security Rules** | `assurance/security` | `platform/crucible/PhoenixStimulation/phoenix_os/security` |
| **Trust Model** | `governance/truth` | `platform/crucible/game/engines/trust.go` |

## 4. Conclusion

The repository currently suffers from "System Mirroring," where research assets (`PhoenixStimulation`) contain a full shadow copy of the production architecture. This creates significant confusion for both human operators and AI agents when searching for canonical implementations.

**Strategic Mandate**: Any shadow implementation must be explicitly marked as `SHADOW` or `LEGACY`.

---
**Next Step**: Phase 4A.12 Final Synthesis & Strategy Update.

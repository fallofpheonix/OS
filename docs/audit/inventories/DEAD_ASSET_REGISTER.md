# Dead Asset Register

> **Authority Phase**: 4A.8 Dead Asset Register
> **Status**: INVENTORY ONLY
> **Last Updated**: 2026-06-04

This document inventories all identified dead or redundant assets in the PhoenixOS repository. No deletions are performed in this phase.

## 1. Unused Packages (Go)

| Package Path | Layer | Notes |
| :--- | :--- | :--- |
| `platform/crucible/PhoenixSimulation/...` | Crucible | 14 Go files. Not in `go.work`. |
| `platform/crucible/PhoenixAudit/...` | Crucible | Has `go.mod`. Not in `go.work`. |
| `platform/crucible/PhoenixStimulation/...` | Crucible | 507 Go files. Used via `replace` but not as workspace module. Shadow monorepo. |
| `platform/crucible/ParticleStimulator/...` | Crucible | Not in `go.work`. |
| `platform/crucible/PhoenixChampions/...` | Crucible | 7 Go files. Not in `go.work`. |
| `platform/crucible/PhoenixVirtualization/...` | Crucible | 1 Go file. Duplicate implementation. |
| `platform/crucible/PhoenixVirtualizer/...` | Crucible | Python/Node focus. Tools reference only. |

## 2. Restored Assets (Corrected during Audit)

| Asset | Action | Rationale |
| :--- | :--- | :--- |
| `platform/cli` | Created `go.mod` & added to `go.work` | Resolved "Lost Package" status; fixed broken imports to `crucible`. |
| `platform/crucible/**/go.mod` | Rewrote `labs/crucible` -> `platform/crucible` | Resolved "Legacy Compatibility Layer" debt. |

## 3. Unused Binaries

| Binary | Location | Notes |
| :--- | :--- | :--- |
| (None) | `bin/` | Only `phoenixd` remains. |

## 4. Unused Modules (Cross-Language)

| Module | Location | Notes |
| :--- | :--- | :--- |
| `nucleus` | `labs/archive/nucleus` | Legacy core experiments. |
| `terminus` | `labs/archive/terminus` | Legacy UI/CLI research assets. |
| `g0dm0d3` | `labs/archive/g0dm0d3` | Root/Debug utilities. |
| `research` | `labs/archive/research` | Speculative AI/ML research. |

## 5. Unused Documentation

| Document | Location | Notes |
| :--- | :--- | :--- |
| `MASTER_DOCUMENTATION_INDEX.md` | `docs/archive/` | Superseded by Authority Matrix. |
| `DEAD_CODE_REGISTER.md` | `docs/audit/` | Paths refer to non-existent `/core` structure. |
| `EXTRACTION_ORDER.md` | `docs/archive/` | Superseded by Master Roadmap. |
| `GAME_STATUS.md` | `docs/archive/` | Legacy gameplay spec. |

## 6. Duplicate Implementations

| System A | System B | Conflict |
| :--- | :--- | :--- |
| `platform/crucible/PhoenixVirtualization` | `platform/crucible/PhoenixVirtualizer` | Identical READMEs. Shadow implementations of 3D dependency visualization. |
| `foundation/ledger` | `platform/crucible/PhoenixStimulation/phoenix_os/ledger` | Production ledger vs. Shadow/Simulation ledger. |
| `foundation/runtime/authority` | `labs/archive/nucleus/Phoenix.Nucleus/authority` | Production authority vs. Legacy authority. |
| `platform/crucible/simulation` | `platform/crucible/game/simulation` | Generic simulation logic vs. Game-specific simulation logic. |

## 7. Abandoned Experiments

| Experiment | Location | Notes |
| :--- | :--- | :--- |
| `DeepSpeed`, `IsaacLab`, `ml-agents` | `labs/archive/terminus/.../research/` | Massive external repos imported for research. High entropy. |
| `PhoenixStimulation/archive/research/rejected/` | `platform/crucible/...` | Explicitly marked as rejected research. |

## 8. Legacy Compatibility Layers

| Layer | Location | Notes |
| :--- | :--- | :--- |
| `platform/crucible/DIRECTORY_NOTES.md` | `platform/crucible/` | References "Labs/Crucible" instead of "Platform/Crucible". |

**Conclusion**: Repository inventory complete. Ready for systematic cleanup in Phase 4B.

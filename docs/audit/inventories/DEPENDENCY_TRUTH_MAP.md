# Dependency Truth Map

> **Authority Phase**: 4A.9 Dependency Truth Map
> **Status**: AUDIT COMPLETE
> **Last Updated**: 2026-06-04

This document maps the authoritative dependency graph of the PhoenixOS substrate, identifying boundary violations and coupling.

## 1. Architectural Layer Enforcement

| Layer | Authority | Boundary Violations |
| :--- | :--- | :--- |
| **Foundation** | High | **CRITICAL**: `runtime/adapters` imports Assurance layer. |
| **Assurance** | Medium | None detected. |
| **Governance** | Medium | None detected. |
| **Cognition** | Low | None detected. |
| **Platform** | Zero | High coupling to Foundation and Crucible. |

## 2. Critical Boundary Violations

### VIOLATION-001: Runtime -> Assurance
- **Source**: `foundation/runtime/adapters` (RESOVLED: Moved to `platform/os/adapters`)
- **Impact**: RESOLVED. Runtime now builds with zero assurance imports.

## 2.5. Consensus Rules (T4.2)
- **Fork-Choice Rule**: Highest sequence number wins on reconnect. If sequences are tied, local state is preserved to favor liveness.

## 3. Crucible Coupling Audit

The `crucible` module (currently in `platform/`) is a "High-Trust Verification" asset with the following dependencies:
- **Primary Dependencies**: `foundation/ledger`, `foundation/runtime/authority`.
- **Internal Coupling**: High internal coupling between `simulation`, `hypothesis`, and `verification`.
- **External Consumers**: `platform/cli` (WARDEN.EXE) depends on `crucible/game/*` for simulation and adjudicating anomalies.

## 4. Platform (OS/CLI) Coupling

The `platform/os` subproject is the primary consumer of the substrate:
- **Dependencies**: Imports almost all layers (Assurance, Foundation, Governance).
- **Broken Imports**: `platform/os/phoenix_os/ai/feature.go` refers to missing `phoenixmind-core/contracts`.
- **Lost Package**: `platform/cli` is not in `go.work` and has no `go.mod`, but imports `crucible/game`.

## 5. Circular Dependency Check
- **Status**: PASSED. No circular import paths detected in the current workspace.

---
**Next Step**: Phase 4A.10 Crucible Classification Audit.

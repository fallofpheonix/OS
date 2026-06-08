# Extraction Blockers

> **Authority Phase**: 4B.2 Canonical Ownership & Extraction Blockers
> **Status**: ACTIVE REGISTRY
> **Last Updated**: 2026-06-04

This document tracks all identified architectural coupling that prevents the independent extraction of core modules into separate repositories.

## 1. Top-Level Blockers

| ID | Blocker | Source | Target | Impact |
| :--- | :--- | :--- | :--- | :--- |
| **BLOCKER-001** | Reverse Authority | `foundation/runtime/adapters/` | `assurance/security`, `assurance/validation` | Prevents Runtime extraction; violates contract-first layering. |
| **BLOCKER-002** | Path Divergence | `platform/os/phoenix_os/ai/` | `platform/os/phoenixmind-core/` | Broken build due to `pheonix` vs `phoenix` typo/drift. |
| **BLOCKER-003** | System Mirroring | `platform/crucible/PhoenixStimulation/` | `foundation/*`, `assurance/*` | High cognitive load and agent confusion due to shadow implementations. |
| **BLOCKER-004** | Missing Contracts | `platform/os/` | `phoenixmind-core/contracts` | Prevents Platform OS from compiling. |
| **BLOCKER-005** | No SemanticValidator | `foundation/runtime/bus/applier.go` | `ledger.AddEntry()` | RESOLVED. PhysicsValidator implemented and wired into Applier. |
| **BLOCKER-006** | Non-Deterministic RNG | `foundation/runtime/security/control/fsm.go` | `math/rand` | Breaks replay determinism in security paths; stochastic behavior removed pending ledger-sourced random values. |

## 2. Technical Debt Backlog

- **Unused Go Workspace Modules**: `platform/crucible/` subprojects are missing from `go.work` but contain 500+ Go files.
- **Legacy Path Imports**: Many `go.mod` files still reference `labs/crucible` instead of `platform/crucible`. (Partially resolved in 4B.1).

---
**Conclusion**: Blocker-001 (Reverse Authority) is the primary target for Phase 4B.4.

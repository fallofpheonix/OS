# Runtime Extraction Plan

> **Authority Phase**: 4B.3 Shadow System Dependency Audit
> **Status**: PROPOSAL
> **Last Updated**: 2026-06-04

This document defines the steps required to resolve **BLOCKER-001 (Reverse Authority)** and enable the independent extraction of the `foundation/runtime` module.

## 1. Problem Analysis

The `foundation/runtime/adapters` package acts as a "Composition Layer," wrapping concrete implementations from `assurance/` (Security Warden, Replay Engine) to satisfy `foundation/contracts/`. 

**The Violation**: A low-level substrate (`foundation`) is importing a high-level verification suite (`assurance`), creating a circular dependency at the conceptual level and blocking modular extraction.

## 2. Targeted Solution: Adapter Migration

The `adapters` package does not belong in the Foundation layer. It belongs where the system is integrated.

### Proposed Move
- **Source**: `foundation/runtime/adapters/`
- **Destination**: `platform/os/adapters/` (or `integration/adapters/`)

### Impact Analysis
- **Build**: `platform/os/cmd/phoenixd` already imports these adapters. It will need to be updated to the new local path.
- **Extraction**: `foundation/runtime` will no longer depend on `assurance/`. It will only depend on `foundation/contracts`, achieving 100% extraction readiness.

## 3. Execution Steps (Phase 4B.4)

1. **Relocate Code**: Move `foundation/runtime/adapters/*.go` to `platform/os/adapters/`.
2. **Update Namespace**: Change package name from `adapters` to `platform_adapters` (or similar) if required to avoid conflicts.
3. **Rewrite Imports**: Update `platform/os/cmd/phoenixd/main.go` and other callers to point to the new location.
4. **Verification**: Run `go build` and `go test` for both `foundation/runtime` and `platform/os`.

---
**Strategic Conclusion**: Successful relocation of the adapters will promote the Runtime module to **VERIFIED (Extractable)** status.

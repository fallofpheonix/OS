# Feature Status Matrix

> **Authority Phase**: 4B.2 Canonical Ownership & Extraction Blockers
> **Status**: UPDATED
> **Last Updated**: 2026-06-04

This matrix defines the functional state of the PhoenixOS substrate, verified via `go build` and `go test`.

## 1. Core Substrate (Foundation & Assurance)

| Subsystem | Compiles | Tests Pass | Extraction Ready | Status | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Ledger** | YES | YES | YES | **VERIFIED** | Verified via `TestEndToEndPersistenceAndRecovery`. |
| **Runtime (Core)** | YES | YES | **NO** | **FUNCTIONAL** | Blocked by Blocker-001 (Runtime -> Assurance). |
| **Security (Guard)** | YES | YES | YES | **VERIFIED** | Fast-path enforcement active; `CheckQuorum` hardened. |
| **Validation** | YES | YES | YES | **VERIFIED** | Vertical slice proven from Kernel to Disk. |
| **Vertical Slice** | YES | YES | YES | **PROVEN** | `eBPF -> Bus -> Applier -> Persistor -> Disk -> Replay`. |
| **Kernel (Probes)** | YES | YES | YES | **VERIFIED** | eBPF probes wired to Bus via bootstrap. |
| **Distributed** | YES | YES | NO | **EXPERIMENTAL** | Gossiping authenticated events; voting planned. |

## 2. Intelligence & Governance

| Subsystem | Compiles | Tests Pass | Integrated | Status | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Cognition (Root)** | YES | YES | NO | **EXPERIMENTAL** | Foundational memory/knowledge graph. |
| **PhoenixMind** | YES | YES | YES | **PARTIAL** | Integrated into Platform OS; path drift resolved. |
| **Arbiter** | YES | NO | NO | **UNVERIFIED** | Logic exists but lacks validation. |
| **Truth** | YES | NO | NO | **UNVERIFIED** | Logic exists but lacks validation. |

## 3. Interfaces & Simulation

| Subsystem | Compiles | Tests Pass | Integrated | Status | Notes |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Platform OS** | YES | NO | YES | **FUNCTIONAL** | Now compiles after PhoenixMind path correction. |
| **WARDEN CLI** | YES | N/A | YES | **VERIFIED** | Stable Auditor Interface. |
| **Crucible (Core)** | YES | YES | YES | **VERIFIED** | Simulation engines functional. |
| **Simulation Lab** | NO | NO | NO | **SHADOW** | `PhoenixStimulation` is a shadow monorepo; frozen. |

## 4. Verification Legend
- **VERIFIED**: Compiles, tests pass, and integrated into primary flow.
- **FUNCTIONAL**: Compiles and passes tests, but has boundary or extraction issues.
- **PARTIAL**: Implementation exists and mostly works, but lacks full coverage.
- **EXPERIMENTAL**: Prototyped but unproven stability/integration.
- **UNVERIFIED**: Code exists but functional status is unknown (no tests/integration).
- **SHADOW**: Duplicate implementation; not for production use.
- **BROKEN**: Fails build or critical integration tests.

---
**Strategic Target**: Resolving Blocker-001 to promote Runtime from FUNCTIONAL to VERIFIED (Extractable).

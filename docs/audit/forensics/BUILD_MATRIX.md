# Phoenix Matrix: Build Matrix

**Status:** RECOVERY AUDIT
**Date:** 2026-06-05

This matrix tracks the buildability and testability of every module in the monorepo.
Rule: **UNKNOWN = BROKEN** until proven otherwise.

| Module | Build | Test | Race | Status | Notes |
| :--- | :---: | :---: | :---: | :--- | :--- |
| `foundation/contracts` | PASS | PASS | PASS | **PROVEN** | Base contracts. |
| `foundation/events` | PASS | PASS | PASS | **PROVEN** | Event definitions. |
| `foundation/ledger` | PASS | PASS | PASS | **PROVEN** | Durable ledger. |
| `foundation/math` | PASS | PASS | PASS | **PROVEN** | Fixed-point math. |
| `foundation/observability` | PASS | PASS | PASS | **PROVEN** | Logging/Metrics. |
| `foundation/runtime` | PASS | PASS | PASS | **PROVEN** | Execution substrate. |
| `foundation/runtime/kernel` | PASS | PASS | PASS | **PROVEN** | eBPF capture pipeline. |
| `foundation/security/identity` | PASS | PASS | PASS | **PROVEN** | Validator identity. |
| `governance/arbiter` | PASS | PASS | PASS | **PROVEN** | Policy adjudication. |
| `governance/truth` | PASS | PASS | PASS | **PROVEN** | Evidence interpretation. |
| `platform/cli` | PASS | PASS | PASS | **PROVEN** | WARDEN.EXE interface. |
| `platform/ui/Service` | PASS | PASS | PASS | **PROVEN** | UI backend. |
| `platform/os` | PASS | PASS | PASS | **PROVEN** | Phoenix Matrix Daemon. |
| `game` | PASS | PASS | PASS | **PROVEN** | Simulation core. |
| `assurance/security` | PASS | PASS | PASS | **PROVEN** | Warden FSM (RESTORED). |
| `assurance/validation` | PASS | FAIL | FAIL | **WORKING** | Legacy tests. |
| `foundation/distributed` | ARCHIVED | - | - | **ARCHIVED** | Speculative consensus. |
| `platform/crucible` | ARCHIVED | - | - | **ARCHIVED** | Speculative test harness. |
| `cognition` | ARCHIVED | - | - | **ARCHIVED** | Experimental AI. |

---
*Authorized by Phoenix Sovereign Governance*

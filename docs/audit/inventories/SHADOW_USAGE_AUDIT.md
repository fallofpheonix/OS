# Shadow Usage Audit

> **Authority Phase**: 4B.3 Shadow System Dependency Audit
> **Status**: AUDIT COMPLETE
> **Last Updated**: 2026-06-04

This document maps all production and platform dependencies on shadow or legacy implementations.

## 1. Primary Shadow Dependencies

| Consumer | Shadow Asset | Type | Impact |
| :--- | :--- | :--- | :--- |
| `platform/cli` (WARDEN.EXE) | `crucible/game/` | Import | Direct coupling to gamified simulation engines. |
| `platform/os` (Phoenix OS) | `crucible/PhoenixSimulation` | `replace` rule | Pulls in shadow simulation for testing/integration. |
| `assurance/security/...` | `physics`, `entropy_engine` | Shadow `replace` | Legacy shadow-style imports in `integrated_model` and `physics`. |
| `platform/crucible/PhoenixRedteam` | `PhoenixStimulation/stress` | Import | Depends on shadow monorepo for scenario execution. |

## 2. Shadow Monorepo (PhoenixStimulation)

The `PhoenixStimulation` module is a **Complete Shadow Substrate**. It contains internal copies of:
- **Ledger**: `phoenix_os/ledger`
- **Kernel**: `phoenix_os/kernel`
- **Security**: `phoenix_os/security`
- **Simulation**: `experimental/simulation_lab`

**Audit Finding**: `PhoenixStimulation` imports `foundation/runtime/bus` but otherwise operates as a parallel universe. No production code imports `PhoenixStimulation` packages directly, except for `replace` rules in platform modules.

## 3. Replace Directive Audit

| Module | Shadow Replace | Status |
| :--- | :--- | :--- |
| `platform/os/go.mod` | `PhoenixSimulation` -> `platform/crucible/PhoenixSimulation` | ACTIVE |
| `platform/cli/go.mod` | `platform/crucible` -> `../crucible` | ACTIVE |
| `assurance/security/integrated_model/go.mod` | `phoenix/security/physics` -> `../physics` | ACTIVE (Internal) |

---
**Strategic Conclusion**: Shadow implementations are isolated behind `replace` rules and limited to the Platform layer. The Foundation and Assurance layers are 95% clean of shadow imports, with the exception of legacy models in `assurance/security/integrated_model`.

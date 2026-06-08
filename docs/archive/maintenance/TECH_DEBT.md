---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS: Technical Debt & Future Substrate

This document tracks documented debt and architectural optimizations identified after the v1.0 sovereign delivery.

## Resolved Debt (v1.0 Delivery)

| ID | Category | Description | Status |
| :--- | :--- | :--- | :--- |
| **TD-001** | Architecture | No root README or consolidated entry point. | **FIXED** |
| **TD-002** | Maintenance | Opaque repository structure with orphan files. | **FIXED** |
| **TD-003** | Security | Unverified boot sequence and implicit authority. | **FIXED** |
| **TD-004** | Truth | Non-deterministic state reconstruction. | **FIXED** |
| **TD-005** | Build | Incomplete documentation pass broke core compilation. | **FIXED** |

## Active Substrate Debt

| ID | Category | Priority | Location | Description | Status |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **TD-101** | Performance | `core/Nucleus` | Needs HDF5 vector optimizations for high-density event streams. | pending |
| **TD-102** | Formal | `core/Formal` | Expand TLA+ models to include full Byzantine subversion scenarios. | research |
| **TD-103** | Concurrency | `core/Bus` | Verify cross-domain memory mappings at 100k+ TPS. | pending |
| **TD-104** | Abstraction | `core/Cognition`| Tight coupling between Causal DAG and SQLite-Vec needs generic persistence interface. | wip |
| **TD-105** | Security | `core/Security` | Critical security interfaces (ModelGuard, PromptGuard, ThreatDetector) are implemented as skeletons; require concrete logic integration. | wip |

## Architectural Risks

1.  **Hardware Side-Channels:** Bit-perfect replay assumes perfect logic; however, timing side-channels on bare-metal are currently out of scope for the software-defined constitution.
2.  **Federation Scale:** Multi-node proof exchange is verified for small clusters (3-5 nodes) but unproven for 100+ nodes.

## Maintenance Roadmap

1.  **[P0]** Continuous Formal Verification (Scaling TLC checker).
2.  **[P1]** Implementation of `honest-failure` static analyzer.
3.  **[P2]** HDF5 Vector integration for forensic-scale storage.

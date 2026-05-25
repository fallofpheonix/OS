# Component Map: The Phoenix Matrix

This document maps directories and packages in the Go workspace to their active system layers and integrated sub-repositories.

---

## 1. Sub-Repository Structure

PhoenixOS is composed of 4 integrated sub-repositories under `03_repositories/integrated/`.

| Sub-Repository | Go Module Path | Layer Mapping | Description |
| :--- | :--- | :--- | :--- |
| **`phoenix-contracts`** | `github.com/fallofpheonix/phoenix-contracts` | Cross-Layer | Core interface definitions, event schemas, versioning matrices. |
| **`phoenix-runtime`** | `github.com/fallofpheonix/phoenix-runtime` | **L1: Substrate** | eBPF probes, syscall bindings, fast-path Guard adaptor. |
| **`phoenix-logic`** | `github.com/fallofpheonix/phoenix-logic` | **L2: Truth** | Telemetry ingestion, normalizer, replay engine, and TruthLedger. |
| **`phoenix-control`** | `github.com/fallofpheonix/phoenix-control` | **L3: Control** | Warden state FSM, State Registry, and cost-benefit Arbiter. |

---

## 2. Directory Layout & Package Mappings

### `phoenix_os/` Root Subsystems
Ties the core subsystems together to run as a single deterministic userland binary (Go PID1).

- **`ai/`**: Advisory-only interface wrappers for local/cloud LLMs.
- **`boot/`**: System boot sequencing and integrity checksum calculation.
- **`bus/`**: Priority-queue telemetry event bus and normalizer.
- **`common/`**: Shared core primitives:
  - `logical_clock/`: Thread-safe monotonic clock.
  - `resource/`: Bounded memory and worker pool allocators.
  - `serialization/`: Canonical JSON and Bencode implementations.
  - `snapshot/`: Graph snapshot utilities.
- **`game/`**: Simulated SOC game server used in the Determinism Lab to trigger threat scenarios.
- **`gov/`**: Local repository validator for validating axioms, import rules, and documentation.
- **`ledger/`**: High-level wrapper matching the `truth.TruthLedger` API.

---

## 3. Subsystem Dependency Rules

> [!WARNING]
> To preserve the Phase F0 lock, no core logic packages (`truth`, `warden`, `arbiter`) may import `ai`, `nexus`, or `memory`. This rule is enforced at compile time via `05_tools/validate_imports.py`.

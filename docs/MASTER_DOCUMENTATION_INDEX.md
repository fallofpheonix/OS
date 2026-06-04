---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Master Documentation Index (v1.0)

## 1. Core Architecture
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`MASTER_SPECIFICATION.md`](./architecture/MASTER_SPECIFICATION.md)** | Contract-First System Design & Flow | YES | CANONICAL |
| **[`FINAL_ARCHITECTURE.md`](./architecture/FINAL_ARCHITECTURE.md)** | Authoritative Architecture Direction | YES | CANONICAL |
| **[`ARCHITECTURE_REPLACEMENT_PLAN.md`](./architecture/ARCHITECTURE_REPLACEMENT_PLAN.md)** | Executable Replacement Roadmap | YES | CANONICAL |
| **[`ARCHITECTURE_DEBT.md`](./architecture/ARCHITECTURE_DEBT.md)** | Boundary Debt Register | YES | CANONICAL |
| **[`FRACTAL_CYCLES_AND_LAWS.md`](./architecture/FRACTAL_CYCLES_AND_LAWS.md)** | Core Axioms | YES | CANONICAL |

## 2. Governance & Rules
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`MASTER_INVARIANTS.md`](./MASTER_INVARIANTS.md)** | Absolute System Constraints | YES | CANONICAL |
| **[`PhoenixOS-Constitution-v1.md`](./governance/PhoenixOS-Constitution-v1.md)** | Executive Invariants| YES | CANONICAL |
| **[`BOUNDARY_RULES.md`](./governance/BOUNDARY_RULES.md)** | Layering & Autonomy Rules | YES | CANONICAL |
| **[`DEPENDENCY_RULES.md`](./governance/DEPENDENCY_RULES.md)** | Import & Dependency Laws | YES | CANONICAL |
| **[`LAYERING_RULES.md`](./governance/LAYERING_RULES.md)** | Strict Layering Architecture | YES | CANONICAL |
| **[`EXTRACTION_RULES.md`](./governance/EXTRACTION_RULES.md)** | Modularization Requirements | YES | CANONICAL |
| **[`AI_AUTONOMY_RULES.md`](./governance/AI_AUTONOMY_RULES.md)** | AI Safety & Limits | YES | CANONICAL |
| **[`REPOSITORY_CONSTITUTION.md`](./governance/REPOSITORY_CONSTITUTION.md)** | Engineering Rules | YES | CANONICAL |
| **[`GLOSSARY.md`](./governance/GLOSSARY.md)** | Universal Terminology | YES | CANONICAL |

## 3. Failure & Recovery
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`FAILURE_CATALOG.md`](./failure/FAILURE_CATALOG.md)** | Failure Modes & Severity | YES | CANONICAL |
| **[`CATASTROPHIC_FAILURES.md`](./failure/CATASTROPHIC_FAILURES.md)** | S1 Failure Protocols | YES | CANONICAL |
| **[`RECOVERY_PATHS.md`](./failure/RECOVERY_PATHS.md)** | Automated Recovery Procedures | YES | CANONICAL |
| **[`ROLLBACK_MODEL.md`](./failure/ROLLBACK_MODEL.md)** | State Rollback Mechanism | YES | CANONICAL |
| **[`CHAOS_TEST_PLAN.md`](./failure/CHAOS_TEST_PLAN.md)** | Resilience Stress Testing | YES | CANONICAL |
| **[`SYSTEM_STATUS.md`](./status/SYSTEM_STATUS.md)** | Real-time Health | YES | CANONICAL |

## 4. Roadmaps & Status
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`MASTER_PHOENIX_ROADMAP.md`](./roadmap/MASTER_PHOENIX_ROADMAP.md)** | Execution Phases | YES | CANONICAL |
| **[`EXTRACTION_ORDER.md`](./extraction/EXTRACTION_ORDER.md)** | Module Extraction Roadmap | YES | CANONICAL |
| **[`GAME_STATUS.md`](./game/GAME_STATUS.md)** | WARDEN.EXE Maturity | YES | CANONICAL |

## 5. Formal Specifications (SPEC)
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`SPEC-001-System-State.md`](./specifications/SPEC-001-System-State.md)** | Hierarchical System State | YES | CANONICAL |
| **[`SPEC-002-Runtime-State.md`](./specifications/SPEC-002-Runtime-State.md)** | Virtual Memory & VM State | YES | CANONICAL |
| **[`SPEC-003-Ledger-State.md`](./specifications/SPEC-003-Ledger-State.md)** | Block Structure & Consensus | YES | CANONICAL |
| **[`SPEC-004-Memory-State.md`](./specifications/SPEC-004-Memory-State.md)** | Truth Model & Fact Schema | YES | CANONICAL |
| **[`SPEC-005-Governance-State.md`](./specifications/SPEC-005-Governance-State.md)** | Policy Engine State | YES | CANONICAL |
| **[`SPEC-006-Intent-State.md`](./specifications/SPEC-006-Intent-State.md)** | Intent & Actuation Schema | YES | CANONICAL |
| **[`SPEC-007-Capability-Model.md`](./specifications/SPEC-007-Capability-Model.md)** | Security & Token Model | YES | CANONICAL |
| **[`EVENT_LIFECYCLE_SPEC.md`](./specifications/EVENT_LIFECYCLE_SPEC.md)** | Schema & Evolution | YES | CANONICAL |
| **[`PSCRIPT_SEMANTICS.md`](./specifications/PSCRIPT_SEMANTICS.md)** | P-Script Language Laws | YES | CANONICAL |

## 6. Verification & Traceability
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`TRACEABILITY_MATRIX.md`](./verification/TRACEABILITY_MATRIX.md)** | Invariants -> Specs -> Tests -> Recovery | YES | CANONICAL |
| **[`CLAIMS_AUDIT.md`](./verification/CLAIMS_AUDIT.md)** | Architectural Claim Classification | YES | CANONICAL |
| **[`INVARIANT_TO_CODE_MATRIX.md`](./verification/INVARIANT_TO_CODE_MATRIX.md)** | Invariant -> Code Package Mapping | YES | CANONICAL |
| **[`SPECIFICATION_COVERAGE_MATRIX.md`](./verification/SPECIFICATION_COVERAGE_MATRIX.md)** | Spec Rule Implementation Status | YES | CANONICAL |
| **[`INVARIANT_VERIFICATION_MATRIX.md`](./verification/INVARIANT_VERIFICATION_MATRIX.md)** | Invariant-to-Test Mapping | YES | CANONICAL |
| **[`FAILURE_TO_TEST_MATRIX.md`](./verification/FAILURE_TO_TEST_MATRIX.md)** | Error-to-Recovery Mapping | YES | CANONICAL |
| **[`VERIFICATION_MANDATE.md`](./specifications/VERIFICATION_MANDATE.md)** | AC & Requirements | YES | CANONICAL |

## 7. Theoretical Foundations
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`CONFIDENCE_MATH.md`](./phoenixmind/theory/CONFIDENCE_MATH.md)** | Cognition Mathematical Models | YES | CANONICAL |
| **[`REWARD_MODEL.md`](./phoenixmind/theory/REWARD_MODEL.md)** | RL & Task Reward Functions | YES | CANONICAL |
| **[`MEMORY_SCORING.md`](./phoenixmind/theory/MEMORY_SCORING.md)** | Memory Lifecycle Scoring | YES | CANONICAL |
| **[`TASK_PRIORITIZATION.md`](./phoenixmind/theory/TASK_PRIORITIZATION.md)** | Cognitive Scheduling | YES | CANONICAL |
| **[`DRIFT_DETECTION.md`](./phoenixmind/theory/DRIFT_DETECTION.md)** | Divergence Correction | YES | CANONICAL |
| **[`SELF_EVOLUTION_GATING.md`](./phoenixmind/theory/SELF_EVOLUTION_GATING.md)** | Self-Modification Safety | YES | CANONICAL |
| **[`WORLD_STATE.md`](../game/docs/specifications/WORLD_STATE.md)** | Formal Game World State | YES | CANONICAL |
| **[`REPLAY_SPEC.md`](../game/docs/specifications/REPLAY_SPEC.md)** | Game Replay Laws | YES | CANONICAL |
| **[`SAVE_FORMAT.md`](../game/docs/specifications/SAVE_FORMAT.md)** | Binary Save Schema | YES | CANONICAL |
| **[`NETWORK_PROTOCOL.md`](../game/docs/specifications/NETWORK_PROTOCOL.md)** | Sync & Comms Protocol | YES | CANONICAL |
| **[`SCORING_SPEC.md`](../game/docs/specifications/SCORING_SPEC.md)** | Competitive Metrics | YES | CANONICAL |
| **[`LEVEL_SPEC.md`](../game/docs/specifications/LEVEL_SPEC.md)** | Level & Environment Data | YES | CANONICAL |
| **[`MISSION_SPEC.md`](../game/docs/specifications/MISSION_SPEC.md)** | Objective Lifecycle | YES | CANONICAL |

## 7. Maintenance & Inventory
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`DOCUMENT_INVENTORY.md`](./inventory/DOCUMENT_INVENTORY.md)** | File Tracking | YES | CANONICAL |
| **[`FEATURE_STATUS_MATRIX.md`](./inventory/FEATURE_STATUS_MATRIX.md)** | Verification Audit | YES | CANONICAL |
| **[`CHANGELOG.md`](./maintenance/CHANGELOG.md)** | Release History | YES | ACTIVE |

---
**Last Verification Date:** 2026-06-03
**Total Documents Audited:** 48

## 6. Architecture Maps
| Document | Purpose | Canonical Source | Status |
| :--- | :--- | :--- | :--- |
| **[`MASTER_SYSTEM_MAP.md`](./MASTER_SYSTEM_MAP.md)** | Layer Hierarchy & Dependency Graph | YES | CANONICAL |
| **[`MASTER_DEPENDENCY_MAP.md`](./MASTER_DEPENDENCY_MAP.md)** | Machine-Readable Dependency Matrix | YES | CANONICAL |
| **[`MASTER_SUBPROJECT_INDEX.md`](./MASTER_SUBPROJECT_INDEX.md)** | Complete Subproject Inventory | YES | CANONICAL |

## 7. Per-Subproject Documentation
Every active subproject maintains its own `docs/` directory with:
- `README.md` — Purpose, ownership, quick start
- `SYSTEM_MAP.md` — Components and relationships
- `CURRENT_STATE.md` — Implementation status
- `TARGET_STATE.md` — Future goals
- `MIGRATION_PATH.md` — Steps from current to target
- `EXTRACTION_READINESS.md` — Readiness for independent release

See [MASTER_SUBPROJECT_INDEX.md](./MASTER_SUBPROJECT_INDEX.md) for links to all subproject docs.

---
**Last Verification Date:** 2026-06-04
**Total Documents Audited:** 140+

# Phoenix Documentation Authority Matrix

> **Authority Phase**: 4A.7 Labs Classification & Restoration Audit
> **Status**: IN-PROGRESS (Restoration Complete)
> **Last Updated**: 2026-06-04

## 1. Governance & Authority Rules

| Level | Definition | Authority |
| :--- | :--- | :--- |
| **CANONICAL** | The single source of truth for a concept or system. | Highest |
| **REFERENCE** | Derived from canonical sources (reports, summaries, maps). | Medium |
| **AUDIT** | Verification reports, coverage matrices, and implementation audits. | Medium |
| **RESEARCH** | Conceptual, speculative, or exploratory work. | Low |
| **HISTORICAL** | Read-only records of past states or decisions. | Low |
| **ARCHIVE** | Obsolete or superseded documents. | Zero |

---

## 2. Core Architecture Authority

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `MASTER_SPECIFICATION.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `FINAL_ARCHITECTURE.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `ARCHITECTURE_REPLACEMENT_PLAN.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `ARCHITECTURE_DEBT.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `FRACTAL_CYCLES_AND_LAWS.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `CLEANUP_REPORT.md` | AUDIT | `docs/audit/` | VERIFIED |
| `DEAD_CODE_REGISTER.md` | AUDIT | `docs/audit/` | VERIFIED |
| `DEPENDENCY_REPORT.md` | AUDIT | `docs/audit/` | VERIFIED |

## 3. Governance Authority

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `MASTER_INVARIANTS.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `PhoenixOS-Constitution-v1.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `BOUNDARY_RULES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `DEPENDENCY_RULES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `LAYERING_RULES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `EXTRACTION_RULES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `AI_AUTONOMY_RULES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `REPOSITORY_CONSTITUTION.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `GLOSSARY.md` | CANONICAL | `docs/canonical/` | VERIFIED |

## 4. Failure & Status Authority

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `FAILURE_CATALOG.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `CATASTROPHIC_FAILURES.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `RECOVERY_PATHS.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `ROLLBACK_MODEL.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `SYSTEM_STATUS.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `CHAOS_TEST_PLAN.md` | CANONICAL | `docs/canonical/` | VERIFIED |

## 5. Verification Authority

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `TRACEABILITY_MATRIX.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `CLAIMS_AUDIT.md` | AUDIT | `docs/audit/` | VERIFIED |
| `SPECIFICATION_COVERAGE_MATRIX.md` | AUDIT | `docs/audit/` | VERIFIED |
| `FAILURE_TO_TEST_MATRIX.md` | AUDIT | `docs/audit/` | VERIFIED |
| `INVARIANT_TO_CODE_MATRIX.md` | REFERENCE | `docs/reference/` | VERIFIED |
| `INVARIANT_VERIFICATION_MATRIX.md` | REFERENCE | `docs/reference/` | VERIFIED |

## 6. Subproject Documentation (Distributed)

| Subproject | Location | Authority |
| :--- | :--- | :--- |
| `Foundation/*` | `foundation/*/docs/` | REFERENCE |
| `Assurance/*` | `assurance/*/docs/` | REFERENCE |
| `Governance/*` | `governance/*/docs/` | REFERENCE |
| `Cognition` | `cognition/docs/` | REFERENCE |
| `Cognition (Future)` | `cognition/docs/future/` | RESEARCH |
| `Platform/*` | `platform/*/docs/` | REFERENCE |

## 7. Roadmaps & Planning

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `MASTER_PHOENIX_ROADMAP.md` | CANONICAL | `docs/canonical/` | VERIFIED |
| `EXTRACTION_ORDER.md` | ARCHIVE | `docs/archive/` | VERIFIED |

## 8. Game Documentation Authority

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `VISION.md` | CANONICAL | `game/docs/` | VERIFIED |
| `GAMEPLAY_LOOP.md` | CANONICAL | `game/docs/` | VERIFIED |
| `WORLD_STATE.md` | CANONICAL | `game/docs/specifications/` | VERIFIED |
| `Research/*` | RESEARCH | `game/docs/research/` | RESTORED |

## 9. Research & Experimental

| Document | Authority Level | Location | Status |
| :--- | :--- | :--- | :--- |
| `Labs` | RESEARCH | `labs/` | INVENTORY REQUIRED |
| `Crucible` | REFERENCE (TEMP) | `platform/crucible/` | CLASSIFICATION REQUIRED |

---
**Conclusion**: Documentation locality restored. Moving to Phase 4A.7 Labs Inventory.

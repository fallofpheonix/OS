# Governance Audit Report

## 1. Unenforced Invariants
- High-level invariants from `governance/constitution/INVARIANTS.md` (e.g., Memory Invariants, Identity Invariants, Branch Invariants) lack mapping and enforcement in `invariants.yaml`.
- Behavioral and Operational invariants defined in `invariants.yaml` are stubbed out in `invariant_engine.py` (e.g., method-level pattern checks, operational runtime deferrals).
- Rules from `control-plane/ARCHITECTURE_RULES.md` and `control-plane/DEPENDENCY_POLICY.md` (e.g., "Modules cannot depend on apps", "SDKs only expose contracts") are missing from executable `invariants.yaml`.

## 2. Governance Drift
- `INVARIANTS.md` serves as a philosophical rulebook, while `invariants.yaml` is highly specific to structural checks in `astraeus-core`. There is no synchronization between the written constitution and the executable rules.
- Operational invariants are declared but explicitly ignored (deferred) by the `InvariantEngine` during static evaluation.

## 3. Duplicate Rules
- Structural architectural boundaries are described redundantly across `control-plane/ARCHITECTURE_RULES.md`, `control-plane/DEPENDENCY_POLICY.md`, and `governance/architecture/DEPENDENCY_RULES.md`.
- Contract rules are split between `ARCHITECTURE_RULES.md` and `DEPENDENCY_POLICY.md`.

## 4. Missing CI Enforcement
- CI workflows are entirely absent (`.github/workflows/` or equivalent CI configurations do not exist in the repository).
- The `InvariantEngine` is not executed automatically on PRs or merges to prevent invalid architecture topologies, cyclic dependencies, or contract violations.

## 5. Architecture Gaps
- The `InvariantEngine` currently only evaluates structural dependencies locally via `ArchitectureGraph`. It lacks integration with the broader repository structure.
- "Every project has its own environment in environments/" is stated in `ARCHITECTURE_RULES.md` but is not programmatically validated.
- "No localhost hardcoding" is a stated rule but no linter or engine checks for it.

## 6. Dependency Violations
- Unenforced dependency rules mean violations are highly likely to exist unchecked:
  - "Modules cannot depend on apps."
  - "Infrastructure cannot depend on experiments."
  - "Research projects cannot directly touch production systems."
- A scan is required to confirm actual current violations, but the lack of CI gating makes the codebase fundamentally vulnerable.

## 7. Missing Contracts
- `control-plane/contracts/` exists with various schemas, but `DEPENDENCY_POLICY.md` says "All shared contracts must be registered in control-plane/contracts/." The `astraeus-core` currently holds its own contracts inside `workspace/active/astraeus-core/contracts/`, which fragments the contract governance and violates the registration rule.

## 8. Undocumented Ownership
- No `CODEOWNERS` or clear metadata file enforces ownership over `governance/`, `infrastructure/`, `control-plane/`, or `workspace/active/` modules. Ownership mapping is purely implicit.

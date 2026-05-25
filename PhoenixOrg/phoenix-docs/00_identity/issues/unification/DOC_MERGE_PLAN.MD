# DOC_MERGE_PLAN.md

This plan details the final multi-part execution for the Documentation Unification Pass.

## Phase 1: Core Identity
- Merge `identity/axioms.md` into `identity/system_identity.md`.
- Move `identity/axioms.md` to version history.

## Phase 2: Architecture Models
- Populate `architecture/state_model.md` using archived `FSM_ARCHITECTURE.md`.
- Populate `architecture/replay_spec.md` using archived `REPLAY_ARCHITECTURE.md`.
- Populate `architecture/truth_model.md` using archived `LEDGER_ARCHITECTURE.md`.
- Populate `architecture/decision_model.md` using archived `DECISION_PIPELINE.md`.

## Phase 3: Security & Boundaries
- Populate `security/threat_model.md` using archived `02-threat-model.md`.
- Ensure `security/boundaries.md` contains rules from `BOUNDARY_RULES.md`.

## Phase 4: Validation & Reports
- Merge `validation/validation_rules.md` with `VALIDATION_MATRIX.md`.
- Move all redundant reports to version history.

## Phase 5: External Governance
- Ensure `external_governance/external_policy.md` is fully populated.
- Sync `external_repo_matrix.md` with current `external/` folder contents.

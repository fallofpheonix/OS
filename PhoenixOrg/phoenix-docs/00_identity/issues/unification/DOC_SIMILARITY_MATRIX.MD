# DOC_SIMILARITY_MATRIX.md

This matrix identifies overlapping documentation that must be merged into the 25 canonical files.

| Canonical Doc | Redundant / Overlapping Sources |
| :--- | :--- |
| `system_identity.md` | `AXIOMS.md`, `PROJECT_VISION.md`, `PHASE_MATRIX.md` |
| `roadmap.md` | `MASTER_ROADMAP.md`, `MATURATION_ROADMAP.md`, `very_long_term_plan.md` |
| `phase_lock.md` | `PHASE_STATUS.md`, `MASTER_STATUS.md`, `CHECKPOINTS.md` |
| `state_model.md` | `FSM_ARCHITECTURE.md`, `WARDEN.md`, `CONTAINMENT.md` |
| `replay_spec.md` | `REPLAY_ARCHITECTURE.md`, `REPLAY.md`, `EVIDENCE_SPEC.md` |
| `truth_model.md` | `LEDGER_ARCHITECTURE.md`, `TRUTH_LEDGER.md`, `LEDGER_SPEC.md` |
| `decision_model.md` | `DECISION_PIPELINE.md`, `ARBITER.md`, `POLICY_ARCHITECTURE.md` |
| `boundary_rules.md` | `ILLEGAL_PATHS.md`, `SECURITY_BOUNDARIES.md`, `MODULE_BOUNDARIES.md` |
| `validation_rules.md` | `TEST_PLAN.md`, `VALIDATION_MATRIX.md`, `VALIDATION_COVERAGE.md` |
| `test_matrix.md` | `SERVER_TEST_MATRIX.md`, `NODE_TEST_MATRIX.md`, `TEST_INDEX.md` |
| `build_spec.md` | `BUILD_PIPELINE.md`, `BUILD_SPEC.md` (old), `lfs-build.md` |
| `threat_model.md` | `ATTACK_MATRIX.md`, `ATTACK_SURFACE.md`, `ZERO_TRUST_MODEL.md` |
| `risk_matrix.md` | `RISK_REGISTER.md`, `MVP_risk_register.md`, `RISKS.md` |
| `external_policy.md` | `EXTERNAL_REPOS.md`, `REPO_HYGIENE.md`, `repo_classification.md` |
| `research_policy.md` | `RESEARCH_INDEX.md`, `MASTER_RESEARCH.md`, `THEORY_NOTES.md` |
| `formal_plan.md` | `PROOFS.md`, `QUANTUM_AUDIT.md`, `formal-verification` (various) |

## Strategy
1. Extract unique value from "Sources".
2. Append or integrate into "Canonical Doc".
3. Move "Sources" to `archive/documentation/version_history/`.

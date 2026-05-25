# Validation Pipeline

Defines stages and gates that outputs must pass before becoming authoritative.

Stages
------
- Syntax checks
- Semantic validation (type, API compatibility)
- Invariant enforcement (run `INVARIANTS.md` validators)
- Runtime validation (integration and smoke tests)
- Confidence scoring and critic review

Gating
------
- Validation failures must prevent commit and trigger repair planning.
- Certain violations escalate immediately to human approval workflows.

Next steps
----------
- Map validators to CI jobs and runtime validators.

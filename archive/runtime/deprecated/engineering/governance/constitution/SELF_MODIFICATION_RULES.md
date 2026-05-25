# Self-Modification Rules

Controls and limits for autonomous self-modification and adaptation.

Sections
--------
- Allowed Self-Changes
- Forbidden Self-Changes
- Simulation Requirements
- Evolution Sandboxing
- Rollback Requirements
- Human Approval Boundaries

Key Rules
---------
- No autonomous modification may directly alter constitutional documents or invariants.
- All self-modification proposals must be simulated offline (dream/replay), include counterfactual evaluation, and provide rollback steps.
- Self-modification to production branches requires human multi-signer approval and staged rollout with canary gating.

Next steps
----------
- Implement simulation + counterfactual evaluator as part of the repair planner.

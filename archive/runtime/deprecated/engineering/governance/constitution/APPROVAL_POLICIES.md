# Approval Policies

Human authority model and approval workflows for high-risk operations.

Sections
--------
- Dangerous Operations
- Destructive Operations
- Autonomous Boundaries
- Human Override
- Emergency Recovery

Approval model
--------------
- Low-risk changes: single owner approval + CI green.
- Medium-risk changes: two approvers from different teams + CI + replay validation.
- High-risk changes (protected modules, self-modification, constitutional amendments): multi-signer (3+) approval, ADR, staged rollout, automated rollback tests.

Emergency procedures
--------------------
- Emergency remediation may be executed by a designated on-call with post-hoc ADR and incident report.

Next steps
----------
- Implement approval metadata in `MASTER_REPO_INDEX.yaml` and CI gating.

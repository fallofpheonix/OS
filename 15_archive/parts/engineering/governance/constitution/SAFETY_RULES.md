# Safety Rules

Operational safety policies and required behaviors for mutation, execution, and data handling.

Sections
--------
- Mutation Safety
- Command Execution Safety
- Cloud Safety
- Data Safety
- Autonomous Action Safety
- Recovery Safety

Mutation Safety
---------------
- All destructive mutations require a checkpoint and an explicit approval step documented in `APPROVAL_POLICIES.md`.
- Mutations to protected modules must include a rollback plan and automated tests that validate invariants.

Command Execution Safety
------------------------
- Arbitrary shell commands executed by the system must run in an isolated sandbox and be logged to the event journal.
- Dangerous commands (file-system deletes, package manager operations) require higher approval levels.

Cloud Safety
------------
- Cloud model calls may be used for inference only; they must not store persistent state.
- Secrets and credentials must never be sent to third-party models or stored in plain text.

Data Safety
-----------
- Sensitive data must be redacted in logs and replay traces by default.
- Backups must be encrypted and integrity-checked.

Autonomous Action Safety
------------------------
- The system may abstain when uncertainty exceeds configured thresholds.
- High-risk interventions require human approval and/or multi-signer gates.

Recovery Safety
---------------
- Every mutation must define a recovery strategy: checkpoint location, rollback commands, and verification tests.

Next steps
----------
- Add enforcement hooks in CI to scan mutation PRs for recovery artifacts and tests.

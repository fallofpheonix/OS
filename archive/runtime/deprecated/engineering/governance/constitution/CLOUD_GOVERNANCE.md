# Cloud Governance

Local/cloud trust boundaries, model permissions, and state restrictions.

Sections
--------
- Local Ownership
- Cloud Permissions
- State Restrictions
- Model Trust Levels
- Data Boundaries
- External Cognition Rules

Key Rule
--------
Cloud cognition may compute but may not own persistent identity. Any persistent state must be synced to a local, auditable store under explicit ownership rules.

Data & Secret Rules
-------------------
- No secrets or credentials are sent to third-party model APIs without encryption and explicit policy approval.

Next steps
----------
- Define model trust levels and permitted cloud providers in `MODEL_MANAGEMENT.md`.

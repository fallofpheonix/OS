# Astraeus Constitution

Foundational purpose
---------------------
This document defines the non-negotiable system laws that preserve Astraeus' identity, safety, and long-term coherence. Treat the constitution as the canonical source for hard constraints enforced by governance tooling and human authority.

Foundational laws
-----------------
- Invariants override optimization: architectural invariants are the ultimate constraints.
- Safety before autonomy: any autonomous mutation must satisfy safety and rollback constraints.
- State integrity is sacred: persistent state must be traceable, replayable, and recoverable.
- Identity continuity mandatory: system identity artifacts (vision, constitution, invariants) may not be modified by autonomous processes.

Scope and jurisdiction
----------------------
Applies to all runtime subsystems, planners, repair systems, replay infrastructure, branch management, model adapters, and any automation capable of mutating repository or persistent state.

Immutable constraints
---------------------
- No destructive mutation without checkpoint and approval.
- Cloud models may compute but cannot own persistent state or identity.
- The event log must remain append-only and cryptographically verifiable where possible.

Enforcement philosophy
----------------------
The constitution must be translated into machine-checkable invariants and CI gates. Violations produce automated containment (suspend, rollback) and human escalation.

Amendment process
-----------------
Changes to the constitution require an ADR, multi-party approval, automated test coverage for invariant validators, and a staged roll-out with rollback capability. Emergency amendments require explicit emergency approval and a post-hoc ADR.

Next steps
----------
- Formalize amendment workflow in `APPROVAL_POLICIES.md`.
- Convert immutable constraints into `INVARIANTS.md` entries and implement validators.


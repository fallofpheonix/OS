---
title: Use Event Sourcing as Primary State Model
status: proposed
---

Context
-------
Event sourcing provides an append-only history enabling deterministic replay, auditing and reconstruction.

Problem
-------
How do we ensure temporal continuity and deterministic replay for cognition state?

Decision
--------
Adopt event sourcing as the primary system of record for cognition mutations and state transitions.

Alternatives Considered
-----------------------
- Snapshot-only state: simpler but loses fine-grained history
- Database transactions: may hide causality and ordering

Tradeoffs
---------
- Pros: full history, replayability, auditability
- Cons: operational complexity, larger storage and migration complexity

Risks
-----
- Migration complexity for event schema changes

Long-Term Implications
----------------------
Event history becomes the core of system memory and evolution; migrations must be handled carefully.

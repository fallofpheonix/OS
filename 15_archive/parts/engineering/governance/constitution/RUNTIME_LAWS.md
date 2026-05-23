# Runtime Laws

Execution-level guarantees, ordering, and repair boundaries.

Sections
--------
- Execution Ordering
- Replay Guarantees
- Event Sourcing Rules
- Failure Propagation
- Repair Boundaries
- Scheduling Rules

Key Laws
--------
- No task output is authoritative until validated by the `validator` subsystem.
- Event journal is the single source of truth for mutation-producing operations.
- Repair loops are limited to a configurable retry count (default = 3) and must include rollback plans.
- Failed replay reduces branch trust score; severely degraded branches are quarantined.

Scheduling rules
----------------
- Heavyweight inference tasks require scheduler approval and an exclusive lock on node resources.
- Short deterministic tasks may execute concurrently under constrained resource budgets.

Next steps
----------
- Implement runtime checks for task validation states and branch trust scoring.

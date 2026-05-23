# Runtime Architecture

Defines the execution kernel, scheduler, event bus, replay engine, and rollback semantics.

Contents
--------
- Execution Pipeline
- Task Lifecycle
- Scheduler and Resource Management
- Event Bus and Storage
- Replay Engine
- Rollback Engine
- Failure Handling and Containment

Execution guarantees
--------------------
- Deterministic replay: every mutation-producing operation must be restorable from the event journal.
- Validation before commit: no output becomes authoritative until validated by the validator subsystem.

Scheduler rules
---------------
- Heavyweight inference tasks require exclusive node locks and checkpointed state.
- Lightweight tasks may run concurrently under configured quotas.

Next steps
----------
- Document scheduler implementation, lock manager, and example sequence diagrams.

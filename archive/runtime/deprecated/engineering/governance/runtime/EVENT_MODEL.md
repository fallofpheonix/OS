# Event Model

Defines event types, ordering, storage, and replay semantics used by the runtime.

Event philosophy
----------------
Every meaningful state transition or mutation emits an event into the append-only event journal. Events are the single source of truth for replay and recovery.

Event categories
----------------
- Task lifecycle events (TASK_CREATED, TASK_STARTED, TASK_COMPLETED, TASK_FAILED)
- Validation events (VALIDATION_PASSED, VALIDATION_FAILED)
- Repair events (REPAIR_TRIGGERED, REPAIR_COMPLETED)
- Mutation events (MUTATION_STAGED, MUTATION_COMMITTED, MUTATION_ROLLED_BACK)
- Branch events (BRANCH_CREATED, BRANCH_MERGED)
- Session events (SESSION_STARTED, SESSION_RESUMED, SESSION_TERMINATED)

Ordering and integrity
----------------------
- Events must include causal metadata (event id, parent event ids, timestamp, origin).
- Event journal must be append-only; integrity checks (hash chaining) are recommended.

Next steps
----------
- Define canonical JSON schemas for each event type and implement validators.

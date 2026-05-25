# Task Lifecycle

Defines the state machine for tasks executed by the cognition runtime.

States
------
- PENDING — created but not ready
- READY — all dependencies satisfied
- RUNNING — actively executing
- VALIDATING — undergoing validation and critic checks
- COMPLETED — success and persisted
- FAILED — terminal failure awaiting repair
- REPAIRING — repair in progress
- ARCHIVED — historic, read-only

Transitions
-----------
- Allowed transitions and forbidden jumps must be enforced by the scheduler and recorded as events.
- Retry semantics: configurable per-task with default max retries = 3.

Failure handling
----------------
- Classification into transient vs persistent failures.
- Transient failures may be retried automatically; persistent failures trigger repair planning.

Next steps
----------
- Add formal state machine diagram and JSON schema for task events.

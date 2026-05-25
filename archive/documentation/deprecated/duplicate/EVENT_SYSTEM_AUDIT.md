# EVENT SYSTEM AUDIT REPORT

## 1. Event Schema & Coverage
*   **Active Coverage:** `EventAction` covers high-level lifecycle stages (RUN_STARTED, TASK_STARTED, etc.) and recently added mutation/rollback events.
*   **Gaps in Granularity:**
    *   **Planner Decisions:** No events for *how* the planner arrived at a DAG (e.g., retrieval context used, model reasoning before JSON generation).
    *   **Repair Logic:** `REPAIR_PLANNED` exists, but the internal state of the `RepairPlanner` (ranking of failures, context selection) is opaque.
    *   **Critic Reasoning:** `VALIDATION_FAILED` contains the critic's output, but the semantic relationship between the failure and the architectural context isn't tracked as a first-class event.
    *   **Queue State:** `ExecutionQueue` manages the `topological_batches`, but the *transition* of a task from `PENDING` to `READY` to `RUNNING` is partially inferred.

## 2. Hidden State Transitions
*   **RunState Mutation:** `RunState` and `TaskRecord` are modified directly in-memory and then dumped to `run.json`. These mutations are side-effects of orchestration, not derived from events.
*   **Status Management:** `state.set_status(task_id, status)` updates the state, but there isn't always a 1:1 event for every status flip (e.g., `TaskStatus.BLOCKED` or `TaskStatus.PENDING` during subtree resets).
*   **Snapshot Manifests:** `SnapshotEngine` creates JSON manifests that exist outside the event stream.

## 3. Nondeterministic State Mutations
*   **In-Process Memory:** `ControlPlane` holds `topology_engine` and `graph` in-memory. If the process restarts, this state is lost. Replay currently relies on `run.json` (the *result* of state) rather than the *sequence* of events.
*   **System Time Dependency:** `run_id` and timestamps are derived from `datetime.now(UTC)`, which is fine for logging but makes strict deterministic clock replay difficult without mocking `time`.

## 4. Replay Gaps & Inconsistencies
*   **Missing Inputs:** `EventBus` does not record the full input prompt or retrieval context for every model call—only high-level metadata. Replaying the model call requires finding the right artifact file.
*   **State vs. Event:** Replay currently uses `run.json` (state) as the source of truth for task status. For true event sourcing, `run.json` should be a *projection* of the event log.
*   **Rollback Replay:** `MUTATION_ROLLED_BACK` events exist but don't contain enough metadata (like the specific backup hash restored) to reconstruct the exact filesystem state at that timestamp without the journal.

## Conclusion
The current system is "Event Logging," not "Event Sourcing." To support temporal replay, we must transition to a model where **events are the authoritative source of truth**, and the `RunState` is a projection that can be deterministically rebuilt by folding events.

# EVENT SOURCING ARCHITECTURE

## Core Principle: State as a Projection
In Astraeus, the `RunState` (and the final `run.json`) must not be the authoritative record of truth. Instead, the `EventBus` (and the `journal.jsonl`) record every atomic transition. The `RunState` is a **projection** that is updated in response to events.

### 1. Unified Event Log
All system transitions must be recorded in the append-only `events.jsonl`. This includes:
*   **Orchestration:** Plan generation, task status flips, model calls, validation results.
*   **Mutation:** Filesystem changes (linked to the Journal).
*   **Repair:** Failure detection and repair planning decisions.

### 2. State Reconstruction (Replay)
To reconstruct the system state at any point in time `T`:
1.  Initialize a clean `RunState`.
2.  Read all events from the log where `timestamp <= T`.
3.  Apply each event to the `RunState` using a deterministic projector function.

### 3. Temporal Consistency
*   **Event ID Chaining:** (Optional) Events can include the hash of the previous event to ensure the timeline hasn't been tampered with.
*   **Determinism:** System decisions must rely exclusively on state derived from the event log or immutable inputs.

## Expanded Event Schema

| Event Action | Description | Metadata |
|---|---|---|
| `PLAN_GENERATED` | Planner output received | Full DAG and reasoning |
| `TASK_STATUS_UPDATED` | Explicit status flip (PENDING, RUNNING, etc.) | `old_status`, `new_status` |
| `MODEL_INPUT_CAPTURED` | The exact prompt and context sent to LLM | `prompt`, `context_ids` |
| `REPAIR_DECISION_MADE` | Why a specific repair path was chosen | `failure_id`, `strategy` |
| `SNAPSHOT_COMMITTED` | Snapshot finalized | `snapshot_id`, `hash` |

## Implementation Strategy
1.  **Refactor `RunState`:** Move mutation logic into a `project_event(event)` method.
2.  **Instrument `ControlPlane`:** Ensure every meaningful code path emits an event *before* updating local state.
3.  **Bridge Journal:** Ensure `MUTATION_APPLIED` events contain enough metadata to link back to the `JournalEntry`.

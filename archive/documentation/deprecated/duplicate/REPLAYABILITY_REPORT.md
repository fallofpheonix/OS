# REPLAYABILITY & OBSERVABILITY REPORT

## Overview
Phase 4 focused on ensuring every runtime action is traceable, attributable, and deterministic enough for eventual full-state replayability. The system now records granular event streams and filesystem journals to maintain a complete history of execution.

## Observability Enhancements

1.  **Extended Event Schema:**
    *   Added new `EventAction` types to `events/schema.py`:
        *   `MUTATION_APPLIED`: Emitted when the `TransactionRunner` successfully applies a `DiffPlan` to the project root.
        *   `MUTATION_ROLLED_BACK`: Emitted when the `RollbackEngine` reverts specific file paths.
    *   This closes the critical observability gap identified in Phase 1 by formally linking the orchestration event bus to the filesystem journal, ensuring all changes are tracked in both streams.

2.  **Telemetry & Tracing:**
    *   The `EventBus` persists all events, including latency and results, to an append-only JSONL format (`events.jsonl`).
    *   This provides a structured execution trace for every task DAG execution.

3.  **Attribution:**
    *   Every entry in the `FilesystemJournal` includes a `run_id`, allowing a strict mapping from a specific file mutation back to the user request and task that initiated it.

## Replayability Status (Current Gaps)
The `ReplayEngine` currently validates the consistency of the artifact stream (ensuring all tasks have corresponding events and output artifacts). 

**Remaining Work for Full Replay:**
To achieve deterministic state-machine replay, the `ReplayEngine` must be expanded to:
1.  Initialize a mock or isolated `ControlPlane` environment.
2.  Feed the exact `DiffPlan` and intermediate outputs from the journal/event bus to verify that the orchestration logic deterministically produces the exact same downstream decisions (routing, DAG construction, validation).
3.  Compare the resulting final state against the persisted final state.

*Conclusion:* The observability primitives are fully active, laying the groundwork for verifiable deterministic replay.

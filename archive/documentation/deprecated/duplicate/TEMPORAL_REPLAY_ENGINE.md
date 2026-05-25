# TEMPORAL REPLAY ENGINE

## Overview
The `ReplayEngine` is the core component responsible for deterministic reconstruction of system state from the append-only event log. It implements the **Fold** pattern: `State = fold(EventLog, Projector)`.

## Features

### 1. State Reconstruction
The engine can rebuild a full `RunState` for any `run_id`.
*   **Targeted Replay:** `reconstruct_state(run_id)`
*   **Temporal Replay:** `reconstruct_state(run_id, up_to_timestamp="...")` - Rebuilds the state as it existed at a specific point in time.

### 2. Consistency Verification
The engine can verify that the persisted `run.json` (the cache of final state) is consistent with the event history. This detects "hidden mutations" that were not journaled as events.

### 3. Timeline Reconstruction
The engine provides a base for building a chronological timeline of all system actions, including:
*   Task lifecycle transitions.
*   Model calls and their results.
*   Filesystem mutations (linked via metadata).
*   Repair decisions.

## Replay Flow
1.  **Locate Log:** Identify the `events.jsonl` associated with the `run_id`.
2.  **Initialize:** Create a fresh `RunState`.
3.  **Project:** Iterate through events, calling `state.apply_event(event)` for each.
4.  **Finalize:** Return the reconstructed state.

## Future Expansion
*   **Branch Replay:** Support replaying alternative timelines from a specific event index.
*   **Mutation Replay:** Replay actual filesystem changes using the backups stored in the journal.
*   **Query Layer:** Add a temporal query interface (e.g., "Show me the state of `task_1` after the first repair attempt").

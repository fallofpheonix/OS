# TEMPORAL QUERY DESIGN

## Overview
The Temporal Query Layer provides an interface to reason about the system's history. It abstracts the event replay process into a set of high-level queries for historical state and lineage.

## Query Types

### 1. State-at-Time (Point-in-Time)
Retrieve the full system state as it existed at timestamp `T`.
*   **API:** `get_state(run_id, timestamp)`
*   **Use Case:** Debugging a failure by inspecting the exact state of all tasks before a specific error event.

### 2. State Diff
Compare the state at time `T1` with the state at time `T2`.
*   **API:** `diff_states(run_id, t1, t2)`
*   **Use Case:** Understanding the "blast radius" of a repair operation or a subtree reset.

### 3. Execution Lineage (Ancestry)
Trace the origin of a specific state or artifact.
*   **API:** `trace_ancestry(run_id, task_id)`
*   **Use Case:** Identifying which dependency or model decision led to a specific invalid code generation.

### 4. Mutation Timeline
Reconstruct the chronological history of filesystem changes for a specific file or directory.
*   **API:** `file_history(run_id, path)`
*   **Use Case:** Seeing all versions of `main.py` produced during a multi-task run.

## Proposed API (runtime/temporal_query.py)

```python
class TemporalQueryLayer:
    def __init__(self, replay_engine: ReplayEngine):
        self.replay = replay_engine

    def get_task_history(self, run_id: str, task_id: str) -> list[TaskStatus]:
        """Show every status transition for a specific task."""
        state = self.replay.reconstruct_state(run_id)
        # Replay would need to capture a list of transitions in TaskRecord
        pass

    def get_mutation_lineage(self, run_id: str, file_path: str) -> list[JournalEntry]:
        """Show every journaled change to a specific file."""
        pass
```

## Integration with CLI
The query layer will power new CLI commands:
*   `astraeus replay --run-id <id> --time <T>`
*   `astraeus history --file <path>`
*   `astraeus trace --task-id <id>`

# EXECUTION_OWNERSHIP_MODEL.md

## Overview
To safely orchestrate distributed, concurrent engineering operations, Astraeus must enforce strict boundaries around who owns specific operations, resources, and state.

## 1. Run Ownership (Orchestrator)
- **Authority**: A single logical run (`run_id`) is owned by exactly **one Control-Plane instance (the Coordinator)** at any given time.
- **Worker Boundaries**: Sub-tasks (`task_id`) within a run can be dispatched to distributed workers, but workers DO NOT own the run state. They are stateless executors.
- **State Management**: The Coordinator is the sole source of truth for the `RunState`. Workers receive immutable input payloads and return results; they do not directly mutate `RunState`.

## 2. Mutation Ownership (Transaction Sandbox)
- **Authority**: The **MutationSandbox** is the exclusive owner of project root mutations.
- **Transaction Ownership**: When a task proposes a mutation, the Sandbox claims ownership of that operation via a distributed lease/lock. No other task can mutate the target paths until the transaction is committed or aborted.
- **Distributed Isolation**: A worker process only builds a `DiffPlan` and validates it within a local staging directory. The final commit to `project_root` is marshaled exclusively by the authorized MutationSandbox.

## 3. Event Ownership (Event Bus & Sequence Coordinator)
- **Authority**: The **EventBus** transitions from a simple file appender to an append-only Sequence Coordinator.
- **Ordering**: In a distributed setup, events are sent to a centralized event broker (e.g., Redis, a dedicated DB table, or an isolated sequencer process) which assigns global, monotonically increasing sequence numbers.
- **Replay**: Replay mechanisms trust the global sequence number, not local timestamps, to reconstruct state.

## 4. Rollback Authority (Safety Substrate)
- **Authority**: The **RollbackEngine** requires exclusive global authority over the repository before executing.
- **Conflict Prevention**: A rollback operation requires a Global Write Lock on the repository. While a rollback is executing, all other mutations (from any worker or coordinator) are strictly blocked.
- **Drift Checks**: Before rollback, the engine still validates hashes, serving as the final defense against external (non-Astraeus) mutations.

## 5. Implementation Strategy
1. **Centralized RunState Management**: Refactor `orchestrator/queue.py` so workers yield results back to the queue, and only the main orchestrator loop mutates `self.state`.
2. **Global Event Sequencing**: Introduce an abstraction for event sequencing that can be backed by SQLite/Redis instead of just a local file with thread locks.
3. **Repository Leases**: Implement a file-based or DB-based lease mechanism for repository locks (Read/Write) to serialize commits and rollbacks.

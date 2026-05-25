# Execution Flow

Defines the canonical end-to-end cognition execution loop and ordering guarantees.

Sections
--------
- Prompt Intake
- Task Decomposition
- DAG Construction
- Model Routing
- Execution Scheduling
- Validation
- Repair
- Persistence & Event Journal
- Replay
- Completion

Canonical sequence
------------------
Prompt -> Planner -> Task DAG -> Scheduler -> Executor -> Validator -> Repair or Persist -> Memory Consolidation

Required guarantees
-------------------
- Synchronous points where validation must complete before commit.
- Event emission timing for every mutation-producing operation.
- Checkpointing rules prior to heavy mutations.

Next steps
----------
- Add sequence diagrams and example traces for common workflows.
# Execution Flow

Describe the DAG lifecycle, validation pipeline, event sourcing, replay, and rollback primitives.

Sections
--------
- Prompt → Planner → Router → Executor → Validator → Repair
- Event journal and replay semantics
- Branch execution and merge rules
- Retry and rollback policies

TODO
----
- Add sequence diagrams and example event traces.

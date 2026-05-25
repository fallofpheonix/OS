# Principles

Core principles that guide design, implementation, and governance across the Astraeus ecosystem. These are intended as durable truths that survive implementation changes.

Local-first cognition
---------------------
The system prioritizes local, reproducible computation and state ownership. Local runtime owns persistent identity and is the canonical source for session state unless explicitly delegated.

State over prompts
------------------
Persistent structured state (beliefs, subgoal graphs, memory traces) is the primary input to planning and repair; prompt engineering is a tactical surface, not the foundation.

Invariants over heuristics
-------------------------
Encode architectural constraints and safety rules as invariants that are enforced programmatically; heuristics may guide decisions but must not violate invariants.

Repository grounding mandatory
----------------------------
All planning, repair, and generation must default to repository-grounded context (symbol graph, imports, ownership). Unseen APIs require explicit hypothesis and testing.

Memory persistence required
-------------------------
Execution traces, failures, and consolidated abstractions must be persisted with replayable event histories to support debugging, consolidation, and causal analysis.

Safety before autonomy
----------------------
Autonomous mutations, especially to critical modules, require approval workflows, checkpoints, and rollback plans. Abstention is a valid action.

Additional principles
---------------------
- Deterministic replayability is mandatory for all mutation-producing operations.
- Minimal surprise: the system must prefer conservative changes with explicit rationale.
- Human-in-the-loop: the system defers to owners and human reviewers on high-risk interventions.

Next steps
----------
- Convert these principles into automated checks in CI (invariant validators, replay tests).
- Add examples showing how each principle maps to runtime behavior.

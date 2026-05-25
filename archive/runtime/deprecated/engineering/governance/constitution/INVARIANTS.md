# Invariants

This file enumerates machine-enforceable architectural truths. Each invariant is written to be testable, auditable, and automatable. Invariants are the primary enforcement mechanism of the Constitution.

Invariant format
----------------
Each invariant entry SHOULD include the following fields:

- ID: unique identifier (e.g., INV-RUNTIME-001)
- Title: short descriptive title
- Rule: machine-verifiable rule text
- Reasoning: why it exists and what it prevents
- Enforcement Layer: runtime/CI/validator name
- Severity: critical/high/medium/low
- Violation Consequences: automatic actions and escalation
- Recovery Strategy: how to recover or rollback

Example invariant
-----------------
ID: INV-RUNTIME-001

Title: Single heavyweight inference

Rule: Only one active heavyweight inference process may execute locally at a time.

Reasoning: Prevents memory exhaustion and resource contention that can corrupt session state.

Enforcement Layer: runtime scheduler lock (scheduler/lock-manager)

Severity: Critical

Violation Consequences: suspend new inference requests, trigger checkpoint and rollback of affected tasks, open incident.

Recovery Strategy: restore from last checkpoint; require manual review before resuming heavy inference on the node.

Invariant categories
-------------------

A. Runtime Invariants
- Event log append-only
- Task transitions must emit events

B. Memory Invariants
- Semantic memory cannot overwrite episodic memory
- Replay history must reconstruct exact prior state

C. Repository Invariants
- No circular imports between runtime and memory layers
- Protected modules require approval for mutation

D. Branch Invariants
- Branches must preserve full causal history
- Merges require invariant validation

E. Identity Invariants
- Persistent memory ownership remains local-first
- System identity docs cannot be modified by autonomous agents

F. Evolution Invariants
- Adaptive policies may not weaken constitutional constraints
- Adaptive routing must preserve deterministic replay metadata

Next steps
----------
- Convert each category item above into `ID`-tagged invariants and add enforcement stubs in the validator config.


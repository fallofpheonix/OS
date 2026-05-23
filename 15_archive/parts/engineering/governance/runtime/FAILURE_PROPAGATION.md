# Failure Propagation

Defines rules for how failures propagate through tasks, branches, and the runtime and how to contain cascades.

Sections
--------
- Local failure containment
- Dependency failure escalation
- Branch-level isolation of failures
- Repair escalation and human intervention thresholds

Containment principles
----------------------
- Fail fast and isolate: on invariant violation, suspend dependent tasks and quarantine branches as needed.

Next steps
----------
- Provide examples of failure scenarios and containment flows.

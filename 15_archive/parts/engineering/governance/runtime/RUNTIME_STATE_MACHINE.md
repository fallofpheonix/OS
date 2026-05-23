# Runtime State Machine

Defines the top-level runtime operating states and allowed transitions.

States
------
- IDLE
- PLANNING
- EXECUTING
- VALIDATING
- REPAIRING
- ROLLING_BACK
- SYNCHRONIZING
- RECOVERING

Rules
-----
- Transitions must be recorded as events to ensure observability.
- State transitions may be inhibited by invariant violations or resource constraints.

Next steps
----------
- Provide a state transition diagram and mapping to scheduler behavior.

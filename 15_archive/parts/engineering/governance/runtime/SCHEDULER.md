# Scheduler

Defines task scheduling, resource allocation, and concurrency rules.

Responsibilities
----------------
- Enforce concurrency and resource limits
- Grant exclusive locks for heavyweight inference
- Prioritize tasks by policy, uncertainty, and owner
- Integrate with the router for model-affinity scheduling

Rules
-----
- Only one heavyweight local inference may execute simultaneously per node by default.
- Adaptive scheduling can change priorities but must preserve replay metadata and checkpoints.

Next steps
----------
- Document current scheduler implementation and configuration knobs.

# Design Values

Engineering tradeoffs and guiding preferences.

Reliability vs Speed
---------------------
Prefer reliable, auditable changes even if slower to implement.

Safety vs Autonomy
------------------
Favor safe abstention and human escalation over risky automated mutation.

Local vs Cloud
--------------
Local runtime retains identity; cloud resources are used for scalable compute but never for owning persistent state.

Determinism vs Exploration
--------------------------
Support experimental exploration, but ensure determinism for mutation-producing operations.

Stability vs Evolution
----------------------
Enable conservative evolution with ADRs and migration paths rather than rapid ungoverned changes.

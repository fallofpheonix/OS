---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# DEPENDENCY POLICY

## Rules
- Modules cannot depend on apps.
- Infrastructure cannot depend on experiments.
- SDKs only expose contracts (interfaces).
- Research projects cannot directly touch production systems.
- Apps only compose modules.

## Additional Interoperability Rules
- Services communicate only through contracts.
- Modules expose interfaces, never internal implementation.
- Shared state across apps is forbidden.
- Every reusable component must have standalone tests.
- Fork-derived code must be marked explicitly.
- No localhost hardcoding: use internal DNS or env-based discovery.

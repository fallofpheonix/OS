# PhoenixOS: System Identity

**Definition:** Deterministic Cyber Defense Runtime with physics-inspired metrics.
**Core Philosophy:** Security as a managed runtime process, focusing on deterministic verification and containment.

## Axioms
1. **Determinism is sacred.** No reliance on non-deterministic primitives.
2. **Replay is authoritative.** Replay governs semantics; logs and metrics are secondary.
3. **AI is advisory.** AI informs, but never directly controls the kernel or actuation FSM.
4. **Control must remain bounded.** Actuation is rate-limited, isolated, and reversible.
5. **Telemetry correctness > AI sophistication.** Precise, replayable telemetry is the foundation.
6. **Never scale instability.** Single-node stability must be achieved before distributed scaling.

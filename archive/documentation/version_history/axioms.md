# PhoenixOS: Six Immutable Axioms

These axioms govern all development and runtime behavior within PhoenixOS.

1. **Determinism is sacred.** No reliance on non-deterministic primitives (unordered maps, race-dependent ordering).
2. **Replay is authoritative.** Replay governs semantics; logs, metrics, and AI-outputs are secondary.
3. **AI is advisory.** AI informs, but never directly controls kernel or actuation FSM.
4. **Control must remain bounded.** Actuation is rate-limited, isolated, and reversible.
5. **Telemetry correctness > AI sophistication.** Precise, replayable telemetry is the foundation.
6. **Never scale instability.** Single-node stability must be achieved before distributed scaling.

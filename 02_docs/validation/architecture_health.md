# Architecture Health

Summary
- Layering: Mostly consistent with telemetry → normalization → graph → evidence → control → physics → game → AI → swarm.
- Observations: Some documents place control and game ordering ambiguously; kernel work is scheduled before full replay/evidence validation in places — suggest reordering to always validate in userspace first.

Checks
- Circular dependencies: None obvious in docs, but code coupling between `phoenix_os/warden` and kernel hooks should be audited.
- Module isolation: Good separation between telemetry, trace, and evidence; control currently mixes strategy and actuation in some modules.

Recommendations
- Enforce clear contracts: `telemetry schema` (RFC-001), `trace ingestion` (RFC-006), `warden API` (RFC-005).
- Add interface mocks for kernel behavior to allow offline testing.
- Gate kernel patches behind evidence replay and determinism tests.

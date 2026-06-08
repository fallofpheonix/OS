---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# ADR-002: Warden FSM Design

## Status
Accepted

## Date
2026-05-31

## Context
The system needs a deterministic security state machine to govern containment and enforcement decisions. The FSM must be simple, verifiable, and resistant to oscillation.

## Decision
Implement a 5-state FSM with strict ladder transitions: SAFE → WATCH → SUSPICIOUS → CRITICAL → COMPROMISED. Each state can only transition to adjacent states (with de-escalation paths).

## Consequences

### Easier
- Simple, verifiable state transitions
- No state skipping prevents catastrophic jumps
- TLA+ model checking possible

### More Difficult
- Slow escalation through intermediate states
- No direct jump from SAFE to COMPROMISED
- Recovery requires multiple steps

## Alternatives Considered
1. **Binary safe/unsafe** — Rejected: too coarse
2. **Free-form states** — Rejected: too complex to verify
3. **Numeric severity levels** — Rejected: no clear transition rules

## References
- [PhoenixGuard README](../../PhoenixGuard/README.md)
- [GuardFSM.tla](../../PhoenixFormal/tla/GuardFSM.tla)

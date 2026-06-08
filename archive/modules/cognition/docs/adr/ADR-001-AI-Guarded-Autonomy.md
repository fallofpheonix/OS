---\nStatus: Planned\nImplementation: 20%\nConfidence: Conceptual\n---\n# ADR-001: AI Guarded Autonomy

## Status
Revised (2026-06-01)

## Date
2026-05-31

## Context
AI systems can provide valuable insights for security decisions, but direct AI control of kernel-level actuation poses significant risks including hallucination, prompt injection, and adversarial manipulation. However, a purely advisory role limits the system's effectiveness during critical anomalies.

## Decision
AI systems (PhoenixMind) possess Guarded Autonomy. They may propose edits and changes to critical conditions through the AdvisoryEnvelope mechanism. All high-impact actuations require explicit human-in-the-loop approval via a UserPermissionGate before the Warden FSM executes them. The AI is a collaborative participant but cannot bypass user authorization for critical changes.

## Consequences

### Easier
- Clear separation between AI reasoning and system control
- Formal verification of actuation paths
- Protection against AI hallucination

### More Difficult
- Slower response time (AI → approval → actuation)
- More complex approval workflow
- Requires robust approval chain

## Alternatives Considered
1. **Direct AI control** — Rejected: too risky
2. **AI with kill switch** — Rejected: still allows direct control
3. **AI behind approval gate** — Accepted (current design)

## References
- [PhoenixMind README](../../PhoenixMind/README.md)
- [GEMINI.md](../../PhoenixMind/GEMINI.md)
- [advisory/publisher.go](../../PhoenixMind/advisory/publisher.go)

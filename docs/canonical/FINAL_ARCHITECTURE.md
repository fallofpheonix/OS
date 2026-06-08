---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Final Architecture

## Executive Summary
Phoenix is a contract-first system. Executable evidence shows the real failure mode is contract ownership drift, not repository count.

The governing principle is:

> Contract -> Adapter -> Implementation

Only contract packages define public interfaces. Everything else implements them.

## Contract Ownership

### contracts/events/v1
Owns:
- Event
- EventEnvelope
- event metadata
- serialization rules
- compatibility matrix

### contracts/replay/v1
Owns:
- ReplayEngine
- Snapshot
- Reconstruction
- replay semantics

### contracts/security/v1
Owns:
- Actuator
- Containment
- Escalation
- security event types

## Versioning Policy
- Patch: no API changes.
- Minor: backward-compatible changes only.
- Major: breaking changes allowed only in a new major contract version.

## Architecture Fitness Functions
CI must enforce the following checks:

- FIT-001: No duplicate Event definitions.
- FIT-002: Validation may not import runtime internals.
- FIT-003: Guard may not import implementation internals.
- FIT-004: Contracts may not import implementation packages.
- FIT-005: No circular dependencies across bounded contexts.
- FIT-006: Tests may not access private engine state.

## Ownership Matrix
- contracts/events: Architecture Team
- contracts/replay: Validation Team
- contracts/security: Security Team
- runtime: Core Runtime Team
- ledger: Ledger Team
- validation: Validation Team
- guard: Security Team
- truth: Verification Team
- arbiter: Governance Team
- cognition: Intelligence Team
- platform: Platform Team
- observability: Observability Team
- labs: Research Team

Only owners may approve contract changes.

## Non-Split Zones
Keep these together until independent release cadence or runtime boundaries are proven:

- phoenix-cognition: memory, reasoning, planning
- phoenix-platform: api, ui, cli, service layer
- phoenix-infrastructure: storage, network, deployment

## Extraction Readiness Score
Extract only when all are true:
- Contract Stability >= 90
- Dependency Isolation >= 80
- No circular dependencies
- Contract compatibility suite passing

## Roadmap
### Phase 1: Contract Stabilization
- Create contract packages.
- Create compatibility test suites.
- Add boundary enforcement to CI.
- Resolve current drift in validation and guard.

### Phase 2: Foundation Extraction
- Extract runtime, events, and ledger only after contract stability is proven.

### Phase 3: Domain Extraction
- Extract validation, guard, truth, arbiter, cognition.

### Phase 4: Platform and Observability Extraction
- Extract platform, observability, and labs only when justified.

## Architectural Debt Register
Document every temporary exception in [ARCHITECTURE_DEBT.md](./ARCHITECTURE_DEBT.md).
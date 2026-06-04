---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Architecture Replacement Plan

## Objective
Replace the implicit architecture with an enforced contract-first architecture.

Success is measured by behavior, not by documentation:

- contracts become authoritative
- adapters become mandatory
- implementations become replaceable
- CI prevents architectural drift
- the current failure surfaces disappear

## Milestone 1: Establish Contract Layer

Create authoritative contract packages under:

- [core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1](../../core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1)
- [core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1](../../core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1)
- [core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1](../../core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1)

Deliverables:

- Event
- EventEnvelope
- Serializer
- EventVersion
- ReplayEngine
- Snapshot
- Reconstructor
- Actuator
- Containment
- Escalation

Acceptance criteria:

- No public interface outside `contracts/*`

## Milestone 2: Create Compatibility Test Suites

Create executable contract suites under:

- `contract-tests/events`
- `contract-tests/replay`
- `contract-tests/security`

Example checks:

- FIT-EVENT-001: Single canonical Event
- FIT-EVENT-002: EventEnvelope compatibility
- FIT-EVENT-003: Serialization determinism
- FIT-REPLAY-001: Replay reconstruction equivalence
- FIT-REPLAY-002: Snapshot consistency
- FIT-SEC-001: Actuator compliance
- FIT-SEC-002: Containment compliance

Acceptance criteria:

- `go test ./contract-tests/...`

## Milestone 3: PhoenixValidation Migration

Remove direct dependency on legacy event and replay implementations.

Current failure surfaces:

- EventEnvelope vs Event
- Replay vs Reconstruct
- currentState access

Deliverables:

- Event adapter: proto EventEnvelope -> contracts/events/v1.Event
- Replay adapter: legacy replay -> contracts/replay/v1.ReplayEngine

Acceptance criteria:

- `go test ./PhoenixValidation/...`
- no internal engine access

## Milestone 4: PhoenixGuard Migration

Remove implementation-owned actuator definitions.

Current failure surface:

- Actuator vs Kill

Deliverable:

- contracts/security/v1.Actuator

Acceptance criteria:

- `go test ./PhoenixGuard/...`
- guard implementations comply with contract-defined interfaces

## Milestone 5: Architecture Enforcement

Prevent reintroduction of the old architecture.

CI checks:

- FIT-ARCH-001: No duplicate Event definitions
- FIT-ARCH-002: Contracts cannot import implementations
- FIT-ARCH-003: Validation cannot import runtime internals
- FIT-ARCH-004: Guard cannot import implementation internals
- FIT-ARCH-005: No circular dependencies

Acceptance criteria:

- all checks run automatically in CI
- any violation fails merge

## Milestone 6: Internal PhoenixCore Separation

Create bounded contexts before repository extraction.

Target structure:

- contracts
- runtime
- events
- ledger
- adapters

Remove direct implementation imports such as:

- runtime -> event implementation imports
- ledger -> runtime implementation imports
- validation -> internal engine imports

Replace them with:

- contract -> adapter -> implementation

Acceptance criteria:

- dependency graph contains no forbidden imports

## Milestone 7: Extraction Readiness

Measure readiness with:

- Contract Stability
- Dependency Isolation
- Coverage
- Release Independence

Extraction gate:

- Contract Stability >= 90
- Dependency Isolation >= 80
- Coverage >= 70
- Contract Suite Passing
- No Circular Dependencies

## Milestone 8: Repository Extraction

Extraction order:

1. phoenix-events
2. phoenix-runtime
3. phoenix-ledger
4. phoenix-validation
5. phoenix-guard
6. phoenix-truth
7. phoenix-arbiter
8. phoenix-cognition
9. phoenix-platform
10. phoenix-observability
11. phoenix-labs

## Weekly Scorecard

Track each area weekly:

- Contract Packages
- Compatibility Tests
- Validation Migration
- Guard Migration
- Architecture Tests
- Boundary Enforcement
- PhoenixCore Separation
- Extraction Readiness

## Definition of Done

Architecture replacement is complete only when all are true:

- contracts/events/v1 exists
- contracts/replay/v1 exists
- contracts/security/v1 exists
- contract compatibility suites pass
- PhoenixValidation uses contracts
- PhoenixGuard uses contracts
- architecture fitness functions run in CI
- forbidden imports fail CI
- PhoenixCore bounded contexts exist
- first extraction wave completed successfully

Final status:

- Architecturally Proven
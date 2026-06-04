---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Master Specification (Contract-First)

## 0. Executive Context
PhoenixOS is a contract-first deterministic substrate. Executable evidence showed the primary risk is contract ownership drift, not repository count or bulk implementation failure.

The governing principle is:

> Contract -> Adapter -> Implementation

Only contract packages define public interfaces. Implementations must adapt to contracts, never define competing public interfaces.

### Contract Scope
- Events: canonical event model, event envelope, serialization, compatibility rules.
- Replay: reconstruction interface, snapshot interface, replay semantics.
- Security: actuator interface, containment interface, escalation contract.

---

## 1. Architectural Core
PhoenixOS is organized around bounded contexts with explicit ownership. Internal organization must follow contract-first rules.

### 1.1 Core Contexts
- contracts: public interfaces and compatibility rules.
- runtime: lifecycle, scheduler, clock, process model.
- events: event schema, versioning, serialization.
- ledger: immutable audit trail, hash chain, state history.
- validation: invariants, proofs, replay verification, chaos/fuzzing.
- security: warden, containment, actuation.
- truth: consistency and verification.
- arbiter: governance and policy.
- cognition: memory, reasoning, planning, coordination.
- platform: api, ui, cli, service layer.
- observability: trace, metrics, lineage, monitoring.
- labs: experiments, research, prototypes.

### 1.2 Trust Boundaries
- PRIVILEGED: kernel probes, platform deployment hooks, boot validation.
- TRUSTED: runtime, ledger, contract adapters, warden enforcement.
- SEMI-TRUSTED: cognitive outputs and advisories.
- UNTRUSTED: operator input, external HTTP APIs, research artifacts.

---

## 2. Event Lifecycle & Execution Flow
1. Contract-defined event is produced by an adapter.
2. Runtime validates and routes the event through the bus.
3. Ledger records the event and maintains hash lineage.
4. Replay verifies deterministic reconstruction against the replay contract.
5. Security contracts drive containment decisions through the warden.
6. Recovery uses the ledger and snapshot contract only.

---

## 3. World State & Determinism
### 3.1 Reconstruction Requirement
Given a Genesis state and an ordered contract-compliant event stream, PhoenixOS must reconstruct the same final state deterministically.

### 3.2 Canonical State
- Genesis state is immutable.
- Ledger history is append-only.
- Replay output is deterministic for identical inputs.

### 3.3 Derived State
- Authority state
- Security posture
- Cognitive projections

---

## 4. Contract Ownership Rules
Only contract packages may define public interfaces.

### 4.1 contracts/events
Owns:
- Event
- EventEnvelope
- serialization rules
- compatibility matrix

### 4.2 contracts/replay
Owns:
- ReplayEngine
- Snapshot
- Reconstruction
- replay semantics

### 4.3 contracts/security
Owns:
- Actuator
- Containment
- Escalation
- security event types

---

## 5. Contract Versioning Policy
- Patch: no API changes.
- Minor: backward-compatible additions only.
- Major: breaking changes allowed with a new major version.

Supported versions must have compatibility tests. Every implementation must pass the full contract suite for every supported version.

---

## 6. Architecture Fitness Functions
The following checks are CI-enforced architectural tests:

- FIT-001: No duplicate Event definitions.
- FIT-002: Validation may not import runtime internals.
- FIT-003: Guard may not import implementation internals.
- FIT-004: Contracts may not import implementation packages.
- FIT-005: No circular dependencies across bounded contexts.
- FIT-006: Tests may not access private engine state.

---

## 7. Ownership Matrix
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

Only owners may approve changes to their boundary contracts.

---

## 8. Non-Split Zones
Keep together until evidence proves independent release cadence or runtime boundaries:

- phoenix-cognition: memory, reasoning, planning
- phoenix-platform: api, ui, cli, service layer
- phoenix-infrastructure: storage, network, deployment

---

## 9. Extraction Readiness
Extraction is allowed only when:
- contract stability >= 90
- dependency isolation >= 80
- contract suite passes
- no circular dependencies remain

---

## 10. Cross-Context Communication
All cross-context communication must pass through versioned contracts and adapters.

No implementation may depend directly on another implementation's private types, fields, or runtime internals.

---

## 11. Phase Order
1. Stabilize contracts.
2. Extract foundation contexts.
3. Extract domain contexts.
4. Extract platform and observability contexts.
5. Split further only if scale or release cadence justifies it.

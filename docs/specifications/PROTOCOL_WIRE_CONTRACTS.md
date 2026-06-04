---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Protocol Wire Contracts (v1.1)

This document defines the canonical wire contracts and message structures for the PhoenixOS ecosystem. All communications within the substrate must adhere to these deterministic protocols.

## 1. Protobuf Package Layout
All messages belong to versioned contract packages.

- `contracts/events/v1`
- `contracts/replay/v1`
- `contracts/security/v1`

### Mandatory Fields:
- `schema_version`
- `created_at`
- `replay_sequence`
- `validation_hash`

### Wire Rules:
- No nullable ambiguity.
- No untyped payloads (`map[string]interface{}`).
- All repeated collections must be ordered and deterministic.
- Breaking changes require a major version bump.
- Contract packages are the source of truth; implementations must adapt to them.

---

## 2. Message Specifications

### 2.1 Event Envelope
- `event_id`: Globally unique identifier.
- `parent_event`: Causal predecessor (DAG link).
- `causal_chain`: Array of ancestor hashes.
- `evidence_hash`: Hash of supporting artifacts.
- `trust_score`: Calculated confidence value.
- `signature`: Cryptographic signature of payload.

### 2.2 Ledger Entry
- `entry_id`: Sequence identifier.
- `evidence_record`: Signed telemetry or fact.
- `state_before` / `state_after`: System state hashes for verification.
- `transition_proof`: Quorum signatures or invariant proofs.

### 2.3 Advisory Envelope (AI Output)
- `advisory_id`: Identifier for the recommendation.
- `recommendation`: Proposed system transition.
- `risk_score`: Potential impact evaluation.
- `forbidden_actions`: Explicit constraints for the session.
- **Rule:** AI may ONLY emit advisories; it may NEVER execute actions directly.

### 2.4 Enforcement Request
- `request_id`: Identifier for the actuation.
- `target`: PID or Resource ID.
- `action`: Class from the Containment Ladder.
- `rollback_plan`: Inverse steps for failure recovery.
- **Rule:** All enforcement must be bounded by a timeout.

---

## 3. OpenAPI REST Groups
Root: `/api/v1`
- `/events`: Stream of signed events.
- `/ledger`: Merkle-DAG query interface.
- `/advisories`: AI-generated recommendations.
- `/enforcement`: Actuation command channel.
- `/trace`: Causal graph traversal.
- `/replay`: Deterministic state reconstruction control.

---

## 4. Event Bus Priority Topics
- `phoenix.events.critical`: Bypasses non-critical queues; immediate processing.
- `phoenix.events.high`: Preserves order within correlation IDs.
- `phoenix.events.normal`: Preserves replay sequence order.
- `phoenix.events.background`: Lowest priority; audit-safe retention.

---
## 4. Compatibility Policy
- Minor versions must remain backward compatible.
- Major versions may break compatibility only when a new versioned contract package is introduced.
- Patch releases may not change APIs.

**Status:** Canonical baseline for future code generation.

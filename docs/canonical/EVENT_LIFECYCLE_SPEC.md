---
Status: Implemented
Implementation: 100%
Confidence: Proven
---
# PhoenixOS Event Lifecycle Specification (v1.2)

## 1. Canonical EventID Generation (EVENT-ID-001)
To ensure bit-perfect Replay stability and prevent collisions across system restarts without relying on non-deterministic wall clocks, every `EventID` MUST be derived using the following deterministic formula:

```text
EventID = SHA256(
    GenesisHash   || 
    LedgerHeight  || 
    ValidatorID   || 
    SequenceNo    || 
    PayloadHash
)
```

- **GenesisHash**: The immutable root hash of the system (prevents cross-network collisions).
- **LedgerHeight**: The finalized block height at the time of event proposal.
- **ValidatorID**: The public key of the proposing node/identity.
- **SequenceNo**: A monotonic counter incremented by the validator within the current block/height.
- **PayloadHash**: SHA-256 of the canonical event payload.

## 2. Canonical Event Schema
An `Event` is the atomic, immutable record of an authorized action or state transition.

| Field | Type | Description |
|-------|------|-------------|
| `EventID` | `string` | Deterministic unique identifier (see EVENT-ID-001). |
| `ParentID` | `string` | ID of the causal predecessor event (DAG link). |
| `AuthorityID` | `string` | Verified authority under which this event occurred. |
| `IdentityID` | `string` | Identity (keypair/node) that generated the event. |
| `LogicalTime` | `uint64` | Monotonically increasing Lamport timestamp (reconstructed during Replay). |
| `Evidence` | `[]string` | Artifact Hashes justifying the action. |
| `Signature` | `string` | Cryptographic signature over the payload. |
| `Payload` | `bytes` | Domain-specific state transition data. |

---

## 3. Artifact & Checkpoint Schemas
### 3.1 Artifact Manifest
| Field | Type | Description |
|-------|------|-------------|
| `Hash` | `string` | SHA-256 of the content. |
| `Version` | `string` | SemVer identifier. |
| `Signer` | `string` | IdentityID that verified the artifact. |

### 3.2 Checkpoint Format
| Field | Type | Description |
|-------|------|-------------|
| `StateHash` | `string` | Canonical SHA-256 hash of entire system state (STATE-001). |
| `ReplayOffset`| `uint64` | Ledger index/logical time of the snapshot. |

---

## 4. Async Actuation & Completion Lifecycle (EVENT-ACT-001)
Actuations follow a **PENDING → COMPLETION** pattern. The Replay engine maintains a `PendingMap[CauseID]` to pair these events.

### 4.1 Eviction Policy
To prevent memory exhaustion from uncompleted or failed actuations, the `PendingMap` implements a strictly-timed eviction lifecycle:

| State | Threshold | Action |
|-------|-----------|--------|
| **WARN** | $T = 2$ epochs | Emit `WARDEN_STALL_WARNING` telemetry with `CauseID`. |
| **STALL** | $T = 5$ epochs | If no `ACTUATOR_ALIVE` heartbeat within 1 epoch, mark `STALLED_FAILURE`. |
| **EVICT** | $T = 5$ epochs | Remove from `PendingMap`; subsequent COMPLETIONs are rejected. |

- **Heartbeat Requirement**: The async actuator goroutine MUST emit an `ACTUATOR_ALIVE` event every epoch to maintain its entry in the `PendingMap`.

## 5. Schema Evolution & Versioning
To ensure long-term replay sustainability:

- **Additive Changes**: New fields are permitted and optional for legacy engines.
- **Breaking Changes**: Fields cannot be removed. New `Event` types must be introduced in a new major contract version.
- **In-Memory Migration**: The replay adapter transforms historical payloads into the current internal target version before application.
- **Deterministic Migrations**: Migration logic must be immutable and bit-perfect.

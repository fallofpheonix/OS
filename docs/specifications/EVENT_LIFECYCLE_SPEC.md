---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Event Lifecycle Specification (v1.1)

## 1. Canonical Event Schema
An `Event` is the atomic, immutable record of an authorized action or state transition.

| Field | Type | Description |
|-------|------|-------------|
| `EventID` | `string` | Globally unique identifier (UUID/ULID). |
| `ParentID` | `string` | ID of the causal predecessor event (DAG link). |
| `AuthorityID` | `string` | Verified authority under which this event occurred. |
| `IdentityID` | `string` | Identity (keypair/node) that generated the event. |
| `LogicalTime` | `uint64` | Monotonically increasing Lamport timestamp. |
| `Evidence` | `[]string` | Artifact Hashes justifying the action. |
| `Signature` | `string` | Cryptographic signature over the payload. |
| `Payload` | `bytes` | Domain-specific state transition data. |

---

## 2. Artifact & Checkpoint Schemas
### 2.1 Artifact Manifest
| Field | Type | Description |
|-------|------|-------------|
| `Hash` | `string` | SHA-256 of the content. |
| `Version` | `string` | SemVer identifier. |
| `Signer` | `string` | IdentityID that verified the artifact. |

### 2.2 Checkpoint Format
| Field | Type | Description |
|-------|------|-------------|
| `StateHash` | `string` | Canonical SHA-256 hash of entire system state. |
| `ReplayOffset`| `uint64` | Ledger index/logical time of the snapshot. |

---

## 3. Schema Evolution & Versioning
To ensure long-term replay sustainability:

- **Additive Changes**: New fields are permitted and optional for legacy engines.
- **Breaking Changes**: Fields cannot be removed. New `Event` types must be introduced in a new major contract version.
- **In-Memory Migration**: The replay adapter transforms historical payloads into the current internal target version before application.
- **Deterministic Migrations**: Migration logic must be immutable and bit-perfect.

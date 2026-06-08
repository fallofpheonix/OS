---
Status: Draft
Implementation: 0%
Confidence: High
---
# SPEC-001: Canonical System State

> This specification defines the authoritative, hierarchical structure of the Phoenix System State. 
> Any state that cannot be represented within this schema is "non-canonical" and must be treated as volatile/ephemeral.

## 1. Root State Hierarchy

```text
SystemState {
    Header {
        Version: SemVer
        Tick: uint64 (Authoritative Time)
        GenesisHash: SHA256
        StateHash: SHA256 (Merkle Root of entire state)
    }
    Layers {
        Foundation: FoundationState
        Assurance: AssuranceState
        Governance: GovernanceState
        Cognition: CognitionState
        Platform: PlatformState
    }
    Game: GameWorldState
}
```

## 2. Layered State Definitions

### 2.1 FoundationState (STATE-001: Consensus-Critical)
- **LedgerIndex**: Current block height.
- **UncommittedEvents**: Queue of events awaiting ledgering.
- **Registry**: Map of active Identities and their public keys.

### 2.2 AssuranceState (STATE-001: Consensus-Critical)
- **Capabilities**: Active capability tokens and their attenuation status.
- **WardenPolicies**: Hash of currently active security policies.

### 2.3 GovernanceState (STATE-001: Consensus-Critical)
- **TruthModel**: Merkle tree of "Proven" facts derived from the Ledger.
- **EvolutionGating**: Status of current self-modification proposals.

### 2.4 ObservabilityState (STATE-002: Volatile/Telemetry)
- **SubsystemTelemetry**: High-volume, non-deterministic performance metrics.
- **WardenHistory**: Transient logs of shadow-mode evaluations.
- **Audit Integrity**: Every entry in the `WardenHistory` ring buffer MUST include a `PrevHash` pointer to form an internal hash-chain.
- **Anchoring**: The HEAD hash of the `WardenHistory` chain MUST be included in the next `CommitQC`'s metadata field (not the `StateHash`) to provide a verifiable anchor for forensic investigators.
- **NOTE**: Observability State MUST NOT be included in the `StateHash` calculation to prevent Replay Divergence.

## 3. Serialization Rules

- **Deterministic Encoding**: Canonical JSON or Protobuf with sorted keys. 
- **Padding/Alignment**: Fixed-width fields where possible to ensure identical byte representations across architectures.
- **Floating Point**: Prohibited. Must use fixed-point `Dec64` for all decimal values (INV-017).
- **Stable Hashing**: Merkle root computation must strictly follow the `SubsystemManifest` order.

## 4. State Transitions

1. `ProcessEvent(E, S_t) -> (S_t+1, I)`
2. Where `E` is an authorized event.
3. `S_t` is the current state.
4. `S_t+1` is the new state.
5. `I` is the set of generated Intents.

--- 
*Refer to [SPEC-002: Runtime State](./SPEC-002-Runtime-State.md) for in-memory implementation details.*

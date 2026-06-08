---
Status: Draft
Implementation: 0%
Confidence: High
---
# SPEC-008: Root Authority & Governance Hierarchy

> This specification defines the origin, rotation, and recovery of the system's Root Authority (Sovereign Trust).

## 1. The Root of Trust (ROOT-001)
The **Genesis Root Authority** is established during the `Genesis Event`. 
- The public key of the first **Root Sovereign** is hardcoded in the `genesis.json` of the substrate.
- This key authorizes the initial validator set (Epoch 0).

## 2. Authority Hierarchy
PhoenixOS employs a **Stacked Authority** model to resolve circular trust:

1. **Layer 0: Genesis Artifact**: Immutable commitment to the first Root.
2. **Layer 1: Root Sovereign**: Holds the power to authorize `CONSTITUTION_UPDATE` and `ROOT_ROTATION` events.
3. **Layer 2: Consensus Quorum**: Finalizes all state changes. **Rule**: A Consensus Quorum cannot override a Root Sovereign signature unless the system is in `COMPROMISED` state and a `RECOVERY_QUORUM` (f+1 signatures from different physical hosts) is reached.
4. **Layer 3: Subsystems**: Warden, Mind, Ledger.

## 3. Root Lifecycle (ROOT-002, ROOT-003, ROOT-004)

### 3.1 Root Rotation
The current Root Sovereign can authorize a `ROOT_ROTATION` event to transition authority to a new keypair. This event must be ledgered and signed by both the old and new keys to be valid.

### 3.2 Revocation & Recovery (ROOT-005)
If a Root key is lost or compromised:
- **Revocation**: A `ROOT_REVOCATION` event can be triggered by a **Consensus Supermajority (3/4 weight)**. This immediately moves the system to `COMPROMISED`.
- **Escalation Fallback**: If the 3/4 threshold is unreachable within `T=EpochTimeout` (due to node failure or partition), the system enters a **DEGRADED_SOVEREIGNTY** state. In this state, a standard `2f+1` quorum suffices for revocation, but all resulting actions are marked `UNVERIFIED_RECOVERY` in the ledger and require manual audit at the next stable checkpoint.
- **Recovery**: Transitioning out of `COMPROMISED` after a Root revocation requires a **Consensus Quorum Proof** containing a new Root Key commitment, verified against an out-of-band "Safety Goal" checkpoint.

## 4. Multi-Root Configurations
- Only one **Active Root** can exist per `Epoch`.
- **Shadow Roots** (Backup keys) can be committed to the Governance state but possess no actuation authority until a `ROOT_ROTATION` or `RECOVERY` event occurs.

---
*Refer to [docs/canonical/SPEC-005-Governance-State.md](./SPEC-005-Governance-State.md) for data structures.*

# SPEC-001: Validator Epoch Rules

## Overview
To ensure protocol stability and bit-perfect replay of historical consensus, validator set changes in PhoenixOS are subject to **Epoch Activation**. 

## Rules
1.  **Epoch Duration:** An epoch lasts for a fixed number of logical ticks (default: 100).
2.  **Staging Changes:** When a `MsgEvent` or `EventUpdateValidator` is processed at height `H`, the change is recorded in the `PendingValidators` list of the `WorldState`.
3.  **Activation Window:** Staged changes are promoted to the `ActiveValidators` set ONLY at the next epoch boundary (`ws.Tick % EpochInterval == 0`).
4.  **Voting Rights:** Only nodes in the `ActiveValidators` set for the current epoch are authorized to sign events, cast votes, or form Quorum Certificates (QCs).
5.  **Historical Stability:** This ensures that for the duration of an epoch, the validator set is immutable, preventing round-level equivocation during membership transitions.

---

# SPEC-002: QC Persistence Rules

## Overview
Quorum Certificates (QCs) represent proof of consensus for a specific state transition. To maintain a verifiable forensic record, QCs must be durable.

## Rules
1.  **Ledger Embedding:** Every `Event` in the ledger MAY contain a `QuorumCertificate`.
2.  **Required for Finality:** Events at the `Commit` phase of HotStuff MUST include a QC before being considered finalized in the substrate.
3.  **Digest Independence:** The `QuorumCertificate` is excluded from the `Event.Digest` to avoid circular dependencies (the QC is a proof *of* the digest).
4.  **Verification on Replay:** The Replay Engine MUST verify the signatures within the QC against the `ActiveValidators` set as it existed when the event was recorded.

---

# SPEC-003: Historical QC Verification

## Overview
A node replaying the ledger from genesis must be able to verify that every historical consensus decision was valid according to the rules of its epoch.

## Rules
1.  **Causal Set Retrieval:** The validator set used for QC verification must be the `ActiveValidators` set associated with the `Event.Tick`.
2.  **Threshold Enforcement:** A QC is only valid if it contains unique signatures from a 2f+1 quorum of the historical active set.
3.  **State Transition Matching:** The `QC.StateHash` MUST match the `StateHash` produced by replaying the event. Discrepancies indicate a semantic divergence and must halt replay.

---

# SPEC-004: View Change Rules (Draft)

## Overview
HotStuff safety relies on deterministic leader transitions and view change proofs.

## Rules
1.  **Leader Selection:** Derived via `ElectLeader(validators, round)` using round-robin of the sorted active set.
2.  **View Synchronization:** Nodes increment their `ViewNumber` upon receiving a QC for the previous view or a Timeout Certificate (TC).
3.  **Safe-Node Rule:** A node only votes for a proposal in view `V` if the proposal extends the highest QC it has locked or if it sees a QC for a higher view.

---

# SPEC-005: Timeout Certificate Rules (Draft)

## Overview
Liveness in the presence of faulty leaders is guaranteed by Timeout Certificates.

## Rules
1.  **Local Timeout:** Nodes trigger a local timer upon entering a new round.
2.  **Timeout Aggregation:** If the timer expires, the node broadcasts a `SignedTimeout` message.
3.  **TC Formation:** A Timeout Certificate is formed when 2f+1 timeout messages for the same view are collected.
4.  **View Advance:** Receiving a valid TC allows nodes to safely transition to the next view and elect a new leader.

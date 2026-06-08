# PHOENIX MATRIX: BFT CONSENSUS SPECIFICATION

**Status:** AUTHORITATIVE
**Document ID:** BFT-SPEC-004
**Consensus Family:** HotStuff (Pipelined)

## 1. Protocol Objects

### 1.1 Quorum Certificate (QC)
A QC is the authoritative proof of agreement for a specific state.
```go
type QuorumCertificate struct {
	Epoch            uint64            `json:"epoch"`
	Round            uint32            `json:"round"`
	Index            uint64            `json:"index"`
	ValidatorSetHash Hash              `json:"validator_set_hash"`
	StateHash        Hash              `json:"state_hash"`
	Signatures       map[string][]byte `json:"signatures"` // hex(PublicKey) -> Ed25519 Signature
}
```

### 1.2 Vote
Individual validator agreement on a proposed round.
```go
type Vote struct {
	Epoch     uint64 `json:"epoch"`
	Round     uint32 `json:"round"`
	Index     uint64 `json:"index"`
	StateHash Hash   `json:"state_hash"`
}
```

## 2. Authentication & Finality (BFT-009, BFT-010)

- **What is Signed?** 
    1. **Event Hash**: Validators sign the `SignedEnvelope.Digest()` which includes the raw event payload.
    2. **State Hash**: Validators sign the `Vote`, committing to the `Post-State Root` (STATE-001) resulting from the event application.
- **What is Finalized?** 
    - A **Block** (sequence of events) and its corresponding **State Root** are finalized only upon the collection of $2f + 1$ Commit votes (forming a `CommitQC`).

## 3. Leader Selection (VRF)

To prevent targeted DoS and eclipse attacks, leader selection is pseudo-randomized using a **Verifiable Random Function (VRF)**.

- **Algorithm**: ECVRF-ED25519-SHA512-ELLIGATOR2.
- **Seed**: The `StateHash` of the last finalized block (Index $N-1$).
- **Selection**: $Leader = \text{argmin}(VRF_{Proof}(Seed || Round || Epoch))$.
- **Verification**: The leader MUST include their `VRF_Proof` in the `SignedEnvelope`. Validators reject any proposal without a valid proof matching the deterministic selection formula.
- **Timing**: Proposals must be received within $\Delta$ of the round start. Withholding a VRF proof is treated as a leader failure and penalized via **Reputation Docking**.

## 4. The HotStuff Pipelined Loop
Phoenix OS uses a 3-chain finality rule:

0.  **INGESTION (MEMPOOL):** Events MUST pass `EventApplier.Validate()` at mempool ingestion time. Events that fail validation are rejected immediately and never proposed. A block containing a semantically invalid event indicates a Byzantine leader and MUST be rejected by all honest validators; however, the ingestion gate ensures that semantic validation is not a bottleneck during the proposal stage.
1.  **PREPARE:** Leader proposes a new event with VRF proof. Validators verify the VRF proof, signature, and include a `Vote`.
2.  **PRE-COMMIT:** Collection of $2f + 1$ Prepare votes forms a `PrepareQC`.
3.  **COMMIT:** Collection of $2f + 1$ Pre-Commit votes forms a `PreCommitQC`.
4.  **FINALITY (DECIDE):** Collection of $2f + 1$ Commit votes forms a `CommitQC`. The block is committed to the Ledger.

## 5. Validator Governance (BFT-012, BFT-013)

- **Set Transitions**: `UPDATE_VALIDATOR` events must be included in a finalized block to take effect.
- **Authority Boundary**: 
    - The **Old Validator Set** MUST sign the transition block (containing the update) with a quorum of $2f_{old} + 1$.
    - The **New Validator Set** gains authority strictly at `Round 0` of the next `Epoch`.
- **Historical Verification**: During Replay, signatures are verified against the **Historical Validator Set** identified by the `ValidatorSetHash` in the QC.
- **Algorithm Migration**: Genesis trust anchors (Ed25519) can be rotated via a `PROTOCOL_UPGRADE` event signed by the old keys, committing to a new Root Public Key for the new algorithm. Replay verifies the chain from the last **Trusted Snapshot**.

## 6. Persistence & Recovery

- **Event Uniqueness (BFT-011)**: No two valid blocks can contain the same `EventID`. The Ledger enforces global uniqueness based on the deterministic `EventID` formula (EVENT-ID-001).
- **Restart Safety**: Validator sequences and Lamport clocks are reconstructed from the last finalized `CommitQC` in the Ledger.

---
*Authorized by Phoenix Sovereign Governance*

# PhoenixOS State Architecture (Authoritative v3)

> **Status:** FROZEN (Zero-Trust Implementation Prerequisite)
> **Mandate:** INV-001, INV-004, INV-022

## 1. Universal Object Model (ARCH-301)

The smallest object that crosses every layer is the **`Event`**. It is defined in `internal/contracts` and represents a pure semantic intent (e.g., Spawn, Move).

---

## 2. Verified Transition (VT-001 .. VT-004)

### **VT-001: The Nature of Transition**
A transition is a **[A] Event** that has been enriched with its execution context (Height, Epoch). It is the semantic payload that the State layer consumes to evolve the `worldState`.

### **VT-002: Construction & Ownership**
The **[C] Runtime Dispatcher** constructs the `VerifiedTransition`. It acts as the "Air-Gap" between the Authority layer (Consensus) and the Physics layer (State).

### **VT-003/004: Persistence & Reconstruction**
*   **Persistence:** **[B] NO.** We do not persist transitions directly; we persist the **`FinalizedBlock`**.
*   **Replay:** **[A] YES.** Replay reconstructs the transition by extracting the `Event` from the persisted block and wrapping it in the current height/epoch context.

---

## 3. Finalized Block (FB-001 .. FB-005)

### **FB-001: Block Schema**
The authoritative unit of the Ledger is the **`FinalizedBlock`**:
*   `Height` (uint64): Monotonic height.
*   `Epoch` (uint64): Consensus epoch.
*   `Round` (uint32): HotStuff round.
*   `Proposer` (NodeID): Identity of the proposer.
*   `PrevBlockHash` (Hash): Link to history.
*   `Events` ([]Event): Batch of semantic intents.
*   `StateRoot` (Hash): Commitment to the state *after* executing these events.
*   `QC` (QuorumCertificate): The proof of 2f+1 agreement.

### **FB-002 .. FB-005: Consensus Invariants**
*   **FB-002 (Height):** **[A] YES.** `Height` participates in the `BlockHash` to prevent height-substitution attacks.
*   **FB-003 (StateRoot):** **[A] YES.** `StateRoot` participates in the `QC` to prove agreement on execution results.
*   **FB-004 (Voting):** Validators vote on the **Block Digest** (which includes both `ProposalHash` and `StateRoot`).
*   **FB-005 (Replay):** **[B] NO.** Replay MUST verify the `QC` for every block unless explicitly running in a "Non-Authoritative Audit" mode.

---

## 4. StateHash Domain (HASH-001 .. HASH-003)

### **HASH-001: The Proof Domain**
The `StateHash` proves **[C] Physics + Governance + Forensics**.

### **HASH-002/003: Sensitivity**
*   **Validator Change:** **[A] YES.** StateHash changes even if physics is identical.
*   **Sequence Change:** **[A] YES.** StateHash changes to anchor the state to its causal provenance.

---

## 5. Replay Scope (REP-101 .. REP-103)

### **REP-101 .. REP-103: Reconstruction Boundaries**
*   **REP-101 (Discarded Proposals):** **[B] NO.** Discarded branches are not replayed.
*   **REP-102 (Scope):** Replay reconstructs the **Finalized History**.
*   **REP-103 (Audit Log):** **[A] YES.** A separate **Consensus Audit Log** is required for debugging Liveness/View-Changes, but it is decoupled from the authoritative State Ledger.

---

## 6. Dependency Graph (Refined)

```text
       internal/contracts  (Event, Block, Hash)
               ↑
       internal/state      (Physics / WorldState)
               ↑
       game/engine         (VM / Physics Rules)
               ↑
       internal/runtime    (Dispatcher / StateGuard)
               ↑
       internal/consensus  (Authority / HotStuff)
               ↑
       game/replay         (Orchestration / Loop)
```

**Boundary Rules:**
1. `internal/state` is a leaf (imports only contracts/math).
2. `internal/runtime` mediates all traffic to `state`.
3. `internal/consensus` never touches `state` directly; it submits blocks to `runtime`.

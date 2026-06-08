# PhoenixOS Block Architecture (Authoritative)

> **Status:** DRAFT (Architectural Resolution)
> **Prerequisite:** internal/contracts implementation
> **Mandate:** INV-001, INV-010, INV-021

## 1. Dependency Analysis (DEP-001 .. DEP-002)

### **DEP-001: Autonomous Replay**
*   **Can replay rebuild state without consensus?** **[A] YES.**
*   **Rationale:** Replay is a consumer of **FinalizedBlocks** stored in the Ledger. Since a `FinalizedBlock` contains its own proof (QC), the Replay Engine only needs the `internal/state` (the destination) and `internal/contracts` (the type definitions) to reconstruct state. It does not need the active `consensus` FSM (Prepare/PreCommit logic).

### **DEP-002: Historical Verification**
*   **Can replay verify historical QCs without consensus?** **[A] YES.**
*   **Resolution:** The logic for **Quorum Verification** (checking 2f+1 signatures against a validator set) is moved to **`internal/contracts`** or a dedicated `internal/crypto` leaf. The `consensus` package remains responsible for the *formation* of QCs, while the rest of the system remains responsible for the *verification* of QCs.

---

## 2. Runtime Boundary (RT-001 .. RT-004)

### **RT-001: Authority Verification**
*   **Can Runtime verify authority?** **[B] NO.**
*   **Rationale:** Authority verification (Signature/QC checks) happens in **Consensus** (for incoming mesh traffic) or **Replay** (for ledger traffic). Runtime is a pass-through "Dispatcher."

### **RT-002: Reordering**
*   **Can Runtime reorder transitions?** **[B] NO.**
*   **Rationale:** Reordering is a **Consensus** concern (the Ordering Buffer). Runtime executes blocks in the strict order it receives them.

### **RT-003: Rejection**
*   **Can Runtime reject a block?** **[A] YES.**
*   **Rationale:** Only for **Physics/Semantic violations** (e.g., "Invalid VM Opcode" or "Entity out of bounds"). It CANNOT reject for authority reasons.

---

## 3. StateRoot Semantics (ROOT-001 .. ROOT-003)

### **ROOT-001: Execution Ownership**
*   **Who executes proposals?** **Every Validator.** 
*   **The Flow:** 1. Leader proposes Block -> 2. Every Validator executes the events locally via **`game/engine`** -> 3. Every Validator derives the same **`StateRoot`** -> 4. Validators vote on the `(ProposalHash + StateRoot)` tuple.

### **ROOT-003: Calculation Source**
*   **Where is StateRoot calculated?** **[A] State/Engine.**
*   **Rationale:** Only the Physics Layer has the internal `worldState` visibility required to produce a root hash. Consensus is "State Blind."

---

## 4. Event & Block Lifecycle (EVT, BLK, REP)

### **EVT-001 .. EVT-004: Event Invariants**
*   **EVT-001 (Immutability):** **[B] NO.** Events are immutable once signed.
*   **EVT-002 .. 004 (Context):** **[B] NO.** The `Event` carries only semantic data. `Height`, `Epoch`, and `Sequence` are part of the **`FinalizedBlock`** or **`SignedEnvelope`** container. This keeps the semantic unit clean.

### **BLK-001 .. BLK-003: Block Invariants**
*   **BLK-001 (Multi-Validator):** **[A] YES.** A block can contain events from multiple validators (batched by the leader).
*   **BLK-002 (Conflicts):** **[B] NO.** Conflicting events within a single block are prohibited.
*   **BLK-003 (Conflict Validation):** **[A] Engine.** The Engine's semantic validation gate detects and rejects conflicting physics intents during the local execution phase.

### **REP-201 .. REP-203: Replay Proofs**
*   **REP-201 (Validator Reconstruction):** **[A] YES.** Replay Advance is tick-by-tick; historical sets are always re-materialized.
*   **REP-202 (StateRoot Verification):** **[A] YES.** Replay re-executes the block and asserts `ActualStateRoot == Block.StateRoot`.
*   **REP-203 (Forgery Detection):** **[A] YES.** Replay verifies the embedded `QC` against the re-materialized historical validator set.

---

## 5. Corrected Dependency Graph

```text
       internal/contracts  (Event, FinalizedBlock, QC, Hash)
               ↑
       internal/state      (Physics / worldState)
               ↑
       game/engine         (Physics Rules / OpCodes)
               ↑
       internal/consensus  (Authority Formation / HotStuff)
               ↑
       internal/runtime    (Dispatcher / Transport)
               ↑
       game/replay         (Forensic Reconstruction)
```

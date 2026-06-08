# PhoenixOS Contract Freeze (Revision 2)

> **Status:** FROZEN (Universal Schema Prerequisite)
> **Mandate:** INV-004, INV-010, INV-022

## 1. Identity Model (ID-001 .. ID-103)

### **ID-001: NodeID Definition**
*   **Definition:** A `NodeID` is a **[A] Public Key Hash** (SHA-256 of the Ed25519 Public Key).
*   **Properties:** 32-byte fixed length. Participating in all digests where identity is required.
*   **Relationship:** `ValidatorID` and `NodeID` are **[A] Identical**.

---

## 2. Event Model (EVT-301 .. EVT-403)

### **EVT-301/402: Universal Payload**
The `Event` struct is refactored to use a **generic Payload []byte**. This decouples the substrate from the game/physics model and allows for `UPDATE_VALIDATOR` or other system events to be encoded without schema changes.

```go
type Event struct {
    Version uint16
    Type    EventType
    Payload []byte // Opaque semantic data
}
```

### **EVT-403: Schema Evolution**
Evolution is handled via the `Version` field (2-byte header). Deserializers for a specific `Type` will check the `Version` before interpreting the `Payload`.

---

## 3. QC Model (QC-401 .. QC-404)

### **QC-401: Deterministic Signatures**
The `map[string][]byte` is **[B] Forbidden**. It is replaced by a sorted slice of `SignatureEntry`.

```go
type SignatureEntry struct {
    ValidatorID NodeID
    Signature   []byte
}
```

*   **Ordering (QC-403):** Entries are sorted lexicographically by `ValidatorID` (raw bytes).
*   **Uniqueness (QC-404):** Duplicate validator signatures within a single QC are prohibited.

---

## 4. Block Model (BID-001 .. BLK-103)

### **BLK-101: BlockID (The Canonical Anchor)**
The `BlockID` is the **Hash(BlockHeader)**. The header commits to:
`Height | Epoch | Round | Proposer | PrevBlockHash | MerkleRoot(Events) | StateRoot`.

*   **Participation (BLK-102):** **[A] YES.** `StateRoot` participates in `BlockID` to bind consensus agreement to execution results.
*   **Exclusion (BLK-103):** **[B] NO.** The `QC` does NOT participate in the `BlockID` (the QC is a proof *of* the ID).

---

## 5. Frozen Schema Decisions

| ID | Category | Frozen Resolution |
| :--- | :--- | :--- |
| **CONTRACT-011** | **NodeID Encoding** | Fixed 32-byte binary. |
| **CONTRACT-012** | **QC Ordering** | Lexicographical by NodeID. |
| **CONTRACT-013** | **Event Extensibility**| Generic `Payload []byte`. |
| **CONTRACT-015** | **BlockID Definition**| Hash of Header (Includes StateRoot). |
| **NAMING** | **Transition** | Rename `VerifiedTransition` to `AppliedEvent`. |

---

## 6. Dependency Graph (Frozen)

```text
       foundation/math      (Arithmetic)
               ↑
       internal/contracts  (Canonical Types: Event, Block, QC)
               ↑
       internal/protocol   (Behavior: Digests, Serialization)
               ↑
       internal/state      (Physics / WorldState)
               ↑
       game/engine         (VM / Physics Rules)
               ↑
       internal/consensus  (Authority Formation / HotStuff)
```

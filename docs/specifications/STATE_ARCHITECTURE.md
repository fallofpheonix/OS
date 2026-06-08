# PhoenixOS State Architecture (Authoritative)

> **Status:** FROZEN (Revision 2 - Zero-Trust Resolved)
> **Mandate:** INV-001, INV-004, INV-022

## 1. Ownership & Authority Map

| Component | Authoritative Owner | Responsibility |
| :--- | :--- | :--- |
| **Physics State** | `internal/state` | Entity positions, statuses, and physics logic. |
| **Governance State** | `internal/consensus` | Active validator set and epoch boundaries. |
| **Forensic State** | `internal/state` | Last-seen sequences and tick counters. |
| **Mutation Authority** | `internal/state.StateGuard` | The serial gate for all state transitions. |
| **Contract Types** | `internal/contracts` | Canonical `Event` and `Hash` definitions. |
| **Reconstruction Authority** | `game/replay` | The mechanism for replaying the Ledger into State. |
| **Persistence Authority** | `internal/ledger` | Durable storage of SignedEnvelopes. |

---

## 2. The Authority Boundary (Authority Token Pattern)

### **Q301/302: Canonical Event Ownership**
The `Event` type is moved to **`internal/contracts`**. This ensures both `state` and `consensus` can import it without circular dependencies or ownership ambiguity.

### **Q305: Verified Mutation (The Authority Token)**
To prevent the "Proof-Bit Bypass" and solve the Go unexported method limitation, `state.StateGuard` uses an **Authority Token**.

```go
// internal/state/guard.go
type AuthorityToken struct {
    _ struct{} // Unexported field prevents external instantiation
}

func (g *StateGuard) Apply(token AuthorityToken, e contracts.Event) error {
    // Only callers holding a valid token can trigger a mutation
}
```

*   **Initialization:** `internal/state` provides a `RequestAuthorityToken()` function that MUST only be called during the system's one-time secure boot or by a package with explicitly authorized metadata.
*   **Enforcement:** Since `AuthorityToken` cannot be instantiated by `consensus`, the consensus layer must receive it from the bootloader, creating a chain of trust.

---

## 3. Forensic State Invariants

### **SEQ-201 .. SEQ-203: The Forensic Anchor**
*   Sequence tracking is **Forensic State**.
*   **StateHash Participation:** Sequences and Validator Membership MUST participate in the `StateHash`.
*   **Invariant:** `StateHash = Hash(Physics + Governance + Forensics)`. This ensures the root hash commits to the entire provenance of the node, not just the physical coordinates.

---

## 4. Persistence Object Model

### **AUTH-501: The Authoritative Log Unit**
*   **Authoritative Object:** **[C] `SignedEnvelope`.**
*   **Strip Policy (EVENT-004):** The `Event` struct is stripped of `Signature`, `Validator`, and `QC`. Those fields are moved to the `SignedEnvelope` container.
*   **Rationale:** The `Event` is a semantic physics unit; the `Envelope` is an authority unit. Replay preserves the `Envelope` to prove that every historical state transition was signed by an authorized validator at the time of execution.

---

## 5. Dependency Graph (Corrected)

```text
       internal/contracts  (Event, Hash)
               ↑
       internal/state      (The Guard / WorldState)
               ↑
       game/engine         (Physics Rules / Opcodes)
               ↑
       internal/consensus  (Authority / Signatures)
               ↑
       game/replay         (Orchestration / Ledger Loop)
```

**Boundary Rules:**
1. `internal/state` imports `internal/contracts`.
2. `internal/consensus` imports `internal/contracts` AND `internal/state` (to get the Guard).
3. `game/replay` imports `internal/consensus` and `internal/state` to drive the loop.
4. `internal/state` MUST NOT import `internal/consensus` or `game/replay`.

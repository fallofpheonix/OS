# PhoenixOS Protocol Encoding (Revision 2)

> **Status:** FROZEN (Serialization Prerequisite)
> **Mandate:** INV-010, INV-022, INV-023

## 1. Domain-Separated Digests (DIGEST-001 .. DIGEST-003)

To prevent cross-protocol collision attacks, every digest type uses a unique 1-byte domain prefix. All numeric fields are Big-Endian.

### **PROTOCOL-002: Canonical Event Digest**
`Prefix(0x01) | Version(2) | Type(2) | PayloadLen(2) | Payload(N)`

### **PROTOCOL-003: Canonical BlockHeader Digest**
`Prefix(0x02) | Version(2) | Height(8) | Epoch(8) | Round(4) | Proposer(32) | PrevBlockHash(32) | MerkleRoot(32) | StateRoot(32)`

### **PROTOCOL-004: Canonical QC Digest**
`Prefix(0x03) | Version(2) | Epoch(8) | Round(4) | Height(8) | BlockID(32) | StateRoot(32) | ValidatorSetHash(32) | SigCount(2) | [SignatureEntries...]`

---

## 2. Merkle Tree Specification (RFC 6962)

PhoenixOS implements the RFC 6962 standard for deterministic event roots.

*   **Empty Tree:** `SHA256("")` (Fixed: `e3b0c442...`)
*   **Leaf Hash:** `SHA256(0x00 | EventDigest)`
*   **Internal Hash:** `SHA256(0x01 | LeftHash | RightHash)`
*   **Odd Node Rule:** If a level has an odd number of nodes, the last node is **carried up** to the next level without hashing.

---

## 3. Protocol Invariants (QC-501 .. QC-503)

*   **QC-501 (Scale):** Maximum of 1024 signatures per QC.
*   **QC-502 (Ordering):** Signature entries MUST be sorted by `ValidatorID` (raw bytes) before serialization. The protocol layer rejects out-of-order signatures.
*   **QC-503 (Duplicates):** Duplicate validator signatures are detected and rejected by the **Verifier** and **Consensus Engine**. 

---

## 4. Identity & Payload (ID-201 .. PAYLOAD-101)

*   **NodeID (ID-201):** Standardized as the **raw 32-byte Ed25519 Public Key**. No secondary hashing.
*   **Payload (PAYLOAD-101):** Payloads MUST be encoded using **Canonical Binary Encoding** defined per `EventType`. Map-to-JSON is forbidden for consensus-critical payloads.
*   **Length Width:** `PayloadLen` is `uint16` (64 KB cap) to minimize header overhead while supporting all planned OS telemetry.

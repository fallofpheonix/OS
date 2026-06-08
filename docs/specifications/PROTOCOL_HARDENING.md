# PhoenixOS Protocol Hardening (Revision 3)

> **Status:** FROZEN (Substrate Hardening Prerequisite)
> **Mandate:** INV-010, INV-022, INV-023

## 1. Merkle Tree & Proof Invariants (PROTOCOL-014, 017)

### **Authoritative Carry-Up Logic**
Odd nodes at any level are promoted without hashing.

```python
def GetMerkleRoot(leaves):
    if not leaves: return SHA256("")
    nodes = [SHA256(0x00 | leaf) for leaf in leaves]
    while len(nodes) > 1:
        next_level = []
        for i in range(0, len(nodes), 2):
            if i + 1 < len(nodes):
                next_level.append(SHA256(0x01 | nodes[i] | nodes[i+1]))
            else:
                next_level.append(nodes[i])
        nodes = next_level
    return nodes[0]
```

### **Proof Generation (Bottom-Up)**
```python
def GenerateMerkleProof(leaves, target_index):
    nodes = [SHA256(0x00 | leaf) for leaf in leaves]
    proof = []
    curr_idx = target_index
    while len(nodes) > 1:
        is_right_child = (curr_idx % 2 == 1)
        is_last_node = (curr_idx == len(nodes) - 1)
        if is_right_child:
            proof.append(nodes[curr_idx - 1])
        elif not is_last_node:
            proof.append(nodes[curr_idx + 1])
        # Build next level
        next_level = []
        for i in range(0, len(nodes), 2):
            if i + 1 < len(nodes):
                next_level.append(SHA256(0x01 | nodes[i] | nodes[i+1]))
            else:
                next_level.append(nodes[i])
        nodes = next_level
        curr_idx //= 2
    return proof
```

### **Proof Verification (Level-Driven)**
Verification is driven by tree depth, not proof length, to correctly handle carry-up levels.
```python
def VerifyCarryUpProof(leaf_hash, index, total_leaves, proof_hashes, expected_root):
    current_hash = SHA256(0x00 | leaf_hash)
    curr_idx, curr_total, proof_ptr = index, total_leaves, 0
    while curr_total > 1:
        is_right_child = (curr_idx % 2 == 1)
        is_last_node = (curr_idx == curr_total - 1)
        if not (not is_right_child and is_last_node):
            if proof_ptr >= len(proof_hashes): return False
            sibling = proof_hashes[proof_ptr]
            proof_ptr += 1
            if is_right_child: current_hash = SHA256(0x01 | sibling | current_hash)
            else: current_hash = SHA256(0x01 | current_hash | sibling)
        curr_idx //= 2
        curr_total = (curr_total + 1) // 2
    return current_hash == expected_root and proof_ptr == len(proof_hashes)
```

---

## 2. Deterministic Execution Law (PROTOCOL-020)

To support **Execute-Before-Vote**, all state-mutation code MUST satisfy the following rules. Failure to do so is a **Consensus Fault**.

1.  **Zero Float64:** Usage of `float32/64` is strictly prohibited. Use `phxmath.FixedPoint` with truncation.
2.  **Clock Isolation:** Logic MUST NOT call `time.Now()`. The only source of time is the `Tick` counter.
3.  **Entropy Isolation:** Randomness MUST only be derived from the `Seed` field via a deterministic PRNG.
4.  **Iteration Stability:** Map iterations MUST be sorted by key before mutation or hashing.

---

## 3. Resource Limits & Validation (PROTOCOL-016, 021)

### **Independent Constraints**
*   **Primary:** Block payload size MUST be ≤ 1,048,576 bytes.
*   **Secondary:** Total events in a block MUST be ≤ 1,024.
*   **Rejection:** Rejection occurs at the first boundary violated.

### **Failure Codes (Append-Only)**
Existing codes cannot be reused or reinterpreted.
*   `0x01`: ERR_DECODE
*   `0x02`: ERR_VERSION
*   `0x03`: ERR_LIMIT
*   `0x04`: ERR_MERKLE
*   `0x05`: ERR_QC
*   `0x06`: ERR_STATE
*   `0x07`: ERR_DETERMINISM (Execution violation)

### **Refined Validation Order**
1. Decode -> 2. Version -> 3. Size Limits -> 4. Event Decoding -> 5. Merkle Root -> 6. Local Execution -> 7. StateRoot Assert -> 8. QC Signature Check -> 9. Consensus Commit.

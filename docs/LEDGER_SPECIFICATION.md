# PHOENIX MATRIX: LEDGER SPECIFICATION

**Status:** AUTHORITATIVE
**Document ID:** LEDGER-SPEC-001
**Source of Truth:** Event Ledger

## 1. Event Formal Specification
Events are the atomic units of state transition. 

### 1.1 Binary Encoding Rules
*   **Byte Order:** BigEndian for all numeric types.
*   **Variable Length Fields:** MUST be prefixed with a `uint16` length descriptor.
*   **Fixed Fields:** MUST appear first in the `Digest()`.

### 1.2 Resource Constraints
*   **Maximum Event Size:** 1 MB (including payload). 
    *   *Rationale:* Prevents memory-exhaustion DoS and keeps gossip propagation latency within `Δ`.
*   **Payload Type:** JSON for flexibility, but hashed deterministically as raw bytes.

### 1.3 Versioning & Evolution
*   **Protocol Versioning:** Semantic versioning (e.g., `1.0.1`).
*   **Event Versioning:** Monotonic `uint16`. 
*   **Mixed Clusters:** A node MUST reject any event version higher than its own implementation. Replay nodes MUST maintain historical `migrateEvent` logic to handle legacy versions.
*   **Algorithm Migration:** SHA-256 is the currently mandated algorithm. Future migrations will require a network-wide Governance QC.

## 2. Tick Semantics
*   **Definition:** Ticks represent deterministic simulation steps (discrete time).
*   **Increment Rule:** Ticks increment **per consensus round**.
    *   *Rationale:* This ensures all events within a single round share the same temporal reality, simplifying parallel physics evaluation.
*   **Monotonicity:** `event.Tick` MUST be greater than or equal to `ws.Tick`.
*   **Maximum Jump:** `ws.Tick + 100`. Events leaping beyond this threshold are rejected.

## 3. Snapshot Specification

### 3.1 Snapshot Header
Every snapshot file MUST begin with a canonical header for integrity verification.
```go
type SnapshotHeader struct {
	Index            uint64 `json:"index"`             // Last finalized index included
	Tick             uint64 `json:"tick"`              // Current simulation tick
	StateHash        Hash   `json:"state_hash"`        // Deterministic projection hash
	LedgerHash       Hash   `json:"ledger_hash"`       // Hash of the last event
	ProtocolVersion  string `json:"protocol_version"`  // Engine version at capture
	ValidatorSetHash Hash   `json:"validator_set_hash"` // Merkle root of authorized keys
}
```

### 3.2 Restoration Policy
1.  Verify `SnapshotHeader.StateHash` against the locally reconstructed state.
2.  Verify `ProtocolVersion` compatibility.
3.  Load `LastSeenSequences` (Crucial: prevents Crash Rollback Attacks).
4.  Begin replaying the "Ledger Tail" (events with `Index > SnapshotHeader.Index`).

---
*Authorized by Phoenix Sovereign Governance*

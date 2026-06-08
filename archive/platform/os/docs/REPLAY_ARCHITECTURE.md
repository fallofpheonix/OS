---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Replay Architecture

Run A

↓

Input Hash

↓

Replay Session

↓

Output Hash

↓

Cross Run Verify

↓

Truth Evidence

---

## Storage

HOT

recent execution

WARM

history

COLD

archive

---

## Authority

logical clock

ordering

hash chain

verification

---

## Technical Implementation (Phase F1)

### ReplayRunner
The core orchestration engine that verifies execution traces against the `TruthLedger`. It performs three levels of verification:
1. **Sequence Verification:** Ensures `LogicalTick` continuity.
2. **Cryptographic Integrity:** Re-computes SHA-256 hashes for every entry using the `Ledger`'s canonical logic.
3. **Lineage Verification:** Validates that each entry's `ParentIDs` correctly link to the previous step's hash.

### ReplayProvider Interface
Decouples the engine from trace storage, allowing for mocks, file-based logs, or live stream ingestion.
```go
type ReplayProvider interface {
    LoadTrace(traceID string) ([]ledger.LedgerEntry, error)
}
```

### Forensic Reporting
`ReplayIdentity` generates detailed `ReplayReports` identifying the exact `DivergenceTick` and `DivergenceType` (HASH_MISMATCH, SEQUENCE_GAP, etc.) in case of state deviation.

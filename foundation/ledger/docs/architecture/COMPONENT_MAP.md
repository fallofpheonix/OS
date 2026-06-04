---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger — Component Map

> Last verified: 2026-06-04

The `foundation/ledger` package provides the immutable forensic history of state transitions within Phoenix OS. It implements sequential chains, Merkle DAG structures, and multi-sensor quorum calculations.

## Component Breakdown

```
foundation/ledger/
├── go.mod                     # Module configuration
├── chain.go                   # Linear hash chain implementation
├── reality.go                 # Reality Confidence Theorem consensus
├── src/                       # Core ledger v2 engine
│   ├── ledger.go              # Evidence Merkle DAG, hashing, verification
│   └── (tests for allocation, pruning, rollback)
├── genesis.go                 # Genesis block bootstrapping
├── persistor.go               # Disk serialization logic
├── snapshot.go                # Ledger state checkpointing
└── cmd/                       # Standalone ledger CLI
    └── main.go
```

### Component Details

1. **`Chain` (`chain.go`)**:
   - Manages a sequence of `Event` logs. Enforces sequence continuity ($S_{t} = S_{t-1} + 1$) and Parent-Hash matching.

2. **`Reality` (`reality.go`)**:
   - Implements the Reality Confidence Theorem. Aggregates multi-sensor observations, resolves conflicts via weighted voting, and outputs a confidence score.

3. **`Ledger` (`src/ledger.go`)**:
   - Core Ledger engine. Manages `LedgerEntry` structs inside an Evidence Merkle DAG.
   - Enforces memory resource bounding through a `ResourceAllocator`.
   - Supports checkpointing, historic pruning, and transaction rollback.

---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger — Control Flow

> Last verified: 2026-06-04

Enforces transaction safety through write locks and rollback states.

## Rollback Control Flow

```mermaid
stateDiagram-v2
    [*] --> LockLedger: Initiate Rollback
    LockLedger --> FindCheckpoint: Search Entries by Hash
    FindCheckpoint --> IdentifyOutdatedEntries: Filter ticks > checkpoint.LogicalTick
    IdentifyOutdatedEntries --> DeallocateMemory: Release bytes to Allocator
    DeallocateMemory --> DeleteEntries: Remove from Entries map
    DeleteEntries --> ResetHeads: Set Heads to Checkpoint Hash
    ResetHeads --> UnlockLedger: Rollback Success
    UnlockLedger --> [*]
```

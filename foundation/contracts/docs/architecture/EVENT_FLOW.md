---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Event Flow

> Last verified: 2026-06-04

This document defines the lifecycle flow of events through the contracts interfaces.

## Event Processing Pipeline

```mermaid
sequenceDiagram
    participant Sender
    participant Validator
    participant Ledger
    participant Replay

    Sender->>Validator: Emit raw event JSON
    Validator->>Validator: Unmarshal via Serializer
    Validator->>Validator: Verify Event signature & logical time
    Validator->>Ledger: Append via ILedger.AddEntryV2
    Ledger->>Ledger: Calculate parent hashes & validation hashes
    Ledger->>Replay: Feed EventEnvelope to ReplayEngine
    Replay->>Replay: Replay state modification
```

## Guarantees

1. **Causal Ordering**: Every event envelope has a parent event hash. Replay engines process events chronologically using logical timestamps.
2. **Non-repudiation**: Events must contain signatures that match the authority ID.

---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Event Flow

> Last verified: 2026-06-04

Coordinates event propagation through consensus nodes and replay executors.

```mermaid
sequenceDiagram
    participant Proposer
    participant Consensus
    participant Ledger
    participant Replay

    Proposer->>Consensus: Propose transactions block
    Consensus->>Consensus: BFT Quorum Voting
    Consensus->>Ledger: Write finalized block events
    Ledger->>Replay: Broadcast event IDs
    Replay->>Replay: Update active memory states
```

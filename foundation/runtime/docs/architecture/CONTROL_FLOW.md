---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Control Flow

> Last verified: 2026-06-04

Describes runtime transaction validation, consensus iterations, and replay loops.

## Replay Loop Control Flow

```mermaid
stateDiagram-v2
    [*] --> Idle: Wait for Replay Trigger
    Idle --> FetchEvents: Load event sequence
    FetchEvents --> ReconstructState: Pass to Reconstructor
    ReconstructState --> AssertInvariants: Run constitution engine
    AssertInvariants --> UpdateStateSnapshot: Save checkpoint
    UpdateStateSnapshot --> Idle
    AssertInvariants --> HaltOnInconsistentState: Violations Found!
    HaltOnInconsistentState --> [*]
```

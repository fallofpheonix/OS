---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Control Flow

> Last verified: 2026-06-04

Defines the escalation states of system containment policies.

## Escalation Control Logic

```mermaid
stateDiagram-v2
    [*] --> LevelNone: Invariant Violations = 0
    LevelNone --> LevelMonitor: Violation Limit Exceeded
    LevelMonitor --> LevelSandbox: Trust Score drops below 0.8
    LevelSandbox --> LevelIsolate: SDI exceeds Critical Threshold
    LevelIsolate --> LevelQuench: Trust Score = 0
    LevelQuench --> Shutdown: Trigger global killswitch
    Shutdown --> [*]
```

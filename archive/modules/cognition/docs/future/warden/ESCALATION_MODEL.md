---
Status: Partial
Implementation: 55%
Confidence: Ready
---
# Warden Guard Rails — Escalation Rules

Rules for mounting the containment ladder.

## Rules
- Single failure of an invariant: Elevate to `LevelMonitor`.
- Network egress violation inside sandbox: Elevate to `LevelIsolate`.
- Memory breach or tampering detection: Immediate transition to `LevelQuench`.

---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Events — Control Flow

> Last verified: 2026-06-04

Event structures are read-only after creation and are thread-safe. Control flow is simple and deterministic.

## Control Flow
1. **Creation**:
   - Build concrete `Event` instance.
   - Attach signatures and parent IDs.
2. **Usage**:
   - Passed down to validation or ledger systems.
   - State transition updates are strictly read-only.

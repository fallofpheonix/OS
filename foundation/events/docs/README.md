---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Events

> **Layer**: Foundation | **Maturity**: Mature | **Owner**: Architecture Team

## Purpose

Event schema, lifecycle management, and serialization. Implements the canonical event types defined by contracts with proper versioning and evolution support.

## Quick Start

```bash
# Build
cd foundation/events && go build ./...

# Test
cd foundation/events && go test ./...
```

## Documentation

| Document | Description |
|----------|-------------|
| [SYSTEM_MAP.md](./SYSTEM_MAP.md) | Components and their relationships |
| [CURRENT_STATE.md](./CURRENT_STATE.md) | What is implemented vs missing |
| [TARGET_STATE.md](./TARGET_STATE.md) | Future goals and requirements |
| [MIGRATION_PATH.md](./MIGRATION_PATH.md) | Steps from current to target state |
| [EXTRACTION_READINESS.md](./EXTRACTION_READINESS.md) | Readiness for independent release |

## Module

```
github.com/fallofpheonix/phoenix/foundation/events
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

---\nStatus: Partial\nImplementation: 70%\nConfidence: Tested\n---\n# Observability

> **Layer**: Foundation | **Maturity**: Partial | **Owner**: Observability Team

## Purpose

System-wide observability infrastructure including metrics collection, tracing, and health reporting for Phoenix subsystems.

## Quick Start

```bash
# Build
cd foundation/observability && go build ./...

# Test
cd foundation/observability && go test ./...
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
github.com/fallofpheonix/phoenix/foundation/observability
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

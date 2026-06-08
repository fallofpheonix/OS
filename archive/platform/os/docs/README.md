---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Platform OS

> **Layer**: Platform | **Maturity**: Partial | **Owner**: Platform Team

## Purpose

PhoenixOS platform layer with CLI, API server, boot system, event bus, game server integration, guard, monitor, recovery, containment, and telemetry. The primary user-facing integration surface with 201 Go source files.

## Quick Start

```bash
# Build
cd platform/os && go build ./...

# Test
cd platform/os && go test ./...
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
github.com/fallofpheonix/phoenix/platform/os
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

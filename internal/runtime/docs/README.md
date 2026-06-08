---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime

> **Layer**: Foundation | **Maturity**: Mature | **Owner**: Core Runtime Team

## Purpose

Core execution engine including the replay engine, event bus, arbiter, AI agents, game server, scheduler, containment, monitoring, and telemetry. The largest subproject in Phoenix with 36 subdirectories.

## Quick Start

```bash
# Build
cd foundation/runtime && go build ./...

# Test
cd foundation/runtime && go test ./...
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
github.com/fallofpheonix/phoenix/foundation/runtime
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

---\nStatus: Partial\nImplementation: 50%\nConfidence: Tested\n---\n# Platform UI

> **Layer**: Platform | **Maturity**: Partial | **Owner**: Platform Team

## Purpose

PhoenixOS graphical shell built with React, TypeScript, and Vite. Provides the desktop environment UI including window management, application integration, and visual theming.

## Quick Start

```bash
# Build
cd platform/ui && go build ./...

# Test
cd platform/ui && go test ./...
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
github.com/fallofpheonix/phoenix/platform/ui
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

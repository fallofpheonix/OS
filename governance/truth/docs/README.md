---\nStatus: Planned\nImplementation: 15%\nConfidence: Conceptual\n---\n# Truth

> **Layer**: Governance | **Maturity**: Early | **Owner**: Verification Team

## Purpose

Truth verification engine that validates system state assertions against evidence. Implements the truth pipeline for distinguishing verified facts from claims.

## Quick Start

```bash
# Build
cd governance/truth && go build ./...

# Test
cd governance/truth && go test ./...
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
github.com/fallofpheonix/phoenix/governance/truth
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

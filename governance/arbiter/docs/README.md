---\nStatus: Planned\nImplementation: 15%\nConfidence: Conceptual\n---\n# Arbiter

> **Layer**: Governance | **Maturity**: Early | **Owner**: Governance Team

## Purpose

Policy arbitration and conflict resolution. Evaluates competing claims and enforces governance decisions based on the Phoenix Constitution.

## Quick Start

```bash
# Build
cd governance/arbiter && go build ./...

# Test
cd governance/arbiter && go test ./...
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
github.com/fallofpheonix/phoenix/governance/arbiter
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

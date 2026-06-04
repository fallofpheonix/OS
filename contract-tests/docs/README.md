---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contract Tests

> **Layer**: Assurance | **Maturity**: Mature | **Owner**: Architecture Team

## Purpose

Contract compatibility test suites that verify all implementations conform to their declared contracts. Tests events, replay, and security contract compliance.

## Quick Start

```bash
# Build
cd contract-tests && go build ./...

# Test
cd contract-tests && go test ./...
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
github.com/fallofpheonix/phoenix/contract-tests
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

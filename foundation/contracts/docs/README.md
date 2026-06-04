---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts

> **Layer**: Foundation | **Maturity**: Mature | **Owner**: Architecture Team

## Purpose

Defines all public interfaces for the Phoenix system. Only contract packages define public types — everything else implements them. Contracts own Event, EventEnvelope, ReplayEngine, Snapshot, Actuator, and Containment types.

## Quick Start

```bash
# Build
cd foundation/contracts && go build ./...

# Test
cd foundation/contracts && go test ./...
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
github.com/fallofpheonix/phoenix/foundation/contracts
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

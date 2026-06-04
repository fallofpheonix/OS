---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security

> **Layer**: Assurance | **Maturity**: Mature | **Owner**: Security Team

## Purpose

Security enforcement, threat detection, containment actuators, and cryptographic verification. Implements the security contracts with real enforcement mechanisms including certificate validation and boundary enforcement.

## Quick Start

```bash
# Build
cd assurance/security && go build ./...

# Test
cd assurance/security && go test ./...
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
github.com/fallofpheonix/phoenix/assurance/security
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

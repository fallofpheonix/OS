---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger

> **Layer**: Foundation | **Maturity**: Mature | **Owner**: Ledger Team

## Purpose

Append-only forensic record implementing the Evidence Merkle DAG. Provides tamper-proof logging, causal traceability, hash chain integrity, snapshots, pruning, and deterministic replay. Split into root package (Chain, Event types) and src/ subpackage (Ledger, AddEntry).

## Quick Start

```bash
# Build
cd foundation/ledger && go build ./...

# Test
cd foundation/ledger && go test ./...
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
github.com/fallofpheonix/phoenix/foundation/ledger
```

---
*Part of the [Phoenix Master Architecture](../../docs/MASTER_SYSTEM_MAP.md)*

---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Master Dependency Map

> Machine-readable dependency matrix derived from `go.mod` and `go.work` files.
> Last verified: 2026-06-04

## Dependency Matrix

Rows depend on columns. `●` = direct dependency. `○` = transitive dependency.

| Module ↓ depends on → | contracts | events | runtime | ledger | distributed | observability | security | validation | truth | mind | platform/os |
|------------------------|-----------|--------|---------|--------|-------------|---------------|----------|------------|-------|------|-------------|
| **contracts** | — | | | | | | | | | | |
| **events** | ● | — | | | | | | | | | |
| **runtime** | ● | ● | — | ● | | | | | | | |
| **runtime/kernel** | ● | | ● | | | | | | | | |
| **ledger** | | | | — | | | | | | | |
| **distributed** | | | ● | | — | | | | | | |
| **observability** | | | ● | | | — | | | | | |
| **security** | ● | ● | ● | ● | | | — | | | | |
| **validation** | ● | ● | ● | ● | | ○ | ● | — | | ○ | ○ |
| **truth** | | | ● | | | | | | — | | |
| **arbiter** | ● | | | | | | | | | | |
| **cognition** | | | ● | ● | | | | | | | |
| **cognition/mind** | | | ● | ● | ○ | | ● | | | — | |
| **platform/os** | ● | ● | ● | ● | ○ | ○ | ● | | ○ | ● | — |
| **platform/ui** | | | | | | | | | | | |
| **contract-tests** | ● | ● | ● | | | | ● | | | | |

## Layer Violation Checks

The following import patterns are **forbidden**:

```
foundation/* → assurance/*     FORBIDDEN
foundation/* → governance/*    FORBIDDEN
foundation/* → cognition/*     FORBIDDEN
foundation/* → platform/*      FORBIDDEN
contracts/*  → runtime/*       FORBIDDEN
contracts/*  → ledger/*        FORBIDDEN
assurance/*  → cognition/*     FORBIDDEN (except validation → mind via adapters)
```

## Circular Dependency Report

| Status | Description |
|--------|-------------|
| ✅ CLEAN | No circular dependencies detected in `go.work` modules |

## Replace Directives Summary

All inter-module dependencies use `replace` directives pointing to local paths:

```
github.com/fallofpheonix/phoenix/foundation/contracts => ../../foundation/contracts
github.com/fallofpheonix/phoenix/foundation/events    => ../../foundation/events
github.com/fallofpheonix/phoenix/foundation/runtime   => ../../foundation/runtime
github.com/fallofpheonix/phoenix/foundation/ledger    => ../../foundation/ledger
...
```

This pattern enables monorepo development with independent module versioning.

---
*Enforced by `tools/check_boundaries.sh`. See [MASTER_SYSTEM_MAP.md](./MASTER_SYSTEM_MAP.md) for visual layout.*

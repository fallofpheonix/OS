---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Dependency Map

> Last verified: 2026-06-04

The `foundation/contracts` package resides at the absolute bottom of the Phoenix OS dependency stack. It must remain fully isolated to prevent circular dependencies.

## Dependency Rules

```
┌──────────────────────────────────────────────┐
│             foundation/contracts             │
├──────────────────────────────────────────────┤
│  ▲ No imports from other Phoenix packages   │
│  ▼ Imported by: events, runtime, ledger, etc.│
└──────────────────────────────────────────────┘
```

## Permitted Standard Library Imports
- `context`
- `time`
- `encoding/json`
- `errors`
- `fmt`

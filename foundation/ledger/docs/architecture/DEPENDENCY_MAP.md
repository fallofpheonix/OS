---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Ledger — Dependency Map

> Last verified: 2026-06-04

The ledger relies on contract interfaces and core event structures.

## Dependencies

```
┌──────────────────────────────────────────────┐
│              foundation/ledger               │
├──────────────────────────────────────────────┤
│  ▲ Imports: contracts, events                │
│  ▼ Imported by: runtime, validation, etc.    │
└──────────────────────────────────────────────┘
```

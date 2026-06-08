---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Dependency Map

> Last verified: 2026-06-04

`foundation/runtime` is a key orchestrator that depends on contracts and events but is isolated from high-level layers (assurance, platform UI).

## Dependency Matrix

```
┌──────────────────────────────────────────────┐
│              foundation/runtime              │
├──────────────────────────────────────────────┤
│  ▲ Imports: contracts, events, ledger        │
│  ▼ Imported by: validation, security         │
└──────────────────────────────────────────────┘
```

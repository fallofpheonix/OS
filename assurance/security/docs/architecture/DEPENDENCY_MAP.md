---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Dependency Map

> Last verified: 2026-06-04

The security module interacts with Go system boundaries.

## Dependency Matrix

```
┌──────────────────────────────────────────────┐
│              assurance/security              │
├──────────────────────────────────────────────┤
│  ▲ Imports: contracts, events, ledger,       │
│            runtime                           │
│  ▼ Imported by: validation                   │
└──────────────────────────────────────────────┘
```

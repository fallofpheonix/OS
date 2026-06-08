---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# ADR-001: Canonical Contract Source

## Status
Accepted

## Date
2026-05-31

## Context
Multiple packages need to share types and contracts across repository boundaries. Without a canonical source, type duplication and inconsistency become problematic.

## Decision
PhoenixCore is the single source of truth for all cross-boundary types. All other packages MUST import contract types from PhoenixCore. No other package may export cross-boundary types.

## Consequences

### Easier
- Consistent type definitions across the ecosystem
- Single place to update contracts
- Clear ownership and responsibility

### More Difficult
- All packages depend on PhoenixCore
- Changes to PhoenixCore require coordination
- Circular dependencies must be avoided

## Alternatives Considered
1. **Distributed contracts** — Rejected: leads to inconsistency
2. **Code generation** — Rejected: adds complexity
3. **Interface-only contracts** — Rejected: loses type information

## References
- [PhoenixCore README](../../PhoenixCore/README.md)
- [INVARIANTS.md](../../PhoenixCore/INVARIANTS.md)

---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# ADR-001: Documentation Standards

## Status
Accepted

## Date
2026-05-31

## Context
Documentation is scattered across multiple locations with inconsistent formats, naming, and cross-references. This makes it difficult for developers and AI agents to find and use documentation.

## Decision
Standardize documentation with:
- Root-level README.md, CONTEXT.md, WORKING_MODEL.md, DEVELOPMENT_GUIDE.md, AGENTS.md
- Directory-level READMEs for each of the 4 directories
- Package-level CLAUDE.md for each package
- Standardized ADR format in docs/adr/
- Cross-references between related documents

## Consequences

### Easier
- Consistent documentation structure
- Clear ownership and responsibility
- Easy navigation for developers and AI agents

### More Difficult
- More documentation to maintain
- Requires discipline to keep in sync
- Cross-references must be validated

## Alternatives Considered
1. **No documentation standards** — Rejected: leads to chaos
2. **Single README per repo** — Rejected: too little detail
3. **Wiki-based docs** — Rejected: not version-controlled

## References
- [README.md](../../README.md)
- [CONTEXT.md](../../CONTEXT.md)
- [DEVELOPMENT_GUIDE.md](../../DEVELOPMENT_GUIDE.md)

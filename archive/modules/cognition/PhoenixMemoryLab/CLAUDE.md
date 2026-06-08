# PhoenixMemoryLab — Episodic Memory Layer

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
# Open Obsidian vault at PhoenixMemoryLab/
```

## Architecture
PhoenixMemoryLab is an Obsidian-based knowledge vault for episodic memory, semantic storage, and retrieval. It provides historical context for AI reasoning.

## Key Components
- **00_DASHBOARD/** — Project dashboard
- **03_CORE_KNOWLEDGE/** — Core knowledge base
- **04_ENGINEERING/** — Engineering patterns and 23 ADRs
- **05_PROJECTS/** — Project documentation
- **06_FAILURE_LIBRARY/** — Failure case studies
- **10_META/** — Meta-governance rules

## Invariants
- All knowledge must be versioned
- Retrieval must be tracked
- Memory lineage must be explicit
- Access must be controlled

## Research & Reference

Third-party reference projects have been moved to `Phoenix.Terminus/PhoenixExternal/research/`. These are NOT part of PhoenixOS core.

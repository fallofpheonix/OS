# Ecosystem Skeleton Plan

This document outlines the transition from the current local file system structure to a GitHub-backed minimal skeleton.

## Target Structure
```
engineering/
├── active/           # Protected and actively developed repositories
├── archive/          # Local copies of inactive/legacy projects
├── modules/          # Installable components and forks
├── github_cache/     # Temporary workspace for on-demand clones
├── runtime/          # Core runtime environments and configurations
└── ecosystem_os/     # Governance and control plane logic
```

## Mapping Strategy

| Category | Source Path | Target Path | Action |
|----------|-------------|-------------|--------|
| **PROTECTED** | `workspace/active/*`, `brain/` | `active/` | Keep local |
| **CORE_LOCAL** | `workspace/active/*` | `active/` | Keep local |
| **ARCHIVE_ONLY** | `archive/*` | `archive/` | Consolidate |
| **REMOTE_PRIMARY** | Various | GitHub | Offload (Clone on demand) |
| **INSTALLABLE** | `modules/*`, `forks/*` | `modules/` | Register in `module_registry.yaml` |

## Migration Steps

1. **Verify Sync:** Ensure all `REMOTE_PRIMARY` candidates are fully pushed to GitHub (See `github_sync_validation.md`).
2. **Consolidate Archive:** Move all `ARCHIVE_ONLY` repos into the top-level `archive/` folder.
3. **Establish Modules:** Move `modules/` and `forks/` into the target `modules/` structure.
4. **Active Workspace:** Relocate `workspace/active/` contents to top-level `active/`.
5. **Clean Up:** Remove redundant folders (`apps/`, `services/`, `sdk/`, etc.) once contents are categorized or archived.

## Protection Rules
The following directories are **EXEMPT** from removal or archiving:
- `fallofpheonix`
- `astraeus-core`
- `brain`
- `forge-agent`
- `ecosystem_os`
- `runtime`

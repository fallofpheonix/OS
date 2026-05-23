# Cleanup Plan

## 1. Directory Consolidation
- **`archives/` and `archive/`**: Duplicate root-level folders.
  - Action: Move `archives/workspace_old` into `archive/` and delete `archives/`.
- **`shared/` vs `modules/` vs `infrastructure/shared-libraries/`**: Scattered shared code locations.
  - Action: Flag for review in structure report. Move `shared/` contents to `modules/` if they are module-level, or to `infrastructure/shared-libraries/`.
- **`tools/` vs `scripts/`**: Duplicate root-level intent.
  - Action: Move `scripts/` into `infrastructure/scripts/` to unify scripts, and move `tools/` into `infrastructure/tooling/`.

## 2. Root Clutter Mitigation
- **Loose Markdown Files**: The root `engineering/` contains >10 floating `.md` architecture/research documents (`AI_SYSTEM_AUDIT.md`, `MULTI_MODEL_ARCHITECTURE.md`, etc.).
  - Action: Move these loose files to `docs/architecture/` or `docs/research/` to declutter the root workspace.

## 3. Hidden File Scrubbing
- **`.DS_Store`**: Misplaced macOS artifacts found in `workspace/` and `archives/`.
  - Action: Safely remove `.DS_Store` files across the repository.

## 4. Execution
The above safe actions will be performed via shell commands.

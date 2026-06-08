# Operating Rules

## Layer Invariants
- Brain: concepts, failures, decisions, patterns, architecture, research. No executable runtime artifacts.
- Workspace: project code, Docker, configs, `.env` files. No permanent knowledge storage.
- Infrastructure: shared Docker images, scripts, templates, reusable libraries, datasets.
- Environments: isolated Python, Node, Cargo, and other runtime state per project.

## Vault Constraints
- Maximum folder depth: 4.
- Prefer links over deeper nesting.
- Note names use Pascal Case.
- No `node_modules`, `venv`, `build`, binaries, or `dist` in the vault.
- Split notes before they exceed 500 lines.

## Project Parallel-Ready Checklist
- [ ] Unique API port assigned.
- [ ] Unique DB host port assigned.
- [ ] Unique database name.
- [ ] Own `.env` file.
- [ ] Own Docker network or container names prefixed with project name.
- [ ] Own venv / `node_modules` / Cargo target dir.
- [ ] No imports from another project's source code; API calls only.
- [ ] Obsidian note exists under `05_PROJECTS/ACTIVE/project-name/`.

## Friday Review
- [ ] Process `01_CAPTURE/inbox.md`.
- [ ] Create at least one failure note.
- [ ] Create one ADR if an architecture decision was made.
- [ ] Update `10_META/dashboards/Project_Dashboard.md`.
- [ ] Move mastered active-learning material into `03_CORE_KNOWLEDGE/`.
- [ ] Update weak areas in `10_META/dashboards/Home.md`.


# Repository Sanitization Report — Phase 0

Generated: 2026-05-15

Scope
-----
Complete repository sanitation audit focused on identifying dead systems, abandoned experiments, duplicate logic, invalid abstractions, misplaced files, and other repository hygiene issues. This report provides evidence, recommended safe actions, and a prioritized cleanup list.

Method
------
- Grepped repository for markers: `deprecated`, `experimental`, `TODO`, `FIXME`, `mock`, `stub`, `archive`, `legacy`, `prototype`, `raise NotImplementedError`.
- Reviewed key runtime modules: `workspace/active/astraeus-core/transactions`, `orchestrator`, `repo_indexer`, `repair`.
- Reviewed control-plane contracts and schemas under `control-plane/contracts`.
- Noted environment artifacts (`.venv`, `environments/ai-system/venv`) and backup artifacts under `backups/`.

Summary Findings
----------------
- Many documentation files contain TODO placeholders and `ARCHIVED` markers (governance and docs areas).
- Multiple experimental/legacy artifacts exist in `environments/` and `backups/` including full `venv` trees and `.smart-env` dump files.
- Evidence of mock/test artifacts in code and docs (e.g., mock critic in `GOALS.md`, hardcoded test strings referenced in `src/openhuman/...`).
- Event emission roadmap and TODOs are present but many emitters are unimplemented.
- No central authoritative event store; events are file-based and per-run logs.

Detailed Items (required outputs)
--------------------------------
1. Dead Files
   - Candidates: committed virtualenv trees under `/.venv` and `environments/ai-system/venv` (examples found). These should be removed from VCS and archived.
   - Many `Pasted text (*.txt)` and `backups/*/.smart-env/*` large dump files live under `backups/` and appear to be knowledge exports rather than source. Archive these to `/engineering/archive/legacy`.

2. Dead Modules
   - No obvious core-module marked `deprecated` in `workspace/active/astraeus-core`, but several `experimental` modules referenced in `control-plane/sync-engine/repo-registry.yaml` and `control-plane/schemas/repo.schema.json` flagged as `experimental`/`deprecated`.
   - Action: run import analysis to list modules with zero inbound imports (Phase 0.3).

3. Unused Dependencies
   - `environments/ai-system/requirements.txt` contains `Deprecated==1.3.1` (likely placeholder). Suggest pinning actual runtime deps in `requirements.txt` at repo root and removing environment-specific copies.
   - Action: build dependency graph and run static import checker to find unused packages.

4. Duplicate Logic
   - Duplication of guides and roadmaps across `governance/evolution/` and `docs/evolution/` (e.g., `ROADMAP.md` duplicates). Consolidate into `governance/evolution/` canonical files and remove doc duplicates.

5. Circular Dependencies
   - Not yet exhaustively scanned. Many modules import across `runtime`, `orchestrator`, and `repo_indexer` — run static graph analysis next to detect cycles.

6. Misplaced Files
   - `.venv` and `environments/*/venv` are misplaced in repo. Move to `/engineering/archive/legacy` and add `.gitignore` entries.
   - Large backup artifacts in `backups/` should be moved to `/engineering/archive/legacy` or external storage.

7. Invalid Naming
   - Mixed naming conventions: `REPAIR_FLOW.md`, `REPAIR_ARCHITECTURE.md`, `repair_planner.py` — normalize to `snake_case` for code and `Title Case` for docs with consistent paths under `governance/architecture/repair.md`.

8. Broken Architecture Boundaries
   - `orchestrator/engine.py` holds planner, journal, rollback, mutation sandbox, and risk engine — this god-object breaks boundaries. Plan: split responsibilities into `planner/`, `executor/`, `safety/` modules.

9. Technical Debt
   - Many docs marked TODO; ADRs are sparse. The repo contains mock fallbacks that allow tests to pass while bypassing real model execution.

10. Unsafe Runtime Patterns
   - Committed `.venv` and local manpages under `environments` and `.venv` risk secret leakage and make environments non-deterministic.
   - Model fallback to deterministic mocks (documented) is an unsafe pattern for Phase A+ runs.

11. Mock/Fake Systems
   - `GOALS.md` and tests indicate use of a mock critic; search shows `mock` markers. Tag all mock paths explicitly and ensure CI rejects PRs that introduce mock fallbacks into production runs.

12. Unused Configurations
   - `environments/*` contains example and sample configs; move sample configs to `docs/operations/examples` and remove environment-specific copies.

13. Drifted Documentation
   - Duplicate and out-of-sync roadmap/ADR/implementation notes across `docs/` and `governance/`. Consolidation required.

14. Orphaned Components
   - Control-plane `sync-engine` registry YAML references example repos and `experimental` maturity; archive or mark as incubating.

15. Recommended Deletions (candidates — verify before action)
   - Remove committed virtualenv trees (move to archive first): `/.venv`, `environments/ai-system/venv`.
   - Delete large `.smart-env` knowledge dumps from `backups` if they are duplicated elsewhere — archive first.

16. Recommended Moves
   - Move `environments/*/venv` → `/engineering/archive/legacy/`.
   - Move `backups/*/.smart-env/*` → `/engineering/archive/legacy/` (or external storage).
   - Consolidate duplicate roadmap/docs under `governance/evolution/`.

17. Recommended Merges
   - Merge duplicate evolutions and roadmap documents into `governance/evolution/ROADMAP.md` and remove redundant copies in `docs/evolution/`.
   - Merge event emitter TODOs into `control-plane/contracts/events/IMPLEMENTATION-ROADMAP.md` canonical TODO list.

18. Recommended Renames
   - Normalize naming: code modules `snake_case`, docs `UPPER_SNAKE.md` -> `Title Case` with hyphenated filenames where appropriate.

19. Architecture Cleanup Priority
   - P0: Remove/Archive committed virtualenvs; add `.gitignore` updates; create canonical `requirements.txt` at repo root.
   - P0: Consolidate event emission roadmap and implement minimal supervisor lifecycle emitters (control-plane TODOs 1-2).
   - P1: Split `orchestrator` responsibilities into separate modules; implement module boundaries and update imports.
   - P1: Run static import analysis and circular dependency fixer; produce `DEPENDENCY_GRAPH.md`.
   - P2: Convert governance docs into machine-readable invariants and wire CI checks.

20. Risk Classification
   - Critical: Committed virtualenvs and backups with potential secrets; mock fallbacks masking failures; incomplete event coverage.
   - High: Overloaded orchestrator; topology drift risk; missing CI gates.
   - Medium: Duplicate docs; experimental registry entries.
   - Low: Naming inconsistencies and documentation TODOs.

Safe Cleanup Plan (Phase 0.1)
---------------------------
1. Create archive folders (done): `/engineering/archive/deprecated`, `/engineering/archive/experiments`, `/engineering/archive/legacy`.
2. For each candidate deletion, verify references (imports, runtime paths, roadmap dependencies). If any inbound reference exists, do NOT delete; instead move to `/engineering/archive/deprecated` and track for future removal.
3. Remove committed virtualenvs from git history if confirmed not required, but only after archiving and ensuring CI can recreate environment from `requirements.txt`.
4. Replace mock fallbacks with explicit feature flags gated behind `DETERMINISTIC_MOCKS=true` and ensure default in CI is false.

Next Steps (Phase 0.2+0.3)
-------------------------
- Run a full static import analysis to produce `DEPENDENCY_GRAPH.md` and list modules with zero inbound imports.
- Produce `MODULE_BOUNDARIES.md` and `IMPORT_RULES.md` for enforcement.
- Implement CI `env-validate` job to ensure reproducible environments.
- Begin orchestrator split refactor with tests and small incremental PRs.

Evidence & Examples
-------------------
- `GOALS.md` references mock critic and example bad keys.
- `control-plane/contracts/events/IMPLEMENTATION-ROADMAP.md` lists 10 emitter TODOs.
- Search results showed many `TODO` placeholders in `governance/` and `docs/` areas.

Verification Checklist (before any deletion)
------------------------------------------
- Confirm zero active imports using static import analysis.
- Confirm no runtime reference in orchestrator/planner/repair paths.
- Confirm not referenced in any roadmap migration or planned feature.
- Archive to `/engineering/archive/*` rather than immediate deletion if uncertain.

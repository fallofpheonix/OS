---
project-id: healingstone
repo: https://github.com/fallofpheonix/healingstone
status: ACTIVE
audit-verdict: "PROMOTE AS PRIMARY PROJECT"
audit-scores: ED8/AQ8/SC6/RU7/MT8/OR8/PV9/CP9/RU7/EC9/LV8/RS8/OS8
language: Python
started: 2026-03
last-commit: 2026-03-31
tags: [project, existing, 3d-reconstruction, computer-vision]
---

# Healingstone

## What It Is
A 2D/3D fragment reconstruction and alignment pipeline that matches broken fragments using geometric descriptors, Siamese network similarity, and cosine-similarity fallback scoring. Given a set of mesh fragments (.ply) or 2D image fragments, it produces reconstructed assemblies with alignment metrics, similarity matrices, and reproducible artifact contracts. Built for computational archaeology and heritage reconstruction use cases.

## Current State
- **What is working:** mypy strict mode, ruff linting, test coverage reporting, CLI smoke test, artifact output contract (metrics.json + reconstruction.ply + alignment_metrics.json + similarity_matrix.png). This has the most mature quality gate stack of all 7 repos.
- **What is broken or incomplete:** `backup_untracked/` directory committed to repo (must be purged from git history), duplicated schema/core modules between 2D and 3D pipelines, no published sample datasets, no algorithm accuracy benchmarks.
- **Audit-identified defects:**
  - backup_untracked tree in repo — binary/generated artifacts polluting git history
  - Duplicated schema/core modules — 2D and 3D interfaces share logic but have separate copies
- **Production readiness:** PARTIAL — closest to READY of all 7 repos, but artifact violations and missing benchmarks block promotion.

## Architecture

**Current:**
```
healingstone/
├── core/           ← shared descriptor extraction and matching logic
├── pipeline_2d/    ← 2D image fragment pipeline
├── pipeline_3d/    ← 3D mesh fragment pipeline (Open3D)
├── models/         ← Siamese network training/inference
├── fallback/       ← cosine similarity deterministic fallback
├── output/         ← artifact contract: metrics.json, reconstruction.ply, etc.
├── tests/          ← pytest suite with smoke test
├── backup_untracked/ ← VIOLATION: must be removed
└── configs/
```

**Target (from audit):**
```
healingstone/
├── src/healingstone/
│   ├── core/           ← unified descriptor/matching interface
│   ├── pipeline/       ← single pipeline with 2D/3D mode switch
│   ├── models/         ← Siamese network
│   ├── fallback/       ← deterministic scoring
│   └── contracts/      ← output schema validation
├── tests/
├── datasets/           ← sample data via DVC/releases (NOT in repo)
├── benchmarks/         ← accuracy metrics on synthetic + real fragments
├── docs/adr/
└── pyproject.toml
```

**Gap:** The 2D and 3D pipelines need to be unified under a single interface with a mode parameter. The backup_untracked tree must be purged. Sample datasets need to be published externally.

## Tech Stack
- **Language:** Python 3.10+
- **3D Processing:** Open3D
- **ML:** PyTorch (Siamese Networks)
- **Matching:** Cosine similarity (fallback), learned descriptors (primary)
- **Quality:** mypy (strict), ruff, pytest, coverage
- **Output:** JSON (metrics), PLY (3D reconstruction), PNG (similarity heatmap)

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-pipeline]] | Not yet imported | Natural fit for fragment pipeline control-plane |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-009-healingstone-2d-3d-pipeline-split]] — why 2D and 3D were built as separate pipelines
- [[04_ENGINEERING/decision-logs/ADR-010-healingstone-artifact-contract]] — why every run produces a fixed set of output artifacts
- [[04_ENGINEERING/decision-logs/ADR-011-healingstone-mypy-ruff-stack]] — why mypy strict + ruff was chosen as the quality gate

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-03-healingstone-backup-committed]] — backup_untracked directory committed to git, polluting history with binaries
- [[06_FAILURE_LIBRARY/2026-03-healingstone-module-boundary-drift]] — duplicated schema/core modules between 2D and 3D pipelines

## Committed Artifact Violations
- `backup_untracked/` — entire directory tree of backup artifacts committed to git history
- Must be removed with BFG Repo Cleaner or `git filter-repo`

## Refactor Actions
- [ ] Remove `backup_untracked/` tree from git history entirely (BFG or git filter-repo)
- [ ] Normalize 2D and 3D interfaces — unify duplicate schema/core modules into one shared interface
- [ ] Publish sample datasets to GitHub Releases or DVC (NOT in the repo)
- [ ] Write algorithm benchmarks (synthetic + real fragment pairs, with accuracy metrics)
- [ ] Add pyproject.toml with explicit tool config (mypy, ruff, pytest settings)
- [ ] Write architecture diagram at docs/architecture.md

## Spec Kit Status
- specify init: NOT YET RUN
- Reason: Spec Kit is initialized when the first post-onboarding feature begins
- Next feature to spec: Unified 2D/3D pipeline interface with mode parameter
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/healingstone/`

## Linked Concepts
- [[03_CORE_KNOWLEDGE/ai-ml]] — Siamese network learning, descriptor matching
- [[03_CORE_KNOWLEDGE/algorithms]] — cosine similarity, graph matching for fragment alignment

## Promotion Criteria
- [ ] All refactor actions complete
- [ ] backup_untracked purged from git history
- [ ] CI passing: mypy + ruff + pytest + coverage
- [ ] Algorithm benchmark report with reproducible results
- [ ] Architecture diagram at docs/architecture.md
- [ ] All 3 ADRs written and reviewed
- [ ] README follows docs contract: problem, status, quick start, validation, architecture


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

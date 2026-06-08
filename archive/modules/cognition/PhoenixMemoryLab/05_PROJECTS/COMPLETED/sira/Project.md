---
project-id: sira
repo: https://github.com/fallofpheonix/sira
status: ACTIVE
audit-verdict: "PROMOTE AS RESEARCH PROJECT"
audit-scores: ED8/AQ7/SC5/RU7/MT7/OR8/PV8/CP8/RU6/EC8/LV8/RS8/OS7
language: Python
started: 2026-03
last-commit: 2026-03-17
tags: [project, existing, epidemiology, scientific-computing, simulation]
---

# Sira

## What It Is
An epidemic stochastic simulation platform that trains a neural approximation of the SIR epidemiological vector field from Gillespie stochastic simulations, then serves the trained model via FastAPI for real-time inference. The full pipeline: Gillespie stochastic simulations → ensemble averaging → finite-difference derivatives → vector-field neural network training → SINDy sparse regression → FastAPI inference endpoint. Built for computational epidemiology research.

## Current State
- **What is working:** Gillespie simulator, SINDy sparse identification, MLP vector-field approximation, tests passing, FastAPI service for model inference.
- **What is broken or incomplete:** duplicated `submission/` source tree in main branch (leftover from submission to a competition/course), legacy `src/` package boundary conflicts with the main package root.
- **Audit-identified defects:**
  - Duplicated submission/ source tree — full copy of source code exists in a submission directory
  - Legacy src package boundary — two conflicting package roots cause import confusion
- **Production readiness:** PARTIAL — core science pipeline works, but the repo structure needs cleanup before it can be treated as a reusable research tool.

## Architecture

**Current:**
```
sira/
├── src/sira/         ← one package root (legacy)
├── sira/             ← another package root (conflict!)
├── submission/       ← VIOLATION: duplicated source tree
├── models/           ← trained MLP weights
├── data/             ← simulation output data
├── tests/
├── api/              ← FastAPI service
└── notebooks/        ← Jupyter analysis notebooks
```

**Target (from audit):**
```
sira/
├── src/sira/
│   ├── simulator/    ← Gillespie stochastic engine
│   ├── training/     ← MLP vector-field training
│   ├── sindy/        ← SINDy sparse regression
│   ├── api/          ← FastAPI inference service
│   └── utils/
├── tests/
├── experiments/      ← experiment cards with params + expected output
├── benchmarks/       ← OOD evaluation, convergence tests
├── docs/adr/
└── pyproject.toml
```

**Gap:** The duplicated submission/ tree and conflicting package roots must be resolved. Experiment reproducibility needs formal cards.

## Tech Stack
- **Language:** Python 3.10+
- **ML:** PyTorch (MLP), SINDy (sparse identification)
- **Simulation:** Gillespie algorithm (stochastic), ensemble averaging
- **API:** FastAPI + Uvicorn
- **Math:** NumPy, SciPy (finite differences, ODE integration)
- **Testing:** pytest

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-runtime]] | Not yet imported | Could wrap simulation executor with retry logic |
| [[08_MODULES/fpx-pipeline]] | Not yet imported | Natural fit for Gillespie → train → serve pipeline |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-012-sira-sindy-dynamics]] — why SINDy was chosen for sparse identification of dynamical systems
- [[04_ENGINEERING/decision-logs/ADR-013-sira-fastapi-service]] — why FastAPI over Flask/Django for the inference service
- [[04_ENGINEERING/decision-logs/ADR-014-sira-stochastic-model]] — why stochastic (Gillespie) simulation over deterministic ODE solvers

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-03-sira-submission-duplication]] — full source tree duplicated in submission/ directory
- [[06_FAILURE_LIBRARY/2026-03-sira-package-boundary-conflict]] — two conflicting Python package roots cause import confusion

## Committed Artifact Violations
- `submission/` directory — duplicated source tree (move to branch/release tag)
- Trained model weights in `models/` — consider moving to releases

## Refactor Actions
- [ ] Remove duplicated `submission/` directory from main branch (keep in separate branch or release tag)
- [ ] Resolve legacy `src/` package boundary — establish one canonical package root at `src/sira/`
- [ ] Publish reproducible experiment cards (input parameters + expected output + metrics)
- [ ] Add OOD (out-of-distribution) benchmarks — test model at epidemic parameters outside training range
- [ ] Add pyproject.toml with build config
- [ ] Write architecture diagram at docs/architecture.md

## Spec Kit Status
- specify init: NOT YET RUN
- Reason: Spec Kit is initialized when the first post-onboarding feature begins
- Next feature to spec: OOD evaluation framework for neural vector-field approximation
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/sira/`

## Linked Concepts
- [[03_CORE_KNOWLEDGE/ai-ml]] — neural network training, sparse identification
- [[03_CORE_KNOWLEDGE/algorithms]] — stochastic simulation, Gillespie algorithm, Monte Carlo methods

## Promotion Criteria
- [ ] submission/ directory removed from main
- [ ] Single canonical package root established
- [ ] Experiment cards with reproducible results
- [ ] OOD benchmark passing
- [ ] CI passing: lint + test
- [ ] All 3 ADRs written
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

---
project-id: particle-stimulator
repo: https://github.com/fallofpheonix/ParticleStimulator
status: REFACTOR
audit-verdict: "REFACTOR THEN PROMOTE"
audit-scores: ED8/AQ7/SC6/RU6/MT6/OR8/PV8/CP8/RU5/EC7/LV8/RS8/OS6
language: Python
started: 2026-03
last-commit: 2026-03-17
tags: [project, existing, physics-simulation, scientific-computing, webgl]
---

# ParticleStimulator

## What It Is
A Monte Carlo particle physics simulator modelling proton-proton collisions at LHC energies. Full pipeline: beam generation → accelerator transport → hard scattering (QCD 2→2) → parton shower → hadronisation → detector simulation → event reconstruction → physics analysis. Features a React/Three.js frontend that streams live simulation results via WebSocket. This IS the planned "Game Physics Engine" — do NOT build a separate repo.

## Current State
- **What is working:** Staged simulation pipeline (BeamSource → BeamDynamics → CollisionEngine → DetectorSimulator → EventReconstructor → PhysicsAnalyser), WebSocket live event streaming, React/Three.js 3D visualization frontend, Higgs classifier (scikit-learn/XGBoost), FastAPI health/simulate/train/predict endpoints.
- **What is broken or incomplete:** Large archive directory included in main branch (should be a git tag), no typed event schema for WebSocket messages, ML and job queues are local only (no persistence/async), no performance benchmarks, frameworkless HTTP server may limit scalability, no backpressure handling on WebSocket stream.
- **Audit-identified defects:**
  - Large archive included in main branch — history artifacts polluting current code
  - ML/job queues are local only — no async job persistence for long simulations
- **Production readiness:** NOT READY — archive cleanup and typed events are prerequisites.

## Architecture

**Current:**
```
ParticleStimulator/
├── simulation/        ← staged physics pipeline
├── server/            ← frameworkless HTTP + WebSocket server
├── frontend/          ← React + Three.js visualization
├── ml/                ← Higgs classifier (sklearn/XGBoost)
├── archive/           ← VIOLATION: should be git tag, not in main
└── tests/
```

**Target (from audit):**
```
ParticleStimulator/
├── simulation_core/   ← typed staged pipeline with deterministic seeds
├── event_stream/      ← typed WebSocket events with backpressure
├── analysis/          ← ML classifiers + physics analysis
├── api/               ← FastAPI (replace frameworkless HTTP)
├── frontend/          ← React + Three.js
├── benchmarks/        ← particles/sec, WebSocket throughput, frame rate
├── docs/adr/
└── pyproject.toml
```

## Tech Stack
- **Language:** Python 3.10+
- **Physics:** Custom Monte Carlo simulation engine
- **ML:** scikit-learn, XGBoost (Higgs classifier)
- **API:** FastAPI + Uvicorn (partially, some frameworkless HTTP)
- **Real-time:** WebSocket (live event stream)
- **Frontend:** React, Three.js, Vite
- **Testing:** pytest

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-runtime]] | Not yet imported | Executor for simulation jobs |
| [[08_MODULES/fpx-pipeline]] | Not yet imported | Perfect fit for staged simulation pipeline |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-015-particle-websocket-streaming]] — why WebSocket for real-time simulation data streaming
- [[04_ENGINEERING/decision-logs/ADR-016-particle-react-threejs]] — why React + Three.js for 3D particle visualization
- [[04_ENGINEERING/decision-logs/ADR-017-particle-frameworkless-http]] — why frameworkless HTTP was used instead of FastAPI

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-03-particle-archive-in-main]] — large archive directory committed to main branch
- [[06_FAILURE_LIBRARY/2026-03-particle-no-job-persistence]] — simulation jobs are local-only with no async persistence

## Committed Artifact Violations
- `archive/` directory — move to versioned git tag, remove from main
- Check for compiled C binaries or built frontend assets in git

## Refactor Actions
- [ ] Move archive out of main branch into a versioned git tag
- [ ] Add pyproject.toml with explicit tool config
- [ ] Define typed event schema (Pydantic models for all WebSocket events)
- [ ] Replace local job execution with async job queue (Celery or arq)
- [ ] Add run persistence (save simulation states to SQLite or Postgres)
- [ ] Add deterministic seeds (same seed = identical simulation output every time)
- [ ] Write perf benchmark suite (particles/sec, WebSocket throughput, frontend frame rate)
- [ ] Add backpressure-aware WebSocket stream
- [ ] Migrate frameworkless HTTP to FastAPI

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec: Typed WebSocket event schema + backpressure handling
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/particle-stimulator/`

## Linked Concepts
- [[03_CORE_KNOWLEDGE/algorithms]] — Monte Carlo simulation, stochastic processes
- [[03_CORE_KNOWLEDGE/networking]] — WebSocket protocol, real-time streaming, backpressure

## Promotion Criteria
- [ ] Archive moved to git tag, removed from main
- [ ] Typed event schema defined and enforced
- [ ] Async job queue with persistence
- [ ] Deterministic seeds working
- [ ] Performance benchmark suite with published results
- [ ] All 3 ADRs written
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

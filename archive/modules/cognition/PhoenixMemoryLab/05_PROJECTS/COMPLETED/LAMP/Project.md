---
project-id: lamp
repo: https://github.com/fallofpheonix/LAMP
status: ACTIVE
audit-verdict: "PROMOTE AS PRIMARY PROJECT"
audit-scores: ED8/AQ8/SC6/RU7/MT8/OR8/PV9/CP9/RU7/EC8/LV8/RS8/OS8
language: Python
started: 2026-03
last-commit: 2026-03-31
tags: [project, existing, archaeology, gis, geospatial]
---

# LAMP — Late Antiquity Modeling Project

## What It Is
Deterministic geospatial pipelines for archaeological path tracing, 3D viewshed analysis, and visibility-coupled movement inference in the El Bagawat necropolis. Models ancient movement patterns using slope, roughness, surface penalty, path priors, and optional visibility coupling over real terrain data. The system produces reproducible GIS artifacts (rasters, viewshed maps, cost surfaces) for archaeological research publications.

## Current State
- **What is working:** Coherent domain model, well-structured package layout, tests passing, Docker support, CI pipeline, CLI entry points (lamp path-tracing, lamp viewsheds-2d, lamp viewsheds-3d, lamp validate-dataset, lamp ml-diagnostics, lamp security-audit, lamp benchmark-raycast).
- **What is broken or incomplete:** GDAL/environment reproducibility issues (geospatial dependency hell between macOS and Linux), data provenance gap (undocumented data lineage for input terrain/building data), no benchmark datasets published, no formal ADRs, no artifact output contracts.
- **Audit-identified defects:**
  - GDAL/env reproducibility — geospatial libraries behave differently across platforms
  - Data provenance gap — input data sources not formally documented or versioned
- **Production readiness:** PARTIAL — solid domain code, but reproducibility gaps block promotion.

## Architecture

**Current:**
```
LAMP/
├── src/lamp/
│   ├── api/             ← FastAPI endpoints (if any)
│   ├── core/            ← domain logic
│   ├── path_tracing/    ← probabilistic path computation
│   ├── viewsheds/       ← 2D and 3D viewshed generation
│   ├── io/              ← raster/vector I/O
│   └── terrain/         ← terrain analysis utilities
├── tests/
├── Dockerfile
├── configs/
└── data/                ← input data (provenance undocumented)
```

**Target (from audit):**
```
LAMP/
├── src/lamp/
│   ├── api/
│   ├── core/
│   ├── io/
│   ├── terrain/
│   ├── tasks/
│   │   ├── path_tracing/
│   │   └── viewsheds/
│   └── services/
├── datasets/
│   ├── sample/          ← golden GIS fixtures
│   └── contracts/       ← input/output schema validation
├── benchmarks/
│   ├── visibility/
│   └── path_tracing/
├── docs/adr/
└── pyproject.toml
```

## Tech Stack
- **Language:** Python 3.10+
- **Geospatial:** GDAL/OGR, rasterio
- **API:** FastAPI + Uvicorn
- **Persistence:** PostgreSQL + PostGIS
- **Containers:** Docker, Docker Compose
- **Testing:** pytest
- **CI:** GitHub Actions

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-pipeline]] | Not yet imported | Path tracing + viewshed as pipeline stages |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-006-lamp-gdal-geospatial]] — why GDAL was chosen as the geospatial processing library
- [[04_ENGINEERING/decision-logs/ADR-007-lamp-path-tracing-algorithm]] — path tracing algorithm selection and cost surface design
- [[04_ENGINEERING/decision-logs/ADR-008-lamp-docker-reproducibility]] — why Docker was required for environment reproducibility

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-03-lamp-gdal-pinning-failure]] — GDAL version differences between macOS and Linux cause runtime failures
- [[06_FAILURE_LIBRARY/2026-03-lamp-data-provenance-gap]] — input terrain data has no formal provenance documentation

## Committed Artifact Violations
- Check for committed raster outputs (.tif files) in data/ directory
- Run git log scan for committed GIS artifacts

## Refactor Actions
- [ ] Pin GDAL version in Dockerfile and document platform caveats (macOS vs Linux)
- [ ] Create golden GIS fixture datasets (reproducible benchmark input)
- [ ] Write raster alignment contracts (input/output schema validation)
- [ ] Add deterministic CLI output (given same input, output must be byte-identical)
- [ ] Write benchmark report generation (visibility + path_tracing benchmarks)
- [ ] Document data provenance for all input terrain/building datasets
- [ ] Write architecture diagram at docs/architecture.md

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec: Golden benchmark dataset + deterministic output validation
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/lamp/`

## Linked Concepts
- [[03_CORE_KNOWLEDGE/algorithms]] — probabilistic path algorithms, cost surface optimization
- [[03_CORE_KNOWLEDGE/databases]] — PostGIS spatial queries, raster data management

## Promotion Criteria
- [ ] GDAL pinned and cross-platform reproducibility verified
- [ ] Golden benchmark datasets published
- [ ] Deterministic CLI output validated
- [ ] Data provenance documented
- [ ] All 3 ADRs written
- [ ] CI passing: lint + test + Docker build
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

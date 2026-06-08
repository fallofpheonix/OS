---
project-id: udie
repo: https://github.com/fallofpheonix/UDIE
status: ACTIVE
audit-verdict: "PROMOTE AS PRIMARY PROJECT"
audit-scores: ED9/AQ8/SC8/RU7/MT7/OR8/PV9/CP10/RU7/EC9/LV8/RS9/OS7
language: TypeScript, Python
started: 2026-02
last-commit: 2026-03-24
tags: [project, existing, distributed-systems, urban-intelligence, spatial]
---

# UDIE — Universal Disruption Intelligence Engine

## What It Is
A spatial intelligence system that converts volatile, multi-source urban disruption signals into a stable operational risk view. Uses H3 hexagonal spatial indexing, event-sourced persistence with PostgreSQL + PostGIS, and deterministic materialization for high-performance route risk evaluation and city-level intelligence. Includes NestJS backend, Python spatial utilities, Redis hot-path caching, and dual mobile clients (Flutter + iOS native). The most architecturally ambitious project in the portfolio.

## Current State
- **What is working:** Migrations, metrics, benchmarks, CI, infrastructure (Postgres/Redis/Prometheus/Grafana/Kubernetes), mobile clients (Flutter + iOS), dual backend surfaces, ingestion substrate, spatial compute, projections & workers, operational interface.
- **What is broken or incomplete:** Too many migrations (need squashing to clean baseline), possible architecture sprawl (dual backend ownership unclear), no deployment SLOs, no load tests, API contracts not formalized in OpenAPI.
- **Audit-identified defects:**
  - Too many migrations — migration sprawl without squashing strategy
  - Architecture sprawl — dual backend surfaces with unclear ownership
- **Production readiness:** PARTIAL — most complete system, but migration debt and API contract gaps block promotion.

## Architecture

**Current:**
```
UDIE/
├── engine-backend/
│   ├── src/modules/
│   │   ├── ingestion/     ← signal normalization + event log append
│   │   ├── risk/          ← H3-indexed risk field evaluation
│   │   ├── forecast/      ← predictive models
│   │   ├── digital-twin/  ← city simulation
│   │   ├── routing/       ← risk-aware routing
│   │   ├── traffic-control/ ← signal management
│   │   └── metrics/       ← observability
│   ├── src/platform/
│   │   ├── database/      ← PostGIS + migrations
│   │   ├── observability/ ← Prometheus + Grafana
│   │   ├── config/
│   │   ├── auth/
│   │   └── queues/
│   └── migrations/        ← TOO MANY — need squashing
├── client-mobile-flutter/
├── ios/
├── infra/
│   ├── postgres/
│   ├── redis/
│   ├── prometheus/
│   ├── grafana/
│   └── kubernetes/
└── docs/
```

**Target (from audit):**
```
Same structure but with:
├── migrations/
│   ├── baseline/          ← squashed to single clean migration
│   ├── squashed/          ← historical reference
│   └── incremental/       ← new migrations only
├── docs/
│   ├── architecture/
│   ├── api/               ← OpenAPI specs
│   ├── runbooks/
│   ├── adr/
│   └── benchmarks/
```

## Tech Stack
- **Backend:** NestJS (TypeScript), Python (Spatial Utils)
- **Persistence:** PostgreSQL + PostGIS (authoritative), Redis (hot-path caching)
- **Mobile:** Flutter (cross-platform), Swift (iOS native)
- **Indexing:** H3 (Uber's Hexagonal Hierarchical Spatial Index)
- **Observability:** Prometheus, Grafana
- **Infra:** Kubernetes, Docker Compose
- **Patterns:** Event Sourcing, CQRS
- **License:** MIT

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|
| [[08_MODULES/fpx-runtime]] | Not yet imported | Backend executor abstraction |
| [[08_MODULES/fpx-observability]] | Source project — will be extracted FROM this repo | Event sourcing + CQRS patterns |
| [[08_MODULES/smart-rate-limiter]] | Not yet imported | API protection for public endpoints |

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-001-udie-nestjs-over-express]] — why NestJS over Express or Fastify
- [[04_ENGINEERING/decision-logs/ADR-002-udie-postgres-postgis]] — why PostgreSQL + PostGIS over alternatives
- [[04_ENGINEERING/decision-logs/ADR-003-udie-redis-cache-stream]] — why Redis for cache + stream coordination
- [[04_ENGINEERING/decision-logs/ADR-004-udie-flutter-ios-dual-client]] — why both Flutter and iOS native clients
- [[04_ENGINEERING/decision-logs/ADR-005-udie-kubernetes-infra]] — why Kubernetes for infrastructure control-plane

## Known Failure Modes
- [[06_FAILURE_LIBRARY/2026-02-udie-migration-sprawl]] — migrations accumulated without squashing, making rollback dangerous
- [[06_FAILURE_LIBRARY/2026-03-udie-dual-backend-ownership]] — two backend surfaces with unclear module ownership

## Committed Artifact Violations
- Run git log scan for committed artifacts — audit does not flag specifics

## Refactor Actions
- [ ] Consolidate backend ownership — single entry point, not dual surfaces
- [ ] Formalize API contracts (OpenAPI spec for all public endpoints)
- [ ] Run migration squashing strategy (collapse all migrations to clean baseline)
- [ ] Add deployment SLOs (latency p95, error rate, throughput targets)
- [ ] Add load tests (k6 or locust against the API gateway)
- [ ] Add seeded reproducible benchmark city (for demo/testing)
- [ ] Write architecture diagram at docs/architecture.md

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec: Migration squashing + API contract formalization
- When ready: `specify init . --integration claude` inside `~/engineering/workspace/active/udie/`
- **Known Constraints:** Do NOT add new migrations until squashing is done

## Linked Concepts
- [[03_CORE_KNOWLEDGE/distributed-systems]] — event sourcing, CQRS, materialized views
- [[03_CORE_KNOWLEDGE/databases]] — PostgreSQL, PostGIS, spatial indexing, Redis caching
- [[03_CORE_KNOWLEDGE/networking]] — real-time data streaming, API gateway patterns

## Promotion Criteria
- [ ] Migration baseline squashed and clean
- [ ] API contracts in OpenAPI format
- [ ] Backend ownership consolidated (single surface)
- [ ] Load test suite passing at defined SLO targets
- [ ] Seeded benchmark city working
- [ ] Architecture diagram at docs/architecture.md
- [ ] All 5 ADRs written
- [ ] README follows docs contract


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

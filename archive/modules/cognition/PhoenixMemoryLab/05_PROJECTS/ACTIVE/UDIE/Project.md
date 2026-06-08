# Project: UDIE

## One-Liner
UDIE — Universal Disruption Intelligence Engine

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/UDIE`

## Ports
- API: N/A
- DB: N/A

## Database
N/A

## Run Command
N/A - historical project overview

## Dependencies On Other Projects
None

## What I Deliver To Others
None

## Links
- [[03_CORE_KNOWLEDGE/ai-ml/AI]]
- [[04_ENGINEERING/architecture-patterns/Software-Engineering]]
- [[04_ENGINEERING/system-design/System Design]]
- [[04_ENGINEERING/architecture-patterns/Frontend Architecture]]
- [[03_CORE_KNOWLEDGE/security/Security]]
- [[Decisions]]
- [[Mistakes]]

## Current Blockers
None

## Last Worked On
2026-05-12

## Original Overview


**Repository:** [github.com/fallofpheonix/UDIE](https://github.com/fallofpheonix/UDIE)  
**Language:** TypeScript | **License:** MIT | **Forks:** 1 | **Created:** 2026-02-12

---

## Project Summary

UDIE is a spatial intelligence system that converts volatile, multi-source urban disruption signals into a stable operational risk view. It uses H3 spatial indexing, event-sourced persistence, and deterministic materialization for high-performance route risk evaluation and city-level intelligence.

## Core Subsystems

1. **Ingestion Substrate** — Normalizes and appends raw signals to the authoritative event log
2. **Spatial Compute** — H3-indexed aggregation and risk field evaluation
3. **Projections & Workers** — Materializes derived states (risk cells, hotspots) for querying
4. **Operational Interface** — Mobile and web clients for real-time visualization

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | NestJS (TypeScript), Python (Spatial Utils) |
| Persistence | PostgreSQL + PostGIS (Authoritative), Redis (Hot-path caching) |
| Mobile | Swift (iOS Native), Flutter (Cross-platform) |
| Indexing | H3 (Uber's Hexagonal Hierarchical Spatial Index) |

## Skills Demonstrated

`TypeScript`, `NestJS`, `PostgreSQL`, `PostGIS`, `Redis`, `H3 Spatial Indexing`, `Event Sourcing`, `CQRS`, `Real-time Visualization`, `Urban Intelligence`, `Flutter`, `Swift`, `Distributed Systems`

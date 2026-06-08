# Project: AI4MH

## One-Liner
AI4MH — AI for Mental Health

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/AI4MH`

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
- [[03_CORE_KNOWLEDGE/ai-ml/Machine Learning]]
- [[04_ENGINEERING/architecture-patterns/Frontend Architecture]]
- [[Decisions]]
- [[Mistakes]]

## Current Blockers
None

## Last Worked On
2026-05-12

## Original Overview


**Repository:** [github.com/fallofpheonix/AI4MH](https://github.com/fallofpheonix/AI4MH)  
**Language:** Python | **License:** MIT | **Created:** 2026-03-10

---

## Project Summary

Full-stack demo for regional mental-health crisis signal monitoring. Ingests synthetic discussion posts, enriches with sentiment and keyword signals, aggregates regional scores, persists alert state in SQLite, and exposes results through FastAPI API and React dashboard.

## Architecture

```
Frontend (React/Vite) ← HTTP → Backend (FastAPI)
                                  ↓
                      Scoring Pipeline
                      (Sentiment + Keywords)
                                  ↓
                      SQLite Persistence
                                  ↓
                      Alert State Management
```

## API Surface

| Endpoint | Method | Description |
|---|---|---|
| `/ingest?n=30` | POST | Ingest synthetic posts |
| `/posts?limit=20` | GET | List posts |
| `/scores` | GET | Regional scores |
| `/alerts` | GET | Active alerts |
| `/alerts/{id}/ack` | POST | Acknowledge alert |
| `/alerts/{id}/resolve` | POST | Resolve alert |
| `/logs` | GET | System logs |
| `/bias` | GET | Bias diagnostics |

## Deployment

- Docker Compose: backend on `:8000`, frontend on `:80`
- Makefile with standard commands

## Skills Demonstrated

`Python`, `FastAPI`, `React`, `Vite`, `SQLite`, `Sentiment Analysis`, `NLP`, `Docker Compose`, `Full-Stack Development`, `Mental Health Tech`, `Alert Management`

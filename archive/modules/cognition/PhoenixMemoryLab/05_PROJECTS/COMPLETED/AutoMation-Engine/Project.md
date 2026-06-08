# Project: AutoMation-Engine

## One-Liner
AutoMation-Engine

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/AutoMation-Engine`

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


**Repository:** [github.com/fallofpheonix/AutoMation-Engine](https://github.com/fallofpheonix/AutoMation-Engine)  
**Language:** Python | **License:** MIT | **Created:** 2026-04-23

---

## Project Summary

A production-grade task automation system with cross-platform support, built with FastAPI and designed for Windows/macOS automation. Demonstrates advanced backend engineering with intelligent retry logic, error classification, and abstracted executor patterns.

## System Architecture

```
API Layer (FastAPI)
  POST /tasks/execute  →  GET /tasks/{id}  →  GET /metrics
       ↓
Task Orchestrator
  Coordinates: Retry Logic → Executor → Logging
       ↓
┌────────────────┬────────────────────┬────────────────────────┐
│  Retry Engine  │ Error Classifier   │  Executor Interface    │
│  • Backoff     │  • Temp vs Perm    │  • MacExecutor ✓       │
│  • Max retries │  • Classification  │  • WindowsExecutor →   │
│  • Timeout     │                    │  • SimulationMode      │
└────────────────┴────────────────────┴────────────────────────┘
       ↓
SQLite Database (Tasks, Steps, Logs)
```

## Core Features

- **Abstracted Executor Layer** — Swap macOS/Windows implementations without changing API
- **Intelligent Retry Logic** — Exponential backoff only for temporary errors
- **Error Classification** — Temporary vs permanent error detection
- **Complete Logging** — Every action logged with timestamps and duration
- **Timeout Handling** — Prevent tasks from running forever
- **Production Observability** — Metrics endpoint for system health

## API Endpoints

| Endpoint | Method | Description |
|---|---|---|
| `/tasks/execute` | POST | Execute task immediately |
| `/tasks/{task_id}` | GET | Get task status |
| `/tasks/{task_id}/logs` | GET | Get all logs for a task |
| `/metrics` | GET | System metrics (success rate, avg execution time) |
| `/health` | GET | Health check |

## Supported Actions

| Action | Parameters | Example |
|---|---|---|
| `open_app` | `app` (string) | `{"action": "open_app", "app": "Notes"}` |
| `click` | `target` (string/coords) | `{"action": "click", "target": "button_name"}` |
| `type` | `text` (string) | `{"action": "type", "text": "Hello"}` |
| `wait` | `seconds` (float) | `{"action": "wait", "seconds": 2}` |
| `close_app` | `app` (string) | `{"action": "close_app", "app": "Notes"}` |

## Performance Metrics

- Average action duration: 450-500ms
- P99 action duration: 2-3 seconds
- Success rate under 30% failures: 94.7%
- Tests: 18+ passing

## Key Design Decisions

1. **Executor Abstraction** — Enables cross-platform without being macOS-only
2. **Exponential Backoff** — 96% success vs 40% with immediate retry
3. **SQLite** — Tasks persist across server crashes
4. **Simulation Mode** — Validate Windows logic on macOS

## Skills Demonstrated

`Python`, `FastAPI`, `SQLite`, `REST API Design`, `Cross-Platform Architecture`, `Retry Patterns`, `Error Classification`, `Exponential Backoff`, `Test-Driven Development`, `Production Observability`

## Tech Stack

- **Framework:** FastAPI + Uvicorn
- **Database:** SQLite
- **Validation:** Pydantic
- **Testing:** Pytest
- **Windows Support:** pywinauto (designed), psutil

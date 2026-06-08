---
failure-id: 2026-03-particle-no-job-persistence
project: [[05_PROJECTS/ACTIVE/particle-stimulator]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, architecture, reliability]
---
# Failure: Simulation jobs are local-only with no async persistence

## What Was Tried
Running Monte Carlo simulations as in-process tasks within the WebSocket server.

## What Happened
Simulation jobs are tied to the server process. If the server crashes or restarts, all running simulations are lost — no state is persisted. Long-running simulations (millions of events) cannot survive process restarts.

## Root Cause
No job queue or persistence layer was implemented. Simulations run directly in the request handler's process, with results streamed immediately via WebSocket.

## What Was Learned
Long-running computational jobs must be separated from the web server process. Use an async job queue (Celery, arq, or even SQLite-backed) so that jobs survive server restarts and can be monitored independently.

## Prevention / Resolution
- Add async job queue (arq or Celery) for simulation execution
- Persist simulation state to SQLite or Postgres (checkpoint every N events)
- Add job status endpoint (GET /api/jobs/{id}) for monitoring
- Implement graceful shutdown (checkpoint current state before exit)

## Linked Concepts
- [[03_CORE_KNOWLEDGE/distributed-systems]] — job queues, task persistence, graceful degradation
- [[03_CORE_KNOWLEDGE/architecture]] — separation of concerns, web server vs worker pattern

---
failure-id: 2026-02-udie-migration-sprawl
project: [[05_PROJECTS/ACTIVE/udie]]
severity: HIGH
status: OPEN
date-encountered: 2026-02
tags: [failure, database, migrations]
---
# Failure: Migration sprawl without squashing strategy in UDIE

## What Was Tried
Database schema evolved iteratively with each new module (ingestion, risk, forecast, digital-twin, routing, traffic-control, metrics). Each change generated a new migration file.

## What Happened
Migration count grew unbounded. No squashing strategy was implemented. The migration chain became fragile — any failure mid-chain leaves the database in an inconsistent state. Rolling back to a clean state requires running dozens of migrations in sequence.

## Root Cause
No migration lifecycle management. Migrations were created on every schema change without periodic squashing to a clean baseline.

## What Was Learned
Migrations must be squashed to a clean baseline periodically (e.g., every release or every 10 migrations). The squashed baseline becomes the new starting point; incremental migrations build on top.

## Prevention / Resolution
- Implement a three-tier migration strategy: baseline/ (squashed), squashed/ (historical), incremental/ (new)
- Squash all current migrations to a single baseline migration
- Add a CI check that flags when migration count exceeds threshold

## Linked Concepts
- [[03_CORE_KNOWLEDGE/databases]] — schema migration strategies, database versioning
- [[03_CORE_KNOWLEDGE/distributed-systems]] — stateful service deployment, data migration in distributed systems

# FAILURE_RECOVERY_ANALYSIS.md

## Overview
In a distributed orchestration environment, partial failures are inevitable. Astraeus must be capable of surviving worker crashes, split-brain scenarios, and synchronization failures.

## 1. Worker Crashes & Task Reassignment
- **Failure Mode**: A worker process pulls a task, claims a lease, but crashes before emitting `TASK_SUCCEEDED` or `TASK_FAILED`.
- **Detection**: The central coordinator detects the failure via a missed heartbeat or an expired task lease.
- **Recovery**: The coordinator automatically invalidates any partial artifacts produced by the crashed worker, resets the task status to `PENDING`, and re-queues the task. The task's retry counter is NOT incremented (this is an infra failure, not a model/logic failure).

## 2. Split-Brain Orchestration
- **Failure Mode**: Network partitioning causes two coordinators to believe they own the same `run_id`.
- **Detection**: Both attempt to renew the `run_id` lease in the central database/lockfile.
- **Recovery**: The lease mechanism relies on a single source of truth (e.g., SQLite with strict locking, or Redis). The second coordinator will fail to acquire/renew the lease. Upon lease renewal failure, the coordinator MUST instantly suicide (terminate its process) to avoid issuing conflicting commands or mutating shared state.

## 3. Concurrent Rollback Storms
- **Failure Mode**: Multiple independent failures trigger simultaneous rollback requests for the same repository.
- **Detection**: The `RollbackEngine` attempts to acquire the Global Write Lock.
- **Recovery**: The Global Write Lock serializes the rollbacks. The first rollback wins and executes. The subsequent rollbacks wait. When they acquire the lock, they perform a Pre-flight Hash Check. Because the first rollback altered the file hashes, the subsequent rollbacks will fail their drift checks (unless explicitly forced) and safely abort, preventing a corruption storm.

## 4. Partial Replay Recovery
- **Failure Mode**: A crash occurs while the `ReplayEngine` is reconstructing state, leaving the filesystem partially rolled back or partially fast-forwarded.
- **Detection**: Replay tracks the `sequence_id` of the last successfully applied event.
- **Recovery**: Replay operations must be idempotent. If a replay is interrupted, rerunning the replay simply skips events whose `sequence_id` is less than or equal to the last known good state, or safely reapplies them if they are idempotent (e.g., rewriting a file with identical content).

## 5. Event Desynchronization
- **Failure Mode**: Workers emit events out of order due to network latency.
- **Detection**: The central EventBroker receives an event.
- **Recovery**: Events are ordered by the *broker's* assigned `sequence_id`, not the worker's timestamp. If a worker sends `TASK_SUCCEEDED` for task A, but the broker hasn't yet processed `TASK_STARTED` for task A, the broker can queue the event in a staging area until the predecessor arrives, or strictly enforce causal ordering by rejecting out-of-sequence state transitions.

## 6. Journal Corruption
- **Failure Mode**: A sudden power loss mid-write corrupts `journal.jsonl`.
- **Detection**: `FilesystemJournal.verify_integrity()` detects a broken hash chain or incomplete JSON payload.
- **Recovery**: The journal is truncated to the last known valid `previous_hash`. Any orphaned files in the project root must be manually reconciled, but the rollback engine will refuse to act on the corrupted partial entries.

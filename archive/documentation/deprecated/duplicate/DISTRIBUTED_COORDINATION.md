# DISTRIBUTED_COORDINATION.md

## Overview
As Astraeus scales to multiple workers processing tasks in parallel, local process state (`asyncio`, `threading`) is insufficient. This document defines the Distributed Coordination Layer.

## 1. Orchestration Leases
To prevent "split-brain" orchestration where two control-plane instances attempt to manage the same `run_id`:
- **Run Lease**: A coordinator must acquire a lease for a `run_id` (e.g., via SQLite/Postgres with a `locked_by` and `locked_until` column).
- **Lease Renewal**: The active coordinator must periodically heartbeat to renew its lease.
- **Lease Expiry**: If a coordinator crashes, its lease expires, and a standby coordinator can claim the `run_id` and resume execution from the last known state.

## 2. Distributed Event Ordering
- **Sequence Numbers**: The `EventBus` will assign a monotonically increasing `sequence_id` to every event.
- **Atomic Appends**: In a file-based system, events are appended atomically (using `O_APPEND` and strict OS file locks). In a database, an auto-incrementing ID or transaction ensures order.
- **Replay Synchronization**: Replay engines must sort events strictly by `sequence_id`, not by the `timestamp`, which is subject to clock drift across nodes.

## 3. Worker Heartbeats & Task Leases
- **Task Leasing**: When a worker pulls a task from the `ExecutionQueue`, it claims a lease on that `task_id`.
- **Heartbeats**: The worker emits `TASK_HEARTBEAT` events. If heartbeats stop for a configured timeout, the coordinator considers the worker dead, marks the task as `FAILED` (or `RETRY_SCHEDULED`), and re-queues it.

## 4. Coordination Protocol
- **Stateless Workers**: Workers pull inputs (the task payload and necessary repository context), execute generation and validation, and return an outcome artifact. They do not hold the global `RunState` in memory.
- **State Reconciliation**: Only the coordinator updates the global `RunState` based on `TASK_SUCCEEDED` or `TASK_FAILED` events emitted by workers.

## Transition Plan

The implementation of distributed coordination is tracked in the **Master TODO Hierarchy**:
- **Phase G**: Concurrency + Distributed Governance (Worker coordination, leases, distributed ordering).
- **Phase I**: Observability + Operations (Distributed tracing, telemetry).

See **[TODO.md](./TODO.md)** for detailed task breakdown.


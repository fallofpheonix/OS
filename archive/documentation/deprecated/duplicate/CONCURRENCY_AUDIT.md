# CONCURRENCY_AUDIT.md

## Overview
This audit analyzes the Astraeus core substrate for its readiness to handle concurrent orchestration, distributed execution, multi-worker coordination, and safe parallel cognition.

## 1. Asyncio & Threading
- **`orchestrator/queue.py`**: Uses `asyncio.gather` for parallel task execution. However, it mutates a shared, in-memory `RunState` object without formal synchronization. While synchronous dict updates in Python are thread-safe (due to GIL) and async-safe (no `await` yields during the update), this model completely breaks down in a multi-worker distributed setup because `RunState` is localized to a single process.
- **`models/ollama.py`**: Uses a local `threading.Lock` (`_MODEL_GENERATION_LOCK`) to prevent concurrent generation. In a distributed environment with multiple workers hitting the same Ollama instance, this local lock is bypassed, risking OOM errors or timeouts on the model server.

## 2. Locking Mechanisms & Shared State
- **`events/event_bus.py`**: Uses `threading.Lock` before appending to `events.jsonl`. This ensures thread safety within a single process but provides **zero isolation across multiple processes/workers**. Concurrent workers emitting events will corrupt `events.jsonl` or interleave writes unpredictably.
- **In-Memory Subscribers**: Event bus subscribers (`self._subscribers`) are in-memory. A distributed worker cannot notify subscribers running in another process.

## 3. Journal Concurrency & Filesystem Locks
- **`transactions/journal.py`**: Implements `fcntl.flock` (exclusive and shared locks) for `journal.jsonl`. This is a positive step for multi-process safety on a single POSIX machine. However:
  - Advisory locks do not guarantee strict chronological ordering of operations (race to acquire the lock).
  - They may not function correctly over distributed file systems (NFS/EFS) if Astraeus is distributed across nodes.
  - The `_last_entry_hash` is an in-memory cache. If multiple processes write to the journal, process A's cached `_last_entry_hash` will become stale after process B writes, leading to a broken hash chain when process A writes its next entry.

## 4. Rollback Coordination
- **`transactions/rollback.py`**: The `RollbackEngine` performs file restorations without any locking mechanism on the `project_root`. If a rollback happens concurrently with a mutation (or another rollback), file states will become corrupted, and the journal's chain of custody will break.

## 5. Parallel Mutation Governance
- **`transactions/runner.py`**: The `TransactionRunner` stages files locally and then copies them back to `project_root`. There is no path-level or topology-level lock. If two concurrent tasks attempt to mutate the same file, or overlapping directories, the final state is determined by a race condition (last writer wins), violating deterministic execution.

## 6. Event Ordering Inconsistencies
- Event timestamps are generated locally (`datetime.now(UTC)`). In a distributed setup, clock drift between workers will result in out-of-order events.
- Because there is no central sequence coordinator or run ownership, reconstructing execution history from multiple workers will yield replay divergence.

## Summary of Critical Risks
1. **Shared Mutable State**: In-memory `RunState` prevents distributed task execution.
2. **Broken Hash Chains**: Concurrent journal writes by different processes will break the `previous_hash` chain due to stale in-memory states.
3. **Mutation Corruption**: Unlocked transaction commits to `project_root`.
4. **Event Corruption**: Thread-only locks on `event_bus.py` will cause JSONL corruption across processes.
5. **Model Server Overload**: Thread-only locks on Ollama client.

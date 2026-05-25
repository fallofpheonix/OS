# Critical Failure Analysis

This document outlines the hidden failure paths, corruption vectors, and architectural limits of the Astraeus substrate.

## 1. Concurrency and Data Corruption

*   **Event Sourcing Deadlocks:** The `EventBus` uses a blocking `fcntl.flock` on `events.jsonl`. Under high concurrency (e.g., executing a wide DAG of 50 simultaneous tasks), the lock contention will cause severe latency spikes and potential deadlocks if a thread crashes while holding the file lock.
*   **Sequence ID Race Conditions:** The method for determining `sequence_id` involves seeking to the end of the file, parsing the last JSON line, and incrementing. If the file is malformed, corrupted by a partial write, or truncated, the sequence ID resets or crashes the entire bus.
*   **Partial Write Corruption:** If the process is `SIGKILL`'d during `json.dumps(row)` in `AtomicEventLogger` or `EventBus`, the JSONL file will contain a truncated line. The next process reading it will fail `json.loads()`, permanently corrupting the run history.

## 2. Temporal and Replay Failures

*   **Invalid Replay Assumptions:** The `ReplayEngine.reconstruct_state` method merely instantiates a `RunState` object and calls `apply_event()`. It does NOT recreate the filesystem state, the `node_modules` environment, or the external API states that existed at the time of the event. Replay is semantic only, not operational.
*   **Temporal Desynchronization:** Timestamps are recorded using `datetime.now(UTC).strftime`. In a multi-node setup (even multi-process with slight clock drift), event sequence IDs will clash with temporal ordering, destroying causality tracking.

## 3. Mutation and State Collapse

*   **Rollback Storms & Inode Exhaustion:** The `SnapshotEngine` copies all un-ignored files via `shutil.copy2`. In a monolithic repository (e.g., millions of files), creating a snapshot will take minutes, consume gigabytes of disk space, and eventually trigger inode exhaustion.
*   **Hidden Mutable State:** Tasks execute inside `asyncio.to_thread`. If a model generates code that mutates global environment variables, network state, or external databases, the filesystem snapshot cannot roll this back. The "rollback" is dangerously incomplete.
*   **Risk Engine Bypass:** `assess_rollback_risk` in `runtime/snapshots.py` relies on naive path matching (e.g., `p.name in {"Dockerfile", "Makefile"}`). A task modifying `setup.py`, a shell script hidden in a subfolder, or a `.env` file bypasses the "DANGEROUS" risk classification entirely.

## 4. Cognition and Context Explosion

*   **Context Window Overflow:** `_task_prompt` in `engine.py` blindly dumps `dependency_outputs` into the prompt, truncating arbitrarily at `[:6000]` characters. If an upstream task generates a large file, the downstream task receives a silently truncated, invalid dependency context, leading to guaranteed downstream hallucination.
*   **Graph Explosion:** The `SemanticTopologyEngine` is heavily invoked. Re-scanning the repository and building the AST graph synchronously during mutations will scale at $O(N)$ with codebase size, eventually stalling the engine entirely on large enterprise repositories.
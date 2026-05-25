# Failure Simulation Results

The following simulations project the behavior of Astraeus under adversarial, scaled, and degraded conditions based on its current codebase geometry.

## Simulation 1: Concurrent DAG Execution (50+ Node Graph)
**Scenario:** A user requests a refactor touching 50 independent files. The `ExecutionQueue` spawns 50 parallel task threads.
**Result:** **SYSTEM COLLAPSE (I/O Bottleneck)**
**Mechanism:** 
1. 50 threads immediately attempt to emit `TASK_STARTED` to the `EventBus`.
2. 50 threads queue up blocking OS-level file locks via `fcntl.flock` on `events.jsonl`.
3. The underlying Ollama client (default single-instance) receives 50 concurrent HTTP generation requests, instantly OOM-killing the local LLM or triggering aggressive timeout failures.
4. The engine records 50 `TIMEOUT` or `RUNTIME_EXCEPTION` failures, writes them sequentially to the locked file, and emits 50 Help Requests, effectively halting the system.

## Simulation 2: The Rollback Storm
**Scenario:** A task generates an incorrect Python script. The critic flags a `SYNTAX_ERROR`. The system attempts a localized rollback and repair cycle in a repository with 50,000 files.
**Result:** **LATENCY CASCADE**
**Mechanism:**
1. `SnapshotEngine.create()` traverses 50,000 files via `rglob`.
2. It executes `shutil.copy2` 50,000 times to the `/artifacts/snapshots/` directory.
3. This synchronous I/O operation takes ~45 seconds. The task execution loop blocks.
4. If a second task fails shortly after, another 50,000 files are copied. Disk usage spikes by gigabytes per failure. Storage exhaustion is reached within hours of autonomous operation.

## Simulation 3: The Partial Write Corruption
**Scenario:** The host machine experiences a hard power cycle or OOM-killer while `AtomicEventLogger._write_event` is writing to `journal.jsonl`.
**Result:** **REPLAY IMPOSSIBILITY**
**Mechanism:**
1. The JSON string is cut off: `{"run_id": "run_123", "action": "TASK_`
2. Upon restart, `EventBus.read_all()` encounters a `JSONDecodeError`.
3. Because the sequence parser in `EventBus.emit()` scans from the back of the file to find the last valid JSON to increment `sequence_id`, it will fail to parse the trailing garbage.
4. The event bus is permanently bricked for that run until manual human intervention truncates the file.

## Simulation 4: Context Truncation Drift
**Scenario:** Task A generates a 6500-character JSON configuration file. Task B is instructed to read Task A's output and validate it.
**Result:** **SILENT COGNITION FAILURE**
**Mechanism:**
1. `engine.py::_task_prompt` extracts Task A's output.
2. The output is hard-sliced via `[:6000]`.
3. Task B receives a broken, syntactically invalid JSON string.
4. Task B's LLM hallucinates a fix, or fails validation repeatedly.
5. The system initiates a repair cycle on Task B, attempting to fix an error caused invisibly by the orchestration layer. The repair loop exhausts its attempts and fails the run.
# Priority Correction Matrix

The following corrections are ranked by catastrophic risk to the system's viability as an autonomous substrate.

| Priority | Problem | Severity | Failure Type | Fix Complexity | Recommended Action |
|---|---|---|---|---|---|
| **1** | **`shutil.copy2` Snapshotting** | FATAL | I/O Collapse / State Corruption | High | Deprecate `SnapshotEngine`. Integrate native Git branching/stashing, btrfs subvolumes, or OverlayFS for instantaneous $O(1)$ snapshots. |
| **2** | **Blocking File Locks (`fcntl`)** | FATAL | Deadlock / Concurrency Failure | Medium | Replace local JSONL file locks with a robust local database (SQLite WAL mode) or an actual message broker (Redis/Postgres/Kafka) for the Event Bus. |
| **3** | **Prompt Context Truncation (`[:6000]`)** | CRITICAL | Semantic Drift / Hallucination | Low | Implement a proper token-counting sliding window or hierarchical summarization phase for dependency outputs in `engine.py`. |
| **4** | **Naive String-Matching Risk Engine** | CRITICAL | Sandbox Escape / Privilege Escalation | Medium | Replace filename guessing with a rigid capability-based permission model. Block all OS/Network execution outside a strict, isolated container overlay. |
| **5** | **Fake Temporal Replay** | HIGH | False Confidence / Observability Blindness | High | Rename `ReplayEngine` to `LogHydrator`. If true replay is required, integrate strict dependency pinning, mocked network interfaces (VCR.py), and LLM seed enforcement. |
| **6** | **Partial Write JSONL Corruption** | HIGH | Event Log Destruction | Low | Implement atomic file writes for the event log (write to temporary file, `os.rename`), or rely on SQLite's ACID guarantees instead of manual text appending. |
| **7** | **Synchronous LLM `asyncio.to_thread`** | MEDIUM | Resource Exhaustion | Medium | Move LLM execution entirely to an async-native queue structure with backpressure limits to prevent OOM errors on local models. |
| **8** | **Unversioned Event Schemas** | MEDIUM | Long-term Data Loss | Low | Implement schema version tags in `Event` and `RunState` models, and write migration adapters for older event logs. |
# Architectural Weakpoints

A detailed audit of Astraeus's structural design reveals a system heavily reliant on naive abstractions that will shatter under production, distributed, or adversarial workloads.

## 1. Fake Distributed Primitives
**Weakpoint:** `EventBus` and `AtomicEventLogger`.
The architecture claims to use "Event Sourcing," a pattern designed for high-availability distributed systems. However, the implementation is locked strictly to a local filesystem file (`events.jsonl`) using blocking POSIX file locks (`fcntl`). 
**Consequence:** The system cannot scale horizontally. Adding a second worker node on a different machine is impossible without mounting a shared NFS drive, which introduces severe latency, lock contention, and network partition vulnerabilities.

## 2. Dangerous Mutation Semantics
**Weakpoint:** `SnapshotEngine` (`shutil.copy2` based).
True safe-mutation sandboxes use filesystem overlays (OverlayFS), btrfs snapshots, ZFS snapshots, or strict containerized volume mounts. Astraeus uses a python `shutil` loop.
**Consequence:** Filesystem operations are not atomic. A process modifying the repository during the seconds-to-minutes it takes Astraeus to copy the files will result in a corrupted, torn snapshot. Restoring this snapshot will destroy the repository's integrity.

## 3. String-Matching Security
**Weakpoint:** `assess_rollback_risk` in `snapshots.py`.
The "Risk Engine" determines the danger of a mutation by checking if a file name is exactly `"Dockerfile"`, `"Makefile"`, or ends in `.json`/`.yaml` at the root. 
**Consequence:** It is trivial for a malicious or hallucinated LLM prompt to write a destructive shell script to `scripts/exploit.sh`, execute it, and bypass the "DANGEROUS" risk flag entirely because it is not a config file at the root.

## 4. Context Overflow Ignorance
**Weakpoint:** `engine.py` prompt generation (`[:6000]`).
The orchestrator truncates dependency context blindly. It has no token-awareness, no summarization compression, and no hierarchical context retrieval for inter-task communication.
**Consequence:** The system relies on hope that task outputs remain small. As agent autonomy deepens and tasks generate larger artifacts, the context window will silently corrupt upstream data, leading to compounding cognitive failures.

## 5. Non-Deterministic Replay
**Weakpoint:** `ReplayEngine`.
The `ReplayEngine` merely reconstructs a Python `RunState` object from JSON logs. It does not re-execute the LLM calls with fixed seeds, nor does it guarantee the external environment (APIs, filesystem state, time) is identical.
**Consequence:** The replay is an illusion. It is a log viewer, not a temporal execution engine. You cannot use it to deterministically debug a past failure because the environmental context of that failure is lost forever.
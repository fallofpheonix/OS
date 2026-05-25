# RUNTIME AUDIT REPORT

## 1. Nondeterministic Flows & Race Conditions
*   **Orchestration Concurrency:** `ControlPlane` uses `asyncio.Lock` to enforce sequential model execution, but if multiple `ControlPlane` instances or other agents run concurrently, filesystem state mutations could race.
*   **Journal Lock:** `FilesystemJournal` uses standard file appending without a file lock (e.g., `fcntl.flock`). Concurrent runs could interleave JSONL entries, causing corruption.
*   **Transaction Staging:** `TransactionRunner` creates a staging directory `staging_{run_id}`. This avoids some races, but the final copy back to `project_root` is not atomic across multiple files, risking partial commits if interrupted.

## 2. Missing Rollback Guarantees & Unsafe Mutations
*   **Interleave-Unaware Rollback:** The `RollbackEngine` blindly restores backups or unlinks files based strictly on the run's journal entries. If Run A creates `file.py`, and Run B modifies `file.py`, rolling back Run A will delete `file.py`, destroying Run B's work silently.
*   **Missing Pre-Rollback Hash Checks:** `RollbackEngine` does not verify that the file's *current* hash matches the `new_hash` recorded in the journal before rolling it back. This could overwrite unexpected manual user changes.
*   **Coarse Dependency Awareness:** There are no transaction boundaries or dependency graphs between runs.

## 3. Replay Gaps
*   **Superficial Replay:** `ReplayEngine` only verifies the *existence* of event logs and output artifacts. It does not actually *replay* the DAG state machine or verify that the inputs deterministically produce the recorded outputs.
*   **Missing Mutation Traces:** While the `FilesystemJournal` records filesystem changes, the `EventBus` does not contain the unified `DiffPlan` or patch details, disconnecting the orchestrator's state history from the actual mutations.

## 4. Filesystem Safety Gaps
*   **Transaction Atomicity:** As mentioned, multi-file commits in `TransactionRunner.apply()` are sequential `shutil.copy2` calls. A hard crash during the loop leaves the project in a torn state.
*   **Snapshot Engine Isolation:** `SnapshotEngine` (referenced in `runner.py` but relying on journaled backups in `mutation_sandbox.py`) is inconsistently applied.

## 5. Weak Command Gating
*   **Heuristic Risk Engine:** `CommandRiskEngine` uses regex heuristics. While improved recently, it cannot reliably parse complex shell scripts or encoded payloads. It is a defense-in-depth layer, not a strict sandbox boundary.

## Conclusion
The runtime has a solid foundation with the `FilesystemJournal` and `MutationSandbox`, but it lacks strict isolation and verification guarantees. The most critical gap is the **interleaved rollback vulnerability** and the lack of **transaction atomicity**.
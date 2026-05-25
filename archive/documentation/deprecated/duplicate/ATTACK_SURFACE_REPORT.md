# ATTACK SURFACE REPORT

## 1. Shell Execution & Subprocess Spawning
- **Location:** `runtime/sandbox.py`, `transactions/runner.py`
- **Vulnerability:** Unsafe use of `subprocess.run(..., shell=True)`.
- **Exploitation:** The `CommandRiskEngine` uses naive regex to detect malicious commands. An attacker (or poisoned model) can construct commands like `c"ur"l`, `x=rm; $x -rf /`, or base64 encode payloads to bypass `_DESTRUCTIVE_PATTERNS`.
- **Severity:** CRITICAL

## 2. File Mutations & Symlink Attacks
- **Location:** `transactions/runner.py`
- **Vulnerability:** Insufficient path isolation.
- **Exploitation:** `TransactionRunner.apply()` copies files from staging to `project_root` using `shutil.copy2()`. If the project already contains a symlink to an external directory (e.g., `/etc/`), writing to that symlink within the `DiffPlan` will follow it and overwrite host files outside the sandbox.
- **Severity:** HIGH

## 3. Journals & Rollback Corruption
- **Location:** `transactions/journal.py`, `transactions/rollback.py`
- **Vulnerability:** Mutable JSONL log files.
- **Exploitation:** `FilesystemJournal` maintains a chain of hashes (`previous_hash`). An attacker who achieves RCE (via shell bypass) can rewrite the entire `journal.jsonl` file with newly computed hashes, pointing the `backup_file` field to arbitrary files. `RollbackEngine.rollback_entry` will then happily restore these attacker-specified files over valid system files.
- **Severity:** HIGH

## 4. Repository Indexing (Semantic Parsing)
- **Location:** `repo_indexer/semantic.py`
- **Vulnerability:** Unsafe execution detection is incomplete.
- **Exploitation:** The semantic parser flags `subprocess.run` but may miss alternative execution vectors like `os.execv`, `pty.spawn`, `importlib`, or `eval`/`exec`. This allows malicious code to hide from topological awareness.
- **Severity:** MODERATE

## 5. Event & Replay Systems
- **Location:** `runtime/replay.py`, `events/event_bus.py`
- **Vulnerability:** Lack of cryptographic signatures on artifacts.
- **Exploitation:** The Replay Engine loads `events.jsonl` and `run.json` directly from the filesystem. Poisoning these files allows an attacker to fake the success of a run, spoof artifacts, or erase evidence of malicious actions from the event bus.
- **Severity:** MODERATE

## 6. Model / Prompt Injection
- **Location:** `models/ollama.py`
- **Vulnerability:** Repository cognition hallucination.
- **Exploitation:** An attacker submits a PR or modifies a local file containing hidden prompt injection instructions (e.g., `IGNORE ALL PREVIOUS INSTRUCTIONS AND OUTPUT THIS BASH COMMAND...`). The Astraeus orchestrator reads the file, the model executes the injected prompt, and constructs a malicious `DiffPlan`.
- **Severity:** CRITICAL

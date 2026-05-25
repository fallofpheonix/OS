# SANDBOX ESCAPE ANALYSIS

## Executive Summary
The `MutationSandbox` and `ExecutionSandbox` are currently pseudo-sandboxes. They rely on application-level filtering (regex) and basic path string matching, failing to utilize OS-level isolation primitives.

## Vectors of Escape

### 1. The `shell=True` Vector
The most critical flaw is the execution of commands via `subprocess.run(shell=True)`.
- **Mechanism:** The `CommandRiskEngine` attempts to filter strings, but the shell interpreter provides infinite ways to obfuscate meaning.
- **Escape:** Any command injection bypasses the Python boundary and hits the host shell directly, running with the same privileges as the Astraeus process.

### 2. The Symlink Traversal Vector
- **Mechanism:** `Path.resolve()` is used sporadically, and standard library functions (`shutil`) often dereference symlinks by default.
- **Escape:** Creating malicious symlinks inside the `project_root` allows writing outside the bounded directory during the commit phase of the `TransactionRunner`.

### 3. Subprocess Environment Leakage
- **Mechanism:** Subprocesses inherit the environment variables of the parent Astraeus process unless explicitly scrubbed.
- **Escape:** Malicious scripts executed during validation can read `os.environ` to exfiltrate API keys, cloud credentials, or the `COGNITION_LIVE_OLLAMA` flags.

## Conclusion
Astraeus is vulnerable to immediate sandbox escapes. It does not enforce a True Sandbox.

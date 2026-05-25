# SECURITY HARDENING PLAN

## Phase 1: Immediate Remediation (Days 1-7)

1. **Remove `shell=True` Everywhere**
   - Refactor `ExecutionSandbox` and `TransactionRunner._validate`.
   - Parse commands into arrays using `shlex.split`.
   - NEVER pass user- or model-generated strings directly to a shell interpreter.

2. **Harden Path Resolution**
   - Enforce `symlinks=False` (where appropriate to preserve symlinks rather than dereference them) in `shutil` operations.
   - Implement strict path resolution checks before ANY write operation: ensure `os.path.realpath` of the target remains inside `project_root`. Reject symlinks pointing outside entirely.

3. **Scrub Environment Variables**
   - Pass an explicit, minimal `env={}` to all `subprocess.run` calls to prevent credential leakage.

## Phase 2: Medium-Term Hardening (Weeks 2-4)

1. **Cryptographic Signatures for Logs**
   - Implement Ed25519 or HMAC signatures for `journal.jsonl` and `events.jsonl`.
   - The private key should be held in memory, not on disk where a standard RCE can read it.

2. **Replace Regex Risk Engine with OS Isolation**
   - The regex-based `CommandRiskEngine` is mathematically unprovable.
   - Implement Seccomp-bpf filters or AppArmor profiles for all subprocesses.
   - Run validation commands inside minimal Docker/Podman containers or gVisor, mounting the project root as read-only (or ephemeral tmpfs).

## Phase 3: Long-Term Architecture (Months 1-3)

1. **Zero Trust Orchestration**
   - Enforce strict privilege separation. The orchestrator runs as a low-privileged user; the execution sandbox runs as a separate, heavily restricted user (e.g., `nobody` or a dedicated namespace).

2. **Adversarial Prompt Filtering**
   - Implement an LLM firewall (e.g., Llama Guard) to filter incoming repository content and outgoing model commands for prompt injection and malicious intents.

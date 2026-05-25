# ZERO TRUST ARCHITECTURE

## Vision
Astraeus must operate under the assumption that the repository is hostile, the models are compromised, and the execution environment is constantly under attack.

## Core Principles

### 1. Ephemeral & Isolated Execution
No command executed by Astraeus shall run on the bare metal host. All `subprocess` calls must be directed to isolated, ephemeral namespaces (e.g., containers, MicroVMs) with strict egress network policies (Deny All by default).

### 2. Cryptographic Provenance
Every state change, journal entry, and event must be signed. Rollbacks must require cryptographic verification of the journal chain against an offline or protected public key.

### 3. Least Privilege File Mutators
The process that plans mutations (the LLM client/orchestrator) MUST NOT be the same process that writes to the filesystem. A dedicated Mutator Daemon should receive signed `DiffPlan` RPCs, validate path boundaries rigorously (with chroot-like path resolution), and perform the write as a separate user.

### 4. Semantic Intent Verification
Model outputs (commands, code patches) must be verified by a secondary, independent system (a Critic Model or deterministic linter) BEFORE execution, checking for anomalous behaviors or obfuscation techniques.

### 5. Deterministic State Verification
Replay engines must not trust `.jsonl` files natively. They must recalculate the state transitions from signed artifacts and verify the structural invariants dynamically upon load.

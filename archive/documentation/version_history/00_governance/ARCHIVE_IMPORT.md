# Archive Import Ledger

This document tracks all code and assets restored from `15_archive/` into the active PhoenixOS codebase. Restoration requires audit against current axioms and Stage A hardening standards.

## Restoration Protocol
1. **Source Mapping:** Document exact source path and commit/date.
2. **Rationalization:** Why is this being restored?
3. **Hardening Check:**
    - [ ] Determinism Audit (No non-deterministic primitives).
    - [ ] Interface Alignment (Matches `phoenix_os/contracts`).
    - [ ] Test Coverage (Must have unit tests).
    - [ ] Security Audit (No secrets, secure IPC).
4. **Replacement Plan:** Is this temporary or permanent?

---

## Import Log

### 2026-05-24: Guard Runtime Python Daemon
- **Source:** `15_archive/guard_runtime/`
- **Active Path:** `05_tools/guard_runtime_py/`
- **Reason:** Restoring Fast-Path telemetry ingestion and HMAC-SHA256 evidence signing for L1 integrity.
- **Hardening Status:**
    - [X] Determinism: Pure functional ingestion logic.
    - [X] Interface Alignment: Updated to match V2 Ledger schemas.
    - [X] Test Coverage: Passing `pytest` suite.
    - [X] Security: Verified HMAC signing and socket permissions.
- **Replacement Plan:** Permanent fast-path component; potential Go rewrite for performance if L1 overhead exceeds 50ms.

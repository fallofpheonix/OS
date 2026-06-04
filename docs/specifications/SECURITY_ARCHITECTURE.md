---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Security Architecture (v1.1)

This document defines the formal security posture of PhoenixOS, including trust boundaries, threat models, containment strategies, and security contract ownership.

## 1. Trust Boundary Map

### 1.1 Root of Trust (PRIVILEGED)
- **Subsystems:** PhoenixKernel, Constitution Engine, Boot Validator.
- **Assumptions:** Cryptographic contracts and signed Genesis state are absolute.
- **Recovery:** Fail-closed if Genesis integrity is breached.

### 1.2 Execution Substrate (TRUSTED)
- **Subsystems:** runtime, ledger, security contracts, guard implementations.
- **Assumptions:** Bit-perfect determinism; all transitions recorded in Merkle-DAG.
- **Isolation:** Single-process user-space enforcement with kernel hooks.

### 1.3 Cognitive Layer (SEMI-TRUSTED)
- **Subsystems:** cognition (Mind, Memory, Reasoning).
- **Assumptions:** Advisory ONLY. No direct write access to Ledger or Enforcement.
- **Containment:** Outputs are sanitized via advisory envelopes and contract adapters.

### 1.4 External Interface (UNTRUSTED)
- **Subsystems:** platform (dashboard, api, cli), LLM Oracle.
- **Assumptions:** All inputs are forged until validated by contract-aware guards.

---

## 2. STRIDE Threat Matrix

| Threat | Impact | Mitigation |
| :--- | :--- | :--- |
| **Spoofing** | Forged node/manifest identity | ed25519 Certificate Pinning |
| **Tampering** | Altered Ledger/Event history | Merkle-DAG Hash Chain + Quorum |
| **Repudiation**| Denied actuation / deployment | Non-repudiable Signatures |
| **Disclosure** | Telemetry / Credential leakage| RBAC + Secret Quarantining |
| **DoS** | FSM / Queue thrashing | Priority Lanes + Rate Limiting |
| **Escalation** | AI direct actuation | Guard Gate + Advisory Envelope |

---

## 3. Key Attack Vectors & Mitigations

### 3.1 Kernel Compromise
- **Goal:** Unauthorized syscall control.
- **Mitigation:** eBPF LSM pinning, signed BPF objects, and isolated loader paths.

### 3.2 Ledger Corruption
- **Goal:** Historical rewrite.
- **Mitigation:** Append-only persistence, cumulative hash checkpoints, and node-reputation scoring.

### 3.3 Replay Divergence
- **Goal:** Non-deterministic state reconstruction.
- **Mitigation:** Seed locking, monotonic sequence enforcement, and bit-perfect state hashing.

---

## 4. Safety & Liveness Properties
- **Safety:** The Ledger shall never mutate committed history.
- **Safety:** The Warden shall never execute an unapproved action class.
- **Liveness:** Every signed event shall eventually commit or move to DLQ.
- **Liveness:** Node resurrection shall eventually stabilize to the authoritative state.

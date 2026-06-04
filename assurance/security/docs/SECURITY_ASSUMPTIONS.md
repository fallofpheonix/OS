---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Phoenix Security — Assumptions & Trust Anchors

> Last verified: 2026-06-04

The Phoenix OS zero-trust architecture relies on the following foundational assumptions. If any of these assumptions are violated, the security guarantees of the system are void.

## 1. Cryptographic Trust Anchors
- **Assumption**: The root keys of authority validators are stored securely (e.g. in Hardware Security Modules or secure enclaves).
- **Rationale**: Compromising authority private keys allows attackers to sign arbitrary state transition events.

## 2. Kernel Integrity
- **Assumption**: The host Linux kernel is secure and has not been pre-compromised.
- **Rationale**: If the host kernel is compromised, eBPF probes can be bypassed or manipulated.

## 3. Hardware Safety
- **Assumption**: The hardware executing the runtime is free from CPU timing attacks or side-channel exploits (e.g. Rowhammer, Spectre).
- **Rationale**: Physical isolation is required for bare-metal deployments.

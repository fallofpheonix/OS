# DOCUMENT OWNERSHIP & TRACEABILITY MATRIX

**Status:** ACTIVE
**Implementation:** 100%
**Maintenance:** CRITICAL

## 1. Core Specification Ownership

| Document | Owner | Criticality | Last Validated Step |
| :--- | :--- | :--- | :--- |
| `docs/BFT_ASSUMPTIONS.md` | Consensus Team | CRITICAL | 6.6 |
| `docs/BFT_SPECIFICATION.md` | Consensus Team | CRITICAL | 6.6 |
| `docs/LEDGER_SPECIFICATION.md` | Ledger Team | CRITICAL | 6.6 |
| `docs/ARCHITECTURE_BASELINE.md` | Governance | HIGH | 6.6 |
| `docs/SECURITY.md` | Security Team | CRITICAL | 7.0 (Planned) |

## 2. Directory Ownership

| Path | Owner | Description |
| :--- | :--- | :--- |
| `foundation/math` | Math/Physics | Fixed-point arithmetic and deterministic geometry. |
| `foundation/ledger` | Ledger | Canonical event schema, persistence, and replayer. |
| `foundation/security` | Security | Identity management, encryption, and Keystore. |
| `game/engine` | Engine | Deterministic VM and WorldState. |
| `game/consensus` | Consensus | BFT Protocol objects and signature logic. |
| `game/multiplayer` | Networking | libp2p GossipSub transport. |
| `archive/` | Governance | Legacy artifacts and archived shadow systems. |

## 3. Implementation Traceability

| Requirement | Artifact | Status |
| :--- | :--- | :--- |
| LEDGER-SPEC-001 | `foundation/ledger/src/persist.go` | IMPLEMENTED |
| BFT-SPEC-001 | `docs/BFT_ASSUMPTIONS.md` | AUTHORITATIVE |
| BFT-SPEC-002 | `game/consensus/certificate.go` | HARDENED |
| CONSENSUS-009 | `game/engine/vm.go` | AUTHORITATIVE |
| CONSENSUS-015 | `foundation/security/identity` | IMPLEMENTED |

---
*Authorized by Phoenix Sovereign Governance*

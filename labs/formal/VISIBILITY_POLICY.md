---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# VISIBILITY POLICY

## Core Principles
- Public repos are **Products** (Reputation, proof, reusable assets).
- Private repos are **Workshops** (Internal control-plane, unstable research, security-sensitive).

## Visibility Matrix
| Type | Visibility | Rationale |
| :--- | :--- | :--- |
| Core Modules | PUBLIC | Demonstrates systems thinking and modularity. |
| SDKs | PUBLIC | Clean architectural surface area. |
| Developer Tools | PUBLIC | Infrastructure maturity and automation proof. |
| Research/Experiments | PRIVATE | Unstable, experimental, noisy. |
| Orchestration | PRIVATE | Operational nervous system; attack surface risk. |
| Forks | PRIVATE | Extraction mines; unstable and derivative. |
| Production Infra | PRIVATE | Security and IP sensitive. |

## Promotion Requirements (Private -> Public)
- README.md, LICENSE, docs/, examples/, tests/
- Working CI and clear module boundaries.
- Sanitized (no secrets, no internal IPs).
- 2+ consumers for modules.

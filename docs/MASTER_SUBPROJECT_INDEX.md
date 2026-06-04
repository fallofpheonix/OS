---
Status: Implemented
Implementation: 100%
Confidence: Proven
---
# Phoenix Master Subproject Index

> Complete inventory of all Phoenix subprojects, their maturity, ownership, and documentation links.
> Last verified: 2026-06-04

## Subproject Inventory

### Foundation Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Contracts** | `foundation/contracts/` | 9 | Mature | Architecture Team | [docs](../foundation/contracts/docs/) |
| **Events** | `foundation/events/` | 1 | Mature | Architecture Team | [docs](../foundation/events/docs/) |
| **Runtime** | `foundation/runtime/` | 173 | Mature | Core Runtime Team | [docs](../foundation/runtime/docs/) |
| **Ledger** | `foundation/ledger/` | 16 | Mature | Ledger Team | [docs](../foundation/ledger/docs/) |
| **Distributed** | `foundation/distributed/` | 12 | Partial | Core Runtime Team | [docs](../foundation/distributed/docs/) |
| **Observability** | `foundation/observability/` | 6 | Partial | Observability Team | [docs](../foundation/observability/docs/) |

### Assurance Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Validation** | `assurance/validation/` | 123 | Mature | Validation Team | [docs](../assurance/validation/docs/) |
| **Security** | `assurance/security/` | 30 | Mature | Security Team | [docs](../assurance/security/docs/) |

### Governance Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Truth** | `governance/truth/` | 6 | Early | Verification Team | [docs](../governance/truth/docs/) |
| **Arbiter** | `governance/arbiter/` | 2 | Early | Governance Team | [docs](../governance/arbiter/docs/) |

### Cognition & PhoenixMind Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Cognition (root)** | `cognition/` | 34 | Early | Intelligence Team | [docs](../cognition/docs/) |
| **Mind (Subproject)** | `cognition/mind/` | nested | Early | Intelligence Team | [docs](../cognition/mind/docs/) |
| **PhoenixMind (Core)** | `docs/phoenixmind/` | — | Research | Intelligence Team | [docs](./phoenixmind/) |

### Platform Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Platform OS** | `platform/os/` | 201 | Partial | Platform Team | [docs](../platform/os/docs/) |
| **Platform UI** | `platform/ui/` | React app | Partial | Platform Team | [docs](../platform/ui/docs/) |

### Game Layer

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Game (Specification)** | `game/` | 0 | Planned | Gameplay Team | [docs](../game/docs/) |

### Cross-Cutting

| Subproject | Path | .go Files | Maturity | Owner | Docs |
|-----------|------|-----------|----------|-------|------|
| **Contract Tests** | `contract-tests/` | ~10 | Mature | Architecture Team | [docs](../contract-tests/docs/) |

### Future / Empty

| Subproject | Path | Status | Notes |
|-----------|------|--------|-------|
| **Infrastructure** | `infrastructure/` | Empty | Storage, network, deployment — no code yet |
| **Labs** | `labs/` | Experimental | Research, crucible, formal methods |

## Maturity Definitions

| Level | Criteria |
|-------|----------|
| **Mature** | Has contracts, tests pass, build succeeds, can be independently verified |
| **Partial** | Has implementation code, may lack tests or contracts |
| **Early** | Has minimal code, architecture is conceptual |
| **Empty** | Directory exists, no implementation code |
| **Experimental** | Research or prototype code, not production-bound |

## Extraction Readiness Summary

| Subproject | Contract Stability | Dep Isolation | Circular Deps | Tests Pass | Score |
|-----------|-------------------|---------------|---------------|------------|-------|
| contracts | 95% | 100% | None | ✅ | **Ready** |
| events | 90% | 100% | None | ✅ | **Ready** |
| runtime | 85% | 90% | None | ✅ | Near-Ready |
| ledger | 90% | 95% | None | ✅ | **Ready** |
| validation | 80% | 85% | None | ✅ | Near-Ready |
| security | 80% | 85% | None | ✅ | Near-Ready |
| truth | 50% | 70% | None | ✅ | Not Ready |
| arbiter | 40% | 60% | None | — | Not Ready |
| cognition | 30% | 50% | None | ✅ | Not Ready |
| platform/os | 40% | 40% | None | ✅ | Not Ready |

---
*See [MASTER_SYSTEM_MAP.md](./MASTER_SYSTEM_MAP.md) for the visual architecture.*
*See [MASTER_DEPENDENCY_MAP.md](./MASTER_DEPENDENCY_MAP.md) for the dependency matrix.*

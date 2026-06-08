# Crucible Classification Audit

> **Authority Phase**: 4A.10 Crucible Classification Audit
> **Status**: PROPOSAL
> **Last Updated**: 2026-06-04

This document analyzes the multi-faceted nature of the Crucible module and proposes a final architectural home.

## 1. Module Composition

Crucible is currently a heterogeneous mix of three distinct asset classes:

| Category | Components | Purpose |
| :--- | :--- | :--- |
| **Verification** | `verification/`, `simulation/`, `hypothesis/`, `security/`, `PhoenixRedteam/` | Adversarial stress testing and red-teaming of the substrate. |
| **Gamification** | `game/ecs/`, `game/engines/`, `game/simulation/`, `game/context/` | The "WARDEN.EXE" logic used by the CLI for gamified auditing. |
| **Research** | `game/bitburner-src/`, `game/robocode/`, `game/mindustry/`, etc. | Massive external repository imports used as research references. |

## 2. Architectural Conflicts

- **Platform Violation**: Crucible's presence in `platform/` is an artifact of a dependency fix. It contains core-runtime logic (engines, ECS) which does not belong in the platform layer.
- **Assurance Overlap**: The verification and simulation tools are functionally identical to the `assurance/` layer's mission.
- **Game Intent**: The gamified interface is the only part that conceptually aligns with the `game/` layer.

## 3. Final Classification Proposal

Crucible should be **decomposed** rather than moved as a single unit:

### A. Move to `assurance/simulation/`
- `platform/crucible/verification`
- `platform/crucible/simulation`
- `platform/crucible/hypothesis`
- `platform/crucible/security`
- `platform/crucible/PhoenixRedteam`
*Rationale*: These are functional verification assets for the sovereign substrate.

### B. Move to `game/crucible/` (or `game/warden/`)
- `platform/crucible/game/ecs`
- `platform/crucible/game/engines`
- `platform/crucible/game/simulation`
- `platform/crucible/game/context`
*Rationale*: This is the internal "Game Engine" for the auditor interface.

### C. Move to `docs/research/assets/`
- All external repository clones (`bitburner-src`, `robocode`, etc.).
*Rationale*: These are read-only research assets that pollute the active code search space and increase build entropy.

## 4. Immediate Decision
Until decomposition is executed, the entire module is classified as **ASSURANCE**.

**New Path (PROPOSED)**: `assurance/crucible/`

---
**Next Step**: Phase 4A.11 Duplicate Systems Audit.

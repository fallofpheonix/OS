# Brain Dashboard

## Current Phase
**Module Extraction** — extracting reusable modules from completed GitHub projects before building new ones.

## Primary Focus
Module extraction: [[05_PROJECTS/REUSABLE_MODULES/fpx-runtime]] → [[05_PROJECTS/REUSABLE_MODULES/fpx-pipeline]]

## Secondary
Banking App planning (blocked on P0 module extraction)

## Research
[[05_PROJECTS/COMPLETED/Noesis/Project]] — knowledge graph primitives for future Personal Knowledge Graph project

---

## Reusable Modules (8 registered)

| Module | Priority | Source | Status |
|--------|----------|--------|--------|
| [[05_PROJECTS/REUSABLE_MODULES/fpx-runtime\|fpx-runtime]] | 🔴 P0 | AutoMation-Engine | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/fpx-pipeline\|fpx-pipeline]] | 🔴 P0 | AutoMation-Engine | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/smart-rate-limiter\|smart-rate-limiter]] | 🟡 P1 | SmartAPILimiter | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/trustlab-primitives\|trustlab-primitives]] | 🟡 P1 | SecureForg + TrustLab | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/fpx-observability\|fpx-observability]] | 🟢 P2 | UDIE | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/agentskill\|agentskill]] | 🟢 P2 | agentskill | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/fraud-detector-ml\|fraud-detector-ml]] | 🟢 P2 | New (Banking App) | NOT_STARTED |
| [[05_PROJECTS/REUSABLE_MODULES/pfi-ingestion\|pfi-ingestion]] | 🔵 P3 | AI-PFI | NOT_STARTED |

---

## Weak Areas (review weekly)
- Distributed consensus
- Event sourcing implementation details
- C FFI / Python bindings
- Networking internals

---

## Quick Links
[[06_FAILURE_LIBRARY]] | [[04_ENGINEERING/debugging-patterns]] | [[08_REFERENCE/commands]]

---

## Completed Projects (21)
See [[10_META/dashboards/Project_Dashboard]]

---

## Cancelled (duplication detected)
- ~~brute-force-guard~~ → use [[05_PROJECTS/REUSABLE_MODULES/smart-rate-limiter]]
- ~~Game Physics Engine~~ → extend [[05_PROJECTS/COMPLETED/ParticleStimulator/Project]]
- ~~Fraud Detection standalone~~ → build as [[05_PROJECTS/REUSABLE_MODULES/fraud-detector-ml]] inside Banking App

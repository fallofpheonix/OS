# Project: agentskill

## One-Liner
agentskill

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/agentskill`

## Ports
- API: N/A
- DB: N/A

## Database
N/A

## Run Command
N/A - historical project overview

## Dependencies On Other Projects
None

## What I Deliver To Others
None

## Links
- [[03_CORE_KNOWLEDGE/ai-ml/AI]]
- [[04_ENGINEERING/architecture-patterns/Software-Engineering]]
- [[04_ENGINEERING/system-design/System Design]]
- [[03_CORE_KNOWLEDGE/ai-ml/Machine Learning]]
- [[04_ENGINEERING/architecture-patterns/Frontend Architecture]]
- [[03_CORE_KNOWLEDGE/security/Security]]
- [[Decisions]]
- [[Mistakes]]

## Current Blockers
None

## Last Worked On
2026-05-12

## Original Overview


**Repository:** [github.com/fallofpheonix/agentskill](https://github.com/fallofpheonix/agentskill)  
**Language:** Python | **Created:** 2026-04-06

---

## Project Summary

A developer toolchain for deterministic Agentfile parsing, validation, and generation. Includes a tri-engine contract system, CI/CD validation scripts, and security gate tooling. Classified as a toolchain, not a deployable service.

## Modules

| Path | Role |
|---|---|
| `backend/agentman` | Deterministic Agentfile parser, validator, and generator |
| `system/tri-engine` | Tri-engine contract and enforced stage schema |
| `ci_cd` | Local validation entrypoints |
| `security` | Security gate scripts (pip-audit, secret scanning) |
| `docs` | Full documentation set |
| `tests` | Pytest suite + repo-integrity assertions |

## Execution Guarantees

- **Build:** Deterministic bundle generation with `control-plane.json` + `stages.yaml`
- **Test:** ≥70% coverage for core package
- **Security:** pip-audit + regex-based secret scanning
- **CI/CD:** GitHub Actions on pull_request and push to main

## Skills Demonstrated

`Python`, `CLI Tooling`, `CI/CD Pipelines`, `GitHub Actions`, `Security Auditing`, `pip-audit`, `Secret Scanning`, `Deterministic Builds`, `Test Coverage`, `Developer Toolchain`

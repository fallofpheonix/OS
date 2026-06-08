---
project-id: {{slug}}
repo: https://github.com/fallofpheonix/{{repo-name}}
status: ACTIVE | REFACTOR | PROMOTED | ARCHIVED
audit-verdict: "{{verdict}}"
language: {{language}}
started: {{YYYY-MM}}
last-commit: {{YYYY-MM-DD}}
tags: [project, existing, {{domain-tag}}]
---

# {{Project Name}}

## What It Is
One paragraph. What does this project do, who is it for, and what problem does it solve?

## Current State
- What is working:
- What is broken or incomplete:
- Audit-identified defects:
- Production readiness: NOT READY / PARTIAL / READY

## Architecture
Current vs target architecture. Document the gap.

## Tech Stack
Languages, frameworks, libraries, databases, infrastructure.

## Modules Used
| Module | Status | Notes |
|--------|--------|-------|

## Retroactive ADRs
- [[04_ENGINEERING/decision-logs/ADR-NNN-slug]]
Minimum: 3 per project.

## Known Failure Modes
- [[06_FAILURE_LIBRARY/YYYY-MM-project-slug]]
Minimum: 2 per project.

## Committed Artifact Violations
List binaries, generated files, data artifacts in repo.

## Refactor Actions
- [ ] Action item from audit

## Spec Kit Status
- specify init: NOT YET RUN
- Next feature to spec:
- When ready: specify init . --integration claude

## Linked Concepts
- [[03_CORE_KNOWLEDGE/]]
Minimum: 2 links.

## Promotion Criteria
- [ ] All refactor actions complete
- [ ] CI passing
- [ ] Architecture diagram exists
- [ ] ADRs written
- [ ] README follows docs contract

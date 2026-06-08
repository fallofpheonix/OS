# Project: web-application-firewall

## One-Liner
Web Application Firewall.

## Status
COMPLETED

## Repo
`~/engineering/workspace/active/web-application-firewall`

## Ports
- API: localhost:{port}
- DB: localhost:{db_port}

## Database
{db_name}

## Run Command
`cd ~/engineering/workspace/active/web-application-firewall && docker compose up -d`

## Modules Used
- [[08_MODULES/xss-scanner]]
- [[08_MODULES/sqli-detector]]
- [[08_MODULES/waf-core]] (produced)

## Spec Kit

### Constitution Summary
One paragraph: the governing principles for this project (copy key points from `.specify/memory/constitution.md`).

### Features Shipped

| # | Feature Branch | Spec | Status | ADR Extracted | Failures Extracted |
|---|---------------|------|--------|---------------|-------------------|
| 001 | {branch-name} | `.specify/specs/001-{name}/spec.md` | SHIPPED | [[ADR: {decision}]] | [[Failure: {title}]] |

### Active Feature
```text
Feature: {name}
Branch: {branch}
Current phase: IDEA
Blocked on: none
```

## Dependencies On Other Projects
None

## What I Deliver To Others
None

## Brain Links
- [[Architecture]]
- [[03_CORE_KNOWLEDGE/HTTP Filtering]]
- [[03_CORE_KNOWLEDGE/ModSecurity Rules]]
- [[03_CORE_KNOWLEDGE/OWASP Top 10]]

## Current Blockers
None

## Last Worked On
2026-05-12


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

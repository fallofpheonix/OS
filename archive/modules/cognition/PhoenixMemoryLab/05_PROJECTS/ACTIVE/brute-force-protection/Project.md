# Project: brute-force-protection

## One-Liner
Brute force guard.

## Status
COMPLETED

## Repo
`~/engineering/workspace/active/brute-force-protection`

## Ports
- API: localhost:{port}
- DB: localhost:{db_port}

## Database
{db_name}

## Run Command
`cd ~/engineering/workspace/active/brute-force-protection && docker compose up -d`

## Modules Used
- [[08_MODULES/brute-force-guard]] (produced)

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
- [[03_CORE_KNOWLEDGE/Rate Limiting]]
- [[03_CORE_KNOWLEDGE/Account Lockout]]
- [[03_CORE_KNOWLEDGE/IP Blocking]]
- [[03_CORE_KNOWLEDGE/Redis]]

## Current Blockers
None

## Last Worked On
2026-05-12


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

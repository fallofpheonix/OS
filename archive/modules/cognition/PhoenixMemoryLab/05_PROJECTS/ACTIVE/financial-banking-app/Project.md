# Project: ledger-core

## One-Liner
Financial Banking App.

## Status
COMPLETED

## Repo
`~/engineering/workspace/active/ledger-core`

## Ports
- API: localhost:{port}
- DB: localhost:{db_port}

## Database
{db_name}

## Run Command
`cd ~/engineering/workspace/active/ledger-core && docker compose up -d`

## Modules Used
- [[08_MODULES/auth-bcrypt]]
- [[08_MODULES/totp-2fa]]
- [[08_MODULES/jwt-analyzer]]
- [[08_MODULES/rbac-middleware]]
- [[08_MODULES/brute-force-guard]]
- [[08_MODULES/waf-core]]
- [[08_MODULES/impossible-travel]]
- [[08_MODULES/fraud-detector-ml]]

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
- [[03_CORE_KNOWLEDGE/Next.js App Router]]
- [[03_CORE_KNOWLEDGE/Plaid OAuth]]
- [[03_CORE_KNOWLEDGE/Database Transactions]]
- [[03_CORE_KNOWLEDGE/ACID Properties]]

## Current Blockers
None

## Last Worked On
2026-05-12


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

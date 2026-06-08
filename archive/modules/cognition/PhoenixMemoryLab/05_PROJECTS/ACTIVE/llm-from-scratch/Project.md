# Project: llm-from-scratch

## One-Liner
LLM from scratch.

## Status
IDEA / PLANNING / ACTIVE / BLOCKED / MAINTENANCE / COMPLETED / ARCHIVED / FAILED

## Repo
`~/engineering/workspace/active/llm-from-scratch`

## Ports
- API: localhost:{port}
- DB: localhost:{db_port}

## Database
{db_name}

## Run Command
`cd ~/engineering/workspace/active/llm-from-scratch && docker compose up -d`

## Modules Used
None

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
- [[03_CORE_KNOWLEDGE/Attention Mechanism]]
- [[03_CORE_KNOWLEDGE/Tokenization (BPE)]]
- [[03_CORE_KNOWLEDGE/Positional Encoding]]
- [[03_CORE_KNOWLEDGE/Layer Normalization]]
- [[03_CORE_KNOWLEDGE/Autoregressive Generation]]

## Current Blockers
None

## Last Worked On
2026-05-12


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

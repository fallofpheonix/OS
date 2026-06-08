# Project: intrusion-detection-system

## One-Liner
IDS.

## Status
IDEA / PLANNING / ACTIVE / BLOCKED / MAINTENANCE / COMPLETED / ARCHIVED / FAILED

## Repo
`~/engineering/workspace/active/intrusion-detection-system`

## Ports
- API: localhost:{port}
- DB: localhost:{db_port}

## Database
{db_name}

## Run Command
`cd ~/engineering/workspace/active/intrusion-detection-system && docker compose up -d`

## Modules Used
- [[08_MODULES/packet-sniffer]]
- [[08_MODULES/malware-sig-scanner]]
- [[08_MODULES/wifi-monitor]]

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
- [[03_CORE_KNOWLEDGE/Anomaly Detection]]
- [[03_CORE_KNOWLEDGE/Signature Matching]]
- [[03_CORE_KNOWLEDGE/Snort Rules]]

## Current Blockers
None

## Last Worked On
2026-05-12


## Independence Checklist
- [ ] Unique port assigned?
- [ ] Dedicated DB or schema?
- [ ] Unique `.env` injected?
- [ ] Isolated Docker network?

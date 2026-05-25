# Operations: Incident Response Workflow

Use this guide to triage, escalate, and resolve security or performance incidents.

## 1. Severity Levels

| Severity | Definition | Target Resolution |
| :--- | :--- | :--- |
| **SEV 1** | Critical security breach or FSM loop lockup | $< 10$ minutes |
| **SEV 2** | Ledger divergence or high event drop rate | $< 30$ minutes |
| **SEV 3** | Non-blocking metric delay or logging warnings | $< 2$ hours |

## 2. Escalation Workflow
1. **Detect:** Automated alerts from metrics or dashboards.
2. **Isolate:** Stop target services or trigger manual container locks.
3. **Replay:** Extract telemetry logs and replay the incident offline in the chaos lab.
4. **Fix:** Apply hotfix and redeploy under strict version verification rules.

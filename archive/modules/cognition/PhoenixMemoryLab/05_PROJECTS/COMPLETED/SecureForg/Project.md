# Project: SecureForg

## One-Liner
SecureForg (Sentinel-Scribe)

## Status
COMPLETED

## Repo
`~/engineering/workspace/archived/SecureForg`

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


**Repository:** [github.com/fallofpheonix/SecureForg](https://github.com/fallofpheonix/SecureForg)  
**Language:** Python | **Created:** 2026-04-05

---

## Project Summary

Runtime vulnerability detection tool that goes beyond static analysis. Verifies code by executing a safe baseline and fixed attack payloads, then comparing behavior to detect runtime-triggerable flaws.

## How It Works

1. Executes target code with benign inputs → baseline behavior
2. Runs fixed attack payloads (SQL injection, command injection, code injection)
3. Compares runtime behavior between baseline and attack
4. Reports vulnerability status with payload-level detail

## Comparison vs Static Analysis

| Tool | Detects runtime exploit |
|---|---|
| Bandit (static) | No |
| Sentinel-Scribe (this) | Yes |

## Supported Vulnerability Types

- ✅ Injection vulnerabilities
- ✅ Runtime-triggerable flaws
- ❌ Memory corruption
- ❌ Multi-step exploits
- ❌ Network-based attacks

## Structure

```
core/       — executor, detector, validator
analysis/   — payloads, AST analyzer
app/        — CLI and UI interfaces
examples/   — sql_vuln, cmd_vuln, code_vuln, fixed_sql
```

## Skills Demonstrated

`Python`, `Security Analysis`, `Runtime Vulnerability Detection`, `AST Analysis`, `Sandboxed Execution`, `Payload Testing`, `CLI Development`

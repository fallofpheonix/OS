# PHOENIX ARCHIVE INDEX

**Status:** ACTIVE
**Implementation:** 100%
**Maintenance:** REQUIRED

## 1. Overview
This directory contains legacy prototypes, experimental artifacts, and deprecated system layers. 
Code in this directory IS NOT included in production builds and should be treated as read-only reference.

## 2. Archive Categories

### labs/ [RESEARCH]
*   **g0dm0d3**: Early AI/LLM integration experiments. (Status: ABANDONED)
*   **nucleus**: Initial sovereign kernel prototype. (Status: PROTOTYPE)
*   **terminus**: Early terminal interface experiments. (Status: ABANDONED)
*   **formal**: Early formal verification attempts. (Status: ACTIVE_RESEARCH)

### platform/ [PROTOTYPES]
*   **crucible**: Early execution environment. Superseded by `foundation/runtime`.
*   **os**: Legacy OS integration layers. Superseded by `platform/os`. (Wait, I moved `platform/os` here, checking...)

### data/ [LEGACY]
*   **warden_slice.jsonl**: Early verification trace data.

### tests/ [DEPRECATED]
*   **chaos.test**: Early chaos engineering experiments.
*   **invariants.test**: Initial invariant checks.
*   **PhoenixGuard.test**: Early guardian runtime tests.
*   **replay.test**: Initial replay logic verification.

## 3. Reference Policy
*   **DO NOT** import packages from `archive/`.
*   **DO NOT** run tests from `archive/`.
*   **OWNER**: Governance/Security Team.

---
*Authorized by Phoenix Sovereign Governance*

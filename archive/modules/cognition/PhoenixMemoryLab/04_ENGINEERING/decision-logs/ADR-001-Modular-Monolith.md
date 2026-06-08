# ADR-001: Modular Monolith Architecture for AI Assistant

**Date:** 2026-05-12  
**Status:** ACCEPTED  
**Authors:** Engineering System

---

## Context

The AI Assistant project is the primary architecture laboratory and will likely run for 3+ months, grow into multiple systems, and produce reusable modules.

Initial implementation must choose between:
1. Microservices (distributed from day 1)
2. Monolith (single binary, monolithic structure)
3. Modular Monolith (single repo, clean internal boundaries)

Each choice has fundamentally different long-term consequences for:
- Iteration speed
- Debugging complexity
- Reusable module extraction
- Testability
- Operational overhead

---

## Decision

**MODULAR MONOLITH.**

Single repository, single Python codebase, but with strict subsystem boundaries:

```
interfaces/      → CLI only
  ↓
core/           → Orchestration only
  ↓
runtime/        → Execution only
  ↓
storage/        → Persistence only
  ↓
infrastructure/ → Shared low-level
```

No circular dependencies.
No cross-subsystem leakage.
But runs as one process.

---

## Rationale

### Why NOT Microservices
**Microservices are premature.**

- Don't know final boundaries yet
- Every service requires:
  - Separate deployment
  - Separate testing
  - Separate debugging
  - IPC overhead
  - Operational complexity
- Adding deployment complexity before proving the architecture works is wrong

**True story:** Most failed early-stage projects over-engineer distribution.

### Why NOT Simple Monolith
**Simple monolith (no boundaries) creates:**

- Everything imports everything
- Circular dependencies appear
- Testability collapses (can't mock subsystems)
- Extraction impossible later
- Debugging chaos

### Why Modular Monolith

**Combines benefits:**
- Single process → easy local development
- Clean boundaries → extraction-ready
- Testability → can mock each layer
- No deployment overhead → iterate fast
- Reusable module extraction → clear APIs

**Forces architectural discipline** without operational complexity.

**Proven pattern** at companies that scale:
- Early: modular monolith
- Mature: extract services where pressure exists
- Not: guess services upfront

---

## Consequences

### POSITIVE

1. **Fast Iteration**
   - Local development simple
   - No IPC debugging
   - Single deployment unit
   - Rapid feedback loop

2. **Clean Extraction**
   - If logger repeats in 2 projects → extract to module
   - If tool-runtime becomes shared → extract to package
   - If control-plane logic is reusable → extract to library
   - Clear boundaries enable this

3. **Debuggability**
   - Single codebase
   - No distributed debugging
   - Stack traces make sense
   - Easy to trace execution flow

4. **Testability**
   - Can mock runtime for core tests
   - Can mock storage for control-plane tests
   - Can test CLI without actual execution
   - Clean interfaces support this

### NEGATIVE

1. **Will Outgrow This**
   - Eventually distributed execution is needed
   - But THEN, not now
   - Decision reversible after patterns clear

2. **Requires Discipline**
   - Clean boundaries don't enforce themselves
   - Must actively prevent leakage
   - Code reviews must enforce
   - Circular dependencies can creep

3. **Single Point of Failure**
   - One subsystem crash affects whole system
   - Mitigated by: good error handling, timeout isolation
   - Not a blocker for Phase 1-2

---

## Alternatives Considered

### Alternative 1: Microservices from Start
**Why rejected:**

Don't know subsystem boundaries yet. Premature distribution adds:
- Async/IPC complexity
- Deployment overhead
- Operational monitoring requirements
- Debugging complexity
- Testing harness complexity

Better to prove architecture works first, then distribute where needed.

### Alternative 2: Simple Monolith (No Boundaries)
**Why rejected:**

Works initially but:
- Everything couples to everything
- Extraction becomes painful refactoring
- Testing becomes integration-only
- Boundaries become implicit (bad)

Forces architecture decisions immediately instead of discovering them.

---

## Implementation Implications

### What This Means for Code

1. **Directory structure enforces boundaries**
   ```
   forge-agent/
   ├── interfaces/
   │   └── cli/        # User I/O only
   ├── core/
   │   └── agent/      # Decision logic only
   ├── runtime/
   │   └── shell/      # Execution only
   ├── storage/        # Persistence only
   └── infrastructure/ # Low-level shared
   ```

2. **Import rules (enforced by code review)**
   - `interfaces` → can import `core`
   - `core` → can import `runtime`, `storage`, `infrastructure`
   - `runtime` → can import `infrastructure` ONLY
   - `storage` → can import `infrastructure` ONLY
   - NO CIRCULAR IMPORTS

3. **Test strategy**
   - Unit: test each subsystem independently (mock dependencies)
   - Integration: test boundaries between layers
   - E2E: test full flow (CLI → execution → response)

---

## Migration Path

If in Phase 3-4 we need distribution:

1. Extract runner as separate service (clear boundary)
2. Other subsystems stay in monolith
3. Use simple message queue (Redis, RabbitMQ) for IPC
4. Replace function calls with async messages

But this only happens if pressure exists.

---

## Related Decisions

- [[ADR-002-Python-First]] — Enables fast modular iteration
- [[ADR-003-Foundation-First]] — Architecture proves itself before adding AI

---

## Review Schedule

- **Week 1:** Validate that boundaries work in practice
- **Week 2:** Check for coupling violations
- **Month 1:** Decide if extraction-ready
- **Month 3:** Revisit if distribution pressure appears

---

## Status History
- 2026-05-12: ACCEPTED (initial architecture decision)

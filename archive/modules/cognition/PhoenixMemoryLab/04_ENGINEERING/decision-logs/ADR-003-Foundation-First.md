# ADR-003: Foundation First — No AI Until Architecture Validates

**Date:** 2026-05-12  
**Status:** ACCEPTED  
**Authors:** Engineering System

---

## Context

The AI Assistant project can proceed down two paths:

**Path A: Build AI first**
- Add local LLM immediately (day 1-2)
- Build reasoning loops
- Optimize prompts
- Hope architecture works

**Path B: Prove architecture first**
- Build basic CLI → Orchestrator → Runtime → Response flow
- Validate clean boundaries
- Add AI once foundation is proven
- Then reasoning can assume solid foundation

Path B is better.

---

## Decision

**Build foundation BEFORE adding LLM.**

Timeline:
- **Days 1-6:** CLI, orchestrator, runtime executor, file/git tools, logging
- **Day 7+:** First working system proven
- **Week 2:** Add local LLM integration
- **Week 3+:** Add memory and reasoning

---

## Rationale

### Why This Matters

**Most AI systems fail because:**

1. Weak control-plane → agent decisions are unreliable
2. Unsafe execution → bugs disappear into LLM hallucinations
3. No logging → can't debug what went wrong (LLM or system?)
4. Bad state management → memory pollutes reasoning
5. Unmeasurable architecture → can't improve

**Adding LLM too early masks all these problems.**

The LLM becomes a black box that hides architectural failures.

### What Foundation First Gives You

1. **Clear Debugging**
   - Is this a shell execution bug or LLM reasoning bug?
   - Can reproduce without AI involved
   - Isolate actual failures

2. **Provable Architecture**
   - If `run ls` works, control-plane works
   - If file tools work, isolation works
   - If logging works, observability works
   - AI adds reasoning, not magic

3. **Better Prompts Later**
   - Once you know what the system can actually do
   - Once you know failure modes
   - Write prompts that work with real constraints

4. **Measurable Impact**
   - Know baseline behavior (non-AI)
   - Measure AI impact separately
   - Debug independently

---

## Consequences

### POSITIVE

1. **Architecture Proven Before Reasoning**
   - If system works with dumb executor, it'll work with smart one
   - Failures are addressable, not AI magic

2. **Testing Is Simple**
   - Test orchestrator without AI
   - Test runtime without reasoning
   - Test tools without intelligence
   - Add AI and re-test

3. **Debugging Is Clear**
   - Execute `run ls` manually
   - Execute via system
   - Compare outputs
   - No "LLM chose wrong tool" ambiguity

4. **Modular Extraction Clear**
   - Tool runtime boundary: crystal clear
   - Orchestrator boundary: proven
   - Memory boundary: understood
   - Can extract safely

### NEGATIVE

1. **Slower Time to "AI Features"**
   - First week is plumbing
   - Second week adds LLM
   - Trade: proven system vs. fast hype

2. **More Code Before AI**
   - CLI, executor, logger, orchestrator
   - But all of it reusable
   - Not wasted work

---

## Alternatives Considered

### Alternative 1: LLM from Day 1
**Why rejected:**

- Masks architectural problems
- Can't tell if failure is control-plane or inference
- Makes debugging impossible
- Prevents module extraction

Example: "Tool executor times out" → could be:
- Actually slow shell command (architectural issue)
- LLM timing out requests (LLM issue)
- Message queue broken (not tested yet)
→ Can't tell which without foundation proven

### Alternative 2: Parallel Track (Foundation + LLM)
**Why rejected:**

Splitting focus:
- Debugging gets confused (which system failed?)
- Priorities conflict
- Architecture gets rushed (to add AI faster)

Better to finish one, then the other.

---

## Implementation Plan

### Phase 1: Foundation (Days 1-7)

**Goal:** User input → validated execution → structured response

Files:
```
interfaces/cli/main.py          # Parse user input
core/agent/orchestrator.py      # Route requests
runtime/shell/executor.py       # Execute safely
runtime/file_ops.py             # File operations
infrastructure/config/
infrastructure/logging/
tests/test_executor.py          # Verify core flow
```

Success: `python -m interfaces.cli run ls` works reliably

### Phase 1b: Validation (Days 7-10)

**Goal:** Prove architecture holds

Tasks:
- [ ] Test file operations
- [ ] Test git operations
- [ ] Test timeout handling
- [ ] Test error recovery
- [ ] Write first ADRs
- [ ] Write first failure notes

Success: Multiple tools working, failures captured in learning system

### Phase 2: LLM Integration (Days 11-21)

**Goal:** Add reasoning on proven foundation

Changes:
- Add local LLM wrapper (core/llm/client.py)
- Add memory system (storage/memory/)
- Update orchestrator to use LLM for planning
- Add prompt templates

Now the system:
- Uses real reasoning
- Has proven control-plane underneath
- Has observable failure modes
- Can be debugged independently

### Phase 2b: Expansion (Days 22+)

Add:
- Advanced memory patterns
- Tool chains (tool A → tool B)
- Error recovery with reasoning
- Performance optimization

---

## Decision Gates

### Gate 1: After Day 3
**Question:** Does basic CLI routing work?
**If NO:** Debug architecture before proceeding
**If YES:** Continue to Day 6

### Gate 2: After Day 6
**Question:** Do tools work independently?
**If NO:** Fix tool isolation before LLM
**If YES:** Ready for LLM integration

### Gate 3: After Day 10
**Question:** Can you run realistic tasks?
**If NO:** Fix execution model
**If YES:** Safe to add LLM

---

## Related Decisions

- [[ADR-001-Modular-Monolith]] — Enables clean layers for this
- [[ADR-002-Python-First]] — Python foundation fastest to build

---

## Long-term Vision

This approach compounds:

**Month 1:** Proven foundation + reasoning
**Month 2:** Stable system + first modules extracted
**Month 3:** Second project reuses modules, validates extraction
**Month 4+:** Ecosystem stabilizes

vs.

**Month 1:** AI hype, architectural chaos
**Month 2:** Debugging nightmares
**Month 3:** Rewrite from scratch

---

## Review Schedule

- **Day 4:** Checkpoint on basic architecture
- **Day 7:** Checkpoint on foundation completion
- **Day 10:** Gate before LLM addition
- **Week 2:** After LLM, validate integration doesn't break foundation

---

## Status History
- 2026-05-12: ACCEPTED (foundation-first strategy)

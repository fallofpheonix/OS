# Extraction Log

**Purpose:** Track modules extracted from projects to prevent duplication and maintain ecosystem coherence.

**Prevents:**
- Duplicated logic across projects
- Premature abstraction before patterns stabilize
- Lost institutional knowledge about what went into modules
- Tight coupling between projects

**Rules:**
1. Only extract after pattern repeats in 2+ projects
2. Document the extraction at the time it happens
3. Update consumers list when new projects adopt the module
4. Mark API stability level clearly
5. Link back to source projects for context

---

## Extraction Candidates (Phase 1 Monitoring)
[Watching for these patterns to emerge — NOT YET EXTRACTED]

### From AI Assistant Project
- **Logger abstraction:** Repeated in every subsystem
  - Status: CANDIDATE (appears in infrastructure/, runtime/, core/)
  - Wait for: Second project reuse attempt
  
- **Config system:** Settings pattern
  - Status: CANDIDATE
  - Wait for: Validation that pattern is stable

- **Shell executor:** Runtime isolation
  - Status: CANDIDATE
  - Wait for: Second system needing execution abstraction

### From [Future Secondary Project]
[To be populated as Phase 2 begins]

---

## Extracted Modules
[Will populate in Phase 3 when extraction pressure becomes clear]

### [Module Name]
- **Extracted:** [Date]
- **Source:** [[Project Name]]
- **Why:** [Pattern appeared in X projects]
- **Consumers:** [[Project A]], [[Project B]]
- **Location:** `modules/[name]/`
- **Status:** EXPERIMENTAL | EVOLVING | STABLE
- **Last Updated:** [Date]

---

## Reuse Validation (Phase 4+)
[Will populate as second/third projects consume extracted modules]

### [Module Name]
- **Initial Consumer:** [[Project A]]
- **Second Consumer:** [[Project B]]
- **Lessons Learned:**
  - What was right about the abstraction?
  - What needed adjustment?
  - Was the API boundary correct?

---

## Ecosystem Health
- **Total Modules:** 0 (Phase 0)
- **Extraction Candidates:** 3
- **Projects Producing Modules:** 0
- **Module Reuse Rate:** N/A (Phase 0)

---

## Next Review
- Check extraction candidates after AI Assistant reaches Phase 2
- Evaluate for first extraction after secondary project starts

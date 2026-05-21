# Stage 0: Milestone Tracker

## Progress Overview

**Total Milestones:** 24 mandatory + 8 optional
**Total Time (Mandatory):** 175h across 9 weeks
**Start Date:** [FILL IN]
**Target Completion:** [FILL IN + 9 weeks]

---

## Tracking Template

For each milestone, update:
- **Status:** Pending → In Progress → Completed → Verified
- **Start Date:** When work began
- **Est. Finish:** Planned completion
- **Act. Finish:** When actually finished
- **Hours Spent:** Actual time invested
- **Variance:** Estimate vs. Actual (flag if >20%)
- **Notes:** Issues, resource recommendations, insights

---

## Week 1: Computer Architecture (25h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| A1: CPU Pipeline & ISA | Pending | — | — | — | 8h | — |
| A2: Memory Hierarchy | Pending | — | — | — | 8h | — |
| A3: Interrupts & Exceptions | Pending | — | — | — | 6h | — |
| A4: Registers & Addressing | Pending | — | — | — | 3h | — |
| **Week 1 Total** | — | — | — | — | **25h** | — |

---

## Week 2: Assembly (20h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| B1: Assembly Syntax & Instructions | Pending | — | — | — | 8h | — |
| B2: Calling Conventions & Stack | Pending | — | — | — | 7h | — |
| B3: ELF Binary Format | Pending | — | — | — | 5h | — |
| **Week 2 Total** | — | — | — | — | **20h** | — |

---

## Week 3: OS Concepts (25h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| C1: Process Model | Pending | — | — | — | 8h | — |
| C2: Thread Model | Pending | — | — | — | 7h | — |
| C3: Scheduling Algorithms (optional) | Pending | — | — | — | 6h | — |
| C4: System Calls & Privilege Transition | Pending | — | — | — | 4h | — |
| **Week 3 Total** | — | — | — | — | **25h** | — |

---

## Week 4: Memory Management (20h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| D1: Virtual Memory Concepts | Pending | — | — | — | 8h | — |
| D2: Paging & Page Faults | Pending | — | — | — | 7h | — |
| D3: Heap & Stack Layout | Pending | — | — | — | 5h | — |
| **Week 4 Total** | — | — | — | — | **20h** | — |

---

## Week 5: Filesystems (15h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| E1: Inode Structure & Operations | Pending | — | — | — | 7h | — |
| E2: VFS & Filesystem Types (optional) | Pending | — | — | — | 5h | — |
| E3: Disk I/O & Page Cache (optional) | Pending | — | — | — | 3h | — |
| **Week 5 Total** | — | — | — | — | **15h** | — |

---

## Weeks 6-7: Linux Internals (35h) — CRITICAL BOTTLENECK

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| F1: Boot Sequence | Pending | — | — | — | 10h | — |
| F2: Kernel Architecture & Subsystems | Pending | — | — | — | 8h | — |
| F3: System Call Path | Pending | — | — | — | 7h | — |
| F4: Loadable Kernel Modules (LKM) | Pending | — | — | — | 6h | — |
| F5: Device Drivers & Device Model (optional) | Pending | — | — | — | 4h | — |
| **Weeks 6-7 Total** | — | — | — | — | **35h** | — |

---

## Week 8: Networking (20h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| G1: TCP/IP Stack Layers | Pending | — | — | — | 8h | — |
| G2: Sockets API & Protocol Basics | Pending | — | — | — | 8h | — |
| G3: DNS & Routing Basics (optional) | Pending | — | — | — | 4h | — |
| **Week 8 Total** | — | — | — | — | **20h** | — |

---

## Week 9: Security Fundamentals (20h)

| Milestone | Status | Start | Est. | Act. | Hours | Notes |
|-----------|--------|-------|------|------|-------|-------|
| H1: CIA Triad & Threat Modeling | Pending | — | — | — | 7h | — |
| H2: Access Control & Authentication | Pending | — | — | — | 6h | — |
| H3: Logging & Audit (optional) | Pending | — | — | — | 4h | — |
| H4: Cryptography Basics (optional) | Pending | — | — | — | 3h | — |
| **Week 9 Total** | — | — | — | — | **20h** | — |

---

## Development Environment (Distributed, 10h total)

| Milestone | Status | Start | Est. | Act. | Hours | Week | Notes |
|-----------|--------|-------|------|------|-------|------|-------|
| I1: Compiler & Build System | Pending | — | — | — | 4h | 1-2 | — |
| I2: Debugging & Profiling | Pending | — | — | — | 4h | 3-5 | — |
| I3: Version Control & Collaboration (optional) | Pending | — | — | — | 2h | 1 | — |

---

## Mandatory Gate Milestones (Must Complete)

```
[W1] A1: CPU Pipeline & ISA                    ✓ GATE
[W1] A2: Memory Hierarchy                      ✓ GATE
[W1] A3: Interrupts & Exceptions               ✓ GATE
[W1] A4: Registers & Addressing                ✓ GATE
[W2] B1: Assembly Syntax & Instructions        ✓ GATE
[W2] B2: Calling Conventions & Stack           ✓ GATE
[W2] B3: ELF Binary Format                     ✓ GATE
[W3] C1: Process Model                         ✓ GATE
[W3] C2: Thread Model                          ✓ GATE
[W3] C4: System Calls & Privilege Transition   ✓ GATE
[W4] D1: Virtual Memory Concepts               ✓ GATE
[W4] D2: Paging & Page Faults                  ✓ GATE
[W4] D3: Heap & Stack Layout                   ✓ GATE
[W5] E1: Inode Structure & Operations          ✓ GATE
[W6] F1: Boot Sequence                         ✓ GATE
[W6] F2: Kernel Architecture & Subsystems      ✓ GATE
[W7] F3: System Call Path                      ✓ GATE
[W7] F4: Loadable Kernel Modules               ✓ GATE
[W8] G1: TCP/IP Stack Layers                   ✓ GATE
[W8] G2: Sockets API & Protocol Basics         ✓ GATE
[W9] H1: CIA Triad & Threat Modeling           ✓ GATE
[W9] H2: Access Control & Authentication       ✓ GATE
[W1-9] I1: Compiler & Build System             ✓ GATE
[W1-9] I2: Debugging & Profiling               ✓ GATE
```

**Total: 24 mandatory gates**

---

## Optional Milestones (Accelerators for Stage 1)

```
[W3] C3: Scheduling Algorithms
[W5] E2: VFS & Filesystem Types
[W5] E3: Disk I/O & Page Cache
[W6] F5: Device Drivers & Device Model
[W8] G3: DNS & Routing Basics
[W9] H3: Logging & Audit
[W9] H4: Cryptography Basics
[W1] I3: Version Control & Collaboration
```

**Total: 8 optional milestones (~35h if completed)**

---

## Completion Criteria

### ✓ Milestone Verified
- Checkpoint completed (explanation/code/tool/analysis)
- Verification method confirmed
- Hours logged
- Variance < 20%

### ⚠ Milestone Blocked
- Dependency not complete
- Resource unavailable
- Unexpected complexity
- **Action:** Update dependency_matrix.md, document blocker

### ✗ Milestone Failed
- Checkpoint not reproducible
- Verification inconclusive
- Hours exceed 120% of estimate
- **Action:** Redo milestone, update estimate

---

## Variance Analysis

### When to Alert (>20% over estimate)

| Milestone | Estimate | Alert Threshold |
|-----------|----------|-----------------|
| A1 | 8h | > 9.6h |
| A2 | 8h | > 9.6h |
| A3 | 6h | > 7.2h |
| A4 | 3h | > 3.6h |
| B1 | 8h | > 9.6h |
| B2 | 7h | > 8.4h |
| B3 | 5h | > 6h |
| C1 | 8h | > 9.6h |
| C2 | 7h | > 8.4h |
| C4 | 4h | > 4.8h |
| D1 | 8h | > 9.6h |
| D2 | 7h | > 8.4h |
| D3 | 5h | > 6h |
| E1 | 7h | > 8.4h |
| F1 | 10h | > 12h |
| F2 | 8h | > 9.6h |
| F3 | 7h | > 8.4h |
| F4 | 6h | > 7.2h |
| G1 | 8h | > 9.6h |
| G2 | 8h | > 9.6h |
| H1 | 7h | > 8.4h |
| H2 | 6h | > 7.2h |
| I1 | 4h | > 4.8h |
| I2 | 4h | > 4.8h |

### Variance Response

**If single milestone > 20%:**
- Document root cause
- Decide: Is estimate wrong, or is topic harder?
- Update future similar estimates
- Continue (don't block)

**If 2+ consecutive milestones > 20%:**
- Possible: Wrong learning material
- Action: Switch resources (see STAGE_0_REPOSITORIES.md)
- Reassess schedule

**If cumulative variance > 10%:**
- Pace has slipped
- Action: Reduce optional milestones, or extend timeline
- Flag for decision point

---

## Checkpoint Verification Log

### A1: CPU Pipeline & ISA
**Status:** Pending
**Verification:** [ ] Draw pipeline diagram [ ] Explain instruction encoding
**Due:** End of Week 1
**Completed:** — / — / — (date)

### B1: Assembly Syntax & Instructions
**Status:** Pending
**Verification:** [ ] Disassemble C function [ ] Write loop in asm
**Due:** End of Week 2
**Completed:** — / — / —

### C1: Process Model
**Status:** Pending
**Verification:** [ ] Draw state diagram [ ] Explain transitions
**Due:** End of Week 3
**Completed:** — / — / —

### D1: Virtual Memory Concepts
**Status:** Pending
**Verification:** [ ] Trace address translation [ ] Explain TLB miss
**Due:** End of Week 4
**Completed:** — / — / —

### E1: Inode Structure & Operations
**Status:** Pending
**Verification:** [ ] Use stat/ls -i [ ] Explain hard links
**Due:** End of Week 5
**Completed:** — / — / —

### F1: Boot Sequence
**Status:** Pending
**Verification:** [ ] Build kernel [ ] Trace boot output
**Due:** End of Week 6
**Completed:** — / — / —

### F4: Loadable Kernel Modules (LKM)
**Status:** Pending
**Verification:** [ ] Write module [ ] Load/verify [ ] Remove
**Due:** End of Week 7
**Completed:** — / — / —

### G1: TCP/IP Stack Layers
**Status:** Pending
**Verification:** [ ] Draw layer diagram [ ] Trace packet
**Due:** End of Week 8
**Completed:** — / — / —

### G2: Sockets API & Protocol Basics
**Status:** Pending
**Verification:** [ ] Write TCP echo server [ ] Use netstat
**Due:** End of Week 8
**Completed:** — / — / —

### H1: CIA Triad & Threat Modeling
**Status:** Pending
**Verification:** [ ] Create threat model [ ] Identify threats
**Due:** End of Week 9
**Completed:** — / — / —

### H2: Access Control & Authentication
**Status:** Pending
**Verification:** [ ] Explain permission bits [ ] Trace access check
**Due:** End of Week 9
**Completed:** — / — / —

### I1: Compiler & Build System
**Status:** Pending
**Verification:** [ ] Compile program [ ] Inspect assembly
**Due:** End of Week 2
**Completed:** — / — / —

### I2: Debugging & Profiling
**Status:** Pending
**Verification:** [ ] Debug with gdb [ ] Inspect memory/registers
**Due:** End of Week 5
**Completed:** — / — / —

---

## Summary Statistics

| Stat | Value |
|------|-------|
| Mandatory Milestones | 24 |
| Optional Milestones | 8 |
| Total Planned Hours | 175h (mandatory) |
| Actual Hours Spent | — |
| Variance (cumulative) | — |
| Week on Critical Path | Weeks 6-7 (Linux) |
| Fastest Possible Completion | 8 weeks (parallel) |
| Recommended Completion | 9 weeks (sequential) |
| Gate Status | — / 24 complete |

---

## Print/Export Format

**Copy this tracker to your progress file (e.g., progress.json, progress.txt) for weekly updates.**

Update weekly:
- Which milestones progressed?
- Which milestones verified?
- Cumulative hours vs. plan?
- Blockers encountered?
- Resource changes needed?

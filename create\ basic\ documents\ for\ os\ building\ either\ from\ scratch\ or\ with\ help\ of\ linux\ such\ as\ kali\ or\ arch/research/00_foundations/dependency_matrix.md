# Stage 0: Dependency Matrix

## Matrix Legend

```
█ = Hard dependency (must complete before)
▓ = Soft dependency (helpful context)
░ = Related topic (optional)
  = No dependency
```

---

## Dependency Matrix

|           | Arch | Asm | OS | Mem | FS | Linux | Net | Sec | DevE |
|-----------|------|-----|----|----|----|----|----|----|------|
| **Arch**  |      | █   | █  | █  |    |    |    |    |      |
| **Asm**   |      |     | █  | █  | ▓  | █  |    |    | ░    |
| **OS**    |      |     |    | █  | █  | █  | █  | ▓  |      |
| **Mem**   |      |     |    |    | █  | █  | ▓  | ▓  | ░    |
| **FS**    |      |     |    |    |    | █  |    | ░  |      |
| **Linux** |      |     |    |    |    |    | ░  | █  | █    |
| **Net**   |      |     |    |    |    |    |    | ▓  | ░    |
| **Sec**   |      |     |    |    |    |    |    |    | ▓    |
| **DevE**  |      |     |    |    |    |    |    |    |      |

---

## Topic Dependencies (Detailed)

### ARCHITECTURE (A)
- **Blocks:** Asm (B), OS (C), Mem (D)
- **Blocked by:** None
- **Order:** Week 1

### ASSEMBLY (B)
- **Blocks:** OS (C), Mem (D), Linux (F)
- **Depends on:** Arch (A)
- **Helpful for:** DevE (I) — understand compiled code
- **Order:** Week 2

### OS CONCEPTS (C)
- **Blocks:** Mem (D), FS (E), Linux (F), Net (G)
- **Depends on:** Arch (A), Asm (B)
- **Blocked by:** None (Arch, Asm must be done)
- **Order:** Week 3

### MEMORY (D)
- **Blocks:** FS (E), Linux (F)
- **Depends on:** Arch (A), Asm (B), OS (C)
- **Helpful for:** Net (G), Sec (H)
- **Order:** Week 4

### FILESYSTEMS (E)
- **Blocks:** Linux (F)
- **Depends on:** Mem (D), OS (C)
- **Order:** Week 5

### LINUX INTERNALS (F)
- **Blocks:** Sec (H), DevE (I)
- **Depends on:** Asm (B), OS (C), Mem (D), FS (E)
- **Related to:** Net (G) — both kernel subsystems
- **Critical path:** This is the bottleneck (35h)
- **Order:** Weeks 6-7

### NETWORKING (G)
- **Blocks:** None (gates Stage 1)
- **Depends on:** OS (C)
- **Helped by:** Linux (F), Mem (D) — understand kernel stack
- **Related to:** Sec (H) — TLS/encryption
- **Can run parallel to:** Linux (F) starting week 6
- **Order:** Week 8 (can overlap weeks 6-7)

### SECURITY (H)
- **Blocks:** Stage 1
- **Depends on:** Linux (F), Net (G)
- **Related to:** Mem (D), OS (C) — access control foundations
- **Order:** Week 9

### DEV ENVIRONMENT (I)
- **Blocks:** Stage 1 execution
- **Depends on:** Linux (F), Asm (B)
- **Related to:** All (integration tool)
- **Order:** Weeks 1-9 (progressive setup)

---

## Critical Path

**Longest chain (determines minimum time):**

```
Architecture (25h)
    ↓
Assembly (20h)
    ↓
OS Concepts (25h)
    ↓
Memory (20h)
    ↓
Filesystems (15h)
    ↓
Linux Internals (35h) ← BOTTLENECK
    ↓
Total: 140h (critical path)

Additional (non-blocking):
+ Networking (20h) — can overlap Linux weeks 6-7
+ Security (20h) — follows Linux
+ Dev Env (10h) — distributed across all

Total: 190h
```

---

## Parallelization Strategy

### Cannot Parallelize
- Arch → Asm (Asm depends on Arch knowledge)
- Asm → OS (OS depends on understanding assembly)
- OS → Mem (Memory depends on process/scheduling context)

### Can Parallelize (Save ~20h overall)

**Option 1: Overlap Networking**
- Linux Internals (weeks 6-7): 35h
- Networking (week 8): 20h
- **Parallel opportunity:** Start Networking at end of week 6, overlap week 7
- **Savings:** ~5-10h (if topics don't interfere)
- **Tradeoff:** Cognitive load increases

**Option 2: Progressive Dev Environment**
- Don't wait until week 9
- Set up compiler/gdb in week 1
- Set up build system in week 2
- Set up git in week 3
- **Savings:** 0h (redistributes effort, doesn't eliminate)
- **Benefit:** Practice tools as you learn

**Option 3: Reference vs. Deep Learning**
- VFS/filesystem types (optional)
- Scheduling algorithms (optional)
- DNS/routing (optional)
- Cryptography details (optional)
- **Savings:** ~15h if skipped entirely
- **Risk:** Gaps in Stage 1+ understanding

---

## Gate Analysis

### Hard Gates (Must Complete Before Stage 1)
All 24 mandatory milestones represent hard gates:
- No architecture → can't understand CPU behavior
- No assembly → can't read kernel/driver code
- No process/threads → can't understand concurrency
- No memory management → can't work with allocators/heap
- No Linux internals → can't modify kernel
- No networking → can't add telemetry/EDR
- No security fundamentals → can't design threat model

**Total hard gate time:** 175h

### Soft Gates (Complete But Not Blocking)
Optional milestones that accelerate Stage 1 but don't block:
- Scheduling algorithms
- VFS/filesystem types
- DNS/routing
- Cryptography deep dive
- Git workflows

**Total soft gate time:** 15h

---

## Resequencing Options

### Option A: Strict Sequential (Recommended)
```
Arch → Asm → OS → Mem → FS → Linux → Net → Sec → DevE
```
- **Pros:** Simple, no cognitive overhead, clear ordering
- **Cons:** Linux bottleneck feels long
- **Time:** 190h over 9 weeks

### Option B: Parallel Networking
```
Arch → Asm → OS → Mem → FS → [Linux + Net (weeks 6-8)] → Sec → DevE
```
- **Pros:** Reduces felt duration, avoids Linux cognitive fatigue
- **Cons:** Requires multitasking, may need backtracking
- **Time:** 180h over 8 weeks (~22h/week weeks 6-8)

### Option C: Reference + Deep Dive
```
Skip optional topics, parallelize Dev Environment, fast-track Security
```
- **Pros:** Minimum time to Stage 1 gate
- **Cons:** Risk of shallow understanding
- **Time:** 160h over 7-8 weeks
- **Risk Level:** High

**Recommendation:** Use Option A (sequential). Linux Internals is best learned without distraction.

---

## Circular Dependencies

**None detected.** Dependency graph is a DAG (directed acyclic graph).

---

## Topic Interactions

### Tight Coupling (Must Understand Both Together)
- **Architecture ↔ Assembly:** ISA concepts directly used
- **OS ↔ Memory:** Virtual memory is OS mechanism
- **Linux ↔ Filesystem:** VFS is kernel subsystem
- **Networking ↔ Linux:** Network stack is kernel subsystem

### Loose Coupling (Understand Separately, Apply Together)
- **Memory ↔ Networking:** Both use memory, different concerns
- **Security ↔ Filesystems:** Permissions are filesystem feature
- **DevE ↔ All:** Tools used throughout, not core dependency

---

## Knowledge Layering

### Layer 1: Hardware Foundation (25h)
- Architecture

### Layer 2: Code Execution (20h)
- Assembly

### Layer 3: System Abstractions (25h)
- OS Concepts

### Layer 4: Resource Management (20h)
- Memory

### Layer 5: Storage & Syscalls (15h)
- Filesystems

### Layer 6: Integration (35h)
- Linux Internals (brings layers 1-5 together)

### Layer 7: Connectivity (20h)
- Networking

### Layer 8: Protection (20h)
- Security (brings layers 3-7 together)

### Layer 9: Implementation (10h)
- Dev Environment (how to execute all above)

---

## Next Actions

1. Use **checkpoints.md** to verify completion of each milestone
2. Use **milestone_tracker.md** to track progress
3. Use **weekly_plan.md** to structure learning week-by-week
4. Use **dependency_matrix.md** as reference for reordering if blocked

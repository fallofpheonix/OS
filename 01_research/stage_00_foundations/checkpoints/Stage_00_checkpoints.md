# Stage 0: Compressed Checkpoints & Milestones

## Overview

Compressed from 69 checkpoints to **24 measurable milestones** across 9 core topics. Each milestone has:
- **Clear definition** (what success looks like)
- **Verification method** (how to confirm)
- **Time box** (estimated hours)
- **Blocker status** (gates to Stage 1? yes/no)

---

## Architecture (25h)

### Milestone A1: CPU Pipeline & ISA
**Definition:** Explain CPU fetch-decode-execute cycle, understand x86-64 instruction format
**Verification:** Draw pipeline diagram, explain instruction encoding
**Time:** 8h
**Gate:** Yes (blocks everything)

### Milestone A2: Memory Hierarchy
**Definition:** Explain L1/L2/L3 caches, TLB, main memory, working set
**Verification:** Describe cache miss cost, calculate latency differences
**Time:** 8h
**Gate:** Yes (blocks Memory topic)

### Milestone A3: Interrupts & Exceptions
**Definition:** Understand hardware interrupt handling, exception types, interrupt handlers
**Verification:** Trace an interrupt from hardware through handler
**Time:** 6h
**Gate:** Yes (blocks Linux Internals)

### Milestone A4: Registers & Addressing
**Definition:** Know x86-64 register set, addressing modes
**Verification:** Identify register purposes, calculate address offsets
**Time:** 3h
**Gate:** Yes (blocks Assembly)

---

## Assembly (20h)

### Milestone B1: Assembly Syntax & Instructions
**Definition:** Read/write basic x86-64 assembly (AT&T or Intel syntax)
**Verification:** Disassemble C function, hand-write simple loop in assembly
**Time:** 8h
**Gate:** Yes (blocks OS understanding)

### Milestone B2: Calling Conventions & Stack
**Definition:** Understand x86-64 calling convention, stack frame layout
**Verification:** Trace function call through disassembly, identify saved registers
**Time:** 7h
**Gate:** Yes (blocks Linux Internals)

### Milestone B3: ELF Binary Format
**Definition:** Understand ELF headers, sections, symbols, relocation
**Verification:** Parse ELF with readelf, explain relocations
**Time:** 5h
**Gate:** Yes (blocks Linux Internals)

---

## OS Concepts (25h)

### Milestone C1: Process Model
**Definition:** Understand process abstraction, process state transitions, context switching
**Verification:** Draw state diagram, explain when transitions occur
**Time:** 8h
**Gate:** Yes (blocks Stage 1)

### Milestone C2: Thread Model
**Definition:** Understand threading, kernel vs user threads, thread scheduling
**Verification:** Explain 1:1 vs M:N threading, thread creation cost
**Time:** 7h
**Gate:** Yes (blocks Stage 1)

### Milestone C3: Scheduling Algorithms
**Definition:** Understand FIFO, RR, priority, multilevel feedback queues
**Verification:** Simulate scheduling algorithm on task set, calculate turnaround time
**Time:** 6h
**Gate:** No (reference)

### Milestone C4: System Calls & Privilege Transition
**Definition:** Understand syscall mechanism, user→kernel→user transition
**Verification:** Trace syscall with strace, explain overhead
**Time:** 4h
**Gate:** Yes (blocks Linux Internals)

---

## Memory Management (20h)

### Milestone D1: Virtual Memory Concepts
**Definition:** Understand address translation, page tables, TLB
**Verification:** Trace address translation through multi-level page table
**Time:** 8h
**Gate:** Yes (blocks Stage 1)

### Milestone D2: Paging & Page Faults
**Definition:** Understand demand paging, page replacement, working set
**Verification:** Explain page fault handling, predict paging behavior on workload
**Time:** 7h
**Gate:** Yes (blocks Stage 1)

### Milestone D3: Heap & Stack Layout
**Definition:** Understand memory layout (text, data, heap, stack, BSS)
**Verification:** View process memory with /proc/maps, calculate stack growth
**Time:** 5h
**Gate:** Yes (gates Stage 1 security work)

---

## Filesystems (15h)

### Milestone E1: Inode Structure & Operations
**Definition:** Understand inode (metadata, pointers), file descriptors, hard/soft links
**Verification:** Use stat/ls -i to inspect inodes, explain hard link behavior
**Time:** 7h
**Gate:** Yes (blocks Linux Internals)

### Milestone E2: VFS & Filesystem Types
**Definition:** Understand VFS abstraction, ext4 vs FAT, journaling
**Verification:** Compare ext4/btrfs design, explain VFS layer
**Time:** 5h
**Gate:** No (reference)

### Milestone E3: Disk I/O & Page Cache
**Definition:** Understand page cache role, disk I/O path, fsync semantics
**Verification:** Explain page cache eviction, fsync impact
**Time:** 3h
**Gate:** No (optimization)

---

## Linux Internals (35h)

### Milestone F1: Boot Sequence
**Definition:** Trace boot from BIOS/UEFI through bootloader to first process
**Verification:** Build Linux kernel, trace boot with serial/console output
**Time:** 10h
**Gate:** Yes (gates Stage 1 kernel work)

### Milestone F2: Kernel Architecture & Subsystems
**Definition:** Understand monolithic kernel structure, VFS, scheduler, networking stack
**Verification:** Navigate Linux source tree, locate key subsystems
**Time:** 8h
**Gate:** Yes (gates Stage 1)

### Milestone F3: System Call Path
**Definition:** Trace syscall execution from entry point through dispatch to handler
**Verification:** Instrument syscall with strace, identify handler function in kernel
**Time:** 7h
**Gate:** Yes (gates Stage 1)

### Milestone F4: Loadable Kernel Modules (LKM)
**Definition:** Write, compile, insmod/rmmod a simple kernel module
**Verification:** Build & load module, verify with lsmod, remove cleanly
**Time:** 6h
**Gate:** Yes (gates Stage 1 driver/module work)

### Milestone F5: Device Drivers & Device Model
**Definition:** Understand driver architecture, device tree, udev
**Verification:** Trace driver binding in sysfs, understand device probe
**Time:** 4h
**Gate:** No (Stage 1 focus)

---

## Networking (20h)

### Milestone G1: TCP/IP Stack Layers
**Definition:** Understand OSI model, TCP/IP stack (layers 1-4), packet flow
**Verification:** Draw layer interaction diagram, trace packet through stack
**Time:** 8h
**Gate:** Yes (gates Stage 1 security/telemetry)

### Milestone G2: Sockets API & Protocol Basics
**Definition:** Understand socket types (stream, datagram), TCP/UDP/ICMP
**Verification:** Write simple TCP echo client/server, verify with netstat
**Time:** 8h
**Gate:** Yes (gates Stage 1)

### Milestone G3: DNS & Routing Basics
**Definition:** Understand DNS resolution, IP routing, ARP
**Verification:** Trace DNS query with dig, understand routing table
**Time:** 4h
**Gate:** No (reference)

---

## Security Fundamentals (20h)

### Milestone H1: CIA Triad & Threat Modeling
**Definition:** Understand Confidentiality/Integrity/Availability, threat models (STRIDE), attack trees
**Verification:** Create threat model for simple system, identify threats
**Time:** 7h
**Gate:** Yes (gates Stage 1 security)

### Milestone H2: Access Control & Authentication
**Definition:** Understand DAC/MAC/RBAC, user/group model, file permissions
**Verification:** Explain Unix file permission bits, trace access check in kernel
**Time:** 6h
**Gate:** Yes (gates Stage 1)

### Milestone H3: Logging & Audit
**Definition:** Understand logging, audit trails, accountability
**Verification:** Review system logs, enable kernel audit, trace security events
**Time:** 4h
**Gate:** No (Stage 1 observability)

### Milestone H4: Cryptography Basics
**Definition:** Understand symmetric/asymmetric crypto, hashing, TLS handshake overview
**Verification:** Use openssl to create keys, understand RSA vs AES use cases
**Time:** 3h
**Gate:** No (Stage 1+ optional)

---

## Dev Environment (10h)

### Milestone I1: Compiler & Build System
**Definition:** Use gcc/clang, write Makefiles, understand linking
**Verification:** Compile multi-file C program, inspect generated assembly
**Time:** 4h
**Gate:** Yes (gates Stage 1)

### Milestone I2: Debugging & Profiling
**Definition:** Use gdb effectively (breakpoints, stepping, inspection), understand perf
**Verification:** Debug program, set breakpoints, inspect memory/registers
**Time:** 4h
**Gate:** Yes (gates Stage 1)

### Milestone I3: Version Control & Collaboration
**Definition:** Git basics (clone, branch, commit, push), understand workflows
**Verification:** Clone repo, make branch, commit, push
**Time:** 2h
**Gate:** No (soft skill)

---

## Summary Table

| Topic | Milestones | Hours | Gate Topics | Optional |
|-------|-----------|-------|-----------|----------|
| Architecture | 4 | 25 | 4/4 | 0 |
| Assembly | 3 | 20 | 3/3 | 0 |
| OS Concepts | 4 | 25 | 3/4 | 1 (Scheduling) |
| Memory | 3 | 20 | 3/3 | 0 |
| Filesystems | 3 | 15 | 1/3 | 2 (FS types, cache) |
| Linux | 5 | 35 | 4/5 | 1 (drivers) |
| Networking | 3 | 20 | 2/3 | 1 (routing) |
| Security | 4 | 20 | 2/4 | 2 (logging, crypto) |
| Dev Env | 3 | 10 | 2/3 | 1 (git) |
| **TOTAL** | **32** | **190** | **24 Gate** | **8 Optional** |

---

## Stage 1 Gate Checklist

**Must Complete Before Stage 1:**

- [ ] A1: CPU pipeline & ISA (8h)
- [ ] A2: Memory hierarchy (8h)
- [ ] A3: Interrupts & exceptions (6h)
- [ ] A4: Registers & addressing (3h)
- [ ] B1: Assembly syntax (8h)
- [ ] B2: Calling conventions (7h)
- [ ] B3: ELF format (5h)
- [ ] C1: Process model (8h)
- [ ] C2: Thread model (7h)
- [ ] C4: System calls (4h)
- [ ] D1: Virtual memory (8h)
- [ ] D2: Paging (7h)
- [ ] D3: Memory layout (5h)
- [ ] E1: Inodes (7h)
- [ ] F1: Boot sequence (10h)
- [ ] F2: Kernel architecture (8h)
- [ ] F3: Syscall path (7h)
- [ ] F4: Kernel modules (6h)
- [ ] G1: TCP/IP stack (8h)
- [ ] G2: Sockets API (8h)
- [ ] H1: CIA & threat modeling (7h)
- [ ] H2: Access control (6h)
- [ ] I1: Compiler & build (4h)
- [ ] I2: Debugging (4h)

**Total Gate Time:** ~175h

**Optional (Post-Stage 0):**
- C3: Scheduling algorithms
- E2: VFS & filesystem types
- E3: Disk I/O
- F5: Device drivers
- G3: DNS & routing
- H3: Logging
- H4: Cryptography
- I3: Git

**Integration Projects (Optional, post-Stage 0):**
- Mini shell (uses C1, C2, E1, I1-I2)
- Allocator prototype (uses D1-D3, I1-I2)
- Filesystem explorer (uses E1-E3, I1-I2)
- Network utility (uses G1-G2, I1-I2)

---

## Verification Methods

### Type 1: Explanation (Theoretical)
- Write 1-2 paragraph explanation
- Draw diagram
- Example: "Explain CPU cache hierarchy"

### Type 2: Code (Practical)
- Write/run code
- Trace execution
- Example: "Disassemble C function, identify calling convention"

### Type 3: Tool (System)
- Use standard tools
- Inspect system state
- Example: "Use gdb to inspect memory layout"

### Type 4: Analysis (Synthesis)
- Apply knowledge to scenario
- Make predictions
- Example: "Predict paging behavior on workload"

---

## Effort Validation

Milestone time estimates are **sandboxed** — if you exceed 20% over estimate on any milestone, flag it for replan:

- Estimate: 8h → Alert if > 9.6h
- Estimate: 5h → Alert if > 6h
- Estimate: 3h → Alert if > 3.6h

Replan decision criteria:
- Is the topic harder than expected? (increase depth, reduce breadth elsewhere)
- Is the learning material poor? (switch resources)
- Is the estimate too aggressive? (adjust future estimates)

---

## Next: Use these 24 milestones with milestone_tracker.md

# Stage 0: Foundations - Dependency Graph

## Topic Hierarchy & Prerequisites

```
COMPUTER ARCHITECTURE (ROOT - no prerequisites)
├── CPU Design & Microarchitecture
├── Registers & Register Files
├── Memory Hierarchy (L1/L2/L3 cache)
├── Bus Architecture
├── Interrupts & Exceptions
└── → prerequisite for: Binary/Assembly

BINARY & ASSEMBLY BASICS (depends on: Architecture)
├── Binary/Hex/Octal Number Systems
├── Boolean Logic & Gates
├── Instruction Sets (x86-64, ARM)
├── Assembly Language
├── Calling Conventions
└── → prerequisite for: OS Concepts, Linux Internals

OS CONCEPTS (depends on: Architecture, Binary/Assembly)
├── Process vs Thread
├── Process State (ready, running, blocked)
├── Scheduling Algorithms
├── Context Switching
├── Virtual Memory Concepts
├── Protection Rings (Ring 0, 1, 2, 3)
├── Syscalls & Kernel Boundary
└── → prerequisite for: Memory, Processes, Linux Internals

PROCESSES & THREADS (depends on: OS Concepts)
├── Process Creation (fork, exec)
├── Process Lifecycle
├── Signals & IPC
├── Thread Models (1:1, N:1, M:N)
├── Synchronization Primitives (mutex, semaphore)
├── Deadlock & Livelock
└── → prerequisite for: Memory, Scheduling, Linux Internals

MEMORY MANAGEMENT (depends on: Architecture, OS Concepts, Processes)
├── Virtual Memory
├── MMU (Memory Management Unit)
├── Paging vs Segmentation
├── Page Tables & TLB
├── Swapping & Thrashing
├── Heap vs Stack
├── Memory Protection
├── ASLR
└── → prerequisite for: Filesystems, Linux Internals

FILESYSTEMS (depends on: Memory, OS Concepts)
├── Inode Structure
├── File Descriptors
├── Directory Trees
├── File Permissions (rwx)
├── Filesystem Types (ext4, btrfs, XFS)
├── VFS (Virtual Filesystem Switch)
├── Disk I/O & Buffering
├── Journaling
└── → prerequisite for: Linux Internals, Security Fundamentals

LINUX INTERNALS (depends on: Architecture, Assembly, OS Concepts, Processes, Memory, Filesystems)
├── Boot Process (BIOS/UEFI → bootloader → kernel)
├── ELF Binary Format
├── Kernel Architecture (monolithic, microkernel concepts)
├── System Calls (libc wrappers)
├── Device Drivers
├── Loadable Kernel Modules
├── Kernel Subsystems (VFS, networking, scheduler)
└── → prerequisite for: Networking, Development Environment

NETWORKING FOUNDATIONS (depends on: OS Concepts, Linux Internals)
├── TCP/IP Stack (layers 1-4)
├── Sockets API
├── Protocol Suites (TCP, UDP, ICMP)
├── DNS Resolution
├── Routing & Forwarding
├── ARP
├── Network Interfaces & Bridging
├── TLS/SSL Fundamentals
└── → prerequisite for: Security Fundamentals

SECURITY FUNDAMENTALS (depends on: Everything above)
├── CIA Triad (Confidentiality, Integrity, Availability)
├── Threat Modeling
├── Access Control (DAC, MAC, RBAC)
├── Authentication & Authorization
├── Cryptographic Basics (symmetric, asymmetric, hashing)
├── Logging & Monitoring
├── Vulnerability Classes (buffer overflow, injection)
├── Defense in Depth
└── → prerequisite for: Stage 1+ work

DEVELOPMENT ENVIRONMENT (depends on: Linux Internals, Security Fundamentals)
├── Compiler Toolchain (gcc, clang)
├── Version Control (git)
├── Debugging Tools (gdb, lldb)
├── Build Systems (make, cmake, cargo)
├── Code Analysis (static analysis, linters)
├── Testing Frameworks
├── Container Basics (Docker)
└── → prerequisite for: Implementation phases
```

## Learning Sequence

**Optimal order for Stage 0:**

1. **Computer Architecture** (Week 1)
   - Foundation for all system-level understanding
   - No prerequisites

2. **Binary & Assembly Basics** (Week 2)
   - Builds directly on architecture
   - Essential for kernel/driver work later

3. **OS Concepts** (Week 3-4)
   - Abstract layer over architecture
   - Introduces core abstractions (processes, memory, scheduling)

4. **Processes & Threads** (Week 4-5)
   - Concrete implementation of process abstraction
   - Depends on OS concepts

5. **Memory Management** (Week 5-6)
   - Deep dive into virtual memory
   - Depends on architecture + OS concepts

6. **Filesystems** (Week 6-7)
   - Practical storage abstraction
   - Depends on memory management

7. **Linux Internals** (Week 7-9)
   - Integrates all previous topics
   - Concrete implementation on real kernel

8. **Networking Foundations** (Week 9-10)
   - Parallel to Linux internals
   - Adds socket layer understanding

9. **Security Fundamentals** (Week 10-11)
   - Builds on everything above
   - Introduces threat modeling, access control

10. **Development Environment** (Week 11-12)
    - Practical tools for implementation
    - Integrated throughout, finalized here

## Blocking Dependencies

```
None → Architecture ✓
Architecture → Assembly ✓
OS Concepts ← Architecture, Assembly
Memory ← Architecture, OS, Processes
Filesystems ← Memory
Linux ← All previous topics
Networking ← OS, Linux
Security ← All previous topics (holistic understanding required)
```

## Parallel Learning Paths

These can overlap/run in parallel after prerequisites:

- **Path A:** Architecture → Assembly → OS Concepts → Processes → Memory → Filesystems
- **Path B:** (after week 3) OS Concepts → Linux Internals (can run parallel to Path A weeks 5-6)
- **Path C:** (after week 7) Networking Foundations (parallel to Linux deep dive)
- **Path D:** (after week 9) Security Fundamentals (parallel to networking)

**Estimated Total:** 12 weeks of focused study (~40 hours/week = 480 hours)

## Circular Dependencies

None detected. Dependency graph is acyclic (DAG).

## Critical Path

```
Architecture
  → Binary/Assembly
    → OS Concepts
      → Memory Management
        → Linux Internals ← (critical bottleneck)
          → Security Fundamentals
            → Stage 1
```

Speeding up Linux Internals study is the highest leverage optimization.

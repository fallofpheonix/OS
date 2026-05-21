# Stage 0: Foundations - Topic Resource Mapping

## 1. COMPUTER ARCHITECTURE

**Objective:** Understand CPU design, memory hierarchy, interrupts, caching

**Complexity:** Medium | **Effort:** 40 hours | **Duration:** Week 1

**Prerequisites:** None

**Core Concepts:**
- x86-64 vs ARM ISA differences
- CPU microarchitecture (Fetch-Decode-Execute)
- Cache hierarchy (L1/L2/L3, coherency)
- Memory bus architecture
- Interrupt handling hardware
- TLB (Translation Lookaside Buffer)

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Computer Architecture: A Quantitative Approach* (Patterson & Hennessy) | Industry standard, comprehensive |
| **Book** | *The Art of Computer Systems Performance Analysis* (Jain) | Performance fundamentals |
| **Course** | CMU 15-447 (Architecture) | Video lectures available |
| **Course** | UC Berkeley CS152 | Open source, comprehensive |
| **Reference** | Intel Software Developer Manual (Vol 1-3) | Official ISA specification |
| **Reference** | AMD64 Architecture Programmer's Manual | AMD ISA reference |
| **Tool** | cpuid | CPU feature detection |
| **Tool** | likwid | Performance counter access |
| **Simulator** | gem5 | Full-system simulator for learning |
| **Paper** | "The Roofline Model: An Insightful Visual Performance Model" (Williams et al.) | Performance modeling |

**Development Tools Needed:**
- `cpuid` – inspect CPU capabilities
- `perf` – Linux performance counter tool
- `likwid` – performance counter wrapper
- VM with QEMU/KVM – see x86 in action

**Learning Checkpoints:**
- [ ] Explain CPU pipeline stages
- [ ] Draw cache hierarchy diagram
- [ ] Understand TLB misses & page faults
- [ ] Explain interrupt handling hardware
- [ ] Compare x86-64 vs ARM calling conventions

---

## 2. BINARY & ASSEMBLY BASICS

**Objective:** Read/write assembly, understand binary representation

**Complexity:** Medium-High | **Effort:** 50 hours | **Duration:** Week 2

**Prerequisites:** Computer Architecture

**Core Concepts:**
- Number systems (binary, hex, octal)
- Boolean logic, bitwise operations
- x86-64 assembly syntax (AT&T vs Intel)
- CPU registers (rax, rbx, rsp, etc.)
- Addressing modes (direct, indirect, indexed)
- Calling conventions (cdecl, System V AMD64 ABI)
- Stack frame layout

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Reverse Engineering for Beginners* (Dennis Yurichev) | Free, assembly-focused |
| **Book** | *Programming from the Ground Up* (Jonathan Bartlett) | x86 assembly teaching book |
| **Course** | UC Berkeley CS61C (Machine Structures) | Includes assembly labs |
| **Reference** | x86-64 ABI Specification | Function calling conventions |
| **Tool** | objdump | Disassembler |
| **Tool** | gdb | Debugger with assembly stepping |
| **Tool** | radare2 | Interactive disassembler |
| **Tool** | nasm | Assembler (Intel syntax) |
| **Tool** | gcc -S | Compiler to assembly |
| **Paper** | System V AMD64 ABI | Calling convention standard |

**Development Tools Needed:**
- `gcc`/`clang` – compiler
- `gdb` – debugger (essential for assembly learning)
- `objdump` – disassemble binaries
- `radare2` – interactive RE tool
- Text editor with syntax highlighting

**Learning Checkpoints:**
- [ ] Disassemble a simple C program, understand each instruction
- [ ] Hand-write assembly to sum array elements
- [ ] Explain stack frame layout for function calls
- [ ] Understand x86-64 calling convention in detail
- [ ] Identify buffer overflow vulnerability in assembly

---

## 3. OS CONCEPTS

**Objective:** Understand process/thread abstraction, scheduling, virtual memory, privilege levels

**Complexity:** Medium | **Effort:** 45 hours | **Duration:** Weeks 3-4

**Prerequisites:** Computer Architecture, Binary & Assembly Basics

**Core Concepts:**
- Process vs thread distinction
- Process state diagram (ready, running, waiting, terminated)
- CPU scheduling algorithms (FIFO, RR, priority, multilevel feedback)
- Context switching mechanism
- Virtual memory (address translation, paging)
- Privilege levels (Ring 0-3, user vs kernel)
- System calls & kernel boundary crossing
- Interrupts & exceptions

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Operating Systems: Three Easy Pieces* (Remzi & Andrea Arpaci-Dusseau) | Free online, OS fundamentals |
| **Book** | *Modern Operating Systems* (Tanenbaum & Bos) | Comprehensive reference |
| **Book** | *The Linux Programming Interface* (Michael Kerrisk) | Deep Linux-specific knowledge |
| **Course** | UC Berkeley CS162 (OS) | Video lectures + projects |
| **Course** | MIT 6.S081 (xv6 kernel) | Hands-on x86 kernel |
| **Course** | Stanford CS110 (Principles of OS) | System-level programming |
| **Tool** | strace | Trace system calls |
| **Tool** | /proc filesystem | Inspect process state |
| **Tool** | ps, top, htop | Process monitoring |
| **Paper** | "The Unix Timesharing System" (Ritchie & Thompson) | OS design history |

**Development Tools Needed:**
- `strace` – trace system calls (understanding kernel boundary)
- `gdb` – inspect process state
- `/proc` – kernel-provided process introspection
- C compiler – write OS-level code

**Learning Checkpoints:**
- [ ] Draw complete process state diagram
- [ ] Explain context switch cost (cache effects)
- [ ] Compare scheduling algorithms (FIFO vs RR vs priority)
- [ ] Understand privilege level transitions
- [ ] Trace a simple program with strace, identify syscalls
- [ ] Explain virtual memory page table structure

---

## 4. PROCESSES & THREADS

**Objective:** Create, manage, synchronize processes and threads

**Complexity:** Medium | **Effort:** 40 hours | **Duration:** Weeks 4-5

**Prerequisites:** OS Concepts

**Core Concepts:**
- Process creation (fork, exec, wait)
- Process lifecycle & exit codes
- Signals & signal handlers
- Inter-Process Communication (IPC: pipes, sockets, shared memory)
- Thread creation (pthreads, kernel threads)
- Thread synchronization (mutex, semaphore, condition variables)
- Race conditions & critical sections
- Deadlock & livelock detection

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Unix Network Programming* (Stevens & Rago) | IPC, signals, networking |
| **Book** | *Programming with POSIX Threads* (Butenhof) | pthreads in depth |
| **Course** | Stanford CS110 (Systems Programming) | Process/thread labs |
| **Reference** | man pthreads, man 7 signal | POSIX API documentation |
| **Tool** | strace -f | Trace multi-process programs |
| **Tool** | gdb thread / info threads | Debug threads |
| **Tool** | helgrind (Valgrind) | Detect race conditions |
| **Paper** | "The C10K Problem" (Dan Kegel) | Process/thread scaling |

**Development Tools Needed:**
- C compiler with pthreads support
- `strace -f` – trace forked processes
- `gdb` – debug threads (`info threads`, `thread N`)
- `valgrind --tool=helgrind` – race condition detection

**Learning Checkpoints:**
- [ ] Implement fork/exec/wait correctly
- [ ] Handle SIGCHLD to reap zombie processes
- [ ] Create multi-threaded program with mutex
- [ ] Detect & fix race condition in shared variable
- [ ] Understand deadlock: create & recover from it
- [ ] Use condition variables for thread coordination

---

## 5. MEMORY MANAGEMENT

**Objective:** Understand virtual memory, paging, heap, stack, memory protection

**Complexity:** High | **Effort:** 50 hours | **Duration:** Weeks 5-6

**Prerequisites:** Architecture, OS Concepts, Processes

**Core Concepts:**
- Virtual addressing & memory translation
- Page tables & multi-level page tables
- TLB (Translation Lookaside Buffer)
- Paging vs segmentation
- Demand paging & page faults
- Page replacement algorithms (LRU, FIFO)
- Working set & thrashing
- Heap management (malloc/free, allocators)
- Stack layout & growth
- Memory protection (page-level permissions)
- ASLR (Address Space Layout Randomization)
- Copy-on-Write (CoW)

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Operating Systems: Three Easy Pieces* (Ch. 12-20) | Virtual memory in depth |
| **Book** | *Understanding the Linux Virtual Memory Manager* (Gorman) | Linux-specific deep dive |
| **Course** | UC Berkeley CS162 (Virtual Memory lectures) | Video explanations |
| **Course** | MIT 6.S081 (xv6 pagetables lab) | Hands-on paging implementation |
| **Reference** | Intel Manual Vol. 3 (Paging) | Hardware details |
| **Tool** | vmstat | Virtual memory stats |
| **Tool** | /proc/[pid]/maps | View process memory layout |
| **Tool** | pmap | Process memory visualization |
| **Tool** | valgrind --tool=memcheck | Memory error detection |
| **Paper** | "Virtual Memory in Contemporary Processors" (Intel whitepaper) | Modern caching |

**Development Tools Needed:**
- C compiler
- `gdb` – inspect memory layout
- `valgrind` – detect memory errors
- `/proc/[pid]/maps` – view address space
- `perf` – measure page faults (`perf record -e page-faults`)

**Learning Checkpoints:**
- [ ] Draw 4-level page table & explain address translation
- [ ] Calculate TLB miss rate given working set size
- [ ] Implement simple malloc/free
- [ ] Understand CoW fork behavior
- [ ] Use valgrind to find memory leaks
- [ ] Explain ASLR & why it matters for security
- [ ] Measure paging impact with vmstat

---

## 6. FILESYSTEMS

**Objective:** Understand file storage, inodes, directories, file access control

**Complexity:** Medium | **Effort:** 40 hours | **Duration:** Week 6-7

**Prerequisites:** Memory Management, OS Concepts

**Core Concepts:**
- Inode structure (metadata, permissions, pointers)
- File descriptors (kernel-side, process-side)
- Directory entries & hierarchies
- File permissions (rwx, umask, special bits)
- Hard vs symbolic links
- Filesystem types (ext4, btrfs, XFS differences)
- VFS (Virtual Filesystem Switch) abstraction
- Disk I/O & buffering (page cache)
- Fsync, sync, durability guarantees
- Journaling & crash recovery
- Filesystem optimization (fragmentation, alignment)

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Understanding Linux Kernel Internals* (Rajkumar) | VFS design |
| **Book** | *The Linux Programming Interface* (Ch. 4-5) | Files & directories |
| **Course** | UC Berkeley CS162 (Filesystems lecture) | Design patterns |
| **Course** | MIT 6.S081 (File System labs) | xv6 FS implementation |
| **Reference** | ext4 Documentation | Format specification |
| **Tool** | ls -i | Show inode numbers |
| **Tool** | stat | Inode metadata |
| **Tool** | debugfs | Ext4 inspector |
| **Tool** | hexdump | Raw disk inspection |
| **Tool** | fstab, mount | Filesystem configuration |
| **Paper** | "A Fast File System for Unix" (McKusick et al.) | FFS design principles |

**Development Tools Needed:**
- `ls -i` – inspect inodes
- `stat` – file metadata
- `debugfs` – filesystem introspection
- `hexdump` – raw data inspection
- C compiler – write filesystem utilities

**Learning Checkpoints:**
- [ ] Understand inode structure (data vs metadata pointers)
- [ ] Explain hard link behavior (same inode, ref count)
- [ ] Calculate disk space used accounting for blocks
- [ ] Compare ext4 vs btrfs journaling approaches
- [ ] Use debugfs to inspect filesystem internals
- [ ] Understand page cache role in I/O
- [ ] Explain fsync semantics

---

## 7. LINUX INTERNALS

**Objective:** Understand Linux boot, ELF, kernel modules, drivers, subsystems

**Complexity:** High | **Effort:** 60 hours | **Duration:** Weeks 7-9

**Prerequisites:** Architecture, Assembly, OS Concepts, Memory, Filesystems

**Core Concepts:**
- Boot sequence (BIOS/UEFI → bootloader → kernel → init)
- ELF binary format (headers, sections, symbols)
- Kernel architecture (monolithic, modular subsystems)
- System call interface (INT 0x80, SYSCALL instruction)
- Device drivers (device model, udev)
- Loadable kernel modules (insmod, rmmod, depmod)
- Kernel subsystems (VFS, networking stack, scheduler)
- Kernel debugging (printk, debugfs, kprobes)
- Device tree (ARM, embedded systems)
- ACPI & firmware interaction

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Linux Kernel Development* (Robert Love) | Industry standard reference |
| **Book** | *Understanding the Linux Kernel* (Bovet & Cesati) | Deep internals |
| **Book** | *Linux Device Drivers* (Corbet, Rubini, Kroah-Hartman) | Driver development |
| **Course** | MIT 6.S081 (xv6 + Linux kernel labs) | Hands-on kernel |
| **Course** | Linux Kernel Newbies (eLinux) | Beginner-friendly intro |
| **Source** | Linux kernel source (kernel.org) | Read actual code |
| **Reference** | Linux man pages | Full API documentation |
| **Tool** | kprobes / kretprobes | Dynamic kernel tracing |
| **Tool** | perf | Kernel profiling |
| **Tool** | BCC / eBPF tools | Kernel introspection |
| **Tool** | objdump -D vmlinux | Disassemble kernel |
| **Paper** | Linux kernel documentation (/Documentation) | Official reference |

**Development Tools Needed:**
- Linux source tree (kernel.org)
- GCC toolchain with kernel headers
- `make` / `make menuconfig` – kernel compilation
- `gdb` + vmlinux symbols – kernel debugging
- `perf` – kernel profiling
- `kprobes` – dynamic tracing
- `strace` + `ltrace` – trace system/library calls

**Learning Checkpoints:**
- [ ] Trace boot sequence from BIOS through first process
- [ ] Parse ELF header with readelf/objdump
- [ ] Write & load a simple kernel module
- [ ] Hook a system call with kprobes
- [ ] Understand kernel memory layout (text, data, heap, stack)
- [ ] Trace a syscall from userspace through kernel
- [ ] Read & understand device driver code
- [ ] Explain kernel preemption & scheduling

---

## 8. NETWORKING FOUNDATIONS

**Objective:** Understand TCP/IP stack, sockets, routing, DNS, TLS basics

**Complexity:** Medium | **Effort:** 45 hours | **Duration:** Weeks 9-10

**Prerequisites:** OS Concepts, Linux Internals

**Core Concepts:**
- OSI Model (layers 1-7)
- TCP/IP stack (layers 1-4: physical, datalink, network, transport)
- IP addressing (IPv4, IPv6, subnetting, CIDR)
- TCP (connection establishment, state machine, flow control)
- UDP (connectionless, datagram)
- ICMP (ping, traceroute)
- ARP (address resolution)
- DNS (queries, responses, caching, recursion)
- Routing (forwarding tables, routing protocols overview)
- Sockets API (stream, datagram, raw)
- TLS/SSL (encryption, certificates, handshake overview)
- Network tools & diagnostics

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Unix Network Programming Vol. 1* (Stevens & Rago) | Sockets API, networking details |
| **Book** | *TCP/IP Illustrated Vol. 1* (Stevens) | Packet-level explanations |
| **Book** | *Beej's Guide to Network Programming* | Free, beginner-friendly |
| **Course** | Stanford CS144 (Intro to Networking) | Video lectures available |
| **Course** | UC Berkeley CS168 (Internet Architecture) | Advanced networking |
| **Reference** | RFC 793 (TCP), RFC 791 (IP) | Protocol specifications |
| **Tool** | tcpdump / Wireshark | Packet capture & analysis |
| **Tool** | netstat / ss | Network statistics |
| **Tool** | ping, traceroute, dig, nslookup | Diagnostic tools |
| **Tool** | iperf | Network performance testing |
| **Paper** | "Anatomy of a Large-Scale Hypertextual Search Engine" | Web at scale |

**Development Tools Needed:**
- C compiler
- `tcpdump` / `Wireshark` – packet analysis
- `strace` – trace network syscalls
- `netstat` / `ss` – network stats
- Sockets API (POSIX standard library)

**Learning Checkpoints:**
- [ ] Implement simple TCP echo server/client
- [ ] Capture & analyze TCP three-way handshake with Wireshark
- [ ] Explain TCP flow control (sliding window)
- [ ] Trace DNS query with tcpdump
- [ ] Understand ARP request/response
- [ ] Implement UDP client (connectionless)
- [ ] Explain TLS handshake overview
- [ ] Diagnose network issues with standard tools

---

## 9. SECURITY FUNDAMENTALS

**Objective:** Understand CIA triad, threat modeling, access control, cryptography basics, vulnerabilities

**Complexity:** Medium-High | **Effort:** 50 hours | **Duration:** Weeks 10-11

**Prerequisites:** All previous topics (holistic understanding)

**Core Concepts:**
- CIA Triad (Confidentiality, Integrity, Availability)
- Threat modeling (STRIDE, attack trees)
- Authentication vs authorization
- Access control models (DAC, MAC, RBAC)
- Cryptography fundamentals (symmetric, asymmetric, hashing)
- Key management & exchange
- Digital signatures & certificates
- Common vulnerability classes (CWE-25, OWASP Top 10)
- Secure coding practices
- Vulnerability analysis & exploitation (defensive context)
- Logging & audit trails
- Defense in depth & security controls

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Book** | *Security Engineering* (Ross Anderson) | Comprehensive, free online |
| **Book** | *Secure Coding* (Seacord) | Defensive programming |
| **Book** | *The Web Application Hacker's Handbook* | Web security, pentesting context |
| **Book** | *Cryptography Engineering* (Ferguson, Schneier, Kohno) | Applied cryptography |
| **Course** | Stanford CS155 (Computer Security) | Video lectures |
| **Course** | MIT 6.858 (Computer Systems Security) | Advanced security |
| **Reference** | OWASP Top 10 | Web vulnerability categories |
| **Reference** | CWE-25 (Most Dangerous Software Weaknesses) | Vulnerability taxonomy |
| **Reference** | NIST Cybersecurity Framework | Standards & guidelines |
| **Tool** | OpenSSL | Cryptography & TLS testing |
| **Tool** | Bandit (Python) / clang-analyzer | Static analysis |
| **Tool** | AFL / libFuzzer | Fuzzing (finding vulnerabilities) |
| **Paper** | "The Confused Deputy Problem" | Authorization classic |

**Development Tools Needed:**
- C/Python compiler
- `openssl` – cryptography testing
- Static analysis tools (clang-analyzer, bandit)
- Fuzzing tools (AFL, libFuzzer)
- Debugger for vulnerability analysis

**Learning Checkpoints:**
- [ ] Explain CIA triad with examples
- [ ] Create threat model using STRIDE
- [ ] Understand privilege escalation attacks
- [ ] Implement simple hash function (understand properties)
- [ ] Explain RSA encryption/signing (math overview)
- [ ] Identify buffer overflow vulnerability
- [ ] Use OWASP Top 10 to audit web app
- [ ] Explain defense-in-depth with examples
- [ ] Review code for common security flaws

---

## 10. DEVELOPMENT ENVIRONMENT

**Objective:** Set up toolchain, version control, debugging, build systems, testing

**Complexity:** Low-Medium | **Effort:** 30 hours | **Duration:** Week 11-12

**Prerequisites:** Linux Internals, Security Fundamentals (practical integration)

**Core Concepts:**
- Compiler toolchain (gcc, clang, cross-compilation)
- Assemblers (gas, nasm)
- Linkers & linking (static, dynamic, position-independent code)
- Version control (git, branching, collaboration)
- Build systems (Makefiles, CMake, Cargo for Rust)
- Debuggers (gdb, lldb, debugging techniques)
- Code analysis & linters (clang-analyze, pylint, clippy)
- Testing frameworks (unit, integration, property-based)
- Containers & VMs (Docker, QEMU, VirtualBox)
- Development workflows & best practices

**Key Resources:**

| Type | Resource | Purpose |
|------|----------|---------|
| **Reference** | GNU Compiler Collection Manual | GCC documentation |
| **Reference** | Linker & Loader (John Levine) | Free ebook, linker internals |
| **Course** | Pro Git (Scott Chacon) | Git book, free online |
| **Guide** | GDB Tutorial (Beej's) | Debugging guide |
| **Guide** | Docker Documentation | Container fundamentals |
| **Tool** | gcc/clang | Compiler |
| **Tool** | gdb / lldb | Debugger |
| **Tool** | make / cmake | Build system |
| **Tool** | git | Version control |
| **Tool** | valgrind | Memory debugging |
| **Tool** | clang-analyzer | Static analysis |
| **Tool** | Docker | Container runtime |
| **Tool** | QEMU / VirtualBox | VM emulation |

**Development Tools Setup:**
- **On macOS:** `brew install gcc llvm cmake gdb git valgrind docker`
- **On Linux:** `apt-get install build-essential gdb git valgrind docker.io cmake`
- **IDE:** VS Code + Extensions (C/C++, Rust, Python) OR CLion

**Learning Checkpoints:**
- [ ] Compile C program with gcc, inspect with objdump
- [ ] Clone, branch, commit, push with git
- [ ] Debug program with gdb (breakpoints, stepping, inspection)
- [ ] Write Makefile for multi-file project
- [ ] Run static analyzer on code
- [ ] Create Dockerfile for application
- [ ] Write unit tests with framework (gtest, pytest)
- [ ] Set up pre-commit hooks

---

## Stage 0 Summary Table

| Topic | Duration | Complexity | Effort | Prerequisites | Key Output |
|-------|----------|------------|--------|---------------|-----------|
| Computer Architecture | Week 1 | Medium | 40h | None | Understand CPU hardware |
| Binary & Assembly | Week 2 | Medium-High | 50h | Arch | Disassemble, read asm |
| OS Concepts | Weeks 3-4 | Medium | 45h | Arch, Asm | Process/memory/scheduling |
| Processes & Threads | Weeks 4-5 | Medium | 40h | OS Concepts | Multi-threaded programs |
| Memory Management | Weeks 5-6 | High | 50h | Arch, OS, Processes | Virtual memory internals |
| Filesystems | Weeks 6-7 | Medium | 40h | Memory, OS | Filesystem understanding |
| Linux Internals | Weeks 7-9 | High | 60h | All above | Boot, ELF, modules, drivers |
| Networking Foundations | Weeks 9-10 | Medium | 45h | OS, Linux | TCP/IP, sockets, tools |
| Security Fundamentals | Weeks 10-11 | Medium-High | 50h | All above | Threat modeling, crypto basics |
| Development Environment | Weeks 11-12 | Low-Medium | 30h | Linux, Security | Functional toolchain |

**Total Estimated Time:** 480 hours (~12 weeks at 40h/week)

**Next Steps After Stage 0:**
- Complete all learning checkpoints
- Build 2-3 small projects integrating Stage 0 knowledge
- Then proceed to Stage 1 (Kernel path, System internals)

# Stage 0: Foundations - Repository & Tool Mapping

## Repository Index by Topic

### 1. COMPUTER ARCHITECTURE

**Learning/Reference Repos:**
- **gem5** (https://github.com/gem5/gem5)
  - Full system simulator for architecture learning
  - CPU models, memory hierarchy, caching
  - Language: Python, C++
  - Use: Simulate CPU behavior, experiment with cache parameters
  - Difficulty: Advanced (learning resource)

- **perf-events** (https://github.com/andikleen/pmu-tools)
  - PMU-tools for accessing performance counters
  - Analyze real CPU behavior on Linux
  - Language: Python
  - Use: Access L1/L2/L3 cache stats, branch misprediction rates
  - Difficulty: Medium

- **likwid** (https://github.com/RRZE-HPC/likwid)
  - Performance counter library & tools
  - Lightweight access to PMU (Performance Monitoring Unit)
  - Language: C
  - Use: Benchmark cache behavior, memory bandwidth
  - Difficulty: Medium

**Tools:**
- Intel VTune (proprietary, free for Linux)
- AMD uProf
- Linux `perf` (built-in)

---

### 2. BINARY & ASSEMBLY BASICS

**Essential Repos:**
- **binutils** (https://sourceware.org/binutils/)
  - objdump, nm, strings, readelf
  - Standard disassembly tools
  - Language: C
  - Use: Disassemble binaries, inspect symbols
  - Difficulty: Easy (tool use)

- **GDB** (https://sourceware.org/gdb/)
  - GNU Debugger
  - Step through assembly, inspect registers
  - Language: C
  - Use: Debug & understand program execution
  - Difficulty: Medium (learning curve)

- **radare2** (https://github.com/radareorg/radare2)
  - Interactive disassembler & debugger
  - Reverse engineering framework
  - Language: C
  - Use: Advanced binary analysis
  - Difficulty: Medium-High

- **NASM** (https://github.com/netwide-assembler/nasm)
  - Intel syntax assembler
  - Write assembly programs
  - Language: C
  - Use: Assemble handwritten assembly code
  - Difficulty: Easy-Medium

- **capstone** (https://github.com/aquynh/capstone)
  - Disassembly engine library
  - Available in Python bindings
  - Language: C (with language bindings)
  - Use: Programmatic disassembly
  - Difficulty: Medium

**Learning Repos:**
- **cs-app** (CMU Computer Systems textbook examples)
  - Provided as supplements to textbook
  - Assembly examples & labs
  - Language: C & Assembly
  - Use: Follow along with textbook learning
  - Difficulty: Easy-Medium

---

### 3. OS CONCEPTS

**Educational Kernels:**
- **xv6** (https://github.com/mit-pdos/xv6-public)
  - Minimal Unix-like kernel (MIT course)
  - Small, readable, educational
  - Language: C & x86 Assembly
  - Use: Understand real OS implementation
  - Difficulty: Medium
  - Courses: MIT 6.S081, UC Berkeley CS162

- **Linux kernel (early versions)** (https://github.com/torvalds/linux)
  - Full production kernel (complex)
  - Historical versions for learning (v0.01)
  - Language: C & Assembly
  - Use: Reference current best practices
  - Difficulty: Very High (production code)

**Reference Implementations:**
- **ucore** (https://github.com/chyyuu/ucore_os_lab)
  - Educational OS lab (Tsinghua University)
  - Chinese university course materials
  - Language: C & Assembly
  - Use: Labs & assignments
  - Difficulty: Medium

---

### 4. PROCESSES & THREADS

**Example Repos & Utilities:**
- **Linux kernel** (process subsystem)
  - kernel/sched/, kernel/signal.c
  - Actual implementation of fork, exec, signals
  - Language: C
  - Use: See real scheduling implementation
  - Difficulty: High

- **pthreads examples** (various)
  - Implementation provided by libc (glibc, musl, etc.)
  - Study pthreads source in: https://github.com/bminor/glibc
  - Language: C
  - Use: Understand thread creation/synchronization
  - Difficulty: Medium-High

- **concurrency tutorial repos**
  - Example repositories on GitHub showing fork/threads
  - Language: C
  - Use: Copy patterns for your learning
  - Difficulty: Easy-Medium

---

### 5. MEMORY MANAGEMENT

**Linux Memory Source:**
- **Linux kernel (mm/ directory)**
  - mm/page_alloc.c (page allocator)
  - mm/vmscan.c (page reclamation)
  - mm/page_table_check.c (page table management)
  - Language: C
  - Use: Production memory manager
  - Difficulty: Very High

**Allocator Implementations:**
- **jemalloc** (https://github.com/jemalloc/jemalloc)
  - Production memory allocator (Redis, Firefox)
  - Well-documented, modular design
  - Language: C
  - Use: Study allocation strategies
  - Difficulty: Medium-High

- **tcmalloc** (https://github.com/google/tcmalloc)
  - Google's threaded allocator
  - Performance-focused design
  - Language: C++
  - Use: Understand allocation for multi-threaded apps
  - Difficulty: Medium-High

- **mimalloc** (https://github.com/microsoft/mimalloc)
  - Microsoft's memory allocator
  - Modern, clean implementation
  - Language: C
  - Use: Learn modern allocation techniques
  - Difficulty: Medium

**Educational Projects:**
- **malloc-lab** (CMU)
  - Exercise in malloc/free implementation
  - Part of CS:APP course materials
  - Language: C
  - Use: Implement simple allocator
  - Difficulty: Medium

---

### 6. FILESYSTEMS

**Filesystem Implementations:**
- **Linux kernel (fs/ directory)** (https://github.com/torvalds/linux/tree/master/fs)
  - ext4/, btrfs/, xfs/ subdirectories
  - Production filesystem implementations
  - Language: C
  - Use: Study real FS design
  - Difficulty: Very High

- **FUSE** (https://github.com/libfuse/libfuse)
  - Filesystem in Userspace
  - Write filesystems without kernel coding
  - Language: C
  - Use: Implement custom filesystem
  - Difficulty: Medium

- **ceph** (https://github.com/ceph/ceph)
  - Distributed filesystem
  - Modern storage architecture
  - Language: C++
  - Use: Study distributed FS design
  - Difficulty: Very High

**Educational Filesystems:**
- **ext2fuse** (Educational ext2 implementation)
  - Simple ext2 in userspace
  - Language: C
  - Use: Learn inode structure
  - Difficulty: Medium

---

### 7. LINUX INTERNALS

**Primary Source:**
- **Linux kernel** (https://github.com/torvalds/linux)
  - kernel/ (core subsystems)
  - drivers/ (device drivers)
  - arch/x86/boot/ (bootloader)
  - Language: C & Assembly
  - Use: Read actual kernel code
  - Difficulty: Very High

**Learning-Focused Repos:**
- **Linux kernel newbies** (https://kernelnewbies.org/)
  - Community resources, guides
  - Build kernel step-by-step
  - Language: Documentation + shell
  - Use: Kernel building tutorials
  - Difficulty: Medium

- **systemtap** (https://sourceware.org/systemtap/)
  - Dynamic kernel probing tool
  - Trace kernel behavior without modifying code
  - Language: C, DSL
  - Use: Non-invasive kernel introspection
  - Difficulty: Medium

- **kprobes examples** (in-kernel documentation)
  - kernel/trace/kprobes.c
  - Hooks into kernel functions dynamically
  - Language: C
  - Use: Instrument kernel functions
  - Difficulty: Medium-High

**Bootloader/ELF:**
- **GRUB2** (https://github.com/coreos/grub)
  - Boot loader (reference implementation)
  - Language: C & Assembly
  - Use: Understand boot process
  - Difficulty: High

- **u-boot** (https://github.com/u-boot/u-boot)
  - Embedded bootloader
  - Simpler than GRUB
  - Language: C & Assembly
  - Use: Learn bootloader design
  - Difficulty: Medium-High

---

### 8. NETWORKING FOUNDATIONS

**Linux Network Stack:**
- **Linux kernel (net/ directory)** (https://github.com/torvalds/linux/tree/master/net)
  - net/ipv4/, net/tcp/, net/socket.c
  - Production TCP/IP implementation
  - Language: C
  - Use: See real networking code
  - Difficulty: Very High

**Packet Analysis:**
- **tcpdump** (https://github.com/the-tcpdump-group/tcpdump)
  - Packet capture utility
  - Language: C
  - Use: Capture & analyze traffic
  - Difficulty: Easy-Medium (usage)

- **Wireshark** (https://github.com/wireshark/wireshark)
  - Interactive packet analyzer (GUI)
  - Language: C
  - Use: Visualize protocols
  - Difficulty: Easy (usage)

**Socket Libraries:**
- **libuv** (https://github.com/libuv/libuv)
  - Cross-platform async I/O
  - Event-driven networking
  - Language: C
  - Use: Study async socket patterns
  - Difficulty: Medium

- **QUIC** (https://github.com/quicwg)
  - New transport protocol
  - Modern networking design
  - Language: C/C++
  - Use: Study modern transport layer
  - Difficulty: High

**Learning Repos:**
- **beej-networking-guide examples**
  - Companion to Beej's guide
  - Simple socket client/server
  - Language: C
  - Use: Learn socket API step-by-step
  - Difficulty: Easy

---

### 9. SECURITY FUNDAMENTALS

**Cryptography Libraries:**
- **OpenSSL** (https://github.com/openssl/openssl)
  - Standard cryptography library
  - Symmetric, asymmetric, hashing, TLS
  - Language: C
  - Use: Reference implementation
  - Difficulty: Medium (understanding code)

- **libsodium** (https://github.com/jedisct1/libsodium)
  - Modern cryptography library
  - Cleaner API than OpenSSL
  - Language: C
  - Use: Learn modern crypto best practices
  - Difficulty: Easy-Medium

- **GnuPG** (https://github.com/gpg/gnupg)
  - GPG/PGP implementation
  - Complete encryption toolchain
  - Language: C
  - Use: Understand key management
  - Difficulty: Medium-High

**Vulnerability Tools:**
- **AFL** (https://github.com/google/AFL) – Fuzzing framework
  - American Fuzzy Lop
  - Finds bugs through mutation
  - Language: C
  - Use: Discover vulnerabilities
  - Difficulty: Medium

- **libFuzzer** (https://github.com/google/libfuzzer)
  - In-process fuzzing
  - Integration with LLVM/Clang
  - Language: C++
  - Use: Coverage-guided fuzzing
  - Difficulty: Medium-High

- **AddressSanitizer (ASAN)** (clang/gcc built-in)
  - Memory error detection
  - Runtime checking
  - Language: N/A (compiler feature)
  - Use: Find memory bugs
  - Difficulty: Easy

**Code Analysis:**
- **clang-analyzer** (built into clang)
  - Static analysis
  - Find common bugs
  - Language: N/A (built-in)
  - Use: Catch issues before runtime
  - Difficulty: Easy

- **cppcheck** (https://github.com/danmar/cppcheck)
  - C/C++ static analysis
  - Standalone tool
  - Language: C++
  - Use: Lint C code
  - Difficulty: Easy

---

### 10. DEVELOPMENT ENVIRONMENT

**Compiler Toolchains:**
- **GCC** (https://github.com/gcc-mirror/gcc)
  - GNU Compiler Collection
  - Language: C
  - Use: Compile C/C++/Fortran
  - Difficulty: Easy (usage)

- **LLVM/Clang** (https://github.com/llvm/llvm-project)
  - Modular compiler infrastructure
  - Modern toolchain
  - Language: C++
  - Use: Compile, analyze, optimize
  - Difficulty: Medium (toolchain design)

**Build Systems:**
- **GNU Make** (https://github.com/mirror/make)
  - Standard build system
  - Language: C
  - Use: Automate compilation
  - Difficulty: Easy-Medium

- **CMake** (https://github.com/Kitware/CMake)
  - Cross-platform build generator
  - Language: C++
  - Use: Build complex projects
  - Difficulty: Medium

- **Cargo** (https://github.com/rust-lang/cargo)
  - Rust package manager & build tool
  - Language: Rust
  - Use: Build Rust projects
  - Difficulty: Easy (Rust-specific)

**Version Control:**
- **Git** (https://github.com/git/git)
  - Distributed version control
  - Language: C
  - Use: Track code changes
  - Difficulty: Easy-Medium

**Debuggers:**
- **GDB** (https://sourceware.org/gdb/)
  - GNU Debugger
  - Language: C
  - Use: Debug programs
  - Difficulty: Medium

- **LLDB** (https://github.com/llvm-mirror/lldb)
  - LLVM debugger
  - Language: C++
  - Use: Debug with LLVM stack
  - Difficulty: Medium

**Testing:**
- **Google Test (gtest)** (https://github.com/google/googletest)
  - C++ unit testing framework
  - Language: C++
  - Use: Write & run tests
  - Difficulty: Easy

- **pytest** (https://github.com/pytest-dev/pytest)
  - Python testing framework
  - Language: Python
  - Use: Test Python code
  - Difficulty: Easy

**Containers:**
- **Docker** (https://github.com/moby/moby)
  - Container runtime
  - Language: Go
  - Use: Package applications
  - Difficulty: Easy-Medium

- **QEMU** (https://github.com/qemu/qemu)
  - Machine emulator
  - Language: C
  - Use: Emulate different architectures
  - Difficulty: Medium

---

## Tool Installation Guide

### macOS
```bash
brew install gcc llvm cmake git gdb valgrind docker nasm
brew tap messense/macos-cross-toolchains
brew install arm-linux-gnueabihf-gcc  # For cross-compilation
```

### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install build-essential gcc-doc gdb git valgrind
sudo apt-get install nasm cmake
sudo apt-get install docker.io
sudo apt-get install linux-headers-$(uname -r)
```

### Fedora/RHEL
```bash
sudo dnf install gcc gcc-c++ gdb git cmake make
sudo dnf install valgrind nasm
sudo dnf install docker
sudo dnf install kernel-devel
```

---

## Contribution Difficulty Classification

### Beginner-Friendly (Start Here)
- binutils (add simple features)
- FUSE (write custom filesystem in user space)
- cppcheck (add new checks)
- libsodium (documentation, examples)

### Intermediate (Stage 0 → Stage 1)
- jemalloc (allocator optimizations)
- libuv (networking features)
- BCC/eBPF (performance tools)
- Wireshark (protocol dissectors)

### Advanced (Stage 1+)
- Linux kernel (core subsystems)
- LLVM/Clang (compiler features)
- QEMU (emulation features)
- Ceph (distributed systems)

### Research Only (Reference)
- gem5 (CPU simulation)
- systemtap (kernel instrumentation)
- AFL (fuzzing research)

---

## Repository Learning Path

**Week 1-2 (Architecture & Assembly):**
- Read binutils source (how objdump works)
- Use GDB to step through assembly
- Study radare2 plugins

**Week 3-5 (OS Concepts & Processes):**
- Read xv6 kernel code cover-to-cover
- Study Linux kernel sched/ (5000 lines)
- Examine fork/exec in glibc

**Week 5-7 (Memory & Filesystems):**
- Read jemalloc implementation
- Study Linux kernel mm/ subsystem
- Explore FUSE examples

**Week 7-9 (Linux Internals):**
- Build Linux kernel from source
- Read bootloader code (GRUB/u-boot)
- Study ELF parsing in binutils

**Week 9-11 (Networking & Security):**
- Trace packets with tcpdump/Wireshark
- Study OpenSSL/libsodium implementations
- Run AFL/ASAN on sample code

**Week 11-12 (Development Setup):**
- Master CMake for builds
- Write custom make targets
- Set up Docker environment

---

## Next Steps

After Stage 0:
1. **Complete all learning checkpoints**
2. **Build 2-3 small projects:**
   - Simple shell implementation (processes, pipes)
   - Memory allocator (memory management)
   - File explorer utility (filesystems)
3. **Move to Stage 1 (System Internals & Kernel Path)**

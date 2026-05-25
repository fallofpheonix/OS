# Phase 2: Low-Level Programming

## Overview

Write code that directly interfaces with the operating system and hardware. Understand how programs are built, linked, loaded, executed, debugged, profiled, and constrained by unsafe memory operations.

## Classification

- Stage: `Stage_00_Foundations`
- Type: `FOUNDATIONAL`
- Status: `RESEARCH_ONLY`
- Difficulty: advanced
- Estimated duration: 6-8 weeks
- Depends on:
  - Phase 0 Computer Science Foundations
  - Phase 1 Computer Architecture
- Blocks:
  - Stage 01 System Internals
  - Stage 02 Linux and Distros
  - Stage 04 Security
  - Stage 18 Custom OS

## Languages And Tools

### C

- Direct hardware and OS interface.
- Manual memory management.
- Foundation for OS and systems programming.
- Uses:
  - Memory allocators.
  - Shell.
  - Syscall wrappers.

### Rust

- Memory safety without garbage collection.
- Systems programming with fewer memory vulnerabilities.
- FFI with C.
- Uses:
  - Allocators.
  - Syscall abstractions.

### Assembly x86_64

- Understand generated code.
- Inline assembly in C/Rust.
- Low-level optimization.
- Uses:
  - Register usage.
  - Calling conventions.
  - Bootloaders.

### Python

- Rapid prototyping.
- Testing.
- `ctypes` syscall exploration.
- Performance baseline comparisons.

## Research

### Program Structure And Linking

#### Calling Conventions

- x86_64 System V ABI.
- Argument registers:
  - rdi.
  - rsi.
  - rdx.
  - rcx.
  - r8.
  - r9.
  - stack.
- Return registers:
  - rax.
  - rdx:rax for 128-bit values.
- Caller-saved registers.
- Callee-saved registers.
- Red zone: 128 bytes below `rsp` on x86_64.
- cdecl.
- stdcall.
- fastcall.
- Function performance impact.

#### Stack Frames

- Frame pointer: `rbp`.
- Stack pointer: `rsp`.
- Function prologue.
- Function epilogue.
- Local variable layout.
- Return address storage.
- Stack alignment requirements.

#### Linkers

- Symbol resolution.
- Relocation entries.
- Relocation types.
- Link-time optimization.
- Dead code elimination.
- Symbol versioning.
- Dynamic linking.

#### Loaders

- OS program loader responsibility.
- Text segment.
- Data segment.
- BSS.
- Heap.
- Stack.
- ASLR.
- Dynamic linker/loader.
- Lazy binding.
- PLT.
- GOT.

### Executable Format And Linking

#### ELF

- ELF header.
- Magic.
- Architecture.
- Entry point.
- Program headers.
- Section headers.
- Symbol table.
- String tables.
- Relocation sections.
- Dynamic linking metadata.

#### Static Vs Dynamic Linking

- Static linking.
- Dynamic linking.
- Pros and cons.
- Partial linking.
- Incremental builds.

### System Interface

#### System Calls

- `int 0x80`.
- `sysenter`.
- `syscall`.
- x86_64 syscall convention:
  - `rax`: syscall number.
  - arguments: `rdi`, `rsi`, `rdx`, `r10`, `r8`, `r9`.
- Common syscalls:
  - open.
  - read.
  - write.
  - mmap.
  - mprotect.
  - fork.
  - exit.
- Error handling.
- libc wrappers.

#### Unsafe Memory Operations

- Pointer dereferencing.
- Pointer arithmetic.
- Buffer overflows.
- Stack smashing.
- Undefined behavior in C.
- Memory safety in Rust.
- ASLR.
- DEP/NX.
- Bounds checking.

#### Inline Assembly

- AT&T syntax.
- Intel syntax.
- Constraints:
  - `r`.
  - `m`.
  - `g`.
  - `=`.
  - `&`.
  - `+`.
- Clobber lists.
- Memory barriers.
- C/Rust and assembly integration.
- Platform differences.

## Projects

1. Heap allocator.
2. ELF parser.
3. Mini shell.
4. Rust syscall wrapper library.

## Recommended Resources

| Topic | Resource |
|---|---|
| x86_64 ABI | System V AMD64 ABI |
| ELF Format | ELF Specification |
| Syscalls | man pages and Linux syscall table |
| GDB | GDB manual |
| Rust | Rust Book |

## Learning Outcomes

- Understand how programs are built, linked, and loaded.
- Write low-level C/Rust code with direct OS interfaces.
- Read assembly and calling conventions.
- Implement systems programming projects from scratch.
- Debug at syscall and hardware level.
- Reason about memory layout and safety.


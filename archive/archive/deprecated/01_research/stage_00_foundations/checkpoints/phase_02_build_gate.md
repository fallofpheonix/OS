# Phase 2 Build Gate: Low-Level Programming

## Rule

This gate must pass before implementing allocators, shells, syscall layers, ELF loaders, kernel boot paths, or low-level security labs.

## C Implementation

### Memory Allocator

- [ ] Implement `malloc`/`free` using `mmap`.
- [ ] Support multiple block sizes.
- [ ] Implement coalescing.
- [ ] Track fragmentation.
- [ ] Pass randomized alloc/free stress tests.

### ELF Parser

- [ ] Parse ELF headers.
- [ ] Determine architecture.
- [ ] List sections.
- [ ] List segments.
- [ ] Parse symbol table.
- [ ] Parse relocation entries.
- [ ] Verify dynamic dependencies.

### Mini Shell

- [ ] Parse and execute commands with `fork` and `execve`.
- [ ] Handle pipes.
- [ ] Handle `>`.
- [ ] Handle `<`.
- [ ] Handle `>>`.
- [ ] Implement `cd`.
- [ ] Implement `exit`.
- [ ] Support background execution with `&`.
- [ ] Handle `SIGCHLD` for zombie cleanup.

## Rust Implementation

### Syscall Wrapper Library

- [ ] Wrap `open`.
- [ ] Wrap `read`.
- [ ] Wrap `write`.
- [ ] Wrap `mmap`.
- [ ] Wrap `fork`.
- [ ] Wrap `execve`.
- [ ] Convert `errno` to `Result<T, E>`.
- [ ] Use type-safe file handles.
- [ ] Use type-safe memory regions.
- [ ] Avoid unnecessary copies.

## Assembly And Low-Level

- [ ] Read generated assembly.
- [ ] Identify calling convention usage.
- [ ] Trace register usage.
- [ ] Trace stack usage.
- [ ] Identify prologue and epilogue.
- [ ] Use inline assembly in C or Rust.
- [ ] Implement atomic compare-and-swap.
- [ ] Use memory barriers.
- [ ] Use CPU-specific instructions such as CPUID or RDTSC.

## Analysis And Debugging

- [ ] Step through syscalls with GDB or LLDB.
- [ ] Inspect registers.
- [ ] Inspect stack frames.
- [ ] View text, data, heap, and stack layout.
- [ ] Analyze allocator behavior.
- [ ] Use perf/profiling tools.
- [ ] Identify hot paths.
- [ ] Detect cache misses.
- [ ] Measure syscall overhead.
- [ ] Profile allocator performance.

## Security Analysis

- [ ] Identify buffer overflow vulnerabilities.
- [ ] Explain ASLR.
- [ ] Explain DEP/NX.
- [ ] Trace theoretical exploit flow without weaponized payloads.

## Code Quality

- [ ] Error handling for all syscalls.
- [ ] No memory leaks.
- [ ] Bounds checking.
- [ ] Clear separation of concerns.
- [ ] Unit tests.
- [ ] Integration tests.

## Pass Criteria

- Heap allocator passes stress tests.
- ELF parser can inspect real ELF binaries.
- Mini shell executes commands, pipes, redirects, and cleans zombies.
- Rust syscall wrapper passes unit tests.
- Debugging and profiling notes are linked from lab records.


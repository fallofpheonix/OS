# Phase 2 Labs: Low-Level Programming

## Rule

Research notes stay in `01_research/`. Code and test artifacts belong in `14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/`.

## Project 01: Heap Allocator

Duration: 2-3 weeks.

Implementation path:

`14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/heap_allocator/`

Milestones:

- Bump allocator.
- Free list.
- Split blocks.
- Coalesce blocks.
- Fragmentation metrics.
- Randomized stress test.

## Project 02: ELF Parser

Duration: 1-2 weeks.

Implementation path:

`14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/elf_parser/`

Milestones:

- ELF header parser.
- Program header parser.
- Section header parser.
- Symbol table parser.
- Relocation parser.
- Dynamic dependency listing.

## Project 03: Mini Shell

Duration: 2-3 weeks.

Implementation path:

`14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/mini_shell/`

Milestones:

- Command parser.
- `fork` + `execve`.
- Pipes.
- Redirection.
- Builtins.
- Background execution.
- Signal handling.

## Project 04: Rust Syscall Wrappers

Duration: 1 week.

Implementation path:

`14_experiments/poc/stage_00_foundations/phase_02_low_level_programming/rust_syscalls/`

Milestones:

- File syscall wrappers.
- Memory syscall wrappers.
- Process syscall wrappers.
- `Result<T, E>` error model.
- Unit tests.

## Completion Record

| Project | Status | Evidence path | Notes |
|---|---|---|---|
| Heap allocator | TODO | TBD | `mmap`, free list, coalescing |
| ELF parser | TODO | TBD | Headers, sections, symbols, relocations |
| Mini shell | TODO | TBD | Exec, pipes, redirection, signals |
| Rust syscall wrappers | TODO | TBD | Safe syscall abstractions |


# Phase 1 Labs: Computer Architecture

## Rule

Research notes stay in `01_research/`. Benchmark and implementation evidence belongs in `14_experiments/poc/stage_00_foundations/phase_01_architecture/`.

## Lab 01: Assembly And Calling Convention

Output:

- Compile simple C functions to assembly.
- Identify argument registers.
- Identify return registers.
- Identify caller-saved and callee-saved registers.
- Explain prologue and epilogue.

## Lab 02: Cache Locality

Output:

- Row-major vs column-major benchmark.
- Sequential vs strided access benchmark.
- Cache-line-size notes.
- Miss/hit analysis.

## Lab 03: False Sharing

Output:

- Multi-threaded false-sharing demonstration.
- Padded vs unpadded structure comparison.
- Throughput delta.

## Lab 04: Virtual Memory Walkthrough

Output:

- Virtual address bit-field breakdown.
- Multi-level page table trace.
- TLB hit/miss explanation.

## Lab 05: Interrupt And Exception Flow

Output:

- Interrupt flow diagram.
- Page fault flow diagram.
- Trap/syscall transition notes.

## Lab 06: NUMA Reasoning

Output:

- NUMA topology notes.
- Local vs remote memory access implications.
- Placement strategy for memory-heavy workloads.

## Completion Record

| Lab | Status | Evidence path | Notes |
|---|---|---|---|
| Assembly and calling convention | TODO | TBD | Generated assembly analysis |
| Cache locality | TODO | TBD | Access pattern benchmark |
| False sharing | TODO | TBD | Multi-core cache coherency |
| Virtual memory walkthrough | TODO | TBD | Page table and TLB |
| Interrupt and exception flow | TODO | TBD | IDT/ISR/trap path |
| NUMA reasoning | TODO | TBD | Topology and locality |


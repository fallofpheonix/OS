# Phase 1 Build Gate: Computer Architecture

## Rule

This gate must pass before Stage 01 system internals work involving paging, interrupts, syscalls, kernel mode, telemetry, or scheduler design.

## Conceptual Understanding

- [ ] Interrupt flow: draw complete path from hardware trigger to handler execution to return.
- [ ] Page translation: explain multi-level page table lookup.
- [ ] Address translation: trace a virtual address through translation.
- [ ] Memory access path: trace CPU -> registers -> L1/L2/L3 -> RAM.
- [ ] Cache behavior: explain hits, misses, line size, alignment, and locality.
- [ ] Cache coherency: explain how multiple CPUs maintain consistent cache state.
- [ ] NUMA impact: identify when NUMA matters and how locality-aware programming changes design.

## Low-Level Analysis

- [ ] Read assembly and identify register usage.
- [ ] Explain calling convention effects on generated code.
- [ ] Explain instruction pipeline stages.
- [ ] Explain branch prediction and misprediction cost.
- [ ] Analyze cache behavior of simple programs.
- [ ] Identify false sharing.
- [ ] Draw TLB lookup process for a given address.
- [ ] Identify system calls that cause user/kernel transitions.

## Problem-Solving

- [ ] Optimize array access patterns for cache locality.
- [ ] Identify performance bottlenecks in CPU-bound code.
- [ ] Design data structures around 64-byte x86 cache lines.
- [ ] Explain row-major vs column-major performance differences.

## Pass Criteria

- All conceptual diagrams complete.
- At least one cache-locality benchmark recorded.
- At least one assembly analysis note recorded.
- At least one TLB/page-translation walkthrough recorded.


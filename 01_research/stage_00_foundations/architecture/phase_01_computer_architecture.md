# Phase 1: Computer Architecture

## Overview

Understand how modern CPUs execute instructions, manage memory, handle interrupts, enforce privilege boundaries, and expose performance constraints to software.

## Classification

- Stage: `Stage_00_Foundations`
- Type: `FOUNDATIONAL`
- Status: `RESEARCH_ONLY`
- Difficulty: intermediate
- Estimated duration: 4-5 weeks
- Depends on: Phase 0 Computer Science Foundations
- Blocks:
  - Stage 01 System Internals
  - Stage 09 Telemetry
  - Stage 16 Kernel Extensions
  - Stage 18 Custom OS

## Research

### CPU Fundamentals

#### Processor Components

- General-purpose registers.
- Special-purpose registers.
- Control registers.
- ALU.
- Control unit.
- Microcode.
- Clock cycles.
- Frequency.
- Single-core architecture.
- Multi-core architecture.

#### Instruction Cycle

- Fetch.
- Decode.
- Execute.
- Write-back.
- Pipelining.
- Instruction-level parallelism.
- Branch prediction.
- Speculative execution.
- Out-of-order execution.
- Misprediction penalties.

### Memory Hierarchy

#### Caches

- L1 cache.
- L2 cache.
- L3 cache.
- Cache line size.
- Alignment.
- MESI.
- MOESI.
- Replacement policies:
  - LRU.
  - LFU.
- Cache-oblivious algorithms.
- False sharing.

#### Memory Management

- Virtual memory.
- Physical memory.
- Memory hierarchy cost model.
- Temporal locality.
- Spatial locality.
- Working set.
- Cache misses.

### Virtual Memory And Address Translation

#### MMU

- Page tables.
- Hierarchical page tables.
- Page table entries.
- PTE flags.
- Virtual-to-physical address translation.
- TLB basics.

#### TLB

- TLB hits.
- TLB misses.
- TLB invalidation.
- TLB shootdown.
- TLB replacement policies.
- Micro-TLB.
- Unified TLB.

#### DMA

- DMA controller operation.
- Memory access without CPU intervention.
- DMA coherency issues.

### Interrupt And Exception Handling

#### Interrupts

- Maskable interrupts.
- Non-maskable interrupts.
- Interrupt vectors.
- Interrupt descriptor tables.
- Interrupt handlers.
- ISRs.
- Interrupt priority.
- Interrupt nesting.
- Interrupt context.
- State preservation.

#### Exceptions

- Page faults.
- Divide-by-zero.
- Software exceptions.
- Traps.
- Exception handling flow.
- Recovery mechanisms.

### Privilege Modes And Protection

#### Privilege Rings

- Ring 0 kernel mode.
- Ring 3 user mode.
- System calls.
- Transition mechanisms.
- CPU mode switching overhead.
- PTE permission bits.

#### NUMA

- NUMA topology.
- Local memory latency.
- Remote memory latency.
- NUMA-aware algorithms.
- Node affinity.
- Memory locality.

## Learning Outcomes

- Understand instruction execution at the hardware level.
- Know how caches, TLB, and virtual memory affect performance.
- Trace the memory access path from CPU to RAM.
- Understand interrupt handling and privilege transitions.
- Reason about performance at the architecture level.


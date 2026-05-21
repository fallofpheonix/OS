# Phase 3: Operating Systems

## Overview

Understand how kernels manage processes, memory, I/O, filesystems, firmware handoff, and system resources.

## Classification

- Stage: `Stage_01_System_Internals`
- Type: `FOUNDATIONAL`
- Status: `RESEARCH_ONLY`
- Difficulty: advanced
- Estimated duration: 7-9 weeks
- Depends on:
  - Phase 0 Computer Science Foundations
  - Phase 1 Computer Architecture
  - Phase 2 Low-Level Programming

## Research

### Process Management

#### Processes

- Process control block / process descriptor.
- Process creation:
  - `fork`.
  - `exec`.
  - `clone`.
- Process states:
  - running.
  - ready.
  - blocked.
  - zombie.
- Process hierarchy.
- Parent-child relationships.
- Context isolation.
- Address-space isolation.
- Process termination.
- Cleanup.

#### Threads

- Lightweight processes.
- Heavy processes.
- Thread-local storage.
- Mutexes.
- Semaphores.
- Condition variables.
- Pthreads API.
- Kernel threads.
- User-level threads.
- Thread pool patterns.

#### IPC

- Pipes.
- Named pipes.
- Message queues.
- Shared memory segments.
- Unix domain sockets.
- Network sockets.
- Signals.
- Signal handlers.
- Memory barriers.
- Happens-before relations.

### CPU Scheduling

#### Scheduling Fundamentals

- Throughput.
- Latency.
- Fairness.
- Preemptive scheduling.
- Non-preemptive scheduling.
- Context-switch overhead.
- Response time.
- Turnaround time.
- CPU utilization.

#### Scheduling Algorithms

- FCFS.
- SJF.
- SRTF.
- Round robin.
- Time quantum selection.
- Priority scheduling.
- Aging.
- Multilevel feedback queue.
- Linux CFS.
- Load balancing.
- Cache-aware scheduling.

#### Context Switching

- Register save/restore.
- Flags save/restore.
- Program counter save/restore.
- TLB flush cost.
- CPU cache pollution.
- Interrupt handling during switch.
- Voluntary switches.
- Involuntary switches.

### Virtual Memory

#### Paging

- Page tables.
- Multi-level page-table hierarchies.
- Page allocation.
- Page deallocation.
- Demand paging.
- Lazy allocation.
- Minor page faults.
- Major page faults.
- LRU.
- LFU.
- FIFO.
- Clock algorithm.
- Working set model.
- Thrashing.
- Copy-on-write.

#### Memory Management

- Virtual address-space layout.
- `mmap`.
- Swapping.
- NUMA-aware allocation.
- Per-page permissions.
- Hugepages.
- Transparent huge pages.

### Filesystems And I/O

#### VFS

- VFS abstraction layer.
- Inodes.
- Metadata.
- Dentry cache.
- File descriptor table.
- Buffering.
- Page cache.
- Buffer management.
- Writeback.
- `fsync`.
- Durability guarantees.

#### Filesystem Types

- ext2.
- ext3.
- ext4.
- Journaling.
- Crash recovery.
- Btrfs.
- Copy-on-write filesystems.
- NFS.
- Tmpfs.

#### I/O And Drivers

- IRQ.
- ISR.
- DMA transfers.
- Device I/O.
- Memory-mapped I/O.
- Driver architecture.
- Character devices.
- Block devices.

### Boot And Firmware

#### Boot Sequence

- BIOS initialization.
- UEFI initialization.
- Bootloader responsibilities.
- GRUB.
- Kernel loading.
- Kernel decompression.
- Early kernel initialization.
- Init process.
- Userspace startup.

#### UEFI

- UEFI specification.
- BIOS differences.
- Secure Boot.
- Signature verification.
- UEFI variables.
- Firmware interface.

## Study Resources

- xv6.
- OSDev Wiki.
- Linux kernel documentation.

## Learning Outcomes

- Understand process lifecycle and scheduling algorithms.
- Trace context switching at kernel level.
- Understand virtual memory and paging mechanisms.
- Understand filesystem abstraction.
- Read and modify xv6.
- Debug OS behavior using traces.


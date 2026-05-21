# OS Engineering Curriculum

## Purpose

Map OS-engineering topics to concrete kernel projects, toy-OS exercises, and real-world code references.

This curriculum is the systems foundation for the Cyber AI OS research path.

## 1. Hardware-First Layer

### Topics

- Computer architecture.
- CPU pipelines.
- Branch prediction.
- Memory ordering.
- CPU modes.
- Privilege rings.
- `CR0`, `CR3`, MSRs.
- Memory hierarchy.
- Caches.
- Cache-line alignment.
- False sharing.
- NUMA zones.

### Kernel Concepts

- User mode vs kernel mode.
- Ring 3 to Ring 0 transition.
- Page-table root in `CR3`.
- Memory barriers such as `mb()` and `smp_rmb()`.
- Scheduler and allocator locality decisions.

### Exercise

Write a small kernel or kernel module that records:

- Current CPU ID.
- `CR0`.
- `CR3`.
- Selected MSRs where safe.
- Basic NUMA topology where available.

Output:

```text
cpu_id
cr0
cr3
numa_node
cache_line_size
```

### Exit Criteria

- Runs on QEMU or Linux VM.
- Produces per-core output.
- Does not require undefined privileged access from userspace.

## 2. Memory Management Core

### Topics

- Segmentation.
- Paging.
- Multi-level page tables.
- Virtual memory.
- Page faults.
- Demand paging.
- MMU.
- TLB.
- TLB shootdowns.

### Kernel Concepts

- `x86_64` uses paging as the primary isolation mechanism.
- Legacy segmentation is minimal in long mode.
- Page faults are part of normal VM behavior.
- SMP invalidation requires coordination across cores.

### Exercise

Implement minimal page-table setup in `xv6` or a tiny kernel.

Then:

- Map kernel text/data.
- Map a test userspace page.
- Intentionally access unmapped memory.
- Print page-fault address and error code.

### Exit Criteria

- Page table layout is documented.
- Page fault handler prints diagnostics.
- Invalid access does not silently corrupt state.

## 3. Concurrency And Control Flow

### Topics

- Interrupts.
- DMA.
- IPIs.
- Context switching.
- FPU save/restore.
- Scheduler tick.
- Process scheduling.
- Round-robin.
- CFS.
- Priorities.
- IPC.
- Pipes.
- Signals.
- Shared memory.
- Message queues.

### Kernel Concepts

- IRQ handlers must be bounded.
- DMA buffers require explicit mapping and alignment.
- IPIs coordinate per-core state.
- Scheduler policy must preserve invariants.
- IPC must define ownership and lifetime.

### Exercise

Extend `xv6` or SerenityOS with one of:

- New scheduler policy.
- New IPC primitive such as channels.
- Per-process scheduling attribute.

### Exit Criteria

- Scheduler or IPC tests exist.
- Starvation behavior is documented.
- Blocking and wakeup paths are deterministic.

## 4. Startup, Binaries, And Drivers

### Topics

- Boot process.
- UEFI.
- ACPI.
- ELF loading.
- Device drivers.
- HAL.
- Kernel modules.
- Symbol tables.
- Module dependencies.

### Kernel Concepts

- Firmware passes control and boot metadata.
- ACPI describes hardware and power state.
- ELF loader maps executable segments.
- Driver probe failure must not corrupt global state.
- Kernel modules require symbol resolution and lifecycle rules.

### Exercise

Build:

- Minimal hello-world Linux kernel module.
- Character device.
- Userspace program that opens or `mmap`s the device.

### Exit Criteria

- Module loads and unloads cleanly.
- Device node is usable.
- Kernel logs all lifecycle events.
- Userspace test handles errors.

## 5. Kernel Services And Observability

### Topics

- Syscalls.
- `syscall`.
- `sysenter`.
- Syscall dispatch.
- `strace`.
- Filesystems.
- VFS.
- Inodes.
- Dentries.
- `open/read/write/close`.
- Kernel debugging.
- `kgdb`.
- `kdb`.
- `kdump`.
- `printk`.
- `ftrace`.
- `perf`.

### Kernel Concepts

- Syscalls are ABI contracts.
- VFS separates syscall semantics from filesystem implementation.
- Observability must be designed before debugging production failures.

### Exercise

Add a syscall to `xv6` or a local Linux kernel build.

Write:

- Kernel implementation.
- Userspace wrapper.
- Test harness.
- Failure-case tests.

### Exit Criteria

- Syscall number documented.
- ABI documented.
- Tests cover success and error paths.
- Tracing confirms dispatch path.

## 6. Research Repositories

| Repository | Study Areas |
|---|---|
| Linux Kernel | `fs/`, `mm/`, `kernel/`, `drivers/char/`, `arch/x86/` |
| Linux From Scratch | Toolchain, userspace, kernel image integration |
| xv6 RISC-V | Boot, ELF loading, paging, context switching, IPC, simple FS |
| OSDev Wiki | x86_64 boot, GDT, IDT, paging, drivers |
| SerenityOS | Readable kernel/userland, IPC, FS, scheduler, GUI stack |
| Redox OS | Rust OS design, userspace IPC, filesystem and driver model |

## 16-Week Roadmap

| Week | Focus | Output |
|---:|---|---|
| 1 | C, assembly, build tools | freestanding build and linker script |
| 2 | Boot basics | boot sector or Limine/GRUB kernel entry |
| 3 | Serial logging and panic path | serial logger and halt loop |
| 4 | CPU modes and descriptors | GDT/IDT notes and basic handlers |
| 5 | Paging | initial page tables and page fault handler |
| 6 | Physical memory | page allocator and memory map parser |
| 7 | Kernel heap | allocator tests and failure handling |
| 8 | Interrupts and timer | timer tick and interrupt diagnostics |
| 9 | Scheduler | simple task switch or xv6 scheduler modification |
| 10 | IPC | pipe/channel/message primitive |
| 11 | ELF loading | load simple userspace binary |
| 12 | Syscalls | syscall table and userspace wrapper |
| 13 | Filesystem | initramfs or VFS prototype |
| 14 | Drivers | character driver or framebuffer driver |
| 15 | Observability | tracing, debug logs, crash dump notes |
| 16 | Capstone | bootable tiny OS or xv6 extension report |

## Repo Skeleton

```text
os_course/
├── 01_boot/
│   ├── README.md
│   ├── src/
│   └── tests/
├── 02_paging/
│   ├── README.md
│   ├── notes/
│   └── src/
├── 03_sched/
│   ├── README.md
│   ├── xv6_patch/
│   └── tests/
├── 04_ipc/
│   ├── README.md
│   ├── design.md
│   └── tests/
├── 05_fs/
│   ├── README.md
│   ├── vfs_notes.md
│   └── src/
├── 06_drivers/
│   ├── README.md
│   ├── char_driver/
│   └── userspace_test/
├── 07_syscalls/
│   ├── README.md
│   ├── syscall_table.md
│   └── tests/
└── 08_observability/
    ├── README.md
    ├── tracing.md
    └── debug_workflows.md
```

## README Template For Each Module

```md
# Module Name

## Objective

TBD

## Concepts

- TBD

## Code Targets

- TBD

## Tests

- TBD

## Debug Workflow

- QEMU:
- GDB:
- Logs:

## Exit Criteria

- TBD
```

## Integration With Cyber AI OS

| OS Curriculum Output | Later Cyber AI OS Use |
|---|---|
| Boot and kernel entry | Custom boot modes and trusted startup |
| Paging and faults | Memory isolation, sandboxing, exploit detection |
| Scheduler and IPC | Policy-aware workload control |
| Drivers and HAL | Security sensor integration |
| Syscalls | Auditable user/kernel boundary |
| Filesystems | Forensic artifact collection |
| Observability | Kernel telemetry and AI Phoenix Bus |

## Constraint

Do not treat this as theory-only. Each module must produce:

- One runnable artifact.
- One design note.
- One test or debug workflow.
- One failure analysis section.


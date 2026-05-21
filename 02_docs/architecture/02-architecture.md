# Architecture

## System Flow

```text
Hardware
  -> UEFI / BIOS
  -> Bootloader
  -> Kernel
  -> HAL
  -> Drivers
  -> System Services
  -> Userspace
  -> Applications
```

## Core Modules

1. Boot layer.
2. Memory layer.
3. Process manager.
4. Scheduler.
5. IPC.
6. Filesystem.
7. Networking.
8. Security layer.
9. Package layer.

## Invariants

- Kernel owns physical memory map interpretation.
- Drivers cannot access userspace memory without explicit copy or mapping rules.
- Interrupt context must not perform unbounded blocking work.
- Userspace cannot execute privileged CPU instructions.
- Filesystem writes must be recoverable or explicitly non-journaled.

## Base-System Paths

| Path | Data Flow | Use Case |
|---|---|---|
| Scratch | Bootloader -> custom kernel -> custom userspace | OS research |
| Arch | Arch ISO/profile -> custom packages -> branded live/install image | General Linux-derived OS |
| Kali | Kali live-build -> security package set -> controlled live image | Authorized security testing |
| LFS/Buildroot | Source toolchain -> rootfs -> kernel/userspace integration | Intermediate engineering path |


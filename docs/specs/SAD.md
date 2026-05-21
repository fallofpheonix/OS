# System Architecture Specification

## Purpose

Define the system being built and constrain scope before implementation.

## Target Architecture

| Field | Decision |
|---|---|
| Primary ISA | `x86_64` |
| Secondary ISA | Undefined |
| Boot firmware | UEFI first, BIOS optional |
| Test platform | QEMU |
| Hardware support | Virtual hardware first |

ARM64 and RISC-V are out of scope until the `x86_64` boot and runtime path is stable.

## Kernel Model

Scratch path:

| Model | Status | Notes |
|---|---|---|
| Monolithic | Candidate | Simplest early driver and syscall path |
| Microkernel | Deferred | Higher IPC and driver isolation complexity |
| Hybrid | Preferred initial design | Modular subsystems without full microkernel cost |

Initial kernel constraints:

- Single address-space kernel.
- Single-core first.
- Interrupt-driven timer.
- No dynamic driver loading initially.

## Base System

Linux path:

| Base | Tooling | Init | Use Case |
|---|---|---|---|
| Arch Linux | `archiso`, `pacman` | `systemd` | General custom ISO |
| Kali Linux | `live-build`, `apt` | `systemd` | Authorized security image |
| LFS | Source builds | Configurable | Controlled learning path |
| Buildroot | Buildroot configs | BusyBox/init or custom | Embedded/minimal image |

## Memory Management

Scratch path requirements:

- Parse firmware or bootloader memory map.
- Reserve kernel, bootloader, firmware, MMIO, and module regions.
- Implement physical page allocator.
- Implement virtual memory mapping.
- Define kernel high-half or direct-map layout.
- Add kernel heap after paging is stable.

Initial allocator:

- Bitmap or stack allocator for physical pages.
- Bump allocator for early boot.
- Explicit allocation failure paths.

## Virtual Memory Layout

Proposed `x86_64` layout:

| Region | Purpose |
|---|---|
| Low memory | Boot compatibility, early identity mappings |
| Higher-half kernel | Kernel text/data/bss/heap |
| Direct map | Physical memory mapping |
| Kernel stacks | Per-thread stacks with guard pages |
| User space | Process mappings |
| MMIO | Device mappings, non-cacheable where required |

## File System

| Phase | Design |
|---|---|
| Boot | FAT32 EFI system partition |
| Early runtime | Initramfs or read-only ramfs |
| Linux path | `ext4`, `squashfs`, optional `btrfs` |
| Scratch path | VFS before custom disk filesystem |

VFS invariants:

- Path resolution must be bounded.
- Mount table updates must be synchronized.
- Filesystem drivers must reject malformed metadata.
- Read-only root must be supported.

## Undefined Requirements

- Distribution name.
- License policy.
- Secure Boot support.
- GUI target.
- Persistence model.
- Package signing root.


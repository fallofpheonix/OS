# Bridge Approach: Linux From Scratch

## Purpose

Use LFS to learn real OS composition without writing hardware drivers from scratch.

## Position

LFS sits between:

```text
Arch/Kali remastering
  -> LFS / Buildroot
  -> Scratch kernel
```

## What LFS Provides

- Source-level toolchain build.
- Kernel and userspace integration.
- Filesystem hierarchy construction.
- Init and shell setup.
- Direct visibility into libc, binutils, compiler, and core utilities.

## What LFS Does Not Provide

- Custom kernel architecture.
- Custom scheduler.
- Custom memory manager.
- Hardware driver implementation.
- Package ecosystem by default.

## Recommended Use

Use LFS when the goal is:

- Understand user space and kernel space interaction.
- Control every userspace component.
- Build a reproducible minimal Linux.
- Prepare for later custom kernel work.

Do not use LFS when the immediate goal is:

- Fast desktop ISO.
- Kali-like security toolkit distribution.
- Broad package ecosystem without custom package work.

## Migration Path

```text
LFS rootfs
  -> custom init
  -> custom shell/utilities
  -> custom package metadata
  -> custom boot flow
  -> experimental kernel components
  -> scratch kernel path
```

## Project Constraint

If the project takes the bridge approach, the kernel remains Linux until a separate ADR approves scratch-kernel replacement.


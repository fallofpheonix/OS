# ADR-001: Base System

## Status

Proposed.

## Context

The project can be built as:

1. Fully custom OS from scratch.
2. Linux-derived distribution using Arch or Kali.
3. Intermediate system using Linux From Scratch or Buildroot.

## Option 1: Build From Scratch

Decision:

Build bootloader integration, kernel, drivers, userspace, and package model from scratch.

Pros:

- Full control.
- Deep OS learning.
- Custom kernel behavior.
- No inherited distribution policy.

Cons:

- Long development cycle.
- Large complexity.
- Slow hardware support.
- High debugging cost.

## Option 2: Use Arch Or Kali Base

Decision:

Use Arch or Kali as the first bootable base.

Pros:

- Fast prototype.
- Existing drivers.
- Existing package ecosystem.
- Lower bootstrapping cost.

Cons:

- Less control.
- Kernel dependency.
- Distribution limitations.
- Release drift.

## Option 3: Intermediate LFS/Buildroot Path

Decision:

Use Linux From Scratch or Buildroot to control userspace and image generation while retaining Linux kernel support.

Pros:

- Practical control boundary.
- Better reproducibility than remastering full distributions.
- Easier migration toward custom userspace.

Cons:

- Still Linux-kernel dependent.
- More toolchain work than Arch/Kali.
- Less package convenience.

## Recommendation

Start with Arch-based or Buildroot-based image work.

Use Kali only for authorized security-focused builds.

Move to scratch kernel work after boot, packaging, filesystem layout, and userspace requirements are stable.


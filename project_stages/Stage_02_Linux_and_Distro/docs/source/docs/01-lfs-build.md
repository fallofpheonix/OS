# 1. Custom Linux: LFS-Style Build

## Purpose

Build a minimal custom Linux system using a controlled toolchain, Linux kernel, BusyBox, and root filesystem.

## Overall Plan

1. Set up host build environment.
2. Build or install cross toolchain.
3. Compile Linux kernel into `bzImage`.
4. Build libc or select static BusyBox for first image.
5. Install BusyBox to provide `/bin/sh`, `ls`, `ps`, `mount`, and init utilities.
6. Assemble root filesystem.
7. Create initramfs, disk image, or ISO.
8. Boot in QEMU.

## Key Directories

```text
scripts/    Build scripts for kernel, libc, BusyBox, rootfs, image
config/     passwd, inittab, fstab, shell init, kernel config
rootfs/     Target root filesystem
iso/        ISO staging/output
images/     Generated bootable artifacts
```

## Root Filesystem Layout

```text
rootfs/
├── bin/
├── sbin/
├── usr/bin/
├── usr/sbin/
├── lib/
├── dev/
├── proc/
├── sys/
├── etc/
└── init
```

## Example Build Flow

```sh
./scripts/00-prerequisites.sh
./scripts/01-kernel.sh
./scripts/02-glibc.sh
./scripts/03-busybox.sh
./scripts/04-rootfs.sh
./scripts/05-iso.sh
```

## Using Kali Or Arch As Host

### Kali

Use when the final image needs security tooling, fuzzers, reverse-engineering tools, or penetration-test utilities.

Constraints:

- Do not enable offensive services by default.
- Do not embed credentials.
- Document authorized-use boundary.

### Arch

Use when the build host should be minimal, current, and easy to customize.

Strengths:

- Excellent package documentation.
- Simple package format.
- Good base for custom ISO work through `archiso`.

## References

- Linux From Scratch: https://www.linuxfromscratch.org
- Arch Wiki: https://wiki.archlinux.org
- OSDev Wiki: https://wiki.osdev.org/


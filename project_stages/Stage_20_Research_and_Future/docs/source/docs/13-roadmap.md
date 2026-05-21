# Roadmap

## Phase 1

- [ ] Bootable image.
- [ ] UEFI startup.
- [ ] Kernel entry or Linux login shell.

## Phase 2

- [ ] Memory manager.
- [ ] Scheduler.
- [ ] Interrupt handling.

## Phase 3

- [ ] Shell.
- [ ] Filesystem.
- [ ] Drivers.

## Phase 4

- [ ] Networking.
- [ ] Security.
- [ ] Packages.

## Phase 5

- [ ] GUI.
- [ ] Containers.
- [ ] Multi-user support.

## Recommended Paths

### Pure OS Research

```text
Bootloader -> Kernel -> Memory -> Scheduler -> Filesystem -> Shell
```

Use:

- GRUB or Limine.
- C/C++ with optional Rust.
- QEMU.
- UEFI.
- ELF loader.

### Linux-Derived Distro

```text
Arch base -> Custom ISO -> Custom packages -> Own desktop -> Own installer -> Own branding
```

### Intermediate Engineering

```text
Linux From Scratch -> Buildroot -> Custom kernel patches -> Userspace rewrite
```

For long-term OS engineering, the intermediate path is usually more practical before moving to a fully scratch kernel.


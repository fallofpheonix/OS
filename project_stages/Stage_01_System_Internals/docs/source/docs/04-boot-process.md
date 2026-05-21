# Boot Process

## Flow

```text
Power on
  -> UEFI initialization
  -> Bootloader load
  -> Kernel entry
  -> Memory setup
  -> Interrupt table
  -> Scheduler init
  -> Userspace start
```

## Bootloader Options

| Bootloader | Use Case | Trade-Off |
|---|---|---|
| Limine | Scratch kernel | Simple modern protocol |
| GRUB | Scratch or Linux ISO | Mature, more legacy complexity |
| Custom loader | Research only | High complexity |
| systemd-boot | Linux UEFI images | Simple, Linux-focused |

## Scratch Boot Requirements

- Valid executable kernel image.
- Linker script.
- Boot protocol metadata.
- Stack setup.
- Serial output.
- Panic halt path.

## Linux-Derived Boot Requirements

- Kernel.
- Initramfs.
- Root filesystem.
- Bootloader entry.
- Kernel command line.
- UEFI boot files.

## First Boot Gate

Pass only when:

- VM reaches kernel serial output or login shell.
- Boot log is captured.
- Shutdown is clean.
- Failure path emits diagnostic output.


# From-Scratch OS Development

## Purpose

Provide a staged quick-start for building a minimal operating system from zero.

## Prerequisites

Required knowledge:

- C.
- Assembly.
- Linkers and object files.
- CPU privilege levels.
- Interrupts.
- Paging.
- Basic filesystems.

Required tools:

- `nasm`.
- `x86_64-elf-gcc` or `i686-elf-gcc`.
- `ld` or `lld`.
- `qemu-system-x86_64`.
- `gdb`.
- GRUB or Limine.

## Stage 1: Bootloader

Objective:

- Produce visible output from the earliest boot stage.

BIOS prototype:

- 512-byte boot sector.
- Boot signature `0xAA55`.
- Print diagnostic string.
- Halt predictably.

Modern path:

- Use Limine or GRUB.
- Load ELF kernel.
- Pass memory map and framebuffer details.

Exit criteria:

- QEMU starts image.
- Boot path prints a marker.
- Failure path halts in known state.

## Stage 2: Kernel

Objective:

- Enter kernel C/C++ code and initialize core CPU state.

Initial tasks:

- Link kernel at known address.
- Set stack.
- Initialize serial output.
- Validate boot protocol.
- Set GDT.
- Set IDT.
- Install page fault handler.
- Parse memory map.

Memory tasks:

- Early bump allocator.
- Physical page allocator.
- Virtual memory mapping.
- Kernel heap.

Process tasks:

- Kernel task.
- Timer interrupt.
- Context switch.
- Scheduler.

Exit criteria:

- Kernel prints serial log.
- Page faults dump fault address.
- Timer interrupt fires.
- Allocator passes basic tests.

## Stage 3: Filesystem And Drivers

Driver order:

1. Serial.
2. Timer.
3. Interrupt controller.
4. Keyboard.
5. Framebuffer.
6. Block device.
7. Network.

Filesystem order:

1. Initramfs.
2. VFS.
3. Read-only filesystem.
4. Writable filesystem.

Exit criteria:

- Kernel loads a file from initramfs.
- Keyboard input works.
- Block read path works.

## Stage 4: Userland

Initial userspace:

- Init process.
- Shell.
- Basic utilities.
- Syscall ABI.
- ELF loader.

Core commands:

- `ls`.
- `cat`.
- `ps`.
- `kill`.
- `mount`.
- `reboot`.

Exit criteria:

- Kernel enters Ring 3.
- Init launches shell.
- Shell executes at least one user command.

## Debugging

QEMU:

```sh
qemu-system-x86_64 -kernel build/kernel.elf -serial stdio -s -S
```

GDB:

```sh
gdb build/kernel.elf
target remote :1234
```

Required logs:

- Serial boot log.
- Page fault dump.
- Interrupt markers.
- Memory allocator diagnostics.

## Common Pitfalls

- Accidentally linking host libc.
- Using host headers in kernel code.
- Incorrect linker script alignment.
- Invalid GDT/IDT descriptors.
- Enabling interrupts before handlers are valid.
- Corrupting memory map reserved regions.
- Assuming QEMU behavior matches hardware.

## References

- OSDev Wiki: https://wiki.osdev.org/
- cfenollosa/os-tutorial: https://github.com/cfenollosa/os-tutorial
- Intel Software Developer Manuals.
- AMD64 Architecture Programmer's Manual.


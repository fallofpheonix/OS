# 1. OS From Scratch: Outline

## Tools

- NASM for `src/boot/boot.asm`.
- GCC cross-compiler: `i686-elf-gcc` or `x86_64-elf-gcc`.
- LD linker.
- QEMU for testing.
- Optional GRUB or Limine for later boot stages.

## Bootloader

File:

```text
src/boot/boot.asm
```

Requirements:

- Exactly 512 bytes for legacy BIOS boot-sector prototype.
- Ends with boot signature `0xAA55`.
- Prints a diagnostic message.
- Loads kernel image from disk.
- Jumps to kernel entry.

UEFI work should not use a 512-byte boot sector. Use GRUB, Limine, or a UEFI application instead.

## Kernel

File:

```text
src/kernel/kernel.c
```

Initial requirements:

- Freestanding C.
- No libc.
- No `printf`.
- Writes to VGA text buffer or serial.
- Halts cleanly after output.

Later requirements:

- Protected mode or long mode setup.
- GDT.
- IDT.
- Paging.
- Physical memory allocator.
- Timer interrupt.

## Linker Script

File:

```text
src/linker.ld
```

Responsibilities:

- Place `.text`, `.rodata`, `.data`, and `.bss`.
- Define kernel load address.
- Export kernel entry symbol.
- Keep alignment explicit.

## Build Steps

BIOS flat-image prototype:

```sh
nasm src/boot/boot.asm -f bin -o build/boot.bin
gcc -ffreestanding -m32 -c src/kernel/kernel.c -o build/kernel.o
ld -T src/linker.ld -o build/kernel.bin build/kernel.o
cat build/boot.bin build/kernel.bin > images/os.img
qemu-system-i386 -fda images/os.img
```

GRUB ISO prototype:

```sh
make iso
make run
```

## References

- OSDev Wiki: https://wiki.osdev.org/
- cfenollosa/os-tutorial: https://github.com/cfenollosa/os-tutorial


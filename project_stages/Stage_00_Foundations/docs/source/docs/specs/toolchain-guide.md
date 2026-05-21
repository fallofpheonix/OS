# Build Environment And Toolchain Guide

## Rule

Do not build the scratch kernel with the host system compiler as a normal hosted program.

Reason:

- Host compiler defaults can link host libc.
- Host headers may leak into kernel code.
- ABI assumptions may not match the target runtime.
- Startup files such as `crt0` are invalid for freestanding kernels.

## Host Prerequisites

Debian/Kali:

```sh
sudo apt update
sudo apt install -y build-essential bison flex libgmp3-dev libmpc-dev libmpfr-dev texinfo nasm xorriso grub-pc-bin grub-efi-amd64-bin qemu-system-x86
```

Arch:

```sh
sudo pacman -S --needed base-devel bison flex gmp libmpc mpfr texinfo nasm xorriso grub qemu-full
```

## Cross Compiler

Target triplets:

| Target | Use |
|---|---|
| `i686-elf` | 32-bit protected-mode prototype |
| `x86_64-elf` | 64-bit kernel target |

Required versions must be pinned before release:

| Component | Version |
|---|---|
| Binutils | TBD |
| GCC | TBD |
| GDB | TBD |
| NASM | TBD |
| QEMU | TBD |

## Compiler Flags

Kernel C/C++:

```sh
-ffreestanding
-fno-stack-protector
-fno-pic
-nostdlib
-nostdinc
-Wall
-Wextra
```

Linking:

```sh
-T src/linker.ld
-nostdlib
```

## Build Automation

Required targets:

```sh
make clean
make
make run
```

Expected future targets:

```sh
make toolchain
make kernel
make iso
make qemu
make test
```

## Emulation Environment

BIOS prototype:

```sh
qemu-system-i386 -fda images/os.img
```

UEFI ISO:

```sh
qemu-system-x86_64 \
  -m 4096 \
  -machine q35 \
  -cdrom images/myos.iso \
  -serial stdio \
  -boot d
```

## Failure Considerations

- Treat host header inclusion as a build failure.
- Treat implicit libc linkage as a build failure.
- Record every tool version in release notes.
- Keep generated toolchain under `tools/` or an external pinned prefix.
- Do not require undocumented environment variables.


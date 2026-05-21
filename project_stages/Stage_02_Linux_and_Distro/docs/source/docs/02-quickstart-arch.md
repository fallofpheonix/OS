# 2. Quickstart: Arch Host

## Use Case

Use Arch as the development host or as the base for a custom Linux ISO.

## Host Packages

```sh
sudo pacman -Syu
sudo pacman -S --needed base-devel git qemu-full nasm xorriso grub mtools dosfstools archiso
```

## Custom ISO Path

```sh
cp -r /usr/share/archiso/configs/releng profile
vim profile/packages.x86_64
sudo mkarchiso -v -w build/arch -o images profile
```

## Cross-Build Notes

- Prefer cross compiler for scratch kernel work.
- Keep host headers out of freestanding kernel builds.
- Use `-ffreestanding`.
- Avoid linking libc into early kernel code.

## Chroot Notes

Use `arch-chroot` only for Linux-derived image construction.

Do not use host chroot assumptions for scratch kernel builds.

## QEMU Test

```sh
qemu-system-x86_64 \
  -m 4096 \
  -cdrom images/*.iso \
  -boot d
```

## Risks

- Rolling release drift can break reproducibility.
- Unpinned mirrors produce changing artifacts.
- Pacman keyring issues can block builds.

## Controls

- Save package manifest.
- Record mirror date/source.
- Keep ISO overlay minimal.


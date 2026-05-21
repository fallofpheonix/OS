# Linux-Based OS Customization

## Purpose

Guide customization of an existing Linux base into a tailored OS image using Arch, Kali, Debian, LFS, or Buildroot.

## Base Selection

| Base | Use Case |
|---|---|
| Arch | Minimal rolling base, strong documentation, custom ISO |
| Kali | Security tooling and lab environments |
| Debian | Stable general-purpose base |
| LFS | Learning and source-level control |
| Buildroot | Minimal embedded-style image |

## Arch Linux Setup

Host packages:

```sh
sudo pacman -Syu
sudo pacman -S --needed base-devel git qemu-full archiso xorriso grub mtools dosfstools
```

Custom ISO:

```sh
cp -r /usr/share/archiso/configs/releng profile
vim profile/packages.x86_64
sudo mkarchiso -v -w build/arch -o images profile
```

Customization points:

- `packages.x86_64`.
- `pacman.conf`.
- `airootfs/`.
- `grub/`.
- `profiledef.sh`.

## Kali Linux Customization

Host packages:

```sh
sudo apt update
sudo apt install -y git live-build simple-cdd cdebootstrap curl qemu-system-x86 xorriso
```

Build:

```sh
git clone https://gitlab.com/kalilinux/build-scripts/live-build-config.git
cd live-build-config
./build.sh --variant light
```

Customization points:

- `kali-config/variant-*/package-lists/`.
- `kali-config/common/includes.chroot/`.
- `kali-config/common/hooks/`.

Constraints:

- Offensive tools disabled unless explicitly enabled.
- No embedded credentials.
- No default target lists.
- Authorized-use boundary documented.

## Kernel Compilation

Generic flow:

```sh
make menuconfig
make -j"$(nproc)"
make modules_install
make install
```

Track:

- Kernel version.
- Config file.
- Patches.
- Compiler version.
- Module list.

## Custom ISO Creation

Required contents:

- Kernel.
- Initramfs.
- Root filesystem.
- Bootloader.
- Package manifest.
- OS identity files.
- Service policy.

Validation:

```sh
qemu-system-x86_64 -m 4096 -cdrom images/custom.iso -boot d
```

## Deployment Targets

USB:

- Raw ISO write.
- Verify checksum.
- Test on non-production hardware.

Cloud:

- QCOW2.
- AMI or cloud image conversion.
- Cloud-init policy.

Containers:

- Not a full OS replacement.
- Useful for packaging userspace tools.
- Avoid privileged containers unless required.

## Release Gate

- Clean rebuild.
- Boot test.
- Package manifest.
- Checksum.
- No secrets.
- No unsafe default services.
- Known limitations documented.


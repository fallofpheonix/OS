# 2. Quickstart: Kali Host

## Use Case

Use Kali when the OS image is security-oriented and intended for authorized testing.

Do not use Kali as the default base for a general-purpose OS.

## Host Packages

```sh
sudo apt update
sudo apt install -y git live-build simple-cdd cdebootstrap curl qemu-system-x86 xorriso
```

## Live Image Path

```sh
git clone https://gitlab.com/kalilinux/build-scripts/live-build-config.git
cd live-build-config
./build.sh --variant light
```

## Customization Points

```text
kali-config/variant-*/package-lists/
kali-config/common/includes.chroot/
kali-config/common/hooks/
```

## Package Policy

Start minimal:

```text
kali-linux-core
kali-tools-top10
openssh-client
network-manager
vim
tmux
```

Avoid full tool collections until image size, legal scope, and operational controls are defined.

## Security Requirements

- Inbound services disabled by default.
- No private keys.
- No tokens.
- No VPN profiles.
- No target data.
- Authorized-use banner if distributed.

## QEMU Test

```sh
qemu-system-x86_64 \
  -m 4096 \
  -cdrom images/kali-linux-*-live-amd64.iso \
  -boot d
```

## Risks

- Package bloat.
- Tools with aggressive network behavior.
- Misuse if distributed without policy.
- Repository drift.


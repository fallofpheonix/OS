# Testing

## Test Targets

| Layer | Tests |
|---|---|
| Build | Clean build, artifact existence |
| Boot | QEMU UEFI boot |
| Kernel | Serial output, panic path, memory tests |
| Userspace | Shell commands, process lifecycle |
| Filesystem | Mount, read, write, corruption checks |
| Network | DHCP, route, DNS, ping |
| Security | Credential scan, service scan |

## QEMU UEFI Test

```sh
qemu-system-x86_64 \
  -m 4096 \
  -machine q35 \
  -cdrom images/phoenixos.iso \
  -serial stdio \
  -boot d
```

## Linux-Derived Validation

```sh
uname -a
cat /etc/os-release
ip addr
ip route
systemctl --failed
journalctl -b --no-pager | tail -100
df -h
```

## Release Gate

- Clean rebuild passes.
- QEMU boot passes.
- Manifest generated.
- Checksum generated.
- No unsafe credentials.
- Known limitations documented.


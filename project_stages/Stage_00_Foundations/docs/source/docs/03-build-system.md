# Build System

## Toolchain

| Component | Options |
|---|---|
| C/C++ compiler | `clang`, `gcc` |
| Assembler | `nasm`, `gas` |
| Build executor | `make`, `ninja` |
| Build generator | `cmake` |
| Image creation | `grub-mkrescue`, `xorriso`, `mkarchiso`, `live-build` |
| Virtualization | `qemu-system-x86_64`, VirtualBox |
| Debugging | `gdb`, `lldb` |
| CI | GitHub Actions |

## Repository Layout

```text
my-os/
├── docs/
├── kernel/
├── userspace/
├── boot/
├── drivers/
├── scripts/
└── tools/
```

## Required Build Targets

```sh
make build
make test
make clean
```

## Generated Outputs

- `images/*.iso`
- `images/*.img`
- `images/*.qcow2`
- `build/logs/*.log`
- `manifests/*.txt`

## Reproducibility Rules

- Pin package lists.
- Record tool versions.
- Keep generated images out of git.
- Build from clean checkout before release.
- Keep all image mutations in scripts or tracked config.


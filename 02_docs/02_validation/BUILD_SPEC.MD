# PhoenixOS: Build Specification (Team A)

## 1. Reproducibility Mandate
All PhoenixOS binaries MUST be bit-for-bit identical when built from the same commit.

## 2. Naming Convention
All built binaries in `build/bin/` must be prefixed with `phoenix_`.

## 3. Targets
- **OS:** linux, darwin
- **Arch:** amd64, arm64

## 4. Build Commands
- `go build -o build/bin/phoenix_os ./phoenix_os`
- `05_tools/build_phoenix.sh` (Master Runner)

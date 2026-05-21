# Packaging

## Linux-Derived Path

Arch:

- Use `pacman`.
- Build custom packages with `PKGBUILD`.
- Build ISO with `archiso`.

Kali/Debian:

- Use `apt`.
- Build packages as `.deb`.
- Build live image with `live-build`.

## Scratch Path

Initial package format:

```text
name
version
architecture
files
checksum
dependencies
install script optional
```

## Package Manager Goals

- Install.
- Remove.
- List.
- Verify checksums.
- Resolve dependencies later.

## Repository Policy

- Signed packages.
- Versioned metadata.
- Reproducible package builds where feasible.
- No mutable release artifacts.

## Failure Considerations

- Interrupted install must be recoverable.
- Package database writes must be atomic.
- File ownership conflicts must fail explicitly.
- Signature verification must be mandatory for release builds.


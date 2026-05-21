# Security Model

## Access Control

Initial model:

- UID/GID for Linux-derived path.
- Capability model for scratch path.
- RBAC later if policy complexity requires it.

## Authentication

- Password authentication.
- Token-based service authentication later.
- No biometric layer until core security model is stable.

## Kernel Protection

- NX.
- ASLR.
- SMEP.
- SMAP.
- Read-only kernel text.
- Guard pages.

## Sandboxing

Linux-derived path:

- Namespaces.
- cgroups.
- seccomp.
- Capabilities.

Scratch path:

- Address-space isolation.
- Kernel handles.
- Explicit IPC permissions.

## AI/ML Security Support

AI/ML is a detection and response support layer, not a replacement for kernel isolation, least privilege, package signing, or network policy.

Reference:

- [specs/ai-ml-security-layer.md](specs/ai-ml-security-layer.md)

## Unsafe Defaults

Do not ship:

- Passwordless root login.
- Known SSH password.
- Embedded private keys.
- API tokens.
- VPN profiles.
- Browser sessions.

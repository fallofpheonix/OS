---
Status: Partial
Implementation: 45%
Confidence: Conceptual
---
# L1 OS Core — Capability Spec

L1 handles low-level process lifecycle auditing, filesystem persistence, and network socket configurations.

## Monitored Events
- Process execution (`execve` tracing).
- File write access on critical paths (`/etc/`, `/sys/`, `/boot/`).
- Network bindings and egress connections.

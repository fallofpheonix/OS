# Goals And Scope

## Goals

- Learn OS internals.
- Build custom boot chain.
- Implement kernel initialization.
- Implement scheduler.
- Implement memory manager.
- Add userspace tools.
- Support package management.
- Enable container runtime later.

## Non-Goals

- Immediate production use.
- Full POSIX compatibility.
- Windows compatibility.
- Multi-architecture support initially.
- Broad hardware support in the first build.
- Secure Boot signing before stable boot.

## Initial Milestones

| Milestone | Target |
|---|---|
| M1 | Boot splash or serial output |
| M2 | Kernel initialization |
| M3 | Memory subsystem |
| M4 | Filesystem |
| M5 | Shell |
| M6 | Networking |
| M7 | Package manager |

## Acceptance Criteria

- Reproducible build command.
- Bootable image.
- Login shell or kernel serial console.
- Captured boot log.
- Package or component manifest.
- No unsafe default credentials.


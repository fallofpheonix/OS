# Phase 3 Build Gate: Operating Systems

## Conceptual Understanding

- [ ] Process lifecycle: draw state diagram and explain transitions.
- [ ] Context switching: trace full switch flow and identify saved state.
- [ ] Scheduling: compare FCFS, RR, and MLFQ.
- [ ] Scheduling trade-offs: fairness vs throughput.
- [ ] Virtual memory: explain paging, page faults, and replacement algorithms.
- [ ] VFS abstraction: draw inode, dentry, and file descriptor relationships.

## xv6 Deep Dive

- [ ] Read process structure and context-switching code.
- [ ] Read scheduler implementation.
- [ ] Read page-table setup and memory layout.
- [ ] Read trap and exception handling.
- [ ] Read filesystem implementation.
- [ ] Add a new syscall.
- [ ] Implement priority scheduling.
- [ ] Add simple file operations.

## System Tracing

- [ ] Identify BIOS/UEFI initialization stages.
- [ ] Trace bootloader responsibilities.
- [ ] Explain early kernel initialization.
- [ ] Explain kernel command-line parameters.

## Implementation Projects

- [ ] Process manager in C:
  - [ ] `fork`.
  - [ ] `exec`.
  - [ ] `wait`.
  - [ ] Process states.
  - [ ] Zombie handling.
- [ ] Simple scheduler:
  - [ ] Round robin.
  - [ ] Priority scheduling with aging.
  - [ ] Context-switch overhead measurement.
- [ ] Page table walker:
  - [ ] Multi-level lookup.
  - [ ] Permission and attribute parsing.
  - [ ] Memory error detection.
- [ ] Mini filesystem:
  - [ ] Inode metadata.
  - [ ] Directory entries.
  - [ ] Open/read/write/delete.
  - [ ] Persistence to disk.

## Analysis And Debugging

- [ ] Trace Linux boot with `dmesg`, `journalctl`, and `systemd-analyze`.
- [ ] Inspect processes with `ps`, `pstree`, `top`, `htop`, `/proc/<pid>`, and `strace`.
- [ ] Inspect memory and scheduling with `/proc/sched_debug`, `perf sched`, `vmstat`, `free`, and `smaps`.

## Pass Criteria

- xv6 changes build and run.
- OS diagrams are complete.
- All implementation projects have tests or recorded execution evidence.
- Debugging notes link to command output or captured traces.


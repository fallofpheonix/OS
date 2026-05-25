# Phase 4: Linux Internals

## Overview

Deep dive into Linux kernel architecture, modules, pseudo-filesystems, namespaces, cgroups, capabilities, MAC systems, and tracing.

## Classification

- Stage: `Stage_02_Linux_and_Distros`
- Type: `IMPLEMENTATION`
- Status: `RESEARCH_ONLY`
- Difficulty: expert
- Estimated duration: 8-10 weeks
- Depends on: Phase 3 Operating Systems

## Research

### Kernel Architecture

- Monolithic kernel design.
- Microkernel comparison.
- Kernel subsystems:
  - `mm`.
  - `fs`.
  - `net`.
  - `drivers`.
  - `arch`.
- Kernel entry points.
- Privilege boundaries.
- System call dispatcher.
- Interrupt handlers.
- Exception handlers.

### Kernel Build System

- Kconfig.
- Kbuild.
- Kernel modules.
- Loadable module architecture.
- Kernel versioning.
- Kernel API stability.
- Compilation flags.
- Optimization levels.

### Kernel Modules And Extensibility

- Module entry.
- Module exit.
- Module parameters.
- Symbol export with `EXPORT_SYMBOL`.
- Module dependencies.
- Loading order.
- `/sys/module`.
- Module signing.
- Trusted keys.

### Drivers

- Character drivers.
- Block drivers.
- Network drivers.
- Device model.
- sysfs hierarchy.
- DMA.
- Device state management.

### Pseudo-Filesystems

#### procfs

- `/proc/<pid>`.
- `/proc/stat`.
- `/proc/meminfo`.
- `/proc/sys`.
- `cmdline`.
- `maps`.
- `status`.
- `fd`.
- Limitations and security considerations.

#### sysfs

- `/sys/devices`.
- `/sys/bus`.
- `/sys/class`.
- Kernel object attributes.
- Udev integration.

#### debugfs And tracefs

- Kernel debugging interface.
- Trace event subsystem.
- ftrace configuration.

### Namespaces And Isolation

- PID namespace.
- Network namespace.
- Mount namespace.
- UTS namespace.
- IPC namespace.
- User namespace.
- Cgroup namespace.
- `unshare`.
- `clone` with `CLONE_NEW*`.
- Namespace nesting.
- Namespace file descriptors.

### Control Groups

- cgroup v1 subsystems.
- cgroup v1 hierarchies.
- Task assignment.
- Process migration.
- Resource limits.
- Accounting.
- cgroup v2 unified hierarchy.
- Delegation.
- CPU, memory, IO, RDMA resources.
- Monitoring.
- Cost model.

### Capabilities And Security

- Linux capabilities.
- Effective set.
- Permitted set.
- Inheritable set.
- Bounding set.
- File capabilities.
- Ambient capabilities.
- `CAP_NET_ADMIN`.
- `CAP_DAC_OVERRIDE`.
- `CAP_KILL`.
- SELinux.
- AppArmor.
- Type enforcement.
- Profile-based access control.

### Observability And Tracing

- `ptrace`.
- `strace`.
- `ltrace`.
- ftrace.
- Trace events.
- Event filtering.
- Triggers.
- `trace_printk`.
- `/sys/kernel/tracing`.
- eBPF VM.
- eBPF maps.
- Kprobes.
- Uprobes.
- Tracepoints.
- JIT compilation.
- Networking, observability, and security use cases.
- Audit rules.
- Syscall auditing.
- File access auditing.
- User action tracking.

## Study Resources

- Linux From Scratch.
- eBPF documentation.
- Linux kernel source.
- Kernel documentation.

## Learning Outcomes

- Build and configure custom Linux kernels.
- Write and load kernel modules.
- Use ftrace, eBPF, and strace.
- Understand namespaces and container isolation.
- Manage resource limits with cgroups.
- Debug production system behavior.


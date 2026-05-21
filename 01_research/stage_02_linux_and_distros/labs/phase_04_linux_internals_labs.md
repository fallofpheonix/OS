# Phase 4 Labs: Linux Internals

## Rule

Research stays in `01_research/`. Build artifacts and code evidence belong in `14_experiments/poc/stage_02_linux_and_distros/phase_04_linux_internals/`.

## Lab 01: Custom Kernel Build

- Configure kernel.
- Build kernel.
- Boot kernel.
- Record config delta.
- Record boot logs.

## Lab 02: Kernel Module

- Build character driver.
- Add parameters.
- Export symbol.
- Load/unload.
- Validate cleanup path.

## Lab 03: procfs, sysfs, debugfs, tracefs

- Inspect process and system state.
- Map procfs entries to kernel concepts.
- Inspect sysfs device hierarchy.
- Enable tracefs/ftrace.

## Lab 04: Namespaces

- Create PID, network, mount, UTS, IPC, user, and cgroup namespace examples.
- Record namespace references under `/proc/<pid>/ns`.

## Lab 05: cgroups

- Create limits.
- Move processes.
- Monitor accounting.
- Compare v1 and v2 behavior.

## Lab 06: eBPF And ftrace

- Count syscalls.
- Attach to kprobe or tracepoint.
- Read map data.
- Compare overhead with strace/ftrace.

## Lab 07: Container Internals

- Run container process.
- Inspect namespace mapping.
- Inspect cgroup limits.
- Trace container syscalls.

## Completion Record

| Lab | Status | Evidence path | Notes |
|---|---|---|---|
| Custom kernel build | TODO | TBD | config, boot, logs |
| Kernel module | TODO | TBD | char driver |
| Pseudo-filesystems | TODO | TBD | procfs/sysfs/debugfs/tracefs |
| Namespaces | TODO | TBD | PID/net/mount/UTS/IPC/user/cgroup |
| cgroups | TODO | TBD | v1/v2 limits |
| eBPF and ftrace | TODO | TBD | probes and traces |
| Container internals | TODO | TBD | namespaces/cgroups/syscalls |


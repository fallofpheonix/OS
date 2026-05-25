# Phase 4 Build Gate: Linux Internals

## Kernel Configuration And Compilation

- [ ] Configure kernel with `menuconfig` or `defconfig`.
- [ ] Build custom kernel.
- [ ] Install custom kernel.
- [ ] Boot custom kernel.
- [ ] Disable and enable features, then observe impact.
- [ ] Explain Kconfig syntax.
- [ ] Explain common kernel command-line parameters.
- [ ] Modify boot parameters and observe effects.
- [ ] Debug with `earlycon` and `loglevel`.

## Module Development

- [ ] Write simple character driver module.
- [ ] Implement init function.
- [ ] Implement exit function.
- [ ] Export symbols with `EXPORT_SYMBOL`.
- [ ] Load and unload with `modprobe`.
- [ ] Inspect `/sys/module`.
- [ ] Handle module parameters.
- [ ] Use `pr_debug` and `pr_info`.
- [ ] Read logs with `dmesg`.
- [ ] Clean up correctly on failure.

## Observability And Tracing

- [ ] Trace syscalls with `strace`.
- [ ] Analyze syscall patterns.
- [ ] Filter and follow syscalls.
- [ ] Write simple eBPF syscall-counting probe.
- [ ] Use BCC or libbpf.
- [ ] Read eBPF map data.
- [ ] Attach to kprobes, uprobes, or tracepoints.
- [ ] Measure system behavior non-intrusively.
- [ ] Enable ftrace function tracing.
- [ ] Filter ftrace functions.
- [ ] Set ftrace triggers and actions.
- [ ] Analyze trace output.

## Namespace And Isolation

- [ ] Create network namespace with `ip netns`.
- [ ] Execute shell inside network namespace.
- [ ] Create PID namespace with `unshare`.
- [ ] Create UTS namespace and change hostname.
- [ ] Inspect `/proc/<pid>/ns`.

## Container Internals

- [ ] List cgroup subsystems.
- [ ] Create cgroup hierarchy.
- [ ] Limit memory for a process group.
- [ ] Limit CPU for a process group.
- [ ] Monitor cgroup stats.
- [ ] Explain cgroup v1 vs v2.
- [ ] Run Docker or Podman process.
- [ ] Inspect namespaces with `lsns`.
- [ ] Map container PID to host PID.
- [ ] Trace container syscalls.
- [ ] Observe cgroup limits.

## Analysis And Debugging

- [ ] Read `/proc/<pid>/maps`.
- [ ] Inspect `/proc/stat`.
- [ ] Inspect `/proc/meminfo`.
- [ ] Analyze `/proc/sched_debug`.
- [ ] Use `iostat`.
- [ ] Use `mpstat`.
- [ ] Build Linux From Scratch system.
- [ ] Trace boot sequence with custom kernel.
- [ ] Debug GRUB, initramfs, or rootfs issues.
- [ ] List process capabilities with `getpcaps`.
- [ ] Examine SELinux contexts with `ls -Z`.
- [ ] Explain capability requirements for privileged operations.

## Pass Criteria

- Custom kernel boots.
- Kernel module loads and unloads cleanly.
- At least one ftrace and one eBPF trace are recorded.
- Namespace/cgroup/container isolation notes are complete.
- LFS or equivalent boot analysis is documented.


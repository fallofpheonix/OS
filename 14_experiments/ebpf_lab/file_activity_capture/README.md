# Experiment: eBPF File Activity Capture

## Objective
Prototyping an eBPF-based monitor to capture high-frequency filesystem events (renames, writes) for ransomware detection.

## Goals
- Trace `vfs_write` and `vfs_rename` kernel functions.
- Calculate rename velocity per PID.
- Estimate entropy/randomness of writes (if possible via eBPF or userspace helper).
- Minimize overhead on the VFS layer.

## Implementation Details
- Toolchain: BCC (BPF Compiler Collection) or libbpf.
- Language: C (eBPF) and Python/Go (Loader).

## Expected Output
A stream of JSON events matching the `filesystem_event` schema defined in RFC-001A.
